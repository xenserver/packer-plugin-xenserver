package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
)

type stepAttachVdi struct {
	VdiUuidKey string
	VdiType    VDIType

	vdi *VDI
}

func (self *stepAttachVdi) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(XenAPIClient)

	var vdiUuid string
	if vdiUuidRaw, ok := state.GetOk(self.VdiUuidKey); ok {
		vdiUuid = vdiUuidRaw.(string)
	} else {
		log.Printf("Skipping attach of '%s'", self.VdiUuidKey)
		return multistep.ActionContinue
	}

	var err error
	self.vdi, err = client.GetVdiByUuid(vdiUuid)
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

	err = instance.ConnectVdi(self.vdi, self.VdiType)
	if err != nil {
		ui.Error(fmt.Sprintf("Error attaching VDI '%s': '%s'", vdiUuid, err.Error()))
		return multistep.ActionHalt
	}

	log.Printf("Attached VDI '%s'", vdiUuid)

	return multistep.ActionContinue
}

func (self *stepAttachVdi) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(config)
	client := state.Get("client").(XenAPIClient)
	if config.ShouldKeepInstance(state) {
		return
	}

	if self.vdi == nil {
		return
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		log.Printf("Unable to get VM from UUID '%s': %s", uuid, err.Error())
		return
	}

	vdiUuid := state.Get(self.VdiUuidKey).(string)

	err = instance.DisconnectVdi(self.vdi)
	if err != nil {
		log.Printf("Unable to disconnect VDI '%s': %s", vdiUuid, err.Error())
		return
	}
	log.Printf("Detached VDI '%s'", vdiUuid)
}
