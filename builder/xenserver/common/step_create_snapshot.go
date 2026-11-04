// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepCreateSnapshot struct{}

func (StepCreateSnapshot) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	instance_uuid := state.Get("instance_uuid").(string)

	if !config.CreateSnapshot {
		ui.Say("Skipping Create Snapshot")
		return multistep.ActionContinue
	}

	instance, err := xenapi.VM.GetByUUID(c.session, instance_uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM with UUID '%s': %s", instance_uuid, err.Error()))
		return multistep.ActionHalt
	}

	_, err = xenapi.VM.Snapshot(c.session, instance, config.SnapshotName, []xenapi.VDIRef{})

	if err != nil {
		ui.Error(fmt.Sprintf("failed to create a snapshot of VM '%s' with error: %v", instance_uuid, err))
		return multistep.ActionHalt
	}

	ui.Message("Successfully created snapshot")
	return multistep.ActionContinue
}

func (StepCreateSnapshot) Cleanup(state multistep.StateBag) {
	// Nothing to undo: the snapshot is a requested artifact, and VM.destroy leaves it.
}
