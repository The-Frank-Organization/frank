package engine

import (
	"slices"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/tables"
)

type DetectorKey struct {
	Phase      string
	RecordKind string
}

type DetectorConfig struct {
	AFloor            map[DetectorKey]string
	TargetBranchField string
	ProtectedBranches []string
}

func KnownADetector(reg *fieldspec.Registry, tab *tables.T, cfg DetectorConfig) fieldspec.KnownADetector {
	return func(cand record.Record, fields map[string]string) (string, bool) {
		if member, ok := s2ReferencedGate(reg, tab, cand, fields); ok {
			return member, true
		}
		if member, ok := s3MergeSplit(fields, cfg); ok {
			return member, true
		}
		if member, ok := s1AFloor(fields, cfg); ok {
			return member, true
		}
		return "", false
	}
}

func s1AFloor(fields map[string]string, cfg DetectorConfig) (string, bool) {
	if len(cfg.AFloor) == 0 {
		return "", false
	}
	member, ok := cfg.AFloor[DetectorKey{Phase: fields["PHASE"], RecordKind: fields["record_kind"]}]
	return member, ok && member != ""
}

func s2ReferencedGate(reg *fieldspec.Registry, tab *tables.T, cand record.Record, fields map[string]string) (string, bool) {
	ref := cand.Headers["resolves_gate"]
	if ref == "" {
		ref = fields["resolves_gate"]
	}
	if ref == "" || tab == nil {
		return "", false
	}
	gate, ok := tab.ByRelay[ref]
	if !ok || gate.Envelope.DeliveryState != record.Accepted {
		return "", false
	}
	member := gate.Headers["gate_category"]
	class, raised := reg.ClassifyGateCategory(member, false)
	if class != "A" || raised {
		return "", false
	}
	return member, true
}

func s3MergeSplit(fields map[string]string, cfg DetectorConfig) (string, bool) {
	field := cfg.TargetBranchField
	if field == "" {
		return "", false
	}
	target := fields[field]
	if target == "" || !slices.Contains(cfg.ProtectedBranches, target) {
		return "", false
	}
	return "merge_to_protected", true
}
