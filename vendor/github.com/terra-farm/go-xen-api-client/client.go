//go:generate go run xenapi.go

package xenapi

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/amfranz/go-xmlrpc-client"
)

type APIResult struct {
	Value interface{}
}

func (client *Client) APICall(method string, params ...interface{}) (result APIResult, err error) {
	rpcParams := xmlrpc.Params{
		Params: params,
	}

	rpcResult := xmlrpc.Struct{}

	err = client.rpc.Call(method, rpcParams, &rpcResult)
	if err != nil {
		return
	}

	status, ok := rpcResult["Status"].(string)
	if !ok {
		err = fmt.Errorf("Expected a field named %q with a string value in the response", "Status")
		return
	}

	if status != "Success" {
		details := rpcResult["ErrorDescription"].([]interface{})
		_objtype := ""
		if len(details) > 1 && details[1] != nil {
			_objtype = details[1].(string)
		}
		var _uuid string
		if len(details) > 2 && details[2] != nil {
			_uuid = details[2].(string)
		}
		err = &Error{
			code:    details[0].(string),
			objtype: _objtype, // might be nil
			uuid:    _uuid,    // optional
		}
		return
	}

	result.Value = rpcResult["Value"]
	return
}

func NewClient(url string, transport *http.Transport) (*Client, error) {
	if transport == nil {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	rpc, err := xmlrpc.NewClient(url, transport)
	if err != nil {
		return nil, err
	}

	return prepClient(rpc), nil
}
