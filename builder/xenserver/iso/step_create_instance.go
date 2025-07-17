package iso

import (
	"context"
	"fmt"
	"log"
	"strings"
	"strconv"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
	
	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
)

type stepCreateInstance struct {
	instance *xenapi.VMRef
	vdi      *xenapi.VDIRef
}

func (self *stepCreateInstance) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {

	c := state.Get("client").(*xscommon.Connection)
	config := state.Get("config").(xscommon.Config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Create Instance")

	// Get the template to clone from

	vms, err := xenapi.VM.GetByNameLabel(c.GetSession(), config.CloneTemplate)

	switch {
	case len(vms) == 0:
		ui.Error(fmt.Sprintf("Couldn't find a template with the name-label '%s'. Aborting.", config.CloneTemplate))
		return multistep.ActionHalt
	case len(vms) > 1:
		ui.Error(fmt.Sprintf("Found more than one template with the name '%s'. The name must be unique. Aborting.", config.CloneTemplate))
		return multistep.ActionHalt
	}

	template := vms[0]

	// Clone that VM template
	instance, err := xenapi.VM.Clone(c.GetSession(), template, config.VMName)
	if err != nil {
		ui.Error(fmt.Sprintf("Error cloning VM: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.instance = &instance

	err = xenapi.VM.SetIsATemplate(c.GetSession(), instance, false)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting is_a_template=false: %s", err.Error()))
		return multistep.ActionHalt
	}

	err = xenapi.VM.SetVCPUsMax(c.GetSession(), instance, int(config.VCPUsMax))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM VCPUs Max=%d: %s", config.VCPUsMax, err.Error()))
		return multistep.ActionHalt
	}

	err = xenapi.VM.SetVCPUsAtStartup(c.GetSession(), instance, int(config.VCPUsAtStartup))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM VCPUs At Startup=%d: %s", config.VCPUsAtStartup, err.Error()))
		return multistep.ActionHalt
	}

	memory := int(config.VMMemory * 1024 * 1024)
	err = xenapi.VM.SetMemoryLimits(c.GetSession(), instance, memory, memory, memory, memory)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM memory=%d: %s", memory, err.Error()))
		return multistep.ActionHalt
	}

	// Only set Platform Args when parameter is set, otherwise we overrule the Template Args
	if len(config.PlatformArgs) != 0 {
		err = xenapi.VM.SetPlatform(c.GetSession(), instance, config.PlatformArgs)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM platform: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	// Set Cores per Socket
	coresPerSocket := config.CorePerSocket
	vcpus := config.VCPUs
	if vcpus%coresPerSocket != 0 {
		ui.Error(fmt.Sprintf("%d cores could not fit to %d cores-per-socket topology", vcpus, coresPerSocket))
		return multistep.ActionHalt
	}
	err = xenapi.VM.RemoveFromPlatform(c.GetSession(), instance, "cores-per-socket")
	if err != nil {
		ui.Error(fmt.Sprintf("Error removing cores-per-socket from VM platform: %s", err.Error()))
		return multistep.ActionHalt
	}
	err = xenapi.VM.AddToPlatform(c.GetSession(), instance, "cores-per-socket", strconv.FormatUint(uint64(coresPerSocket), 10))
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting cores-per-socket to %d: %s", coresPerSocket, err.Error()))
		return multistep.ActionHalt
	}

	// Set secure boot
	if config.SecureBoot && config.Firmware == "uefi" {
		ui.Say("Set Secure boot to Auto")
		err = xenapi.VM.RemoveFromPlatform(c.GetSession(), instance, "secureboot")
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM Secure Boot=%t: %s", config.SecureBoot, err.Error()))
			return multistep.ActionHalt
		}
		err = xenapi.VM.AddToPlatform(c.GetSession(), instance, "secureboot", "auto")
	} else {
		ui.Say("Set Secure boot to Disabled")
		err = xenapi.VM.RemoveFromPlatform(c.GetSession(), instance, "secureboot")
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM Secure Boot=%t: %s", config.SecureBoot, err.Error()))
			return multistep.ActionHalt
		}
		err = xenapi.VM.AddToPlatform(c.GetSession(), instance, "secureboot", "false")
	}

	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM Secure Boot=%t: %s", config.SecureBoot, err.Error()))
		return multistep.ActionHalt
	}

	err = xenapi.VM.SetNameDescription(c.GetSession(), instance, config.VMDescription)
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting VM description: %s", err.Error()))
		return multistep.ActionHalt
	}

	if len(config.VMOtherConfig) != 0 {
		vm_other_config, err := xenapi.VM.GetOtherConfig(c.GetSession(), instance)
		if err != nil {
			ui.Error(fmt.Sprintf("Error getting VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}
		for key, value := range config.VMOtherConfig {
			vm_other_config[key] = value
		}
		err = xenapi.VM.SetOtherConfig(c.GetSession(), instance, vm_other_config)
		if err != nil {
			ui.Error(fmt.Sprintf("Error setting VM other-config: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	// Create VDI for the instance
	sr, err := config.GetSR(c)

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Using the following SR for the VM: %s", sr))

	vdi, err := xenapi.VDI.Create(c.GetSession(), xenapi.VDIRecord{
		NameLabel:   "Packer-disk",
		VirtualSize: int(config.DiskSize * 1024 * 1024),
		Type:        "user",
		Sharable:    false,
		ReadOnly:    false,
		SR:          sr,
		OtherConfig: map[string]string{
			"temp": "temp",
		},
	})
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}
	self.vdi = &vdi

	err = xscommon.ConnectVdi(c, instance, vdi, xenapi.VbdTypeDisk)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to connect packer disk VDI: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Connect Network

	var network xenapi.NetworkRef

	if len(config.NetworkNames) == 0 {
		// No network has be specified. Use the management interface
		log.Println("No network name given, attempting to use management interface")
		pifs, err := xenapi.PIF.GetAll(c.GetSession())

		if err != nil {
			ui.Error(fmt.Sprintf("Error getting PIFs: %s", err.Error()))
			return multistep.ActionHalt
		}

		for _, pif := range pifs {
			pif_rec, err := xenapi.PIF.GetRecord(c.GetSession(), pif)

			if err != nil {
				ui.Error(fmt.Sprintf("Error getting PIF record: %s", err.Error()))
				return multistep.ActionHalt
			}

			if pif_rec.Management {
				network = pif_rec.Network
			}

		}

		if string(network) == "" {
			ui.Error("Error: couldn't find management network. Aborting.")
			return multistep.ActionHalt
		}

		log.Printf("Creating VIF on network '%s' on VM '%s'\n", network, instance)
		_, err = xscommon.ConnectNetwork(c, network, instance, "0")

		if err != nil {
			ui.Error(fmt.Sprintf("Failed to create VIF with error: %v", err))
			return multistep.ActionHalt
		}

	} else {
		log.Printf("Using provided network names: %v\n", config.NetworkNames)
		// Look up each network by it's name label
		for i, networkNameLabel := range config.NetworkNames {
			ui.Say(fmt.Sprintf("Finding network '%s'", networkNameLabel))
			networks, err := xenapi.Network.GetByNameLabel(c.GetSession(), networkNameLabel)

			if err != nil {
				ui.Error(fmt.Sprintf("Error occured getting Network by name-label: %s", err.Error()))
				return multistep.ActionHalt
			}
			// If network Name label starts with "Network ", we assume it is a default network
			if len(networks) == 0 && strings.HasPrefix(networkNameLabel, "Network ") {
				
				tmpNetworkNo := strings.TrimPrefix(networkNameLabel, "Network ")
				tmpLabel := "Pool-wide network associated with eth" + tmpNetworkNo
				ui.Say(fmt.Sprintf("No network found with name-label '%s'. This might be a default built-in network '%s'", networkNameLabel, tmpLabel))
				networks, err = xenapi.Network.GetByNameLabel(c.GetSession(), tmpLabel)
				if err != nil {
					ui.Error(fmt.Sprintf("Error occured getting default Network: %s", err.Error()))
					return multistep.ActionHalt
				}
			}

			switch {
				case len(networks) == 0:
					ui.Error(fmt.Sprintf("Couldn't find a network with the specified name-label '%s'. Aborting.", networkNameLabel))
					return multistep.ActionHalt
				case len(networks) > 1:
					ui.Error(fmt.Sprintf("Found more than one network with the name '%s'. The name must be unique. Aborting.", networkNameLabel))
					return multistep.ActionHalt
			}

			//we need the VIF index string
			vifIndexString := fmt.Sprintf("%d", i)
			_, err = xscommon.ConnectNetwork(c, networks[0], instance, vifIndexString)

			if err != nil {
				ui.Say(fmt.Sprintf("Failed to connect VIF with error: %v", err.Error()))
			}
		}
	}

	instanceId, err := xenapi.VM.GetUUID(c.GetSession(), instance)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM UUID: %s", err.Error()))
		return multistep.ActionHalt
	}

	state.Put("instance_uuid", instanceId)
	ui.Say(fmt.Sprintf("Created instance '%s'", instanceId))

	return multistep.ActionContinue
}

func (self *stepCreateInstance) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(xscommon.Config)
	if config.ShouldKeepVM(state) {
		return
	}

	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*xscommon.Connection)

	if self.instance != nil {
		ui.Say("Destroying VM")
		_ = xenapi.VM.HardShutdown(c.GetSession(), *self.instance) // redundant, just in case
		err := xenapi.VM.Destroy(c.GetSession(), *self.instance)
		if err != nil {
			ui.Error(err.Error())
		}
	}

	if self.vdi != nil {
		ui.Say("Destroying VDI")
		err := xenapi.VDI.Destroy(c.GetSession(), *self.vdi)
		if err != nil {
			ui.Error(err.Error())
		}
	}
}
