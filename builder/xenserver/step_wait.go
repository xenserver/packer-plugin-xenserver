package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"time"
)

type stepWait struct{}

func (self *stepWait) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(XenAPIClient)

	ui.Say("Step: Wait for install to complete.")

	instance_id := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(instance_id)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM from UUID %s", instance_id))
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	//Expect install to be configured to shutdown on completion
	err = InterruptibleWait{
		Predicate: func() (bool, error) {
			ui.Say("Waiting for install to complete.")
			power_state, err := instance.GetPowerState()
			return power_state == "Halted", err
		},
		PredicateInterval: 30 * time.Second,
		Timeout:           config.InstallTimeout,
	}.Wait(state)

	if err != nil {
		ui.Error(err.Error())
		ui.Error("Giving up waiting for installation to complete.")
		return multistep.ActionHalt
	}

	ui.Say("Install has completed. Moving on.")

	return multistep.ActionContinue
}

func (self *stepWait) Cleanup(state multistep.StateBag) {}
