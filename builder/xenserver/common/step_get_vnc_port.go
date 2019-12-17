package common

import (
	"fmt"
	"strconv"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
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
		ui.Error(fmt.Sprintf("XS7.5/7.6 no longer support xenstore-read: Try to use 5900.  See https://bugs.xenserver.org/browse/XSO-906"))
		remote_vncport = "5900"
		client := state.Get("client").(xsclient.XenAPIClient)
		hosts, err := client.GetHosts()
		if err != nil {
			ui.Error(fmt.Sprintf("Could not retrieve hosts in the pool: %s", err.Error()))
			return multistep.ActionHalt
		}
		host := hosts[0]
		host_software_versions, err := host.GetSoftwareVersion()
		xs_version := host_software_versions["product_version"].(string)

		if err != nil {
			ui.Error(fmt.Sprintf("Could not get the software version: %s", err.Error()))
			return multistep.ActionHalt
		}
		if xs_version > "7.6.0" {
			ui.Say(fmt.Sprintf("XS8.0+ no longer support xenstore-read: Make sure to install socat for XS8.0+, attempting to use socat"))
			expectedport := fmt.Sprintf("59%s", domid)
			cmd1 := fmt.Sprintf("nohup socat -d -d -lf /tmp/socat-%s TCP4-LISTEN:%s,reuseaddr,fork,tcpwrap=socat,allow-table=all UNIX-CONNECT:/var/run/xen/vnc-%s &>/dev/null &", expectedport, expectedport, domid)
			remote_vncport, err := ExecuteHostSSHCmd(state, cmd1)
			if err != nil {
				ui.Say(fmt.Sprintf("socat not available on XenServer, halting packer ..."))
				return multistep.ActionHalt
			}
			ui.Say(fmt.Sprintf("nohup socat -d -d -lf /tmp/socat-%s TCP4-LISTEN:%s,reuseaddr,fork,tcpwrap=socat,allow-table=all UNIX-CONNECT:/var/run/xen/vnc-%s &>/dev/null &", remote_vncport, expectedport, domid))
			remote_vncport = expectedport
		}
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
