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

type PvsProxyStatus string

const (
	// The proxy is not currently running
	PvsProxyStatusStopped PvsProxyStatus = "stopped"
	// The proxy is setup but has not yet cached anything
	PvsProxyStatusInitialised PvsProxyStatus = "initialised"
	// The proxy is currently caching data
	PvsProxyStatusCaching PvsProxyStatus = "caching"
	// The PVS device is configured to use an incompatible write-cache mode
	PvsProxyStatusIncompatibleWriteCacheMode PvsProxyStatus = "incompatible_write_cache_mode"
	// The PVS protocol in use is not compatible with the PVS proxy
	PvsProxyStatusIncompatibleProtocolVersion PvsProxyStatus = "incompatible_protocol_version"
)

type PVSProxyRecord struct {
	// Unique identifier/object reference
	UUID string
	// PVS site this proxy is part of
	Site PVSSiteRef
	// VIF of the VM using the proxy
	VIF VIFRef
	// true = VM is currently proxied
	CurrentlyAttached bool
	// The run-time status of the proxy
	Status PvsProxyStatus
}

type PVSProxyRef string

// a proxy connects a VM/VIF with a PVS site
type PVSProxyClass struct {
	client *Client
}

// GetAllRecords Return a map of PVS_proxy references to PVS_proxy records for all PVS_proxys known to the system.
func (_class PVSProxyClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSProxyRef]PVSProxyRecord, _err error) {
	_method := "PVS_proxy.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefToPVSProxyRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the PVS_proxys known to the system.
func (_class PVSProxyClass) GetAll(sessionID SessionRef) (_retval []PVSProxyRef, _err error) {
	_method := "PVS_proxy.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy remove (or switch off) a PVS proxy for this VM
func (_class PVSProxyClass) Destroy(sessionID SessionRef, self PVSProxyRef) (_err error) {
	_method := "PVS_proxy.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Configure a VM/VIF to use a PVS proxy
func (_class PVSProxyClass) Create(sessionID SessionRef, site PVSSiteRef, vif VIFRef) (_retval PVSProxyRef, _err error) {
	_method := "PVS_proxy.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_siteArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "site"), site)
	if _err != nil {
		return
	}
	_vifArg, _err := convertVIFRefToXen(fmt.Sprintf("%s(%s)", _method, "VIF"), vif)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _siteArg, _vifArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefToGo(_method + " -> ", _result.Value)
	return
}

// GetStatus Get the status field of the given PVS_proxy.
func (_class PVSProxyClass) GetStatus(sessionID SessionRef, self PVSProxyRef) (_retval PvsProxyStatus, _err error) {
	_method := "PVS_proxy.get_status"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumPvsProxyStatusToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentlyAttached Get the currently_attached field of the given PVS_proxy.
func (_class PVSProxyClass) GetCurrentlyAttached(sessionID SessionRef, self PVSProxyRef) (_retval bool, _err error) {
	_method := "PVS_proxy.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// GetVIF Get the VIF field of the given PVS_proxy.
func (_class PVSProxyClass) GetVIF(sessionID SessionRef, self PVSProxyRef) (_retval VIFRef, _err error) {
	_method := "PVS_proxy.get_VIF"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefToGo(_method + " -> ", _result.Value)
	return
}

// GetSite Get the site field of the given PVS_proxy.
func (_class PVSProxyClass) GetSite(sessionID SessionRef, self PVSProxyRef) (_retval PVSSiteRef, _err error) {
	_method := "PVS_proxy.get_site"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given PVS_proxy.
func (_class PVSProxyClass) GetUUID(sessionID SessionRef, self PVSProxyRef) (_retval string, _err error) {
	_method := "PVS_proxy.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the PVS_proxy instance with the specified UUID.
func (_class PVSProxyClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSProxyRef, _err error) {
	_method := "PVS_proxy.get_by_uuid"
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
	_retval, _err = convertPVSProxyRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given PVS_proxy.
func (_class PVSProxyClass) GetRecord(sessionID SessionRef, self PVSProxyRef) (_retval PVSProxyRecord, _err error) {
	_method := "PVS_proxy.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSProxyRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRecordToGo(_method + " -> ", _result.Value)
	return
}
