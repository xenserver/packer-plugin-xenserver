---
layout: "docs"
page_title: "XenServer-iso Commented Example"
description: |-
  This is a Commented example for how to use the xenserver-iso Packer Builder.
  This Example is written in hcl
---

# What's this 
This Example builds a vanilla Centos 8 Template. No Provisioning or Post-Processing will be done.

If you want to run this example, you have to replace the following variables with your own values.
```hcl
  remote_host     = "xenserver.example.org"
  remote_username = "root"
  remote_password = "very-secret-password"
  sr_iso_name     = "Local-ISO"
  sr_name         = "Local-SR"
  tools_iso_name  = "guest-tools.iso"
```

## Parts of this example
This examples composes the following files from the examples folder

* [centos8-example.pkr.hcl](../../examples/http/centos/centos8-example.pkr.hcl)
* [ks-centos8-example.cfg](../../examples/http/centos/ks-centos8-example.cfg)


## Explanation of centos8-example.pkr.hcl

* `source "xenserver-iso" "example" { } ` directs packer to configure an Artifact named example using the `xenserver-iso` builder

* `iso_url  = "http://mirrors.ocf.berkeley.edu/centos/8.3.2011/isos/x86_64/CentOS-8.3.2011-x86_64-dvd1.iso" ` use the iso obtainable from this url
* `iso_checksum_type = "sha1"` the checksum is of type sha1
* `iso_checksum      = "aaf9d4b3071c16dbbda01dfe06085e5d0fdac76df323e3bbe87cce4318052247"` The checksum for the ISO. 

* `sr_name        = "Local-SR"` store the vmdisk used during install on the SR named `Local-SR`
* `sr_iso_name    = "Local-ISO"` store the iso used during install on the ISO-SR named `Local-ISO`
* `tools_iso_name = "guest-tools.iso"` mount the guest-tools iso with the name `guest-tools.iso`

* `remote_host     = "xenserver.example.org"` the ipadress or fqdn of the xenserver. This should be the pool primary
* `remote_username = "root"` the user with witch to connect.
* `remote_password = "very-secret-password"` the password for the user.

* `vm_name        = "packer-centos8-example"` How packer will name the vm.
* `vm_description = "This is an example."` The description field of the vm
* `vm_memory      = 4096` the Amount of RAM, in MB
* `disk_size      = 4096` the Size of the Primary Disk in MB

* `http_directory = "examples/http/centos8'`  Packer will spin up a http-server for serving files to the installing vm. the kickstart file `ks-centos8-examples.cfg` is in this directory
* `boot_wait      = "10s'`  Wait for 10s after starting the VM before typeing the `boot_command`
* `boot_command   = ["<tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks-centos8-examples.cfg<enter><wait>"'` 
  * The command to type into the installing vm.
  * `<tab>` a tab Character `\t`
  * `text ks=http://` character literals
  * `{{ .HTTPIP }}` will be replaced by the local ip-address.
  * `:` character literal
  * `{{ .HTTPPort }}` will be replaced by the local port. Will be randomly selectet between 8000 and 9000
  * `/ks-centos8-examples.cfg` character literals
  * `<enter>` an enter character
  * `<wait>` wait for 1s

* `ssh_username     = "root"` The ssh user packer uses to connect to the VM
* `ssh_password     = "centos"` The ssh password packer uses to connect to the VM
* `ssh_wait_timeout = "10m"` consider install failed if unable to connect via ssh 10m into the build
* `shutdown_command = "/sbin/shutdown"` After connection via ssh issue this command to shut down the vm
* `output_directory = "packer-centos8-local"` Store the resulting xva file in this directory
* `keep_vm          = "on_success"` Create a template with the vm after a successfull build

* `build { }` Tell packer what to do while builing
* `sources = ["xenserver-iso.example"]` Build the `xenserver.example` without changing any configuration.

## Explanation of ks-centos8-examples.cfg
```
eula --agreed
# agree to the eula of centos

lang en-US.UTF-8
# set the system-locale to en-US.UTF-8

timezone Europe/Berlin
# set the timezone to Europe/Berlin

url --url="http://mirror.centos.org/centos/8.3.2011/BaseOS/x86_64/os/"
# Primary installation mirror

text
skipx
# Install in textmode, do not configure X11

firstboot --disable
# Dont do any configuration on first boot

rootpw --plaintext centos
# set the reboot to "centos"

firewall --enabled --ssh
# enable the firewall, allow ssh

selinux --enforcing
# ensure selinux is in enforcing mode

logging --level=info
# Installation logging level

network  --bootproto=dhcp --device=eth0 --onboot=on
# configure the network with dhcp on install

clearpart --all
zerombr
bootloader --location=mbr
# Clear all partitioning on all disk
# Zero the MBR section
# Install the bootloader into the MBR

part / --asprimary --fstype="ext4" --size=1024 --grow
# / shall the a primary partition with ext4
# with a mimimul size of 1024 MB, but growing to the actual size of the disk

%packages
@base
%end
# install all packages of the core group

reboot --eject
# After finishing eject all CD-Disk and reboot
```
