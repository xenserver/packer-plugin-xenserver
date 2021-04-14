# This is for packer to find the xenserver plugin
# so that you no longer need to install it by hand
packer {
  required_plugins {
   xenserver= {
      version = ">= v0.3.2"
      source = "github.com/ddelnano/xenserver"
    }
  }
}


source "xenserver-iso" "example" {
  #
  # Where to get the iso
  #
  iso_url           = "http://mirrors.ocf.berkeley.edu/centos/8.3.2011/isos/x86_64/CentOS-8.3.2011-x86_64-dvd1.iso"
  iso_checksum_type = "sha1"
  iso_checksum      = "aaf9d4b3071c16dbbda01dfe06085e5d0fdac76df323e3bbe87cce4318052247"
  #
  # Where to store the ISO, vm-disk and where to find the xentools iso
  #
  sr_iso_name    = "Local-ISO"
  sr_name        = "Local-SR"
  tools_iso_name = "guest-tools.iso"
  #
  # How to communicate with the xenserver
  #
  remote_host     = "xenserver.example.org"
  remote_username = "root"
  remote_password = "very-secret-password"
  #
  # Basic info for the vm 
  #
  vm_name        = "packer-centos8-example"
  vm_description = "This is an example."
  vm_memory      = 4096
  disk_size      = 4096
  #
  # For the installation
  #
  http_directory = "examples/http/centos8"
  boot_command   = ["<tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks-centos8-examples.cfg<enter><wait>"]
  boot_wait      = "10s"
  #
  # how packer contacts the vm  
  #
  ssh_username     = "root"
  ssh_password     = "centos"
  ssh_wait_timeout = "10m"
  shutdown_command = "/sbin/shutdown"
  #
  # What to do with the resulting VM
  #
  output_directory = "packer-centos8-local"
  keep_vm          = "on_success"

}
#
# Operations to do while building.
# Any provisioning would be done here
#
build {
  sources = ["xenserver-iso.example"]
}