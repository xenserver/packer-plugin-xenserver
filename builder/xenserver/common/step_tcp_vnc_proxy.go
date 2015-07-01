package common

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"net/textproto"
	"net/url"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/simonfuhrer/go-xenserver-client"
)

type StepTCPVNCProxy struct {
	RemoteConsole func(state multistep.StateBag) (string, error)

	HostPortMin uint
	HostPortMax uint
}

func (s *StepTCPVNCProxy) Run(state multistep.StateBag) multistep.StepAction {

	//config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	consoleLocation := state.Get("instance_console_location").(string)
	consoleuri, err := url.Parse(consoleLocation)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to Parse URL: %s", err.Error()))
		return multistep.ActionHalt
	}

	l, vncHostPort := FindPort(s.HostPortMin, s.HostPortMax)

	if l == nil || vncHostPort == 0 {
		ui.Error("Error: unable to find free host port. Try providing a larger range [host_port_min, host_port_max]")
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Creating a local port forward on local port %d", vncHostPort))

	req := fmt.Sprintf("CONNECT %s?%s&session_id=%s HTTP/1.1", consoleuri.Path, consoleuri.RawQuery, client.Session.(string))

	go start_vnc_proxy(l, req, consoleuri.Host)

	ui.Say(fmt.Sprintf("VNC Proxy listener setup. 127.0.0.1:%d ---> %s:80 with request details %s", vncHostPort, consoleuri.Host, req))

	// Provide the local port to future steps.
	state.Put("local_vnc_port", vncHostPort)
	return multistep.ActionContinue
}

func proxy_connection(local_conn net.Conn, remote_dest, req string) error {
	defer local_conn.Close()
	remote_conn, err := net.Dial("tcp", remote_dest+":80")
	if err != nil {
		log.Printf("local ssh.Dial error: %s", err)
		return err
	}
	defer remote_conn.Close()

	fmt.Fprintf(remote_conn, req+"\r\n\r\n")

	reader := bufio.NewReader(remote_conn)
	tp := textproto.NewReader(reader)
	line, err := tp.ReadLine()
	if err != nil {
		return err
	}
	if line == "HTTP/1.1 200 OK" {
		for {
			line, err := tp.ReadLine()
			if err != nil {
				return err
			}
			if line == "" {
				break
			}
			log.Printf("EE %s\n", line)
		}
	} else {
		err := fmt.Errorf("Error to establish a connection: %s", line)
		//log.Printf("EE %s\n", line)
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		_, err = io.Copy(remote_conn, local_conn)
		if err != nil {
			log.Printf("io.copy failed: %v", err)
		}
		//close(txDone)
		wg.Done()
	}()

	go func() {
		_, err = io.Copy(local_conn, remote_conn)
		if err != nil {
			log.Printf("io.copy failed: %v", err)
		}
		//close(rxDone)
		wg.Done()
	}()

	wg.Wait()

	return nil
}
func start_vnc_proxy(local_listener net.Listener, req, remote_dest string) error {
	for {
		local_connection, err := local_listener.Accept()
		if err != nil {
			log.Printf("Local accept failed: %s", err)
			return err
		}
		// Forward to a remote port
		go proxy_connection(local_connection, remote_dest, req)
	}
}

func (s *StepTCPVNCProxy) Cleanup(state multistep.StateBag) {}
