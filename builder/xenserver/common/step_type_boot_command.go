package common

/* Heavily borrowed from builder/quemu/step_type_boot_command.go */

import (
	"fmt"
	"github.com/mitchellh/go-vnc"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
	"net"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const KeyLeftShift uint = 0xFFE1

type bootCommandTemplateData struct {
	Name     string
	HTTPIP   string
	HTTPPort uint
}

type StepTypeBootCommand struct {
	Tpl *packer.ConfigTemplate
}

func (self *StepTypeBootCommand) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	vnc_port := state.Get("local_vnc_port").(uint)
	http_port := state.Get("http_port").(uint)

	// skip this step if we have nothing to type
	if len(config.BootCommand) == 0 {
		return multistep.ActionContinue
	}

	// Connect to the local VNC port as we have set up a SSH port forward
	ui.Say("Connecting to the VM over VNC")
	ui.Message(fmt.Sprintf("Using local port: %d", vnc_port))
	net_conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", vnc_port))

	if err != nil {
		err := fmt.Errorf("Error connecting to VNC: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer net_conn.Close()

	c, err := vnc.Client(net_conn, &vnc.ClientConfig{Exclusive: true})

	if err != nil {
		err := fmt.Errorf("Error establishing VNC session: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer c.Close()

	log.Printf("Connected to the VNC console: %s", c.DesktopName)

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

	tplData := &bootCommandTemplateData{
		config.VMName,
		localIp,
		http_port,
	}

	ui.Say("Typing boot commands over VNC...")
	for _, command := range config.BootCommand {

		command, err := self.Tpl.Process(command, tplData)
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

		vncSendString(c, command)
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
		c.KeyEvent(keyCode, false)

		if keyShift {
			c.KeyEvent(uint32(KeyLeftShift), false)
		}

		// no matter what, wait a small period
		time.Sleep(50 * time.Millisecond)
	}
}
