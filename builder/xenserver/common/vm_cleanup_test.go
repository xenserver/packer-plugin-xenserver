// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"

	"xenapi"
)

func cleanupState(t *testing.T, results map[string]interface{}) (*multistep.BasicStateBag, *recordingUi, *fakeXAPI) {
	t.Helper()

	backoff := vdiDestroyBackoff
	vdiDestroyBackoff = 0
	t.Cleanup(func() { vdiDestroyBackoff = backoff })

	answers := map[string]interface{}{
		"VM.get_by_uuid":     "OpaqueRef:vm",
		"VM.get_VBDs":        []interface{}{"OpaqueRef:vbd"},
		"VBD.get_record":     map[string]interface{}{"type": "Disk", "VDI": "OpaqueRef:vdi"},
		"VM.get_power_state": "Halted",
		"VM.destroy":         "",
		"VBD.unplug":         "",
		"VDI.get_uuid":       "vdi-uuid",
		"VDI.destroy":        "",
	}
	for method, result := range results {
		answers[method] = result
	}

	fake := newFakeXAPI(t, answers)
	ui := &recordingUi{}
	state := new(multistep.BasicStateBag)
	state.Put("ui", ui)
	state.Put("client", fake.conn())
	state.Put("commonconfig", CommonConfig{KeepVM: "never"})
	state.Put("instance_uuid", "vm-uuid")

	return state, ui, fake
}

func TestVmCleanup_DestroysTheVmAndTheDisksTheCloneDuplicated(t *testing.T) {
	state, ui, fake := cleanupState(t, nil)

	new(VmCleanup).Cleanup(state)

	if args := fake.argsOf(t, "VDI.destroy"); len(args) != 1 || args[0] != "OpaqueRef:vdi" {
		t.Errorf("VDI.destroy%v, want the VM's disk", args)
	}
	if !strings.Contains(ui.text(), "VM and disk cleanup completed") {
		t.Errorf("a clean teardown said nothing: %s", ui.text())
	}
}

func TestVmCleanup_SaysWhatItLeftBehind(t *testing.T) {
	state, ui, fake := cleanupState(t, map[string]interface{}{
		"VDI.destroy": errors.New("VDI_IN_USE"),
	})

	new(VmCleanup).Cleanup(state)

	if tries := len(fake.allArgsOf("VDI.destroy")); tries != vdiDestroyAttempts {
		t.Errorf("VDI.destroy tried %d times, want %d", tries, vdiDestroyAttempts)
	}
	if !strings.Contains(ui.text(), "Cleanup left 1 object(s)") {
		t.Errorf("a disk left on the SR was not reported: %s", ui.text())
	}
	if strings.Contains(ui.text(), "VM and disk cleanup completed") {
		t.Errorf("a teardown that left a disk behind reported success: %s", ui.text())
	}
}

func TestVmCleanup_SaysSoWhenItCannotListTheDisks(t *testing.T) {
	state, ui, fake := cleanupState(t, map[string]interface{}{
		"VM.get_VBDs": errors.New("SESSION_INVALID"),
	})

	new(VmCleanup).Cleanup(state)

	if !fake.wasCalled("VM.destroy") {
		t.Error("the VM was left behind as well, over a disk list it could do without")
	}
	if strings.Contains(ui.text(), "VM and disk cleanup completed") {
		t.Errorf("a teardown that never saw the disks reported success: %s", ui.text())
	}
}

func preCleanupState(t *testing.T, name string, results map[string]interface{}) (*multistep.BasicStateBag, *fakeXAPI) {
	t.Helper()

	fake := newFakeXAPI(t, results)
	state := new(multistep.BasicStateBag)
	state.Put("ui", &recordingUi{})
	state.Put("client", fake.conn())
	state.Put("commonconfig", CommonConfig{VMName: &name})

	return state, fake
}

func TestPreCleanup_RefusesWhenVmNameIsTheSource(t *testing.T) {
	for _, force := range []bool{false, true} {
		t.Run(fmt.Sprintf("force=%v", force), func(t *testing.T) {
			state, fake := preCleanupState(t, "golden", map[string]interface{}{
				"VM.get_by_name_label": []interface{}{"OpaqueRef:source"},
			})

			err := PreCleanup(state, force, TemplateSource("golden"), xenapi.VMRef("OpaqueRef:source"))

			if err == nil {
				t.Fatal("PreCleanup accepted a vm_name naming the build's own source")
			}
			if !strings.Contains(err.Error(), "clone_template 'golden'") {
				t.Errorf("error does not name the source option: %s", err)
			}
			if strings.Contains(err.Error(), "-force flag to destroy it") {
				t.Errorf("error tells the user to -force away their own source: %s", err)
			}
			if fake.wasCalled("VM.destroy") {
				t.Error("the source was destroyed")
			}
		})
	}
}

func TestPreCleanup_DestroysAClashOnlyWithForce(t *testing.T) {
	results := func() map[string]interface{} {
		return map[string]interface{}{
			"VM.get_by_name_label": []interface{}{"OpaqueRef:stale"},
			"VM.get_is_a_template": false,
			"VM.get_VBDs":          []interface{}{},
			"VM.get_power_state":   "Halted",
			"VM.destroy":           "",
		}
	}

	t.Run("force=false", func(t *testing.T) {
		state, fake := preCleanupState(t, "packer-vm", results())

		err := PreCleanup(state, false, TemplateSource("golden"), xenapi.VMRef("OpaqueRef:source"))

		if err == nil || !strings.Contains(err.Error(), "already exists") {
			t.Fatalf("err = %v, want the name clash reported", err)
		}
		if fake.wasCalled("VM.destroy") {
			t.Error("a clash was destroyed without -force")
		}
	})

	t.Run("force=true", func(t *testing.T) {
		state, fake := preCleanupState(t, "packer-vm", results())

		if err := PreCleanup(state, true, TemplateSource("golden"), xenapi.VMRef("OpaqueRef:source")); err != nil {
			t.Fatalf("PreCleanup: %s", err)
		}

		args := fake.argsOf(t, "VM.destroy")
		if len(args) != 1 || args[0] != "OpaqueRef:stale" {
			t.Errorf("VM.destroy%v, want the clashing VM", args)
		}
	})
}

func TestPreCleanup_NoClashDestroysNothing(t *testing.T) {
	state, fake := preCleanupState(t, "packer-vm", map[string]interface{}{
		"VM.get_by_name_label": []interface{}{},
	})

	if err := PreCleanup(state, true, TemplateSource("golden"), xenapi.VMRef("OpaqueRef:source")); err != nil {
		t.Fatalf("PreCleanup: %s", err)
	}
	if fake.wasCalled("VM.destroy") {
		t.Error("something was destroyed with no name clash")
	}
}
