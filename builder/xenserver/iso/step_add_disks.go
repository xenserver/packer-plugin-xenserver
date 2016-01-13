package iso

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type stepAddDisks struct {
	instance *xsclient.VM
	vdiB     *xsclient.VDI
	vdiC     *xsclient.VDI
}


func (self *stepAddDisks) Run(state multistep.StateBag) multistep.StepAction {

	client := state.Get("client").(xsclient.XenAPIClient)
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Add disks")

	sr, err := config.GetSR(client)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	// Add additional disks - this is a terrible hack until my golang gets better
	if (config.DiskBSize !=0) {
		ui.Say(fmt.Sprintf("Creating second disk with size %d", config.DiskBSize))

		vdiB, err := sr.CreateVdi("Packer-disk-b", int64(config.DiskBSize*1024*1024))
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to create packer disk VDI: %s", err.Error()))
		}

		self.vdiB = vdiB

		err = instance.ConnectVdi(vdiB, xsclient.Disk, "")
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
			return multistep.ActionHalt
		}


	}

	if (config.DiskCSize !=0) {
		ui.Say(fmt.Sprintf("Creating third disk with size %d", config.DiskCSize))

		vdiC, err := sr.CreateVdi("Packer-disk-c", int64(config.DiskCSize*1024*1024))
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to create packer disk VDI: %s", err.Error()))
		}

		self.vdiC = vdiC

		err = instance.ConnectVdi(vdiC, xsclient.Disk, "")
		if err != nil {
			ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
			return multistep.ActionHalt
		}

	}

	return multistep.ActionContinue
}

func (self *stepAddDisks) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(config)
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)

	ui.Say("Destroying additional disks")


	//TODO: learn golang and make this better
	if self.vdiB != nil {
		ui.Say("Destroying second VDI")
		err := self.vdiB.Destroy()
		if err != nil {
			ui.Error(err.Error())
		}
	}

	if self.vdiC != nil {
		ui.Say("Destroying third VDI")
		err := self.vdiB.Destroy()
		if err != nil {
			ui.Error(err.Error())
		}
	}

}