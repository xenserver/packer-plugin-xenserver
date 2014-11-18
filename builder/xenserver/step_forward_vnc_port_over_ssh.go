package xenserver


import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "fmt"
    "strconv"
    "log"
)

type stepForwardVncPortOverSsh struct {}


func (self *stepForwardVncPortOverSsh) Run(state multistep.StateBag) multistep.StepAction {

//    client := state.Get("client").(XenAPIClient)

    config := state.Get("config").(config)
    ui := state.Get("ui").(packer.Ui)

    ui.Say("Step: forward the instances VNC port over SSH")

    domid := state.Get("domid").(string)
    cmd := fmt.Sprintf("xenstore-read /local/domain/%s/console/vnc-port", domid)

    remote_vncport, _ := execute_ssh_cmd(cmd, config.HostIp, "22", config.Username, config.Password)

    remote_port, err := strconv.ParseUint(remote_vncport, 10, 16)
    remote_port_uint := uint(remote_port)
    if err != nil {
        log.Fatal(err.Error())
        log.Fatal(fmt.Sprintf("Unable to convert '%s' to an int", remote_vncport))
        return multistep.ActionHalt
    }

    ui.Say("The VNC port is " + remote_vncport)
    // Just take the min port for the moment
    state.Put("local_vnc_port", config.HostPortMin)
    //local_port := strconv.Itoa(int(config.HostPortMin))
    local_port := config.HostPortMin
    ui.Say(fmt.Sprintf("About to setup SSH Port forward setup on local port %d",local_port))


    go ssh_port_forward(local_port, remote_port_uint, "127.0.0.1", config.HostIp, config.Username, config.Password)
    ui.Say("Port forward setup.")

    return multistep.ActionContinue
}


func (self *stepForwardVncPortOverSsh) Cleanup(state multistep.StateBag) {
}
