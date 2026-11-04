// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

const rfbBanner = "RFB 003.008\n"

// xapiConsoleResponse matches xapi's Http.http_200_ok plus the blank line Http_svr.headers adds.
const xapiConsoleResponse = "HTTP/1.1 200 OK\r\n" +
	"Connection: keep-alive\r\n" +
	"Cache-Control: no-cache, no-store\r\n" +
	"\r\n"

type chunkedConn struct {
	chunks  [][]byte
	written bytes.Buffer
}

func newChunkedConn(chunks ...string) *chunkedConn {
	c := &chunkedConn{}
	for _, chunk := range chunks {
		c.chunks = append(c.chunks, []byte(chunk))
	}
	return c
}

func (c *chunkedConn) Read(p []byte) (int, error) {
	if len(c.chunks) == 0 {
		return 0, io.EOF
	}

	n := copy(p, c.chunks[0])
	c.chunks[0] = c.chunks[0][n:]
	if len(c.chunks[0]) == 0 {
		c.chunks = c.chunks[1:]
	}

	return n, nil
}

func (c *chunkedConn) Write(p []byte) (int, error) {
	return c.written.Write(p)
}

func readBanner(t *testing.T, br *bufio.Reader) string {
	t.Helper()

	banner := make([]byte, len(rfbBanner))
	if _, err := io.ReadFull(br, banner); err != nil {
		t.Fatalf("reading the banner: %v", err)
	}

	return string(banner)
}

func TestConnectConsoleProxy_KeepsBannerWhenCoalesced(t *testing.T) {
	conn := newChunkedConn(xapiConsoleResponse + rfbBanner)

	br, statusLine, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s")
	if err != nil {
		t.Fatalf("connectConsoleProxy returned %v, want nil", err)
	}
	if statusLine != "HTTP/1.1 200 OK" {
		t.Errorf("statusLine = %q, want %q", statusLine, "HTTP/1.1 200 OK")
	}
	if got := readBanner(t, br); got != rfbBanner {
		t.Errorf("banner = %q, want %q", got, rfbBanner)
	}
}

func TestConnectConsoleProxy_KeepsBannerWhenSeparate(t *testing.T) {
	conn := newChunkedConn(xapiConsoleResponse, rfbBanner)

	br, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s")
	if err != nil {
		t.Fatalf("connectConsoleProxy returned %v, want nil", err)
	}
	if got := readBanner(t, br); got != rfbBanner {
		t.Errorf("banner = %q, want %q", got, rfbBanner)
	}
}

func TestConnectConsoleProxy_KeepsBannerWhenHeaderSplit(t *testing.T) {
	split := len(xapiConsoleResponse) - 10
	conn := newChunkedConn(xapiConsoleResponse[:split], xapiConsoleResponse[split:]+rfbBanner)

	br, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s")
	if err != nil {
		t.Fatalf("connectConsoleProxy returned %v, want nil", err)
	}
	if got := readBanner(t, br); got != rfbBanner {
		t.Errorf("banner = %q, want %q", got, rfbBanner)
	}
}

// xapi's Http_svr.headers writes the header lines and the terminating blank line separately.
func TestConnectConsoleProxy_KeepsBannerWhenTerminatorArrivesSeparately(t *testing.T) {
	headerLines := "HTTP/1.1 200 OK\r\n" +
		"Connection: keep-alive\r\n" +
		"Cache-Control: no-cache, no-store\r\n"
	conn := newChunkedConn(headerLines, "\r\n"+rfbBanner)

	br, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s")
	if err != nil {
		t.Fatalf("connectConsoleProxy returned %v, want nil", err)
	}
	if got := readBanner(t, br); got != rfbBanner {
		t.Errorf("banner = %q, want %q", got, rfbBanner)
	}
}

func TestConnectConsoleProxy_SendsConnectRequest(t *testing.T) {
	conn := newChunkedConn(xapiConsoleResponse + rfbBanner)

	if _, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:sess"); err != nil {
		t.Fatalf("connectConsoleProxy returned %v, want nil", err)
	}

	want := "CONNECT /console?uuid=abc HTTP/1.0\r\n" +
		"Host: host.example\r\n" +
		"Cookie: session_id=OpaqueRef:sess\r\n" +
		"\r\n"
	if got := conn.written.String(); got != want {
		t.Errorf("request = %q, want %q", got, want)
	}
}

func TestConnectConsoleProxy_RejectsNon2xx(t *testing.T) {
	conn := newChunkedConn("HTTP/1.1 403 Forbidden\r\n\r\nnot a banner")

	_, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s")
	if err == nil {
		t.Fatal("connectConsoleProxy returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "403 Forbidden") {
		t.Errorf("error = %q, want it to name the status", err)
	}
}

func TestConnectConsoleProxy_RejectsNonHTTP(t *testing.T) {
	conn := newChunkedConn(rfbBanner + "\r\n")

	if _, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s"); err == nil {
		t.Fatal("connectConsoleProxy returned nil, want an error")
	}
}

func TestConnectConsoleProxy_ReportsTruncatedHeader(t *testing.T) {
	conn := newChunkedConn("HTTP/1.1 200 OK\r\nConnection: keep-alive\r\n")

	if _, _, err := connectConsoleProxy(conn, "host.example", "/console?uuid=abc", "OpaqueRef:s"); err == nil {
		t.Fatal("connectConsoleProxy returned nil, want an error")
	}
}

func TestParseHTTPStatusCode(t *testing.T) {
	cases := []struct {
		name       string
		statusLine string
		want       int
		wantErr    bool
	}{
		{name: "ok", statusLine: "HTTP/1.1 200 OK", want: 200},
		{name: "no reason phrase", statusLine: "HTTP/1.0 200", want: 200},
		{name: "forbidden", statusLine: "HTTP/1.1 403 Forbidden", want: 403},
		{name: "not http", statusLine: "RFB 003.008", wantErr: true},
		{name: "empty", statusLine: "", wantErr: true},
		{name: "unreadable code", statusLine: "HTTP/1.1 OK", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseHTTPStatusCode(tc.statusLine)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("parseHTTPStatusCode(%q) returned nil, want an error", tc.statusLine)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseHTTPStatusCode(%q) returned %v, want nil", tc.statusLine, err)
			}
			if got != tc.want {
				t.Errorf("parseHTTPStatusCode(%q) = %d, want %d", tc.statusLine, got, tc.want)
			}
		})
	}
}

func TestBufferedConn_DeliversBufferedBytesFirst(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	conn := &bufferedConn{Conn: client, br: bufio.NewReader(strings.NewReader(rfbBanner))}

	got := make([]byte, len(rfbBanner))
	if _, err := io.ReadFull(conn, got); err != nil {
		t.Fatalf("reading the banner: %v", err)
	}
	if string(got) != rfbBanner {
		t.Errorf("banner = %q, want %q", got, rfbBanner)
	}
}

func TestBufferedConn_WriteGivesUpOnAConsoleThatNeverDrains(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close()
	defer server.Close()

	conn := &bufferedConn{
		Conn:         client,
		br:           bufio.NewReader(strings.NewReader("")),
		writeTimeout: 50 * time.Millisecond,
	}

	done := make(chan error, 1)
	go func() {
		_, err := conn.Write([]byte("keystroke"))
		done <- err
	}()

	select {
	case err := <-done:
		if !errors.Is(err, os.ErrDeadlineExceeded) {
			t.Fatalf("write error = %v, want the deadline to have cut it short", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("the write never came back: a keystroke can park on the console forever")
	}
}

type deadlineRecorder struct {
	net.Conn
	mu        sync.Mutex
	deadlines []time.Time
	closed    bool
}

func (c *deadlineRecorder) SetWriteDeadline(t time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.deadlines = append(c.deadlines, t)
	return nil
}

func (c *deadlineRecorder) Write(p []byte) (int, error) { return len(p), nil }

func (c *deadlineRecorder) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.closed = true
	return nil
}

func (c *deadlineRecorder) wasClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

func (c *deadlineRecorder) deadlineCalls() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.deadlines)
}

func TestBufferedConn_WriteLeavesTheDeadlineToCryptoTLS(t *testing.T) {
	rec := &deadlineRecorder{}
	conn := &bufferedConn{Conn: rec, br: bufio.NewReader(strings.NewReader("")), writeTimeout: time.Minute}

	if _, err := conn.Write([]byte("k")); err != nil {
		t.Fatalf("write: %v", err)
	}

	if n := rec.deadlineCalls(); n != 0 {
		t.Errorf("Write set the write deadline %d times, want 0: the deadline is crypto/tls's", n)
	}
}

func TestWithCloseWatchdog_LeavesACallThatReturnedAlone(t *testing.T) {
	rec := &deadlineRecorder{}

	if withCloseWatchdog(rec, 200*time.Millisecond, func() {}) {
		t.Error("the watchdog reported firing on a call that returned at once")
	}
	time.Sleep(400 * time.Millisecond)

	if rec.wasClosed() {
		t.Error("the watchdog closed a connection whose call had already returned")
	}
}

func TestBufferedConn_UnsetTimeoutTakesTheDefault(t *testing.T) {
	if got := (&bufferedConn{}).writeBound(); got != consoleWriteTimeout {
		t.Errorf("writeBound() = %s with nothing set, want the %s default", got, consoleWriteTimeout)
	}
	if got := (&bufferedConn{writeTimeout: time.Second}).writeBound(); got != time.Second {
		t.Errorf("writeBound() = %s, want the 1s it was given", got)
	}
	// Pinned to a literal: every other assertion reads the constant and would miss it drifting.
	if consoleWriteTimeout != 5*time.Minute {
		t.Errorf("consoleWriteTimeout = %s, want 5m", consoleWriteTimeout)
	}
}

func TestHandshakeVNC_TimesOutOnSilentProxy(t *testing.T) {
	client, server := net.Pipe()
	defer server.Close()

	_, err := handshakeVNC(client, 50*time.Millisecond)
	if err == nil {
		t.Fatal("handshakeVNC returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "timed out") {
		t.Errorf("error = %q, want a timeout", err)
	}
}

func TestRetryConsoleConnect_SucceedsFirstTime(t *testing.T) {
	calls := 0

	err := retryConsoleConnect(freshState(), packer.TestUi(t), 5, 0, func() error {
		calls++
		return nil
	})
	if err != nil {
		t.Fatalf("retryConsoleConnect returned %v, want nil", err)
	}
	if calls != 1 {
		t.Errorf("attempt called %d times, want 1", calls)
	}
}

func TestRetryConsoleConnect_RecoversFromTransientFailure(t *testing.T) {
	calls := 0

	err := retryConsoleConnect(freshState(), packer.TestUi(t), 5, 0, func() error {
		calls++
		if calls < 3 {
			return errors.New("error parsing ProtocolVersion")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("retryConsoleConnect returned %v, want nil", err)
	}
	if calls != 3 {
		t.Errorf("attempt called %d times, want 3", calls)
	}
}

func TestRetryConsoleConnect_ReportsLastErrorAfterAllAttempts(t *testing.T) {
	calls := 0

	err := retryConsoleConnect(freshState(), packer.TestUi(t), 3, 0, func() error {
		calls++
		return fmt.Errorf("attempt %d failed", calls)
	})
	if err == nil {
		t.Fatal("retryConsoleConnect returned nil, want an error")
	}
	if calls != 3 {
		t.Errorf("attempt called %d times, want 3", calls)
	}
	if !strings.Contains(err.Error(), "attempt 3 failed") {
		t.Errorf("error = %q, want it to carry the last failure", err)
	}
	if !strings.Contains(err.Error(), "3 attempts") {
		t.Errorf("error = %q, want it to say how many attempts ran", err)
	}
}

func TestRetryConsoleConnect_StopsOnPermanentError(t *testing.T) {
	calls := 0

	err := retryConsoleConnect(freshState(), packer.TestUi(t), 5, 0, func() error {
		calls++
		return fmt.Errorf("wrapped: %w", permanentError{errors.New("certificate signed by unknown authority")})
	})
	if err == nil {
		t.Fatal("retryConsoleConnect returned nil, want an error")
	}
	if calls != 1 {
		t.Errorf("attempt called %d times, want 1", calls)
	}
	if !strings.Contains(err.Error(), "unknown authority") {
		t.Errorf("error = %q, want the original reason", err)
	}
}

func TestRetryConsoleConnect_StopsWhenCancelled(t *testing.T) {
	state := freshState()
	state.Put(multistep.StateCancelled, true)
	calls := 0

	err := retryConsoleConnect(state, packer.TestUi(t), 5, 0, func() error {
		calls++
		return nil
	})
	if _, ok := err.(InterruptedError); !ok {
		t.Fatalf("retryConsoleConnect returned %v, want InterruptedError", err)
	}
	if calls != 0 {
		t.Errorf("attempt called %d times, want 0", calls)
	}
}
