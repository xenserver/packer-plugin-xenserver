---
layout: "docs"
page_title: "XenServer Builder (from an ISO)"
description: |-
  The XenServer Packer builder is able to create XenServer virtual machines and export them either as an XVA or a VDI and create VM templates starting from an ISO image.
---

# XenServer Builder (from an ISO)

Type: `xenserver-iso`

The XenServer Packer builder is able to create [XenServer](https://www.xenserver.org/)
virtual machines and export them either as an XVA or a VDI and create VM templates starting from an ISO image.

The builder builds a virtual machine by creating a new virtual machine
from scratch, booting it, installing an OS, provisioning software within
the OS, then shutting it down. The result of the XenServer builder is a
directory containing all the files necessary to run the virtual machine
portably.


## Configuration Reference

There are many configuration options available for the XenServer builder.
They are organized below into two categories: required and optional. Within
each category, the available options are alphabetized and described.

### Required:

* `iso_checksum` (string) - The checksum for the OS ISO file. Because ISO
  files are so large, this is required and Packer will verify it prior
  to booting a virtual machine with the ISO attached. The type of
  the checksum is specified within the checksum field as a prefix, ex:
  "md5:{$checksum}". The type of the checksum can also be omitted and
  Packer will try to infer it based on string length. Valid values are
  "none", "{$checksum}", "md5:{$checksum}", "sha1:{$checksum}",
  "sha256:{$checksum}", "sha512:{$checksum}" or "file:{$path}". Here is a
  list of valid checksum values:
   * md5:090992ba9fd140077b0661cb75f7ce13
   * 090992ba9fd140077b0661cb75f7ce13
   * sha1:ebfb681885ddf1234c18094a45bbeafd91467911
   * ebfb681885ddf1234c18094a45bbeafd91467911
   * sha256:ed363350696a726b7932db864dda019bd2017365c9e299627830f06954643f93
   * ed363350696a726b7932db864dda019bd2017365c9e299627830f06954643f93
   * file:http://releases.ubuntu.com/20.04/SHA256SUMS
   * file:file://./local/path/file.sum
   * file:./local/path/file.sum
   * none
  Although the checksum will not be verified when it is set to "none",
  this is not recommended since these files can be very large and
  corruption does happen from time to time.


* `iso_url` (string) - A URL to the ISO containing the installation image.
  This URL can be either an HTTP URL or a file URL (or path to a file).
  If this is an HTTP URL, Packer will download it and cache it between
  runs.

* `remote_host` (string) - The host of the Xenserver / XCP-ng pool primary. Typically these will be specified through environment variables as seen in the [examples](../../../examples).

* `remote_username` (string) - The XenServer username used to access the remote machine.

* `remote_password` (string) - The XenServer password for access to the remote machine.

* `ssh_username` (string) - The username to use to SSH into the machine
  once the OS is installed.

### Optional:

* `boot_command` (array of strings) - This is an array of commands to type
  when the virtual machine is first booted. The goal of these commands should
  be to type just enough to initialize the operating system installer. Special
  keys can be typed as well, and are covered in the section below on the boot
  command. If this is not specified, it is assumed the installer will start
  itself. See the [Ubuntu](../../../examples/ubuntu) and [centos](../../../examples/centos) examples to see how these are used to launch autoinstall and kickstart respectively.

* `boot_wait` (string) - The time to wait after booting the initial virtual
  machine before typing the `boot_command`. The value of this should be
  a duration. Examples are "5s" and "1m30s" which will cause Packer to wait
  five seconds and one minute 30 seconds, respectively. If this isn't specified,
  the default is 10 seconds.

* `clone_template` (string) - The template to clone. Defaults to "Other install
  media", this is "other", but you can get _dramatic_ performance improvements
  by setting this to the proper value. To view all available values for this
  run `xe template-list`. Setting the correct value hints to XenServer how to
  optimize the virtual hardware to work best with that operating system.

* `disk_size` (integer) - The size, in megabytes, of the hard disk to create
  for the VM. By default, this is 40000 (about 40 GB).

* `floppy_files` (array of strings) - A list of files to place onto a floppy
  disk that is attached when the VM is booted. This is most useful
  for unattended Windows installs, which look for an `Autounattend.xml` file
  on removable media. By default, no floppy will be attached. All files
  listed in this setting get placed into the root directory of the floppy
  and the floppy is attached as the first floppy device. Currently, no
  support exists for creating sub-directories on the floppy. Wildcard
  characters (\*, ?, and []) are allowed. Directory names are also allowed,
  which will add all the files found in the directory to the floppy.

* `format` (string) - Either "xva", "vdi_raw" or "none", this specifies the
  output format of the exported virtual machine. This defaults to "xva". Set to
  "vdi_raw" to export just the raw disk image. Set to "none" to export nothing;
  this is only useful with "keep_vm" set to "always" or "on_success".

* `http_directory` (string) - Path to a directory to serve using an HTTP
  server. The files in this directory will be available over HTTP that will
  be requestable from the virtual machine. This is useful for hosting
  kickstart files and so on. By default this is "", which means no HTTP
  server will be started. The address and port of the HTTP server will be
  available as variables in `boot_command`. This is covered in more detail
  below.

* `http_port_min` and `http_port_max` (integer) - These are the minimum and
  maximum port to use for the HTTP server started to serve the `http_directory`.
  Because Packer often runs in parallel, Packer will choose a randomly available
  port in this range to run the HTTP server. If you want to force the HTTP
  server to be on one port, make this minimum and maximum port the same.
  By default the values are 8000 and 9000, respectively.

* `install_timeout` (string) - The amount of time to wait after booting the VM
  for the installer to shut itself down.
  If it doesn't shut down in this time, it is an error. By default, the timeout
  is "200m", or over three hours.

* `iso_urls` (array of strings) - Multiple URLs for the ISO to download.
  Packer will try these in order. If anything goes wrong attempting to download
  or while downloading a single URL, it will move on to the next. All URLs
  must point to the same file (same checksum). By default this is empty
  and `iso_url` is used. Only one of `iso_url` or `iso_urls` can be specified.

* `keep_vm` (string) - Determine when to keep the VM and when to clean it up. This
  can be "always", "never" or "on_success". By default this is "never", and Packer
  always deletes the VM regardless of whether the process succeeded and an artifact
  was produced. "always" asks Packer to leave the VM at the end of the process
  regardless of success. "on_success" requests that the VM only be cleaned up if an
  artifact was produced. The latter is useful for debugging templates that fail.

* `network_names` (array of strings) - A list of networks identified by their name label which 
  will be used for the VM during creation. The first network will correspond to the VM's
  first network interface (VIF), the second will corespond to the second VIF and so on.

* `output_directory` (string) - This is the path to the directory where the
  resulting virtual machine will be created. This may be relative or absolute.
  If relative, the path is relative to the working directory when `packer`
  is executed. This directory must not exist or be empty prior to running the builder.
  By default this is "output-BUILDNAME" where "BUILDNAME" is the name
  of the build.

* `platform_args` (object of key/value strings) - The platform args.
  Defaults to 
```javascript
{
    "viridian": "false",
    "nx": "true",
    "pae": "true",
    "apic": "true",
    "timeoffset": "0",
    "acpi": "1",
    "cores-per-socket": "1"
}
```

* `shutdown_command` (string) - The command to use to gracefully shut down
  the machine once all the provisioning is done. If this is omitted, packer
  will shut down the VM gracefully through the Xen api's vm shutdown command. Unless
  you have special requirements this should typically be left to its default.

* `ssh_host_port_min` and `ssh_host_port_max` (integer) - The minimum and
  maximum port to use for the SSH port on the host machine which is forwarded
  to the SSH port on the guest machine. Because Packer often runs in parallel,
  Packer will choose a randomly available port in this range to use as the
  host port.

* `ssh_key_path` (string) - Path to a private key to use for authenticating
  with SSH. By default this is not set (key-based auth won't be used).
  The associated public key is expected to already be configured on the
  VM being prepared by some other process (kickstart, etc.).

* `ssh_password` (string) - The password for `ssh_username` to use to
  authenticate with SSH. By default this is the empty string.

* `ssh_port` (integer) - The port that SSH will be listening on in the guest
  virtual machine. By default this is 22.

* `ssh_wait_timeout` (string) - The duration to wait for SSH to become
  available. By default this is "20m", or 20 minutes. Note that this should
  be quite long since the timer begins as soon as the virtual machine is booted.

* `tools_iso_name` (string) - The name of the XenServer Tools ISO. Defaults to
  "xs-tools.iso".

* `vm_description` (string) - The description of the new virtual
  machine. By default this is the empty string.

* `vm_name` (string) - This is the name of the new virtual
  machine, without the file extension. By default this is
  "packer-BUILDNAME-TIMESTAMP", where "BUILDNAME" is the name of the build.

* `vcpus_max` (integer) - The maximum number of VCPUs for the VM.
  By default this is 1.

* `vcpus_atstartup` (integer) - The number of startup VCPUs for the VM.
  By default this is 1.

* `vm_memory` (integer) - The size, in megabytes, of the amount of memory to
  allocate for the VM. By default, this is 1024 (1 GB).

## Differences with other Packer builders

Currently the XenServer builder has some quirks when compared with other Packer builders.

The builder currently only works remotely.

The installer is expected to shut down the VM to indicate that it has completed. This is in contrast to other builders, which instead detect completion by a successful SSH connection. The reason for this difference is that currently the builder has no way of knowing what the IP address of the VM is without starting it on the HIMN.

## Boot Command

The `boot_command` configuration is very important: it specifies the keys
to type when the virtual machine is first booted in order to start the
OS installer. This command is typed after `boot_wait`, which gives the
virtual machine some time to actually load the ISO.

As documented above, the `boot_command` is an array of strings. The
strings are all typed in sequence. It is an array only to improve readability
within the template.

The boot command is "typed" character for character over a VNC connection
to the machine, simulating a human actually typing the keyboard. There are
a set of special keys available. If these are in your boot command, they
will be replaced by the proper key:

* `<bs>` - Backspace

* `<del>` - Delete

* `<enter>` and `<return>` - Simulates an actual "enter" or "return" keypress.

* `<esc>` - Simulates pressing the escape key.

* `<tab>` - Simulates pressing the tab key.

* `<f1>` - `<f12>` - Simulates pressing a function key.

* `<up>` `<down>` `<left>` `<right>` - Simulates pressing an arrow key.

* `<spacebar>` - Simulates pressing the spacebar.

* `<insert>` - Simulates pressing the insert key.

* `<home>` `<end>` - Simulates pressing the home and end keys.

* `<pageUp>` `<pageDown>` - Simulates pressing the page up and page down keys.

* `<wait>` `<wait5>` `<wait10>` - Adds a 1, 5 or 10 second pause before sending any additional keys. This
  is useful if you have to generally wait for the UI to update before typing more.

In addition to the special keys, each command to type is treated as a
configuration template.
The available variables are:

* `HTTPIP` and `HTTPPort` - The IP and port, respectively of an HTTP server
  that is started serving the directory specified by the `http_directory`
  configuration parameter. If `http_directory` isn't specified, these will be
  blank!

See the [examples](../../../examples) for working boot commands.
