// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"
)

type fakeBootCheck struct {
	firmware    string
	firmwareErr error
	vtpms       int
	vtpmErr     error
	reads       int
}

func (f *fakeBootCheck) Firmware() (string, error) {
	f.reads++
	if f.firmwareErr != nil {
		return "", f.firmwareErr
	}
	return f.firmware, nil
}

func (f *fakeBootCheck) VTPMCount() (int, error) {
	if f.vtpmErr != nil {
		return 0, f.vtpmErr
	}
	return f.vtpms, nil
}

func TestEnsureVMCanStart_RejectsABIOSVMThatInheritedAVTPM(t *testing.T) {
	api := &fakeBootCheck{vtpms: 1}

	firmware, err := ensureVMCanStart(api, "bios")

	if err == nil {
		t.Fatal("a BIOS VM with a vTPM cannot start, so the build must be stopped here")
	}
	if firmware != "bios" {
		t.Errorf("firmware = %q, want the value the VM will boot with reported even on failure", firmware)
	}
	msg := strings.ToLower(err.Error())
	for _, want := range []string{"vtpm", "firmware", "uefi", "source"} {
		if !strings.Contains(msg, want) {
			t.Errorf("error should mention %q, got: %s", want, err)
		}
	}
	if !strings.Contains(msg, "remove firmware") {
		t.Errorf("a build that set firmware should be told it can unset it, got: %s", err)
	}
	if !strings.Contains(msg, "vtpm = false") || !strings.Contains(msg, "clone builder") {
		t.Errorf("the message must name vtpm = false and which builder it works on, got: %s", err)
	}
	if strings.Contains(msg, "platform:vtpm") {
		t.Errorf("the message must not still claim platform:vtpm defeats the removal, got: %s", err)
	}
}

func TestEnsureVMCanStart_ChecksTheInheritedFirmwareWhenNoneIsConfigured(t *testing.T) {
	api := &fakeBootCheck{firmware: "", vtpms: 1}

	firmware, err := ensureVMCanStart(api, "")

	if err == nil {
		t.Fatal("the template is BIOS and has a vTPM, so this VM cannot start either")
	}
	if firmware != "bios" {
		t.Errorf("firmware = %q, want an absent boot param reported as \"bios\"", firmware)
	}
	if !strings.Contains(err.Error(), "inherited") {
		t.Errorf("the build set neither option, so the error has to say where both came from, got: %s", err)
	}
}

func TestEnsureVMCanStart_AllowsAUEFITemplateWithAVTPM(t *testing.T) {
	api := &fakeBootCheck{firmware: "uefi", vtpms: 1}

	firmware, err := ensureVMCanStart(api, "")

	if err != nil {
		t.Fatalf("a UEFI VM with a vTPM boots fine: %s", err)
	}
	if firmware != "uefi" {
		t.Errorf("firmware = %q, want the template's \"uefi\" reported back", firmware)
	}
}

func TestEnsureVMCanStart_PrefersTheConfiguredFirmwareOverTheTemplates(t *testing.T) {
	api := &fakeBootCheck{firmware: "uefi", vtpms: 1}

	if _, err := ensureVMCanStart(api, "bios"); err == nil {
		t.Error("firmware = \"bios\" overrides a uefi template, so the vTPM stops the boot")
	}
	if api.reads != 0 {
		t.Errorf("read the VM's firmware %d times, want 0 when the build named one", api.reads)
	}
}

func TestEnsureVMCanStart_LeavesValuesXAPIWillRejectAlone(t *testing.T) {
	for _, firmware := range []string{"BIOS", "UEFI", "Uefi", "efi", "legacy"} {
		t.Run(firmware, func(t *testing.T) {
			reported, err := ensureVMCanStart(&fakeBootCheck{vtpms: 1}, firmware)
			if err != nil {
				t.Errorf("VM.start reports %q itself, and names the value: %s", firmware, err)
			}
			if reported != firmware {
				t.Errorf("firmware = %q, want the build's own %q reported back unchanged", reported, firmware)
			}
		})
	}
}

func TestEnsureVMCanStart_LeavesAnUnrecognisedInheritedFirmwareAlone(t *testing.T) {
	reported, err := ensureVMCanStart(&fakeBootCheck{firmware: "UEFI", vtpms: 1}, "")
	if err != nil {
		t.Errorf("the source's own firmware is XAPI's to reject: %s", err)
	}
	if reported != "UEFI" {
		t.Errorf("firmware = %q, want the source's own \"UEFI\" reported back unchanged", reported)
	}
}

func TestEnsureVMCanStart_AllowsEverythingThatCanActuallyBoot(t *testing.T) {
	cases := map[string]struct {
		configured string
		api        *fakeBootCheck
		want       string
	}{
		"bios without a vTPM":            {"bios", &fakeBootCheck{}, "bios"},
		"uefi with a vTPM":               {"uefi", &fakeBootCheck{vtpms: 1}, "uefi"},
		"a bios template without a vTPM": {"", &fakeBootCheck{firmware: "bios"}, "bios"},
		"a uefi template without a vTPM": {"", &fakeBootCheck{firmware: "uefi"}, "uefi"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			firmware, err := ensureVMCanStart(tc.api, tc.configured)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
			if firmware != tc.want {
				t.Errorf("firmware = %q, want %q", firmware, tc.want)
			}
		})
	}
}

func TestEnsureVMCanStart_LetsTheBuildRunWhenItCannotRead(t *testing.T) {
	unreadable := errors.New("API error: code 1, message SESSION_INVALID")

	cases := map[string]struct {
		api  *fakeBootCheck
		want string
	}{
		"the firmware cannot be read":   {&fakeBootCheck{firmwareErr: unreadable, vtpms: 1}, ""},
		"the vtpm state cannot be read": {&fakeBootCheck{vtpmErr: unreadable}, "bios"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			firmware, err := ensureVMCanStart(tc.api, "")
			if err != nil {
				t.Errorf("a failed read must not fail the build here, got: %s", err)
			}
			if firmware != tc.want {
				t.Errorf("firmware = %q, want %q", firmware, tc.want)
			}
		})
	}
}
