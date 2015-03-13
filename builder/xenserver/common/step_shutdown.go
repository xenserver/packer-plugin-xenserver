package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	"time"
)

type StepShutdown struct{}

func (StepShutdown) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)
	instance_uuid := state.Get("instance_uuid").(string)

	instance, err := client.GetVMByUuid(instance_uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM with UUID '%s': %s", instance_uuid, err.Error()))
		return multistep.ActionHalt
	}

	ui.Say("Step: Shutting down VM")

	// Shutdown the VM
	success := func() bool {
		if config.ShutdownCommand != "" {
			ui.Message("Executing shutdown command...")

			_, err := ExecuteGuestSSHCmd(state, config.ShutdownCommand)
			if err != nil {
				ui.Error(fmt.Sprintf("Shutdown command failed: %s", err.Error()))
				return false
			}

			ui.Message("Waiting for VM to enter Halted state...")

			err = InterruptibleWait{
				Predicate: func() (bool, error) {
					power_state, err := instance.GetPowerState()
					return power_state == "Halted", err
				},
				PredicateInterval: 5 * time.Second,
				Timeout:           300 * time.Second,
			}.Wait(state)

			if err != nil {
				ui.Error(fmt.Sprintf("Error waiting for VM to halt: %s", err.Error()))
				return false
			}

		} else {
			ui.Message("Attempting to cleanly shutdown the VM...")

			err = instance.CleanShutdown()
			if err != nil {
				ui.Error(fmt.Sprintf("Could not shut down VM: %s", err.Error()))
				return false
			}

		}
		return true
	}()

	if !success {
		ui.Say("WARNING: Forcing hard shutdown of the VM...")
		err = instance.HardShutdown()
		if err != nil {
			ui.Error(fmt.Sprintf("Could not hard shut down VM -- giving up: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	ui.Message("Successfully shut down VM")
	return multistep.ActionContinue
}

func (StepShutdown) Cleanup(state multistep.StateBag) {}
