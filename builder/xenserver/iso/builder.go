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
	hconfig "github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
	xsclient "github.com/xenserver/go-xenserver-client"
	xscommon "github.com/xenserver/packer-builder-xenserver/builder/xenserver/common"
)

type config struct {
	common.PackerConfig   `mapstructure:",squash"`
	xscommon.CommonConfig `mapstructure:",squash"`

	VCPUsMax       uint              `mapstructure:"vcpus_max"`
	VCPUsAtStartup uint              `mapstructure:"vcpus_atstartup"`
	VMMemory       uint              `mapstructure:"vm_memory"`
	DiskSize       uint              `mapstructure:"disk_size"`
	CloneTemplate  string            `mapstructure:"clone_template"`
	VMOtherConfig  map[string]string `mapstructure:"vm_other_config"`

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

type Builder struct {
	config config
	runner multistep.Runner
}

func (self *Builder) Prepare(raws ...interface{}) (params []string, retErr error) {

	var errs *packer.MultiError

	err := hconfig.Decode(&self.config, &hconfig.DecodeOpts{
		Interpolate: true,
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
		errs, self.config.CommonConfig.Prepare(&self.config.ctx, &self.config.PackerConfig)...)
	errs = packer.MultiErrorAppend(errs, self.config.SSHConfig.Prepare(&self.config.ctx)...)

	// Set default values

	if self.config.RawInstallTimeout == "" {
		self.config.RawInstallTimeout = "200m"
	}

	if self.config.DiskSize == 0 {
		self.config.DiskSize = 40000
	}

	if self.config.VCPUsMax == 0 {
		self.config.VCPUsMax = 1
	}

	if self.config.VCPUsAtStartup == 0 {
		self.config.VCPUsAtStartup = 1
	}

	if self.config.VCPUsAtStartup > self.config.VCPUsMax {
		self.config.VCPUsAtStartup = self.config.VCPUsMax
	}

	if self.config.VMMemory == 0 {
		self.config.VMMemory = 1024
	}

	if self.config.CloneTemplate == "" {
		self.config.CloneTemplate = "Other install media"
	}

	if len(self.config.PlatformArgs) == 0 {
		pargs := make(map[string]string)
		pargs["viridian"] = "false"
		pargs["nx"] = "true"
		pargs["pae"] = "true"
		pargs["apic"] = "true"
		pargs["timeoffset"] = "0"
		pargs["acpi"] = "1"
		self.config.PlatformArgs = pargs
	}

	// Template substitution

	templates := map[string]*string{
		"clone_template":    &self.config.CloneTemplate,
		"iso_checksum":      &self.config.ISOChecksum,
		"iso_checksum_type": &self.config.ISOChecksumType,
		"iso_url":           &self.config.ISOUrl,
		"iso_name":          &self.config.ISOName,
		"install_timeout":   &self.config.RawInstallTimeout,
	}
	for i := range self.config.ISOUrls {
		templates[fmt.Sprintf("iso_urls[%d]", i)] = &self.config.ISOUrls[i]
	}

	// Validation

	self.config.InstallTimeout, err = time.ParseDuration(self.config.RawInstallTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed to parse install_timeout: %s", err))
	}

	if self.config.ISOName == "" {

		// If ISO name is not specified, assume a URL and checksum has been provided.

		if self.config.ISOChecksumType == "" {
			errs = packer.MultiErrorAppend(
				errs, errors.New("The iso_checksum_type must be specified."))
		} else {
			self.config.ISOChecksumType = strings.ToLower(self.config.ISOChecksumType)
			if self.config.ISOChecksumType != "none" {
				if self.config.ISOChecksum == "" {
					errs = packer.MultiErrorAppend(
						errs, errors.New("Due to the file size being large, an iso_checksum is required."))
				} else {
					self.config.ISOChecksum = strings.ToLower(self.config.ISOChecksum)
				}

				if hash := common.HashForType(self.config.ISOChecksumType); hash == nil {
					errs = packer.MultiErrorAppend(
						errs, fmt.Errorf("Unsupported checksum type: %s", self.config.ISOChecksumType))
				}

			}
		}

		if len(self.config.ISOUrls) == 0 {
			if self.config.ISOUrl == "" {
				errs = packer.MultiErrorAppend(
					errs, errors.New("One of iso_url or iso_urls must be specified."))
			} else {
				self.config.ISOUrls = []string{self.config.ISOUrl}
			}
		} else if self.config.ISOUrl != "" {
			errs = packer.MultiErrorAppend(
				errs, errors.New("Only one of iso_url or iso_urls may be specified."))
		}

		for i, url := range self.config.ISOUrls {
			self.config.ISOUrls[i], err = common.DownloadableURL(url)
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

func (self *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	//Setup XAPI client
	client := xsclient.NewXenAPIClient(self.config.HostIp, self.config.Username, self.config.Password)

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
	state.Put("config", self.config)
	state.Put("commonconfig", self.config.CommonConfig)
	state.Put("hook", hook)
	state.Put("ui", ui)

	httpReqChan := make(chan string, 1)

	//Build the steps
	download_steps := []multistep.Step{
		&common.StepDownload{
			Checksum:     self.config.ISOChecksum,
			ChecksumType: self.config.ISOChecksumType,
			Description:  "ISO",
			ResultKey:    "iso_path",
			Url:          self.config.ISOUrls,
		},
	}

	steps := []multistep.Step{
		&xscommon.StepPrepareOutputDir{
			Force: self.config.PackerForce,
			Path:  self.config.OutputDir,
		},
		&common.StepCreateFloppy{
			Files: self.config.FloppyFiles,
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
				if len(self.config.ISOUrls) > 0 {
					return path.Base(self.config.ISOUrls[0])
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
			VdiName:    self.config.ToolsIsoName,
			VdiUuidKey: "tools_vdi_uuid",
		},
		&xscommon.StepFindVdi{
			VdiName:    self.config.ISOName,
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
		new(xscommon.StepSetVmHostSshAddress),
		new(xscommon.StepGetVNCPort),
		&xscommon.StepForwardPortOverSSH{
			RemotePort:  xscommon.InstanceVNCPort,
			RemoteDest:  xscommon.InstanceVNCIP,
			HostPortMin: self.config.HostPortMin,
			HostPortMax: self.config.HostPortMax,
			ResultKey:   "local_vnc_port",
		},
		new(xscommon.StepBootWait),
		&xscommon.StepTypeBootCommand{
			Ctx: self.config.ctx,
		},
		&xscommon.StepWaitForIP{
			Chan:    httpReqChan,
			Timeout: self.config.InstallTimeout, // @todo change this
		},
		&xscommon.StepForwardPortOverSSH{
			RemotePort:  xscommon.InstanceSSHPort,
			RemoteDest:  xscommon.InstanceSSHIP,
			HostPortMin: self.config.HostPortMin,
			HostPortMax: self.config.HostPortMax,
			ResultKey:   "local_ssh_port",
		},
		&communicator.StepConnect{
			Config:    &self.config.SSHConfig.Comm,
			Host:      xscommon.CommHost,
			SSHConfig: xscommon.SSHConfigFunc(self.config.CommonConfig.SSHConfig),
			SSHPort:   xscommon.SSHPort,
		},
		new(common.StepProvision),
		new(xscommon.StepShutdown),
		&xscommon.StepDetachVdi{
			VdiUuidKey: "iso_vdi_uuid",
		},
		&xscommon.StepDetachVdi{
			VdiUuidKey: "isoname_vdi_uuid",
		},
		&xscommon.StepDetachVdi{
			VdiUuidKey: "tools_vdi_uuid",
		},
		&xscommon.StepDetachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
		},
		new(xscommon.StepConvertToTemplate),
		new(xscommon.StepExport),
	}

	if self.config.ISOName == "" {
		steps = append(download_steps, steps...)
	}

	self.runner = &multistep.BasicRunner{Steps: steps}
	self.runner.Run(state)

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

	artifact, _ := xscommon.NewArtifact(self.config.OutputDir)

	return artifact, nil
}

func (self *Builder) Cancel() {
	if self.runner != nil {
		log.Println("Cancelling the step runner...")
		self.runner.Cancel()
	}
	fmt.Println("Cancelling the builder")
}
