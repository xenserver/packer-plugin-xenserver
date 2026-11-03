// Copyright (c) 2026. Citrix Systems, Inc.

package clone

import "testing"

func TestCloneConfig_BootsTheGuest_FalseForCommunicatorNoneWithNothingStaged(t *testing.T) {
	var c CloneConfig
	c.Comm.Type = "none"

	if c.bootsTheGuest() {
		t.Error("nothing logs in, nothing runs and nothing halts the guest: it must not be started")
	}
}

func TestCloneConfig_BootsTheGuest_TrueForEveryRealCommunicator(t *testing.T) {
	for _, commType := range []string{"ssh", "winrm"} {
		var c CloneConfig
		c.Comm.Type = commType

		if !c.bootsTheGuest() {
			t.Errorf("communicator %q logs in to the guest, so the guest has to be running", commType)
		}
	}
}

func TestCloneConfig_BootsTheGuest_TrueWhenCDFilesIsStaged(t *testing.T) {
	var c CloneConfig
	c.Comm.Type = "none"
	c.CDFiles = []string{"Autounattend.xml"}

	if !c.bootsTheGuest() {
		t.Error("cd_files gives the guest something to run, so it has to be started")
	}
}

func TestCloneConfig_BootsTheGuest_TrueWhenBootCommandIsSet(t *testing.T) {
	var c CloneConfig
	c.Comm.Type = "none"
	c.BootCommand = []string{"<enter>"}

	if !c.bootsTheGuest() {
		t.Error("boot_command types at the guest, so it has to be started")
	}
}

func TestCloneConfig_BootsTheGuest_FalseWhenOnlyHTTPDirIsSet(t *testing.T) {
	var c CloneConfig
	c.Comm.Type = "none"
	c.HTTPDir = "/tmp/does-not-matter"

	if c.bootsTheGuest() {
		t.Error("http_directory alone tells the guest nothing, so it is not a reason to boot")
	}
}
