package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepCreateInstance struct {
	instance *VM
	vdi      *VDI
}

func (self *stepCreateInstance) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(XenAPIClient)
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Create Instance")

	// Get the template to clone from

	vms, err := client.GetVMByNameLabel(config.CloneTemplate)

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
	instance, err := template.Clone(config.InstanceName)
	if err != nil {
		ui.Error(fmt.Sprintf("Error cloning VM: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.instance = instance

	err = instance.SetIsATemplate(false)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting is_a_template=false: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = instance.SetStaticMemoryRange(config.InstanceMemory, config.InstanceMemory)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM memory=%s: %s", config.InstanceMemory, err.Error()))
		return multistep.ActionHalt
	}

	instance.SetPlatform(config.PlatformArgs)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM platform: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Create VDI for the instance
	var sr *SR

	if config.SrName == "" {
		// Find the default SR
		default_sr, err := client.GetDefaultSR()
		sr = default_sr

		if err != nil {
			ui.Error(fmt.Sprintf("Error getting default SR: %s", err.Error()))
			return multistep.ActionHalt
		}

	} else {
		// Use the provided name label to find the SR to use
		srs, err := client.GetSRByNameLabel(config.SrName)

		if err != nil {
			ui.Error(fmt.Sprintf("Error getting default SR: %s", err.Error()))
			return multistep.ActionHalt
		}

		switch {
		case len(srs) == 0:
			ui.Error(fmt.Sprintf("Couldn't find a SR with the specified name-label '%s'. Aborting.", config.SrName))
			return multistep.ActionHalt
		case len(srs) > 1:
			ui.Error(fmt.Sprintf("Found more than one SR with the name '%s'. The name must be unique. Aborting.", config.SrName))
			return multistep.ActionHalt
		}

		sr = srs[0]
	}

	vdi, err := sr.CreateVdi("Packer-disk", config.RootDiskSize)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.vdi = vdi

	err = instance.ConnectVdi(vdi, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Connect Network

	var network *Network

	if config.NetworkName == "" {
		// No network has be specified. Use the management interface
		network = new(Network)
		network.Ref = ""
		network.Client = &client

		pifs, err := client.GetPIFs()

		if err != nil {
			ui.Error(fmt.Sprintf("Error getting PIFs: %s", err.Error()))
			return multistep.ActionHalt
		}

		for _, pif := range pifs {
			pif_rec, err := pif.GetRecord()

			if err != nil {
				ui.Error(fmt.Sprintf("Error getting PIF record: %s", err.Error()))
				return multistep.ActionHalt
			}

			if pif_rec["management"].(bool) {
				network.Ref = pif_rec["network"].(string)
			}

		}

		if network.Ref == "" {
			ui.Error("Error: couldn't find management network. Aborting.")
			return multistep.ActionHalt
		}

	} else {
		// Look up the network by it's name label

		networks, err := client.GetNetworkByNameLabel(config.NetworkName)

		if err != nil {
			ui.Error(fmt.Sprintf("Error occured getting Network by name-label: %s", err.Error()))
			return multistep.ActionHalt
		}

		switch {
		case len(networks) == 0:
			ui.Error(fmt.Sprintf("Couldn't find a network with the specified name-label '%s'. Aborting.", config.NetworkName))
			return multistep.ActionHalt
		case len(networks) > 1:
			ui.Error(fmt.Sprintf("Found more than one SR with the name '%s'. The name must be unique. Aborting.", config.NetworkName))
			return multistep.ActionHalt
		}

		network = networks[0]
	}

	if err != nil {
		ui.Say(err.Error())
	}
	_, err = instance.ConnectNetwork(network, "0")

	if err != nil {
		ui.Say(err.Error())
	}

	// Connect the ISO
	//iso_vdi_uuid := state.Get("iso_vdi_uuid").(string)

	isos, err := client.GetVdiByNameLabel(config.IsoName)

	switch {
	case len(isos) == 0:
		ui.Error(fmt.Sprintf("Couldn't find an ISO named '%s'. Aborting", config.IsoName))
		return multistep.ActionHalt
	case len(isos) > 1:
		ui.Error(fmt.Sprintf("Found more than one VDI with name '%s'. Name must be unique. Aborting.", config.IsoName))
		return multistep.ActionHalt
	}

	iso := isos[0]

	//iso, _ := client.GetVdiByUuid(config.IsoUuid)
	//ui.Say("Using VDI: " + iso_vdi_uuid)
	//iso, _ := client.GetVdiByUuid(iso_vdi_uuid)

	err = instance.ConnectVdi(iso, true)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to connect ISO VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	instanceId, err := instance.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}

	state.Put("instance_uuid", instanceId)
	ui.Say(fmt.Sprintf("Created instance '%s'", instanceId))

	return multistep.ActionContinue
}

func (self *stepCreateInstance) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(config)
	if config.ShouldKeepInstance(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)

	if self.instance != nil {
		ui.Say("Destroying VM")
		_ = self.instance.HardShutdown()
		err := self.instance.Destroy()
		if err != nil {
			ui.Error(err.Error())
		}
	}

	if self.vdi != nil {
		ui.Say("Destroying VDI")
		err := self.vdi.Destroy()
		if err != nil {
			ui.Error(err.Error())
		}
	}
}
