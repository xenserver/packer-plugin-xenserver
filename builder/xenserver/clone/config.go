// Copyright (c) 2026. Citrix Systems, Inc.

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type CloneConfig
package clone

import (
	"github.com/hashicorp/packer-plugin-sdk/common"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

// The install-media options (`iso_url`, `iso_name`, checksums, `guest_tools_iso`) are absent on purpose.

// CloneConfig is the xenserver-vm-clone builder's configuration: the options
// shared with the iso builder, plus `clone_template`, `clone_vm` and `vtpm`.
type CloneConfig struct {
	common.PackerConfig   `mapstructure:",squash"`
	xscommon.CommonConfig `mapstructure:",squash"`

	// Name label of the template to clone. Exactly one of `clone_template` and
	// `clone_vm` must be set, and the name must match exactly one template.
	CloneTemplate string `mapstructure:"clone_template"`

	// Name label of the VM or snapshot to clone. Exactly one of `clone_template`
	// and `clone_vm` must be set, and the name must match exactly one VM or
	// snapshot. A VM must be halted; snapshot a running VM and name the snapshot
	// instead. A checkpoint cannot be combined with `vm_sr`.
	CloneVM string `mapstructure:"clone_vm"`

	// Attach a virtual TPM. Omit it to keep whatever the source has. `true` adds
	// one, and requires `firmware = "uefi"`, set explicitly. `false` removes the
	// inherited vTPM, with anything sealed to it, and cannot be undone.
	VTPM *bool `mapstructure:"vtpm"`
}

// Prepare has already rejected setting both or neither, so an unset clone_vm means clone_template.
func (c CloneConfig) cloneSource() xscommon.CloneSource {
	if c.CloneVM != "" {
		return xscommon.VMSource(c.CloneVM)
	}
	return xscommon.TemplateSource(c.CloneTemplate)
}

// VTPMEnabled reports whether vtpm is true; check VTPM for nil where "keep the source's" matters.
func (c CloneConfig) VTPMEnabled() bool {
	return c.VTPM != nil && *c.VTPM
}

func (c CloneConfig) removesVTPM() bool {
	return c.VTPM != nil && !*c.VTPM
}

// Booting a clone with nothing to run would only wait out install_timeout, so it is exported unstarted.
func (c CloneConfig) bootsTheGuest() bool {
	return c.Comm.Type != "none" ||
		len(c.CDFiles) > 0 ||
		len(c.BootCommand) > 0
}
