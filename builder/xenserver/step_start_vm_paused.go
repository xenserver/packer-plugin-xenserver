package xenserver


import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
)

type stepStartVmPaused struct {}

func (self *stepStartVmPaused) Run(state multistep.StateBag) multistep.StepAction {

    client := state.Get("client").(XenAPIClient)
    ui := state.Get("ui").(packer.Ui)

    ui.Say("Step: Start VM Paused")

    instance, _ := client.GetVMByUuid(state.Get("instance_uuid").(string))

    instance.Start(true, false) 

    domid, _ := instance.GetDomainId()
    state.Put("domid", domid)

    return multistep.ActionContinue
}


func (self *stepStartVmPaused) Cleanup(state multistep.StateBag) {
}
