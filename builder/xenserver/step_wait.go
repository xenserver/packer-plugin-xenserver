package xenserver

import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "time"
    "reflect"
    "github.com/nilshell/xmlrpc"
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

    //Eject ISO from drive and start VM
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

    ui.Say("Starting VM...")
    instance.Start(false, false)


    vm_ip := ""
    for i:=0; i < 10; i++ {
        ref, _ := instance.GetGuestMetricsRef()

        if ref != "OpaqueRef:NULL" {
            metrics, _ := instance.GetGuestMetrics()
            // todo: xmlrpc shouldn't be needed here
            networks := metrics["networks"].(xmlrpc.Struct)
            for k, v := range networks {
                if k == "0/ip" && v.(string) != "" {
                    vm_ip = v.(string)
                    break
                }
            }

        }

        // Check if an IP has been returned yet
        if vm_ip != "" {
            break
        }

        ui.Say("Wait for IP address...")
        time.Sleep(10*time.Second)
    }


    // Pass on the VM's IP
    state.Put("ssh_address", vm_ip)
    ui.Say("Found the VM's IP " + vm_ip)


    return multistep.ActionContinue
}

func (self *stepWait) Cleanup(state multistep.StateBag) {}

