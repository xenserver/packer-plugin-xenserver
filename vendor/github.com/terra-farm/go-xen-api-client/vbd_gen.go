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

type VbdOperations string

const (
	// Attempting to attach this VBD to a VM
	VbdOperationsAttach VbdOperations = "attach"
	// Attempting to eject the media from this VBD
	VbdOperationsEject VbdOperations = "eject"
	// Attempting to insert new media into this VBD
	VbdOperationsInsert VbdOperations = "insert"
	// Attempting to hotplug this VBD
	VbdOperationsPlug VbdOperations = "plug"
	// Attempting to hot unplug this VBD
	VbdOperationsUnplug VbdOperations = "unplug"
	// Attempting to forcibly unplug this VBD
	VbdOperationsUnplugForce VbdOperations = "unplug_force"
	// Attempting to pause a block device backend
	VbdOperationsPause VbdOperations = "pause"
	// Attempting to unpause a block device backend
	VbdOperationsUnpause VbdOperations = "unpause"
)

type VbdType string

const (
	// VBD will appear to guest as CD
	VbdTypeCD VbdType = "CD"
	// VBD will appear to guest as disk
	VbdTypeDisk VbdType = "Disk"
	// VBD will appear as a floppy
	VbdTypeFloppy VbdType = "Floppy"
)

type VbdMode string

const (
	// only read-only access will be allowed
	VbdModeRO VbdMode = "RO"
	// read-write access will be allowed
	VbdModeRW VbdMode = "RW"
)

type VBDRecord struct {
	// Unique identifier/object reference
	UUID string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VbdOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VbdOperations
	// the virtual machine
	VM VMRef
	// the virtual disk
	VDI VDIRef
	// device seen by the guest e.g. hda1
	Device string
	// user-friendly device name e.g. 0,1,2,etc.
	Userdevice string
	// true if this VBD is bootable
	Bootable bool
	// the mode the VBD should be mounted with
	Mode VbdMode
	// how the VBD will appear to the guest (e.g. disk or CD)
	Type VbdType
	// true if this VBD will support hot-unplug
	Unpluggable bool
	// true if a storage level lock was acquired
	StorageLock bool
	// if true this represents an empty drive
	Empty bool
	// additional configuration
	OtherConfig map[string]string
	// is the device currently attached (erased on reboot)
	CurrentlyAttached bool
	// error/success code associated with last attach-operation (erased on reboot)
	StatusCode int
	// error/success information associated with last attach-operation status (erased on reboot)
	StatusDetail string
	// Device runtime properties
	RuntimeProperties map[string]string
	// QoS algorithm to use
	QosAlgorithmType string
	// parameters for chosen QoS algorithm
	QosAlgorithmParams map[string]string
	// supported QoS algorithms for this VBD
	QosSupportedAlgorithms []string
	// metrics associated with this VBD
	Metrics VBDMetricsRef
}

type VBDRef string

// A virtual block device
type VBDClass struct {
	client *Client
}

// GetAllRecords Return a map of VBD references to VBD records for all VBDs known to the system.
func (_class VBDClass) GetAllRecords(sessionID SessionRef) (_retval map[VBDRef]VBDRecord, _err error) {
	_method := "VBD.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefToVBDRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VBDs known to the system.
func (_class VBDClass) GetAll(sessionID SessionRef) (_retval []VBDRef, _err error) {
	_method := "VBD.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}

// SetMode Sets the mode of the VBD. The power_state of the VM must be halted.
func (_class VBDClass) SetMode(sessionID SessionRef, self VBDRef, value VbdMode) (_err error) {
	_method := "VBD.set_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVbdModeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// AssertAttachable Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
func (_class VBDClass) AssertAttachable(sessionID SessionRef, self VBDRef) (_err error) {
	_method := "VBD.assert_attachable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// UnplugForce Forcibly unplug the specified VBD
func (_class VBDClass) UnplugForce(sessionID SessionRef, self VBDRef) (_err error) {
	_method := "VBD.unplug_force"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Unplug Hot-unplug the specified VBD, dynamically unattaching it from the running VM
//
// Errors:
//  DEVICE_DETACH_REJECTED - The VM rejected the attempt to detach the device.
//  DEVICE_ALREADY_DETACHED - The device is not currently attached
func (_class VBDClass) Unplug(sessionID SessionRef, self VBDRef) (_err error) {
	_method := "VBD.unplug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Plug Hotplug the specified VBD, dynamically attaching it to the running VM
func (_class VBDClass) Plug(sessionID SessionRef, self VBDRef) (_err error) {
	_method := "VBD.plug"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Insert Insert new media into the device
//
// Errors:
//  VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
//  VBD_NOT_EMPTY - Operation could not be performed because the drive is not empty
func (_class VBDClass) Insert(sessionID SessionRef, vbd VBDRef, vdi VDIRef) (_err error) {
	_method := "VBD.insert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vbdArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "vbd"), vbd)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vbdArg, _vdiArg)
	return
}

// Eject Remove the media from the device and leave it empty
//
// Errors:
//  VBD_NOT_REMOVABLE_MEDIA - Media could not be ejected because it is not removable
//  VBD_IS_EMPTY - Operation could not be performed because the drive is empty
func (_class VBDClass) Eject(sessionID SessionRef, vbd VBDRef) (_err error) {
	_method := "VBD.eject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vbdArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "vbd"), vbd)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vbdArg)
	return
}

// RemoveFromQosAlgorithmParams Remove the given key and its corresponding value from the qos/algorithm_params field of the given VBD.  If the key is not in that Map, then do nothing.
func (_class VBDClass) RemoveFromQosAlgorithmParams(sessionID SessionRef, self VBDRef, key string) (_err error) {
	_method := "VBD.remove_from_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToQosAlgorithmParams Add the given key-value pair to the qos/algorithm_params field of the given VBD.
func (_class VBDClass) AddToQosAlgorithmParams(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	_method := "VBD.add_to_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetQosAlgorithmParams Set the qos/algorithm_params field of the given VBD.
func (_class VBDClass) SetQosAlgorithmParams(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	_method := "VBD.set_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetQosAlgorithmType Set the qos/algorithm_type field of the given VBD.
func (_class VBDClass) SetQosAlgorithmType(sessionID SessionRef, self VBDRef, value string) (_err error) {
	_method := "VBD.set_qos_algorithm_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromOtherConfig Remove the given key and its corresponding value from the other_config field of the given VBD.  If the key is not in that Map, then do nothing.
func (_class VBDClass) RemoveFromOtherConfig(sessionID SessionRef, self VBDRef, key string) (_err error) {
	_method := "VBD.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToOtherConfig Add the given key-value pair to the other_config field of the given VBD.
func (_class VBDClass) AddToOtherConfig(sessionID SessionRef, self VBDRef, key string, value string) (_err error) {
	_method := "VBD.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetOtherConfig Set the other_config field of the given VBD.
func (_class VBDClass) SetOtherConfig(sessionID SessionRef, self VBDRef, value map[string]string) (_err error) {
	_method := "VBD.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetUnpluggable Set the unpluggable field of the given VBD.
func (_class VBDClass) SetUnpluggable(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	_method := "VBD.set_unpluggable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetType Set the type field of the given VBD.
func (_class VBDClass) SetType(sessionID SessionRef, self VBDRef, value VbdType) (_err error) {
	_method := "VBD.set_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVbdTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetBootable Set the bootable field of the given VBD.
func (_class VBDClass) SetBootable(sessionID SessionRef, self VBDRef, value bool) (_err error) {
	_method := "VBD.set_bootable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetUserdevice Set the userdevice field of the given VBD.
func (_class VBDClass) SetUserdevice(sessionID SessionRef, self VBDRef, value string) (_err error) {
	_method := "VBD.set_userdevice"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMetrics Get the metrics field of the given VBD.
func (_class VBDClass) GetMetrics(sessionID SessionRef, self VBDRef) (_retval VBDMetricsRef, _err error) {
	_method := "VBD.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

// GetQosSupportedAlgorithms Get the qos/supported_algorithms field of the given VBD.
func (_class VBDClass) GetQosSupportedAlgorithms(sessionID SessionRef, self VBDRef) (_retval []string, _err error) {
	_method := "VBD.get_qos_supported_algorithms"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetQosAlgorithmParams Get the qos/algorithm_params field of the given VBD.
func (_class VBDClass) GetQosAlgorithmParams(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	_method := "VBD.get_qos_algorithm_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetQosAlgorithmType Get the qos/algorithm_type field of the given VBD.
func (_class VBDClass) GetQosAlgorithmType(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	_method := "VBD.get_qos_algorithm_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetRuntimeProperties Get the runtime_properties field of the given VBD.
func (_class VBDClass) GetRuntimeProperties(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	_method := "VBD.get_runtime_properties"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetStatusDetail Get the status_detail field of the given VBD.
func (_class VBDClass) GetStatusDetail(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	_method := "VBD.get_status_detail"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetStatusCode Get the status_code field of the given VBD.
func (_class VBDClass) GetStatusCode(sessionID SessionRef, self VBDRef) (_retval int, _err error) {
	_method := "VBD.get_status_code"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetCurrentlyAttached Get the currently_attached field of the given VBD.
func (_class VBDClass) GetCurrentlyAttached(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	_method := "VBD.get_currently_attached"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetOtherConfig Get the other_config field of the given VBD.
func (_class VBDClass) GetOtherConfig(sessionID SessionRef, self VBDRef) (_retval map[string]string, _err error) {
	_method := "VBD.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetEmpty Get the empty field of the given VBD.
func (_class VBDClass) GetEmpty(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	_method := "VBD.get_empty"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetStorageLock Get the storage_lock field of the given VBD.
func (_class VBDClass) GetStorageLock(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	_method := "VBD.get_storage_lock"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUnpluggable Get the unpluggable field of the given VBD.
func (_class VBDClass) GetUnpluggable(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	_method := "VBD.get_unpluggable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetType Get the type field of the given VBD.
func (_class VBDClass) GetType(sessionID SessionRef, self VBDRef) (_retval VbdType, _err error) {
	_method := "VBD.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVbdTypeToGo(_method + " -> ", _result.Value)
	return
}

// GetMode Get the mode field of the given VBD.
func (_class VBDClass) GetMode(sessionID SessionRef, self VBDRef) (_retval VbdMode, _err error) {
	_method := "VBD.get_mode"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVbdModeToGo(_method + " -> ", _result.Value)
	return
}

// GetBootable Get the bootable field of the given VBD.
func (_class VBDClass) GetBootable(sessionID SessionRef, self VBDRef) (_retval bool, _err error) {
	_method := "VBD.get_bootable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUserdevice Get the userdevice field of the given VBD.
func (_class VBDClass) GetUserdevice(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	_method := "VBD.get_userdevice"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetDevice Get the device field of the given VBD.
func (_class VBDClass) GetDevice(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	_method := "VBD.get_device"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVDI Get the VDI field of the given VBD.
func (_class VBDClass) GetVDI(sessionID SessionRef, self VBDRef) (_retval VDIRef, _err error) {
	_method := "VBD.get_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// GetVM Get the VM field of the given VBD.
func (_class VBDClass) GetVM(sessionID SessionRef, self VBDRef) (_retval VMRef, _err error) {
	_method := "VBD.get_VM"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentOperations Get the current_operations field of the given VBD.
func (_class VBDClass) GetCurrentOperations(sessionID SessionRef, self VBDRef) (_retval map[string]VbdOperations, _err error) {
	_method := "VBD.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVbdOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAllowedOperations Get the allowed_operations field of the given VBD.
func (_class VBDClass) GetAllowedOperations(sessionID SessionRef, self VBDRef) (_retval []VbdOperations, _err error) {
	_method := "VBD.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVbdOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

// GetUUID Get the uuid field of the given VBD.
func (_class VBDClass) GetUUID(sessionID SessionRef, self VBDRef) (_retval string, _err error) {
	_method := "VBD.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Destroy Destroy the specified VBD instance.
func (_class VBDClass) Destroy(sessionID SessionRef, self VBDRef) (_err error) {
	_method := "VBD.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new VBD instance, and return its handle. The constructor args are: VM*, VDI*, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
func (_class VBDClass) Create(sessionID SessionRef, args VBDRecord) (_retval VBDRef, _err error) {
	_method := "VBD.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVBDRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the VBD instance with the specified UUID.
func (_class VBDClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VBDRef, _err error) {
	_method := "VBD.get_by_uuid"
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
	_retval, _err = convertVBDRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VBD.
func (_class VBDClass) GetRecord(sessionID SessionRef, self VBDRef) (_retval VBDRecord, _err error) {
	_method := "VBD.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVBDRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRecordToGo(_method + " -> ", _result.Value)
	return
}
