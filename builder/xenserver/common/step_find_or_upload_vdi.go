package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepFindOrUploadVdi struct {
	StepUploadVdi
}

func (self *StepFindOrUploadVdi) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	vdiName := self.VdiNameFunc()

	ui.Say(fmt.Sprintf("Attemping to find VDI '%s'", vdiName))

	vdis, err := c.client.VDI.GetByNameLabel(c.session, vdiName)
	if err != nil {
		ui.Error(fmt.Sprintf("Failed to find VDI '%s' by name label: %s", vdiName, err.Error()))
		return multistep.ActionHalt
	}

	if len(vdis) > 1 {
		ui.Error(fmt.Sprintf("Found more than one VDI with name '%s'. Name must be unique", vdiName))
		return multistep.ActionHalt
	} else if len(vdis) == 1 {

		vdi := vdis[0]

		vdiUuid, err := c.client.VDI.GetUUID(c.session, vdi)
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", vdiName, err.Error()))
			return multistep.ActionHalt
		}
		state.Put(self.VdiUuidKey, vdiUuid)
		return multistep.ActionContinue
	}
	return self.uploadVdi(ctx, state)
}
