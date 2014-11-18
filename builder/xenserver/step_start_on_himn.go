package xenserver

import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "time"
    "log"
)

type stepStartOnHIMN struct{}


/*
 * This step starts the installed guest on the Host Internal Management Network
 * as there exists an API to obtain the IP allocated to the VM by XAPI.
 * This in turn will allow Packer to SSH into the VM, provided NATing has been
 * enabled on the host.
 *
 */

func (self *stepStartOnHIMN) Run(state multistep.StateBag) multistep.StepAction {

    ui := state.Get("ui").(packer.Ui)
    client := state.Get("client").(XenAPIClient)

    ui.Say("Step: Start VM on the Host Internal Mangement Network")

    instance := state.Get("instance").(*VM)

    // Find the HIMN Ref
    networks, err := client.GetNetworkByNameLabel("Host internal management network")
    if err != nil || len(networks) == 0 {
        log.Fatal("Unable to find a host internal management network")
        log.Fatal(err.Error())
        return multistep.ActionHalt
    }


    himn := networks[0]

    // Create a VIF for the HIMN
    himn_vif, err := instance.ConnectNetwork(himn, "0")
    if err != nil {
        log.Fatal("Error creating VIF")
        log.Fatal(err.Error())
        return multistep.ActionHalt
    }

    // Start the VM
    instance.Start(false, false)


    var himn_iface_ip string = ""

    // Obtain the allocated IP
    for i:=0; i < 10; i++ {
        ips, _ := himn.GetAssignedIPs()
        log.Printf("IPs: %s", ips)
        log.Printf("Ref: %s", instance.Ref)

        //Check for instance.Ref in map
        if vm_ip, ok := ips[himn_vif.Ref]; ok {
            ui.Say("Found the VM's IP " + vm_ip)
            himn_iface_ip = vm_ip
            break
        }

        ui.Say("Wait for IP address...")
        time.Sleep(10*time.Second)

    }

    if himn_iface_ip != "" {
        state.Put("himn_ssh_address", himn_iface_ip)
        ui.Say("Stored VM's IP " + himn_iface_ip)
    } else {
        log.Fatal("Unable to find an IP on the Host-internal management interface")
        return multistep.ActionHalt
    }

    return multistep.ActionContinue

}

func (self *stepStartOnHIMN) Cleanup(state multistep.StateBag) {}


func himnSSHIP (state multistep.StateBag) (string, error) {
    ip := state.Get("himn_ssh_address").(string)
    return ip, nil
}


func himnSSHPort (state multistep.StateBag) (uint, error) {
    return 22, nil
}
