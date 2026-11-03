// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"

	"xenapi"
)

// diskSRs is an interface only so that tests can fake XenAPI.
type diskSRs interface {
	// SRs reports the SR each of the VM's disks is on, one entry per disk.
	SRs() ([]xenapi.SRRef, error)
}

type xapiDiskSRs struct {
	c  *Connection
	vm xenapi.VMRef
}

func (d xapiDiskSRs) SRs() ([]xenapi.SRRef, error) {
	vbds, err := xenapi.VM.GetVBDs(d.c.session, d.vm)
	if err != nil {
		return nil, err
	}

	srs := make([]xenapi.SRRef, 0, len(vbds))
	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(d.c.session, vbd)
		if err != nil {
			return nil, err
		}
		// As in ClonedDisks, but an unreadable VBD is an error: it decides clone versus copy.
		if rec.Empty || rec.Type == xenapi.VbdTypeCD {
			continue
		}
		sr, err := xenapi.VDI.GetSR(d.c.session, rec.VDI)
		if err != nil {
			return nil, err
		}
		srs = append(srs, sr)
	}
	return srs, nil
}

// TemplateNeedsCopy reports whether target needs VM.copy, as VM.clone keeps each disk on its own SR.
func TemplateNeedsCopy(c *Connection, template xenapi.VMRef, target xenapi.SRRef) (bool, error) {
	return templateNeedsCopy(xapiDiskSRs{c: c, vm: template}, target)
}

func templateNeedsCopy(api diskSRs, target xenapi.SRRef) (bool, error) {
	srs, err := api.SRs()
	if err != nil {
		return false, fmt.Errorf("could not read which SRs the source's disks are on: %w", err)
	}

	for _, sr := range srs {
		if sr != target {
			return true, nil
		}
	}

	return false, nil
}
