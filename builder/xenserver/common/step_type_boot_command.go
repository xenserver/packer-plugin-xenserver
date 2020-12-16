package common

/* Heavily borrowed from builder/quemu/step_type_boot_command.go */

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/mitchellh/go-vnc"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

const KeyLeftShift uint = 0xFFE1

type bootCommandTemplateData struct {
	Name     string
	HTTPIP   string
	HTTPPort uint
}

type StepTypeBootCommand struct {
	Ctx interpolate.Context
}

func (self *StepTypeBootCommand) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	httpPort := state.Get("http_port").(uint)

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

	ui.Say("Connecting to the VM console VNC over xapi")
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:443", config.HostIp))

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

	locationPieces := strings.SplitAfter(location, "/")
	consoleLocation := strings.TrimSpace(fmt.Sprintf("/%s", locationPieces[len(locationPieces)-1]))
	httpReq := fmt.Sprintf("CONNECT %s HTTP/1.0\r\nCookie: session_id=%s\r\n\r\n", consoleLocation, c.session)
	fmt.Printf("Sending the follow http req: %v", httpReq)

	ui.Say(fmt.Sprintf("Making HTTP request to initiate VNC connection: %s", httpReq))
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

	ui.Say(fmt.Sprintf("Received response: %s", string(buffer)))

	vncClient, err := vnc.Client(tlsConn, &vnc.ClientConfig{Exclusive: true})

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

	self.Ctx.Data = &bootCommandTemplateData{
		config.VMName,
		localIp,
		httpPort,
	}

	ui.Say("Typing boot commands over VNC...")
	for _, command := range config.BootCommand {

		command, err := interpolate.Render(command, &self.Ctx)
		if err != nil {
			err := fmt.Errorf("Error preparing boot command: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		// Check for interrupts
		if _, ok := state.GetOk(multistep.StateCancelled); ok {
			return multistep.ActionHalt
		}

		vncSendString(vncClient, command)
	}

	ui.Say("Finished typing.")

	return multistep.ActionContinue
}

func (self *StepTypeBootCommand) Cleanup(multistep.StateBag) {}

// Taken from qemu's builder plugin - not an exported function.
func vncSendString(c *vnc.ClientConn, original string) {
	// Scancodes reference: https://github.com/qemu/qemu/blob/master/ui/vnc_keysym.h
	special := make(map[string]uint32)
	special["<bs>"] = 0xFF08
	special["<del>"] = 0xFFFF
	special["<enter>"] = 0xFF0D
	special["<esc>"] = 0xFF1B
	special["<f1>"] = 0xFFBE
	special["<f2>"] = 0xFFBF
	special["<f3>"] = 0xFFC0
	special["<f4>"] = 0xFFC1
	special["<f5>"] = 0xFFC2
	special["<f6>"] = 0xFFC3
	special["<f7>"] = 0xFFC4
	special["<f8>"] = 0xFFC5
	special["<f9>"] = 0xFFC6
	special["<f10>"] = 0xFFC7
	special["<f11>"] = 0xFFC8
	special["<f12>"] = 0xFFC9
	special["<return>"] = 0xFF0D
	special["<tab>"] = 0xFF09
	special["<up>"] = 0xFF52
	special["<down>"] = 0xFF54
	special["<left>"] = 0xFF51
	special["<right>"] = 0xFF53
	special["<spacebar>"] = 0x020
	special["<insert>"] = 0xFF63
	special["<home>"] = 0xFF50
	special["<end>"] = 0xFF57
	special["<pageUp>"] = 0xFF55
	special["<pageDown>"] = 0xFF56

	shiftedChars := "~!@#$%^&*()_+{}|:\"<>?"

	// TODO(mitchellh): Ripe for optimizations of some point, perhaps.
	for len(original) > 0 {
		var keyCode uint32
		keyShift := false

		if strings.HasPrefix(original, "<wait>") {
			log.Printf("Special code '<wait>' found, sleeping one second")
			time.Sleep(1 * time.Second)
			original = original[len("<wait>"):]
			continue
		}

		if strings.HasPrefix(original, "<wait5>") {
			log.Printf("Special code '<wait5>' found, sleeping 5 seconds")
			time.Sleep(5 * time.Second)
			original = original[len("<wait5>"):]
			continue
		}

		if strings.HasPrefix(original, "<wait10>") {
			log.Printf("Special code '<wait10>' found, sleeping 10 seconds")
			time.Sleep(10 * time.Second)
			original = original[len("<wait10>"):]
			continue
		}

		for specialCode, specialValue := range special {
			if strings.HasPrefix(original, specialCode) {
				log.Printf("Special code '%s' found, replacing with: %d", specialCode, specialValue)
				keyCode = specialValue
				original = original[len(specialCode):]
				break
			}
		}

		if keyCode == 0 {
			r, size := utf8.DecodeRuneInString(original)
			original = original[size:]
			keyCode = uint32(r)
			keyShift = unicode.IsUpper(r) || strings.ContainsRune(shiftedChars, r)

			log.Printf("Sending char '%c', code %d, shift %v", r, keyCode, keyShift)
		}

		if keyShift {
			c.KeyEvent(uint32(KeyLeftShift), true)
		}

		c.KeyEvent(keyCode, true)
		time.Sleep(time.Second / 10)
		c.KeyEvent(keyCode, false)
		time.Sleep(time.Second / 10)

		if keyShift {
			c.KeyEvent(uint32(KeyLeftShift), false)
		}

		// no matter what, wait a small period
		time.Sleep(50 * time.Millisecond)
	}
}
