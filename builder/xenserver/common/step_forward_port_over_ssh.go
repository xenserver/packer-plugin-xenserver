package common

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type StepForwardPortOverSSH struct {
	RemotePort func(state multistep.StateBag) (uint, error)
	RemoteDest func(state multistep.StateBag) (string, error)

	HostPortMin uint
	HostPortMax uint

	ResultKey string
}

func (s *StepForwardPortOverSSH) Run(state multistep.StateBag) multistep.StepAction {

	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	// Find a free local port:

	l, sshHostPort := FindPort(s.HostPortMin, s.HostPortMax)

	if l == nil || sshHostPort == 0 {
		ui.Error("Error: unable to find free host port. Try providing a larger range [host_port_min, host_port_max]")
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Creating a local port forward over SSH on local port %d", sshHostPort))

	remotePort, _ := s.RemotePort(state)
	remoteDest, _ := s.RemoteDest(state)

	go ssh_port_forward(l, remotePort, remoteDest, config.HostIp, config.Username, config.Password)
	ui.Say(fmt.Sprintf("Port forward setup. %d ---> %s:%d on %s", sshHostPort, remoteDest, remotePort, config.HostIp))

	// Provide the local port to future steps.
	state.Put(s.ResultKey, sshHostPort)

	return multistep.ActionContinue
}

func (s *StepForwardPortOverSSH) Cleanup(state multistep.StateBag) {}
