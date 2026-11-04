// Copyright (c) 2026. Citrix Systems, Inc.

package clone

import (
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

func TestConnectStepDialsTheConfiguredHost(t *testing.T) {
	var b Builder
	config := testConfig()
	config["communicator"] = "winrm"
	config["winrm_username"] = "Administrator"
	config["winrm_password"] = "packer"
	config["winrm_host"] = "10.0.0.9"

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	state := new(multistep.BasicStateBag)
	state.Put("instance_ssh_address", "192.168.1.5")

	host, err := b.connectStep().Host(state)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if host != "10.0.0.9" {
		t.Errorf("host = %q, want the configured winrm_host", host)
	}
}

func TestConnectStepFallsBackToTheDiscoveredIP(t *testing.T) {
	var b Builder
	config := testConfig()
	config["communicator"] = "ssh"
	config["ssh_username"] = "root"

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	state := new(multistep.BasicStateBag)
	state.Put("instance_ssh_address", "192.168.1.5")

	host, err := b.connectStep().Host(state)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if host != "192.168.1.5" {
		t.Errorf("host = %q, want the discovered IP", host)
	}
}

func TestRemovesVTPMOnlyOnAnExplicitFalse(t *testing.T) {
	unset := testConfig()

	off := testConfig()
	off["vtpm"] = false

	on := testConfig()
	on["vtpm"] = true
	on["firmware"] = "uefi"

	for name, tc := range map[string]struct {
		config map[string]interface{}
		want   bool
	}{
		"unset":          {unset, false},
		"explicit false": {off, true},
		"explicit true":  {on, false},
	} {
		var b Builder
		if _, _, err := b.Prepare(tc.config); err != nil {
			t.Errorf("%s: should not have error: %s", name, err)
			continue
		}
		if got := b.config.removesVTPM(); got != tc.want {
			t.Errorf("%s: removesVTPM() = %t, want %t", name, got, tc.want)
		}
	}
}

func TestConnectStepDoesNotOverrideTheConfiguredPort(t *testing.T) {
	var b Builder
	config := testConfig()
	config["communicator"] = "ssh"
	config["ssh_username"] = "root"
	config["ssh_port"] = 2222

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	step := b.connectStep()
	if step.SSHPort != nil {
		t.Error("a port callback is called on every attempt and its answer overwrites ssh_port")
	}
	if step.Config.SSHPort != 2222 {
		t.Errorf("ssh_port = %d, want the configured 2222", step.Config.SSHPort)
	}
}
