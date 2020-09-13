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

func convertBondRefToBondRecordMapToGo(context string, input interface{}) (goMap map[BondRef]BondRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[BondRef]BondRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertBondRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertBondRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertDRTaskRefToDRTaskRecordMapToGo(context string, input interface{}) (goMap map[DRTaskRef]DRTaskRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[DRTaskRef]DRTaskRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertDRTaskRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertDRTaskRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertFeatureRefToFeatureRecordMapToGo(context string, input interface{}) (goMap map[FeatureRef]FeatureRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[FeatureRef]FeatureRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertFeatureRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertFeatureRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertGPUGroupRefToGPUGroupRecordMapToGo(context string, input interface{}) (goMap map[GPUGroupRef]GPUGroupRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[GPUGroupRef]GPUGroupRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertGPUGroupRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertGPUGroupRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPBDRefToPBDRecordMapToGo(context string, input interface{}) (goMap map[PBDRef]PBDRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PBDRef]PBDRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPBDRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPBDRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPCIRefToPCIRecordMapToGo(context string, input interface{}) (goMap map[PCIRef]PCIRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PCIRef]PCIRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPCIRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPCIRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPGPURefToPGPURecordMapToGo(context string, input interface{}) (goMap map[PGPURef]PGPURecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PGPURef]PGPURecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPGPURefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPGPURecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPIFRefToPIFRecordMapToGo(context string, input interface{}) (goMap map[PIFRef]PIFRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PIFRef]PIFRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPIFRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPIFRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPIFMetricsRefToPIFMetricsRecordMapToGo(context string, input interface{}) (goMap map[PIFMetricsRef]PIFMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PIFMetricsRef]PIFMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPIFMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPIFMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPUSBRefToPUSBRecordMapToGo(context string, input interface{}) (goMap map[PUSBRef]PUSBRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PUSBRef]PUSBRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPUSBRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPUSBRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPVSCacheStorageRefToPVSCacheStorageRecordMapToGo(context string, input interface{}) (goMap map[PVSCacheStorageRef]PVSCacheStorageRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PVSCacheStorageRef]PVSCacheStorageRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPVSCacheStorageRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPVSCacheStorageRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPVSProxyRefToPVSProxyRecordMapToGo(context string, input interface{}) (goMap map[PVSProxyRef]PVSProxyRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PVSProxyRef]PVSProxyRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPVSProxyRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPVSProxyRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPVSServerRefToPVSServerRecordMapToGo(context string, input interface{}) (goMap map[PVSServerRef]PVSServerRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PVSServerRef]PVSServerRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPVSServerRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPVSServerRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPVSSiteRefToPVSSiteRecordMapToGo(context string, input interface{}) (goMap map[PVSSiteRef]PVSSiteRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PVSSiteRef]PVSSiteRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPVSSiteRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPVSSiteRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSDNControllerRefToSDNControllerRecordMapToGo(context string, input interface{}) (goMap map[SDNControllerRef]SDNControllerRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SDNControllerRef]SDNControllerRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSDNControllerRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSDNControllerRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSMRefToSMRecordMapToGo(context string, input interface{}) (goMap map[SMRef]SMRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SMRef]SMRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSMRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSRRefToSRRecordMapToGo(context string, input interface{}) (goMap map[SRRef]SRRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SRRef]SRRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSRRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSRRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertUSBGroupRefToUSBGroupRecordMapToGo(context string, input interface{}) (goMap map[USBGroupRef]USBGroupRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[USBGroupRef]USBGroupRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertUSBGroupRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertUSBGroupRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVBDRefToVBDRecordMapToGo(context string, input interface{}) (goMap map[VBDRef]VBDRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VBDRef]VBDRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVBDRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVBDRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVBDMetricsRefToVBDMetricsRecordMapToGo(context string, input interface{}) (goMap map[VBDMetricsRef]VBDMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VBDMetricsRef]VBDMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVBDMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVBDMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVDIRefToSRRefMapToXen(context string, goMap map[VDIRef]SRRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVDIRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertSRRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVDIRefToVDIRecordMapToGo(context string, input interface{}) (goMap map[VDIRef]VDIRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VDIRef]VDIRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVDIRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVDIRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVGPURefToGPUGroupRefMapToXen(context string, goMap map[VGPURef]GPUGroupRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVGPURefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertGPUGroupRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVGPURefToVGPURecordMapToGo(context string, input interface{}) (goMap map[VGPURef]VGPURecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VGPURef]VGPURecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVGPURefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVGPURecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVGPUTypeRefToVGPUTypeRecordMapToGo(context string, input interface{}) (goMap map[VGPUTypeRef]VGPUTypeRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VGPUTypeRef]VGPUTypeRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVGPUTypeRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVGPUTypeRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVGPUTypeRefToIntMapToGo(context string, input interface{}) (goMap map[VGPUTypeRef]int, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VGPUTypeRef]int, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVGPUTypeRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertIntToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVIFRefToVIFRecordMapToGo(context string, input interface{}) (goMap map[VIFRef]VIFRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VIFRef]VIFRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVIFRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVIFRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVIFRefToNetworkRefMapToXen(context string, goMap map[VIFRef]NetworkRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVIFRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertNetworkRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVIFRefToStringMapToGo(context string, input interface{}) (goMap map[VIFRef]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VIFRef]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVIFRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVIFRefToStringMapToXen(context string, goMap map[VIFRef]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVIFRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVIFMetricsRefToVIFMetricsRecordMapToGo(context string, input interface{}) (goMap map[VIFMetricsRef]VIFMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VIFMetricsRef]VIFMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVIFMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVIFMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVLANRefToVLANRecordMapToGo(context string, input interface{}) (goMap map[VLANRef]VLANRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VLANRef]VLANRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVLANRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVLANRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToStringToStringMapMapToGo(context string, input interface{}) (goMap map[VMRef]map[string]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMRef]map[string]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToStringMapToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToVMRecordMapToGo(context string, input interface{}) (goMap map[VMRef]VMRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMRef]VMRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToStringSetMapToGo(context string, input interface{}) (goMap map[VMRef][]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMRef][]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringSetToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToStringMapToXen(context string, goMap map[VMRef]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVMRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVMPPRefToVMPPRecordMapToGo(context string, input interface{}) (goMap map[VMPPRef]VMPPRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMPPRef]VMPPRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMPPRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMPPRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMSSRefToVMSSRecordMapToGo(context string, input interface{}) (goMap map[VMSSRef]VMSSRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMSSRef]VMSSRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMSSRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMSSRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMApplianceRefToVMApplianceRecordMapToGo(context string, input interface{}) (goMap map[VMApplianceRef]VMApplianceRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMApplianceRef]VMApplianceRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMApplianceRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMApplianceRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMGuestMetricsRefToVMGuestMetricsRecordMapToGo(context string, input interface{}) (goMap map[VMGuestMetricsRef]VMGuestMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMGuestMetricsRef]VMGuestMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMGuestMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMGuestMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMMetricsRefToVMMetricsRecordMapToGo(context string, input interface{}) (goMap map[VMMetricsRef]VMMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMMetricsRef]VMMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVUSBRefToVUSBRecordMapToGo(context string, input interface{}) (goMap map[VUSBRef]VUSBRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VUSBRef]VUSBRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVUSBRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVUSBRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertBlobRefToBlobRecordMapToGo(context string, input interface{}) (goMap map[BlobRef]BlobRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[BlobRef]BlobRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertBlobRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertBlobRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertConsoleRefToConsoleRecordMapToGo(context string, input interface{}) (goMap map[ConsoleRef]ConsoleRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[ConsoleRef]ConsoleRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertConsoleRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertConsoleRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertCrashdumpRefToCrashdumpRecordMapToGo(context string, input interface{}) (goMap map[CrashdumpRef]CrashdumpRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[CrashdumpRef]CrashdumpRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertCrashdumpRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertCrashdumpRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertEnumVMOperationsToStringMapToGo(context string, input interface{}) (goMap map[VMOperations]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMOperations]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertEnumVMOperationsToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertEnumVMOperationsToStringMapToXen(context string, goMap map[VMOperations]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertEnumVMOperationsToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertHostRefToHostRecordMapToGo(context string, input interface{}) (goMap map[HostRef]HostRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostRef]HostRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostRefToStringSetMapToGo(context string, input interface{}) (goMap map[HostRef][]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostRef][]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringSetToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostCPURefToHostCPURecordMapToGo(context string, input interface{}) (goMap map[HostCPURef]HostCPURecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostCPURef]HostCPURecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostCPURefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostCPURecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostCrashdumpRefToHostCrashdumpRecordMapToGo(context string, input interface{}) (goMap map[HostCrashdumpRef]HostCrashdumpRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostCrashdumpRef]HostCrashdumpRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostCrashdumpRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostCrashdumpRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostMetricsRefToHostMetricsRecordMapToGo(context string, input interface{}) (goMap map[HostMetricsRef]HostMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostMetricsRef]HostMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostPatchRefToHostPatchRecordMapToGo(context string, input interface{}) (goMap map[HostPatchRef]HostPatchRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostPatchRef]HostPatchRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostPatchRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostPatchRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertIntToFloatMapToGo(context string, input interface{}) (goMap map[int]float64, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[int]float64, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertIntToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertFloatToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertIntToIntMapToGo(context string, input interface{}) (goMap map[int]int, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[int]int, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertIntToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertIntToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertIntToStringSetMapToGo(context string, input interface{}) (goMap map[int][]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[int][]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertIntToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringSetToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertMessageRefToMessageRecordMapToGo(context string, input interface{}) (goMap map[MessageRef]MessageRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[MessageRef]MessageRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertMessageRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertMessageRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertNetworkRefToNetworkRecordMapToGo(context string, input interface{}) (goMap map[NetworkRef]NetworkRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[NetworkRef]NetworkRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertNetworkRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertNetworkRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPoolRefToPoolRecordMapToGo(context string, input interface{}) (goMap map[PoolRef]PoolRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PoolRef]PoolRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPoolRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPoolRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPoolPatchRefToPoolPatchRecordMapToGo(context string, input interface{}) (goMap map[PoolPatchRef]PoolPatchRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PoolPatchRef]PoolPatchRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPoolPatchRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPoolPatchRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPoolUpdateRefToPoolUpdateRecordMapToGo(context string, input interface{}) (goMap map[PoolUpdateRef]PoolUpdateRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PoolUpdateRef]PoolUpdateRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPoolUpdateRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPoolUpdateRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertRoleRefToRoleRecordMapToGo(context string, input interface{}) (goMap map[RoleRef]RoleRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[RoleRef]RoleRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertRoleRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertRoleRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSecretRefToSecretRecordMapToGo(context string, input interface{}) (goMap map[SecretRef]SecretRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SecretRef]SecretRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSecretRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSecretRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToBlobRefMapToGo(context string, input interface{}) (goMap map[string]BlobRef, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]BlobRef, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertBlobRefToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToBlobRefMapToXen(context string, goMap map[string]BlobRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertBlobRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumHostAllowedOperationsMapToGo(context string, input interface{}) (goMap map[string]HostAllowedOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]HostAllowedOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumHostAllowedOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumNetworkOperationsMapToGo(context string, input interface{}) (goMap map[string]NetworkOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]NetworkOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumNetworkOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumNetworkOperationsMapToXen(context string, goMap map[string]NetworkOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumNetworkOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumPoolAllowedOperationsMapToGo(context string, input interface{}) (goMap map[string]PoolAllowedOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]PoolAllowedOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumPoolAllowedOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumStorageOperationsMapToGo(context string, input interface{}) (goMap map[string]StorageOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]StorageOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumStorageOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumTaskAllowedOperationsMapToGo(context string, input interface{}) (goMap map[string]TaskAllowedOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]TaskAllowedOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumTaskAllowedOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVbdOperationsMapToGo(context string, input interface{}) (goMap map[string]VbdOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VbdOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVbdOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVbdOperationsMapToXen(context string, goMap map[string]VbdOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVbdOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVdiOperationsMapToGo(context string, input interface{}) (goMap map[string]VdiOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VdiOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVdiOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVdiOperationsMapToXen(context string, goMap map[string]VdiOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVdiOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVifOperationsMapToGo(context string, input interface{}) (goMap map[string]VifOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VifOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVifOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVifOperationsMapToXen(context string, goMap map[string]VifOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVifOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVMApplianceOperationMapToGo(context string, input interface{}) (goMap map[string]VMApplianceOperation, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VMApplianceOperation, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVMApplianceOperationToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVMApplianceOperationMapToXen(context string, goMap map[string]VMApplianceOperation) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVMApplianceOperationToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVMOperationsMapToGo(context string, input interface{}) (goMap map[string]VMOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VMOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVMOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVMOperationsMapToXen(context string, goMap map[string]VMOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVMOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVusbOperationsMapToGo(context string, input interface{}) (goMap map[string]VusbOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VusbOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVusbOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToIntMapToGo(context string, input interface{}) (goMap map[string]int, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]int, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertIntToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToStringMapToGo(context string, input interface{}) (goMap map[string]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToStringMapToXen(context string, goMap map[string]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertSubjectRefToSubjectRecordMapToGo(context string, input interface{}) (goMap map[SubjectRef]SubjectRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SubjectRef]SubjectRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSubjectRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSubjectRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertTaskRefToTaskRecordMapToGo(context string, input interface{}) (goMap map[TaskRef]TaskRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[TaskRef]TaskRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertTaskRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertTaskRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertTunnelRefToTunnelRecordMapToGo(context string, input interface{}) (goMap map[TunnelRef]TunnelRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[TunnelRef]TunnelRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertTunnelRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertTunnelRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertBondRecordToGo(context string, input interface{}) (record BondRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  masterValue, ok := rpcStruct["master"]
	if ok && masterValue != nil {
  	record.Master, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "master"), masterValue)
		if err != nil {
			return
		}
	}
  slavesValue, ok := rpcStruct["slaves"]
	if ok && slavesValue != nil {
  	record.Slaves, err = convertPIFRefSetToGo(fmt.Sprintf("%s.%s", context, "slaves"), slavesValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  primarySlaveValue, ok := rpcStruct["primary_slave"]
	if ok && primarySlaveValue != nil {
  	record.PrimarySlave, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "primary_slave"), primarySlaveValue)
		if err != nil {
			return
		}
	}
  modeValue, ok := rpcStruct["mode"]
	if ok && modeValue != nil {
  	record.Mode, err = convertEnumBondModeToGo(fmt.Sprintf("%s.%s", context, "mode"), modeValue)
		if err != nil {
			return
		}
	}
  propertiesValue, ok := rpcStruct["properties"]
	if ok && propertiesValue != nil {
  	record.Properties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "properties"), propertiesValue)
		if err != nil {
			return
		}
	}
  linksUpValue, ok := rpcStruct["links_up"]
	if ok && linksUpValue != nil {
  	record.LinksUp, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "links_up"), linksUpValue)
		if err != nil {
			return
		}
	}
	return
}

func convertBondRefSetToGo(context string, input interface{}) (slice []BondRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]BondRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertBondRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertBondRefToGo(context string, input interface{}) (ref BondRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = BondRef(value)
	}
	return
}

func convertBondRefToXen(context string, ref BondRef) (string, error) {
	return string(ref), nil
}

func convertDRTaskRecordToGo(context string, input interface{}) (record DRTaskRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  introducedSRsValue, ok := rpcStruct["introduced_SRs"]
	if ok && introducedSRsValue != nil {
  	record.IntroducedSRs, err = convertSRRefSetToGo(fmt.Sprintf("%s.%s", context, "introduced_SRs"), introducedSRsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertDRTaskRefSetToGo(context string, input interface{}) (slice []DRTaskRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]DRTaskRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertDRTaskRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertDRTaskRefToGo(context string, input interface{}) (ref DRTaskRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = DRTaskRef(value)
	}
	return
}

func convertDRTaskRefToXen(context string, ref DRTaskRef) (string, error) {
	return string(ref), nil
}

func convertFeatureRecordToGo(context string, input interface{}) (record FeatureRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  enabledValue, ok := rpcStruct["enabled"]
	if ok && enabledValue != nil {
  	record.Enabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enabled"), enabledValue)
		if err != nil {
			return
		}
	}
  experimentalValue, ok := rpcStruct["experimental"]
	if ok && experimentalValue != nil {
  	record.Experimental, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "experimental"), experimentalValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
	return
}

func convertFeatureRefSetToGo(context string, input interface{}) (slice []FeatureRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]FeatureRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertFeatureRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertFeatureRefToGo(context string, input interface{}) (ref FeatureRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = FeatureRef(value)
	}
	return
}

func convertFeatureRefToXen(context string, ref FeatureRef) (string, error) {
	return string(ref), nil
}

func convertGPUGroupRecordToGo(context string, input interface{}) (record GPUGroupRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  pgpusValue, ok := rpcStruct["PGPUs"]
	if ok && pgpusValue != nil {
  	record.PGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "PGPUs"), pgpusValue)
		if err != nil {
			return
		}
	}
  vgpusValue, ok := rpcStruct["VGPUs"]
	if ok && vgpusValue != nil {
  	record.VGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "VGPUs"), vgpusValue)
		if err != nil {
			return
		}
	}
  gpuTypesValue, ok := rpcStruct["GPU_types"]
	if ok && gpuTypesValue != nil {
  	record.GPUTypes, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "GPU_types"), gpuTypesValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  allocationAlgorithmValue, ok := rpcStruct["allocation_algorithm"]
	if ok && allocationAlgorithmValue != nil {
  	record.AllocationAlgorithm, err = convertEnumAllocationAlgorithmToGo(fmt.Sprintf("%s.%s", context, "allocation_algorithm"), allocationAlgorithmValue)
		if err != nil {
			return
		}
	}
  supportedVGPUTypesValue, ok := rpcStruct["supported_VGPU_types"]
	if ok && supportedVGPUTypesValue != nil {
  	record.SupportedVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "supported_VGPU_types"), supportedVGPUTypesValue)
		if err != nil {
			return
		}
	}
  enabledVGPUTypesValue, ok := rpcStruct["enabled_VGPU_types"]
	if ok && enabledVGPUTypesValue != nil {
  	record.EnabledVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_VGPU_types"), enabledVGPUTypesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertGPUGroupRefSetToGo(context string, input interface{}) (slice []GPUGroupRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]GPUGroupRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertGPUGroupRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertGPUGroupRefToGo(context string, input interface{}) (ref GPUGroupRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = GPUGroupRef(value)
	}
	return
}

func convertGPUGroupRefToXen(context string, ref GPUGroupRef) (string, error) {
	return string(ref), nil
}

func convertLVHDRecordToGo(context string, input interface{}) (record LVHDRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
	return
}

func convertLVHDRefToGo(context string, input interface{}) (ref LVHDRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = LVHDRef(value)
	}
	return
}

func convertLVHDRefToXen(context string, ref LVHDRef) (string, error) {
	return string(ref), nil
}

func convertPBDRecordToGo(context string, input interface{}) (record PBDRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  srValue, ok := rpcStruct["SR"]
	if ok && srValue != nil {
  	record.SR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "SR"), srValue)
		if err != nil {
			return
		}
	}
  deviceConfigValue, ok := rpcStruct["device_config"]
	if ok && deviceConfigValue != nil {
  	record.DeviceConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "device_config"), deviceConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPBDRecordToXen(context string, record PBDRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["host"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "host"), record.Host)
  if err != nil {
		return
	}
  rpcStruct["SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "SR"), record.SR)
  if err != nil {
		return
	}
  rpcStruct["device_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "device_config"), record.DeviceConfig)
  if err != nil {
		return
	}
  rpcStruct["currently_attached"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "currently_attached"), record.CurrentlyAttached)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertPBDRefSetToGo(context string, input interface{}) (slice []PBDRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PBDRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPBDRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPBDRefToGo(context string, input interface{}) (ref PBDRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PBDRef(value)
	}
	return
}

func convertPBDRefToXen(context string, ref PBDRef) (string, error) {
	return string(ref), nil
}

func convertPCIRecordToGo(context string, input interface{}) (record PCIRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  classNameValue, ok := rpcStruct["class_name"]
	if ok && classNameValue != nil {
  	record.ClassName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "class_name"), classNameValue)
		if err != nil {
			return
		}
	}
  vendorNameValue, ok := rpcStruct["vendor_name"]
	if ok && vendorNameValue != nil {
  	record.VendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_name"), vendorNameValue)
		if err != nil {
			return
		}
	}
  deviceNameValue, ok := rpcStruct["device_name"]
	if ok && deviceNameValue != nil {
  	record.DeviceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device_name"), deviceNameValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  pciIDValue, ok := rpcStruct["pci_id"]
	if ok && pciIDValue != nil {
  	record.PciID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "pci_id"), pciIDValue)
		if err != nil {
			return
		}
	}
  dependenciesValue, ok := rpcStruct["dependencies"]
	if ok && dependenciesValue != nil {
  	record.Dependencies, err = convertPCIRefSetToGo(fmt.Sprintf("%s.%s", context, "dependencies"), dependenciesValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  subsystemVendorNameValue, ok := rpcStruct["subsystem_vendor_name"]
	if ok && subsystemVendorNameValue != nil {
  	record.SubsystemVendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subsystem_vendor_name"), subsystemVendorNameValue)
		if err != nil {
			return
		}
	}
  subsystemDeviceNameValue, ok := rpcStruct["subsystem_device_name"]
	if ok && subsystemDeviceNameValue != nil {
  	record.SubsystemDeviceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subsystem_device_name"), subsystemDeviceNameValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPCIRefSetToGo(context string, input interface{}) (slice []PCIRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PCIRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPCIRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPCIRefSetToXen(context string, slice []PCIRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPCIRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertPCIRefToGo(context string, input interface{}) (ref PCIRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PCIRef(value)
	}
	return
}

func convertPCIRefToXen(context string, ref PCIRef) (string, error) {
	return string(ref), nil
}

func convertPGPURecordToGo(context string, input interface{}) (record PGPURecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  pciValue, ok := rpcStruct["PCI"]
	if ok && pciValue != nil {
  	record.PCI, err = convertPCIRefToGo(fmt.Sprintf("%s.%s", context, "PCI"), pciValue)
		if err != nil {
			return
		}
	}
  gpuGroupValue, ok := rpcStruct["GPU_group"]
	if ok && gpuGroupValue != nil {
  	record.GPUGroup, err = convertGPUGroupRefToGo(fmt.Sprintf("%s.%s", context, "GPU_group"), gpuGroupValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  supportedVGPUTypesValue, ok := rpcStruct["supported_VGPU_types"]
	if ok && supportedVGPUTypesValue != nil {
  	record.SupportedVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "supported_VGPU_types"), supportedVGPUTypesValue)
		if err != nil {
			return
		}
	}
  enabledVGPUTypesValue, ok := rpcStruct["enabled_VGPU_types"]
	if ok && enabledVGPUTypesValue != nil {
  	record.EnabledVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_VGPU_types"), enabledVGPUTypesValue)
		if err != nil {
			return
		}
	}
  residentVGPUsValue, ok := rpcStruct["resident_VGPUs"]
	if ok && residentVGPUsValue != nil {
  	record.ResidentVGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "resident_VGPUs"), residentVGPUsValue)
		if err != nil {
			return
		}
	}
  supportedVGPUMaxCapacitiesValue, ok := rpcStruct["supported_VGPU_max_capacities"]
	if ok && supportedVGPUMaxCapacitiesValue != nil {
  	record.SupportedVGPUMaxCapacities, err = convertVGPUTypeRefToIntMapToGo(fmt.Sprintf("%s.%s", context, "supported_VGPU_max_capacities"), supportedVGPUMaxCapacitiesValue)
		if err != nil {
			return
		}
	}
  dom0AccessValue, ok := rpcStruct["dom0_access"]
	if ok && dom0AccessValue != nil {
  	record.Dom0Access, err = convertEnumPgpuDom0AccessToGo(fmt.Sprintf("%s.%s", context, "dom0_access"), dom0AccessValue)
		if err != nil {
			return
		}
	}
  isSystemDisplayDeviceValue, ok := rpcStruct["is_system_display_device"]
	if ok && isSystemDisplayDeviceValue != nil {
  	record.IsSystemDisplayDevice, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_system_display_device"), isSystemDisplayDeviceValue)
		if err != nil {
			return
		}
	}
  compatibilityMetadataValue, ok := rpcStruct["compatibility_metadata"]
	if ok && compatibilityMetadataValue != nil {
  	record.CompatibilityMetadata, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "compatibility_metadata"), compatibilityMetadataValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPGPURefSetToGo(context string, input interface{}) (slice []PGPURef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PGPURef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPGPURefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPGPURefToGo(context string, input interface{}) (ref PGPURef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PGPURef(value)
	}
	return
}

func convertPGPURefToXen(context string, ref PGPURef) (string, error) {
	return string(ref), nil
}

func convertPIFRecordToGo(context string, input interface{}) (record PIFRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok && deviceValue != nil {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  networkValue, ok := rpcStruct["network"]
	if ok && networkValue != nil {
  	record.Network, err = convertNetworkRefToGo(fmt.Sprintf("%s.%s", context, "network"), networkValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  macValue, ok := rpcStruct["MAC"]
	if ok && macValue != nil {
  	record.MAC, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "MAC"), macValue)
		if err != nil {
			return
		}
	}
  mtuValue, ok := rpcStruct["MTU"]
	if ok && mtuValue != nil {
  	record.MTU, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "MTU"), mtuValue)
		if err != nil {
			return
		}
	}
  vlanValue, ok := rpcStruct["VLAN"]
	if ok && vlanValue != nil {
  	record.VLAN, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VLAN"), vlanValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok && metricsValue != nil {
  	record.Metrics, err = convertPIFMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  physicalValue, ok := rpcStruct["physical"]
	if ok && physicalValue != nil {
  	record.Physical, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "physical"), physicalValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  ipConfigurationModeValue, ok := rpcStruct["ip_configuration_mode"]
	if ok && ipConfigurationModeValue != nil {
  	record.IPConfigurationMode, err = convertEnumIPConfigurationModeToGo(fmt.Sprintf("%s.%s", context, "ip_configuration_mode"), ipConfigurationModeValue)
		if err != nil {
			return
		}
	}
  ipValue, ok := rpcStruct["IP"]
	if ok && ipValue != nil {
  	record.IP, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "IP"), ipValue)
		if err != nil {
			return
		}
	}
  netmaskValue, ok := rpcStruct["netmask"]
	if ok && netmaskValue != nil {
  	record.Netmask, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "netmask"), netmaskValue)
		if err != nil {
			return
		}
	}
  gatewayValue, ok := rpcStruct["gateway"]
	if ok && gatewayValue != nil {
  	record.Gateway, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "gateway"), gatewayValue)
		if err != nil {
			return
		}
	}
  dnsValue, ok := rpcStruct["DNS"]
	if ok && dnsValue != nil {
  	record.DNS, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "DNS"), dnsValue)
		if err != nil {
			return
		}
	}
  bondSlaveOfValue, ok := rpcStruct["bond_slave_of"]
	if ok && bondSlaveOfValue != nil {
  	record.BondSlaveOf, err = convertBondRefToGo(fmt.Sprintf("%s.%s", context, "bond_slave_of"), bondSlaveOfValue)
		if err != nil {
			return
		}
	}
  bondMasterOfValue, ok := rpcStruct["bond_master_of"]
	if ok && bondMasterOfValue != nil {
  	record.BondMasterOf, err = convertBondRefSetToGo(fmt.Sprintf("%s.%s", context, "bond_master_of"), bondMasterOfValue)
		if err != nil {
			return
		}
	}
  vlanMasterOfValue, ok := rpcStruct["VLAN_master_of"]
	if ok && vlanMasterOfValue != nil {
  	record.VLANMasterOf, err = convertVLANRefToGo(fmt.Sprintf("%s.%s", context, "VLAN_master_of"), vlanMasterOfValue)
		if err != nil {
			return
		}
	}
  vlanSlaveOfValue, ok := rpcStruct["VLAN_slave_of"]
	if ok && vlanSlaveOfValue != nil {
  	record.VLANSlaveOf, err = convertVLANRefSetToGo(fmt.Sprintf("%s.%s", context, "VLAN_slave_of"), vlanSlaveOfValue)
		if err != nil {
			return
		}
	}
  managementValue, ok := rpcStruct["management"]
	if ok && managementValue != nil {
  	record.Management, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "management"), managementValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  disallowUnplugValue, ok := rpcStruct["disallow_unplug"]
	if ok && disallowUnplugValue != nil {
  	record.DisallowUnplug, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "disallow_unplug"), disallowUnplugValue)
		if err != nil {
			return
		}
	}
  tunnelAccessPIFOfValue, ok := rpcStruct["tunnel_access_PIF_of"]
	if ok && tunnelAccessPIFOfValue != nil {
  	record.TunnelAccessPIFOf, err = convertTunnelRefSetToGo(fmt.Sprintf("%s.%s", context, "tunnel_access_PIF_of"), tunnelAccessPIFOfValue)
		if err != nil {
			return
		}
	}
  tunnelTransportPIFOfValue, ok := rpcStruct["tunnel_transport_PIF_of"]
	if ok && tunnelTransportPIFOfValue != nil {
  	record.TunnelTransportPIFOf, err = convertTunnelRefSetToGo(fmt.Sprintf("%s.%s", context, "tunnel_transport_PIF_of"), tunnelTransportPIFOfValue)
		if err != nil {
			return
		}
	}
  ipv6ConfigurationModeValue, ok := rpcStruct["ipv6_configuration_mode"]
	if ok && ipv6ConfigurationModeValue != nil {
  	record.Ipv6ConfigurationMode, err = convertEnumIpv6ConfigurationModeToGo(fmt.Sprintf("%s.%s", context, "ipv6_configuration_mode"), ipv6ConfigurationModeValue)
		if err != nil {
			return
		}
	}
  ipv6Value, ok := rpcStruct["IPv6"]
	if ok && ipv6Value != nil {
  	record.IPv6, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "IPv6"), ipv6Value)
		if err != nil {
			return
		}
	}
  ipv6GatewayValue, ok := rpcStruct["ipv6_gateway"]
	if ok && ipv6GatewayValue != nil {
  	record.Ipv6Gateway, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ipv6_gateway"), ipv6GatewayValue)
		if err != nil {
			return
		}
	}
  primaryAddressTypeValue, ok := rpcStruct["primary_address_type"]
	if ok && primaryAddressTypeValue != nil {
  	record.PrimaryAddressType, err = convertEnumPrimaryAddressTypeToGo(fmt.Sprintf("%s.%s", context, "primary_address_type"), primaryAddressTypeValue)
		if err != nil {
			return
		}
	}
  managedValue, ok := rpcStruct["managed"]
	if ok && managedValue != nil {
  	record.Managed, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "managed"), managedValue)
		if err != nil {
			return
		}
	}
  propertiesValue, ok := rpcStruct["properties"]
	if ok && propertiesValue != nil {
  	record.Properties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "properties"), propertiesValue)
		if err != nil {
			return
		}
	}
  capabilitiesValue, ok := rpcStruct["capabilities"]
	if ok && capabilitiesValue != nil {
  	record.Capabilities, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "capabilities"), capabilitiesValue)
		if err != nil {
			return
		}
	}
  igmpSnoopingStatusValue, ok := rpcStruct["igmp_snooping_status"]
	if ok && igmpSnoopingStatusValue != nil {
  	record.IgmpSnoopingStatus, err = convertEnumPifIgmpStatusToGo(fmt.Sprintf("%s.%s", context, "igmp_snooping_status"), igmpSnoopingStatusValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPIFRefSetToGo(context string, input interface{}) (slice []PIFRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PIFRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPIFRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPIFRefSetToXen(context string, slice []PIFRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPIFRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertPIFRefToGo(context string, input interface{}) (ref PIFRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PIFRef(value)
	}
	return
}

func convertPIFRefToXen(context string, ref PIFRef) (string, error) {
	return string(ref), nil
}

func convertPIFMetricsRecordToGo(context string, input interface{}) (record PIFMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  ioReadKbsValue, ok := rpcStruct["io_read_kbs"]
	if ok && ioReadKbsValue != nil {
  	record.IoReadKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_read_kbs"), ioReadKbsValue)
		if err != nil {
			return
		}
	}
  ioWriteKbsValue, ok := rpcStruct["io_write_kbs"]
	if ok && ioWriteKbsValue != nil {
  	record.IoWriteKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_write_kbs"), ioWriteKbsValue)
		if err != nil {
			return
		}
	}
  carrierValue, ok := rpcStruct["carrier"]
	if ok && carrierValue != nil {
  	record.Carrier, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "carrier"), carrierValue)
		if err != nil {
			return
		}
	}
  vendorIDValue, ok := rpcStruct["vendor_id"]
	if ok && vendorIDValue != nil {
  	record.VendorID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_id"), vendorIDValue)
		if err != nil {
			return
		}
	}
  vendorNameValue, ok := rpcStruct["vendor_name"]
	if ok && vendorNameValue != nil {
  	record.VendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_name"), vendorNameValue)
		if err != nil {
			return
		}
	}
  deviceIDValue, ok := rpcStruct["device_id"]
	if ok && deviceIDValue != nil {
  	record.DeviceID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device_id"), deviceIDValue)
		if err != nil {
			return
		}
	}
  deviceNameValue, ok := rpcStruct["device_name"]
	if ok && deviceNameValue != nil {
  	record.DeviceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device_name"), deviceNameValue)
		if err != nil {
			return
		}
	}
  speedValue, ok := rpcStruct["speed"]
	if ok && speedValue != nil {
  	record.Speed, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "speed"), speedValue)
		if err != nil {
			return
		}
	}
  duplexValue, ok := rpcStruct["duplex"]
	if ok && duplexValue != nil {
  	record.Duplex, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "duplex"), duplexValue)
		if err != nil {
			return
		}
	}
  pciBusPathValue, ok := rpcStruct["pci_bus_path"]
	if ok && pciBusPathValue != nil {
  	record.PciBusPath, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "pci_bus_path"), pciBusPathValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPIFMetricsRefSetToGo(context string, input interface{}) (slice []PIFMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PIFMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPIFMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPIFMetricsRefToGo(context string, input interface{}) (ref PIFMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PIFMetricsRef(value)
	}
	return
}

func convertPIFMetricsRefToXen(context string, ref PIFMetricsRef) (string, error) {
	return string(ref), nil
}

func convertPUSBRecordToGo(context string, input interface{}) (record PUSBRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  usbGroupValue, ok := rpcStruct["USB_group"]
	if ok && usbGroupValue != nil {
  	record.USBGroup, err = convertUSBGroupRefToGo(fmt.Sprintf("%s.%s", context, "USB_group"), usbGroupValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  pathValue, ok := rpcStruct["path"]
	if ok && pathValue != nil {
  	record.Path, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "path"), pathValue)
		if err != nil {
			return
		}
	}
  vendorIDValue, ok := rpcStruct["vendor_id"]
	if ok && vendorIDValue != nil {
  	record.VendorID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_id"), vendorIDValue)
		if err != nil {
			return
		}
	}
  vendorDescValue, ok := rpcStruct["vendor_desc"]
	if ok && vendorDescValue != nil {
  	record.VendorDesc, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_desc"), vendorDescValue)
		if err != nil {
			return
		}
	}
  productIDValue, ok := rpcStruct["product_id"]
	if ok && productIDValue != nil {
  	record.ProductID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "product_id"), productIDValue)
		if err != nil {
			return
		}
	}
  productDescValue, ok := rpcStruct["product_desc"]
	if ok && productDescValue != nil {
  	record.ProductDesc, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "product_desc"), productDescValue)
		if err != nil {
			return
		}
	}
  serialValue, ok := rpcStruct["serial"]
	if ok && serialValue != nil {
  	record.Serial, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "serial"), serialValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  descriptionValue, ok := rpcStruct["description"]
	if ok && descriptionValue != nil {
  	record.Description, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "description"), descriptionValue)
		if err != nil {
			return
		}
	}
  passthroughEnabledValue, ok := rpcStruct["passthrough_enabled"]
	if ok && passthroughEnabledValue != nil {
  	record.PassthroughEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "passthrough_enabled"), passthroughEnabledValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPUSBRefSetToGo(context string, input interface{}) (slice []PUSBRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PUSBRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPUSBRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPUSBRefToGo(context string, input interface{}) (ref PUSBRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PUSBRef(value)
	}
	return
}

func convertPUSBRefToXen(context string, ref PUSBRef) (string, error) {
	return string(ref), nil
}

func convertPVSCacheStorageRecordToGo(context string, input interface{}) (record PVSCacheStorageRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  srValue, ok := rpcStruct["SR"]
	if ok && srValue != nil {
  	record.SR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "SR"), srValue)
		if err != nil {
			return
		}
	}
  siteValue, ok := rpcStruct["site"]
	if ok && siteValue != nil {
  	record.Site, err = convertPVSSiteRefToGo(fmt.Sprintf("%s.%s", context, "site"), siteValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok && sizeValue != nil {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  vdiValue, ok := rpcStruct["VDI"]
	if ok && vdiValue != nil {
  	record.VDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "VDI"), vdiValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPVSCacheStorageRecordToXen(context string, record PVSCacheStorageRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["host"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "host"), record.Host)
  if err != nil {
		return
	}
  rpcStruct["SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "SR"), record.SR)
  if err != nil {
		return
	}
  rpcStruct["site"], err = convertPVSSiteRefToXen(fmt.Sprintf("%s.%s", context, "site"), record.Site)
  if err != nil {
		return
	}
  rpcStruct["size"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "size"), record.Size)
  if err != nil {
		return
	}
  rpcStruct["VDI"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "VDI"), record.VDI)
  if err != nil {
		return
	}
	return
}

func convertPVSCacheStorageRefSetToGo(context string, input interface{}) (slice []PVSCacheStorageRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PVSCacheStorageRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPVSCacheStorageRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPVSCacheStorageRefToGo(context string, input interface{}) (ref PVSCacheStorageRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PVSCacheStorageRef(value)
	}
	return
}

func convertPVSCacheStorageRefToXen(context string, ref PVSCacheStorageRef) (string, error) {
	return string(ref), nil
}

func convertPVSProxyRecordToGo(context string, input interface{}) (record PVSProxyRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  siteValue, ok := rpcStruct["site"]
	if ok && siteValue != nil {
  	record.Site, err = convertPVSSiteRefToGo(fmt.Sprintf("%s.%s", context, "site"), siteValue)
		if err != nil {
			return
		}
	}
  vifValue, ok := rpcStruct["VIF"]
	if ok && vifValue != nil {
  	record.VIF, err = convertVIFRefToGo(fmt.Sprintf("%s.%s", context, "VIF"), vifValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  statusValue, ok := rpcStruct["status"]
	if ok && statusValue != nil {
  	record.Status, err = convertEnumPvsProxyStatusToGo(fmt.Sprintf("%s.%s", context, "status"), statusValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPVSProxyRefSetToGo(context string, input interface{}) (slice []PVSProxyRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PVSProxyRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPVSProxyRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPVSProxyRefToGo(context string, input interface{}) (ref PVSProxyRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PVSProxyRef(value)
	}
	return
}

func convertPVSProxyRefToXen(context string, ref PVSProxyRef) (string, error) {
	return string(ref), nil
}

func convertPVSServerRecordToGo(context string, input interface{}) (record PVSServerRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  addressesValue, ok := rpcStruct["addresses"]
	if ok && addressesValue != nil {
  	record.Addresses, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "addresses"), addressesValue)
		if err != nil {
			return
		}
	}
  firstPortValue, ok := rpcStruct["first_port"]
	if ok && firstPortValue != nil {
  	record.FirstPort, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "first_port"), firstPortValue)
		if err != nil {
			return
		}
	}
  lastPortValue, ok := rpcStruct["last_port"]
	if ok && lastPortValue != nil {
  	record.LastPort, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "last_port"), lastPortValue)
		if err != nil {
			return
		}
	}
  siteValue, ok := rpcStruct["site"]
	if ok && siteValue != nil {
  	record.Site, err = convertPVSSiteRefToGo(fmt.Sprintf("%s.%s", context, "site"), siteValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPVSServerRefSetToGo(context string, input interface{}) (slice []PVSServerRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PVSServerRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPVSServerRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPVSServerRefToGo(context string, input interface{}) (ref PVSServerRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PVSServerRef(value)
	}
	return
}

func convertPVSServerRefToXen(context string, ref PVSServerRef) (string, error) {
	return string(ref), nil
}

func convertPVSSiteRecordToGo(context string, input interface{}) (record PVSSiteRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  pvsUUIDValue, ok := rpcStruct["PVS_uuid"]
	if ok && pvsUUIDValue != nil {
  	record.PVSUUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PVS_uuid"), pvsUUIDValue)
		if err != nil {
			return
		}
	}
  cacheStorageValue, ok := rpcStruct["cache_storage"]
	if ok && cacheStorageValue != nil {
  	record.CacheStorage, err = convertPVSCacheStorageRefSetToGo(fmt.Sprintf("%s.%s", context, "cache_storage"), cacheStorageValue)
		if err != nil {
			return
		}
	}
  serversValue, ok := rpcStruct["servers"]
	if ok && serversValue != nil {
  	record.Servers, err = convertPVSServerRefSetToGo(fmt.Sprintf("%s.%s", context, "servers"), serversValue)
		if err != nil {
			return
		}
	}
  proxiesValue, ok := rpcStruct["proxies"]
	if ok && proxiesValue != nil {
  	record.Proxies, err = convertPVSProxyRefSetToGo(fmt.Sprintf("%s.%s", context, "proxies"), proxiesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPVSSiteRefSetToGo(context string, input interface{}) (slice []PVSSiteRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PVSSiteRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPVSSiteRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPVSSiteRefToGo(context string, input interface{}) (ref PVSSiteRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PVSSiteRef(value)
	}
	return
}

func convertPVSSiteRefToXen(context string, ref PVSSiteRef) (string, error) {
	return string(ref), nil
}

func convertSDNControllerRecordToGo(context string, input interface{}) (record SDNControllerRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  protocolValue, ok := rpcStruct["protocol"]
	if ok && protocolValue != nil {
  	record.Protocol, err = convertEnumSdnControllerProtocolToGo(fmt.Sprintf("%s.%s", context, "protocol"), protocolValue)
		if err != nil {
			return
		}
	}
  addressValue, ok := rpcStruct["address"]
	if ok && addressValue != nil {
  	record.Address, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "address"), addressValue)
		if err != nil {
			return
		}
	}
  portValue, ok := rpcStruct["port"]
	if ok && portValue != nil {
  	record.Port, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "port"), portValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSDNControllerRefSetToGo(context string, input interface{}) (slice []SDNControllerRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SDNControllerRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSDNControllerRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSDNControllerRefToGo(context string, input interface{}) (ref SDNControllerRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SDNControllerRef(value)
	}
	return
}

func convertSDNControllerRefToXen(context string, ref SDNControllerRef) (string, error) {
	return string(ref), nil
}

func convertSMRecordToGo(context string, input interface{}) (record SMRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  vendorValue, ok := rpcStruct["vendor"]
	if ok && vendorValue != nil {
  	record.Vendor, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor"), vendorValue)
		if err != nil {
			return
		}
	}
  copyrightValue, ok := rpcStruct["copyright"]
	if ok && copyrightValue != nil {
  	record.Copyright, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "copyright"), copyrightValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  requiredAPIVersionValue, ok := rpcStruct["required_api_version"]
	if ok && requiredAPIVersionValue != nil {
  	record.RequiredAPIVersion, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "required_api_version"), requiredAPIVersionValue)
		if err != nil {
			return
		}
	}
  configurationValue, ok := rpcStruct["configuration"]
	if ok && configurationValue != nil {
  	record.Configuration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "configuration"), configurationValue)
		if err != nil {
			return
		}
	}
  capabilitiesValue, ok := rpcStruct["capabilities"]
	if ok && capabilitiesValue != nil {
  	record.Capabilities, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "capabilities"), capabilitiesValue)
		if err != nil {
			return
		}
	}
  featuresValue, ok := rpcStruct["features"]
	if ok && featuresValue != nil {
  	record.Features, err = convertStringToIntMapToGo(fmt.Sprintf("%s.%s", context, "features"), featuresValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  driverFilenameValue, ok := rpcStruct["driver_filename"]
	if ok && driverFilenameValue != nil {
  	record.DriverFilename, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "driver_filename"), driverFilenameValue)
		if err != nil {
			return
		}
	}
  requiredClusterStackValue, ok := rpcStruct["required_cluster_stack"]
	if ok && requiredClusterStackValue != nil {
  	record.RequiredClusterStack, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "required_cluster_stack"), requiredClusterStackValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSMRefSetToGo(context string, input interface{}) (slice []SMRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SMRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSMRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSMRefToGo(context string, input interface{}) (ref SMRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SMRef(value)
	}
	return
}

func convertSMRefToXen(context string, ref SMRef) (string, error) {
	return string(ref), nil
}

func convertSRRecordToGo(context string, input interface{}) (record SRRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumStorageOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumStorageOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vdisValue, ok := rpcStruct["VDIs"]
	if ok && vdisValue != nil {
  	record.VDIs, err = convertVDIRefSetToGo(fmt.Sprintf("%s.%s", context, "VDIs"), vdisValue)
		if err != nil {
			return
		}
	}
  pbdsValue, ok := rpcStruct["PBDs"]
	if ok && pbdsValue != nil {
  	record.PBDs, err = convertPBDRefSetToGo(fmt.Sprintf("%s.%s", context, "PBDs"), pbdsValue)
		if err != nil {
			return
		}
	}
  virtualAllocationValue, ok := rpcStruct["virtual_allocation"]
	if ok && virtualAllocationValue != nil {
  	record.VirtualAllocation, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "virtual_allocation"), virtualAllocationValue)
		if err != nil {
			return
		}
	}
  physicalUtilisationValue, ok := rpcStruct["physical_utilisation"]
	if ok && physicalUtilisationValue != nil {
  	record.PhysicalUtilisation, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "physical_utilisation"), physicalUtilisationValue)
		if err != nil {
			return
		}
	}
  physicalSizeValue, ok := rpcStruct["physical_size"]
	if ok && physicalSizeValue != nil {
  	record.PhysicalSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "physical_size"), physicalSizeValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  contentTypeValue, ok := rpcStruct["content_type"]
	if ok && contentTypeValue != nil {
  	record.ContentType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "content_type"), contentTypeValue)
		if err != nil {
			return
		}
	}
  sharedValue, ok := rpcStruct["shared"]
	if ok && sharedValue != nil {
  	record.Shared, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "shared"), sharedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok && tagsValue != nil {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  smConfigValue, ok := rpcStruct["sm_config"]
	if ok && smConfigValue != nil {
  	record.SmConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "sm_config"), smConfigValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok && blobsValue != nil {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  localCacheEnabledValue, ok := rpcStruct["local_cache_enabled"]
	if ok && localCacheEnabledValue != nil {
  	record.LocalCacheEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "local_cache_enabled"), localCacheEnabledValue)
		if err != nil {
			return
		}
	}
  introducedByValue, ok := rpcStruct["introduced_by"]
	if ok && introducedByValue != nil {
  	record.IntroducedBy, err = convertDRTaskRefToGo(fmt.Sprintf("%s.%s", context, "introduced_by"), introducedByValue)
		if err != nil {
			return
		}
	}
  clusteredValue, ok := rpcStruct["clustered"]
	if ok && clusteredValue != nil {
  	record.Clustered, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "clustered"), clusteredValue)
		if err != nil {
			return
		}
	}
  isToolsSrValue, ok := rpcStruct["is_tools_sr"]
	if ok && isToolsSrValue != nil {
  	record.IsToolsSr, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_tools_sr"), isToolsSrValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSRRefSetToGo(context string, input interface{}) (slice []SRRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SRRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSRRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSRRefSetToXen(context string, slice []SRRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSRRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertSRRefToGo(context string, input interface{}) (ref SRRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SRRef(value)
	}
	return
}

func convertSRRefToXen(context string, ref SRRef) (string, error) {
	return string(ref), nil
}

func convertUSBGroupRecordToGo(context string, input interface{}) (record USBGroupRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  pusbsValue, ok := rpcStruct["PUSBs"]
	if ok && pusbsValue != nil {
  	record.PUSBs, err = convertPUSBRefSetToGo(fmt.Sprintf("%s.%s", context, "PUSBs"), pusbsValue)
		if err != nil {
			return
		}
	}
  vusbsValue, ok := rpcStruct["VUSBs"]
	if ok && vusbsValue != nil {
  	record.VUSBs, err = convertVUSBRefSetToGo(fmt.Sprintf("%s.%s", context, "VUSBs"), vusbsValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertUSBGroupRefSetToGo(context string, input interface{}) (slice []USBGroupRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]USBGroupRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertUSBGroupRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertUSBGroupRefToGo(context string, input interface{}) (ref USBGroupRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = USBGroupRef(value)
	}
	return
}

func convertUSBGroupRefToXen(context string, ref USBGroupRef) (string, error) {
	return string(ref), nil
}

func convertVBDRecordToGo(context string, input interface{}) (record VBDRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumVbdOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumVbdOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  vdiValue, ok := rpcStruct["VDI"]
	if ok && vdiValue != nil {
  	record.VDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "VDI"), vdiValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok && deviceValue != nil {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  userdeviceValue, ok := rpcStruct["userdevice"]
	if ok && userdeviceValue != nil {
  	record.Userdevice, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "userdevice"), userdeviceValue)
		if err != nil {
			return
		}
	}
  bootableValue, ok := rpcStruct["bootable"]
	if ok && bootableValue != nil {
  	record.Bootable, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "bootable"), bootableValue)
		if err != nil {
			return
		}
	}
  modeValue, ok := rpcStruct["mode"]
	if ok && modeValue != nil {
  	record.Mode, err = convertEnumVbdModeToGo(fmt.Sprintf("%s.%s", context, "mode"), modeValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertEnumVbdTypeToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  unpluggableValue, ok := rpcStruct["unpluggable"]
	if ok && unpluggableValue != nil {
  	record.Unpluggable, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "unpluggable"), unpluggableValue)
		if err != nil {
			return
		}
	}
  storageLockValue, ok := rpcStruct["storage_lock"]
	if ok && storageLockValue != nil {
  	record.StorageLock, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "storage_lock"), storageLockValue)
		if err != nil {
			return
		}
	}
  emptyValue, ok := rpcStruct["empty"]
	if ok && emptyValue != nil {
  	record.Empty, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "empty"), emptyValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  statusCodeValue, ok := rpcStruct["status_code"]
	if ok && statusCodeValue != nil {
  	record.StatusCode, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "status_code"), statusCodeValue)
		if err != nil {
			return
		}
	}
  statusDetailValue, ok := rpcStruct["status_detail"]
	if ok && statusDetailValue != nil {
  	record.StatusDetail, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "status_detail"), statusDetailValue)
		if err != nil {
			return
		}
	}
  runtimePropertiesValue, ok := rpcStruct["runtime_properties"]
	if ok && runtimePropertiesValue != nil {
  	record.RuntimeProperties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "runtime_properties"), runtimePropertiesValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmTypeValue, ok := rpcStruct["qos_algorithm_type"]
	if ok && qosAlgorithmTypeValue != nil {
  	record.QosAlgorithmType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), qosAlgorithmTypeValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmParamsValue, ok := rpcStruct["qos_algorithm_params"]
	if ok && qosAlgorithmParamsValue != nil {
  	record.QosAlgorithmParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), qosAlgorithmParamsValue)
		if err != nil {
			return
		}
	}
  qosSupportedAlgorithmsValue, ok := rpcStruct["qos_supported_algorithms"]
	if ok && qosSupportedAlgorithmsValue != nil {
  	record.QosSupportedAlgorithms, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), qosSupportedAlgorithmsValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok && metricsValue != nil {
  	record.Metrics, err = convertVBDMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVBDRecordToXen(context string, record VBDRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVbdOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVbdOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["VDI"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "VDI"), record.VDI)
  if err != nil {
		return
	}
  rpcStruct["device"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "device"), record.Device)
  if err != nil {
		return
	}
  rpcStruct["userdevice"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "userdevice"), record.Userdevice)
  if err != nil {
		return
	}
  rpcStruct["bootable"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "bootable"), record.Bootable)
  if err != nil {
		return
	}
  rpcStruct["mode"], err = convertEnumVbdModeToXen(fmt.Sprintf("%s.%s", context, "mode"), record.Mode)
  if err != nil {
		return
	}
  rpcStruct["type"], err = convertEnumVbdTypeToXen(fmt.Sprintf("%s.%s", context, "type"), record.Type)
  if err != nil {
		return
	}
  rpcStruct["unpluggable"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "unpluggable"), record.Unpluggable)
  if err != nil {
		return
	}
  rpcStruct["storage_lock"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "storage_lock"), record.StorageLock)
  if err != nil {
		return
	}
  rpcStruct["empty"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "empty"), record.Empty)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["currently_attached"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "currently_attached"), record.CurrentlyAttached)
  if err != nil {
		return
	}
  rpcStruct["status_code"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "status_code"), record.StatusCode)
  if err != nil {
		return
	}
  rpcStruct["status_detail"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "status_detail"), record.StatusDetail)
  if err != nil {
		return
	}
  rpcStruct["runtime_properties"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "runtime_properties"), record.RuntimeProperties)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_type"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), record.QosAlgorithmType)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), record.QosAlgorithmParams)
  if err != nil {
		return
	}
  rpcStruct["qos_supported_algorithms"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), record.QosSupportedAlgorithms)
  if err != nil {
		return
	}
  rpcStruct["metrics"], err = convertVBDMetricsRefToXen(fmt.Sprintf("%s.%s", context, "metrics"), record.Metrics)
  if err != nil {
		return
	}
	return
}

func convertVBDRefSetToGo(context string, input interface{}) (slice []VBDRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VBDRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVBDRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVBDRefSetToXen(context string, slice []VBDRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVBDRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVBDRefToGo(context string, input interface{}) (ref VBDRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VBDRef(value)
	}
	return
}

func convertVBDRefToXen(context string, ref VBDRef) (string, error) {
	return string(ref), nil
}

func convertVBDMetricsRecordToGo(context string, input interface{}) (record VBDMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  ioReadKbsValue, ok := rpcStruct["io_read_kbs"]
	if ok && ioReadKbsValue != nil {
  	record.IoReadKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_read_kbs"), ioReadKbsValue)
		if err != nil {
			return
		}
	}
  ioWriteKbsValue, ok := rpcStruct["io_write_kbs"]
	if ok && ioWriteKbsValue != nil {
  	record.IoWriteKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_write_kbs"), ioWriteKbsValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVBDMetricsRefSetToGo(context string, input interface{}) (slice []VBDMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VBDMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVBDMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVBDMetricsRefToGo(context string, input interface{}) (ref VBDMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VBDMetricsRef(value)
	}
	return
}

func convertVBDMetricsRefToXen(context string, ref VBDMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVDIRecordToGo(context string, input interface{}) (record VDIRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumVdiOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumVdiOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  srValue, ok := rpcStruct["SR"]
	if ok && srValue != nil {
  	record.SR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "SR"), srValue)
		if err != nil {
			return
		}
	}
  vbdsValue, ok := rpcStruct["VBDs"]
	if ok && vbdsValue != nil {
  	record.VBDs, err = convertVBDRefSetToGo(fmt.Sprintf("%s.%s", context, "VBDs"), vbdsValue)
		if err != nil {
			return
		}
	}
  crashDumpsValue, ok := rpcStruct["crash_dumps"]
	if ok && crashDumpsValue != nil {
  	record.CrashDumps, err = convertCrashdumpRefSetToGo(fmt.Sprintf("%s.%s", context, "crash_dumps"), crashDumpsValue)
		if err != nil {
			return
		}
	}
  virtualSizeValue, ok := rpcStruct["virtual_size"]
	if ok && virtualSizeValue != nil {
  	record.VirtualSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "virtual_size"), virtualSizeValue)
		if err != nil {
			return
		}
	}
  physicalUtilisationValue, ok := rpcStruct["physical_utilisation"]
	if ok && physicalUtilisationValue != nil {
  	record.PhysicalUtilisation, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "physical_utilisation"), physicalUtilisationValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertEnumVdiTypeToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  sharableValue, ok := rpcStruct["sharable"]
	if ok && sharableValue != nil {
  	record.Sharable, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "sharable"), sharableValue)
		if err != nil {
			return
		}
	}
  readOnlyValue, ok := rpcStruct["read_only"]
	if ok && readOnlyValue != nil {
  	record.ReadOnly, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "read_only"), readOnlyValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  storageLockValue, ok := rpcStruct["storage_lock"]
	if ok && storageLockValue != nil {
  	record.StorageLock, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "storage_lock"), storageLockValue)
		if err != nil {
			return
		}
	}
  locationValue, ok := rpcStruct["location"]
	if ok && locationValue != nil {
  	record.Location, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "location"), locationValue)
		if err != nil {
			return
		}
	}
  managedValue, ok := rpcStruct["managed"]
	if ok && managedValue != nil {
  	record.Managed, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "managed"), managedValue)
		if err != nil {
			return
		}
	}
  missingValue, ok := rpcStruct["missing"]
	if ok && missingValue != nil {
  	record.Missing, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "missing"), missingValue)
		if err != nil {
			return
		}
	}
  parentValue, ok := rpcStruct["parent"]
	if ok && parentValue != nil {
  	record.Parent, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "parent"), parentValue)
		if err != nil {
			return
		}
	}
  xenstoreDataValue, ok := rpcStruct["xenstore_data"]
	if ok && xenstoreDataValue != nil {
  	record.XenstoreData, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "xenstore_data"), xenstoreDataValue)
		if err != nil {
			return
		}
	}
  smConfigValue, ok := rpcStruct["sm_config"]
	if ok && smConfigValue != nil {
  	record.SmConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "sm_config"), smConfigValue)
		if err != nil {
			return
		}
	}
  isASnapshotValue, ok := rpcStruct["is_a_snapshot"]
	if ok && isASnapshotValue != nil {
  	record.IsASnapshot, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), isASnapshotValue)
		if err != nil {
			return
		}
	}
  snapshotOfValue, ok := rpcStruct["snapshot_of"]
	if ok && snapshotOfValue != nil {
  	record.SnapshotOf, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "snapshot_of"), snapshotOfValue)
		if err != nil {
			return
		}
	}
  snapshotsValue, ok := rpcStruct["snapshots"]
	if ok && snapshotsValue != nil {
  	record.Snapshots, err = convertVDIRefSetToGo(fmt.Sprintf("%s.%s", context, "snapshots"), snapshotsValue)
		if err != nil {
			return
		}
	}
  snapshotTimeValue, ok := rpcStruct["snapshot_time"]
	if ok && snapshotTimeValue != nil {
  	record.SnapshotTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "snapshot_time"), snapshotTimeValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok && tagsValue != nil {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  allowCachingValue, ok := rpcStruct["allow_caching"]
	if ok && allowCachingValue != nil {
  	record.AllowCaching, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "allow_caching"), allowCachingValue)
		if err != nil {
			return
		}
	}
  onBootValue, ok := rpcStruct["on_boot"]
	if ok && onBootValue != nil {
  	record.OnBoot, err = convertEnumOnBootToGo(fmt.Sprintf("%s.%s", context, "on_boot"), onBootValue)
		if err != nil {
			return
		}
	}
  metadataOfPoolValue, ok := rpcStruct["metadata_of_pool"]
	if ok && metadataOfPoolValue != nil {
  	record.MetadataOfPool, err = convertPoolRefToGo(fmt.Sprintf("%s.%s", context, "metadata_of_pool"), metadataOfPoolValue)
		if err != nil {
			return
		}
	}
  metadataLatestValue, ok := rpcStruct["metadata_latest"]
	if ok && metadataLatestValue != nil {
  	record.MetadataLatest, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "metadata_latest"), metadataLatestValue)
		if err != nil {
			return
		}
	}
  isToolsIsoValue, ok := rpcStruct["is_tools_iso"]
	if ok && isToolsIsoValue != nil {
  	record.IsToolsIso, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_tools_iso"), isToolsIsoValue)
		if err != nil {
			return
		}
	}
  cbtEnabledValue, ok := rpcStruct["cbt_enabled"]
	if ok && cbtEnabledValue != nil {
  	record.CbtEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "cbt_enabled"), cbtEnabledValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVDIRecordToXen(context string, record VDIRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVdiOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVdiOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "SR"), record.SR)
  if err != nil {
		return
	}
  rpcStruct["VBDs"], err = convertVBDRefSetToXen(fmt.Sprintf("%s.%s", context, "VBDs"), record.VBDs)
  if err != nil {
		return
	}
  rpcStruct["crash_dumps"], err = convertCrashdumpRefSetToXen(fmt.Sprintf("%s.%s", context, "crash_dumps"), record.CrashDumps)
  if err != nil {
		return
	}
  rpcStruct["virtual_size"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "virtual_size"), record.VirtualSize)
  if err != nil {
		return
	}
  rpcStruct["physical_utilisation"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "physical_utilisation"), record.PhysicalUtilisation)
  if err != nil {
		return
	}
  rpcStruct["type"], err = convertEnumVdiTypeToXen(fmt.Sprintf("%s.%s", context, "type"), record.Type)
  if err != nil {
		return
	}
  rpcStruct["sharable"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "sharable"), record.Sharable)
  if err != nil {
		return
	}
  rpcStruct["read_only"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "read_only"), record.ReadOnly)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["storage_lock"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "storage_lock"), record.StorageLock)
  if err != nil {
		return
	}
  rpcStruct["location"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "location"), record.Location)
  if err != nil {
		return
	}
  rpcStruct["managed"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "managed"), record.Managed)
  if err != nil {
		return
	}
  rpcStruct["missing"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "missing"), record.Missing)
  if err != nil {
		return
	}
  rpcStruct["parent"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "parent"), record.Parent)
  if err != nil {
		return
	}
  rpcStruct["xenstore_data"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "xenstore_data"), record.XenstoreData)
  if err != nil {
		return
	}
  rpcStruct["sm_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "sm_config"), record.SmConfig)
  if err != nil {
		return
	}
  rpcStruct["is_a_snapshot"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), record.IsASnapshot)
  if err != nil {
		return
	}
  rpcStruct["snapshot_of"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "snapshot_of"), record.SnapshotOf)
  if err != nil {
		return
	}
  rpcStruct["snapshots"], err = convertVDIRefSetToXen(fmt.Sprintf("%s.%s", context, "snapshots"), record.Snapshots)
  if err != nil {
		return
	}
  rpcStruct["snapshot_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "snapshot_time"), record.SnapshotTime)
  if err != nil {
		return
	}
  rpcStruct["tags"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "tags"), record.Tags)
  if err != nil {
		return
	}
  rpcStruct["allow_caching"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "allow_caching"), record.AllowCaching)
  if err != nil {
		return
	}
  rpcStruct["on_boot"], err = convertEnumOnBootToXen(fmt.Sprintf("%s.%s", context, "on_boot"), record.OnBoot)
  if err != nil {
		return
	}
  rpcStruct["metadata_of_pool"], err = convertPoolRefToXen(fmt.Sprintf("%s.%s", context, "metadata_of_pool"), record.MetadataOfPool)
  if err != nil {
		return
	}
  rpcStruct["metadata_latest"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "metadata_latest"), record.MetadataLatest)
  if err != nil {
		return
	}
  rpcStruct["is_tools_iso"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_tools_iso"), record.IsToolsIso)
  if err != nil {
		return
	}
  rpcStruct["cbt_enabled"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "cbt_enabled"), record.CbtEnabled)
  if err != nil {
		return
	}
	return
}

func convertVDIRefSetToGo(context string, input interface{}) (slice []VDIRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VDIRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVDIRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVDIRefSetToXen(context string, slice []VDIRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVDIRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVDIRefToGo(context string, input interface{}) (ref VDIRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VDIRef(value)
	}
	return
}

func convertVDIRefToXen(context string, ref VDIRef) (string, error) {
	return string(ref), nil
}

func convertVGPURecordToGo(context string, input interface{}) (record VGPURecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  gpuGroupValue, ok := rpcStruct["GPU_group"]
	if ok && gpuGroupValue != nil {
  	record.GPUGroup, err = convertGPUGroupRefToGo(fmt.Sprintf("%s.%s", context, "GPU_group"), gpuGroupValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok && deviceValue != nil {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertVGPUTypeRefToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  residentOnValue, ok := rpcStruct["resident_on"]
	if ok && residentOnValue != nil {
  	record.ResidentOn, err = convertPGPURefToGo(fmt.Sprintf("%s.%s", context, "resident_on"), residentOnValue)
		if err != nil {
			return
		}
	}
  scheduledToBeResidentOnValue, ok := rpcStruct["scheduled_to_be_resident_on"]
	if ok && scheduledToBeResidentOnValue != nil {
  	record.ScheduledToBeResidentOn, err = convertPGPURefToGo(fmt.Sprintf("%s.%s", context, "scheduled_to_be_resident_on"), scheduledToBeResidentOnValue)
		if err != nil {
			return
		}
	}
  compatibilityMetadataValue, ok := rpcStruct["compatibility_metadata"]
	if ok && compatibilityMetadataValue != nil {
  	record.CompatibilityMetadata, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "compatibility_metadata"), compatibilityMetadataValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVGPURefSetToGo(context string, input interface{}) (slice []VGPURef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VGPURef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPURefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVGPURefSetToXen(context string, slice []VGPURef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPURefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVGPURefToGo(context string, input interface{}) (ref VGPURef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VGPURef(value)
	}
	return
}

func convertVGPURefToXen(context string, ref VGPURef) (string, error) {
	return string(ref), nil
}

func convertVGPUTypeRecordToGo(context string, input interface{}) (record VGPUTypeRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vendorNameValue, ok := rpcStruct["vendor_name"]
	if ok && vendorNameValue != nil {
  	record.VendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_name"), vendorNameValue)
		if err != nil {
			return
		}
	}
  modelNameValue, ok := rpcStruct["model_name"]
	if ok && modelNameValue != nil {
  	record.ModelName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "model_name"), modelNameValue)
		if err != nil {
			return
		}
	}
  framebufferSizeValue, ok := rpcStruct["framebuffer_size"]
	if ok && framebufferSizeValue != nil {
  	record.FramebufferSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "framebuffer_size"), framebufferSizeValue)
		if err != nil {
			return
		}
	}
  maxHeadsValue, ok := rpcStruct["max_heads"]
	if ok && maxHeadsValue != nil {
  	record.MaxHeads, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "max_heads"), maxHeadsValue)
		if err != nil {
			return
		}
	}
  maxResolutionXValue, ok := rpcStruct["max_resolution_x"]
	if ok && maxResolutionXValue != nil {
  	record.MaxResolutionX, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "max_resolution_x"), maxResolutionXValue)
		if err != nil {
			return
		}
	}
  maxResolutionYValue, ok := rpcStruct["max_resolution_y"]
	if ok && maxResolutionYValue != nil {
  	record.MaxResolutionY, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "max_resolution_y"), maxResolutionYValue)
		if err != nil {
			return
		}
	}
  supportedOnPGPUsValue, ok := rpcStruct["supported_on_PGPUs"]
	if ok && supportedOnPGPUsValue != nil {
  	record.SupportedOnPGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "supported_on_PGPUs"), supportedOnPGPUsValue)
		if err != nil {
			return
		}
	}
  enabledOnPGPUsValue, ok := rpcStruct["enabled_on_PGPUs"]
	if ok && enabledOnPGPUsValue != nil {
  	record.EnabledOnPGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_on_PGPUs"), enabledOnPGPUsValue)
		if err != nil {
			return
		}
	}
  vgpusValue, ok := rpcStruct["VGPUs"]
	if ok && vgpusValue != nil {
  	record.VGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "VGPUs"), vgpusValue)
		if err != nil {
			return
		}
	}
  supportedOnGPUGroupsValue, ok := rpcStruct["supported_on_GPU_groups"]
	if ok && supportedOnGPUGroupsValue != nil {
  	record.SupportedOnGPUGroups, err = convertGPUGroupRefSetToGo(fmt.Sprintf("%s.%s", context, "supported_on_GPU_groups"), supportedOnGPUGroupsValue)
		if err != nil {
			return
		}
	}
  enabledOnGPUGroupsValue, ok := rpcStruct["enabled_on_GPU_groups"]
	if ok && enabledOnGPUGroupsValue != nil {
  	record.EnabledOnGPUGroups, err = convertGPUGroupRefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_on_GPU_groups"), enabledOnGPUGroupsValue)
		if err != nil {
			return
		}
	}
  implementationValue, ok := rpcStruct["implementation"]
	if ok && implementationValue != nil {
  	record.Implementation, err = convertEnumVgpuTypeImplementationToGo(fmt.Sprintf("%s.%s", context, "implementation"), implementationValue)
		if err != nil {
			return
		}
	}
  identifierValue, ok := rpcStruct["identifier"]
	if ok && identifierValue != nil {
  	record.Identifier, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "identifier"), identifierValue)
		if err != nil {
			return
		}
	}
  experimentalValue, ok := rpcStruct["experimental"]
	if ok && experimentalValue != nil {
  	record.Experimental, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "experimental"), experimentalValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVGPUTypeRefSetToGo(context string, input interface{}) (slice []VGPUTypeRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VGPUTypeRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPUTypeRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVGPUTypeRefSetToXen(context string, slice []VGPUTypeRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPUTypeRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVGPUTypeRefToGo(context string, input interface{}) (ref VGPUTypeRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VGPUTypeRef(value)
	}
	return
}

func convertVGPUTypeRefToXen(context string, ref VGPUTypeRef) (string, error) {
	return string(ref), nil
}

func convertVIFRecordToGo(context string, input interface{}) (record VIFRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumVifOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumVifOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok && deviceValue != nil {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  networkValue, ok := rpcStruct["network"]
	if ok && networkValue != nil {
  	record.Network, err = convertNetworkRefToGo(fmt.Sprintf("%s.%s", context, "network"), networkValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  macValue, ok := rpcStruct["MAC"]
	if ok && macValue != nil {
  	record.MAC, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "MAC"), macValue)
		if err != nil {
			return
		}
	}
  mtuValue, ok := rpcStruct["MTU"]
	if ok && mtuValue != nil {
  	record.MTU, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "MTU"), mtuValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  statusCodeValue, ok := rpcStruct["status_code"]
	if ok && statusCodeValue != nil {
  	record.StatusCode, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "status_code"), statusCodeValue)
		if err != nil {
			return
		}
	}
  statusDetailValue, ok := rpcStruct["status_detail"]
	if ok && statusDetailValue != nil {
  	record.StatusDetail, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "status_detail"), statusDetailValue)
		if err != nil {
			return
		}
	}
  runtimePropertiesValue, ok := rpcStruct["runtime_properties"]
	if ok && runtimePropertiesValue != nil {
  	record.RuntimeProperties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "runtime_properties"), runtimePropertiesValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmTypeValue, ok := rpcStruct["qos_algorithm_type"]
	if ok && qosAlgorithmTypeValue != nil {
  	record.QosAlgorithmType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), qosAlgorithmTypeValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmParamsValue, ok := rpcStruct["qos_algorithm_params"]
	if ok && qosAlgorithmParamsValue != nil {
  	record.QosAlgorithmParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), qosAlgorithmParamsValue)
		if err != nil {
			return
		}
	}
  qosSupportedAlgorithmsValue, ok := rpcStruct["qos_supported_algorithms"]
	if ok && qosSupportedAlgorithmsValue != nil {
  	record.QosSupportedAlgorithms, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), qosSupportedAlgorithmsValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok && metricsValue != nil {
  	record.Metrics, err = convertVIFMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  macAutogeneratedValue, ok := rpcStruct["MAC_autogenerated"]
	if ok && macAutogeneratedValue != nil {
  	record.MACAutogenerated, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "MAC_autogenerated"), macAutogeneratedValue)
		if err != nil {
			return
		}
	}
  lockingModeValue, ok := rpcStruct["locking_mode"]
	if ok && lockingModeValue != nil {
  	record.LockingMode, err = convertEnumVifLockingModeToGo(fmt.Sprintf("%s.%s", context, "locking_mode"), lockingModeValue)
		if err != nil {
			return
		}
	}
  ipv4AllowedValue, ok := rpcStruct["ipv4_allowed"]
	if ok && ipv4AllowedValue != nil {
  	record.Ipv4Allowed, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ipv4_allowed"), ipv4AllowedValue)
		if err != nil {
			return
		}
	}
  ipv6AllowedValue, ok := rpcStruct["ipv6_allowed"]
	if ok && ipv6AllowedValue != nil {
  	record.Ipv6Allowed, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ipv6_allowed"), ipv6AllowedValue)
		if err != nil {
			return
		}
	}
  ipv4ConfigurationModeValue, ok := rpcStruct["ipv4_configuration_mode"]
	if ok && ipv4ConfigurationModeValue != nil {
  	record.Ipv4ConfigurationMode, err = convertEnumVifIpv4ConfigurationModeToGo(fmt.Sprintf("%s.%s", context, "ipv4_configuration_mode"), ipv4ConfigurationModeValue)
		if err != nil {
			return
		}
	}
  ipv4AddressesValue, ok := rpcStruct["ipv4_addresses"]
	if ok && ipv4AddressesValue != nil {
  	record.Ipv4Addresses, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ipv4_addresses"), ipv4AddressesValue)
		if err != nil {
			return
		}
	}
  ipv4GatewayValue, ok := rpcStruct["ipv4_gateway"]
	if ok && ipv4GatewayValue != nil {
  	record.Ipv4Gateway, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ipv4_gateway"), ipv4GatewayValue)
		if err != nil {
			return
		}
	}
  ipv6ConfigurationModeValue, ok := rpcStruct["ipv6_configuration_mode"]
	if ok && ipv6ConfigurationModeValue != nil {
  	record.Ipv6ConfigurationMode, err = convertEnumVifIpv6ConfigurationModeToGo(fmt.Sprintf("%s.%s", context, "ipv6_configuration_mode"), ipv6ConfigurationModeValue)
		if err != nil {
			return
		}
	}
  ipv6AddressesValue, ok := rpcStruct["ipv6_addresses"]
	if ok && ipv6AddressesValue != nil {
  	record.Ipv6Addresses, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ipv6_addresses"), ipv6AddressesValue)
		if err != nil {
			return
		}
	}
  ipv6GatewayValue, ok := rpcStruct["ipv6_gateway"]
	if ok && ipv6GatewayValue != nil {
  	record.Ipv6Gateway, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ipv6_gateway"), ipv6GatewayValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVIFRecordToXen(context string, record VIFRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVifOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVifOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["device"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "device"), record.Device)
  if err != nil {
		return
	}
  rpcStruct["network"], err = convertNetworkRefToXen(fmt.Sprintf("%s.%s", context, "network"), record.Network)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["MAC"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "MAC"), record.MAC)
  if err != nil {
		return
	}
  rpcStruct["MTU"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "MTU"), record.MTU)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["currently_attached"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "currently_attached"), record.CurrentlyAttached)
  if err != nil {
		return
	}
  rpcStruct["status_code"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "status_code"), record.StatusCode)
  if err != nil {
		return
	}
  rpcStruct["status_detail"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "status_detail"), record.StatusDetail)
  if err != nil {
		return
	}
  rpcStruct["runtime_properties"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "runtime_properties"), record.RuntimeProperties)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_type"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), record.QosAlgorithmType)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), record.QosAlgorithmParams)
  if err != nil {
		return
	}
  rpcStruct["qos_supported_algorithms"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), record.QosSupportedAlgorithms)
  if err != nil {
		return
	}
  rpcStruct["metrics"], err = convertVIFMetricsRefToXen(fmt.Sprintf("%s.%s", context, "metrics"), record.Metrics)
  if err != nil {
		return
	}
  rpcStruct["MAC_autogenerated"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "MAC_autogenerated"), record.MACAutogenerated)
  if err != nil {
		return
	}
  rpcStruct["locking_mode"], err = convertEnumVifLockingModeToXen(fmt.Sprintf("%s.%s", context, "locking_mode"), record.LockingMode)
  if err != nil {
		return
	}
  rpcStruct["ipv4_allowed"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "ipv4_allowed"), record.Ipv4Allowed)
  if err != nil {
		return
	}
  rpcStruct["ipv6_allowed"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "ipv6_allowed"), record.Ipv6Allowed)
  if err != nil {
		return
	}
  rpcStruct["ipv4_configuration_mode"], err = convertEnumVifIpv4ConfigurationModeToXen(fmt.Sprintf("%s.%s", context, "ipv4_configuration_mode"), record.Ipv4ConfigurationMode)
  if err != nil {
		return
	}
  rpcStruct["ipv4_addresses"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "ipv4_addresses"), record.Ipv4Addresses)
  if err != nil {
		return
	}
  rpcStruct["ipv4_gateway"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "ipv4_gateway"), record.Ipv4Gateway)
  if err != nil {
		return
	}
  rpcStruct["ipv6_configuration_mode"], err = convertEnumVifIpv6ConfigurationModeToXen(fmt.Sprintf("%s.%s", context, "ipv6_configuration_mode"), record.Ipv6ConfigurationMode)
  if err != nil {
		return
	}
  rpcStruct["ipv6_addresses"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "ipv6_addresses"), record.Ipv6Addresses)
  if err != nil {
		return
	}
  rpcStruct["ipv6_gateway"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "ipv6_gateway"), record.Ipv6Gateway)
  if err != nil {
		return
	}
	return
}

func convertVIFRefSetToGo(context string, input interface{}) (slice []VIFRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VIFRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVIFRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVIFRefSetToXen(context string, slice []VIFRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVIFRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVIFRefToGo(context string, input interface{}) (ref VIFRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VIFRef(value)
	}
	return
}

func convertVIFRefToXen(context string, ref VIFRef) (string, error) {
	return string(ref), nil
}

func convertVIFMetricsRecordToGo(context string, input interface{}) (record VIFMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  ioReadKbsValue, ok := rpcStruct["io_read_kbs"]
	if ok && ioReadKbsValue != nil {
  	record.IoReadKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_read_kbs"), ioReadKbsValue)
		if err != nil {
			return
		}
	}
  ioWriteKbsValue, ok := rpcStruct["io_write_kbs"]
	if ok && ioWriteKbsValue != nil {
  	record.IoWriteKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_write_kbs"), ioWriteKbsValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVIFMetricsRefSetToGo(context string, input interface{}) (slice []VIFMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VIFMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVIFMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVIFMetricsRefToGo(context string, input interface{}) (ref VIFMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VIFMetricsRef(value)
	}
	return
}

func convertVIFMetricsRefToXen(context string, ref VIFMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVLANRecordToGo(context string, input interface{}) (record VLANRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  taggedPIFValue, ok := rpcStruct["tagged_PIF"]
	if ok && taggedPIFValue != nil {
  	record.TaggedPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "tagged_PIF"), taggedPIFValue)
		if err != nil {
			return
		}
	}
  untaggedPIFValue, ok := rpcStruct["untagged_PIF"]
	if ok && untaggedPIFValue != nil {
  	record.UntaggedPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "untagged_PIF"), untaggedPIFValue)
		if err != nil {
			return
		}
	}
  tagValue, ok := rpcStruct["tag"]
	if ok && tagValue != nil {
  	record.Tag, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "tag"), tagValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVLANRefSetToGo(context string, input interface{}) (slice []VLANRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VLANRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVLANRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVLANRefToGo(context string, input interface{}) (ref VLANRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VLANRef(value)
	}
	return
}

func convertVLANRefToXen(context string, ref VLANRef) (string, error) {
	return string(ref), nil
}

func convertVMRecordToGo(context string, input interface{}) (record VMRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumVMOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumVMOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  powerStateValue, ok := rpcStruct["power_state"]
	if ok && powerStateValue != nil {
  	record.PowerState, err = convertEnumVMPowerStateToGo(fmt.Sprintf("%s.%s", context, "power_state"), powerStateValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  userVersionValue, ok := rpcStruct["user_version"]
	if ok && userVersionValue != nil {
  	record.UserVersion, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "user_version"), userVersionValue)
		if err != nil {
			return
		}
	}
  isATemplateValue, ok := rpcStruct["is_a_template"]
	if ok && isATemplateValue != nil {
  	record.IsATemplate, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_a_template"), isATemplateValue)
		if err != nil {
			return
		}
	}
  isDefaultTemplateValue, ok := rpcStruct["is_default_template"]
	if ok && isDefaultTemplateValue != nil {
  	record.IsDefaultTemplate, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_default_template"), isDefaultTemplateValue)
		if err != nil {
			return
		}
	}
  suspendVDIValue, ok := rpcStruct["suspend_VDI"]
	if ok && suspendVDIValue != nil {
  	record.SuspendVDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "suspend_VDI"), suspendVDIValue)
		if err != nil {
			return
		}
	}
  residentOnValue, ok := rpcStruct["resident_on"]
	if ok && residentOnValue != nil {
  	record.ResidentOn, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "resident_on"), residentOnValue)
		if err != nil {
			return
		}
	}
  affinityValue, ok := rpcStruct["affinity"]
	if ok && affinityValue != nil {
  	record.Affinity, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "affinity"), affinityValue)
		if err != nil {
			return
		}
	}
  memoryOverheadValue, ok := rpcStruct["memory_overhead"]
	if ok && memoryOverheadValue != nil {
  	record.MemoryOverhead, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_overhead"), memoryOverheadValue)
		if err != nil {
			return
		}
	}
  memoryTargetValue, ok := rpcStruct["memory_target"]
	if ok && memoryTargetValue != nil {
  	record.MemoryTarget, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_target"), memoryTargetValue)
		if err != nil {
			return
		}
	}
  memoryStaticMaxValue, ok := rpcStruct["memory_static_max"]
	if ok && memoryStaticMaxValue != nil {
  	record.MemoryStaticMax, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_static_max"), memoryStaticMaxValue)
		if err != nil {
			return
		}
	}
  memoryDynamicMaxValue, ok := rpcStruct["memory_dynamic_max"]
	if ok && memoryDynamicMaxValue != nil {
  	record.MemoryDynamicMax, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_dynamic_max"), memoryDynamicMaxValue)
		if err != nil {
			return
		}
	}
  memoryDynamicMinValue, ok := rpcStruct["memory_dynamic_min"]
	if ok && memoryDynamicMinValue != nil {
  	record.MemoryDynamicMin, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_dynamic_min"), memoryDynamicMinValue)
		if err != nil {
			return
		}
	}
  memoryStaticMinValue, ok := rpcStruct["memory_static_min"]
	if ok && memoryStaticMinValue != nil {
  	record.MemoryStaticMin, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_static_min"), memoryStaticMinValue)
		if err != nil {
			return
		}
	}
  vcpusParamsValue, ok := rpcStruct["VCPUs_params"]
	if ok && vcpusParamsValue != nil {
  	record.VCPUsParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_params"), vcpusParamsValue)
		if err != nil {
			return
		}
	}
  vcpusMaxValue, ok := rpcStruct["VCPUs_max"]
	if ok && vcpusMaxValue != nil {
  	record.VCPUsMax, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VCPUs_max"), vcpusMaxValue)
		if err != nil {
			return
		}
	}
  vcpusAtStartupValue, ok := rpcStruct["VCPUs_at_startup"]
	if ok && vcpusAtStartupValue != nil {
  	record.VCPUsAtStartup, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VCPUs_at_startup"), vcpusAtStartupValue)
		if err != nil {
			return
		}
	}
  actionsAfterShutdownValue, ok := rpcStruct["actions_after_shutdown"]
	if ok && actionsAfterShutdownValue != nil {
  	record.ActionsAfterShutdown, err = convertEnumOnNormalExitToGo(fmt.Sprintf("%s.%s", context, "actions_after_shutdown"), actionsAfterShutdownValue)
		if err != nil {
			return
		}
	}
  actionsAfterRebootValue, ok := rpcStruct["actions_after_reboot"]
	if ok && actionsAfterRebootValue != nil {
  	record.ActionsAfterReboot, err = convertEnumOnNormalExitToGo(fmt.Sprintf("%s.%s", context, "actions_after_reboot"), actionsAfterRebootValue)
		if err != nil {
			return
		}
	}
  actionsAfterCrashValue, ok := rpcStruct["actions_after_crash"]
	if ok && actionsAfterCrashValue != nil {
  	record.ActionsAfterCrash, err = convertEnumOnCrashBehaviourToGo(fmt.Sprintf("%s.%s", context, "actions_after_crash"), actionsAfterCrashValue)
		if err != nil {
			return
		}
	}
  consolesValue, ok := rpcStruct["consoles"]
	if ok && consolesValue != nil {
  	record.Consoles, err = convertConsoleRefSetToGo(fmt.Sprintf("%s.%s", context, "consoles"), consolesValue)
		if err != nil {
			return
		}
	}
  vifsValue, ok := rpcStruct["VIFs"]
	if ok && vifsValue != nil {
  	record.VIFs, err = convertVIFRefSetToGo(fmt.Sprintf("%s.%s", context, "VIFs"), vifsValue)
		if err != nil {
			return
		}
	}
  vbdsValue, ok := rpcStruct["VBDs"]
	if ok && vbdsValue != nil {
  	record.VBDs, err = convertVBDRefSetToGo(fmt.Sprintf("%s.%s", context, "VBDs"), vbdsValue)
		if err != nil {
			return
		}
	}
  vusbsValue, ok := rpcStruct["VUSBs"]
	if ok && vusbsValue != nil {
  	record.VUSBs, err = convertVUSBRefSetToGo(fmt.Sprintf("%s.%s", context, "VUSBs"), vusbsValue)
		if err != nil {
			return
		}
	}
  crashDumpsValue, ok := rpcStruct["crash_dumps"]
	if ok && crashDumpsValue != nil {
  	record.CrashDumps, err = convertCrashdumpRefSetToGo(fmt.Sprintf("%s.%s", context, "crash_dumps"), crashDumpsValue)
		if err != nil {
			return
		}
	}
  vtpmsValue, ok := rpcStruct["VTPMs"]
	if ok && vtpmsValue != nil {
  	record.VTPMs, err = convertVTPMRefSetToGo(fmt.Sprintf("%s.%s", context, "VTPMs"), vtpmsValue)
		if err != nil {
			return
		}
	}
  pvBootloaderValue, ok := rpcStruct["PV_bootloader"]
	if ok && pvBootloaderValue != nil {
  	record.PVBootloader, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_bootloader"), pvBootloaderValue)
		if err != nil {
			return
		}
	}
  pvKernelValue, ok := rpcStruct["PV_kernel"]
	if ok && pvKernelValue != nil {
  	record.PVKernel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_kernel"), pvKernelValue)
		if err != nil {
			return
		}
	}
  pvRamdiskValue, ok := rpcStruct["PV_ramdisk"]
	if ok && pvRamdiskValue != nil {
  	record.PVRamdisk, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_ramdisk"), pvRamdiskValue)
		if err != nil {
			return
		}
	}
  pvArgsValue, ok := rpcStruct["PV_args"]
	if ok && pvArgsValue != nil {
  	record.PVArgs, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_args"), pvArgsValue)
		if err != nil {
			return
		}
	}
  pvBootloaderArgsValue, ok := rpcStruct["PV_bootloader_args"]
	if ok && pvBootloaderArgsValue != nil {
  	record.PVBootloaderArgs, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_bootloader_args"), pvBootloaderArgsValue)
		if err != nil {
			return
		}
	}
  pvLegacyArgsValue, ok := rpcStruct["PV_legacy_args"]
	if ok && pvLegacyArgsValue != nil {
  	record.PVLegacyArgs, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_legacy_args"), pvLegacyArgsValue)
		if err != nil {
			return
		}
	}
  hvmBootPolicyValue, ok := rpcStruct["HVM_boot_policy"]
	if ok && hvmBootPolicyValue != nil {
  	record.HVMBootPolicy, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "HVM_boot_policy"), hvmBootPolicyValue)
		if err != nil {
			return
		}
	}
  hvmBootParamsValue, ok := rpcStruct["HVM_boot_params"]
	if ok && hvmBootParamsValue != nil {
  	record.HVMBootParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "HVM_boot_params"), hvmBootParamsValue)
		if err != nil {
			return
		}
	}
  hvmShadowMultiplierValue, ok := rpcStruct["HVM_shadow_multiplier"]
	if ok && hvmShadowMultiplierValue != nil {
  	record.HVMShadowMultiplier, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "HVM_shadow_multiplier"), hvmShadowMultiplierValue)
		if err != nil {
			return
		}
	}
  platformValue, ok := rpcStruct["platform"]
	if ok && platformValue != nil {
  	record.Platform, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "platform"), platformValue)
		if err != nil {
			return
		}
	}
  pciBusValue, ok := rpcStruct["PCI_bus"]
	if ok && pciBusValue != nil {
  	record.PCIBus, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PCI_bus"), pciBusValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  domidValue, ok := rpcStruct["domid"]
	if ok && domidValue != nil {
  	record.Domid, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "domid"), domidValue)
		if err != nil {
			return
		}
	}
  domarchValue, ok := rpcStruct["domarch"]
	if ok && domarchValue != nil {
  	record.Domarch, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "domarch"), domarchValue)
		if err != nil {
			return
		}
	}
  lastBootCPUFlagsValue, ok := rpcStruct["last_boot_CPU_flags"]
	if ok && lastBootCPUFlagsValue != nil {
  	record.LastBootCPUFlags, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "last_boot_CPU_flags"), lastBootCPUFlagsValue)
		if err != nil {
			return
		}
	}
  isControlDomainValue, ok := rpcStruct["is_control_domain"]
	if ok && isControlDomainValue != nil {
  	record.IsControlDomain, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_control_domain"), isControlDomainValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok && metricsValue != nil {
  	record.Metrics, err = convertVMMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  guestMetricsValue, ok := rpcStruct["guest_metrics"]
	if ok && guestMetricsValue != nil {
  	record.GuestMetrics, err = convertVMGuestMetricsRefToGo(fmt.Sprintf("%s.%s", context, "guest_metrics"), guestMetricsValue)
		if err != nil {
			return
		}
	}
  lastBootedRecordValue, ok := rpcStruct["last_booted_record"]
	if ok && lastBootedRecordValue != nil {
  	record.LastBootedRecord, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "last_booted_record"), lastBootedRecordValue)
		if err != nil {
			return
		}
	}
  recommendationsValue, ok := rpcStruct["recommendations"]
	if ok && recommendationsValue != nil {
  	record.Recommendations, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "recommendations"), recommendationsValue)
		if err != nil {
			return
		}
	}
  xenstoreDataValue, ok := rpcStruct["xenstore_data"]
	if ok && xenstoreDataValue != nil {
  	record.XenstoreData, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "xenstore_data"), xenstoreDataValue)
		if err != nil {
			return
		}
	}
  haAlwaysRunValue, ok := rpcStruct["ha_always_run"]
	if ok && haAlwaysRunValue != nil {
  	record.HaAlwaysRun, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_always_run"), haAlwaysRunValue)
		if err != nil {
			return
		}
	}
  haRestartPriorityValue, ok := rpcStruct["ha_restart_priority"]
	if ok && haRestartPriorityValue != nil {
  	record.HaRestartPriority, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ha_restart_priority"), haRestartPriorityValue)
		if err != nil {
			return
		}
	}
  isASnapshotValue, ok := rpcStruct["is_a_snapshot"]
	if ok && isASnapshotValue != nil {
  	record.IsASnapshot, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), isASnapshotValue)
		if err != nil {
			return
		}
	}
  snapshotOfValue, ok := rpcStruct["snapshot_of"]
	if ok && snapshotOfValue != nil {
  	record.SnapshotOf, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "snapshot_of"), snapshotOfValue)
		if err != nil {
			return
		}
	}
  snapshotsValue, ok := rpcStruct["snapshots"]
	if ok && snapshotsValue != nil {
  	record.Snapshots, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "snapshots"), snapshotsValue)
		if err != nil {
			return
		}
	}
  snapshotTimeValue, ok := rpcStruct["snapshot_time"]
	if ok && snapshotTimeValue != nil {
  	record.SnapshotTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "snapshot_time"), snapshotTimeValue)
		if err != nil {
			return
		}
	}
  transportableSnapshotIDValue, ok := rpcStruct["transportable_snapshot_id"]
	if ok && transportableSnapshotIDValue != nil {
  	record.TransportableSnapshotID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "transportable_snapshot_id"), transportableSnapshotIDValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok && blobsValue != nil {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok && tagsValue != nil {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  blockedOperationsValue, ok := rpcStruct["blocked_operations"]
	if ok && blockedOperationsValue != nil {
  	record.BlockedOperations, err = convertEnumVMOperationsToStringMapToGo(fmt.Sprintf("%s.%s", context, "blocked_operations"), blockedOperationsValue)
		if err != nil {
			return
		}
	}
  snapshotInfoValue, ok := rpcStruct["snapshot_info"]
	if ok && snapshotInfoValue != nil {
  	record.SnapshotInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "snapshot_info"), snapshotInfoValue)
		if err != nil {
			return
		}
	}
  snapshotMetadataValue, ok := rpcStruct["snapshot_metadata"]
	if ok && snapshotMetadataValue != nil {
  	record.SnapshotMetadata, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "snapshot_metadata"), snapshotMetadataValue)
		if err != nil {
			return
		}
	}
  parentValue, ok := rpcStruct["parent"]
	if ok && parentValue != nil {
  	record.Parent, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "parent"), parentValue)
		if err != nil {
			return
		}
	}
  childrenValue, ok := rpcStruct["children"]
	if ok && childrenValue != nil {
  	record.Children, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "children"), childrenValue)
		if err != nil {
			return
		}
	}
  biosStringsValue, ok := rpcStruct["bios_strings"]
	if ok && biosStringsValue != nil {
  	record.BiosStrings, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "bios_strings"), biosStringsValue)
		if err != nil {
			return
		}
	}
  protectionPolicyValue, ok := rpcStruct["protection_policy"]
	if ok && protectionPolicyValue != nil {
  	record.ProtectionPolicy, err = convertVMPPRefToGo(fmt.Sprintf("%s.%s", context, "protection_policy"), protectionPolicyValue)
		if err != nil {
			return
		}
	}
  isSnapshotFromVmppValue, ok := rpcStruct["is_snapshot_from_vmpp"]
	if ok && isSnapshotFromVmppValue != nil {
  	record.IsSnapshotFromVmpp, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_snapshot_from_vmpp"), isSnapshotFromVmppValue)
		if err != nil {
			return
		}
	}
  snapshotScheduleValue, ok := rpcStruct["snapshot_schedule"]
	if ok && snapshotScheduleValue != nil {
  	record.SnapshotSchedule, err = convertVMSSRefToGo(fmt.Sprintf("%s.%s", context, "snapshot_schedule"), snapshotScheduleValue)
		if err != nil {
			return
		}
	}
  isVmssSnapshotValue, ok := rpcStruct["is_vmss_snapshot"]
	if ok && isVmssSnapshotValue != nil {
  	record.IsVmssSnapshot, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_vmss_snapshot"), isVmssSnapshotValue)
		if err != nil {
			return
		}
	}
  applianceValue, ok := rpcStruct["appliance"]
	if ok && applianceValue != nil {
  	record.Appliance, err = convertVMApplianceRefToGo(fmt.Sprintf("%s.%s", context, "appliance"), applianceValue)
		if err != nil {
			return
		}
	}
  startDelayValue, ok := rpcStruct["start_delay"]
	if ok && startDelayValue != nil {
  	record.StartDelay, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "start_delay"), startDelayValue)
		if err != nil {
			return
		}
	}
  shutdownDelayValue, ok := rpcStruct["shutdown_delay"]
	if ok && shutdownDelayValue != nil {
  	record.ShutdownDelay, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "shutdown_delay"), shutdownDelayValue)
		if err != nil {
			return
		}
	}
  orderValue, ok := rpcStruct["order"]
	if ok && orderValue != nil {
  	record.Order, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "order"), orderValue)
		if err != nil {
			return
		}
	}
  vgpusValue, ok := rpcStruct["VGPUs"]
	if ok && vgpusValue != nil {
  	record.VGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "VGPUs"), vgpusValue)
		if err != nil {
			return
		}
	}
  attachedPCIsValue, ok := rpcStruct["attached_PCIs"]
	if ok && attachedPCIsValue != nil {
  	record.AttachedPCIs, err = convertPCIRefSetToGo(fmt.Sprintf("%s.%s", context, "attached_PCIs"), attachedPCIsValue)
		if err != nil {
			return
		}
	}
  suspendSRValue, ok := rpcStruct["suspend_SR"]
	if ok && suspendSRValue != nil {
  	record.SuspendSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "suspend_SR"), suspendSRValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  generationIDValue, ok := rpcStruct["generation_id"]
	if ok && generationIDValue != nil {
  	record.GenerationID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "generation_id"), generationIDValue)
		if err != nil {
			return
		}
	}
  hardwarePlatformVersionValue, ok := rpcStruct["hardware_platform_version"]
	if ok && hardwarePlatformVersionValue != nil {
  	record.HardwarePlatformVersion, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "hardware_platform_version"), hardwarePlatformVersionValue)
		if err != nil {
			return
		}
	}
  hasVendorDeviceValue, ok := rpcStruct["has_vendor_device"]
	if ok && hasVendorDeviceValue != nil {
  	record.HasVendorDevice, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "has_vendor_device"), hasVendorDeviceValue)
		if err != nil {
			return
		}
	}
  requiresRebootValue, ok := rpcStruct["requires_reboot"]
	if ok && requiresRebootValue != nil {
  	record.RequiresReboot, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "requires_reboot"), requiresRebootValue)
		if err != nil {
			return
		}
	}
  referenceLabelValue, ok := rpcStruct["reference_label"]
	if ok && referenceLabelValue != nil {
  	record.ReferenceLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "reference_label"), referenceLabelValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMRecordToXen(context string, record VMRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVMOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVMOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["power_state"], err = convertEnumVMPowerStateToXen(fmt.Sprintf("%s.%s", context, "power_state"), record.PowerState)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["user_version"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "user_version"), record.UserVersion)
  if err != nil {
		return
	}
  rpcStruct["is_a_template"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_a_template"), record.IsATemplate)
  if err != nil {
		return
	}
  rpcStruct["is_default_template"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_default_template"), record.IsDefaultTemplate)
  if err != nil {
		return
	}
  rpcStruct["suspend_VDI"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "suspend_VDI"), record.SuspendVDI)
  if err != nil {
		return
	}
  rpcStruct["resident_on"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "resident_on"), record.ResidentOn)
  if err != nil {
		return
	}
  rpcStruct["affinity"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "affinity"), record.Affinity)
  if err != nil {
		return
	}
  rpcStruct["memory_overhead"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_overhead"), record.MemoryOverhead)
  if err != nil {
		return
	}
  rpcStruct["memory_target"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_target"), record.MemoryTarget)
  if err != nil {
		return
	}
  rpcStruct["memory_static_max"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_static_max"), record.MemoryStaticMax)
  if err != nil {
		return
	}
  rpcStruct["memory_dynamic_max"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_dynamic_max"), record.MemoryDynamicMax)
  if err != nil {
		return
	}
  rpcStruct["memory_dynamic_min"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_dynamic_min"), record.MemoryDynamicMin)
  if err != nil {
		return
	}
  rpcStruct["memory_static_min"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_static_min"), record.MemoryStaticMin)
  if err != nil {
		return
	}
  rpcStruct["VCPUs_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "VCPUs_params"), record.VCPUsParams)
  if err != nil {
		return
	}
  rpcStruct["VCPUs_max"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "VCPUs_max"), record.VCPUsMax)
  if err != nil {
		return
	}
  rpcStruct["VCPUs_at_startup"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "VCPUs_at_startup"), record.VCPUsAtStartup)
  if err != nil {
		return
	}
  rpcStruct["actions_after_shutdown"], err = convertEnumOnNormalExitToXen(fmt.Sprintf("%s.%s", context, "actions_after_shutdown"), record.ActionsAfterShutdown)
  if err != nil {
		return
	}
  rpcStruct["actions_after_reboot"], err = convertEnumOnNormalExitToXen(fmt.Sprintf("%s.%s", context, "actions_after_reboot"), record.ActionsAfterReboot)
  if err != nil {
		return
	}
  rpcStruct["actions_after_crash"], err = convertEnumOnCrashBehaviourToXen(fmt.Sprintf("%s.%s", context, "actions_after_crash"), record.ActionsAfterCrash)
  if err != nil {
		return
	}
  rpcStruct["consoles"], err = convertConsoleRefSetToXen(fmt.Sprintf("%s.%s", context, "consoles"), record.Consoles)
  if err != nil {
		return
	}
  rpcStruct["VIFs"], err = convertVIFRefSetToXen(fmt.Sprintf("%s.%s", context, "VIFs"), record.VIFs)
  if err != nil {
		return
	}
  rpcStruct["VBDs"], err = convertVBDRefSetToXen(fmt.Sprintf("%s.%s", context, "VBDs"), record.VBDs)
  if err != nil {
		return
	}
  rpcStruct["VUSBs"], err = convertVUSBRefSetToXen(fmt.Sprintf("%s.%s", context, "VUSBs"), record.VUSBs)
  if err != nil {
		return
	}
  rpcStruct["crash_dumps"], err = convertCrashdumpRefSetToXen(fmt.Sprintf("%s.%s", context, "crash_dumps"), record.CrashDumps)
  if err != nil {
		return
	}
  rpcStruct["VTPMs"], err = convertVTPMRefSetToXen(fmt.Sprintf("%s.%s", context, "VTPMs"), record.VTPMs)
  if err != nil {
		return
	}
  rpcStruct["PV_bootloader"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_bootloader"), record.PVBootloader)
  if err != nil {
		return
	}
  rpcStruct["PV_kernel"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_kernel"), record.PVKernel)
  if err != nil {
		return
	}
  rpcStruct["PV_ramdisk"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_ramdisk"), record.PVRamdisk)
  if err != nil {
		return
	}
  rpcStruct["PV_args"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_args"), record.PVArgs)
  if err != nil {
		return
	}
  rpcStruct["PV_bootloader_args"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_bootloader_args"), record.PVBootloaderArgs)
  if err != nil {
		return
	}
  rpcStruct["PV_legacy_args"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_legacy_args"), record.PVLegacyArgs)
  if err != nil {
		return
	}
  rpcStruct["HVM_boot_policy"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "HVM_boot_policy"), record.HVMBootPolicy)
  if err != nil {
		return
	}
  rpcStruct["HVM_boot_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "HVM_boot_params"), record.HVMBootParams)
  if err != nil {
		return
	}
  rpcStruct["HVM_shadow_multiplier"], err = convertFloatToXen(fmt.Sprintf("%s.%s", context, "HVM_shadow_multiplier"), record.HVMShadowMultiplier)
  if err != nil {
		return
	}
  rpcStruct["platform"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "platform"), record.Platform)
  if err != nil {
		return
	}
  rpcStruct["PCI_bus"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PCI_bus"), record.PCIBus)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["domid"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "domid"), record.Domid)
  if err != nil {
		return
	}
  rpcStruct["domarch"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "domarch"), record.Domarch)
  if err != nil {
		return
	}
  rpcStruct["last_boot_CPU_flags"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "last_boot_CPU_flags"), record.LastBootCPUFlags)
  if err != nil {
		return
	}
  rpcStruct["is_control_domain"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_control_domain"), record.IsControlDomain)
  if err != nil {
		return
	}
  rpcStruct["metrics"], err = convertVMMetricsRefToXen(fmt.Sprintf("%s.%s", context, "metrics"), record.Metrics)
  if err != nil {
		return
	}
  rpcStruct["guest_metrics"], err = convertVMGuestMetricsRefToXen(fmt.Sprintf("%s.%s", context, "guest_metrics"), record.GuestMetrics)
  if err != nil {
		return
	}
  rpcStruct["last_booted_record"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "last_booted_record"), record.LastBootedRecord)
  if err != nil {
		return
	}
  rpcStruct["recommendations"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "recommendations"), record.Recommendations)
  if err != nil {
		return
	}
  rpcStruct["xenstore_data"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "xenstore_data"), record.XenstoreData)
  if err != nil {
		return
	}
  rpcStruct["ha_always_run"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "ha_always_run"), record.HaAlwaysRun)
  if err != nil {
		return
	}
  rpcStruct["ha_restart_priority"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "ha_restart_priority"), record.HaRestartPriority)
  if err != nil {
		return
	}
  rpcStruct["is_a_snapshot"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), record.IsASnapshot)
  if err != nil {
		return
	}
  rpcStruct["snapshot_of"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "snapshot_of"), record.SnapshotOf)
  if err != nil {
		return
	}
  rpcStruct["snapshots"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "snapshots"), record.Snapshots)
  if err != nil {
		return
	}
  rpcStruct["snapshot_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "snapshot_time"), record.SnapshotTime)
  if err != nil {
		return
	}
  rpcStruct["transportable_snapshot_id"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "transportable_snapshot_id"), record.TransportableSnapshotID)
  if err != nil {
		return
	}
  rpcStruct["blobs"], err = convertStringToBlobRefMapToXen(fmt.Sprintf("%s.%s", context, "blobs"), record.Blobs)
  if err != nil {
		return
	}
  rpcStruct["tags"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "tags"), record.Tags)
  if err != nil {
		return
	}
  rpcStruct["blocked_operations"], err = convertEnumVMOperationsToStringMapToXen(fmt.Sprintf("%s.%s", context, "blocked_operations"), record.BlockedOperations)
  if err != nil {
		return
	}
  rpcStruct["snapshot_info"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "snapshot_info"), record.SnapshotInfo)
  if err != nil {
		return
	}
  rpcStruct["snapshot_metadata"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "snapshot_metadata"), record.SnapshotMetadata)
  if err != nil {
		return
	}
  rpcStruct["parent"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "parent"), record.Parent)
  if err != nil {
		return
	}
  rpcStruct["children"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "children"), record.Children)
  if err != nil {
		return
	}
  rpcStruct["bios_strings"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "bios_strings"), record.BiosStrings)
  if err != nil {
		return
	}
  rpcStruct["protection_policy"], err = convertVMPPRefToXen(fmt.Sprintf("%s.%s", context, "protection_policy"), record.ProtectionPolicy)
  if err != nil {
		return
	}
  rpcStruct["is_snapshot_from_vmpp"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_snapshot_from_vmpp"), record.IsSnapshotFromVmpp)
  if err != nil {
		return
	}
  rpcStruct["snapshot_schedule"], err = convertVMSSRefToXen(fmt.Sprintf("%s.%s", context, "snapshot_schedule"), record.SnapshotSchedule)
  if err != nil {
		return
	}
  rpcStruct["is_vmss_snapshot"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_vmss_snapshot"), record.IsVmssSnapshot)
  if err != nil {
		return
	}
  rpcStruct["appliance"], err = convertVMApplianceRefToXen(fmt.Sprintf("%s.%s", context, "appliance"), record.Appliance)
  if err != nil {
		return
	}
  rpcStruct["start_delay"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "start_delay"), record.StartDelay)
  if err != nil {
		return
	}
  rpcStruct["shutdown_delay"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "shutdown_delay"), record.ShutdownDelay)
  if err != nil {
		return
	}
  rpcStruct["order"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "order"), record.Order)
  if err != nil {
		return
	}
  rpcStruct["VGPUs"], err = convertVGPURefSetToXen(fmt.Sprintf("%s.%s", context, "VGPUs"), record.VGPUs)
  if err != nil {
		return
	}
  rpcStruct["attached_PCIs"], err = convertPCIRefSetToXen(fmt.Sprintf("%s.%s", context, "attached_PCIs"), record.AttachedPCIs)
  if err != nil {
		return
	}
  rpcStruct["suspend_SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "suspend_SR"), record.SuspendSR)
  if err != nil {
		return
	}
  rpcStruct["version"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "version"), record.Version)
  if err != nil {
		return
	}
  rpcStruct["generation_id"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "generation_id"), record.GenerationID)
  if err != nil {
		return
	}
  rpcStruct["hardware_platform_version"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "hardware_platform_version"), record.HardwarePlatformVersion)
  if err != nil {
		return
	}
  rpcStruct["has_vendor_device"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "has_vendor_device"), record.HasVendorDevice)
  if err != nil {
		return
	}
  rpcStruct["requires_reboot"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "requires_reboot"), record.RequiresReboot)
  if err != nil {
		return
	}
  rpcStruct["reference_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "reference_label"), record.ReferenceLabel)
  if err != nil {
		return
	}
	return
}

func convertVMRefSetToGo(context string, input interface{}) (slice []VMRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMRefSetToXen(context string, slice []VMRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVMRefToGo(context string, input interface{}) (ref VMRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMRef(value)
	}
	return
}

func convertVMRefToXen(context string, ref VMRef) (string, error) {
	return string(ref), nil
}

func convertVMPPRecordToGo(context string, input interface{}) (record VMPPRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  isPolicyEnabledValue, ok := rpcStruct["is_policy_enabled"]
	if ok && isPolicyEnabledValue != nil {
  	record.IsPolicyEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_policy_enabled"), isPolicyEnabledValue)
		if err != nil {
			return
		}
	}
  backupTypeValue, ok := rpcStruct["backup_type"]
	if ok && backupTypeValue != nil {
  	record.BackupType, err = convertEnumVmppBackupTypeToGo(fmt.Sprintf("%s.%s", context, "backup_type"), backupTypeValue)
		if err != nil {
			return
		}
	}
  backupRetentionValueValue, ok := rpcStruct["backup_retention_value"]
	if ok && backupRetentionValueValue != nil {
  	record.BackupRetentionValue, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "backup_retention_value"), backupRetentionValueValue)
		if err != nil {
			return
		}
	}
  backupFrequencyValue, ok := rpcStruct["backup_frequency"]
	if ok && backupFrequencyValue != nil {
  	record.BackupFrequency, err = convertEnumVmppBackupFrequencyToGo(fmt.Sprintf("%s.%s", context, "backup_frequency"), backupFrequencyValue)
		if err != nil {
			return
		}
	}
  backupScheduleValue, ok := rpcStruct["backup_schedule"]
	if ok && backupScheduleValue != nil {
  	record.BackupSchedule, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "backup_schedule"), backupScheduleValue)
		if err != nil {
			return
		}
	}
  isBackupRunningValue, ok := rpcStruct["is_backup_running"]
	if ok && isBackupRunningValue != nil {
  	record.IsBackupRunning, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_backup_running"), isBackupRunningValue)
		if err != nil {
			return
		}
	}
  backupLastRunTimeValue, ok := rpcStruct["backup_last_run_time"]
	if ok && backupLastRunTimeValue != nil {
  	record.BackupLastRunTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "backup_last_run_time"), backupLastRunTimeValue)
		if err != nil {
			return
		}
	}
  archiveTargetTypeValue, ok := rpcStruct["archive_target_type"]
	if ok && archiveTargetTypeValue != nil {
  	record.ArchiveTargetType, err = convertEnumVmppArchiveTargetTypeToGo(fmt.Sprintf("%s.%s", context, "archive_target_type"), archiveTargetTypeValue)
		if err != nil {
			return
		}
	}
  archiveTargetConfigValue, ok := rpcStruct["archive_target_config"]
	if ok && archiveTargetConfigValue != nil {
  	record.ArchiveTargetConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "archive_target_config"), archiveTargetConfigValue)
		if err != nil {
			return
		}
	}
  archiveFrequencyValue, ok := rpcStruct["archive_frequency"]
	if ok && archiveFrequencyValue != nil {
  	record.ArchiveFrequency, err = convertEnumVmppArchiveFrequencyToGo(fmt.Sprintf("%s.%s", context, "archive_frequency"), archiveFrequencyValue)
		if err != nil {
			return
		}
	}
  archiveScheduleValue, ok := rpcStruct["archive_schedule"]
	if ok && archiveScheduleValue != nil {
  	record.ArchiveSchedule, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "archive_schedule"), archiveScheduleValue)
		if err != nil {
			return
		}
	}
  isArchiveRunningValue, ok := rpcStruct["is_archive_running"]
	if ok && isArchiveRunningValue != nil {
  	record.IsArchiveRunning, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_archive_running"), isArchiveRunningValue)
		if err != nil {
			return
		}
	}
  archiveLastRunTimeValue, ok := rpcStruct["archive_last_run_time"]
	if ok && archiveLastRunTimeValue != nil {
  	record.ArchiveLastRunTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "archive_last_run_time"), archiveLastRunTimeValue)
		if err != nil {
			return
		}
	}
  vmsValue, ok := rpcStruct["VMs"]
	if ok && vmsValue != nil {
  	record.VMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "VMs"), vmsValue)
		if err != nil {
			return
		}
	}
  isAlarmEnabledValue, ok := rpcStruct["is_alarm_enabled"]
	if ok && isAlarmEnabledValue != nil {
  	record.IsAlarmEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_alarm_enabled"), isAlarmEnabledValue)
		if err != nil {
			return
		}
	}
  alarmConfigValue, ok := rpcStruct["alarm_config"]
	if ok && alarmConfigValue != nil {
  	record.AlarmConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "alarm_config"), alarmConfigValue)
		if err != nil {
			return
		}
	}
  recentAlertsValue, ok := rpcStruct["recent_alerts"]
	if ok && recentAlertsValue != nil {
  	record.RecentAlerts, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "recent_alerts"), recentAlertsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMPPRecordToXen(context string, record VMPPRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["is_policy_enabled"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_policy_enabled"), record.IsPolicyEnabled)
  if err != nil {
		return
	}
  rpcStruct["backup_type"], err = convertEnumVmppBackupTypeToXen(fmt.Sprintf("%s.%s", context, "backup_type"), record.BackupType)
  if err != nil {
		return
	}
  rpcStruct["backup_retention_value"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "backup_retention_value"), record.BackupRetentionValue)
  if err != nil {
		return
	}
  rpcStruct["backup_frequency"], err = convertEnumVmppBackupFrequencyToXen(fmt.Sprintf("%s.%s", context, "backup_frequency"), record.BackupFrequency)
  if err != nil {
		return
	}
  rpcStruct["backup_schedule"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "backup_schedule"), record.BackupSchedule)
  if err != nil {
		return
	}
  rpcStruct["is_backup_running"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_backup_running"), record.IsBackupRunning)
  if err != nil {
		return
	}
  rpcStruct["backup_last_run_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "backup_last_run_time"), record.BackupLastRunTime)
  if err != nil {
		return
	}
  rpcStruct["archive_target_type"], err = convertEnumVmppArchiveTargetTypeToXen(fmt.Sprintf("%s.%s", context, "archive_target_type"), record.ArchiveTargetType)
  if err != nil {
		return
	}
  rpcStruct["archive_target_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "archive_target_config"), record.ArchiveTargetConfig)
  if err != nil {
		return
	}
  rpcStruct["archive_frequency"], err = convertEnumVmppArchiveFrequencyToXen(fmt.Sprintf("%s.%s", context, "archive_frequency"), record.ArchiveFrequency)
  if err != nil {
		return
	}
  rpcStruct["archive_schedule"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "archive_schedule"), record.ArchiveSchedule)
  if err != nil {
		return
	}
  rpcStruct["is_archive_running"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_archive_running"), record.IsArchiveRunning)
  if err != nil {
		return
	}
  rpcStruct["archive_last_run_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "archive_last_run_time"), record.ArchiveLastRunTime)
  if err != nil {
		return
	}
  rpcStruct["VMs"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "VMs"), record.VMs)
  if err != nil {
		return
	}
  rpcStruct["is_alarm_enabled"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_alarm_enabled"), record.IsAlarmEnabled)
  if err != nil {
		return
	}
  rpcStruct["alarm_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "alarm_config"), record.AlarmConfig)
  if err != nil {
		return
	}
  rpcStruct["recent_alerts"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "recent_alerts"), record.RecentAlerts)
  if err != nil {
		return
	}
	return
}

func convertVMPPRefSetToGo(context string, input interface{}) (slice []VMPPRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMPPRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMPPRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMPPRefToGo(context string, input interface{}) (ref VMPPRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMPPRef(value)
	}
	return
}

func convertVMPPRefToXen(context string, ref VMPPRef) (string, error) {
	return string(ref), nil
}

func convertVMSSRecordToGo(context string, input interface{}) (record VMSSRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  enabledValue, ok := rpcStruct["enabled"]
	if ok && enabledValue != nil {
  	record.Enabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enabled"), enabledValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertEnumVmssTypeToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  retainedSnapshotsValue, ok := rpcStruct["retained_snapshots"]
	if ok && retainedSnapshotsValue != nil {
  	record.RetainedSnapshots, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "retained_snapshots"), retainedSnapshotsValue)
		if err != nil {
			return
		}
	}
  frequencyValue, ok := rpcStruct["frequency"]
	if ok && frequencyValue != nil {
  	record.Frequency, err = convertEnumVmssFrequencyToGo(fmt.Sprintf("%s.%s", context, "frequency"), frequencyValue)
		if err != nil {
			return
		}
	}
  scheduleValue, ok := rpcStruct["schedule"]
	if ok && scheduleValue != nil {
  	record.Schedule, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "schedule"), scheduleValue)
		if err != nil {
			return
		}
	}
  lastRunTimeValue, ok := rpcStruct["last_run_time"]
	if ok && lastRunTimeValue != nil {
  	record.LastRunTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_run_time"), lastRunTimeValue)
		if err != nil {
			return
		}
	}
  vmsValue, ok := rpcStruct["VMs"]
	if ok && vmsValue != nil {
  	record.VMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "VMs"), vmsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMSSRecordToXen(context string, record VMSSRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["enabled"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "enabled"), record.Enabled)
  if err != nil {
		return
	}
  rpcStruct["type"], err = convertEnumVmssTypeToXen(fmt.Sprintf("%s.%s", context, "type"), record.Type)
  if err != nil {
		return
	}
  rpcStruct["retained_snapshots"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "retained_snapshots"), record.RetainedSnapshots)
  if err != nil {
		return
	}
  rpcStruct["frequency"], err = convertEnumVmssFrequencyToXen(fmt.Sprintf("%s.%s", context, "frequency"), record.Frequency)
  if err != nil {
		return
	}
  rpcStruct["schedule"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "schedule"), record.Schedule)
  if err != nil {
		return
	}
  rpcStruct["last_run_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "last_run_time"), record.LastRunTime)
  if err != nil {
		return
	}
  rpcStruct["VMs"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "VMs"), record.VMs)
  if err != nil {
		return
	}
	return
}

func convertVMSSRefSetToGo(context string, input interface{}) (slice []VMSSRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMSSRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMSSRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMSSRefToGo(context string, input interface{}) (ref VMSSRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMSSRef(value)
	}
	return
}

func convertVMSSRefToXen(context string, ref VMSSRef) (string, error) {
	return string(ref), nil
}

func convertVMApplianceRecordToGo(context string, input interface{}) (record VMApplianceRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumVMApplianceOperationSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumVMApplianceOperationMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vmsValue, ok := rpcStruct["VMs"]
	if ok && vmsValue != nil {
  	record.VMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "VMs"), vmsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMApplianceRecordToXen(context string, record VMApplianceRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVMApplianceOperationSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVMApplianceOperationMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["VMs"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "VMs"), record.VMs)
  if err != nil {
		return
	}
	return
}

func convertVMApplianceRefSetToGo(context string, input interface{}) (slice []VMApplianceRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMApplianceRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMApplianceRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMApplianceRefToGo(context string, input interface{}) (ref VMApplianceRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMApplianceRef(value)
	}
	return
}

func convertVMApplianceRefToXen(context string, ref VMApplianceRef) (string, error) {
	return string(ref), nil
}

func convertVMGuestMetricsRecordToGo(context string, input interface{}) (record VMGuestMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  osVersionValue, ok := rpcStruct["os_version"]
	if ok && osVersionValue != nil {
  	record.OSVersion, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "os_version"), osVersionValue)
		if err != nil {
			return
		}
	}
  pvDriversVersionValue, ok := rpcStruct["PV_drivers_version"]
	if ok && pvDriversVersionValue != nil {
  	record.PVDriversVersion, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "PV_drivers_version"), pvDriversVersionValue)
		if err != nil {
			return
		}
	}
  pvDriversUpToDateValue, ok := rpcStruct["PV_drivers_up_to_date"]
	if ok && pvDriversUpToDateValue != nil {
  	record.PVDriversUpToDate, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "PV_drivers_up_to_date"), pvDriversUpToDateValue)
		if err != nil {
			return
		}
	}
  memoryValue, ok := rpcStruct["memory"]
	if ok && memoryValue != nil {
  	record.Memory, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "memory"), memoryValue)
		if err != nil {
			return
		}
	}
  disksValue, ok := rpcStruct["disks"]
	if ok && disksValue != nil {
  	record.Disks, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "disks"), disksValue)
		if err != nil {
			return
		}
	}
  networksValue, ok := rpcStruct["networks"]
	if ok && networksValue != nil {
  	record.Networks, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "networks"), networksValue)
		if err != nil {
			return
		}
	}
  otherValue, ok := rpcStruct["other"]
	if ok && otherValue != nil {
  	record.Other, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other"), otherValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  liveValue, ok := rpcStruct["live"]
	if ok && liveValue != nil {
  	record.Live, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "live"), liveValue)
		if err != nil {
			return
		}
	}
  canUseHotplugVbdValue, ok := rpcStruct["can_use_hotplug_vbd"]
	if ok && canUseHotplugVbdValue != nil {
  	record.CanUseHotplugVbd, err = convertEnumTristateTypeToGo(fmt.Sprintf("%s.%s", context, "can_use_hotplug_vbd"), canUseHotplugVbdValue)
		if err != nil {
			return
		}
	}
  canUseHotplugVifValue, ok := rpcStruct["can_use_hotplug_vif"]
	if ok && canUseHotplugVifValue != nil {
  	record.CanUseHotplugVif, err = convertEnumTristateTypeToGo(fmt.Sprintf("%s.%s", context, "can_use_hotplug_vif"), canUseHotplugVifValue)
		if err != nil {
			return
		}
	}
  pvDriversDetectedValue, ok := rpcStruct["PV_drivers_detected"]
	if ok && pvDriversDetectedValue != nil {
  	record.PVDriversDetected, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "PV_drivers_detected"), pvDriversDetectedValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMGuestMetricsRefSetToGo(context string, input interface{}) (slice []VMGuestMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMGuestMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMGuestMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMGuestMetricsRefToGo(context string, input interface{}) (ref VMGuestMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMGuestMetricsRef(value)
	}
	return
}

func convertVMGuestMetricsRefToXen(context string, ref VMGuestMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVMMetricsRecordToGo(context string, input interface{}) (record VMMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  memoryActualValue, ok := rpcStruct["memory_actual"]
	if ok && memoryActualValue != nil {
  	record.MemoryActual, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_actual"), memoryActualValue)
		if err != nil {
			return
		}
	}
  vcpusNumberValue, ok := rpcStruct["VCPUs_number"]
	if ok && vcpusNumberValue != nil {
  	record.VCPUsNumber, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VCPUs_number"), vcpusNumberValue)
		if err != nil {
			return
		}
	}
  vcpusUtilisationValue, ok := rpcStruct["VCPUs_utilisation"]
	if ok && vcpusUtilisationValue != nil {
  	record.VCPUsUtilisation, err = convertIntToFloatMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_utilisation"), vcpusUtilisationValue)
		if err != nil {
			return
		}
	}
  vcpusCPUValue, ok := rpcStruct["VCPUs_CPU"]
	if ok && vcpusCPUValue != nil {
  	record.VCPUsCPU, err = convertIntToIntMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_CPU"), vcpusCPUValue)
		if err != nil {
			return
		}
	}
  vcpusParamsValue, ok := rpcStruct["VCPUs_params"]
	if ok && vcpusParamsValue != nil {
  	record.VCPUsParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_params"), vcpusParamsValue)
		if err != nil {
			return
		}
	}
  vcpusFlagsValue, ok := rpcStruct["VCPUs_flags"]
	if ok && vcpusFlagsValue != nil {
  	record.VCPUsFlags, err = convertIntToStringSetMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_flags"), vcpusFlagsValue)
		if err != nil {
			return
		}
	}
  stateValue, ok := rpcStruct["state"]
	if ok && stateValue != nil {
  	record.State, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "state"), stateValue)
		if err != nil {
			return
		}
	}
  startTimeValue, ok := rpcStruct["start_time"]
	if ok && startTimeValue != nil {
  	record.StartTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "start_time"), startTimeValue)
		if err != nil {
			return
		}
	}
  installTimeValue, ok := rpcStruct["install_time"]
	if ok && installTimeValue != nil {
  	record.InstallTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "install_time"), installTimeValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  hvmValue, ok := rpcStruct["hvm"]
	if ok && hvmValue != nil {
  	record.Hvm, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "hvm"), hvmValue)
		if err != nil {
			return
		}
	}
  nestedVirtValue, ok := rpcStruct["nested_virt"]
	if ok && nestedVirtValue != nil {
  	record.NestedVirt, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "nested_virt"), nestedVirtValue)
		if err != nil {
			return
		}
	}
  nomigrateValue, ok := rpcStruct["nomigrate"]
	if ok && nomigrateValue != nil {
  	record.Nomigrate, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "nomigrate"), nomigrateValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMMetricsRefSetToGo(context string, input interface{}) (slice []VMMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMMetricsRefToGo(context string, input interface{}) (ref VMMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMMetricsRef(value)
	}
	return
}

func convertVMMetricsRefToXen(context string, ref VMMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVTPMRecordToGo(context string, input interface{}) (record VTPMRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  backendValue, ok := rpcStruct["backend"]
	if ok && backendValue != nil {
  	record.Backend, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "backend"), backendValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVTPMRecordToXen(context string, record VTPMRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["backend"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "backend"), record.Backend)
  if err != nil {
		return
	}
	return
}

func convertVTPMRefSetToGo(context string, input interface{}) (slice []VTPMRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VTPMRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVTPMRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVTPMRefSetToXen(context string, slice []VTPMRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVTPMRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVTPMRefToGo(context string, input interface{}) (ref VTPMRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VTPMRef(value)
	}
	return
}

func convertVTPMRefToXen(context string, ref VTPMRef) (string, error) {
	return string(ref), nil
}

func convertVUSBRecordToGo(context string, input interface{}) (record VUSBRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumVusbOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumVusbOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  usbGroupValue, ok := rpcStruct["USB_group"]
	if ok && usbGroupValue != nil {
  	record.USBGroup, err = convertUSBGroupRefToGo(fmt.Sprintf("%s.%s", context, "USB_group"), usbGroupValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok && currentlyAttachedValue != nil {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVUSBRefSetToGo(context string, input interface{}) (slice []VUSBRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VUSBRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVUSBRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVUSBRefSetToXen(context string, slice []VUSBRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVUSBRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVUSBRefToGo(context string, input interface{}) (ref VUSBRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VUSBRef(value)
	}
	return
}

func convertVUSBRefToXen(context string, ref VUSBRef) (string, error) {
	return string(ref), nil
}

func convertBlobRecordToGo(context string, input interface{}) (record BlobRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok && sizeValue != nil {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  publicValue, ok := rpcStruct["public"]
	if ok && publicValue != nil {
  	record.Public, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "public"), publicValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  mimeTypeValue, ok := rpcStruct["mime_type"]
	if ok && mimeTypeValue != nil {
  	record.MimeType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "mime_type"), mimeTypeValue)
		if err != nil {
			return
		}
	}
	return
}

func convertBlobRefSetToGo(context string, input interface{}) (slice []BlobRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]BlobRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertBlobRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertBlobRefToGo(context string, input interface{}) (ref BlobRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = BlobRef(value)
	}
	return
}

func convertBlobRefToXen(context string, ref BlobRef) (string, error) {
	return string(ref), nil
}

func convertBoolToGo(context string, input interface{}) (value bool, err error) {
	if input == nil {
		return
	}
	value, ok := input.(bool)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "bool", context, reflect.TypeOf(input), input)
	}
	return
}

func convertBoolToXen(context string, value bool) (bool, error) {
	return value, nil
}

func convertConsoleRecordToGo(context string, input interface{}) (record ConsoleRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  protocolValue, ok := rpcStruct["protocol"]
	if ok && protocolValue != nil {
  	record.Protocol, err = convertEnumConsoleProtocolToGo(fmt.Sprintf("%s.%s", context, "protocol"), protocolValue)
		if err != nil {
			return
		}
	}
  locationValue, ok := rpcStruct["location"]
	if ok && locationValue != nil {
  	record.Location, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "location"), locationValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertConsoleRecordToXen(context string, record ConsoleRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["protocol"], err = convertEnumConsoleProtocolToXen(fmt.Sprintf("%s.%s", context, "protocol"), record.Protocol)
  if err != nil {
		return
	}
  rpcStruct["location"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "location"), record.Location)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertConsoleRefSetToGo(context string, input interface{}) (slice []ConsoleRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]ConsoleRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertConsoleRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertConsoleRefSetToXen(context string, slice []ConsoleRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertConsoleRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertConsoleRefToGo(context string, input interface{}) (ref ConsoleRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = ConsoleRef(value)
	}
	return
}

func convertConsoleRefToXen(context string, ref ConsoleRef) (string, error) {
	return string(ref), nil
}

func convertCrashdumpRecordToGo(context string, input interface{}) (record CrashdumpRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok && vmValue != nil {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  vdiValue, ok := rpcStruct["VDI"]
	if ok && vdiValue != nil {
  	record.VDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "VDI"), vdiValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertCrashdumpRefSetToGo(context string, input interface{}) (slice []CrashdumpRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]CrashdumpRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertCrashdumpRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertCrashdumpRefSetToXen(context string, slice []CrashdumpRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertCrashdumpRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertCrashdumpRefToGo(context string, input interface{}) (ref CrashdumpRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = CrashdumpRef(value)
	}
	return
}

func convertCrashdumpRefToXen(context string, ref CrashdumpRef) (string, error) {
	return string(ref), nil
}

func convertDataSourceRecordSetToGo(context string, input interface{}) (slice []DataSourceRecord, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]DataSourceRecord, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertDataSourceRecordToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertDataSourceRecordToGo(context string, input interface{}) (record DataSourceRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  enabledValue, ok := rpcStruct["enabled"]
	if ok && enabledValue != nil {
  	record.Enabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enabled"), enabledValue)
		if err != nil {
			return
		}
	}
  standardValue, ok := rpcStruct["standard"]
	if ok && standardValue != nil {
  	record.Standard, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "standard"), standardValue)
		if err != nil {
			return
		}
	}
  unitsValue, ok := rpcStruct["units"]
	if ok && unitsValue != nil {
  	record.Units, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "units"), unitsValue)
		if err != nil {
			return
		}
	}
  minValue, ok := rpcStruct["min"]
	if ok && minValue != nil {
  	record.Min, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "min"), minValue)
		if err != nil {
			return
		}
	}
  maxValue, ok := rpcStruct["max"]
	if ok && maxValue != nil {
  	record.Max, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "max"), maxValue)
		if err != nil {
			return
		}
	}
  valueValue, ok := rpcStruct["value"]
	if ok && valueValue != nil {
  	record.Value, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "value"), valueValue)
		if err != nil {
			return
		}
	}
	return
}

func convertTimeToGo(context string, input interface{}) (value time.Time, err error) {
	if input == nil {
		return
	}
	value, ok := input.(time.Time)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "time.Time", context, reflect.TypeOf(input), input)
	}
	return
}

func convertTimeToXen(context string, value time.Time) (time.Time, error) {
	return value, nil
}

func convertEnumAfterApplyGuidanceSetToGo(context string, input interface{}) (slice []AfterApplyGuidance, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]AfterApplyGuidance, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumAfterApplyGuidanceToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumAfterApplyGuidanceToGo(context string, input interface{}) (value AfterApplyGuidance, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "restartHVM":
      value = AfterApplyGuidanceRestartHVM
    case "restartPV":
      value = AfterApplyGuidanceRestartPV
    case "restartHost":
      value = AfterApplyGuidanceRestartHost
    case "restartXAPI":
      value = AfterApplyGuidanceRestartXAPI
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "AfterApplyGuidance", context)
	}
	return
}

func convertEnumAllocationAlgorithmToGo(context string, input interface{}) (value AllocationAlgorithm, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "breadth_first":
      value = AllocationAlgorithmBreadthFirst
    case "depth_first":
      value = AllocationAlgorithmDepthFirst
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "AllocationAlgorithm", context)
	}
	return
}

func convertEnumAllocationAlgorithmToXen(context string, value AllocationAlgorithm) (string, error) {
	return string(value), nil
}

func convertEnumBondModeToGo(context string, input interface{}) (value BondMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "balance-slb":
      value = BondModeBalanceSlb
    case "active-backup":
      value = BondModeActiveBackup
    case "lacp":
      value = BondModeLacp
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "BondMode", context)
	}
	return
}

func convertEnumBondModeToXen(context string, value BondMode) (string, error) {
	return string(value), nil
}

func convertEnumClsToGo(context string, input interface{}) (value Cls, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "VM":
      value = ClsVM
    case "Host":
      value = ClsHost
    case "SR":
      value = ClsSR
    case "Pool":
      value = ClsPool
    case "VMPP":
      value = ClsVMPP
    case "VMSS":
      value = ClsVMSS
    case "PVS_proxy":
      value = ClsPVSProxy
    case "VDI":
      value = ClsVDI
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "Cls", context)
	}
	return
}

func convertEnumClsToXen(context string, value Cls) (string, error) {
	return string(value), nil
}

func convertEnumConsoleProtocolToGo(context string, input interface{}) (value ConsoleProtocol, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "vt100":
      value = ConsoleProtocolVt100
    case "rfb":
      value = ConsoleProtocolRfb
    case "rdp":
      value = ConsoleProtocolRdp
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "ConsoleProtocol", context)
	}
	return
}

func convertEnumConsoleProtocolToXen(context string, value ConsoleProtocol) (string, error) {
	return string(value), nil
}

func convertEnumEventOperationToGo(context string, input interface{}) (value EventOperation, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "add":
      value = EventOperationAdd
    case "del":
      value = EventOperationDel
    case "mod":
      value = EventOperationMod
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "EventOperation", context)
	}
	return
}

func convertEnumHostAllowedOperationsSetToGo(context string, input interface{}) (slice []HostAllowedOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostAllowedOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumHostAllowedOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumHostAllowedOperationsToGo(context string, input interface{}) (value HostAllowedOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "provision":
      value = HostAllowedOperationsProvision
    case "evacuate":
      value = HostAllowedOperationsEvacuate
    case "shutdown":
      value = HostAllowedOperationsShutdown
    case "reboot":
      value = HostAllowedOperationsReboot
    case "power_on":
      value = HostAllowedOperationsPowerOn
    case "vm_start":
      value = HostAllowedOperationsVMStart
    case "vm_resume":
      value = HostAllowedOperationsVMResume
    case "vm_migrate":
      value = HostAllowedOperationsVMMigrate
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "HostAllowedOperations", context)
	}
	return
}

func convertEnumHostDisplayToGo(context string, input interface{}) (value HostDisplay, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "enabled":
      value = HostDisplayEnabled
    case "disable_on_reboot":
      value = HostDisplayDisableOnReboot
    case "disabled":
      value = HostDisplayDisabled
    case "enable_on_reboot":
      value = HostDisplayEnableOnReboot
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "HostDisplay", context)
	}
	return
}

func convertEnumHostDisplayToXen(context string, value HostDisplay) (string, error) {
	return string(value), nil
}

func convertEnumIPConfigurationModeToGo(context string, input interface{}) (value IPConfigurationMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "None":
      value = IPConfigurationModeNone
    case "DHCP":
      value = IPConfigurationModeDHCP
    case "Static":
      value = IPConfigurationModeStatic
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "IPConfigurationMode", context)
	}
	return
}

func convertEnumIPConfigurationModeToXen(context string, value IPConfigurationMode) (string, error) {
	return string(value), nil
}

func convertEnumIpv6ConfigurationModeToGo(context string, input interface{}) (value Ipv6ConfigurationMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "None":
      value = Ipv6ConfigurationModeNone
    case "DHCP":
      value = Ipv6ConfigurationModeDHCP
    case "Static":
      value = Ipv6ConfigurationModeStatic
    case "Autoconf":
      value = Ipv6ConfigurationModeAutoconf
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "Ipv6ConfigurationMode", context)
	}
	return
}

func convertEnumIpv6ConfigurationModeToXen(context string, value Ipv6ConfigurationMode) (string, error) {
	return string(value), nil
}

func convertEnumLivepatchStatusToGo(context string, input interface{}) (value LivepatchStatus, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "ok_livepatch_complete":
      value = LivepatchStatusOkLivepatchComplete
    case "ok_livepatch_incomplete":
      value = LivepatchStatusOkLivepatchIncomplete
    case "ok":
      value = LivepatchStatusOk
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "LivepatchStatus", context)
	}
	return
}

func convertEnumNetworkDefaultLockingModeToGo(context string, input interface{}) (value NetworkDefaultLockingMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "unlocked":
      value = NetworkDefaultLockingModeUnlocked
    case "disabled":
      value = NetworkDefaultLockingModeDisabled
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "NetworkDefaultLockingMode", context)
	}
	return
}

func convertEnumNetworkDefaultLockingModeToXen(context string, value NetworkDefaultLockingMode) (string, error) {
	return string(value), nil
}

func convertEnumNetworkOperationsSetToGo(context string, input interface{}) (slice []NetworkOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]NetworkOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumNetworkOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumNetworkOperationsSetToXen(context string, slice []NetworkOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumNetworkOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumNetworkOperationsToGo(context string, input interface{}) (value NetworkOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attaching":
      value = NetworkOperationsAttaching
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "NetworkOperations", context)
	}
	return
}

func convertEnumNetworkOperationsToXen(context string, value NetworkOperations) (string, error) {
	return string(value), nil
}

func convertEnumNetworkPurposeSetToGo(context string, input interface{}) (slice []NetworkPurpose, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]NetworkPurpose, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumNetworkPurposeToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumNetworkPurposeSetToXen(context string, slice []NetworkPurpose) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumNetworkPurposeToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumNetworkPurposeToGo(context string, input interface{}) (value NetworkPurpose, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "nbd":
      value = NetworkPurposeNbd
    case "insecure_nbd":
      value = NetworkPurposeInsecureNbd
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "NetworkPurpose", context)
	}
	return
}

func convertEnumNetworkPurposeToXen(context string, value NetworkPurpose) (string, error) {
	return string(value), nil
}

func convertEnumOnBootToGo(context string, input interface{}) (value OnBoot, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "reset":
      value = OnBootReset
    case "persist":
      value = OnBootPersist
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "OnBoot", context)
	}
	return
}

func convertEnumOnBootToXen(context string, value OnBoot) (string, error) {
	return string(value), nil
}

func convertEnumOnCrashBehaviourToGo(context string, input interface{}) (value OnCrashBehaviour, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "destroy":
      value = OnCrashBehaviourDestroy
    case "coredump_and_destroy":
      value = OnCrashBehaviourCoredumpAndDestroy
    case "restart":
      value = OnCrashBehaviourRestart
    case "coredump_and_restart":
      value = OnCrashBehaviourCoredumpAndRestart
    case "preserve":
      value = OnCrashBehaviourPreserve
    case "rename_restart":
      value = OnCrashBehaviourRenameRestart
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "OnCrashBehaviour", context)
	}
	return
}

func convertEnumOnCrashBehaviourToXen(context string, value OnCrashBehaviour) (string, error) {
	return string(value), nil
}

func convertEnumOnNormalExitToGo(context string, input interface{}) (value OnNormalExit, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "destroy":
      value = OnNormalExitDestroy
    case "restart":
      value = OnNormalExitRestart
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "OnNormalExit", context)
	}
	return
}

func convertEnumOnNormalExitToXen(context string, value OnNormalExit) (string, error) {
	return string(value), nil
}

func convertEnumPgpuDom0AccessToGo(context string, input interface{}) (value PgpuDom0Access, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "enabled":
      value = PgpuDom0AccessEnabled
    case "disable_on_reboot":
      value = PgpuDom0AccessDisableOnReboot
    case "disabled":
      value = PgpuDom0AccessDisabled
    case "enable_on_reboot":
      value = PgpuDom0AccessEnableOnReboot
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PgpuDom0Access", context)
	}
	return
}

func convertEnumPifIgmpStatusToGo(context string, input interface{}) (value PifIgmpStatus, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "enabled":
      value = PifIgmpStatusEnabled
    case "disabled":
      value = PifIgmpStatusDisabled
    case "unknown":
      value = PifIgmpStatusUnknown
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PifIgmpStatus", context)
	}
	return
}

func convertEnumPoolAllowedOperationsSetToGo(context string, input interface{}) (slice []PoolAllowedOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolAllowedOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumPoolAllowedOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumPoolAllowedOperationsToGo(context string, input interface{}) (value PoolAllowedOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "ha_enable":
      value = PoolAllowedOperationsHaEnable
    case "ha_disable":
      value = PoolAllowedOperationsHaDisable
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PoolAllowedOperations", context)
	}
	return
}

func convertEnumPrimaryAddressTypeToGo(context string, input interface{}) (value PrimaryAddressType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "IPv4":
      value = PrimaryAddressTypeIPv4
    case "IPv6":
      value = PrimaryAddressTypeIPv6
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PrimaryAddressType", context)
	}
	return
}

func convertEnumPrimaryAddressTypeToXen(context string, value PrimaryAddressType) (string, error) {
	return string(value), nil
}

func convertEnumPvsProxyStatusToGo(context string, input interface{}) (value PvsProxyStatus, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "stopped":
      value = PvsProxyStatusStopped
    case "initialised":
      value = PvsProxyStatusInitialised
    case "caching":
      value = PvsProxyStatusCaching
    case "incompatible_write_cache_mode":
      value = PvsProxyStatusIncompatibleWriteCacheMode
    case "incompatible_protocol_version":
      value = PvsProxyStatusIncompatibleProtocolVersion
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PvsProxyStatus", context)
	}
	return
}

func convertEnumSdnControllerProtocolToGo(context string, input interface{}) (value SdnControllerProtocol, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "ssl":
      value = SdnControllerProtocolSsl
    case "pssl":
      value = SdnControllerProtocolPssl
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "SdnControllerProtocol", context)
	}
	return
}

func convertEnumSdnControllerProtocolToXen(context string, value SdnControllerProtocol) (string, error) {
	return string(value), nil
}

func convertEnumStorageOperationsSetToGo(context string, input interface{}) (slice []StorageOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]StorageOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumStorageOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumStorageOperationsToGo(context string, input interface{}) (value StorageOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "scan":
      value = StorageOperationsScan
    case "destroy":
      value = StorageOperationsDestroy
    case "forget":
      value = StorageOperationsForget
    case "plug":
      value = StorageOperationsPlug
    case "unplug":
      value = StorageOperationsUnplug
    case "update":
      value = StorageOperationsUpdate
    case "vdi_create":
      value = StorageOperationsVdiCreate
    case "vdi_introduce":
      value = StorageOperationsVdiIntroduce
    case "vdi_destroy":
      value = StorageOperationsVdiDestroy
    case "vdi_resize":
      value = StorageOperationsVdiResize
    case "vdi_clone":
      value = StorageOperationsVdiClone
    case "vdi_snapshot":
      value = StorageOperationsVdiSnapshot
    case "vdi_mirror":
      value = StorageOperationsVdiMirror
    case "vdi_enable_cbt":
      value = StorageOperationsVdiEnableCbt
    case "vdi_disable_cbt":
      value = StorageOperationsVdiDisableCbt
    case "vdi_data_destroy":
      value = StorageOperationsVdiDataDestroy
    case "vdi_list_changed_blocks":
      value = StorageOperationsVdiListChangedBlocks
    case "vdi_set_on_boot":
      value = StorageOperationsVdiSetOnBoot
    case "pbd_create":
      value = StorageOperationsPbdCreate
    case "pbd_destroy":
      value = StorageOperationsPbdDestroy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "StorageOperations", context)
	}
	return
}

func convertEnumTaskAllowedOperationsSetToGo(context string, input interface{}) (slice []TaskAllowedOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]TaskAllowedOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumTaskAllowedOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumTaskAllowedOperationsToGo(context string, input interface{}) (value TaskAllowedOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "cancel":
      value = TaskAllowedOperationsCancel
    case "destroy":
      value = TaskAllowedOperationsDestroy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "TaskAllowedOperations", context)
	}
	return
}

func convertEnumTaskStatusTypeToGo(context string, input interface{}) (value TaskStatusType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "pending":
      value = TaskStatusTypePending
    case "success":
      value = TaskStatusTypeSuccess
    case "failure":
      value = TaskStatusTypeFailure
    case "cancelling":
      value = TaskStatusTypeCancelling
    case "cancelled":
      value = TaskStatusTypeCancelled
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "TaskStatusType", context)
	}
	return
}

func convertEnumTaskStatusTypeToXen(context string, value TaskStatusType) (string, error) {
	return string(value), nil
}

func convertEnumTristateTypeToGo(context string, input interface{}) (value TristateType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "yes":
      value = TristateTypeYes
    case "no":
      value = TristateTypeNo
    case "unspecified":
      value = TristateTypeUnspecified
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "TristateType", context)
	}
	return
}

func convertEnumUpdateAfterApplyGuidanceSetToGo(context string, input interface{}) (slice []UpdateAfterApplyGuidance, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]UpdateAfterApplyGuidance, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumUpdateAfterApplyGuidanceToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumUpdateAfterApplyGuidanceToGo(context string, input interface{}) (value UpdateAfterApplyGuidance, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "restartHVM":
      value = UpdateAfterApplyGuidanceRestartHVM
    case "restartPV":
      value = UpdateAfterApplyGuidanceRestartPV
    case "restartHost":
      value = UpdateAfterApplyGuidanceRestartHost
    case "restartXAPI":
      value = UpdateAfterApplyGuidanceRestartXAPI
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "UpdateAfterApplyGuidance", context)
	}
	return
}

func convertEnumVbdModeToGo(context string, input interface{}) (value VbdMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "RO":
      value = VbdModeRO
    case "RW":
      value = VbdModeRW
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VbdMode", context)
	}
	return
}

func convertEnumVbdModeToXen(context string, value VbdMode) (string, error) {
	return string(value), nil
}

func convertEnumVbdOperationsSetToGo(context string, input interface{}) (slice []VbdOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VbdOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVbdOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVbdOperationsSetToXen(context string, slice []VbdOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVbdOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVbdOperationsToGo(context string, input interface{}) (value VbdOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attach":
      value = VbdOperationsAttach
    case "eject":
      value = VbdOperationsEject
    case "insert":
      value = VbdOperationsInsert
    case "plug":
      value = VbdOperationsPlug
    case "unplug":
      value = VbdOperationsUnplug
    case "unplug_force":
      value = VbdOperationsUnplugForce
    case "pause":
      value = VbdOperationsPause
    case "unpause":
      value = VbdOperationsUnpause
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VbdOperations", context)
	}
	return
}

func convertEnumVbdOperationsToXen(context string, value VbdOperations) (string, error) {
	return string(value), nil
}

func convertEnumVbdTypeToGo(context string, input interface{}) (value VbdType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "CD":
      value = VbdTypeCD
    case "Disk":
      value = VbdTypeDisk
    case "Floppy":
      value = VbdTypeFloppy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VbdType", context)
	}
	return
}

func convertEnumVbdTypeToXen(context string, value VbdType) (string, error) {
	return string(value), nil
}

func convertEnumVdiOperationsSetToGo(context string, input interface{}) (slice []VdiOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VdiOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVdiOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVdiOperationsSetToXen(context string, slice []VdiOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVdiOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVdiOperationsToGo(context string, input interface{}) (value VdiOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "clone":
      value = VdiOperationsClone
    case "copy":
      value = VdiOperationsCopy
    case "resize":
      value = VdiOperationsResize
    case "resize_online":
      value = VdiOperationsResizeOnline
    case "snapshot":
      value = VdiOperationsSnapshot
    case "mirror":
      value = VdiOperationsMirror
    case "destroy":
      value = VdiOperationsDestroy
    case "forget":
      value = VdiOperationsForget
    case "update":
      value = VdiOperationsUpdate
    case "force_unlock":
      value = VdiOperationsForceUnlock
    case "generate_config":
      value = VdiOperationsGenerateConfig
    case "enable_cbt":
      value = VdiOperationsEnableCbt
    case "disable_cbt":
      value = VdiOperationsDisableCbt
    case "data_destroy":
      value = VdiOperationsDataDestroy
    case "list_changed_blocks":
      value = VdiOperationsListChangedBlocks
    case "set_on_boot":
      value = VdiOperationsSetOnBoot
    case "blocked":
      value = VdiOperationsBlocked
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VdiOperations", context)
	}
	return
}

func convertEnumVdiOperationsToXen(context string, value VdiOperations) (string, error) {
	return string(value), nil
}

func convertEnumVdiTypeToGo(context string, input interface{}) (value VdiType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "system":
      value = VdiTypeSystem
    case "user":
      value = VdiTypeUser
    case "ephemeral":
      value = VdiTypeEphemeral
    case "suspend":
      value = VdiTypeSuspend
    case "crashdump":
      value = VdiTypeCrashdump
    case "ha_statefile":
      value = VdiTypeHaStatefile
    case "metadata":
      value = VdiTypeMetadata
    case "redo_log":
      value = VdiTypeRedoLog
    case "rrd":
      value = VdiTypeRrd
    case "pvs_cache":
      value = VdiTypePvsCache
    case "cbt_metadata":
      value = VdiTypeCbtMetadata
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VdiType", context)
	}
	return
}

func convertEnumVdiTypeToXen(context string, value VdiType) (string, error) {
	return string(value), nil
}

func convertEnumVgpuTypeImplementationToGo(context string, input interface{}) (value VgpuTypeImplementation, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "passthrough":
      value = VgpuTypeImplementationPassthrough
    case "nvidia":
      value = VgpuTypeImplementationNvidia
    case "gvt_g":
      value = VgpuTypeImplementationGvtG
    case "mxgpu":
      value = VgpuTypeImplementationMxgpu
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VgpuTypeImplementation", context)
	}
	return
}

func convertEnumVifIpv4ConfigurationModeToGo(context string, input interface{}) (value VifIpv4ConfigurationMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "None":
      value = VifIpv4ConfigurationModeNone
    case "Static":
      value = VifIpv4ConfigurationModeStatic
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VifIpv4ConfigurationMode", context)
	}
	return
}

func convertEnumVifIpv4ConfigurationModeToXen(context string, value VifIpv4ConfigurationMode) (string, error) {
	return string(value), nil
}

func convertEnumVifIpv6ConfigurationModeToGo(context string, input interface{}) (value VifIpv6ConfigurationMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "None":
      value = VifIpv6ConfigurationModeNone
    case "Static":
      value = VifIpv6ConfigurationModeStatic
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VifIpv6ConfigurationMode", context)
	}
	return
}

func convertEnumVifIpv6ConfigurationModeToXen(context string, value VifIpv6ConfigurationMode) (string, error) {
	return string(value), nil
}

func convertEnumVifLockingModeToGo(context string, input interface{}) (value VifLockingMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "network_default":
      value = VifLockingModeNetworkDefault
    case "locked":
      value = VifLockingModeLocked
    case "unlocked":
      value = VifLockingModeUnlocked
    case "disabled":
      value = VifLockingModeDisabled
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VifLockingMode", context)
	}
	return
}

func convertEnumVifLockingModeToXen(context string, value VifLockingMode) (string, error) {
	return string(value), nil
}

func convertEnumVifOperationsSetToGo(context string, input interface{}) (slice []VifOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VifOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVifOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVifOperationsSetToXen(context string, slice []VifOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVifOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVifOperationsToGo(context string, input interface{}) (value VifOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attach":
      value = VifOperationsAttach
    case "plug":
      value = VifOperationsPlug
    case "unplug":
      value = VifOperationsUnplug
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VifOperations", context)
	}
	return
}

func convertEnumVifOperationsToXen(context string, value VifOperations) (string, error) {
	return string(value), nil
}

func convertEnumVMApplianceOperationSetToGo(context string, input interface{}) (slice []VMApplianceOperation, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMApplianceOperation, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMApplianceOperationToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVMApplianceOperationSetToXen(context string, slice []VMApplianceOperation) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMApplianceOperationToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVMApplianceOperationToGo(context string, input interface{}) (value VMApplianceOperation, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "start":
      value = VMApplianceOperationStart
    case "clean_shutdown":
      value = VMApplianceOperationCleanShutdown
    case "hard_shutdown":
      value = VMApplianceOperationHardShutdown
    case "shutdown":
      value = VMApplianceOperationShutdown
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VMApplianceOperation", context)
	}
	return
}

func convertEnumVMApplianceOperationToXen(context string, value VMApplianceOperation) (string, error) {
	return string(value), nil
}

func convertEnumVMOperationsSetToGo(context string, input interface{}) (slice []VMOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVMOperationsSetToXen(context string, slice []VMOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVMOperationsToGo(context string, input interface{}) (value VMOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "snapshot":
      value = VMOperationsSnapshot
    case "clone":
      value = VMOperationsClone
    case "copy":
      value = VMOperationsCopy
    case "create_template":
      value = VMOperationsCreateTemplate
    case "revert":
      value = VMOperationsRevert
    case "checkpoint":
      value = VMOperationsCheckpoint
    case "snapshot_with_quiesce":
      value = VMOperationsSnapshotWithQuiesce
    case "provision":
      value = VMOperationsProvision
    case "start":
      value = VMOperationsStart
    case "start_on":
      value = VMOperationsStartOn
    case "pause":
      value = VMOperationsPause
    case "unpause":
      value = VMOperationsUnpause
    case "clean_shutdown":
      value = VMOperationsCleanShutdown
    case "clean_reboot":
      value = VMOperationsCleanReboot
    case "hard_shutdown":
      value = VMOperationsHardShutdown
    case "power_state_reset":
      value = VMOperationsPowerStateReset
    case "hard_reboot":
      value = VMOperationsHardReboot
    case "suspend":
      value = VMOperationsSuspend
    case "csvm":
      value = VMOperationsCsvm
    case "resume":
      value = VMOperationsResume
    case "resume_on":
      value = VMOperationsResumeOn
    case "pool_migrate":
      value = VMOperationsPoolMigrate
    case "migrate_send":
      value = VMOperationsMigrateSend
    case "get_boot_record":
      value = VMOperationsGetBootRecord
    case "send_sysrq":
      value = VMOperationsSendSysrq
    case "send_trigger":
      value = VMOperationsSendTrigger
    case "query_services":
      value = VMOperationsQueryServices
    case "shutdown":
      value = VMOperationsShutdown
    case "call_plugin":
      value = VMOperationsCallPlugin
    case "changing_memory_live":
      value = VMOperationsChangingMemoryLive
    case "awaiting_memory_live":
      value = VMOperationsAwaitingMemoryLive
    case "changing_dynamic_range":
      value = VMOperationsChangingDynamicRange
    case "changing_static_range":
      value = VMOperationsChangingStaticRange
    case "changing_memory_limits":
      value = VMOperationsChangingMemoryLimits
    case "changing_shadow_memory":
      value = VMOperationsChangingShadowMemory
    case "changing_shadow_memory_live":
      value = VMOperationsChangingShadowMemoryLive
    case "changing_VCPUs":
      value = VMOperationsChangingVCPUs
    case "changing_VCPUs_live":
      value = VMOperationsChangingVCPUsLive
    case "assert_operation_valid":
      value = VMOperationsAssertOperationValid
    case "data_source_op":
      value = VMOperationsDataSourceOp
    case "update_allowed_operations":
      value = VMOperationsUpdateAllowedOperations
    case "make_into_template":
      value = VMOperationsMakeIntoTemplate
    case "import":
      value = VMOperationsImport
    case "export":
      value = VMOperationsExport
    case "metadata_export":
      value = VMOperationsMetadataExport
    case "reverting":
      value = VMOperationsReverting
    case "destroy":
      value = VMOperationsDestroy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VMOperations", context)
	}
	return
}

func convertEnumVMOperationsToXen(context string, value VMOperations) (string, error) {
	return string(value), nil
}

func convertEnumVMPowerStateToGo(context string, input interface{}) (value VMPowerState, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "Halted":
      value = VMPowerStateHalted
    case "Paused":
      value = VMPowerStatePaused
    case "Running":
      value = VMPowerStateRunning
    case "Suspended":
      value = VMPowerStateSuspended
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VMPowerState", context)
	}
	return
}

func convertEnumVMPowerStateToXen(context string, value VMPowerState) (string, error) {
	return string(value), nil
}

func convertEnumVmppArchiveFrequencyToGo(context string, input interface{}) (value VmppArchiveFrequency, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "never":
      value = VmppArchiveFrequencyNever
    case "always_after_backup":
      value = VmppArchiveFrequencyAlwaysAfterBackup
    case "daily":
      value = VmppArchiveFrequencyDaily
    case "weekly":
      value = VmppArchiveFrequencyWeekly
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppArchiveFrequency", context)
	}
	return
}

func convertEnumVmppArchiveFrequencyToXen(context string, value VmppArchiveFrequency) (string, error) {
	return string(value), nil
}

func convertEnumVmppArchiveTargetTypeToGo(context string, input interface{}) (value VmppArchiveTargetType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "none":
      value = VmppArchiveTargetTypeNone
    case "cifs":
      value = VmppArchiveTargetTypeCifs
    case "nfs":
      value = VmppArchiveTargetTypeNfs
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppArchiveTargetType", context)
	}
	return
}

func convertEnumVmppArchiveTargetTypeToXen(context string, value VmppArchiveTargetType) (string, error) {
	return string(value), nil
}

func convertEnumVmppBackupFrequencyToGo(context string, input interface{}) (value VmppBackupFrequency, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "hourly":
      value = VmppBackupFrequencyHourly
    case "daily":
      value = VmppBackupFrequencyDaily
    case "weekly":
      value = VmppBackupFrequencyWeekly
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppBackupFrequency", context)
	}
	return
}

func convertEnumVmppBackupFrequencyToXen(context string, value VmppBackupFrequency) (string, error) {
	return string(value), nil
}

func convertEnumVmppBackupTypeToGo(context string, input interface{}) (value VmppBackupType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "snapshot":
      value = VmppBackupTypeSnapshot
    case "checkpoint":
      value = VmppBackupTypeCheckpoint
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppBackupType", context)
	}
	return
}

func convertEnumVmppBackupTypeToXen(context string, value VmppBackupType) (string, error) {
	return string(value), nil
}

func convertEnumVmssFrequencyToGo(context string, input interface{}) (value VmssFrequency, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "hourly":
      value = VmssFrequencyHourly
    case "daily":
      value = VmssFrequencyDaily
    case "weekly":
      value = VmssFrequencyWeekly
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmssFrequency", context)
	}
	return
}

func convertEnumVmssFrequencyToXen(context string, value VmssFrequency) (string, error) {
	return string(value), nil
}

func convertEnumVmssTypeToGo(context string, input interface{}) (value VmssType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "snapshot":
      value = VmssTypeSnapshot
    case "checkpoint":
      value = VmssTypeCheckpoint
    case "snapshot_with_quiesce":
      value = VmssTypeSnapshotWithQuiesce
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmssType", context)
	}
	return
}

func convertEnumVmssTypeToXen(context string, value VmssType) (string, error) {
	return string(value), nil
}

func convertEnumVusbOperationsSetToGo(context string, input interface{}) (slice []VusbOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VusbOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVusbOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVusbOperationsToGo(context string, input interface{}) (value VusbOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attach":
      value = VusbOperationsAttach
    case "plug":
      value = VusbOperationsPlug
    case "unplug":
      value = VusbOperationsUnplug
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VusbOperations", context)
	}
	return
}

func convertEventRecordSetToGo(context string, input interface{}) (slice []EventRecord, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]EventRecord, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEventRecordToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEventRecordToGo(context string, input interface{}) (record EventRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  idValue, ok := rpcStruct["id"]
	if ok && idValue != nil {
  	record.ID, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "id"), idValue)
		if err != nil {
			return
		}
	}
  timestampValue, ok := rpcStruct["timestamp"]
	if ok && timestampValue != nil {
  	record.Timestamp, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp"), timestampValue)
		if err != nil {
			return
		}
	}
  classValue, ok := rpcStruct["class"]
	if ok && classValue != nil {
  	record.Class, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "class"), classValue)
		if err != nil {
			return
		}
	}
  operationValue, ok := rpcStruct["operation"]
	if ok && operationValue != nil {
  	record.Operation, err = convertEnumEventOperationToGo(fmt.Sprintf("%s.%s", context, "operation"), operationValue)
		if err != nil {
			return
		}
	}
  refValue, ok := rpcStruct["ref"]
	if ok && refValue != nil {
  	record.Ref, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ref"), refValue)
		if err != nil {
			return
		}
	}
  objUUIDValue, ok := rpcStruct["obj_uuid"]
	if ok && objUUIDValue != nil {
  	record.ObjUUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "obj_uuid"), objUUIDValue)
		if err != nil {
			return
		}
	}
	return
}

func convertFloatToGo(context string, input interface{}) (value float64, err error) {
	if input == nil {
		return
	}
	value, ok := input.(float64)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "float64", context, reflect.TypeOf(input), input)
	}
	return
}

func convertFloatToXen(context string, value float64) (float64, error) {
	return value, nil
}

func convertHostRecordToGo(context string, input interface{}) (record HostRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  memoryOverheadValue, ok := rpcStruct["memory_overhead"]
	if ok && memoryOverheadValue != nil {
  	record.MemoryOverhead, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_overhead"), memoryOverheadValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumHostAllowedOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumHostAllowedOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  apiVersionMajorValue, ok := rpcStruct["API_version_major"]
	if ok && apiVersionMajorValue != nil {
  	record.APIVersionMajor, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "API_version_major"), apiVersionMajorValue)
		if err != nil {
			return
		}
	}
  apiVersionMinorValue, ok := rpcStruct["API_version_minor"]
	if ok && apiVersionMinorValue != nil {
  	record.APIVersionMinor, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "API_version_minor"), apiVersionMinorValue)
		if err != nil {
			return
		}
	}
  apiVersionVendorValue, ok := rpcStruct["API_version_vendor"]
	if ok && apiVersionVendorValue != nil {
  	record.APIVersionVendor, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "API_version_vendor"), apiVersionVendorValue)
		if err != nil {
			return
		}
	}
  apiVersionVendorImplementationValue, ok := rpcStruct["API_version_vendor_implementation"]
	if ok && apiVersionVendorImplementationValue != nil {
  	record.APIVersionVendorImplementation, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "API_version_vendor_implementation"), apiVersionVendorImplementationValue)
		if err != nil {
			return
		}
	}
  enabledValue, ok := rpcStruct["enabled"]
	if ok && enabledValue != nil {
  	record.Enabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enabled"), enabledValue)
		if err != nil {
			return
		}
	}
  softwareVersionValue, ok := rpcStruct["software_version"]
	if ok && softwareVersionValue != nil {
  	record.SoftwareVersion, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "software_version"), softwareVersionValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  capabilitiesValue, ok := rpcStruct["capabilities"]
	if ok && capabilitiesValue != nil {
  	record.Capabilities, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "capabilities"), capabilitiesValue)
		if err != nil {
			return
		}
	}
  cpuConfigurationValue, ok := rpcStruct["cpu_configuration"]
	if ok && cpuConfigurationValue != nil {
  	record.CPUConfiguration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "cpu_configuration"), cpuConfigurationValue)
		if err != nil {
			return
		}
	}
  schedPolicyValue, ok := rpcStruct["sched_policy"]
	if ok && schedPolicyValue != nil {
  	record.SchedPolicy, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "sched_policy"), schedPolicyValue)
		if err != nil {
			return
		}
	}
  supportedBootloadersValue, ok := rpcStruct["supported_bootloaders"]
	if ok && supportedBootloadersValue != nil {
  	record.SupportedBootloaders, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "supported_bootloaders"), supportedBootloadersValue)
		if err != nil {
			return
		}
	}
  residentVMsValue, ok := rpcStruct["resident_VMs"]
	if ok && residentVMsValue != nil {
  	record.ResidentVMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "resident_VMs"), residentVMsValue)
		if err != nil {
			return
		}
	}
  loggingValue, ok := rpcStruct["logging"]
	if ok && loggingValue != nil {
  	record.Logging, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "logging"), loggingValue)
		if err != nil {
			return
		}
	}
  pifsValue, ok := rpcStruct["PIFs"]
	if ok && pifsValue != nil {
  	record.PIFs, err = convertPIFRefSetToGo(fmt.Sprintf("%s.%s", context, "PIFs"), pifsValue)
		if err != nil {
			return
		}
	}
  suspendImageSrValue, ok := rpcStruct["suspend_image_sr"]
	if ok && suspendImageSrValue != nil {
  	record.SuspendImageSr, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "suspend_image_sr"), suspendImageSrValue)
		if err != nil {
			return
		}
	}
  crashDumpSrValue, ok := rpcStruct["crash_dump_sr"]
	if ok && crashDumpSrValue != nil {
  	record.CrashDumpSr, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "crash_dump_sr"), crashDumpSrValue)
		if err != nil {
			return
		}
	}
  crashdumpsValue, ok := rpcStruct["crashdumps"]
	if ok && crashdumpsValue != nil {
  	record.Crashdumps, err = convertHostCrashdumpRefSetToGo(fmt.Sprintf("%s.%s", context, "crashdumps"), crashdumpsValue)
		if err != nil {
			return
		}
	}
  patchesValue, ok := rpcStruct["patches"]
	if ok && patchesValue != nil {
  	record.Patches, err = convertHostPatchRefSetToGo(fmt.Sprintf("%s.%s", context, "patches"), patchesValue)
		if err != nil {
			return
		}
	}
  updatesValue, ok := rpcStruct["updates"]
	if ok && updatesValue != nil {
  	record.Updates, err = convertPoolUpdateRefSetToGo(fmt.Sprintf("%s.%s", context, "updates"), updatesValue)
		if err != nil {
			return
		}
	}
  pbdsValue, ok := rpcStruct["PBDs"]
	if ok && pbdsValue != nil {
  	record.PBDs, err = convertPBDRefSetToGo(fmt.Sprintf("%s.%s", context, "PBDs"), pbdsValue)
		if err != nil {
			return
		}
	}
  hostCPUsValue, ok := rpcStruct["host_CPUs"]
	if ok && hostCPUsValue != nil {
  	record.HostCPUs, err = convertHostCPURefSetToGo(fmt.Sprintf("%s.%s", context, "host_CPUs"), hostCPUsValue)
		if err != nil {
			return
		}
	}
  cpuInfoValue, ok := rpcStruct["cpu_info"]
	if ok && cpuInfoValue != nil {
  	record.CPUInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "cpu_info"), cpuInfoValue)
		if err != nil {
			return
		}
	}
  hostnameValue, ok := rpcStruct["hostname"]
	if ok && hostnameValue != nil {
  	record.Hostname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "hostname"), hostnameValue)
		if err != nil {
			return
		}
	}
  addressValue, ok := rpcStruct["address"]
	if ok && addressValue != nil {
  	record.Address, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "address"), addressValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok && metricsValue != nil {
  	record.Metrics, err = convertHostMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  licenseParamsValue, ok := rpcStruct["license_params"]
	if ok && licenseParamsValue != nil {
  	record.LicenseParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "license_params"), licenseParamsValue)
		if err != nil {
			return
		}
	}
  haStatefilesValue, ok := rpcStruct["ha_statefiles"]
	if ok && haStatefilesValue != nil {
  	record.HaStatefiles, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ha_statefiles"), haStatefilesValue)
		if err != nil {
			return
		}
	}
  haNetworkPeersValue, ok := rpcStruct["ha_network_peers"]
	if ok && haNetworkPeersValue != nil {
  	record.HaNetworkPeers, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ha_network_peers"), haNetworkPeersValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok && blobsValue != nil {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok && tagsValue != nil {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  externalAuthTypeValue, ok := rpcStruct["external_auth_type"]
	if ok && externalAuthTypeValue != nil {
  	record.ExternalAuthType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "external_auth_type"), externalAuthTypeValue)
		if err != nil {
			return
		}
	}
  externalAuthServiceNameValue, ok := rpcStruct["external_auth_service_name"]
	if ok && externalAuthServiceNameValue != nil {
  	record.ExternalAuthServiceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "external_auth_service_name"), externalAuthServiceNameValue)
		if err != nil {
			return
		}
	}
  externalAuthConfigurationValue, ok := rpcStruct["external_auth_configuration"]
	if ok && externalAuthConfigurationValue != nil {
  	record.ExternalAuthConfiguration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "external_auth_configuration"), externalAuthConfigurationValue)
		if err != nil {
			return
		}
	}
  editionValue, ok := rpcStruct["edition"]
	if ok && editionValue != nil {
  	record.Edition, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "edition"), editionValue)
		if err != nil {
			return
		}
	}
  licenseServerValue, ok := rpcStruct["license_server"]
	if ok && licenseServerValue != nil {
  	record.LicenseServer, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "license_server"), licenseServerValue)
		if err != nil {
			return
		}
	}
  biosStringsValue, ok := rpcStruct["bios_strings"]
	if ok && biosStringsValue != nil {
  	record.BiosStrings, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "bios_strings"), biosStringsValue)
		if err != nil {
			return
		}
	}
  powerOnModeValue, ok := rpcStruct["power_on_mode"]
	if ok && powerOnModeValue != nil {
  	record.PowerOnMode, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "power_on_mode"), powerOnModeValue)
		if err != nil {
			return
		}
	}
  powerOnConfigValue, ok := rpcStruct["power_on_config"]
	if ok && powerOnConfigValue != nil {
  	record.PowerOnConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "power_on_config"), powerOnConfigValue)
		if err != nil {
			return
		}
	}
  localCacheSrValue, ok := rpcStruct["local_cache_sr"]
	if ok && localCacheSrValue != nil {
  	record.LocalCacheSr, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "local_cache_sr"), localCacheSrValue)
		if err != nil {
			return
		}
	}
  chipsetInfoValue, ok := rpcStruct["chipset_info"]
	if ok && chipsetInfoValue != nil {
  	record.ChipsetInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "chipset_info"), chipsetInfoValue)
		if err != nil {
			return
		}
	}
  pcisValue, ok := rpcStruct["PCIs"]
	if ok && pcisValue != nil {
  	record.PCIs, err = convertPCIRefSetToGo(fmt.Sprintf("%s.%s", context, "PCIs"), pcisValue)
		if err != nil {
			return
		}
	}
  pgpusValue, ok := rpcStruct["PGPUs"]
	if ok && pgpusValue != nil {
  	record.PGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "PGPUs"), pgpusValue)
		if err != nil {
			return
		}
	}
  pusbsValue, ok := rpcStruct["PUSBs"]
	if ok && pusbsValue != nil {
  	record.PUSBs, err = convertPUSBRefSetToGo(fmt.Sprintf("%s.%s", context, "PUSBs"), pusbsValue)
		if err != nil {
			return
		}
	}
  sslLegacyValue, ok := rpcStruct["ssl_legacy"]
	if ok && sslLegacyValue != nil {
  	record.SslLegacy, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ssl_legacy"), sslLegacyValue)
		if err != nil {
			return
		}
	}
  guestVCPUsParamsValue, ok := rpcStruct["guest_VCPUs_params"]
	if ok && guestVCPUsParamsValue != nil {
  	record.GuestVCPUsParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "guest_VCPUs_params"), guestVCPUsParamsValue)
		if err != nil {
			return
		}
	}
  displayValue, ok := rpcStruct["display"]
	if ok && displayValue != nil {
  	record.Display, err = convertEnumHostDisplayToGo(fmt.Sprintf("%s.%s", context, "display"), displayValue)
		if err != nil {
			return
		}
	}
  virtualHardwarePlatformVersionsValue, ok := rpcStruct["virtual_hardware_platform_versions"]
	if ok && virtualHardwarePlatformVersionsValue != nil {
  	record.VirtualHardwarePlatformVersions, err = convertIntSetToGo(fmt.Sprintf("%s.%s", context, "virtual_hardware_platform_versions"), virtualHardwarePlatformVersionsValue)
		if err != nil {
			return
		}
	}
  controlDomainValue, ok := rpcStruct["control_domain"]
	if ok && controlDomainValue != nil {
  	record.ControlDomain, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "control_domain"), controlDomainValue)
		if err != nil {
			return
		}
	}
  updatesRequiringRebootValue, ok := rpcStruct["updates_requiring_reboot"]
	if ok && updatesRequiringRebootValue != nil {
  	record.UpdatesRequiringReboot, err = convertPoolUpdateRefSetToGo(fmt.Sprintf("%s.%s", context, "updates_requiring_reboot"), updatesRequiringRebootValue)
		if err != nil {
			return
		}
	}
  featuresValue, ok := rpcStruct["features"]
	if ok && featuresValue != nil {
  	record.Features, err = convertFeatureRefSetToGo(fmt.Sprintf("%s.%s", context, "features"), featuresValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostRefSetToGo(context string, input interface{}) (slice []HostRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostRefSetToXen(context string, slice []HostRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertHostRefToGo(context string, input interface{}) (ref HostRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostRef(value)
	}
	return
}

func convertHostRefToXen(context string, ref HostRef) (string, error) {
	return string(ref), nil
}

func convertHostCPURecordToGo(context string, input interface{}) (record HostCPURecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  numberValue, ok := rpcStruct["number"]
	if ok && numberValue != nil {
  	record.Number, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "number"), numberValue)
		if err != nil {
			return
		}
	}
  vendorValue, ok := rpcStruct["vendor"]
	if ok && vendorValue != nil {
  	record.Vendor, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor"), vendorValue)
		if err != nil {
			return
		}
	}
  speedValue, ok := rpcStruct["speed"]
	if ok && speedValue != nil {
  	record.Speed, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "speed"), speedValue)
		if err != nil {
			return
		}
	}
  modelnameValue, ok := rpcStruct["modelname"]
	if ok && modelnameValue != nil {
  	record.Modelname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "modelname"), modelnameValue)
		if err != nil {
			return
		}
	}
  familyValue, ok := rpcStruct["family"]
	if ok && familyValue != nil {
  	record.Family, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "family"), familyValue)
		if err != nil {
			return
		}
	}
  modelValue, ok := rpcStruct["model"]
	if ok && modelValue != nil {
  	record.Model, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "model"), modelValue)
		if err != nil {
			return
		}
	}
  steppingValue, ok := rpcStruct["stepping"]
	if ok && steppingValue != nil {
  	record.Stepping, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "stepping"), steppingValue)
		if err != nil {
			return
		}
	}
  flagsValue, ok := rpcStruct["flags"]
	if ok && flagsValue != nil {
  	record.Flags, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "flags"), flagsValue)
		if err != nil {
			return
		}
	}
  featuresValue, ok := rpcStruct["features"]
	if ok && featuresValue != nil {
  	record.Features, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "features"), featuresValue)
		if err != nil {
			return
		}
	}
  utilisationValue, ok := rpcStruct["utilisation"]
	if ok && utilisationValue != nil {
  	record.Utilisation, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "utilisation"), utilisationValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostCPURefSetToGo(context string, input interface{}) (slice []HostCPURef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostCPURef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostCPURefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostCPURefToGo(context string, input interface{}) (ref HostCPURef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostCPURef(value)
	}
	return
}

func convertHostCPURefToXen(context string, ref HostCPURef) (string, error) {
	return string(ref), nil
}

func convertHostCrashdumpRecordToGo(context string, input interface{}) (record HostCrashdumpRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  timestampValue, ok := rpcStruct["timestamp"]
	if ok && timestampValue != nil {
  	record.Timestamp, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp"), timestampValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok && sizeValue != nil {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostCrashdumpRefSetToGo(context string, input interface{}) (slice []HostCrashdumpRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostCrashdumpRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostCrashdumpRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostCrashdumpRefToGo(context string, input interface{}) (ref HostCrashdumpRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostCrashdumpRef(value)
	}
	return
}

func convertHostCrashdumpRefToXen(context string, ref HostCrashdumpRef) (string, error) {
	return string(ref), nil
}

func convertHostMetricsRecordToGo(context string, input interface{}) (record HostMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  memoryTotalValue, ok := rpcStruct["memory_total"]
	if ok && memoryTotalValue != nil {
  	record.MemoryTotal, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_total"), memoryTotalValue)
		if err != nil {
			return
		}
	}
  memoryFreeValue, ok := rpcStruct["memory_free"]
	if ok && memoryFreeValue != nil {
  	record.MemoryFree, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_free"), memoryFreeValue)
		if err != nil {
			return
		}
	}
  liveValue, ok := rpcStruct["live"]
	if ok && liveValue != nil {
  	record.Live, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "live"), liveValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok && lastUpdatedValue != nil {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostMetricsRefSetToGo(context string, input interface{}) (slice []HostMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostMetricsRefToGo(context string, input interface{}) (ref HostMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostMetricsRef(value)
	}
	return
}

func convertHostMetricsRefToXen(context string, ref HostMetricsRef) (string, error) {
	return string(ref), nil
}

func convertHostPatchRecordToGo(context string, input interface{}) (record HostPatchRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok && hostValue != nil {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  appliedValue, ok := rpcStruct["applied"]
	if ok && appliedValue != nil {
  	record.Applied, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "applied"), appliedValue)
		if err != nil {
			return
		}
	}
  timestampAppliedValue, ok := rpcStruct["timestamp_applied"]
	if ok && timestampAppliedValue != nil {
  	record.TimestampApplied, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp_applied"), timestampAppliedValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok && sizeValue != nil {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  poolPatchValue, ok := rpcStruct["pool_patch"]
	if ok && poolPatchValue != nil {
  	record.PoolPatch, err = convertPoolPatchRefToGo(fmt.Sprintf("%s.%s", context, "pool_patch"), poolPatchValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostPatchRefSetToGo(context string, input interface{}) (slice []HostPatchRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostPatchRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostPatchRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostPatchRefToGo(context string, input interface{}) (ref HostPatchRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostPatchRef(value)
	}
	return
}

func convertHostPatchRefToXen(context string, ref HostPatchRef) (string, error) {
	return string(ref), nil
}

func convertIntSetToGo(context string, input interface{}) (slice []int, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]int, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertIntToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertIntToGo(context string, input interface{}) (value int, err error) {
	strValue, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
  	value, err = strconv.Atoi(strValue)
	}
	return
}

func convertIntToXen(context string, value int) (string, error) {
	return strconv.Itoa(value), nil
}

func convertMessageRecordToGo(context string, input interface{}) (record MessageRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameValue, ok := rpcStruct["name"]
	if ok && nameValue != nil {
  	record.Name, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name"), nameValue)
		if err != nil {
			return
		}
	}
  priorityValue, ok := rpcStruct["priority"]
	if ok && priorityValue != nil {
  	record.Priority, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "priority"), priorityValue)
		if err != nil {
			return
		}
	}
  clsValue, ok := rpcStruct["cls"]
	if ok && clsValue != nil {
  	record.Cls, err = convertEnumClsToGo(fmt.Sprintf("%s.%s", context, "cls"), clsValue)
		if err != nil {
			return
		}
	}
  objUUIDValue, ok := rpcStruct["obj_uuid"]
	if ok && objUUIDValue != nil {
  	record.ObjUUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "obj_uuid"), objUUIDValue)
		if err != nil {
			return
		}
	}
  timestampValue, ok := rpcStruct["timestamp"]
	if ok && timestampValue != nil {
  	record.Timestamp, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp"), timestampValue)
		if err != nil {
			return
		}
	}
  bodyValue, ok := rpcStruct["body"]
	if ok && bodyValue != nil {
  	record.Body, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "body"), bodyValue)
		if err != nil {
			return
		}
	}
	return
}

func convertMessageRefSetToGo(context string, input interface{}) (slice []MessageRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]MessageRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertMessageRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertMessageRefToGo(context string, input interface{}) (ref MessageRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = MessageRef(value)
	}
	return
}

func convertMessageRefToXen(context string, ref MessageRef) (string, error) {
	return string(ref), nil
}

func convertNetworkRecordToGo(context string, input interface{}) (record NetworkRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumNetworkOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumNetworkOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vifsValue, ok := rpcStruct["VIFs"]
	if ok && vifsValue != nil {
  	record.VIFs, err = convertVIFRefSetToGo(fmt.Sprintf("%s.%s", context, "VIFs"), vifsValue)
		if err != nil {
			return
		}
	}
  pifsValue, ok := rpcStruct["PIFs"]
	if ok && pifsValue != nil {
  	record.PIFs, err = convertPIFRefSetToGo(fmt.Sprintf("%s.%s", context, "PIFs"), pifsValue)
		if err != nil {
			return
		}
	}
  mtuValue, ok := rpcStruct["MTU"]
	if ok && mtuValue != nil {
  	record.MTU, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "MTU"), mtuValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  bridgeValue, ok := rpcStruct["bridge"]
	if ok && bridgeValue != nil {
  	record.Bridge, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "bridge"), bridgeValue)
		if err != nil {
			return
		}
	}
  managedValue, ok := rpcStruct["managed"]
	if ok && managedValue != nil {
  	record.Managed, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "managed"), managedValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok && blobsValue != nil {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok && tagsValue != nil {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  defaultLockingModeValue, ok := rpcStruct["default_locking_mode"]
	if ok && defaultLockingModeValue != nil {
  	record.DefaultLockingMode, err = convertEnumNetworkDefaultLockingModeToGo(fmt.Sprintf("%s.%s", context, "default_locking_mode"), defaultLockingModeValue)
		if err != nil {
			return
		}
	}
  assignedIpsValue, ok := rpcStruct["assigned_ips"]
	if ok && assignedIpsValue != nil {
  	record.AssignedIps, err = convertVIFRefToStringMapToGo(fmt.Sprintf("%s.%s", context, "assigned_ips"), assignedIpsValue)
		if err != nil {
			return
		}
	}
  purposeValue, ok := rpcStruct["purpose"]
	if ok && purposeValue != nil {
  	record.Purpose, err = convertEnumNetworkPurposeSetToGo(fmt.Sprintf("%s.%s", context, "purpose"), purposeValue)
		if err != nil {
			return
		}
	}
	return
}

func convertNetworkRecordToXen(context string, record NetworkRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumNetworkOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumNetworkOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["VIFs"], err = convertVIFRefSetToXen(fmt.Sprintf("%s.%s", context, "VIFs"), record.VIFs)
  if err != nil {
		return
	}
  rpcStruct["PIFs"], err = convertPIFRefSetToXen(fmt.Sprintf("%s.%s", context, "PIFs"), record.PIFs)
  if err != nil {
		return
	}
  rpcStruct["MTU"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "MTU"), record.MTU)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["bridge"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "bridge"), record.Bridge)
  if err != nil {
		return
	}
  rpcStruct["managed"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "managed"), record.Managed)
  if err != nil {
		return
	}
  rpcStruct["blobs"], err = convertStringToBlobRefMapToXen(fmt.Sprintf("%s.%s", context, "blobs"), record.Blobs)
  if err != nil {
		return
	}
  rpcStruct["tags"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "tags"), record.Tags)
  if err != nil {
		return
	}
  rpcStruct["default_locking_mode"], err = convertEnumNetworkDefaultLockingModeToXen(fmt.Sprintf("%s.%s", context, "default_locking_mode"), record.DefaultLockingMode)
  if err != nil {
		return
	}
  rpcStruct["assigned_ips"], err = convertVIFRefToStringMapToXen(fmt.Sprintf("%s.%s", context, "assigned_ips"), record.AssignedIps)
  if err != nil {
		return
	}
  rpcStruct["purpose"], err = convertEnumNetworkPurposeSetToXen(fmt.Sprintf("%s.%s", context, "purpose"), record.Purpose)
  if err != nil {
		return
	}
	return
}

func convertNetworkRefSetToGo(context string, input interface{}) (slice []NetworkRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]NetworkRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertNetworkRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertNetworkRefToGo(context string, input interface{}) (ref NetworkRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = NetworkRef(value)
	}
	return
}

func convertNetworkRefToXen(context string, ref NetworkRef) (string, error) {
	return string(ref), nil
}

func convertPoolRecordToGo(context string, input interface{}) (record PoolRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  masterValue, ok := rpcStruct["master"]
	if ok && masterValue != nil {
  	record.Master, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "master"), masterValue)
		if err != nil {
			return
		}
	}
  defaultSRValue, ok := rpcStruct["default_SR"]
	if ok && defaultSRValue != nil {
  	record.DefaultSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "default_SR"), defaultSRValue)
		if err != nil {
			return
		}
	}
  suspendImageSRValue, ok := rpcStruct["suspend_image_SR"]
	if ok && suspendImageSRValue != nil {
  	record.SuspendImageSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "suspend_image_SR"), suspendImageSRValue)
		if err != nil {
			return
		}
	}
  crashDumpSRValue, ok := rpcStruct["crash_dump_SR"]
	if ok && crashDumpSRValue != nil {
  	record.CrashDumpSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "crash_dump_SR"), crashDumpSRValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  haEnabledValue, ok := rpcStruct["ha_enabled"]
	if ok && haEnabledValue != nil {
  	record.HaEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_enabled"), haEnabledValue)
		if err != nil {
			return
		}
	}
  haConfigurationValue, ok := rpcStruct["ha_configuration"]
	if ok && haConfigurationValue != nil {
  	record.HaConfiguration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "ha_configuration"), haConfigurationValue)
		if err != nil {
			return
		}
	}
  haStatefilesValue, ok := rpcStruct["ha_statefiles"]
	if ok && haStatefilesValue != nil {
  	record.HaStatefiles, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ha_statefiles"), haStatefilesValue)
		if err != nil {
			return
		}
	}
  haHostFailuresToTolerateValue, ok := rpcStruct["ha_host_failures_to_tolerate"]
	if ok && haHostFailuresToTolerateValue != nil {
  	record.HaHostFailuresToTolerate, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "ha_host_failures_to_tolerate"), haHostFailuresToTolerateValue)
		if err != nil {
			return
		}
	}
  haPlanExistsForValue, ok := rpcStruct["ha_plan_exists_for"]
	if ok && haPlanExistsForValue != nil {
  	record.HaPlanExistsFor, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "ha_plan_exists_for"), haPlanExistsForValue)
		if err != nil {
			return
		}
	}
  haAllowOvercommitValue, ok := rpcStruct["ha_allow_overcommit"]
	if ok && haAllowOvercommitValue != nil {
  	record.HaAllowOvercommit, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_allow_overcommit"), haAllowOvercommitValue)
		if err != nil {
			return
		}
	}
  haOvercommittedValue, ok := rpcStruct["ha_overcommitted"]
	if ok && haOvercommittedValue != nil {
  	record.HaOvercommitted, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_overcommitted"), haOvercommittedValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok && blobsValue != nil {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok && tagsValue != nil {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  guiConfigValue, ok := rpcStruct["gui_config"]
	if ok && guiConfigValue != nil {
  	record.GuiConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "gui_config"), guiConfigValue)
		if err != nil {
			return
		}
	}
  healthCheckConfigValue, ok := rpcStruct["health_check_config"]
	if ok && healthCheckConfigValue != nil {
  	record.HealthCheckConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "health_check_config"), healthCheckConfigValue)
		if err != nil {
			return
		}
	}
  wlbURLValue, ok := rpcStruct["wlb_url"]
	if ok && wlbURLValue != nil {
  	record.WlbURL, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "wlb_url"), wlbURLValue)
		if err != nil {
			return
		}
	}
  wlbUsernameValue, ok := rpcStruct["wlb_username"]
	if ok && wlbUsernameValue != nil {
  	record.WlbUsername, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "wlb_username"), wlbUsernameValue)
		if err != nil {
			return
		}
	}
  wlbEnabledValue, ok := rpcStruct["wlb_enabled"]
	if ok && wlbEnabledValue != nil {
  	record.WlbEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "wlb_enabled"), wlbEnabledValue)
		if err != nil {
			return
		}
	}
  wlbVerifyCertValue, ok := rpcStruct["wlb_verify_cert"]
	if ok && wlbVerifyCertValue != nil {
  	record.WlbVerifyCert, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "wlb_verify_cert"), wlbVerifyCertValue)
		if err != nil {
			return
		}
	}
  redoLogEnabledValue, ok := rpcStruct["redo_log_enabled"]
	if ok && redoLogEnabledValue != nil {
  	record.RedoLogEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "redo_log_enabled"), redoLogEnabledValue)
		if err != nil {
			return
		}
	}
  redoLogVdiValue, ok := rpcStruct["redo_log_vdi"]
	if ok && redoLogVdiValue != nil {
  	record.RedoLogVdi, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "redo_log_vdi"), redoLogVdiValue)
		if err != nil {
			return
		}
	}
  vswitchControllerValue, ok := rpcStruct["vswitch_controller"]
	if ok && vswitchControllerValue != nil {
  	record.VswitchController, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vswitch_controller"), vswitchControllerValue)
		if err != nil {
			return
		}
	}
  restrictionsValue, ok := rpcStruct["restrictions"]
	if ok && restrictionsValue != nil {
  	record.Restrictions, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "restrictions"), restrictionsValue)
		if err != nil {
			return
		}
	}
  metadataVDIsValue, ok := rpcStruct["metadata_VDIs"]
	if ok && metadataVDIsValue != nil {
  	record.MetadataVDIs, err = convertVDIRefSetToGo(fmt.Sprintf("%s.%s", context, "metadata_VDIs"), metadataVDIsValue)
		if err != nil {
			return
		}
	}
  haClusterStackValue, ok := rpcStruct["ha_cluster_stack"]
	if ok && haClusterStackValue != nil {
  	record.HaClusterStack, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ha_cluster_stack"), haClusterStackValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumPoolAllowedOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumPoolAllowedOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  guestAgentConfigValue, ok := rpcStruct["guest_agent_config"]
	if ok && guestAgentConfigValue != nil {
  	record.GuestAgentConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "guest_agent_config"), guestAgentConfigValue)
		if err != nil {
			return
		}
	}
  cpuInfoValue, ok := rpcStruct["cpu_info"]
	if ok && cpuInfoValue != nil {
  	record.CPUInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "cpu_info"), cpuInfoValue)
		if err != nil {
			return
		}
	}
  policyNoVendorDeviceValue, ok := rpcStruct["policy_no_vendor_device"]
	if ok && policyNoVendorDeviceValue != nil {
  	record.PolicyNoVendorDevice, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "policy_no_vendor_device"), policyNoVendorDeviceValue)
		if err != nil {
			return
		}
	}
  livePatchingDisabledValue, ok := rpcStruct["live_patching_disabled"]
	if ok && livePatchingDisabledValue != nil {
  	record.LivePatchingDisabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "live_patching_disabled"), livePatchingDisabledValue)
		if err != nil {
			return
		}
	}
  igmpSnoopingEnabledValue, ok := rpcStruct["igmp_snooping_enabled"]
	if ok && igmpSnoopingEnabledValue != nil {
  	record.IgmpSnoopingEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "igmp_snooping_enabled"), igmpSnoopingEnabledValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPoolRefSetToGo(context string, input interface{}) (slice []PoolRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPoolRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPoolRefToGo(context string, input interface{}) (ref PoolRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PoolRef(value)
	}
	return
}

func convertPoolRefToXen(context string, ref PoolRef) (string, error) {
	return string(ref), nil
}

func convertPoolPatchRecordToGo(context string, input interface{}) (record PoolPatchRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok && sizeValue != nil {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  poolAppliedValue, ok := rpcStruct["pool_applied"]
	if ok && poolAppliedValue != nil {
  	record.PoolApplied, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "pool_applied"), poolAppliedValue)
		if err != nil {
			return
		}
	}
  hostPatchesValue, ok := rpcStruct["host_patches"]
	if ok && hostPatchesValue != nil {
  	record.HostPatches, err = convertHostPatchRefSetToGo(fmt.Sprintf("%s.%s", context, "host_patches"), hostPatchesValue)
		if err != nil {
			return
		}
	}
  afterApplyGuidanceValue, ok := rpcStruct["after_apply_guidance"]
	if ok && afterApplyGuidanceValue != nil {
  	record.AfterApplyGuidance, err = convertEnumAfterApplyGuidanceSetToGo(fmt.Sprintf("%s.%s", context, "after_apply_guidance"), afterApplyGuidanceValue)
		if err != nil {
			return
		}
	}
  poolUpdateValue, ok := rpcStruct["pool_update"]
	if ok && poolUpdateValue != nil {
  	record.PoolUpdate, err = convertPoolUpdateRefToGo(fmt.Sprintf("%s.%s", context, "pool_update"), poolUpdateValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPoolPatchRefSetToGo(context string, input interface{}) (slice []PoolPatchRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolPatchRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPoolPatchRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPoolPatchRefToGo(context string, input interface{}) (ref PoolPatchRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PoolPatchRef(value)
	}
	return
}

func convertPoolPatchRefToXen(context string, ref PoolPatchRef) (string, error) {
	return string(ref), nil
}

func convertPoolUpdateRecordToGo(context string, input interface{}) (record PoolUpdateRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok && versionValue != nil {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  installationSizeValue, ok := rpcStruct["installation_size"]
	if ok && installationSizeValue != nil {
  	record.InstallationSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "installation_size"), installationSizeValue)
		if err != nil {
			return
		}
	}
  keyValue, ok := rpcStruct["key"]
	if ok && keyValue != nil {
  	record.Key, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "key"), keyValue)
		if err != nil {
			return
		}
	}
  afterApplyGuidanceValue, ok := rpcStruct["after_apply_guidance"]
	if ok && afterApplyGuidanceValue != nil {
  	record.AfterApplyGuidance, err = convertEnumUpdateAfterApplyGuidanceSetToGo(fmt.Sprintf("%s.%s", context, "after_apply_guidance"), afterApplyGuidanceValue)
		if err != nil {
			return
		}
	}
  vdiValue, ok := rpcStruct["vdi"]
	if ok && vdiValue != nil {
  	record.Vdi, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "vdi"), vdiValue)
		if err != nil {
			return
		}
	}
  hostsValue, ok := rpcStruct["hosts"]
	if ok && hostsValue != nil {
  	record.Hosts, err = convertHostRefSetToGo(fmt.Sprintf("%s.%s", context, "hosts"), hostsValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  enforceHomogeneityValue, ok := rpcStruct["enforce_homogeneity"]
	if ok && enforceHomogeneityValue != nil {
  	record.EnforceHomogeneity, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enforce_homogeneity"), enforceHomogeneityValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPoolUpdateRefSetToGo(context string, input interface{}) (slice []PoolUpdateRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolUpdateRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPoolUpdateRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPoolUpdateRefToGo(context string, input interface{}) (ref PoolUpdateRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PoolUpdateRef(value)
	}
	return
}

func convertPoolUpdateRefToXen(context string, ref PoolUpdateRef) (string, error) {
	return string(ref), nil
}

func convertRoleRecordToGo(context string, input interface{}) (record RoleRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  subrolesValue, ok := rpcStruct["subroles"]
	if ok && subrolesValue != nil {
  	record.Subroles, err = convertRoleRefSetToGo(fmt.Sprintf("%s.%s", context, "subroles"), subrolesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertRoleRefSetToGo(context string, input interface{}) (slice []RoleRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]RoleRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertRoleRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertRoleRefSetToXen(context string, slice []RoleRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertRoleRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertRoleRefToGo(context string, input interface{}) (ref RoleRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = RoleRef(value)
	}
	return
}

func convertRoleRefToXen(context string, ref RoleRef) (string, error) {
	return string(ref), nil
}

func convertSecretRecordToGo(context string, input interface{}) (record SecretRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  valueValue, ok := rpcStruct["value"]
	if ok && valueValue != nil {
  	record.Value, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "value"), valueValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSecretRecordToXen(context string, record SecretRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["value"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "value"), record.Value)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertSecretRefSetToGo(context string, input interface{}) (slice []SecretRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SecretRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSecretRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSecretRefToGo(context string, input interface{}) (ref SecretRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SecretRef(value)
	}
	return
}

func convertSecretRefToXen(context string, ref SecretRef) (string, error) {
	return string(ref), nil
}

func convertSessionRecordToGo(context string, input interface{}) (record SessionRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  thisHostValue, ok := rpcStruct["this_host"]
	if ok && thisHostValue != nil {
  	record.ThisHost, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "this_host"), thisHostValue)
		if err != nil {
			return
		}
	}
  thisUserValue, ok := rpcStruct["this_user"]
	if ok && thisUserValue != nil {
  	record.ThisUser, err = convertUserRefToGo(fmt.Sprintf("%s.%s", context, "this_user"), thisUserValue)
		if err != nil {
			return
		}
	}
  lastActiveValue, ok := rpcStruct["last_active"]
	if ok && lastActiveValue != nil {
  	record.LastActive, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_active"), lastActiveValue)
		if err != nil {
			return
		}
	}
  poolValue, ok := rpcStruct["pool"]
	if ok && poolValue != nil {
  	record.Pool, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "pool"), poolValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  isLocalSuperuserValue, ok := rpcStruct["is_local_superuser"]
	if ok && isLocalSuperuserValue != nil {
  	record.IsLocalSuperuser, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_local_superuser"), isLocalSuperuserValue)
		if err != nil {
			return
		}
	}
  subjectValue, ok := rpcStruct["subject"]
	if ok && subjectValue != nil {
  	record.Subject, err = convertSubjectRefToGo(fmt.Sprintf("%s.%s", context, "subject"), subjectValue)
		if err != nil {
			return
		}
	}
  validationTimeValue, ok := rpcStruct["validation_time"]
	if ok && validationTimeValue != nil {
  	record.ValidationTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "validation_time"), validationTimeValue)
		if err != nil {
			return
		}
	}
  authUserSidValue, ok := rpcStruct["auth_user_sid"]
	if ok && authUserSidValue != nil {
  	record.AuthUserSid, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "auth_user_sid"), authUserSidValue)
		if err != nil {
			return
		}
	}
  authUserNameValue, ok := rpcStruct["auth_user_name"]
	if ok && authUserNameValue != nil {
  	record.AuthUserName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "auth_user_name"), authUserNameValue)
		if err != nil {
			return
		}
	}
  rbacPermissionsValue, ok := rpcStruct["rbac_permissions"]
	if ok && rbacPermissionsValue != nil {
  	record.RbacPermissions, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "rbac_permissions"), rbacPermissionsValue)
		if err != nil {
			return
		}
	}
  tasksValue, ok := rpcStruct["tasks"]
	if ok && tasksValue != nil {
  	record.Tasks, err = convertTaskRefSetToGo(fmt.Sprintf("%s.%s", context, "tasks"), tasksValue)
		if err != nil {
			return
		}
	}
  parentValue, ok := rpcStruct["parent"]
	if ok && parentValue != nil {
  	record.Parent, err = convertSessionRefToGo(fmt.Sprintf("%s.%s", context, "parent"), parentValue)
		if err != nil {
			return
		}
	}
  originatorValue, ok := rpcStruct["originator"]
	if ok && originatorValue != nil {
  	record.Originator, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "originator"), originatorValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSessionRefToGo(context string, input interface{}) (ref SessionRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SessionRef(value)
	}
	return
}

func convertSessionRefToXen(context string, ref SessionRef) (string, error) {
	return string(ref), nil
}

func convertStringSetToGo(context string, input interface{}) (slice []string, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]string, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertStringToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertStringSetToXen(context string, slice []string) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertStringToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertStringToGo(context string, input interface{}) (value string, err error) {
	if input == nil {
		return
	}
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	}
	return
}

func convertStringToXen(context string, value string) (string, error) {
	return value, nil
}

func convertSubjectRecordToGo(context string, input interface{}) (record SubjectRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  subjectIdentifierValue, ok := rpcStruct["subject_identifier"]
	if ok && subjectIdentifierValue != nil {
  	record.SubjectIdentifier, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subject_identifier"), subjectIdentifierValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  rolesValue, ok := rpcStruct["roles"]
	if ok && rolesValue != nil {
  	record.Roles, err = convertRoleRefSetToGo(fmt.Sprintf("%s.%s", context, "roles"), rolesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSubjectRecordToXen(context string, record SubjectRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["subject_identifier"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "subject_identifier"), record.SubjectIdentifier)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["roles"], err = convertRoleRefSetToXen(fmt.Sprintf("%s.%s", context, "roles"), record.Roles)
  if err != nil {
		return
	}
	return
}

func convertSubjectRefSetToGo(context string, input interface{}) (slice []SubjectRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SubjectRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSubjectRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSubjectRefToGo(context string, input interface{}) (ref SubjectRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SubjectRef(value)
	}
	return
}

func convertSubjectRefToXen(context string, ref SubjectRef) (string, error) {
	return string(ref), nil
}

func convertTaskRecordToGo(context string, input interface{}) (record TaskRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok && nameLabelValue != nil {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok && nameDescriptionValue != nil {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok && allowedOperationsValue != nil {
  	record.AllowedOperations, err = convertEnumTaskAllowedOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok && currentOperationsValue != nil {
  	record.CurrentOperations, err = convertStringToEnumTaskAllowedOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  createdValue, ok := rpcStruct["created"]
	if ok && createdValue != nil {
  	record.Created, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "created"), createdValue)
		if err != nil {
			return
		}
	}
  finishedValue, ok := rpcStruct["finished"]
	if ok && finishedValue != nil {
  	record.Finished, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "finished"), finishedValue)
		if err != nil {
			return
		}
	}
  statusValue, ok := rpcStruct["status"]
	if ok && statusValue != nil {
  	record.Status, err = convertEnumTaskStatusTypeToGo(fmt.Sprintf("%s.%s", context, "status"), statusValue)
		if err != nil {
			return
		}
	}
  residentOnValue, ok := rpcStruct["resident_on"]
	if ok && residentOnValue != nil {
  	record.ResidentOn, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "resident_on"), residentOnValue)
		if err != nil {
			return
		}
	}
  progressValue, ok := rpcStruct["progress"]
	if ok && progressValue != nil {
  	record.Progress, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "progress"), progressValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok && atypeValue != nil {
  	record.Type, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  resultValue, ok := rpcStruct["result"]
	if ok && resultValue != nil {
  	record.Result, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "result"), resultValue)
		if err != nil {
			return
		}
	}
  errorInfoValue, ok := rpcStruct["error_info"]
	if ok && errorInfoValue != nil {
  	record.ErrorInfo, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "error_info"), errorInfoValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  subtaskOfValue, ok := rpcStruct["subtask_of"]
	if ok && subtaskOfValue != nil {
  	record.SubtaskOf, err = convertTaskRefToGo(fmt.Sprintf("%s.%s", context, "subtask_of"), subtaskOfValue)
		if err != nil {
			return
		}
	}
  subtasksValue, ok := rpcStruct["subtasks"]
	if ok && subtasksValue != nil {
  	record.Subtasks, err = convertTaskRefSetToGo(fmt.Sprintf("%s.%s", context, "subtasks"), subtasksValue)
		if err != nil {
			return
		}
	}
  backtraceValue, ok := rpcStruct["backtrace"]
	if ok && backtraceValue != nil {
  	record.Backtrace, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "backtrace"), backtraceValue)
		if err != nil {
			return
		}
	}
	return
}

func convertTaskRefSetToGo(context string, input interface{}) (slice []TaskRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]TaskRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertTaskRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertTaskRefToGo(context string, input interface{}) (ref TaskRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = TaskRef(value)
	}
	return
}

func convertTaskRefToXen(context string, ref TaskRef) (string, error) {
	return string(ref), nil
}

func convertTunnelRecordToGo(context string, input interface{}) (record TunnelRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  accessPIFValue, ok := rpcStruct["access_PIF"]
	if ok && accessPIFValue != nil {
  	record.AccessPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "access_PIF"), accessPIFValue)
		if err != nil {
			return
		}
	}
  transportPIFValue, ok := rpcStruct["transport_PIF"]
	if ok && transportPIFValue != nil {
  	record.TransportPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "transport_PIF"), transportPIFValue)
		if err != nil {
			return
		}
	}
  statusValue, ok := rpcStruct["status"]
	if ok && statusValue != nil {
  	record.Status, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "status"), statusValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertTunnelRefSetToGo(context string, input interface{}) (slice []TunnelRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]TunnelRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertTunnelRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertTunnelRefToGo(context string, input interface{}) (ref TunnelRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = TunnelRef(value)
	}
	return
}

func convertTunnelRefToXen(context string, ref TunnelRef) (string, error) {
	return string(ref), nil
}

func convertUserRecordToGo(context string, input interface{}) (record UserRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok && uuidValue != nil {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  shortNameValue, ok := rpcStruct["short_name"]
	if ok && shortNameValue != nil {
  	record.ShortName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "short_name"), shortNameValue)
		if err != nil {
			return
		}
	}
  fullnameValue, ok := rpcStruct["fullname"]
	if ok && fullnameValue != nil {
  	record.Fullname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "fullname"), fullnameValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok && otherConfigValue != nil {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertUserRecordToXen(context string, record UserRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = xmlrpc.Struct{}
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["short_name"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "short_name"), record.ShortName)
  if err != nil {
		return
	}
  rpcStruct["fullname"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "fullname"), record.Fullname)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertUserRefToGo(context string, input interface{}) (ref UserRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = UserRef(value)
	}
	return
}

func convertUserRefToXen(context string, ref UserRef) (string, error) {
	return string(ref), nil
}

func convertVdiNbdServerInfoRecordSetToGo(context string, input interface{}) (slice []VdiNbdServerInfoRecord, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VdiNbdServerInfoRecord, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVdiNbdServerInfoRecordToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVdiNbdServerInfoRecordToGo(context string, input interface{}) (record VdiNbdServerInfoRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  exportnameValue, ok := rpcStruct["exportname"]
	if ok && exportnameValue != nil {
  	record.Exportname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "exportname"), exportnameValue)
		if err != nil {
			return
		}
	}
  addressValue, ok := rpcStruct["address"]
	if ok && addressValue != nil {
  	record.Address, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "address"), addressValue)
		if err != nil {
			return
		}
	}
  portValue, ok := rpcStruct["port"]
	if ok && portValue != nil {
  	record.Port, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "port"), portValue)
		if err != nil {
			return
		}
	}
  certValue, ok := rpcStruct["cert"]
	if ok && certValue != nil {
  	record.Cert, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "cert"), certValue)
		if err != nil {
			return
		}
	}
  subjectValue, ok := rpcStruct["subject"]
	if ok && subjectValue != nil {
  	record.Subject, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subject"), subjectValue)
		if err != nil {
			return
		}
	}
	return
}
