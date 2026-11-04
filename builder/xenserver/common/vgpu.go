// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"strings"

	"xenapi"
)

// XAPI reads device "0" as the next free slot, so a vGPU inherited by a clone does not clash.
const firstVGPUDevice = "0"

// AttachedVGPU records the vGPU an attach created, with its type and GPU group.
type AttachedVGPU struct {
	VGPU     xenapi.VGPURef
	VGPUType xenapi.VGPUTypeRef
	GPUGroup xenapi.GPUGroupRef
}

type vgpuAPI interface {
	AllTypes() ([]xenapi.VGPUTypeRef, error)
	ModelName(vgpuType xenapi.VGPUTypeRef) (string, error)
	EnabledGroups(vgpuType xenapi.VGPUTypeRef) ([]xenapi.GPUGroupRef, error)
	Create(group xenapi.GPUGroupRef, vgpuType xenapi.VGPUTypeRef) (xenapi.VGPURef, error)
}

type xapiVGPU struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (v xapiVGPU) AllTypes() ([]xenapi.VGPUTypeRef, error) {
	return xenapi.VGPUType.GetAll(v.session)
}

func (v xapiVGPU) ModelName(vgpuType xenapi.VGPUTypeRef) (string, error) {
	return xenapi.VGPUType.GetModelName(v.session, vgpuType)
}

func (v xapiVGPU) EnabledGroups(vgpuType xenapi.VGPUTypeRef) ([]xenapi.GPUGroupRef, error) {
	return xenapi.VGPUType.GetEnabledOnGPUGroups(v.session, vgpuType)
}

func (v xapiVGPU) Create(group xenapi.GPUGroupRef, vgpuType xenapi.VGPUTypeRef) (xenapi.VGPURef, error) {
	return xenapi.VGPU.Create(v.session, v.vm, group, firstVGPUDevice, nil, vgpuType)
}

// AttachVGPU adds a vGPU of the first type whose model name contains profile; inherited vGPUs stay.
func AttachVGPU(c *Connection, vm xenapi.VMRef, profile string, say func(string)) (AttachedVGPU, error) {
	return attachVGPU(xapiVGPU{session: c.GetSession(), vm: vm}, profile, say)
}

func attachVGPU(api vgpuAPI, profile string, say func(string)) (AttachedVGPU, error) {
	vgpuType, modelName, err := findVGPUType(api, profile)
	if err != nil {
		return AttachedVGPU{}, err
	}
	say(fmt.Sprintf("Found profile with name: %s, ref: %s", modelName, string(vgpuType)))

	groups, err := api.EnabledGroups(vgpuType)
	if err != nil {
		return AttachedVGPU{}, fmt.Errorf("could not read the GPU groups that have the vGPU type %q enabled: %w", modelName, err)
	}

	// XAPI keeps a type record for every model the pool has seen, so a match may have no enabled group.
	if len(groups) == 0 {
		return AttachedVGPU{}, fmt.Errorf(
			"the vGPU type %q matches the requested vgpu_profile %q, but no GPU group has that type enabled, "+
				"so no vGPU can be attached; enable the type on a GPU group that holds a physical GPU supporting it, "+
				"or ask for a vgpu_profile the pool can serve",
			modelName, profile)
	}

	// No ranking by remaining capacity: a full first group fails only when the VM starts.
	group := groups[0]

	say("Attaching vGPU to VM")
	vgpu, err := api.Create(group, vgpuType)
	if err != nil {
		return AttachedVGPU{}, fmt.Errorf("could not attach a vGPU of type %q to the VM: %w", modelName, err)
	}

	return AttachedVGPU{VGPU: vgpu, VGPUType: vgpuType, GPUGroup: group}, nil
}

func findVGPUType(api vgpuAPI, profile string) (xenapi.VGPUTypeRef, string, error) {
	types, err := api.AllTypes()
	if err != nil {
		return "", "", fmt.Errorf("could not list the pool's vGPU types: %w", err)
	}

	want := strings.ToLower(profile)
	for _, vgpuType := range types {
		modelName, err := api.ModelName(vgpuType)
		if err != nil {
			return "", "", fmt.Errorf("could not read a vGPU type's model name: %w", err)
		}

		if strings.Contains(strings.ToLower(modelName), want) {
			return vgpuType, modelName, nil
		}
	}

	return "", "", fmt.Errorf("no vGPU type on this pool has a model name matching the requested vgpu_profile %q", profile)
}
