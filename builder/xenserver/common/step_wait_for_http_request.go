package common

import (
	"fmt"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepWaitForHTTPRequest struct {
	Chan    <-chan string
	Timeout time.Duration
}

func (self *StepWaitForHTTPRequest) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: Wait for install to complete.")

	timeout := time.After(self.Timeout)
	var ip string
	select {
	case ip = <-self.Chan:
	case <-timeout:
		ui.Error("Timed out. Giving up waiting for installation to complete.")
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Got IP address '%s'", ip))
	state.Put("instance_ssh_address", ip)

	return multistep.ActionContinue
}

func (self *StepWaitForHTTPRequest) Cleanup(state multistep.StateBag) {}

func InstanceSSHIP(state multistep.StateBag) (string, error) {
	ip := state.Get("instance_ssh_address").(string)
	return ip, nil
}

func InstanceSSHPort(state multistep.StateBag) (uint, error) {
	return 22, nil
}
