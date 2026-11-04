// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"
)

type fakeVendorDevice struct {
	current bool
	getErr  error
	setErr  error
	writes  []bool
}

func (f *fakeVendorDevice) Get() (bool, error) {
	if f.getErr != nil {
		return false, f.getErr
	}
	return f.current, nil
}

func (f *fakeVendorDevice) Set(value bool) error {
	f.writes = append(f.writes, value)
	if f.setErr != nil {
		return f.setErr
	}
	f.current = value
	return nil
}

func TestApplyVendorDevice_TurnsTheDeviceOnWhenTheSourceHasItOff(t *testing.T) {
	api := &fakeVendorDevice{current: false}

	changed, err := applyVendorDevice(api, true)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !changed {
		t.Error("changed = false, want true: the VM had the device off")
	}
	if len(api.writes) != 1 || api.writes[0] != true {
		t.Errorf("writes = %v, want exactly one write of true", api.writes)
	}
}

func TestApplyVendorDevice_TurnsTheDeviceOffWhenAskedTo(t *testing.T) {
	api := &fakeVendorDevice{current: true}

	changed, err := applyVendorDevice(api, false)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if !changed {
		t.Error("changed = false, want true: the VM had the device on")
	}
	if len(api.writes) != 1 || api.writes[0] != false {
		t.Errorf("writes = %v, want exactly one write of false", api.writes)
	}
}

func TestApplyVendorDevice_LeavesASourceThatAlreadyMatchesAlone(t *testing.T) {
	for _, want := range []bool{true, false} {
		api := &fakeVendorDevice{current: want}

		changed, err := applyVendorDevice(api, want)

		if err != nil {
			t.Errorf("want=%t: unexpected error: %s", want, err)
			continue
		}
		if changed {
			t.Errorf("want=%t: changed = true, but the VM was already in that state", want)
		}
		if len(api.writes) != 0 {
			t.Errorf("want=%t: wrote %v, but nothing needed changing", want, api.writes)
		}
	}
}

func TestApplyVendorDevice_ReportsAReadFailureAndWritesNothing(t *testing.T) {
	readFailed := errors.New("API error: code 1, message SESSION_INVALID")
	api := &fakeVendorDevice{getErr: readFailed}

	changed, err := applyVendorDevice(api, true)

	if err == nil {
		t.Fatal("a failed read must be reported")
	}
	if !errors.Is(err, readFailed) {
		t.Errorf("error should wrap the XAPI failure, got: %s", err)
	}
	if !strings.Contains(err.Error(), "has_vendor_device") {
		t.Errorf("error should name the option the user can act on, got: %s", err)
	}
	if changed {
		t.Error("changed = true after a failed read")
	}
	if len(api.writes) != 0 {
		t.Errorf("wrote %v despite not knowing the current value", api.writes)
	}
}

func TestApplyVendorDevice_ReportsAWriteFailure(t *testing.T) {
	writeFailed := errors.New("API error: code 1, message VM_BAD_POWER_STATE")
	api := &fakeVendorDevice{current: false, setErr: writeFailed}

	changed, err := applyVendorDevice(api, true)

	if err == nil {
		t.Fatal("a failed write must be reported")
	}
	if !errors.Is(err, writeFailed) {
		t.Errorf("error should wrap the XAPI failure, got: %s", err)
	}
	if !strings.Contains(err.Error(), "has_vendor_device") {
		t.Errorf("error should name the option the user can act on, got: %s", err)
	}
	if changed {
		t.Error("changed = true after a failed write")
	}
}
