//go:generate mapstructure-to-hcl2 -type Config
package xva

import (
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/template/interpolate"
	xscommon "github.com/xenserver/packer-builder-xenserver/builder/xenserver/common"
)

type Config struct {
	common.PackerConfig   `mapstructure:",squash"`
	xscommon.CommonConfig `mapstructure:",squash"`

	SourcePath string `mapstructure:"source_path"`
	VMMemory   uint   `mapstructure:"vm_memory"`

	PlatformArgs map[string]string `mapstructure:"platform_args"`

	ctx interpolate.Context
}
