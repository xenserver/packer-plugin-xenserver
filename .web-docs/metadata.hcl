# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# Details on using this Integration template can be found at
# https://github.com/hashicorp/integration-template
# This metadata.hcl file and the adjacent `components` docs directory should be
# kept in a `.web-docs` directory at the root of the plugin repository.
integration {
  name        = "XenServer"
  description = "A plugin for creating XenServer virtual machine images."
  identifier  = "packer/xenserver/xenserver"
  # No hcp-ready flag: the artifact reports no HCP Packer metadata
  # (LocalArtifact.State returns nil).
  docs {
    process_docs    = true
    readme_location = "./README.md"
    external_url    = "https://github.com/xenserver/packer-plugin-xenserver"
  }
  license {
    type = "MPL-2.0"
    url  = "https://github.com/xenserver/packer-plugin-xenserver/blob/main/LICENSE"
  }
  component {
    type = "builder"
    name = "XenServer ISO"
    slug = "xenserver-vm-iso"
  }
  component {
    type = "builder"
    name = "XenServer Clone"
    slug = "xenserver-vm-clone"
  }
}
