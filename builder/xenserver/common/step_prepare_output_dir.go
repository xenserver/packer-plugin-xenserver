package common

// Not commonsteps.StepOutputDir from the SDK: that one discards the os.RemoveAll error below.

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepPrepareOutputDir struct {
	Force bool
	Path  string

	cleanup bool
}

func (self *StepPrepareOutputDir) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	if _, err := os.Stat(self.Path); err == nil {
		if self.Force {
			ui.Say("Deleting previous output directory...")
			if err := os.RemoveAll(self.Path); err != nil {
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		} else {
			entries, err := os.ReadDir(self.Path)
			if err != nil {
				err = fmt.Errorf("output directory is not usable: %w", err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
			// Reuse an empty directory: a `format = "none"` build leaves one behind.
			if len(entries) > 0 {
				err := fmt.Errorf("output directory is not empty: %s\n\nUse the force flag to delete it prior to building.", self.Path)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}
	}

	if err := os.MkdirAll(self.Path, 0755); err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	// Only once the directory is this build's doing is Cleanup entitled to delete it.
	self.cleanup = true

	f, err := os.Create(filepath.Join(self.Path, "_packer_perm_check"))
	if err == nil {
		// NewArtifact reports every file left here, so a surviving probe ships as an artifact.
		if err = f.Close(); err == nil {
			err = os.Remove(f.Name())
		}
	}
	if err != nil {
		err = fmt.Errorf("could not write to the output directory %s: %w", self.Path, err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (self *StepPrepareOutputDir) Cleanup(state multistep.StateBag) {
	if !self.cleanup {
		return
	}

	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)

	if cancelled || halted {
		ui := state.Get("ui").(packer.Ui)

		ui.Say("Deleting output directory...")
		var err error
		for i := 0; i < 5; i++ {
			if i > 0 {
				time.Sleep(2 * time.Second)
			}
			if err = os.RemoveAll(self.Path); err == nil {
				break
			}
			log.Printf("Error removing output dir: %s", err)
		}
		if err != nil {
			ui.Error("Could not delete output directory " + self.Path + ": " + err.Error())
		}
	}
}
