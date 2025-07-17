//go:generate packer-sdc mapstructure-to-hcl2 -type Config
package common

import (
	"time"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	CommonConfig        `mapstructure:",squash"`

	VCPUs       uint              `mapstructure:"vcpus"`
	VCPUsMax    uint             
	VCPUsAtStartup uint             
	CorePerSocket  uint              `mapstructure:"cores_per_socket"`
	VMMemory       uint              `mapstructure:"vm_memory"`
	DiskSize       uint              `mapstructure:"disk_size"`
	CloneTemplate  string            `mapstructure:"clone_template"`
	VMOtherConfig  map[string]string `mapstructure:"vm_other_config"`
	VGPUProfile    string            `mapstructure:"vgpu_profile"`

	ISOChecksum     string   `mapstructure:"iso_checksum"`
	ISOChecksumType string   `mapstructure:"iso_checksum_type"`
	ISOUrls         []string `mapstructure:"iso_urls"`
	ISOUrl          string   `mapstructure:"iso_url"`
	ISOName         string   `mapstructure:"iso_name"`

	PlatformArgs map[string]string `mapstructure:"platform_args"`

	InstallTimeout time.Duration        `mapstructure:"install_timeout"`
	SourcePath        string        `mapstructure:"source_path"`

	Firmware          string `mapstructure:"firmware"`
	SecureBoot        bool   `mapstructure:"secure_boot"`
	VTPMEnabled       bool   `mapstructure:"vTPM"`
	ConvertToTemplate bool   `mapstructure:"convert_to_template"`
	CreateSnapshot    bool   `mapstructure:"create_snapshot"`
	SnapshotName      string `mapstructure:"snapshot_name"`
	ctx               interpolate.Context
}

func (c Config) GetInterpContext() *interpolate.Context {
	return &c.ctx
}
