// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

// StepAttachvTPM adds a vTPM when asked and the clone_template has none; Run, not the caller, sets VTPM.
type StepAttachvTPM struct {
	VTPM       xenapi.VTPMRef
	EnablevTPM bool
}

func (self *StepAttachvTPM) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)

	if !self.EnablevTPM {
		ui.Say("vTPM is not enabled, skipping")
		return multistep.ActionContinue
	}
	ui.Say("Step: Attaching vTPM")

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	outcome, err := ReconcileVTPM(c, instance, true)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to attach a vTPM to VM '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}
	if outcome != VTPMCreated {
		ui.Say("VM already has a vTPM, leaving it alone")
		return multistep.ActionContinue
	}
	ui.Say("Created a new vTPM")

	// Recorded only when this step created it, so Cleanup never destroys an inherited one.
	vTPMs, err := xenapi.VM.GetVTPMs(c.session, instance)
	if err != nil {
		log.Printf("Unable to read back the new vTPM of VM '%s': %s", uuid, err.Error())
	} else if len(vTPMs) == 1 {
		self.VTPM = vTPMs[0]
	}

	return multistep.ActionContinue
}

func (self *StepAttachvTPM) Cleanup(state multistep.StateBag) {
	config := state.Get("commonconfig").(CommonConfig)
	c := state.Get("client").(*Connection)
	if config.ShouldKeepVM(state) {
		return
	}

	if self.VTPM == "" {
		return
	}

	err := xenapi.VTPM.Destroy(c.session, self.VTPM)
	if err != nil {
		log.Printf("Unable to destroy vTPM: %s", err.Error())
		return
	}
	log.Printf("Destroyed vTPM '%s'", self.VTPM)
}
