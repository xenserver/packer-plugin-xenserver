# XenServer Packer.io builder

This builder plugin extends packer.io to support building images for XenServer.

It contains 3 different builders:
- ISO -> With the ISO builder you can create a VM from scratch by auto-installing the OS.
- Clone -> With the Clone builder you can clone an existing Template / VM for further customization.
- XVA -> With the XVA builder you can import an existing Template / VM for further customization.

## Status

At the time of this writing the packer builder has been verified to work with XenServer 8.4.

The following list contains things that are incomplete but will be worked on soon:
- The documentation is still an on-going process
- XVA builder is untested
- Lots of dead code to remove from upstream

## Using the builder

The packer builder can be installed via `packer init` as long as the packer template includes the following in it's `pkr.hcl` file
```
packer {
  required_plugins {
   xenserver= {
      version = ">= v0.1.0"
      source = "github.com/xenserver/xenserver"
    }
  }
}
```

The following command will install the packer plugin using the Ubuntu example provided in this repository.

```
packer init examples/ubuntu/ubuntu-2004.pkr.hcl
```

If you are using an older version of packer or are still using json templates you will need to download the relevant release from the project's [releases page](https://github.com/ddelnano/packer-builder-xenserver/releases) and copy the binary to `~/.packer.d/plugins/packer-builder-xenserver-iso`.

## Developing the builder

### Dependencies
* Packer >= v1.7.1 (https://packer.io)
* XenServer > 8.4
* Golang 1.24.1

## Compile the plugin

Once you have installed Packer, you must compile this plugin and install the
resulting binary.

```shell
# Just build the Plugin
$ make
# Build the plugin and copy it to ~/.packer.d/plugins/
$ make dev

#Manual steps:
$ go build -o packer-plugin-xenserver

# Add the builder to the location packer expects it to be installed in
$ mkdir -p ~/.packer.d/plugins/
$ cp packer-builder-xenserver-iso  ~/.packer.d/plugins/packer-builder-xenserver-iso
```

# Documentation

For complete documentation on configuration commands, see [the
xenserver-iso docs](docs/builders/xenserver-iso.mdx)

## Special thanks

Thanks @ddelnano for the the community support.