// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"xenapi"
)

// other_config keys that XenServer provisions a diskless template's disks from, then deletes.
const (
	provisionDisksKey  = "disks"
	provisionScriptKey = "postinstall"

	// primaryDiskDevice is the system disk's userdevice, both on a VBD and in the provisioning XML.
	primaryDiskDevice = "0"
)

// provisionXML is the <provision> document in other_config:disks.
type provisionXML struct {
	XMLName xml.Name `xml:"provision"`
	Disks   []struct {
		Device string `xml:"device,attr"`
		Size   string `xml:"size,attr"`
	} `xml:"disk"`
}

// TemplateDiskSize reports the primary disk size, in bytes, that clone_template specifies.
func TemplateDiskSize(c *Connection, src CloneSource, vm xenapi.VMRef, record xenapi.VMRecord) (int, error) {
	return templateDiskSize(xapiPrimaryDisk{session: c.GetSession(), vm: vm}, src, record.OtherConfig)
}

func templateDiskSize(api primaryDiskAPI, src CloneSource, otherConfig map[string]string) (int, error) {
	size, err := provisionedDiskSize(otherConfig[provisionDisksKey])
	if err != nil {
		return 0, fmt.Errorf("could not read the provisioning disks of %s: %w", src, err)
	}
	if size > 0 {
		return size, nil
	}

	// Unlike XenCenter, fall back to the template's own disk when the XML lists no device 0.
	vdi, found, err := primaryDiskVDI(api)
	if err != nil {
		return 0, fmt.Errorf("could not read the disks of %s: %w", src, err)
	}
	if found {
		if size, err = api.VirtualSize(vdi); err != nil {
			return 0, fmt.Errorf("could not read the primary disk's size: %w", err)
		}
		if size > 0 {
			return size, nil
		}
	}

	return 0, fmt.Errorf("%s supplies no primary disk size: it describes no disk to provision "+
		"and carries none of its own. Set disk_size_gb, which is then the size of the disk this "+
		"builder installs onto.", src)
}

func provisionedDiskSize(doc string) (int, error) {
	if strings.TrimSpace(doc) == "" {
		return 0, nil
	}

	var provision provisionXML
	if err := xml.Unmarshal([]byte(doc), &provision); err != nil {
		return 0, err
	}

	for _, disk := range provision.Disks {
		if disk.Device != primaryDiskDevice {
			continue
		}

		// The size is in bytes: xapi passes it straight to VDI.create.
		size, err := strconv.ParseInt(disk.Size, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("the disk on device %s has an unreadable size %q",
				primaryDiskDevice, disk.Size)
		}
		if size <= 0 {
			return 0, fmt.Errorf("the disk on device %s has a size of %d bytes",
				primaryDiskDevice, size)
		}

		return int(size), nil
	}

	return 0, nil
}

// EnsureCloneSourceHasPrimaryDisk rejects disk_size_gb for a source with no disk to grow.
func EnsureCloneSourceHasPrimaryDisk(c *Connection, src CloneSource, vm xenapi.VMRef) error {
	return ensureCloneSourceHasPrimaryDisk(xapiPrimaryDisk{session: c.GetSession(), vm: vm}, src)
}

func ensureCloneSourceHasPrimaryDisk(api primaryDiskAPI, src CloneSource) error {
	if _, found, err := primaryDiskVDI(api); err != nil || found {
		return err
	}

	return fmt.Errorf("disk_size_gb is set, but %s has no disk on device %s: this builder grows the "+
		"disk its source already carries and never creates one. Remove disk_size_gb, or name a source "+
		"that has an installed disk.", src, primaryDiskDevice)
}

// StripProvisioningKeys removes the provisioning keys, which would break `xe vm-install` from the template.
func StripProvisioningKeys(c *Connection, vm xenapi.VMRef) ([]string, error) {
	otherConfig, err := xenapi.VM.GetOtherConfig(c.GetSession(), vm)
	if err != nil {
		return nil, fmt.Errorf("could not read the VM's other_config: %w", err)
	}

	var stripped []string
	for _, key := range []string{provisionDisksKey, provisionScriptKey} {
		if _, ok := otherConfig[key]; !ok {
			continue
		}
		if err := xenapi.VM.RemoveFromOtherConfig(c.GetSession(), vm, key); err != nil {
			return stripped, fmt.Errorf("could not remove other_config:%s from the VM: %w", key, err)
		}
		stripped = append(stripped, key)
	}

	return stripped, nil
}
