// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"net"
	"testing"
)

func listeningHost(t *testing.T, l net.Listener) string {
	t.Helper()
	host, _, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		t.Fatalf("could not read the listening address %q: %v", l.Addr(), err)
	}
	return host
}

func TestFindPort_BindsToTheRequestedAddress(t *testing.T) {
	l, port := FindPort("127.0.0.1", 8000, 9000)

	if l == nil {
		t.Fatal("no port was available in 8000-9000")
	}
	defer l.Close()

	tcpAddr, ok := l.Addr().(*net.TCPAddr)
	if !ok || uint(tcpAddr.Port) != port {
		t.Errorf("port = %d, but the listener is on %v", port, l.Addr())
	}
	if host := listeningHost(t, l); host != "127.0.0.1" {
		t.Errorf("listening on %q, want 127.0.0.1", host)
	}
}

func TestFindPort_EmptyAddressListensOnEveryInterface(t *testing.T) {
	l, _ := FindPort("", 8000, 9000)

	if l == nil {
		t.Fatal("no port was available in 8000-9000")
	}
	defer l.Close()

	host := listeningHost(t, l)
	ip := net.ParseIP(host)
	if ip == nil || !ip.IsUnspecified() {
		t.Errorf("listening on %q, want the unspecified address", host)
	}
}

func TestFindPort_ReturnsNothingWhenTheRangeIsExhausted(t *testing.T) {
	busy, port := FindPort("127.0.0.1", 8000, 9000)
	if busy == nil {
		t.Fatal("no port was available in 8000-9000")
	}
	defer busy.Close()

	l, got := FindPort("127.0.0.1", port, port)

	if l != nil {
		l.Close()
		t.Error("a range containing only a taken port must not yield a listener")
	}
	if got != 0 {
		t.Errorf("port = %d, want 0 to signal failure", got)
	}
}
