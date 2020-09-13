//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenapi

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type DataSourceRecord struct {
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// true if the data source is being logged
	Enabled bool
	// true if the data source is enabled by default. Non-default data sources cannot be disabled
	Standard bool
	// the units of the value
	Units string
	// the minimum value of the data source
	Min float64
	// the maximum value of the data source
	Max float64
	// current value of the data source
	Value float64
}

type DataSourceRef string

// Data sources for logging in RRDs
type DataSourceClass struct {
	client *Client
}
