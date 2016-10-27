package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type StepConvertToTemplate struct{}

func (self *StepConvertToTemplate) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	if !config.ConvertToTemplate {
		return multistep.ActionContinue
	}

	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	ui.Say("Step: Convert to template")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	err = instance.SetIsATemplate(true)
	if err != nil {
		ui.Error(fmt.Sprintf("Error converting VM to a template: %s", err.Error()))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (self *StepConvertToTemplate) Cleanup(state multistep.StateBag) {}
