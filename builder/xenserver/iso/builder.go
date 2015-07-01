package iso

import (
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/helper/communicator"
	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
	xscommon "github.com/rdobson/packer-builder-xenserver/builder/xenserver/common"
	xsclient "github.com/simonfuhrer/go-xenserver-client"
)

type Builder struct {
	config Config
	runner multistep.Runner
}

type Config struct {
	common.PackerConfig   `mapstructure:",squash"`
	xscommon.CommonConfig `mapstructure:",squash"`

	VMMemory      uint   `mapstructure:"vm_memory"`
	DiskSize      uint   `mapstructure:"disk_size"`
	CloneTemplate string `mapstructure:"clone_template"`

	ISOChecksum     string   `mapstructure:"iso_checksum"`
	ISOChecksumType string   `mapstructure:"iso_checksum_type"`
	ISOUrls         []string `mapstructure:"iso_urls"`
	ISOUrl          string   `mapstructure:"iso_url"`
	ISOName         string   `mapstructure:"iso_name"`

	PlatformArgs map[string]string `mapstructure:"platform_args"`

	RawInstallTimeout string        `mapstructure:"install_timeout"`
	InstallTimeout    time.Duration ``

	ctx interpolate.Context
}

func (b *Builder) Prepare(raws ...interface{}) (params []string, retErr error) {

	var errs *packer.MultiError

	err := config.Decode(&b.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &b.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"boot_command",
			},
		},
	}, raws...)

	if err != nil {
		packer.MultiErrorAppend(errs, err)
	}

	errs = packer.MultiErrorAppend(
		errs, b.config.CommonConfig.Prepare(&b.config.ctx, &b.config.PackerConfig)...)
	errs = packer.MultiErrorAppend(errs, b.config.SSHConfig.Prepare(&b.config.ctx)...)

	// Set default values

	if b.config.RawInstallTimeout == "" {
		b.config.RawInstallTimeout = "200m"
	}

	if b.config.DiskSize == 0 {
		b.config.DiskSize = 40000
	}

	if b.config.VMMemory == 0 {
		b.config.VMMemory = 1024
	}

	if b.config.CloneTemplate == "" {
		b.config.CloneTemplate = "Other install media"
	}

	if len(b.config.PlatformArgs) == 0 {
		pargs := make(map[string]string)
		pargs["viridian"] = "false"
		pargs["nx"] = "true"
		pargs["pae"] = "true"
		pargs["apic"] = "true"
		pargs["timeoffset"] = "0"
		pargs["acpi"] = "1"
		b.config.PlatformArgs = pargs
	}

	// Template substitution

	templates := map[string]*string{
		"clone_template":    &b.config.CloneTemplate,
		"network_name":      &b.config.NetworkName,
		"iso_checksum":      &b.config.ISOChecksum,
		"iso_checksum_type": &b.config.ISOChecksumType,
		"iso_url":           &b.config.ISOUrl,
		"iso_name":          &b.config.ISOName,
		"install_timeout":   &b.config.RawInstallTimeout,
	}
	for i := range b.config.ISOUrls {
		templates[fmt.Sprintf("iso_urls[%d]", i)] = &b.config.ISOUrls[i]
	}

	// Validation

	b.config.InstallTimeout, err = time.ParseDuration(b.config.RawInstallTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed to parse install_timeout: %s", err))
	}

	if b.config.ISOName == "" {

		// If ISO name is not specified, assume a URL and checksum has been provided.

		if b.config.ISOChecksumType == "" {
			errs = packer.MultiErrorAppend(
				errs, errors.New("The iso_checksum_type must be specified."))
		} else {
			b.config.ISOChecksumType = strings.ToLower(b.config.ISOChecksumType)
			if b.config.ISOChecksumType != "none" {
				if b.config.ISOChecksum == "" {
					errs = packer.MultiErrorAppend(
						errs, errors.New("Due to the file size being large, an iso_checksum is required."))
				} else {
					b.config.ISOChecksum = strings.ToLower(b.config.ISOChecksum)
				}

				if hash := common.HashForType(b.config.ISOChecksumType); hash == nil {
					errs = packer.MultiErrorAppend(
						errs, fmt.Errorf("Unsupported checksum type: %s", b.config.ISOChecksumType))
				}

			}
		}

		if len(b.config.ISOUrls) == 0 {
			if b.config.ISOUrl == "" {
				errs = packer.MultiErrorAppend(
					errs, errors.New("One of iso_url or iso_urls must be specified."))
			} else {
				b.config.ISOUrls = []string{b.config.ISOUrl}
			}
		} else if b.config.ISOUrl != "" {
			errs = packer.MultiErrorAppend(
				errs, errors.New("Only one of iso_url or iso_urls may be specified."))
		}

		for i, url := range b.config.ISOUrls {
			b.config.ISOUrls[i], err = common.DownloadableURL(url)
			if err != nil {
				errs = packer.MultiErrorAppend(
					errs, fmt.Errorf("Failed to parse iso_urls[%d]: %s", i, err))
			}
		}
	} else {

		// An ISO name has been provided. It should be attached from an available SR.

	}

	if len(errs.Errors) > 0 {
		retErr = errors.New(errs.Error())
	}

	return nil, retErr

}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	//Setup XAPI client
	client := xsclient.NewXenAPIClient(b.config.HostIp, b.config.Username, b.config.Password)

	err := client.Login()
	if err != nil {
		return nil, err.(error)
	}
	ui.Say("XAPI client session established")

	client.GetHosts()

	//Share state between the other steps using a statebag
	state := new(multistep.BasicStateBag)
	state.Put("cache", cache)
	state.Put("client", client)
	state.Put("config", b.config)
	state.Put("commonconfig", b.config.CommonConfig)
	state.Put("hook", hook)
	state.Put("ui", ui)

	httpReqChan := make(chan string, 1)

	//Build the steps
	download_steps := []multistep.Step{
		&common.StepDownload{
			Checksum:     b.config.ISOChecksum,
			ChecksumType: b.config.ISOChecksumType,
			Description:  "ISO",
			ResultKey:    "iso_path",
			Url:          b.config.ISOUrls,
		},
	}

	steps := []multistep.Step{
		&xscommon.StepPrepareOutputDir{
			Force: b.config.PackerForce,
			Path:  b.config.OutputDir,
		},
		&common.StepCreateFloppy{
			Files: b.config.FloppyFiles,
		},
		&xscommon.StepHTTPServer{
			Chan: httpReqChan,
		},
		&xscommon.StepUploadVdi{
			VdiNameFunc: func() string {
				return "Packer-floppy-disk"
			},
			ImagePathFunc: func() string {
				if floppyPath, ok := state.GetOk("floppy_path"); ok {
					return floppyPath.(string)
				}
				return ""
			},
			VdiUuidKey: "floppy_vdi_uuid",
		},
		&xscommon.StepUploadVdi{
			VdiNameFunc: func() string {
				if len(b.config.ISOUrls) > 0 {
					return path.Base(b.config.ISOUrls[0])
				}
				return ""
			},
			ImagePathFunc: func() string {
				if isoPath, ok := state.GetOk("iso_path"); ok {
					return isoPath.(string)
				}
				return ""
			},
			VdiUuidKey: "iso_vdi_uuid",
		},
		&xscommon.StepFindVdi{
			VdiName:    b.config.ToolsIsoName,
			VdiUuidKey: "tools_vdi_uuid",
		},
		&xscommon.StepFindVdi{
			VdiName:    b.config.ISOName,
			VdiUuidKey: "isoname_vdi_uuid",
		},
		new(stepCreateInstance),
		&xscommon.StepAttachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
			VdiType:    xsclient.Floppy,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "iso_vdi_uuid",
			VdiType:    xsclient.CD,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "isoname_vdi_uuid",
			VdiType:    xsclient.CD,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "tools_vdi_uuid",
			VdiType:    xsclient.CD,
		},
		new(xscommon.StepStartVmPaused),
		new(xscommon.StepGetConsoleLocation),
		&xscommon.StepTCPVNCProxy{
			RemoteConsole: xscommon.InstanceConsoleLocation,
			HostPortMin:   b.config.HostPortMin,
			HostPortMax:   b.config.HostPortMax,
		},
		new(xscommon.StepBootWait),
		&xscommon.StepTypeBootCommand{
			BootCommand: b.config.BootCommand,
			VMName:      b.config.VMName,
			Ctx:         b.config.ctx,
		},
		&xscommon.StepWaitForIP{
			Chan:    httpReqChan,
			Timeout: b.config.InstallTimeout, // @todo change this
		},
		&xscommon.StepForwardPortOverSSH{
			RemotePort:  xscommon.InstanceSSHPort,
			RemoteDest:  xscommon.InstanceSSHIP,
			HostPortMin: b.config.HostPortMin,
			HostPortMax: b.config.HostPortMax,
			ResultKey:   "local_ssh_port",
		},
		&communicator.StepConnect{
			Config:    &b.config.SSHConfig.Comm,
			Host:      xscommon.CommHost,
			SSHConfig: xscommon.SSHConfigFunc(b.config.CommonConfig.SSHConfig),
			SSHPort:   xscommon.SSHPort,
		},
		new(xscommon.StepShutdown),
		&xscommon.StepDetachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
		},
		&xscommon.StepDetachVdi{
			VdiUuidKey: "iso_vdi_uuid",
		},
		new(xscommon.StepStartVmPaused),
		new(xscommon.StepBootWait),
		&communicator.StepConnectSSH{
			Config:    &b.config.SSHConfig.Comm,
			Host:      xscommon.CommHost,
			SSHConfig: xscommon.SSHConfigFunc(b.config.CommonConfig.SSHConfig),
			SSHPort:   xscommon.SSHPort,
		},
		new(common.StepProvision),
		new(xscommon.StepShutdown),
		&xscommon.StepDetachVdi{
			VdiUuidKey: "tools_vdi_uuid",
		},
		new(xscommon.StepExport),
	}

	if b.config.ISOName == "" {
		steps = append(download_steps, steps...)
	}

	b.runner = &multistep.BasicRunner{Steps: steps}
	b.runner.Run(state)

	if rawErr, ok := state.GetOk("error"); ok {
		return nil, rawErr.(error)
	}

	// If we were interrupted or cancelled, then just exit.
	if _, ok := state.GetOk(multistep.StateCancelled); ok {
		return nil, errors.New("Build was cancelled.")
	}
	if _, ok := state.GetOk(multistep.StateHalted); ok {
		return nil, errors.New("Build was halted.")
	}

	artifact, _ := xscommon.NewArtifact(b.config.OutputDir)

	return artifact, nil
}

func (b *Builder) Cancel() {
	if b.runner != nil {
		log.Println("Cancelling the step runner...")
		b.runner.Cancel()
	}
	fmt.Println("Cancelling the builder")
}
