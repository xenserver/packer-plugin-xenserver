// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func uploadVdiState(t *testing.T, f *fakeXAPI, config CommonConfig) (*multistep.BasicStateBag, *packersdk.MockUi, string) {
	t.Helper()

	imagePath := filepath.Join(t.TempDir(), "cddata.iso")
	if err := os.WriteFile(imagePath, []byte("not really an iso"), 0o600); err != nil {
		t.Fatal(err)
	}

	ui := &packersdk.MockUi{}
	state := new(multistep.BasicStateBag)
	state.Put("ui", ui)
	state.Put("client", f.conn())
	state.Put("commonconfig", config)

	return state, ui, imagePath
}

// Run halts at the upload, which the fake cannot serve, but only the SR chosen before it matters.
func uploadedTo(t *testing.T, f *fakeXAPI, config CommonConfig) string {
	t.Helper()

	state, _, imagePath := uploadVdiState(t, f, config)
	step := &StepUploadVdi{
		VdiNameFunc:   func() string { return "Packer-cd-disk" },
		ImagePathFunc: func() string { return imagePath },
		VdiUuidKey:    "cd_vdi_uuid",
	}
	step.Run(context.Background(), state)

	record, ok := f.argsOf(t, "VDI.create")[0].(map[string]interface{})
	if !ok {
		t.Fatal("VDI.create was not passed a record")
	}
	if readOnly, _ := record["read_only"].(bool); readOnly {
		t.Error("the uploaded VDI is read-only, so the SR it needs is not the one this test pins")
	}

	sr, _ := record["SR"].(string)

	return sr
}

func TestStepUploadVdi_CreatesTheVdiOnTheVMSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"SR.get_by_name_label": []interface{}{"OpaqueRef:vm-sr"},
		"VDI.create":           "OpaqueRef:vdi",
		"VDI.get_uuid":         "cd-vdi-uuid",
	})

	if sr := uploadedTo(t, f, CommonConfig{VMSR: "Local storage", ISOSR: "ISO library"}); sr != "OpaqueRef:vm-sr" {
		t.Errorf("the VDI was created on %q, want OpaqueRef:vm-sr", sr)
	}
	if args := f.argsOf(t, "SR.get_by_name_label"); len(args) != 1 || args[0] != "Local storage" {
		t.Errorf("resolved SR %v, want [Local storage]", args)
	}
}

func TestStepUploadVdi_FallsBackToThePoolDefaultSR(t *testing.T) {
	results := poolWith("OpaqueRef:default-sr")
	results["VDI.create"] = "OpaqueRef:vdi"
	results["VDI.get_uuid"] = "cd-vdi-uuid"
	f := newFakeXAPI(t, results)

	if sr := uploadedTo(t, f, CommonConfig{ISOSR: "ISO library"}); sr != "OpaqueRef:default-sr" {
		t.Errorf("the VDI was created on %q, want OpaqueRef:default-sr", sr)
	}
	if f.wasCalled("SR.get_by_name_label") {
		t.Error("a name-label was resolved despite vm_sr being unset")
	}
}

func TestStepUploadVdi_SkipsAnUnsetImage(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{})

	state, _, _ := uploadVdiState(t, f, CommonConfig{})
	step := &StepUploadVdi{
		VdiNameFunc:   func() string { return "Packer-cd-disk" },
		ImagePathFunc: func() string { return "" },
		VdiUuidKey:    "cd_vdi_uuid",
	}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatal("expected a build without cd_files to skip the upload")
	}
	if called := f.methodsCalled(); len(called) != 0 {
		t.Errorf("expected no XAPI call, got %v", called)
	}
}

func TestStepUploadVdi_ReportsAnUnresolvableSR(t *testing.T) {
	f := newFakeXAPI(t, map[string]interface{}{
		"SR.get_by_name_label": []interface{}{},
	})

	state, ui, imagePath := uploadVdiState(t, f, CommonConfig{VMSR: "No such SR"})
	step := &StepUploadVdi{
		VdiNameFunc:   func() string { return "Packer-cd-disk" },
		ImagePathFunc: func() string { return imagePath },
		VdiUuidKey:    "cd_vdi_uuid",
	}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatal("expected the step to halt when the SR cannot be resolved")
	}
	if !strings.Contains(ui.ErrorMessage, "No such SR") {
		t.Errorf("error %q does not quote the SR name that could not be resolved", ui.ErrorMessage)
	}
	if f.wasCalled("VDI.create") {
		t.Error("a VDI was created despite the SR lookup failing")
	}
}

func cleanupUploaded(t *testing.T, config CommonConfig, vbds []interface{}) *fakeXAPI {
	t.Helper()

	f := newFakeXAPI(t, map[string]interface{}{
		"VDI.get_by_uuid": "OpaqueRef:cd-vdi",
		"VDI.get_VBDs":    vbds,
		"VDI.destroy":     "",
	})

	state, _, _ := uploadVdiState(t, f, config)
	state.Put("cd_vdi_uuid", "cd-vdi-uuid")

	step := &StepUploadVdi{
		VdiNameFunc:   func() string { return "Packer-cd-disk" },
		ImagePathFunc: func() string { return "" },
		VdiUuidKey:    "cd_vdi_uuid",
	}
	step.Cleanup(state)

	return f
}

func TestStepUploadVdi_DestroysADetachedVdiOnABuildThatKeepsTheVM(t *testing.T) {
	f := cleanupUploaded(t, CommonConfig{Format: "none", KeepVM: "always"}, []interface{}{})

	if !f.wasCalled("VDI.destroy") {
		t.Error("the uploaded VDI is attached to nothing and was left on the SR")
	}
}

func TestStepUploadVdi_KeepsAVdiThatIsStillAttached(t *testing.T) {
	f := cleanupUploaded(t, CommonConfig{KeepVM: "always"}, []interface{}{"OpaqueRef:cd-vbd"})

	if f.wasCalled("VDI.destroy") {
		t.Error("a VDI still plugged into a VM the build is keeping was destroyed")
	}
}
