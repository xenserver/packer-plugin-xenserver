package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepStartVmPaused struct{}

func (self *stepStartVmPaused) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(XenAPIClient)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Start VM Paused")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	err = instance.Start(true, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to start VM with UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	domid, err := instance.GetDomainId()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get domid of VM with UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}
	state.Put("domid", domid)

	return multistep.ActionContinue
}

func (self *stepStartVmPaused) Cleanup(state multistep.StateBag) {
}
