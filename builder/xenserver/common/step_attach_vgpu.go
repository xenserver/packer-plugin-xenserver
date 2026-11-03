// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepAttachvGPU struct {
	VGPUName string
}

func (self *StepAttachvGPU) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)

	if len(self.VGPUName) == 0 {
		ui.Say("No vGPU requested, skipping")
		return multistep.ActionContinue
	}
	ui.Say("Step: Find and Attach vGPU: " + self.VGPUName)

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.GetSession(), uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	if _, err := AttachVGPU(c, instance, self.VGPUName, ui.Say); err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (self *StepAttachvGPU) Cleanup(state multistep.StateBag) {}
