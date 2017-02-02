package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type StepSetVmHostSshAddress struct{}

func (self *StepSetVmHostSshAddress) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(xsclient.XenAPIClient)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Set SSH address to VM host IP")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	host, err := instance.GetResidentOn()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM Host for VM '%s': %s", uuid, err.Error()))
	}

	address, err := host.GetAddress()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get address from VM Host: %s", err.Error()))
	}

	state.Put("ssh_address", address)
	ui.Say(fmt.Sprintf("Set host SSH address to '%s'.", address))

	return multistep.ActionContinue
}

func (self *StepSetVmHostSshAddress) Cleanup(state multistep.StateBag) {}
