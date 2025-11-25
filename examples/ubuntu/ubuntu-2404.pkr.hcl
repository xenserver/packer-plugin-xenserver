packer {
  required_plugins {
   xenserver= {
      version = ">= v0.7.4"
      source = "github.com/vatesfr/packer-plugin-xenserver"
    }
  }
}

# The ubuntu_version value determines what Ubuntu iso URL and sha256 hash we lookup. Updating
# this will allow a new version to be pulled in.
data "null" "ubuntu_version" {
  input = "24.04"
}

locals {
  timestamp = regex_replace(timestamp(), "[- TZ:]", "")
  ubuntu_version = data.null.ubuntu_version.output

  # Update this map depending on what templates are available on your Xen server.
  ubuntu_template_name = {
    24.04 = "Ubuntu Focal Fossa 20.04"
  }
}

# TODO(ddelnano): Update this to use a local once https://github.com/hashicorp/packer/issues/11011
# is fixed.
data "http" "ubuntu_sha_and_release" {
  url = "https://releases.ubuntu.com/${data.null.ubuntu_version.output}/SHA256SUMS"
}

local "ubuntu_sha256" {
  expression = regex("([A-Za-z0-9]+)[\\s\\*]+ubuntu-.*server", data.http.ubuntu_sha_and_release.body)
}

variable "remote_host" {
  type        = string
  description = "The ip or fqdn of your XenServer. This will be pulled from the env var 'PKR_VAR_remote_host'"
  sensitive   = true
  default     = null
}

variable "remote_password" {
  type        = string
  description = "The password used to interact with your XenServer. This will be pulled from the env var 'PKR_VAR_remote_password'"
  sensitive   = true
  default     = null
}

variable "remote_username" {
  type        = string
  description = "The username used to interact with your XenServer. This will be pulled from the env var 'PKR_VAR_remote_username'"
  sensitive   = true
  default     = null
}

variable "sr_iso_name" {
  type        = string
  description = "The name of the SR packer will use to store the installation ISO"
  default     = ""
}

variable "sr_name" {
  type        = string
  description = "The name of the SR packer will use to create the VM"
  default     = ""
}

source "xenserver-iso" "ubuntu-2404" {
  iso_checksum    = "sha256:${local.ubuntu_sha256.0}"
  iso_url         = "https://releases.ubuntu.com/${local.ubuntu_version}/ubuntu-${local.ubuntu_version}-live-server-amd64.iso"

  sr_name         = var.sr_name
  sr_iso_name     = var.sr_iso_name
  remote_host     = var.remote_host
  remote_password = var.remote_password
  remote_username = var.remote_username

  clone_template  = local.ubuntu_template_name[data.null.ubuntu_version.output]
  vm_name         = "ubuntu-${data.null.ubuntu_version.output}-gold"
  vm_description  = "Built at ${local.timestamp}"
  vm_memory       = 8192
  disk_size       = 10240
  vcpus_max       = 4
  vcpus_atstartup = 4

  boot_command =  [
    "c<wait5>",
    "set gfxpayload=keep<enter>",
    "linux /casper/vmlinuz autoinstall ---<enter>",
    "initrd /casper/initrd<enter>",
    "boot<enter>"
  ]
  cd_files = [
    "examples/http/ubuntu-2404/meta-data",
    "examples/http/ubuntu-2404/user-data"
  ]

  # The xenserver plugin needs to SSH in to the new VM, so we give it
  # the information to do so
  ssh_username            = "testuser"
  ssh_password            = "ubuntu"
  ssh_wait_timeout        = "60000s"
  ssh_handshake_attempts  = 10000


  output_directory = null
  keep_vm          = "always"
}

build {
  sources = ["xenserver-iso.ubuntu-2404"]

  # Things to do on the new VM once it's past first reboot:

  # Wait for cloud-init to finish everything. We need to do this as a
  # packer provisioner to prevent packer-plugin-xenserver from shutting
  # the VM down before all cloud-init processing is complete.
  provisioner "shell" {
    inline = [
      "echo ubuntu | sudo -S cloud-init status --wait"
    ]
    valid_exit_codes = [0, 2]
  }
}
