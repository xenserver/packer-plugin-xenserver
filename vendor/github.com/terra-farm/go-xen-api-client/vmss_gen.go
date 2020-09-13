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

type VmssFrequency string

const (
	// Hourly snapshots
	VmssFrequencyHourly VmssFrequency = "hourly"
	// Daily snapshots
	VmssFrequencyDaily VmssFrequency = "daily"
	// Weekly snapshots
	VmssFrequencyWeekly VmssFrequency = "weekly"
)

type VmssType string

const (
	// The snapshot is a disk snapshot
	VmssTypeSnapshot VmssType = "snapshot"
	// The snapshot is a checkpoint
	VmssTypeCheckpoint VmssType = "checkpoint"
	// The snapshot is a VSS
	VmssTypeSnapshotWithQuiesce VmssType = "snapshot_with_quiesce"
)

type VMSSRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// enable or disable this snapshot schedule
	Enabled bool
	// type of the snapshot schedule
	Type VmssType
	// maximum number of snapshots that should be stored at any time
	RetainedSnapshots int
	// frequency of taking snapshot from snapshot schedule
	Frequency VmssFrequency
	// schedule of the snapshot containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	Schedule map[string]string
	// time of the last snapshot
	LastRunTime time.Time
	// all VMs attached to this snapshot schedule
	VMs []VMRef
}

type VMSSRef string

// VM Snapshot Schedule
type VMSSClass struct {
	client *Client
}

// GetAllRecords Return a map of VMSS references to VMSS records for all VMSSs known to the system.
func (_class VMSSClass) GetAllRecords(sessionID SessionRef) (_retval map[VMSSRef]VMSSRecord, _err error) {
	_method := "VMSS.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefToVMSSRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VMSSs known to the system.
func (_class VMSSClass) GetAll(sessionID SessionRef) (_retval []VMSSRef, _err error) {
	_method := "VMSS.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefSetToGo(_method + " -> ", _result.Value)
	return
}

// SetType 
func (_class VMSSClass) SetType(sessionID SessionRef, self VMSSRef, value VmssType) (_err error) {
	_method := "VMSS.set_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmssTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetLastRunTime 
func (_class VMSSClass) SetLastRunTime(sessionID SessionRef, self VMSSRef, value time.Time) (_err error) {
	_method := "VMSS.set_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveFromSchedule 
func (_class VMSSClass) RemoveFromSchedule(sessionID SessionRef, self VMSSRef, key string) (_err error) {
	_method := "VMSS.remove_from_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToSchedule 
func (_class VMSSClass) AddToSchedule(sessionID SessionRef, self VMSSRef, key string, value string) (_err error) {
	_method := "VMSS.add_to_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetSchedule 
func (_class VMSSClass) SetSchedule(sessionID SessionRef, self VMSSRef, value map[string]string) (_err error) {
	_method := "VMSS.set_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetFrequency Set the value of the frequency field
func (_class VMSSClass) SetFrequency(sessionID SessionRef, self VMSSRef, value VmssFrequency) (_err error) {
	_method := "VMSS.set_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmssFrequencyToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetRetainedSnapshots 
func (_class VMSSClass) SetRetainedSnapshots(sessionID SessionRef, self VMSSRef, value int) (_err error) {
	_method := "VMSS.set_retained_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SnapshotNow This call executes the snapshot schedule immediately
func (_class VMSSClass) SnapshotNow(sessionID SessionRef, vmss VMSSRef) (_retval string, _err error) {
	_method := "VMSS.snapshot_now"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmssArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "vmss"), vmss)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmssArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// SetEnabled Set the enabled field of the given VMSS.
func (_class VMSSClass) SetEnabled(sessionID SessionRef, self VMSSRef, value bool) (_err error) {
	_method := "VMSS.set_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameDescription Set the name/description field of the given VMSS.
func (_class VMSSClass) SetNameDescription(sessionID SessionRef, self VMSSRef, value string) (_err error) {
	_method := "VMSS.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetNameLabel Set the name/label field of the given VMSS.
func (_class VMSSClass) SetNameLabel(sessionID SessionRef, self VMSSRef, value string) (_err error) {
	_method := "VMSS.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVMs Get the VMs field of the given VMSS.
func (_class VMSSClass) GetVMs(sessionID SessionRef, self VMSSRef) (_retval []VMRef, _err error) {
	_method := "VMSS.get_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetLastRunTime Get the last_run_time field of the given VMSS.
func (_class VMSSClass) GetLastRunTime(sessionID SessionRef, self VMSSRef) (_retval time.Time, _err error) {
	_method := "VMSS.get_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSchedule Get the schedule field of the given VMSS.
func (_class VMSSClass) GetSchedule(sessionID SessionRef, self VMSSRef) (_retval map[string]string, _err error) {
	_method := "VMSS.get_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetFrequency Get the frequency field of the given VMSS.
func (_class VMSSClass) GetFrequency(sessionID SessionRef, self VMSSRef) (_retval VmssFrequency, _err error) {
	_method := "VMSS.get_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmssFrequencyToGo(_method + " -> ", _result.Value)
	return
}

// GetRetainedSnapshots Get the retained_snapshots field of the given VMSS.
func (_class VMSSClass) GetRetainedSnapshots(sessionID SessionRef, self VMSSRef) (_retval int, _err error) {
	_method := "VMSS.get_retained_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetType Get the type field of the given VMSS.
func (_class VMSSClass) GetType(sessionID SessionRef, self VMSSRef) (_retval VmssType, _err error) {
	_method := "VMSS.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmssTypeToGo(_method + " -> ", _result.Value)
	return
}

// GetEnabled Get the enabled field of the given VMSS.
func (_class VMSSClass) GetEnabled(sessionID SessionRef, self VMSSRef) (_retval bool, _err error) {
	_method := "VMSS.get_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameDescription Get the name/description field of the given VMSS.
func (_class VMSSClass) GetNameDescription(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	_method := "VMSS.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given VMSS.
func (_class VMSSClass) GetNameLabel(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	_method := "VMSS.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given VMSS.
func (_class VMSSClass) GetUUID(sessionID SessionRef, self VMSSRef) (_retval string, _err error) {
	_method := "VMSS.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the VMSS instances with the given label.
func (_class VMSSClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMSSRef, _err error) {
	_method := "VMSS.get_by_name_label"
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
	_retval, _err = convertVMSSRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the specified VMSS instance.
func (_class VMSSClass) Destroy(sessionID SessionRef, self VMSSRef) (_err error) {
	_method := "VMSS.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new VMSS instance, and return its handle. The constructor args are: name_label, name_description, enabled, type*, retained_snapshots, frequency*, schedule (* = non-optional).
func (_class VMSSClass) Create(sessionID SessionRef, args VMSSRecord) (_retval VMSSRef, _err error) {
	_method := "VMSS.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMSSRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the VMSS instance with the specified UUID.
func (_class VMSSClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMSSRef, _err error) {
	_method := "VMSS.get_by_uuid"
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
	_retval, _err = convertVMSSRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VMSS.
func (_class VMSSClass) GetRecord(sessionID SessionRef, self VMSSRef) (_retval VMSSRecord, _err error) {
	_method := "VMSS.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMSSRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMSSRecordToGo(_method + " -> ", _result.Value)
	return
}
