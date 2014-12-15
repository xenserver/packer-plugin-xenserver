package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepRemoveDevices struct{}

func (self *stepRemoveDevices) Run(state multistep.StateBag) multistep.StepAction {
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

	// Eject ISO from drive
	vbds, err := instance.GetVBDs()
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VBDs"))
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	for _, vbd := range vbds {
		rec, err := vbd.GetRecord()
		if err != nil {
			ui.Error(fmt.Sprintf("Could not get record for VBD"))
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		// Hack - should encapsulate this in the client really
		// This is needed because we can't guarentee the type
		// returned by the xmlrpc lib will be string
		if recType, ok := rec["type"].(string); ok {
			if recType == "CD" {
				ui.Say("Ejecting CD...")
				vbd.Eject()
			}
		} else {
			break
		}
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

func (self *stepRemoveDevices) Cleanup(state multistep.StateBag) {}
