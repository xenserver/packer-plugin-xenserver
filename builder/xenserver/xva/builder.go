package xva

import (
	"errors"
	"fmt"
	"log"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
	xscommon "github.com/rdobson/packer-builder-xenserver/builder/xenserver/common"
)

type config struct {
	common.PackerConfig   `mapstructure:",squash"`
	xscommon.CommonConfig `mapstructure:",squash"`

	SourcePath string `mapstructure:"source_path"`
	VMMemory      uint   `mapstructure:"vm_memory"`

	PlatformArgs map[string]string `mapstructure:"platform_args"`

	tpl *packer.ConfigTemplate
}

type Builder struct {
	config config
	runner multistep.Runner
}

func (self *Builder) Prepare(raws ...interface{}) (params []string, retErr error) {

	md, err := common.DecodeConfig(&self.config, raws...)
	if err != nil {
		return nil, err
	}

	self.config.tpl, err = packer.NewConfigTemplate()
	if err != nil {
		return nil, err
	}
	self.config.tpl.UserVars = self.config.PackerUserVars

	errs := common.CheckUnusedConfig(md)
	errs = packer.MultiErrorAppend(
		errs, self.config.CommonConfig.Prepare(self.config.tpl, &self.config.PackerConfig)...)

	// Set default values

	if self.config.VMMemory == 0 {
		self.config.VMMemory = 1024
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
		"source_path":   &self.config.SourcePath,
		"network_name":   &self.config.NetworkName,
	}

	for n, ptr := range templates {
		var err error
		*ptr, err = self.config.tpl.Process(*ptr, nil)
		if err != nil {
			errs = packer.MultiErrorAppend(errs, fmt.Errorf("Error processing %s: %s", n, err))
		}
	}

	// Validation

	if self.config.SourcePath == "" {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("A source_path must be specified"))
	}

	if len(errs.Errors) > 0 {
		retErr = errors.New(errs.Error())
	}

	return nil, retErr

}

func (self *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	//Setup XAPI client
	client := xscommon.NewXenAPIClient(self.config.HostIp, self.config.Username, self.config.Password)

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

	//Build the steps
	steps := []multistep.Step{
		&xscommon.StepPrepareOutputDir{
			Force: self.config.PackerForce,
			Path:  self.config.OutputDir,
		},
		&common.StepCreateFloppy{
			Files: self.config.FloppyFiles,
		},
		new(xscommon.StepHTTPServer),
		&xscommon.StepUploadVdi{
			VdiName: "Packer-floppy-disk",
			ImagePathFunc: func() string {
				if floppyPath, ok := state.GetOk("floppy_path"); ok {
					return floppyPath.(string)
				}
				return ""
			},
			VdiUuidKey: "floppy_vdi_uuid",
		},
		&xscommon.StepFindVdi{
			VdiName:    self.config.ToolsIsoName,
			VdiUuidKey: "tools_vdi_uuid",
		},
		new(stepImportInstance),
		&xscommon.StepAttachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
			VdiType:    xscommon.Floppy,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "tools_vdi_uuid",
			VdiType:    xscommon.CD,
		},
		new(xscommon.StepRemoveDevices),
		new(xscommon.StepStartOnHIMN),
		&xscommon.StepForwardPortOverSSH{
			RemotePort:  xscommon.HimnSSHPort,
			RemoteDest:  xscommon.HimnSSHIP,
			HostPortMin: self.config.HostPortMin,
			HostPortMax: self.config.HostPortMax,
			ResultKey:   "local_ssh_port",
		},
		&common.StepConnectSSH{
			SSHAddress:     xscommon.SSHLocalAddress,
			SSHConfig:      xscommon.SSHConfig,
			SSHWaitTimeout: self.config.SSHWaitTimeout,
		},
		new(common.StepProvision),
		&xscommon.StepDetachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
		},
		&xscommon.StepDetachVdi{
			VdiUuidKey: "iso_vdi_uuid",
		},
		new(xscommon.StepShutdownAndExport),
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
