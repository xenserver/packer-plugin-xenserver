package common

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepDetachVdi struct {
	VdiUuidKey string
}

func (self *StepDetachVdi) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)

	var vdiUuid string
	if vdiUuidRaw, ok := state.GetOk(self.VdiUuidKey); ok {
		vdiUuid = vdiUuidRaw.(string)
	} else {
		log.Printf("Skipping detach of '%s'", self.VdiUuidKey)
		return multistep.ActionContinue
	}

	vdi, err := c.client.VDI.GetByUUID(c.session, vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VDI from UUID '%s': %s", vdiUuid, err.Error()))
		return multistep.ActionHalt
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := c.client.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	err = DisconnectVdi(c, instance, vdi)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to detach VDI '%s': %s", vdiUuid, err.Error()))
		//return multistep.ActionHalt
		return multistep.ActionContinue
	}

	log.Printf("Detached VDI '%s'", vdiUuid)

	return multistep.ActionContinue
}

func (self *StepDetachVdi) Cleanup(state multistep.StateBag) {}
