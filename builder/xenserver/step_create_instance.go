package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
)

type stepCreateInstance struct {
	InstanceId string
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
		log.Fatal(fmt.Sprintf("Couldn't find a template with the name-label '%s'. Aborting.", config.CloneTemplate))
		return multistep.ActionHalt
	case len(vms) > 1:
		log.Fatal(fmt.Sprintf("Found more than one template with the name '%s'. The name must be unique. Aborting.", config.CloneTemplate))
		return multistep.ActionHalt
	}

	template := vms[0]

	// Clone that VM template
	instance, _ := template.Clone(config.InstanceName)
	instance.SetIsATemplate(false)
	instance.SetStaticMemoryRange(config.InstanceMemory, config.InstanceMemory)
	instance.SetPlatform(config.PlatformArgs)

	// Create VDI for the instance
	var sr *SR

	if config.SrName == "" {
		// Find the default SR
		default_sr, err := client.GetDefaultSR()
		sr = default_sr

		if err != nil {
			log.Fatal(fmt.Sprintf("Error getting default SR: %s", err.Error()))
			return multistep.ActionHalt
		}

	} else {
		// Use the provided name label to find the SR to use
		srs, err := client.GetSRByNameLabel(config.SrName)

		if err != nil {
			log.Fatal(fmt.Sprintf("Error getting default SR: %s", err.Error()))
			return multistep.ActionHalt
		}

		switch {
		case len(srs) == 0:
			log.Fatal(fmt.Sprintf("Couldn't find a SR with the specified name-label '%s'. Aborting.", config.SrName))
			return multistep.ActionHalt
		case len(srs) > 1:
			log.Fatal(fmt.Sprintf("Found more than one SR with the name '%s'. The name must be unique. Aborting.", config.SrName))
			return multistep.ActionHalt
		}

		sr = srs[0]
	}

	vdi, _ := sr.CreateVdi("Packer-disk", config.RootDiskSize)

	instance.ConnectVdi(vdi, false)

	// Connect Network

	var network *Network

	if config.NetworkName == "" {
		// No network has be specified. Use the management interface
		network = new(Network)
		network.Ref = ""
		network.Client = &client

		pifs, err := client.GetPIFs()

		if err != nil {
			log.Fatal(fmt.Sprintf("Error getting PIFs %s", err.Error()))
			return multistep.ActionHalt
		}

		for _, pif := range pifs {
			pif_rec, err := pif.GetRecord()

			if err != nil {
				log.Fatal(fmt.Sprintf("Error getting PIF record: %s", err.Error()))
				return multistep.ActionHalt
			}

			if pif_rec["management"].(bool) {
				network.Ref = pif_rec["network"].(string)
			}

		}

		if network.Ref == "" {
			log.Fatal("Error: couldn't find management network. Aborting.")
			return multistep.ActionHalt
		}

	} else {
		// Look up the network by it's name label

		networks, err := client.GetNetworkByNameLabel(config.NetworkName)

		if err != nil {
			log.Fatal(fmt.Sprintf("Error occured getting Network by name-label: %s", err.Error()))
			return multistep.ActionHalt
		}

		switch {
		case len(networks) == 0:
			log.Fatal(fmt.Sprintf("Couldn't find a network with the specified name-label '%s'. Aborting.", config.NetworkName))
			return multistep.ActionHalt
		case len(networks) > 1:
			log.Fatal(fmt.Sprintf("Found more than one SR with the name '%s'. The name must be unique. Aborting.", config.NetworkName))
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
		log.Fatal(fmt.Sprintf("Couldn't find an ISO named '%s'. Aborting", config.IsoName))
		return multistep.ActionHalt
	case len(isos) > 1:
		log.Fatal(fmt.Sprintf("Found more than one VDI with name '%s'. Name must be unique. Aborting.", config.IsoName))
		return multistep.ActionHalt
	}

	iso := isos[0]

	//iso, _ := client.GetVdiByUuid(config.IsoUuid)
	//ui.Say("Using VDI: " + iso_vdi_uuid)
	//iso, _ := client.GetVdiByUuid(iso_vdi_uuid)
	instance.ConnectVdi(iso, true)

	// Stash the VM reference
	self.InstanceId, _ = instance.GetUuid()
	state.Put("instance_uuid", self.InstanceId)
	state.Put("instance", instance)
	ui.Say(fmt.Sprintf("Created instance '%s'", self.InstanceId))

	return multistep.ActionContinue
}

func (self *stepCreateInstance) Cleanup(state multistep.StateBag) {

	//    client := state.Get("client").(*XenAPIClient)
	//    config := state.Get("config").(config)
	//    ui := state.Get("ui").(packer.Ui)

	// If instance hasn't been created, we have nothing to do.
	if self.InstanceId == "" {
		return
	}

	// @todo: destroy the created instance.

	return
}
