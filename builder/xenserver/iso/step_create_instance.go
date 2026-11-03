// Copyright (c) 2026. Citrix Systems, Inc.

package iso

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

type stepCreateInstance struct {
	instance *xenapi.VMRef
	vdi      *xenapi.VDIRef
}

// keptSizing leaves out the CPU topology, which ReconcileCPUTopology reports itself.
func keptSizing(config ISOConfig) []string {
	var kept []string

	if config.VCPUs == nil {
		kept = append(kept, "vCPU count")
	}
	if config.VMMemoryGB == nil {
		kept = append(kept, "memory")
	}

	return kept
}

func (self *stepCreateInstance) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*xscommon.Connection)
	config := state.Get("config").(ISOConfig)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Create Instance")

	template := state.Get("source_vm").(xenapi.VMRef)

	err := xscommon.PreCleanup(state, config.PackerForce,
		xscommon.TemplateSource(config.CloneTemplate), template)
	if err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	instance, err := xenapi.VM.Clone(c.GetSession(), template, config.Name())
	if err != nil {
		ui.Error(fmt.Sprintf("Error cloning '%s': %s", config.CloneTemplate, err.Error()))
		return multistep.ActionHalt
	}
	self.instance = &instance

	err = xenapi.VM.SetIsATemplate(c.GetSession(), instance, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting is_a_template=false: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Before the install: Windows offers the PV drivers through Windows Update only if it
	// saw the vendor device while installing.
	if _, err = xscommon.ApplyVendorDevice(c, instance, config.HasVendorDeviceEnabled()); err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Unset sizing options keep what VM.clone copied from clone_template.
	if kept := keptSizing(config); len(kept) > 0 {
		ui.Say(fmt.Sprintf("Keeping the %s of '%s'", strings.Join(kept, ", "), config.CloneTemplate))
	}

	if config.VCPUs != nil {
		// XAPI enforces VCPUs_at_startup <= VCPUs_max in both setters: drop to 1 first, or a
		// vcpus below the template's count is refused.
		err = xenapi.VM.SetVCPUsAtStartup(c.GetSession(), instance, 1)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=1: %s", err.Error()))
			return multistep.ActionHalt
		}

		err = xenapi.VM.SetVCPUsMax(c.GetSession(), instance, int(*config.VCPUs))
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM VCPUs Max=%d: %s", *config.VCPUs, err.Error()))
			return multistep.ActionHalt
		}

		err = xenapi.VM.SetVCPUsAtStartup(c.GetSession(), instance, int(*config.VCPUs))
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=%d: %s", *config.VCPUs, err.Error()))
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
	ui.Say(topology.Describe(config.CloneTemplate))

	secureBoot, err := xscommon.ReconcileSecureBoot(c, instance, config.SecureBoot)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	ui.Say(secureBoot.Describe(config.CloneTemplate))

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

	sr, err := config.GetSR(c)

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	diskBytes := state.Get("disk_size_bytes").(int)
	if config.DiskSizeGB == nil {
		ui.Say(fmt.Sprintf("Installing onto a %s disk, the size '%s' asks for",
			xscommon.FormatGB(diskBytes), config.CloneTemplate))
	}

	ui.Say(fmt.Sprintf("Using the following SR for the VM: %s", sr))

	vdi, err := xenapi.VDI.Create(c.GetSession(), xenapi.VDIRecord{
		NameLabel:   "Packer-disk",
		VirtualSize: diskBytes,
		Type:        "user",
		Sharable:    false,
		ReadOnly:    false,
		SR:          sr,
		OtherConfig: map[string]string{
			"temp": "temp",
		},
	})
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.vdi = &vdi

	err = xscommon.ConnectVdi(c, instance, vdi, xenapi.VbdTypeDisk)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	// VM.clone copies the template's VIFs, so network_names has to destroy them first or
	// VIF.create fails with DEVICE_ALREADY_EXISTS.
	if len(config.NetworkNames) == 0 {
		inheritedVIFs, err := xscommon.CountVIFs(c, instance)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		if inheritedVIFs > 0 {
			ui.Say(fmt.Sprintf("No network_names given; keeping the %d NIC(s) inherited from the template", inheritedVIFs))
			log.Printf("Template supplied %d VIF(s); not creating a management VIF\n", inheritedVIFs)
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

func (self *stepCreateInstance) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(ISOConfig)
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*xscommon.Connection)

	// VmCleanup has usually destroyed both already, so a failure here is an expected HANDLE_INVALID.
	if self.instance != nil {
		ui.Say("Destroying VM")
		_ = xenapi.VM.HardShutdown(c.GetSession(), *self.instance) // redundant, just in case
		if err := xenapi.VM.Destroy(c.GetSession(), *self.instance); err != nil {
			log.Printf("Unable to destroy the VM: %s", err.Error())
		}
	}

	if self.vdi != nil {
		ui.Say("Destroying VDI")
		if err := xenapi.VDI.Destroy(c.GetSession(), *self.vdi); err != nil {
			log.Printf("Unable to destroy the Packer disk: %s", err.Error())
		}
	}
}
