package xva

import (
	"fmt"
	"os"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	xscommon "github.com/xenserver/packer-builder-xenserver/builder/xenserver/common"
)

type stepImportInstance struct {
	instance *xsclient.VM
}

func (self *stepImportInstance) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(xsclient.XenAPIClient)
	config := state.Get("config").(config)
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

	instanceId, err := instance.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}
	state.Put("instance_uuid", instanceId)

	self.instance, err = client.GetVMByUuid(instanceId)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", instanceId, err.Error()))
		return multistep.ActionHalt
	}

	self.instance.SetDescription(config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Imported instance '%s'", instanceId))

	return multistep.ActionContinue
}

func (self *stepImportInstance) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(config)
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)

	// We want to perform a 'xe vm-uninstall'
	// Uninstall a VM. This operation will destroy those VDIs that are marked RW and connected to this VM only.
	if self.instance != nil {
		vbds, err := self.instance.GetVBDs()
		if err != nil {
			ui.Error(err.Error())
		}

		// Destroy VDIs before destroying VM, since vm.Destroy() also destroys VBDs,
		// in which case we will be unable to find the VDIs
		for _, vbd := range vbds {
			vbd_rec, err := vbd.GetRecord()
			if err != nil {
				ui.Error(err.Error())
				continue
			}
			if vbd_rec["mode"].(string) == "RW" {
				vdi, err := vbd.GetVDI()
				if err != nil {
					ui.Error(err.Error())
					continue
				}
				vdi_vbds, err := vdi.GetVBDs()
				if err != nil {
					ui.Error(err.Error())
					continue
				}
				// If connected to this VM only
				if len(vdi_vbds) <= 1 {
					ui.Say("Destroying VDI")
					err = vdi.Destroy()
					if err != nil {
						ui.Error(err.Error())
					}
				}
			}
		}

		ui.Say("Destroying VM")
		_ = self.instance.HardShutdown()
		err = self.instance.Destroy()
		if err != nil {
			ui.Error(err.Error())
		}
	}
}
