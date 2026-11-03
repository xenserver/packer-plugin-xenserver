// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"

	"xenapi"
)

type fakeDiskSRs struct {
	srs []xenapi.SRRef
	err error
}

func (f fakeDiskSRs) SRs() ([]xenapi.SRRef, error) {
	return f.srs, f.err
}

func TestTemplateNeedsCopy(t *testing.T) {
	const target = xenapi.SRRef("OpaqueRef:target")

	for name, tc := range map[string]struct {
		srs  []xenapi.SRRef
		want bool
	}{
		"single disk elsewhere":       {[]xenapi.SRRef{"OpaqueRef:other"}, true},
		"single disk already on it":   {[]xenapi.SRRef{target}, false},
		"every disk already on it":    {[]xenapi.SRRef{target, target}, false},
		"one of several is elsewhere": {[]xenapi.SRRef{target, "OpaqueRef:other"}, true},
		"no disks at all":             {[]xenapi.SRRef{}, false},
	} {
		got, err := templateNeedsCopy(fakeDiskSRs{srs: tc.srs}, target)
		if err != nil {
			t.Errorf("%s: unexpected error: %s", name, err)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: templateNeedsCopy() = %t, want %t", name, got, tc.want)
		}
	}
}

func TestTemplateNeedsCopy_ReportsAFailedRead(t *testing.T) {
	readFailed := errors.New("API error: code 1, message HANDLE_INVALID")

	needsCopy, err := templateNeedsCopy(fakeDiskSRs{err: readFailed}, "OpaqueRef:target")

	if err == nil {
		t.Fatal("a failed read must be reported")
	}
	if !errors.Is(err, readFailed) {
		t.Errorf("error should wrap the XAPI failure, got: %s", err)
	}
	if !strings.Contains(err.Error(), "SR") {
		t.Errorf("error should say what could not be read, got: %s", err)
	}
	if needsCopy {
		t.Error("a failed read must not be reported as a reason to copy")
	}
}

func TestXapiDiskSRs_ReadsWhatACopyWouldDuplicate(t *testing.T) {
	srOf := map[string]string{
		"OpaqueRef:vdi-disk":   "OpaqueRef:sr-local",
		"OpaqueRef:vdi-floppy": "OpaqueRef:sr-shared",
		"OpaqueRef:vdi-iso":    "OpaqueRef:sr-iso",
	}
	records := map[string]interface{}{
		"OpaqueRef:vbd-disk":   map[string]interface{}{"type": "Disk", "VDI": "OpaqueRef:vdi-disk"},
		"OpaqueRef:vbd-floppy": map[string]interface{}{"type": "Floppy", "VDI": "OpaqueRef:vdi-floppy"},
		"OpaqueRef:vbd-cd":     map[string]interface{}{"type": "CD", "VDI": "OpaqueRef:vdi-iso"},
		// A floppy, not a CD, so that the empty-VBD rule is tested on its own.
		"OpaqueRef:vbd-empty": map[string]interface{}{"type": "Floppy", "empty": true, "VDI": "OpaqueRef:NULL"},
	}

	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_VBDs": []interface{}{
			"OpaqueRef:vbd-disk", "OpaqueRef:vbd-cd", "OpaqueRef:vbd-floppy", "OpaqueRef:vbd-empty",
		},
		"VBD.get_record": func(args []interface{}) interface{} {
			return records[args[0].(string)]
		},
		"VDI.get_SR": func(args []interface{}) interface{} {
			return srOf[args[0].(string)]
		},
	})

	srs, err := xapiDiskSRs{c: fake.conn(), vm: "OpaqueRef:vm"}.SRs()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	want := []xenapi.SRRef{"OpaqueRef:sr-local", "OpaqueRef:sr-shared"}
	if len(srs) != len(want) {
		t.Fatalf("SRs() = %v, want %v", srs, want)
	}
	for i, sr := range srs {
		if sr != want[i] {
			t.Errorf("SRs()[%d] = %q, want %q", i, sr, want[i])
		}
	}
}

func TestXapiDiskSRs_NeverLooksUpACDsISO(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_VBDs": []interface{}{"OpaqueRef:vbd-cd"},
		"VBD.get_record": func([]interface{}) interface{} {
			return map[string]interface{}{"type": "CD", "VDI": "OpaqueRef:vdi-iso"}
		},
		"VDI.get_SR": func([]interface{}) interface{} { return "OpaqueRef:sr-iso" },
	})

	srs, err := xapiDiskSRs{c: fake.conn(), vm: "OpaqueRef:vm"}.SRs()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(srs) != 0 {
		t.Errorf("SRs() = %v, want none", srs)
	}
	if fake.wasCalled("VDI.get_SR") {
		t.Error("a CD's ISO must not be read: a copy leaves it where it is")
	}
}

func TestXapiDiskSRs_ReportsAFailedRead(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_VBDs": errors.New("HANDLE_INVALID"),
	})

	if _, err := (xapiDiskSRs{c: fake.conn(), vm: "OpaqueRef:vm"}).SRs(); err == nil {
		t.Fatal("a failed read must be reported")
	}
}
