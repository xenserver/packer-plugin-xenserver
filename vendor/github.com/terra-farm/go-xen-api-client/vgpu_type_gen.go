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

type VgpuTypeImplementation string

const (
	// Pass through an entire physical GPU to a guest
	VgpuTypeImplementationPassthrough VgpuTypeImplementation = "passthrough"
	// vGPU using NVIDIA hardware
	VgpuTypeImplementationNvidia VgpuTypeImplementation = "nvidia"
	// vGPU using Intel GVT-g
	VgpuTypeImplementationGvtG VgpuTypeImplementation = "gvt_g"
	// vGPU using AMD MxGPU
	VgpuTypeImplementationMxgpu VgpuTypeImplementation = "mxgpu"
)

type VGPUTypeRecord struct {
	// Unique identifier/object reference
	UUID string
	// Name of VGPU vendor
	VendorName string
	// Model name associated with the VGPU type
	ModelName string
	// Framebuffer size of the VGPU type, in bytes
	FramebufferSize int
	// Maximum number of displays supported by the VGPU type
	MaxHeads int
	// Maximum resolution (width) supported by the VGPU type
	MaxResolutionX int
	// Maximum resolution (height) supported by the VGPU type
	MaxResolutionY int
	// List of PGPUs that support this VGPU type
	SupportedOnPGPUs []PGPURef
	// List of PGPUs that have this VGPU type enabled
	EnabledOnPGPUs []PGPURef
	// List of VGPUs of this type
	VGPUs []VGPURef
	// List of GPU groups in which at least one PGPU supports this VGPU type
	SupportedOnGPUGroups []GPUGroupRef
	// List of GPU groups in which at least one have this VGPU type enabled
	EnabledOnGPUGroups []GPUGroupRef
	// The internal implementation of this VGPU type
	Implementation VgpuTypeImplementation
	// Key used to identify VGPU types and avoid creating duplicates - this field is used internally and not intended for interpretation by API clients
	Identifier string
	// Indicates whether VGPUs of this type should be considered experimental
	Experimental bool
}

type VGPUTypeRef string

// A type of virtual GPU
type VGPUTypeClass struct {
	client *Client
}

// GetAllRecords Return a map of VGPU_type references to VGPU_type records for all VGPU_types known to the system.
func (_class VGPUTypeClass) GetAllRecords(sessionID SessionRef) (_retval map[VGPUTypeRef]VGPUTypeRecord, _err error) {
	_method := "VGPU_type.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefToVGPUTypeRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VGPU_types known to the system.
func (_class VGPUTypeClass) GetAll(sessionID SessionRef) (_retval []VGPUTypeRef, _err error) {
	_method := "VGPU_type.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetExperimental Get the experimental field of the given VGPU_type.
func (_class VGPUTypeClass) GetExperimental(sessionID SessionRef, self VGPUTypeRef) (_retval bool, _err error) {
	_method := "VGPU_type.get_experimental"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIdentifier Get the identifier field of the given VGPU_type.
func (_class VGPUTypeClass) GetIdentifier(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_identifier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetImplementation Get the implementation field of the given VGPU_type.
func (_class VGPUTypeClass) GetImplementation(sessionID SessionRef, self VGPUTypeRef) (_retval VgpuTypeImplementation, _err error) {
	_method := "VGPU_type.get_implementation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVgpuTypeImplementationToGo(_method + " -> ", _result.Value)
	return
}

// GetEnabledOnGPUGroups Get the enabled_on_GPU_groups field of the given VGPU_type.
func (_class VGPUTypeClass) GetEnabledOnGPUGroups(sessionID SessionRef, self VGPUTypeRef) (_retval []GPUGroupRef, _err error) {
	_method := "VGPU_type.get_enabled_on_GPU_groups"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetSupportedOnGPUGroups Get the supported_on_GPU_groups field of the given VGPU_type.
func (_class VGPUTypeClass) GetSupportedOnGPUGroups(sessionID SessionRef, self VGPUTypeRef) (_retval []GPUGroupRef, _err error) {
	_method := "VGPU_type.get_supported_on_GPU_groups"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertGPUGroupRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetVGPUs Get the VGPUs field of the given VGPU_type.
func (_class VGPUTypeClass) GetVGPUs(sessionID SessionRef, self VGPUTypeRef) (_retval []VGPURef, _err error) {
	_method := "VGPU_type.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetEnabledOnPGPUs Get the enabled_on_PGPUs field of the given VGPU_type.
func (_class VGPUTypeClass) GetEnabledOnPGPUs(sessionID SessionRef, self VGPUTypeRef) (_retval []PGPURef, _err error) {
	_method := "VGPU_type.get_enabled_on_PGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetSupportedOnPGPUs Get the supported_on_PGPUs field of the given VGPU_type.
func (_class VGPUTypeClass) GetSupportedOnPGPUs(sessionID SessionRef, self VGPUTypeRef) (_retval []PGPURef, _err error) {
	_method := "VGPU_type.get_supported_on_PGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetMaxResolutionY Get the max_resolution_y field of the given VGPU_type.
func (_class VGPUTypeClass) GetMaxResolutionY(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_max_resolution_y"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMaxResolutionX Get the max_resolution_x field of the given VGPU_type.
func (_class VGPUTypeClass) GetMaxResolutionX(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_max_resolution_x"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMaxHeads Get the max_heads field of the given VGPU_type.
func (_class VGPUTypeClass) GetMaxHeads(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_max_heads"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetFramebufferSize Get the framebuffer_size field of the given VGPU_type.
func (_class VGPUTypeClass) GetFramebufferSize(sessionID SessionRef, self VGPUTypeRef) (_retval int, _err error) {
	_method := "VGPU_type.get_framebuffer_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetModelName Get the model_name field of the given VGPU_type.
func (_class VGPUTypeClass) GetModelName(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_model_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVendorName Get the vendor_name field of the given VGPU_type.
func (_class VGPUTypeClass) GetVendorName(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_vendor_name"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given VGPU_type.
func (_class VGPUTypeClass) GetUUID(sessionID SessionRef, self VGPUTypeRef) (_retval string, _err error) {
	_method := "VGPU_type.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByUUID Get a reference to the VGPU_type instance with the specified UUID.
func (_class VGPUTypeClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VGPUTypeRef, _err error) {
	_method := "VGPU_type.get_by_uuid"
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
	_retval, _err = convertVGPUTypeRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VGPU_type.
func (_class VGPUTypeClass) GetRecord(sessionID SessionRef, self VGPUTypeRef) (_retval VGPUTypeRecord, _err error) {
	_method := "VGPU_type.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVGPUTypeRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPUTypeRecordToGo(_method + " -> ", _result.Value)
	return
}
