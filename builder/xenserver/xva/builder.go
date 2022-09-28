package xva

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	commonsteps "github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	hconfig "github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	xsclient "github.com/terra-farm/go-xen-api-client"
	xscommon "github.com/xenserver/packer-builder-xenserver/builder/xenserver/common"
)

type Builder struct {
	config xscommon.Config
	runner multistep.Runner
}

func (self *Builder) ConfigSpec() hcldec.ObjectSpec { return self.config.FlatMapstructure().HCL2Spec() }

func (self *Builder) Prepare(raws ...interface{}) (params []string, warns []string, retErr error) {

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
		errs = packer.MultiErrorAppend(errs, err)
	}

	errs = packer.MultiErrorAppend(
		errs, self.config.CommonConfig.Prepare(self.config.GetInterpContext(), &self.config.PackerConfig)...)

	// Set default values
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

	// Validation

	if self.config.SourcePath == "" {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("A source_path must be specified"))
	}

	if len(errs.Errors) > 0 {
		retErr = errors.New(errs.Error())
	}

	return nil, nil, retErr

}

func (self *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	//Setup XAPI client
	c, err := xscommon.NewXenAPIClient(self.config.HostIp, self.config.Username, self.config.Password)

	if err != nil {
		return nil, err
	}

	ui.Say("XAPI client session established")

	c.GetClient().Host.GetAll(c.GetSessionRef())

	//Share state between the other steps using a statebag
	state := new(multistep.BasicStateBag)
	state.Put("client", c)
	// state.Put("config", self.config)
	state.Put("commonconfig", self.config.CommonConfig)
	state.Put("hook", hook)
	state.Put("ui", ui)

	httpReqChan := make(chan string, 1)

	//Build the steps
	steps := []multistep.Step{
		&xscommon.StepPrepareOutputDir{
			Force: self.config.PackerForce,
			Path:  self.config.OutputDir,
		},
		&commonsteps.StepCreateFloppy{
			Files: self.config.FloppyFiles,
		},
		new(xscommon.StepHTTPServer),
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
		&xscommon.StepFindVdi{
			VdiName:    self.config.ToolsIsoName,
			VdiUuidKey: "tools_vdi_uuid",
		},
		new(stepImportInstance),
		&xscommon.StepAttachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
			VdiType:    xsclient.VbdTypeFloppy,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "tools_vdi_uuid",
			VdiType:    xsclient.VbdTypeCD,
		},
		new(xscommon.StepStartVmPaused),
		new(xscommon.StepSetVmHostSshAddress),
		new(xscommon.StepBootWait),
		&xscommon.StepTypeBootCommand{
			Ctx: *self.config.GetInterpContext(),
		},
		&xscommon.StepWaitForIP{
			Chan:    httpReqChan,
			Timeout: 300 * time.Minute, /*self.config.InstallTimeout*/ // @todo change this
		},
		&communicator.StepConnect{
			Config:    &self.config.SSHConfig.Comm,
			Host:      xscommon.CommHost,
			SSHConfig: xscommon.SSHConfigFunc(self.config.CommonConfig.SSHConfig),
			SSHPort:   xscommon.SSHPort,
		},
		new(commonsteps.StepProvision),
		new(xscommon.StepShutdown),
		&xscommon.StepDetachVdi{
			VdiUuidKey: "floppy_vdi_uuid",
		},
		&xscommon.StepDetachVdi{
			VdiUuidKey: "tools_vdi_uuid",
		},
		new(xscommon.StepExport),
	}

	self.runner = &multistep.BasicRunner{Steps: steps}
	self.runner.Run(ctx, state)

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
