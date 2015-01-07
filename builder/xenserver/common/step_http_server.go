package common

// Taken from mitchellh/packer/builder/qemu/step_http_server.go

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
	"net"
	"net/http"
)

// This step creates and runs the HTTP server that is serving files from the
// directory specified by the 'http_directory` configuration parameter in the
// template.
//
// Uses:
//   config *config
//   ui     packer.Ui
//
// Produces:
//   http_port int - The port the HTTP server started on.
type StepHTTPServer struct {
	Chan chan<- string

	l net.Listener
}

type IPSnooper struct {
	ch      chan<- string
	handler http.Handler
}

func (snooper IPSnooper) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Printf("HTTP: %s %s %s", req.RemoteAddr, req.Method, req.URL)
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err == nil && ip != "" {
		select {
		case snooper.ch <- ip:
			log.Printf("Remembering remote address '%s'", ip)
		default:
			// if ch is already full, don't block waiting to send the address, just drop it
		}
	}
	snooper.handler.ServeHTTP(resp, req)
}

func (s *StepHTTPServer) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	var httpPort uint = 0
	if config.HTTPDir == "" {
		state.Put("http_port", httpPort)
		return multistep.ActionContinue
	}

	s.l, httpPort = FindPort(config.HTTPPortMin, config.HTTPPortMax)

	if s.l == nil || httpPort == 0 {
		ui.Error("Error: unable to find free HTTP server port. Try providing a larger range [http_port_min, http_port_max]")
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Starting HTTP server on port %d", httpPort))

	// Start the HTTP server and run it in the background
	fileServer := http.FileServer(http.Dir(config.HTTPDir))
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", httpPort),
		Handler: IPSnooper{
			ch:      s.Chan,
			handler: fileServer,
		},
	}
	go server.Serve(s.l)

	// Save the address into the state so it can be accessed in the future
	state.Put("http_port", httpPort)

	return multistep.ActionContinue
}

func (s *StepHTTPServer) Cleanup(multistep.StateBag) {
	if s.l != nil {
		// Close the listener so that the HTTP server stops
		s.l.Close()
	}
}
