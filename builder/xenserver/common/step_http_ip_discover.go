// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// StepHTTPIPDiscover publishes "http_ip", the address behind `{{ .HTTPIP }}` and PackerHTTPIP.
type StepHTTPIPDiscover struct{}

func (StepHTTPIPDiscover) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	ip, err := newHTTPIPResolver().resolve(config.HTTPAddress, config.HTTPInterface, config.HostIp)
	if err != nil {
		err = fmt.Errorf("error determining the HTTP server address to advertise to the VM: %w", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("The VM will reach this machine at %s", ip))
	state.Put("http_ip", ip)

	return multistep.ActionContinue
}

func (StepHTTPIPDiscover) Cleanup(multistep.StateBag) {}
