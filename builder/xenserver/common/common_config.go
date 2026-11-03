// Copyright (c) 2026. Citrix Systems, Inc.

//go:generate packer-sdc struct-markdown
package common

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/common"
	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"

	"xenapi"
)

// BytesPerGB converts the GB-denominated size options to the bytes XAPI expects.
const BytesPerGB = 1024 * 1024 * 1024

// CommonConfig holds the options both builders share.
type CommonConfig struct {
	// The XenServer account to log in with.
	Username string `mapstructure:"xenserver_username" required:"true"`
	// The password for `xenserver_username`.
	Password string `mapstructure:"xenserver_password" required:"true"`
	// Address of the pool coordinator (master). An `https://` prefix is optional;
	// `http://` is rejected.
	HostIp string `mapstructure:"xenserver_host" required:"true"`

	// Name of the new VM. Defaults to `packer-BUILDNAME-TIMESTAMP`.
	VMName *string `mapstructure:"vm_name"`
	// Description of the new VM. Defaults to empty.
	VMDescription string `mapstructure:"vm_description"`
	// Name label of the SR for the VM's disks. On the iso builder it defaults to
	// the pool's default SR, and also holds the `cd_files` CD and a downloaded
	// ISO, so it must be writable. On the clone builder it defaults to the
	// source's SR, and naming another SR copies the disks in full.
	VMSR string `mapstructure:"vm_sr"`
	// Name label of the SR that holds the ISO named by `iso_name`. It is only
	// read, so a read-only ISO library works.
	ISOSR string `mapstructure:"iso_sr"`
	// Files and directories to put on a CD attached at boot, such as
	// `Autounattend.xml`. Works under BIOS and UEFI. Wildcards are not expanded.
	CDFiles []string `mapstructure:"cd_files"`
	// Networks to attach, by name label, in order. They replace the source's
	// networks. Defaults to the source's networks, or the management network if it
	// has none.
	NetworkNames []string `mapstructure:"network_names"`

	// Number of virtual CPUs. Defaults to the source's. On the iso builder it is
	// required with the generic `Other install media` template, which has none.
	VCPUs *uint `mapstructure:"vcpus"`
	// Cores per socket. Must divide `vcpus`, giving at most 16 sockets. Defaults
	// to the source's, and is ignored unless `vcpus` is set too.
	CorePerSocket *uint `mapstructure:"cores_per_socket"`
	// Memory in GiB, as a fixed allocation with no ballooning. Defaults to the
	// source's. On the iso builder it is required with `Other install media`.
	VMMemoryGB *uint `mapstructure:"vm_memory_gb"`
	// Size of the primary disk in GiB. The iso builder creates the disk at this
	// size; the clone builder can only grow the source's disk, so a smaller size
	// fails the build. Defaults to the source's. On the iso builder it is required
	// with `Other install media`.
	DiskSizeGB *uint `mapstructure:"disk_size_gb"`
	// Key/value pairs added to the VM's `other-config`. `xs_packer_plugin` and
	// `xs_packer_plugin_version` are reserved: an entry with either name is
	// overwritten.
	VMOtherConfig map[string]string `mapstructure:"vm_other_config"`
	// vGPU profile to attach, matched as a case-insensitive substring of the
	// pool's vGPU type names. The build fails if no type matches.
	VGPUProfile string `mapstructure:"vgpu_profile"`

	// `bios` or `uefi`. Defaults to the source's firmware. The build fails early
	// if the source cannot boot that way, for example `bios` with a source that
	// has a vTPM.
	Firmware string `mapstructure:"firmware"`
	// Enable UEFI Secure Boot. `true` requires `firmware = "uefi"`, set
	// explicitly; `false` turns Secure Boot off. Defaults to the source's setting.
	// On a pool with no UEFI certificates installed, the VM starts with Secure
	// Boot off.
	SecureBoot *bool `mapstructure:"secure_boot"`
	// Give the VM the device through which Windows Update offers the XenServer PV
	// drivers. Windows needs it for a clean shutdown, reboot and suspend. Defaults
	// to `true` on the iso builder, and to the source's setting on the clone
	// builder, where changing it on an installed Windows guest is likely to stop
	// it booting.
	HasVendorDevice *bool `mapstructure:"has_vendor_device"`

	// Keys typed over VNC to start the installer. See the Boot Command section.
	BootCommand []string `mapstructure:"boot_command"`
	// Command run inside the guest over WinRM to shut it down after provisioning.
	// Without it, the VM is shut down through the XenAPI.
	ShutdownCommand string `mapstructure:"shutdown_command"`

	// How long to wait after power-on before typing `boot_command`. Defaults to
	// 5s.
	BootWait time.Duration `mapstructure:"boot_wait"`
	// How many times to send `boot_command`, `boot_key_group_interval` apart.
	// Defaults to 20 on a UEFI iso build whose `boot_command` is a single
	// `<enter>`, `<return>` or `<spacebar>`, and to 1 otherwise.
	BootKeyGroups int `mapstructure:"boot_key_groups"`
	// Time between `boot_command` repetitions. Required when `boot_key_groups` is
	// above 1. Defaults to 1s with the UEFI default above.
	BootKeyGroupInterval time.Duration `mapstructure:"boot_key_group_interval"`
	// How long to wait for the guest: with `communicator = "winrm"`, until it
	// reports an address; with `communicator = "none"`, until it shuts down.
	// Defaults to 200m.
	InstallTimeout time.Duration `mapstructure:"install_timeout"`

	// Directory served over HTTP to the VM. Defaults to none.
	HTTPDir string `mapstructure:"http_directory"`
	// IP address the HTTP server binds to, also handed to the guest as `.HTTPIP`.
	// Defaults to every interface, with `.HTTPIP` taken from the route to
	// `xenserver_host`.
	HTTPAddress string `mapstructure:"http_bind_address"`
	// Network interface whose IPv4 address is handed to the guest as `.HTTPIP`.
	// Use it instead of `http_bind_address`, not with it.
	HTTPInterface string `mapstructure:"http_interface"`

	SSHConfig `mapstructure:",squash"`

	// Directory for the exported files. Defaults to `output-BUILDNAME`. It must be
	// missing or empty; `-force` empties it first, and also destroys any VM or
	// template on the pool named `vm_name`.
	OutputDir string `mapstructure:"output_directory"`
	// What the build produces: `xva` (the default), `vdi_raw`, `vdi_vhd`, or
	// `none`, which converts the finished VM into a template instead of exporting
	// a file. `xva_compressed` is a deprecated alias of `xva`.
	Format string `mapstructure:"format"`
	// When to keep the VM on the pool: `never`, `always` or `on_success`. Defaults
	// to `on_success` with `format = "none"`, and to `never` otherwise.
	KeepVM string `mapstructure:"keep_vm"`
	// Snapshot the VM after it shuts down, before it is exported. Defaults to
	// false. The build never deletes the snapshot.
	CreateSnapshot bool `mapstructure:"create_snapshot"`
	// Name label of that snapshot. Only valid with `create_snapshot = true`.
	SnapshotName string `mapstructure:"snapshot_name"`

	// Defaults to false. When true, the builder does not verify the XenServer TLS
	// certificate at all (chain, name and expiry), so the connection can be
	// intercepted. Use only in a trusted lab. For a host with a self-signed
	// certificate, set `ca_cert_path` to its certificate instead.
	//
	// Falls back to the `XENSERVER_INSECURE` environment variable, with a warning
	// when that turns verification off. An explicit `insecure = false` in the
	// template is not overridden by the environment.
	Insecure *bool `mapstructure:"insecure"`
	// Path to the PEM-encoded CA certificate used to verify the XenServer TLS
	// connection. Required unless `insecure` is true. Falls back to the
	// `XENSERVER_CA_CERT_PATH` environment variable.
	CACertPath string `mapstructure:"ca_cert_path"`

	ctx interpolate.Context
}

// Environment variables that Prepare falls back to for connection options the template leaves unset.
const (
	EnvHost       = "XENSERVER_HOST"
	EnvUsername   = "XENSERVER_USERNAME"
	EnvPassword   = "XENSERVER_PASSWORD"
	EnvCACertPath = "XENSERVER_CA_CERT_PATH"
	EnvInsecure   = "XENSERVER_INSECURE"
)

// InsecureEnabled reports whether the build asked to skip TLS verification.
func (c CommonConfig) InsecureEnabled() bool {
	return c.Insecure != nil && *c.Insecure
}

func (c *CommonConfig) applyConnectionEnv() (warns []string, err error) {
	// Read before the environment fills it in: only a ca_cert_path from the template overrides XENSERVER_INSECURE.
	pinnedByTemplate := c.CACertPath != ""

	applyEnvFallbacks([]envFallback{
		{&c.HostIp, EnvHost},
		{&c.Username, EnvUsername},
		{&c.Password, EnvPassword},
		{&c.CACertPath, EnvCACertPath},
	})

	if c.Insecure == nil {
		if raw := os.Getenv(EnvInsecure); raw != "" {
			v, parseErr := strconv.ParseBool(raw)
			if parseErr != nil {
				// Never guess: a misread value could silently turn verification off.
				return warns, fmt.Errorf("%s must be a boolean (for example \"true\" or \"0\"), got '%s'", EnvInsecure, raw)
			}
			// A template that pins a certificate wants it verified; a leftover XENSERVER_INSECURE must not undo that.
			if v && pinnedByTemplate {
				warns = append(warns, fmt.Sprintf("%s=%s was ignored because this build sets ca_cert_path, "+
					"so the pool's certificate is still verified. Set insecure = true in the template to skip "+
					"verification.", EnvInsecure, raw))
			} else {
				c.Insecure = &v
				if v {
					warns = append(warns, fmt.Sprintf("%s=%s turns off verification of the pool's certificate "+
						"for this build. Unset it to verify, or set insecure = true in the template to skip "+
						"verification on purpose.", EnvInsecure, raw))
				}
			}
		}
	}

	return warns, nil
}

func validateHost(host string) error {
	switch {
	case strings.HasPrefix(strings.ToLower(host), "http://"):
		return fmt.Errorf("xenserver_host must not use http://: the XenAPI login and the session "+
			"reference it returns would travel in plaintext. Use https:// or a bare host name, got '%s'", host)
	case strings.HasPrefix(strings.ToLower(host), "https://"), !strings.Contains(host, "://"):
		return nil
	default:
		return fmt.Errorf("xenserver_host must be a host name, an IP address, or an https:// URL, got '%s'", host)
	}
}

// SecureBootEnabled reports whether the build asked for Secure Boot.
func (c CommonConfig) SecureBootEnabled() bool {
	return c.SecureBoot != nil && *c.SecureBoot
}

// HasVendorDeviceEnabled reports whether the VM gets the vendor device, treating unset as true.
// For the iso builder only: on a clone, writing the flag can stop the installed guest from booting.
func (c CommonConfig) HasVendorDeviceEnabled() bool {
	return c.HasVendorDevice == nil || *c.HasVendorDevice
}

// Name returns the VM's name-label.
func (c CommonConfig) Name() string {
	if c.VMName == nil {
		return ""
	}
	return *c.VMName
}

// GetInterpContext returns the shared interpolation context; with a value receiver it would point into a copy.
func (c *CommonConfig) GetInterpContext() *interpolate.Context {
	return &c.ctx
}

func (c *CommonConfig) Prepare(pc *common.PackerConfig) (warns []string, errs []error) {
	if c.BootWait == 0 {
		c.BootWait = 5 * time.Second
	}

	// Left at zero for StepTypeBootCommand to default: that needs a XenAPI session, and an explicit 1 must stay 1.
	if c.BootKeyGroups < 0 {
		errs = append(errs, fmt.Errorf("boot_key_groups must be at least 1, got %d", c.BootKeyGroups))
	}

	if c.BootKeyGroups > 1 && c.BootKeyGroupInterval == 0 {
		errs = append(errs, errors.New("boot_key_groups is greater than 1 but "+
			"boot_key_group_interval is not set, so every repetition is sent at once and covers "+
			"no more of the boot prompt than one send would. Set boot_key_group_interval to the "+
			"gap between repetitions, or remove boot_key_groups."))
	}

	if c.InstallTimeout == 0 {
		c.InstallTimeout = 200 * time.Minute
	}

	if c.CDFiles == nil {
		c.CDFiles = make([]string, 0)
	}

	if c.OutputDir == "" {
		c.OutputDir = fmt.Sprintf("output-%s", pc.PackerBuildName)
	}

	if c.VMName != nil && *c.VMName == "" {
		errs = append(errs, errors.New("vm_name must not be empty: it is the VM's name-label and the name of "+
			"the exported file. Remove vm_name to get the default packer-{{build_name}}-{{timestamp}}."))
	}
	for _, o := range []struct {
		name  string
		value *uint
	}{
		{"vcpus", c.VCPUs},
		{"cores_per_socket", c.CorePerSocket},
		{"vm_memory_gb", c.VMMemoryGB},
		{"disk_size_gb", c.DiskSizeGB},
	} {
		if o.value != nil && *o.value == 0 {
			errs = append(errs, fmt.Errorf("%s must be greater than 0. Remove it to take the builder's "+
				"default, which for the clone builder is whatever the source has.", o.name))
		}
	}

	// mapstructure decodes an absent list to nil but [] to non-nil; packer validate decodes an unknown list to [] too.
	if c.NetworkNames != nil && len(c.NetworkNames) == 0 {
		errs = append(errs, errors.New("network_names names no network. Omit the option to take the "+
			"builder's default networking, or list at least one network name-label."))
	}
	for i, name := range c.NetworkNames {
		// Checked here because the build destroys the inherited VIFs before it resolves these names.
		if strings.TrimSpace(name) == "" {
			errs = append(errs, fmt.Errorf("network_names[%d] names no network. Every entry has to "+
				"be the name-label of a network on the pool.", i))
		}
	}

	if c.VMName == nil {
		// The SDK interpolates before Prepare, so render {{timestamp}} here; InitTime is what it expands to.
		name := fmt.Sprintf("packer-%s-%d", pc.PackerBuildName, interpolate.InitTime.Unix())
		c.VMName = &name
	}

	if c.Format == "" {
		c.Format = "xva"
	}

	if c.Format == "xva_compressed" {
		warns = append(warns, "xva_compressed is deprecated: XVA exports are always compressed. "+
			"Use \"xva\" instead — the two are equivalent.")
		c.Format = "xva"
	}

	// With format = "none" the VM is the output, so a "never" default would delete it.
	if c.KeepVM == "" {
		if c.Format == "none" {
			c.KeepVM = "on_success"
		} else {
			c.KeepVM = "never"
		}
	}

	envWarns, err := c.applyConnectionEnv()
	warns = append(warns, envWarns...)
	if err != nil {
		errs = append(errs, err)
	}

	// Packer only fills this filter from HCL sensitive values, so a password from the environment is added here.
	packersdk.LogSecretFilter.Set(c.Password)

	if c.Username == "" {
		errs = append(errs, fmt.Errorf("xenserver_username must be specified (or set %s).", EnvUsername))
	}

	if c.Password == "" {
		errs = append(errs, fmt.Errorf("xenserver_password must be specified (or set %s).", EnvPassword))
	}

	if c.HostIp == "" {
		errs = append(errs, fmt.Errorf("xenserver_host must be specified (or set %s).", EnvHost))
	} else if err := validateHost(c.HostIp); err != nil {
		errs = append(errs, err)
	}

	switch {
	case c.InsecureEnabled() && c.CACertPath != "":
		warns = append(warns, fmt.Sprintf("insecure = true overrides ca_cert_path '%s': the pool's "+
			"certificate is not verified and that file is never read. Remove insecure to verify "+
			"against it.", c.CACertPath))
	case c.InsecureEnabled():
		// Nothing to verify against, by request.
	case c.CACertPath == "":
		errs = append(errs, fmt.Errorf("ca_cert_path must be specified when insecure is false. "+
			"Point it at the pool's certificate (or set %s), or set insecure = true to skip "+
			"verification, which is for non-production use only.", EnvCACertPath))
	default:
		// Load it now: the goSDK calls log.Fatal on a bad CA certificate, killing the plugin without a Packer error.
		if _, err := loadCACertPool(c.CACertPath); err != nil {
			errs = append(errs, err)
		}
	}

	for _, f := range c.CDFiles {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			errs = append(errs, fmt.Errorf("cd_files: '%s' does not exist", f))
		}
	}

	errs = append(errs, validateBootCommandWaits(c.BootCommand)...)

	if c.HTTPAddress != "" && c.HTTPInterface != "" {
		errs = append(errs, errors.New("only one of http_bind_address or http_interface may be specified"))
	}

	// It is also handed to the guest as {{ .HTTPIP }}, so only a literal address works.
	if c.HTTPAddress != "" && net.ParseIP(c.HTTPAddress) == nil {
		errs = append(errs, fmt.Errorf("http_bind_address must be an IP address, got '%s'", c.HTTPAddress))
	}

	if c.ShutdownCommand != "" && c.Comm.Type == "none" {
		errs = append(errs, errors.New("shutdown_command requires a communicator: it runs inside the guest, "+
			"so it cannot be used with communicator = \"none\". Remove it; where the build does boot the "+
			"guest, have the unattended installation shut it down instead"))
	}

	if c.Comm.Type == "none" {
		warns = append(warns, "communicator = \"none\": the build never logs in to the guest, "+
			"so no provisioner runs.")
	}

	// "xva_compressed" is absent on purpose: it was normalized to "xva" above.
	switch c.Format {
	case "xva", "vdi_raw", "vdi_vhd", "none":
	default:
		errs = append(errs, errors.New("format must be one of 'xva', 'vdi_raw', 'vdi_vhd', 'none'"))
	}

	switch c.KeepVM {
	case "always", "never", "on_success":
	default:
		errs = append(errs, errors.New("keep_vm must be one of 'always', 'never', 'on_success'"))
	}

	// XAPI and both builders expect lowercase. Whitespace-only is left for the switch to reject, not read as unset.
	if lowered := strings.ToLower(strings.TrimSpace(c.Firmware)); lowered != "" {
		c.Firmware = lowered
	}

	switch c.Firmware {
	case "", "bios", "uefi":
	default:
		errs = append(errs, fmt.Errorf("firmware must be one of 'bios', 'uefi', got '%s'", c.Firmware))
	}

	if c.Format == "none" && c.KeepVM == "never" {
		errs = append(errs, errors.New("format = \"none\" with keep_vm = \"never\" produces no output: "+
			"the build VM is converted to a template and then destroyed. Remove keep_vm to keep the "+
			"template, since format = \"none\" defaults it to \"on_success\", or pick an export format."))
	}

	if c.SnapshotName != "" && !c.CreateSnapshot {
		errs = append(errs, errors.New("snapshot_name is set but create_snapshot is not, so no "+
			"snapshot is taken and snapshot_name has no effect. Set create_snapshot = true to take "+
			"one, or remove snapshot_name."))
	}

	return warns, errs
}

// Uint returns a pointer to v.
func Uint(v uint) *uint {
	return &v
}

// ValidateCPUTopology rejects a vcpus and cores_per_socket pair that XenServer cannot build a VM from.
func ValidateCPUTopology(vcpus, coresPerSocket *uint) error {
	// Prepare reports a zero under the option's own name.
	if vcpus == nil || coresPerSocket == nil || *vcpus == 0 || *coresPerSocket == 0 {
		return nil
	}

	if *vcpus%*coresPerSocket != 0 {
		return fmt.Errorf("cores_per_socket (%d) must divide vcpus (%d) exactly: XenServer derives the "+
			"socket count from the two and refuses to start a VM whose vcpus is not a multiple of it",
			*coresPerSocket, *vcpus)
	}

	// XAPI does not check this, so the VM would fail below the toolstack: Xen allows at most maxSockets.
	if sockets := *vcpus / *coresPerSocket; sockets > maxSockets {
		return fmt.Errorf("cores_per_socket (%d) over vcpus (%d) gives the VM %d sockets, more than the %d "+
			"XenServer supports: raise cores_per_socket, or lower vcpus", *coresPerSocket, *vcpus,
			sockets, maxSockets)
	}

	return nil
}

// ShouldKeepVM reports whether cleanup must leave the build VM in place.
func (c CommonConfig) ShouldKeepVM(state multistep.StateBag) bool {
	switch c.KeepVM {
	case "always":
		return true
	case "never":
		return false
	case "on_success":
		_, cancelled := state.GetOk(multistep.StateCancelled)
		_, halted := state.GetOk(multistep.StateHalted)
		return !(cancelled || halted)
	default:
		panic(fmt.Sprintf("Unknown keep_vm value '%s'", c.KeepVM))
	}
}

const nullRef = "OpaqueRef:NULL"

func (config CommonConfig) GetSR(c *Connection) (xenapi.SRRef, error) {
	if config.VMSR != "" {
		return findSRByNameLabel(c, config.VMSR)
	}

	hostRef, err := c.session.GetThisHost(c.ref)
	if err != nil {
		return "", err
	}

	pools, err := xenapi.Pool.GetAllRecords(c.session)
	if err != nil {
		return "", err
	}

	for _, pool := range pools {
		if pool.Master == hostRef {
			if pool.DefaultSR == "" || pool.DefaultSR == nullRef {
				return "", errors.New("the pool has no default SR: set vm_sr to the name-label of " +
					"the SR the build should use")
			}
			return pool.DefaultSR, nil
		}
	}

	return "", fmt.Errorf("no pool on this connection has host '%s' as its primary, so its default SR "+
		"was never looked at: point xenserver_host at the pool primary, or set vm_sr", hostRef)
}

func findSRByNameLabel(c *Connection, nameLabel string) (xenapi.SRRef, error) {
	var srRef xenapi.SRRef

	srs, err := xenapi.SR.GetByNameLabel(c.session, nameLabel)
	if err != nil {
		return srRef, err
	}

	switch {
	case len(srs) == 0:
		return srRef, fmt.Errorf("Couldn't find an SR with the specified name-label '%s'", nameLabel)
	case len(srs) > 1:
		return srRef, fmt.Errorf("Found more than one SR with the name '%s'. The name must be unique", nameLabel)
	}

	return srs[0], nil
}
