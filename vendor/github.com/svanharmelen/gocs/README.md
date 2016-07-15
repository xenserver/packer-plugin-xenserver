# gocs

gocs is a CloudStack client that enables Go programs to interact with the CloudStack API in a simple and uniform way.
The package itself is smart enough to learn the available CloudStack API commands and the required and optional
parameters of all available commands (using the listApis command), so it can be used with any given version of
CloudStack.

Based on the processed info about all available commands and their details, the request functions will check
if all required parameters have a value and if all other used parameters are valid parameters for the requested
command before actually passing the request to the API. In case of any errors during this validation or when
executing the actual request, detailed error info will be returned to the requesting function making it very
easy to understand and fix the problem.


## Usage

Install/Update the Go package:
```
go get -u github.com/svanharmelen/gocs
```

Add this to your Go program:
```
import "github.com/svanharmelen/gocs"
```


## Documentation

Package documentation is dynamically generated from the source code at [GoDoc](http://godoc.org/github.com/svanharmelen/gocs)
