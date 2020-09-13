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

type Client struct {
	rpc *xmlrpc.Client
	Session SessionClass
	Auth AuthClass
	Subject SubjectClass
	Role RoleClass
	Task TaskClass
	Event EventClass
	Pool PoolClass
	PoolPatch PoolPatchClass
	PoolUpdate PoolUpdateClass
	VM VMClass
	VMMetrics VMMetricsClass
	VMGuestMetrics VMGuestMetricsClass
	VMPP VMPPClass
	VMSS VMSSClass
	VMAppliance VMApplianceClass
	DRTask DRTaskClass
	Host HostClass
	HostCrashdump HostCrashdumpClass
	HostPatch HostPatchClass
	HostMetrics HostMetricsClass
	HostCPU HostCPUClass
	Network NetworkClass
	VIF VIFClass
	VIFMetrics VIFMetricsClass
	PIF PIFClass
	PIFMetrics PIFMetricsClass
	Bond BondClass
	VLAN VLANClass
	SM SMClass
	SR SRClass
	LVHD LVHDClass
	VDI VDIClass
	VBD VBDClass
	VBDMetrics VBDMetricsClass
	PBD PBDClass
	Crashdump CrashdumpClass
	VTPM VTPMClass
	Console ConsoleClass
	User UserClass
	DataSource DataSourceClass
	Blob BlobClass
	Message MessageClass
	Secret SecretClass
	Tunnel TunnelClass
	PCI PCIClass
	PGPU PGPUClass
	GPUGroup GPUGroupClass
	VGPU VGPUClass
	VGPUType VGPUTypeClass
	PVSSite PVSSiteClass
	PVSServer PVSServerClass
	PVSProxy PVSProxyClass
	PVSCacheStorage PVSCacheStorageClass
	Feature FeatureClass
	SDNController SDNControllerClass
	VdiNbdServerInfo VdiNbdServerInfoClass
	PUSB PUSBClass
	USBGroup USBGroupClass
	VUSB VUSBClass
}

func prepClient(rpc *xmlrpc.Client) *Client {
	var client Client
	client.rpc = rpc
	client.Session = SessionClass{&client}
	client.Auth = AuthClass{&client}
	client.Subject = SubjectClass{&client}
	client.Role = RoleClass{&client}
	client.Task = TaskClass{&client}
	client.Event = EventClass{&client}
	client.Pool = PoolClass{&client}
	client.PoolPatch = PoolPatchClass{&client}
	client.PoolUpdate = PoolUpdateClass{&client}
	client.VM = VMClass{&client}
	client.VMMetrics = VMMetricsClass{&client}
	client.VMGuestMetrics = VMGuestMetricsClass{&client}
	client.VMPP = VMPPClass{&client}
	client.VMSS = VMSSClass{&client}
	client.VMAppliance = VMApplianceClass{&client}
	client.DRTask = DRTaskClass{&client}
	client.Host = HostClass{&client}
	client.HostCrashdump = HostCrashdumpClass{&client}
	client.HostPatch = HostPatchClass{&client}
	client.HostMetrics = HostMetricsClass{&client}
	client.HostCPU = HostCPUClass{&client}
	client.Network = NetworkClass{&client}
	client.VIF = VIFClass{&client}
	client.VIFMetrics = VIFMetricsClass{&client}
	client.PIF = PIFClass{&client}
	client.PIFMetrics = PIFMetricsClass{&client}
	client.Bond = BondClass{&client}
	client.VLAN = VLANClass{&client}
	client.SM = SMClass{&client}
	client.SR = SRClass{&client}
	client.LVHD = LVHDClass{&client}
	client.VDI = VDIClass{&client}
	client.VBD = VBDClass{&client}
	client.VBDMetrics = VBDMetricsClass{&client}
	client.PBD = PBDClass{&client}
	client.Crashdump = CrashdumpClass{&client}
	client.VTPM = VTPMClass{&client}
	client.Console = ConsoleClass{&client}
	client.User = UserClass{&client}
	client.DataSource = DataSourceClass{&client}
	client.Blob = BlobClass{&client}
	client.Message = MessageClass{&client}
	client.Secret = SecretClass{&client}
	client.Tunnel = TunnelClass{&client}
	client.PCI = PCIClass{&client}
	client.PGPU = PGPUClass{&client}
	client.GPUGroup = GPUGroupClass{&client}
	client.VGPU = VGPUClass{&client}
	client.VGPUType = VGPUTypeClass{&client}
	client.PVSSite = PVSSiteClass{&client}
	client.PVSServer = PVSServerClass{&client}
	client.PVSProxy = PVSProxyClass{&client}
	client.PVSCacheStorage = PVSCacheStorageClass{&client}
	client.Feature = FeatureClass{&client}
	client.SDNController = SDNControllerClass{&client}
	client.VdiNbdServerInfo = VdiNbdServerInfoClass{&client}
	client.PUSB = PUSBClass{&client}
	client.USBGroup = USBGroupClass{&client}
	client.VUSB = VUSBClass{&client}
	return &client
}
