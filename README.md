# XenServer packer.io builder

This builder plugin extends packer.io to support building images for XenServer. 

You can check out packer [here](https://packer.io).


## Dependencies
* Packer >= 0.7.2 (https://packer.io)
* XenServer > 6.2 (http://xenserver.org)
* Golang (tested with 1.2.1) 


## Install Go

Follow these instructions and install golang on your system:
* https://golang.org/doc/install

## Install Packer

Clone the Packer repository:

```shell
git clone https://github.com/mitchellh/packer.git
```

Then follow the [instructions to build and install a development version of Packer](https://github.com/mitchellh/packer#developing-packer).

## Compile the plugin

Once you have installed Packer, you must compile this plugin and install the
resulting binary.

```shell
cd $GOROOT
mkdir -p src/github.com/rdobson/
cd src/github.com/rdobson
git clone https://github.com/rdobson/packer-builder-xenserver.git
cd packer-builder-xenserver
./build.sh
```

If the build is successful, you should now have a `packer-builder-xenserver` binary
in your `$GOPATH/bin` directory and you are ready to get going with packer.

## Centos 6.6 Example

Once you've setup the above, you are good to go with an example. 

To get you started, there is an example config file which you can use:
[`examples/centos-6.6.json`](https://github.com/rdobson/packer-builder-xenserver/blob/master/examples/centos-6.6.json)

The example is functional, once suitable `remote_host`, `remote_username` and
`remote_password` configurations have been substituted.

A brief explanation of what the config parameters mean:
 * `type` - specifies the builder type. This is 'xenserver-iso', for installing
   a VM from scratch, or 'xenserver-xva' to import existing XVA as a starting
   point.
 * `remote_username` - the username for the XenServer host being used.
 * `remote_password` - the password for the XenServer host being used.
 * `remote_host` - the IP for the XenServer host being used.
 * `vm_name` - the name that should be given to the created VM.
 * `vm_memory` - the static memory configuration for the VM, in MB.
 * `disk_size` - the size of the disk the VM should be created with, in MB.
 * `iso_name` - the name of the ISO visible on a ISO SR connected to the XenServer host.
 * `http_directory` - the path to a local directory to serve up over http.
 * `ssh_username` - the username set by the installer for the instance.
 * `ssh_password` - the password set by the installer for the instance.
 * `boot_command` - a list of commands to be sent to the instance over VNC.

Note, the `http_directory` parameter is only required if you
want Packer to serve up files over HTTP. In this example, the templated variables
`{{ .HTTPIP }}` and `{{ .HTTPPort }}` will be substituted for the local IP and
the port that Packer starts its HTTP service on.

Once you've updated the config file with your own parameters, you can use packer
to build this VM with the following command:

```
packer build centos-6.6.json
```

# Documentation

For complete documentation on configuration commands, see either [the
xenserver-iso docs](https://github.com/rdobson/packer-builder-xenserver/blob/master/docs/builders/xenserver-iso.html.markdown) or [the xenserver-xva docs](https://github.com/rdobson/packer-builder-xenserver/blob/master/docs/builders/xenserver-xva.html.markdown).
