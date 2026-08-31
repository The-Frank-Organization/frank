package fixtures_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/derived"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	frankrecover "github.com/The-Frank-Organization/frank/internal/recover"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16DerivedFoldIsPerRecordAcceptedOnlyAndOrderFree(t *testing.T) {
	sourceA := h16DerivedSource("z-source", "1", "gate")
	sourceB := h16DerivedSource("a-source", "1", "seat_mint")
	legacy := h16DerivedSource("legacy", "", "gate")
	future := h16DerivedSource("future", "2", "gate")
	rejectedFuture := h16DerivedSource("rejected-future", "2", "gate")
	rejectedFuture.Envelope.DeliveryState = record.Rejected
	advance := h16DerivedRecord("m-advance", "derived-work-transition", map[string]any{
		"source_relay_id": "z-source", "kind": "cursor_advance", "completed_hooks": []string{"gate", "approval"},
	})
	attempt := h16DerivedRecord("b-attempt", "derived-work-attempt", map[string]any{
		"source_relay_id": "a-source", "hook": "mint", "state": "running_or_unknown", "predecessor": "none",
	})

	forward := []record.Record{sourceA, sourceB, legacy, future, rejectedFuture, advance, attempt}
	reverse := append([]record.Record(nil), forward...)
	for left, right := 0, len(reverse)-1; left < right; left, right = left+1, right-1 {
		reverse[left], reverse[right] = reverse[right], reverse[left]
	}
	want := map[string]derived.WorkStatus{
		"z-source": {Cursor: []string{}, Status: ""},
		"a-source": {Cursor: []string{"mint"}, Status: "unknown"},
		"future":   {Cursor: []string{"gate", "approval"}, Status: "unknown"},
	}
	if got := derived.Fold(forward); !reflect.DeepEqual(got, want) {
		t.Fatalf("forward fold=%#v, want %#v", got, want)
	}
	if got := derived.Fold(reverse); !reflect.DeepEqual(got, want) {
		t.Fatalf("reverse fold=%#v, want %#v", got, want)
	}
	if got := derived.Cursor(legacy); len(got) != 2 {
		t.Fatalf("pure legacy cursor=%v, want gate+approval before membership classification", got)
	}
}

func TestH16DerivedFoldKeepsTwoConcurrentOpenItems(t *testing.T) {
	records := []record.Record{
		h16DerivedSource("source-one", "1", "gate"),
		h16DerivedSource("source-two", "1", "gate"),
	}
	got := derived.Fold(records)
	for _, id := range []string{"source-one", "source-two"} {
		if !reflect.DeepEqual(got[id], derived.WorkStatus{Cursor: []string{"gate", "approval"}, Status: "pending"}) {
			t.Fatalf("%s fold=%+v", id, got[id])
		}
	}
}

func TestH16DerivedFoldFailedParkPrecedesUnknownMarkerInEveryOrder(t *testing.T) {
	source := h16DerivedSource("parked-with-marker", "1", "seat_mint")
	marker := h16DerivedRecord("open-marker", "derived-work-attempt", map[string]any{
		"source_relay_id": source.Envelope.RelayID, "hook": "mint", "state": "running_or_unknown", "predecessor": "none",
	})
	park := h16DerivedRecord("failed-park", "derived-work-transition", map[string]any{
		"source_relay_id": source.Envelope.RelayID, "kind": "parked", "reason": "operator ceiling",
	})
	for _, records := range [][]record.Record{{source, marker, park}, {park, marker, source}} {
		if got := derived.Fold(records)[source.Envelope.RelayID]; got.Status != "failed" {
			t.Fatalf("fold=%+v, want failed park to precede unresolved marker", got)
		}
	}
}

func TestH16DecisionConsumesJournalWhileDerivedCursorRemainsOpen(t *testing.T) {
	root := t.TempDir()
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("open journal: %v", err)
	}
	intakeID, err := journal.Append(intake.Cmd{Seat: "operator", Role: "operator", Verb: "submit", Payload: json.RawMessage(`{}`)})
	if err != nil {
		t.Fatalf("append command: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	source := h16DerivedSource("source-consumed", "1", "gate")
	source.Envelope.IntakeID = intakeID
	if _, err := st.Commit(source, nil); err != nil {
		t.Fatalf("commit source: %v", err)
	}
	unconsumed, err := intake.Unconsumed(context.Background(), journal, st)
	if err != nil {
		t.Fatalf("Unconsumed: %v", err)
	}
	if len(unconsumed) != 0 {
		t.Fatalf("decision command remained unconsumed: %+v", unconsumed)
	}
	status := derived.Fold([]record.Record{source})["source-consumed"]
	if status.Status != "pending" {
		t.Fatalf("derived state=%+v, want pending despite consumed command", status)
	}
}

func TestH16LoopStampsAndBlindReplayAdvancesToTerminalIdempotently(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers: map[string]string{
				"PHASE": "SITREP", "SUBJECT": "h16 blind source", "resolves_gate": "gate-target",
			},
		}, nil, nil
	}, engine.TestReady())
	var gateCalls, approvalCalls int
	loop.AfterGateResolution = func(record.Record) error {
		gateCalls++
		if gateCalls == 1 {
			return errors.New("injected first gate failure")
		}
		return nil
	}
	loop.AfterApprovalResolution = func(record.Record) error {
		approvalCalls++
		return nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)
	cmd := intake.Cmd{IntakeID: "h16-blind-replay", Seat: "operator", Role: "operator", Verb: "submit"}
	first := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "pending", false)
	source, err := st.Read(first.RelayID)
	if err != nil {
		t.Fatalf("read source: %v", err)
	}
	if source.Headers["hook_contract"] != "1" {
		t.Fatalf("source hook_contract=%q, want 1", source.Headers["hook_contract"])
	}

	second := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, second, record.Accepted, "complete", true)
	before, err := st.Records()
	if err != nil {
		t.Fatalf("records before idempotent replay: %v", err)
	}
	third := h16SubmitLoop(t, loop, cmd)
	after, err := st.Records()
	if err != nil {
		t.Fatalf("records after idempotent replay: %v", err)
	}
	assertH16PostCommitState(t, third, record.Accepted, "complete", true)
	if len(after) != len(before) || gateCalls != 2 || approvalCalls != 1 {
		t.Fatalf("idempotency records=%d->%d gate=%d approval=%d", len(before), len(after), gateCalls, approvalCalls)
	}
	if status := derived.Fold(after)[first.RelayID]; status.Status != "" || len(status.Cursor) != 0 {
		t.Fatalf("terminal fold=%+v", status)
	}
}

func TestH16RecoveryProcessorStampIsBeforeItsCommit(t *testing.T) {
	mainBytes, err := os.ReadFile(filepath.Join("..", "..", "cmd", "frank", "main.go"))
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	processAt := strings.Index(string(mainBytes), "process := func(cmd intake.Cmd) error")
	stampAt := strings.Index(string(mainBytes[processAt:]), "derived.Stamp(&rec)")
	commitAt := strings.Index(string(mainBytes[processAt:]), "st.Commit(rec, intents)")
	if processAt < 0 || stampAt < 0 || commitAt < 0 || stampAt > commitAt {
		t.Fatalf("recovery processor stamp ordering process=%d stamp=%d commit=%d", processAt, stampAt, commitAt)
	}
}

func TestH16FirstRecoveryCommitIsStampedAndRebuildsPendingFromCanonicalBytes(t *testing.T) {
	for _, class := range []string{"gate", "seat_mint"} {
		t.Run(class, func(t *testing.T) {
			root := t.TempDir()
			sources := s8ConfigSources(t, false)
			if err := store.Init(root, sources); err != nil {
				t.Fatalf("init store: %v", err)
			}
			pinned, err := config.Load(store.StoreRootConfigPaths(root))
			if err != nil {
				t.Fatalf("load pinned config: %v", err)
			}
			journal, err := intake.Open(root)
			if err != nil {
				t.Fatalf("open journal: %v", err)
			}
			intakeID, err := journal.Append(intake.Cmd{Seat: "operator", Role: "operator", Verb: "submit", Payload: json.RawMessage(`{}`)})
			if err != nil {
				t.Fatalf("append command: %v", err)
			}
			_, err = frankrecover.RunWithProcessor(root, pinned, func(cmd intake.Cmd) error {
				rec := h16DerivedSource("", "", class)
				rec.Envelope.IntakeID = cmd.IntakeID
				derived.Stamp(&rec)
				if _, commitErr := (&store.Store{Root: root}).Commit(rec, nil); commitErr != nil {
					return commitErr
				}
				return errors.New("injected failure after recovery decision commit")
			})
			if err == nil {
				t.Fatal("recovery processor unexpectedly completed")
			}
			st, err := store.Open(root)
			if err != nil {
				t.Fatalf("open post-failure store: %v", err)
			}
			records, err := st.Records()
			if err != nil {
				t.Fatalf("records: %v", err)
			}
			var source record.Record
			for _, rec := range records {
				if rec.Envelope.IntakeID == intakeID {
					source = rec
					break
				}
			}
			if source.Envelope.RelayID == "" || source.Headers["hook_contract"] != "1" {
				t.Fatalf("recovery source=%+v, want stamped committed decision", source)
			}
			if err := os.RemoveAll(filepath.Join(root, "redo")); err != nil {
				t.Fatalf("remove redo: %v", err)
			}
			if err := os.RemoveAll(filepath.Join(root, "projections")); err != nil {
				t.Fatalf("remove projections: %v", err)
			}
			records, err = st.Records()
			if err != nil {
				t.Fatalf("canonical records after substrate deletion: %v", err)
			}
			if status := derived.Fold(records)[source.Envelope.RelayID]; status.Status != "pending" {
				t.Fatalf("restart fold=%+v, want pending", status)
			}
			result, err := frankrecover.Run(root, pinned)
			if err != nil || result.Ready == nil {
				t.Fatalf("next restart result=%+v err=%v", result, err)
			}
		})
	}
}

func h16DerivedSource(id, contract, class string) record.Record {
	headers := map[string]string{"PHASE": "SITREP", "SUBJECT": id}
	if contract != "" {
		headers["hook_contract"] = contract
	}
	if class == "gate" {
		headers["resolves_gate"] = "gate-target"
	} else {
		headers["record_kind"] = class
	}
	return record.Record{Envelope: record.Envelope{RelayID: id, From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: headers}
}

func h16DerivedRecord(id, kind string, body map[string]any) record.Record {
	data, _ := json.Marshal(body)
	return record.Record{
		Envelope: record.Envelope{RelayID: id, From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": kind}, Body: string(data),
	}
}
