package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepBootWait struct{}

func (self *stepBootWait) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(XenAPIClient)
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)

	instance, _ := client.GetVMByUuid(state.Get("instance_uuid").(string))
	ui.Say("Unpausing VM " + state.Get("instance_uuid").(string))
	instance.Unpause()

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

func (self *stepBootWait) Cleanup(state multistep.StateBag) {}
