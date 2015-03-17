package common

import (
	"fmt"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/nilshell/xmlrpc"
	xsclient "github.com/xenserver/go-xenserver-client"
)

type StepWaitForIP struct {
	Chan    <-chan string
	Timeout time.Duration
}

func (self *StepWaitForIP) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	ui.Say("Step: Wait for VM's IP to become known to us.")

	uuid := state.Get("instance_uuid").(string)
	instance, err := client.GetVMByUuid(uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	var ip string
	err = InterruptibleWait{
		Timeout:           self.Timeout,
		PredicateInterval: 5 * time.Second,
		Predicate: func() (result bool, err error) {
			// first check if we got any HTTP requests
			select {
			case ip = <-self.Chan:
				ui.Message(fmt.Sprintf("Got IP '%s' from HTTP request", ip))
				return true, nil
			default:
			}

			// failing that, look for PV IP
			metrics, err := instance.GetGuestMetrics()
			if err != nil {
				return false, err
			}
			if metrics != nil {
				networks := metrics["networks"].(xmlrpc.Struct)
				if ipRaw, ok := networks["0/ip"]; ok {
					if ip = ipRaw.(string); ip != "" {
						ui.Message(fmt.Sprintf("Got IP '%s' from XenServer tools", ip))
						return true, nil
					}
				}
			}

			return false, nil
		},
	}.Wait(state)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get IP address of VM: %s", err.Error()))
		// @todo: give advice on what went wrong (no HTTP server? no PV drivers?)
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Got IP address '%s'", ip))
	state.Put("instance_ssh_address", ip)

	return multistep.ActionContinue
}

func (self *StepWaitForIP) Cleanup(state multistep.StateBag) {}

func InstanceSSHIP(state multistep.StateBag) (string, error) {
	ip := state.Get("instance_ssh_address").(string)
	return ip, nil
}

func InstanceSSHPort(state multistep.StateBag) (uint, error) {
	return 22, nil
}
