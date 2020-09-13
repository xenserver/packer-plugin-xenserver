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

type SessionRecord struct {
	// Unique identifier/object reference
	UUID string
	// Currently connected host
	ThisHost HostRef
	// Currently connected user
	ThisUser UserRef
	// Timestamp for last time session was active
	LastActive time.Time
	// True if this session relates to a intra-pool login, false otherwise
	Pool bool
	// additional configuration
	OtherConfig map[string]string
	// true iff this session was created using local superuser credentials
	IsLocalSuperuser bool
	// references the subject instance that created the session. If a session instance has is_local_superuser set, then the value of this field is undefined.
	Subject SubjectRef
	// time when session was last validated
	ValidationTime time.Time
	// the subject identifier of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserSid string
	// the subject name of the user that was externally authenticated. If a session instance has is_local_superuser set, then the value of this field is undefined.
	AuthUserName string
	// list with all RBAC permissions for this session
	RbacPermissions []string
	// list of tasks created using the current session
	Tasks []TaskRef
	// references the parent session that created this session
	Parent SessionRef
	// a key string provided by a API user to distinguish itself from other users sharing the same login name
	Originator string
}

type SessionRef string

// A session
type SessionClass struct {
	client *Client
}

// LogoutSubjectIdentifier Log out all sessions associated to a user subject-identifier, except the session associated with the context calling this function
func (_class SessionClass) LogoutSubjectIdentifier(sessionID SessionRef, subjectIdentifier string) (_err error) {
	_method := "session.logout_subject_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_subjectIdentifierArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "subject_identifier"), subjectIdentifier)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _subjectIdentifierArg)
	return
}

// GetAllSubjectIdentifiers Return a list of all the user subject-identifiers of all existing sessions
func (_class SessionClass) GetAllSubjectIdentifiers(sessionID SessionRef) (_retval []string, _err error) {
	_method := "session.get_all_subject_identifiers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// LocalLogout Log out of local session.
func (_class SessionClass) LocalLogout(sessionID SessionRef) (_err error) {
	_method := "session.local_logout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// CreateFromDbFile 
func (_class SessionClass) CreateFromDbFile(sessionID SessionRef, filename string) (_retval SessionRef, _err error) {
	_method := "session.create_from_db_file"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_filenameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "filename"), filename)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _filenameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

// SlaveLocalLoginWithPassword Authenticate locally against a slave in emergency mode. Note the resulting sessions are only good for use on this host.
func (_class SessionClass) SlaveLocalLoginWithPassword(uname string, pwd string) (_retval SessionRef, _err error) {
	_method := "session.slave_local_login_with_password"
	_unameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uname"), uname)
	if _err != nil {
		return
	}
	_pwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "pwd"), pwd)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _unameArg, _pwdArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

// ChangePassword Change the account password; if your session is authenticated with root priviledges then the old_pwd is validated and the new_pwd is set regardless
func (_class SessionClass) ChangePassword(sessionID SessionRef, oldPwd string, newPwd string) (_err error) {
	_method := "session.change_password"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_oldPwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "old_pwd"), oldPwd)
	if _err != nil {
		return
	}
	_newPwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_pwd"), newPwd)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _oldPwdArg, _newPwdArg)
	return
}

// Logout Log out of a session
func (_class SessionClass) Logout(sessionID SessionRef) (_err error) {
	_method := "session.logout"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg)
	return
}

// LoginWithPassword Attempt to authenticate the user, returning a session reference if successful
//
// Errors:
//  SESSION_AUTHENTICATION_FAILED - The credentials given by the user are incorrect, so access has been denied, and you have not been issued a session handle.
//  HOST_IS_SLAVE - You cannot make regular API calls directly on a slave. Please pass API calls via the master host.
func (_class SessionClass) LoginWithPassword(uname string, pwd string, version string, originator string) (_retval SessionRef, _err error) {
	_method := "session.login_with_password"
	_unameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uname"), uname)
	if _err != nil {
		return
	}
	_pwdArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "pwd"), pwd)
	if _err != nil {
		return
	}
	_versionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "version"), version)
	if _err != nil {
		return
	}
	_originatorArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "originator"), originator)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _unameArg, _pwdArg, _versionArg, _originatorArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

// RemoveFromOtherConfig Remove the given key and its corresponding value from the other_config field of the given session.  If the key is not in that Map, then do nothing.
func (_class SessionClass) RemoveFromOtherConfig(sessionID SessionRef, self SessionRef, key string) (_err error) {
	_method := "session.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// AddToOtherConfig Add the given key-value pair to the other_config field of the given session.
func (_class SessionClass) AddToOtherConfig(sessionID SessionRef, self SessionRef, key string, value string) (_err error) {
	_method := "session.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// SetOtherConfig Set the other_config field of the given session.
func (_class SessionClass) SetOtherConfig(sessionID SessionRef, self SessionRef, value map[string]string) (_err error) {
	_method := "session.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// GetOriginator Get the originator field of the given session.
func (_class SessionClass) GetOriginator(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_originator"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetParent Get the parent field of the given session.
func (_class SessionClass) GetParent(sessionID SessionRef, self SessionRef) (_retval SessionRef, _err error) {
	_method := "session.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

// GetTasks Get the tasks field of the given session.
func (_class SessionClass) GetTasks(sessionID SessionRef, self SessionRef) (_retval []TaskRef, _err error) {
	_method := "session.get_tasks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetRbacPermissions Get the rbac_permissions field of the given session.
func (_class SessionClass) GetRbacPermissions(sessionID SessionRef, self SessionRef) (_retval []string, _err error) {
	_method := "session.get_rbac_permissions"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetAuthUserName Get the auth_user_name field of the given session.
func (_class SessionClass) GetAuthUserName(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_auth_user_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetAuthUserSid Get the auth_user_sid field of the given session.
func (_class SessionClass) GetAuthUserSid(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_auth_user_sid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetValidationTime Get the validation_time field of the given session.
func (_class SessionClass) GetValidationTime(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	_method := "session.get_validation_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}

// GetSubject Get the subject field of the given session.
func (_class SessionClass) GetSubject(sessionID SessionRef, self SessionRef) (_retval SubjectRef, _err error) {
	_method := "session.get_subject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSubjectRefToGo(_method + " -> ", _result.Value)
	return
}

// GetIsLocalSuperuser Get the is_local_superuser field of the given session.
func (_class SessionClass) GetIsLocalSuperuser(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	_method := "session.get_is_local_superuser"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetOtherConfig Get the other_config field of the given session.
func (_class SessionClass) GetOtherConfig(sessionID SessionRef, self SessionRef) (_retval map[string]string, _err error) {
	_method := "session.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetPool Get the pool field of the given session.
func (_class SessionClass) GetPool(sessionID SessionRef, self SessionRef) (_retval bool, _err error) {
	_method := "session.get_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetLastActive Get the last_active field of the given session.
func (_class SessionClass) GetLastActive(sessionID SessionRef, self SessionRef) (_retval time.Time, _err error) {
	_method := "session.get_last_active"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}

// GetThisUser Get the this_user field of the given session.
func (_class SessionClass) GetThisUser(sessionID SessionRef, self SessionRef) (_retval UserRef, _err error) {
	_method := "session.get_this_user"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertUserRefToGo(_method + " -> ", _result.Value)
	return
}

// GetThisHost Get the this_host field of the given session.
func (_class SessionClass) GetThisHost(sessionID SessionRef, self SessionRef) (_retval HostRef, _err error) {
	_method := "session.get_this_host"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given session.
func (_class SessionClass) GetUUID(sessionID SessionRef, self SessionRef) (_retval string, _err error) {
	_method := "session.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the session instance with the specified UUID.
func (_class SessionClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SessionRef, _err error) {
	_method := "session.get_by_uuid"
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
	_retval, _err = convertSessionRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given session.
func (_class SessionClass) GetRecord(sessionID SessionRef, self SessionRef) (_retval SessionRecord, _err error) {
	_method := "session.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSessionRecordToGo(_method + " -> ", _result.Value)
	return
}
