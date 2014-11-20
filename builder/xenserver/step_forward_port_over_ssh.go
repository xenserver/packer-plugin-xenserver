package xenserver

import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "log"
    "net"
    "fmt"
)

type stepForwardPortOverSSH struct {

    RemotePort func (state multistep.StateBag) (uint, error)
    RemoteDest func (state multistep.StateBag) (string, error)

    HostPortMin uint
    HostPortMax uint

    ResultKey string
}


func (self *stepForwardPortOverSSH) Run(state multistep.StateBag) multistep.StepAction {

    config := state.Get("config").(config)
    ui := state.Get("ui").(packer.Ui)

    // Find a free local port:

    log.Printf("Looking for an available port between %d and %d",
                self.HostPortMin,
                self.HostPortMax)


    var sshHostPort uint
    var foundPort bool

    foundPort = false

    for i := self.HostPortMin; i < self.HostPortMax; i++ {
        sshHostPort = i
        log.Printf("Trying port: %d", sshHostPort)
        l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", sshHostPort))
        if err == nil {
            l.Close()
            foundPort = true
            break
        }

    }

    if !foundPort {
        log.Fatal("Error: unable to find free host port. Try providing a larger range")
        return multistep.ActionHalt
    }


    ui.Say(fmt.Sprintf("Creating a local port forward over SSH on local port %d", sshHostPort))

    remotePort, _ := self.RemotePort(state)
    remoteDest, _ := self.RemoteDest(state)


    go ssh_port_forward(sshHostPort, remotePort, remoteDest, config.HostIp, config.Username, config.Password)
    ui.Say(fmt.Sprintf("Port forward setup. %d ---> %s:%d on %s", sshHostPort, remoteDest, remotePort, config.HostIp))

    // Provide the local port to future steps.
    state.Put(self.ResultKey, sshHostPort)

    return multistep.ActionContinue
}

func (self *stepForwardPortOverSSH) Cleanup(state multistep.StateBag) {}
