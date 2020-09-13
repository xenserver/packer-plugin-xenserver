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

type SubjectRecord struct {
	// Unique identifier/object reference
	UUID string
	// the subject identifier, unique in the external directory service
	SubjectIdentifier string
	// additional configuration
	OtherConfig map[string]string
	// the roles associated with this subject
	Roles []RoleRef
}

type SubjectRef string

// A user or group that can log in xapi
type SubjectClass struct {
	client *Client
}

// GetAllRecords Return a map of subject references to subject records for all subjects known to the system.
func (_class SubjectClass) GetAllRecords(sessionID SessionRef) (_retval map[SubjectRef]SubjectRecord, _err error) {
	_method := "subject.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToSubjectRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the subjects known to the system.
func (_class SubjectClass) GetAll(sessionID SessionRef) (_retval []SubjectRef, _err error) {
	_method := "subject.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetPermissionsNameLabel This call returns a list of permission names given a subject
func (_class SubjectClass) GetPermissionsNameLabel(sessionID SessionRef, self SubjectRef) (_retval []string, _err error) {
	_method := "subject.get_permissions_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromRoles This call removes a role from a subject
func (_class SubjectClass) RemoveFromRoles(sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	_method := "subject.remove_from_roles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_roleArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "role"), role)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _roleArg)
	return
}

// AddToRoles This call adds a new role to a subject
func (_class SubjectClass) AddToRoles(sessionID SessionRef, self SubjectRef, role RoleRef) (_err error) {
	_method := "subject.add_to_roles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_roleArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "role"), role)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _roleArg)
	return
}

// GetRoles Get the roles field of the given subject.
func (_class SubjectClass) GetRoles(sessionID SessionRef, self SubjectRef) (_retval []RoleRef, _err error) {
	_method := "subject.get_roles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetOtherConfig Get the other_config field of the given subject.
func (_class SubjectClass) GetOtherConfig(sessionID SessionRef, self SubjectRef) (_retval map[string]string, _err error) {
	_method := "subject.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// GetSubjectIdentifier Get the subject_identifier field of the given subject.
func (_class SubjectClass) GetSubjectIdentifier(sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	_method := "subject.get_subject_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given subject.
func (_class SubjectClass) GetUUID(sessionID SessionRef, self SubjectRef) (_retval string, _err error) {
	_method := "subject.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Destroy Destroy the specified subject instance.
func (_class SubjectClass) Destroy(sessionID SessionRef, self SubjectRef) (_err error) {
	_method := "subject.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new subject instance, and return its handle. The constructor args are: subject_identifier, other_config (* = non-optional).
func (_class SubjectClass) Create(sessionID SessionRef, args SubjectRecord) (_retval SubjectRef, _err error) {
	_method := "subject.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertSubjectRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the subject instance with the specified UUID.
func (_class SubjectClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SubjectRef, _err error) {
	_method := "subject.get_by_uuid"
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
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given subject.
func (_class SubjectClass) GetRecord(sessionID SessionRef, self SubjectRef) (_retval SubjectRecord, _err error) {
	_method := "subject.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSubjectRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRecordToGo(_method + " -> ", _result.Value)
	return
}
