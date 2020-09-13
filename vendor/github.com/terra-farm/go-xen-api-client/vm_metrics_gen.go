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

type VMMetricsRecord struct {
	// Unique identifier/object reference
	UUID string
	// Guest's actual memory (bytes)
	MemoryActual int
	// Current number of VCPUs
	VCPUsNumber int
	// Utilisation for all of guest's current VCPUs
	VCPUsUtilisation map[int]float64
	// VCPU to PCPU map
	VCPUsCPU map[int]int
	// The live equivalent to VM.VCPUs_params
	VCPUsParams map[string]string
	// CPU flags (blocked,online,running)
	VCPUsFlags map[int][]string
	// The state of the guest, eg blocked, dying etc
	State []string
	// Time at which this VM was last booted
	StartTime time.Time
	// Time at which the VM was installed
	InstallTime time.Time
	// Time at which this information was last updated
	LastUpdated time.Time
	// additional configuration
	OtherConfig map[string]string
	// hardware virtual machine
	Hvm bool
	// VM supports nested virtualisation
	NestedVirt bool
	// VM is immobile and can't migrate between hosts
	Nomigrate bool
}

type VMMetricsRef string

// The metrics associated with a VM
type VMMetricsClass struct {
	client *Client
}

// GetAllRecords Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system.
func (_class VMMetricsClass) GetAllRecords(sessionID SessionRef) (_retval map[VMMetricsRef]VMMetricsRecord, _err error) {
	_method := "VM_metrics.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefToVMMetricsRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VM_metrics instances known to the system.
func (_class VMMetricsClass) GetAll(sessionID SessionRef) (_retval []VMMetricsRef, _err error) {
	_method := "VM_metrics.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefSetToGo(_method + " -> ", _result.Value)
	return
}

// RemoveFromOtherConfig Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
func (_class VMMetricsClass) RemoveFromOtherConfig(sessionID SessionRef, self VMMetricsRef, key string) (_err error) {
	_method := "VM_metrics.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToOtherConfig Add the given key-value pair to the other_config field of the given VM_metrics.
func (_class VMMetricsClass) AddToOtherConfig(sessionID SessionRef, self VMMetricsRef, key string, value string) (_err error) {
	_method := "VM_metrics.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetOtherConfig Set the other_config field of the given VM_metrics.
func (_class VMMetricsClass) SetOtherConfig(sessionID SessionRef, self VMMetricsRef, value map[string]string) (_err error) {
	_method := "VM_metrics.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNomigrate Get the nomigrate field of the given VM_metrics.
func (_class VMMetricsClass) GetNomigrate(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	_method := "VM_metrics.get_nomigrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNestedVirt Get the nested_virt field of the given VM_metrics.
func (_class VMMetricsClass) GetNestedVirt(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	_method := "VM_metrics.get_nested_virt"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetHvm Get the hvm field of the given VM_metrics.
func (_class VMMetricsClass) GetHvm(sessionID SessionRef, self VMMetricsRef) (_retval bool, _err error) {
	_method := "VM_metrics.get_hvm"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetOtherConfig Get the other_config field of the given VM_metrics.
func (_class VMMetricsClass) GetOtherConfig(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	_method := "VM_metrics.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetLastUpdated Get the last_updated field of the given VM_metrics.
func (_class VMMetricsClass) GetLastUpdated(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	_method := "VM_metrics.get_last_updated"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetInstallTime Get the install_time field of the given VM_metrics.
func (_class VMMetricsClass) GetInstallTime(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	_method := "VM_metrics.get_install_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetStartTime Get the start_time field of the given VM_metrics.
func (_class VMMetricsClass) GetStartTime(sessionID SessionRef, self VMMetricsRef) (_retval time.Time, _err error) {
	_method := "VM_metrics.get_start_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetState Get the state field of the given VM_metrics.
func (_class VMMetricsClass) GetState(sessionID SessionRef, self VMMetricsRef) (_retval []string, _err error) {
	_method := "VM_metrics.get_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVCPUsFlags Get the VCPUs/flags field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsFlags(sessionID SessionRef, self VMMetricsRef) (_retval map[int][]string, _err error) {
	_method := "VM_metrics.get_VCPUs_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}

// GetVCPUsParams Get the VCPUs/params field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsParams(sessionID SessionRef, self VMMetricsRef) (_retval map[string]string, _err error) {
	_method := "VM_metrics.get_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVCPUsCPU Get the VCPUs/CPU field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsCPU(sessionID SessionRef, self VMMetricsRef) (_retval map[int]int, _err error) {
	_method := "VM_metrics.get_VCPUs_CPU"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToIntMapToGo(_method + " -> ", _result.Value)
	return
}

// GetVCPUsUtilisation Get the VCPUs/utilisation field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsUtilisation(sessionID SessionRef, self VMMetricsRef) (_retval map[int]float64, _err error) {
	_method := "VM_metrics.get_VCPUs_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToFloatMapToGo(_method + " -> ", _result.Value)
	return
}

// GetVCPUsNumber Get the VCPUs/number field of the given VM_metrics.
func (_class VMMetricsClass) GetVCPUsNumber(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	_method := "VM_metrics.get_VCPUs_number"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMemoryActual Get the memory/actual field of the given VM_metrics.
func (_class VMMetricsClass) GetMemoryActual(sessionID SessionRef, self VMMetricsRef) (_retval int, _err error) {
	_method := "VM_metrics.get_memory_actual"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given VM_metrics.
func (_class VMMetricsClass) GetUUID(sessionID SessionRef, self VMMetricsRef) (_retval string, _err error) {
	_method := "VM_metrics.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the VM_metrics instance with the specified UUID.
func (_class VMMetricsClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMMetricsRef, _err error) {
	_method := "VM_metrics.get_by_uuid"
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
	_retval, _err = convertVMMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VM_metrics.
func (_class VMMetricsClass) GetRecord(sessionID SessionRef, self VMMetricsRef) (_retval VMMetricsRecord, _err error) {
	_method := "VM_metrics.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMMetricsRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRecordToGo(_method + " -> ", _result.Value)
	return
}
