// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"net/url"
	"strings"
	"testing"
)

const testSessionRef = "OpaqueRef:11111111-2222-3333-4444-555555555555"

func TestRedactURL_HidesTheSessionReference(t *testing.T) {
	raw := "https://10.0.0.1/export?use_compression=true&uuid=abc-123&session_id=" + testSessionRef

	redacted := redactURL(raw)

	if strings.Contains(redacted, testSessionRef) {
		t.Errorf("the session reference survived redaction: %s", redacted)
	}
	if !strings.Contains(redacted, "uuid=abc-123") || !strings.Contains(redacted, "/export") {
		t.Errorf("redaction removed more than the credential: %s", redacted)
	}
}

func TestRedactURL_HidesBasicCredentials(t *testing.T) {
	redacted := redactURL("https://root:hunter2@10.0.0.1/export_raw_vdi?vdi=abc-123")

	if strings.Contains(redacted, "hunter2") {
		t.Errorf("the password survived redaction: %s", redacted)
	}
}

// Named literally: looping over sensitiveQueryParams would pass however short that list got.
func TestRedactURL_HidesEveryCredentialQueryParam(t *testing.T) {
	for _, key := range []string{"session_id", "password", "pwd"} {
		redacted := redactURL("https://10.0.0.1/export?" + key + "=hunter2")

		if strings.Contains(redacted, "hunter2") {
			t.Errorf("%s survived redaction: %s", key, redacted)
		}
	}
}

func TestRedactURL_SaysNothingAboutAnUnparseableURL(t *testing.T) {
	redacted := redactURL("https://10.0.0.1/export?session_id=" + testSessionRef + "\x7f\x00")

	if strings.Contains(redacted, testSessionRef) {
		t.Errorf("an unparsable URL was echoed back: %s", redacted)
	}
}

func TestRedactURLError_RewritesTheURLItCarries(t *testing.T) {
	raw := "https://10.0.0.1/export?session_id=" + testSessionRef
	err := redactURLError(&url.Error{Op: "Get", URL: raw, Err: errors.New("connection refused")})

	if strings.Contains(err.Error(), testSessionRef) {
		t.Errorf("the session reference reached the error message: %s", err)
	}
	if !strings.Contains(err.Error(), "connection refused") {
		t.Errorf("the cause was lost: %s", err)
	}
}

func TestRedactURLError_LeavesOtherErrorsAlone(t *testing.T) {
	cause := errors.New("something else went wrong")

	if got := redactURLError(cause); got != cause {
		t.Errorf("redactURLError() = %v, want the error unchanged", got)
	}
}

func TestRedactArgs_HidesTheXePassword(t *testing.T) {
	args := xeExportArgs("10.0.0.1", "root", "hunter2", "abc-123", "compress=true", "out/vm.xva")

	redacted := redactArgs(args)

	joined := strings.Join(redacted, " ")
	if strings.Contains(joined, "hunter2") {
		t.Errorf("the password survived redaction: %v", redacted)
	}
	for _, want := range []string{"vm-export", "vm=abc-123", "filename=out/vm.xva", "-u root"} {
		if !strings.Contains(joined, want) {
			t.Errorf("redaction removed more than the credential: %q missing from %s", want, joined)
		}
	}
}

func TestRedactArgs_DoesNotMutateItsInput(t *testing.T) {
	args := xeExportArgs("10.0.0.1", "root", "hunter2", "abc-123", "compress=true", "out/vm.xva")

	_ = redactArgs(args)

	if !strings.Contains(strings.Join(args, " "), "hunter2") {
		t.Error("redactArgs modified the argv in place; xe would be run without its password")
	}
}
