// Copyright (c) 2026. Citrix Systems, Inc.

package iso

import (
	"context"
	"errors"
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	commonsteps "github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	hconfig "github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	"xenapi"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

type Builder struct {
	config ISOConfig
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
		warns = append(warns, fmt.Sprintf("communicator = \"none\": the build waits up to %s "+
			"(install_timeout) for the guest to power itself off, and fails if it does not. Make the "+
			"unattended installation shut the guest down when it finishes.", self.config.InstallTimeout))
	}

	if self.config.GuestToolsISO != "" {
		warns = append(warns, "guest_tools_iso is deprecated: XenServer 8 and later ship no guest-tools ISO. "+
			"Set it only against XCP-ng or legacy pools that still publish one.")
	}

	// Required rather than defaulted: the template decides the VM's virtual hardware.
	if strings.TrimSpace(self.config.CloneTemplate) == "" {
		errs = packer.MultiErrorAppend(errs, errors.New("clone_template must be specified: the "+
			"name-label of the template to clone as the VM's starting point. "+
			xscommon.TemplateLookupHint))
	}

	// Unset sizing keeps the template's; genericTemplate has none, and no default suits every guest.
	for _, option := range missingGenericSizing(self.config) {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("%s must be specified with clone_template "+
			"'%s': that template describes no guest OS, so it supplies no %s for the VM to inherit.",
			option, genericTemplate, option))
	}

	// Only a warning: Prepare cannot tell whether the ISO installs Windows. Zero is already
	// an error from CommonConfig.Prepare.
	if self.config.VMMemoryGB != nil && *self.config.VMMemoryGB > 0 &&
		*self.config.VMMemoryGB < minWindowsMemoryGB {
		warns = append(warns, fmt.Sprintf("vm_memory_gb = %d is below the %d GB Microsoft documents as the "+
			"minimum for installing any supported Windows guest; 1 GB is the exact configuration it documents "+
			"as failing to install Windows Server 2025. Windows Setup will not report this: it stalls, and the "+
			"build fails after install_timeout. Ignore this if the ISO installs something else.",
			*self.config.VMMemoryGB, minWindowsMemoryGB))
	}

	// Skipped without clone_template: the build has already failed above.
	if self.config.CorePerSocket != nil && self.config.VCPUs == nil &&
		strings.TrimSpace(self.config.CloneTemplate) != "" {
		warns = append(warns, fmt.Sprintf("cores_per_socket is ignored without vcpus: the VM keeps the CPU "+
			"topology of clone_template '%s'. Set vcpus as well to choose the topology.", self.config.CloneTemplate))
	}

	if err := xscommon.ValidateCPUTopology(self.config.VCPUs, self.config.CorePerSocket); err != nil {
		errs = packer.MultiErrorAppend(errs, err)
	}

	// firmware is deliberately not defaulted: unset keeps the template's (Windows 11 needs uefi).

	hasISOUrl := self.config.RawSingleISOUrl != "" || len(self.config.ISOUrls) > 0
	switch {
	case self.config.ISOName != "" && hasISOUrl:
		errs = packer.MultiErrorAppend(errs, errors.New("only one of iso_name or iso_url/iso_urls may be "+
			"specified: iso_name takes an ISO already present on iso_sr, iso_url downloads one and uploads "+
			"it to the pool."))

	case self.config.ISOName == "" && !hasISOUrl:
		errs = packer.MultiErrorAppend(errs, errors.New("one of iso_name or iso_url/iso_urls must be "+
			"specified: iso_name names an ISO already present on iso_sr, iso_url an ISO to download."))

	case hasISOUrl:
		// The SDK folds iso_url into iso_urls and normalizes iso_checksum and iso_target_extension.
		isoWarns, isoErrs := self.config.ISOConfig.Prepare(self.config.GetInterpContext())
		warns = append(warns, isoWarns...)
		errs = packer.MultiErrorAppend(errs, isoErrs...)

	default:
		if unused := setDownloadOptions(&self.config.ISOConfig); len(unused) > 0 {
			errs = packer.MultiErrorAppend(errs, fmt.Errorf("iso_name attaches an ISO already present on "+
				"iso_sr, so nothing is downloaded: the download options set here (%s) apply to "+
				"iso_url/iso_urls only.", strings.Join(unused, ", ")))
		}
	}

	// A downloaded ISO is uploaded to vm_sr, so only iso_name needs iso_sr.
	if self.config.ISOName != "" && self.config.ISOSR == "" {
		errs = packer.MultiErrorAppend(errs, errors.New("iso_sr must be specified with iso_name: it names "+
			"the SR the install ISO is looked up in."))
	}

	// firmware must be explicit: Prepare has no XenAPI session to read the template's.
	if self.config.VTPMEnabled && self.config.Firmware != "uefi" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("vtpm can only be enabled with firmware = \"uefi\", which must be set explicitly."))
	}

	if self.config.SecureBootEnabled() && self.config.Firmware != "uefi" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("secure_boot can only be enabled with firmware = \"uefi\", which must be set explicitly."))
	}

	if len(errs.Errors) > 0 {
		retErr = errors.New(errs.Error())
	}

	return nil, warns, retErr

}

// missingGenericSizing leaves out cores_per_socket: a VM without a topology gets one core per socket.
func missingGenericSizing(c ISOConfig) []string {
	if c.CloneTemplate != genericTemplate {
		return nil
	}

	var missing []string
	for _, option := range []struct {
		name  string
		value *uint
	}{
		{"vcpus", c.VCPUs},
		{"vm_memory_gb", c.VMMemoryGB},
		{"disk_size_gb", c.DiskSizeGB},
	} {
		if option.value == nil {
			missing = append(missing, option.name)
		}
	}

	return missing
}

// bootCommandStep sweeps the UEFI "Press any key to boot from CD" prompt, which a single key
// press cannot be timed to hit.
func bootCommandStep(ctx interpolate.Context) *xscommon.StepTypeBootCommand {
	return &xscommon.StepTypeBootCommand{
		Ctx:               ctx,
		SweepUEFICDPrompt: true,
	}
}

func setDownloadOptions(c *commonsteps.ISOConfig) []string {
	var set []string
	for option, value := range map[string]string{
		"iso_checksum":         c.ISOChecksum,
		"iso_target_path":      c.TargetPath,
		"iso_target_extension": c.TargetExtension,
	} {
		if value != "" {
			set = append(set, option)
		}
	}
	sort.Strings(set)

	return set
}

// finalSteps snapshots before converting to a template: XAPI refuses VM.snapshot on a template.
func finalSteps() []multistep.Step {
	return []multistep.Step{
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
			VdiUuidKey: "cd_vdi_uuid",
		},
		new(xscommon.StepCreateSnapshot),
		new(xscommon.StepSetVmToTemplate),
		new(xscommon.StepExport),
	}
}

func (self *Builder) downloadSteps() []multistep.Step {
	return []multistep.Step{
		&commonsteps.StepDownload{
			Checksum:    self.config.ISOChecksum,
			Description: "ISO",
			ResultKey:   "iso_path",
			Url:         self.config.ISOUrls,
			Extension:   self.config.TargetExtension,
			TargetPath:  self.config.TargetPath,
		},
	}
}

func (self *Builder) connectStep() *communicator.StepConnect {
	return xscommon.ConnectStep(&self.config.SSHConfig)
}

func (self *Builder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	c, err := xscommon.NewXenAPIClient(self.config.HostIp, self.config.Username, self.config.Password, self.config.InsecureEnabled(), self.config.CACertPath)

	if err != nil {
		return nil, err
	}
	defer c.Logout()
	ui.Say("XAPI client session established")

	// Resolved before the steps so that an unusable template fails before any ISO download.
	src := xscommon.ISOTemplateSource(self.config.CloneTemplate)
	sourceVM, sourceRecord, err := xscommon.ResolveCloneSource(c, src)
	if err != nil {
		return nil, err
	}

	// removesVtpm is false: vtpm = false here only declines to attach one.
	if err := xscommon.CloneSourceSupportsFirmware(src, sourceRecord, self.config.Firmware, false); err != nil {
		return nil, err
	}

	// Bytes, not GB: a template disk that is not a whole number of GB must not be rounded down.
	diskBytes := 0
	if self.config.DiskSizeGB != nil {
		diskBytes = int(*self.config.DiskSizeGB) * xscommon.BytesPerGB
	} else if diskBytes, err = xscommon.TemplateDiskSize(c, src, sourceVM, sourceRecord); err != nil {
		return nil, err
	}

	state := new(multistep.BasicStateBag)
	state.Put("client", c)
	state.Put("source_vm", sourceVM)
	state.Put("disk_size_bytes", diskBytes)
	state.Put("config", self.config)
	state.Put("commonconfig", self.config.CommonConfig)
	state.Put("hook", hook)
	state.Put("ui", ui)

	httpReqChan := make(chan string, 1)

	steps := []multistep.Step{
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
			VdiName:    self.config.GuestToolsISO,
			VdiUuidKey: "tools_vdi_uuid",
		},
		&xscommon.StepFindVdi{
			VdiName:    self.config.ISOName,
			SRName:     self.config.ISOSR,
			VdiUuidKey: "isoname_vdi_uuid",
		},
		new(stepCreateInstance),
		&xscommon.StepAttachvTPM{
			EnablevTPM: self.config.VTPMEnabled,
		},
		&xscommon.StepAttachvGPU{
			VGPUName: self.config.VGPUProfile,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "iso_vdi_uuid",
			VdiType:    xenapi.VbdTypeCD,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "isoname_vdi_uuid",
			VdiType:    xenapi.VbdTypeCD,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "cd_vdi_uuid",
			VdiType:    xenapi.VbdTypeCD,
		},
		&xscommon.StepAttachVdi{
			VdiUuidKey: "tools_vdi_uuid",
			VdiType:    xenapi.VbdTypeCD,
		},
		new(xscommon.StepStartVmPaused),
		new(xscommon.StepBootWait),
		bootCommandStep(*self.config.GetInterpContext()),
		xscommon.InstallWaitStep(self.config.Comm.Type, httpReqChan, self.config.InstallTimeout),
		self.connectStep(),
		new(commonsteps.StepProvision),
	}
	steps = append(steps, finalSteps()...)

	// Before downloadSteps, so that a refused output directory fails before a multi-GB download.
	preSteps := []multistep.Step{
		&xscommon.StepPrepareOutputDir{
			Force: self.config.PackerForce,
			Path:  self.config.OutputDir,
		},
	}
	if self.config.ISOName == "" {
		preSteps = append(preSteps, self.downloadSteps()...)
	}
	steps = append(preSteps, steps...)

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
