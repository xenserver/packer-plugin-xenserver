![Continuous Integration](https://github.com/terra-farm/go-xen-api-client/workflows/Continuous%20Integration/badge.svg)

# Go XenAPI client library

This is a client library for the Xapi toolstack
(http://xapi-project.github.io/).

This library covers the entire [XenAPI](https://xapi-project.github.io/xen-api/)
and I have successfully used it to implement a Terraform plugin that interfaces
Citrix XenServer. That being said, this library is not production-ready yet.
Use it at your own risk, and don't expect everything in this library to work
out of the box.

## Usage example

The following example demonstrates how to instruct XenServer to start a VM with
a given name label:

```go
package main

import (
    "fmt"
    "github.com/terra-farm/go-xen-api-client"
)

const XEN_API_URL string = "https://IP.OF.XEN.SERVER"
const XEN_API_USERNAME string = "USERNAME"
const XEN_API_PASSWORD string = "PASSWORD"
const VM_NAME_LABEL = "VM NAME LABEL"

func main() {
    xapi, err := xenapi.NewClient(XEN_API_URL, nil)
    if err != nil {
        panic(err)
    }

    session, err := xapi.Session.LoginWithPassword(XEN_API_USERNAME, XEN_API_PASSWORD, "1.0", "example")
    if err != nil {
        panic(err)
    }

    vms, err := xapi.VM.GetByNameLabel(session, VM_NAME_LABEL)
    if err != nil {
        panic(err)
    }

    if len(vms) == 0 {
        panic(fmt.Errorf("No VM template with name label %q has been found", VM_NAME_LABEL))
    }

    if len(vms) > 1 {
        panic(fmt.Errorf("More than one VM with name label %q has been found", VM_NAME_LABEL))
    }

    vm := vms[0]

    xapi.VM.Start(session, vm, false, false)
    if err != nil {
        panic(err)
    }

    err = xapi.Session.Logout(session)
    if err != nil {
        panic(err)
    }
}
```

## Project status

The most important missing pieces before this library is production-ready are:

- A strategy how to handle the differences in the XenAPI versions.
- Tests, at least for the various data type conversions.
- Embed XenAPI documentation as GoDoc in the generated code.
- Better error messages.
- Usage examples.

Contributions welcome!

Please note that I want to keep this library lean. I envision it to merely
provide a one-to-one mapping of XenAPI functions to Go functions. Because of
this, I will likely not accept pull requests that implement higher level
functionality.

## Implementation notes

Most of the code in this library is generated from a description of the XenAPI.
This description is the file `xenapi.json`, the source of which is the XenAPI
documentation at http://xapi-project.github.io/:

- https://github.com/xapi-project/xapi-project.github.io/tree/master/_data

The list of error code constants in `error.go` is borrowed from xapi-projects
OCaml client:

- https://github.com/xapi-project/xen-api/blob/master/ocaml/idl/api_errors.ml

The JSON file contains the lifecycle of published classes, fields and messages.
Each of the release names can be mapped back to a version listed here:

- https://xapi-project.github.io/xen-api/releases.html

## Regenerating API after xenapi.json update
If XenAPI was updated, it is required to regenerate all of files with a new API description. In order to do that one needs to follow these steps:
- Get newest `xenapi.json` from the link above.
- Delete old generated APIs using `rm *_gen.go`
- Generate new API with `go generate`
