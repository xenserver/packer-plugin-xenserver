package xenserver


import (
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
    "os"
    "strconv"
    "os/exec"
    "log"
)

type stepUploadIso struct {}

func (self *stepUploadIso) Run(state multistep.StateBag) multistep.StepAction {

    client := state.Get("client").(XenAPIClient)
    config := state.Get("config").(config)
    ui := state.Get("ui").(packer.Ui)

    ui.Say("Step: Upload ISO to server")
    iso_path := state.Get("iso_path").(string)

    // Determine the ISO's filesize
    file, err := os.Open(iso_path)
    if err != nil {
        ui.Error(err.Error())
        return multistep.ActionHalt
    }
    stat, err := file.Stat()

    if err != nil {
        ui.Error(err.Error())
        return multistep.ActionHalt
    }

    iso_filesize := stat.Size()

    // Create a VDI with the write size
    srs, err := client.GetSRByNameLabel(config.SrName)

    sr := srs[0]

    if err != nil {
        ui.Error(err.Error())
        return multistep.ActionHalt
    }

    filesize_str := strconv.FormatInt(iso_filesize, 10)
    log.Printf("Filesize of the ISO is %d", filesize_str)
    vdi, err := sr.CreateVdi("Packer Gen " + stat.Name(), filesize_str)

    if err != nil {
        ui.Error(err.Error())
        return multistep.ActionHalt
    }

    // Upload the ISO to this VDI
    vdi_uuid, _ := vdi.GetUuid()
    host_url := "https://" +  client.Host
    log.Printf("Host URL: %s", host_url)
    ui.Say("Uploading ISO to " + vdi_uuid)
    out, err := exec.Command("/usr/bin/importdisk.py", host_url, client.Username, client.Password, vdi_uuid, iso_path).CombinedOutput()
    log.Printf("Output: %s", out)

    if err != nil {
        log.Printf("%s", err)
        ui.Error(err.Error())
        return multistep.ActionHalt
    }

    // Stash the vdi uuid to be used in preference
    state.Put("iso_vdi_uuid", vdi_uuid)


    return multistep.ActionContinue
}


func (self *stepUploadIso) Cleanup(state multistep.StateBag) {
}
