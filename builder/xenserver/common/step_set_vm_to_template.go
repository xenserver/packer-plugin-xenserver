// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepSetVmToTemplate struct{}

func (StepSetVmToTemplate) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	instance_uuid := state.Get("instance_uuid").(string)

	// A `format = "none"` build's output is the template itself.
	if config.Format != "none" {
		ui.Say("Skipping Convert to Template")
		return multistep.ActionContinue
	}

	instance, err := xenapi.VM.GetByUUID(c.session, instance_uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM with UUID '%s': %s", instance_uuid, err.Error()))
		return multistep.ActionHalt
	}

	// VM.clone copied these from clone_template.
	stripped, err := StripProvisioningKeys(c, instance)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	for _, key := range stripped {
		ui.Message(fmt.Sprintf("Removed the inherited other_config:%s: the template ships the disk this "+
			"build installed onto, and does not provision a fresh one", key))
	}

	err = xenapi.VM.SetIsATemplate(c.session, instance, true)

	if err != nil {
		ui.Error(fmt.Sprintf("failed to set VM '%s' as a template with error: %v", instance_uuid, err))
		return multistep.ActionHalt
	}

	ui.Message("Successfully set VM as a template")
	return multistep.ActionContinue
}

func (StepSetVmToTemplate) Cleanup(state multistep.StateBag) {}
