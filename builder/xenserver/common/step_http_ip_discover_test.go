// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func discoverState(t *testing.T, config CommonConfig) *multistep.BasicStateBag {
	t.Helper()
	state := new(multistep.BasicStateBag)
	state.Put("commonconfig", config)
	state.Put("ui", packersdk.TestUi(t))
	return state
}

func TestStepHTTPIPDiscover_PublishesTheAddressForLaterSteps(t *testing.T) {
	state := discoverState(t, CommonConfig{HTTPAddress: "192.168.7.7"})

	if action := (StepHTTPIPDiscover{}).Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run returned %v, want ActionContinue", action)
	}

	ip, ok := state.GetOk("http_ip")
	if !ok {
		t.Fatal("Run must publish http_ip into the state bag")
	}
	if ip != "192.168.7.7" {
		t.Errorf("http_ip = %v, want the configured bind address", ip)
	}
}

func TestStepHTTPIPDiscover_DerivesTheAddressFromThePoolRoute(t *testing.T) {
	state := discoverState(t, CommonConfig{HostIp: "127.0.0.1"})

	if action := (StepHTTPIPDiscover{}).Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run returned %v, want ActionContinue", action)
	}

	if ip, _ := state.GetOk("http_ip"); ip != "127.0.0.1" {
		t.Errorf("http_ip = %v, want the route-derived source address", ip)
	}
}

func TestStepHTTPIPDiscover_HaltsWhenTheAddressCannotBeFound(t *testing.T) {
	state := discoverState(t, CommonConfig{HTTPInterface: "definitely-not-a-nic"})

	if action := (StepHTTPIPDiscover{}).Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run returned %v, want ActionHalt", action)
	}

	if _, ok := state.GetOk("error"); !ok {
		t.Error("halting without recording an error leaves the build with no explanation")
	}
}
