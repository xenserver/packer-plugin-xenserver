// Copyright (c) 2026. Citrix Systems, Inc.

package clone

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	hconfig "github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	"xenapi"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

type Builder struct {
	config CloneConfig
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

	// CommonConfig.Prepare validates against the communicator, so default it first.
	self.config.SSHConfig.ApplyCommunicatorDefault()

	commonWarns, commonErrs := self.config.CommonConfig.Prepare(&self.config.PackerConfig)
	warns = append(warns, commonWarns...)
	errs = packer.MultiErrorAppend(errs, commonErrs...)
	errs = packer.MultiErrorAppend(errs, self.config.SSHConfig.Prepare(self.config.GetInterpContext())...)

	if self.config.Comm.Type == "none" {
		if self.config.bootsTheGuest() {
			warns = append(warns, fmt.Sprintf("communicator = \"none\": the build waits up to %s "+
				"(install_timeout) for the guest to power itself off, and fails if it does not.",
				self.config.InstallTimeout))
		} else {
			warns = append(warns, "communicator = \"none\" with no cd_files and no boot_command: "+
				"nothing would run inside the guest, so the build never starts it and exports the "+
				"clone as it came off the source. Set cd_files or boot_command to give it something "+
				"to run, or a communicator to provision it.")
		}
	}

	switch {
	case self.config.CloneTemplate != "" && self.config.CloneVM != "":
		errs = packer.MultiErrorAppend(errs, errors.New("only one of clone_template or clone_vm may be "+
			"specified: clone_template names a source template, clone_vm a source VM."))

	case self.config.CloneTemplate == "" && self.config.CloneVM == "":
		errs = packer.MultiErrorAppend(errs, errors.New("one of clone_template or clone_vm must be "+
			"specified: clone_template names a source template, clone_vm a source VM."))
	}

	if self.config.CorePerSocket != nil && self.config.VCPUs == nil {
		warns = append(warns, "cores_per_socket is set but vcpus is not, so the clone keeps the source's "+
			"CPU topology and cores_per_socket has no effect. Set vcpus to apply it.")
	}

	if err := xscommon.ValidateCPUTopology(self.config.VCPUs, self.config.CorePerSocket); err != nil {
		errs = packer.MultiErrorAppend(errs, err)
	}

	// A clone inherits its firmware, which Prepare has no session to read.
	if self.config.SecureBootEnabled() && self.config.Firmware != "uefi" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("secure_boot can only be enabled with firmware = \"uefi\", which must be set explicitly on a clone."))
	}

	// XAPI only rejects a vTPM on BIOS at VM.start, with a bare NOT_IMPLEMENTED.
	if self.config.VTPMEnabled() && self.config.Firmware != "uefi" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("vtpm can only be enabled with firmware = \"uefi\", which must be set explicitly on a clone."))
	}

	if len(errs.Errors) > 0 {
		retErr = errors.New(errs.Error())
	}

	return nil, warns, retErr

}

func (self *Builder) connectStep() *communicator.StepConnect {
	return xscommon.ConnectStep(&self.config.SSHConfig)
}

// finalSteps snapshots before converting to a template: XAPI refuses VM.snapshot on a template.
func finalSteps() []multistep.Step {
	return []multistep.Step{
		new(xscommon.StepShutdown),
		&xscommon.StepDetachVdi{
			VdiUuidKey: "cd_vdi_uuid",
		},
		new(xscommon.StepCreateSnapshot),
		new(xscommon.StepSetVmToTemplate),
		new(xscommon.StepExport),
	}
}

func (self *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	c, err := xscommon.NewXenAPIClient(self.config.HostIp, self.config.Username, self.config.Password, self.config.InsecureEnabled(), self.config.CACertPath)

	if err != nil {
		return nil, err
	}
	defer c.Logout()
	ui.Say("XAPI client session established")

	// Before the steps, so that an unusable source fails before -force deletes the output directory.
	src := self.config.cloneSource()
	sourceVM, sourceRecord, err := xscommon.ResolveCloneSource(c, src)
	if err != nil {
		return nil, err
	}

	if err := xscommon.EnsureCloneSourceIsHalted(src, sourceRecord, self.config.VMSR != ""); err != nil {
		return nil, err
	}

	if self.config.DiskSizeGB != nil {
		if err := xscommon.EnsureCloneSourceHasPrimaryDisk(c, src, sourceVM); err != nil {
			return nil, err
		}
	}

	// Only on the path that boots the VM: the other one never writes firmware to the clone.
	if self.config.bootsTheGuest() {
		if err := xscommon.CloneSourceSupportsFirmware(src, sourceRecord,
			self.config.Firmware, self.config.removesVTPM()); err != nil {
			return nil, err
		}
	}

	state := new(multistep.BasicStateBag)
	state.Put("client", c)
	state.Put("source_vm", sourceVM)
	state.Put("config", self.config)
	state.Put("commonconfig", self.config.CommonConfig)
	state.Put("hook", hook)
	state.Put("ui", ui)

	httpReqChan := make(chan string, 1)

	steps := []multistep.Step{
		&xscommon.StepPrepareOutputDir{
			Force: self.config.PackerForce,
			Path:  self.config.OutputDir,
		},
		&commonsteps.StepCreateCD{
			Files: self.config.CDFiles,
			Label: "cddata",
		},
		new(xscommon.StepHTTPIPDiscover),
		&xscommon.StepHTTPServer{
			Chan: httpReqChan,
		},
		&xscommon.StepUploadVdi{
			VdiNameFunc: func() string {
				return "Packer-cd-disk"
			},
			ImagePathFunc: func() string {
				if cdPath, ok := state.GetOk("cd_path"); ok {
					return cdPath.(string)
				}
				return ""
			},
			VdiUuidKey: "cd_vdi_uuid",
		},
		new(stepCloneInstance),
		&xscommon.StepAttachvGPU{
			VGPUName: self.config.VGPUProfile,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "cd_vdi_uuid",
			VdiType:    xenapi.VbdTypeCD,
		},
	}

	// StepConnect runs even without a boot: with communicator = "none" it only puts the SDK's
	// null communicator in the state bag, which StepProvision needs.
	if self.config.bootsTheGuest() {
		steps = append(steps,
			new(xscommon.StepStartVmPaused),
			new(xscommon.StepBootWait),
			&xscommon.StepTypeBootCommand{
				Ctx: *self.config.GetInterpContext(),
			},
			xscommon.InstallWaitStep(self.config.Comm.Type, httpReqChan, self.config.InstallTimeout),
		)
	} else {
		ui.Say("Nothing would run inside this guest, so the build does not start it: the clone " +
			"is exported as it came off the source. Set cd_files or boot_command to give it " +
			"something to run, or a communicator to provision it.")
	}

	steps = append(steps, self.connectStep(), new(commonsteps.StepProvision))
	steps = append(steps, finalSteps()...)

	self.runner = &multistep.BasicRunner{Steps: steps}
	self.runner.Run(ctx, state)

	if rawErr, ok := state.GetOk("error"); ok {
		return nil, rawErr.(error)
	}

	if _, ok := state.GetOk(multistep.StateCancelled); ok {
		return nil, errors.New("Build was cancelled.")
	}
	if _, ok := state.GetOk(multistep.StateHalted); ok {
		return nil, errors.New("Build was halted.")
	}

	artifact, err := xscommon.NewArtifact(self.config.OutputDir)
	if err != nil {
		return nil, fmt.Errorf("could not list the exported files in %s: %w", self.config.OutputDir, err)
	}

	return artifact, nil
}
