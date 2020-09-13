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

type RoleRecord struct {
	// Unique identifier/object reference
	UUID string
	// a short user-friendly name for the role
	NameLabel string
	// what this role is for
	NameDescription string
	// a list of pointers to other roles or permissions
	Subroles []RoleRef
}

type RoleRef string

// A set of permissions associated with a subject
type RoleClass struct {
	client *Client
}

// GetAllRecords Return a map of role references to role records for all roles known to the system.
func (_class RoleClass) GetAllRecords(sessionID SessionRef) (_retval map[RoleRef]RoleRecord, _err error) {
	_method := "role.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefToRoleRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the roles known to the system.
func (_class RoleClass) GetAll(sessionID SessionRef) (_retval []RoleRef, _err error) {
	_method := "role.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetByPermissionNameLabel This call returns a list of roles given a permission name
func (_class RoleClass) GetByPermissionNameLabel(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	_method := "role.get_by_permission_name_label"
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
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetByPermission This call returns a list of roles given a permission
func (_class RoleClass) GetByPermission(sessionID SessionRef, permission RoleRef) (_retval []RoleRef, _err error) {
	_method := "role.get_by_permission"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_permissionArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "permission"), permission)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _permissionArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetPermissionsNameLabel This call returns a list of permission names given a role
func (_class RoleClass) GetPermissionsNameLabel(sessionID SessionRef, self RoleRef) (_retval []string, _err error) {
	_method := "role.get_permissions_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetPermissions This call returns a list of permissions given a role
func (_class RoleClass) GetPermissions(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	_method := "role.get_permissions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSubroles Get the subroles field of the given role.
func (_class RoleClass) GetSubroles(sessionID SessionRef, self RoleRef) (_retval []RoleRef, _err error) {
	_method := "role.get_subroles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameDescription Get the name/description field of the given role.
func (_class RoleClass) GetNameDescription(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	_method := "role.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given role.
func (_class RoleClass) GetNameLabel(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	_method := "role.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given role.
func (_class RoleClass) GetUUID(sessionID SessionRef, self RoleRef) (_retval string, _err error) {
	_method := "role.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the role instances with the given label.
func (_class RoleClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []RoleRef, _err error) {
	_method := "role.get_by_name_label"
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
	_retval, _err = convertRoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the role instance with the specified UUID.
func (_class RoleClass) GetByUUID(sessionID SessionRef, uuid string) (_retval RoleRef, _err error) {
	_method := "role.get_by_uuid"
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
	_retval, _err = convertRoleRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given role.
func (_class RoleClass) GetRecord(sessionID SessionRef, self RoleRef) (_retval RoleRecord, _err error) {
	_method := "role.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertRoleRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertRoleRecordToGo(_method + " -> ", _result.Value)
	return
}
