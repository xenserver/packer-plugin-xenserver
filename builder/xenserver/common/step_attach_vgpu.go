package common

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepAttachvGPU struct {
	VGPU     xenapi.VGPURef
	VGPUType xenapi.VGPUTypeRef
	GPUGroup xenapi.GPUGroupRef

	VGPUName string
}

func (self *StepAttachvGPU) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	c := state.Get("client").(*Connection)
	ui := state.Get("ui").(packer.Ui)
	IsFound := false
	ui.Say("Step: Find and Attach vGPU: " + self.VGPUName)

	if len(self.VGPUName) == 0 {
		ui.Say("No vGPU requested, skipping")
		return multistep.ActionContinue
	}

	vgputypes, err := xenapi.VGPUType.GetAll(c.session)
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	for i := 0; i < len(vgputypes); i++ {
		vgpuname, err := xenapi.VGPUType.GetModelName(c.session, vgputypes[i])
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		if strings.Contains(strings.ToLower(vgpuname), strings.ToLower(self.VGPUName)) {
			ui.Say(fmt.Sprintf("Found profile with name: %s, uuid: %s", vgpuname, string(vgputypes[i])))
			IsFound = true
			self.VGPUType = vgputypes[i]
			break
		}
	}

	if IsFound {
		gpugroups, err := xenapi.VGPUType.GetEnabledOnGPUGroups(c.session, self.VGPUType)
		if err != nil {
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		if len(gpugroups) > 0 {
			// Select the first Group for now
			self.GPUGroup = gpugroups[0]

			// Get the VM we are going to attach the GPU to
			uuid := state.Get("instance_uuid").(string)
			instance, err := xenapi.VM.GetByUUID(c.session, uuid)
			if err != nil {
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
			ui.Say("Attaching vGPU to VM")
			vgpu, err := xenapi.VGPU.Create(c.session, instance, self.GPUGroup, "0", nil, self.VGPUType)
			if err != nil {
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
			self.VGPU = vgpu
		}
	} else {
		ui.Error("Request vGPU not found")
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (self *StepAttachvGPU) Cleanup(state multistep.StateBag) {}
