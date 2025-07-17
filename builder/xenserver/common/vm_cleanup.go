package common

import (
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	
	"xenapi"
)

type VmCleanup struct{}

func (self *VmCleanup) Cleanup(state multistep.StateBag) {
	config := state.Get("commonconfig").(CommonConfig)
	c := state.Get("client").(*Connection)

	if config.ShouldKeepVM(state) {
		return
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return
	}

	err = xenapi.VM.HardShutdown(c.session, instance)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to force shutdown VM '%s': %s", uuid, err.Error()))
	}
}

func PreCleanup(state multistep.StateBag, force bool) error {
	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("commonconfig").(CommonConfig)

	// Let's find existing VMs with the same name
	vms, err := xenapi.VM.GetByNameLabel(c.session, config.VMName)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VM from Name '%s': %s", config.VMName, err.Error()))
		return err
	}

	if force && len(vms) == 1 {
		// We found a VM and the force flag is set, so remove it
		ui.Say(fmt.Sprintf("The VM / Template %s already exists, but deleting it due to -force flag", config.VMName))

		vmstate, err := xenapi.VM.GetPowerState(c.session, vms[0])
		if err != nil {
			return fmt.Errorf("Error getting powerstate of VM %s: %v", config.VMName, err)
		}
		if vmstate == xenapi.VMPowerStateHalted || vmstate == xenapi.VMPowerStateRunning {
			// Shutdown the VM
			err = xenapi.VM.Shutdown(c.session, vms[0])
			if err != nil {
				return fmt.Errorf("Error shutting down %s: %v", config.VMName, err)
			}
		}

		// Destroy the VM
		err = xenapi.VM.Destroy(c.session, vms[0])
		if err != nil {
			return fmt.Errorf("Error destroying %s: %v", config.VMName, err)
		}
	}
	if !force && len(vms) > 0 {
		return fmt.Errorf("%s already exists, you can use -force flag to destroy it: %v", config.VMName, err)
	}

	return nil
}
