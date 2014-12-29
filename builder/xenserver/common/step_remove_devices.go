package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepRemoveDevices struct{}

func (self *StepRemoveDevices) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(XenAPIClient)

	ui.Say("Step: Remove devices from VM")

	instance_id := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(instance_id)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM from UUID %s", instance_id))
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Destroy all connected VIFs
	vifs, err := instance.GetVIFs()
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VIFs"))
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	for _, vif := range vifs {
		ui.Message("Destroying VIF " + vif.Ref)
		vif.Destroy()
	}

	return multistep.ActionContinue
}

func (self *StepRemoveDevices) Cleanup(state multistep.StateBag) {}
