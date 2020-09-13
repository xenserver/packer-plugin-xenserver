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

type VdiOperations string

const (
	// Cloning the VDI
	VdiOperationsClone VdiOperations = "clone"
	// Copying the VDI
	VdiOperationsCopy VdiOperations = "copy"
	// Resizing the VDI
	VdiOperationsResize VdiOperations = "resize"
	// Resizing the VDI which may or may not be online
	VdiOperationsResizeOnline VdiOperations = "resize_online"
	// Snapshotting the VDI
	VdiOperationsSnapshot VdiOperations = "snapshot"
	// Mirroring the VDI
	VdiOperationsMirror VdiOperations = "mirror"
	// Destroying the VDI
	VdiOperationsDestroy VdiOperations = "destroy"
	// Forget about the VDI
	VdiOperationsForget VdiOperations = "forget"
	// Refreshing the fields of the VDI
	VdiOperationsUpdate VdiOperations = "update"
	// Forcibly unlocking the VDI
	VdiOperationsForceUnlock VdiOperations = "force_unlock"
	// Generating static configuration
	VdiOperationsGenerateConfig VdiOperations = "generate_config"
	// Enabling changed block tracking for a VDI
	VdiOperationsEnableCbt VdiOperations = "enable_cbt"
	// Disabling changed block tracking for a VDI
	VdiOperationsDisableCbt VdiOperations = "disable_cbt"
	// Deleting the data of the VDI
	VdiOperationsDataDestroy VdiOperations = "data_destroy"
	// Exporting a bitmap that shows the changed blocks between two VDIs
	VdiOperationsListChangedBlocks VdiOperations = "list_changed_blocks"
	// Setting the on_boot field of the VDI
	VdiOperationsSetOnBoot VdiOperations = "set_on_boot"
	// Operations on this VDI are temporarily blocked
	VdiOperationsBlocked VdiOperations = "blocked"
)

type VdiType string

const (
	// a disk that may be replaced on upgrade
	VdiTypeSystem VdiType = "system"
	// a disk that is always preserved on upgrade
	VdiTypeUser VdiType = "user"
	// a disk that may be reformatted on upgrade
	VdiTypeEphemeral VdiType = "ephemeral"
	// a disk that stores a suspend image
	VdiTypeSuspend VdiType = "suspend"
	// a disk that stores VM crashdump information
	VdiTypeCrashdump VdiType = "crashdump"
	// a disk used for HA storage heartbeating
	VdiTypeHaStatefile VdiType = "ha_statefile"
	// a disk used for HA Pool metadata
	VdiTypeMetadata VdiType = "metadata"
	// a disk used for a general metadata redo-log
	VdiTypeRedoLog VdiType = "redo_log"
	// a disk that stores SR-level RRDs
	VdiTypeRrd VdiType = "rrd"
	// a disk that stores PVS cache data
	VdiTypePvsCache VdiType = "pvs_cache"
	// Metadata about a snapshot VDI that has been deleted: the set of blocks that changed between some previous version of the disk and the version tracked by the snapshot.
	VdiTypeCbtMetadata VdiType = "cbt_metadata"
)

type OnBoot string

const (
	// When a VM containing this VDI is started, the contents of the VDI are reset to the state they were in when this flag was last set.
	OnBootReset OnBoot = "reset"
	// Standard behaviour.
	OnBootPersist OnBoot = "persist"
)

type VDIRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VdiOperations
	// links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VdiOperations
	// storage repository in which the VDI resides
	SR SRRef
	// list of vbds that refer to this disk
	VBDs []VBDRef
	// list of crash dumps that refer to this disk
	CrashDumps []CrashdumpRef
	// size of disk as presented to the guest (in bytes). Note that, depending on storage backend type, requested size may not be respected exactly
	VirtualSize int
	// amount of physical space that the disk image is currently taking up on the storage repository (in bytes)
	PhysicalUtilisation int
	// type of the VDI
	Type VdiType
	// true if this disk may be shared
	Sharable bool
	// true if this disk may ONLY be mounted read-only
	ReadOnly bool
	// additional configuration
	OtherConfig map[string]string
	// true if this disk is locked at the storage level
	StorageLock bool
	// location information
	Location string
	// 
	Managed bool
	// true if SR scan operation reported this VDI as not present on disk
	Missing bool
	// This field is always null. Deprecated
	Parent VDIRef
	// data to be inserted into the xenstore tree (/local/domain/0/backend/vbd/<domid>/<device-id>/sm-data) after the VDI is attached. This is generally set by the SM backends on vdi_attach.
	XenstoreData map[string]string
	// SM dependent data
	SmConfig map[string]string
	// true if this is a snapshot.
	IsASnapshot bool
	// Ref pointing to the VDI this snapshot is of.
	SnapshotOf VDIRef
	// List pointing to all the VDIs snapshots.
	Snapshots []VDIRef
	// Date/time when this snapshot was created.
	SnapshotTime time.Time
	// user-specified tags for categorization purposes
	Tags []string
	// true if this VDI is to be cached in the local cache SR
	AllowCaching bool
	// The behaviour of this VDI on a VM boot
	OnBoot OnBoot
	// The pool whose metadata is contained in this VDI
	MetadataOfPool PoolRef
	// Whether this VDI contains the latest known accessible metadata for the pool
	MetadataLatest bool
	// Whether this VDI is a Tools ISO
	IsToolsIso bool
	// True if changed blocks are tracked for this VDI
	CbtEnabled bool
}

type VDIRef string

// A virtual disk image
type VDIClass struct {
	client *Client
}

// GetAllRecords Return a map of VDI references to VDI records for all VDIs known to the system.
func (_class VDIClass) GetAllRecords(sessionID SessionRef) (_retval map[VDIRef]VDIRecord, _err error) {
	_method := "VDI.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToVDIRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VDIs known to the system.
func (_class VDIClass) GetAll(sessionID SessionRef) (_retval []VDIRef, _err error) {
	_method := "VDI.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetNbdInfo Get details specifying how to access this VDI via a Network Block Device server. For each of a set of NBD server addresses on which the VDI is available, the return value set contains a vdi_nbd_server_info object that contains an exportname to request once the NBD connection is established, and connection details for the address. An empty list is returned if there is no network that has a PIF on a host with access to the relevant SR, or if no such network has been assigned an NBD-related purpose in its purpose field. To access the given VDI, any of the vdi_nbd_server_info objects can be used to make a connection to a server, and then the VDI will be available by requesting the exportname.
//
// Errors:
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
func (_class VDIClass) GetNbdInfo(sessionID SessionRef, self VDIRef) (_retval []VdiNbdServerInfoRecord, _err error) {
	_method := "VDI.get_nbd_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVdiNbdServerInfoRecordSetToGo(_method + " -> ", _result.Value)
	return
}

// ListChangedBlocks Compare two VDIs in 64k block increments and report which blocks differ. This operation is not allowed when vdi_to is attached to a VM.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
func (_class VDIClass) ListChangedBlocks(sessionID SessionRef, vdiFrom VDIRef, vdiTo VDIRef) (_retval string, _err error) {
	_method := "VDI.list_changed_blocks"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiFromArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi_from"), vdiFrom)
	if _err != nil {
		return
	}
	_vdiToArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi_to"), vdiTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiFromArg, _vdiToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// DataDestroy Delete the data of the snapshot VDI, but keep its changed block tracking metadata. When successful, this call changes the type of the VDI to cbt_metadata. This operation is idempotent: calling it on a VDI of type cbt_metadata results in a no-op, and no error will be thrown.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
//  VDI_NO_CBT_METADATA - The requested operation is not allowed because the specified VDI does not have changed block tracking metadata.
//  VDI_IN_USE - This operation cannot be performed because this VDI is in use by some other operation
//  VDI_IS_A_PHYSICAL_DEVICE - The operation cannot be performed on physical device
func (_class VDIClass) DataDestroy(sessionID SessionRef, self VDIRef) (_err error) {
	_method := "VDI.data_destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// DisableCbt Disable changed block tracking for the VDI. This call is only allowed on VDIs that support enabling CBT. It is an idempotent operation - disabling CBT for a VDI for which CBT is not enabled results in a no-op, and no error will be thrown.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
//  VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the 'on-boot=reset' mode, or on VMs having such VDIs.
func (_class VDIClass) DisableCbt(sessionID SessionRef, self VDIRef) (_err error) {
	_method := "VDI.disable_cbt"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// EnableCbt Enable changed block tracking for the VDI. This call is idempotent - enabling CBT for a VDI for which CBT is already enabled results in a no-op, and no error will be thrown.
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
//  VDI_MISSING - This operation cannot be performed because the specified VDI could not be found on the storage substrate
//  SR_NOT_ATTACHED - The SR is not attached.
//  SR_HAS_NO_PBDS - The SR has no attached PBDs
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VDI_INCOMPATIBLE_TYPE - This operation cannot be performed because the specified VDI is of an incompatible type (eg: an HA statefile cannot be attached to a guest)
//  VDI_ON_BOOT_MODE_INCOMPATIBLE_WITH_OPERATION - This operation is not permitted on VDIs in the 'on-boot=reset' mode, or on VMs having such VDIs.
func (_class VDIClass) EnableCbt(sessionID SessionRef, self VDIRef) (_err error) {
	_method := "VDI.enable_cbt"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// PoolMigrate Migrate a VDI, which may be attached to a running guest, to a different SR. The destination SR must be visible to the guest.
func (_class VDIClass) PoolMigrate(sessionID SessionRef, vdi VDIRef, sr SRRef, options map[string]string) (_retval VDIRef, _err error) {
	_method := "VDI.pool_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _srArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// ReadDatabasePoolUUID Check the VDI cache for the pool UUID of the database on this VDI.
func (_class VDIClass) ReadDatabasePoolUUID(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	_method := "VDI.read_database_pool_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// OpenDatabase Load the metadata found on the supplied VDI and return a session reference which can be used in XenAPI calls to query its contents.
func (_class VDIClass) OpenDatabase(sessionID SessionRef, self VDIRef) (_retval SessionRef, _err error) {
	_method := "VDI.open_database"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetAllowCaching Set the value of the allow_caching parameter. This value can only be changed when the VDI is not attached to a running VM. The caching behaviour is only affected by this flag for VHD-based VDIs that have one parent and no child VHDs. Moreover, caching only takes place when the host running the VM containing this VDI has a nominated SR for local caching.
func (_class VDIClass) SetAllowCaching(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	_method := "VDI.set_allow_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetOnBoot Set the value of the on_boot parameter. This value can only be changed when the VDI is not attached to a running VM.
func (_class VDIClass) SetOnBoot(sessionID SessionRef, self VDIRef, value OnBoot) (_err error) {
	_method := "VDI.set_on_boot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnBootToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetNameDescription Set the name description of the VDI. This can only happen when its SR is currently attached.
func (_class VDIClass) SetNameDescription(sessionID SessionRef, self VDIRef, value string) (_err error) {
	_method := "VDI.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetNameLabel Set the name label of the VDI. This can only happen when then its SR is currently attached.
func (_class VDIClass) SetNameLabel(sessionID SessionRef, self VDIRef, value string) (_err error) {
	_method := "VDI.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetMetadataOfPool Records the pool whose metadata is contained by this VDI.
func (_class VDIClass) SetMetadataOfPool(sessionID SessionRef, self VDIRef, value PoolRef) (_err error) {
	_method := "VDI.set_metadata_of_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetSnapshotTime Sets the snapshot time of this VDI.
func (_class VDIClass) SetSnapshotTime(sessionID SessionRef, self VDIRef, value time.Time) (_err error) {
	_method := "VDI.set_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetSnapshotOf Sets the VDI of which this VDI is a snapshot
func (_class VDIClass) SetSnapshotOf(sessionID SessionRef, self VDIRef, value VDIRef) (_err error) {
	_method := "VDI.set_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetIsASnapshot Sets whether this VDI is a snapshot
func (_class VDIClass) SetIsASnapshot(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	_method := "VDI.set_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetPhysicalUtilisation Sets the VDI's physical_utilisation field
func (_class VDIClass) SetPhysicalUtilisation(sessionID SessionRef, self VDIRef, value int) (_err error) {
	_method := "VDI.set_physical_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetVirtualSize Sets the VDI's virtual_size field
func (_class VDIClass) SetVirtualSize(sessionID SessionRef, self VDIRef, value int) (_err error) {
	_method := "VDI.set_virtual_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetMissing Sets the VDI's missing field
func (_class VDIClass) SetMissing(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	_method := "VDI.set_missing"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetReadOnly Sets the VDI's read_only field
func (_class VDIClass) SetReadOnly(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	_method := "VDI.set_read_only"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetSharable Sets the VDI's sharable field
func (_class VDIClass) SetSharable(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	_method := "VDI.set_sharable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Forget Removes a VDI record from the database
func (_class VDIClass) Forget(sessionID SessionRef, vdi VDIRef) (_err error) {
	_method := "VDI.forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	return
}

// SetManaged Sets the VDI's managed field
func (_class VDIClass) SetManaged(sessionID SessionRef, self VDIRef, value bool) (_err error) {
	_method := "VDI.set_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// Copy Copy either a full VDI or the block differences between two VDIs into either a fresh VDI or an existing VDI.
//
// Errors:
//  VDI_READONLY - The operation required write access but this VDI is read-only
//  VDI_TOO_SMALL - The VDI is too small. Please resize it to at least the minimum size.
//  VDI_NOT_SPARSE - The VDI is not stored using a sparse format. It is not possible to query and manipulate only the changed blocks (or 'block differences' or 'disk deltas') between two VDIs. Please select a VDI which uses a sparse-aware technology such as VHD.
func (_class VDIClass) Copy(sessionID SessionRef, vdi VDIRef, sr SRRef, baseVdi VDIRef, intoVdi VDIRef) (_retval VDIRef, _err error) {
	_method := "VDI.copy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_baseVdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "base_vdi"), baseVdi)
	if _err != nil {
		return
	}
	_intoVdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "into_vdi"), intoVdi)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _srArg, _baseVdiArg, _intoVdiArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// Update Ask the storage backend to refresh the fields in the VDI object
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
func (_class VDIClass) Update(sessionID SessionRef, vdi VDIRef) (_err error) {
	_method := "VDI.update"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	return
}

// DbForget Removes a VDI record from the database
func (_class VDIClass) DbForget(sessionID SessionRef, vdi VDIRef) (_err error) {
	_method := "VDI.db_forget"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg)
	return
}

// DbIntroduce Create a new VDI record in the database only
func (_class VDIClass) DbIntroduce(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef, cbtEnabled bool) (_retval VDIRef, _err error) {
	_method := "VDI.db_introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "SR"), sr)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertEnumVdiTypeToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_sharableArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "sharable"), sharable)
	if _err != nil {
		return
	}
	_readOnlyArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "read_only"), readOnly)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_locationArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "location"), location)
	if _err != nil {
		return
	}
	_xenstoreDataArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "xenstore_data"), xenstoreData)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_virtualSizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "virtual_size"), virtualSize)
	if _err != nil {
		return
	}
	_physicalUtilisationArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "physical_utilisation"), physicalUtilisation)
	if _err != nil {
		return
	}
	_metadataOfPoolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "metadata_of_pool"), metadataOfPool)
	if _err != nil {
		return
	}
	_isASnapshotArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "is_a_snapshot"), isASnapshot)
	if _err != nil {
		return
	}
	_snapshotTimeArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_time"), snapshotTime)
	if _err != nil {
		return
	}
	_snapshotOfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_of"), snapshotOf)
	if _err != nil {
		return
	}
	_cbtEnabledArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "cbt_enabled"), cbtEnabled)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg, _nameLabelArg, _nameDescriptionArg, _srArg, _atypeArg, _sharableArg, _readOnlyArg, _otherConfigArg, _locationArg, _xenstoreDataArg, _smConfigArg, _managedArg, _virtualSizeArg, _physicalUtilisationArg, _metadataOfPoolArg, _isASnapshotArg, _snapshotTimeArg, _snapshotOfArg, _cbtEnabledArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// Introduce Create a new VDI record in the database only
//
// Errors:
//  SR_OPERATION_NOT_SUPPORTED - The SR backend does not support the operation (check the SR's allowed operations)
func (_class VDIClass) Introduce(sessionID SessionRef, uuid string, nameLabel string, nameDescription string, sr SRRef, atype VdiType, sharable bool, readOnly bool, otherConfig map[string]string, location string, xenstoreData map[string]string, smConfig map[string]string, managed bool, virtualSize int, physicalUtilisation int, metadataOfPool PoolRef, isASnapshot bool, snapshotTime time.Time, snapshotOf VDIRef) (_retval VDIRef, _err error) {
	_method := "VDI.introduce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_nameLabelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_label"), nameLabel)
	if _err != nil {
		return
	}
	_nameDescriptionArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name_description"), nameDescription)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "SR"), sr)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertEnumVdiTypeToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_sharableArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "sharable"), sharable)
	if _err != nil {
		return
	}
	_readOnlyArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "read_only"), readOnly)
	if _err != nil {
		return
	}
	_otherConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "other_config"), otherConfig)
	if _err != nil {
		return
	}
	_locationArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "location"), location)
	if _err != nil {
		return
	}
	_xenstoreDataArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "xenstore_data"), xenstoreData)
	if _err != nil {
		return
	}
	_smConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "sm_config"), smConfig)
	if _err != nil {
		return
	}
	_managedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "managed"), managed)
	if _err != nil {
		return
	}
	_virtualSizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "virtual_size"), virtualSize)
	if _err != nil {
		return
	}
	_physicalUtilisationArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "physical_utilisation"), physicalUtilisation)
	if _err != nil {
		return
	}
	_metadataOfPoolArg, _err := convertPoolRefToXen(fmt.Sprintf("%s(%s)", _method, "metadata_of_pool"), metadataOfPool)
	if _err != nil {
		return
	}
	_isASnapshotArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "is_a_snapshot"), isASnapshot)
	if _err != nil {
		return
	}
	_snapshotTimeArg, _err := convertTimeToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_time"), snapshotTime)
	if _err != nil {
		return
	}
	_snapshotOfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot_of"), snapshotOf)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg, _nameLabelArg, _nameDescriptionArg, _srArg, _atypeArg, _sharableArg, _readOnlyArg, _otherConfigArg, _locationArg, _xenstoreDataArg, _smConfigArg, _managedArg, _virtualSizeArg, _physicalUtilisationArg, _metadataOfPoolArg, _isASnapshotArg, _snapshotTimeArg, _snapshotOfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// ResizeOnline Resize the VDI which may or may not be attached to running guests.
func (_class VDIClass) ResizeOnline(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	_method := "VDI.resize_online"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_sizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "size"), size)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg, _sizeArg)
	return
}

// Resize Resize the VDI.
func (_class VDIClass) Resize(sessionID SessionRef, vdi VDIRef, size int) (_err error) {
	_method := "VDI.resize"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_sizeArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "size"), size)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vdiArg, _sizeArg)
	return
}

// Clone Take an exact copy of the VDI and return a reference to the new disk. If any driver_params are specified then these are passed through to the storage-specific substrate driver that implements the clone operation. NB the clone lives in the same Storage Repository as its parent.
func (_class VDIClass) Clone(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	_method := "VDI.clone"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_driverParamsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "driver_params"), driverParams)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _driverParamsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// Snapshot Take a read-only snapshot of the VDI, returning a reference to the snapshot. If any driver_params are specified then these are passed through to the storage-specific substrate driver that takes the snapshot. NB the snapshot lives in the same Storage Repository as its parent.
func (_class VDIClass) Snapshot(sessionID SessionRef, vdi VDIRef, driverParams map[string]string) (_retval VDIRef, _err error) {
	_method := "VDI.snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vdiArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "vdi"), vdi)
	if _err != nil {
		return
	}
	_driverParamsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "driver_params"), driverParams)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vdiArg, _driverParamsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// RemoveTags Remove the given value from the tags field of the given VDI.  If the value is not in that Set, then do nothing.
func (_class VDIClass) RemoveTags(sessionID SessionRef, self VDIRef, value string) (_err error) {
	_method := "VDI.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddTags Add the given value to the tags field of the given VDI.  If the value is already in that Set, then do nothing.
func (_class VDIClass) AddTags(sessionID SessionRef, self VDIRef, value string) (_err error) {
	_method := "VDI.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetTags Set the tags field of the given VDI.
func (_class VDIClass) SetTags(sessionID SessionRef, self VDIRef, value []string) (_err error) {
	_method := "VDI.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// RemoveFromSmConfig Remove the given key and its corresponding value from the sm_config field of the given VDI.  If the key is not in that Map, then do nothing.
func (_class VDIClass) RemoveFromSmConfig(sessionID SessionRef, self VDIRef, key string) (_err error) {
	_method := "VDI.remove_from_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToSmConfig Add the given key-value pair to the sm_config field of the given VDI.
func (_class VDIClass) AddToSmConfig(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	_method := "VDI.add_to_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetSmConfig Set the sm_config field of the given VDI.
func (_class VDIClass) SetSmConfig(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	_method := "VDI.set_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromXenstoreData Remove the given key and its corresponding value from the xenstore_data field of the given VDI.  If the key is not in that Map, then do nothing.
func (_class VDIClass) RemoveFromXenstoreData(sessionID SessionRef, self VDIRef, key string) (_err error) {
	_method := "VDI.remove_from_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToXenstoreData Add the given key-value pair to the xenstore_data field of the given VDI.
func (_class VDIClass) AddToXenstoreData(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	_method := "VDI.add_to_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetXenstoreData Set the xenstore_data field of the given VDI.
func (_class VDIClass) SetXenstoreData(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	_method := "VDI.set_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromOtherConfig Remove the given key and its corresponding value from the other_config field of the given VDI.  If the key is not in that Map, then do nothing.
func (_class VDIClass) RemoveFromOtherConfig(sessionID SessionRef, self VDIRef, key string) (_err error) {
	_method := "VDI.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToOtherConfig Add the given key-value pair to the other_config field of the given VDI.
func (_class VDIClass) AddToOtherConfig(sessionID SessionRef, self VDIRef, key string, value string) (_err error) {
	_method := "VDI.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetOtherConfig Set the other_config field of the given VDI.
func (_class VDIClass) SetOtherConfig(sessionID SessionRef, self VDIRef, value map[string]string) (_err error) {
	_method := "VDI.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetCbtEnabled Get the cbt_enabled field of the given VDI.
func (_class VDIClass) GetCbtEnabled(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_cbt_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIsToolsIso Get the is_tools_iso field of the given VDI.
func (_class VDIClass) GetIsToolsIso(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_is_tools_iso"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMetadataLatest Get the metadata_latest field of the given VDI.
func (_class VDIClass) GetMetadataLatest(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_metadata_latest"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMetadataOfPool Get the metadata_of_pool field of the given VDI.
func (_class VDIClass) GetMetadataOfPool(sessionID SessionRef, self VDIRef) (_retval PoolRef, _err error) {
	_method := "VDI.get_metadata_of_pool"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPoolRefToGo(_method + " -> ", _result.Value)
	return
}

// GetOnBoot Get the on_boot field of the given VDI.
func (_class VDIClass) GetOnBoot(sessionID SessionRef, self VDIRef) (_retval OnBoot, _err error) {
	_method := "VDI.get_on_boot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnBootToGo(_method + " -> ", _result.Value)
	return
}

// GetAllowCaching Get the allow_caching field of the given VDI.
func (_class VDIClass) GetAllowCaching(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_allow_caching"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetTags Get the tags field of the given VDI.
func (_class VDIClass) GetTags(sessionID SessionRef, self VDIRef) (_retval []string, _err error) {
	_method := "VDI.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSnapshotTime Get the snapshot_time field of the given VDI.
func (_class VDIClass) GetSnapshotTime(sessionID SessionRef, self VDIRef) (_retval time.Time, _err error) {
	_method := "VDI.get_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSnapshots Get the snapshots field of the given VDI.
func (_class VDIClass) GetSnapshots(sessionID SessionRef, self VDIRef) (_retval []VDIRef, _err error) {
	_method := "VDI.get_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetSnapshotOf Get the snapshot_of field of the given VDI.
func (_class VDIClass) GetSnapshotOf(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	_method := "VDI.get_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIsASnapshot Get the is_a_snapshot field of the given VDI.
func (_class VDIClass) GetIsASnapshot(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSmConfig Get the sm_config field of the given VDI.
func (_class VDIClass) GetSmConfig(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	_method := "VDI.get_sm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetXenstoreData Get the xenstore_data field of the given VDI.
func (_class VDIClass) GetXenstoreData(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	_method := "VDI.get_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetParent Get the parent field of the given VDI.
func (_class VDIClass) GetParent(sessionID SessionRef, self VDIRef) (_retval VDIRef, _err error) {
	_method := "VDI.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetMissing Get the missing field of the given VDI.
func (_class VDIClass) GetMissing(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_missing"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetManaged Get the managed field of the given VDI.
func (_class VDIClass) GetManaged(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_managed"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetLocation Get the location field of the given VDI.
func (_class VDIClass) GetLocation(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	_method := "VDI.get_location"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetStorageLock Get the storage_lock field of the given VDI.
func (_class VDIClass) GetStorageLock(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_storage_lock"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetOtherConfig Get the other_config field of the given VDI.
func (_class VDIClass) GetOtherConfig(sessionID SessionRef, self VDIRef) (_retval map[string]string, _err error) {
	_method := "VDI.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetReadOnly Get the read_only field of the given VDI.
func (_class VDIClass) GetReadOnly(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_read_only"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetSharable Get the sharable field of the given VDI.
func (_class VDIClass) GetSharable(sessionID SessionRef, self VDIRef) (_retval bool, _err error) {
	_method := "VDI.get_sharable"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetType Get the type field of the given VDI.
func (_class VDIClass) GetType(sessionID SessionRef, self VDIRef) (_retval VdiType, _err error) {
	_method := "VDI.get_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVdiTypeToGo(_method + " -> ", _result.Value)
	return
}

// GetPhysicalUtilisation Get the physical_utilisation field of the given VDI.
func (_class VDIClass) GetPhysicalUtilisation(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	_method := "VDI.get_physical_utilisation"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVirtualSize Get the virtual_size field of the given VDI.
func (_class VDIClass) GetVirtualSize(sessionID SessionRef, self VDIRef) (_retval int, _err error) {
	_method := "VDI.get_virtual_size"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetCrashDumps Get the crash_dumps field of the given VDI.
func (_class VDIClass) GetCrashDumps(sessionID SessionRef, self VDIRef) (_retval []CrashdumpRef, _err error) {
	_method := "VDI.get_crash_dumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetVBDs Get the VBDs field of the given VDI.
func (_class VDIClass) GetVBDs(sessionID SessionRef, self VDIRef) (_retval []VBDRef, _err error) {
	_method := "VDI.get_VBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}

// GetSR Get the SR field of the given VDI.
func (_class VDIClass) GetSR(sessionID SessionRef, self VDIRef) (_retval SRRef, _err error) {
	_method := "VDI.get_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

// GetCurrentOperations Get the current_operations field of the given VDI.
func (_class VDIClass) GetCurrentOperations(sessionID SessionRef, self VDIRef) (_retval map[string]VdiOperations, _err error) {
	_method := "VDI.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVdiOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAllowedOperations Get the allowed_operations field of the given VDI.
func (_class VDIClass) GetAllowedOperations(sessionID SessionRef, self VDIRef) (_retval []VdiOperations, _err error) {
	_method := "VDI.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVdiOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

// GetNameDescription Get the name/description field of the given VDI.
func (_class VDIClass) GetNameDescription(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	_method := "VDI.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given VDI.
func (_class VDIClass) GetNameLabel(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	_method := "VDI.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given VDI.
func (_class VDIClass) GetUUID(sessionID SessionRef, self VDIRef) (_retval string, _err error) {
	_method := "VDI.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the VDI instances with the given label.
func (_class VDIClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VDIRef, _err error) {
	_method := "VDI.get_by_name_label"
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
	_retval, _err = convertVDIRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the specified VDI instance.
func (_class VDIClass) Destroy(sessionID SessionRef, self VDIRef) (_err error) {
	_method := "VDI.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new VDI instance, and return its handle. The constructor args are: name_label, name_description, SR*, virtual_size*, type*, sharable*, read_only*, other_config*, xenstore_data, sm_config, tags (* = non-optional).
func (_class VDIClass) Create(sessionID SessionRef, args VDIRecord) (_retval VDIRef, _err error) {
	_method := "VDI.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVDIRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the VDI instance with the specified UUID.
func (_class VDIClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VDIRef, _err error) {
	_method := "VDI.get_by_uuid"
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
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VDI.
func (_class VDIClass) GetRecord(sessionID SessionRef, self VDIRef) (_retval VDIRecord, _err error) {
	_method := "VDI.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRecordToGo(_method + " -> ", _result.Value)
	return
}
