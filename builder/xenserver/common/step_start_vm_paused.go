package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepStartVmPaused struct {
	VmCleanup
}

func (self *StepStartVmPaused) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(Config)

	ui.Say("Step: Start VM Paused")

	uuid := state.Get("instance_uuid").(string)
	instance, err := c.client.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	// note that here "cd" means boot from hard drive ('c') first, then CDROM ('d')
	err = c.client.VM.SetHVMBootPolicy(c.session, instance, "BIOS order")

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to set HVM boot params: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = c.client.VM.SetHVMBootParams(c.session, instance, map[string]string{"order": "cd", "firmware": config.Firmware})
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to set HVM boot params: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = c.client.VM.Start(c.session, instance, true, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to start VM with UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	domid, err := c.client.VM.GetDomid(c.session, instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get domid of VM with UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}
	state.Put("domid", domid)

	return multistep.ActionContinue
}
