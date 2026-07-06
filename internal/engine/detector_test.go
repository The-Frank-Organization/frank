package engine_test

import (
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/tables"
)

func TestKnownADetectorSourcesAndPrecedence(t *testing.T) {
	reg := loadEngineRegistry(t)
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "gate-a", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"gate_category": "authz_security"},
	})
	cfg := engine.DetectorConfig{
		AFloor: map[engine.DetectorKey]string{
			{Phase: "PLAN", RecordKind: "diagnostics"}: "product_semantics",
		},
		TargetBranchField: "target_branch",
		ProtectedBranches: []string{"main"},
	}
	detect := engine.KnownADetector(reg, tab, cfg)

	tests := []struct {
		name   string
		rec    record.Record
		fields map[string]string
		want   string
		hit    bool
	}{
		{
			name: "S1 floor",
			fields: map[string]string{
				"PHASE":       "PLAN",
				"record_kind": "diagnostics",
			},
			want: "product_semantics",
			hit:  true,
		},
		{
			name:   "S2 referenced gate",
			rec:    record.Record{Headers: map[string]string{"resolves_gate": "gate-a"}},
			fields: map[string]string{"resolves_gate": "gate-a"},
			want:   "authz_security",
			hit:    true,
		},
		{
			name:   "S3 merge split",
			fields: map[string]string{"target_branch": "main"},
			want:   "merge_to_protected",
			hit:    true,
		},
		{
			name: "S2 beats S3 and S1",
			rec:  record.Record{Headers: map[string]string{"resolves_gate": "gate-a"}},
			fields: map[string]string{
				"PHASE":         "PLAN",
				"record_kind":   "diagnostics",
				"resolves_gate": "gate-a",
				"target_branch": "main",
			},
			want: "authz_security",
			hit:  true,
		},
		{
			name: "S3 beats S1",
			fields: map[string]string{
				"PHASE":         "PLAN",
				"record_kind":   "diagnostics",
				"target_branch": "main",
			},
			want: "merge_to_protected",
			hit:  true,
		},
		{
			name:   "no source hit",
			fields: map[string]string{"PHASE": "SITREP"},
			hit:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			member, hit := detect(tc.rec, tc.fields)
			if member != tc.want || hit != tc.hit {
				t.Fatalf("detect = (%q, %v), want (%q, %v)", member, hit, tc.want, tc.hit)
			}
		})
	}
}

func TestKnownADetectorIgnoresNonAGateReferences(t *testing.T) {
	reg := loadEngineRegistry(t)
	tab := tables.New()
	for _, rec := range []record.Record{
		{Envelope: record.Envelope{RelayID: "gate-b", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"gate_category": "routing"}},
		{Envelope: record.Envelope{RelayID: "gate-other", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"gate_category": "other"}},
		{Envelope: record.Envelope{RelayID: "gate-held", DeliveryState: record.Held, SchemaVersion: 1}, Headers: map[string]string{"gate_category": "authz_security"}},
	} {
		tab.OnCommit(rec)
	}
	detect := engine.KnownADetector(reg, tab, engine.DetectorConfig{})
	for _, ref := range []string{"gate-b", "gate-other", "gate-held", "missing"} {
		t.Run(ref, func(t *testing.T) {
			member, hit := detect(record.Record{Headers: map[string]string{"resolves_gate": ref}}, map[string]string{"resolves_gate": ref})
			if member != "" || hit {
				t.Fatalf("detect = (%q, %v), want no hit", member, hit)
			}
		})
	}
}

func loadEngineRegistry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	reg, err := fieldspec.Load("../fieldspec/registry.json")
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return reg
}
