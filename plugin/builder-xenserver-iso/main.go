package main

import (
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/xenserver/packer-builder-xenserver/builder/xenserver/iso"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(iso.Builder))
	server.Serve()
}
