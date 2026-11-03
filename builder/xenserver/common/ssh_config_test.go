// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

var guestCredentialEnvVars = []string{
	EnvSSHUsername, EnvSSHPassword, EnvSSHPrivateKeyFile, EnvWinRMUsername, EnvWinRMPassword,
}

func prepareSSH(t *testing.T, c *SSHConfig) []error {
	t.Helper()
	return prepareSSHWithEnv(t, c, nil)
}

func prepareSSHWithEnv(t *testing.T, c *SSHConfig, env map[string]string) []error {
	t.Helper()
	for _, key := range guestCredentialEnvVars {
		t.Setenv(key, env[key])
	}
	return c.Prepare(&interpolate.Context{})
}

func writePrivateKey(t *testing.T) string {
	t.Helper()

	_, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("generating a test key: %v", err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatalf("marshalling a test key: %v", err)
	}

	path := filepath.Join(t.TempDir(), "id_ed25519")
	if err := os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der}), 0o600); err != nil {
		t.Fatalf("writing a test key: %v", err)
	}
	return path
}

func TestSSHConfigPrepare_DoesNotRequireAnSSHUsernameWithoutACommunicator(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "none"

	if errs := prepareSSH(t, c); len(errs) > 0 {
		t.Fatalf("communicator = \"none\" must not require an ssh_username: %v", errs)
	}
}

func TestSSHConfigPrepare_SSHRequiresAnSSHUsernameExactlyOnce(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "ssh"

	errs := prepareSSH(t, c)

	var found int
	for _, err := range errs {
		if strings.Contains(err.Error(), "ssh_username") {
			found++
		}
	}
	if found != 1 {
		t.Errorf("communicator = \"ssh\" should report a missing ssh_username once, got %d of %v", found, errs)
	}
}

func TestSSHConfigPrepare_WinRMNeedsNoSSHUsername(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"
	c.Comm.WinRMUser = "Administrator"
	c.Comm.WinRMPassword = "secret"

	errs := prepareSSH(t, c)

	if len(errs) > 0 {
		t.Fatalf("a WinRM template carrying winrm credentials should be accepted: %v", errs)
	}
	if c.Comm.WinRMPort != 5985 {
		t.Errorf("winrm_port = %d, want the SDK default 5985", c.Comm.WinRMPort)
	}
}

func TestSSHConfigPrepare_WinRMStillRequiresAWinRMUsername(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"

	errs := prepareSSH(t, c)

	var found bool
	for _, err := range errs {
		if strings.Contains(err.Error(), "winrm_username") {
			found = true
		}
	}
	if !found {
		t.Errorf("communicator = \"winrm\" must be told to set winrm_username, got %v", errs)
	}
}

func TestSSHConfigPrepare_WinRMRequiresAWinRMPassword(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"
	c.Comm.WinRMUser = "Administrator"

	errs := prepareSSH(t, c)

	var found bool
	for _, err := range errs {
		if strings.Contains(err.Error(), "winrm_password") {
			found = true
		}
	}
	if !found {
		t.Errorf("communicator = \"winrm\" must be told to set winrm_password, got %v", errs)
	}
}

func prepareComm(t *testing.T, commType string) string {
	t.Helper()
	c := &SSHConfig{}
	c.Comm.Type = commType
	_ = prepareSSH(t, c)
	return c.Comm.Type
}

func TestSSHConfigPrepare_UnsetCommunicatorDefaultsToNone(t *testing.T) {
	if got := prepareComm(t, ""); got != "none" {
		t.Errorf("unset communicator = %q, want %q", got, "none")
	}
}

func TestSSHConfigPrepare_UnsetCommunicatorNeedsNoCredentials(t *testing.T) {
	c := &SSHConfig{}
	if errs := prepareSSH(t, c); len(errs) != 0 {
		t.Errorf("a template with no communicator should be accepted, got %v", errs)
	}
}

func TestSSHConfigPrepare_AnExplicitCommunicatorIsKept(t *testing.T) {
	for _, want := range []string{"ssh", "winrm", "none"} {
		t.Run(want, func(t *testing.T) {
			if got := prepareComm(t, want); got != want {
				t.Errorf("communicator = %q survived Prepare as %q", want, got)
			}
		})
	}
}

func TestSSHConfigPrepare_SSHCredentialsFallBackToTheEnvironment(t *testing.T) {
	keyPath := writePrivateKey(t)

	c := &SSHConfig{}
	c.Comm.Type = "ssh"

	errs := prepareSSHWithEnv(t, c, map[string]string{
		EnvSSHUsername:       "builder",
		EnvSSHPassword:       "from-the-environment",
		EnvSSHPrivateKeyFile: keyPath,
	})

	if len(errs) > 0 {
		t.Fatalf("an ssh template whose credentials are all in the environment should be accepted: %v", errs)
	}
	if c.Comm.SSHUsername != "builder" {
		t.Errorf("SSHUsername = %q, want it from %s", c.Comm.SSHUsername, EnvSSHUsername)
	}
	if c.Comm.SSHPassword != "from-the-environment" {
		t.Errorf("SSHPassword = %q, want it from %s", c.Comm.SSHPassword, EnvSSHPassword)
	}
	if c.Comm.SSHPrivateKeyFile != keyPath {
		t.Errorf("SSHPrivateKeyFile = %q, want it from %s", c.Comm.SSHPrivateKeyFile, EnvSSHPrivateKeyFile)
	}
}

func TestSSHConfigPrepare_WinRMCredentialsFallBackToTheEnvironment(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"

	errs := prepareSSHWithEnv(t, c, map[string]string{
		EnvWinRMUsername: "Administrator",
		EnvWinRMPassword: "from-the-environment",
	})

	if len(errs) > 0 {
		t.Fatalf("a winrm template whose credentials are all in the environment should be accepted: %v", errs)
	}
	if c.Comm.WinRMUser != "Administrator" {
		t.Errorf("WinRMUser = %q, want it from %s", c.Comm.WinRMUser, EnvWinRMUsername)
	}
	if c.Comm.WinRMPassword != "from-the-environment" {
		t.Errorf("WinRMPassword = %q, want it from %s", c.Comm.WinRMPassword, EnvWinRMPassword)
	}
}

func TestSSHConfigPrepare_TemplateGuestCredentialsWinOverTheEnvironment(t *testing.T) {
	keyPath := writePrivateKey(t)

	c := &SSHConfig{}
	c.Comm.Type = "ssh"
	c.Comm.SSHUsername = "from-the-template"
	c.Comm.SSHPassword = "template-password"
	c.Comm.SSHPrivateKeyFile = keyPath

	errs := prepareSSHWithEnv(t, c, map[string]string{
		EnvSSHUsername:       "wrong",
		EnvSSHPassword:       "wrong",
		EnvSSHPrivateKeyFile: filepath.Join(t.TempDir(), "does-not-exist"),
	})

	if len(errs) > 0 {
		t.Fatalf("the template's own credentials should be accepted: %v", errs)
	}
	if c.Comm.SSHUsername != "from-the-template" {
		t.Errorf("SSHUsername = %q, want the template's", c.Comm.SSHUsername)
	}
	if c.Comm.SSHPassword != "template-password" {
		t.Errorf("SSHPassword = %q, want the template's", c.Comm.SSHPassword)
	}
	if c.Comm.SSHPrivateKeyFile != keyPath {
		t.Errorf("SSHPrivateKeyFile = %q, want the template's", c.Comm.SSHPrivateKeyFile)
	}
}

func TestSSHConfigPrepare_TemplateWinRMCredentialsWinOverTheEnvironment(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"
	c.Comm.WinRMUser = "from-the-template"
	c.Comm.WinRMPassword = "template-password"

	errs := prepareSSHWithEnv(t, c, map[string]string{
		EnvWinRMUsername: "wrong",
		EnvWinRMPassword: "wrong",
	})

	if len(errs) > 0 {
		t.Fatalf("the template's own credentials should be accepted: %v", errs)
	}
	if c.Comm.WinRMUser != "from-the-template" {
		t.Errorf("WinRMUser = %q, want the template's", c.Comm.WinRMUser)
	}
	if c.Comm.WinRMPassword != "template-password" {
		t.Errorf("WinRMPassword = %q, want the template's", c.Comm.WinRMPassword)
	}
}

func TestSSHConfigPrepare_EnvironmentCredentialsSatisfyTheRequiredChecks(t *testing.T) {
	t.Run("ssh_username", func(t *testing.T) {
		c := &SSHConfig{}
		c.Comm.Type = "ssh"

		errs := prepareSSHWithEnv(t, c, map[string]string{EnvSSHUsername: "builder"})

		if errorsContain(errs, "ssh_username") {
			t.Errorf("%s should satisfy the ssh_username check, got %v", EnvSSHUsername, errs)
		}
	})

	t.Run("winrm_username", func(t *testing.T) {
		c := &SSHConfig{}
		c.Comm.Type = "winrm"

		errs := prepareSSHWithEnv(t, c, map[string]string{
			EnvWinRMUsername: "Administrator",
			EnvWinRMPassword: "secret",
		})

		if errorsContain(errs, "winrm_username") {
			t.Errorf("%s should satisfy the winrm_username check, got %v", EnvWinRMUsername, errs)
		}
	})

	t.Run("winrm_password", func(t *testing.T) {
		c := &SSHConfig{}
		c.Comm.Type = "winrm"
		c.Comm.WinRMUser = "Administrator"

		errs := prepareSSHWithEnv(t, c, map[string]string{EnvWinRMPassword: "secret"})

		if errorsContain(errs, "winrm_password") {
			t.Errorf("%s should satisfy the winrm_password check, got %v", EnvWinRMPassword, errs)
		}
	})
}

func TestSSHConfigPrepare_MissingWinRMPasswordNamesTheEnvironmentVariable(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"
	c.Comm.WinRMUser = "Administrator"

	errs := prepareSSH(t, c)

	if !errorsContain(errs, EnvWinRMPassword) {
		t.Errorf("the winrm_password error should also name %s, got %v", EnvWinRMPassword, errs)
	}
}

func TestSSHConfigPrepare_SSHEnvironmentDoesNotAffectAWinRMBuild(t *testing.T) {
	c := &SSHConfig{}
	c.Comm.Type = "winrm"
	c.Comm.WinRMUser = "Administrator"
	c.Comm.WinRMPassword = "secret"

	errs := prepareSSHWithEnv(t, c, map[string]string{
		EnvSSHUsername: "builder",
		EnvSSHPassword: "ssh-password",
	})

	if len(errs) > 0 {
		t.Fatalf("an exported SSH credential should not disturb a winrm build: %v", errs)
	}
	if c.Comm.WinRMUser != "Administrator" || c.Comm.WinRMPassword != "secret" {
		t.Errorf("the WinRM credentials changed: user %q, password %q", c.Comm.WinRMUser, c.Comm.WinRMPassword)
	}
}
