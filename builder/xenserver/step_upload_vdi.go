package xenserver

import (
	"crypto/tls"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
	"net/http"
	"os"
	"time"
)

type stepUploadVdi struct {
	VdiName       string
	ImagePathFunc func() string
	VdiUuidKey    string
}

func (self *stepUploadVdi) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(XenAPIClient)

	imagePath := self.ImagePathFunc()
	if imagePath == "" {
		// skip if no disk image to attach
		return multistep.ActionContinue
	}

	ui.Say(fmt.Sprintf("Step: Upload VDI '%s'", self.VdiName))

	// Create VDI for the image
	sr, err := config.GetSR(client)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get SR: %s", err.Error()))
		return multistep.ActionHalt
	}

	// Open the file for reading (NB: putFile closes the file for us)
	fh, err := os.Open(imagePath)
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to open disk image '%s': %s", imagePath, err.Error()))
		return multistep.ActionHalt
	}

	// Get file length
	fstat, err := fh.Stat()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to stat disk image '%s': %s", imagePath, err.Error()))
		return multistep.ActionHalt
	}
	fileLength := fstat.Size()

	// Create the VDI
	vdi, err := sr.CreateVdi(self.VdiName, fmt.Sprintf("%d", fileLength))
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create VDI '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}

	vdiUuid, err := vdi.GetUuid()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to get UUID of VDI '%s': %s", self.VdiName, err.Error()))
		return multistep.ActionHalt
	}
	state.Put(self.VdiUuidKey, vdiUuid)

	task, err := client.CreateTask()
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to create task: %s", err.Error()))
		return multistep.ActionHalt
	}
	defer task.Destroy()

	import_url := fmt.Sprintf("https://%s/import_raw_vdi?vdi=%s&session_id=%s&task_id=%s",
		client.Host,
		vdi.Ref,
		client.Session.(string),
		task.Ref,
	)

	// Define a new transport which allows self-signed certs
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create a client
	httpClient := &http.Client{Transport: tr}

	// Create request and download file
	request, err := http.NewRequest("PUT", import_url, fh)
	request.ContentLength = fileLength

	ui.Say(fmt.Sprintf("PUT disk image '%s'", import_url))

	resp, err := httpClient.Do(request) // Do closes fh for us, according to docs
	if err != nil {
		ui.Error(fmt.Sprintf("Unable to upload disk image: %s", err.Error()))
		return multistep.ActionHalt
	}

	if resp.StatusCode != 200 {
		ui.Error(fmt.Sprintf("Unable to upload disk image: PUT request got non-200 status code: %s", resp.Status))
		return multistep.ActionHalt
	}

	logInterval := 0
	err = InterruptibleWait{
		Predicate: func() (bool, error) {
			status, err := task.GetStatus()
			if err != nil {
				return false, fmt.Errorf("Failed to get task status: %s", err.Error())
			}
			switch status {
			case Pending:
				progress, err := task.GetProgress()
				if err != nil {
					return false, fmt.Errorf("Failed to get progress: %s", err.Error())
				}
				logInterval = logInterval + 1
				if logInterval%5 == 0 {
					log.Printf("Upload %.0f%% complete", progress*100)
				}
				return false, nil
			case Success:
				return true, nil
			case Failure:
				errorInfo, err := task.GetErrorInfo()
				if err != nil {
					errorInfo = []string{fmt.Sprintf("furthermore, failed to get error info: %s", err.Error())}
				}
				return false, fmt.Errorf("Task failed: %s", errorInfo)
			case Cancelling, Cancelled:
				return false, fmt.Errorf("Task cancelled")
			default:
				return false, fmt.Errorf("Unknown task status %v", status)
			}
		},
		PredicateInterval: 1 * time.Second,
		Timeout:           24 * time.Hour,
	}.Wait(state)

	resp.Body.Close()

	if err != nil {
		ui.Error(fmt.Sprintf("Error uploading: %s", err.Error()))

		return multistep.ActionHalt
	}

	log.Printf("Upload complete")

	return multistep.ActionContinue
}

func (self *stepUploadVdi) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(XenAPIClient)

	if config.ShouldKeepInstance(state) {
		return
	}

	vdiUuidRaw, ok := state.GetOk(self.VdiUuidKey)
	if !ok {
		// VDI doesn't exist
		return
	}

	vdiUuid := vdiUuidRaw.(string)
	if vdiUuid == "" {
		// VDI already cleaned up
		return
	}

	vdi, err := client.GetVdiByUuid(vdiUuid)
	if err != nil {
		ui.Error(fmt.Sprintf("Can't get VDI '%s': %s", vdiUuid, err.Error()))
		return
	}

	// an interrupted import_raw_vdi takes a while to release the VDI
	// so try several times
	for i := 0; i < 3; i++ {
		log.Printf("Trying to destroy VDI...")
		err = vdi.Destroy()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		ui.Error(fmt.Sprintf("Can't destroy VDI '%s': %s", vdiUuid, err.Error()))
		return
	}
	ui.Say(fmt.Sprintf("Destroyed VDI '%s'", self.VdiName))

	state.Put(self.VdiUuidKey, "")
}
