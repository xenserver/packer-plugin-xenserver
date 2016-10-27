package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type StepFindVdi struct {
	VdiName       string
	ImagePathFunc func() string
	VdiUuidKey    string
}

func (self *StepFindVdi) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	// Ignore if VdiName is not specified
	if self.VdiName == "" {
		return multistep.ActionContinue
	} else if len(self.VdiName) >= 7 && self.VdiName[:7] == "uuid://" {
		vdiUuid := self.VdiName[7:]
		_, err := client.GetVdiByUuid(vdiUuid)
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to get VDI from UUID '%s': %s", vdiUuid, err.Error()))
			return multistep.ActionHalt
		}
		state.Put(self.VdiUuidKey, vdiUuid)
		return multistep.ActionContinue
	}

	vdis, err := client.GetVdiByNameLabel(self.VdiName)

	switch {
	case len(vdis) == 0:
		ui.Error(fmt.Sprintf("Couldn't find a VDI named '%s'", self.VdiName))
		return multistep.ActionHalt
	case len(vdis) > 1:
		ui.Error(fmt.Sprintf("Found more than one VDI with name '%s'. Name must be unique", self.VdiName))
		return multistep.ActionHalt
	}

	vdi := vdis[0]

	vdiUuid, err := vdi.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}
	state.Put(self.VdiUuidKey, vdiUuid)

	return multistep.ActionContinue
}

func (self *StepFindVdi) Cleanup(state multistep.StateBag) {}
