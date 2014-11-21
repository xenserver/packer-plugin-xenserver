# XenServer packer.io builder

This builder plugin extends packer.io to support building images for XenServer. 

You can check out packer [here](https://packer.io).


## Dependencies
* Packer >= 0.7.2 (https://packer.io)
* XenServer > 6.2 (http://xenserver.org)
* Golang (tested with 1.2.1) 


In order for this integration to work you must also configure NATing in Dom0. 

You can do this by executing the following in Dom0:

```shell
# Install netcat
yum install --enablerepo=base,extras --disablerepo=citrix -y nc
# Setup NAT - NB, this _disable the firewall_ - be careful!
echo 1 > /proc/sys/net/ipv4/ip_forward
/sbin/iptables -F INPUT

/sbin/iptables -t nat -A POSTROUTING -o xenbr0 -j MASQUERADE
/sbin/iptables -A INPUT -i xenbr0 -p tcp -m tcp --dport 53 -j ACCEPT
/sbin/iptables -A INPUT -i xenbr0 -p udp -m udp --dport 53 -j ACCEPT
/sbin/iptables -A FORWARD -i xenbr0 -o xenapi -m state --state RELATED,ESTABLISHED -j ACCEPT
/sbin/iptables -A FORWARD -i xenapi -o xenbr0 -j ACCEPT
```
(Borrowed from: jonludlam/vagrant-xenserver)


## Install Go

Follow these instructions and install golang on your system:
* https://golang.org/doc/install

## Install Packer

Visit https://packer.io and install the latest version of packer. Once the
install has completed, setup an environment variable 'PACKERPATH' pointing
to the installation location. E.g.

```shell
export PACKERPATH=/usr/local/packer
```

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

If the build is successful, you should now have a 'packer-builder-xenserver' binary
in your $PACKERPATH directory and you are ready to get going with packer.

## Centos 6.4 Example

Once you've setup the above, you are good to go with an example. 

To get you started, there is an example config file which you can use `examples/centos-6.4.conf`:

```shell
{
    "builders": [{
        "type": "xenserver",
        "username": "root",
        "password": "hostpassword",
        "host_ip": "10.81.2.105",
        "instance_name": "packer-centos-6-4",
        "instance_memory": "2048000000",
        "root_disk_size": "5000000000",
        "iso_name": "CentOS-6.4-x86_64-minimal.iso",
        "http_directory": "http",
        "local_ip": "10.80.3.223",
        "ssh_username": "root",
        "ssh_password": "vmpassword",
        "boot_command": 
            [
                "<tab><wait>",
                " ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/centos6-ks.cfg<enter>"
            ]
    }]
}
```
Currently it is not (easily) possible to take care of the ISO download and upload,
so you will need to attach an ISO SR to the XenServer host (NFS/CIFS) with the
ISO you want to use for installation. You will then need to specify the name
in the config file (this must be unique).


An explanation of what these parameters are doing:
 * `type` - this specifies the builder. This must be 'xenserver'.
 * `username` - this is the username for the XenServer host being used.
 * `password` - this is the password for the XenServer host being used.
 * `host_ip` - this is the IP for the XenServer host being used.
 * `instance_name` - this is the name that should be given to the created VM.
 * `instance_memory` - this is the static memory configuration for the VM.
 * `root_disk_size` - this is the size of the disk the VM should be created with.
 * `iso_name` - this is the name of the ISO visible on a ISO SR connected to the XenServer host.
 * `http_directory` - the path to a local directory to serve up over http.
 * `local_ip` - the IP on the machine you are running packer that your XenServer can connect too.
 * `ssh_username` - the username set by the installer for the instance.
 * `ssh_password` - the password set by the installer for the instance.
 * `boot_command` - a set of commands to be sent to the instance over VNC.


Note, the `http_directory` and `local_ip` parameters are only required if you
want packer to serve up files over HTTP. In this example, the templated variables
`{{ .HTTPIP }}` and `{{ .HTTPPort }}` will be substituted for the `local_ip` and
the port that packer starts it's HTTP service on.

Once you've updated the config file with your own parameters, you can use packer
to build this VM with the following:

```
packer build centos-6.4.conf
```







