// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// isWildcardBindAddress parses the address, so "::0" and "::ffff:0.0.0.0" are wildcards too.
func isWildcardBindAddress(bindAddress string) bool {
	ip := net.ParseIP(bindAddress)
	return ip == nil || ip.IsUnspecified()
}

type httpIPResolver struct {
	// routeTo reports the local address the kernel would send from when talking to host.
	routeTo func(host string) (string, error)
	// addrOfInterface reports the first usable IPv4 address assigned to a named interface.
	addrOfInterface func(name string) (string, error)
}

func newHTTPIPResolver() httpIPResolver {
	return httpIPResolver{
		routeTo:         localAddrForHost,
		addrOfInterface: addrOfInterface,
	}
}

func (r httpIPResolver) resolve(bindAddress, interfaceName, poolAddress string) (string, error) {
	if !isWildcardBindAddress(bindAddress) {
		return bindAddress, nil
	}

	if interfaceName != "" {
		ip, err := r.addrOfInterface(interfaceName)
		if err != nil {
			return "", fmt.Errorf("cannot use http_interface %q: %w", interfaceName, err)
		}
		return ip, nil
	}

	ip, err := r.routeTo(poolAddress)
	if err != nil {
		return "", fmt.Errorf("could not work out which local address reaches the XenServer pool at %q: %w; "+
			"set http_bind_address to the address the guest should fetch from", poolAddress, err)
	}
	return ip, nil
}

// localAddrForHost dials UDP, which picks a source address without sending; any port works.
func localAddrForHost(host string) (string, error) {
	host = bareHost(host)
	if host == "" {
		return "", errors.New("no XenServer pool address to route towards")
	}

	conn, err := net.Dial("udp", net.JoinHostPort(host, "80"))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	addr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok || addr.IP == nil || addr.IP.IsUnspecified() {
		return "", fmt.Errorf("the route to %q has no local source address", host)
	}
	return addr.IP.String(), nil
}

func addrOfInterface(name string) (string, error) {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}
		if ipv4 := ipNet.IP.To4(); ipv4 != nil {
			return ipv4.String(), nil
		}
	}

	return "", fmt.Errorf("interface %q has no non-loopback IPv4 address", name)
}

// stripScheme matches any "scheme://" because schemes are case-insensitive ("HTTPS://pool").
func stripScheme(host string) string {
	if i := strings.Index(host, "://"); i >= 0 {
		return host[i+len("://"):]
	}
	return host
}

func bareHost(host string) string {
	host = strings.TrimSuffix(stripScheme(host), "/")

	// SplitHostPort fails on a bare IPv6 literal, which is then returned untouched.
	if stripped, _, err := net.SplitHostPort(host); err == nil {
		return stripped
	}

	// JoinHostPort would double the brackets of a bracketed IPv6 literal with no port.
	if len(host) > 1 && host[0] == '[' && host[len(host)-1] == ']' {
		return host[1 : len(host)-1]
	}

	return host
}
