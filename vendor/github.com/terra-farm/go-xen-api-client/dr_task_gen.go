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

type DRTaskRecord struct {
	// Unique identifier/object reference
	UUID string
	// All SRs introduced by this appliance
	IntroducedSRs []SRRef
}

type DRTaskRef string

// DR task
type DRTaskClass struct {
	client *Client
}

// GetAllRecords Return a map of DR_task references to DR_task records for all DR_tasks known to the system.
func (_class DRTaskClass) GetAllRecords(sessionID SessionRef) (_retval map[DRTaskRef]DRTaskRecord, _err error) {
	_method := "DR_task.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefToDRTaskRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the DR_tasks known to the system.
func (_class DRTaskClass) GetAll(sessionID SessionRef) (_retval []DRTaskRef, _err error) {
	_method := "DR_task.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
func (_class DRTaskClass) Destroy(sessionID SessionRef, self DRTaskRef) (_err error) {
	_method := "DR_task.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a disaster recovery task which will query the supplied list of devices
func (_class DRTaskClass) Create(sessionID SessionRef, atype string, deviceConfig map[string]string, whitelist []string) (_retval DRTaskRef, _err error) {
	_method := "DR_task.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_deviceConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "device_config"), deviceConfig)
	if _err != nil {
		return
	}
	_whitelistArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "whitelist"), whitelist)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _atypeArg, _deviceConfigArg, _whitelistArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// GetIntroducedSRs Get the introduced_SRs field of the given DR_task.
func (_class DRTaskClass) GetIntroducedSRs(sessionID SessionRef, self DRTaskRef) (_retval []SRRef, _err error) {
	_method := "DR_task.get_introduced_SRs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given DR_task.
func (_class DRTaskClass) GetUUID(sessionID SessionRef, self DRTaskRef) (_retval string, _err error) {
	_method := "DR_task.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the DR_task instance with the specified UUID.
func (_class DRTaskClass) GetByUUID(sessionID SessionRef, uuid string) (_retval DRTaskRef, _err error) {
	_method := "DR_task.get_by_uuid"
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
	_retval, _err = convertDRTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given DR_task.
func (_class DRTaskClass) GetRecord(sessionID SessionRef, self DRTaskRef) (_retval DRTaskRecord, _err error) {
	_method := "DR_task.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertDRTaskRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDRTaskRecordToGo(_method + " -> ", _result.Value)
	return
}
