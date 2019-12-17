package common

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/template/interpolate"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type CommonConfig struct {
	Username   string `mapstructure:"remote_username"`
	Password   string `mapstructure:"remote_password"`
	HostIp     string `mapstructure:"remote_host"`
	XenSSHPort uint   `mapstructure:"remote_ssh_port"`

	VMName             string   `mapstructure:"vm_name"`
	VMDescription      string   `mapstructure:"vm_description"`
	SrName             string   `mapstructure:"sr_name"`
	FloppyFiles        []string `mapstructure:"floppy_files"`
	NetworkNames       []string `mapstructure:"network_names"`
	ExportNetworkNames []string `mapstructure:"export_network_names"`

	HostPortMin uint `mapstructure:"host_port_min"`
	HostPortMax uint `mapstructure:"host_port_max"`

	PreBootHostScripts   []string `mapstructure:"pre_boot_host_scripts"`
	PreExportHostScripts []string `mapstructure:"pre_export_host_scripts"`
	BootCommand          []string `mapstructure:"boot_command"`
	ShutdownCommand      string   `mapstructure:"shutdown_command"`

	RawBootWait string `mapstructure:"boot_wait"`
	BootWait    time.Duration

	ToolsIsoName string `mapstructure:"tools_iso_name"`

	HTTPDir     string `mapstructure:"http_directory"`
	HTTPPortMin uint   `mapstructure:"http_port_min"`
	HTTPPortMax uint   `mapstructure:"http_port_max"`

	Communicator string `mapstructure:"communicator"`

	//	SSHHostPortMin    uint   `mapstructure:"ssh_host_port_min"`
	//	SSHHostPortMax    uint   `mapstructure:"ssh_host_port_max"`
	SSHKeyPath  string `mapstructure:"ssh_key_path"`
	SSHPassword string `mapstructure:"ssh_password"`
	SSHPort     uint   `mapstructure:"ssh_port"`
	SSHUser     string `mapstructure:"ssh_username"`
	SSHConfig   `mapstructure:",squash"`

	RawSSHWaitTimeout string `mapstructure:"ssh_wait_timeout"`
	SSHWaitTimeout    time.Duration

	ConvertToTemplate bool `mapstructure:"convert_to_template"`
	DestroyVIFs       bool `mapstructure:"destroy_vifs"`
	DiscDrives        int  `mapstructure:"disc_drives"`

	OutputDir string `mapstructure:"output_directory"`
	Format    string `mapstructure:"format"`
	KeepVM    string `mapstructure:"keep_vm"`
	IPGetter  string `mapstructure:"ip_getter"`
}

func (c *CommonConfig) Prepare(ctx *interpolate.Context, pc *common.PackerConfig) []error {
	var err error
	var errs []error

	// Set default values

	if c.XenSSHPort == 0 {
		c.XenSSHPort = 22
	}

	if c.HostPortMin == 0 {
		c.HostPortMin = 5900
	}

	if c.HostPortMax == 0 {
		c.HostPortMax = 6000
	}

	if c.RawBootWait == "" {
		c.RawBootWait = "5s"
	}

	if c.ToolsIsoName == "" {
		c.ToolsIsoName = "xs-tools.iso"
	}

	if c.HTTPPortMin == 0 {
		c.HTTPPortMin = 8000
	}

	if c.HTTPPortMax == 0 {
		c.HTTPPortMax = 9000
	}

	if c.RawSSHWaitTimeout == "" {
		c.RawSSHWaitTimeout = "200m"
	}

	if c.FloppyFiles == nil {
		c.FloppyFiles = make([]string, 0)
	}

	/*
		if c.SSHHostPortMin == 0 {
			c.SSHHostPortMin = 2222
		}

		if c.SSHHostPortMax == 0 {
			c.SSHHostPortMax = 4444
		}
	*/

	if c.SSHPort == 0 {
		c.SSHPort = 22
	}

	if c.RawSSHWaitTimeout == "" {
		c.RawSSHWaitTimeout = "20m"
	}

	if c.OutputDir == "" {
		c.OutputDir = fmt.Sprintf("output-%s", pc.PackerBuildName)
	}

	if c.VMName == "" {
		c.VMName = fmt.Sprintf("packer-%s-{{timestamp}}", pc.PackerBuildName)
	}

	if c.Format == "" {
		c.Format = "xva"
	}

	if c.KeepVM == "" {
		c.KeepVM = "never"
	}

	if c.IPGetter == "" {
		c.IPGetter = "auto"
	}

	// Validation

	if c.Username == "" {
		errs = append(errs, errors.New("remote_username must be specified."))
	}

	if c.Password == "" {
		errs = append(errs, errors.New("remote_password must be specified."))
	}

	if c.HostIp == "" {
		errs = append(errs, errors.New("remote_host must be specified."))
	}

	if c.HostPortMin > c.HostPortMax {
		errs = append(errs, errors.New("the host min port must be less than the max"))
	}

	if c.HTTPPortMin > c.HTTPPortMax {
		errs = append(errs, errors.New("the HTTP min port must be less than the max"))
	}

	c.BootWait, err = time.ParseDuration(c.RawBootWait)
	if err != nil {
		errs = append(errs, fmt.Errorf("Failed to parse boot_wait: %s", err))
	}

	if c.SSHKeyPath != "" {
		if _, err := os.Stat(c.SSHKeyPath); err != nil {
			errs = append(errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		} else if _, err := commonssh.FileSigner(c.SSHKeyPath); err != nil {
			errs = append(errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		}
	}

	/*
		if c.SSHHostPortMin > c.SSHHostPortMax {
			errs = append(errs,
				errors.New("ssh_host_port_min must be less than ssh_host_port_max"))
		}
	*/

	if c.SSHUser == "" && c.Communicator != "winrm" {
		errs = append(errs, errors.New("An ssh_username must be specified."))
	}

	c.SSHWaitTimeout, err = time.ParseDuration(c.RawSSHWaitTimeout)
	if err != nil {
		errs = append(errs, fmt.Errorf("Failed to parse ssh_wait_timeout: %s", err))
	}

	if c.DiscDrives < 0 {
		errs = append(errs, errors.New("disc_drives greater than or equal to 0."))
	}

	switch c.Format {
	case "xva", "xva_compressed", "vdi_raw", "vdi_vhd", "none":
	default:
		errs = append(errs, errors.New("format must be one of 'xva', 'xva_compressed', 'vdi_raw', 'vdi_vhd', 'none'"))
	}

	switch c.KeepVM {
	case "always", "never", "on_success":
	default:
		errs = append(errs, errors.New("keep_vm must be one of 'always', 'never', 'on_success'"))
	}

	switch c.IPGetter {
	case "auto", "tools", "http":
	default:
		errs = append(errs, errors.New("ip_getter must be one of 'auto', 'tools', 'http'"))
	}

	return errs
}

// steps should check config.ShouldKeepVM first before cleaning up the VM
func (c CommonConfig) ShouldKeepVM(state multistep.StateBag) bool {
	switch c.KeepVM {
	case "always":
		return true
	case "never":
		return false
	case "on_success":
		// only keep instance if build was successful
		_, cancelled := state.GetOk(multistep.StateCancelled)
		_, halted := state.GetOk(multistep.StateHalted)
		return !(cancelled || halted)
	default:
		panic(fmt.Sprintf("Unknown keep_vm value '%s'", c.KeepVM))
	}
}

func (config CommonConfig) GetSR(client xsclient.XenAPIClient) (*xsclient.SR, error) {
	if config.SrName == "" {
		// Find the default SR
		return client.GetDefaultSR()

	} else {
		// Use the provided name label to find the SR to use
		srs, err := client.GetSRByNameLabel(config.SrName)

		if err != nil {
			return nil, err
		}

		switch {
		case len(srs) == 0:
			return nil, fmt.Errorf("Couldn't find a SR with the specified name-label '%s'", config.SrName)
		case len(srs) > 1:
			return nil, fmt.Errorf("Found more than one SR with the name '%s'. The name must be unique", config.SrName)
		}

		return srs[0], nil
	}
}
