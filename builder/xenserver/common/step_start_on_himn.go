package common

import (
	gossh "code.google.com/p/go.crypto/ssh"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	"log"
	"time"
)

type StepStartOnHIMN struct{}

/*
 * This step starts the installed guest on the Host Internal Management Network
 * as there exists an API to obtain the IP allocated to the VM by XAPI.
 * This in turn will allow Packer to SSH into the VM, provided NATing has been
 * enabled on the host.
 *
 */

func (self *StepStartOnHIMN) Run(state multistep.StateBag) multistep.StepAction {

	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	ui.Say("Step: Start VM on the Host Internal Mangement Network")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	// Find the HIMN Ref
	networks, err := client.GetNetworkByNameLabel("Host internal management network")
	if err != nil || len(networks) == 0 {
		ui.Error("Unable to find a host internal management network")
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	himn := networks[0]

	// Create a VIF for the HIMN
	himn_vif, err := instance.ConnectNetwork(himn, "0")
	if err != nil {
		ui.Error("Error creating VIF")
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Start the VM
	instance.Start(false, false)

	var himn_iface_ip string = ""

	// Obtain the allocated IP
	err = InterruptibleWait{
		Predicate: func() (found bool, err error) {
			ips, err := himn.GetAssignedIPs()
			if err != nil {
				return false, fmt.Errorf("Can't get assigned IPs: %s", err.Error())
			}
			log.Printf("IPs: %s", ips)
			log.Printf("Ref: %s", instance.Ref)

			//Check for instance.Ref in map
			if vm_ip, ok := ips[himn_vif.Ref]; ok && vm_ip != "" {
				ui.Say("Found the VM's IP: " + vm_ip)
				himn_iface_ip = vm_ip
				return true, nil
			}

			ui.Say("Wait for IP address...")
			return false, nil
		},
		PredicateInterval: 10 * time.Second,
		Timeout:           100 * time.Second,
	}.Wait(state)

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to find an IP on the Host-internal management interface: %s", err.Error()))
		return multistep.ActionHalt
	}

	state.Put("himn_ssh_address", himn_iface_ip)
	ui.Say("Stored VM's IP " + himn_iface_ip)

	// Wait for the VM to boot, and check we can ping this interface

	ping_cmd := fmt.Sprintf("ping -c 1 %s", himn_iface_ip)

	err = InterruptibleWait{
		Predicate: func() (success bool, err error) {
			ui.Message(fmt.Sprintf("Attempting to ping interface: %s", ping_cmd))
			_, err = ExecuteHostSSHCmd(state, ping_cmd)

			switch err.(type) {
			case nil:
				// ping succeeded
				return true, nil
			case *gossh.ExitError:
				// ping failed, try again
				return false, nil
			default:
				// unknown error
				return false, err
			}
		},
		PredicateInterval: 10 * time.Second,
		Timeout:           300 * time.Second,
	}.Wait(state)

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to ping interface. (Has the VM not booted?): %s", err.Error()))
		return multistep.ActionHalt
	}

	ui.Message("Ping success! Continuing...")
	return multistep.ActionContinue

}

func (self *StepStartOnHIMN) Cleanup(state multistep.StateBag) {}

func HimnSSHIP(state multistep.StateBag) (string, error) {
	ip := state.Get("himn_ssh_address").(string)
	return ip, nil
}

func HimnSSHPort(state multistep.StateBag) (uint, error) {
	config := state.Get("commonconfig").(CommonConfig)
	return config.SSHPort, nil
}
