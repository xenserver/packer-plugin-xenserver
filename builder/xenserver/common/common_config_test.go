// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

func boolPtr(v bool) *bool { return &v }

var connectionEnvVars = []string{EnvHost, EnvUsername, EnvPassword, EnvCACertPath, EnvInsecure}

func minimalConfig(format string) *CommonConfig {
	return &CommonConfig{
		Username: "root",
		Password: "secret",
		HostIp:   "10.0.0.1",
		Format:   format,
		Insecure: boolPtr(true), // otherwise Prepare also demands ca_cert_path
	}
}

func prepare(t *testing.T, c *CommonConfig) ([]string, []error) {
	t.Helper()
	return prepareWithEnv(t, c, nil)
}

func prepareWithEnv(t *testing.T, c *CommonConfig, env map[string]string) ([]string, []error) {
	t.Helper()
	for _, key := range connectionEnvVars {
		t.Setenv(key, env[key])
	}
	return c.Prepare(&common.PackerConfig{PackerBuildName: "test"})
}

func errorsContain(errs []error, substring string) bool {
	for _, err := range errs {
		if strings.Contains(err.Error(), substring) {
			return true
		}
	}
	return false
}

func TestCommonConfigPrepare_NormalizesXvaCompressedToXva(t *testing.T) {
	c := minimalConfig("xva_compressed")

	warns, errs := prepare(t, c)

	if len(errs) > 0 {
		t.Fatalf("xva_compressed must remain an accepted input: %v", errs)
	}
	if c.Format != "xva" {
		t.Errorf("Format = %q, want it normalized to \"xva\"", c.Format)
	}

	var found bool
	for _, w := range warns {
		if strings.Contains(w, "xva_compressed is deprecated") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a deprecation warning for xva_compressed, got %v", warns)
	}
}

func TestCommonConfigPrepare_FormatDefaultsToXva(t *testing.T) {
	c := minimalConfig("")

	warns, errs := prepare(t, c)

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.Format != "xva" {
		t.Errorf("Format = %q, want \"xva\"", c.Format)
	}
	if len(warns) > 0 {
		t.Errorf("the default must not warn, got %v", warns)
	}
}

func TestCommonConfigPrepare_OtherFormatsAreUnchanged(t *testing.T) {
	for _, format := range []string{"xva", "vdi_raw", "vdi_vhd", "none"} {
		c := minimalConfig(format)
		// format = "none" is only legal with a keep_vm that keeps the template.
		if format == "none" {
			c.KeepVM = "on_success"
		}

		warns, errs := prepare(t, c)

		if len(errs) > 0 {
			t.Errorf("format %q: unexpected errors: %v", format, errs)
			continue
		}
		if c.Format != format {
			t.Errorf("format %q was rewritten to %q", format, c.Format)
		}
		if len(warns) > 0 {
			t.Errorf("format %q must not warn, got %v", format, warns)
		}
	}
}

func TestCommonConfigPrepare_RejectsUnknownFormat(t *testing.T) {
	c := minimalConfig("tarball")

	if _, errs := prepare(t, c); len(errs) == 0 {
		t.Error("an unrecognised format must be rejected")
	}
}

func TestCommonConfigPrepare_RejectsFormatSpellingVariants(t *testing.T) {
	for _, format := range []string{"XVA", "Xva", " xva", "xva ", " xva ", "VDI_RAW"} {
		c := minimalConfig(format)

		_, errs := prepare(t, c)

		if !errorsContain(errs, "format") {
			t.Errorf("format = %q must be rejected by name, got %v", format, errs)
		}
	}
}

func TestCommonConfigPrepare_RejectsKeepVMSpellingVariants(t *testing.T) {
	for _, keepVM := range []string{"Never", "NEVER", "on-success", "On_Success", " never ", "always "} {
		c := minimalConfig("xva")
		c.KeepVM = keepVM

		_, errs := prepare(t, c)

		if !errorsContain(errs, "keep_vm") {
			t.Errorf("keep_vm = %q must be rejected by name, got %v", keepVM, errs)
		}
	}
}

func TestCommonConfigPrepare_RejectsBindAddressWithInterface(t *testing.T) {
	c := minimalConfig("xva")
	c.HTTPAddress = "10.0.0.5"
	c.HTTPInterface = "eth0"

	_, errs := prepare(t, c)

	if len(errs) == 0 {
		t.Fatal("http_bind_address and http_interface must not be accepted together")
	}
}

func TestCommonConfigPrepare_RejectsANonAddressBindAddress(t *testing.T) {
	c := minimalConfig("xva")
	c.HTTPAddress = "http://10.0.0.5"

	if _, errs := prepare(t, c); len(errs) == 0 {
		t.Error("http_bind_address must be rejected when it is not an IP address")
	}
}

func TestCommonConfigPrepare_AcceptsAnAddressBindAddress(t *testing.T) {
	for _, addr := range []string{"10.0.0.5", "0.0.0.0", "::"} {
		c := minimalConfig("xva")
		c.HTTPAddress = addr

		if _, errs := prepare(t, c); len(errs) > 0 {
			t.Errorf("http_bind_address %q should be accepted: %v", addr, errs)
		}
	}
}

func TestCommonConfig_HasVendorDeviceDefaultsToOn(t *testing.T) {
	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if !c.HasVendorDeviceEnabled() {
		t.Error("has_vendor_device defaulted to off; Windows guests need it on")
	}
}

func TestCommonConfig_HasVendorDeviceHonoursAnExplicitValue(t *testing.T) {
	for _, want := range []bool{true, false} {
		c := minimalConfig("xva")
		c.HasVendorDevice = &want

		if _, errs := prepare(t, c); len(errs) > 0 {
			t.Errorf("has_vendor_device=%t: unexpected errors: %v", want, errs)
			continue
		}
		if got := c.HasVendorDeviceEnabled(); got != want {
			t.Errorf("HasVendorDeviceEnabled() = %t, want %t", got, want)
		}
	}
}

func TestCommonConfigPrepare_AcceptsCDFilesWithoutAnISOSR(t *testing.T) {
	c := minimalConfig("xva")
	c.CDFiles = []string{filepath.Join(t.TempDir(), "Autounattend.xml")}
	if err := os.WriteFile(c.CDFiles[0], []byte("<unattend/>"), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Errorf("unexpected errors: %v", errs)
	}
}

func TestCommonConfigPrepare_AllowsCDFilesWithAnISOSR(t *testing.T) {
	c := minimalConfig("xva")
	c.ISOSR = "ISO library"
	c.CDFiles = []string{filepath.Join(t.TempDir(), "Autounattend.xml")}
	if err := os.WriteFile(c.CDFiles[0], []byte("<unattend/>"), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Errorf("unexpected errors: %v", errs)
	}
}

func TestCommonConfigPrepare_RejectsShutdownCommandWithoutACommunicator(t *testing.T) {
	c := minimalConfig("xva")
	c.ShutdownCommand = "shutdown -h now"
	c.Comm.Type = "none"

	_, errs := prepare(t, c)

	if len(errs) == 0 {
		t.Fatal("shutdown_command with communicator = \"none\" must be rejected")
	}
}

func TestCommonConfigPrepare_AllowsShutdownCommandWithACommunicator(t *testing.T) {
	for _, commType := range []string{"ssh", "winrm"} {
		c := minimalConfig("xva")
		c.ShutdownCommand = "shutdown -h now"
		c.Comm.Type = commType

		if _, errs := prepare(t, c); len(errs) > 0 {
			t.Errorf("communicator %q: shutdown_command should be accepted: %v", commType, errs)
		}
	}
}

func TestCommonConfigPrepare_RejectsNoneFormatWithKeepVMNever(t *testing.T) {
	c := minimalConfig("none")
	c.KeepVM = "never"

	_, errs := prepare(t, c)

	if len(errs) == 0 {
		t.Fatal("format = \"none\" with keep_vm = \"never\" produces no output and must be rejected")
	}

	var found bool
	for _, err := range errs {
		if strings.Contains(err.Error(), "produces no output") {
			found = true
		}
	}
	if !found {
		t.Errorf("the error should explain that the build produces nothing, got: %v", errs)
	}
}

func TestCommonConfigPrepare_NoneFormatDefaultsKeepVMToOnSuccess(t *testing.T) {
	c := minimalConfig("none")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("format = \"none\" on its own must be accepted: %v", errs)
	}
	if c.KeepVM != "on_success" {
		t.Errorf("keep_vm = %q, want \"on_success\" for format = \"none\"", c.KeepVM)
	}
}

func TestCommonConfigPrepare_AllowsNoneFormatWhenTheTemplateIsKept(t *testing.T) {
	for _, keep := range []string{"always", "on_success"} {
		c := minimalConfig("none")
		c.KeepVM = keep

		if _, errs := prepare(t, c); len(errs) > 0 {
			t.Errorf("format = \"none\" with keep_vm = %q should be accepted: %v", keep, errs)
		}
	}
}

func TestCommonConfigPrepare_AllowsKeepVMNeverForExportingFormats(t *testing.T) {
	for _, format := range []string{"xva", "vdi_raw", "vdi_vhd"} {
		c := minimalConfig(format)
		c.KeepVM = "never"

		if _, errs := prepare(t, c); len(errs) > 0 {
			t.Errorf("format %q with keep_vm = \"never\" should be accepted: %v", format, errs)
		}
	}
}

func TestCommonConfigPrepare_WarnsThatACommunicatorlessBuildRunsNoProvisioner(t *testing.T) {
	c := minimalConfig("xva")
	c.Comm.Type = "none"

	warns, errs := prepare(t, c)

	if len(errs) > 0 {
		t.Fatalf("communicator = \"none\" on its own is valid: %v", errs)
	}
	if !anyContains(warns, "never logs in to the guest") || !anyContains(warns, "no provisioner runs") {
		t.Errorf("expected a warning that nothing logs in and no provisioner runs, got %v", warns)
	}
	if anyContains(warns, "install_timeout") {
		t.Errorf("the consequence is the builder's to state, not this one's: %v", warns)
	}
}

func TestCommonConfigPrepare_DoesNotWarnAboutAMissingCommunicator(t *testing.T) {
	for _, commType := range []string{"ssh", "winrm"} {
		c := minimalConfig("xva")
		c.Comm.Type = commType

		warns, _ := prepare(t, c)

		if anyContains(warns, "no provisioner runs") {
			t.Errorf("communicator %q: unexpected communicator-less warning: %v", commType, warns)
		}
	}
}

func TestCommonConfigPrepare_ConnectionOptionsFallBackToTheEnvironment(t *testing.T) {
	caCertPath := writeTestCACert(t)
	c := &CommonConfig{}

	_, errs := prepareWithEnv(t, c, map[string]string{
		EnvHost:       "pool.example.com",
		EnvUsername:   "root",
		EnvPassword:   "from-the-environment",
		EnvCACertPath: caCertPath,
	})

	if len(errs) > 0 {
		t.Fatalf("a config supplied entirely from the environment must be accepted: %v", errs)
	}
	if c.HostIp != "pool.example.com" {
		t.Errorf("HostIp = %q, want it from %s", c.HostIp, EnvHost)
	}
	if c.Username != "root" {
		t.Errorf("Username = %q, want it from %s", c.Username, EnvUsername)
	}
	if c.Password != "from-the-environment" {
		t.Errorf("Password = %q, want it from %s", c.Password, EnvPassword)
	}
	if c.CACertPath != caCertPath {
		t.Errorf("CACertPath = %q, want it from %s", c.CACertPath, EnvCACertPath)
	}
}

func TestCommonConfigPrepare_TemplateValuesWinOverTheEnvironment(t *testing.T) {
	c := minimalConfig("xva")

	_, errs := prepareWithEnv(t, c, map[string]string{
		EnvHost:     "wrong.example.com",
		EnvUsername: "wrong",
		EnvPassword: "wrong",
	})

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.HostIp != "10.0.0.1" || c.Username != "root" || c.Password != "secret" {
		t.Errorf("the environment overrode the template: host=%q user=%q password=%q",
			c.HostIp, c.Username, c.Password)
	}
}

func TestCommonConfigPrepare_ExplicitInsecureFalseSurvivesTheEnvironment(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = boolPtr(false)
	c.CACertPath = writeTestCACert(t)

	_, errs := prepareWithEnv(t, c, map[string]string{EnvInsecure: "true"})

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.InsecureEnabled() {
		t.Error("XENSERVER_INSECURE overrode an explicit insecure = false")
	}
}

func TestCommonConfigPrepare_InsecureFallsBackToTheEnvironment(t *testing.T) {
	c := &CommonConfig{Username: "root", Password: "secret", HostIp: "10.0.0.1"}

	_, errs := prepareWithEnv(t, c, map[string]string{EnvInsecure: "1"})

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if !c.InsecureEnabled() {
		t.Error("XENSERVER_INSECURE=1 did not turn verification off")
	}
}

func TestCommonConfigPrepare_EnvironmentInsecureDoesNotDefeatAPinnedCertificate(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = nil
	c.CACertPath = writeTestCACert(t)

	warns, errs := prepareWithEnv(t, c, map[string]string{EnvInsecure: "true"})

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.InsecureEnabled() {
		t.Error("XENSERVER_INSECURE turned off verification for a build that pinned a certificate")
	}
	if !anyContains(warns, EnvInsecure) {
		t.Errorf("ignoring %s must be said out loud, got %v", EnvInsecure, warns)
	}
}

func TestCommonConfigPrepare_EnvironmentInsecureAppliesWithNoCertificatePinned(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = nil

	warns, errs := prepareWithEnv(t, c, map[string]string{EnvInsecure: "true"})

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if !c.InsecureEnabled() {
		t.Error("XENSERVER_INSECURE=true did not turn verification off")
	}
	if !anyContains(warns, EnvInsecure) {
		t.Errorf("turning verification off from the environment must be said out loud, got %v", warns)
	}
}

func TestCommonConfigPrepare_EnvironmentInsecureFalseDoesNotWarn(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = nil
	c.CACertPath = writeTestCACert(t)

	warns, errs := prepareWithEnv(t, c, map[string]string{EnvInsecure: "false"})

	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if anyContains(warns, EnvInsecure) {
		t.Errorf("XENSERVER_INSECURE=false keeps verification on and needs no warning, got %v", warns)
	}
}

func TestCommonConfigPrepare_RejectsANonBooleanInsecureEnvironmentValue(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = nil

	_, errs := prepareWithEnv(t, c, map[string]string{EnvInsecure: "yes-please"})

	if !errorsContain(errs, EnvInsecure) {
		t.Errorf("a non-boolean %s must be rejected by name, got %v", EnvInsecure, errs)
	}
}

func TestCommonConfigPrepare_RejectsAnInvalidHostFromTheEnvironment(t *testing.T) {
	c := minimalConfig("xva")
	c.HostIp = ""

	_, errs := prepareWithEnv(t, c, map[string]string{EnvHost: "http://10.0.0.1"})

	if !errorsContain(errs, "xenserver_host") {
		t.Errorf("an http:// %s must be rejected by name, got %v", EnvHost, errs)
	}
}

func TestCommonConfigPrepare_RejectsAnUnusableCACertFromTheEnvironment(t *testing.T) {
	notPEM := filepath.Join(t.TempDir(), "not-a-cert.pem")
	if err := os.WriteFile(notPEM, []byte("this is not a certificate\n"), 0600); err != nil {
		t.Fatalf("could not write the fixture: %s", err)
	}

	for name, path := range map[string]string{
		"missing":   filepath.Join(t.TempDir(), "absent.pem"),
		"not a PEM": notPEM,
	} {
		c := minimalConfig("xva")
		c.Insecure = nil
		c.CACertPath = ""

		_, errs := prepareWithEnv(t, c, map[string]string{EnvCACertPath: path})

		if !errorsContain(errs, "ca_cert_path") {
			t.Errorf("a %s %s must be rejected at prepare time, got %v", name, EnvCACertPath, errs)
		}
	}
}

func TestCommonConfigPrepare_RequiresACertificateWhenVerifying(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = nil

	_, errs := prepare(t, c)

	if !errorsContain(errs, "ca_cert_path") {
		t.Fatalf("a verifying build with no certificate must be rejected, got %v", errs)
	}
	if !errorsContain(errs, EnvCACertPath) {
		t.Errorf("the error should also name %s, got %v", EnvCACertPath, errs)
	}
}

func TestCommonConfigPrepare_AcceptsAReadableCACert(t *testing.T) {
	c := minimalConfig("xva")
	c.Insecure = nil
	c.CACertPath = writeTestCACert(t)

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Errorf("a readable PEM certificate should be accepted: %v", errs)
	}
}

func TestCommonConfigPrepare_RejectsAnUnusableCACert(t *testing.T) {
	notPEM := filepath.Join(t.TempDir(), "not-a-cert.pem")
	if err := os.WriteFile(notPEM, []byte("this is not a certificate\n"), 0600); err != nil {
		t.Fatalf("could not write the fixture: %s", err)
	}

	empty := filepath.Join(t.TempDir(), "empty.pem")
	if err := os.WriteFile(empty, nil, 0600); err != nil {
		t.Fatalf("could not write the fixture: %s", err)
	}

	for name, path := range map[string]string{
		"missing":       filepath.Join(t.TempDir(), "absent.pem"),
		"not a PEM":     notPEM,
		"empty":         empty,
		"a private key": writeTestPrivateKeyPEM(t),
	} {
		c := minimalConfig("xva")
		c.Insecure = nil
		c.CACertPath = path

		_, errs := prepare(t, c)

		if !errorsContain(errs, "ca_cert_path") {
			t.Errorf("a %s ca_cert_path must be rejected at prepare time, got %v", name, errs)
		}
	}
}

func anyContains(values []string, substr string) bool {
	for _, v := range values {
		if strings.Contains(v, substr) {
			return true
		}
	}
	return false
}

func TestCommonConfigPrepare_RejectsAPlaintextHost(t *testing.T) {
	for _, host := range []string{"http://10.0.0.1", "HTTP://10.0.0.1", "http://pool.example.com:80"} {
		c := minimalConfig("xva")
		c.HostIp = host

		_, errs := prepare(t, c)

		if !errorsContain(errs, "xenserver_host") {
			t.Errorf("xenserver_host = %q must be rejected, got %v", host, errs)
		}
	}
}

func TestCommonConfigPrepare_AcceptsHTTPSAndBareHosts(t *testing.T) {
	for _, host := range []string{"10.0.0.1", "pool.example.com", "https://pool.example.com", "HTTPS://10.0.0.1"} {
		c := minimalConfig("xva")
		c.HostIp = host

		if _, errs := prepare(t, c); len(errs) > 0 {
			t.Errorf("xenserver_host = %q should be accepted: %v", host, errs)
		}
	}
}

func TestCommonConfigPrepare_RejectsANonHTTPScheme(t *testing.T) {
	c := minimalConfig("xva")
	c.HostIp = "ssh://10.0.0.1"

	if _, errs := prepare(t, c); !errorsContain(errs, "xenserver_host") {
		t.Errorf("an unrecognised scheme must be rejected, got %v", errs)
	}
}

func TestCommonConfigPrepare_DefaultVMNameHasARenderedTimestamp(t *testing.T) {
	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("should not have errors: %v", errs)
	}

	if strings.Contains(c.Name(), "{{") {
		t.Fatalf("vm_name = %q still carries an un-rendered template", c.Name())
	}
	want := fmt.Sprintf("packer-test-%d", interpolate.InitTime.Unix())
	if c.Name() != want {
		t.Errorf("vm_name = %q, want %q", c.Name(), want)
	}
}

func TestCommonConfigPrepare_AnExplicitVMNameIsLeftAlone(t *testing.T) {
	c := minimalConfig("xva")
	name := "my-golden-image"
	c.VMName = &name

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("should not have errors: %v", errs)
	}
	if c.Name() != "my-golden-image" {
		t.Errorf("vm_name = %q, want it untouched", c.Name())
	}
}

func TestCommonConfigPrepare_RejectsANegativeBootKeyGroups(t *testing.T) {
	c := minimalConfig("xva")
	c.BootKeyGroups = -1

	_, errs := prepare(t, c)

	if len(errs) == 0 {
		t.Fatal("boot_key_groups = -1 must be rejected at validate time")
	}
}

func TestCommonConfigPrepare_RequiresAUsername(t *testing.T) {
	c := minimalConfig("xva")
	c.Username = ""

	_, errs := prepare(t, c)

	if !errorsContain(errs, "xenserver_username") {
		t.Fatalf("a build with no xenserver_username must be rejected, got %v", errs)
	}
	if !errorsContain(errs, EnvUsername) {
		t.Errorf("the error should also name %s, got %v", EnvUsername, errs)
	}
}

func TestCommonConfigPrepare_RequiresAPassword(t *testing.T) {
	c := minimalConfig("xva")
	c.Password = ""

	_, errs := prepare(t, c)

	if !errorsContain(errs, "xenserver_password") {
		t.Fatalf("a build with no xenserver_password must be rejected, got %v", errs)
	}
	if !errorsContain(errs, EnvPassword) {
		t.Errorf("the error should also name %s, got %v", EnvPassword, errs)
	}
}

func TestCommonConfigPrepare_RequiresAHost(t *testing.T) {
	c := minimalConfig("xva")
	c.HostIp = ""

	_, errs := prepare(t, c)

	if !errorsContain(errs, "xenserver_host") {
		t.Fatalf("a build with no xenserver_host must be rejected, got %v", errs)
	}
	if !errorsContain(errs, EnvHost) {
		t.Errorf("the error should also name %s, got %v", EnvHost, errs)
	}
}

func TestCommonConfigPrepare_InsecureSkipsCACertValidation(t *testing.T) {
	c := minimalConfig("xva") // minimalConfig already sets insecure = true
	c.CACertPath = filepath.Join(t.TempDir(), "does-not-exist.pem")

	_, errs := prepare(t, c)

	if len(errs) > 0 {
		t.Fatalf("insecure = true must not require or validate ca_cert_path: %v", errs)
	}
	if !c.InsecureEnabled() {
		t.Error("insecure = true was not honoured")
	}
}

func TestCommonConfigPrepare_BootWaitDefaultsToFiveSeconds(t *testing.T) {
	const fiveSeconds = 5 * 1000 * 1000 * 1000

	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.BootWait != fiveSeconds {
		t.Errorf("boot_wait = %d, want %d (5s)", c.BootWait, fiveSeconds)
	}
}

func TestCommonConfigPrepare_InstallTimeoutDefaultsTo200Minutes(t *testing.T) {
	const twoHundredMinutes = 200 * 60 * 1000 * 1000 * 1000

	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.InstallTimeout != twoHundredMinutes {
		t.Errorf("install_timeout = %s, want 200m", c.InstallTimeout)
	}
}

func TestCommonConfigPrepare_BootKeyGroupIntervalIsPreserved(t *testing.T) {
	const interval = 1500 * 1000 * 1000 // 1.5s

	c := minimalConfig("xva")
	c.BootKeyGroupInterval = interval

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.BootKeyGroupInterval != interval {
		t.Errorf("boot_key_group_interval = %s, want 1.5s; Prepare must leave it untouched", c.BootKeyGroupInterval)
	}
}

func TestCommonConfigPrepare_HTTPDirectoryIsPreserved(t *testing.T) {
	c := minimalConfig("xva")
	c.HTTPDir = "http"

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.HTTPDir != "http" {
		t.Errorf("http_directory = %q, want it left as %q", c.HTTPDir, "http")
	}
}

func TestCommonConfigPrepare_AcceptsHTTPInterfaceAlone(t *testing.T) {
	c := minimalConfig("xva")
	c.HTTPInterface = "eth0"

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Errorf("http_interface on its own must be accepted: %v", errs)
	}
	if c.HTTPInterface != "eth0" {
		t.Errorf("http_interface = %q, want it preserved", c.HTTPInterface)
	}
}

func TestCommonConfigPrepare_RejectsUnknownKeepVM(t *testing.T) {
	c := minimalConfig("xva")
	c.KeepVM = "sometimes"

	_, errs := prepare(t, c)

	if !errorsContain(errs, "keep_vm must be one of") {
		t.Errorf("an unrecognised keep_vm must be rejected, got %v", errs)
	}
}

func TestCommonConfigPrepare_KeepVMDefaultsToNever(t *testing.T) {
	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.KeepVM != "never" {
		t.Errorf("keep_vm = %q, want the default \"never\"", c.KeepVM)
	}
}

func TestCommonConfigPrepare_OutputDirDefaultsToOutputBuildName(t *testing.T) {
	c := minimalConfig("xva")

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.OutputDir != "output-test" {
		t.Errorf("output_directory = %q, want the default \"output-test\" (output-<build_name>)", c.OutputDir)
	}
}

func TestCommonConfigPrepare_OutputDirIsPreservedWhenSet(t *testing.T) {
	c := minimalConfig("xva")
	c.OutputDir = "my-artifacts"

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if c.OutputDir != "my-artifacts" {
		t.Errorf("output_directory = %q, want it left as the value set (\"my-artifacts\")", c.OutputDir)
	}
}

func TestCommonConfigPrepare_CreateSnapshotWithSnapshotNameIsAccepted(t *testing.T) {
	c := minimalConfig("xva")
	c.CreateSnapshot = true
	c.SnapshotName = "pre-sysprep"

	if _, errs := prepare(t, c); len(errs) > 0 {
		t.Fatalf("create_snapshot with a snapshot_name must be accepted: %v", errs)
	}
	if !c.CreateSnapshot {
		t.Error("create_snapshot was cleared by Prepare")
	}
	if c.SnapshotName != "pre-sysprep" {
		t.Errorf("snapshot_name = %q, want it preserved as \"pre-sysprep\"", c.SnapshotName)
	}
}

func TestCommonConfigPrepare_RejectsSnapshotNameWithoutCreateSnapshot(t *testing.T) {
	c := minimalConfig("xva")
	c.CreateSnapshot = false
	c.SnapshotName = "pre-sysprep"

	_, errs := prepare(t, c)

	if !errorsContain(errs, "snapshot_name") {
		t.Errorf("snapshot_name set without create_snapshot should be rejected, got %v", errs)
	}
}

func TestCommonConfigPrepare_RejectsANonExistentCDFile(t *testing.T) {
	c := minimalConfig("xva")
	c.CDFiles = []string{filepath.Join(t.TempDir(), "Autounattend.xml")}

	_, errs := prepare(t, c)

	if !errorsContain(errs, "does not exist") {
		t.Errorf("a cd_files entry that does not exist should be rejected, got %v", errs)
	}
}

func TestCommonConfigDecode_TellsAnEmptyNetworkNamesFromAnAbsentOne(t *testing.T) {
	var absent, empty CommonConfig

	if err := config.Decode(&absent, nil, map[string]interface{}{}); err != nil {
		t.Fatal(err)
	}
	if absent.NetworkNames != nil {
		t.Errorf("an absent network_names decoded to %#v, want nil", absent.NetworkNames)
	}

	if err := config.Decode(&empty, nil, map[string]interface{}{"network_names": []interface{}{}}); err != nil {
		t.Fatal(err)
	}
	if empty.NetworkNames == nil || len(empty.NetworkNames) != 0 {
		t.Errorf("an explicit network_names = [] decoded to %#v, want a non-nil empty slice", empty.NetworkNames)
	}
}

func TestCommonConfigPrepare_RejectsAnEmptyNetworkNames(t *testing.T) {
	c := minimalConfig("xva")
	c.NetworkNames = []string{}

	_, errs := prepare(t, c)

	if !errorsContain(errs, "network_names") {
		t.Errorf("network_names = [] should be rejected, got %v", errs)
	}
}

func TestCommonConfigPrepare_AcceptsAnAbsentNetworkNames(t *testing.T) {
	c := minimalConfig("xva")

	_, errs := prepare(t, c)

	if len(errs) > 0 {
		t.Errorf("an absent network_names should be accepted, got %v", errs)
	}
}

func TestCommonConfigPrepare_AcceptsANamedNetwork(t *testing.T) {
	c := minimalConfig("xva")
	c.NetworkNames = []string{"Pool-wide network associated with eth0"}

	_, errs := prepare(t, c)

	if len(errs) > 0 {
		t.Errorf("a named network should be accepted, got %v", errs)
	}
}

func TestCommonConfigPrepare_RejectsAnEmptyNetworkNamesEntry(t *testing.T) {
	for _, name := range []string{"", " "} {
		c := minimalConfig("xva")
		c.NetworkNames = []string{"Pool-wide network associated with eth0", name}

		_, errs := prepare(t, c)

		if !errorsContain(errs, "network_names[1]") {
			t.Errorf("network_names entry %q should be rejected, got %v", name, errs)
		}
	}
}
