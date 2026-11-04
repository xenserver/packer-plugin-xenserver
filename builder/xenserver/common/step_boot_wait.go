// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepBootWait struct{}

func (self *StepBootWait) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	c := state.Get("client").(*Connection)
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	ui.Say("Unpausing VM " + uuid)
	if err := xenapi.VM.Unpause(c.session, instance); err != nil {
		ui.Error(fmt.Sprintf("Unable to unpause VM with UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	if int64(config.BootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", config.BootWait))
		err := InterruptibleWait{Timeout: config.BootWait}.Wait(state)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	return multistep.ActionContinue
}

func (self *StepBootWait) Cleanup(state multistep.StateBag) {}
