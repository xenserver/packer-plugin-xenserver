//gocs is a CloudStack client that enables Go programs to interact with the CloudStack API in a simple and uniform way.
//The package itself is smart enough to learn the available CloudStack API commands and the required and optional
//parameters of all available commands (using the listApis command), so it can be used with any given version of
//CloudStack.
//
//Based on the processed info about all available commands and their details, the request functions will check
//if all required parameters have a value and if all other used parameters are valid parameters for the requested
//command before actually passing the request to the API. In case of any errors during this validation or when
//executing the actual request, detailed error info will be returned to the requesting function making it very
//easy to understand and fix the problem.
//
//A usefull (BETA!) feature when making API requests is that the request functions are able to retieve ID's for parameters
//that need an ID of some kind (e.g. virtualmachine, zone) instead of a name. Every parameter that ends with 'id' can also
//be used as a parameter without the 'id' and a value that is a name rather than an ID. In those cases the request
//functions will take care of retrieving all the needed ID's for you.
//
//NOTE: This only works when the list command for the needed ID (e.g. listVirtualMachines, listZones) doesn't require
//additional parameters like the listTemplates does, and (again) it's still a BETA functionality. When the list command
//contains required parameters you would still need to make a call yourself supplying the correct values. And again, this
//is still, and it's still BETA functionality
//
//See these two small snippets that show two ways to make the same call. Starting with the more elaborate one:
//
//  // Original call using ID's, so need to retrieve all ID's first
//  serviceofferingid, err := cs.Request("listServiceOfferings", "name:small")
//  if err =! nil {
//    return err
//  }
//  templateid, err := cs.Request("listTemplates", "templatefilter:featured, name:Centos 6 x64")
//  if err =! nil {
//    return err
//  }
//  zoneid, err := cs.Request("listZones", "name:myzone")
//  if err =! nil {
//    return err
//  }
//  if id, err := cs.Request("deployVirtualMachine", fmt.Sprintf("serviceofferingid:%s, templateid:%s, zoneid:%s",
//                            serviceofferingid, templateid, zoneid); err != nil {
//    return err
//  }
//  return id
//
//Followed by the easier one:
//
//  // Easier call using only one additional call to retrieve the templateid
//  zoneid, err := cs.Request("listTemplates", "templatefilter:featured, name:Centos 6 x64")
//  if err =! nil {
//    return err
//  }
//  if id, err := cs.Request("deployVirtualMachine", fmt.Sprintf("serviceoffering:small, templateid:%s,
//                            zone:my zone", id)); err != nil {
//    return err
//  }
//  return id
//
//
//The response you get from a request is determined by the type of request you are using. When using a normal
//request the function returns only an ID (if the command itself returns an ID of course) or an empty string.
//Additionally you can use a 'raw' request to get the complete raw JSON received from the API. In the latter
//case unmarshalling can be done using a custom struct that matches the response of the called command that
//fits the needs of the program using this package.
//
//The example below shows how to use the package in a very simple commandline utility. The demo utility
//uses a caching client and a raw synced request which waits for async calls to finish and returns the raw
//JSON response as output:
//
//  package main
//
//  import (
//    "github.com/svanharmelen/gocs"
//    "fmt"
//    "log"
//    "os"
//    "path"
//    "strings"
//  )
//
//  // Off course you better not do this in an actual program :)
//  const (
//    apiurl = "https://cloudstack.url/client/api"
//    apikey = "xxx"
//    secret = "xxx"
//  )
//
//  func main() {
//    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
//      fmt.Printf("Usage: %s command <param:value param:\"value\" ...>\n", path.Base(os.Args[0]))
//      os.Exit(1)
//    }
//
//    // Create a new caching client
//    cs, err := gocs.NewCachingClient(apiurl, apikey, secret, 1, false)
//    if err != nil {
//      log.Fatal(err)
//    }
//
//    // Get the command and join all 'parameter:value' pairs into one comma separated string
//    command := os.Args[1]
//    params := strings.Join(os.Args[2:], ",")
//
//    // Make a raw, but synced request
//    response, err := cs.RawSyncedRequest(command, params)
//    if err != nil {
//      log.Fatal(err)
//    }
//    fmt.Println(string(response))
//  }
package gocs
