package common

import (
	"crypto/tls"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type StepExport struct{}

func downloadFile(url, filename string, ui packer.Ui) (err error) {

	// Create the file
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fh.Close()

	// Define a new transport which allows self-signed certs
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create a client
	client := &http.Client{Transport: tr}

	// Create request and download file

	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var progress uint
	var total uint
	var percentage uint
	var marker_len uint

	progress = uint(0)
	total = uint(resp.ContentLength)
	percentage = uint(0)
	marker_len = uint(5)

	var buffer [4096]byte
	for {
		n, err := resp.Body.Read(buffer[:])
		if err != nil && err != io.EOF {
			return err
		}

		progress += uint(n)

		if _, write_err := fh.Write(buffer[:n]); write_err != nil {
			return write_err
		}

		if err == io.EOF {
			break
		}

		// Increment percentage in multiples of marker_len
		cur_percentage := ((progress * 100 / total) / marker_len) * marker_len
		if cur_percentage > percentage {
			percentage = cur_percentage
			ui.Message(fmt.Sprintf("Downloading... %d%%", percentage))
		}

	}

	return nil
}

func (StepExport) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)
	instance_uuid := state.Get("instance_uuid").(string)
	suffix := ".vhd"
	extrauri := "&format=vhd"

	instance, err := client.GetVMByUuid(instance_uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM with UUID '%s': %s", instance_uuid, err.Error()))
		return multistep.ActionHalt
	}

	if len(config.ExportNetworkNames) > 0 {
		vifs, err := instance.GetVIFs()
		if err != nil {
			ui.Error(fmt.Sprintf("Error occured getting VIFs: %s", err.Error()))
			return multistep.ActionHalt
		}

		for _, vif := range vifs {
			err := vif.Destroy()
			if err != nil {
				ui.Error(fmt.Sprintf("Destroy vif fail: '%s': %s", vif.Ref, err.Error()))
				return multistep.ActionHalt
			}
		}
		for i, networkNameLabel := range config.ExportNetworkNames {
			networks, err := client.GetNetworkByNameLabel(networkNameLabel)

			if err != nil {
				ui.Error(fmt.Sprintf("Error occured getting Network by name-label: %s", err.Error()))
				return multistep.ActionHalt
			}

			switch {
			case len(networks) == 0:
				ui.Error(fmt.Sprintf("Couldn't find a network with the specified name-label '%s'. Aborting.", networkNameLabel))
				return multistep.ActionHalt
			case len(networks) > 1:
				ui.Error(fmt.Sprintf("Found more than one network with the name '%s'. The name must be unique. Aborting.", networkNameLabel))
				return multistep.ActionHalt
			}

			//we need the VIF index string
			vifIndexString := fmt.Sprintf("%d", i)
			_, err = instance.ConnectNetwork(networks[0], vifIndexString)

			if err != nil {
				ui.Say(err.Error())
			}
		}
	}

	ui.Say("Step: export artifact")

	compress_option_xe := "compress=false"
	compress_option_url := ""

	switch config.Format {
	case "none":
		ui.Say("Skipping export")
		return multistep.ActionContinue

	case "xva_compressed":
		compress_option_xe = "compress=true"
		compress_option_url = "use_compression=true&"
		fallthrough
	case "xva":
		// export the VM

		export_filename := fmt.Sprintf("%s/%s.xva", config.OutputDir, config.VMName)
		ui.Say("Exporting to: " + export_filename)

		use_xe := os.Getenv("USE_XE") == "1"
		if xe, e := exec.LookPath("xe"); e == nil && use_xe {
			cmd := exec.Command(
				xe,
				"-s", client.Host,
				"-p", "443",
				"-u", client.Username,
				"-pw", client.Password,
				"vm-export",
				"vm="+instance_uuid,
				compress_option_xe,
				"filename="+export_filename,
			)

			ui.Say(fmt.Sprintf("Getting XVA %+v %+v", cmd.Path, cmd.Args))

			err = cmd.Run()
		} else {
			export_url := fmt.Sprintf("https://%s/export?%suuid=%s&session_id=%s",
				client.Host,
				compress_option_url,
				instance_uuid,
				client.Session.(string),
			)

			ui.Say("Getting XVA " + export_url)
			err = downloadFile(export_url, export_filename, ui)
		}

		if err != nil {
			ui.Error(fmt.Sprintf("Could not download XVA: %s", err.Error()))
			return multistep.ActionHalt
		}

	case "vdi_raw":
		suffix = ".raw"
		extrauri = ""
		fallthrough
	case "vdi_vhd":
		// export the disks

		disks, err := instance.GetDisks()
		if err != nil {
			ui.Error(fmt.Sprintf("Could not get VM disks: %s", err.Error()))
			return multistep.ActionHalt
		}
		for _, disk := range disks {
			disk_uuid, err := disk.GetUuid()
			if err != nil {
				ui.Error(fmt.Sprintf("Could not get disk with UUID '%s': %s", disk_uuid, err.Error()))
				return multistep.ActionHalt
			}

			// Work out XenServer version
			hosts, err := client.GetHosts()

			if err != nil {
				ui.Error(fmt.Sprintf("Could not retrieve hosts in the pool: %s", err.Error()))
				return multistep.ActionHalt
			}
			host := hosts[0]
			host_software_versions, err := host.GetSoftwareVersion()
			xs_version := host_software_versions["product_version"].(string)

			if err != nil {
				ui.Error(fmt.Sprintf("Could not get the software version: %s", err.Error()))
				return multistep.ActionHalt
			}

			var disk_export_url string

			// @todo: check for 6.5 SP1
			if xs_version <= "6.5.0" && config.Format == "vdi_vhd" {
				// Export the VHD using a Transfer VM

				disk_export_url, err = disk.Expose("vhd")

				if err != nil {
					ui.Error(fmt.Sprintf("Failed to expose disk %s: %s", disk_uuid, err.Error()))
					return multistep.ActionHalt
				}

			} else {

				// Use the preferred direct export from XAPI
				// Basic auth in URL request is required as session token is not
				// accepted for some reason.
				// @todo: raise with XAPI team.
				disk_export_url = fmt.Sprintf("https://%s:%s@%s/export_raw_vdi?vdi=%s%s",
					client.Username,
					client.Password,
					client.Host,
					disk_uuid,
					extrauri)

			}

			disk_export_filename := fmt.Sprintf("%s/%s%s", config.OutputDir, disk_uuid, suffix)
			ui.Say("Exporting to: " + disk_export_filename)

			ui.Say("Getting VDI " + disk_export_url)
			err = downloadFile(disk_export_url, disk_export_filename, ui)
			if err != nil {
				ui.Error(fmt.Sprintf("Could not download VDI: %s", err.Error()))
				return multistep.ActionHalt
			}

			// Call unexpose in case a TVM was used. The call is harmless
			// if that is not the case.
			disk.Unexpose()

		}

	default:
		panic(fmt.Sprintf("Unknown export format '%s'", config.Format))
	}

	ui.Say("Download completed: " + config.OutputDir)

	return multistep.ActionContinue
}

func (StepExport) Cleanup(state multistep.StateBag) {}
