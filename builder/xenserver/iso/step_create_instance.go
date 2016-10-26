package iso

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type stepCreateInstance struct {
	instance *xsclient.VM
	vdi      *xsclient.VDI
}

func (self *stepCreateInstance) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(xsclient.XenAPIClient)
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
	instance, err := template.Clone(config.VMName)
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

	err = instance.SetStaticMemoryRange(uint64(config.VMMemory*1024*1024), uint64(config.VMMemory*1024*1024))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM memory=%d: %s", config.VMMemory*1024*1024, err.Error()))
		return multistep.ActionHalt
	}

	err = instance.SetPlatform(config.PlatformArgs)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM platform: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = instance.SetDescription(config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	if len(config.VMOtherConfig) != 0 {
		vm_other_config, err := instance.GetOtherConfig()
		if err != nil {
			ui.Error(fmt.Sprintf("Error getting VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}
		for key, value := range config.VMOtherConfig {
			vm_other_config[key] = value
		}
		err = instance.SetOtherConfig(vm_other_config)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	// Create VDI for the instance

	sr, err := config.GetSR(client)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	vdi, err := sr.CreateVdi(config.VMName+" 0", int64(config.DiskSize*1024*1024))
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.vdi = vdi

	err = instance.ConnectVdi(vdi, xsclient.Disk, "")
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Connect Network

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
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)

	if self.instance != nil {
		ui.Say("Destroying VM")
		_ = self.instance.HardShutdown() // redundant, just in case
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
