package common

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	
	"xenapi"
	version "github.com/xenserver/packer-plugin-xenserver/version"
)

type VDIType int

const (
	_ VDIType = iota
	Disk
	CD
	Floppy
)

type TaskStatusType int

const (
	_ TaskStatusType = iota
	Pending
	Success
	Failure
	Cancelling
	Cancelled
)

func Unpause(c *Connection, vmRef xenapi.VMRef) (err error) {
	err = xenapi.VM.Unpause(c.session, vmRef)
	if err != nil {
		return err
	}
	return
}

func GetDisks(c *Connection, vmRef xenapi.VMRef) (vdis []xenapi.VDIRef, err error) {
	// Return just data disks (non-isos)
	vdis = make([]xenapi.VDIRef, 0)
	vbds, err := xenapi.VM.GetVBDs(c.session, vmRef)
	if err != nil {
		return nil, err
	}

	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(c.session, vbd)
		if err != nil {
			return nil, err
		}
		if rec.Type == "Disk" {

			vdi, err := xenapi.VBD.GetVDI(c.session, vbd)
			if err != nil {
				return nil, err
			}
			vdis = append(vdis, vdi)

		}
	}
	return vdis, nil
}



func ConnectVdi(c *Connection, vmRef xenapi.VMRef, vdiRef xenapi.VDIRef, vbdType xenapi.VbdType) (err error) {

	var mode xenapi.VbdMode
	var unpluggable bool
	var bootable bool
	var t xenapi.VbdType
	switch vbdType {
	case xenapi.VbdTypeCD:
		mode = xenapi.VbdModeRO
		bootable = true
		unpluggable = false
		t = xenapi.VbdTypeCD
	case xenapi.VbdTypeDisk:
		mode = xenapi.VbdModeRW
		bootable = false
		unpluggable = false
		t = xenapi.VbdTypeDisk
	case xenapi.VbdTypeFloppy:
		mode = xenapi.VbdModeRW
		bootable = false
		unpluggable = true
		t = xenapi.VbdTypeFloppy
	}

	vbd_ref, err := xenapi.VBD.Create(c.session, xenapi.VBDRecord{
		VM:         xenapi.VMRef(vmRef),
		VDI:        xenapi.VDIRef(vdiRef),
		Userdevice: "autodetect",
		Empty:      false,
		// OtherConfig: map[string]interface{{}},
		QosAlgorithmType: "",
		// QosAlgorithmParams: map[string]interface{{}},
		Mode:        mode,
		Unpluggable: unpluggable,
		Bootable:    bootable,
		Type:        t,
	})

	if err != nil {
		return err
	}

	fmt.Println("VBD Ref:", vbd_ref)

	uuid, err := xenapi.VBD.GetUUID(c.session, vbd_ref)

	fmt.Println("VBD UUID: ", uuid)
	/*
	   // 2. Plug VBD (Non need - the VM hasn't booted.
	   // @todo - check VM state
	   result = APIResult{}
	   err = self.Client.APICall(&result, "VBD.plug", vbd_ref)

	   if err != nil {
	       return err
	   }
	*/
	return
}

func DisconnectVdi(c *Connection, vmRef xenapi.VMRef, vdi xenapi.VDIRef) error {
	vbds, err := xenapi.VM.GetVBDs(c.session, vmRef)
	if err != nil {
		return fmt.Errorf("Unable to get VM VBDs: %s", err.Error())
	}

	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(c.session, vbd)
		if err != nil {
			return fmt.Errorf("Could not get record for VBD '%s': %s", vbd, err.Error())
		}
		recVdi := rec.VDI
		if recVdi == vdi {
			_ = xenapi.VBD.Unplug(c.session, vbd)
			err = xenapi.VBD.Destroy(c.session, vbd)
			if err != nil {
				return fmt.Errorf("Could not destroy VBD '%s': %s", vbd, err.Error())
			}

			return nil
		} else {
			log.Printf("Could not find VDI record in VBD '%s'", vbd)
		}
	}

	return fmt.Errorf("Could not find VBD for VDI '%s'", vdi)
}


func ConnectNetwork(c *Connection, networkRef xenapi.NetworkRef, vmRef xenapi.VMRef, device string) (*xenapi.VIFRef, error) {
	vif, err := xenapi.VIF.Create(c.session, xenapi.VIFRecord{
		Network:     networkRef,
		VM:          vmRef,
		Device:      device,
		LockingMode: xenapi.VifLockingModeNetworkDefault,
	})

	if err != nil {
		return nil, err
	}
	log.Printf("Created the following VIF: %s", vif)

	return &vif, nil
}




// Expose a VDI using the Transfer VM
// (Legacy VHD export)

type TransferRecord struct {
	UrlFull string `xml:"url_full,attr"`
}

func Expose(c *Connection, vdiRef xenapi.VDIRef, format string) (url string, err error) {

	hosts, err := xenapi.Host.GetAll(c.session)

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not retrieve hosts in the pool: %s", err.Error()))
		return "", err
	}
	host := hosts[0]

	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to get VDI uuid for %s: %s", vdiRef, err.Error()))
		return "", err
	}

	args := make(map[string]string)
	args["transfer_mode"] = "http"
	args["vdi_uuid"] = string(vdiRef)
	args["expose_vhd"] = "true"
	args["network_uuid"] = "management"
	args["timeout_minutes"] = "5"

	handle, err := xenapi.Host.CallPlugin(c.session, host, "transfer", "expose", args)

	if err != nil {
		err = errors.New(fmt.Sprintf("Error whilst exposing VDI %s: %s", vdiRef, err.Error()))
		return "", err
	}

	args = make(map[string]string)
	args["record_handle"] = handle
	record_xml, err := xenapi.Host.CallPlugin(c.session, host, "transfer", "get_record", args)

	if err != nil {
		err = errors.New(fmt.Sprintf("Unable to retrieve transfer record for VDI %s: %s", vdiRef, err.Error()))
		return "", err
	}

	var record TransferRecord
	xml.Unmarshal([]byte(record_xml), &record)

	if record.UrlFull == "" {
		return "", errors.New(fmt.Sprintf("Error: did not parse XML properly: '%s'", record_xml))
	}

	// Handles either raw or VHD formats

	switch format {
	case "vhd":
		url = fmt.Sprintf("%s.vhd", record.UrlFull)

	case "raw":
		url = record.UrlFull
	}

	return
}

func Unexpose(c *Connection, vdiRef xenapi.VDIRef) (err error) {

	disk_uuid, err := xenapi.VDI.GetUUID(c.session, vdiRef)

	if err != nil {
		return err
	}

	hosts, err := xenapi.Host.GetAll(c.session)

	if err != nil {
		err = errors.New(fmt.Sprintf("Could not retrieve hosts in the pool: %s", err.Error()))
		return err
	}

	host := hosts[0]

	args := make(map[string]string)
	args["vdi_uuid"] = disk_uuid

	result, err := xenapi.Host.CallPlugin(c.session, host, "transfer", "unexpose", args)

	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Unexpose result: %s", result))

	return nil
}



// Client Initiator
type Connection struct {
	session  *xenapi.Session
	ref	 	 xenapi.SessionRef
	Host     string
	Username string
	Password string
}

func (c Connection) GetSession() *xenapi.Session {
	return c.session
}

func NewXenAPIClient(host, username, password string) (*Connection, error) {
	log.Printf("XenServer Packer Plugin Version: %s", version.PluginVersion.FormattedVersion())

	session := xenapi.NewSession(&xenapi.ClientOpts{
		URL: "http://" + host,
		Headers: map[string]string{
			"User-Agent": fmt.Sprintf("XenServerPacker/%s", version.PluginVersion.FormattedVersion()),
		},
	})


	ref, err := session.LoginWithPassword(username, password, "1.0", "packer")
	if err != nil {
		return nil, err
	}

	return &Connection{session,ref, host, username, password}, nil
}

func (c *Connection) GetSessionRef() xenapi.SessionRef {
	return c.ref
}
