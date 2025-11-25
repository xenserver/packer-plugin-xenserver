package common

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	xenapi "github.com/terra-farm/go-xen-api-client"
)

type StepShutdown struct{}

func (StepShutdown) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	instance_uuid := state.Get("instance_uuid").(string)

	instance, err := c.client.VM.GetByUUID(c.session, instance_uuid)
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
					power_state, err := c.client.VM.GetPowerState(c.session, instance)
					return power_state == xenapi.VMPowerStateHalted, err
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

			err = c.client.VM.CleanShutdown(c.session, instance)
			if err != nil {
				ui.Error(fmt.Sprintf("Could not shut down VM: %s", err.Error()))
				return false
			}

		}
		return true
	}()

	if !success {
		ui.Say("WARNING: Forcing hard shutdown of the VM...")
		err = c.client.VM.HardShutdown(c.session, instance)
		if err != nil {
			ui.Error(fmt.Sprintf("Could not hard shut down VM -- giving up: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	ui.Message("Successfully shut down VM")
	return multistep.ActionContinue
}

func (StepShutdown) Cleanup(state multistep.StateBag) {}
