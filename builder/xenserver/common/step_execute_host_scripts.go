package common

import (
	"fmt"
	"github.com/dongyuzheng/easyssh"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"path/filepath"
)

type StepExecuteHostScripts struct {
	ScriptType           string
	LocalScripts         []string
	HostScriptFolderPath string
}

func (self *StepExecuteHostScripts) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("commonconfig").(CommonConfig)

	ui.Say(fmt.Sprintf("Step: Execute %s XenServer host scripts", self.ScriptType))

	if len(self.LocalScripts) == 0 {
		ui.Say("No scripts to execute")
		return multistep.ActionContinue
	}

	ssh := &easyssh.MakeConfig{
		Server:        config.HostIp,
		Port:          fmt.Sprintf("%d", config.XenSSHPort),
		User:          config.Username,
		Password:      config.Password,
		OutputHandler: func(s string) { ui.Say(s) },
	}

	uuid := state.Get("instance_uuid").(string)
	self.HostScriptFolderPath = fmt.Sprintf("/tmp/packer_scripts_%s_%s", uuid, self.ScriptType)

	ui.Say(fmt.Sprintf("Creating script folder on XenServer host: %s", self.HostScriptFolderPath))

	cmd := fmt.Sprintf("mkdir -p '%s'", self.HostScriptFolderPath)
	_, sessionErr, err := ssh.Exec(cmd)
	if err != nil {
		ui.Error(fmt.Sprintf("Error executing SSH command: %s", err.Error()))
		return multistep.ActionHalt
	} else if sessionErr != nil {
		ui.Error(fmt.Sprintf("Failed to create script folder: %s\n%s", cmd, sessionErr.Error()))
		return multistep.ActionHalt
	}

	for _, script := range self.LocalScripts {
		scriptBasename := filepath.Base(script)
		scriptPath := fmt.Sprintf("%s/%s", self.HostScriptFolderPath, scriptBasename)

		err = ssh.Upload(script, scriptPath)
		if err != nil {
			ui.Error(fmt.Sprintf("Error uploading script: %s\n%s", script, err.Error()))
			return multistep.ActionHalt
		}

		ui.Say(fmt.Sprintf("Executing %s ...", scriptBasename))
		cmd = fmt.Sprintf("chmod +x '%s' && '%s' '%s'", scriptPath, scriptPath, uuid)
		_, sessionErr, err := ssh.Exec(cmd)
		if err != nil {
			ui.Error(fmt.Sprintf("Error executing SSH command: %s", err.Error()))
			return multistep.ActionHalt
		} else if sessionErr != nil {
			ui.Error(fmt.Sprintf("Failed to execute script: %s\n%s", cmd, sessionErr.Error()))
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (self *StepExecuteHostScripts) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	ui.Say(fmt.Sprintf("Cleaning up %s host scripts", self.ScriptType))
	if self.HostScriptFolderPath == "" {
		ui.Say("No scripts to cleanup")
		return
	}
	config := state.Get("commonconfig").(CommonConfig)
	ssh := &easyssh.MakeConfig{
		Server:   config.HostIp,
		Port:     fmt.Sprintf("%d", config.XenSSHPort),
		User:     config.Username,
		Password: config.Password,
	}
	ui.Say(fmt.Sprintf("Deleting script folder: %s", self.HostScriptFolderPath))
	cmd := fmt.Sprintf("rm -rf '%s'", self.HostScriptFolderPath)
	_, sessionErr, err := ssh.Exec(cmd)
	if err != nil {
		ui.Error(fmt.Sprintf("Error executing SSH command: %s", err.Error()))
	} else if sessionErr != nil {
		ui.Error(fmt.Sprintf("Failed to delete scripts folder: %s\n%s", cmd, sessionErr.Error()))
	}
}
