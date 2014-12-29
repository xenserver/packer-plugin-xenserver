package common

import (
	"fmt"
	"log"
	"net"
)

// FindPort finds and starts listening on a port in the range [portMin, portMax]
// returns the listener and the port number on success, or nil, 0 on failure
func FindPort(portMin uint, portMax uint) (net.Listener, uint) {
	log.Printf("Looking for an available port between %d and %d", portMin, portMax)

	for port := portMin; port <= portMax; port++ {
		log.Printf("Trying port: %d", port)
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			return l, port
		} else {
			log.Printf("Port %d unavailable: %s", port, err.Error())
		}
	}

	return nil, 0
}
