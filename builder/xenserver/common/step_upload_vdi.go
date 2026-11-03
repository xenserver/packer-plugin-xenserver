// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepUploadVdi struct {
	VdiNameFunc   func() string
	ImagePathFunc func() string
	VdiUuidKey    string
}

func (self *StepUploadVdi) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)

	imagePath := self.ImagePathFunc()
	vdiName := self.VdiNameFunc()
	if imagePath == "" {
		// Both builders always register this step; no cd_files or iso_url means nothing to upload.
		return multistep.ActionContinue
	}

	ui.Say(fmt.Sprintf("Step: Upload VDI '%s'", vdiName))

	// A new VDI needs a writable SR, and iso_sr is usually read-only (SR_BACKEND_FAILURE_78).
	sr, err := config.GetSR(c)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to find an SR to upload '%s' to: %v", vdiName, err))
		return multistep.ActionHalt
	}

	fh, err := os.Open(imagePath)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to open disk image '%s': %s", imagePath, err.Error()))
		return multistep.ActionHalt
	}
	defer fh.Close()

	fstat, err := fh.Stat()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to stat disk image '%s': %s", imagePath, err.Error()))
		return multistep.ActionHalt
	}
	fileLength := fstat.Size()

	vdi, err := xenapi.VDI.Create(c.session, xenapi.VDIRecord{
		NameLabel:   vdiName,
		VirtualSize: int(fileLength),
		Type:        "user",
		Sharable:    false,
		ReadOnly:    false,
		SR:          sr,
		OtherConfig: map[string]string{
			"temp": "temp",
		},
	})
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create VDI '%s': %s", vdiName, err.Error()))
		return multistep.ActionHalt
	}

	vdiUuid, err := xenapi.VDI.GetUUID(c.session, vdi)
	if err != nil {
		// Cleanup keys off the UUID, so without this the VDI is orphaned on the SR.
		if destroyErr := xenapi.VDI.Destroy(c.session, vdi); destroyErr != nil {
			log.Printf("Unable to destroy VDI '%s' whose UUID could not be read: %s", vdiName, destroyErr.Error())
		}
		ui.Error(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", vdiName, err.Error()))
		return multistep.ActionHalt
	}
	state.Put(self.VdiUuidKey, vdiUuid)

	_, err = HTTPUpload(fmt.Sprintf("https://%s/import_raw_vdi?vdi=%s&session_id=%s",
		c.Host,
		vdi,
		c.GetSessionRef(),
	), fh, state)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to upload VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (self *StepUploadVdi) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)

	vdiName := self.VdiNameFunc()

	vdiUuidRaw, ok := state.GetOk(self.VdiUuidKey)
	if !ok {
		return
	}

	vdiUuid := vdiUuidRaw.(string)
	if vdiUuid == "" {
		return
	}

	vdi, err := xenapi.VDI.GetByUUID(c.session, vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Can't get VDI '%s': %s", vdiUuid, err.Error()))
		return
	}

	// Decide by VBDs, not keep_vm: a kept VM has usually had this disk detached by StepDetachVdi.
	vbds, err := xenapi.VDI.GetVBDs(c.session, vdi)
	if err != nil {
		ui.Error(fmt.Sprintf("Can't tell whether VDI '%s' is still attached, so leaving it: %s", vdiUuid, err.Error()))
		return
	}
	if len(vbds) > 0 {
		log.Printf("Keeping VDI '%s': still attached by %d VBD(s)", vdiUuid, len(vbds))
		return
	}

	// An interrupted import_raw_vdi takes a while to release the VDI.
	for i := 0; i < 3; i++ {
		log.Printf("Trying to destroy VDI...")
		err = xenapi.VDI.Destroy(c.session, vdi)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		ui.Error(fmt.Sprintf("Can't destroy VDI '%s': %s", vdiUuid, err.Error()))
		return
	}
	ui.Say(fmt.Sprintf("Destroyed VDI '%s'", vdiName))

	state.Put(self.VdiUuidKey, "")
}
