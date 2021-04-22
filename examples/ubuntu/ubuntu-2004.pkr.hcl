packer {
  required_plugins {
   xenserver= {
      version = ">= v0.3.2"
      source = "github.com/ddelnano/xenserver"
    }
  }
}

variable "remote_host" {
  type        = string
  description = "The ip or fqdn of your XenServer. This will be pulled from the env var 'PKR_VAR_XAPI_HOST'"
  sensitive   = true
  default     = null
}

variable "remote_password" {
  type        = string
  description = "The password used to interact with your XenServer. This will be pulled from the env var 'PKR_VAR_XAPI_PASSWORD'"
  sensitive   = true
  default     = null
}

variable "remote_username" {
  type        = string
  description = "The username used to interact with your XenServer. This will be pulled from the env var 'PKR_VAR_XAPI_USERNAME'"
  sensitive   = true
  default     = null

}

variable "sr_iso_name" {
  type        = string
  default     = ""
  description = "The ISO-SR to packer will use"

}

variable "sr_name" {
  type        = string
  default     = ""
  description = "The name of the SR to packer will use"
}

locals {
  timestamp = regex_replace(timestamp(), "[- TZ:]", "") 
}


source "xenserver-iso" "ubuntu-2004" {
  iso_checksum      = "d1f2bf834bbe9bb43faf16f9be992a6f3935e65be0edece1dee2aa6eb1767423"
  iso_checksum_type = "sha256"
  iso_url           = "http://releases.ubuntu.com/20.04/ubuntu-20.04.2-live-server-amd64.iso"

  sr_iso_name    = var.sr_iso_name
  sr_name        = var.sr_name
  tools_iso_name = "guest-tools.iso"

  remote_host     = var.remote_host
  remote_password = var.remote_password
  remote_username = var.remote_username

  vm_name        = "packer-ubuntu-2004-${local.timestamp}"
  vm_description = "Build started: ${local.timestamp}"
  vm_memory      = 4096
  disk_size      = 20000

  http_directory = "examples/http/ubuntu-2004"
  boot_command   = [
    "<esc><f6> autoinstall ds=nocloud-net;s=http://{{ .HTTPIP }}:{{ .HTTPPort }}/<enter><wait>",
    "<f6><wait><esc><wait> autoinstall ds=nocloud-net;s=http://{{ .HTTPIP }}:{{ .HTTPPort }}/<enter><wait>"
  ]
  boot_wait      = "10s"

  ssh_username            = "testuser"
  ssh_password            = "ubuntu"
  ssh_wait_timeout        = "60000s"
  ssh_handshake_attempts  = 10000

  output_directory = "packer-ubuntu-2004-iso"
  keep_vm          = "always"
}

build {
  sources = ["xenserver-iso.ubuntu-2004"]
}