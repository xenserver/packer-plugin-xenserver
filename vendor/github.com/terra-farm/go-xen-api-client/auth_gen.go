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

type AuthRef string

// Management of remote authentication services
type AuthClass struct {
	client *Client
}

// GetGroupMembership This calls queries the external directory service to obtain the transitively-closed set of groups that the the subject_identifier is member of.
func (_class AuthClass) GetGroupMembership(sessionID SessionRef, subjectIdentifier string) (_retval []string, _err error) {
	_method := "auth.get_group_membership"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_subjectIdentifierArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "subject_identifier"), subjectIdentifier)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _subjectIdentifierArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// GetSubjectInformationFromIdentifier This call queries the external directory service to obtain the user information (e.g. username, organization etc) from the specified subject_identifier
func (_class AuthClass) GetSubjectInformationFromIdentifier(sessionID SessionRef, subjectIdentifier string) (_retval map[string]string, _err error) {
	_method := "auth.get_subject_information_from_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_subjectIdentifierArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "subject_identifier"), subjectIdentifier)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _subjectIdentifierArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// GetSubjectIdentifier This call queries the external directory service to obtain the subject_identifier as a string from the human-readable subject_name
func (_class AuthClass) GetSubjectIdentifier(sessionID SessionRef, subjectName string) (_retval string, _err error) {
	_method := "auth.get_subject_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_subjectNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "subject_name"), subjectName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _subjectNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}
