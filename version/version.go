// Copyright (c) 2026. Citrix Systems, Inc.

package version

import (
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var (
	// Version is the main version number.
	Version = "v0.9.1"

	// VersionPrerelease marks Version as a pre-release, such as "dev"; empty means a final release.
	VersionPrerelease = ""

	VersionMetadata = ""

	// PluginVersion is the version the plugin set reports to Packer.
	PluginVersion = version.NewPluginVersion(Version, VersionPrerelease, VersionMetadata)
)
