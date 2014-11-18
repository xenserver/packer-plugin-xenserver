package xenserver


import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "fmt"
)

type stepCreateInstance struct {
    InstanceId string
}

func (self *stepCreateInstance) Run(state multistep.StateBag) multistep.StepAction {

    client := state.Get("client").(XenAPIClient)
    config := state.Get("config").(config)
    ui := state.Get("ui").(packer.Ui)

    ui.Say("Step: Create Instance")


    // Get the template to clone from
    template, _ := client.GetVMByUuid(config.CloneTemplate)

    // Clone that VM template
    instance, _ := template.Clone(config.InstanceName)
    instance.SetIsATemplate(false)
    instance.SetStaticMemoryRange("1024000000", "1024000000")
    instance.SetPlatform(config.PlatformArgs)

    // Create VDI for the instance
    sr, _ := client.GetSRByUuid(config.SrUuid)
    vdi, _ := sr.CreateVdi("Packer-disk", config.RootDiskSize)

    instance.ConnectVdi(vdi, false)

    // Connect Network
    network, err := client.GetNetworkByUuid(config.NetworkUuid)
    if err != nil {
        ui.Say(err.Error())
    }
    _, err = instance.ConnectNetwork(network, "0")

    if err != nil {
        ui.Say(err.Error())
    }

    // Connect the ISO
    //iso_vdi_uuid := state.Get("iso_vdi_uuid").(string)

    iso, _ := client.GetVdiByUuid(config.IsoUuid)
    //ui.Say("Using VDI: " + iso_vdi_uuid)
    //iso, _ := client.GetVdiByUuid(iso_vdi_uuid)
    instance.ConnectVdi(iso, true)

    // Stash the VM reference
    self.InstanceId, _ = instance.GetUuid()
    state.Put("instance_uuid", self.InstanceId)
    state.Put("instance", instance)
    ui.Say(fmt.Sprintf("Created instance '%s'", self.InstanceId))

    return multistep.ActionContinue
}


func (self *stepCreateInstance) Cleanup(state multistep.StateBag) {

//    client := state.Get("client").(*XenAPIClient)
//    config := state.Get("config").(config)
//    ui := state.Get("ui").(packer.Ui)

    // If instance hasn't been created, we have nothing to do.
    if self.InstanceId == "" {
        return
    }

    // @todo: destroy the created instance.

    return
}
