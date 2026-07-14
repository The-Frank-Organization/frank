package engine

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestS10ResummonTimerCrashRefireDedupesBySeatDecisionAndCadenceSlot(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	gate := record.Record{
		Envelope: record.Envelope{RelayID: "timer-gate", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "timer gate", "HUMAN_GATE_REQUIRED": "yes"},
	}
	if _, err := st.Commit(gate, nil); err != nil {
		t.Fatalf("Commit gate: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "park-timer-gate", From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": GateParkedWaitingHuman, "parks_gate": "timer-gate"},
	}, nil); err != nil {
		t.Fatalf("Commit park: %v", err)
	}

	fired := make(chan time.Time, 1)
	fired <- time.Unix(17, 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	loop := New(st, ResummonHandler(nil), TestReady())
	go loop.Run(ctx)
	journal, err := intake.Open(st.Root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	writer, err := intake.NewWriter[Outcome](journal, config.EngineConfig{}, TestReady())
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	go writer.Run(ctx, loop.In)
	scheduler, err := NewResummonScheduler(st, writer, resummonTestSnapshot(t, st))
	if err != nil {
		t.Fatalf("NewResummonScheduler: %v", err)
	}
	scheduler.after = func(time.Duration) <-chan time.Time { return fired }
	input := ResummonInput{
		Seat:          "seat-a",
		DecisionID:    "timer-gate",
		CadenceSlot:   "no-response-0001",
		Reason:        ResummonNoResponse,
		SummonChannel: SummonLocal,
	}
	first, err := scheduler.EmitAfter(ctx, time.Hour, input)
	if err != nil {
		t.Fatalf("EmitAfter: %v", err)
	}
	if first.Deduped || first.Record.Headers["record_kind"] != "resummon_command" || first.Record.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("first resummon = %+v", first)
	}
	if first.ContentHash != ResummonContentHash(input) {
		t.Fatalf("content hash = %q, want deterministic key", first.ContentHash)
	}
	if strings.Contains(first.Record.Body, "approve") || strings.Contains(first.Record.Body, "verdict") || strings.Contains(first.Record.Body, "deadline") {
		t.Fatalf("resummon escalated beyond summon channel: %s", first.Record.Body)
	}

	refireInput := input
	refireInput.Reason = ResummonAnsweredStalled
	refireInput.SummonChannel = SummonLouderLocal
	refired, err := scheduler.Emit(ctx, refireInput)
	if err != nil {
		t.Fatalf("crash-refire Emit: %v", err)
	}
	if !refired.Deduped || refired.Record.Envelope.RelayID != first.Record.Envelope.RelayID || refired.Record.Envelope.IntakeID != first.Record.Envelope.IntakeID {
		t.Fatalf("crash refire = %+v, want original %+v", refired, first)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	count := 0
	for _, rec := range records {
		if rec.Headers["record_kind"] == "resummon_command" && rec.Headers["subject_ref"] == "timer-gate" {
			count++
		}
	}
	if count != 1 {
		t.Fatalf("resummon command count = %d, want 1", count)
	}
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("Build tables: %v", err)
	}
	if got := tab.ContentHash[first.ContentHash]; got != first.Record.Envelope.IntakeID {
		t.Fatalf("A-2 content hash maps to %q, want %q", got, first.Record.Envelope.IntakeID)
	}
	if state := GateState(tab, "timer-gate"); state != GateResummonDue {
		t.Fatalf("gate state = %q, want %q", state, GateResummonDue)
	}

	answeredSlot := refireInput
	answeredSlot.CadenceSlot = "answered-stalled-0001"
	answered, err := scheduler.Emit(ctx, answeredSlot)
	if err != nil {
		t.Fatalf("answered-but-stalled Emit: %v", err)
	}
	if answered.Deduped || answered.Record.Envelope.RelayID == first.Record.Envelope.RelayID || !strings.Contains(answered.Record.Body, `"reason":"answered_but_stalled"`) {
		t.Fatalf("answered-but-stalled timer = %+v", answered)
	}
}

func TestS10ProductionSchedulerArmsParkedGateAndEmitsExactlyOneResummon(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	for _, rec := range []record.Record{
		{
			Envelope: record.Envelope{RelayID: "live-timer-gate", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "live timer gate", "HUMAN_GATE_REQUIRED": "yes"},
		},
		{
			Envelope: record.Envelope{RelayID: "park-live-timer-gate", From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": GateParkedWaitingHuman, "parks_gate": "live-timer-gate"},
		},
	} {
		if _, err := st.Commit(rec, nil); err != nil {
			t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	loop := New(st, ResummonHandler(nil), TestReady())
	journal, err := intake.Open(st.Root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	writer, err := intake.NewWriter[Outcome](journal, config.EngineConfig{}, TestReady())
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	scheduler, err := NewResummonScheduler(st, writer, resummonTestSnapshot(t, st))
	if err != nil {
		t.Fatalf("NewResummonScheduler: %v", err)
	}
	go loop.Run(ctx)
	go writer.Run(ctx, loop.In)
	cadence := ResummonCadence{NoResponse: 0, AnsweredStalled: time.Hour}
	if err := scheduler.ArmParked(ctx, cadence); err != nil {
		t.Fatalf("ArmParked: %v", err)
	}
	if err := scheduler.ArmParked(ctx, cadence); err != nil {
		t.Fatalf("ArmParked repeat: %v", err)
	}
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		records, err := st.Records()
		if err != nil {
			t.Fatalf("Records: %v", err)
		}
		count := 0
		for _, rec := range records {
			if rec.Headers["record_kind"] == "resummon_command" && rec.Headers["subject_ref"] == "live-timer-gate" {
				count++
			}
		}
		if count == 1 {
			return
		}
		if count > 1 {
			t.Fatalf("resummon command count = %d, want exactly one", count)
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("production-composed scheduler did not emit a resummon command")
}

func TestS11StaleSchemaSuppressesOldCadenceAndReplacementRestartsIdentity(t *testing.T) {
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "held-stale-schema-old", DeliveryState: record.Held},
		Headers:  map[string]string{"failing_edge": "stale_schema", "subject_ref": "old"},
	})
	if !gateUnpreservable(tab, "old") {
		t.Fatal("old decision remained eligible for resummon after stale_schema signal")
	}
	if gateUnpreservable(tab, "replacement") {
		t.Fatal("replacement decision inherited old decision suppression")
	}
	replacement := record.Record{Envelope: record.Envelope{RelayID: "replacement", From: "seat-a"}}
	replacementInputs := g4ResummonInputs(replacement)
	if got := replacementInputs[0].CadenceSlot; got != "g4-no-response-1" {
		t.Fatalf("replacement no-response cadence slot = %q, want g4-no-response-1", got)
	}
	if got := replacementInputs[1].CadenceSlot; got != "g4-answered-stalled-1" {
		t.Fatalf("replacement answered-stalled cadence slot = %q, want g4-answered-stalled-1", got)
	}
	oldKey := ResummonContentHash(ResummonInput{Seat: "seat-a", DecisionID: "old", CadenceSlot: "g4-no-response-1"})
	replacementKey := ResummonContentHash(replacementInputs[0])
	if oldKey == replacementKey {
		t.Fatal("replacement decision did not restart cadence under its new identity")
	}
}

func resummonTestSnapshot(t *testing.T, st *store.Store) func() *tables.T {
	t.Helper()
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("Build scheduler snapshot: %v", err)
	}
	live := tables.NewLive(tab)
	return live.Snapshot
}
