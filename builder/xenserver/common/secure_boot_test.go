// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"
)

type fakeSecureBoot struct {
	err    error
	writes []string
}

func (f *fakeSecureBoot) Set(value string) error {
	if f.err != nil {
		return f.err
	}
	f.writes = append(f.writes, value)
	return nil
}

func TestReconcileSecureBoot_WritesNothingWhenTheBuildSaidNothing(t *testing.T) {
	api := &fakeSecureBoot{}

	outcome, err := reconcileSecureBoot(api, nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if outcome != SecureBootKept {
		t.Errorf("outcome = %d, want SecureBootKept", outcome)
	}
	if len(api.writes) != 0 {
		t.Errorf("wrote platform:secureboot %v, want no write at all", api.writes)
	}
}

func TestReconcileSecureBoot_EnablesWithAuto(t *testing.T) {
	api := &fakeSecureBoot{}

	outcome, err := reconcileSecureBoot(api, boolPtr(true))

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if outcome != SecureBootAuto {
		t.Errorf("outcome = %d, want SecureBootAuto", outcome)
	}
	if len(api.writes) != 1 || api.writes[0] != "auto" {
		t.Errorf("wrote %v, want exactly one write of \"auto\"", api.writes)
	}
}

func TestReconcileSecureBoot_DisablesOnAnExplicitFalse(t *testing.T) {
	api := &fakeSecureBoot{}

	outcome, err := reconcileSecureBoot(api, boolPtr(false))

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if outcome != SecureBootDisabled {
		t.Errorf("outcome = %d, want SecureBootDisabled", outcome)
	}
	if len(api.writes) != 1 || api.writes[0] != "false" {
		t.Errorf("wrote %v, want exactly one write of \"false\"", api.writes)
	}
}

func TestReconcileSecureBoot_FailsTheBuildWhenTheWriteFails(t *testing.T) {
	api := &fakeSecureBoot{err: errors.New("API error: code 1, message SESSION_INVALID")}

	outcome, err := reconcileSecureBoot(api, boolPtr(true))

	if err == nil {
		t.Fatal("a failed write leaves the VM in a state the build did not ask for")
	}
	if outcome != SecureBootKept {
		t.Errorf("outcome = %d, want SecureBootKept reported when nothing was written", outcome)
	}
	if !strings.Contains(err.Error(), "auto") || !strings.Contains(err.Error(), "SESSION_INVALID") {
		t.Errorf("the error must name the value and the cause, got: %s", err)
	}
}

func TestSecureBootOutcome_Describe(t *testing.T) {
	kept := SecureBootKept.Describe("Windows 11")
	if !strings.Contains(kept, "Windows 11") {
		t.Errorf("a kept setting must name the source it came from, got: %s", kept)
	}

	for outcome, want := range map[SecureBootOutcome]string{
		SecureBootDisabled: "Disabled",
		SecureBootAuto:     "Auto",
	} {
		if got := outcome.Describe("Windows 11"); !strings.Contains(got, want) {
			t.Errorf("Describe() = %q, want it to mention %q", got, want)
		}
	}
}
