// Copyright (c) 2026. Citrix Systems, Inc.

package main

import (
	"fmt"
	"os"

	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/clone"
	"github.com/xenserver/packer-plugin-xenserver/builder/xenserver/iso"
	"github.com/xenserver/packer-plugin-xenserver/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func pluginSet() *plugin.Set {
	pps := plugin.NewSet()

	pps.RegisterBuilder("vm-iso", new(iso.Builder))
	pps.RegisterBuilder("vm-clone", new(clone.Builder))

	pps.SetVersion(version.PluginVersion)
	return pps
}

func main() {
	err := pluginSet().Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
