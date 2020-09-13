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

type PVSServerRecord struct {
	// Unique identifier/object reference
	UUID string
	// IPv4 addresses of this server
	Addresses []string
	// First UDP port accepted by this server
	FirstPort int
	// Last UDP port accepted by this server
	LastPort int
	// PVS site this server is part of
	Site PVSSiteRef
}

type PVSServerRef string

// individual machine serving provisioning (block) data
type PVSServerClass struct {
	client *Client
}

// GetAllRecords Return a map of PVS_server references to PVS_server records for all PVS_servers known to the system.
func (_class PVSServerClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSServerRef]PVSServerRecord, _err error) {
	_method := "PVS_server.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefToPVSServerRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the PVS_servers known to the system.
func (_class PVSServerClass) GetAll(sessionID SessionRef) (_retval []PVSServerRef, _err error) {
	_method := "PVS_server.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Forget forget a PVS server
func (_class PVSServerClass) Forget(sessionID SessionRef, self PVSServerRef) (_err error) {
	_method := "PVS_server.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Introduce introduce new PVS server
func (_class PVSServerClass) Introduce(sessionID SessionRef, addresses []string, firstPort int, lastPort int, site PVSSiteRef) (_retval PVSServerRef, _err error) {
	_method := "PVS_server.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_addressesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "addresses"), addresses)
	if _err != nil {
		return
	}
	_firstPortArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "first_port"), firstPort)
	if _err != nil {
		return
	}
	_lastPortArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "last_port"), lastPort)
	if _err != nil {
		return
	}
	_siteArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "site"), site)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _addressesArg, _firstPortArg, _lastPortArg, _siteArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefToGo(_method + " -> ", _result.Value)
	return
}

// GetSite Get the site field of the given PVS_server.
func (_class PVSServerClass) GetSite(sessionID SessionRef, self PVSServerRef) (_retval PVSSiteRef, _err error) {
	_method := "PVS_server.get_site"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefToGo(_method + " -> ", _result.Value)
	return
}

// GetLastPort Get the last_port field of the given PVS_server.
func (_class PVSServerClass) GetLastPort(sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	_method := "PVS_server.get_last_port"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// GetFirstPort Get the first_port field of the given PVS_server.
func (_class PVSServerClass) GetFirstPort(sessionID SessionRef, self PVSServerRef) (_retval int, _err error) {
	_method := "PVS_server.get_first_port"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// GetAddresses Get the addresses field of the given PVS_server.
func (_class PVSServerClass) GetAddresses(sessionID SessionRef, self PVSServerRef) (_retval []string, _err error) {
	_method := "PVS_server.get_addresses"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given PVS_server.
func (_class PVSServerClass) GetUUID(sessionID SessionRef, self PVSServerRef) (_retval string, _err error) {
	_method := "PVS_server.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the PVS_server instance with the specified UUID.
func (_class PVSServerClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSServerRef, _err error) {
	_method := "PVS_server.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given PVS_server.
func (_class PVSServerClass) GetRecord(sessionID SessionRef, self PVSServerRef) (_retval PVSServerRecord, _err error) {
	_method := "PVS_server.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSServerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRecordToGo(_method + " -> ", _result.Value)
	return
}
