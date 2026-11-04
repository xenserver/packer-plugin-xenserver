// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"fmt"
	"log"
	"net"
)

// HTTPPortMin/HTTPPortMax bound the HTTP server port; the guest must reach the whole range.
const (
	HTTPPortMin uint = 8000
	HTTPPortMax uint = 9000
)

// FindPort listens on a free port in [portMin, portMax]; address "" means all interfaces.
func FindPort(address string, portMin uint, portMax uint) (net.Listener, uint) {
	log.Printf("Looking for an available port between %d and %d on '%s'", portMin, portMax, address)

	for port := portMin; port <= portMax; port++ {
		log.Printf("Trying port: %d", port)
		l, err := net.Listen("tcp", net.JoinHostPort(address, fmt.Sprint(port)))
		if err == nil {
			return l, port
		} else {
			log.Printf("Port %d unavailable: %s", port, err.Error())
		}
	}

	return nil, 0
}
