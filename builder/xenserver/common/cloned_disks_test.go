// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"testing"

	"xenapi"
)

func TestClonedDisks_TakesEveryDiskAClonePaidFor(t *testing.T) {
	records := map[string]interface{}{
		"OpaqueRef:vbd-disk":   map[string]interface{}{"type": "Disk", "VDI": "OpaqueRef:vdi-disk"},
		"OpaqueRef:vbd-floppy": map[string]interface{}{"type": "Floppy", "VDI": "OpaqueRef:vdi-floppy"},
		"OpaqueRef:vbd-cd":     map[string]interface{}{"type": "CD", "VDI": "OpaqueRef:vdi-iso"},
		"OpaqueRef:vbd-empty":  map[string]interface{}{"type": "CD", "empty": true, "VDI": "OpaqueRef:NULL"},
	}

	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_VBDs": []interface{}{
			"OpaqueRef:vbd-disk", "OpaqueRef:vbd-cd", "OpaqueRef:vbd-floppy", "OpaqueRef:vbd-empty",
		},
		"VBD.get_record": func(args []interface{}) interface{} {
			return records[args[0].(string)]
		},
	})

	disks, err := ClonedDisks(fake.conn(), "OpaqueRef:vm")
	if err != nil {
		t.Fatalf("ClonedDisks: %s", err)
	}

	want := []ClonedDisk{
		{VBD: "OpaqueRef:vbd-disk", VDI: "OpaqueRef:vdi-disk"},
		{VBD: "OpaqueRef:vbd-floppy", VDI: "OpaqueRef:vdi-floppy"},
	}
	if len(disks) != len(want) {
		t.Fatalf("ClonedDisks = %v, want %v", disks, want)
	}
	for i, disk := range disks {
		if disk != want[i] {
			t.Errorf("ClonedDisks[%d] = %v, want %v", i, disk, want[i])
		}
	}
}

func TestClonedDisks_KeepsSweepingPastAnUnreadableVBD(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_VBDs": []interface{}{"OpaqueRef:vbd-gone", "OpaqueRef:vbd-disk"},
		"VBD.get_record": func(args []interface{}) interface{} {
			if args[0] == "OpaqueRef:vbd-gone" {
				return errors.New("HANDLE_INVALID")
			}
			return map[string]interface{}{"type": "Disk", "VDI": "OpaqueRef:vdi-disk"}
		},
	})

	disks, err := ClonedDisks(fake.conn(), "OpaqueRef:vm")
	if err != nil {
		t.Fatalf("ClonedDisks: %s", err)
	}
	if len(disks) != 1 || disks[0].VDI != xenapi.VDIRef("OpaqueRef:vdi-disk") {
		t.Errorf("ClonedDisks = %v, want the readable disk", disks)
	}
}

func TestClonedDisks_ReportsAFailedVBDList(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_VBDs": errors.New("SESSION_INVALID"),
	})

	if _, err := ClonedDisks(fake.conn(), "OpaqueRef:vm"); err == nil {
		t.Fatal("a VM whose VBDs could not be listed was reported as having no disks")
	}
}
