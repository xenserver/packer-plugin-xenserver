package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepBootWait struct{}

func (self *StepBootWait) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	c := state.Get("client").(*Connection)
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	instance, _ := c.client.VM.GetByUUID(c.session, state.Get("instance_uuid").(string))
	ui.Say("Unpausing VM " + state.Get("instance_uuid").(string))
	Unpause(c, instance)

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
