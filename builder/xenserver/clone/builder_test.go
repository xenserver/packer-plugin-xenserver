package clone

import (
	"strings"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/packer"
)

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"remote_host":     "localhost",
		"remote_username": "admin",
		"remote_password": "admin",
		"clone_template":  "SomeTemplate",
		"vm_name":         "foo",
		"ssh_username":    "foo",

		"packer_build_name":      "foo",
		"skip_cert_verification": true,
	}
}

func TestBuilder_ImplementsBuilder(t *testing.T) {
	var raw interface{}
	raw = &Builder{}
	if _, ok := raw.(packer.Builder); !ok {
		t.Error("Builder must implement builder.")
	}
}

func TestBuilderPrepare_Defaults(t *testing.T) {
	var b Builder
	config := testConfig()
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}
}

func TestBuilderPrepare_InvalidKey(t *testing.T) {
	var b Builder
	config := testConfig()

	// Add a random key
	config["i_should_not_be_valid"] = true
	_, warns, err := b.Prepare(config)
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}

func TestBuilderPrepare_MissingCloneTemplate_KeepsOtherErrors(t *testing.T) {
	var b Builder
	config := testConfig()

	// Remove two required fields so that both errors must be accumulated.
	delete(config, "clone_template")
	delete(config, "remote_username")

	_, _, err := b.Prepare(config)
	if err == nil {
		t.Fatal("should have error")
	}

	if !strings.Contains(err.Error(), "Source Template / VM not specified") {
		t.Errorf("expected missing clone_template error, got: %s", err)
	}
	if !strings.Contains(err.Error(), "remote_username must be specified.") {
		t.Errorf("expected missing remote_username error, got: %s", err)
	}
}
