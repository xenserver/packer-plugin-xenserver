// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/packer"
)

func TestUploadLogLine_HidesTheSessionReference(t *testing.T) {
	importURL, err := appendQuery("https://10.0.0.1/import_raw_vdi?vdi=vdi-uuid-456&session_id="+testSessionRef,
		"task_id", "OpaqueRef:task")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	line := uploadLogLine(importURL)

	// appendQuery re-encodes the query, so the reference may appear percent-escaped.
	for _, form := range []string{testSessionRef, url.QueryEscape(testSessionRef)} {
		if strings.Contains(line, form) {
			t.Errorf("the session reference reached the build log: %s", line)
		}
	}
	if !strings.Contains(line, "vdi=vdi-uuid-456") {
		t.Errorf("the log line no longer says which VDI is being uploaded: %s", line)
	}
}

// Both forms: redactURL re-encodes the query, so a leaked reference may be percent-escaped.
func assertNoSessionRef(t *testing.T, err error) {
	t.Helper()

	for _, form := range []string{testSessionRef, url.QueryEscape(testSessionRef)} {
		if strings.Contains(err.Error(), form) {
			t.Errorf("the session reference reached the error: %s", err)
		}
	}
}

func writeImage(t *testing.T) (string, []byte) {
	t.Helper()

	content := []byte("disk image contents")
	path := filepath.Join(t.TempDir(), "image.iso")
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	return path, content
}

func TestRedirectTarget_KeepsEverythingButTheHost(t *testing.T) {
	target, err := redirectTarget(
		"https://192.0.2.119/somewhere/else?session_id=other&vdi=OpaqueRef:theirs",
		"https://192.0.2.116/import_raw_vdi?session_id="+testSessionRef+"&vdi=OpaqueRef:ours")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	want := "https://192.0.2.119/import_raw_vdi?session_id=" + testSessionRef + "&vdi=OpaqueRef:ours"
	if target != want {
		t.Errorf("redirect target = %s, want %s", target, want)
	}
}

func TestRedirectTarget_RejectsALocationNamingNoHost(t *testing.T) {
	_, err := redirectTarget("/import_raw_vdi?session_id="+testSessionRef, "https://192.0.2.116/import_raw_vdi")

	if err == nil {
		t.Fatal("a relative Location leaves nowhere to send the upload, so it must fail")
	}
	assertNoSessionRef(t, err)
}

// XAPI's own responses (ocaml/libs/http-lib/http.ml), sent before it reads any of the body.
const (
	xapiAccepts = "HTTP/1.1 200 OK\r\n" +
		"Connection: close\r\n" +
		"Cache-Control: no-cache, no-store\r\n\r\n"

	xapiRefuses = "HTTP/1.1 403 Forbidden\r\n" +
		"Connection: close\r\n" +
		"Content-Length: 0\r\n\r\n"

	// Ample for a loopback write, short enough not to drag the tests that expect nothing.
	stubBodyWait = 750 * time.Millisecond
)

// XAPI echoes the whole query, session reference included, into the Location.
func xapiRedirectsTo(host string) string {
	return "HTTP/1.1 302 Found\r\n" +
		"Connection: close\r\n" +
		"Cache-Control: no-cache, no-store\r\n" +
		"Location: https://" + host + "/import_raw_vdi?session_id=" + testSessionRef +
		"&vdi=OpaqueRef:ours\r\n\r\n"
}

type exchange struct {
	head string
	body []byte
}

// xapiStub answers before reading the body, as XAPI does; httptest reads the request first.
type xapiStub struct {
	addr      string
	exchanges chan exchange
}

func startXAPI(t *testing.T, response string) *xapiStub {
	t.Helper()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	t.Cleanup(func() { _ = ln.Close() })

	s := &xapiStub{addr: ln.Addr().String(), exchanges: make(chan exchange, maxUploadRedirects+2)}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go s.answer(conn, response)
		}
	}()

	return s
}

func (s *xapiStub) answer(conn net.Conn, response string) {
	defer conn.Close()

	br := bufio.NewReader(conn)
	head, err := readHead(br)
	if err != nil {
		return
	}

	if _, err := io.WriteString(conn, response); err != nil {
		return
	}

	// Read the body even after a refusal or redirect, or "no body was sent" proves nothing.
	announced, err := contentLengthOf(head)
	if err != nil {
		s.exchanges <- exchange{head: head}
		return
	}
	body := make([]byte, announced)
	_ = conn.SetReadDeadline(time.Now().Add(stubBodyWait))
	n, _ := io.ReadFull(br, body)

	s.exchanges <- exchange{head: head, body: body[:n]}
}

func (s *xapiStub) next(t *testing.T) exchange {
	t.Helper()

	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	select {
	case e := <-s.exchanges:
		return e
	case <-timer.C:
		t.Fatal("the upload never reached this host")
		return exchange{}
	}
}

func (s *xapiStub) noMore(t *testing.T) {
	t.Helper()

	timer := time.NewTimer(stubBodyWait)
	defer timer.Stop()

	select {
	case e := <-s.exchanges:
		t.Errorf("the upload made another attempt: %q", e.head)
	case <-timer.C:
	}
}

func readHead(br *bufio.Reader) (string, error) {
	var head strings.Builder
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			return "", err
		}
		head.WriteString(line)
		if line == "\r\n" {
			return head.String(), nil
		}
	}
}

func contentLengthOf(head string) (int64, error) {
	for _, line := range strings.Split(head, "\r\n") {
		if name, value, ok := strings.Cut(line, ":"); ok &&
			strings.EqualFold(strings.TrimSpace(name), "content-length") {
			return strconv.ParseInt(strings.TrimSpace(value), 10, 64)
		}
	}
	return 0, fmt.Errorf("no Content-Length in %q", head)
}

func dialPool(hosts map[string]*xapiStub) dialPoolHost {
	return func(host string) (net.Conn, error) {
		stub, ok := hosts[host]
		if !ok {
			return nil, fmt.Errorf("no host in the pool answers for %s", host)
		}
		return net.Dial("tcp", stub.addr)
	}
}

func upload(t *testing.T, hosts map[string]*xapiStub, contentLength int64) (*http.Response, error) {
	t.Helper()

	path, content := writeImage(t)
	if contentLength == 0 {
		contentLength = int64(len(content))
	}

	return putFollowingRedirects(dialPool(hosts),
		"https://192.0.2.134/import_raw_vdi?session_id="+testSessionRef+"&vdi=OpaqueRef:ours",
		path, contentLength, packer.TestUi(t))
}

func TestPutFollowingRedirects_SendsNoBodyToAHostThatRedirects(t *testing.T) {
	contacted := startXAPI(t, xapiRedirectsTo("192.0.2.119"))
	owner := startXAPI(t, xapiAccepts)

	resp, err := upload(t, map[string]*xapiStub{"192.0.2.134": contacted, "192.0.2.119": owner}, 0)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer resp.Body.Close()

	if sent := contacted.next(t).body; len(sent) != 0 {
		t.Errorf("the host that redirected was sent %d body bytes, want none", len(sent))
	}
	if sent := owner.next(t).body; string(sent) != "disk image contents" {
		t.Errorf("the host that owns the SR received %q", sent)
	}
}

func TestPutFollowingRedirects_SendsNoBodyToAHostThatRefuses(t *testing.T) {
	refusing := startXAPI(t, xapiRefuses)

	resp, err := upload(t, map[string]*xapiStub{"192.0.2.134": refusing}, 0)
	if err != nil {
		t.Fatalf("a refusal is the caller's to report, not an error here: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("status = %s, want 403", resp.Status)
	}
	if sent := refusing.next(t).body; len(sent) != 0 {
		t.Errorf("the host that refused was sent %d body bytes, want none", len(sent))
	}
}

func TestPutFollowingRedirects_SendsTheFileAfterTheAcceptingResponse(t *testing.T) {
	owner := startXAPI(t, xapiAccepts)

	resp, err := upload(t, map[string]*xapiStub{"192.0.2.134": owner}, 0)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %s, want 200", resp.Status)
	}
	if sent := owner.next(t).body; string(sent) != "disk image contents" {
		t.Errorf("the host received %q", sent)
	}
}

func TestPutFollowingRedirects_ReplaysThePutAgainstTheRedirectedHost(t *testing.T) {
	contacted := startXAPI(t, xapiRedirectsTo("192.0.2.119"))
	owner := startXAPI(t, xapiAccepts)

	resp, err := upload(t, map[string]*xapiStub{"192.0.2.134": contacted, "192.0.2.119": owner}, 0)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %s, want the 200 from the host that took the file", resp.Status)
	}

	contacted.next(t)
	replayed := owner.next(t)

	wantLine := "PUT /import_raw_vdi?session_id=" + testSessionRef + "&vdi=OpaqueRef:ours HTTP/1.1"
	if line, _, _ := strings.Cut(replayed.head, "\r\n"); line != wantLine {
		t.Errorf("the redirected request line was %q, want %q", line, wantLine)
	}
	if !strings.Contains(replayed.head, "Host: 192.0.2.119\r\n") {
		t.Errorf("the redirected request was not addressed to the new host: %q", replayed.head)
	}
	if !strings.Contains(replayed.head, "User-Agent: "+pluginUserAgent()+"\r\n") {
		t.Errorf("the upload did not name the plugin to the host: %q", replayed.head)
	}
	if string(replayed.body) != "disk image contents" {
		t.Errorf("the host that owns the SR received %q", replayed.body)
	}
}

func TestPutFollowingRedirects_StopsAfterTooManyRedirects(t *testing.T) {
	looping := startXAPI(t, xapiRedirectsTo("192.0.2.134"))

	_, err := upload(t, map[string]*xapiStub{"192.0.2.134": looping}, 0)

	if err == nil {
		t.Fatal("a redirect loop must end the upload rather than run forever")
	}
	assertNoSessionRef(t, err)

	for i := 0; i <= maxUploadRedirects; i++ {
		looping.next(t)
	}
	looping.noMore(t)
}

func TestPutFollowingRedirects_KeepsTheSessionReferenceOutOfErrors(t *testing.T) {
	path, content := writeImage(t)

	_, err := putFollowingRedirects(
		func(string) (net.Conn, error) { return nil, errors.New("no route to host") },
		"https://192.0.2.134/import_raw_vdi?session_id="+testSessionRef,
		path, int64(len(content)), packer.TestUi(t))

	if err == nil {
		t.Fatal("a host that cannot be reached must fail the upload")
	}
	assertNoSessionRef(t, err)
}

func TestPutFollowingRedirects_RejectsAFileShorterThanAnnounced(t *testing.T) {
	owner := startXAPI(t, xapiAccepts)

	_, err := upload(t, map[string]*xapiStub{"192.0.2.134": owner}, int64(len("disk image contents"))+512)

	if err == nil {
		t.Fatal("a file shorter than the announced length must fail the upload")
	}
	if !strings.Contains(err.Error(), "changed size") {
		t.Errorf("the error should say the file changed size, got: %s", err)
	}
	assertNoSessionRef(t, err)
}

// The stub reads only Content-Length bytes, so it cannot see a client that sends more.
type countingConn struct {
	net.Conn
	written *int64
}

func (c countingConn) Write(b []byte) (int, error) {
	n, err := c.Conn.Write(b)
	*c.written += int64(n)
	return n, err
}

func TestPutFollowingRedirects_SendsNoMoreThanAnnounced(t *testing.T) {
	owner := startXAPI(t, xapiAccepts)
	const announced = 4

	path, content := writeImage(t)
	var written int64
	dial := dialPool(map[string]*xapiStub{"192.0.2.134": owner})

	resp, err := putFollowingRedirects(
		func(host string) (net.Conn, error) {
			conn, err := dial(host)
			if err != nil {
				return nil, err
			}
			return countingConn{Conn: conn, written: &written}, nil
		},
		"https://192.0.2.134/import_raw_vdi?session_id="+testSessionRef,
		path, announced, packer.TestUi(t))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer resp.Body.Close()

	// putFollowingRedirects has returned, so every write it makes is already counted.
	head := owner.next(t).head
	if sent := written - int64(len(head)); sent != announced {
		t.Errorf("the client sent %d of the %d bytes in the file, want the announced %d",
			sent, len(content), announced)
	}
}

func TestPutFollowingRedirects_ClosingTheBodyReleasesTheSocket(t *testing.T) {
	owner := startXAPI(t, xapiAccepts)

	var dialled []net.Conn
	resp, err := putFollowingRedirects(
		func(host string) (net.Conn, error) {
			conn, err := net.Dial("tcp", owner.addr)
			if err == nil {
				dialled = append(dialled, conn)
			}
			return conn, err
		},
		"https://192.0.2.134/import_raw_vdi?session_id="+testSessionRef,
		mustWriteImage(t), int64(len("disk image contents")), packer.TestUi(t))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(dialled) != 1 {
		t.Fatalf("dialled %d times, want 1", len(dialled))
	}
	if err := resp.Body.Close(); err != nil {
		t.Errorf("closing the body reported %s", err)
	}
	if err := dialled[0].SetDeadline(time.Now()); !errors.Is(err, net.ErrClosed) {
		t.Errorf("the socket is still open after the body was closed: %v", err)
	}
	if err := resp.Body.Close(); err != nil {
		t.Errorf("closing the body twice reported %s", err)
	}
}

func mustWriteImage(t *testing.T) string {
	t.Helper()
	path, _ := writeImage(t)
	return path
}
