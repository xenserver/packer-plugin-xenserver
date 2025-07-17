package common

import (
	"context"
	"fmt"
	"time"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepWaitForIP struct {
	VmCleanup
	Chan    <-chan string
	Timeout time.Duration
}

func (self *StepWaitForIP) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	config := state.Get("commonconfig").(CommonConfig)

	ui.Say("Step: Wait for VM's IP to become known to us.")

	uuid := state.Get("instance_uuid").(string)
	instance, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return multistep.ActionHalt
	}

	var ip string
	err = InterruptibleWait{
		Timeout:           self.Timeout,
		PredicateInterval: 5 * time.Second,
		Predicate: func() (result bool, err error) {

			if config.IPGetter == "auto" || config.IPGetter == "http" {

				// Snoop IP from HTTP fetch
				select {
				case ip = <-self.Chan:
					ui.Message(fmt.Sprintf("Got IP '%s' from HTTP request", ip))
					return true, nil
				default:
				}

			}

			if config.IPGetter == "auto" || config.IPGetter == "tools" {

				// Look for PV IP
				m, err := xenapi.VM.GetGuestMetrics(c.session, instance)
				if err != nil {
					return false, err
				}
				if m != "" {
					metrics, err := xenapi.VMGuestMetrics.GetRecord(c.session, m)
					if err != nil {
						return false, err
					}
					networks := metrics.Networks
					var ok bool
					if ip, ok = networks["0/ip"]; ok {
						if ip != "" {
							// Check if it is not a link-local address
							if strings.HasPrefix(ip, "169.254."){
								ui.Message(fmt.Sprintf("Got link-local IP '%s' from XenServer tools, which is not usable for SSH. Waiting for a valid IP address", ip))
							} else {
								ui.Message(fmt.Sprintf("Got IP '%s' from XenServer tools", ip))
								return true, nil
							}
						}
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

func InstanceSSHIP(state multistep.StateBag) (string, error) {
	ip := state.Get("instance_ssh_address").(string)
	return ip, nil
}

func InstanceSSHPort(state multistep.StateBag) (int, error) {
	return 22, nil
}
