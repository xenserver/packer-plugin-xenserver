// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

func appendQuery(urlstring, k, v string) (string, error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		// url.Error's own text quotes the raw URL, which holds the session reference.
		return "", fmt.Errorf("unable to parse URL: %s", redactURLError(err))
	}
	m := u.Query()
	m.Add(k, v)
	u.RawQuery = m.Encode()
	return u.String(), nil
}

// uploadLogLine is separate so a test can check the session reference stays out of the log.
func uploadLogLine(importURL string) string {
	return fmt.Sprintf("PUT '%s'", redactURL(importURL))
}

// XenCenter's MAX_REDIRECTS (XenModel/XenAPI/HTTP.cs).
const maxUploadRedirects = 10

// redirectTarget takes only the authority from location so the VDI written cannot change.
func redirectTarget(location, originalURL string) (string, error) {
	loc, err := url.Parse(location)
	if err != nil {
		return "", fmt.Errorf("could not parse the Location XAPI redirected to: %s", redactURLError(err))
	}
	if loc.Host == "" {
		return "", fmt.Errorf("XAPI redirected the upload to '%s', which names no host", redactURL(location))
	}

	u, err := url.Parse(originalURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse the upload URL: %s", redactURLError(err))
	}
	if loc.Scheme != "" {
		u.Scheme = loc.Scheme
	}
	u.Host = loc.Host

	return u.String(), nil
}

// uploadAnswerTimeout covers only XAPI's answer; a slow multi-GB transfer is not a stall.
const uploadAnswerTimeout = 30 * time.Second

type dialPoolHost func(host string) (net.Conn, error)

// bodyOwningConn closes the socket before the body, whose Close drains it with no deadline.
type bodyOwningConn struct {
	io.ReadCloser
	conn net.Conn
	once sync.Once
	err  error
}

func (b *bodyOwningConn) Close() error {
	b.once.Do(func() {
		b.err = b.conn.Close()
		_ = b.ReadCloser.Close()
	})
	return b.err
}

// putFollowingRedirects reads XAPI's answer before sending the file, because XAPI replies
// and closes before reading the body. Only 302 is followed: 301 and 303 allow a method change.
func putFollowingRedirects(dial dialPoolHost, putURL, path string, contentLength int64, ui packer.Ui) (*http.Response, error) {
	for redirects := 0; ; redirects++ {
		resp, err := putOnce(dial, putURL, path, contentLength, ui)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusFound {
			return resp, nil
		}

		location := resp.Header.Get("Location")
		contacted := resp.Request.URL.Host
		resp.Body.Close()

		if redirects == maxUploadRedirects {
			return nil, fmt.Errorf("gave up after %d redirects; the last named '%s'",
				maxUploadRedirects, redactURL(location))
		}

		ui.Say(fmt.Sprintf("%s does not have the SR attached; following XAPI's redirect", contacted))

		if putURL, err = redirectTarget(location, putURL); err != nil {
			return nil, err
		}
	}
}

// putOnce must keep putURL out of its errors: the query holds the session reference.
func putOnce(dial dialPoolHost, putURL, path string, contentLength int64, ui packer.Ui) (*http.Response, error) {
	u, err := url.Parse(putURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the upload URL: %s", redactURLError(err))
	}

	ui.Say(uploadLogLine(putURL))

	conn, err := dial(u.Host)
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(uploadAnswerTimeout))

	head := fmt.Sprintf("PUT %s HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\n"+
		"Content-Length: %d\r\nConnection: close\r\n\r\n",
		u.RequestURI(), u.Host, pluginUserAgent(), contentLength)
	if _, err := io.WriteString(conn, head); err != nil {
		conn.Close()
		return nil, fmt.Errorf("unable to send the upload request to %s: %s", u.Host, err.Error())
	}

	resp, err := http.ReadResponse(bufio.NewReader(conn),
		&http.Request{Method: http.MethodPut, URL: u, Host: u.Host})
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("no usable reply from %s to the upload: %s", u.Host, err.Error())
	}
	resp.Body = &bodyOwningConn{ReadCloser: resp.Body, conn: conn}

	// A redirect, or a refusal for the caller to report: either way, do not send the file.
	if resp.StatusCode != http.StatusOK {
		return resp, nil
	}

	fh, err := os.Open(path)
	if err != nil {
		resp.Body.Close()
		return nil, fmt.Errorf("Unable to open '%s': %s", path, err.Error())
	}
	defer fh.Close()

	_ = conn.SetDeadline(time.Time{})

	// CopyN, not Copy: XAPI reads exactly Content-Length bytes, and the file may have grown.
	sent, err := io.CopyN(conn, fh, contentLength)
	if err != nil {
		resp.Body.Close()
		if errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("'%s' is %d bytes but the upload announced %d: "+
				"it changed size while the build was running", path, sent, contentLength)
		}
		return nil, fmt.Errorf("unable to send '%s' to %s: %s", path, u.Host, err.Error())
	}

	return resp, nil
}

// whyTheUploadFailed prefers the task's error: a failed import only closes the connection.
func whyTheUploadFailed(c *Connection, task xenapi.TaskRef, err error) error {
	status, statusErr := xenapi.Task.GetStatus(c.session, task)
	if statusErr != nil || status != xenapi.TaskStatusTypeFailure {
		return err
	}

	info, infoErr := xenapi.Task.GetErrorInfo(c.session, task)
	if infoErr != nil || len(info) == 0 {
		return err
	}

	return fmt.Errorf("XAPI failed the import: %s (the connection then failed with: %s)", info, err)
}

func HTTPUpload(import_url string, fh *os.File, state multistep.StateBag) (result string, err error) {
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)

	task, err := xenapi.Task.Create(c.session, "packer-task", "Packer task")
	if err != nil {
		err = fmt.Errorf("Unable to create task: %s", err.Error())
		return
	}
	defer xenapi.Task.Destroy(c.session, task)

	import_task_url, err := appendQuery(import_url, "task_id", string(task))
	if err != nil {
		return
	}

	fstat, err := fh.Stat()
	if err != nil {
		err = fmt.Errorf("Unable to stat '%s': %s", fh.Name(), err.Error())
		return
	}
	fileLength := fstat.Size()

	resp, err := putFollowingRedirects(c.DialTLS, import_task_url, fh.Name(), fileLength, ui)
	if err != nil {
		err = whyTheUploadFailed(c, task, err)
		return
	}

	// Close now rather than defer: the task poll below lasts as long as the import.
	resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("PUT to %s got non-200 status code: %s",
			redactURL(resp.Request.URL.String()), resp.Status)
		return
	}

	logIteration := 0
	err = InterruptibleWait{
		Predicate: func() (bool, error) {
			status, err := xenapi.Task.GetStatus(c.session, task)
			if err != nil {
				return false, fmt.Errorf("Failed to get task status: %s", err.Error())
			}
			switch status {
			case xenapi.TaskStatusTypePending:
				logIteration = logIteration + 1
				if logIteration%5 == 0 {
					progress, err := xenapi.Task.GetProgress(c.session, task)
					if err != nil {
						return false, fmt.Errorf("Failed to get progress: %s", err.Error())
					}
					log.Printf("Upload %.0f%% complete", progress*100)
				}
				return false, nil
			case xenapi.TaskStatusTypeSuccess:
				return true, nil
			case xenapi.TaskStatusTypeFailure:
				errorInfo, err := xenapi.Task.GetErrorInfo(c.session, task)
				if err != nil {
					errorInfo = []string{fmt.Sprintf("furthermore, failed to get error info: %s", err.Error())}
				}
				return false, fmt.Errorf("Task failed: %s", errorInfo)
			case xenapi.TaskStatusTypeCancelling, xenapi.TaskStatusTypeCancelled:
				return false, fmt.Errorf("Task cancelled")
			default:
				return false, fmt.Errorf("Unknown task status %v", status)
			}
		},
		PredicateInterval: 1 * time.Second,
		Timeout:           24 * time.Hour,
	}.Wait(state)

	if err != nil {
		err = fmt.Errorf("Error uploading: %s", err.Error())
		return
	}

	result, err = xenapi.Task.GetResult(c.session, task)
	if err != nil {
		err = fmt.Errorf("Error getting result: %s", err.Error())
		return
	}

	log.Printf("Upload complete")
	return
}
