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

source "xenserver-iso" "centos8-netinstall" {
  iso_checksum      = "sha1:07a8e59c42cc086ec4c49bdce4fae5a17b077dea"
  iso_url           = "http://mirrors.ocf.berkeley.edu/centos/8.3.2011/isos/x86_64/CentOS-8.3.2011-x86_64-boot.iso"

  sr_iso_name    = var.sr_iso_name
  sr_name        = var.sr_name
  tools_iso_name = "guest-tools.iso"

  remote_host     = var.remote_host
  remote_password = var.remote_password
  remote_username = var.remote_username

  vm_name        = "packer-centos8-netinstall-${local.timestamp}"
  vm_description = "Build started: ${local.timestamp}\n This was installed with an external repository"
  vm_memory      = 4096
  disk_size      = 4096

  http_directory = "examples/http/centos8"
  boot_command   = ["<tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks-centos8-netinstall.cfg<enter><wait>"]
  boot_wait      = "10s"

  ssh_username     = "root"
  ssh_password     = "centos"
  ssh_wait_timeout = "10000s"

  output_directory = "packer-centos8-netinstall"
  keep_vm          = "always"
}

build {
  sources = ["xenserver-iso.centos8-netinstall"]
}
