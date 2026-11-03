// Copyright (c) 2026. Citrix Systems, Inc.

package iso

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	commonsteps "github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	xscommon "github.com/xenserver/packer-plugin-xenserver/builder/xenserver/common"
	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/internal/xstest"
)

const testISOSHA256 = "e907d92eeec9df64163a7e454cbc8d7755e8ddc7ed42f99dbc80c40f1a138433"

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"xenserver_host":     "localhost",
		"xenserver_username": "admin",
		"xenserver_password": "admin",
		"vm_name":            "foo",
		"iso_sr":             "ISO library",
		"iso_checksum":       "sha256:" + testISOSHA256,
		"iso_url":            "http://www.google.com/",
		"shutdown_command":   "yes",
		"ssh_username":       "foo",

		"clone_template": "Windows 11",

		"communicator": "ssh",

		"packer_build_name": "foo",
		"insecure":          true,
	}
}

func clearConnectionEnv(t *testing.T) {
	t.Helper()
	for _, key := range []string{
		xscommon.EnvHost,
		xscommon.EnvUsername,
		xscommon.EnvPassword,
		xscommon.EnvCACertPath,
		xscommon.EnvInsecure,
	} {
		t.Setenv(key, "")
	}
}

func TestBuilderPrepare_Defaults(t *testing.T) {
	var b Builder
	config := testConfig()
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	// No tools ISO unless the user names one: XenServer 8 and later ship none.
	if b.config.GuestToolsISO != "" {
		t.Errorf("bad guest tools ISO name: %s", b.config.GuestToolsISO)
	}

	if b.config.Name() == "" {
		t.Errorf("bad vm name: %s", b.config.Name())
	}

	var noName Builder
	nameless := testConfig()
	delete(nameless, "vm_name")
	if _, _, err := noName.Prepare(nameless); err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if !strings.HasPrefix(noName.config.Name(), "packer-foo-") || strings.Contains(noName.config.Name(), "{{") {
		t.Errorf("bad default vm name: %s", noName.config.Name())
	}

	if b.config.Format != "xva" {
		t.Errorf("bad format: %s", b.config.Format)
	}

	if b.config.KeepVM != "never" {
		t.Errorf("bad keep instance: %s", b.config.KeepVM)
	}

	xstest.WantUnset(t, "vcpus", b.config.VCPUs)
	xstest.WantUnset(t, "cores_per_socket", b.config.CorePerSocket)
	xstest.WantUnset(t, "vm_memory_gb", b.config.VMMemoryGB)
	xstest.WantUnset(t, "disk_size_gb", b.config.DiskSizeGB)
}

func TestBuilderPrepare_RenamedKeysCarryTheirValues(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vm_sr"] = "Local storage"
	config["iso_sr"] = "ISO library"
	config["vm_memory_gb"] = 8
	config["disk_size_gb"] = 60
	config["guest_tools_iso"] = "xs-tools.iso"
	config["vtpm"] = true
	config["firmware"] = "uefi" // not a rename; vtpm is refused without it

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	if b.config.HostIp != "localhost" {
		t.Errorf("xenserver_host = %q, want \"localhost\"", b.config.HostIp)
	}
	if b.config.Username != "admin" {
		t.Errorf("xenserver_username = %q, want \"admin\"", b.config.Username)
	}
	if b.config.Password != "admin" {
		t.Errorf("xenserver_password = %q, want \"admin\"", b.config.Password)
	}
	if b.config.VMSR != "Local storage" {
		t.Errorf("vm_sr = %q, want \"Local storage\"", b.config.VMSR)
	}
	if b.config.ISOSR != "ISO library" {
		t.Errorf("iso_sr = %q, want \"ISO library\"", b.config.ISOSR)
	}
	xstest.WantUint(t, "vm_memory_gb", b.config.VMMemoryGB, 8)
	xstest.WantUint(t, "disk_size_gb", b.config.DiskSizeGB, 60)
	if b.config.GuestToolsISO != "xs-tools.iso" {
		t.Errorf("guest_tools_iso = %q, want \"xs-tools.iso\"", b.config.GuestToolsISO)
	}
	if !b.config.VTPMEnabled {
		t.Error("vtpm = false, want true")
	}
}

func TestBuilderPrepare_CACertPathReplacesServerCert(t *testing.T) {
	clearConnectionEnv(t)
	caCertPath := xstest.WriteCACert(t)

	var b Builder
	config := testConfig()
	delete(config, "insecure")
	config["ca_cert_path"] = caCertPath

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	if b.config.CACertPath != caCertPath {
		t.Errorf("ca_cert_path = %q, want %q", b.config.CACertPath, caCertPath)
	}
}

func TestBuilderPrepare_MissingCACertPathNamesTheNewOption(t *testing.T) {
	clearConnectionEnv(t)

	var b Builder
	config := testConfig()
	delete(config, "insecure")

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("a build that verifies certificates but names none must be rejected")
	}
	if !strings.Contains(err.Error(), "ca_cert_path") {
		t.Errorf("error should name ca_cert_path, got: %s", err)
	}
}

func TestBuilderPrepare_ConnectionOptionsFallBackToTheEnvironment(t *testing.T) {
	clearConnectionEnv(t)
	t.Setenv(xscommon.EnvHost, "pool.example.com")
	t.Setenv(xscommon.EnvUsername, "root")
	t.Setenv(xscommon.EnvPassword, "from-the-environment")
	t.Setenv(xscommon.EnvInsecure, "true")

	var b Builder
	config := testConfig()
	for _, key := range []string{"xenserver_host", "xenserver_username", "xenserver_password", "insecure"} {
		delete(config, key)
	}

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	if b.config.HostIp != "pool.example.com" {
		t.Errorf("xenserver_host = %q, want it from the environment", b.config.HostIp)
	}
	if b.config.Username != "root" {
		t.Errorf("xenserver_username = %q, want it from the environment", b.config.Username)
	}
	if b.config.Password != "from-the-environment" {
		t.Errorf("xenserver_password = %q, want it from the environment", b.config.Password)
	}
	if !b.config.InsecureEnabled() {
		t.Error("insecure did not come from the environment")
	}
}

func TestBuilderPrepare_RejectsOldKeys(t *testing.T) {
	oldKeys := map[string]interface{}{
		"remote_host":     "localhost",
		"remote_username": "admin",
		"remote_password": "admin",
		"sr_name":         "Local storage",
		"sr_iso_name":     "ISO library",
		// Ignoring it instead would make unverified templates look verified.
		"iso_checksum_type":      "sha256",
		"vm_memory":              1024,
		"disk_size":              40000,
		"tools_iso_name":         "xs-tools.iso",
		"server_cert":            "/etc/ssl/certs/pool.pem",
		"server_cert_file":       "/etc/ssl/certs/pool.pem",
		"skip_cert_verification": true,
	}

	for key, value := range oldKeys {
		var b Builder
		config := testConfig()
		config[key] = value

		_, _, err := b.Prepare(config)
		if err == nil {
			t.Errorf("%s: the old name must be rejected as unknown", key)
			continue
		}
		if !strings.Contains(err.Error(), key) {
			t.Errorf("%s: error should name the unknown key, got: %s", key, err)
		}
	}
}

func TestBuilderPrepare_OldVTPMSpellingStillDecodes(t *testing.T) {
	var b Builder
	config := testConfig()
	config["firmware"] = "uefi"
	config["vTPM"] = true

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("mapstructure matches keys case-insensitively, so vTPM should still decode: %s", err)
	}
	if !b.config.VTPMEnabled {
		t.Error("vTPM did not reach the vtpm field")
	}
}

func TestBuilderPrepare_RejectsDroppedKeys(t *testing.T) {
	droppedKeys := map[string]interface{}{
		"floppy_files":         []string{"foo"},
		"export_network_names": []string{"foo"},
		"ip_getter":            "tools",
		"platform_args":        map[string]string{"viridian": "false"},
		"ssh_skip_nat_mapping": true,
		"host_port_min":        5900,
		"host_port_max":        6000,
		"http_port_min":        8000,
		"http_port_max":        9000,
		"ssh_host_port_min":    2222,
		"ssh_host_port_max":    4444,
		"source_path":          "/tmp/foo.xva",
		"convert_to_template":  true,
		"ssh_key_path":         "/tmp/id_rsa",
		"vcpus_max":            4,
		"vcpus_atstartup":      2,
	}

	for key, value := range droppedKeys {
		var b Builder
		config := testConfig()
		config[key] = value

		_, _, err := b.Prepare(config)
		if err == nil {
			t.Errorf("%s: a dropped option must be rejected", key)
			continue
		}
		if !strings.Contains(err.Error(), key) {
			t.Errorf("%s: error should name the unknown key, got: %s", key, err)
		}
	}
}

func TestBuilderPrepare_RejectsMalformedDurations(t *testing.T) {
	for _, key := range []string{
		"boot_wait",
		"boot_key_group_interval",
		"install_timeout",
		"winrm_timeout",
		"pause_before_connecting",
	} {
		var b Builder
		config := testConfig()
		config[key] = "5 seconds"

		_, _, err := b.Prepare(config)
		if err == nil {
			t.Errorf("%s: a malformed duration must be rejected", key)
			continue
		}
		if !strings.Contains(err.Error(), key) {
			t.Errorf("%s: error should name the field, got: %s", key, err)
		}
	}
}

func TestBuilderPrepare_RejectsAnUnknownCommunicator(t *testing.T) {
	var b Builder
	config := testConfig()
	config["communicator"] = "telnet"

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("an unrecognised communicator must be rejected")
	}
	if !strings.Contains(err.Error(), "telnet") {
		t.Errorf("error should name the rejected value, got: %s", err)
	}
}

func TestBuilderPrepare_RejectsAnUnparseableISOURL(t *testing.T) {
	var b Builder
	config := testConfig()
	config["iso_url"] = "://no-scheme"

	if _, _, err := b.Prepare(config); err == nil {
		t.Fatal("an unparsable iso_url must be rejected")
	}
}

func TestBuilderPrepare_VCPUs(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vcpus"] = 4

	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	xstest.WantUint(t, "vcpus", b.config.VCPUs, 4)
	xstest.WantUnset(t, "cores_per_socket", b.config.CorePerSocket)
}

func TestBuilderPrepare_SSHPrivateKeyFileIsAccepted(t *testing.T) {
	var b Builder
	config := testConfig()
	config["ssh_private_key_file"] = "/nonexistent/id_rsa"

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("the SDK validates ssh_private_key_file and should reject a missing file")
	}
	// An unknown key would fail with an "unknown configuration key" error instead.
	if strings.Contains(err.Error(), "unknown configuration key") {
		t.Errorf("ssh_private_key_file must be a known option, got: %s", err)
	}
}

func TestBuilderPrepare_GuestToolsISO(t *testing.T) {
	var b Builder
	config := testConfig()

	config["guest_tools_iso"] = "xs-tools.iso"
	_, warns, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	if b.config.GuestToolsISO != "xs-tools.iso" {
		t.Errorf("bad guest tools ISO name: %s", b.config.GuestToolsISO)
	}

	if len(warns) != 1 || !strings.Contains(warns[0], "guest_tools_iso is deprecated") {
		t.Errorf("expected a deprecation warning, got: %#v", warns)
	}
}

func TestBuilderPrepare_DiskSizeGB(t *testing.T) {
	var b Builder
	config := testConfig()

	delete(config, "disk_size_gb")
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("bad err: %s", err)
	}

	xstest.WantUnset(t, "disk_size_gb", b.config.DiskSizeGB)

	config["disk_size_gb"] = 60
	b = Builder{}
	_, warns, err = b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	xstest.WantUint(t, "disk_size_gb", b.config.DiskSizeGB, 60)
}

func TestBuilderPrepare_Format(t *testing.T) {
	var b Builder
	config := testConfig()

	config["format"] = "foo"
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err == nil {
		t.Fatal("should have error")
	}

	for _, format := range []string{"xva", "vdi_raw"} {
		config["format"] = format
		b = Builder{}
		_, warns, err = b.Prepare(config)
		if len(warns) > 0 {
			t.Fatalf("bad: %#v", warns)
		}
		if err != nil {
			t.Fatalf("%s: should not have error: %s", format, err)
		}
		if b.config.Format != format {
			t.Errorf("format = %q, want %q", b.config.Format, format)
		}
	}
}

func TestBuilderPrepare_SharedOptionsMovedIntoCommonConfig(t *testing.T) {
	var b Builder
	config := testConfig()
	config["cores_per_socket"] = 2
	config["vcpus"] = 4
	config["vgpu_profile"] = "GRID T4-2Q"
	config["install_timeout"] = "45m"
	config["firmware"] = "uefi"
	config["secure_boot"] = true
	config["create_snapshot"] = true
	config["snapshot_name"] = "pre-export"
	config["vm_other_config"] = map[string]string{"owner": "packer"}

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	xstest.WantUint(t, "cores_per_socket", b.config.CorePerSocket, 2)
	if b.config.VGPUProfile != "GRID T4-2Q" {
		t.Errorf("vgpu_profile = %q, want \"GRID T4-2Q\"", b.config.VGPUProfile)
	}
	if b.config.InstallTimeout != 45*time.Minute {
		t.Errorf("install_timeout = %s, want 45m", b.config.InstallTimeout)
	}
	if b.config.Firmware != "uefi" {
		t.Errorf("firmware = %q, want \"uefi\"", b.config.Firmware)
	}
	if !b.config.SecureBootEnabled() {
		t.Error("secure_boot = false, want true")
	}
	if !b.config.CreateSnapshot {
		t.Error("create_snapshot = false, want true")
	}
	if b.config.SnapshotName != "pre-export" {
		t.Errorf("snapshot_name = %q, want \"pre-export\"", b.config.SnapshotName)
	}
	if b.config.VMOtherConfig["owner"] != "packer" {
		t.Errorf("vm_other_config = %v, want owner=packer", b.config.VMOtherConfig)
	}
}

func TestBuilderPrepare_SecureBootRequiresUEFI(t *testing.T) {
	for _, firmware := range []string{"", "bios"} {
		var b Builder
		config := testConfig()
		config["secure_boot"] = true
		if firmware != "" {
			config["firmware"] = firmware
		}

		_, _, err := b.Prepare(config)
		if err == nil {
			t.Errorf("firmware=%q: secure_boot without UEFI must be rejected", firmware)
			continue
		}
		if !strings.Contains(err.Error(), "secure_boot") {
			t.Errorf("firmware=%q: error should name secure_boot, got: %s", firmware, err)
		}
	}

	var b Builder
	config := testConfig()
	config["firmware"] = "uefi"
	config["secure_boot"] = true
	if _, _, err := b.Prepare(config); err != nil {
		t.Errorf("secure_boot with uefi should be accepted: %s", err)
	}
}

func TestBuilderPrepare_UnsetSecureBootIsOff(t *testing.T) {
	var b Builder
	config := testConfig()

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.SecureBootEnabled() {
		t.Error("secure_boot defaulted to on")
	}
}

func TestBuilderPrepare_HasVendorDevice(t *testing.T) {
	var b Builder
	config := testConfig()

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.HasVendorDevice != nil {
		t.Errorf("has_vendor_device = %t, want the pointer left nil when no template names it", *b.config.HasVendorDevice)
	}
	if !b.config.HasVendorDeviceEnabled() {
		t.Error("an unset has_vendor_device must still mean on")
	}

	for _, want := range []bool{true, false} {
		b = Builder{}
		config = testConfig()
		config["has_vendor_device"] = want

		if _, _, err := b.Prepare(config); err != nil {
			t.Fatalf("has_vendor_device=%t: should not have error: %s", want, err)
		}
		if b.config.HasVendorDevice == nil {
			t.Fatalf("has_vendor_device=%t never reached the field", want)
		}
		if got := b.config.HasVendorDeviceEnabled(); got != want {
			t.Errorf("HasVendorDeviceEnabled() = %t, want %t", got, want)
		}
	}
}

func TestBuilderPrepare_InstallTimeoutDefault(t *testing.T) {
	var b Builder
	config := testConfig()

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	if b.config.InstallTimeout != 200*time.Minute {
		t.Errorf("install_timeout = %s, want 200m", b.config.InstallTimeout)
	}
}

func TestBuilderPrepare_InvalidKey(t *testing.T) {
	var b Builder
	config := testConfig()

	config["i_should_not_be_valid"] = true
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}

func TestBuilderPrepare_ISOChecksumForms(t *testing.T) {
	sums := filepath.Join(t.TempDir(), "SHA256SUMS")
	if err := os.WriteFile(sums, []byte(testISOSHA256+"  test.iso\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name     string
		checksum string
		want     string
		wantWarn bool
	}{
		{"typed", "sha256:" + testISOSHA256, "sha256:" + testISOSHA256, false},
		{"untyped", testISOSHA256, "sha256:" + testISOSHA256, false},
		{"mixed case", "SHA256:" + strings.ToUpper(testISOSHA256), "sha256:" + testISOSHA256, false},
		{"file", "file:" + filepath.ToSlash(sums), "sha256:" + testISOSHA256, false},
		{"none", "none", "none", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var b Builder
			config := testConfig()
			config["iso_url"] = "http://example.com/test.iso"
			config["iso_checksum"] = tc.checksum

			_, warns, err := b.Prepare(config)
			if err != nil {
				t.Fatalf("should not have error: %s", err)
			}
			if b.config.ISOChecksum != tc.want {
				t.Errorf("iso_checksum = %q, want %q", b.config.ISOChecksum, tc.want)
			}
			if gotWarn := len(warns) > 0; gotWarn != tc.wantWarn {
				t.Errorf("warnings %v, want a warning: %t", warns, tc.wantWarn)
			}
		})
	}
}

func TestBuilderPrepare_RejectsAnUnusableISOChecksum(t *testing.T) {
	cases := map[string]string{
		"empty":                       "",
		"unknown type":                "crc32:" + testISOSHA256,
		"not hex":                     "sha256:not-a-hash",
		"no algorithm of that length": "abcdef",
		// The classic mistake of naming the ISO instead of its checksum file.
		"an ISO name": "ubuntu-24.04.iso",
		// The file: form is read at prepare time, before any download.
		"file: pointing at nothing": "file:" + filepath.ToSlash(filepath.Join(t.TempDir(), "absent.txt")),
	}

	for name, checksum := range cases {
		t.Run(name, func(t *testing.T) {
			var b Builder
			config := testConfig()
			config["iso_checksum"] = checksum

			if _, _, err := b.Prepare(config); err == nil {
				t.Fatalf("iso_checksum = %q should be rejected", checksum)
			}
		})
	}
}

func TestBuilderPrepare_ISOUrls(t *testing.T) {
	t.Run("iso_url becomes the whole list", func(t *testing.T) {
		var b Builder
		config := testConfig()
		config["iso_url"] = "http://www.packer.io/test.iso"

		if _, _, err := b.Prepare(config); err != nil {
			t.Fatalf("should not have error: %s", err)
		}

		expected := []string{"http://www.packer.io/test.iso"}
		if !reflect.DeepEqual(b.config.ISOUrls, expected) {
			t.Fatalf("iso_urls = %#v, want %#v", b.config.ISOUrls, expected)
		}
	})

	t.Run("iso_urls keeps its order", func(t *testing.T) {
		var b Builder
		config := testConfig()
		delete(config, "iso_url")
		config["iso_urls"] = []string{
			"http://www.packer.io/test.iso",
			"http://www.hashicorp.com/test.iso",
		}

		if _, _, err := b.Prepare(config); err != nil {
			t.Fatalf("should not have error: %s", err)
		}

		expected := []string{
			"http://www.packer.io/test.iso",
			"http://www.hashicorp.com/test.iso",
		}
		if !reflect.DeepEqual(b.config.ISOUrls, expected) {
			t.Fatalf("iso_urls = %#v, want %#v", b.config.ISOUrls, expected)
		}
	})

	t.Run("both is rejected", func(t *testing.T) {
		var b Builder
		config := testConfig()
		config["iso_urls"] = []string{"http://www.packer.io/test.iso"}

		if _, _, err := b.Prepare(config); err == nil {
			t.Fatal("iso_url together with iso_urls must be rejected")
		}
	})
}

func TestBuilderPrepare_ISONameAndISOUrlAreExclusive(t *testing.T) {
	var b Builder
	config := testConfig()
	config["iso_name"] = "Win11.iso"

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("iso_name together with iso_url must be rejected")
	}
	if !strings.Contains(err.Error(), "only one of iso_name or iso_url") {
		t.Errorf("error should say the two are exclusive, got: %s", err)
	}
}

func TestBuilderPrepare_RequiresAnISONameOrAnISOUrl(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "iso_url")
	delete(config, "iso_checksum")

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("a build with no installation media must be rejected")
	}
	if !strings.Contains(err.Error(), "one of iso_name or iso_url") {
		t.Errorf("error should name both ways of supplying an ISO, got: %s", err)
	}
}

func TestBuilderPrepare_ISONameNeedsNoChecksum(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "iso_url")
	delete(config, "iso_checksum")
	config["iso_name"] = "Win11.iso"

	_, warns, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if len(warns) > 0 {
		t.Fatalf("unexpected warnings: %#v", warns)
	}
	if b.config.ISOName != "Win11.iso" {
		t.Errorf("iso_name = %q, want Win11.iso", b.config.ISOName)
	}
}

func TestBuilderPrepare_RequiresAnISOSR(t *testing.T) {
	// iso_sr is normally a read-only ISO library, so a downloaded ISO goes to vm_sr.
	t.Run("url mode needs none", func(t *testing.T) {
		var b Builder
		config := testConfig()
		delete(config, "iso_sr")

		if _, _, err := b.Prepare(config); err != nil {
			t.Fatalf("iso_url without iso_sr should be accepted, got: %v", err)
		}
	})

	t.Run("name mode", func(t *testing.T) {
		var b Builder
		config := testConfig()
		delete(config, "iso_sr")
		delete(config, "iso_url")
		delete(config, "iso_checksum")
		config["iso_name"] = "Win11.iso"

		_, _, err := b.Prepare(config)
		if err == nil || !strings.Contains(err.Error(), "iso_sr must be specified") {
			t.Fatalf("iso_name without iso_sr should be rejected, got: %v", err)
		}
	})
}

func TestBuilderPrepare_ISOTargetOptions(t *testing.T) {
	var b Builder
	config := testConfig()
	config["iso_target_path"] = "cache/ubuntu.iso"
	config["iso_target_extension"] = "ISO"

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.TargetPath != "cache/ubuntu.iso" {
		t.Errorf("iso_target_path = %q, want cache/ubuntu.iso", b.config.TargetPath)
	}
	if b.config.TargetExtension != "iso" {
		t.Errorf("iso_target_extension = %q, want iso", b.config.TargetExtension)
	}
}

func TestBuilderPrepare_RejectsDownloadOptionsInNameMode(t *testing.T) {
	cases := map[string]string{
		"iso_checksum":         "sha256:" + testISOSHA256,
		"iso_target_path":      "cache/win11.iso",
		"iso_target_extension": "iso",
	}

	for option, value := range cases {
		t.Run(option, func(t *testing.T) {
			var b Builder
			config := testConfig()
			delete(config, "iso_url")
			delete(config, "iso_checksum")
			config["iso_name"] = "Win11.iso"
			config[option] = value

			_, _, err := b.Prepare(config)
			if err == nil {
				t.Fatalf("%s should be rejected alongside iso_name", option)
			}
			if !strings.Contains(err.Error(), option) {
				t.Errorf("error should name %s, got: %s", option, err)
			}
		})
	}
}

func TestBuilderDownloadSteps_CarryTheResolvedISOOptions(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "iso_url")
	config["iso_urls"] = []string{"http://mirror.example.com/test.iso", "http://backup.example.com/test.iso"}
	config["iso_checksum"] = "SHA256:" + strings.ToUpper(testISOSHA256)
	config["iso_target_path"] = "cache/test.iso"
	config["iso_target_extension"] = "img"

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	steps := b.downloadSteps()
	if len(steps) != 1 {
		t.Fatalf("got %d download steps, want 1", len(steps))
	}
	step, ok := steps[0].(*commonsteps.StepDownload)
	if !ok {
		t.Fatalf("download step is a %T, want *commonsteps.StepDownload", steps[0])
	}

	if step.Checksum != "sha256:"+testISOSHA256 {
		t.Errorf("checksum = %q, want the canonical sha256:%s", step.Checksum, testISOSHA256)
	}
	if !reflect.DeepEqual(step.Url, config["iso_urls"]) {
		t.Errorf("urls = %v, want %v in order", step.Url, config["iso_urls"])
	}
	if step.TargetPath != "cache/test.iso" {
		t.Errorf("target path = %q, want cache/test.iso", step.TargetPath)
	}
	if step.Extension != "img" {
		t.Errorf("extension = %q, want img", step.Extension)
	}
}

func TestBuilderPrepare_ISOTargetExtensionDefaultsToISO(t *testing.T) {
	var b Builder
	config := testConfig()

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.TargetExtension != "iso" {
		t.Errorf("iso_target_extension = %q, want iso", b.config.TargetExtension)
	}
}

func TestBuilderPrepare_KeepVM(t *testing.T) {
	var b Builder
	config := testConfig()

	config["keep_vm"] = "foo"
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err == nil {
		t.Fatal("should have error")
	}

	config["keep_vm"] = "always"
	b = Builder{}
	_, warns, err = b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
}

func TestBuilderPrepare_DefaultCommunicatorIsNone(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "communicator")
	delete(config, "ssh_username")
	delete(config, "shutdown_command")

	_, _, err := b.Prepare(config)

	if err != nil {
		t.Fatalf("a template with no communicator should be accepted: %s", err)
	}
	if b.config.Comm.Type != "none" {
		t.Errorf("communicator = %q, want \"none\"", b.config.Comm.Type)
	}
}

func TestBuilderPrepare_CommunicatorNoneWaitsForTheGuestToPowerItselfOff(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "communicator")
	delete(config, "ssh_username")
	delete(config, "shutdown_command")
	delete(config, "cd_files")
	delete(config, "boot_command")

	_, warns, err := b.Prepare(config)

	if err != nil {
		t.Fatalf("an unattended install with no communicator is the supported case: %s", err)
	}
	var found bool
	for _, w := range warns {
		if strings.Contains(w, "power itself off") && strings.Contains(w, "install_timeout") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a warning naming the self-shutdown requirement and install_timeout, got %v", warns)
	}
}

func TestBuilderPrepare_DefaultCommunicatorRejectsShutdownCommand(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "communicator")

	_, _, err := b.Prepare(config)

	if err == nil {
		t.Fatal("shutdown_command with the default communicator should be rejected")
	}
	if !strings.Contains(err.Error(), "shutdown_command") {
		t.Fatalf("error should name shutdown_command; got: %s", err)
	}
}

func TestBuilderPrepare_ExplicitWinRMIsAccepted(t *testing.T) {
	var b Builder
	config := testConfig()
	config["communicator"] = "winrm"
	config["winrm_username"] = "Administrator"
	config["winrm_password"] = "secret"

	_, _, err := b.Prepare(config)

	if err != nil {
		t.Fatalf("a complete WinRM template should be accepted: %s", err)
	}
}

func TestBuilderPrepare_WinRMNeedsNoSSHUsername(t *testing.T) {
	var b Builder
	config := testConfig()
	config["communicator"] = "winrm"
	config["winrm_username"] = "Administrator"
	config["winrm_password"] = "secret"
	delete(config, "ssh_username")

	_, _, err := b.Prepare(config)

	if err != nil {
		t.Fatalf("a WinRM install with winrm credentials and no ssh_username should be accepted: %s", err)
	}
}

func TestBuilderPrepare_CommunicatorNoneNeedsNoSSHUsername(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "ssh_username")
	delete(config, "shutdown_command")
	config["communicator"] = "none"

	_, _, err := b.Prepare(config)

	if err != nil {
		t.Fatalf("communicator = \"none\" without an ssh_username should be accepted: %s", err)
	}
}

func TestBuilderPrepare_VTPMWithoutUEFIIsRejected(t *testing.T) {
	for _, firmware := range []string{"", "bios"} {
		var b Builder
		config := testConfig()
		config["vtpm"] = true
		if firmware != "" {
			config["firmware"] = firmware
		}

		_, _, err := b.Prepare(config)

		if err == nil {
			t.Errorf("firmware=%q: vtpm without UEFI must be rejected", firmware)
			continue
		}
		if !strings.Contains(err.Error(), "vtpm") || !strings.Contains(err.Error(), "uefi") {
			t.Errorf("firmware=%q: error should name vtpm and uefi, got: %s", firmware, err)
		}
	}
}

func TestBuilderPrepare_FirmwareCaseIsNormalised(t *testing.T) {
	var b Builder
	config := testConfig()
	config["firmware"] = "UEFI"
	config["secure_boot"] = true

	_, _, err := b.Prepare(config)

	if err != nil {
		t.Fatalf("firmware = \"UEFI\" with secure_boot must be accepted (UEFI is UEFI); got: %s", err)
	}
	// Accepting it is not enough: XAPI takes only the lower-case spelling.
	if b.config.Firmware != "uefi" {
		t.Errorf("firmware = %q, want it normalised to \"uefi\"", b.config.Firmware)
	}
}

func TestBuilderPrepare_FirmwareEnumIsValidated(t *testing.T) {
	var b Builder
	config := testConfig()
	config["firmware"] = "efi" // neither bios nor uefi

	_, _, err := b.Prepare(config)

	if err == nil {
		t.Fatal("firmware = \"efi\" must be rejected: firmware must be one of bios/uefi")
	}
	if !strings.Contains(err.Error(), "firmware") {
		t.Errorf("error should name firmware, got: %s", err)
	}
}

func TestBuilderPrepare_CorePerSocketIsNotDefaultedFromVCPUs(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vcpus"] = 2

	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	xstest.WantUint(t, "vcpus", b.config.VCPUs, 2)
	xstest.WantUnset(t, "cores_per_socket", b.config.CorePerSocket)
}

func TestBuilderPrepare_VMDescription(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vm_description"] = "built by packer"

	_, _, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.VMDescription != "built by packer" {
		t.Errorf("vm_description = %q, want \"built by packer\"", b.config.VMDescription)
	}
}

func TestBuilderPrepare_CorePerSocketExceedingVCPUsIsRejected(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vcpus"] = 4
	config["cores_per_socket"] = 8

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("cores_per_socket greater than vcpus is an impossible topology and must be rejected at validate time")
	}
	if !strings.Contains(err.Error(), "cores_per_socket") {
		t.Errorf("error should name cores_per_socket, got: %s", err)
	}
}

func TestBuilderPrepare_TooManySocketsIsRejected(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vcpus"] = 64
	config["cores_per_socket"] = 1

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("64 sockets is beyond what XenServer supports and must be rejected at validate time")
	}
	if !strings.Contains(err.Error(), "sockets") {
		t.Errorf("error should say how many sockets the pair gives, got: %s", err)
	}
}

func TestBuilderPrepare_UnsetFirmwareIsLeftToTheTemplate(t *testing.T) {
	var b Builder
	config := testConfig()

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.Firmware != "" {
		t.Errorf("firmware = %q, want it left empty so the template's value is kept", b.config.Firmware)
	}
}

func TestBuilderPrepare_ExplicitSecureBootFalseNeedsNoUEFI(t *testing.T) {
	var b Builder
	config := testConfig()
	config["secure_boot"] = false

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("secure_boot = false with no firmware set should be accepted: %s", err)
	}
	if b.config.SecureBoot == nil {
		t.Fatal("secure_boot = false never reached the field")
	}
	if b.config.SecureBootEnabled() {
		t.Error("secure_boot = false must not report Secure Boot enabled")
	}
}

func TestBuilderPrepare_CloneTemplateIsRequired(t *testing.T) {
	for _, tc := range []struct {
		name  string
		apply func(map[string]interface{})
	}{
		{"absent", func(c map[string]interface{}) { delete(c, "clone_template") }},
		{"empty", func(c map[string]interface{}) { c["clone_template"] = "" }},
		{"whitespace", func(c map[string]interface{}) { c["clone_template"] = " " }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var b Builder
			config := testConfig()
			tc.apply(config)

			_, _, err := b.Prepare(config)
			if err == nil {
				t.Fatal("a build with no clone_template should not prepare")
			}
			if !strings.Contains(err.Error(), "clone_template") {
				t.Errorf("error %q does not name clone_template", err)
			}
			if !strings.Contains(err.Error(), xscommon.TemplateLookupHint) {
				t.Errorf("error %q does not say how to find a template", err)
			}
		})
	}
}

func TestBuilderPrepare_MissingCloneTemplateDoesNotWarnAboutAnEmptyName(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "clone_template")
	delete(config, "vcpus")
	config["cores_per_socket"] = 3

	_, warns, err := b.Prepare(config)
	if err == nil {
		t.Fatal("a build with no clone_template should not prepare")
	}
	for _, w := range warns {
		if strings.Contains(w, "clone_template ''") {
			t.Errorf("warning quotes an empty clone_template: %s", w)
		}
	}
}

func TestBuilderPrepare_CloneTemplatePassthrough(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = "Windows 11"

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if b.config.CloneTemplate != "Windows 11" {
		t.Errorf("clone_template = %q, want \"Windows 11\"", b.config.CloneTemplate)
	}
}

func TestBuilderPrepare_ISONameAndISOUrlsAreExclusive(t *testing.T) {
	var b Builder
	config := testConfig()
	delete(config, "iso_url")
	config["iso_name"] = "Win11.iso"
	config["iso_urls"] = []string{
		"http://mirror.example.com/test.iso",
		"http://backup.example.com/test.iso",
	}

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("iso_name together with iso_urls must be rejected")
	}
	if !strings.Contains(err.Error(), "only one of iso_name or iso_url") {
		t.Errorf("error should say the two are exclusive, got: %s", err)
	}
}

func TestBuilderPrepare_NetworkNamesDecodeInOrder(t *testing.T) {
	var b Builder
	config := testConfig()
	config["network_names"] = []string{"management", "storage", "guest"}

	if _, _, err := b.Prepare(config); err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	want := []string{"management", "storage", "guest"}
	if !reflect.DeepEqual(b.config.NetworkNames, want) {
		t.Errorf("network_names = %#v, want %#v in order", b.config.NetworkNames, want)
	}
}

func TestFinalStepsSnapshotsBeforeConverting(t *testing.T) {
	var order []string
	for _, step := range finalSteps() {
		order = append(order, fmt.Sprintf("%T", step))
	}

	want := []string{
		"*common.StepShutdown",
		"*common.StepDetachVdi",
		"*common.StepDetachVdi",
		"*common.StepDetachVdi",
		"*common.StepDetachVdi",
		"*common.StepCreateSnapshot",
		"*common.StepSetVmToTemplate",
		"*common.StepExport",
	}
	if !reflect.DeepEqual(order, want) {
		t.Errorf("final steps = %v, want %v", order, want)
	}
}

func TestBuilderPrepare_ExplicitZeroSizingIsRejected(t *testing.T) {
	for _, tc := range []struct {
		option string
		value  interface{}
	}{
		{"vm_name", ""},
		{"vcpus", 0},
		{"cores_per_socket", 0},
		{"vm_memory_gb", 0},
		{"disk_size_gb", 0},
	} {
		t.Run(tc.option, func(t *testing.T) {
			var b Builder
			config := testConfig()
			config[tc.option] = tc.value

			_, _, err := b.Prepare(config)
			if err == nil {
				t.Fatalf("%s = %#v must be rejected, not replaced by a default", tc.option, tc.value)
			}
			if !strings.Contains(err.Error(), tc.option) {
				t.Errorf("error should name %s, got: %s", tc.option, err)
			}
		})
	}
}

func TestBuilderPrepare_NamedCloneTemplateKeepsItsSizing(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = "Windows 11"

	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	xstest.WantUnset(t, "vcpus", b.config.VCPUs)
	xstest.WantUnset(t, "cores_per_socket", b.config.CorePerSocket)
	xstest.WantUnset(t, "vm_memory_gb", b.config.VMMemoryGB)

	xstest.WantUnset(t, "disk_size_gb", b.config.DiskSizeGB)
}

func TestBuilderPrepare_GenericTemplateRequiresItsOwnSizing(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = genericTemplate

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("the generic template supplies no sizing, so a build naming it must set its own")
	}
	for _, option := range []string{"vcpus", "vm_memory_gb", "disk_size_gb"} {
		if !strings.Contains(err.Error(), option) {
			t.Errorf("error should name %s, got: %s", option, err)
		}
	}
	if strings.Contains(err.Error(), "cores_per_socket") {
		t.Errorf("cores_per_socket is not required on the generic template, got: %s", err)
	}
}

func TestBuilderPrepare_GenericTemplateWithFullSizingIsAccepted(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = genericTemplate
	config["vcpus"] = 2
	config["vm_memory_gb"] = 4
	config["disk_size_gb"] = 64

	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	xstest.WantUint(t, "vcpus", b.config.VCPUs, 2)
	xstest.WantUint(t, "vm_memory_gb", b.config.VMMemoryGB, 4)
	xstest.WantUint(t, "disk_size_gb", b.config.DiskSizeGB, 64)
}

func TestBuilderPrepare_NamedCloneTemplateStillTakesExplicitSizing(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = "Windows 11"
	config["vcpus"] = 4
	config["disk_size_gb"] = 100

	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	xstest.WantUint(t, "vcpus", b.config.VCPUs, 4)
	xstest.WantUint(t, "disk_size_gb", b.config.DiskSizeGB, 100)
	xstest.WantUnset(t, "cores_per_socket", b.config.CorePerSocket)
	xstest.WantUnset(t, "vm_memory_gb", b.config.VMMemoryGB)
}

func TestBuilderPrepare_CorePerSocketWithoutVCPUsIsKeptFromTheCloneTemplate(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = "Windows 11"
	config["cores_per_socket"] = 2

	_, warns, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	xstest.WantUnset(t, "vcpus", b.config.VCPUs)
	if len(warns) != 1 || !strings.Contains(warns[0], "cores_per_socket") {
		t.Errorf("cores_per_socket without vcpus should warn that it is ignored, got: %#v", warns)
	}
}

func TestKeptSizing(t *testing.T) {
	for _, tc := range []struct {
		name   string
		config ISOConfig
		want   []string
	}{
		{"nothing set", ISOConfig{}, []string{"vCPU count", "memory"}},
		{
			"everything set",
			sizedConfig(xscommon.Uint(2), xscommon.Uint(4)),
			nil,
		},
		{
			"vcpus alone",
			sizedConfig(xscommon.Uint(2), nil),
			[]string{"memory"},
		},
		{
			"memory alone",
			sizedConfig(nil, xscommon.Uint(4)),
			[]string{"vCPU count"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := keptSizing(tc.config); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("keptSizing() = %#v, want %#v", got, tc.want)
			}
		})
	}
}

func sizedConfig(vcpus, memoryGB *uint) ISOConfig {
	var c ISOConfig
	c.VCPUs = vcpus
	c.VMMemoryGB = memoryGB
	return c
}

func TestBuilderPrepare_InheritedSizingIsSilent(t *testing.T) {
	var b Builder
	config := testConfig()
	config["clone_template"] = "Windows 11"

	_, warns, err := b.Prepare(config)
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
	if len(warns) > 0 {
		t.Errorf("inherited sizing is not the build's to warn about, got: %#v", warns)
	}
}

func TestBuilderPrepare_MemoryBelowTheWindowsMinimumWarns(t *testing.T) {
	for _, tc := range []struct {
		memoryGB int
		wantWarn bool
	}{
		{1, true},
		{2, false},
		{4, false},
	} {
		t.Run(fmt.Sprintf("%dGB", tc.memoryGB), func(t *testing.T) {
			var b Builder
			config := testConfig()
			config["vm_memory_gb"] = tc.memoryGB

			_, warns, err := b.Prepare(config)
			if err != nil {
				t.Fatalf("should not have error: %s", err)
			}

			var warned bool
			for _, w := range warns {
				if strings.Contains(w, "vm_memory_gb") {
					warned = true
				}
			}
			if warned != tc.wantWarn {
				t.Errorf("vm_memory_gb = %d warned = %v, want %v (warnings: %#v)",
					tc.memoryGB, warned, tc.wantWarn, warns)
			}
		})
	}
}

func TestBuilderPrepare_ZeroMemoryDoesNotAlsoWarn(t *testing.T) {
	var b Builder
	config := testConfig()
	config["vm_memory_gb"] = 0

	_, warns, err := b.Prepare(config)
	if err == nil {
		t.Fatal("vm_memory_gb = 0 must be rejected")
	}

	for _, w := range warns {
		if strings.Contains(w, "vm_memory_gb") {
			t.Errorf("vm_memory_gb = 0 should not also warn, got: %s", w)
		}
	}
}

func TestMissingGenericSizing(t *testing.T) {
	for _, tc := range []struct {
		name   string
		config ISOConfig
		want   []string
	}{
		{"named template, nothing set", templateConfig("Windows 11", nil, nil, nil), nil},
		{
			"generic template, nothing set",
			templateConfig(genericTemplate, nil, nil, nil),
			[]string{"vcpus", "vm_memory_gb", "disk_size_gb"},
		},
		{
			"generic template, everything set",
			templateConfig(genericTemplate, xscommon.Uint(2), xscommon.Uint(4), xscommon.Uint(64)),
			nil,
		},
		{
			"generic template, memory alone",
			templateConfig(genericTemplate, nil, xscommon.Uint(4), nil),
			[]string{"vcpus", "disk_size_gb"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := missingGenericSizing(tc.config); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("missingGenericSizing() = %#v, want %#v", got, tc.want)
			}
		})
	}
}

func templateConfig(template string, vcpus, memoryGB, diskGB *uint) ISOConfig {
	c := sizedConfig(vcpus, memoryGB)
	c.CloneTemplate = template
	c.DiskSizeGB = diskGB
	return c
}

func TestBootCommandStep_SweepsTheUEFICDPrompt(t *testing.T) {
	if !bootCommandStep(interpolate.Context{}).SweepUEFICDPrompt {
		t.Error("the iso builder must ask for the UEFI CD boot prompt to be swept")
	}
}
