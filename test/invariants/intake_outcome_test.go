package invariants_test

import (
	"context"
	"encoding/json"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestLawIntakeOutcomeOneToOne(t *testing.T) {
	_, law := requireCatalogLaw(t, "TestLawIntakeOutcomeOneToOne")
	for _, required := range []string{
		"at most one outcome per intake at all times",
		"settled means exactly one",
		"pending-zero is valid",
		"recovery and duplicate retries do not double-emit",
	} {
		if !strings.Contains(law.Claim, required) {
			t.Fatalf("row-9 claim %q missing %q", law.Claim, required)
		}
	}
	assertRecoveryReenqueueSelection(t)

	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open intake store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("open intake journal: %v", err)
	}
	ready := engine.TestReady()
	writer, err := intake.NewWriter[engine.Outcome](journal, config.EngineConfig{}, ready)
	if err != nil {
		t.Fatalf("new intake writer: %v", err)
	}
	var executions atomic.Int32
	firstStarted := make(chan struct{}, 1)
	releaseFirst := make(chan struct{})
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		count := executions.Add(1)
		if count == 1 {
			firstStarted <- struct{}{}
			<-releaseFirst
		}
		return record.Record{
			Envelope: record.Envelope{
				From:          cmd.Seat,
				Role:          cmd.Role,
				DeliveryState: record.Accepted,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "intake outcome"},
		}, nil, nil
	}, ready)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)
	go writer.Run(ctx, loop.In)

	duplicate := intake.Cmd{
		Seat:    "s7.implementer",
		Role:    "implementer",
		Verb:    "submit",
		Payload: json.RawMessage(`{"same":true}`),
	}
	firstReply, firstID, err := writer.Submit(ctx, duplicate)
	if err != nil {
		t.Fatalf("submit first duplicate: %v", err)
	}
	select {
	case <-firstStarted:
	case <-time.After(3 * time.Second):
		t.Fatal("first duplicate did not reach handler")
	}
	secondReply, secondID, err := writer.Submit(ctx, duplicate)
	if err != nil {
		t.Fatalf("submit concurrent duplicate: %v", err)
	}
	if firstID == "" || secondID != firstID {
		t.Fatalf("concurrent duplicate intake IDs = %q/%q", firstID, secondID)
	}
	close(releaseFirst)
	first := waitOutcome(t, firstReply)
	second := waitOutcome(t, secondReply)
	if first != second || first.IntakeID != firstID || first.RelayID == "" {
		t.Fatalf("coalesced outcomes = %+v / %+v, intake=%s", first, second, firstID)
	}
	if got := executions.Load(); got != 1 {
		t.Fatalf("concurrent duplicate executions = %d, want 1", got)
	}

	replayReply, replayID, err := writer.Submit(ctx, duplicate)
	if err != nil {
		t.Fatalf("submit completed duplicate retry: %v", err)
	}
	replayed := waitOutcome(t, replayReply)
	if replayID != firstID || replayed != first {
		t.Fatalf("completed duplicate replay = id %q outcome %+v, want %q %+v", replayID, replayed, firstID, first)
	}
	if got := executions.Load(); got != 1 {
		t.Fatalf("completed duplicate re-executed handler %d times", got)
	}

	pendingID, err := journal.Append(intake.Cmd{
		Seat:    "s7.implementer",
		Role:    "implementer",
		Verb:    "submit",
		Payload: json.RawMessage(`{"pending":true}`),
	})
	if err != nil {
		t.Fatalf("append pending intake: %v", err)
	}
	pending, err := intake.Unconsumed(ctx, journal, st)
	if err != nil {
		t.Fatalf("derive pending intake: %v", err)
	}
	if len(pending) != 1 || pending[0].IntakeID != pendingID {
		t.Fatalf("pending intake set = %+v, want only %s", pending, pendingID)
	}
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("build outcome table with pending intake: %v", err)
	}
	if _, exists := tab.OutcomeByIntake[pendingID]; exists {
		t.Fatalf("pending intake %s has an outcome before settlement", pendingID)
	}

	firstRecovery := make(chan engine.Outcome, 1)
	loop.In <- intake.Job[engine.Outcome]{Cmd: pending[0], ReplyCh: firstRecovery}
	recovered := waitOutcome(t, firstRecovery)
	if recovered.IntakeID != pendingID || recovered.RelayID == "" || recovered.State != record.Accepted {
		t.Fatalf("first recovery outcome = %+v", recovered)
	}
	recordsAfterFirstRecovery := mustRecords(t, st)
	if got := executions.Load(); got != 2 {
		t.Fatalf("handler executions after pending recovery = %d, want 2", got)
	}

	secondRecovery := make(chan engine.Outcome, 1)
	loop.In <- intake.Job[engine.Outcome]{Cmd: pending[0], ReplyCh: secondRecovery}
	recoveredAgain := waitOutcome(t, secondRecovery)
	if recoveredAgain != recovered {
		t.Fatalf("second recovery outcome = %+v, want replay %+v", recoveredAgain, recovered)
	}
	recordsAfterSecondRecovery := mustRecords(t, st)
	if len(recordsAfterSecondRecovery) != len(recordsAfterFirstRecovery) {
		t.Fatalf("second recovery emitted another record: before=%d after=%d", len(recordsAfterFirstRecovery), len(recordsAfterSecondRecovery))
	}
	if got := executions.Load(); got != 2 {
		t.Fatalf("second recovery re-executed handler %d times", got)
	}

	if unconsumed, err := intake.Unconsumed(ctx, journal, st); err != nil {
		t.Fatalf("derive final pending set: %v", err)
	} else if len(unconsumed) != 0 {
		t.Fatalf("settled intakes still pending: %+v", unconsumed)
	}
	assertIntakeOutcomeCardinality(t, journal, recordsAfterSecondRecovery)
}

func assertRecoveryReenqueueSelection(t *testing.T) {
	t.Helper()
	root := t.TempDir()
	pinned := initPinnedStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open recovery-selection store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("open recovery-selection journal: %v", err)
	}
	settledID, err := journal.Append(intake.Cmd{
		Seat: "s7.implementer", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"settled":true}`),
	})
	if err != nil {
		t.Fatalf("append settled recovery intake: %v", err)
	}
	pendingID, err := journal.Append(intake.Cmd{
		Seat: "s7.implementer", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"pending":true}`),
	})
	if err != nil {
		t.Fatalf("append pending recovery intake: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "settled-recovery-outcome",
			From:          "s7.implementer",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			IntakeID:      settledID,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "settled recovery outcome"},
	}, nil); err != nil {
		t.Fatalf("commit settled recovery outcome: %v", err)
	}
	if tab, err := tables.Build(st); err != nil {
		t.Fatalf("build recovery-selection table: %v", err)
	} else if _, exists := tab.OutcomeByIntake[pendingID]; exists {
		t.Fatalf("pending recovery intake %s has an outcome", pendingID)
	}

	processorCalls := map[string]int{}
	result, err := frankrecover.RunWithProcessor(root, pinned, func(cmd intake.Cmd) error {
		processorCalls[cmd.IntakeID]++
		return nil
	})
	if err != nil {
		t.Fatalf("run production recovery selection: %v", err)
	}
	if result.Ready == nil || result.Diag != nil {
		t.Fatalf("production recovery result = %+v, want Ready", result)
	}
	if processorCalls[settledID] != 0 || processorCalls[pendingID] != 1 || len(processorCalls) != 1 {
		t.Fatalf("recovery processor calls = %v, want only %s exactly once and %s zero", processorCalls, pendingID, settledID)
	}
}

func mustRecords(t *testing.T, st *store.Store) []record.Record {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("read records: %v", err)
	}
	return records
}

func assertIntakeOutcomeCardinality(t *testing.T, journal *intake.Journal, records []record.Record) {
	t.Helper()
	entries, err := journal.ReadAll()
	if err != nil {
		t.Fatalf("read journal cardinality: %v", err)
	}
	outcomes := map[string]int{}
	for _, rec := range records {
		if rec.Envelope.IntakeID == "" {
			t.Fatalf("outcome %s has empty intake_id", rec.Envelope.RelayID)
		}
		outcomes[rec.Envelope.IntakeID]++
		if outcomes[rec.Envelope.IntakeID] > 1 {
			t.Fatalf("intake %s has %d outcomes", rec.Envelope.IntakeID, outcomes[rec.Envelope.IntakeID])
		}
	}
	seenEntries := map[string]bool{}
	for _, entry := range entries {
		if entry.IntakeID == "" {
			t.Fatal("journal entry has empty intake_id")
		}
		if seenEntries[entry.IntakeID] {
			t.Fatalf("duplicate journal intake_id %s", entry.IntakeID)
		}
		seenEntries[entry.IntakeID] = true
		if outcomes[entry.IntakeID] != 1 {
			t.Fatalf("settled intake %s has %d outcomes, want exactly 1", entry.IntakeID, outcomes[entry.IntakeID])
		}
	}
}
