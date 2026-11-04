// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"strings"
	"testing"

	"github.com/xenserver/packer-plugin-xenserver/version"
)

func TestBuildProvenanceOtherConfig_StampsWhenNothingElseIsSet(t *testing.T) {
	got := BuildProvenanceOtherConfig(nil, nil)

	if got[ProvenanceKey] != ProvenanceValue {
		t.Errorf("%s = %q, want %q", ProvenanceKey, got[ProvenanceKey], ProvenanceValue)
	}
	if got[ProvenanceVersionKey] != PluginVersion() {
		t.Errorf("%s = %q, want %q", ProvenanceVersionKey, got[ProvenanceVersionKey], PluginVersion())
	}
	if len(got) != 2 {
		t.Errorf("map has %d keys, want only the 2 provenance keys: %v", len(got), got)
	}
}

func TestBuildProvenanceOtherConfig_PreservesUserAndInheritedKeys(t *testing.T) {
	inherited := map[string]string{"template_origin": "golden-2026-08"}
	user := map[string]string{"owner": "packer", "cost_centre": "42"}

	got := BuildProvenanceOtherConfig(inherited, user)

	for key, want := range map[string]string{
		"template_origin": "golden-2026-08",
		"owner":           "packer",
		"cost_centre":     "42",
	} {
		if got[key] != want {
			t.Errorf("%s = %q, want %q", key, got[key], want)
		}
	}
	if got[ProvenanceKey] != ProvenanceValue {
		t.Errorf("%s = %q, want %q", ProvenanceKey, got[ProvenanceKey], ProvenanceValue)
	}
}

func TestBuildProvenanceOtherConfig_ReservedKeysBeatUserKeys(t *testing.T) {
	user := map[string]string{
		ProvenanceKey:        "false",
		ProvenanceVersionKey: "9.9.9",
	}

	got := BuildProvenanceOtherConfig(nil, user)

	if got[ProvenanceKey] != ProvenanceValue {
		t.Errorf("%s = %q, want the plugin value %q to win", ProvenanceKey, got[ProvenanceKey], ProvenanceValue)
	}
	if got[ProvenanceVersionKey] != PluginVersion() {
		t.Errorf("%s = %q, want the plugin value %q to win", ProvenanceVersionKey, got[ProvenanceVersionKey], PluginVersion())
	}
}

func TestBuildProvenanceOtherConfig_OverwritesAStaleInheritedStamp(t *testing.T) {
	inherited := map[string]string{
		ProvenanceKey:        ProvenanceValue,
		ProvenanceVersionKey: "0.0.1",
	}

	got := BuildProvenanceOtherConfig(inherited, nil)

	if got[ProvenanceVersionKey] != PluginVersion() {
		t.Errorf("%s = %q, want the running plugin version %q", ProvenanceVersionKey, got[ProvenanceVersionKey], PluginVersion())
	}
}

func TestBuildProvenanceOtherConfig_DoesNotMutateItsInputs(t *testing.T) {
	inherited := map[string]string{"kept": "yes"}
	user := map[string]string{"owner": "packer"}

	BuildProvenanceOtherConfig(inherited, user)

	if len(inherited) != 1 {
		t.Errorf("inherited map was modified: %v", inherited)
	}
	if len(user) != 1 {
		t.Errorf("user map was modified: %v", user)
	}
}

// FormattedVersion() re-renders the parsed semver, dropping the constant's "v" prefix.
func TestPluginVersion_MatchesTheInCodeVersionConstant(t *testing.T) {
	want := strings.TrimPrefix(version.Version, "v")
	if version.VersionPrerelease != "" {
		want += "-" + version.VersionPrerelease
	}

	if got := PluginVersion(); got != want {
		t.Errorf("PluginVersion() = %q, want %q", got, want)
	}
}
