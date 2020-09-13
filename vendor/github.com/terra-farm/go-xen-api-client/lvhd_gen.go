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

type LVHDRecord struct {
	// Unique identifier/object reference
	UUID string
}

type LVHDRef string

// LVHD SR specific operations
type LVHDClass struct {
	client *Client
}

// EnableThinProvisioning Upgrades an LVHD SR to enable thin-provisioning. Future VDIs created in this SR will be thinly-provisioned, although existing VDIs will be left alone. Note that the SR must be attached to the SRmaster for upgrade to work.
func (_class LVHDClass) EnableThinProvisioning(sessionID SessionRef, host HostRef, sr SRRef, initialAllocation int, allocationQuantum int) (_retval string, _err error) {
	_method := "LVHD.enable_thin_provisioning"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "SR"), sr)
	if _err != nil {
		return
	}
	_initialAllocationArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "initial_allocation"), initialAllocation)
	if _err != nil {
		return
	}
	_allocationQuantumArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "allocation_quantum"), allocationQuantum)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _hostArg, _srArg, _initialAllocationArg, _allocationQuantumArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given LVHD.
func (_class LVHDClass) GetUUID(sessionID SessionRef, self LVHDRef) (_retval string, _err error) {
	_method := "LVHD.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertLVHDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the LVHD instance with the specified UUID.
func (_class LVHDClass) GetByUUID(sessionID SessionRef, uuid string) (_retval LVHDRef, _err error) {
	_method := "LVHD.get_by_uuid"
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
	_retval, _err = convertLVHDRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given LVHD.
func (_class LVHDClass) GetRecord(sessionID SessionRef, self LVHDRef) (_retval LVHDRecord, _err error) {
	_method := "LVHD.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertLVHDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertLVHDRecordToGo(_method + " -> ", _result.Value)
	return
}
