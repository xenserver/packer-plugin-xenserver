// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"

	"xenapi"
)

// "auto", not "true": VM.start resolves it to true or false against the pool's UEFI certificates.
const (
	secureBootAuto = "auto"
	secureBootOff  = "false"
)

// SecureBootOutcome says what setting Secure Boot actually did.
type SecureBootOutcome int

const (
	// SecureBootKept means the build did not set secure_boot, so the inherited value stands.
	SecureBootKept SecureBootOutcome = iota
	// SecureBootDisabled means Secure Boot was turned off.
	SecureBootDisabled
	// SecureBootAuto means Secure Boot was set to auto, enforced only if the pool has UEFI certificates.
	SecureBootAuto
)

// Describe reports the outcome for the build log; source names the clone_template.
func (o SecureBootOutcome) Describe(source string) string {
	switch o {
	case SecureBootDisabled:
		return "Set Secure Boot to Disabled"
	case SecureBootAuto:
		return "Set Secure Boot to Auto"
	}
	return fmt.Sprintf("Keeping the Secure Boot setting inherited from '%s'", source)
}

type secureBootAPI interface {
	Set(value string) error
}

type xapiSecureBoot struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (s xapiSecureBoot) Set(value string) error {
	return setPlatformKey(s.session, s.vm, "secureboot", value)
}

// ReconcileSecureBoot applies secure_boot to a halted clone; nil writes nothing, keeping the inherited value.
func ReconcileSecureBoot(c *Connection, vm xenapi.VMRef, secureBoot *bool) (SecureBootOutcome, error) {
	return reconcileSecureBoot(xapiSecureBoot{session: c.GetSession(), vm: vm}, secureBoot)
}

func reconcileSecureBoot(api secureBootAPI, secureBoot *bool) (SecureBootOutcome, error) {
	if secureBoot == nil {
		return SecureBootKept, nil
	}

	// Prepare already rejects secure_boot = true without UEFI firmware.
	value, outcome := secureBootOff, SecureBootDisabled
	if *secureBoot {
		value, outcome = secureBootAuto, SecureBootAuto
	}

	if err := api.Set(value); err != nil {
		return SecureBootKept, fmt.Errorf("could not set the VM's secure boot state to %q: %w", value, err)
	}
	return outcome, nil
}
