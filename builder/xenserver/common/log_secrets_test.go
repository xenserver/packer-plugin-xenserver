// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"strings"
	"testing"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func redacts(secret string) bool {
	return !strings.Contains(packersdk.LogSecretFilter.FilterString("before "+secret+" after"), secret)
}

func TestCommonConfigPrepare_RedactsAPasswordTakenFromTheEnvironment(t *testing.T) {
	const password = "pool-password-only-this-test-uses"

	c := minimalConfig("xva")
	c.Password = ""

	if _, errs := prepareWithEnv(t, c, map[string]string{EnvPassword: password}); len(errs) > 0 {
		t.Fatalf("Prepare: %v", errs)
	}

	if !redacts(password) {
		t.Errorf("%s reaches packer.log unredacted", EnvPassword)
	}
}

func TestSSHConfigPrepare_RedactsGuestPasswordsTakenFromTheEnvironment(t *testing.T) {
	for _, tc := range []struct {
		communicator string
		variable     string
		password     string
	}{
		{"ssh", EnvSSHPassword, "ssh-password-only-this-test-uses"},
		{"winrm", EnvWinRMPassword, "winrm-password-only-this-test-uses"},
	} {
		t.Run(tc.communicator, func(t *testing.T) {
			c := &SSHConfig{}
			c.Comm.Type = tc.communicator
			c.Comm.SSHUsername = "packer"
			c.Comm.WinRMUser = "Administrator"

			if errs := prepareSSHWithEnv(t, c, map[string]string{tc.variable: tc.password}); len(errs) > 0 {
				t.Fatalf("Prepare: %v", errs)
			}

			if !redacts(tc.password) {
				t.Errorf("%s reaches packer.log unredacted", tc.variable)
			}
		})
	}
}
