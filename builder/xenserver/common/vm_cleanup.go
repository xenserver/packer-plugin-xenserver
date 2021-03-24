package common

import (
	"fmt"
	"log"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

type VmCleanup struct{}

func (self *VmCleanup) Cleanup(state multistep.StateBag) {
	config := state.Get("commonconfig").(CommonConfig)
	c := state.Get("client").(*Connection)

	if config.ShouldKeepVM(state) {
		return
	}

	uuid := state.Get("instance_uuid").(string)
	instance, err := c.client.VM.GetByUUID(c.session, uuid)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to get VM from UUID '%s': %s", uuid, err.Error()))
		return
	}

	err = c.client.VM.HardShutdown(c.session, instance)
	if err != nil {
		log.Printf(fmt.Sprintf("Unable to force shutdown VM '%s': %s", uuid, err.Error()))
	}
}
