package common

/* Taken from https://raw.githubusercontent.com/mitchellh/packer/master/builder/qemu/step_prepare_output_dir.go */

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
	"os"
	"time"
)

type StepPrepareOutputDir struct {
	Force bool
	Path  string
}

func (self *StepPrepareOutputDir) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	if _, err := os.Stat(self.Path); err == nil && self.Force {
		ui.Say("Deleting previous output directory...")
		os.RemoveAll(self.Path)
	}

	if err := os.MkdirAll(self.Path, 0755); err != nil {
		state.Put("error", err)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (self *StepPrepareOutputDir) Cleanup(state multistep.StateBag) {
	_, cancelled := state.GetOk(multistep.StateCancelled)
	_, halted := state.GetOk(multistep.StateHalted)

	if cancelled || halted {
		ui := state.Get("ui").(packer.Ui)

		ui.Say("Deleting output directory...")
		for i := 0; i < 5; i++ {
			err := os.RemoveAll(self.Path)
			if err == nil {
				break
			}

			log.Printf("Error removing output dir: %s", err)
			time.Sleep(2 * time.Second)
		}
	}
}
