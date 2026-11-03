// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

// recordingUi exists because MockUi keeps only the last line of each kind and is not goroutine-safe.
type recordingUi struct {
	packersdk.MockUi

	mu    sync.Mutex
	lines []string
}

func (u *recordingUi) Say(message string)     { u.record(message) }
func (u *recordingUi) Message(message string) { u.record(message) }
func (u *recordingUi) Error(message string)   { u.record(message) }

func (u *recordingUi) record(message string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.lines = append(u.lines, message)
}

func (u *recordingUi) text() string {
	u.mu.Lock()
	defer u.mu.Unlock()

	return strings.Join(u.lines, "\n")
}

func TestShutdownProgress_ReportsOnSchedule(t *testing.T) {
	p := newShutdownProgress(time.Minute, time.Hour)

	for _, elapsed := range []time.Duration{30 * time.Second, 59 * time.Second} {
		if line := p.observe(elapsed, xenapi.VMPowerStateRunning, time.Time{}); line != "" {
			t.Errorf("at %s: reported %q before the first report was due", elapsed, line)
		}
	}
	if p.observe(time.Minute, xenapi.VMPowerStateRunning, time.Time{}) == "" {
		t.Error("at 1m: silent, want the first report")
	}
	if line := p.observe(90*time.Second, xenapi.VMPowerStateRunning, time.Time{}); line != "" {
		t.Errorf("at 1m30s: reported %q between reports", line)
	}
	if p.observe(2*time.Minute, xenapi.VMPowerStateRunning, time.Time{}) == "" {
		t.Error("at 2m: silent, want the second report")
	}
}

func TestShutdownProgress_SkipsReportsItSleptThrough(t *testing.T) {
	p := newShutdownProgress(time.Minute, time.Hour)

	if p.observe(5*time.Minute, xenapi.VMPowerStateRunning, time.Time{}) == "" {
		t.Fatal("at 5m: silent, want a report")
	}
	if line := p.observe(5*time.Minute+time.Second, xenapi.VMPowerStateRunning, time.Time{}); line != "" {
		t.Errorf("at 5m1s: reported %q catching up on the reports it slept through", line)
	}
}

func TestShutdownProgress_ReportNamesElapsedAndPowerState(t *testing.T) {
	p := newShutdownProgress(time.Minute, time.Hour)

	line := p.observe(time.Minute+13*time.Millisecond, xenapi.VMPowerStateRunning, time.Time{})

	for _, want := range []string{"1m0s", "1h0m0s", "Running"} {
		if !strings.Contains(line, want) {
			t.Errorf("report %q does not mention %q", line, want)
		}
	}
}

func TestShutdownProgress_StopsReportingPastTheDeadline(t *testing.T) {
	p := newShutdownProgress(time.Minute, 90*time.Second)

	if line := p.observe(2*time.Minute, xenapi.VMPowerStateRunning, time.Time{}); line != "" {
		t.Errorf("past install_timeout: reported %q", line)
	}
}

func TestShutdownProgress_CountsDomainStarts(t *testing.T) {
	p := newShutdownProgress(time.Minute, time.Hour)
	start := time.Date(2026, 9, 28, 9, 0, 0, 0, time.UTC)

	p.observe(time.Second, xenapi.VMPowerStateRunning, start)
	p.observe(2*time.Second, xenapi.VMPowerStateRunning, start)
	p.observe(3*time.Second, xenapi.VMPowerStateRunning, time.Time{})
	p.observe(4*time.Second, xenapi.VMPowerStateRunning, start.Add(10*time.Minute))

	if p.starts != 2 {
		t.Errorf("starts = %d, want 2", p.starts)
	}
}

func TestShutdownProgress_DomainSummary(t *testing.T) {
	cases := map[int]string{
		0: "unknown",
		1: "started once",
		3: "started 3 times",
	}

	for starts, want := range cases {
		p := newShutdownProgress(time.Minute, time.Hour)
		p.starts = starts

		if summary := p.domainSummary(); !strings.Contains(summary, want) {
			t.Errorf("%d starts: summary = %q, want it to mention %q", starts, summary, want)
		}
	}
}

func TestDomainStartTime_ReadsTheCurrentDomainsStartTime(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_metrics":            "OpaqueRef:metrics",
		"VM_metrics.get_start_time": "20260928T09:00:00Z",
	})

	got := domainStartTime(fake.conn(), "OpaqueRef:vm")

	want := time.Date(2026, 9, 28, 9, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("start time = %s, want %s", got, want)
	}
}

func TestDomainStartTime_IsZeroWhenXAPICannotSay(t *testing.T) {
	cases := map[string]map[string]interface{}{
		"the VM has no metrics object": {
			"VM.get_metrics": nullRef,
		},
		"the metrics reference is empty": {
			"VM.get_metrics": "",
		},
		"the metrics lookup fails": {
			"VM.get_metrics": errors.New("HANDLE_INVALID"),
		},
		"the start time lookup fails": {
			"VM.get_metrics":            "OpaqueRef:metrics",
			"VM_metrics.get_start_time": errors.New("HANDLE_INVALID"),
		},
		"the domain has not started": {
			"VM.get_metrics":            "OpaqueRef:metrics",
			"VM_metrics.get_start_time": "19700101T00:00:00Z",
		},
	}

	for name, results := range cases {
		t.Run(name, func(t *testing.T) {
			fake := newFakeXAPI(t, results)

			if got := domainStartTime(fake.conn(), "OpaqueRef:vm"); !got.IsZero() {
				t.Errorf("start time = %s, want the zero time", got)
			}
		})
	}
}

func TestDomainStartTime_DoesNotChaseANullMetricsRef(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{"VM.get_metrics": nullRef})

	domainStartTime(fake.conn(), "OpaqueRef:vm")

	if fake.wasCalled("VM_metrics.get_start_time") {
		t.Error("the start time was looked up on a null metrics reference")
	}
}

func shutdownStepState(t *testing.T, results map[string]interface{}) (*multistep.BasicStateBag, *recordingUi) {
	t.Helper()

	results["VM.get_by_uuid"] = "OpaqueRef:vm"

	ui := &recordingUi{}
	state := new(multistep.BasicStateBag)
	state.Put("ui", ui)
	state.Put("client", newFakeXAPI(t, results).conn())
	state.Put("instance_uuid", "7a1e1a1e-0000-0000-0000-000000000001")

	return state, ui
}

func TestStepWaitForGuestShutdown_ReportsProgressWhileItWaits(t *testing.T) {
	state, ui := shutdownStepState(t, map[string]interface{}{
		"VM.get_power_state":        "Running",
		"VM.get_metrics":            "OpaqueRef:metrics",
		"VM_metrics.get_start_time": "20260928T09:00:00Z",
	})
	step := &StepWaitForGuestShutdown{
		Timeout:      600 * time.Millisecond,
		pollInterval: 10 * time.Millisecond,
		reportEvery:  50 * time.Millisecond,
	}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt on a guest that never halts", action)
	}

	out := ui.text()
	if !strings.Contains(out, "Waited") {
		t.Errorf("nothing was reported while waiting:\n%s", out)
	}
	if !strings.Contains(out, "The guest did not shut itself down within 600ms (install_timeout): the domain has started once") {
		t.Errorf("the timeout message does not carry the domain count:\n%s", out)
	}
	if !strings.Contains(out, "This VM is about to be destroyed") {
		t.Errorf("the timeout message does not say the VM is going away:\n%s", out)
	}
	if !strings.Contains(out, `Set keep_vm = "always"`) {
		t.Errorf("the timeout message does not say how to keep the VM for diagnosis:\n%s", out)
	}
}

func TestKeepVMAdvice_OnlyAlwaysKeepsTheVMPastThisTimeout(t *testing.T) {
	for _, keepVM := range []string{"never", "on_success", ""} {
		state := new(multistep.BasicStateBag)
		state.Put("commonconfig", CommonConfig{KeepVM: keepVM})

		if advice := keepVMAdvice(state); !strings.Contains(advice, "about to be destroyed") {
			t.Errorf("keep_vm = %q: %q should say the VM is going away", keepVM, advice)
		}
	}

	state := new(multistep.BasicStateBag)
	state.Put("commonconfig", CommonConfig{KeepVM: "always"})

	if advice := keepVMAdvice(state); strings.Contains(advice, "about to be destroyed") {
		t.Errorf(`keep_vm = "always": %q claims the VM is going away`, advice)
	}
}

func TestKeepVMAdvice_SurvivesAStateBagWithoutTheConfig(t *testing.T) {
	if advice := keepVMAdvice(new(multistep.BasicStateBag)); advice == "" {
		t.Error("no advice given when commonconfig is absent")
	}
}

func TestStepWaitForGuestShutdown_ContinuesOnceTheGuestHalts(t *testing.T) {
	state, ui := shutdownStepState(t, map[string]interface{}{
		"VM.get_power_state":        "Halted",
		"VM.get_metrics":            "OpaqueRef:metrics",
		"VM_metrics.get_start_time": "20260928T09:00:00Z",
	})
	step := &StepWaitForGuestShutdown{
		Timeout:      5 * time.Second,
		pollInterval: 10 * time.Millisecond,
		reportEvery:  50 * time.Millisecond,
	}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run = %v, want ActionContinue", action)
	}
	if out := ui.text(); !strings.Contains(out, "The guest has shut itself down") {
		t.Errorf("the guest halting was not reported:\n%s", out)
	}
}

func TestStepWaitForGuestShutdown_HaltsWhenThePowerStateCannotBeRead(t *testing.T) {
	state, ui := shutdownStepState(t, map[string]interface{}{
		"VM.get_power_state": errors.New("HANDLE_INVALID"),
	})
	step := &StepWaitForGuestShutdown{
		Timeout:      5 * time.Second,
		pollInterval: 10 * time.Millisecond,
	}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}
	if out := ui.text(); !strings.Contains(out, "Could not read the VM's power state") {
		t.Errorf("the XAPI failure was not reported:\n%s", out)
	}
}
