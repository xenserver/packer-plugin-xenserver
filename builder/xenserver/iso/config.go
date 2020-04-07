//go:generate mapstructure-to-hcl2 -type Config
package iso

import (
	"time"

	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/template/interpolate"
	xscommon "github.com/xenserver/packer-builder-xenserver/builder/xenserver/common"
)

type Config struct {
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
