// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"

	"xenapi"

	"github.com/xenserver/packer-plugin-xenserver/version"
)

const (
	// ProvenanceKey marks a VM or template as built by this plugin.
	ProvenanceKey = "xs_packer_plugin"

	// ProvenanceVersionKey records the plugin release that produced the object.
	ProvenanceVersionKey = "xs_packer_plugin_version"

	// ProvenanceValue is the value written to ProvenanceKey.
	ProvenanceValue = "true"
)

// PluginVersion returns the release string stamped into other-config.
func PluginVersion() string {
	return version.PluginVersion.FormattedVersion()
}

// BuildProvenanceOtherConfig returns a new map of current, then userSupplied, then the
// provenance keys, last so they override a stale stamp copied by VM.clone and any user key.
func BuildProvenanceOtherConfig(current, userSupplied map[string]string) map[string]string {
	merged := make(map[string]string, len(current)+len(userSupplied)+2)
	for key, value := range current {
		merged[key] = value
	}
	for key, value := range userSupplied {
		merged[key] = value
	}
	merged[ProvenanceKey] = ProvenanceValue
	merged[ProvenanceVersionKey] = PluginVersion()
	return merged
}

// StampVM writes the provenance keys and userSupplied onto instance; call it for every VM.
func StampVM(c *Connection, instance xenapi.VMRef, userSupplied map[string]string) error {
	current, err := xenapi.VM.GetOtherConfig(c.GetSession(), instance)
	if err != nil {
		return fmt.Errorf("error getting VM other-config: %s", err.Error())
	}

	err = xenapi.VM.SetOtherConfig(c.GetSession(), instance, BuildProvenanceOtherConfig(current, userSupplied))
	if err != nil {
		return fmt.Errorf("error setting VM other-config: %s", err.Error())
	}

	return nil
}
