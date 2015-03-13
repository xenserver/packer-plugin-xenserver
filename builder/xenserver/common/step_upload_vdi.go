package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	"log"
	"os"
	"time"
)

type StepUploadVdi struct {
	VdiName       string
	ImagePathFunc func() string
	VdiUuidKey    string
}

func (self *StepUploadVdi) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	imagePath := self.ImagePathFunc()
	if imagePath == "" {
		// skip if no disk image to attach
		return multistep.ActionContinue
	}

	ui.Say(fmt.Sprintf("Step: Upload VDI '%s'", self.VdiName))

	// Create VDI for the image
	sr, err := config.GetSR(client)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Open the file for reading (NB: HTTPUpload closes the file for us)
	fh, err := os.Open(imagePath)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to open disk image '%s': %s", imagePath, err.Error()))
		return multistep.ActionHalt
	}

	// Get file length
	fstat, err := fh.Stat()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to stat disk image '%s': %s", imagePath, err.Error()))
		return multistep.ActionHalt
	}
	fileLength := fstat.Size()

	// Create the VDI
	vdi, err := sr.CreateVdi(self.VdiName, fileLength)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create VDI '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}

	vdiUuid, err := vdi.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}
	state.Put(self.VdiUuidKey, vdiUuid)

	_, err = HTTPUpload(fmt.Sprintf("https://%s/import_raw_vdi?vdi=%s&session_id=%s",
		client.Host,
		vdi.Ref,
		client.Session.(string),
	), fh, state)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to upload VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (self *StepUploadVdi) Cleanup(state multistep.StateBag) {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	if config.ShouldKeepVM(state) {
		return
	}

	vdiUuidRaw, ok := state.GetOk(self.VdiUuidKey)
	if !ok {
		// VDI doesn't exist
		return
	}

	vdiUuid := vdiUuidRaw.(string)
	if vdiUuid == "" {
		// VDI already cleaned up
		return
	}

	vdi, err := client.GetVdiByUuid(vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Can't get VDI '%s': %s", vdiUuid, err.Error()))
		return
	}

	// an interrupted import_raw_vdi takes a while to release the VDI
	// so try several times
	for i := 0; i < 3; i++ {
		log.Printf("Trying to destroy VDI...")
		err = vdi.Destroy()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		ui.Error(fmt.Sprintf("Can't destroy VDI '%s': %s", vdiUuid, err.Error()))
		return
	}
	ui.Say(fmt.Sprintf("Destroyed VDI '%s'", self.VdiName))

	state.Put(self.VdiUuidKey, "")
}
