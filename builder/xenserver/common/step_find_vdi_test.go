// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func findVdiState(t *testing.T, f *fakeXAPI) (*multistep.BasicStateBag, *packersdk.MockUi) {
	t.Helper()

	ui := &packersdk.MockUi{}
	state := new(multistep.BasicStateBag)
	state.Put("ui", ui)
	state.Put("client", f.conn())

	return state, ui
}

func vdiSRs(srOf map[string]string) func([]interface{}) interface{} {
	return func(params []interface{}) interface{} {
		return srOf[params[0].(string)]
	}
}

func TestStepFindVdi_PicksTheVdiOnTheNamedSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"VDI.get_by_name_label": []interface{}{"OpaqueRef:vdi-elsewhere", "OpaqueRef:vdi-on-iso-sr"},
		"SR.get_by_name_label":  []interface{}{"OpaqueRef:iso-sr"},
		"VDI.get_SR": vdiSRs(map[string]string{
			"OpaqueRef:vdi-elsewhere": "OpaqueRef:other-sr",
			"OpaqueRef:vdi-on-iso-sr": "OpaqueRef:iso-sr",
		}),
		"VDI.get_uuid": "the-right-uuid",
	})

	state, ui := findVdiState(t, f)
	step := &StepFindVdi{VdiName: "Win11.iso", SRName: "ISO library", VdiUuidKey: "isoname_vdi_uuid"}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("step halted: %s", ui.ErrorMessage)
	}

	if uuid, _ := state.Get("isoname_vdi_uuid").(string); uuid != "the-right-uuid" {
		t.Errorf("got VDI uuid %q, want the-right-uuid", uuid)
	}
	if args := f.argsOf(t, "SR.get_by_name_label"); len(args) != 1 || args[0] != "ISO library" {
		t.Errorf("resolved SR %v, want [ISO library]", args)
	}
	if args := f.argsOf(t, "VDI.get_uuid"); len(args) != 1 || args[0] != "OpaqueRef:vdi-on-iso-sr" {
		t.Errorf("took the uuid of %v, want [OpaqueRef:vdi-on-iso-sr]", args)
	}
}

func TestStepFindVdi_RejectsAVdiOnAnotherSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"VDI.get_by_name_label": []interface{}{"OpaqueRef:vdi-elsewhere"},
		"SR.get_by_name_label":  []interface{}{"OpaqueRef:iso-sr"},
		"VDI.get_SR":            vdiSRs(map[string]string{"OpaqueRef:vdi-elsewhere": "OpaqueRef:other-sr"}),
		"VDI.get_uuid":          "must-not-be-used",
	})

	state, ui := findVdiState(t, f)
	step := &StepFindVdi{VdiName: "Win11.iso", SRName: "ISO library", VdiUuidKey: "isoname_vdi_uuid"}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatal("expected the step to halt for a VDI on another SR")
	}
	if !strings.Contains(ui.ErrorMessage, "Couldn't find a VDI") || !strings.Contains(ui.ErrorMessage, "ISO library") {
		t.Errorf("error %q does not report the VDI as missing from the SR that was searched", ui.ErrorMessage)
	}
	if _, ok := state.GetOk("isoname_vdi_uuid"); ok {
		t.Error("a VDI uuid was stored despite the halt")
	}
}

func TestStepFindVdi_ReportsTwoVdisOnTheSameSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"VDI.get_by_name_label": []interface{}{"OpaqueRef:one", "OpaqueRef:two"},
		"SR.get_by_name_label":  []interface{}{"OpaqueRef:iso-sr"},
		"VDI.get_SR": vdiSRs(map[string]string{
			"OpaqueRef:one": "OpaqueRef:iso-sr",
			"OpaqueRef:two": "OpaqueRef:iso-sr",
		}),
	})

	state, ui := findVdiState(t, f)
	step := &StepFindVdi{VdiName: "Win11.iso", SRName: "ISO library", VdiUuidKey: "isoname_vdi_uuid"}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatal("expected the step to halt for two VDIs of the same name")
	}
	if !strings.Contains(ui.ErrorMessage, "more than one") || !strings.Contains(ui.ErrorMessage, "ISO library") {
		t.Errorf("error %q does not report an ambiguous name within the SR", ui.ErrorMessage)
	}
}

func TestStepFindVdi_SearchesThePoolWithoutAnSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"VDI.get_by_name_label": []interface{}{"OpaqueRef:tools"},
		"VDI.get_uuid":          "tools-uuid",
	})

	state, ui := findVdiState(t, f)
	step := &StepFindVdi{VdiName: "guest-tools.iso", VdiUuidKey: "tools_vdi_uuid"}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("step halted: %s", ui.ErrorMessage)
	}
	if uuid, _ := state.Get("tools_vdi_uuid").(string); uuid != "tools-uuid" {
		t.Errorf("got VDI uuid %q, want tools-uuid", uuid)
	}
	if f.wasCalled("SR.get_by_name_label") {
		t.Error("an SR was resolved for a pool-wide lookup")
	}
}

func TestStepFindVdi_ReportsALookupFailure(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"VDI.get_by_name_label": errors.New("SESSION_INVALID"),
	})

	state, ui := findVdiState(t, f)
	step := &StepFindVdi{VdiName: "Win11.iso", SRName: "ISO library", VdiUuidKey: "isoname_vdi_uuid"}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatal("expected the step to halt when the lookup fails")
	}
	if !strings.Contains(ui.ErrorMessage, "SESSION_INVALID") {
		t.Errorf("error %q hides the reason the lookup failed", ui.ErrorMessage)
	}
	if strings.Contains(ui.ErrorMessage, "Couldn't find") {
		t.Errorf("error %q reports a failed lookup as a missing ISO", ui.ErrorMessage)
	}
}

func TestStepFindVdi_IgnoresAnUnsetName(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{})

	state, _ := findVdiState(t, f)
	step := &StepFindVdi{VdiName: "", SRName: "ISO library", VdiUuidKey: "isoname_vdi_uuid"}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatal("expected an unset iso_name to be skipped")
	}
	if called := f.methodsCalled(); len(called) != 0 {
		t.Errorf("expected no XAPI call, got %v", called)
	}
}
