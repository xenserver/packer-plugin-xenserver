# XCP-ng packer.io builder

This builder plugin extends packer.io to support building images for XCP-ng.

This is a fork of the original builder since the original project was abandoned and no longer compiled with recent versions of Go or worked with Xenserver 7.6 and later.

It improves the original project in the following ways:
1. Developed alongside the [Xenorchestra terraform provider](https://github.com/vatesfr/terraform-provider-xenorchestra) to ensure the hashicorp ecosystem is interoperable.
2. Reimplements how the boot commands are sent over VNC to be compatible with later versions of Xenserver (Citrix hypervisor) and XCP-ng

## Status

At the time of this writing the packer builder has been verified to work with Xenserver 7.6 and can launch VMs with the packer output through the xenorchestra terraform provider.

The following list contains things that are incomplete but will be worked on soon:

- The documentation is still in an inconsistent state with upstream
- XVA builder is untested
- Lots of dead code to remove from upstream

## Using the builder

The packer builder can be installed via `packer init` as long as the packer template includes the following in it's `pkr.hcl` file
```
packer {
  required_plugins {
   xenserver= {
      version = ">= v0.6.0"
      source = "github.com/vatesfr/packer-plugin-xenserver"
    }
  }
}
```

The following command will install the packer plugin using the Ubuntu example provided in this repository.

```
packer init examples/ubuntu/ubuntu-2004.pkr.hcl
```

If you are using an older version of packer or are still using json templates you will need to download the relevant release from the project's [releases page](https://github.com/vatesfr/packer-builder-xenserver/releases) and copy the binary to `~/.packer.d/plugins/packer-builder-xenserver-iso`.

## Developing the builder

### Dependencies
* Packer >= v1.7.1 (https://packer.io)
* XCP-ng / Citrix Hypervisor > 7.6
* Golang 1.20

## Compile the plugin

Once you have installed Packer, you must compile this plugin and install the
resulting binary.

Documentation for Plugins directory: [Official Docs](https://developer.hashicorp.com/packer/docs/configure#packer-s-plugin-directory)

### Linux/MacOS

```shell
go build -o packer-plugin-xenserver

# Add the plugin to the location packer expects it to be installed in
mkdir -p ~/.packer.d/plugins/
cp packer-plugin-xenserver  ~/.packer.d/plugins
```

### Windows (Powershell)

```powershell
go build -o packer-plugin-xenserver

mkdir "%APPDATA%\packer.d\plugins"
cp packer-plugin-xenserver  "%APPDATA%\packer.d\plugins"
```

# Documentation

For complete documentation on configuration commands, see [the
xenserver-iso docs](docs/builders/iso/xenserver-iso.html.markdown)

## Support

You can discuss any issues you have or feature requests in [Discord](https://discord.gg/ZpNq8ez).

If you'd like to support my effort on the project, please consider buying me a coffee

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/ddelnano)
