package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/nilshell/xmlrpc"
	xsclient "github.com/xenserver/go-xenserver-client"
	"strings"
)

type StepConfigureDiscDrives struct{}

func (self *StepConfigureDiscDrives) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("commonconfig").(CommonConfig)
	client := state.Get("client").(xsclient.XenAPIClient)
	ui.Say("Step: Configure disc drives")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	vbds, err := instance.GetVBDs()
	if err != nil {
		ui.Error(fmt.Sprintf("Error getting VBDs: %s", err.Error()))
		return multistep.ActionHalt
	}

	var current_number_of_disc_drives int = 0
	for _, vbd := range vbds {
		vbd_rec, err := vbd.GetRecord()
		if err != nil {
			ui.Error(fmt.Sprintf("Error getting VBD record: %s", err.Error()))
			return multistep.ActionHalt
		}
		if vbd_rec["type"].(string) == "CD" {
			if current_number_of_disc_drives < config.DiscDrives {
				ui.Say("Ejecting disc drive")
				err = vbd.Eject()
				if err != nil && !strings.Contains(err.Error(), "VBD_IS_EMPTY") {
					ui.Error(fmt.Sprintf("Error ejecting VBD: %s", err.Error()))
					return multistep.ActionHalt
				}
				current_number_of_disc_drives++
			} else {
				ui.Say("Destroying excess disc drive")
				_ = vbd.Eject()
				_ = vbd.Unplug()
				err = vbd.Destroy()
				if err != nil {
					ui.Error(fmt.Sprintf("Error destroying VBD: %s", err.Error()))
					return multistep.ActionHalt
				}
			}
		}
	}

	if current_number_of_disc_drives < config.DiscDrives {
		vbd_rec := make(xmlrpc.Struct)
		vbd_rec["VM"] = instance.Ref
		vbd_rec["VDI"] = "OpaqueRef:NULL"
		vbd_rec["userdevice"] = "autodetect"
		vbd_rec["empty"] = true
		vbd_rec["other_config"] = make(xmlrpc.Struct)
		vbd_rec["qos_algorithm_type"] = ""
		vbd_rec["qos_algorithm_params"] = make(xmlrpc.Struct)
		vbd_rec["mode"] = "RO"
		vbd_rec["bootable"] = true
		vbd_rec["unpluggable"] = false
		vbd_rec["type"] = "CD"
		var failures int = 0
		for current_number_of_disc_drives < config.DiscDrives {
			ui.Say("Creating disc drive")

			result := xsclient.APIResult{}
			err := client.APICall(&result, "VBD.create", vbd_rec)

			if err != nil {
				failures++
				if failures < 3 {
					ui.Error("Error creating disc drive. Retrying...")
					continue
				} else {
					ui.Error("Failed to create disc drive after 3 tries.")
					return multistep.ActionHalt
				}
			}

			vbd_ref := result.Value.(string)

			result = xsclient.APIResult{}
			err = client.APICall(&result, "VBD.get_uuid", vbd_ref)

			if err != nil {
				failures++
				if failures < 3 {
					ui.Error("Error verifying disc drive. Retrying...")
					continue
				} else {
					ui.Error("Failed to create disc drive after 3 tries.")
					return multistep.ActionHalt
				}
			}

			current_number_of_disc_drives++
		}
	}

	return multistep.ActionContinue
}

func (self *StepConfigureDiscDrives) Cleanup(state multistep.StateBag) {}
