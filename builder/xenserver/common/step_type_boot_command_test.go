// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

func freshState() *multistep.BasicStateBag {
	return new(multistep.BasicStateBag)
}

func newStep(t *testing.T, name string) *StepTypeBootCommand {
	t.Helper()
	s := &StepTypeBootCommand{}
	s.Ctx = interpolate.Context{}
	s.Ctx.Data = &bootCommandTemplateData{Name: name, HTTPIP: "10.0.0.1", HTTPPort: 8080}
	return s
}

func TestTypeCommands_SendsOnceForOneGroup(t *testing.T) {
	commands := []string{"a", "b", "c"}

	var got []string
	action := newStep(t, "vm").typeCommands(context.Background(), freshState(), packer.TestUi(t), 1, 0, commands, func(s string) error {
		got = append(got, s)
		return nil
	})

	if action != multistep.ActionContinue {
		t.Fatalf("action = %v, want ActionContinue", action)
	}
	if len(got) != 3 {
		t.Errorf("send called %d times, want 3; got %v", len(got), got)
	}
}

func TestTypeCommands_RepeatsNGroups(t *testing.T) {
	commands := []string{"x", "y"}

	var got []string
	action := newStep(t, "vm").typeCommands(context.Background(), freshState(), packer.TestUi(t), 3, 0, commands, func(s string) error {
		got = append(got, s)
		return nil
	})

	if action != multistep.ActionContinue {
		t.Fatalf("action = %v, want ActionContinue", action)
	}
	want := 6
	if len(got) != want {
		t.Errorf("send called %d times, want %d; got %v", len(got), want, got)
	}
}

func TestTypeCommands_StopsOnCancellation(t *testing.T) {
	commands := []string{"spacebar"}

	state := freshState()
	state.Put(multistep.StateCancelled, true)

	var sendCount int
	action := newStep(t, "vm").typeCommands(context.Background(), state, packer.TestUi(t), 5, 0, commands, func(string) error {
		sendCount++
		return nil
	})

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt", action)
	}
	if sendCount != 0 {
		t.Errorf("send called %d times after cancellation, want 0", sendCount)
	}
}

func TestTypeCommands_SendsWhatItWasGiven(t *testing.T) {
	commands := []string{"name={{ .Name }} ip={{ .HTTPIP }} port={{ .HTTPPort }}"}

	var got []string
	action := newStep(t, "myvm").typeCommands(context.Background(), freshState(), packer.TestUi(t), 1, 0, commands, func(s string) error {
		got = append(got, s)
		return nil
	})

	if action != multistep.ActionContinue {
		t.Fatalf("action = %v, want ActionContinue", action)
	}
	if len(got) != 1 || got[0] != commands[0] {
		t.Errorf("sent command = %q, want %q", got, commands[0])
	}
}

func TestTypeCommands_ZeroGroupsNeverSends(t *testing.T) {
	var sendCount int
	newStep(t, "vm").typeCommands(context.Background(), freshState(), packer.TestUi(t), 0, 0, []string{"k"}, func(string) error { sendCount++; return nil })

	if sendCount != 0 {
		t.Errorf("send called %d times for zero groups, want 0", sendCount)
	}
}

// Prepare has no XenAPI session to read the firmware, so it must leave the count unset.
func TestCommonConfigPrepare_LeavesAnUnsetBootKeyGroupsAtZero(t *testing.T) {
	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.BootKeyGroups != 0 {
		t.Errorf("BootKeyGroups = %d after Prepare, want it left at 0", c.BootKeyGroups)
	}
}

func TestCommonConfigPrepare_BootKeyGroupsRespectedWhenSet(t *testing.T) {
	c := minimalConfig("xva")
	c.BootKeyGroups = 10
	c.BootKeyGroupInterval = time.Second

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.BootKeyGroups != 10 {
		t.Errorf("BootKeyGroups = %d after Prepare, want 10", c.BootKeyGroups)
	}
}

func TestParseWaitToken_Accepts(t *testing.T) {
	tests := []struct {
		in     string
		want   time.Duration
		length int
	}{
		{"<wait>", time.Second, 6},
		{"<wait5>", 5 * time.Second, 7},
		{"<wait10>", 10 * time.Second, 8},
		{"<wait2>", 2 * time.Second, 7},
		{"<wait90>", 90 * time.Second, 8},
		{"<wait0>", 0, 7},
		{"<wait500ms>", 500 * time.Millisecond, 11},
		{"<wait2m>", 2 * time.Minute, 8},
		{"<wait1m30s>", 90 * time.Second, 11},
		{"<wait3>rest<enter>", 3 * time.Second, 7},
	}

	for _, tt := range tests {
		got, length, ok := parseWaitToken(tt.in)
		if !ok {
			t.Errorf("parseWaitToken(%q) = not a wait token, want %s", tt.in, tt.want)
			continue
		}
		if got != tt.want {
			t.Errorf("parseWaitToken(%q) = %s, want %s", tt.in, got, tt.want)
		}
		if length != tt.length {
			t.Errorf("parseWaitToken(%q) consumed %d bytes, want %d", tt.in, length, tt.length)
		}
	}
}

func TestParseWaitToken_Rejects(t *testing.T) {
	for _, in := range []string{"", "hello", "<enter>", "<waitress>", "<wait 5>", "<wait-5>", "<wait5", "<wait5m30>"} {
		if d, length, ok := parseWaitToken(in); ok {
			t.Errorf("parseWaitToken(%q) = %s over %d bytes, want it rejected", in, d, length)
		}
	}
}

func TestValidateBootCommandWaits(t *testing.T) {
	good := []string{"<wait><wait10>linux ks=http://10.0.0.1/ks.cfg<enter>", "<wait1m30s><esc>", "fine<wait>", "no waits here"}
	if errs := validateBootCommandWaits(good); len(errs) > 0 {
		t.Errorf("unexpected errors: %v", errs)
	}

	bad := []string{"<wait2x><enter>", "<wait five>", "<wait><enter><wait2x>", "<wait2x>a<wait3y>"}
	errs := validateBootCommandWaits(bad)
	if len(errs) != 5 {
		t.Fatalf("got %d errors, want 5: %v", len(errs), errs)
	}
	if !strings.Contains(errs[0].Error(), "<wait2x>") {
		t.Errorf("error %q should name the offending token", errs[0])
	}
}

type keyEvent struct {
	keysym uint32
	down   bool
}

type fakeTyper struct {
	events []keyEvent
	err    error
}

func (f *fakeTyper) KeyEvent(keysym uint32, down bool) error {
	if f.err != nil {
		return f.err
	}
	f.events = append(f.events, keyEvent{keysym, down})
	return nil
}

func TestVncSendString_ReportsADroppedConsole(t *testing.T) {
	typer := &fakeTyper{err: errors.New("connection reset")}

	err := vncSendString(freshState(), typer, "ab")

	if err == nil {
		t.Fatal("a failed KeyEvent must reach the caller, got nil")
	}
	if !strings.Contains(err.Error(), "connection reset") {
		t.Errorf("error = %q, want it to carry the console's own error", err)
	}
}

func TestTypeCommands_HaltsAndRecordsTheErrorWhenSendingFails(t *testing.T) {
	state := freshState()

	var sent int
	action := newStep(t, "vm").typeCommands(context.Background(), state, packer.TestUi(t), 2, 0, []string{"a", "b"}, func(string) error {
		sent++
		return errors.New("connection reset")
	})

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt", action)
	}
	if sent != 1 {
		t.Errorf("send called %d times after a failure, want 1", sent)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Error("the build's error must be in the state bag, or Run reports \"Build was halted.\"")
	}
}

func TestTypeCommands_LeavesACancellationToRun(t *testing.T) {
	state := freshState()

	action := newStep(t, "vm").typeCommands(context.Background(), state, packer.TestUi(t), 1, 0, []string{"a"}, func(string) error {
		state.Put(multistep.StateCancelled, true)
		return InterruptedError{}
	})

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); ok {
		t.Error("a cancelled build was reported as a boot command failure")
	}
}

func TestTypeCommands_LeavesACancellationToRunBeforeStateCancelledIsSet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	state := freshState()

	action := newStep(t, "vm").typeCommands(ctx, state, packer.TestUi(t), 1, 0, []string{"a"}, func(string) error {
		cancel()
		return fmt.Errorf("write tcp 10.0.0.1:443: %w", net.ErrClosed)
	})

	if action != multistep.ActionHalt {
		t.Fatalf("action = %v, want ActionHalt", action)
	}
	if err, ok := state.GetOk("error"); ok {
		t.Errorf("state carries %v, want the cancellation left for the builder to report", err)
	}
}

func TestVncSendString_ConsumesTheWaitAndTypesTheRest(t *testing.T) {
	typer := &fakeTyper{}

	if err := vncSendString(freshState(), typer, "<wait1ms>a"); err != nil {
		t.Fatalf("vncSendString: %v", err)
	}

	want := []keyEvent{{'a', true}, {'a', false}}
	if !reflect.DeepEqual(typer.events, want) {
		t.Errorf("key events = %v, want %v", typer.events, want)
	}
}

func TestVncSendString_StopsAtACancelledBuild(t *testing.T) {
	state := freshState()
	state.Put(multistep.StateCancelled, true)
	typer := &fakeTyper{}

	if err := vncSendString(state, typer, "<wait1h>a"); err == nil {
		t.Error("a cancelled build must be reported to the caller, got nil")
	}

	if len(typer.events) != 0 {
		t.Errorf("typed %v after cancellation, want nothing", typer.events)
	}
}

func TestCommonConfigPrepare_RejectsAMalformedWaitToken(t *testing.T) {
	c := minimalConfig("xva")
	c.BootCommand = []string{"<wait2s><enter>", "<wait2seconds><enter>"}

	_, errs := prepare(t, c)

	if len(errs) != 1 {
		t.Fatalf("got %d errors, want 1: %v", len(errs), errs)
	}
	if !strings.Contains(errs[0].Error(), "<wait2seconds>") {
		t.Errorf("error %q should name the offending token", errs[0])
	}
}

func TestRenderBootCommand_SubstitutesTheDiscoveredHTTPAddress(t *testing.T) {
	commands := []string{
		"linux ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks.cfg",
		"hostname={{ .Name }}<enter>",
	}
	data := bootCommandTemplateData{Name: "packer-build", HTTPIP: "10.11.12.13", HTTPPort: 8042}

	rendered, err := renderBootCommand(commands, interpolate.Context{}, data)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if want := "linux ks=http://10.11.12.13:8042/ks.cfg"; rendered[0] != want {
		t.Errorf("rendered[0] = %q, want %q", rendered[0], want)
	}
	if want := "hostname=packer-build<enter>"; rendered[1] != want {
		t.Errorf("rendered[1] = %q, want %q", rendered[1], want)
	}
}

func TestRenderBootCommand_ReportsATemplateErrorBeforeTypingAnything(t *testing.T) {
	commands := []string{"fine", "{{ .NoSuchField }}"}

	rendered, err := renderBootCommand(commands, interpolate.Context{}, bootCommandTemplateData{})

	if err == nil {
		t.Fatal("expected an error for an unknown template field")
	}
	if rendered != nil {
		t.Errorf("no commands should be returned on failure, got %v", rendered)
	}
	if !strings.Contains(err.Error(), "boot command 2 of 2") {
		t.Errorf("error %q should identify which boot command failed", err)
	}
}

func TestRenderBootCommand_LeavesTheCallersContextAlone(t *testing.T) {
	ctx := interpolate.Context{}

	if _, err := renderBootCommand([]string{"x"}, ctx, bootCommandTemplateData{HTTPIP: "10.0.0.1"}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ctx.Data != nil {
		t.Errorf("the caller's context was mutated: Data = %v", ctx.Data)
	}
}

func TestBootKeyGroups_Resolution(t *testing.T) {
	dismiss := []string{"<spacebar>"}

	for name, tc := range map[string]struct {
		config       CommonConfig
		firmware     string
		sweep        bool
		commands     []string
		wantGroups   int
		wantInterval time.Duration
	}{
		"a UEFI iso build with no count sweeps the CD prompt": {
			firmware: firmwareUEFI, sweep: true, commands: dismiss,
			wantGroups: uefiCDPromptGroups, wantInterval: uefiCDPromptInterval,
		},
		"a single <enter> is swept too": {
			firmware: firmwareUEFI, sweep: true, commands: []string{"<enter>"},
			wantGroups: uefiCDPromptGroups, wantInterval: uefiCDPromptInterval,
		},
		"an explicit 1 sends once even on UEFI": {
			config: CommonConfig{BootKeyGroups: 1}, firmware: firmwareUEFI, sweep: true, commands: dismiss,
			wantGroups: 1,
		},
		"an explicit count wins over the sweep": {
			config:   CommonConfig{BootKeyGroups: 3, BootKeyGroupInterval: 2 * time.Second},
			firmware: firmwareUEFI, sweep: true, commands: dismiss,
			wantGroups: 3, wantInterval: 2 * time.Second,
		},
		"an interval named without a count survives the sweep": {
			config:   CommonConfig{BootKeyGroupInterval: 3 * time.Second},
			firmware: firmwareUEFI, sweep: true, commands: dismiss,
			wantGroups: uefiCDPromptGroups, wantInterval: 3 * time.Second,
		},
		"a BIOS iso build still sends once": {
			firmware: firmwareBIOS, sweep: true, commands: dismiss,
			wantGroups: 1,
		},
		"an unreadable firmware still sends once": {
			sweep: true, commands: dismiss,
			wantGroups: 1,
		},
		"a builder that boots no installer CD never sweeps": {
			firmware: firmwareUEFI, sweep: false, commands: dismiss,
			wantGroups: 1,
		},
		"an explicit count is honoured where there is no sweep": {
			config: CommonConfig{BootKeyGroups: 4, BootKeyGroupInterval: time.Second},
			sweep:  false, commands: dismiss,
			wantGroups: 4, wantInterval: time.Second,
		},
		"a command that edits a boot menu is never swept": {
			firmware: firmwareUEFI, sweep: true,
			commands: []string{"<wait><esc><wait>", "e<wait>", "<down><down><down><end>",
				" autoinstall ds=nocloud-net;s=http://10.0.0.1:8000/", "<wait><f10>"},
			wantGroups: 1,
		},
		"a command carrying a wait is never swept": {
			firmware: firmwareUEFI, sweep: true, commands: []string{"<spacebar><wait10>"},
			wantGroups: 1,
		},
		"a command that picks a menu entry is never swept": {
			firmware: firmwareUEFI, sweep: true, commands: []string{"<down><enter>"},
			wantGroups: 1,
		},
		"a command that is the empty string is never swept": {
			firmware: firmwareUEFI, sweep: true, commands: []string{""},
			wantGroups: 1,
		},
		"an empty boot command is never swept": {
			firmware: firmwareUEFI, sweep: true,
			wantGroups: 1,
		},
		"an explicit zero is the same as unset": {
			config: CommonConfig{BootKeyGroups: 0}, firmware: firmwareUEFI, sweep: true, commands: dismiss,
			wantGroups: uefiCDPromptGroups, wantInterval: uefiCDPromptInterval,
		},
	} {
		groups, interval := bootKeyGroups(tc.config, tc.firmware, tc.sweep, tc.commands)

		if groups != tc.wantGroups {
			t.Errorf("%s: groups = %d, want %d", name, groups, tc.wantGroups)
		}
		if interval != tc.wantInterval {
			t.Errorf("%s: interval = %s, want %s", name, interval, tc.wantInterval)
		}
	}
}

func TestSweepableBootCommand(t *testing.T) {
	for name, tc := range map[string]struct {
		commands []string
		want     bool
	}{
		"a single <spacebar>": {[]string{"<spacebar>"}, true},
		"a single <enter>":    {[]string{"<enter>"}, true},
		"a single <return>":   {[]string{"<return>"}, true},

		"a menu entry chosen and committed": {[]string{"<down><enter>"}, false},
		"a dismissal and a commit":          {[]string{"<spacebar><enter>"}, false},
		"keys spread over elements":         {[]string{"<esc>", "<f10>"}, false},

		"a typed character":               {[]string{"a"}, false},
		"a key and a typed character":     {[]string{"<esc>e"}, false},
		"a kernel argument":               {[]string{"<tab> text ks=http://x/ks.cfg<enter>"}, false},
		"a wait, which costs time to run": {[]string{"<spacebar><wait10>"}, false},
		"a bare wait":                     {[]string{"<wait>"}, false},
		"an unknown token, typed as text": {[]string{"<nosuchkey>"}, false},
		"an unterminated token":           {[]string{"<spacebar"}, false},

		"an element that is the empty string": {[]string{""}, false},
		"no commands at all":                  {nil, false},
	} {
		if got := sweepableBootCommand(tc.commands); got != tc.want {
			t.Errorf("%s: sweepableBootCommand(%q) = %v, want %v", name, tc.commands, got, tc.want)
		}
	}
}

func TestSweepKeys_AreKeysTheSenderKnows(t *testing.T) {
	for key := range sweepKeys {
		if _, ok := specialKeys[key]; !ok {
			t.Errorf("%s is swept as a key press but the sender types it as text", key)
		}
	}
}

func TestSweepKeys_ExcludeEveryOtherKey(t *testing.T) {
	swept := map[string]bool{"<enter>": true, "<return>": true, "<spacebar>": true}

	for token := range specialKeys {
		if got := sweepableBootCommand([]string{token}); got != swept[token] {
			t.Errorf("sweepableBootCommand([%q]) = %v, want %v", token, got, swept[token])
		}
	}
}

func TestBootKeyGroups_SweepOutlastsTheCDPromptWindow(t *testing.T) {
	groups, interval := bootKeyGroups(CommonConfig{}, firmwareUEFI, true, []string{"<spacebar>"})

	if span := time.Duration(groups-1) * interval; span < 15*time.Second {
		t.Errorf("the sweep covers %s, which is too short to cover a ~6s window at an "+
			"offset that varies by host", span)
	}
}

func TestCommonConfigPrepare_RejectsRepetitionsWithNoInterval(t *testing.T) {
	c := minimalConfig("xva")
	c.BootKeyGroups = 20

	_, errs := prepare(t, c)

	if !errorsContain(errs, "boot_key_group_interval") {
		t.Errorf("boot_key_groups without an interval should be rejected, got %v", errs)
	}
}

func TestCommonConfigPrepare_AcceptsOneGroupWithNoInterval(t *testing.T) {
	c := minimalConfig("xva")
	c.BootKeyGroups = 1

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Errorf("one group needs no interval, got %v", errs)
	}
}

// Expected keysyms are spelled out: ranging over specialKeys would pass with two values swapped.
func TestVncSendString_PressesEverySpecialKey(t *testing.T) {
	keysyms := map[string]uint32{
		"<bs>": 0xFF08, "<del>": 0xFFFF, "<enter>": 0xFF0D, "<esc>": 0xFF1B,
		"<f1>": 0xFFBE, "<f2>": 0xFFBF, "<f3>": 0xFFC0, "<f4>": 0xFFC1,
		"<f5>": 0xFFC2, "<f6>": 0xFFC3, "<f7>": 0xFFC4, "<f8>": 0xFFC5,
		"<f9>": 0xFFC6, "<f10>": 0xFFC7, "<f11>": 0xFFC8, "<f12>": 0xFFC9,
		"<return>": 0xFF0D, "<tab>": 0xFF09, "<up>": 0xFF52, "<down>": 0xFF54,
		"<left>": 0xFF51, "<right>": 0xFF53, "<spacebar>": 0x020, "<insert>": 0xFF63,
		"<home>": 0xFF50, "<end>": 0xFF57, "<pageUp>": 0xFF55, "<pageDown>": 0xFF56,
	}

	if len(specialKeys) != len(keysyms) {
		t.Fatalf("specialKeys has %d tokens, this test knows %d", len(specialKeys), len(keysyms))
	}

	for token, keysym := range keysyms {
		t.Run(token, func(t *testing.T) {
			typer := &fakeTyper{}
			if err := vncSendString(freshState(), typer, token); err != nil {
				t.Fatalf("vncSendString(%q) returned %v, want nil", token, err)
			}

			want := []keyEvent{{keysym, true}, {keysym, false}}
			if !reflect.DeepEqual(typer.events, want) {
				t.Errorf("%q typed %v, want %v", token, typer.events, want)
			}
		})
	}
}

func TestVncSendString_SaysTheConsoleClosedWithoutBlamingTheGuest(t *testing.T) {
	typer := &fakeTyper{err: fmt.Errorf("write tcp 10.0.0.1:443: %w", net.ErrClosed)}

	err := vncSendString(freshState(), typer, "a")
	if err == nil {
		t.Fatal("vncSendString returned nil, want an error")
	}

	for _, want := range []string{"console channel went away", "rejected a message", "did not refuse"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error = %q, want it to mention %q", err, want)
		}
	}
	if strings.Contains(err.Error(), "stopped typing the boot command") {
		t.Errorf("error = %q, want the closed-console branch rather than the generic wrapper", err)
	}
	if !errors.Is(err, net.ErrClosed) {
		t.Errorf("error = %q, want the cause still reachable through the chain", err)
	}
}

func TestVncSendString_NamesItsOwnCapWhenTheConsoleStalls(t *testing.T) {
	typer := &fakeTyper{err: &net.OpError{Op: "write", Err: os.ErrDeadlineExceeded}}

	err := vncSendString(freshState(), typer, "a")
	if err == nil {
		t.Fatal("vncSendString returned nil, want an error")
	}

	if !strings.Contains(err.Error(), "this plugin's own") {
		t.Errorf("error = %q, want it to own the cap rather than blame the guest", err)
	}
	if strings.Contains(err.Error(), "stopped typing the boot command") {
		t.Errorf("error = %q, want the timeout branch rather than the generic wrapper", err)
	}
	if !errors.Is(err, os.ErrDeadlineExceeded) {
		t.Errorf("error = %q, want the cause still reachable through the chain", err)
	}
}

func runWithConsoleLocation(t *testing.T, location string) (multistep.StepAction, multistep.StateBag) {
	t.Helper()

	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_by_uuid":  "OpaqueRef:vm",
		"VM.get_consoles": []interface{}{"OpaqueRef:console"},
		// Lowercase: that is how the generated SDK spells this class.
		"console.get_location": location,
	})

	state := freshState()
	state.Put("commonconfig", CommonConfig{BootCommand: []string{"a"}})
	state.Put("ui", packer.TestUi(t))
	state.Put("client", fake.conn())
	state.Put("http_port", 8080)
	state.Put("instance_uuid", "abc")

	return newStep(t, "vm").Run(context.Background(), state), state
}

func TestStepTypeBootCommand_ReportsAConsoleLocationWithNoHost(t *testing.T) {
	action, state := runWithConsoleLocation(t, "/console?uuid=abc")

	if action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}

	err, _ := state.Get("error").(error)
	if err == nil {
		t.Fatal("Run recorded no error")
	}
	if !strings.Contains(err.Error(), "/console?uuid=abc") {
		t.Errorf("error = %q, want it to name the location it could not read", err)
	}
	if strings.Contains(err.Error(), "%!w") {
		t.Errorf("error = %q, want no verb for an absent cause", err)
	}
}

func TestStepTypeBootCommand_CarriesTheParseFailureIntoTheReport(t *testing.T) {
	action, state := runWithConsoleLocation(t, "https://host\x7f/console")

	if action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}

	err, _ := state.Get("error").(error)
	if err == nil {
		t.Fatal("Run recorded no error")
	}

	var urlErr *url.Error
	if !errors.As(err, &urlErr) {
		t.Errorf("error = %q, want url.Parse's own failure still in the chain", err)
	}
	// url.Error quotes the location itself, so a base message that names it too says it twice.
	if n := strings.Count(err.Error(), "/console"); n != 1 {
		t.Errorf("error = %q names the location %d times, want 1", err, n)
	}
}
