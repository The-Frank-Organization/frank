package journal

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestClosedRecordUnionRoundTripsCanonicalLines(t *testing.T) {
	records := allKindRecords(t)
	if len(records) != len(RecordKinds()) {
		t.Fatalf("fixture kinds = %d, union kinds = %d", len(records), len(RecordKinds()))
	}
	for _, record := range records {
		t.Run(record.Kind, func(t *testing.T) {
			finalized, err := FinalizeRecord(record)
			if err != nil {
				t.Fatalf("FinalizeRecord: %v", err)
			}
			encoded, err := MarshalRecord(finalized)
			if err != nil {
				t.Fatalf("MarshalRecord: %v", err)
			}
			decoded, err := DecodeRecord(encoded)
			if err != nil {
				t.Fatalf("DecodeRecord: %v", err)
			}
			reencoded, err := MarshalRecord(decoded)
			if err != nil {
				t.Fatalf("re-MarshalRecord: %v", err)
			}
			if !bytes.Equal(reencoded, encoded) {
				t.Fatalf("round trip changed bytes:\n got: %s\nwant: %s", reencoded, encoded)
			}
			if len(encoded) == 0 || encoded[len(encoded)-1] == '\n' {
				t.Fatalf("MarshalRecord returned line framing: %q", encoded)
			}
		})
	}
}

func TestClosedRecordUnionRejectsUnknownKindAndMembers(t *testing.T) {
	base := mustFinalized(t, recordForKind(t, KindToolCall))
	base.Kind = "reasoning_replay"
	if _, err := MarshalRecord(base); err == nil {
		t.Fatal("MarshalRecord accepted an unknown kind")
	}

	base = mustFinalized(t, recordForKind(t, KindToolCall))
	base.Fields["provider_secret"] = rawString(t, "must-not-ride")
	if _, err := MarshalRecord(base); err == nil {
		t.Fatal("MarshalRecord accepted an unknown payload member")
	}
}

func TestRecordDigestIsSelfExcludingAndDetectsContentDrift(t *testing.T) {
	record := mustFinalized(t, recordForKind(t, KindProviderOutput))
	if ok, err := VerifyRecordDigest(record); err != nil || !ok {
		t.Fatalf("VerifyRecordDigest(clean) = %v, %v", ok, err)
	}
	record.Fields["content"] = rawString(t, "changed opaque bytes")
	if ok, err := VerifyRecordDigest(record); err != nil || ok {
		t.Fatalf("VerifyRecordDigest(mutated) = %v, %v", ok, err)
	}
}

func TestOpaqueProviderItemPassesThroughBytePreservedAndUnparsed(t *testing.T) {
	opaque := rawString(t, `opaque_item_passthrough:{"provider_native":true}`)
	record := recordForKind(t, KindProviderOutput)
	record.Fields["content"] = append(json.RawMessage(nil), opaque...)
	decoded, err := DecodeRecord(mustMarshal(t, mustFinalized(t, record)))
	if err != nil {
		t.Fatalf("DecodeRecord: %v", err)
	}
	if !bytes.Equal(decoded.Fields["content"], opaque) {
		t.Fatalf("opaque content changed: got %s, want %s", decoded.Fields["content"], opaque)
	}
}

func TestTransitionTableIsTotalAndNonVacuous(t *testing.T) {
	subjects := TransitionSubjects()
	rows := TransitionRows()
	wantStates := []TrustState{TrustClean, TrustDigestMismatch, TrustStructuralFailure}
	seen := make(map[string]int)
	actions := make(map[TransitionAction]bool)
	for _, row := range rows {
		seen[row.Subject+"/"+string(row.State)]++
		actions[row.Action] = true
	}
	for _, subject := range subjects {
		for _, state := range wantStates {
			key := subject + "/" + string(state)
			if seen[key] != 1 {
				t.Errorf("transition %s has %d rows, want exactly 1", key, seen[key])
			}
		}
	}
	if len(rows) != len(subjects)*len(wantStates) {
		t.Fatalf("transition rows = %d, want %d", len(rows), len(subjects)*len(wantStates))
	}
	for _, action := range []TransitionAction{ActionContinue, ActionDegrade, ActionResolveObjective} {
		if !actions[action] {
			t.Errorf("transition action %q is unexercised", action)
		}
	}
}

func TestRoundMarkerDigestAndMembership(t *testing.T) {
	members := []Record{
		mustFinalized(t, withCoordinates(recordForKind(t, KindToolCall), "1", "generation-a", "turn-1", "0")),
		mustFinalized(t, withCoordinates(recordForKind(t, KindToolResult), "2", "generation-a", "turn-1", "0")),
	}
	marker, err := BuildRoundMarker("3", "generation-a", "turn-1", "0", "3", members)
	if err != nil {
		t.Fatalf("BuildRoundMarker: %v", err)
	}
	if err := HonourRoundMarker(marker, members, nil); err != nil {
		t.Fatalf("HonourRoundMarker: %v", err)
	}

	crossTurn := append([]Record(nil), members...)
	crossTurn[1].TurnID = "turn-2"
	crossTurn[1] = mustFinalized(t, crossTurn[1])
	if err := HonourRoundMarker(marker, crossTurn, nil); err == nil {
		t.Fatal("HonourRoundMarker accepted a cross-turn interval")
	}

	outOfInterval := mustFinalized(t, withCoordinates(recordForKind(t, KindInputItem), "4", "generation-a", "turn-1", "0"))
	if err := HonourRoundMarker(marker, members, []Record{outOfInterval}); err == nil {
		t.Fatal("HonourRoundMarker accepted same-key member outside interval")
	}
}

func allKindRecords(t *testing.T) []Record {
	t.Helper()
	result := make([]Record, 0, len(RecordKinds()))
	for _, kind := range RecordKinds() {
		result = append(result, recordForKind(t, kind))
	}
	return result
}

func recordForKind(t *testing.T, kind string) Record {
	t.Helper()
	record := Record{Seq: "1", Kind: kind, GenerationID: "generation-a", TSMonotonic: "1", Fields: map[string]json.RawMessage{}}
	switch kind {
	case KindRunOpen:
		record.Seq = "0"
		record.TSMonotonic = "0"
		record.Fields = map[string]json.RawMessage{
			"run_id":              rawString(t, "run-1"),
			"run_manifest_digest": rawString(t, strings.Repeat("a", 64)),
			"turn_epoch":          rawString(t, "1"),
			"create_auth_id":      rawString(t, strings.Repeat("b", 32)),
		}
	case KindTurnScope:
		record.TurnID = "turn-1"
		record.Fields = map[string]json.RawMessage{"admission_ref_kind": rawString(t, "operator_input")}
	case KindObjectiveRef:
		record.TurnID = "turn-1"
		record.Fields = map[string]json.RawMessage{
			"objective_locator": rawString(t, "operator-input"),
			"constraint_refs":   json.RawMessage(`[]`),
		}
	case KindWorkspaceSnapshot:
		record.TurnID = "turn-1"
		record.Fields = map[string]json.RawMessage{
			"workspace_root_id": rawString(t, "workspace-1"),
			"snapshot_id":       rawString(t, "snapshot-1"),
		}
	case KindInputItem:
		record.TurnID, record.RoundIndex = "turn-1", "0"
		record.Fields = map[string]json.RawMessage{
			"role":       rawString(t, "assistant"),
			"item_index": rawString(t, "0"),
			"content":    rawString(t, "hello"),
		}
	case KindToolCall:
		record.TurnID, record.RoundIndex = "turn-1", "0"
		record.Fields = map[string]json.RawMessage{
			"tool_call_id":          rawString(t, "call-1"),
			"canonical_tool_name":   rawString(t, "read"),
			"canonical_args_digest": rawString(t, strings.Repeat("c", 64)),
			"args":                  json.RawMessage(`{"path":"README.md"}`),
		}
	case KindToolResult:
		record.TurnID, record.RoundIndex = "turn-1", "0"
		record.Fields = map[string]json.RawMessage{
			"tool_call_id": rawString(t, "call-1"),
			"content":      rawString(t, "result"),
			"truncated":    json.RawMessage(`false`),
		}
	case KindProviderOutput:
		record.TurnID, record.RoundIndex = "turn-1", "0"
		record.Fields = map[string]json.RawMessage{
			"attempt_id": rawString(t, "attempt-1"),
			"item_index": rawString(t, "0"),
			"content":    rawString(t, "opaque-provider-item"),
		}
	case KindCompactionEvent:
		record.TurnID, record.RoundIndex = "turn-1", "0"
		record.Fields = map[string]json.RawMessage{
			"tier":             rawString(t, "evict"),
			"template_id":      rawString(t, "tool-result-eviction"),
			"template_version": rawString(t, "v1"),
			"affected_seq":     json.RawMessage(`["1"]`),
		}
	case KindRoundMarker:
		record.TurnID = "turn-1"
		record.Fields = map[string]json.RawMessage{
			"round_index":   rawString(t, "0"),
			"first_seq":     rawString(t, "0"),
			"last_seq":      rawString(t, "0"),
			"marker_digest": rawString(t, strings.Repeat("d", 64)),
		}
	default:
		t.Fatalf("unhandled kind %q", kind)
	}
	return record
}

func withCoordinates(record Record, seq, generationID, turnID, roundIndex string) Record {
	record.Seq = seq
	record.GenerationID = generationID
	record.TurnID = turnID
	record.RoundIndex = roundIndex
	record.TSMonotonic = seq
	return record
}

func rawString(t *testing.T, value string) json.RawMessage {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("json.Marshal(%q): %v", value, err)
	}
	return raw
}

func mustFinalized(t *testing.T, record Record) Record {
	t.Helper()
	finalized, err := FinalizeRecord(record)
	if err != nil {
		t.Fatalf("FinalizeRecord(%s): %v", record.Kind, err)
	}
	return finalized
}

func mustMarshal(t *testing.T, record Record) []byte {
	t.Helper()
	encoded, err := MarshalRecord(record)
	if err != nil {
		t.Fatalf("MarshalRecord(%s): %v", record.Kind, err)
	}
	return encoded
}

func TestRecordKindsAreStable(t *testing.T) {
	want := []string{
		KindRunOpen, KindTurnScope, KindObjectiveRef, KindWorkspaceSnapshot,
		KindInputItem, KindToolCall, KindToolResult, KindProviderOutput,
		KindCompactionEvent, KindRoundMarker,
	}
	if !reflect.DeepEqual(RecordKinds(), want) {
		t.Fatalf("RecordKinds = %v, want %v", RecordKinds(), want)
	}
}
