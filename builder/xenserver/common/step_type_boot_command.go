package common

/* Heavily borrowed from builder/quemu/step_type_boot_command.go */

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/bootcommand"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/mitchellh/go-vnc"
)

const KeyLeftShift uint32 = 0xFFE1

type bootCommandTemplateData struct {
	Name     string
	HTTPIP   string
	HTTPPort uint
}

type StepTypeBootCommand struct {
	Ctx interpolate.Context
}

func (step *StepTypeBootCommand) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	httpPort := state.Get("http_port").(int)

	// skip this step if we have nothing to type
	if len(config.BootCommand) == 0 {
		return multistep.ActionContinue
	}

	vmRef, err := c.client.VM.GetByNameLabel(c.session, config.VMName)

	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if len(vmRef) != 1 {
		ui.Error(fmt.Sprintf("expected to find a single VM, instead found '%d'. Ensure the VM name is unique", len(vmRef)))
	}

	consoles, err := c.client.VM.GetConsoles(c.session, vmRef[0])
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if len(consoles) != 1 {
		ui.Error(fmt.Sprintf("expected to find a VM console, instead found '%d'. Ensure there is only one console", len(consoles)))
		return multistep.ActionHalt
	}

	location, err := c.client.Console.GetLocation(c.session, consoles[0])
	if err != nil {
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	locationPieces := strings.SplitAfter(location, "/")
	consoleHost := strings.TrimSuffix(locationPieces[2], "/")
	ui.Say("Connecting to VNC over XAPI...")
	log.Printf("Connecting to host: %s", consoleHost)
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:443", consoleHost))

	if err != nil {
		err := fmt.Errorf("Error connecting to VNC: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer conn.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	tlsConn := tls.Client(conn, tlsConfig)

	consoleLocation := strings.TrimSpace(fmt.Sprintf("/%s", locationPieces[len(locationPieces)-1]))
	httpReq := fmt.Sprintf("CONNECT %s HTTP/1.0\r\nHost: %s\r\nCookie: session_id=%s\r\n\r\n", consoleLocation, consoleHost, c.session)
	fmt.Printf("Sending the follow http req: %v", httpReq)

	ui.Message(fmt.Sprintf("Making HTTP request to initiate VNC connection: %s", httpReq))
	_, err = io.WriteString(tlsConn, httpReq)

	if err != nil {
		err := fmt.Errorf("failed to start vnc session: %v", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	buffer := make([]byte, 10000)
	_, err = tlsConn.Read(buffer)
	if err != nil && err != io.EOF {
		err := fmt.Errorf("failed to read vnc session response: %v", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Message(fmt.Sprintf("Received response: %s", string(buffer)))

	vncClient, err := vnc.Client(tlsConn, &vnc.ClientConfig{Exclusive: !config.PackerDebug})

	if err != nil {
		err := fmt.Errorf("Error establishing VNC session: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer vncClient.Close()

	log.Printf("Connected to the VNC console: %s", vncClient.DesktopName)

	// find local ip
	envVar, err := ExecuteHostSSHCmd(state, "echo $SSH_CLIENT")
	if err != nil {
		ui.Error(fmt.Sprintf("Error detecting local IP: %s", err))
		return multistep.ActionHalt
	}
	if envVar == "" {
		ui.Error("Error detecting local IP: $SSH_CLIENT was empty")
		return multistep.ActionHalt
	}
	localIp := strings.Split(envVar, " ")[0]
	ui.Message(fmt.Sprintf("Found local IP: %s", localIp))

	step.Ctx.Data = &bootCommandTemplateData{
		config.VMName,
		localIp,
		uint(httpPort),
	}

	d := bootcommand.NewVNCDriver(vncClient, time.Second/10)

	ui.Say("Typing boot commands over VNC...")
	for _, command := range config.BootCommand {
		if len(command) == 0 {
			continue
		}
		command, err := interpolate.Render(command, &step.Ctx)
		if err != nil {
			err := fmt.Errorf("Error preparing boot command: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		seq, err := bootcommand.GenerateExpressionSequence(command)
		if err != nil {
			err := fmt.Errorf("Error generating boot command: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		if err := seq.Do(ctx, d); err != nil {
			err := fmt.Errorf("Error running boot command: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

	}

	return multistep.ActionContinue
}

func (step *StepTypeBootCommand) Cleanup(multistep.StateBag) {}
