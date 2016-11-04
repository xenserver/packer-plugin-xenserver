package xva

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type stepInstantiateTemplate struct {
	instance *xsclient.VM
}

func (self *stepInstantiateTemplate) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(config)

	if config.SourceTemplate == "" {
		return multistep.ActionContinue
	}

	client := state.Get("client").(xsclient.XenAPIClient)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Instantiate Template")

	var template *xsclient.VM
	var err error

	if len(config.SourceTemplate) >= 7 && config.SourceTemplate[:7] == "uuid://" {
		templateUuid := config.SourceTemplate[7:]

		template, err = client.GetVMByUuid(templateUuid)
		if err != nil {
			ui.Error(fmt.Sprintf("Could not get template with UUID '%s': %s", templateUuid, err.Error()))
			ui.Error("Defaulting to use \"source_path\".")
			return multistep.ActionContinue
		}
	} else {
		templates, err := client.GetVMByNameLabel(config.SourceTemplate)
		if err != nil {
			ui.Error(fmt.Sprintf("Error getting template: %s", err.Error()))
			ui.Error("Defaulting to use \"source_path\".")
			return multistep.ActionContinue
		}

		switch {
		case len(templates) == 0:
			ui.Error(fmt.Sprintf("Couldn't find a template with the name-label '%s'.", config.SourceTemplate))
			ui.Error("Defaulting to use \"source_path\".")
			return multistep.ActionContinue
		case len(templates) > 1:
			ui.Error(fmt.Sprintf("Found more than one template with the name '%s'. The name must be unique.", config.SourceTemplate))
			ui.Error("Defaulting to use \"source_path\".")
			return multistep.ActionContinue
		}

		template = templates[0]
	}

	self.instance, err = template.Clone(config.VMName)
	if err != nil {
		ui.Error(fmt.Sprintf("Error cloning template: %s", err.Error()))
		return multistep.ActionHalt
	}

	instanceId, err := self.instance.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}
	state.Put("instance_uuid", instanceId)

	err = self.instance.SetIsATemplate(false)
	if err != nil {
		ui.Error(fmt.Sprintf("Error converting template to a VM: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = self.instance.SetDescription(config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Instantiated template '%s'", instanceId))

	return multistep.ActionContinue
}

func (self *stepInstantiateTemplate) Cleanup(state multistep.StateBag) {
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
