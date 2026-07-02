The XenServer Packer plugin builds XenServer virtual machine images from ISO installers, existing VMs or templates, and XVA appliances.

### Installation

To install this plugin, copy and paste this code into your Packer configuration, then run [`packer init`](https://developer.hashicorp.com/packer/docs/commands/init).

```hcl
packer {
  required_plugins {
    xenserver = {
      source  = "github.com/xenserver/xenserver"
      version = ">= 0.1.0"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
packer plugins install github.com/xenserver/xenserver
```

### Components

This plugin supports three XenServer builder workflows: creating a VM from installation media, cloning an existing VM or template, and importing an existing XVA appliance.

#### Builders

- [xenserver-iso](/packer/integrations/xenserver/xenserver/latest/components/builder/xenserver-iso) - Creates a new XenServer VM from ISO media, provisions it, and exports the resulting image.
- [xenserver-clone](/packer/integrations/xenserver/xenserver/latest/components/builder/xenserver-clone) - Clones an existing XenServer VM or template, provisions it, and exports the result.
- [xenserver-xva](/packer/integrations/xenserver/xenserver/latest/components/builder/xenserver-xva) - Imports an existing XVA appliance, provisions it, and exports the updated image.

### Examples

Working examples for Ubuntu, CentOS, and Windows live in the [examples directory](https://github.com/xenserver/packer-plugin-xenserver/tree/main/examples).
