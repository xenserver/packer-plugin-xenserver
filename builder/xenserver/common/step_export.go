// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

type StepExport struct{}

const minExportSize = 512

// A plain tar (an uncompressed XVA) is accepted too, in case a pool ignores use_compression=true.
const tarMagicOffset = 257

const previewBytes = 8

const downloadBufferBytes = 1 << 20

const progressMarkerPercent = 5

var (
	gzipMagic = []byte{0x1f, 0x8b}
	tarMagic  = []byte("ustar")
)

// exportXVAURL embeds the session reference, so the URL must be redacted before printing.
func exportXVAURL(host, compressOption, vmUUID string, sessionRef xenapi.SessionRef) string {
	return fmt.Sprintf("https://%s/export?%suuid=%s&session_id=%s", host, compressOption, vmUUID, sessionRef)
}

// exportVDIURL uses session_id: HTTP Basic userinfo is lost on XAPI's redirect to another host.
func exportVDIURL(host string, sessionRef xenapi.SessionRef, vdiUUID, extraURI string) string {
	return fmt.Sprintf("https://%s/export_raw_vdi?session_id=%s&vdi=%s%s", host, sessionRef, vdiUUID, extraURI)
}

func xeExportArgs(host, username, password, vmUUID, compressOption, filename string) []string {
	return []string{
		"-s", host,
		"-p", "443",
		"-u", username,
		"-pw", password,
		"vm-export",
		"vm=" + vmUUID,
		compressOption,
		"filename=" + filename,
	}
}

func exportLogLine(exportURL string) string {
	return "Getting XVA " + redactURL(exportURL)
}

func xeExportLogLine(path string, args []string) string {
	return fmt.Sprintf("Getting XVA %+v %+v", path, redactArgs(args))
}

// XAPI sends the 200 before the body, so a failed export still gets a 200 and xe still exits 0.
func validateExportedFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("the export wrote no file: %w", err)
	}
	if info.Size() < minExportSize {
		return fmt.Errorf("the export wrote %d bytes to %s, which is too small to be an artifact", info.Size(), path)
	}
	return nil
}

func validateExportedXVA(path string) error {
	if err := validateExportedFile(path); err != nil {
		return err
	}

	fh, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("could not open the exported XVA: %w", err)
	}
	defer fh.Close()

	header := make([]byte, tarMagicOffset+len(tarMagic))
	if _, err := io.ReadFull(fh, header); err != nil {
		return fmt.Errorf("could not read the start of %s: %w", path, err)
	}

	if bytes.HasPrefix(header, gzipMagic) {
		return nil
	}
	if bytes.Equal(header[tarMagicOffset:], tarMagic) {
		return nil
	}

	return fmt.Errorf("%s is not an XVA: it starts with %x, which is neither a gzip nor a tar header", path, header[:previewBytes])
}

// downloadFile takes the caller's client so TLS verification matches the XenAPI session's.
func downloadFile(ctx context.Context, client *http.Client, url, filename string, ui packer.Ui) (err error) {

	fh, err := os.Create(filename)
	if err != nil {
		return err
	}
	// A failed Close can lose the tail of the artifact.
	defer func() {
		if cerr := fh.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return redactURLError(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		// net/url redacts only userinfo, but the session reference is in the query string.
		return redactURLError(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("export request got non-200 status code: %s", resp.Status)
	}

	var progress, percentage uint
	total := uint(resp.ContentLength)
	showProgress := resp.ContentLength > 0

	// Exports run to tens of gigabytes; a small buffer would cost a syscall pair per few KiB.
	var buffer [downloadBufferBytes]byte
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

		if showProgress {
			cur_percentage := ((progress * 100 / total) / progressMarkerPercent) * progressMarkerPercent
			if cur_percentage > percentage {
				percentage = cur_percentage
				ui.Message(fmt.Sprintf("Downloading... %d%%", percentage))
			}
		}
	}

	return nil
}

func (StepExport) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("client").(*Connection)
	instance_uuid := state.Get("instance_uuid").(string)

	ui.Say("Step: export artifact")

	switch config.Format {
	case "none":
		ui.Say("Skipping export")
		return multistep.ActionContinue

	case "xva":
		// Always gzip (imports detect it); CommonConfig.Prepare normalizes "xva_compressed" to "xva".
		compress_option_xe := "compress=true"
		compress_option_url := "use_compression=true&"

		export_filename := fmt.Sprintf("%s/%s.xva", config.OutputDir, config.Name())

		var err error
		use_xe := os.Getenv("USE_XE") == "1"
		if xe, e := exec.LookPath("xe"); e == nil && use_xe {
			// xe takes the password in argv, readable by other local processes; the log line redacts it.
			args := xeExportArgs(c.Host, c.Username, c.Password, instance_uuid, compress_option_xe, export_filename)
			cmd := exec.CommandContext(ctx, xe, args...)

			var stderr bytes.Buffer
			cmd.Stderr = &stderr

			ui.Say(xeExportLogLine(cmd.Path, cmd.Args))

			if err = cmd.Run(); err != nil && stderr.Len() > 0 {
				err = fmt.Errorf("%w: %s", err, stderr.String())
			}
		} else {
			export_url := exportXVAURL(c.Host, compress_option_url, instance_uuid, c.GetSessionRef())

			ui.Say(exportLogLine(export_url))
			err = downloadFile(ctx, c.HTTPClient(), export_url, export_filename, ui)
		}

		if err != nil {
			ui.Error(fmt.Sprintf("Could not download XVA: %s", err.Error()))
			return multistep.ActionHalt
		}

		if err := validateExportedXVA(export_filename); err != nil {
			ui.Error(fmt.Sprintf("The XVA export did not produce a usable artifact: %s", err.Error()))
			return multistep.ActionHalt
		}

	case "vdi_raw", "vdi_vhd":
		suffix, extrauri := ".raw", ""
		if config.Format == "vdi_vhd" {
			suffix, extrauri = ".vhd", "&format=vhd"
		}

		instance, err := xenapi.VM.GetByUUID(c.session, instance_uuid)
		if err != nil {
			ui.Error(fmt.Sprintf("Could not get VM with UUID '%s': %s", instance_uuid, err.Error()))
			return multistep.ActionHalt
		}

		disks, err := GetDisks(c, instance)
		if err != nil {
			ui.Error(fmt.Sprintf("Could not get VM disks: %s", err.Error()))
			return multistep.ActionHalt
		}
		for _, disk := range disks {
			disk_uuid, err := xenapi.VDI.GetUUID(c.session, disk)
			if err != nil {
				ui.Error(fmt.Sprintf("Could not get the UUID of disk '%s': %s", disk, err.Error()))
				return multistep.ActionHalt
			}

			disk_export_url := exportVDIURL(c.Host, c.GetSessionRef(), disk_uuid, extrauri)

			disk_export_filename := fmt.Sprintf("%s/%s%s", config.OutputDir, disk_uuid, suffix)

			ui.Say(fmt.Sprintf("Getting VDI %s", disk_uuid))
			err = downloadFile(ctx, c.HTTPClient(), disk_export_url, disk_export_filename, ui)
			if err != nil {
				ui.Error(fmt.Sprintf("Could not download VDI: %s", err.Error()))
				return multistep.ActionHalt
			}

			if err := validateExportedFile(disk_export_filename); err != nil {
				ui.Error(fmt.Sprintf("The VDI export did not produce a usable artifact: %s", err.Error()))
				return multistep.ActionHalt
			}

		}

	default:
		panic(fmt.Sprintf("Unknown export format '%s'", config.Format))
	}

	ui.Say("Download completed: " + config.OutputDir)

	return multistep.ActionContinue
}

func (StepExport) Cleanup(state multistep.StateBag) {}
