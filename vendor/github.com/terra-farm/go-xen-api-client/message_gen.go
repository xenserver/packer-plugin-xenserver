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

type Cls string

const (
	// VM
	ClsVM Cls = "VM"
	// Host
	ClsHost Cls = "Host"
	// SR
	ClsSR Cls = "SR"
	// Pool
	ClsPool Cls = "Pool"
	// VMPP
	ClsVMPP Cls = "VMPP"
	// VMSS
	ClsVMSS Cls = "VMSS"
	// PVS_proxy
	ClsPVSProxy Cls = "PVS_proxy"
	// VDI
	ClsVDI Cls = "VDI"
)

type MessageRecord struct {
	// Unique identifier/object reference
	UUID string
	// The name of the message
	Name string
	// The message priority, 0 being low priority
	Priority int
	// The class of the object this message is associated with
	Cls Cls
	// The uuid of the object this message is associated with
	ObjUUID string
	// The time at which the message was created
	Timestamp time.Time
	// The body of the message
	Body string
}

type MessageRef string

// An message for the attention of the administrator
type MessageClass struct {
	client *Client
}

// GetAllRecordsWhere 
func (_class MessageClass) GetAllRecordsWhere(sessionID SessionRef, expr string) (_retval map[MessageRef]MessageRecord, _err error) {
	_method := "message.get_all_records_where"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_exprArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "expr"), expr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _exprArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAllRecords 
func (_class MessageClass) GetAllRecords(sessionID SessionRef) (_retval map[MessageRef]MessageRecord, _err error) {
	_method := "message.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID 
func (_class MessageClass) GetByUUID(sessionID SessionRef, uuid string) (_retval MessageRef, _err error) {
	_method := "message.get_by_uuid"
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
	_retval, _err = convertMessageRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord 
func (_class MessageClass) GetRecord(sessionID SessionRef, self MessageRef) (_retval MessageRecord, _err error) {
	_method := "message.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertMessageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRecordToGo(_method + " -> ", _result.Value)
	return
}

// GetSince 
func (_class MessageClass) GetSince(sessionID SessionRef, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	_method := "message.get_since"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_sinceArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "since"), since)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _sinceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll 
func (_class MessageClass) GetAll(sessionID SessionRef) (_retval []MessageRef, _err error) {
	_method := "message.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get 
func (_class MessageClass) Get(sessionID SessionRef, cls Cls, objUUID string, since time.Time) (_retval map[MessageRef]MessageRecord, _err error) {
	_method := "message.get"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_clsArg, _err := convertEnumClsToXen(fmt.Sprintf("%s(%s)", _method, "cls"), cls)
	if _err != nil {
		return
	}
	_objUUIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "obj_uuid"), objUUID)
	if _err != nil {
		return
	}
	_sinceArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "since"), since)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _clsArg, _objUUIDArg, _sinceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToMessageRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// Destroy 
func (_class MessageClass) Destroy(sessionID SessionRef, self MessageRef) (_err error) {
	_method := "message.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertMessageRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create 
func (_class MessageClass) Create(sessionID SessionRef, name string, priority int, cls Cls, objUUID string, body string) (_retval MessageRef, _err error) {
	_method := "message.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_priorityArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "priority"), priority)
	if _err != nil {
		return
	}
	_clsArg, _err := convertEnumClsToXen(fmt.Sprintf("%s(%s)", _method, "cls"), cls)
	if _err != nil {
		return
	}
	_objUUIDArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "obj_uuid"), objUUID)
	if _err != nil {
		return
	}
	_bodyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "body"), body)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _nameArg, _priorityArg, _clsArg, _objUUIDArg, _bodyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertMessageRefToGo(_method + " -> ", _result.Value)
	return
}
