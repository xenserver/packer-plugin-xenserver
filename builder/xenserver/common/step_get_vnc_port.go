package common

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"strconv"
)

type StepGetVNCPort struct{}

func (self *StepGetVNCPort) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: forward the instances VNC port over SSH")

	domid := state.Get("domid").(string)
	cmd := fmt.Sprintf("xenstore-read /local/domain/%s/console/vnc-port", domid)

	remote_vncport, err := ExecuteHostSSHCmd(state, cmd)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get VNC port (is the VM running?): %s", err.Error()))
		return multistep.ActionHalt
	}

	remote_port, err := strconv.ParseUint(remote_vncport, 10, 16)

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to convert '%s' to an int", remote_vncport))
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("instance_vnc_port", uint(remote_port))

	return multistep.ActionContinue
}

func (self *StepGetVNCPort) Cleanup(state multistep.StateBag) {
}

func InstanceVNCPort(state multistep.StateBag) (uint, error) {
	vncPort := state.Get("instance_vnc_port").(uint)
	return vncPort, nil
}

func InstanceVNCIP(state multistep.StateBag) (string, error) {
	// The port is in Dom0, so we want to forward from localhost
	return "127.0.0.1", nil
}
