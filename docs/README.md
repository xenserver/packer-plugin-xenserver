<!-- markdownlint-disable first-line-h1 no-inline-html -->

The XenServer plugin lets you build XenServer virtual machines and
templates with [HashiCorp Packer](https://www.packer.io). It connects to the pool
coordinator over the XenAPI — XenCenter is not required — and is tested against
XenServer 9.

### Installation

To install this plugin, copy and paste this block into your Packer configuration,
then run [`packer init`](/packer/docs/commands/init).

```hcl
packer {
  required_plugins {
    xenserver = {
      version = ">= 1.0.0"
      source  = "github.com/xenserver/xenserver"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage the installation of
this plugin.

```sh
packer plugins install github.com/xenserver/xenserver
```

Verify the installed plugins with:

```sh
packer plugins installed
```

### Components

The plugin includes the following builders for creating virtual machine images,
depending on your desired strategy:

#### Builders

- [xenserver-vm-iso](/packer/integrations/xenserver/xenserver/latest/components/builder/xenserver-vm-iso) -
  Starts from an ISO file and uses the XenAPI to build a virtual machine on a
  remote XenServer host or pool, driving an unattended operating system install.

- [xenserver-vm-clone](/packer/integrations/xenserver/xenserver/latest/components/builder/xenserver-vm-clone) -
  Clones an existing template or VM as the starting point, then re-provisions and
  saves it — for patching or customizing a golden image without a full reinstall.

### Upgrading from the earlier XenServer plugin

@include 'xenserver/common/UpgradeIntro.mdx'
#### Builder types

@include 'xenserver/common/UpgradeBuilderTypes.mdx'
#### Renamed options

@include 'xenserver/common/UpgradeRenamedOptions.mdx'
#### Removed options

@include 'xenserver/common/UpgradeRemovedOptions.mdx'
### Compatibility

| Component | Supported |
| --------- | --------- |
| XenServer | 9. Earlier versions and XCP-ng are not supported in this release. |
| Packer    | 1.7.0 and later. |
| Guest OS  | Windows 11 (UEFI + vTPM) and Windows Server 2016 / 2019 / 2022 / 2025, connected to over WinRM. Linux guests are not supported in this release. |
| Platforms | Released plugin binaries for the host running `packer build`: `linux`, `windows`, `darwin` and `freebsd`, each on `amd64` and `arm64`. No 32-bit (`386` / `arm`) binaries are published. |

