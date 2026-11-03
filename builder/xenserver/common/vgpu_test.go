// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"

	"xenapi"
)

type fakeVGPU struct {
	modelNames    map[xenapi.VGPUTypeRef]string
	enabledGroups map[xenapi.VGPUTypeRef][]xenapi.GPUGroupRef

	allTypes      []xenapi.VGPUTypeRef
	allTypesErr   error
	modelNameErr  error
	enabledErr    error
	createErr     error
	createdOnType xenapi.VGPUTypeRef
	createdGroup  xenapi.GPUGroupRef
	creates       int
}

func (f *fakeVGPU) AllTypes() ([]xenapi.VGPUTypeRef, error) {
	if f.allTypesErr != nil {
		return nil, f.allTypesErr
	}
	return f.allTypes, nil
}

func (f *fakeVGPU) ModelName(vgpuType xenapi.VGPUTypeRef) (string, error) {
	if f.modelNameErr != nil {
		return "", f.modelNameErr
	}
	return f.modelNames[vgpuType], nil
}

func (f *fakeVGPU) EnabledGroups(vgpuType xenapi.VGPUTypeRef) ([]xenapi.GPUGroupRef, error) {
	if f.enabledErr != nil {
		return nil, f.enabledErr
	}
	return f.enabledGroups[vgpuType], nil
}

func (f *fakeVGPU) Create(group xenapi.GPUGroupRef, vgpuType xenapi.VGPUTypeRef) (xenapi.VGPURef, error) {
	f.creates++
	if f.createErr != nil {
		return "", f.createErr
	}
	f.createdOnType = vgpuType
	f.createdGroup = group
	return "opaque-vgpu", nil
}

func poolWithT4() *fakeVGPU {
	return &fakeVGPU{
		allTypes: []xenapi.VGPUTypeRef{"opaque-passthrough", "opaque-t4-2q"},
		modelNames: map[xenapi.VGPUTypeRef]string{
			"opaque-passthrough": "passthrough",
			"opaque-t4-2q":       "GRID T4-2Q",
		},
		enabledGroups: map[xenapi.VGPUTypeRef][]xenapi.GPUGroupRef{
			"opaque-passthrough": {"opaque-group-0"},
			"opaque-t4-2q":       {"opaque-group-0", "opaque-group-1"},
		},
	}
}

func TestAttachVGPU_AttachesTheMatchedProfileFromTheFirstEnabledGroup(t *testing.T) {
	api := poolWithT4()

	attached, err := attachVGPU(api, "t4-2q", discard)

	if err != nil {
		t.Fatalf("attachVGPU() = %v, want nil", err)
	}
	if api.creates != 1 {
		t.Fatalf("created %d vGPUs, want 1", api.creates)
	}
	if api.createdOnType != "opaque-t4-2q" {
		t.Errorf("created a vGPU of type %q, want the matched type", api.createdOnType)
	}
	if api.createdGroup != "opaque-group-0" {
		t.Errorf("created a vGPU from group %q, want the first enabled group", api.createdGroup)
	}
	want := AttachedVGPU{VGPU: "opaque-vgpu", VGPUType: "opaque-t4-2q", GPUGroup: "opaque-group-0"}
	if attached != want {
		t.Errorf("attachVGPU() = %+v, want %+v", attached, want)
	}
}

func TestAttachVGPU_NoEnabledGPUGroupStopsTheBuild(t *testing.T) {
	api := poolWithT4()
	api.enabledGroups["opaque-t4-2q"] = nil

	_, err := attachVGPU(api, "t4-2q", discard)

	if err == nil {
		t.Fatal("attachVGPU() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "GRID T4-2Q") {
		t.Errorf("error %q does not name the vGPU type that has no GPU group", err)
	}
	if !strings.Contains(err.Error(), "t4-2q") {
		t.Errorf("error %q does not name the requested vgpu_profile", err)
	}
	if api.creates != 0 {
		t.Errorf("created %d vGPUs, want none", api.creates)
	}
}

func TestAttachVGPU_UnknownProfileStopsTheBuild(t *testing.T) {
	api := poolWithT4()

	_, err := attachVGPU(api, "A16-4Q", discard)

	if err == nil {
		t.Fatal("attachVGPU() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "A16-4Q") {
		t.Errorf("error %q does not name the missing profile", err)
	}
	if api.creates != 0 {
		t.Errorf("created %d vGPUs, want none", api.creates)
	}
}

func TestAttachVGPU_ModelNameFailureStopsTheBuild(t *testing.T) {
	api := poolWithT4()
	api.modelNameErr = errors.New("API error: code 1, message HANDLE_INVALID, data [1]")

	_, err := attachVGPU(api, "t4-2q", discard)

	if err == nil {
		t.Fatal("attachVGPU() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "HANDLE_INVALID") {
		t.Errorf("error %q does not carry the XAPI cause", err)
	}
	if api.creates != 0 {
		t.Errorf("created %d vGPUs, want none", api.creates)
	}
}

func TestAttachVGPU_CreateFailureStopsTheBuild(t *testing.T) {
	api := poolWithT4()
	api.createErr = errors.New("API error: code 1, message VGPU_TYPE_NOT_ENABLED, data [1]")

	_, err := attachVGPU(api, "t4-2q", discard)

	if err == nil {
		t.Fatal("attachVGPU() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "VGPU_TYPE_NOT_ENABLED") {
		t.Errorf("error %q does not carry the XAPI cause", err)
	}
}

func TestAttachVGPU_TypeListingFailureStopsTheBuild(t *testing.T) {
	api := poolWithT4()
	api.allTypesErr = errors.New("API error: code 1, message SESSION_INVALID, data [1]")

	_, err := attachVGPU(api, "t4-2q", discard)

	if err == nil {
		t.Fatal("attachVGPU() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "SESSION_INVALID") {
		t.Errorf("error %q does not carry the XAPI cause", err)
	}
	if api.creates != 0 {
		t.Errorf("created %d vGPUs, want none", api.creates)
	}
}

func TestAttachVGPU_EnabledGroupLookupFailureStopsTheBuild(t *testing.T) {
	api := poolWithT4()
	api.enabledErr = errors.New("API error: code 1, message HANDLE_INVALID, data [1]")

	_, err := attachVGPU(api, "t4-2q", discard)

	if err == nil {
		t.Fatal("attachVGPU() = nil, want an error")
	}
	if !strings.Contains(err.Error(), "HANDLE_INVALID") {
		t.Errorf("error %q does not carry the XAPI cause", err)
	}
	if api.creates != 0 {
		t.Errorf("created %d vGPUs, want none", api.creates)
	}
}
