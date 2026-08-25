package resume

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/worker/jcs"
	"github.com/jackli/frank/internal/worker/journal"
)

func TestAllFiveFirstActionBranchesAreReachable(t *testing.T) {
	now := time.Unix(100, 0)
	cases := []struct {
		name    string
		entry   SettlementEntry
		content ContentObservation
		want    FirstAction
	}{
		{"clean positive", SettlementEntry{Kind: KindTool, Class: SettledWithContent, RunID: "r", TurnID: "t", EffectID: "call", ArgsDigest: digest("a"), SettledAt: now}, ContentObservation{Present: true, DigestValid: true, DurableAt: now.Add(-time.Second), InspectedAt: now.Add(time.Second)}, ContinueModelLoop},
		{"determinate terminal", SettlementEntry{Kind: KindProvider, Class: DeterminateNoResume, RunID: "r", TurnID: "t", EffectID: "attempt", Terminal: "transport_failed", SettledAt: now}, ContentObservation{InspectedAt: now.Add(time.Second)}, SurfaceAndTerminalize},
		{"uncertain tool", SettlementEntry{Kind: KindTool, Class: Uncertain, RunID: "r", TurnID: "t", EffectID: "call", ArgsDigest: digest("a"), SettledAt: now}, ContentObservation{InspectedAt: now.Add(time.Second)}, SurfaceUncertainTool},
		{"uncertain provider", SettlementEntry{Kind: KindProvider, Class: Uncertain, RunID: "r", TurnID: "t", EffectID: "attempt", SettledAt: now}, ContentObservation{InspectedAt: now.Add(time.Second)}, SurfaceUncertainProvider},
		{"content lost", SettlementEntry{Kind: KindProvider, Class: SettledWithContent, RunID: "r", TurnID: "t", EffectID: "attempt", SettledAt: now}, ContentObservation{Present: false, InspectedAt: now.Add(time.Second)}, ReDerive},
	}
	seen := map[FirstAction]bool{}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			decision, err := Reconcile(testCase.entry, testCase.content)
			if err != nil {
				t.Fatalf("Reconcile: %v", err)
			}
			if decision.FirstAction != testCase.want {
				t.Fatalf("first action = %q, want %q", decision.FirstAction, testCase.want)
			}
			seen[decision.FirstAction] = true
		})
	}
	if len(seen) != 5 {
		t.Fatalf("reached %d first-action branches, want 5", len(seen))
	}
}

func TestTrustWindowViolationsNeverTrustContent(t *testing.T) {
	settled := time.Unix(100, 0)
	entry := SettlementEntry{Kind: KindProvider, Class: SettledWithContent, RunID: "r", TurnID: "t", EffectID: "attempt", SettledAt: settled}
	cases := []ContentObservation{
		{Present: true, DigestValid: true, DurableAt: settled.Add(time.Second), InspectedAt: settled.Add(2 * time.Second)},
		{Present: true, DigestValid: true, DurableAt: settled.Add(-time.Second), InspectedAt: settled.Add(-2 * time.Second)},
		{Present: true, DigestValid: false, DurableAt: settled.Add(-time.Second), InspectedAt: settled.Add(time.Second)},
	}
	for _, observation := range cases {
		decision, err := Reconcile(entry, observation)
		if err == nil || decision.Trusted || decision.FirstAction == ContinueModelLoop {
			t.Fatalf("trust-window violation was accepted: decision=%+v err=%v", decision, err)
		}
	}
}

func TestSettlementManifestClosedUnionAndIdentityGrain(t *testing.T) {
	manifest := map[string]any{
		"settlement_manifest_v": 1, "run_id": "run", "produced_for_turn_id": "next",
		"entries": []any{
			map[string]any{"kind": "tool", "class": "settled_with_content", "run_id": "run", "turn_id": "source-a", "tool_call_id": "same", "args_digest": digest("args-a"), "terminal": "EXECUTED"},
			map[string]any{"kind": "tool", "class": "settled_with_content", "run_id": "run", "turn_id": "source-b", "tool_call_id": "same", "args_digest": digest("args-b"), "terminal": "EXECUTED"},
			map[string]any{"kind": "provider", "class": "uncertain", "run_id": "run", "turn_id": "source-a", "attempt_id": "attempt", "terminal": "UNKNOWN_PROVIDER_OUTCOME"},
		},
	}
	raw, _ := json.Marshal(manifest)
	canonical, _ := jcs.Canonicalize(raw)
	decoded, err := DecodeSettlementManifest(canonical)
	if err != nil || len(decoded.Entries) != 3 {
		t.Fatalf("DecodeSettlementManifest = %+v, %v", decoded, err)
	}
	badShapes := []map[string]any{
		{"kind": "unknown", "class": "uncertain", "run_id": "run", "turn_id": "t", "attempt_id": "a", "terminal": "UNKNOWN"},
		{"kind": "provider", "class": "uncertain", "run_id": "run", "turn_id": "t", "attempt_id": "a", "args_digest": digest("illegal"), "terminal": "UNKNOWN"},
		{"kind": "tool", "class": "settled_with_content", "run_id": "run", "turn_id": "t", "tool_call_id": "c", "args_digest": digest("a"), "terminal": "EXECUTED", "extra": true},
	}
	for _, bad := range badShapes {
		candidate := map[string]any{"settlement_manifest_v": 1, "run_id": "run", "produced_for_turn_id": "next", "entries": []any{bad}}
		raw, _ := json.Marshal(candidate)
		canonical, _ := jcs.Canonicalize(raw)
		if partial, err := DecodeSettlementManifest(canonical); err == nil || len(partial.Entries) != 0 {
			t.Fatalf("bad union returned partial manifest: %+v err=%v", partial, err)
		}
	}
}

func TestContentReadyOrderingAndDispositionGate(t *testing.T) {
	t0 := time.Unix(100, 0)
	frame, err := ContentReady(ContentReadyInput{Seq: "7", RunID: "r", TurnEpoch: "2", TurnID: "t", AttemptID: "a", RoundIdentity: digest("round"), SeqHWM: "5", GenerationID: "g", ContentFsyncAt: t0, MarkerFsyncAt: t0.Add(time.Second), EmitAt: t0.Add(2 * time.Second)})
	if err != nil || frame.Type != "content_ready" || frame.ReplyTo != "" {
		t.Fatalf("ContentReady = %+v, %v", frame, err)
	}
	bad := ContentReadyInput{Seq: "7", RunID: "r", TurnEpoch: "2", TurnID: "t", AttemptID: "a", RoundIdentity: digest("round"), SeqHWM: "5", GenerationID: "g", ContentFsyncAt: t0.Add(2 * time.Second), MarkerFsyncAt: t0, EmitAt: t0.Add(time.Second)}
	if _, err := ContentReady(bad); err == nil {
		t.Fatal("content_ready accepted before content/marker durability")
	}

	gate := NewWorkGate("r", "t", "2")
	report, err := gate.Report("8", DispositionResumable)
	if err != nil || report.Type != "report_resume_disposition" || gate.WorkAllowed() {
		t.Fatalf("report gate = %+v allowed=%v err=%v", report, gate.WorkAllowed(), err)
	}
	if err := gate.Commit(DispositionDegraded, ResumeActionReDerive); err != nil || !gate.WorkAllowed() || gate.CommittedDisposition() != DispositionDegraded {
		t.Fatalf("committed pair not adopted: allowed=%v disposition=%q err=%v", gate.WorkAllowed(), gate.CommittedDisposition(), err)
	}
}

func TestPrefixOracleStopsAtFrozenFullKeyAndExcludesRecordDigest(t *testing.T) {
	identity := journal.Identity{RunID: "run-1", RunManifestDigest: strings.Repeat("a", 64), CreateAuthID: strings.Repeat("b", 32)}
	records := []journal.Record{
		mustRecord(t, journal.KindRunOpen, "0", "", "", map[string]any{"run_id": "run-1", "run_manifest_digest": strings.Repeat("a", 64), "turn_epoch": "1", "create_auth_id": strings.Repeat("b", 32)}),
		mustRecord(t, journal.KindToolCall, "1", "turn-old", "0", map[string]any{"tool_call_id": "c0", "canonical_tool_name": "read", "canonical_args_digest": digest("args0"), "args": map[string]any{}}),
	}
	records = append(records, mustMarker(t, "2", "turn-old", "0", records[1:2]))
	records = append(records, mustRecord(t, journal.KindProviderOutput, "3", "turn-target", "0", map[string]any{"attempt_id": "a1", "item_index": "0", "content": map[string]any{"opaque": true}}))
	records = append(records, mustMarker(t, "4", "turn-target", "0", records[3:4]))
	targetCount := len(records)
	records = append(records, mustRecord(t, journal.KindInputItem, "5", "later-turn", "0", map[string]any{"role": "user", "item_index": "0", "content": "later"}))
	data := journalBytes(t, records...)

	actual, err := ExtractCanonicalPrefix(data, identity, "turn-target", "0")
	if err != nil {
		t.Fatalf("ExtractCanonicalPrefix: %v", err)
	}
	if bytes.Contains(actual, []byte("record_digest")) || bytes.Contains(actual, []byte("later-turn")) {
		t.Fatalf("oracle leaked excluded bytes: %s", actual)
	}
	if lines := bytes.Count(actual, []byte{'\n'}); lines != targetCount {
		t.Fatalf("oracle lines = %d, want %d", lines, targetCount)
	}
	if _, err := ExtractCanonicalPrefix(data, identity, "turn-old", "9"); err == nil {
		t.Fatal("bare-index or missing full-key marker was accepted")
	}
}

func TestEditedSessionClassDistinctnessAndNonPromotion(t *testing.T) {
	external := mustRecord(t, journal.KindProviderOutput, "1", "turn", "0", map[string]any{"attempt_id": "a", "item_index": "0", "content": "opaque"})
	model := mustRecord(t, journal.KindToolCall, "1", "turn", "0", map[string]any{"tool_call_id": "c", "canonical_tool_name": "read", "canonical_args_digest": digest("x"), "args": map[string]any{}})

	if result, err := AssessJournalEdit(external, external); err != nil || result.Disposition != DispositionDegraded || result.ResumeAction != ResumeActionReDerive || result.PresentAsOriginalTruth {
		t.Fatalf("external edit promoted: %+v err=%v", result, err)
	}
	if result, err := AssessJournalEdit(model, model); err != nil || result.Disposition != DispositionResumable {
		t.Fatalf("model transcript class not distinct: %+v err=%v", result, err)
	}
	if result, err := AssessJournalEdit(external, model); err == nil || result.Disposition != DispositionDegraded || result.PresentAsOriginalTruth {
		t.Fatalf("cross-class edit promoted: %+v err=%v", result, err)
	}
	if _, err := AssessEdit(SourceM10Settlement, external, external); err == nil {
		t.Fatal("m-10 settlement row entered the journal-only edit surface")
	}
}

func digest(seed string) string {
	sum := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(sum[:])
}

func mustRecord(t *testing.T, kind, seq, turnID, roundIndex string, fields map[string]any) journal.Record {
	t.Helper()
	raw := make(map[string]json.RawMessage, len(fields))
	for name, value := range fields {
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		raw[name] = encoded
	}
	record, err := journal.FinalizeRecord(journal.Record{Seq: seq, Kind: kind, GenerationID: "generation-a", TurnID: turnID, RoundIndex: roundIndex, TSMonotonic: seq, Fields: raw})
	if err != nil {
		t.Fatalf("FinalizeRecord(%s): %v", kind, err)
	}
	return record
}

func mustMarker(t *testing.T, seq, turnID, roundIndex string, members []journal.Record) journal.Record {
	t.Helper()
	marker, err := journal.BuildRoundMarker(seq, "generation-a", turnID, roundIndex, seq, members)
	if err != nil {
		t.Fatal(err)
	}
	return marker
}

func journalBytes(t *testing.T, records ...journal.Record) []byte {
	t.Helper()
	var output bytes.Buffer
	for _, record := range records {
		encoded, err := journal.MarshalRecord(record)
		if err != nil {
			t.Fatal(err)
		}
		output.Write(encoded)
		output.WriteByte('\n')
	}
	return output.Bytes()
}
