package xenserver


import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "fmt"
    "code.google.com/p/go.crypto/ssh"
    "bytes"
    "net"
    "io"
    "strconv"
    "log"
    "strings"
)

type stepForwardVncPortOverSsh struct {}


func execute_ssh_cmd (cmd, host, port, username, password string) (stdout string, err error) {
    // Setup connection config
    config := &ssh.ClientConfig {
        User: username,
        Auth: []ssh.AuthMethod {
            ssh.Password(password),
        },
    }

    client, err := ssh.Dial("tcp", host + ":" + port, config)

    if err != nil {
        return "", err
    }

    //Create session
    session, err := client.NewSession()

    if err != nil {
        return "", err
    }

    defer session.Close()

    var b bytes.Buffer
    session.Stdout = &b
    if err := session.Run(cmd); err != nil {
        return "", err
    }

    session.Close()
    return strings.Trim(b.String(), "\n"), nil
}

func forward(local_conn net.Conn, config *ssh.ClientConfig, server, remote_port string) {
    ssh_client_conn, err := ssh.Dial("tcp", server + ":22", config)
    if err != nil {
        log.Fatalf("local ssh.Dial error: %s", err)
    }

    ssh_conn, err := ssh_client_conn.Dial("tcp", "127.0.0.1:" + remote_port) 
    if err != nil {
        log.Fatalf("ssh.Dial error: %s", err)
    }

    go func() {
        _, err = io.Copy(ssh_conn, local_conn)
        if err != nil {
            log.Fatalf("io.copy failed: %v", err)
        }
    }()

    go func() {
        _, err = io.Copy(local_conn, ssh_conn)
        if err != nil {
            log.Fatalf("io.copy failed: %v", err)
        }
    }()
}

func ssh_port_forward(local_port, remote_port, host, username, password string) (err error) {

    config := &ssh.ClientConfig {
        User: username,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
    }

    // Listen on a local port
    local_listener, err := net.Listen("tcp", "127.0.0.1:" + local_port)
    if err != nil {
        log.Fatalf("Local listen failed: %s", err)
        return err
    }

    for {
        local_connection, err := local_listener.Accept()

        if err != nil {
            log.Fatalf("Local accept failed: %s", err)
            return err
        }


        // Forward to a remote port
        go forward(local_connection, config, host, remote_port)
    }



    return nil
}

func (self *stepForwardVncPortOverSsh) Run(state multistep.StateBag) multistep.StepAction {

//    client := state.Get("client").(XenAPIClient)

    config := state.Get("config").(config)
    ui := state.Get("ui").(packer.Ui)

    ui.Say("Step: forward the instances VNC port over SSH")

    domid := state.Get("domid").(string)
    cmd := fmt.Sprintf("xenstore-read /local/domain/%s/console/vnc-port", domid)

    remote_vncport, _ := execute_ssh_cmd(cmd, config.HostIp, "22", config.Username, config.Password)

    ui.Say("The VNC port is " + remote_vncport)
    // Just take the min port for the moment
    state.Put("local_vnc_port", config.VncPortMin)
    local_port := strconv.Itoa(int(config.VncPortMin))
    ui.Say("About to setup SSH Port forward setup on local port " + local_port)


    go ssh_port_forward(local_port, remote_vncport, config.HostIp, config.Username, config.Password)
    ui.Say("Port forward setup.")

    return multistep.ActionContinue
}


func (self *stepForwardVncPortOverSsh) Cleanup(state multistep.StateBag) {
}
