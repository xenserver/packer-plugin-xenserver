package xva

import (
	"fmt"
	"os"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/simonfuhrer/go-xenserver-client"
	xscommon "github.com/rdobson/packer-builder-xenserver/builder/xenserver/common"
)

type stepImportInstance struct {
	instance *xsclient.VM
	vdi      *xsclient.VDI
}

func (s *stepImportInstance) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(xsclient.XenAPIClient)
	config := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Import Instance")

	// find the SR
	sr, err := config.GetSR(client)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Open the file for reading (NB: httpUpload closes the file for us)
	fh, err := os.Open(config.SourcePath)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to open XVA '%s': %s", config.SourcePath, err.Error()))
		return multistep.ActionHalt
	}

	result, err := xscommon.HTTPUpload(fmt.Sprintf("https://%s/import?session_id=%s&sr_id=%s",
		client.Host,
		client.Session.(string),
		sr.Ref,
	), fh, state)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to upload VDI: %s", err.Error()))
		return multistep.ActionHalt
	}
	if result == nil {
		ui.Error("XAPI did not reply with an instance reference")
		return multistep.ActionHalt
	}

	instance := xsclient.VM(*result)

	/*
		err = instance.SetStaticMemoryRange(config.VMMemory*1024*1024, config.VMMemory*1024*1024)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM memory=%d: %s", config.VMMemory*1024*1024, err.Error()))
			return multistep.ActionHalt
		}

		instance.SetPlatform(config.PlatformArgs)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM platform: %s", err.Error()))
			return multistep.ActionHalt
		}

		// Connect Network

		var network *xscommon.Network

		if config.NetworkName == "" {
			// No network has be specified. Use the management interface
			network = new(xscommon.Network)
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
				ui.Error(fmt.Sprintf("Found more than one network with the name '%s'. The name must be unique. Aborting.", config.NetworkName))
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

	*/

	instanceId, err := instance.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}
	state.Put("instance_uuid", instanceId)

	instance.SetDescription(config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Imported instance '%s'", instanceId))

	return multistep.ActionContinue
}

func (s *stepImportInstance) Cleanup(state multistep.StateBag) {
	/*
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
	*/
}
