// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

func outputDirState() (*multistep.BasicStateBag, *packersdk.MockUi) {
	ui := &packersdk.MockUi{}
	state := new(multistep.BasicStateBag)
	state.Put("ui", ui)

	return state, ui
}

func writeFile(t *testing.T, path string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("previous build"), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestPrepareOutputDirHaltsOnANonEmptyDirectoryWithoutForce(t *testing.T) {
	path := t.TempDir()
	foreign := filepath.Join(path, "someone-elses.xva")
	writeFile(t, foreign)

	state, ui := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}

	err, ok := state.GetOk("error")
	if !ok {
		t.Fatal("no error was put in the state bag")
	}
	if msg := err.(error).Error(); !strings.Contains(msg, "force flag") {
		t.Errorf("error %q does not point at the force flag", msg)
	}
	if !strings.Contains(ui.ErrorMessage, "force flag") {
		t.Errorf("ui error %q does not point at the force flag", ui.ErrorMessage)
	}

	state.Put(multistep.StateHalted, true)
	step.Cleanup(state)

	if _, err := os.Stat(foreign); err != nil {
		t.Errorf("Cleanup deleted a directory the build never took over: %s", err)
	}
}

func TestPrepareOutputDirReusesAnExistingEmptyDirectory(t *testing.T) {
	path := t.TempDir()

	state, _ := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run = %v, want ActionContinue", action)
	}
	if _, ok := state.GetOk("error"); ok {
		t.Error("an empty output directory was refused")
	}
}

func TestPrepareOutputDirHaltsWhenThePathIsNotADirectory(t *testing.T) {
	path := filepath.Join(t.TempDir(), "in-the-way")
	writeFile(t, path)

	state, ui := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}
	if !ui.ErrorCalled {
		t.Error("the failure was never reported on the ui")
	}

	state.Put(multistep.StateHalted, true)
	step.Cleanup(state)

	if _, err := os.Stat(path); err != nil {
		t.Errorf("Cleanup deleted a file the build never created: %s", err)
	}
}

func TestPrepareOutputDirHaltsWhenTheDirectoryCannotBeCreated(t *testing.T) {
	blocker := filepath.Join(t.TempDir(), "not-a-directory")
	writeFile(t, blocker)
	path := filepath.Join(blocker, "output-test")

	state, ui := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Error("no error was put in the state bag")
	}
	if !ui.ErrorCalled {
		t.Error("the failure was never reported on the ui")
	}

	state.Put(multistep.StateHalted, true)
	step.Cleanup(state)

	for _, said := range ui.SayMessages {
		if strings.Contains(said.Message, "Deleting output directory") {
			t.Error("Cleanup took over a directory Run never created")
		}
	}
}

func TestPrepareOutputDirCleanupLeavesADanglingSymlinkAlone(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "output-test")
	if err := os.Symlink(filepath.Join(dir, "nowhere"), path); err != nil {
		t.Skipf("symlinks are unavailable here: %s", err)
	}

	state, _ := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	// os.Stat follows the dangling link and reports ENOENT, so only MkdirAll fails.
	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}

	state.Put(multistep.StateHalted, true)
	step.Cleanup(state)

	if _, err := os.Lstat(path); err != nil {
		t.Errorf("Cleanup deleted a path the build never created: %s", err)
	}
}

func TestPrepareOutputDirHaltsWhenForceCannotEmptyTheDirectory(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("os.Chmod does not deny directory writes on Windows")
	}
	if os.Geteuid() == 0 {
		t.Skip("root unlinks from a read-only directory regardless")
	}

	path := filepath.Join(t.TempDir(), "output-test")
	previous := filepath.Join(path, "previous.xva")
	writeFile(t, previous)
	if err := os.Chmod(path, 0o555); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(path, 0o755) })

	state, ui := outputDirState()
	step := &StepPrepareOutputDir{Path: path, Force: true}

	if action := step.Run(context.Background(), state); action != multistep.ActionHalt {
		t.Fatalf("Run = %v, want ActionHalt", action)
	}
	if _, ok := state.GetOk("error"); !ok {
		t.Error("no error was put in the state bag")
	}
	if !ui.ErrorCalled {
		t.Error("the failure was never reported on the ui")
	}

	if _, err := os.Stat(previous); err != nil {
		t.Errorf("the build carried on over the previous artifact: %s", err)
	}
}

func TestPrepareOutputDirForceEmptiesAnExistingDirectory(t *testing.T) {
	path := t.TempDir()
	previous := filepath.Join(path, "previous.xva")
	writeFile(t, previous)

	state, _ := outputDirState()
	step := &StepPrepareOutputDir{Path: path, Force: true}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run = %v, want ActionContinue", action)
	}

	if _, err := os.Stat(previous); !os.IsNotExist(err) {
		t.Errorf("the previous artifact survived -force: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Errorf("output directory was not recreated: %s", err)
	}
}

func TestPrepareOutputDirLeavesNoPermissionProbeBehind(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output-test")
	state, _ := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run = %v, want ActionContinue", action)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		t.Fatalf("output directory was not created: %s", err)
	}
	if len(entries) != 0 {
		t.Errorf("output directory is not empty after Run: %v", entries)
	}
}

func TestPrepareOutputDirCleanupRemovesTheDirectoryOnHalt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output-test")
	state, _ := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run = %v, want ActionContinue", action)
	}
	writeFile(t, filepath.Join(path, "half-written.xva"))

	state.Put(multistep.StateHalted, true)
	step.Cleanup(state)

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("output directory survived a halted build: %v", err)
	}
}

func TestPrepareOutputDirCleanupKeepsTheDirectoryOnSuccess(t *testing.T) {
	path := filepath.Join(t.TempDir(), "output-test")
	state, _ := outputDirState()
	step := &StepPrepareOutputDir{Path: path}

	if action := step.Run(context.Background(), state); action != multistep.ActionContinue {
		t.Fatalf("Run = %v, want ActionContinue", action)
	}
	exported := filepath.Join(path, "built.xva")
	writeFile(t, exported)

	step.Cleanup(state)

	if _, err := os.Stat(exported); err != nil {
		t.Errorf("Cleanup deleted the artifact of a successful build: %s", err)
	}
}
