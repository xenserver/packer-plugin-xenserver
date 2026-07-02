# XenServer Packer plugin

This repository provides a Packer plugin for XenServer.

It includes three builders:
- ISO - Create a VM from installation media.
- Clone - Clone an existing template or VM for further customization.
- XVA - Import an existing XVA appliance for further customization.

## Status

At the time of writing, the plugin has been verified to work with XenServer 8.4.

The following list contains things that are incomplete but will be worked on soon:
- The documentation is still evolving.
- XVA builder is untested.
- Some dead code remains to be removed from upstream.

## Using the plugin

The plugin can be installed via `packer init` as long as the Packer template includes the following in its `pkr.hcl` file:
```
packer {
  required_plugins {
    xenserver = {
      version = ">= 0.1.0"
      source = "github.com/xenserver/xenserver"
    }
  }
}
```

The following command will install the packer plugin using the Ubuntu example provided in this repository.

```
packer init examples/ubuntu/ubuntu-2004.pkr.hcl
```

If you are using an older version of Packer or are still using JSON templates, download the relevant release from the project's [releases page](https://github.com/xenserver/packer-plugin-xenserver/releases) and install it with `packer plugins install --path /path/to/packer-plugin-xenserver github.com/xenserver/xenserver`.

## Developing the plugin

### Dependencies
* Packer >= v1.7.1 (https://packer.io)
* XenServer > 8.4
* Golang 1.24.1

## Compile the plugin

Once you have installed Packer, you must compile this plugin and install the
resulting binary.

```shell
# Build the plugin binary
$ make build

# Manual steps:
$ go build -o packer-plugin-xenserver
$ packer plugins install --path packer-plugin-xenserver github.com/xenserver/xenserver
```

# Documentation

For complete documentation on configuration commands, see [the
xenserver-iso docs](docs/builders/xenserver-iso.mdx)

## Special thanks

Thanks @ddelnano for the community support.