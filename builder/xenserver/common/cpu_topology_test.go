// Copyright (c) 2026. Citrix Systems, Inc.

package common

import (
	"errors"
	"strings"
	"testing"
)

type fakeTopology struct {
	value    string
	found    bool
	readErr  error
	writeErr error
	written  []uint
}

func (f *fakeTopology) CoresPerSocket() (string, bool, error) {
	if f.readErr != nil {
		return "", false, f.readErr
	}
	return f.value, f.found, nil
}

func (f *fakeTopology) SetCoresPerSocket(cores uint) error {
	f.written = append(f.written, cores)
	return f.writeErr
}

func inheritedTopology(value string) *fakeTopology {
	return &fakeTopology{value: value, found: true}
}

func TestReconcileCPUTopology(t *testing.T) {
	for _, tc := range []struct {
		name           string
		api            *fakeTopology
		vcpus          *uint
		coresPerSocket *uint
		wantKind       topologyKind
		wantWritten    []uint
	}{
		{"nothing set", inheritedTopology("4"), nil, nil, topologyUntouched, nil},
		{"cores_per_socket alone", inheritedTopology("4"), nil, Uint(2), topologyUntouched, nil},
		{"both set", inheritedTopology("4"), Uint(8), Uint(2), topologyWritten, []uint{2}},
		{"vcpus alone, no inherited topology", &fakeTopology{}, Uint(3), nil, topologyAbsent, nil},
		{"vcpus alone, inherited topology divides", inheritedTopology("4"), Uint(8), nil, topologyKept, nil},
		{"vcpus alone, inherited topology is the vcpus", inheritedTopology("4"), Uint(4), nil, topologyKept, nil},
		{"vcpus alone, inherited topology does not divide", inheritedTopology("4"), Uint(6), nil, topologyReset, []uint{1}},
		{"vcpus alone, inherited topology exceeds the vcpus", inheritedTopology("4"), Uint(2), nil, topologyReset, []uint{1}},
		{"inherited topology is zero", inheritedTopology("0"), Uint(4), nil, topologyReset, []uint{1}},
		{"inherited topology is unreadable", inheritedTopology("many"), Uint(4), nil, topologyReset, []uint{1}},
		{"inherited topology is an empty value", inheritedTopology(""), Uint(4), nil, topologyReset, []uint{1}},

		{"inherited topology divides but overruns the socket limit", inheritedTopology("1"), Uint(32), nil, topologyReset, []uint{2}},
		{"no inherited topology, and one core per socket overruns it", &fakeTopology{}, Uint(32), nil, topologyReset, []uint{2}},
		{"no topology under the socket limit divides the vcpus", &fakeTopology{}, Uint(17), nil, topologyReset, []uint{17}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			outcome, err := reconcileCPUTopology(tc.api, tc.vcpus, tc.coresPerSocket)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if outcome.kind != tc.wantKind {
				t.Errorf("outcome kind = %d, want %d", outcome.kind, tc.wantKind)
			}
			if len(tc.api.written) != len(tc.wantWritten) {
				t.Fatalf("wrote %#v, want %#v", tc.api.written, tc.wantWritten)
			}
			for i, cores := range tc.wantWritten {
				if tc.api.written[i] != cores {
					t.Errorf("wrote %#v, want %#v", tc.api.written, tc.wantWritten)
				}
			}
		})
	}
}

func TestReconcileCPUTopology_ResetSaysWhatItDidAndWhy(t *testing.T) {
	outcome, err := reconcileCPUTopology(inheritedTopology("4"), Uint(6), nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	message := outcome.Describe("Windows 11")
	for _, want := range []string{"4", "Windows 11", "cores_per_socket"} {
		if !strings.Contains(message, want) {
			t.Errorf("%q does not mention %q", message, want)
		}
	}
}

func TestReconcileCPUTopology_ResetWithoutAnInheritedTopologySaysSo(t *testing.T) {
	outcome, err := reconcileCPUTopology(&fakeTopology{}, Uint(32), nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	message := outcome.Describe("Other install media")
	for _, want := range []string{"Other install media", "32", "16", "cores_per_socket"} {
		if !strings.Contains(message, want) {
			t.Errorf("%q does not mention %q", message, want)
		}
	}
	if strings.Contains(message, "Keeping") {
		t.Errorf("a reset topology must not read as a kept one: %q", message)
	}
}

func TestReconcileCPUTopology_AbsentTopologyIsNotDescribedAsKept(t *testing.T) {
	outcome, err := reconcileCPUTopology(&fakeTopology{}, Uint(4), nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if message := outcome.Describe("Other install media"); strings.Contains(message, "Keeping") {
		t.Errorf("a source with no topology has none to keep: %q", message)
	}
}

func TestReconcileCPUTopology_EmptyInheritedValueIsNotDescribedAsAbsent(t *testing.T) {
	outcome, err := reconcileCPUTopology(inheritedTopology(""), Uint(4), nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	message := outcome.Describe("win11-base")
	if strings.Contains(message, "carries no CPU topology") {
		t.Errorf("the key is there, it is its value that is unusable: %q", message)
	}
	for _, want := range []string{"win11-base", "cores_per_socket"} {
		if !strings.Contains(message, want) {
			t.Errorf("%q does not mention %q", message, want)
		}
	}
	if strings.Contains(message, "16") {
		t.Errorf("4 vcpus at one core per socket is 4 sockets, not more than 16: %q", message)
	}
}

func TestReconcileCPUTopology_UntouchedTopologyClaimsNothing(t *testing.T) {
	outcome, err := reconcileCPUTopology(&fakeTopology{}, nil, nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if message := outcome.Describe("Other install media"); strings.Contains(message, "Keeping") {
		t.Errorf("the VM was never read, so there is no topology to report keeping: %q", message)
	}
}

func TestReconcileCPUTopology_KeptTopologyNamesTheCoresPerSocket(t *testing.T) {
	outcome, err := reconcileCPUTopology(inheritedTopology("4"), Uint(8), nil)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	message := outcome.Describe("win11-base")
	for _, want := range []string{"Keeping", "win11-base", "4"} {
		if !strings.Contains(message, want) {
			t.Errorf("%q does not mention %q", message, want)
		}
	}
}

func TestValidateCPUTopology(t *testing.T) {
	for _, tc := range []struct {
		name           string
		vcpus          *uint
		coresPerSocket *uint
		wantErr        string
	}{
		{"neither set", nil, nil, ""},
		{"vcpus alone", Uint(4), nil, ""},
		{"cores_per_socket alone", nil, Uint(4), ""},
		{"a zero is CommonConfig.Prepare's to report", Uint(0), Uint(4), ""},
		{"one socket", Uint(4), Uint(4), ""},
		{"four sockets", Uint(8), Uint(2), ""},
		{"the socket limit exactly", Uint(16), Uint(1), ""},
		{"not a multiple", Uint(6), Uint(4), "must divide"},
		{"cores per socket above the vcpus", Uint(2), Uint(4), "must divide"},
		{"one socket too many", Uint(17), Uint(1), "sockets"},
		{"twice the socket limit", Uint(64), Uint(2), "sockets"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateCPUTopology(tc.vcpus, tc.coresPerSocket)

			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}
				return
			}
			if err == nil {
				t.Fatal("want an error")
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("error %q does not mention %q", err, tc.wantErr)
			}
			if !strings.Contains(err.Error(), "cores_per_socket") {
				t.Errorf("error %q does not name the option to change", err)
			}
		})
	}
}

func TestTopologyFits(t *testing.T) {
	for _, tc := range []struct {
		vcpus, coresPerSocket uint
		want                  bool
	}{
		{4, 1, true},
		{4, 4, true},
		{6, 4, false},
		{4, 0, false},
		{16, 1, true},
		{17, 1, false},
		{32, 1, false},
		{32, 2, true},
	} {
		if got := topologyFits(tc.vcpus, tc.coresPerSocket); got != tc.want {
			t.Errorf("topologyFits(%d, %d) = %v, want %v", tc.vcpus, tc.coresPerSocket, got, tc.want)
		}
	}
}

func TestFallbackCoresPerSocket(t *testing.T) {
	for _, tc := range []struct{ vcpus, want uint }{
		{0, 1},
		{1, 1},
		{16, 1},
		{17, 17},
		{24, 2},
		{32, 2},
		{64, 4},
	} {
		got := fallbackCoresPerSocket(tc.vcpus)
		if got != tc.want {
			t.Errorf("fallbackCoresPerSocket(%d) = %d, want %d", tc.vcpus, got, tc.want)
		}
		if tc.vcpus > 0 && !topologyFits(tc.vcpus, got) {
			t.Errorf("fallbackCoresPerSocket(%d) = %d, which does not itself fit", tc.vcpus, got)
		}
	}
}

func TestReconcileCPUTopology_ReportsAPoolFailure(t *testing.T) {
	read := &fakeTopology{readErr: errors.New("HOST_IS_SLAVE")}
	if _, err := reconcileCPUTopology(read, Uint(4), nil); err == nil {
		t.Error("a failure reading the VM's platform must not read as a kept topology")
	}

	write := &fakeTopology{writeErr: errors.New("MAP_DUPLICATE_KEY"), value: "4", found: true}
	if _, err := reconcileCPUTopology(write, Uint(6), nil); err == nil {
		t.Error("a failure writing cores-per-socket must not read as a reset topology")
	}
}
