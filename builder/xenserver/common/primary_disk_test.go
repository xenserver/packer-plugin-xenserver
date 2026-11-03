// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"

	"xenapi"
)

type fakeDisks struct {
	vbds      []xenapi.VBDRef
	records   map[xenapi.VBDRef]xenapi.VBDRecord
	sizes     map[xenapi.VDIRef]int
	listErr   error
	recordErr error
	sizeErr   error
	resizeErr error
	resized   map[xenapi.VDIRef]int
}

func (f *fakeDisks) ListVBDs() ([]xenapi.VBDRef, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.vbds, nil
}

func (f *fakeDisks) VBDRecord(vbd xenapi.VBDRef) (xenapi.VBDRecord, error) {
	if f.recordErr != nil {
		return xenapi.VBDRecord{}, f.recordErr
	}
	return f.records[vbd], nil
}

func (f *fakeDisks) VirtualSize(vdi xenapi.VDIRef) (int, error) {
	if f.sizeErr != nil {
		return 0, f.sizeErr
	}
	return f.sizes[vdi], nil
}

func (f *fakeDisks) Resize(vdi xenapi.VDIRef, bytes int) error {
	if f.resized == nil {
		f.resized = map[xenapi.VDIRef]int{}
	}
	f.resized[vdi] = bytes
	return f.resizeErr
}

func oneDisk(sizeGB int) *fakeDisks {
	return &fakeDisks{
		vbds: []xenapi.VBDRef{"OpaqueRef:vbd-0"},
		records: map[xenapi.VBDRef]xenapi.VBDRecord{
			"OpaqueRef:vbd-0": {Type: xenapi.VbdTypeDisk, Userdevice: "0", VDI: "OpaqueRef:vdi-0"},
		},
		sizes: map[xenapi.VDIRef]int{"OpaqueRef:vdi-0": sizeGB * BytesPerGB},
	}
}

func TestSizePrimaryDisk_GrowsTheDiskOnDeviceZero(t *testing.T) {
	api := oneDisk(24)

	outcome, err := sizePrimaryDisk(api, 40*BytesPerGB)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if outcome != DiskGrown {
		t.Errorf("outcome = %v, want DiskGrown", outcome)
	}
	if got := api.resized["OpaqueRef:vdi-0"]; got != 40*BytesPerGB {
		t.Errorf("resized to %d bytes, want %d", got, 40*BytesPerGB)
	}
}

func TestSizePrimaryDisk_RejectsASizeSmallerThanTheSource(t *testing.T) {
	api := oneDisk(40)

	outcome, err := sizePrimaryDisk(api, 24*BytesPerGB)

	if err == nil {
		t.Fatal("a disk_size_gb below the source's disk size must fail the build, not be ignored")
	}
	if outcome != DiskUnchanged {
		t.Errorf("outcome = %v, want DiskUnchanged", outcome)
	}
	if len(api.resized) != 0 {
		t.Errorf("wrote %v; a shrink must not reach VDI.resize", api.resized)
	}

	for _, want := range []string{"disk_size_gb", "24 GB", "40 GB"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
}

func TestSizePrimaryDisk_LeavesADiskThatAlreadyMatchesAlone(t *testing.T) {
	api := oneDisk(40)

	outcome, err := sizePrimaryDisk(api, 40*BytesPerGB)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if outcome != DiskUnchanged {
		t.Errorf("outcome = %v, want DiskUnchanged", outcome)
	}
	if len(api.resized) != 0 {
		t.Errorf("wrote %v to a disk that was already the right size", api.resized)
	}
}

func TestSizePrimaryDisk_TouchesNeitherDataDisksNorCDs(t *testing.T) {
	api := &fakeDisks{
		vbds: []xenapi.VBDRef{"OpaqueRef:cd", "OpaqueRef:data", "OpaqueRef:root"},
		records: map[xenapi.VBDRef]xenapi.VBDRecord{
			"OpaqueRef:cd":   {Type: xenapi.VbdTypeCD, Userdevice: "3", VDI: "OpaqueRef:vdi-cd"},
			"OpaqueRef:data": {Type: xenapi.VbdTypeDisk, Userdevice: "1", VDI: "OpaqueRef:vdi-data"},
			"OpaqueRef:root": {Type: xenapi.VbdTypeDisk, Userdevice: "0", VDI: "OpaqueRef:vdi-root"},
		},
		sizes: map[xenapi.VDIRef]int{
			"OpaqueRef:vdi-cd":   1 * BytesPerGB,
			"OpaqueRef:vdi-data": 8 * BytesPerGB,
			"OpaqueRef:vdi-root": 24 * BytesPerGB,
		},
	}

	outcome, err := sizePrimaryDisk(api, 40*BytesPerGB)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if outcome != DiskGrown {
		t.Errorf("outcome = %v, want DiskGrown", outcome)
	}
	if len(api.resized) != 1 {
		t.Fatalf("resized %v, want only the device 0 disk", api.resized)
	}
	if _, ok := api.resized["OpaqueRef:vdi-root"]; !ok {
		t.Errorf("resized %v, want OpaqueRef:vdi-root", api.resized)
	}
}

func TestSizePrimaryDisk_ReportsASourceWithNoDiskOnDeviceZero(t *testing.T) {
	api := &fakeDisks{
		vbds: []xenapi.VBDRef{"OpaqueRef:cd"},
		records: map[xenapi.VBDRef]xenapi.VBDRecord{
			"OpaqueRef:cd": {Type: xenapi.VbdTypeCD, Userdevice: "0", VDI: "OpaqueRef:vdi-cd"},
		},
		sizes: map[xenapi.VDIRef]int{"OpaqueRef:vdi-cd": 1 * BytesPerGB},
	}

	_, err := sizePrimaryDisk(api, 40*BytesPerGB)

	if err == nil {
		t.Fatal("disk_size_gb with nothing to apply it to must be reported")
	}
	if !strings.Contains(err.Error(), "disk_size_gb") {
		t.Errorf("error should name the option the user can act on, got: %s", err)
	}
	if len(api.resized) != 0 {
		t.Errorf("wrote %v; a CD is not the primary disk", api.resized)
	}
}

func TestSizePrimaryDisk_ReportsReadFailuresAndWritesNothing(t *testing.T) {
	readFailed := errors.New("API error: code 1, message SESSION_INVALID")

	cases := map[string]func(*fakeDisks){
		"listing the VBDs":     func(f *fakeDisks) { f.listErr = readFailed },
		"reading a VBD record": func(f *fakeDisks) { f.recordErr = readFailed },
		"reading the size":     func(f *fakeDisks) { f.sizeErr = readFailed },
	}

	for name, breakIt := range cases {
		t.Run(name, func(t *testing.T) {
			api := oneDisk(24)
			breakIt(api)

			outcome, err := sizePrimaryDisk(api, 40*BytesPerGB)

			if err == nil {
				t.Fatal("a failed read must be reported")
			}
			if !errors.Is(err, readFailed) {
				t.Errorf("error should wrap the XAPI failure, got: %s", err)
			}
			if outcome != DiskUnchanged {
				t.Errorf("outcome = %v after a failed read", outcome)
			}
			if len(api.resized) != 0 {
				t.Errorf("wrote %v despite not knowing the current size", api.resized)
			}
		})
	}
}

func TestSizePrimaryDisk_ReportsAResizeFailure(t *testing.T) {
	resizeFailed := errors.New("API error: code 1, message SR_NO_SPACE")
	api := oneDisk(24)
	api.resizeErr = resizeFailed

	outcome, err := sizePrimaryDisk(api, 40*BytesPerGB)

	if err == nil {
		t.Fatal("a failed resize must be reported")
	}
	if !errors.Is(err, resizeFailed) {
		t.Errorf("error should wrap the XAPI failure, got: %s", err)
	}
	if outcome != DiskUnchanged {
		t.Errorf("outcome = %v after a failed resize", outcome)
	}
}

func TestSizePrimaryDisk_SuggestedMinimumIsNeverAShrink(t *testing.T) {
	api := oneDisk(40)
	api.sizes["OpaqueRef:vdi-0"] = 40*BytesPerGB + 1

	_, err := sizePrimaryDisk(api, 24*BytesPerGB)

	if err == nil {
		t.Fatal("a disk_size_gb below the source's disk size must be reported")
	}
	if !strings.Contains(err.Error(), "41 or more") {
		t.Errorf("error should round the suggested minimum up to 41, got: %s", err)
	}
}
