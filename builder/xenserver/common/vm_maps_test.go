// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"testing"
)

func TestSetPlatformKey_RemovesBeforeAdding(t *testing.T) {
	var fake *fakeXAPI
	var removedFirst bool
	fake = newFakeXAPI(t, map[string]interface{}{
		"VM.remove_from_platform": "",
		"VM.add_to_platform": func([]interface{}) interface{} {
			removedFirst = fake.wasCalled("VM.remove_from_platform")
			return ""
		},
	})

	if err := setPlatformKey(fake.conn().GetSession(), "OpaqueRef:vm", "secureboot", "true"); err != nil {
		t.Fatalf("setPlatformKey: %s", err)
	}

	if !removedFirst {
		t.Error("the key was added without being removed first; XAPI raises MAP_DUPLICATE_KEY")
	}
	if args := fake.argsOf(t, "VM.add_to_platform"); len(args) != 3 || args[1] != "secureboot" || args[2] != "true" {
		t.Errorf("VM.add_to_platform%v, want the key and value", args)
	}
}

func TestSetPlatformKey_StopsWhenTheRemoveFails(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.remove_from_platform": errors.New("SESSION_INVALID"),
		"VM.add_to_platform":      "",
	})

	if err := setPlatformKey(fake.conn().GetSession(), "OpaqueRef:vm", "secureboot", "true"); err == nil {
		t.Fatal("setPlatformKey reported success after the remove failed")
	}
	if fake.wasCalled("VM.add_to_platform") {
		t.Error("the key was added after the remove failed")
	}
}

func TestSetHVMBootParam_RemovesBeforeAdding(t *testing.T) {
	var fake *fakeXAPI
	var removedFirst bool
	fake = newFakeXAPI(t, map[string]interface{}{
		"VM.remove_from_HVM_boot_params": "",
		"VM.add_to_HVM_boot_params": func([]interface{}) interface{} {
			removedFirst = fake.wasCalled("VM.remove_from_HVM_boot_params")
			return ""
		},
	})

	if err := setHVMBootParam(fake.conn().GetSession(), "OpaqueRef:vm", "firmware", "uefi"); err != nil {
		t.Fatalf("setHVMBootParam: %s", err)
	}

	if !removedFirst {
		t.Error("the key was added without being removed first; XAPI raises MAP_DUPLICATE_KEY")
	}
	if args := fake.argsOf(t, "VM.add_to_HVM_boot_params"); len(args) != 3 || args[1] != "firmware" || args[2] != "uefi" {
		t.Errorf("VM.add_to_HVM_boot_params%v, want the key and value", args)
	}
}
