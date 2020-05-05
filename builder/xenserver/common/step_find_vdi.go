package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type StepFindVdi struct {
	VdiName            string
	VdiUuidKey         string
	PreviousResult     string
	PreviousStepAction multistep.StepAction
	ErrorFunc          func(errString string) multistep.StepAction
}

func (self *StepFindVdi) ErrorHandler(msg string) multistep.StepAction {
	self.PreviousResult = "FAILURE"
	self.PreviousStepAction = self.ErrorFunc(msg)
	return self.PreviousStepAction
}

func (self *StepFindVdi) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	// Ignore if VdiName is not specified
	if self.VdiName == "" {
		return multistep.ActionContinue
	}

	if self.PreviousResult != "" {
		return self.PreviousStepAction
	}

	self.PreviousResult = "SUCCESS"
	self.PreviousStepAction = multistep.ActionContinue

	if self.ErrorFunc == nil {
		self.ErrorFunc = func(errString string) multistep.StepAction {
			ui.Error(errString)
			return multistep.ActionHalt
		}
	}

	if len(self.VdiName) >= 7 && self.VdiName[:7] == "uuid://" {
		vdiUuid := self.VdiName[7:]
		_, err := client.GetVdiByUuid(vdiUuid)
		if err != nil {
			return self.ErrorHandler(fmt.Sprintf("Unable to get VDI from UUID '%s': %s", vdiUuid, err.Error()))
		}
		state.Put(self.VdiUuidKey, vdiUuid)
		return self.PreviousStepAction
	}

	vdis, err := client.GetVdiByNameLabel(self.VdiName)

	switch {
	case len(vdis) == 0:
		return self.ErrorHandler(fmt.Sprintf("Couldn't find a VDI named '%s'", self.VdiName))
	case len(vdis) > 1:
		return self.ErrorHandler(fmt.Sprintf("Found more than one VDI with name '%s'. Name must be unique", self.VdiName))
	}

	vdi := vdis[0]

	vdiUuid, err := vdi.GetUuid()
	if err != nil {
		return self.ErrorHandler(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", self.VdiName, err.Error()))
	}
	state.Put(self.VdiUuidKey, vdiUuid)
	return self.PreviousStepAction
}

func (self *StepFindVdi) Cleanup(state multistep.StateBag) {}
