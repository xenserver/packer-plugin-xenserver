// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepFindVdi struct {
	VdiName    string
	VdiUuidKey string
	// SRName limits the lookup to one SR; empty means the whole pool (guest_tools_iso is not on iso_sr).
	SRName string
}

func (self *StepFindVdi) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)

	// Both callers wire an optional config value, so an unset name means there is no VDI to look up.
	if self.VdiName == "" {
		return multistep.ActionContinue
	}

	scope := "in the pool"
	if self.SRName != "" {
		scope = fmt.Sprintf("in SR '%s'", self.SRName)
	}

	vdis, err := xenapi.VDI.GetByNameLabel(c.session, self.VdiName)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to look up VDIs named '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}

	if self.SRName != "" {
		vdis, err = vdisInSR(c, vdis, self.SRName)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	switch {
	case len(vdis) == 0:
		ui.Error(fmt.Sprintf("Couldn't find a VDI named '%s' %s", self.VdiName, scope))
		return multistep.ActionHalt
	case len(vdis) > 1:
		ui.Error(fmt.Sprintf("Found more than one VDI with name '%s' %s. Name must be unique", self.VdiName, scope))
		return multistep.ActionHalt
	}

	vdi := vdis[0]

	vdiUuid, err := xenapi.VDI.GetUUID(c.session, vdi)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}
	state.Put(self.VdiUuidKey, vdiUuid)

	return multistep.ActionContinue
}

// VDI.get_by_name_label searches the whole pool, so filter out same-named VDIs on other SRs.
func vdisInSR(c *Connection, vdis []xenapi.VDIRef, srName string) ([]xenapi.VDIRef, error) {
	sr, err := findSRByNameLabel(c, srName)
	if err != nil {
		return nil, err
	}

	var matched []xenapi.VDIRef
	for _, vdi := range vdis {
		vdiSR, err := xenapi.VDI.GetSR(c.session, vdi)
		if err != nil {
			return nil, fmt.Errorf("unable to get the SR of VDI '%s': %s", vdi, err.Error())
		}
		if vdiSR == sr {
			matched = append(matched, vdi)
		}
	}

	return matched, nil
}

func (self *StepFindVdi) Cleanup(state multistep.StateBag) {}
