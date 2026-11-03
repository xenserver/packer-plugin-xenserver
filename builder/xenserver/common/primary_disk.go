// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"fmt"
	"strconv"

	"xenapi"
)

// DiskOutcome says what sizing the primary disk actually did.
type DiskOutcome int

const (
	// DiskUnchanged means nothing was resized; it is also returned with an error.
	DiskUnchanged DiskOutcome = iota
	// DiskGrown means the disk was resized upwards.
	DiskGrown
)

// primaryDiskAPI is an interface only so that tests can stub XenAPI directly.
type primaryDiskAPI interface {
	// ListVBDs reports the VM's block devices, disks and CDs alike.
	ListVBDs() ([]xenapi.VBDRef, error)
	// VBDRecord reads one block device.
	VBDRecord(vbd xenapi.VBDRef) (xenapi.VBDRecord, error)
	// VirtualSize reports a VDI's size in bytes.
	VirtualSize(vdi xenapi.VDIRef) (int, error)
	// Resize sets a VDI's size in bytes.
	Resize(vdi xenapi.VDIRef, bytes int) error
}

type xapiPrimaryDisk struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (d xapiPrimaryDisk) ListVBDs() ([]xenapi.VBDRef, error) {
	return xenapi.VM.GetVBDs(d.session, d.vm)
}

func (d xapiPrimaryDisk) VBDRecord(vbd xenapi.VBDRef) (xenapi.VBDRecord, error) {
	return xenapi.VBD.GetRecord(d.session, vbd)
}

func (d xapiPrimaryDisk) VirtualSize(vdi xenapi.VDIRef) (int, error) {
	return xenapi.VDI.GetVirtualSize(d.session, vdi)
}

func (d xapiPrimaryDisk) Resize(vdi xenapi.VDIRef, bytes int) error {
	return xenapi.VDI.Resize(d.session, vdi, bytes)
}

// SizePrimaryDisk grows a halted VM's primary disk to wantBytes; XAPI no longer resizes online.
func SizePrimaryDisk(c *Connection, vm xenapi.VMRef, wantBytes int) (DiskOutcome, error) {
	return sizePrimaryDisk(xapiPrimaryDisk{session: c.GetSession(), vm: vm}, wantBytes)
}

func sizePrimaryDisk(api primaryDiskAPI, wantBytes int) (DiskOutcome, error) {
	vdi, found, err := primaryDiskVDI(api)
	if err != nil {
		return DiskUnchanged, err
	}
	if !found {
		return DiskUnchanged, errors.New("disk_size_gb is set, but the source has no disk on device 0 to size")
	}

	haveBytes, err := api.VirtualSize(vdi)
	if err != nil {
		return DiskUnchanged, fmt.Errorf("could not read the primary disk's size: %w", err)
	}

	switch {
	case wantBytes < haveBytes:
		return DiskUnchanged, fmt.Errorf("disk_size_gb is %s, but the source's primary disk is already %s, "+
			"and a disk can only be grown, never shrunk. Remove disk_size_gb to keep the source's size, "+
			"or set it to %d or more.", FormatGB(wantBytes), FormatGB(haveBytes), smallestGB(haveBytes))

	case wantBytes == haveBytes:
		return DiskUnchanged, nil
	}

	if err := api.Resize(vdi, wantBytes); err != nil {
		return DiskUnchanged, fmt.Errorf("could not grow the primary disk to %s: %w", FormatGB(wantBytes), err)
	}

	return DiskGrown, nil
}

func primaryDiskVDI(api primaryDiskAPI) (xenapi.VDIRef, bool, error) {
	vbds, err := api.ListVBDs()
	if err != nil {
		return "", false, fmt.Errorf("could not list the VM's block devices: %w", err)
	}

	for _, vbd := range vbds {
		rec, err := api.VBDRecord(vbd)
		if err != nil {
			return "", false, fmt.Errorf("could not read a VBD record: %w", err)
		}

		if rec.Type == xenapi.VbdTypeDisk && rec.Userdevice == primaryDiskDevice {
			return rec.VDI, true, nil
		}
	}

	return "", false, nil
}

// FormatGB renders a byte count in GB, the unit of disk_size_gb.
func FormatGB(bytes int) string {
	if bytes%BytesPerGB == 0 {
		return strconv.Itoa(bytes/BytesPerGB) + " GB"
	}
	return strconv.FormatFloat(float64(bytes)/BytesPerGB, 'f', 2, 64) + " GB"
}

// smallestGB is the lowest whole disk_size_gb that is not a shrink.
func smallestGB(bytes int) int {
	return (bytes + BytesPerGB - 1) / BytesPerGB
}
