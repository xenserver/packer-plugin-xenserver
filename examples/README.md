## Examples

In order for new users to get up and running with the packer builder, a few examples of building a machine image with popular distros have been created.

In order to see an exhaustive list of configuration options for the packer builder please see the [following documentation](../docs/builders/xenserver-iso.html.markdown). This doc will focus on the details relevant to the particular distro.

### Running the examples

In order to run this example you will need to perform the following steps:
1. Export those vars:
```
PKR_VAR_remote_host
PKR_VAR_remote_password
PKR_VAR_remote_username
PKR_VAR_sr_name
PKR_VAR_sr_iso_name
``` 
`PKR_VAR_remote_host` must be the resource pool primary.


2. Run `packer build  path/to/defenition.pkr.hcl`   
so for example:
`packer build  examples/centos/centos8-netinstall.pkr.hcl`


### Ubuntu

This example has not yet updated to HCL.

In order to run this example you will need to perform the following steps:
1. Export the `XAPI_HOST`, `XAPI_USERNAME` and `XAPI_PASSWORD` environment variables to the current shell. Note: The `XAPI_HOST` must be the resource pool primary.
2. Run the `packer build` command specifying the storage repositories to use for the ISO upload and for the VM created during the build.

```
# Replace sr_name and sr_iso_name with your storage repositories names
packer build -debug  --var sr_name='Local storage' --var sr_iso_name=LocalISO examples/centos8.json

# Do the same variable replacement for the ubuntu example as well.
packer build -debug  --var sr_name='Local storage' --var sr_iso_name=LocalISO examples/ubuntu-2004.json
```


The Ubuntu example uses the [autoinstall tool](https://ubuntu.com/server/docs/install/autoinstallhttps://ubuntu.com/server/docs/install/autoinstall) to configure the VM template. Please see the [autoinstall docs](https://ubuntu.com/server/docs/install/autoinstall-reference) for an exhaustive list of what is supported.

Packer will create a http server to serve the files as specified from the `http_directory` specified in the builder configuration. This is where the [user-data](http/ubuntu-2004/user-data) and [meta-data](http/ubuntu-2004/meta-data) for autoinstall must be present.

### Centos

The Centos examples use kickstart files to configure the VM template. Please see the [kickstart documentation](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/installation_guide/sect-kickstart-syntax) for the options that are supported.

Packer will create a http server to serve the files as specified from the `http_directory` specified in the builder configuration. This is where the [kickstart config](http/centos8/ks-centos8.cfg) file must be present.
