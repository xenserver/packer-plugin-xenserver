package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type StepConfigureNetworking struct{}

func (self *StepConfigureNetworking) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	ui.Say("Step: Configure Networking")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	vifs, err := instance.GetVIFs()
	if err != nil {
		ui.Error(fmt.Sprintf("Error getting VIFs: %s", err.Error()))
		return multistep.ActionHalt
	}

	if len(vifs) != 0 {
		ui.Say("VM already has VIFs. Skipping step.")
		return multistep.ActionContinue
	}

	ui.Say("VM has no VIFs. Automatically configuring networking.")

	var network *xsclient.Network

	if len(config.NetworkNames) == 0 {
		// No network has be specified. Use the management interface
		network = new(xsclient.Network)
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

		_, err = instance.ConnectNetwork(network, "0")

		if err != nil {
			ui.Say(err.Error())
		}

	} else {
		// Look up each network by it's name label
		for i, networkNameLabel := range config.NetworkNames {
			networks, err := client.GetNetworkByNameLabel(networkNameLabel)

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
			_, err = instance.ConnectNetwork(networks[0], vifIndexString)

			if err != nil {
				ui.Say(err.Error())
			}
		}
	}

	return multistep.ActionContinue
}

func (self *StepConfigureNetworking) Cleanup(state multistep.StateBag) {}