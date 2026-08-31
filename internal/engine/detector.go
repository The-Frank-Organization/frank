package engine

import (
	"encoding/json"
	"fmt"
	"slices"

	frankconfig "github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/tables"
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

// DetectorConfig wires Step-1 known-A detection for exactly S1 + S2 + S3
// plus the hardcoded other->A fail-safe. With TargetBranchField unset, the
// S3 mechanism is present but input-atom-pending and therefore inert.
type pinnedDetectorConfig struct {
	Detector *struct {
		AFloor []struct {
			Phase      string `json:"phase"`
			RecordKind string `json:"record_kind"`
			Member     string `json:"member"`
		} `json:"a_floor"`
		TargetBranchField string   `json:"target_branch_field"`
		ProtectedBranches []string `json:"protected_branches"`
	} `json:"detector"`
}

func DetectorConfigFromPinned(pinned *frankconfig.Pinned) (DetectorConfig, error) {
	if pinned == nil {
		return DetectorConfig{}, nil
	}
	engineBytes := pinned.Members["engine"]
	if len(engineBytes) == 0 {
		return DetectorConfig{}, nil
	}
	var raw pinnedDetectorConfig
	if err := json.Unmarshal(engineBytes, &raw); err != nil {
		return DetectorConfig{}, fmt.Errorf("parse detector config: %w", err)
	}
	if raw.Detector == nil {
		return DetectorConfig{}, nil
	}

	cfg := DetectorConfig{
		TargetBranchField: raw.Detector.TargetBranchField,
		ProtectedBranches: append([]string(nil), raw.Detector.ProtectedBranches...),
	}
	if len(raw.Detector.AFloor) == 0 {
		return cfg, nil
	}
	if pinned.Registry == nil {
		return DetectorConfig{}, fmt.Errorf("detector a_floor requires fieldspec registry")
	}
	cfg.AFloor = make(map[DetectorKey]string, len(raw.Detector.AFloor))
	for i, row := range raw.Detector.AFloor {
		if row.Member == "" {
			return DetectorConfig{}, fmt.Errorf("detector a_floor[%d]: member required", i)
		}
		if !slices.Contains(pinned.Registry.GateCategory["A"], row.Member) {
			return DetectorConfig{}, fmt.Errorf("detector a_floor[%d]: member %q not in gate_category_A", i, row.Member)
		}
		cfg.AFloor[DetectorKey{Phase: row.Phase, RecordKind: row.RecordKind}] = row.Member
	}
	return cfg, nil
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
