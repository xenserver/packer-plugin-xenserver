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
	ui := state.Get("ui").(packer.Ui)

	if config.ShouldKeepVM(state) {
		return
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return
	}

	// Get all VBDs (Virtual Block Devices) attached to the VM
	vbds, err := xenapi.VM.GetVBDs(c.session, instance)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VBDs for VM '%s': %s", uuid, err.Error()))
		vbds = []xenapi.VBDRef{}
	}

	// Collect all VDIs before destroying the VM
	var vdis []xenapi.VDIRef
	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(c.session, vbd)
		if err != nil {
			log.Printf(fmt.Sprintf("Unable to get VBD record: %s", err.Error()))
			continue
		}
		// Skip empty VBDs (like CD drives)
		if rec.VDI != "" {
			vdis = append(vdis, rec.VDI)
		}
	}

	ui.Say(fmt.Sprintf("Found %d disk(s) attached to VM", len(vdis)))

	// Shutdown the VM if it's running
	vmstate, err := xenapi.VM.GetPowerState(c.session, instance)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VM power state '%s': %s", uuid, err.Error()))
	}

	if vmstate == xenapi.VMPowerStateRunning {
		ui.Say(fmt.Sprintf("Shutting down VM on cleanup: %s", uuid))
		err = xenapi.VM.Shutdown(c.session, instance)
		if err != nil {
			log.Printf(fmt.Sprintf("Unable to shutdown VM '%s': %s", uuid, err.Error()))
			// Try hard shutdown if normal shutdown fails
			err = xenapi.VM.HardShutdown(c.session, instance)
			if err != nil {
				log.Printf(fmt.Sprintf("Unable to hard shutdown VM '%s': %s", uuid, err.Error()))
			}
		}
	}

	// Destroy the VM
	ui.Say(fmt.Sprintf("Destroying VM on cleanup: %s", uuid))
	err = xenapi.VM.Destroy(c.session, instance)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to destroy VM '%s': %s", uuid, err.Error()))
	}

	// Destroy all VDIs (disks) attached to the VM
	// Do this after VM destruction to ensure disks are detached
	ui.Say(fmt.Sprintf("Destroying %d disk(s)...", len(vdis)))
	for i, vdi := range vdis {
		vdiUuid, err := xenapi.VDI.GetUUID(c.session, vdi)
		if err != nil {
			log.Printf(fmt.Sprintf("Unable to get VDI UUID: %s", err.Error()))
			continue
		}

		// Unplug the VBD first if it still exists
		if i < len(vbds) {
			_ = xenapi.VBD.Unplug(c.session, vbds[i])
		}

		ui.Say(fmt.Sprintf("Destroying VDI (disk): %s", vdiUuid))

		// Retry destroying the VDI up to 3 times
		var lastErr error
		for attempt := 1; attempt <= 3; attempt++ {
			err = xenapi.VDI.Destroy(c.session, vdi)
			if err == nil {
				ui.Say(fmt.Sprintf("Successfully destroyed VDI: %s", vdiUuid))
				lastErr = nil
				break
			}
			lastErr = err
			if attempt < 3 {
				log.Printf(fmt.Sprintf("Attempt %d to destroy VDI '%s' failed: %s. Retrying...", attempt, vdiUuid, err.Error()))
			}
		}
		if lastErr != nil {
			log.Printf(fmt.Sprintf("Unable to destroy VDI '%s' after 3 attempts: %s", vdiUuid, lastErr.Error()))
		}
	}
	ui.Say("VM and disk cleanup completed")
}

func PreCleanup(state multistep.StateBag, force bool) error {
	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("commonconfig").(CommonConfig)

	ui.Say("Step: PreCleanup - Checking for existing VMs with name '" + config.VMName + "'")

	// Let's find existing VMs with the same name
	vms, err := xenapi.VM.GetByNameLabel(c.session, config.VMName)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VM from Name '%s': %s", config.VMName, err.Error()))
		return err
	}

	ui.Say(fmt.Sprintf("PreCleanup: Found %d existing VM(s)", len(vms)))

	if force && len(vms) > 0 {
		// We found one or more VMs and the force flag is set, so remove them all (but not templates)
		ui.Say(fmt.Sprintf("Found %d VM(s) / Template(s) named %s. Deleting due to -force flag", len(vms), config.VMName))

		for _, vm := range vms {
			// Check if this is a template
			isTemplate, err := xenapi.VM.GetIsATemplate(c.session, vm)
			if err != nil {
				return fmt.Errorf("Error checking if %s is a template: %v", config.VMName, err)
			}

			if isTemplate {
				ui.Say(fmt.Sprintf("Skipping template %s (only deleting VMs, not templates)", config.VMName))
				continue
			}

			// Get all VBDs (Virtual Block Devices) attached to the VM
			vbds, err := xenapi.VM.GetVBDs(c.session, vm)
			if err != nil {
				log.Printf(fmt.Sprintf("Unable to get VBDs for VM '%s': %s", config.VMName, err.Error()))
				vbds = []xenapi.VBDRef{}
			}

			// Collect all VDIs before destroying the VM
			var vdis []xenapi.VDIRef
			for _, vbd := range vbds {
				rec, err := xenapi.VBD.GetRecord(c.session, vbd)
				if err != nil {
					log.Printf(fmt.Sprintf("Unable to get VBD record: %s", err.Error()))
					continue
				}
				// Skip empty VBDs (like CD drives)
				if rec.VDI != "" {
					vdis = append(vdis, rec.VDI)
				}
			}

			vmstate, err := xenapi.VM.GetPowerState(c.session, vm)
			if err != nil {
				return fmt.Errorf("Error getting powerstate of VM %s: %v", config.VMName, err)
			}
			ui.Say(fmt.Sprintf("VM %s is in state: %v", config.VMName, vmstate))

			if vmstate == xenapi.VMPowerStateRunning {
				// Shutdown the VM only if it's running
				ui.Say(fmt.Sprintf("Shutting down running VM: %s", config.VMName))
				err = xenapi.VM.Shutdown(c.session, vm)
				if err != nil {
					return fmt.Errorf("Error shutting down %s: %v", config.VMName, err)
				}
				ui.Say(fmt.Sprintf("Successfully shut down VM: %s", config.VMName))
			}

			// Destroy the VM
			ui.Say(fmt.Sprintf("Destroying VM: %s", config.VMName))
			err = xenapi.VM.Destroy(c.session, vm)
			if err != nil {
				return fmt.Errorf("Error destroying %s: %v", config.VMName, err)
			}
			ui.Say(fmt.Sprintf("Successfully destroyed VM: %s", config.VMName))

			// Destroy all VDIs (disks) attached to the VM
			ui.Say(fmt.Sprintf("Destroying %d disk(s) for VM %s", len(vdis), config.VMName))
			for i, vdi := range vdis {
				vdiUuid, err := xenapi.VDI.GetUUID(c.session, vdi)
				if err != nil {
					log.Printf(fmt.Sprintf("Unable to get VDI UUID: %s", err.Error()))
					continue
				}

				// Unplug the VBD first if it still exists
				if i < len(vbds) {
					_ = xenapi.VBD.Unplug(c.session, vbds[i])
				}

				ui.Say(fmt.Sprintf("Destroying VDI (disk): %s", vdiUuid))

				// Retry destroying the VDI up to 3 times
				var lastErr error
				for attempt := 1; attempt <= 3; attempt++ {
					err = xenapi.VDI.Destroy(c.session, vdi)
					if err == nil {
						ui.Say(fmt.Sprintf("Successfully destroyed VDI: %s", vdiUuid))
						lastErr = nil
						break
					}
					lastErr = err
					if attempt < 3 {
						log.Printf(fmt.Sprintf("Attempt %d to destroy VDI '%s' failed: %s. Retrying...", attempt, vdiUuid, err.Error()))
					}
				}
				if lastErr != nil {
					log.Printf(fmt.Sprintf("Unable to destroy VDI '%s' after 3 attempts: %s", vdiUuid, lastErr.Error()))
				}
			}
		}
	}
	if !force && len(vms) > 0 {
		return fmt.Errorf("%s already exists, you can use -force flag to destroy it: %v", config.VMName, err)
	}

	return nil
}
