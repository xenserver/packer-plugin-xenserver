// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"testing"
	"time"
)

func TestInstallWaitStep_WaitsForTheGuestToPowerItselfOffWithoutACommunicator(t *testing.T) {
	step := InstallWaitStep("none", make(chan string, 1), 42*time.Minute)

	halt, ok := step.(*StepWaitForGuestShutdown)
	if !ok {
		t.Fatalf("communicator \"none\": got %T, want *StepWaitForGuestShutdown", step)
	}
	if halt.Timeout != 42*time.Minute {
		t.Errorf("Timeout = %s, want 42m", halt.Timeout)
	}
}

func TestInstallWaitStep_WaitsForAnIPForEveryRealCommunicator(t *testing.T) {
	for _, commType := range []string{"ssh", "winrm"} {
		ch := make(chan string, 1)

		step := InstallWaitStep(commType, ch, 42*time.Minute)

		waitForIP, ok := step.(*StepWaitForIP)
		if !ok {
			t.Fatalf("communicator %q: got %T, want *StepWaitForIP", commType, step)
		}
		if waitForIP.Timeout != 42*time.Minute {
			t.Errorf("communicator %q: Timeout = %s, want 42m", commType, waitForIP.Timeout)
		}
		if waitForIP.Chan != (<-chan string)(ch) {
			t.Errorf("communicator %q: the HTTP snoop channel was not passed through", commType)
		}
	}
}
