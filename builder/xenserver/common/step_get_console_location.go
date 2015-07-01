package common

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/simonfuhrer/go-xenserver-client"
)

type StepGetConsoleLocation struct{}

func (s *StepGetConsoleLocation) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)
	ui.Say("Step: get instances VNC Location via XAPI")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	consoles, err := instance.GetConsoles()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VNC consoles (is the VM running?): %s", err.Error()))
		return multistep.ActionHalt
	}
	var consoleLocation string
	for _, console := range consoles {
		consoleRec, err := console.GetRecord()
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to get VNC consoles (is the VM running?): %s", err.Error()))
			return multistep.ActionHalt
		}
		if consoleRec["protocol"].(string) == "rfb" {
			consoleLocation = consoleRec["location"].(string)
			break
		}

	}
	if consoleLocation == "" {
		ui.Error("Unable to get VNC consoles (is the VM running?")
		return multistep.ActionHalt
	}
	ui.Say(fmt.Sprintf("Store instance_console_location: %s", consoleLocation))
	state.Put("instance_console_location", consoleLocation)

	return multistep.ActionContinue
}

func (s *StepGetConsoleLocation) Cleanup(state multistep.StateBag) {
}

func InstanceConsoleLocation(state multistep.StateBag) (string, error) {
	consoleLocation := state.Get("instance_console_location").(string)
	return consoleLocation, nil
}

func InstanceVNCIP(state multistep.StateBag) (string, error) {
	// The port is in Dom0, so we want to forward from localhost
	return "127.0.0.1", nil
}
