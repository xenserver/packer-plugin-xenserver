// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"xenapi"

	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/internal/xstest"
)

func writeTestCACert(t *testing.T) string {
	t.Helper()
	return xstest.WriteCACert(t)
}

func writeTestPrivateKeyPEM(t *testing.T) string {
	t.Helper()
	return xstest.WritePrivateKeyPEM(t)
}

func TestNewTLSConfig_VerifiesAgainstTheGivenCA(t *testing.T) {
	config, err := newTLSConfig(false, writeTestCACert(t))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if config.InsecureSkipVerify {
		t.Error("verification is off despite insecure = false")
	}
	if config.RootCAs == nil {
		t.Error("ca_cert_path was not loaded into the trust roots")
	}
	if config.MinVersion != tls.VersionTLS12 {
		t.Errorf("MinVersion = %#x, want TLS 1.2", config.MinVersion)
	}
}

func TestNewTLSConfig_HonoursInsecure(t *testing.T) {
	config, err := newTLSConfig(true, "")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !config.InsecureSkipVerify {
		t.Error("insecure = true did not turn verification off")
	}
	if config.MinVersion != tls.VersionTLS12 {
		t.Errorf("MinVersion = %#x, want TLS 1.2 even when not verifying", config.MinVersion)
	}
}

func TestNewTLSConfig_RejectsAnUnusableCACert(t *testing.T) {
	if _, err := newTLSConfig(false, filepath.Join(t.TempDir(), "absent.pem")); err == nil {
		t.Error("a missing ca_cert_path should be an error, not a silent empty trust store")
	}

	notPEM := filepath.Join(t.TempDir(), "not-a-cert.pem")
	if err := os.WriteFile(notPEM, []byte("nope"), 0600); err != nil {
		t.Fatalf("could not write the fixture: %s", err)
	}
	if _, err := newTLSConfig(false, notPEM); err == nil {
		t.Error("a ca_cert_path holding no certificate should be an error")
	}
}

func TestConnectionTLSConfig_ReturnsACopy(t *testing.T) {
	config, err := newTLSConfig(true, "")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	c := &Connection{tlsConfig: config}

	c.TLSConfig().ServerName = "console.example.com"

	if c.tlsConfig.ServerName != "" {
		t.Error("mutating the returned config changed the connection's own policy")
	}
}

func TestConnectionString_PrintsNoSecrets(t *testing.T) {
	c := &Connection{
		Host:     "10.0.0.1",
		Username: "root",
		Password: "hunter2",
		ref:      "OpaqueRef:11111111-2222-3333-4444-555555555555",
	}

	for _, verb := range []string{"%v", "%s", "%+v", "%#v"} {
		rendered := fmt.Sprintf(verb, c)

		if strings.Contains(rendered, "hunter2") {
			t.Errorf("%s printed the password: %s", verb, rendered)
		}
		if strings.Contains(rendered, "OpaqueRef") {
			t.Errorf("%s printed the session reference: %s", verb, rendered)
		}
	}
}

func TestNormalizeXenAPIHost(t *testing.T) {
	for _, tc := range []struct{ in, wantURL, wantHost string }{
		{"10.0.0.1", "https://10.0.0.1", "10.0.0.1"},
		{"10.0.0.1:8443", "https://10.0.0.1:8443", "10.0.0.1:8443"},
		{"https://xs.example.com", "https://xs.example.com", "xs.example.com"},
		{"https://xs.example.com/", "https://xs.example.com", "xs.example.com"},
		{"HTTPS://xs.example.com", "HTTPS://xs.example.com", "xs.example.com"},
		{"Https://xs.example.com:443/", "Https://xs.example.com:443", "xs.example.com:443"},
	} {
		gotURL, gotHost, err := normalizeXenAPIHost(tc.in)

		if err != nil {
			t.Errorf("normalizeXenAPIHost(%q): unexpected error: %s", tc.in, err)
			continue
		}
		if gotURL != tc.wantURL || gotHost != tc.wantHost {
			t.Errorf("normalizeXenAPIHost(%q) = (%q, %q), want (%q, %q)",
				tc.in, gotURL, gotHost, tc.wantURL, tc.wantHost)
		}
	}

	for _, in := range []string{"http://10.0.0.1", "HTTP://10.0.0.1", "ftp://10.0.0.1"} {
		if _, _, err := normalizeXenAPIHost(in); err == nil {
			t.Errorf("normalizeXenAPIHost(%q) should refuse a session that is not over https", in)
		}
	}
}

func TestConnectionHTTPClient_VerifiesTheServer(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()

	serverURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	trusted := x509.NewCertPool()
	trusted.AddCert(server.Certificate())

	for name, tc := range map[string]struct {
		config    *tls.Config
		wantTrust bool
	}{
		"a certificate in ca_cert_path": {
			config:    &tls.Config{MinVersion: tls.VersionTLS12, RootCAs: trusted},
			wantTrust: true,
		},
		"a certificate not in ca_cert_path": {
			config:    &tls.Config{MinVersion: tls.VersionTLS12, RootCAs: x509.NewCertPool()},
			wantTrust: false,
		},
		"insecure": {
			// #nosec G402 -- the point of the case is that the option works.
			config:    &tls.Config{MinVersion: tls.VersionTLS12, InsecureSkipVerify: true},
			wantTrust: true,
		},
	} {
		resp, err := (&Connection{tlsConfig: tc.config}).HTTPClient().Get(server.URL)
		if err == nil {
			resp.Body.Close()
		}

		if accepted := err == nil; accepted != tc.wantTrust {
			t.Errorf("%s: HTTPClient trusted = %v, want %v (%v)", name, accepted, tc.wantTrust, err)
		}

		conn, err := (&Connection{tlsConfig: tc.config}).DialTLS(serverURL.Host)
		if err == nil {
			conn.Close()
		}

		if accepted := err == nil; accepted != tc.wantTrust {
			t.Errorf("%s: DialTLS trusted = %v, want %v (%v)", name, accepted, tc.wantTrust, err)
		}
	}
}

// The error text is asserted because the call still fails with either guard removed.
func TestNewXenAPIClient_RefusesUnsafeArguments(t *testing.T) {
	for name, tc := range map[string]struct {
		host     string
		insecure bool
		want     string
	}{
		"a plaintext host": {
			host: "http://10.0.0.1", insecure: true,
			want: "must be over https",
		},
		"verifying against nothing": {
			host: "10.0.0.1", insecure: false,
			want: "no ca_cert_path",
		},
	} {
		_, err := NewXenAPIClient(tc.host, "root", "secret", tc.insecure, "")

		if err == nil {
			t.Errorf("%s: should have been refused before any credential was sent", name)
			continue
		}
		if !strings.Contains(err.Error(), tc.want) {
			t.Errorf("%s: error %q should say %q", name, err, tc.want)
		}
	}
}

// NewXenAPIClient must reject the path before xenapi.NewSession calls log.Fatal on it.
func TestNewXenAPIClient_RejectsAnUnusableCACert(t *testing.T) {
	badPath := filepath.Join(t.TempDir(), "absent.pem")

	_, err := NewXenAPIClient("10.0.0.1", "root", "secret", false, badPath)

	if err == nil {
		t.Fatal("a verifying connection with an unreadable ca_cert_path must be refused before dialling")
	}
	if !strings.Contains(err.Error(), "ca_cert_path") {
		t.Errorf("error %q should name ca_cert_path", err)
	}
}

func TestConnectVdi_MarksOnlyDataDisksAsOwned(t *testing.T) {
	for name, tc := range map[string]struct {
		vbdType   xenapi.VbdType
		wantOwner bool
	}{
		"a data disk": {vbdType: xenapi.VbdTypeDisk, wantOwner: true},
		"a CD":        {vbdType: xenapi.VbdTypeCD, wantOwner: false},
	} {
		f := newFakeXAPI(t, map[string]interface{}{"VBD.create": "OpaqueRef:vbd"})

		if err := ConnectVdi(f.conn(), "OpaqueRef:vm", "OpaqueRef:vdi", tc.vbdType); err != nil {
			t.Fatalf("%s: unexpected error: %s", name, err)
		}

		record, ok := f.argsOf(t, "VBD.create")[0].(map[string]interface{})
		if !ok {
			t.Fatalf("%s: VBD.create was not passed a record", name)
		}
		otherConfig, _ := record["other_config"].(map[string]interface{})
		_, gotOwner := otherConfig["owner"]

		if gotOwner != tc.wantOwner {
			t.Errorf("%s: other-config:owner present = %v, want %v", name, gotOwner, tc.wantOwner)
		}
	}
}

func TestPoolHostAddress_PortsEveryFormOfAuthority(t *testing.T) {
	for _, tc := range []struct{ host, want string }{
		{"192.0.2.134", "192.0.2.134:443"},
		{"192.0.2.134:443", "192.0.2.134:443"},
		{"xenhost.example.com", "xenhost.example.com:443"},
		{"[fd00::1]", "[fd00::1]:443"},
		{"[fd00::1]:443", "[fd00::1]:443"},
		{"192.0.2.134:", "192.0.2.134:443"},
		{"[fd00::1]:", "[fd00::1]:443"},
	} {
		if got := poolHostAddress(tc.host); got != tc.want {
			t.Errorf("poolHostAddress(%q) = %q, want %q", tc.host, got, tc.want)
		}
	}
}
