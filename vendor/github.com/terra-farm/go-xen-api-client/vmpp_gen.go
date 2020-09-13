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

type VmppBackupType string

const (
	// The backup is a snapshot
	VmppBackupTypeSnapshot VmppBackupType = "snapshot"
	// The backup is a checkpoint
	VmppBackupTypeCheckpoint VmppBackupType = "checkpoint"
)

type VmppBackupFrequency string

const (
	// Hourly backups
	VmppBackupFrequencyHourly VmppBackupFrequency = "hourly"
	// Daily backups
	VmppBackupFrequencyDaily VmppBackupFrequency = "daily"
	// Weekly backups
	VmppBackupFrequencyWeekly VmppBackupFrequency = "weekly"
)

type VmppArchiveFrequency string

const (
	// Never archive
	VmppArchiveFrequencyNever VmppArchiveFrequency = "never"
	// Archive after backup
	VmppArchiveFrequencyAlwaysAfterBackup VmppArchiveFrequency = "always_after_backup"
	// Daily archives
	VmppArchiveFrequencyDaily VmppArchiveFrequency = "daily"
	// Weekly backups
	VmppArchiveFrequencyWeekly VmppArchiveFrequency = "weekly"
)

type VmppArchiveTargetType string

const (
	// No target config
	VmppArchiveTargetTypeNone VmppArchiveTargetType = "none"
	// CIFS target config
	VmppArchiveTargetTypeCifs VmppArchiveTargetType = "cifs"
	// NFS target config
	VmppArchiveTargetTypeNfs VmppArchiveTargetType = "nfs"
)

type VMPPRecord struct {
	// Unique identifier/object reference
	UUID string
	// a human-readable name
	NameLabel string
	// a notes field containing human-readable description
	NameDescription string
	// enable or disable this policy
	IsPolicyEnabled bool
	// type of the backup sub-policy
	BackupType VmppBackupType
	// maximum number of backups that should be stored at any time
	BackupRetentionValue int
	// frequency of the backup schedule
	BackupFrequency VmppBackupFrequency
	// schedule of the backup containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	BackupSchedule map[string]string
	// true if this protection policy's backup is running
	IsBackupRunning bool
	// time of the last backup
	BackupLastRunTime time.Time
	// type of the archive target config
	ArchiveTargetType VmppArchiveTargetType
	// configuration for the archive, including its 'location', 'username', 'password'
	ArchiveTargetConfig map[string]string
	// frequency of the archive schedule
	ArchiveFrequency VmppArchiveFrequency
	// schedule of the archive containing 'hour', 'min', 'days'. Date/time-related information is in Local Timezone
	ArchiveSchedule map[string]string
	// true if this protection policy's archive is running
	IsArchiveRunning bool
	// time of the last archive
	ArchiveLastRunTime time.Time
	// all VMs attached to this protection policy
	VMs []VMRef
	// true if alarm is enabled for this policy
	IsAlarmEnabled bool
	// configuration for the alarm
	AlarmConfig map[string]string
	// recent alerts
	RecentAlerts []string
}

type VMPPRef string

// VM Protection Policy
type VMPPClass struct {
	client *Client
}

// GetAllRecords Return a map of VMPP references to VMPP records for all VMPPs known to the system.
func (_class VMPPClass) GetAllRecords(sessionID SessionRef) (_retval map[VMPPRef]VMPPRecord, _err error) {
	_method := "VMPP.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToVMPPRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// GetAll Return a list of all the VMPPs known to the system.
func (_class VMPPClass) GetAll(sessionID SessionRef) (_retval []VMPPRef, _err error) {
	_method := "VMPP.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefSetToGo(_method + " -> ", _result.Value)
	return
}

// SetArchiveLastRunTime 
func (_class VMPPClass) SetArchiveLastRunTime(sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	_method := "VMPP.set_archive_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetBackupLastRunTime 
func (_class VMPPClass) SetBackupLastRunTime(sessionID SessionRef, self VMPPRef, value time.Time) (_err error) {
	_method := "VMPP.set_backup_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromAlarmConfig 
func (_class VMPPClass) RemoveFromAlarmConfig(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	_method := "VMPP.remove_from_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromArchiveSchedule 
func (_class VMPPClass) RemoveFromArchiveSchedule(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	_method := "VMPP.remove_from_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromArchiveTargetConfig 
func (_class VMPPClass) RemoveFromArchiveTargetConfig(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	_method := "VMPP.remove_from_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// RemoveFromBackupSchedule 
func (_class VMPPClass) RemoveFromBackupSchedule(sessionID SessionRef, self VMPPRef, key string) (_err error) {
	_method := "VMPP.remove_from_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToAlarmConfig 
func (_class VMPPClass) AddToAlarmConfig(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	_method := "VMPP.add_to_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToArchiveSchedule 
func (_class VMPPClass) AddToArchiveSchedule(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	_method := "VMPP.add_to_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToArchiveTargetConfig 
func (_class VMPPClass) AddToArchiveTargetConfig(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	_method := "VMPP.add_to_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// AddToBackupSchedule 
func (_class VMPPClass) AddToBackupSchedule(sessionID SessionRef, self VMPPRef, key string, value string) (_err error) {
	_method := "VMPP.add_to_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetAlarmConfig 
func (_class VMPPClass) SetAlarmConfig(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	_method := "VMPP.set_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetIsAlarmEnabled Set the value of the is_alarm_enabled field
func (_class VMPPClass) SetIsAlarmEnabled(sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	_method := "VMPP.set_is_alarm_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetArchiveTargetConfig 
func (_class VMPPClass) SetArchiveTargetConfig(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	_method := "VMPP.set_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetArchiveTargetType Set the value of the archive_target_config_type field
func (_class VMPPClass) SetArchiveTargetType(sessionID SessionRef, self VMPPRef, value VmppArchiveTargetType) (_err error) {
	_method := "VMPP.set_archive_target_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppArchiveTargetTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetArchiveSchedule 
func (_class VMPPClass) SetArchiveSchedule(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	_method := "VMPP.set_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetArchiveFrequency Set the value of the archive_frequency field
func (_class VMPPClass) SetArchiveFrequency(sessionID SessionRef, self VMPPRef, value VmppArchiveFrequency) (_err error) {
	_method := "VMPP.set_archive_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppArchiveFrequencyToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetBackupSchedule 
func (_class VMPPClass) SetBackupSchedule(sessionID SessionRef, self VMPPRef, value map[string]string) (_err error) {
	_method := "VMPP.set_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetBackupFrequency Set the value of the backup_frequency field
func (_class VMPPClass) SetBackupFrequency(sessionID SessionRef, self VMPPRef, value VmppBackupFrequency) (_err error) {
	_method := "VMPP.set_backup_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppBackupFrequencyToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetBackupRetentionValue 
func (_class VMPPClass) SetBackupRetentionValue(sessionID SessionRef, self VMPPRef, value int) (_err error) {
	_method := "VMPP.set_backup_retention_value"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetAlerts This call fetches a history of alerts for a given protection policy
func (_class VMPPClass) GetAlerts(sessionID SessionRef, vmpp VMPPRef, hoursFromNow int) (_retval []string, _err error) {
	_method := "VMPP.get_alerts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmppArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "vmpp"), vmpp)
	if _err != nil {
		return
	}
	_hoursFromNowArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "hours_from_now"), hoursFromNow)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmppArg, _hoursFromNowArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// ArchiveNow This call archives the snapshot provided as a parameter
func (_class VMPPClass) ArchiveNow(sessionID SessionRef, snapshot VMRef) (_retval string, _err error) {
	_method := "VMPP.archive_now"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_snapshotArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot"), snapshot)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _snapshotArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// ProtectNow This call executes the protection policy immediately
func (_class VMPPClass) ProtectNow(sessionID SessionRef, vmpp VMPPRef) (_retval string, _err error) {
	_method := "VMPP.protect_now"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmppArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "vmpp"), vmpp)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmppArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// SetBackupType Set the backup_type field of the given VMPP.
func (_class VMPPClass) SetBackupType(sessionID SessionRef, self VMPPRef, value VmppBackupType) (_err error) {
	_method := "VMPP.set_backup_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVmppBackupTypeToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// SetIsPolicyEnabled Set the is_policy_enabled field of the given VMPP.
func (_class VMPPClass) SetIsPolicyEnabled(sessionID SessionRef, self VMPPRef, value bool) (_err error) {
	_method := "VMPP.set_is_policy_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetNameDescription Set the name/description field of the given VMPP.
func (_class VMPPClass) SetNameDescription(sessionID SessionRef, self VMPPRef, value string) (_err error) {
	_method := "VMPP.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// SetNameLabel Set the name/label field of the given VMPP.
func (_class VMPPClass) SetNameLabel(sessionID SessionRef, self VMPPRef, value string) (_err error) {
	_method := "VMPP.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetRecentAlerts Get the recent_alerts field of the given VMPP.
func (_class VMPPClass) GetRecentAlerts(sessionID SessionRef, self VMPPRef) (_retval []string, _err error) {
	_method := "VMPP.get_recent_alerts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetAlarmConfig Get the alarm_config field of the given VMPP.
func (_class VMPPClass) GetAlarmConfig(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	_method := "VMPP.get_alarm_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIsAlarmEnabled Get the is_alarm_enabled field of the given VMPP.
func (_class VMPPClass) GetIsAlarmEnabled(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	_method := "VMPP.get_is_alarm_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetVMs Get the VMs field of the given VMPP.
func (_class VMPPClass) GetVMs(sessionID SessionRef, self VMPPRef) (_retval []VMRef, _err error) {
	_method := "VMPP.get_VMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetArchiveLastRunTime Get the archive_last_run_time field of the given VMPP.
func (_class VMPPClass) GetArchiveLastRunTime(sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	_method := "VMPP.get_archive_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIsArchiveRunning Get the is_archive_running field of the given VMPP.
func (_class VMPPClass) GetIsArchiveRunning(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	_method := "VMPP.get_is_archive_running"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetArchiveSchedule Get the archive_schedule field of the given VMPP.
func (_class VMPPClass) GetArchiveSchedule(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	_method := "VMPP.get_archive_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetArchiveFrequency Get the archive_frequency field of the given VMPP.
func (_class VMPPClass) GetArchiveFrequency(sessionID SessionRef, self VMPPRef) (_retval VmppArchiveFrequency, _err error) {
	_method := "VMPP.get_archive_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppArchiveFrequencyToGo(_method + " -> ", _result.Value)
	return
}

// GetArchiveTargetConfig Get the archive_target_config field of the given VMPP.
func (_class VMPPClass) GetArchiveTargetConfig(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	_method := "VMPP.get_archive_target_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetArchiveTargetType Get the archive_target_type field of the given VMPP.
func (_class VMPPClass) GetArchiveTargetType(sessionID SessionRef, self VMPPRef) (_retval VmppArchiveTargetType, _err error) {
	_method := "VMPP.get_archive_target_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppArchiveTargetTypeToGo(_method + " -> ", _result.Value)
	return
}

// GetBackupLastRunTime Get the backup_last_run_time field of the given VMPP.
func (_class VMPPClass) GetBackupLastRunTime(sessionID SessionRef, self VMPPRef) (_retval time.Time, _err error) {
	_method := "VMPP.get_backup_last_run_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetIsBackupRunning Get the is_backup_running field of the given VMPP.
func (_class VMPPClass) GetIsBackupRunning(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	_method := "VMPP.get_is_backup_running"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetBackupSchedule Get the backup_schedule field of the given VMPP.
func (_class VMPPClass) GetBackupSchedule(sessionID SessionRef, self VMPPRef) (_retval map[string]string, _err error) {
	_method := "VMPP.get_backup_schedule"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetBackupFrequency Get the backup_frequency field of the given VMPP.
func (_class VMPPClass) GetBackupFrequency(sessionID SessionRef, self VMPPRef) (_retval VmppBackupFrequency, _err error) {
	_method := "VMPP.get_backup_frequency"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppBackupFrequencyToGo(_method + " -> ", _result.Value)
	return
}

// GetBackupRetentionValue Get the backup_retention_value field of the given VMPP.
func (_class VMPPClass) GetBackupRetentionValue(sessionID SessionRef, self VMPPRef) (_retval int, _err error) {
	_method := "VMPP.get_backup_retention_value"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetBackupType Get the backup_type field of the given VMPP.
func (_class VMPPClass) GetBackupType(sessionID SessionRef, self VMPPRef) (_retval VmppBackupType, _err error) {
	_method := "VMPP.get_backup_type"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVmppBackupTypeToGo(_method + " -> ", _result.Value)
	return
}

// GetIsPolicyEnabled Get the is_policy_enabled field of the given VMPP.
func (_class VMPPClass) GetIsPolicyEnabled(sessionID SessionRef, self VMPPRef) (_retval bool, _err error) {
	_method := "VMPP.get_is_policy_enabled"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameDescription Get the name/description field of the given VMPP.
func (_class VMPPClass) GetNameDescription(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	_method := "VMPP.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetNameLabel Get the name/label field of the given VMPP.
func (_class VMPPClass) GetNameLabel(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	_method := "VMPP.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetUUID Get the uuid field of the given VMPP.
func (_class VMPPClass) GetUUID(sessionID SessionRef, self VMPPRef) (_retval string, _err error) {
	_method := "VMPP.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
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

// GetByNameLabel Get all the VMPP instances with the given label.
func (_class VMPPClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMPPRef, _err error) {
	_method := "VMPP.get_by_name_label"
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
	_retval, _err = convertVMPPRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy Destroy the specified VMPP instance.
func (_class VMPPClass) Destroy(sessionID SessionRef, self VMPPRef) (_err error) {
	_method := "VMPP.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create Create a new VMPP instance, and return its handle. The constructor args are: name_label, name_description, is_policy_enabled, backup_type, backup_retention_value, backup_frequency, backup_schedule, archive_target_type, archive_target_config, archive_frequency, archive_schedule, is_alarm_enabled, alarm_config (* = non-optional).
func (_class VMPPClass) Create(sessionID SessionRef, args VMPPRecord) (_retval VMPPRef, _err error) {
	_method := "VMPP.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMPPRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}

// GetByUUID Get a reference to the VMPP instance with the specified UUID.
func (_class VMPPClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMPPRef, _err error) {
	_method := "VMPP.get_by_uuid"
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
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}

// GetRecord Get a record containing the current state of the given VMPP.
func (_class VMPPClass) GetRecord(sessionID SessionRef, self VMPPRef) (_retval VMPPRecord, _err error) {
	_method := "VMPP.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRecordToGo(_method + " -> ", _result.Value)
	return
}
