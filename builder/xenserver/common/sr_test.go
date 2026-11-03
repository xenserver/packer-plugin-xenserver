// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"strings"
	"testing"

	"xenapi"
)

func poolWith(defaultSR string) map[string]interface{} {
	return map[string]interface{}{
		"session.get_this_host": "OpaqueRef:this-host",
		"pool.get_all_records": map[string]interface{}{
			"OpaqueRef:pool": map[string]interface{}{
				"master":     "OpaqueRef:this-host",
				"default_SR": defaultSR,
			},
		},
	}
}

func TestGetSR_ResolvesTheVMSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"SR.get_by_name_label": []interface{}{"OpaqueRef:vm-sr"},
	})

	config := CommonConfig{VMSR: "Local storage", ISOSR: "ISO library"}

	sr, err := config.GetSR(f.conn())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if sr != xenapi.SRRef("OpaqueRef:vm-sr") {
		t.Errorf("got SR %q, want OpaqueRef:vm-sr", sr)
	}

	// The fake returns this SR for any name, so only the name shows that iso_sr was not used.
	args := f.argsOf(t, "SR.get_by_name_label")
	if len(args) != 1 || args[0] != "Local storage" {
		t.Errorf("looked up %v, want [Local storage]", args)
	}
}

func TestGetSR_FallsBackToThePoolDefaultSR(t *testing.T) {
	// Resolvable, so wrongly resolving iso_sr shows up as the wrong SR rather than as an error.
	results := poolWith("OpaqueRef:default-sr")
	results["SR.get_by_name_label"] = []interface{}{"OpaqueRef:iso-sr"}
	f := newFakeXAPI(t, results)

	sr, err := CommonConfig{ISOSR: "ISO library"}.GetSR(f.conn())
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if sr != xenapi.SRRef("OpaqueRef:default-sr") {
		t.Errorf("got SR %q, want OpaqueRef:default-sr", sr)
	}
	if f.wasCalled("SR.get_by_name_label") {
		t.Error("a name-label was resolved despite vm_sr being unset")
	}
}

func TestGetSR_ReportsAPoolWithNoDefaultSR(t *testing.T) {
	for _, nullSR := range []string{"", nullRef} {
		f := newFakeXAPI(t, poolWith(nullSR))

		sr, err := CommonConfig{}.GetSR(f.conn())
		if err == nil {
			t.Fatalf("default_SR %q: expected an error, got SR %q", nullSR, sr)
		}
		if !strings.Contains(err.Error(), "vm_sr") {
			t.Errorf("default_SR %q: error %q does not name the option that fixes it", nullSR, err)
		}
	}
}

func TestFindSRByNameLabel_ReportsANameThatMatchesNothing(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"SR.get_by_name_label": []interface{}{},
	})

	_, err := findSRByNameLabel(f.conn(), "No such SR")
	if err == nil {
		t.Fatal("expected an error for a name no SR has")
	}
	if !strings.Contains(err.Error(), "No such SR") {
		t.Errorf("error %q does not quote the name that was not found", err)
	}
}

func TestFindSRByNameLabel_ReportsAnAmbiguousName(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"SR.get_by_name_label": []interface{}{"OpaqueRef:one", "OpaqueRef:two"},
	})

	_, err := findSRByNameLabel(f.conn(), "Local storage")
	if err == nil {
		t.Fatal("expected an error for a name two SRs share")
	}
	if !strings.Contains(err.Error(), "more than one") {
		t.Errorf("error %q does not say the name is ambiguous", err)
	}
}
