package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/rdobson/packer-builder-xenserver/builder/xenserver/iso"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(iso.Builder))
	server.Serve()
}
