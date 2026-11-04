// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/packer"

	"xenapi"
)

const testExportHost = "10.0.0.1"

func TestExportXVAURL_AuthenticatesWithTheSessionReference(t *testing.T) {
	raw := exportXVAURL(testExportHost, "use_compression=true&", "vm-uuid-123", xenapi.SessionRef(testSessionRef))

	parsed, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("the export URL does not parse: %s", err)
	}
	if got := parsed.Query().Get("session_id"); got != testSessionRef {
		t.Errorf("session_id = %q, want the session reference itself", got)
	}
	if got := parsed.Query().Get("uuid"); got != "vm-uuid-123" {
		t.Errorf("uuid = %q, want \"vm-uuid-123\"", got)
	}
	if got := parsed.Query().Get("use_compression"); got != "true" {
		t.Errorf("use_compression = %q, want \"true\"", got)
	}
	if parsed.Scheme != "https" {
		t.Errorf("scheme = %q, want https", parsed.Scheme)
	}
}

func TestExportVDIURL_CarriesNoPassword(t *testing.T) {
	raw := exportVDIURL(testExportHost, xenapi.SessionRef(testSessionRef), "vdi-uuid-456", "&format=vhd")

	parsed, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("the export URL does not parse: %s", err)
	}
	if parsed.User != nil {
		t.Errorf("the URL carries userinfo: %s", parsed.User)
	}
	if got := parsed.Query().Get("session_id"); got != testSessionRef {
		t.Errorf("session_id = %q, want the session reference", got)
	}
	if got := parsed.Query().Get("vdi"); got != "vdi-uuid-456" {
		t.Errorf("vdi = %q, want \"vdi-uuid-456\"", got)
	}
	if got := parsed.Query().Get("format"); got != "vhd" {
		t.Errorf("format = %q, want \"vhd\"", got)
	}
}

func TestExportVDIURL_OmitsTheFormatForRaw(t *testing.T) {
	raw := exportVDIURL(testExportHost, xenapi.SessionRef(testSessionRef), "vdi-uuid-456", "")

	parsed, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("the export URL does not parse: %s", err)
	}
	if parsed.Query().Has("format") {
		t.Errorf("a raw export should name no format, got %q", parsed.Query().Get("format"))
	}
}

func TestExportURLs_AreSafeToLog(t *testing.T) {
	for name, raw := range map[string]string{
		"xva": exportXVAURL(testExportHost, "", "vm-uuid-123", xenapi.SessionRef(testSessionRef)),
		"vdi": exportVDIURL(testExportHost, xenapi.SessionRef(testSessionRef), "vdi-uuid-456", ""),
	} {
		if redacted := redactURL(raw); strings.Contains(redacted, testSessionRef) {
			t.Errorf("%s: the session reference survived redaction: %s", name, redacted)
		}
	}
}

func TestExportLogLine_HidesTheSessionReference(t *testing.T) {
	exportURL := exportXVAURL(testExportHost, "use_compression=true&", "vm-uuid-123", xenapi.SessionRef(testSessionRef))

	line := exportLogLine(exportURL)

	if strings.Contains(line, testSessionRef) {
		t.Errorf("the session reference reached the build log: %s", line)
	}
	if !strings.Contains(line, "vm-uuid-123") {
		t.Errorf("the log line no longer says what is being exported: %s", line)
	}
}

func TestXeExportLogLine_HidesThePassword(t *testing.T) {
	args := xeExportArgs(testExportHost, "root", "hunter2", "vm-uuid-123", "compress=true", "out/vm.xva")

	line := xeExportLogLine("/usr/bin/xe", args)

	if strings.Contains(line, "hunter2") {
		t.Errorf("the pool password reached the build log: %s", line)
	}
	if !strings.Contains(line, "vm-export") {
		t.Errorf("the log line no longer says what is being run: %s", line)
	}
}

func TestDownloadFile_RejectsANonOKResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = io.WriteString(w, "<error class=\"session_invalid\"/>")
	}))
	defer server.Close()

	target := filepath.Join(t.TempDir(), "export.xva")
	err := downloadFile(context.Background(), server.Client(), server.URL, target, packer.TestUi(t))

	if err == nil {
		t.Fatal("a 403 was reported as a successful export")
	}
	if !strings.Contains(err.Error(), "403") {
		t.Errorf("error %q should report the status that failed", err)
	}
}

// Unlike the upload, which follows only 302, the export is a GET and must follow 301 and 303 too.
func TestDownloadFile_FollowsARedirectToTheHostHoldingTheVDI(t *testing.T) {
	const exported = "exported bytes"

	owner := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, exported)
	}))
	defer owner.Close()

	for _, status := range []int{http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther} {
		contacted := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, owner.URL+"/export_raw_vdi", status)
		}))

		target := filepath.Join(t.TempDir(), "disk.raw")
		err := downloadFile(context.Background(), contacted.Client(), contacted.URL+"/export_raw_vdi", target, packer.TestUi(t))
		contacted.Close()
		if err != nil {
			t.Errorf("%d: unexpected error: %s", status, err)
			continue
		}

		written, err := os.ReadFile(target)
		if err != nil {
			t.Errorf("%d: unexpected error: %s", status, err)
			continue
		}
		if string(written) != exported {
			t.Errorf("%d: the export wrote %q, want the redirected host's body", status, written)
		}
	}
}

func writeExport(t *testing.T, name string, content []byte) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, content, 0o600); err != nil {
		t.Fatalf("could not write the test artifact: %s", err)
	}
	return path
}

func incompressible(t *testing.T, n int) []byte {
	t.Helper()

	buf := make([]byte, n)
	_, _ = rand.New(rand.NewSource(1)).Read(buf)
	return buf
}

func TestValidateExportedXVA_RejectsAnErrorDocument(t *testing.T) {
	path := writeExport(t, "export.xva", []byte("<error class=\"session_invalid\"/>"))

	err := validateExportedXVA(path)

	if err == nil {
		t.Fatal("an XML error document was accepted as an XVA")
	}
	if !strings.Contains(err.Error(), "too small") {
		t.Errorf("error %q should say the artifact is too small", err)
	}
}

func TestValidateExportedXVA_RejectsAnEmptyFile(t *testing.T) {
	path := writeExport(t, "export.xva", nil)

	if err := validateExportedXVA(path); err == nil {
		t.Fatal("an empty file was accepted as an XVA")
	}
}

func TestValidateExportedXVA_RejectsAFileThatIsBigEnoughButNotAnArchive(t *testing.T) {
	body := []byte("<html><body>Internal Server Error</body></html>")
	body = append(body, bytes.Repeat([]byte(" "), minExportSize)...)
	path := writeExport(t, "export.xva", body)

	err := validateExportedXVA(path)

	if err == nil {
		t.Fatal("an HTML error page was accepted as an XVA")
	}
	if !strings.Contains(err.Error(), "not an XVA") {
		t.Errorf("error %q should say the artifact is not an XVA", err)
	}
}

func TestValidateExportedXVA_AcceptsAGzipStream(t *testing.T) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write(incompressible(t, 4096)); err != nil {
		t.Fatalf("could not build the gzip stream: %s", err)
	}
	if err := zw.Close(); err != nil {
		t.Fatalf("could not close the gzip stream: %s", err)
	}
	path := writeExport(t, "export.xva", buf.Bytes())

	if err := validateExportedXVA(path); err != nil {
		t.Errorf("a gzip stream was rejected: %s", err)
	}
}

func TestValidateExportedXVA_AcceptsAnUncompressedTar(t *testing.T) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	content := incompressible(t, 2048)
	if err := tw.WriteHeader(&tar.Header{Name: "ova.xml", Mode: 0o600, Size: int64(len(content))}); err != nil {
		t.Fatalf("could not build the tar header: %s", err)
	}
	if _, err := tw.Write(content); err != nil {
		t.Fatalf("could not build the tar: %s", err)
	}
	if err := tw.Close(); err != nil {
		t.Fatalf("could not close the tar: %s", err)
	}
	path := writeExport(t, "export.xva", buf.Bytes())

	if err := validateExportedXVA(path); err != nil {
		t.Errorf("an uncompressed tar was rejected: %s", err)
	}
}

func TestValidateExportedXVA_RejectsAMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "never-written.xva")

	if err := validateExportedXVA(path); err == nil {
		t.Fatal("a missing artifact was reported as a successful export")
	}
}

func TestValidateExportedFile_RejectsAnErrorDocumentButAcceptsADiskImage(t *testing.T) {
	tooSmall := writeExport(t, "disk.vhd", []byte("<error class=\"vdi_io_error\"/>"))
	if err := validateExportedFile(tooSmall); err == nil {
		t.Error("an XML error document was accepted as a disk image")
	}

	bigEnough := writeExport(t, "disk.raw", incompressible(t, minExportSize))
	if err := validateExportedFile(bigEnough); err != nil {
		t.Errorf("a disk image was rejected: %s", err)
	}
}

func TestDownloadFile_A200ThatIsNotAnExportIsCaughtByTheArtifactCheck(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// XAPI has already written the status line by the time the export fails.
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "<error class=\"vdi_missing\"/>")
	}))
	defer server.Close()

	target := filepath.Join(t.TempDir(), "export.xva")

	if err := downloadFile(context.Background(), server.Client(), server.URL, target, packer.TestUi(t)); err != nil {
		t.Fatalf("the download itself should succeed on a 200: %s", err)
	}
	if err := validateExportedXVA(target); err == nil {
		t.Fatal("a 200 carrying an error document was reported as a successful export")
	}
}

// CommonConfig.Prepare accepts "HTTPS://", so stripping the scheme must ignore case.
func TestExportXVAURL_FromAHostThatWasWrittenWithAScheme(t *testing.T) {
	host := stripScheme("HTTPS://xs.example.com:443")

	raw := exportXVAURL(host, "", "vm-uuid-123", xenapi.SessionRef(testSessionRef))

	parsed, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("the export URL does not parse: %s", err)
	}
	if parsed.Host != "xs.example.com:443" {
		t.Errorf("host = %q, want \"xs.example.com:443\"", parsed.Host)
	}
	if parsed.Path != "/export" {
		t.Errorf("path = %q, want \"/export\"", parsed.Path)
	}
}
