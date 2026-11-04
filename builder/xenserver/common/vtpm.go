// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"strings"

	"xenapi"
)

// VTPMOutcome says what reconciling a VM's vTPM actually did.
type VTPMOutcome int

const (
	// VTPMUnchanged means the VM already had what the build asked for.
	VTPMUnchanged VTPMOutcome = iota
	// VTPMCreated means a vTPM was added, because the VM had none.
	VTPMCreated
	// VTPMDestroyed means an inherited vTPM was removed, with its secret.
	VTPMDestroyed
	// VTPMDeclarationCleared means only platform:vtpm was cleared; the VM had no vTPM object.
	VTPMDeclarationCleared
)

type vtpmAPI interface {
	List() ([]xenapi.VTPMRef, error)
	Create() (xenapi.VTPMRef, error)
	Destroy(ref xenapi.VTPMRef) error
	Declared() (bool, error)
	ClearDeclaration() error
}

type xapiVTPM struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (v xapiVTPM) List() ([]xenapi.VTPMRef, error) {
	return xenapi.VM.GetVTPMs(v.session, v.vm)
}

// Create asks for a copyable vTPM: a unique one's contents are never copied to clones.
func (v xapiVTPM) Create() (xenapi.VTPMRef, error) {
	return xenapi.VTPM.Create(v.session, v.vm, false)
}

func (v xapiVTPM) Destroy(ref xenapi.VTPMRef) error {
	return xenapi.VTPM.Destroy(v.session, ref)
}

func (v xapiVTPM) Declared() (bool, error) {
	platform, err := xenapi.VM.GetPlatform(v.session, v.vm)
	if err != nil {
		return false, err
	}
	return platformDeclaresVTPM(platform), nil
}

// platformDeclaresVTPM matches XAPI's Vm_platform.is_true: "true" or "1", case-insensitively.
func platformDeclaresVTPM(platform map[string]string) bool {
	switch strings.ToLower(platform["vtpm"]) {
	case "true", "1":
		return true
	default:
		return false
	}
}

func (v xapiVTPM) ClearDeclaration() error {
	return xenapi.VM.RemoveFromPlatform(v.session, v.vm, "vtpm")
}

// ReconcileVTPM gives a halted, freshly cloned VM the vTPM state the build asked for.
func ReconcileVTPM(c *Connection, vm xenapi.VMRef, want bool) (VTPMOutcome, error) {
	return reconcileVTPM(xapiVTPM{session: c.GetSession(), vm: vm}, want)
}

// reconcileVTPM must not touch a VM that already matches: a second vTPM raises
// VTPM_MAX_AMOUNT_REACHED, and destroying one deletes its secret irrecoverably.
func reconcileVTPM(api vtpmAPI, want bool) (VTPMOutcome, error) {
	current, err := api.List()
	if err != nil {
		return VTPMUnchanged, fmt.Errorf("could not read the VM's vtpm state: %w", err)
	}

	if want {
		// An existing vTPM is attached whatever platform:vtpm says, so there is nothing else to write.
		if len(current) == 0 {
			if _, err := api.Create(); err != nil {
				return VTPMUnchanged, fmt.Errorf("could not create the vtpm the build asked for: %w", err)
			}
			return VTPMCreated, nil
		}
		return VTPMUnchanged, nil
	}

	// VM.clone copies platform:vtpm, from which XAPI makes a fresh vTPM at VM.start, so clear it too.
	declared, err := api.Declared()
	if err != nil {
		return VTPMUnchanged, fmt.Errorf("could not read the VM's platform:vtpm state: %w", err)
	}
	if len(current) == 0 && !declared {
		return VTPMUnchanged, nil
	}

	for _, ref := range current {
		if err := api.Destroy(ref); err != nil {
			return VTPMUnchanged, fmt.Errorf("could not destroy the vtpm inherited from the source: %w", err)
		}
	}

	outcome := VTPMDeclarationCleared
	if len(current) > 0 {
		outcome = VTPMDestroyed
	}
	if declared {
		if err := api.ClearDeclaration(); err != nil {
			if len(current) == 0 {
				outcome = VTPMUnchanged
			}
			return outcome, fmt.Errorf("could not clear the VM's platform:vtpm, from which XenServer would "+
				"create a fresh vtpm at start: %w", err)
		}
	}
	return outcome, nil
}
