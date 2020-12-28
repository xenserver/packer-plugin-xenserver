# XenServer packer.io builder

This builder plugin extends packer.io to support building images for XenServer. 

This is a fork of the original builder since the original project was abandoned and no longer compilied with recent versions of Go or worked with Xenserver 7.6 and later.

It improves the original project in the following ways:
2. Developed alongside the [Xenorchestra terraform provider](https://github.com/ddelnano/terraform-provider-xenorchestra) to ensure the hashicorp ecosystem is interoperable.
3. Reimplements how the boot commands are sent over VNC to be compatible with later versions of Xenserver (Citrix hypervisor) and XCP

## Status

At the time of this writing the packer builder has been verified to work with Xenserver 7.6 and can launch VMs with the packer output through the xenorchestra terraform provider.

The following list contains things that are incomplete but will be worked on soon:
- The documentation is still in an inconsistent state with upstream
- Examples that are easier for new users to get up and running quickly
- XVA builder is untested
- Lots of dead code to remove from upstream

## Using the builder

Download the relevant release from the project's [releases page](https://github.com/ddelnano/packer-builder-xenserver/releases) and copy the binary to `~/.packer.d/plugin/packer-builder-xenserver-iso`.

## Developing the builder

### Dependencies
* Packer >= v0.10.2 (https://packer.io)
* XenServer / Citrix Hypervisor > 7.6
* Golang 1.14

## Compile the plugin

Once you have installed Packer, you must compile this plugin and install the
resulting binary.

```shell
$ go build github.com/xenserver/packer-builder-xenserver/plugin/builder-xenserver-iso

# Add the builder to the location packer expects it to be installed in
$ mkdir -p ~/.packer.d/plugins/
$ cp builder-xenserver-iso  ~/.packer.d/plugins/packer-builder-xenserver-iso
```

# Documentation

For complete documentation on configuration commands, see [the
xenserver-iso docs](docs/builders/xenserver-iso.html.markdown)
