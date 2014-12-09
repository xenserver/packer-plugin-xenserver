package xenserver

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"strconv"
)

type stepGetVNCPort struct{}

func (self *stepGetVNCPort) Run(state multistep.StateBag) multistep.StepAction {

	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Step: forward the instances VNC port over SSH")

	domid := state.Get("domid").(string)
	cmd := fmt.Sprintf("xenstore-read /local/domain/%s/console/vnc-port", domid)

	remote_vncport, _ := execute_ssh_cmd(cmd, config.HostIp, "22", config.Username, config.Password)

	remote_port, err := strconv.ParseUint(remote_vncport, 10, 16)

	if err != nil {
		ui.Error(fmt.Sprintf("Unable to convert '%s' to an int", remote_vncport))
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("instance_vnc_port", uint(remote_port))

	return multistep.ActionContinue
}

func (self *stepGetVNCPort) Cleanup(state multistep.StateBag) {
}

func instanceVNCPort(state multistep.StateBag) (uint, error) {
	vncPort := state.Get("instance_vnc_port").(uint)
	return vncPort, nil
}

func instanceVNCIP(state multistep.StateBag) (string, error) {
	// The port is in Dom0, so we want to forward from localhost
	return "127.0.0.1", nil
}
