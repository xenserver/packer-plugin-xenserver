package xenserver

import (
    "fmt"
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "time"
)


type stepBootWait struct{}

func (self *stepBootWait) Run(state multistep.StateBag) multistep.StepAction {
    client := state.Get("client").(XenAPIClient)
    config := state.Get("config").(config)
    ui := state.Get("ui").(packer.Ui)

    instance, _ := client.GetVMByUuid(state.Get("instance_uuid").(string))
    ui.Say("Unpausing VM " + state.Get("instance_uuid").(string))
    instance.Unpause()

    if int64(config.BootWait) > 0 {
        ui.Say(fmt.Sprintf("Waiting %s for boot...", config.BootWait))
        time.Sleep(config.BootWait)
    }
    return multistep.ActionContinue
}

func (self *stepBootWait) Cleanup(state multistep.StateBag) {}

