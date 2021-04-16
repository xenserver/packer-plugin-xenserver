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
After changing those, run 
`packer init examples/http/centos/centos8-example.pkr.hcl` to download the xenserver plugin
and 
`packer build examples/http/centos/centos8-example.pkr.hcl` to create the vm

## Parts of this example
This examples composes the following files from the examples folder

* [centos8-example.pkr.hcl](../../examples/http/centos/centos8-example.pkr.hcl)
* [ks-centos8-example.cfg](../../examples/http/centos/ks-centos8-example.cfg)

Both file contain comments explaining what they do


