// Copyright (c) 2026. Citrix Systems, Inc.

package clone

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"xenapi"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

type stepCloneInstance struct {
	instance *xenapi.VMRef
	disks    []xscommon.ClonedDisk
}

func (self *stepCloneInstance) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*xscommon.Connection)
	config := state.Get("config").(CloneConfig)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Create Instance")

	src := config.cloneSource()
	source := state.Get("source_vm").(xenapi.VMRef)

	err := xscommon.PreCleanup(state, config.PackerForce, src, source)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	// VM.clone keeps each disk on its current SR; only VM.copy takes a target SR, and it
	// copies every disk in full, so it is used only when the disks are not on vm_sr already.
	var targetSR xenapi.SRRef
	needsCopy := false
	if config.VMSR != "" {
		targetSR, err = config.GetSR(c)
		if err != nil {
			ui.Error(fmt.Sprintf("Could not resolve vm_sr: %s", err.Error()))
			state.Put("error", err)
			return multistep.ActionHalt
		}
		needsCopy, err = xscommon.TemplateNeedsCopy(c, source, targetSR)
		if err != nil {
			ui.Error(err.Error())
			state.Put("error", err)
			return multistep.ActionHalt
		}
	}

	var instance xenapi.VMRef
	if needsCopy {
		ui.Say(fmt.Sprintf("Copying '%s' onto SR '%s'. Not all of its disks are there already, and the only "+
			"call that places them copies every one of them in full rather than cloning it copy-on-write, "+
			"so this takes as long as the data takes to write. Leave vm_sr unset to clone in place instead; "+
			"on a build with cd_files that also sends the CD to the pool's default SR.",
			src.Name, config.VMSR))
		instance, err = xenapi.VM.Copy(c.GetSession(), source, config.Name(), targetSR)
	} else {
		instance, err = xenapi.VM.Clone(c.GetSession(), source, config.Name())
	}
	if err != nil {
		// XAPI refuses a copy where a clone would succeed (an ISO target SR, a VDI with
		// on_boot=reset), and its error names neither vm_sr nor the copy.
		hint := ""
		if needsCopy {
			hint = " (this build copies rather than clones because vm_sr names an SR the source is not on)"
		}
		ui.Error(fmt.Sprintf("Error creating the VM from '%s'%s: %s", src.Name, hint, err.Error()))
		state.Put("error", err)
		return multistep.ActionHalt
	}
	self.instance = &instance

	// Only for Cleanup, so a failed read is not worth failing an otherwise good clone over.
	self.disks, err = xscommon.ClonedDisks(c, instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not list the cloned VM's disks, so cleanup may leave them behind: %s", err.Error()))
	}

	err = xenapi.VM.SetIsATemplate(c.GetSession(), instance, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting is_a_template=false: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Only an explicit value is written: changing the flag under an installed Windows can stop it booting.
	if config.HasVendorDevice != nil {
		vendorDeviceChanged, err := xscommon.ApplyVendorDevice(c, instance, *config.HasVendorDevice)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		if vendorDeviceChanged {
			ui.Say(fmt.Sprintf("Set has_vendor_device=%t on the clone; '%s' had it the other way round. "+
				"If that source already has Windows installed, this is likely to stop it booting: "+
				"remove has_vendor_device to keep the source's setting.",
				*config.HasVendorDevice, src.Name))
		}
	}

	// VM.clone copies the source's vTPM, so both true and false must act here, while the VM is halted.
	if config.VTPM != nil {
		outcome, err := xscommon.ReconcileVTPM(c, instance, config.VTPMEnabled())
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		switch outcome {
		case xscommon.VTPMCreated:
			ui.Say(fmt.Sprintf("Attached a vTPM: '%s' had none", src.Name))
		case xscommon.VTPMDestroyed:
			ui.Say(fmt.Sprintf("Destroyed the vTPM inherited from '%s'. Anything sealed to it, such as a "+
				"BitLocker key, is now unrecoverable in this build's VM. The source is untouched. "+
				"Remove vtpm = false to keep the inherited vTPM.", src.Name))
		case xscommon.VTPMDeclarationCleared:
			ui.Say(fmt.Sprintf("Cleared platform:vtpm, which '%s' declares: the clone had no vTPM to remove, "+
				"but XenServer would have created one from that key at start.", src.Name))
		}
	}

	if config.VCPUs != nil {
		vcpus := *config.VCPUs
		// XAPI enforces VCPUs_at_startup <= VCPUs_max in both setters: drop to 1 first, or a
		// vcpus below the source's count is refused.
		err = xenapi.VM.SetVCPUsAtStartup(c.GetSession(), instance, 1)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=1: %s", err.Error()))
			return multistep.ActionHalt
		}

		err = xenapi.VM.SetVCPUsMax(c.GetSession(), instance, int(vcpus))
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM VCPUs Max=%d: %s", vcpus, err.Error()))
			return multistep.ActionHalt
		}

		err = xenapi.VM.SetVCPUsAtStartup(c.GetSession(), instance, int(vcpus))
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=%d: %s", vcpus, err.Error()))
			return multistep.ActionHalt
		}
	}

	if config.VMMemoryGB != nil {
		memory := int(*config.VMMemoryGB) * xscommon.BytesPerGB
		err = xenapi.VM.SetMemoryLimits(c.GetSession(), instance, memory, memory, memory, memory)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM memory=%d: %s", memory, err.Error()))
			return multistep.ActionHalt
		}
	}

	// After the vCPU count: the topology is only valid against the final VCPUs_max.
	topology, err := xscommon.ReconcileCPUTopology(c, instance, config.VCPUs, config.CorePerSocket)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	ui.Say(topology.Describe(src.Name))

	secureBoot, err := xscommon.ReconcileSecureBoot(c, instance, config.SecureBoot)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	ui.Say(secureBoot.Describe(src.Name))

	if config.VMDescription != "" {
		err = xenapi.VM.SetNameDescription(c.GetSession(), instance, config.VMDescription)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	// Unconditional: the VM must be identifiable as plugin-built even without vm_other_config.
	err = xscommon.StampVM(c, instance, config.VMOtherConfig)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// VDI.resize needs the VM halted, as it still is here.
	if config.DiskSizeGB != nil {
		outcome, err := xscommon.SizePrimaryDisk(c, instance, int(*config.DiskSizeGB)*xscommon.BytesPerGB)
		if err != nil {
			ui.Error(fmt.Sprintf("%s (the source is '%s')", err.Error(), src.Name))
			return multistep.ActionHalt
		}
		if outcome == xscommon.DiskGrown {
			ui.Say(fmt.Sprintf("Grew the primary disk to %d GB", *config.DiskSizeGB))
		}
	}

	// A clone inherits the source's VIFs, so network_names has to destroy them first or
	// VIF.create fails with DEVICE_ALREADY_EXISTS.
	if len(config.NetworkNames) == 0 {
		inheritedVIFs, err := xscommon.CountVIFs(c, instance)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		if inheritedVIFs > 0 {
			ui.Say(fmt.Sprintf("No network_names given; keeping the %d NIC(s) inherited from the source", inheritedVIFs))
			log.Printf("Source supplied %d VIF(s); not creating a management VIF\n", inheritedVIFs)
		} else {
			log.Println("No network name given, attempting to use management interface")
			if err := xscommon.AttachManagementNetwork(c, instance, ui.Say); err != nil {
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}

	} else {
		if err := xscommon.DetachAllVIFs(c, instance, ui.Say); err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		log.Printf("Using provided network names: %v\n", config.NetworkNames)
		if err := xscommon.AttachNamedNetworks(c, instance, config.NetworkNames, ui.Say); err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	instanceId, err := xenapi.VM.GetUUID(c.GetSession(), instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}

	state.Put("instance_uuid", instanceId)
	ui.Say(fmt.Sprintf("Created instance '%s'", instanceId))

	return multistep.ActionContinue
}

// Cleanup destroys the cloned VM and its disks, which VM.destroy leaves behind.
func (self *stepCloneInstance) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(CloneConfig)
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*xscommon.Connection)

	if self.instance != nil {
		ui.Say("Destroying VM")
		_ = xenapi.VM.HardShutdown(c.GetSession(), *self.instance) // redundant, just in case
		if err := xenapi.VM.Destroy(c.GetSession(), *self.instance); err != nil {
			log.Printf("Unable to destroy the cloned VM: %s", err.Error())
		}
	}

	// Whether or not VM.destroy succeeded: the disks are left behind either way.
	if len(self.disks) > 0 {
		ui.Say(fmt.Sprintf("Destroying %d disk(s)...", len(self.disks)))
		for _, disk := range self.disks {
			if err := xenapi.VDI.Destroy(c.GetSession(), disk.VDI); err != nil {
				log.Printf("Unable to destroy VDI: %s", err.Error())
			}
		}
	}
}
