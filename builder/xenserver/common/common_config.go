package common

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	"xenapi"
)

type CommonConfig struct {
	Username string `mapstructure:"remote_username"`
	Password string `mapstructure:"remote_password"`
	HostIp   string `mapstructure:"remote_host"`

	VMName             string   `mapstructure:"vm_name"`
	VMDescription      string   `mapstructure:"vm_description"`
	SrName             string   `mapstructure:"sr_name"`
	SrISOName          string   `mapstructure:"sr_iso_name"`
	FloppyFiles        []string `mapstructure:"floppy_files"`
	CDFiles            []string `mapstructure:"cd_files"`
	NetworkNames       []string `mapstructure:"network_names"`
	ExportNetworkNames []string `mapstructure:"export_network_names"`

	HostPortMin uint `mapstructure:"host_port_min"`
	HostPortMax uint `mapstructure:"host_port_max"`

	BootCommand     []string `mapstructure:"boot_command"`
	ShutdownCommand string   `mapstructure:"shutdown_command"`

	BootWait time.Duration `mapstructure:"boot_wait"`

	ToolsIsoName string `mapstructure:"tools_iso_name"`

	HTTPDir     string `mapstructure:"http_directory"`
	HTTPPortMin uint   `mapstructure:"http_port_min"`
	HTTPPortMax uint   `mapstructure:"http_port_max"`

	SSHConfig `mapstructure:",squash"`

	OutputDir string `mapstructure:"output_directory"`
	Format    string `mapstructure:"format"`
	KeepVM    string `mapstructure:"keep_vm"`
	IPGetter  string `mapstructure:"ip_getter"`

	SkipCertVerification bool   `mapstructure:"skip_cert_verification"`
	ServerCert           string `mapstructure:"server_cert"`
}

func (c *CommonConfig) Prepare(ctx *interpolate.Context, pc *common.PackerConfig) []error {
	var errs []error

	// Set default values
	if c.HostPortMin == 0 {
		c.HostPortMin = 5900
	}

	if c.HostPortMax == 0 {
		c.HostPortMax = 6000
	}

	if c.BootWait == 0 {
		c.BootWait = 5 * time.Second
	}

	if c.HTTPPortMin == 0 {
		c.HTTPPortMin = 8000
	}

	if c.HTTPPortMax == 0 {
		c.HTTPPortMax = 9000
	}

	if c.FloppyFiles == nil {
		c.FloppyFiles = make([]string, 0)
	}

	if c.CDFiles == nil {
		c.CDFiles = make([]string, 0)
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

	if !c.SkipCertVerification && c.ServerCert == "" {
		errs = append(errs, errors.New("server_cert must be specified when skip_cert_verification is false."))
	}

	for _, f := range c.CDFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			errs = append(errs, fmt.Errorf("cd_files: '%s' does not exist", f))
		}
	}

	if c.HostPortMin > c.HostPortMax {
		errs = append(errs, errors.New("the host min port must be less than the max"))
	}

	if c.HTTPPortMin > c.HTTPPortMax {
		errs = append(errs, errors.New("the HTTP min port must be less than the max"))
	}

	switch c.Format {
	case "xva", "xva_compressed", "vdi_raw", "vdi_vhd", "none":
	default:
		errs = append(errs, errors.New("format must be one of 'xva', 'vdi_raw', 'vdi_vhd', 'none'"))
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

func (config CommonConfig) GetSR(c *Connection) (xenapi.SRRef, error) {
	var srRef xenapi.SRRef
	if config.SrName == "" {
		hostRef, err := c.session.GetThisHost(c.ref)

		if err != nil {
			return srRef, err
		}

		pools, err := xenapi.Pool.GetAllRecords(c.session)

		if err != nil {
			return srRef, err
		}

		for _, pool := range pools {
			if pool.Master == hostRef {
				return pool.DefaultSR, nil
			}
		}

		return srRef, errors.New(fmt.Sprintf("failed to find default SR on host '%s'", hostRef))

	} else {
		// Use the provided name label to find the SR to use
		srs, err := xenapi.SR.GetByNameLabel(c.session, config.SrName)

		if err != nil {
			return srRef, err
		}

		switch {
		case len(srs) == 0:
			return srRef, fmt.Errorf("Couldn't find a SR with the specified name-label '%s'", config.SrName)
		case len(srs) > 1:
			return srRef, fmt.Errorf("Found more than one SR with the name '%s'. The name must be unique", config.SrName)
		}

		return srs[0], nil
	}
}

func (config CommonConfig) GetISOSR(c *Connection) (xenapi.SRRef, error) {
	var srRef xenapi.SRRef
	if config.SrISOName == "" {
		return srRef, errors.New("sr_iso_name must be specified in the packer configuration")

	} else {
		// Use the provided name label to find the SR to use
		srs, err := xenapi.SR.GetByNameLabel(c.session, config.SrName)

		if err != nil {
			return srRef, err
		}

		switch {
		case len(srs) == 0:
			return srRef, fmt.Errorf("Couldn't find a SR with the specified name-label '%s'", config.SrName)
		case len(srs) > 1:
			return srRef, fmt.Errorf("Found more than one SR with the name '%s'. The name must be unique", config.SrName)
		}

		return srs[0], nil
	}
}
