
#this tells packer that it needs the `xenserver` plugin to run this script.  
packer {
  required_plugins {
   xenserver= {
      # specifically a version higher than 0.3.2
      version = ">= v0.3.2"
      # which can be found on github
      source = "github.com/ddelnano/xenserver"
    }
  }
}

  
# packer configures an Artifact named "example" using the "xenserver-iso" builder
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
  sr_iso_name    = "Local-ISO" # store the iso used during install on the ISO-SR named `Local-ISO`
  sr_name        = "Local-SR" # store the vmdisk used during install on the SR named `Local-SR`
  tools_iso_name = "guest-tools.iso" # mount the guest-tools iso with the name `guest-tools.iso`
  #
  # How to communicate with the xenserver
  #
  remote_host     = "xenserver.example.org" # the ipadress or fqdn of the xenserver. This should be the pool primary
  remote_username = "root" # the user with which to connect.
  remote_password = "very-secret-password" # the password for the user.
  #
  # Basic info for the vm 
  # 
  vm_name        = "packer-centos8-example" # how packer will name the vm.
  vm_description = "This is an example." # the description field of the vm
  vm_memory      = 4096 # the Amount of RAM, in MB
  disk_size      = 4096  # the Amount of RAM, in MB
  #
  # For the installation
  #
  http_directory = "examples/http/centos8" # Packer will spin up a http-server for serving files to the installing vm. the kickstart file `ks-centos8-examples.cfg` is in this directory
  boot_wait      = "10s" # Wait for 10s after starting the VM before typeing the `boot_command`
  boot_command   = ["<tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks-centos8-examples.cfg<enter><wait>"]
    # The command to type into the installing vm.
    # `<tab>` a tab Character `\t`
    # `text ks=http://` character literals
    # `{{ .HTTPIP }}` will be replaced by the local ip-address.
    # `:` character literal
    # `{{ .HTTPPort }}` will be replaced by the local port. Will be randomly selectet between 8000 and 9000
    # `/ks-centos8-examples.cfg` character literals
    # `<enter>` an enter character
    # `<wait>` wait for 1s

  #
  # how packer contacts the vm  
  #
  ssh_username     = "root" # The ssh user packer uses to connect to the VM
  ssh_password     = "centos" # The ssh password packer uses to connect to the VM
  ssh_wait_timeout = "10m" # consider install failed if unable to connect via ssh 10m into the build
  shutdown_command = "/sbin/shutdown" #  After connection via ssh issue this command to shut down the vm
  #
  # What to do with the resulting VM
  #
  output_directory = "packer-centos8-local" #  Store the resulting xva file in this directory
  keep_vm          = "on_success" # Create a template with the vm after a successfull build
composes
}
#
# Operations to do while building.
# Any provisioning would be done here
#
build {
  sources = ["xenserver-iso.example"]
}
