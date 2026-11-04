// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"log"

	"xenapi"
)

// ClonedDisk is a disk that VM.clone duplicated, with its VBD so that a teardown can unplug it first.
type ClonedDisk struct {
	VBD xenapi.VBDRef
	VDI xenapi.VDIRef
}

// ClonedDisks lists the disks VM.clone duplicated for a VM: as in xapi, every VBD except empty
// ones and CDs, which keep the original ISO. Floppies are included, so GetDisks is too narrow.
func ClonedDisks(c *Connection, vm xenapi.VMRef) ([]ClonedDisk, error) {
	vbds, err := xenapi.VM.GetVBDs(c.session, vm)
	if err != nil {
		return nil, err
	}

	disks := make([]ClonedDisk, 0, len(vbds))
	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(c.session, vbd)
		if err != nil {
			log.Printf("Unable to get VBD record: %s", err.Error())
			continue
		}
		if rec.Empty || rec.Type == xenapi.VbdTypeCD {
			continue
		}
		disks = append(disks, ClonedDisk{VBD: vbd, VDI: rec.VDI})
	}

	return disks, nil
}
