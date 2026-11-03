// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"xenapi"
)

// XenCenter displays the pool-wide network on eth3 as "Network 3"; XAPI's own label is longer.
const defaultNetworkLabelPrefix = "Network "

type namedNetworkAPI interface {
	FindByNameLabel(label string) ([]xenapi.NetworkRef, error)
	AttachVIF(network xenapi.NetworkRef, device string) error
}

type xapiNamedNetwork struct {
	c  *Connection
	vm xenapi.VMRef
}

func (n xapiNamedNetwork) FindByNameLabel(label string) ([]xenapi.NetworkRef, error) {
	return xenapi.Network.GetByNameLabel(n.c.GetSession(), label)
}

func (n xapiNamedNetwork) AttachVIF(network xenapi.NetworkRef, device string) error {
	_, err := ConnectNetwork(n.c, network, n.vm, device)
	return err
}

// AttachNamedNetworks attaches one VIF per name in device order; every failure is fatal.
func AttachNamedNetworks(c *Connection, vm xenapi.VMRef, names []string, say func(string)) error {
	return attachNamedNetworks(xapiNamedNetwork{c: c, vm: vm}, names, say)
}

func attachNamedNetworks(api namedNetworkAPI, names []string, say func(string)) error {
	for i, label := range names {
		say(fmt.Sprintf("Finding network '%s'", label))

		networks, err := api.FindByNameLabel(label)
		if err != nil {
			return fmt.Errorf("could not look up the network name-label %q: %w", label, err)
		}

		if len(networks) == 0 && strings.HasPrefix(label, defaultNetworkLabelPrefix) {
			builtIn := poolWideNetworkLabel(label)
			say(fmt.Sprintf("No network found with name-label '%s'. This might be a default built-in network '%s'", label, builtIn))

			networks, err = api.FindByNameLabel(builtIn)
			if err != nil {
				return fmt.Errorf("could not look up the default network name-label %q: %w", builtIn, err)
			}
		}

		switch {
		case len(networks) == 0:
			return fmt.Errorf("could not find a network with the name-label %q", label)
		case len(networks) > 1:
			return fmt.Errorf("found %d networks with the name-label %q; the name must be unique", len(networks), label)
		}

		if err := api.AttachVIF(networks[0], strconv.Itoa(i)); err != nil {
			return fmt.Errorf("could not attach the VM to network %q on VIF device %d: %w", label, i, err)
		}
	}

	return nil
}

func poolWideNetworkLabel(label string) string {
	return "Pool-wide network associated with eth" + strings.TrimPrefix(label, defaultNetworkLabelPrefix)
}

type vmVIFAPI interface {
	ExistingVIFs() ([]xenapi.VIFRef, error)
	DestroyVIF(vif xenapi.VIFRef) error
}

type xapiVMVIFs struct {
	c  *Connection
	vm xenapi.VMRef
}

func (v xapiVMVIFs) ExistingVIFs() ([]xenapi.VIFRef, error) {
	return xenapi.VM.GetVIFs(v.c.GetSession(), v.vm)
}

func (v xapiVMVIFs) DestroyVIF(vif xenapi.VIFRef) error {
	return xenapi.VIF.Destroy(v.c.GetSession(), vif)
}

// CountVIFs returns how many VIFs the VM has, including any VM.clone copied from the template.
func CountVIFs(c *Connection, vm xenapi.VMRef) (int, error) {
	return countVIFs(xapiVMVIFs{c: c, vm: vm})
}

func countVIFs(api vmVIFAPI) (int, error) {
	vifs, err := api.ExistingVIFs()
	if err != nil {
		return 0, fmt.Errorf("could not list the VM's VIFs: %w", err)
	}

	return len(vifs), nil
}

// DetachAllVIFs destroys every VIF so that network_names does not hit DEVICE_ALREADY_EXISTS.
func DetachAllVIFs(c *Connection, vm xenapi.VMRef, say func(string)) error {
	return detachAllVIFs(xapiVMVIFs{c: c, vm: vm}, say)
}

func detachAllVIFs(api vmVIFAPI, say func(string)) error {
	vifs, err := api.ExistingVIFs()
	if err != nil {
		return fmt.Errorf("could not list the VM's VIFs: %w", err)
	}

	if len(vifs) == 0 {
		return nil
	}

	say(fmt.Sprintf("Removing %d VIF(s) inherited from the source before attaching the requested networks", len(vifs)))

	for _, vif := range vifs {
		if err := api.DestroyVIF(vif); err != nil {
			return fmt.Errorf("could not destroy the inherited VIF %q: %w", vif, err)
		}
	}

	return nil
}

// AttachManagementNetwork gives the VM one VIF, on device 0, on the pool's management network.
func AttachManagementNetwork(c *Connection, vm xenapi.VMRef, say func(string)) error {
	pifs, err := xenapi.PIF.GetAll(c.GetSession())
	if err != nil {
		return fmt.Errorf("could not list the pool's PIFs: %w", err)
	}

	var network xenapi.NetworkRef
	// Every host has a management PIF; the loop deliberately keeps the last one's network.
	for _, pif := range pifs {
		rec, err := xenapi.PIF.GetRecord(c.GetSession(), pif)
		if err != nil {
			return fmt.Errorf("could not read the record of PIF %q: %w", pif, err)
		}
		if rec.Management {
			network = rec.Network
		}
	}

	if network == "" {
		return errors.New("could not find the management network to attach the VM to")
	}

	say("Creating a VIF on the pool's management network")

	if _, err := ConnectNetwork(c, network, vm, "0"); err != nil {
		return fmt.Errorf("could not attach the VM to the management network: %w", err)
	}

	return nil
}
