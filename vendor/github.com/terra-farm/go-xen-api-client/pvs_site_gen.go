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

type PVSSiteRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// Unique identifier of the PVS site, as configured in PVS
	PVSUUID string
	// The SR used by PVS proxy for the cache
	CacheStorage []PVSCacheStorageRef
	// The set of PVS servers in the site
	Servers []PVSServerRef
	// The set of proxies associated with the site
	Proxies []PVSProxyRef
}

type PVSSiteRef string

// machines serving blocks of data for provisioning VMs
type PVSSiteClass struct {
	client *Client
}

// GetAllRecords Return a map of PVS_site references to PVS_site records for all PVS_sites known to the system.
func (_class PVSSiteClass) GetAllRecords(sessionID SessionRef) (_retval map[PVSSiteRef]PVSSiteRecord, _err error) {
	_method := "PVS_site.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefToPVSSiteRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the PVS_sites known to the system.
func (_class PVSSiteClass) GetAll(sessionID SessionRef) (_retval []PVSSiteRef, _err error) {
	_method := "PVS_site.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefSetToGo(_method + " -> ", _result.Value)
	return
}

// SetPVSUUID Update the PVS UUID of the PVS site
func (_class PVSSiteClass) SetPVSUUID(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	_method := "PVS_site.set_PVS_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Forget Remove a site's meta data
//
// Errors:
//  PVS_SITE_CONTAINS_RUNNING_PROXIES - The PVS site contains running proxies.
//  PVS_SITE_CONTAINS_SERVERS - The PVS site contains servers and cannot be forgotten.
func (_class PVSSiteClass) Forget(sessionID SessionRef, self PVSSiteRef) (_err error) {
	_method := "PVS_site.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Introduce Introduce new PVS site
func (_class PVSSiteClass) Introduce(sessionID SessionRef, nameLabel string, nameDescription string, pvsUUID string) (_retval PVSSiteRef, _err error) {
	_method := "PVS_site.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_pvsUUIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "PVS_uuid"), pvsUUID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameLabelArg, _nameDescriptionArg, _pvsUUIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefToGo(_method + " -> ", _result.Value)
	return
}

// SetNameDescription Set the name/description field of the given PVS_site.
func (_class PVSSiteClass) SetNameDescription(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	_method := "PVS_site.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameLabel Set the name/label field of the given PVS_site.
func (_class PVSSiteClass) SetNameLabel(sessionID SessionRef, self PVSSiteRef, value string) (_err error) {
	_method := "PVS_site.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// GetProxies Get the proxies field of the given PVS_site.
func (_class PVSSiteClass) GetProxies(sessionID SessionRef, self PVSSiteRef) (_retval []PVSProxyRef, _err error) {
	_method := "PVS_site.get_proxies"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSProxyRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetServers Get the servers field of the given PVS_site.
func (_class PVSSiteClass) GetServers(sessionID SessionRef, self PVSSiteRef) (_retval []PVSServerRef, _err error) {
	_method := "PVS_site.get_servers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSServerRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetCacheStorage Get the cache_storage field of the given PVS_site.
func (_class PVSSiteClass) GetCacheStorage(sessionID SessionRef, self PVSSiteRef) (_retval []PVSCacheStorageRef, _err error) {
	_method := "PVS_site.get_cache_storage"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSCacheStorageRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetPVSUUID Get the PVS_uuid field of the given PVS_site.
func (_class PVSSiteClass) GetPVSUUID(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	_method := "PVS_site.get_PVS_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameDescription Get the name/description field of the given PVS_site.
func (_class PVSSiteClass) GetNameDescription(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	_method := "PVS_site.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given PVS_site.
func (_class PVSSiteClass) GetNameLabel(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	_method := "PVS_site.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given PVS_site.
func (_class PVSSiteClass) GetUUID(sessionID SessionRef, self PVSSiteRef) (_retval string, _err error) {
	_method := "PVS_site.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the PVS_site instances with the given label.
func (_class PVSSiteClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []PVSSiteRef, _err error) {
	_method := "PVS_site.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the PVS_site instance with the specified UUID.
func (_class PVSSiteClass) GetByUUID(sessionID SessionRef, uuid string) (_retval PVSSiteRef, _err error) {
	_method := "PVS_site.get_by_uuid"
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
	_retval, _err = convertPVSSiteRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given PVS_site.
func (_class PVSSiteClass) GetRecord(sessionID SessionRef, self PVSSiteRef) (_retval PVSSiteRecord, _err error) {
	_method := "PVS_site.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertPVSSiteRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPVSSiteRecordToGo(_method + " -> ", _result.Value)
	return
}
