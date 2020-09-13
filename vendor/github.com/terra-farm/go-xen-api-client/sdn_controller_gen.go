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

type SdnControllerProtocol string

const (
	// Active ssl connection
	SdnControllerProtocolSsl SdnControllerProtocol = "ssl"
	// Passive ssl connection
	SdnControllerProtocolPssl SdnControllerProtocol = "pssl"
)

type SDNControllerRecord struct {
	// Unique identifier/object reference
	UUID string
	// Protocol to connect with SDN controller
	Protocol SdnControllerProtocol
	// IP address of the controller
	Address string
	// TCP port of the controller
	Port int
}

type SDNControllerRef string

// Describes the SDN controller that is to connect with the pool
type SDNControllerClass struct {
	client *Client
}

// GetAllRecords Return a map of SDN_controller references to SDN_controller records for all SDN_controllers known to the system.
func (_class SDNControllerClass) GetAllRecords(sessionID SessionRef) (_retval map[SDNControllerRef]SDNControllerRecord, _err error) {
	_method := "SDN_controller.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRefToSDNControllerRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the SDN_controllers known to the system.
func (_class SDNControllerClass) GetAll(sessionID SessionRef) (_retval []SDNControllerRef, _err error) {
	_method := "SDN_controller.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Forget Remove the OVS manager of the pool and destroy the db record.
func (_class SDNControllerClass) Forget(sessionID SessionRef, self SDNControllerRef) (_err error) {
	_method := "SDN_controller.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Introduce Introduce an SDN controller to the pool.
func (_class SDNControllerClass) Introduce(sessionID SessionRef, protocol SdnControllerProtocol, address string, port int) (_retval SDNControllerRef, _err error) {
	_method := "SDN_controller.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_protocolArg, _err := convertEnumSdnControllerProtocolToXen(fmt.Sprintf("%s(%s)", _method, "protocol"), protocol)
	if _err != nil {
		return
	}
	_addressArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "address"), address)
	if _err != nil {
		return
	}
	_portArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "port"), port)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _protocolArg, _addressArg, _portArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRefToGo(_method + " -> ", _result.Value)
	return
}

// GetPort Get the port field of the given SDN_controller.
func (_class SDNControllerClass) GetPort(sessionID SessionRef, self SDNControllerRef) (_retval int, _err error) {
	_method := "SDN_controller.get_port"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetAddress Get the address field of the given SDN_controller.
func (_class SDNControllerClass) GetAddress(sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	_method := "SDN_controller.get_address"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetProtocol Get the protocol field of the given SDN_controller.
func (_class SDNControllerClass) GetProtocol(sessionID SessionRef, self SDNControllerRef) (_retval SdnControllerProtocol, _err error) {
	_method := "SDN_controller.get_protocol"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumSdnControllerProtocolToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given SDN_controller.
func (_class SDNControllerClass) GetUUID(sessionID SessionRef, self SDNControllerRef) (_retval string, _err error) {
	_method := "SDN_controller.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the SDN_controller instance with the specified UUID.
func (_class SDNControllerClass) GetByUUID(sessionID SessionRef, uuid string) (_retval SDNControllerRef, _err error) {
	_method := "SDN_controller.get_by_uuid"
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
	_retval, _err = convertSDNControllerRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given SDN_controller.
func (_class SDNControllerClass) GetRecord(sessionID SessionRef, self SDNControllerRef) (_retval SDNControllerRecord, _err error) {
	_method := "SDN_controller.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertSDNControllerRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSDNControllerRecordToGo(_method + " -> ", _result.Value)
	return
}
