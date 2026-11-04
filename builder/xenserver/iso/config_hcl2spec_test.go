// Copyright (c) 2026. Citrix Systems, Inc.

package iso

import "testing"

// These options only ever configured the dom0 SSH port forward, which no longer exists.
// The plugin is released fresh, so they are gone rather than deprecated, and a template
// still setting one must fail loudly. Packer rejects unknown keys itself; what this
// guards is that nothing reintroduces them into the HCL2 spec.
func TestISOConfigSpec_HasNoOptionsForTheRemovedTunnel(t *testing.T) {
	spec := (&FlatISOConfig{}).HCL2Spec()

	for _, key := range []string{
		"host_port_min",
		"host_port_max",
		"ssh_host_port_min",
		"ssh_host_port_max",
		"ssh_skip_nat_mapping",
	} {
		if _, found := spec[key]; found {
			t.Errorf("%q is still an accepted option, but the dom0 SSH tunnel it configured is gone", key)
		}
	}
}

// The replacements must be present, or a template pinning the address the guest calls
// back on is rejected as an unknown key.
func TestISOConfigSpec_ExposesTheHTTPAddressOptions(t *testing.T) {
	spec := (&FlatISOConfig{}).HCL2Spec()

	for _, key := range []string{"http_bind_address", "http_interface"} {
		if _, found := spec[key]; !found {
			t.Errorf("%q is missing from the HCL2 spec", key)
		}
	}
}

// The install-media options are the SDK's now. iso_checksum_type went with the
// hand-rolled fields, replaced by the type prefix of iso_checksum; iso_target_path and
// iso_target_extension come with the SDK struct and must be accepted.
func TestISOConfigSpec_UsesTheSDKInstallMediaOptions(t *testing.T) {
	spec := (&FlatISOConfig{}).HCL2Spec()

	if _, found := spec["iso_checksum_type"]; found {
		t.Error("iso_checksum_type is still accepted; the type belongs in iso_checksum as a \"sha256:\" prefix")
	}

	for _, key := range []string{
		"iso_checksum",
		"iso_url",
		"iso_urls",
		"iso_target_path",
		"iso_target_extension",
	} {
		if _, found := spec[key]; !found {
			t.Errorf("%q is missing from the HCL2 spec", key)
		}
	}
}

// host_ssh_enabled_timeout came in with StepEnableHostSSH, which went away with the rest
// of the dom0 SSH path; the option must not linger as an accepted no-op.
func TestISOConfigSpec_HasNoHostSSHTimeout(t *testing.T) {
	spec := (&FlatISOConfig{}).HCL2Spec()

	if _, found := spec["host_ssh_enabled_timeout"]; found {
		t.Error("host_ssh_enabled_timeout is still accepted, but nothing enables dom0 SSH any more")
	}
}
