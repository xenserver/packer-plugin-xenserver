package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	"log"
)

type StepDetachVdi struct {
	VdiUuidKey string
}

func (self *StepDetachVdi) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	client := state.Get("client").(xsclient.XenAPIClient)

	var vdiUuid string
	if vdiUuidRaw, ok := state.GetOk(self.VdiUuidKey); ok {
		vdiUuid = vdiUuidRaw.(string)
	} else {
		log.Printf("Skipping detach of '%s'", self.VdiUuidKey)
		return multistep.ActionContinue
	}

	vdi, err := client.GetVdiByUuid(vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VDI from UUID '%s': %s", vdiUuid, err.Error()))
		return multistep.ActionHalt
	}

	if config.KeepDiskDrive && self.VdiUuidKey == "tools_vdi_uuid" {
		ui.Say(fmt.Sprintf("Ejecting '%s' and keeping DVD Drive", self.VdiUuidKey))
		vbds, err := vdi.GetVBDs()
		if err != nil {
			ui.Error(fmt.Sprintf("Error getting VBDs: %s", err.Error()))
			return multistep.ActionHalt
		}
		for _, vbd := range vbds {
			err = vbd.Eject()
			if err != nil {
				ui.Error(fmt.Sprintf("Error ejecting VBD: %s", err.Error()))
				return multistep.ActionHalt
			}
		}
		return multistep.ActionContinue
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

func (self *StepDetachVdi) Cleanup(state multistep.StateBag) {}
