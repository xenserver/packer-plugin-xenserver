// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"sort"
	"strings"
	"testing"

	"xenapi"
)

var win11Template = bootModeFacts{
	isHVM:        true,
	firmware:     "uefi",
	declaresVtpm: true,
}

var bootedUEFIVM = bootModeFacts{
	isHVM:     true,
	firmware:  "uefi",
	hasBooted: true,
}

var bootedBIOSVM = bootModeFacts{
	isHVM:     true,
	firmware:  "bios",
	hasBooted: true,
}

func TestCloneSourceRejectsFirmware_RefusesBIOSOnAVTPMSource(t *testing.T) {
	err := cloneSourceRejectsFirmware(TemplateSource("Windows 11"), "bios", win11Template, false)

	if err == nil {
		t.Fatal("XenCenter disables BIOS for a uefi source carrying a vTPM, so this must be refused too")
	}
	msg := strings.ToLower(err.Error())
	for _, want := range []string{"windows 11", "bios", "uefi", "vtpm"} {
		if !strings.Contains(msg, want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
	// Only the clone builder's vtpm = false removes a vTPM that the source declares.
	if !strings.Contains(msg, "vtpm = false") || !strings.Contains(msg, "clone builder") {
		t.Errorf("the message must offer vtpm = false and say where it works, got: %s", err)
	}
}

func TestCloneSourceRejectsFirmware_AllowsBIOSOnAVTPMSourceTheBuildStrips(t *testing.T) {
	if err := cloneSourceRejectsFirmware(TemplateSource("Windows 11"), "bios", win11Template, true); err != nil {
		t.Errorf("vtpm = false removes the vTPM and platform:vtpm before VM.start: %s", err)
	}
}

func TestCloneSourceRejectsFirmware_StillRefusesASettledUEFISourceThatLosesItsVTPM(t *testing.T) {
	settledWithVtpm := bootModeFacts{isHVM: true, firmware: "uefi", hasBooted: true, declaresVtpm: true}

	err := cloneSourceRejectsFirmware(TemplateSource("win11-golden"), "bios", settledWithVtpm, true)

	if err == nil {
		t.Fatal("the settled boot mode is a separate arm, and vtpm = false does not answer it")
	}
	if !strings.Contains(err.Error(), "already settled") {
		t.Errorf("the refusal should be the settled-source one, got: %s", err)
	}
}

func TestCloneSourceRejectsFirmware_RefusesBIOSOnASettledUEFISource(t *testing.T) {
	err := cloneSourceRejectsFirmware(TemplateSource("win11-golden"), "bios", bootedUEFIVM, false)

	if err == nil {
		t.Fatal("a source that has already run under UEFI cannot be re-pointed at BIOS")
	}
	msg := strings.ToLower(err.Error())
	for _, want := range []string{"win11-golden", "bios", "uefi", "already settled"} {
		if !strings.Contains(msg, want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
	if strings.Contains(msg, "vtpm") {
		t.Errorf("the settled-source message is not about the vTPM, got: %s", err)
	}
	if !strings.Contains(msg, "last_boot_cpu_flags") {
		t.Errorf("the message must name the field it read, got: %s", err)
	}
}

func TestCloneSourceRejectsFirmware_RefusesUEFIOnASettledBIOSSource(t *testing.T) {
	err := cloneSourceRejectsFirmware(TemplateSource("ubuntu-golden"), "uefi", bootedBIOSVM, false)

	if err == nil {
		t.Fatal("a source that has already run under BIOS cannot be re-pointed at UEFI")
	}
	msg := strings.ToLower(err.Error())
	for _, want := range []string{"ubuntu-golden", "bios", "uefi", "already settled"} {
		if !strings.Contains(msg, want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
}

func TestCloneSourceRejectsFirmware_ReportsTheSettledBootModeFirst(t *testing.T) {
	src := bootModeFacts{isHVM: true, firmware: "uefi", hasBooted: true, declaresVtpm: true}

	err := cloneSourceRejectsFirmware(TemplateSource("win11-sysprepped"), "bios", src, false)
	if err == nil {
		t.Fatal("a source that is both settled and declares a vTPM must still be refused")
	}
	if !strings.Contains(strings.ToLower(err.Error()), "already settled") {
		t.Errorf("the settled boot mode is the blocker that survives every remedy, got: %s", err)
	}
}

func TestCloneSourceRejectsFirmware_AllowsWhatXenCenterAllows(t *testing.T) {
	cases := map[string]struct {
		firmware string
		src      bootModeFacts
	}{
		"a PV source, which has no firmware to contradict": {
			firmware: "bios", src: bootModeFacts{isHVM: false, firmware: "uefi"}},
		"an HVM source already set to bios": {
			firmware: "bios", src: bootModeFacts{isHVM: true, firmware: "bios"}},
		"an HVM source with no firmware at all": {
			firmware: "bios", src: bootModeFacts{isHVM: true}},
		"a uefi source that is unsettled and has no vTPM": {
			firmware: "bios", src: bootModeFacts{isHVM: true, firmware: "uefi"}},
		"bios on a settled bios source": {
			firmware: "bios", src: bootedBIOSVM},
		"uefi on a settled uefi source": {
			firmware: "uefi", src: bootedUEFIVM},
		"bios on a settled PV source": {
			firmware: "bios", src: bootModeFacts{hasBooted: true}},
		"uefi on a settled PV source": {
			firmware: "uefi", src: bootModeFacts{hasBooted: true}},
		"bios on a bios source declaring a vTPM": {
			firmware: "bios", src: bootModeFacts{isHVM: true, firmware: "bios", declaresVtpm: true}},
		"bios on a PV source declaring a vTPM": {
			firmware: "bios", src: bootModeFacts{declaresVtpm: true}},
		"bios on a PV source whose stale params say uefi and vtpm": {
			firmware: "bios", src: bootModeFacts{firmware: "uefi", declaresVtpm: true}},
		"uefi on a uefi source declaring a vTPM": {
			firmware: "uefi", src: win11Template},
		"uefi on an unsettled HVM source with no firmware at all": {
			firmware: "uefi", src: bootModeFacts{isHVM: true}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if err := cloneSourceRejectsFirmware(TemplateSource("a source"), tc.firmware, tc.src, false); err != nil {
				t.Errorf("unexpected error: %s", err)
			}
		})
	}
}

func TestCloneSourceRejectsFirmware_LeavesAnUnsetFirmwareAlone(t *testing.T) {
	for name, src := range map[string]bootModeFacts{
		"a vTPM template":       win11Template,
		"a settled uefi source": bootedUEFIVM,
		"a settled bios source": bootedBIOSVM,
	} {
		t.Run(name, func(t *testing.T) {
			if err := cloneSourceRejectsFirmware(TemplateSource("a source"), "", src, false); err != nil {
				t.Errorf("unexpected error: %s", err)
			}
		})
	}
}

func TestCloneSourceRejectsFirmware_LeavesValuesXAPIWillRejectAlone(t *testing.T) {
	for _, firmware := range []string{"BIOS", "Bios", "UEFI", "legacy"} {
		t.Run(firmware, func(t *testing.T) {
			if err := cloneSourceRejectsFirmware(TemplateSource("Windows 11"), firmware, win11Template, false); err != nil {
				t.Errorf("%q is not bios or uefi, so it is not an override to refuse: %s", firmware, err)
			}
			if err := cloneSourceRejectsFirmware(TemplateSource("ubuntu-golden"), firmware, bootedBIOSVM, false); err != nil {
				t.Errorf("%q is not bios or uefi, so it is not an override to refuse: %s", firmware, err)
			}
		})
	}
}

func TestBootModeFactsFromRecord_ReadsTheSourceTheWayXenCenterDoes(t *testing.T) {
	cases := map[string]struct {
		record xenapi.VMRecord
		want   bootModeFacts
	}{
		"a Windows 11 template": {
			record: xenapi.VMRecord{
				HVMBootPolicy: "BIOS order",
				HVMBootParams: map[string]string{"firmware": "uefi", "order": "cdn"},
				IsATemplate:   true,
				Platform:      map[string]string{"vtpm": "true", "secureboot": "true"},
			},
			want: bootModeFacts{isHVM: true, firmware: "uefi", declaresVtpm: true},
		},
		"firmware and vtpm in capitals": {
			record: xenapi.VMRecord{
				HVMBootPolicy: "BIOS order",
				HVMBootParams: map[string]string{"firmware": " UEFI "},
				IsATemplate:   true,
				Platform:      map[string]string{"vtpm": "TRUE"},
			},
			want: bootModeFacts{isHVM: true, firmware: "uefi", declaresVtpm: true},
		},
		"a VM that has been started": {
			record: xenapi.VMRecord{
				HVMBootPolicy:    "BIOS order",
				HVMBootParams:    map[string]string{"firmware": "uefi"},
				LastBootCPUFlags: map[string]string{"vendor": "GenuineIntel", "features": "..."},
			},
			want: bootModeFacts{isHVM: true, firmware: "uefi", hasBooted: true},
		},
		"a VM declaring a vTPM is read like a template that does": {
			record: xenapi.VMRecord{
				HVMBootPolicy: "BIOS order",
				HVMBootParams: map[string]string{"firmware": "uefi"},
				Platform:      map[string]string{"vtpm": "true"},
			},
			want: bootModeFacts{isHVM: true, firmware: "uefi", declaresVtpm: true},
		},
		// XAPI's Vm_platform.is_true accepts "1" too.
		"platform saying 1": {
			record: xenapi.VMRecord{
				HVMBootPolicy: "BIOS order",
				HVMBootParams: map[string]string{"firmware": "uefi"},
				IsATemplate:   true,
				Platform:      map[string]string{"vtpm": "1"},
			},
			want: bootModeFacts{isHVM: true, firmware: "uefi", declaresVtpm: true},
		},
		"platform saying anything but true": {
			record: xenapi.VMRecord{
				HVMBootPolicy: "BIOS order",
				HVMBootParams: map[string]string{"firmware": "uefi"},
				IsATemplate:   true,
				Platform:      map[string]string{"vtpm": "false"},
			},
			want: bootModeFacts{isHVM: true, firmware: "uefi"},
		},
		"a template carrying no firmware at all": {
			record: xenapi.VMRecord{HVMBootPolicy: "BIOS order", IsATemplate: true},
			want:   bootModeFacts{isHVM: true},
		},
		"a PV template": {
			record: xenapi.VMRecord{PVBootloader: "pygrub", IsATemplate: true},
			want:   bootModeFacts{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := bootModeFactsFromRecord(tc.record); got != tc.want {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestEnsureCloneSourceIsHalted_RefusesEveryStateButHalted(t *testing.T) {
	for _, state := range []xenapi.VMPowerState{
		xenapi.VMPowerStateRunning,
		xenapi.VMPowerStatePaused,
		xenapi.VMPowerStateSuspended,
	} {
		t.Run(string(state), func(t *testing.T) {
			err := EnsureCloneSourceIsHalted(VMSource("build-src"), xenapi.VMRecord{PowerState: state}, false)

			if err == nil {
				t.Fatalf("a %s source cannot be cloned", state)
			}
			for _, want := range []string{"clone_vm 'build-src'", string(state), "snapshot"} {
				if !strings.Contains(err.Error(), want) {
					t.Errorf("error should mention %q, got: %s", want, err)
				}
			}
		})
	}
}

func TestEnsureCloneSourceIsHalted_AllowsAHaltedVM(t *testing.T) {
	err := EnsureCloneSourceIsHalted(VMSource("build-src"),
		xenapi.VMRecord{PowerState: xenapi.VMPowerStateHalted}, false)
	if err != nil {
		t.Errorf("halted is what both calls require: %s", err)
	}
}

func TestEnsureCloneSourceIsHalted_SkipsATemplate(t *testing.T) {
	err := EnsureCloneSourceIsHalted(TemplateSource("Windows 11"),
		xenapi.VMRecord{IsATemplate: true, PowerState: xenapi.VMPowerStateRunning}, true)
	if err != nil {
		t.Errorf("a template is not subject to the power-state check: %s", err)
	}
}

// VM.clone shuts the clone of a checkpoint down itself, so the build still gets a halted VM.
func TestEnsureCloneSourceIsHalted_SkipsASuspendedCheckpoint(t *testing.T) {
	err := EnsureCloneSourceIsHalted(VMSource("nightly-checkpoint"),
		xenapi.VMRecord{IsATemplate: true, IsASnapshot: true, PowerState: xenapi.VMPowerStateSuspended}, false)
	if err != nil {
		t.Errorf("a snapshot is not subject to the power-state check: %s", err)
	}
}

func TestEnsureCloneSourceIsHalted_RefusesASuspendedCheckpointWhenTheBuildMayCopy(t *testing.T) {
	err := EnsureCloneSourceIsHalted(VMSource("nightly-checkpoint"),
		xenapi.VMRecord{IsATemplate: true, IsASnapshot: true, PowerState: xenapi.VMPowerStateSuspended}, true)
	if err == nil {
		t.Fatal("VM.copy cannot finish on a suspended snapshot")
	}
	for _, want := range []string{"clone_vm 'nightly-checkpoint'", "vm_sr", "VM_IS_TEMPLATE"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
}

func TestEnsureCloneSourceIsHalted_AllowsAHaltedSnapshotEvenWhenTheBuildMayCopy(t *testing.T) {
	err := EnsureCloneSourceIsHalted(VMSource("nightly"),
		xenapi.VMRecord{IsATemplate: true, IsASnapshot: true, PowerState: xenapi.VMPowerStateHalted}, true)
	if err != nil {
		t.Errorf("a halted snapshot needs no hard_shutdown after the copy: %s", err)
	}
}

func TestCloneSourceMatches_SeparatesTemplatesVMsAndSnapshots(t *testing.T) {
	realTemplate := xenapi.VMRecord{IsATemplate: true}
	realVM := xenapi.VMRecord{}
	snapshot := xenapi.VMRecord{IsATemplate: true, IsASnapshot: true}
	dom0 := xenapi.VMRecord{IsControlDomain: true}

	cases := map[string]struct {
		src    CloneSource
		record xenapi.VMRecord
		want   bool
	}{
		"clone_template takes a template":         {TemplateSource("x"), realTemplate, true},
		"clone_template refuses a VM":             {TemplateSource("x"), realVM, false},
		"clone_template refuses a snapshot":       {TemplateSource("x"), snapshot, false},
		"clone_vm takes a VM":                     {VMSource("x"), realVM, true},
		"clone_vm takes a snapshot":               {VMSource("x"), snapshot, true},
		"clone_vm refuses a template":             {VMSource("x"), realTemplate, false},
		"clone_vm refuses a control domain":       {VMSource("x"), dom0, false},
		"clone_template refuses a control domain": {TemplateSource("x"), dom0, false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if got := tc.src.matches(tc.record); got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestKindOf_NamesASnapshotAsItself(t *testing.T) {
	cases := map[string]xenapi.VMRecord{
		"template":       {IsATemplate: true},
		"VM":             {},
		"snapshot":       {IsATemplate: true, IsASnapshot: true},
		"control domain": {IsControlDomain: true},
	}

	for want, record := range cases {
		if got := kindOf(record); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestCloneSource_NamesTheOptionTheBuildSet(t *testing.T) {
	if got := VMSource("build-src").String(); got != "clone_vm 'build-src'" {
		t.Errorf("got %q", got)
	}
	if got := TemplateSource("Windows 11").String(); got != "clone_template 'Windows 11'" {
		t.Errorf("got %q", got)
	}
	if got := VMSource("x").alternative; got != "clone_template" {
		t.Errorf("the other option for clone_vm is clone_template, got %q", got)
	}
	if got := TemplateSource("x").alternative; got != "clone_vm" {
		t.Errorf("the other option for clone_template is clone_vm, got %q", got)
	}
}

func resolverFake(t *testing.T, records map[string]map[string]interface{}) *fakeXAPI {
	t.Helper()

	refs := make([]interface{}, 0, len(records))
	for ref := range records {
		refs = append(refs, ref)
	}
	sort.Slice(refs, func(i, j int) bool { return refs[i].(string) < refs[j].(string) })

	return newFakeXAPI(t, map[string]interface{}{
		"VM.get_by_name_label": refs,
		"VM.get_record": func(args []interface{}) interface{} {
			return records[args[0].(string)]
		},
	})
}

func TestResolveCloneSource_PicksTheObjectOfTheOptionsOwnKind(t *testing.T) {
	records := map[string]map[string]interface{}{
		"OpaqueRef:the-template": {"is_a_template": true, "power_state": "Halted"},
		"OpaqueRef:the-vm":       {"power_state": "Running"},
		"OpaqueRef:the-snapshot": {"is_a_template": true, "is_a_snapshot": true, "power_state": "Suspended"},
	}

	isoPath := ISOTemplateSource("golden")

	for name, src := range map[string]CloneSource{
		"clone_template":                 TemplateSource("golden"),
		"clone_template on the iso path": isoPath,
	} {
		t.Run(name, func(t *testing.T) {
			ref, record, err := ResolveCloneSource(resolverFake(t, records).conn(), src)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if ref != "OpaqueRef:the-template" {
				t.Errorf("ref = %q, want the template", ref)
			}
			if !record.IsATemplate || record.IsASnapshot {
				t.Errorf("record is not the template's: %+v", record)
			}
		})
	}

	t.Run("clone_vm prefers the VM", func(t *testing.T) {
		vmOnly := map[string]map[string]interface{}{
			"OpaqueRef:the-template": records["OpaqueRef:the-template"],
			"OpaqueRef:the-vm":       records["OpaqueRef:the-vm"],
		}
		ref, record, err := ResolveCloneSource(resolverFake(t, vmOnly).conn(), VMSource("golden"))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if ref != "OpaqueRef:the-vm" {
			t.Errorf("ref = %q, want the VM", ref)
		}
		if record.PowerState != xenapi.VMPowerStateRunning {
			t.Errorf("the record must be the resolved object's, got power state %q", record.PowerState)
		}
	})

	t.Run("clone_vm takes the snapshot", func(t *testing.T) {
		snapshotOnly := map[string]map[string]interface{}{
			"OpaqueRef:the-template": records["OpaqueRef:the-template"],
			"OpaqueRef:the-snapshot": records["OpaqueRef:the-snapshot"],
		}
		ref, _, err := ResolveCloneSource(resolverFake(t, snapshotOnly).conn(), VMSource("golden"))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if ref != "OpaqueRef:the-snapshot" {
			t.Errorf("ref = %q, want the snapshot", ref)
		}
	})
}

func TestResolveCloneSource_ReturnsTheResolvedObjectsOwnRecord(t *testing.T) {
	fake := resolverFake(t, map[string]map[string]interface{}{
		"OpaqueRef:win11": {
			"is_a_template":   true,
			"HVM_boot_policy": "BIOS order",
			"HVM_boot_params": map[string]interface{}{"firmware": "uefi"},
			"platform":        map[string]interface{}{"vtpm": "true"},
		},
	})

	_, record, err := ResolveCloneSource(fake.conn(), TemplateSource("Windows 11"))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	facts := bootModeFactsFromRecord(record)
	if !facts.isHVM || facts.firmware != firmwareUEFI || !facts.declaresVtpm {
		t.Errorf("the record did not carry the source's boot mode: %+v", facts)
	}
}

func TestResolveCloneSource_AsksForTheNameTheOptionNamed(t *testing.T) {
	fake := resolverFake(t, map[string]map[string]interface{}{
		"OpaqueRef:x": {"is_a_template": true},
	})

	if _, _, err := ResolveCloneSource(fake.conn(), TemplateSource("Windows 11")); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	args := fake.argsOf(t, "VM.get_by_name_label")
	if len(args) != 1 || args[0] != "Windows 11" {
		t.Errorf("looked up %v, want the name clone_template gave", args)
	}
}

func TestResolveCloneSource_RefusesWhatItCannotResolveToOneObject(t *testing.T) {
	template := map[string]interface{}{"is_a_template": true}

	cases := map[string]struct {
		src     CloneSource
		records map[string]map[string]interface{}
		want    []string
	}{
		"nothing of that name": {
			TemplateSource("missing"), nil,
			[]string{"could not find a template named 'missing'"},
		},
		"two of the same kind": {
			TemplateSource("golden"),
			map[string]map[string]interface{}{"OpaqueRef:a": template, "OpaqueRef:b": template},
			[]string{"found 2 templates named 'golden'", "unique among templates"},
		},
		"only the other kind": {
			TemplateSource("build-src"),
			map[string]map[string]interface{}{"OpaqueRef:vm": {}},
			[]string{"clone_template 'build-src'", "names a VM, not a template", "Use clone_vm instead"},
		},
		"a snapshot named in clone_template": {
			TemplateSource("nightly"),
			map[string]map[string]interface{}{
				"OpaqueRef:snap": {"is_a_template": true, "is_a_snapshot": true},
			},
			[]string{"names a snapshot, not a template", "Use clone_vm instead"},
		},
		"a control domain named in clone_vm": {
			VMSource("Control domain on host: host1"),
			map[string]map[string]interface{}{"OpaqueRef:dom0": {"is_control_domain": true}},
			[]string{"names a control domain", "no source here to build from"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, _, err := ResolveCloneSource(resolverFake(t, tc.records).conn(), tc.src)
			if err == nil {
				t.Fatal("expected an error naming what went wrong")
			}
			for _, want := range tc.want {
				if !strings.Contains(err.Error(), want) {
					t.Errorf("error should mention %q, got: %s", want, err)
				}
			}
		})
	}
}

func TestResolveCloneSource_ALookupFailureIsFatal(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_by_name_label": errors.New("SESSION_INVALID"),
	})

	if _, _, err := ResolveCloneSource(fake.conn(), TemplateSource("golden")); err == nil {
		t.Fatal("expected the lookup failure to be reported")
	} else if !strings.Contains(err.Error(), "SESSION_INVALID") {
		t.Errorf("error should carry XAPI's own: %s", err)
	}
}

func TestResolveCloneSource_ARecordFailureIsFatal(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_by_name_label": []interface{}{"OpaqueRef:x"},
		"VM.get_record":        errors.New("HANDLE_INVALID"),
	})

	if _, _, err := ResolveCloneSource(fake.conn(), TemplateSource("golden")); err == nil {
		t.Fatal("expected the record read to be reported")
	} else if !strings.Contains(err.Error(), "HANDLE_INVALID") {
		t.Errorf("error should carry XAPI's own: %s", err)
	}
}
