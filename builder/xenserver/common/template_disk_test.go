// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"xenapi"
)

const provisionOneDisk = `<provision><disk device="0" size="25769803776" sr="" bootable="true" type="system"/></provision>`

func TestProvisionedDiskSize(t *testing.T) {
	for _, tc := range []struct {
		name    string
		doc     string
		want    int
		wantErr bool
	}{
		{"no document", "", 0, false},
		{"whitespace only", "  \n", 0, false},
		{"one disk on device 0", provisionOneDisk, 24 * BytesPerGB, false},
		{
			"device 0 among several",
			`<provision><disk device="1" size="1073741824" sr="" bootable="false" type="user"/>` +
				`<disk device="0" size="25769803776" sr="" bootable="true" type="system"/></provision>`,
			24 * BytesPerGB,
			false,
		},
		{
			"no device 0",
			`<provision><disk device="1" size="1073741824" sr="" bootable="false" type="user"/></provision>`,
			0,
			false,
		},
		{
			"unreadable size",
			`<provision><disk device="0" size="24GB" sr="" bootable="true" type="system"/></provision>`,
			0,
			true,
		},
		{
			"negative size",
			`<provision><disk device="0" size="-1" sr="" bootable="true" type="system"/></provision>`,
			0,
			true,
		},
		{"not a provision document", `<disks><disk device="0" size="1"/></disks>`, 0, true},
		{"malformed", `<provision><disk device="0"`, 0, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := provisionedDiskSize(tc.doc)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("provisionedDiskSize(%q) = %d, want an error", tc.doc, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if got != tc.want {
				t.Errorf("provisionedDiskSize() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestTemplateDiskSize_ReadsTheProvisioningXML(t *testing.T) {
	size, err := templateDiskSize(&fakeDisks{}, TemplateSource("CentOS 9"),
		map[string]string{provisionDisksKey: provisionOneDisk})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if size != 24*BytesPerGB {
		t.Errorf("templateDiskSize() = %d, want %d", size, 24*BytesPerGB)
	}
}

func TestTemplateDiskSize_FallsBackToTheTemplatesOwnDisk(t *testing.T) {
	size, err := templateDiskSize(oneDisk(40), TemplateSource("packer-base"), nil)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if size != 40*BytesPerGB {
		t.Errorf("templateDiskSize() = %d, want %d", size, 40*BytesPerGB)
	}
}

func TestTemplateDiskSize_FallsBackWhenTheXMLNamesNoPrimaryDisk(t *testing.T) {
	otherConfig := map[string]string{
		provisionDisksKey: `<provision><disk device="1" size="1073741824" sr="" bootable="false" type="user"/></provision>`,
	}

	size, err := templateDiskSize(oneDisk(40), TemplateSource("packer-base"), otherConfig)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if size != 40*BytesPerGB {
		t.Errorf("templateDiskSize() = %d, want %d", size, 40*BytesPerGB)
	}
}

func TestTemplateDiskSize_FailsWhenNothingSuppliesASize(t *testing.T) {
	_, err := templateDiskSize(&fakeDisks{}, TemplateSource("Other install media"), nil)

	if err == nil {
		t.Fatal("a source with no provisioning XML and no disks supplies no size")
	}
	for _, want := range []string{"Other install media", "disk_size_gb"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %q", err, want)
		}
	}
}

func TestTemplateDiskSize_ReportsAPoolFailure(t *testing.T) {
	api := &fakeDisks{listErr: errors.New("HOST_IS_SLAVE")}

	if _, err := templateDiskSize(api, TemplateSource("CentOS 9"), nil); err == nil {
		t.Fatal("a failure listing the template's disks must not read as a missing size")
	}
}

func TestEnsureCloneSourceHasPrimaryDisk(t *testing.T) {
	if err := ensureCloneSourceHasPrimaryDisk(oneDisk(24), VMSource("win11-base")); err != nil {
		t.Errorf("a source with a disk on device 0 should be accepted: %s", err)
	}

	err := ensureCloneSourceHasPrimaryDisk(&fakeDisks{}, TemplateSource("Other install media"))
	if err == nil {
		t.Fatal("disk_size_gb against a source with no disk on device 0 must be refused")
	}
	for _, want := range []string{"disk_size_gb", "Other install media"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %q", err, want)
		}
	}
}

func TestStripProvisioningKeys(t *testing.T) {
	otherConfig := map[string]string{
		provisionDisksKey:  provisionOneDisk,
		provisionScriptKey: "/opt/xensource/packages/post-install-scripts/debian-postinstall",
		"owner":            "packer",
	}
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_other_config":         otherConfig,
		"VM.remove_from_other_config": "",
	})

	stripped, err := StripProvisioningKeys(fake.conn(), xenapi.VMRef("OpaqueRef:vm"))

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if want := []string{provisionDisksKey, provisionScriptKey}; !reflect.DeepEqual(stripped, want) {
		t.Errorf("stripped = %#v, want %#v", stripped, want)
	}

	var removed []string
	for _, args := range fake.allArgsOf("VM.remove_from_other_config") {
		if len(args) != 2 {
			t.Fatalf("remove_from_other_config args = %#v, want a reference and a key", args)
		}
		removed = append(removed, args[1].(string))
	}
	if want := []string{provisionDisksKey, provisionScriptKey}; !reflect.DeepEqual(removed, want) {
		t.Errorf("removed %#v, want %#v -- vm_other_config keys are not this step's to touch", removed, want)
	}
}

func TestStripProvisioningKeys_WritesNothingWhenThereIsNothingToStrip(t *testing.T) {
	fake := newFakeXAPI(t, map[string]interface{}{
		"VM.get_other_config": map[string]string{"owner": "packer"},
	})

	stripped, err := StripProvisioningKeys(fake.conn(), xenapi.VMRef("OpaqueRef:vm"))

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if len(stripped) != 0 {
		t.Errorf("stripped = %#v, want nothing", stripped)
	}
	if fake.wasCalled("VM.remove_from_other_config") {
		t.Error("a VM with neither key must not be written to")
	}
}
