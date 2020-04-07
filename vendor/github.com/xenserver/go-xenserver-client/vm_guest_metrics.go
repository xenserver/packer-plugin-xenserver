package client

import (
	"github.com/nilshell/xmlrpc"
)

type VM_Guest_Metrics XenAPIObject

func (self *VM_Guest_Metrics) GetNetworks() (ipMap map[string]string, err error) {
	ipMap = make(map[string]string, 0)
	result := APIResult{}
	err = self.Client.APICall(&result, "VM_guest_metrics.get_networks", self.Ref)
	if err != nil {
		return ipMap, err
	}
	for k, v := range result.Value.(xmlrpc.Struct) {
		ipMap[k] = v.(string)
	}
	return ipMap, nil
}
