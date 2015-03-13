package common

import (
	"crypto/tls"
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	xsclient "github.com/xenserver/go-xenserver-client"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func appendQuery(urlstring, k, v string) (string, error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		return "", fmt.Errorf("Unable to parse URL '%s': %s", urlstring, err.Error())
	}
	m := u.Query()
	m.Add(k, v)
	u.RawQuery = m.Encode()
	return u.String(), err
}

func HTTPUpload(import_url string, fh *os.File, state multistep.StateBag) (result *xsclient.XenAPIObject, err error) {
	ui := state.Get("ui").(packer.Ui)
	client := state.Get("client").(xsclient.XenAPIClient)

	task, err := client.CreateTask()
	if err != nil {
		err = fmt.Errorf("Unable to create task: %s", err.Error())
		return
	}
	defer task.Destroy()

	import_task_url, err := appendQuery(import_url, "task_id", task.Ref)
	if err != nil {
		return
	}

	// Get file length
	fstat, err := fh.Stat()
	if err != nil {
		err = fmt.Errorf("Unable to stat '%s': %s", fh.Name(), err.Error())
		return
	}
	fileLength := fstat.Size()

	// Define a new transport which allows self-signed certs
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Create a client
	httpClient := &http.Client{Transport: tr}

	// Create request and download file
	request, err := http.NewRequest("PUT", import_task_url, fh)
	request.ContentLength = fileLength

	ui.Say(fmt.Sprintf("PUT '%s'", import_task_url))

	resp, err := httpClient.Do(request) // Do closes fh for us, according to docs
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("PUT request got non-200 status code: %s", resp.Status)
		return
	}

	logIteration := 0
	err = InterruptibleWait{
		Predicate: func() (bool, error) {
			status, err := task.GetStatus()
			if err != nil {
				return false, fmt.Errorf("Failed to get task status: %s", err.Error())
			}
			switch status {
			case xsclient.Pending:
				progress, err := task.GetProgress()
				if err != nil {
					return false, fmt.Errorf("Failed to get progress: %s", err.Error())
				}
				logIteration = logIteration + 1
				if logIteration%5 == 0 {
					log.Printf("Upload %.0f%% complete", progress*100)
				}
				return false, nil
			case xsclient.Success:
				return true, nil
			case xsclient.Failure:
				errorInfo, err := task.GetErrorInfo()
				if err != nil {
					errorInfo = []string{fmt.Sprintf("furthermore, failed to get error info: %s", err.Error())}
				}
				return false, fmt.Errorf("Task failed: %s", errorInfo)
			case xsclient.Cancelling, xsclient.Cancelled:
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
		err = fmt.Errorf("Error uploading: %s", err.Error())
		return
	}

	result, err = task.GetResult()
	if err != nil {
		err = fmt.Errorf("Error getting result: %s", err.Error())
		return
	}

	log.Printf("Upload complete")
	return
}
