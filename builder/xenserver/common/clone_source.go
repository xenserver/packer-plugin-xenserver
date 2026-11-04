// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"strings"

	"xenapi"
)

// CloneSource is the template or VM a build clones from, and the option that named it.
type CloneSource struct {
	Option     string
	Kind       string
	Name       string
	IsTemplate bool

	alternative string
}

const TemplateLookupHint = "Run `xe template-list params=name-label` on the pool, or " +
	"look at the template list in XenCenter's New VM wizard, to see what is available. " +
	"The name must match exactly one template."

func TemplateSource(name string) CloneSource {
	return CloneSource{Option: "clone_template", Kind: "template", Name: name, IsTemplate: true,
		alternative: "clone_vm"}
}

// ISOTemplateSource is TemplateSource for the iso builder, which has no clone_vm.
func ISOTemplateSource(name string) CloneSource {
	s := TemplateSource(name)
	s.alternative = "the xenserver-vm-clone builder's clone_vm"
	return s
}

func VMSource(name string) CloneSource {
	return CloneSource{Option: "clone_vm", Kind: "VM", Name: name, alternative: "clone_template"}
}

func (s CloneSource) String() string {
	return fmt.Sprintf("%s '%s'", s.Option, s.Name)
}

// matches counts a snapshot as a VM even though XAPI also sets is_a_template on it, and
// rejects a control domain, which XAPI refuses to clone.
func (s CloneSource) matches(record xenapi.VMRecord) bool {
	if s.IsTemplate {
		return record.IsATemplate && !record.IsASnapshot
	}
	return (!record.IsATemplate || record.IsASnapshot) && !record.IsControlDomain
}

func kindOf(record xenapi.VMRecord) string {
	switch {
	case record.IsControlDomain:
		return "control domain"
	case record.IsASnapshot:
		return "snapshot"
	case record.IsATemplate:
		return "template"
	default:
		return "VM"
	}
}

// ResolveCloneSource resolves the source's name-label to exactly one object of its kind.
func ResolveCloneSource(c *Connection, src CloneSource) (xenapi.VMRef, xenapi.VMRecord, error) {
	var none xenapi.VMRecord

	vms, err := xenapi.VM.GetByNameLabel(c.GetSession(), src.Name)
	if err != nil {
		return "", none, fmt.Errorf("could not look up the %s name-label '%s': %w", src.Kind, src.Name, err)
	}

	var matches []xenapi.VMRef
	var records []xenapi.VMRecord
	var wrongKind []xenapi.VMRecord
	for _, vm := range vms {
		record, err := xenapi.VM.GetRecord(c.GetSession(), vm)
		if err != nil {
			return "", none, fmt.Errorf("could not read the record of a VM named '%s': %w", src.Name, err)
		}
		if src.matches(record) {
			matches = append(matches, vm)
			records = append(records, record)
		} else {
			wrongKind = append(wrongKind, record)
		}
	}

	switch {
	case len(matches) == 1:
		return matches[0], records[0], nil

	case len(matches) > 1:
		return "", none, fmt.Errorf("found %d %ss named '%s'. %s needs a name that is unique among %ss",
			len(matches), src.Kind, src.Name, src.Option, src.Kind)

	case len(wrongKind) > 0:
		if wrongKind[0].IsControlDomain {
			return "", none, fmt.Errorf("%s names a control domain, which XAPI will neither clone nor "+
				"snapshot nor shut down. There is no source here to build from",
				src)
		}
		return "", none, fmt.Errorf("%s names a %s, not a %s. Use %s instead",
			src, kindOf(wrongKind[0]), src.Kind, src.alternative)

	default:
		return "", none, fmt.Errorf("could not find a %s named '%s'", src.Kind, src.Name)
	}
}

// EnsureCloneSourceIsHalted rejects a source that the build cannot turn into a halted VM.
func EnsureCloneSourceIsHalted(src CloneSource, record xenapi.VMRecord, mayCopy bool) error {
	if src.IsTemplate || record.PowerState == xenapi.VMPowerStateHalted {
		return nil
	}

	if record.IsASnapshot {
		if !mayCopy {
			return nil
		}
		return fmt.Errorf("%s names a suspended snapshot, and vm_sr may send this build through VM.copy, "+
			"which XAPI cannot finish on one: the copy inherits is_a_template from the snapshot, so the "+
			"hard_shutdown XAPI runs to settle its power state is refused with VM_IS_TEMPLATE -- once every "+
			"disk has been copied in full, and with the half-built VM left behind. Leave vm_sr unset to clone "+
			"in place, or name a snapshot of a halted VM, which copies without this step", src)
	}

	return fmt.Errorf("%s is %s, and this build needs a halted source: VM.clone and VM.copy take a halted "+
		"VM only, and raise VM_BAD_POWER_STATE on anything else. A suspended one is accepted solely on a pool "+
		"whose other_config sets allow_clone_suspended_vm, and there the clone arrives suspended, which the "+
		"build cannot start. Shut it down, or snapshot it and name the snapshot in %s instead -- a snapshot "+
		"clones while the VM keeps running, at the cost of an image only crash-consistent as of the snapshot",
		src, record.PowerState, src.Option)
}

// bootModeFacts holds what XenCenter's boot-mode rule reads from a clone source's record.
type bootModeFacts struct {
	isHVM bool

	// firmware is HVM_boot_params["firmware"]; XAPI treats an absent key as BIOS.
	firmware string

	// hasBooted reads last_boot_CPU_flags, which clone, copy and import carry over. That is
	// intended: whatever laid out the disk settles the boot mode.
	hasBooted bool

	// declaresVtpm is XenCenter's IsVtpmTemplate without its is_a_template check, because
	// XAPI creates a vTPM from platform:vtpm at start whether or not the source is a template.
	declaresVtpm bool
}

// CloneSourceSupportsFirmware rejects a firmware override that the clone source cannot boot with.
func CloneSourceSupportsFirmware(src CloneSource, record xenapi.VMRecord, firmware string, removesVtpm bool) error {
	return cloneSourceRejectsFirmware(src, firmware, bootModeFactsFromRecord(record), removesVtpm)
}

func bootModeFactsFromRecord(record xenapi.VMRecord) bootModeFacts {
	return bootModeFacts{
		isHVM:        record.HVMBootPolicy != "",
		firmware:     strings.ToLower(strings.TrimSpace(record.HVMBootParams["firmware"])),
		hasBooted:    len(record.LastBootCPUFlags) > 0,
		declaresVtpm: platformDeclaresVTPM(record.Platform),
	}
}

const settledNote = " (What is read is last_boot_CPU_flags. XenServer writes it on a VM's first start and " +
	"never clears it, and VM.clone, VM.copy and import all carry it across -- so a template made from a " +
	"VM that ran counts as settled even though the template itself was never started.)"

// cloneSourceRejectsFirmware ports XenCenter's BootModesControl.UpdateControls. It leaves out
// VM.SupportsUefiBoot(), which would reject every source without a supports-uefi recommendation.
func cloneSourceRejectsFirmware(src CloneSource, firmware string, facts bootModeFacts, removesVtpm bool) error {
	// VM.start rejects any other value with INVALID_VALUE; unset keeps the source's own firmware.
	if firmware != firmwareBIOS && firmware != firmwareUEFI {
		return nil
	}

	// HVM_boot_params means nothing to a PV guest.
	if !facts.isHVM {
		return nil
	}

	if facts.firmware == firmwareUEFI {
		switch {
		case firmware != firmwareBIOS:
			return nil

		case facts.hasBooted:
			return fmt.Errorf("%s boots with UEFI and its boot mode is already settled, so "+
				"firmware = %q cannot work: a source that has run under UEFI carries the boot files a UEFI "+
				"guest loads, and a BIOS guest has no firmware to load them with. The VM would start, find "+
				"nothing to boot, and the build would wait for a guest that never comes up. Remove firmware "+
				"to keep the source's %q, or clone from a source that already boots BIOS.%s",
				src, firmwareBIOS, firmwareUEFI, settledNote)

		// XAPI checks for a vTPM only at start, and the build will have removed it by then.
		case facts.declaresVtpm && !removesVtpm:
			return fmt.Errorf("%s boots with UEFI and declares a vTPM (platform:vtpm = "+
				"true), so firmware = %q cannot work: XenServer implements the vTPM for UEFI guests only, "+
				"and VM.start refuses a BIOS VM that has one (NOT_IMPLEMENTED [Booting BIOS VM with VTPMs "+
				"attached]). Remove firmware to keep the source's %q, or pick a source that boots "+
				"BIOS. On the clone builder, vtpm = false builds without one instead, which destroys "+
				"whatever the inherited vTPM had sealed to it.",
				src, firmwareBIOS, firmwareUEFI)
		}

		return nil
	}

	// No vTPM arm here: a BIOS source cannot have a vTPM.
	if firmware == firmwareUEFI && facts.hasBooted {
		return fmt.Errorf("%s boots with BIOS and its boot mode is already settled, so "+
			"firmware = %q cannot work: a source that has run under BIOS has no EFI system partition, and "+
			"UEFI firmware has nothing to boot from. The VM would start into the UEFI shell, and the build "+
			"would wait for a guest that never comes up. Remove firmware to keep the source's %q, or clone "+
			"from a source that already boots UEFI.%s",
			src, firmwareUEFI, firmwareBIOS, settledNote)
	}

	return nil
}
