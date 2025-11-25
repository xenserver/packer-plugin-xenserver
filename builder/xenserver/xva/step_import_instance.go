package xva

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	xsclient "github.com/terra-farm/go-xen-api-client"
	xscommon "github.com/xenserver/packer-builder-xenserver/builder/xenserver/common"
)

type stepImportInstance struct {
	instance xsclient.VMRef
	vdi      xsclient.VDIRef
}

func (self *stepImportInstance) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*xscommon.Connection)
	config := state.Get("config").(xscommon.Config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Import Instance")

	if config.SourcePath == "" {
		log.Println("Skipping importing instance - no `source_path` configured.")
		return multistep.ActionContinue
	}

	// find the SR
	srs, err := c.GetClient().SR.GetAll(c.GetSessionRef())
	sr := srs[0]
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
		c.Host,
		c.GetSession(),
		sr,
	), fh, state)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to upload VDI: %s", err.Error()))
		return multistep.ActionHalt
	}
	if result == "" {
		ui.Error("XAPI did not reply with an instance reference")
		return multistep.ActionHalt
	}

	instance := xsclient.VMRef(result)

	instanceId, err := c.GetClient().VM.GetUUID(c.GetSessionRef(), instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}
	state.Put("instance_uuid", instanceId)

	err = c.GetClient().VM.SetVCPUsMax(c.GetSessionRef(), instance, int(config.VCPUsMax))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM VCPUs Max=%d: %s", config.VCPUsMax, err.Error()))
		return multistep.ActionHalt
	}

	err = c.GetClient().VM.SetVCPUsAtStartup(c.GetSessionRef(), instance, int(config.VCPUsAtStartup))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=%d: %s", config.VCPUsAtStartup, err.Error()))
		return multistep.ActionHalt
	}

	err = c.GetClient().VM.SetNameDescription(c.GetSessionRef(), instance, config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = xscommon.AddVMTags(c, instance, config.VMTags)
	if err != nil {
		ui.Error(fmt.Sprintf("Failed to add tags: %s", err.Error()))
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Imported instance '%s'", instanceId))

	return multistep.ActionContinue
}

func (self *stepImportInstance) Cleanup(state multistep.StateBag) {
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
