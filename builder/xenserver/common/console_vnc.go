// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/mitchellh/go-vnc"
)

const (
	// The console proxy answers before the guest's VNC server is up, so early attempts can fail transiently.
	consoleConnectAttempts = 5
	consoleConnectBackoff  = 3 * time.Second

	// Bounds a proxy that accepts the connection and then sends nothing.
	consoleHandshakeTimeout = 30 * time.Second

	// consoleWriteTimeout only catches a socket that never drains: a timed-out write ends the build.
	consoleWriteTimeout = 5 * time.Minute
)

// permanentError marks a failure that retrying cannot fix.
type permanentError struct{ err error }

func (e permanentError) Error() string { return e.err.Error() }

func (e permanentError) Unwrap() error { return e.err }

// withCloseWatchdog closes conn if fn outlasts timeout and reports whether it did. It does
// not set a deadline: crypto/tls sets its own around the closeNotify alert in Close.
func withCloseWatchdog(conn net.Conn, timeout time.Duration, fn func()) bool {
	var mu sync.Mutex
	done := false
	expired := false

	watchdog := time.AfterFunc(timeout, func() {
		mu.Lock()
		defer mu.Unlock()
		if done {
			return
		}
		expired = true
		conn.Close()
	})

	fn()

	mu.Lock()
	done = true
	fired := expired
	mu.Unlock()
	watchdog.Stop()

	return fired
}

// bufferedConn reads from br, which may already hold RFB bytes read along with the CONNECT response.
type bufferedConn struct {
	net.Conn
	br           *bufio.Reader
	writeTimeout time.Duration
}

func (c *bufferedConn) Read(p []byte) (int, error) {
	return c.br.Read(p)
}

func (c *bufferedConn) writeBound() time.Duration {
	if c.writeTimeout <= 0 {
		return consoleWriteTimeout
	}
	return c.writeTimeout
}

func (c *bufferedConn) Write(p []byte) (int, error) {
	bound := c.writeBound()

	var n int
	var err error
	if withCloseWatchdog(c.Conn, bound, func() { n, err = c.Conn.Write(p) }) {
		return n, fmt.Errorf("the VM console accepted no data for %s: %w", bound, os.ErrDeadlineExceeded)
	}

	return n, err
}

func parseHTTPStatusCode(statusLine string) (int, error) {
	fields := strings.Fields(statusLine)
	if len(fields) < 2 || !strings.HasPrefix(fields[0], "HTTP/") {
		return 0, fmt.Errorf("the xapi console proxy did not answer with an HTTP status line: %q", statusLine)
	}

	code, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, fmt.Errorf("the xapi console proxy returned an unreadable HTTP status: %q", statusLine)
	}

	return code, nil
}

func readHTTPHeader(br *bufio.Reader) (string, error) {
	tp := textproto.NewReader(br)

	statusLine, err := tp.ReadLine()
	if err != nil {
		return "", fmt.Errorf("failed to read the console CONNECT response: %w", err)
	}

	for {
		line, err := tp.ReadLine()
		if err != nil {
			return statusLine, fmt.Errorf("failed to read the console CONNECT response header: %w", err)
		}
		if line == "" {
			return statusLine, nil
		}
	}
}

// connectConsoleProxy reads the CONNECT response a line at a time and returns the buffered
// reader: xapi sends the RFB banner right after the header, often in the same TLS record.
func connectConsoleProxy(conn io.ReadWriter, consoleHost, consoleLocation, sessionRef string) (*bufio.Reader, string, error) {
	// The request carries the session reference in its Cookie header, so never log it.
	request := fmt.Sprintf("CONNECT %s HTTP/1.0\r\nHost: %s\r\nCookie: session_id=%s\r\n\r\n",
		consoleLocation, consoleHost, sessionRef)

	if _, err := io.WriteString(conn, request); err != nil {
		return nil, "", fmt.Errorf("failed to send the console CONNECT request: %w", err)
	}

	br := bufio.NewReader(conn)
	statusLine, err := readHTTPHeader(br)
	if err != nil {
		return nil, statusLine, err
	}

	code, err := parseHTTPStatusCode(statusLine)
	if err != nil {
		return nil, statusLine, err
	}
	if code < 200 || code > 299 {
		return nil, statusLine, fmt.Errorf("the xapi console proxy refused the connection: %s", statusLine)
	}

	return br, statusLine, nil
}

// handshakeVNC avoids a read deadline, which would later expire under vnc.Client's reader goroutine.
func handshakeVNC(conn net.Conn, timeout time.Duration) (*vnc.ClientConn, error) {
	var vncClient *vnc.ClientConn
	var err error

	timedOut := withCloseWatchdog(conn, timeout, func() {
		vncClient, err = vnc.Client(conn, &vnc.ClientConfig{Exclusive: true})
	})

	if timedOut {
		// vnc.Client may have succeeded just before the watchdog closed the connection.
		if err == nil {
			vncClient.Close()
		}
		return nil, fmt.Errorf("timed out after %s waiting for the VM console VNC handshake", timeout)
	}
	if err != nil {
		return nil, fmt.Errorf("error establishing VNC session: %w", err)
	}

	return vncClient, nil
}

func dialConsoleVNC(ctx context.Context, ui packer.Ui, c *Connection, consoleHost, consoleLocation string) (*vnc.ClientConn, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(consoleHost, "443"), consoleHandshakeTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the VM console: %w", err)
	}

	connected := false
	defer func() {
		if !connected {
			conn.Close()
		}
	}()

	// A deadline is safe until the VNC client starts its read loop; it is cleared before then.
	if err := conn.SetDeadline(time.Now().Add(consoleHandshakeTimeout)); err != nil {
		return nil, fmt.Errorf("failed to set a deadline on the VM console connection: %w", err)
	}

	// tls.Client needs ServerName set, and the console may be on a pool member other than xenserver_host.
	tlsConfig := c.TLSConfig()
	tlsConfig.ServerName = consoleHost
	tlsConn := tls.Client(conn, tlsConfig)

	// Handshake now so a certificate problem is reported as one; only a trust failure is permanent.
	if err := tlsConn.HandshakeContext(ctx); err != nil {
		wrapped := describeTLSError(consoleHost,
			fmt.Errorf("failed to establish a TLS connection to the VM console: %w", err))
		var verifyErr *tls.CertificateVerificationError
		if errors.As(err, &verifyErr) {
			return nil, permanentError{wrapped}
		}
		return nil, wrapped
	}

	ui.Say(fmt.Sprintf("Making HTTP request to initiate VNC connection: CONNECT %s", consoleLocation))
	br, statusLine, err := connectConsoleProxy(tlsConn, consoleHost, consoleLocation, string(c.GetSessionRef()))
	if err != nil {
		return nil, err
	}
	ui.Say(fmt.Sprintf("Received response: %s", statusLine))

	if err := conn.SetDeadline(time.Time{}); err != nil {
		return nil, fmt.Errorf("failed to clear the deadline on the VM console connection: %w", err)
	}

	vncClient, err := handshakeVNC(&bufferedConn{Conn: tlsConn, br: br}, consoleHandshakeTimeout)
	if err != nil {
		return nil, err
	}

	connected = true
	return vncClient, nil
}

func retryConsoleConnect(state multistep.StateBag, ui packer.Ui, attempts int, backoff time.Duration, attempt func() error) error {
	var lastErr error

	for i := 1; i <= attempts; i++ {
		if _, cancelled := state.GetOk(multistep.StateCancelled); cancelled {
			return InterruptedError{}
		}

		err := attempt()
		if err == nil {
			return nil
		}

		var permanent permanentError
		if errors.As(err, &permanent) {
			return err
		}
		lastErr = err

		if i == attempts {
			break
		}

		ui.Say(fmt.Sprintf("VM console VNC attempt %d of %d failed (%s). Retrying in %s...",
			i, attempts, err, backoff))
		if waitErr := (InterruptibleWait{Timeout: backoff}).Wait(state); waitErr != nil {
			return waitErr
		}
	}

	return fmt.Errorf("could not establish a VNC session with the VM console after %d attempts: %w",
		attempts, lastErr)
}

func establishConsoleVNC(ctx context.Context, state multistep.StateBag, ui packer.Ui, c *Connection, consoleHost, consoleLocation string) (*vnc.ClientConn, error) {
	var vncClient *vnc.ClientConn

	err := retryConsoleConnect(state, ui, consoleConnectAttempts, consoleConnectBackoff, func() error {
		var attemptErr error
		vncClient, attemptErr = dialConsoleVNC(ctx, ui, c, consoleHost, consoleLocation)
		return attemptErr
	})
	if err != nil {
		return nil, err
	}

	return vncClient, nil
}
