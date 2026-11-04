// Copyright (c) 2026. Citrix Systems, Inc.

package common

/* Heavily borrowed from hashicorp/packer/builder/qemu/step_type_boot_command.go */

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	"xenapi"
)

const KeyLeftShift uint = 0xFFE1

// The UEFI "Press any key to boot from CD" prompt stays open for only about six seconds, at
// an offset that varies with firmware boot time, so the key press is swept across it.
// 20 presses 1 s apart is the tested setting.
const (
	uefiCDPromptGroups   = 20
	uefiCDPromptInterval = time.Second
)

type bootCommandTemplateData struct {
	Name     string
	HTTPIP   string
	HTTPPort uint
}

type StepTypeBootCommand struct {
	Ctx interpolate.Context
	// SweepUEFICDPrompt is set only by builders whose guest boots an installer CD.
	SweepUEFICDPrompt bool
}

// Repeats can reach Windows Boot Manager's menu, where <enter> takes the default.
var sweepKeys = map[string]bool{
	"<enter>":    true,
	"<return>":   true,
	"<spacebar>": true,
}

// sweepableBootCommand allows only a single key: repeating typed text or a selection such as
// "<down><enter>" lands in whatever the first pass started, and a <wait> skews the sweep timing.
func sweepableBootCommand(commands []string) bool {
	return len(commands) == 1 && sweepKeys[commands[0]]
}

// bootKeyGroups treats a zero BootKeyGroups as "no count set", so Prepare must not default it.
func bootKeyGroups(config CommonConfig, firmware string, sweepUEFICDPrompt bool, commands []string) (int, time.Duration) {
	if config.BootKeyGroups != 0 {
		return config.BootKeyGroups, config.BootKeyGroupInterval
	}

	if sweepUEFICDPrompt && firmware == firmwareUEFI && sweepableBootCommand(commands) {
		// The interval is only a default: a build that set one without a count meant it.
		if config.BootKeyGroupInterval != 0 {
			return uefiCDPromptGroups, config.BootKeyGroupInterval
		}
		return uefiCDPromptGroups, uefiCDPromptInterval
	}

	return 1, config.BootKeyGroupInterval
}

func (self *StepTypeBootCommand) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	httpPort := state.Get("http_port").(int)

	if len(config.BootCommand) == 0 {
		return multistep.ActionContinue
	}

	uuid := state.Get("instance_uuid").(string)
	vmRef, err := xenapi.VM.GetByUUID(c.session, uuid)
	if err != nil {
		err = fmt.Errorf("unable to get VM from UUID '%s': %w", uuid, err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	consoles, err := xenapi.VM.GetConsoles(c.session, vmRef)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if len(consoles) != 1 {
		err := fmt.Errorf("expected to find a VM console, instead found %d; ensure there is only one console", len(consoles))
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	location, err := xenapi.Console.GetLocation(c.session, consoles[0])
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	trimmed := strings.TrimSpace(location)
	consoleURL, parseErr := url.Parse(trimmed)
	if parseErr != nil || consoleURL.Hostname() == "" {
		// url.Parse already quotes the location in its error, so only the other message does.
		err := fmt.Errorf("could not read a host and path out of the console location %q", trimmed)
		if parseErr != nil {
			err = fmt.Errorf("could not read a host and path out of the console location: %w", parseErr)
		}
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	consoleHost := consoleURL.Hostname()
	consoleLocation := consoleURL.RequestURI()

	ui.Say(fmt.Sprintf("Connecting to the VM console VNC over xapi via %s", consoleHost))
	vncClient, err := establishConsoleVNC(ctx, state, ui, c, consoleHost, consoleLocation)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer vncClient.Close()

	// Closing the console unblocks a parked write; otherwise Run never returns and no Cleanup runs.
	stopOnCancel := context.AfterFunc(ctx, func() { vncClient.Close() })
	defer stopOnCancel()

	log.Printf("Connected to the VNC console: %s", vncClient.DesktopName)

	localIp, ok := state.Get("http_ip").(string)
	if !ok || localIp == "" {
		err := fmt.Errorf("no HTTP server address is known for this build; StepHTTPIPDiscover must run before the boot command")
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	ui.Message(fmt.Sprintf("Boot command will point the VM at %s", localIp))

	commands, err := renderBootCommand(config.BootCommand, self.Ctx, bootCommandTemplateData{
		Name:     config.Name(),
		HTTPIP:   localIp,
		HTTPPort: uint(httpPort),
	})
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	firmware, _ := state.Get("firmware").(string)
	groups, interval := bootKeyGroups(config, firmware, self.SweepUEFICDPrompt, commands)
	switch {
	case config.BootKeyGroups != 0:
	case groups > 1:
		ui.Message(fmt.Sprintf("UEFI guest: sending the boot command %d times, at least %s apart, "+
			"to cover the \"Press any key to boot from CD\" window. Set boot_key_groups = 1 to "+
			"send it once.", groups, interval))
	case self.SweepUEFICDPrompt && firmware == firmwareUEFI:
		ui.Message("UEFI guest: sending the boot command once. Only a boot_command of a single " +
			"<enter>, <return> or <spacebar> is repeated by default, because every repeat lands " +
			"in whatever the first one started. Set boot_key_groups and boot_key_group_interval " +
			"to sweep this command anyway.")
	}
	if groups > 1 && len(commands) == 1 && commands[0] == "<spacebar>" {
		ui.Message("A repeated <spacebar> stops the countdown of a Windows Boot Manager menu, " +
			"which Windows Server install media can show, and the build then waits on it until " +
			"install_timeout. boot_command = [\"<enter>\"] takes the menu's default instead.")
	}

	ui.Say("Typing boot commands over VNC...")
	action := self.typeCommands(ctx, state, ui, groups, interval, commands, func(cmd string) error {
		return vncSendString(state, vncClient, cmd)
	})

	if action == multistep.ActionContinue {
		ui.Say("Finished typing.")
	}
	return action
}

func (self *StepTypeBootCommand) typeCommands(
	ctx context.Context,
	state multistep.StateBag,
	ui packer.Ui,
	groups int,
	interval time.Duration,
	commands []string,
	send func(string) error,
) multistep.StepAction {
	// Check ctx.Err() too: StateCancelled is put by another goroutine and can lag the console close.
	cancelled := func() bool {
		if ctx.Err() != nil {
			return true
		}
		_, ok := state.GetOk(multistep.StateCancelled)
		return ok
	}

	for group := 0; group < groups; group++ {
		if cancelled() {
			return multistep.ActionHalt
		}

		for _, command := range commands {
			if cancelled() {
				return multistep.ActionHalt
			}

			// Halt: a half-typed boot command would otherwise surface only as an install timeout.
			if err := send(command); err != nil {
				// Ctrl-C closes the console mid-keystroke; leave the cancellation for the builder to report.
				if cancelled() {
					return multistep.ActionHalt
				}
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}

		if group < groups-1 && interval > 0 {
			ui.Message(fmt.Sprintf("Waiting %s before next boot command group...", interval))
			if err := (InterruptibleWait{Timeout: interval}).Wait(state); err != nil {
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}
	}
	return multistep.ActionContinue
}

// renderBootCommand renders every command up front: a half-typed boot command leaves the
// installer unrecoverable. ctx is taken by value so setting Data does not leak to the caller.
func renderBootCommand(commands []string, ctx interpolate.Context, data bootCommandTemplateData) ([]string, error) {
	ctx.Data = &data

	rendered := make([]string, 0, len(commands))
	for i, command := range commands {
		out, err := interpolate.Render(command, &ctx)
		if err != nil {
			return nil, fmt.Errorf("error preparing boot command %d of %d: %w", i+1, len(commands), err)
		}
		rendered = append(rendered, out)
	}

	return rendered, nil
}

func (self *StepTypeBootCommand) Cleanup(multistep.StateBag) {}

// parseWaitToken returns the duration and byte length of a leading <wait...> token, in the
// syntax other Packer builders use: <wait> is one second and a bare number is seconds.
func parseWaitToken(s string) (time.Duration, int, bool) {
	if !strings.HasPrefix(s, "<wait") {
		return 0, 0, false
	}

	end := strings.Index(s, ">")
	if end < 0 {
		return 0, 0, false
	}

	arg := s[len("<wait"):end]
	if arg == "" {
		return time.Second, end + 1, true
	}
	if _, err := strconv.Atoi(arg); err == nil {
		arg += "s"
	}

	d, err := time.ParseDuration(arg)
	if err != nil || d < 0 {
		return 0, 0, false
	}
	return d, end + 1, true
}

func validateBootCommandWaits(commands []string) (errs []error) {
	for i, command := range commands {
		for rest := command; ; {
			at := strings.Index(rest, "<wait")
			if at < 0 {
				break
			}
			rest = rest[at:]

			if _, length, ok := parseWaitToken(rest); ok {
				rest = rest[length:]
				continue
			}

			end := strings.Index(rest, ">")
			if end < 0 {
				break
			}
			errs = append(errs, fmt.Errorf("boot_command %d: '%s' is not a valid wait token and would be "+
				"typed into the guest as text. Use <wait>, a number of seconds such as <wait5>, or a "+
				"duration such as <wait500ms>, <wait2m> or <wait1m30s>", i+1, rest[:end+1]))
			rest = rest[end+1:]
		}
	}
	return errs
}

type keyTyper interface {
	KeyEvent(keysym uint32, down bool) error
}

// Keysym reference: https://github.com/qemu/qemu/blob/master/ui/vnc_keysym.h
var specialKeys = map[string]uint32{
	"<bs>":       0xFF08,
	"<del>":      0xFFFF,
	"<enter>":    0xFF0D,
	"<esc>":      0xFF1B,
	"<f1>":       0xFFBE,
	"<f2>":       0xFFBF,
	"<f3>":       0xFFC0,
	"<f4>":       0xFFC1,
	"<f5>":       0xFFC2,
	"<f6>":       0xFFC3,
	"<f7>":       0xFFC4,
	"<f8>":       0xFFC5,
	"<f9>":       0xFFC6,
	"<f10>":      0xFFC7,
	"<f11>":      0xFFC8,
	"<f12>":      0xFFC9,
	"<return>":   0xFF0D,
	"<tab>":      0xFF09,
	"<up>":       0xFF52,
	"<down>":     0xFF54,
	"<left>":     0xFF51,
	"<right>":    0xFF53,
	"<spacebar>": 0x020,
	"<insert>":   0xFF63,
	"<home>":     0xFF50,
	"<end>":      0xFF57,
	"<pageUp>":   0xFF55,
	"<pageDown>": 0xFF56,
}

// Taken from the QEMU builder plugin, which does not export it.
func vncSendString(state multistep.StateBag, c keyTyper, original string) error {
	special := specialKeys

	shiftedChars := "~!@#$%^&*()_+{}|:\"<>?"

	// Return KeyEvent errors: otherwise a dropped console goes unnoticed until install_timeout.
	send := func(keysym uint32, down bool) error {
		if err := c.KeyEvent(keysym, down); err != nil {
			// go-vnc closes the connection on any read failure without saying why, so name both causes.
			if errors.Is(err, net.ErrClosed) {
				return fmt.Errorf("the console connection closed part-way through the boot "+
					"command: either the VM console channel went away or the VNC client's "+
					"reader rejected a message from it. The guest did not refuse the "+
					"keystroke: %w", err)
			}

			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				return fmt.Errorf("the VM console stopped accepting data part-way through the "+
					"boot command; the cap that ended it is this plugin's own rather than "+
					"anything the guest reported: %w", err)
			}

			return fmt.Errorf("stopped typing the boot command: %w", err)
		}
		return nil
	}

	for len(original) > 0 {
		var keyCode uint32
		keyShift := false

		if d, length, ok := parseWaitToken(original); ok {
			log.Printf("Special code '%s' found, sleeping %s", original[:length], d)
			// A wait can last minutes, so it must not ignore Ctrl-C.
			if err := (InterruptibleWait{Timeout: d}).Wait(state); err != nil {
				return fmt.Errorf("stopped typing the boot command: %w", err)
			}
			original = original[length:]
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

		// Never log typed text: a boot command can carry installer passwords.
		if keyCode == 0 {
			r, size := utf8.DecodeRuneInString(original)
			original = original[size:]
			keyCode = uint32(r)
			keyShift = unicode.IsUpper(r) || strings.ContainsRune(shiftedChars, r)
		}

		if keyShift {
			if err := send(uint32(KeyLeftShift), true); err != nil {
				return err
			}
		}

		if err := send(keyCode, true); err != nil {
			return err
		}
		time.Sleep(time.Second / 10)
		if err := send(keyCode, false); err != nil {
			return err
		}
		time.Sleep(time.Second / 10)

		if keyShift {
			if err := send(uint32(KeyLeftShift), false); err != nil {
				return err
			}
		}

		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
