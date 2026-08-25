//go:build frank_test_reduced_limits

package appipc

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
)

func TestReducedLimitsArtifactAndAssertions(t *testing.T) {
	if FrameMax != 4_096 || AdmissionRefEncMax != 8_192 || ManifestMax != 2_048 || EntryMax != 512 {
		t.Fatalf("reduced primary limits = (%d, %d, %d, %d)", FrameMax, AdmissionRefEncMax, ManifestMax, EntryMax)
	}
	if ParkedMax != 1_024 || MaxParkedRowsPerRun != 4 || ParkedRowMax != 256 || PathMaxM10 != 512 || OverheadMax != 1_024 {
		t.Fatalf("reduced aggregate limits do not match the owner-final table")
	}
	if AttemptAckMembersMax != 512 || ChainDepthMax != 4 || IDMax != 64 {
		t.Fatalf("reduced remaining limits do not match the owner-final table")
	}
	if got := AdmissionRefEncMax + ManifestMax + ParkedMax + PathMaxM10 + OverheadMax; got <= FrameMax {
		t.Fatalf("flipped P1 sum = %d, want > FRAME_MAX %d", got, FrameMax)
	}
	if got := AttemptAckMembersMax + MaxParkedRowsPerRun*ParkedRowMax; got > FrameMax {
		t.Fatalf("P2 sum = %d, exceeds FRAME_MAX %d", got, FrameMax)
	}

	artifact := ReducedLimitsArtifact()
	digest := sha256.Sum256(artifact)
	if got := hex.EncodeToString(digest[:]); got != "33e67a09fa68000b6c01d728165f1cc5df6327509d8a4d70ce11f99b2a301403" {
		t.Fatalf("reduced artifact SHA-256 = %s", got)
	}
	var decoded struct {
		Mode  string         `json:"mode"`
		Table map[string]int `json:"table"`
	}
	if err := json.Unmarshal(artifact, &decoded); err != nil {
		t.Fatalf("decode reduced artifact: %v", err)
	}
	if decoded.Mode != "test-reduced-limits" || decoded.Table["FRAME_MAX"] != FrameMax || decoded.Table["E_MAX"] != EntryMax {
		t.Fatalf("reduced artifact is not the selected table: %#v", decoded)
	}
}

func TestReducedContinuationWitnessExactFitAndOneByteOver(t *testing.T) {
	skeleton := reducedWitnessFrame("")
	if len(skeleton) != 770 {
		t.Fatalf("|S_cont| = %d, want 770", len(skeleton))
	}
	if got := sha256.Sum256(skeleton); hex.EncodeToString(got[:]) != "60a40957f033255cdfe24b0d1b77a82805726fe727f45625576816c435bfe2f3" {
		t.Fatalf("S_cont SHA-256 = %x", got)
	}

	lstar := FrameMax - len(skeleton)
	if lstar != 3_326 {
		t.Fatalf("L* = %d, want 3326", lstar)
	}
	exact := reducedWitnessFrame(strings.Repeat("a", lstar))
	over := reducedWitnessFrame(strings.Repeat("a", lstar+1))
	if len(exact) != FrameMax || ClassifyTurnOpenSize(len(exact), true) != TurnOpenSizeFits {
		t.Fatalf("exact witness size/disposition = %d/%q", len(exact), ClassifyTurnOpenSize(len(exact), true))
	}
	if got := sha256.Sum256(exact); hex.EncodeToString(got[:]) != "740553dfa398b769f75f8a6358fbbf1c98c94699e431a6323ac444fa7dbc4806" {
		t.Fatalf("W SHA-256 = %x", got)
	}
	if len(over) != FrameMax+1 || ClassifyTurnOpenSize(len(over), true) != TurnOpenResumeOverflow {
		t.Fatalf("over witness size/disposition = %d/%q", len(over), ClassifyTurnOpenSize(len(over), true))
	}

	for _, n := range []int{0, 1, 2, 1_000} {
		if got := len(reducedWitnessFrame(strings.Repeat("a", n))); got != len(skeleton)+n {
			t.Fatalf("step-by-one at n=%d: got %d, want %d", n, got, len(skeleton)+n)
		}
	}
	for n, want := range map[int]int{lstar: 3_367, lstar + 1: 3_368} {
		encoded, err := MarshalJCS(map[string]any{"kind": "operator_input", "task_input": strings.Repeat("a", n)})
		if err != nil {
			t.Fatalf("encode admission_ref: %v", err)
		}
		if len(encoded) != want || len(encoded) > AdmissionRefEncMax {
			t.Fatalf("admission_ref(%d) length = %d, want %d within %d", n, len(encoded), want, AdmissionRefEncMax)
		}
	}
}

func reducedWitnessFrame(taskInput string) []byte {
	providerEntry := map[string]any{
		"kind": "provider", "class": "uncertain", "run_id": "run-000000000000000000000000",
		"turn_id": "turn-00000000000000000000", "attempt_id": "att-0000000000000001", "terminal": "completed",
	}
	manifest := map[string]any{
		"settlement_manifest_v": 1, "run_id": "run-000000000000000000000000",
		"produced_for_turn_id": "turn-00000000000000000001", "entries": []any{providerEntry},
	}
	frame := map[string]any{
		"v": 1, "chan": "ctrl-w", "type": "turn_open", "seq": "0",
		"run_id": "run-000000000000000000000000", "turn_epoch": "1",
		"body": map[string]any{
			"turn_id":        "turn-00000000000000000001",
			"admission_ref":  map[string]any{"kind": "operator_input", "task_input": taskInput},
			"parked_unknown": []any{}, "run_disposition": "resume", "create_auth_id": strings.Repeat("0", 32),
			"session_log_path":    "/var/run/frank/run-000000000000000000000000/session.log",
			"settlement_manifest": manifest, "predecessor_turn_id": "turn-00000000000000000000",
		},
	}
	encoded, err := MarshalJCS(frame)
	if err != nil {
		panic(err)
	}
	return encoded
}
