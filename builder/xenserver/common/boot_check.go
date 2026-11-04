// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"log"

	"xenapi"
)

// XAPI matches HVM_boot_params["firmware"] against these exactly; an absent key means BIOS.
const (
	firmwareBIOS = "bios"
	firmwareUEFI = "uefi"
)

// bootCheckAPI is read-only: it runs in the step that starts the VM, where nothing may be changed.
type bootCheckAPI interface {
	Firmware() (string, error)
	VTPMCount() (int, error)
}

type xapiBootCheck struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (b xapiBootCheck) Firmware() (string, error) {
	params, err := xenapi.VM.GetHVMBootParams(b.session, b.vm)
	if err != nil {
		return "", err
	}
	return params["firmware"], nil
}

func (b xapiBootCheck) VTPMCount() (int, error) {
	vtpms, err := xenapi.VM.GetVTPMs(b.session, b.vm)
	if err != nil {
		return 0, err
	}
	return len(vtpms), nil
}

// EnsureVMCanStart returns a halted VM's boot firmware and rejects BIOS with a vTPM, which XenServer cannot start.
func EnsureVMCanStart(c *Connection, vm xenapi.VMRef, firmware string) (string, error) {
	return ensureVMCanStart(xapiBootCheck{session: c.GetSession(), vm: vm}, firmware)
}

func ensureVMCanStart(api bootCheckAPI, firmware string) (string, error) {
	configured := firmware != ""

	if !configured {
		inherited, err := api.Firmware()
		if err != nil {
			// This check only improves an XAPI error, so failing to read must not fail the build.
			log.Printf("Could not read the VM's firmware, skipping the BIOS/vTPM check: %s\n", err)
			return "", nil
		}
		if firmware = inherited; firmware == "" {
			firmware = firmwareBIOS
		}
	}

	if firmware == firmwareUEFI {
		return firmware, nil
	}

	// XAPI rejects any other value with INVALID_VALUE at VM.start, an error that already names it.
	if firmware != firmwareBIOS {
		log.Printf("VM.start will reject firmware %q outright, skipping the BIOS/vTPM check\n", firmware)
		return firmware, nil
	}

	count, err := api.VTPMCount()
	if err != nil {
		log.Printf("Could not read the VM's vtpm state, skipping the BIOS/vTPM check: %s\n", err)
		return firmware, nil
	}
	if count == 0 {
		return firmware, nil
	}

	origin := "Both were inherited from the source the VM was cloned from; the build set neither."
	remedy := fmt.Sprintf("Set firmware = %q, or use a source without a vTPM.", firmwareUEFI)
	if configured {
		origin = fmt.Sprintf("The build set firmware = %q; the vTPM was inherited from the source the VM "+
			"was cloned from.", firmware)
		remedy = fmt.Sprintf("Remove firmware to keep the source's own, set firmware = %q, or use a source "+
			"without a vTPM.", firmwareUEFI)
	}

	// vtpm = false helps only on the clone builder: on the iso builder StepAttachvTPM merely skips.
	return firmware, fmt.Errorf("this VM has a vTPM and will boot with firmware %q, which XenServer refuses "+
		"to start: NOT_IMPLEMENTED [Booting BIOS VM with VTPMs attached]. %s %s On the clone builder, vtpm = "+
		"false removes the inherited vTPM instead, which resolves this without changing the firmware.",
		firmware, origin, remedy)
}
