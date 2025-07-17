package common

import (
	"errors"
	"os"
	"fmt"
	"github.com/hashicorp/packer-plugin-sdk/communicator"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type SSHConfig struct {
	Comm                communicator.Config `mapstructure:",squash"`

	SSHKeyPath  		string `mapstructure:"ssh_key_path"`
	SSHHostPortMin    	uint `mapstructure:"ssh_host_port_min"`
	SSHHostPortMax    	uint `mapstructure:"ssh_host_port_max"`
	SSHSkipNatMapping 	bool `mapstructure:"ssh_skip_nat_mapping"`
}

func (c *SSHConfig) Prepare(ctx *interpolate.Context) []error {
	var errs []error

	errs = append(errs, c.Comm.Prepare(ctx)...)
	if c.SSHHostPortMin == 0 {
		c.SSHHostPortMin = 2222
	}

	if c.SSHHostPortMax == 0 {
		c.SSHHostPortMax = 4444
	}

	if c.SSHKeyPath != "" {
		if _, err := os.Stat(c.SSHKeyPath); err != nil {
			errs = append(errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		} else if _, err := FileSigner(c.SSHKeyPath); err != nil {
			errs = append(errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		}
		// TODO: backwards compatibility, write fixer instead
		c.Comm.SSHPrivateKeyFile = c.SSHKeyPath
	}

	if c.Comm.SSHUsername == "" {
		errs = append(errs, errors.New("An ssh_username must be specified."))
	}

	
	if c.SSHHostPortMin > c.SSHHostPortMax {
		errs = append(errs,
			errors.New("ssh_host_port_min must be less than ssh_host_port_max"))
	}

	return errs
}
