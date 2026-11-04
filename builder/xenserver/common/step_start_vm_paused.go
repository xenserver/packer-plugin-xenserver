// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepStartVmPaused struct {
	VmCleanup
}

func (self *StepStartVmPaused) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("commonconfig").(CommonConfig)

	ui.Say("Step: Start VM Paused")

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	err = xenapi.VM.SetHVMBootPolicy(c.session, instance, "BIOS order")

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to set HVM boot policy: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Set one key at a time: SetHVMBootParams replaces the whole map, inherited firmware included.
	// "cd" boots the hard drive ('c') first, then the CDROM ('d').
	if err = setHVMBootParam(c.session, instance, "order", "cd"); err != nil {
		ui.Error(fmt.Sprintf("Unable to set HVM boot param 'order': %s", err.Error()))
		return multistep.ActionHalt
	}

	// Run the check even with firmware unset: the VM then boots with the source's firmware.
	firmware, err := EnsureVMCanStart(c, instance, config.Firmware)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// StepTypeBootCommand reads this to decide whether to sweep the UEFI CD prompt.
	state.Put("firmware", firmware)

	switch {
	case config.Firmware != "":
		if err = setHVMBootParam(c.session, instance, "firmware", config.Firmware); err != nil {
			ui.Error(fmt.Sprintf("Unable to set HVM boot param 'firmware': %s", err.Error()))
			return multistep.ActionHalt
		}

	case firmware != "":
		// VM.clone already copied the source's firmware; it decides whether the install media can boot.
		ui.Say(fmt.Sprintf("Keeping the source's firmware (%s)", firmware))
	}

	err = xenapi.VM.Start(c.session, instance, true, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to start VM with UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}
