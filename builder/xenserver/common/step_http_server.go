// Copyright (c) 2026. Citrix Systems, Inc.

package common

// Taken from hashicorp/packer/builder/qemu/step_http_server.go

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// StepHTTPServer serves `http_directory` to the guest and puts "http_port" in the state bag.
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
		}
	}
	snooper.handler.ServeHTTP(resp, req)
}

func (s *StepHTTPServer) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("commonconfig").(CommonConfig)
	ui := state.Get("ui").(packer.Ui)

	var httpPort uint = 0
	if config.HTTPDir == "" {
		// The provision steps assert an int, so this cannot be a uint.
		state.Put("http_port", int(httpPort))
		return multistep.ActionContinue
	}

	s.l, httpPort = FindPort(config.HTTPAddress, HTTPPortMin, HTTPPortMax)

	if s.l == nil || httpPort == 0 {
		err := fmt.Errorf("unable to find a free HTTP server port between %d and %d", HTTPPortMin, HTTPPortMax)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Say(fmt.Sprintf("Starting HTTP server on %s", s.l.Addr()))

	fileServer := http.FileServer(http.Dir(config.HTTPDir))
	server := &http.Server{
		Handler: IPSnooper{
			ch:      s.Chan,
			handler: fileServer,
		},
	}
	go server.Serve(s.l)

	// The provision steps assert an int, so this cannot be a uint.
	state.Put("http_port", int(httpPort))

	return multistep.ActionContinue
}

func (s *StepHTTPServer) Cleanup(multistep.StateBag) {
	if s.l != nil {
		s.l.Close()
	}
}
