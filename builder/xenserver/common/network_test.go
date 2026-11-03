// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"

	"xenapi"
)

type fakeNamedNetwork struct {
	byLabel   map[string][]xenapi.NetworkRef
	lookupErr map[string]error
	attachErr map[xenapi.NetworkRef]error

	lookedUp []string
	attached []attachedVIF
}

type attachedVIF struct {
	network xenapi.NetworkRef
	device  string
}

func (f *fakeNamedNetwork) FindByNameLabel(label string) ([]xenapi.NetworkRef, error) {
	f.lookedUp = append(f.lookedUp, label)
	if err := f.lookupErr[label]; err != nil {
		return nil, err
	}
	return f.byLabel[label], nil
}

func (f *fakeNamedNetwork) AttachVIF(network xenapi.NetworkRef, device string) error {
	if err := f.attachErr[network]; err != nil {
		return err
	}
	f.attached = append(f.attached, attachedVIF{network: network, device: device})
	return nil
}

func discard(string) {}

func TestAttachNamedNetworks_OneVIFPerNameInOrder(t *testing.T) {
	api := &fakeNamedNetwork{byLabel: map[string][]xenapi.NetworkRef{
		"public":  {"opaque-public"},
		"storage": {"opaque-storage"},
	}}

	err := attachNamedNetworks(api, []string{"public", "storage"}, discard)

	if err != nil {
		t.Fatalf("attachNamedNetworks() = %v, want nil", err)
	}
	want := []attachedVIF{
		{network: "opaque-public", device: "0"},
		{network: "opaque-storage", device: "1"},
	}
	if len(api.attached) != len(want) {
		t.Fatalf("attached %d VIFs, want %d: %v", len(api.attached), len(want), api.attached)
	}
	for i, w := range want {
		if api.attached[i] != w {
			t.Errorf("VIF %d = %+v, want %+v", i, api.attached[i], w)
		}
	}
}

func TestAttachNamedNetworks_AVIFFailureStopsTheBuild(t *testing.T) {
	api := &fakeNamedNetwork{
		byLabel: map[string][]xenapi.NetworkRef{
			"public":  {"opaque-public"},
			"storage": {"opaque-storage"},
			"backup":  {"opaque-backup"},
		},
		attachErr: map[xenapi.NetworkRef]error{
			"opaque-storage": errors.New("API error: code 1, message DEVICE_ALREADY_EXISTS, data [1]"),
		},
	}

	err := attachNamedNetworks(api, []string{"public", "storage", "backup"}, discard)

	if err == nil {
		t.Fatal("attachNamedNetworks() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "storage") {
		t.Errorf("error %q does not name the network that failed", err)
	}
	if !strings.Contains(err.Error(), "DEVICE_ALREADY_EXISTS") {
		t.Errorf("error %q does not carry the XAPI cause", err)
	}
	if len(api.attached) != 1 || api.attached[0].network != "opaque-public" {
		t.Errorf("attached %v, want only the network preceding the failure", api.attached)
	}
}

func TestAttachNamedNetworks_UnknownNameLabelIsFatal(t *testing.T) {
	api := &fakeNamedNetwork{byLabel: map[string][]xenapi.NetworkRef{}}

	err := attachNamedNetworks(api, []string{"absent"}, discard)

	if err == nil {
		t.Fatal("attachNamedNetworks() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "absent") {
		t.Errorf("error %q does not name the missing network", err)
	}
	if len(api.attached) != 0 {
		t.Errorf("attached %v, want nothing", api.attached)
	}
}

func TestAttachNamedNetworks_AmbiguousNameLabelIsFatal(t *testing.T) {
	api := &fakeNamedNetwork{byLabel: map[string][]xenapi.NetworkRef{
		"double": {"opaque-a", "opaque-b"},
	}}

	err := attachNamedNetworks(api, []string{"double"}, discard)

	if err == nil {
		t.Fatal("attachNamedNetworks() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "double") {
		t.Errorf("error %q does not name the ambiguous network", err)
	}
	if len(api.attached) != 0 {
		t.Errorf("attached %v, want nothing", api.attached)
	}
}

func TestAttachNamedNetworks_FallsBackToThePoolWideLabel(t *testing.T) {
	api := &fakeNamedNetwork{byLabel: map[string][]xenapi.NetworkRef{
		"Pool-wide network associated with eth3": {"opaque-eth3"},
	}}

	err := attachNamedNetworks(api, []string{"Network 3"}, discard)

	if err != nil {
		t.Fatalf("attachNamedNetworks() = %v, want nil", err)
	}
	if len(api.attached) != 1 || api.attached[0].network != "opaque-eth3" {
		t.Fatalf("attached %v, want the pool-wide network on eth3", api.attached)
	}
	if len(api.lookedUp) != 2 || api.lookedUp[0] != "Network 3" {
		t.Errorf("looked up %v, want the given label first and then the pool-wide one", api.lookedUp)
	}
}

func TestAttachNamedNetworks_NoFallbackForOtherLabels(t *testing.T) {
	api := &fakeNamedNetwork{byLabel: map[string][]xenapi.NetworkRef{}}

	err := attachNamedNetworks(api, []string{"whatever"}, discard)

	if err == nil {
		t.Fatal("attachNamedNetworks() = nil, want an error")
	}
	if len(api.lookedUp) != 1 {
		t.Errorf("looked up %v, want a single lookup and no pool-wide fallback", api.lookedUp)
	}
}

func TestAttachNamedNetworks_LookupFailureIsFatal(t *testing.T) {
	api := &fakeNamedNetwork{
		byLabel:   map[string][]xenapi.NetworkRef{},
		lookupErr: map[string]error{"public": errors.New("session invalid")},
	}

	err := attachNamedNetworks(api, []string{"public"}, discard)

	if err == nil {
		t.Fatal("attachNamedNetworks() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "session invalid") {
		t.Errorf("error %q does not carry the cause", err)
	}
}

func TestAttachNamedNetworks_PoolWideLookupFailureIsFatal(t *testing.T) {
	api := &fakeNamedNetwork{
		byLabel:   map[string][]xenapi.NetworkRef{},
		lookupErr: map[string]error{"Pool-wide network associated with eth3": errors.New("session invalid")},
	}

	err := attachNamedNetworks(api, []string{"Network 3"}, discard)

	if err == nil {
		t.Fatal("attachNamedNetworks() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "session invalid") {
		t.Errorf("error %q does not carry the cause", err)
	}
}

func TestAttachNamedNetworks_NoNamesAttachesNothing(t *testing.T) {
	api := &fakeNamedNetwork{byLabel: map[string][]xenapi.NetworkRef{}}

	err := attachNamedNetworks(api, nil, discard)

	if err != nil {
		t.Fatalf("attachNamedNetworks() = %v, want nil", err)
	}
	if len(api.attached) != 0 {
		t.Errorf("attached %v, want nothing", api.attached)
	}
}

type fakeVMVIFs struct {
	vifs       []xenapi.VIFRef
	listErr    error
	destroyErr map[xenapi.VIFRef]error

	destroyed []xenapi.VIFRef
}

func (f *fakeVMVIFs) ExistingVIFs() ([]xenapi.VIFRef, error) {
	if f.listErr != nil {
		return nil, f.listErr
	}
	return f.vifs, nil
}

func (f *fakeVMVIFs) DestroyVIF(vif xenapi.VIFRef) error {
	if err := f.destroyErr[vif]; err != nil {
		return err
	}
	f.destroyed = append(f.destroyed, vif)
	return nil
}

func TestDetachAllVIFs_RemovesEveryInheritedVIF(t *testing.T) {
	api := &fakeVMVIFs{vifs: []xenapi.VIFRef{"vif0", "vif1", "vif2"}}

	if err := detachAllVIFs(api, discard); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(api.destroyed) != 3 {
		t.Fatalf("destroyed %d VIFs, want 3: %v", len(api.destroyed), api.destroyed)
	}
	for i, want := range []xenapi.VIFRef{"vif0", "vif1", "vif2"} {
		if api.destroyed[i] != want {
			t.Errorf("destroyed[%d] = %q, want %q", i, api.destroyed[i], want)
		}
	}
}

func TestDetachAllVIFs_NoVIFsIsANoOpAndSaysNothing(t *testing.T) {
	api := &fakeVMVIFs{}
	said := 0

	if err := detachAllVIFs(api, func(string) { said++ }); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(api.destroyed) != 0 {
		t.Errorf("destroyed %v, want none", api.destroyed)
	}
	if said != 0 {
		t.Errorf("emitted %d messages, want 0 for a VM with no VIFs", said)
	}
}

func TestDetachAllVIFs_ADestroyFailureIsFatal(t *testing.T) {
	api := &fakeVMVIFs{
		vifs:       []xenapi.VIFRef{"vif0", "vif1"},
		destroyErr: map[xenapi.VIFRef]error{"vif1": errors.New("DEVICE_DETACH_REJECTED")},
	}

	err := detachAllVIFs(api, discard)
	if err == nil {
		t.Fatal("expected an error when a VIF cannot be destroyed")
	}
	if !strings.Contains(err.Error(), "vif1") {
		t.Errorf("error %q does not name the VIF that failed", err)
	}
}

func TestDetachAllVIFs_AListFailureIsFatal(t *testing.T) {
	api := &fakeVMVIFs{listErr: errors.New("SESSION_INVALID")}

	if err := detachAllVIFs(api, discard); err == nil {
		t.Fatal("expected an error when the VIF list cannot be read")
	}
}

func TestCountVIFs_ReportsTheInheritedCount(t *testing.T) {
	api := &fakeVMVIFs{vifs: []xenapi.VIFRef{"vif0", "vif1"}}

	got, err := countVIFs(api)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 2 {
		t.Errorf("countVIFs = %d, want 2", got)
	}
}

func TestCountVIFs_AListFailureIsFatal(t *testing.T) {
	api := &fakeVMVIFs{listErr: errors.New("SESSION_INVALID")}

	if _, err := countVIFs(api); err == nil {
		t.Fatal("expected an error when the VIF list cannot be read")
	}
}
