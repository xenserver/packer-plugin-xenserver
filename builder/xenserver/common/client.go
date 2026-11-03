// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"xenapi"

	version "github.com/xenserver/packer-plugin-xenserver/version"
)

// GetDisks returns the VDI of every disk-type VBD, root disk included.
func GetDisks(c *Connection, vmRef xenapi.VMRef) (vdis []xenapi.VDIRef, err error) {
	vdis = make([]xenapi.VDIRef, 0)
	vbds, err := xenapi.VM.GetVBDs(c.session, vmRef)
	if err != nil {
		return nil, err
	}

	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(c.session, vbd)
		if err != nil {
			return nil, err
		}
		if rec.Type == xenapi.VbdTypeDisk {
			vdis = append(vdis, rec.VDI)
		}
	}
	return vdis, nil
}

// ConnectVdi attaches a VDI to a halted VM without VBD.plug; XAPI plugs the VBD at VM.Start.
func ConnectVdi(c *Connection, vmRef xenapi.VMRef, vdiRef xenapi.VDIRef, vbdType xenapi.VbdType) error {
	mode, bootable := xenapi.VbdModeRW, false
	if vbdType == xenapi.VbdTypeCD {
		mode, bootable = xenapi.VbdModeRO, true
	}

	// other-config:owner makes vm-uninstall and XenCenter delete the VDI with the VM. Data
	// disks only: vm-uninstall does not check the type, so it would destroy a CD's ISO too.
	var otherConfig map[string]string
	if vbdType == xenapi.VbdTypeDisk {
		otherConfig = map[string]string{"owner": ""}
	}

	_, err := xenapi.VBD.Create(c.session, xenapi.VBDRecord{
		VM:               vmRef,
		VDI:              vdiRef,
		Userdevice:       "autodetect",
		Empty:            false,
		QosAlgorithmType: "",
		Mode:             mode,
		Unpluggable:      false,
		Bootable:         bootable,
		Type:             vbdType,
		OtherConfig:      otherConfig,
	})

	return err
}

// DisconnectVdi unplugs and destroys the VBD attaching vdi to vmRef. The VDI itself is
// left alone.
func DisconnectVdi(c *Connection, vmRef xenapi.VMRef, vdi xenapi.VDIRef) error {
	vbds, err := xenapi.VM.GetVBDs(c.session, vmRef)
	if err != nil {
		return fmt.Errorf("unable to get VM VBDs: %w", err)
	}

	for _, vbd := range vbds {
		rec, err := xenapi.VBD.GetRecord(c.session, vbd)
		if err != nil {
			return fmt.Errorf("could not get record for VBD '%s': %w", vbd, err)
		}
		if rec.VDI == vdi {
			_ = xenapi.VBD.Unplug(c.session, vbd)
			if err := xenapi.VBD.Destroy(c.session, vbd); err != nil {
				return fmt.Errorf("could not destroy VBD '%s': %w", vbd, err)
			}

			return nil
		}
	}

	return fmt.Errorf("could not find VBD for VDI '%s'", vdi)
}

func ConnectNetwork(c *Connection, networkRef xenapi.NetworkRef, vmRef xenapi.VMRef, device string) (*xenapi.VIFRef, error) {
	vif, err := xenapi.VIF.Create(c.session, xenapi.VIFRecord{
		Network:     networkRef,
		VM:          vmRef,
		Device:      device,
		LockingMode: xenapi.VifLockingModeNetworkDefault,
	})

	if err != nil {
		return nil, err
	}
	log.Printf("Created the following VIF: %s", vif)

	return &vif, nil
}

type Connection struct {
	session  *xenapi.Session
	ref      xenapi.SessionRef
	Host     string
	Username string
	Password string
	// TLS policy for the HTTP and VNC channels; the goSDK builds the XenAPI session's own.
	tlsConfig *tls.Config
}

func (c Connection) GetSession() *xenapi.Session {
	return c.session
}

// Logout ends the session so it does not linger on the pool master until it times out.
func (c Connection) Logout() {
	if c.session == nil {
		return
	}
	if err := c.session.Logout(); err != nil {
		log.Printf("Unable to log out of the XAPI session: %s", err.Error())
	}
}

// String omits the password and session reference so that formatting cannot leak them.
func (c Connection) String() string {
	return fmt.Sprintf("Connection{Host: %s, Username: %s}", c.Host, c.Username)
}

// GoString does the same for %#v.
func (c Connection) GoString() string {
	return c.String()
}

// TLSConfig returns a copy of the TLS policy that the caller may modify.
func (c *Connection) TLSConfig() *tls.Config {
	return c.tlsConfig.Clone()
}

// HTTPClient returns an http.Client that verifies the pool like the XenAPI session does.
func (c *Connection) HTTPClient() *http.Client {
	return &http.Client{Transport: &http.Transport{TLSClientConfig: c.TLSConfig()}}
}

const poolDialTimeout = 30 * time.Second

// DialTLS dials host[:port] in the pool under the same TLS policy as HTTPClient.
func (c *Connection) DialTLS(host string) (net.Conn, error) {
	if strings.TrimSpace(host) == "" {
		return nil, errors.New("no host to connect to")
	}

	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: poolDialTimeout}, "tcp",
		poolHostAddress(host), c.TLSConfig())
	if err != nil {
		return nil, describeTLSError(host, err)
	}

	return conn, nil
}

func pluginUserAgent() string {
	return fmt.Sprintf("XenServerPacker/%s", version.PluginVersion.FormattedVersion())
}

// poolHostAddress adds port 443 if host has none. url.URL.Host keeps IPv6 brackets, which
// JoinHostPort would double, and SplitHostPort accepts an empty port ("host:").
func poolHostAddress(host string) string {
	if name, port, err := net.SplitHostPort(host); err == nil {
		if port == "" {
			port = "443"
		}
		return net.JoinHostPort(name, port)
	}

	return net.JoinHostPort(strings.Trim(host, "[]"), "443")
}

// loadCACertPool is also called by Prepare: the goSDK calls log.Fatal on a bad ca_cert_path.
func loadCACertPool(caCertPath string) (*x509.CertPool, error) {
	pemBytes, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, fmt.Errorf("could not read ca_cert_path '%s': %w", caCertPath, err)
	}

	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(pemBytes) {
		return nil, fmt.Errorf("ca_cert_path '%s' contains no PEM-encoded certificate. "+
			"`xe host-get-server-certificate host=<uuid>` writes one out in the expected form", caCertPath)
	}

	return pool, nil
}

// newTLSConfig mirrors the TLS config in goSDK/jsonrpc_client.go; keep the two in sync.
func newTLSConfig(insecure bool, caCertPath string) (*tls.Config, error) {
	config := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	if insecure {
		// #nosec G402 -- the build asked for this explicitly with `insecure`.
		config.InsecureSkipVerify = true
		return config, nil
	}

	pool, err := loadCACertPool(caCertPath)
	if err != nil {
		return nil, err
	}
	config.RootCAs = pool

	return config, nil
}

func describeTLSError(host string, err error) error {
	var unknownAuthority x509.UnknownAuthorityError
	var hostname x509.HostnameError

	switch {
	case errors.As(err, &unknownAuthority):
		return fmt.Errorf("TLS verification of '%s' failed: %w.\n"+
			"A XenServer host's default certificate is self-signed, so it has to be trusted explicitly: "+
			"write it out with `xe host-get-server-certificate host=<uuid>` and point ca_cert_path at it. "+
			"For a pool, concatenate every host's certificate into one PEM file — an export or console "+
			"connection can be served by any host in the pool, not just the one you connected to. "+
			"Setting insecure = true skips verification altogether and is for non-production use only",
			host, err)
	case errors.As(err, &hostname):
		return fmt.Errorf("TLS verification of '%s' failed: %w.\n"+
			"Reach the host by a name or address its certificate is issued for, or reissue the "+
			"certificate to cover the one you use", host, err)
	}

	return err
}

// normalizeXenAPIHost returns the session URL and the bare host that other URLs are built
// from, derived together so that both agree on where the case-insensitive scheme ends.
func normalizeXenAPIHost(host string) (sessionURL, urlHost string, err error) {
	// Test for "://", not an "http" prefix, so a host named "httpserver" still gets a scheme.
	if !strings.Contains(host, "://") {
		host = "https://" + host
	}
	host = strings.TrimSuffix(host, "/")

	if !strings.HasPrefix(strings.ToLower(host), "https://") {
		return "", "", fmt.Errorf("refusing to connect to '%s': the XenAPI session must be over https", host)
	}

	return host, stripScheme(host), nil
}

func NewXenAPIClient(host, username, password string, insecure bool, caCertPath string) (*Connection, error) {
	log.Printf("XenServer Packer Plugin Version: %s", version.PluginVersion.FormattedVersion())

	host, hostForURLs, err := normalizeXenAPIHost(host)
	if err != nil {
		return nil, err
	}
	if !insecure && caCertPath == "" {
		return nil, errors.New("refusing to connect: certificate verification is on but no ca_cert_path was given")
	}

	tlsConfig, err := newTLSConfig(insecure, caCertPath)
	if err != nil {
		return nil, err
	}

	opts := &xenapi.ClientOpts{
		URL: host,
		Headers: map[string]string{
			"User-Agent": pluginUserAgent(),
		},
	}
	if !insecure {
		// ServerCert is the goSDK's own field name; it takes the CA certificate path.
		opts.SecureOpts = &xenapi.SecureOpts{
			ServerCert: caCertPath,
		}
	}
	session := xenapi.NewSession(opts)

	ref, err := session.LoginWithPassword(username, password, "1.0", "packer")
	if err != nil {
		return nil, describeTLSError(host, err)
	}

	return &Connection{
		session:   session,
		ref:       ref,
		Host:      hostForURLs,
		Username:  username,
		Password:  password,
		tlsConfig: tlsConfig,
	}, nil
}

func (c *Connection) GetSessionRef() xenapi.SessionRef {
	return c.ref
}
