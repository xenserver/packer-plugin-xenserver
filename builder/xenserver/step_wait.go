package xenserver

import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "time"
    "reflect"
)


type stepWait struct{}

func (self *stepWait) Run(state multistep.StateBag) multistep.StepAction {
    ui := state.Get("ui").(packer.Ui)
    client := state.Get("client").(XenAPIClient)

    ui.Say("Step: Wait for install to complete.")


    //Expect install to be configured to shutdown on completion

    instance_id := state.Get("instance_uuid").(string)
    instance, _ := client.GetVMByUuid(instance_id)

    for {
        time.Sleep(30 * time.Second)
        ui.Say("Waiting for VM install...")

        power_state, _ := instance.GetPowerState()
        if power_state == "Halted" {
            ui.Say("Install has completed. Moving on.")
            break
        }
    }

    // Eject ISO from drive
    vbds, _ := instance.GetVBDs()
    for _, vbd := range vbds {
        rec, _ := vbd.GetRecord()

        // Hack - should encapsulate this in the client really
        // This is needed because we can't guarentee the type
        // returned by the xmlrpc lib will be string
        switch reflect.TypeOf(rec["type"]).Kind() {
            case reflect.String:
                if rec["type"].(string) == "CD" {
                    ui.Say("Ejecting CD...")
                    vbd.Eject()
                }
            default:
                break
        }
    }

    // Destroy all connected VIFs
    vifs, _ := instance.GetVIFs()
    for _, vif := range vifs {
        ui.Message("Destroying VIF " + vif.Ref)
        vif.Destroy()
    }

    return multistep.ActionContinue
}

func (self *stepWait) Cleanup(state multistep.StateBag) {}

