// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type fakeCommunicator struct {
	started    []string
	exitStatus int
	startErr   error
	neverExits bool
}

func (f *fakeCommunicator) Start(_ context.Context, cmd *packersdk.RemoteCmd) error {
	f.started = append(f.started, cmd.Command)
	if f.startErr != nil {
		return f.startErr
	}
	if f.neverExits {
		return nil
	}
	cmd.SetExited(f.exitStatus)
	return nil
}

func (f *fakeCommunicator) Upload(string, io.Reader, *os.FileInfo) error { return nil }
func (f *fakeCommunicator) UploadDir(string, string, []string) error     { return nil }
func (f *fakeCommunicator) Download(string, io.Writer) error             { return nil }
func (f *fakeCommunicator) DownloadDir(string, string, []string) error   { return nil }

func shutdownState(t *testing.T, comm packersdk.Communicator) *multistep.BasicStateBag {
	t.Helper()
	state := new(multistep.BasicStateBag)
	if comm != nil {
		state.Put("communicator", comm)
	}
	return state
}

func TestRunInGuest_RunsTheCommandOverTheCommunicator(t *testing.T) {
	comm := &fakeCommunicator{}
	state := shutdownState(t, comm)

	err := runInGuest(context.Background(), state, packersdk.TestUi(t), "shutdown -h now")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(comm.started) != 1 || comm.started[0] != "shutdown -h now" {
		t.Errorf("commands started = %v, want the configured shutdown command", comm.started)
	}
}

func TestRunInGuest_TreatsADisconnectAsTheCommandWorking(t *testing.T) {
	comm := &fakeCommunicator{exitStatus: packersdk.CmdDisconnect}
	state := shutdownState(t, comm)

	if err := runInGuest(context.Background(), state, packersdk.TestUi(t), "poweroff"); err != nil {
		t.Errorf("a disconnect is the expected outcome of a shutdown command, got %v", err)
	}
}

func TestRunInGuest_ReportsANonZeroExitStatus(t *testing.T) {
	comm := &fakeCommunicator{exitStatus: 127}
	state := shutdownState(t, comm)

	err := runInGuest(context.Background(), state, packersdk.TestUi(t), "no-such-command -h now")

	if err == nil {
		t.Fatal("a failing shutdown command must be reported")
	}
	if !strings.Contains(err.Error(), "127") {
		t.Errorf("error %q should carry the exit status", err)
	}
}

func TestRunInGuest_ReportsAFailureToStart(t *testing.T) {
	comm := &fakeCommunicator{startErr: errors.New("connection refused")}
	state := shutdownState(t, comm)

	err := runInGuest(context.Background(), state, packersdk.TestUi(t), "poweroff")

	if err == nil {
		t.Fatal("a command that cannot be started must be reported")
	}
	if !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("error %q should keep the underlying cause", err)
	}
}

func TestRunShutdownCommand_IsBoundedByTheGracefulShutdownBudget(t *testing.T) {
	comm := &fakeCommunicator{neverExits: true}
	state := shutdownState(t, comm)

	const budget = 200 * time.Millisecond
	start := time.Now()
	done := make(chan time.Duration, 1)
	go func() {
		done <- runShutdownCommand(context.Background(), state, packersdk.TestUi(t), "shutdown -h now", budget)
	}()

	var remaining time.Duration
	select {
	case remaining = <-done:
	case <-time.After(10 * time.Second):
		// Generous: the point is that it returns at all, not that it is prompt.
		t.Fatal("a shutdown command that never returns was not bounded by the budget")
	}
	elapsed := time.Since(start)
	if elapsed < budget {
		t.Errorf("returned after %s, before the %s budget was spent", elapsed, budget)
	}
	if remaining > 0 {
		t.Errorf("remaining = %s, want nothing left after the command used the whole budget", remaining)
	}
}

func TestRunShutdownCommand_LeavesTheUnusedBudgetForTheVMToHalt(t *testing.T) {
	comm := &fakeCommunicator{exitStatus: packersdk.CmdDisconnect}
	state := shutdownState(t, comm)

	const budget = 5 * time.Second
	remaining := runShutdownCommand(context.Background(), state, packersdk.TestUi(t), "poweroff", budget)

	if remaining <= 0 {
		t.Fatalf("remaining = %s, want the unused budget carried over to the halt wait", remaining)
	}
	if remaining > budget {
		t.Errorf("remaining = %s, want no more than the %s budget", remaining, budget)
	}
	if budget-remaining > time.Second {
		t.Errorf("a command that exited immediately consumed %s of the budget", budget-remaining)
	}
}

func TestRunInGuest_ReportsAMissingCommunicator(t *testing.T) {
	state := shutdownState(t, nil)

	err := runInGuest(context.Background(), state, packersdk.TestUi(t), "poweroff")

	if err == nil {
		t.Fatal("expected an error when no communicator has been established")
	}
	if !strings.Contains(err.Error(), "communicator") {
		t.Errorf("error %q should say that no communicator is available", err)
	}
}
