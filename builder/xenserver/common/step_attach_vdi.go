package common

import (
	"fmt"
	"log"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/simonfuhrer/go-xenserver-client"
)

type StepAttachVdi struct {
	VdiUuidKey string
	VdiType    xsclient.VDIType

	vdi *xsclient.VDI
}

func (s *StepAttachVdi) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	var vdiUuid string
	if vdiUuidRaw, ok := state.GetOk(s.VdiUuidKey); ok {
		vdiUuid = vdiUuidRaw.(string)
	} else {
		log.Printf("Skipping attach of '%s'", s.VdiUuidKey)
		return multistep.ActionContinue
	}

	var err error
	s.vdi, err = client.GetVdiByUuid(vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VDI from UUID '%s': %s", vdiUuid, err.Error()))
		return multistep.ActionHalt
	}
	log.Printf("VDIUUIDKEY: %s", s.VdiUuidKey)
	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	err = instance.ConnectVdi(s.vdi, s.VdiType, "")
	if err != nil {
		ui.Error(fmt.Sprintf("Error attaching VDI '%s': '%s'", vdiUuid, err.Error()))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *StepAttachVdi) Cleanup(state multistep.StateBag) {
	config := state.Get("commonconfig").(CommonConfig)
	client := state.Get("client").(xsclient.XenAPIClient)
	if config.ShouldKeepVM(state) {
		return
	}

	if s.vdi == nil {
		return
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		log.Printf("Unable to get VM from UUID '%s': %s", uuid, err.Error())
		return
	}

	vdiUuid := state.Get(s.VdiUuidKey).(string)

	err = instance.DisconnectVdi(s.vdi)
	if err != nil {
		log.Printf("Unable to disconnect VDI '%s': %s", vdiUuid, err.Error())
		return
	}
	log.Printf("Detached VDI '%s'", vdiUuid)
}
