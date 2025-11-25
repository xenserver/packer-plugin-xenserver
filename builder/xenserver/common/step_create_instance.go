package common

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	xenapi "github.com/terra-farm/go-xen-api-client"
	xsclient "github.com/terra-farm/go-xen-api-client"
)

type StepCreateInstance struct {
	// The XVA builder assumes it will boot an instance with an OS installed on its disks
	// while the ISO builder needs packer to create a disk for an OS to be installed on.
	AssumePreInstalledOS bool

	instance *xsclient.VMRef
	vdi      *xsclient.VDIRef
}

func (self *StepCreateInstance) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*Connection)
	config := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Create Instance")

	// Get the template to clone from

	vms, err := c.GetClient().VM.GetByNameLabel(c.GetSessionRef(), config.CloneTemplate)

	switch {
	case len(vms) == 0:
		ui.Error(fmt.Sprintf("Couldn't find a template with the name-label '%s'. Aborting.", config.CloneTemplate))
		return multistep.ActionHalt
	case len(vms) > 1:
		ui.Error(fmt.Sprintf("Found more than one template with the name '%s'. The name must be unique. Aborting.", config.CloneTemplate))
		return multistep.ActionHalt
	}

	template := vms[0]

	// Clone that VM template
	instance, err := c.GetClient().VM.Clone(c.GetSessionRef(), template, config.VMName)
	if err != nil {
		ui.Error(fmt.Sprintf("Error cloning VM: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.instance = &instance

	err = c.GetClient().VM.SetIsATemplate(c.GetSessionRef(), instance, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting is_a_template=false: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = c.GetClient().VM.SetVCPUsMax(c.GetSessionRef(), instance, int(config.VCPUsMax))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM VCPUs Max=%d: %s", config.VCPUsMax, err.Error()))
		return multistep.ActionHalt
	}

	err = c.GetClient().VM.SetVCPUsAtStartup(c.GetSessionRef(), instance, int(config.VCPUsAtStartup))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=%d: %s", config.VCPUsAtStartup, err.Error()))
		return multistep.ActionHalt
	}

	memory := int(config.VMMemory * 1024 * 1024)
	err = c.GetClient().VM.SetMemoryLimits(c.GetSessionRef(), instance, memory, memory, memory, memory)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM memory=%d: %s", memory, err.Error()))
		return multistep.ActionHalt
	}

	err = c.GetClient().VM.SetPlatform(c.GetSessionRef(), instance, config.PlatformArgs)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM platform: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = c.GetClient().VM.SetNameDescription(c.GetSessionRef(), instance, config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	if len(config.VMOtherConfig) != 0 {
		vm_other_config, err := c.GetClient().VM.GetOtherConfig(c.GetSessionRef(), instance)
		if err != nil {
			ui.Error(fmt.Sprintf("Error getting VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}
		for key, value := range config.VMOtherConfig {
			vm_other_config[key] = value
		}
		err = c.GetClient().VM.SetOtherConfig(c.GetSessionRef(), instance, vm_other_config)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	if !self.AssumePreInstalledOS {
		err = c.GetClient().VM.RemoveFromOtherConfig(c.GetSessionRef(), instance, "disks")
		if err != nil {
			ui.Error(fmt.Sprintf("Error removing disks from VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}

		// Create VDI for the instance
		sr, err := config.GetSR(c)

		if err != nil {
			ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
			return multistep.ActionHalt
		}

		ui.Say(fmt.Sprintf("Using the following SR for the VM: %s", sr))

		vdi, err := c.GetClient().VDI.Create(c.GetSessionRef(), xenapi.VDIRecord{
			NameLabel:   config.DiskName,
			VirtualSize: int(config.DiskSize * 1024 * 1024),
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

		err = ConnectVdi(c, instance, vdi, xsclient.VbdTypeDisk)
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	// Connect Network

	var network xsclient.NetworkRef

	if len(config.NetworkNames) == 0 {
		// No network has be specified. Use the management interface
		log.Println("No network name given, attempting to use management interface")
		pifs, err := c.GetClient().PIF.GetAll(c.GetSessionRef())

		if err != nil {
			ui.Error(fmt.Sprintf("Error getting PIFs: %s", err.Error()))
			return multistep.ActionHalt
		}

		for _, pif := range pifs {
			pif_rec, err := c.GetClient().PIF.GetRecord(c.GetSessionRef(), pif)

			if err != nil {
				ui.Error(fmt.Sprintf("Error getting PIF record: %s", err.Error()))
				return multistep.ActionHalt
			}

			if pif_rec.Management {
				network = pif_rec.Network
			}

		}

		if string(network) == "" {
			ui.Error("Error: couldn't find management network. Aborting.")
			return multistep.ActionHalt
		}

		log.Printf("Creating VIF on network '%s' on VM '%s'\n", network, instance)
		_, err = ConnectNetwork(c, network, instance, "0")

		if err != nil {
			ui.Error(fmt.Sprintf("Failed to create VIF with error: %v", err))
			return multistep.ActionHalt
		}

	} else {
		log.Printf("Using provided network names: %v\n", config.NetworkNames)
		// Look up each network by it's name label
		for i, networkNameLabel := range config.NetworkNames {
			networks, err := c.GetClient().Network.GetByNameLabel(c.GetSessionRef(), networkNameLabel)

			if err != nil {
				ui.Error(fmt.Sprintf("Error occured getting Network by name-label: %s", err.Error()))
				return multistep.ActionHalt
			}

			switch {
			case len(networks) == 0:
				ui.Error(fmt.Sprintf("Couldn't find a network with the specified name-label '%s'. Aborting.", networkNameLabel))
				return multistep.ActionHalt
			case len(networks) > 1:
				ui.Error(fmt.Sprintf("Found more than one network with the name '%s'. The name must be unique. Aborting.", networkNameLabel))
				return multistep.ActionHalt
			}

			//we need the VIF index string
			vifIndexString := fmt.Sprintf("%d", i)
			_, err = ConnectNetwork(c, networks[0], instance, vifIndexString)

			if err != nil {
				ui.Say(fmt.Sprintf("Failed to connect VIF with error: %v", err.Error()))
			}
		}
	}

	err = AddVMTags(c, instance, config.VMTags)
	if err != nil {
		ui.Error(fmt.Sprintf("Failed to add tags: %s", err.Error()))
		return multistep.ActionHalt
	}

	instanceId, err := c.GetClient().VM.GetUUID(c.GetSessionRef(), instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}

	state.Put("instance_uuid", instanceId)
	ui.Say(fmt.Sprintf("Created instance '%s'", instanceId))

	return multistep.ActionContinue
}

func (self *StepCreateInstance) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(Config)
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)

	if self.instance != nil {
		ui.Say("Destroying VM")
		_ = c.GetClient().VM.HardShutdown(c.GetSessionRef(), *self.instance) // redundant, just in case
		err := c.GetClient().VM.Destroy(c.GetSessionRef(), *self.instance)
		if err != nil {
			ui.Error(err.Error())
		}
	}

	if self.vdi != nil {
		ui.Say("Destroying VDI")
		err := c.GetClient().VDI.Destroy(c.GetSessionRef(), *self.vdi)
		if err != nil {
			ui.Error(err.Error())
		}
	}
}
