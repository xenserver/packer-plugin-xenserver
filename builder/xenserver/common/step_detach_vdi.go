package common

import (
	"fmt"
	"log"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/simonfuhrer/go-xenserver-client"
)

type StepDetachVdi struct {
	VdiUuidKey string
}

func (s *StepDetachVdi) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	var vdiUuid string
	if vdiUuidRaw, ok := state.GetOk(s.VdiUuidKey); ok {
		vdiUuid = vdiUuidRaw.(string)
	} else {
		log.Printf("Skipping detach of '%s'", s.VdiUuidKey)
		return multistep.ActionContinue
	}

	vdi, err := client.GetVdiByUuid(vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VDI from UUID '%s': %s", vdiUuid, err.Error()))
		return multistep.ActionHalt
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	err = instance.DisconnectVdi(vdi)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to detach VDI '%s': %s", vdiUuid, err.Error()))
		//return multistep.ActionHalt
		return multistep.ActionContinue
	}

	log.Printf("Detached VDI '%s'", vdiUuid)

	return multistep.ActionContinue
}

func (s *StepDetachVdi) Cleanup(state multistep.StateBag) {}
