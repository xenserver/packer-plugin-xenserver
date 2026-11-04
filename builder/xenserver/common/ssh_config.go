// Copyright (c) 2026. Citrix Systems, Inc.

//go:generate packer-sdc struct-markdown
package common

import (
	"errors"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/communicator"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

// SSHConfig carries the communicator options. The communicator defaults to `none`; a
// build that connects to the guest sets `communicator = "winrm"`.
type SSHConfig struct {
	Comm communicator.Config `mapstructure:",squash"`
}

// Environment variables that Prepare falls back to for guest credentials the template leaves unset.
const (
	EnvSSHUsername       = "XENSERVER_SSH_USERNAME"
	EnvSSHPassword       = "XENSERVER_SSH_PASSWORD"
	EnvSSHPrivateKeyFile = "XENSERVER_SSH_PRIVATE_KEY_FILE"
	EnvWinRMUsername     = "XENSERVER_WINRM_USERNAME"
	EnvWinRMPassword     = "XENSERVER_WINRM_PASSWORD"
)

// ApplyCommunicatorDefault sets an unset communicator to "none"; call it before CommonConfig.Prepare.
func (c *SSHConfig) ApplyCommunicatorDefault() {
	if c.Comm.Type == "" {
		c.Comm.Type = "none"
	}
}

// applyGuestCredentialEnv runs for every communicator: the SDK only reads the credentials of the one in use.
func (c *SSHConfig) applyGuestCredentialEnv() {
	applyEnvFallbacks([]envFallback{
		{&c.Comm.SSHUsername, EnvSSHUsername},
		{&c.Comm.SSHPassword, EnvSSHPassword},
		{&c.Comm.SSHPrivateKeyFile, EnvSSHPrivateKeyFile},
		{&c.Comm.WinRMUser, EnvWinRMUsername},
		{&c.Comm.WinRMPassword, EnvWinRMPassword},
	})
}

func (c *SSHConfig) Prepare(ctx *interpolate.Context) []error {
	// Before c.Comm.Prepare, which would default to "ssh" and demand an ssh_username.
	c.ApplyCommunicatorDefault()

	// Also before c.Comm.Prepare, which checks the usernames and loads ssh_private_key_file.
	c.applyGuestCredentialEnv()

	// Packer only fills this filter from HCL sensitive values, so add environment passwords; empty ones are ignored.
	packersdk.LogSecretFilter.Set(c.Comm.SSHPassword, c.Comm.WinRMPassword)

	errs := c.Comm.Prepare(ctx)

	// The SDK also accepts the docker types, which would fail at StepConnect only after the install.
	switch c.Comm.Type {
	case "ssh", "winrm", "none":
	default:
		errs = append(errs, errors.New("communicator must be one of \"ssh\", \"winrm\" or \"none\", got \""+
			c.Comm.Type+"\""))
	}

	// The SDK checks winrm_username but not winrm_password, and no later step can supply the password.
	if c.Comm.Type == "winrm" && c.Comm.WinRMPassword == "" {
		errs = append(errs, fmt.Errorf("winrm_password must be specified when communicator = \"winrm\" "+
			"(or set %s): the WinRM communicator has no other way to authenticate, so the build would "+
			"install and then fail at the connect step.", EnvWinRMPassword))
	}

	return errs
}
