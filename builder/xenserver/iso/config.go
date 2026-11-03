// Copyright (c) 2026. Citrix Systems, Inc.

//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type ISOConfig
package iso

import (
	"github.com/hashicorp/packer-plugin-sdk/common"
	commonsteps "github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

// genericTemplate has no sizing to inherit, so the build must supply its own.
const genericTemplate = "Other install media"

// minWindowsMemoryGB is Windows' documented installation minimum; below it, Setup stalls instead of failing.
const minWindowsMemoryGB = 2

// ISOConfig is the xenserver-vm-iso builder's configuration: the options shared
// with the clone builder, Packer's ISO download options, and the ones declared
// here.
type ISOConfig struct {
	common.PackerConfig   `mapstructure:",squash"`
	xscommon.CommonConfig `mapstructure:",squash"`
	commonsteps.ISOConfig `mapstructure:",squash"`

	// Name label of the ISO on `iso_sr`. Set this or `iso_url`/`iso_urls`, not
	// both.
	ISOName string `mapstructure:"iso_name"`

	// Deprecated: not supported on XenServer 9, which ships no guest-tools ISO.
	// Leave it unset.
	GuestToolsISO string `mapstructure:"guest_tools_iso"`

	// Name label of the XenServer template the VM is cloned from, such as
	// `Windows 11`. The VM takes the template's firmware, Secure Boot, vTPM and
	// sizing unless the build sets them. Name `Other install media` to define the
	// hardware yourself; `vcpus`, `vm_memory_gb` and `disk_size_gb` are then
	// required. It must name a template, not a VM or snapshot.
	CloneTemplate string `mapstructure:"clone_template" required:"true"`

	// Attach a virtual TPM. Requires `firmware = "uefi"`, set explicitly. Defaults
	// to false, which adds none but keeps any vTPM setup `clone_template` has.
	VTPMEnabled bool `mapstructure:"vtpm"`
}
