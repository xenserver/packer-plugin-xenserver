package main

import (
	"fmt"
	"os"

	"github.com/xenserver/packer-builder-xenserver/builder/xenserver/iso"
	"github.com/xenserver/packer-builder-xenserver/builder/xenserver/xva"
	"github.com/xenserver/packer-builder-xenserver/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("iso", new(iso.Builder))
	pps.RegisterBuilder("xva", new(xva.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
