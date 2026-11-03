// Copyright (c) 2026. Citrix Systems, Inc.

package common

import "xenapi"

// XAPI's add_to_ setters raise MAP_DUPLICATE_KEY on an existing key; remove_from_ ignores a missing one.

func setPlatformKey(session *xenapi.Session, vm xenapi.VMRef, key, value string) error {
	if err := xenapi.VM.RemoveFromPlatform(session, vm, key); err != nil {
		return err
	}
	return xenapi.VM.AddToPlatform(session, vm, key, value)
}

func setHVMBootParam(session *xenapi.Session, vm xenapi.VMRef, key, value string) error {
	if err := xenapi.VM.RemoveFromHVMBootParams(session, vm, key); err != nil {
		return err
	}
	return xenapi.VM.AddToHVMBootParams(session, vm, key, value)
}
