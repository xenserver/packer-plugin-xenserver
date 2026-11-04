// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"

	"xenapi"
)

type vendorDeviceAPI interface {
	Get() (bool, error)
	Set(value bool) error
}

type xapiVendorDevice struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (v xapiVendorDevice) Get() (bool, error) {
	return xenapi.VM.GetHasVendorDevice(v.session, v.vm)
}

func (v xapiVendorDevice) Set(value bool) error {
	return xenapi.VM.SetHasVendorDevice(v.session, v.vm, value)
}

// ApplyVendorDevice sets has_vendor_device on a halted clone and reports whether it changed.
func ApplyVendorDevice(c *Connection, vm xenapi.VMRef, want bool) (changed bool, err error) {
	return applyVendorDevice(xapiVendorDevice{session: c.GetSession(), vm: vm}, want)
}

// applyVendorDevice writes only on a mismatch: changing the flag under installed Windows can crash it.
func applyVendorDevice(api vendorDeviceAPI, want bool) (changed bool, err error) {
	current, err := api.Get()
	if err != nil {
		return false, fmt.Errorf("could not read the VM's has_vendor_device: %w", err)
	}

	if current == want {
		return false, nil
	}

	if err := api.Set(want); err != nil {
		return false, fmt.Errorf("could not set has_vendor_device=%t on the VM: %w", want, err)
	}

	return true, nil
}
