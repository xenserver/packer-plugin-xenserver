package xenserver

/* Taken from https://raw.githubusercontent.com/mitchellh/packer/master/builder/qemu/step_prepare_output_dir.go */

import (
	"crypto/tls"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"io"
	"net/http"
	"os"
	"time"
)

type stepShutdownAndExport struct{}

func downloadFile(url, filename string) (err error) {

	// Create the file
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}

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
	io.Copy(fh, resp.Body)

	return nil
}

func (stepShutdownAndExport) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(XenAPIClient)
	instance_uuid := state.Get("instance_uuid").(string)

	instance, err := client.GetVMByUuid(instance_uuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Could not get VM with UUID '%s': %s", instance_uuid, err.Error()))
		return multistep.ActionHalt
	}

	ui.Say("Step: Shutdown and export VPX")

	// Shutdown the VM
	success := func() bool {
		if config.ShutdownCommand != "" {
			ui.Say("Executing shutdown command...")

			_, err := execute_ssh_cmd(config.ShutdownCommand, config.HostIp, "22", config.Username, config.Password)
			if err != nil {
				ui.Error(fmt.Sprintf("Shutdown command failed: %s", err.Error()))
				return false
			}

			ui.Say("Waiting for VM to enter Halted state...")

			err = InterruptibleWait{
				Predicate: func() (bool, error) {
					power_state, err := instance.GetPowerState()
					return power_state == "Halted", err
				},
				PredicateInterval: 5 * time.Second,
				Timeout:           300 * time.Second,
			}.Wait(state)

			if err != nil {
				ui.Error(fmt.Sprintf("Error waiting for VM to halt: %s", err.Error()))
				return false
			}

		} else {
			ui.Say("Attempting to cleanly shutdown the VM...")

			err = instance.CleanShutdown()
			if err != nil {
				ui.Error(fmt.Sprintf("Could not shut down VM: %s", err.Error()))
				return false
			}

		}
		return true
	}()

	if !success {
		ui.Say("Forcing hard shutdown of the VM...")
		err = instance.HardShutdown()
		if err != nil {
			ui.Error(fmt.Sprintf("Could not hard shut down VM -- giving up: %s", err.Error()))
			return multistep.ActionHalt
		}
	}

	ui.Say("Successfully shut down VM")

	switch config.ExportFormat {
	case "xva":
		// export the VM

		export_url := fmt.Sprintf("https://%s/export?uuid=%s&session_id=%s",
			client.Host,
			instance_uuid,
			client.Session.(string),
		)

		export_filename := fmt.Sprintf("%s/%s.xva", config.OutputDir, config.VMName)

		ui.Say("Getting XVA " + export_url)
		err = downloadFile(export_url, export_filename)
		if err != nil {
			ui.Error(fmt.Sprintf("Could not download XVA: %s", err.Error()))
			return multistep.ActionHalt
		}

	case "vdi_raw":
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

			// Basic auth in URL request is required as session token is not
			// accepted for some reason.
			// @todo: raise with XAPI team.
			disk_export_url := fmt.Sprintf("https://%s:%s@%s/export_raw_vdi?vdi=%s",
				client.Username,
				client.Password,
				client.Host,
				disk_uuid,
			)

			disk_export_filename := fmt.Sprintf("%s/%s.raw", config.OutputDir, disk_uuid)

			ui.Say("Getting VDI " + disk_export_url)
			err = downloadFile(disk_export_url, disk_export_filename)
			if err != nil {
				ui.Error(fmt.Sprintf("Could not download VDI: %s", err.Error()))
				return multistep.ActionHalt
			}
		}

	default:
		panic(fmt.Sprintf("Unknown export_format '%s'", config.ExportFormat))
	}

	ui.Say("Download completed: " + config.OutputDir)

	return multistep.ActionContinue
}

func (stepShutdownAndExport) Cleanup(state multistep.StateBag) {
}
