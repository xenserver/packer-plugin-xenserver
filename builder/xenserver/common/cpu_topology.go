// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"strconv"

	"xenapi"
)

const coresPerSocketKey = "cores-per-socket"

// maxSockets is Xen's limit on sockets per VM, which XenCenter enforces and XAPI does not.
const maxSockets = 16

type topologyKind int

// topologyUntouched is the zero value: the VM was never read.
const (
	topologyUntouched topologyKind = iota
	topologyKept
	topologyWritten
	topologyReset
	topologyAbsent
)

// TopologyOutcome says what writing the CPU topology actually did.
type TopologyOutcome struct {
	kind      topologyKind
	cores     uint
	vcpus     uint
	inherited string
	// Separate from inherited: a key holding an empty value reads the same as no key.
	found     bool
	effective uint
}

// Describe reports the outcome for the build log; source names the clone_template.
func (o TopologyOutcome) Describe(source string) string {
	switch o.kind {
	case topologyWritten:
		return fmt.Sprintf("Set the CPU topology to %d core(s) per socket", o.cores)

	case topologyKept:
		return fmt.Sprintf("Keeping the CPU topology of '%s': %s core(s) per socket", source, o.inherited)

	case topologyAbsent:
		return fmt.Sprintf("'%s' carries no CPU topology of its own, so the VM gets one core per socket", source)

	case topologyReset:
		switch {
		case !o.found:
			return fmt.Sprintf("'%s' carries no CPU topology of its own, and one core per socket would give "+
				"the %d vcpus this build asked for more than %d sockets, so the VM gets %d core(s) per "+
				"socket. Set cores_per_socket to choose the topology.", source, o.vcpus, maxSockets, o.cores)

		case o.effective == 0:
			return fmt.Sprintf("The cores per socket of '%s' is %q, which is not a usable number of cores, so "+
				"the VM gets %d core(s) per socket. Set cores_per_socket to choose the topology.",
				source, o.inherited, o.cores)
		}
		return fmt.Sprintf("The %s cores per socket of '%s' cannot be kept for the %d vcpus this build asked "+
			"for, so the VM gets %d core(s) per socket. Set cores_per_socket to choose the topology.",
			o.inherited, source, o.vcpus, o.cores)
	}
	return fmt.Sprintf("Leaving the CPU topology of '%s' as it is", source)
}

// topologyFits mirrors XenCenter's VM.ValidVCPUConfiguration.
func topologyFits(vcpus, coresPerSocket uint) bool {
	return coresPerSocket > 0 && vcpus%coresPerSocket == 0 && vcpus/coresPerSocket <= maxSockets
}

// fallbackCoresPerSocket is the fewest cores per socket that fits, as XenCenter's New VM wizard picks.
func fallbackCoresPerSocket(vcpus uint) uint {
	if vcpus == 0 {
		return 1
	}

	for cores := uint(1); cores < vcpus; cores++ {
		if topologyFits(vcpus, cores) {
			return cores
		}
	}

	return vcpus
}

// cpuTopologyAPI is an interface only so that tests can fake XenAPI.
type cpuTopologyAPI interface {
	// CoresPerSocket reads platform:cores-per-socket; found is false when the key is absent.
	CoresPerSocket() (string, bool, error)
	// SetCoresPerSocket writes platform:cores-per-socket, replacing any existing value.
	SetCoresPerSocket(cores uint) error
}

type xapiCPUTopology struct {
	session *xenapi.Session
	vm      xenapi.VMRef
}

func (t xapiCPUTopology) CoresPerSocket() (string, bool, error) {
	platform, err := xenapi.VM.GetPlatform(t.session, t.vm)
	if err != nil {
		return "", false, err
	}

	value, found := platform[coresPerSocketKey]

	return value, found, nil
}

func (t xapiCPUTopology) SetCoresPerSocket(cores uint) error {
	return setPlatformKey(t.session, t.vm, coresPerSocketKey, strconv.FormatUint(uint64(cores), 10))
}

// ReconcileCPUTopology applies the build's CPU topology to a halted VM; call it after writing VCPUs_max.
func ReconcileCPUTopology(c *Connection, vm xenapi.VMRef, vcpus, coresPerSocket *uint) (TopologyOutcome, error) {
	return reconcileCPUTopology(xapiCPUTopology{session: c.GetSession(), vm: vm}, vcpus, coresPerSocket)
}

func reconcileCPUTopology(api cpuTopologyAPI, vcpus, coresPerSocket *uint) (TopologyOutcome, error) {
	// cores_per_socket alone is ignored, as Prepare warns: XenServer refuses to start a VM whose
	// cores per socket does not divide its vcpus, and the inherited vcpus are not known here.
	if vcpus == nil {
		return TopologyOutcome{}, nil
	}

	if coresPerSocket != nil {
		if err := api.SetCoresPerSocket(*coresPerSocket); err != nil {
			return TopologyOutcome{}, fmt.Errorf("could not set the VM's cores per socket to %d: %w",
				*coresPerSocket, err)
		}

		return TopologyOutcome{kind: topologyWritten, cores: *coresPerSocket}, nil
	}

	inherited, found, err := api.CoresPerSocket()
	if err != nil {
		return TopologyOutcome{}, fmt.Errorf("could not read the VM's CPU topology: %w", err)
	}

	// XenServer treats an absent key as one core per socket and refuses to start a VM with an unparsable one.
	effective := uint(1)
	if found {
		cores, err := strconv.ParseUint(inherited, 10, 32)
		if err != nil {
			effective = 0
		} else {
			effective = uint(cores)
		}
	}

	if topologyFits(*vcpus, effective) {
		if !found {
			return TopologyOutcome{kind: topologyAbsent}, nil
		}
		return TopologyOutcome{kind: topologyKept, inherited: inherited}, nil
	}

	cores := fallbackCoresPerSocket(*vcpus)
	if err := api.SetCoresPerSocket(cores); err != nil {
		return TopologyOutcome{}, fmt.Errorf("could not set the VM's cores per socket to %d: %w", cores, err)
	}

	return TopologyOutcome{
		kind: topologyReset, cores: cores, vcpus: *vcpus,
		inherited: inherited, found: found, effective: effective,
	}, nil
}
