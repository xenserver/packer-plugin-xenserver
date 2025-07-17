package common

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepAttachvTPM struct {
	VTPM     xenapi.VTPMRef
	EnablevTPM bool}

func (self *StepAttachvTPM) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)

	if( self.EnablevTPM == false) {
		ui.Say("vTPM is not enabled, skipping")
		return multistep.ActionContinue
	}
	ui.Say("Step: Attaching vTPM")

	// Check if the VM already has a vTPM attached
	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}
	vTPMs, err := xenapi.VM.GetVTPMs(c.session, instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get vTPMs for VM '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}
	if len(vTPMs) == 0 {
		ui.Say("No existing vTPM found, creating a new one")
		vTPM, err := xenapi.VTPM.Create(c.session, instance, false)
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to create vTPM for VM '%s': %s", uuid, err.Error()))
			return multistep.ActionHalt
		}
		self.VTPM = vTPM
		ui.Say(fmt.Sprintf("Created new vTPM: %s", self.VTPM))
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

	// Destroy the vTPM if it was created
	if self.VTPM != "" {
		err := xenapi.VTPM.Destroy(c.session, self.VTPM)
		if err != nil {
			log.Printf("Unable to destroy vTPM: %s", err.Error())
			return
		}
		log.Printf("Destroyed vTPM '%s'", self.VTPM)
	}
}
