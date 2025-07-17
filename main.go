package main

import (
	"fmt"
	"os"

	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/clone"
	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/iso"
	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/xva"
	"github.com/xenserver/packer-plugin-xenserver/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()

	// Add builders to the plugin set
	pps.RegisterBuilder("xva", new(xva.Builder))
	pps.RegisterBuilder("iso", new(iso.Builder))
	pps.RegisterBuilder("clone", new(clone.Builder))

	// Set the plugin version
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
