# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "XenServer"
  description = "The XenServer Packer plugin builds XenServer virtual machine images from ISO installers, existing VMs or templates, and XVA appliances."
  identifier = "packer/xenserver/xenserver"
  docs {
    process_docs = true
    readme_location = "./README.md"
    external_url = "https://github.com/xenserver/packer-plugin-xenserver"
  }
  license {
    type = "MPL-2.0"
    url = "https://github.com/xenserver/packer-plugin-xenserver/blob/main/LICENSE"
  }
  component {
    type = "builder"
    name = "XenServer ISO"
    slug = "xenserver-iso"
  }
  component {
    type = "builder"
    name = "XenServer Clone"
    slug = "xenserver-clone"
  }
  component {
    type = "builder"
    name = "XenServer XVA"
    slug = "xenserver-xva"
  }
}
