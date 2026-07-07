package engine_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestLoopProcessesFIFOAndRepliesAfterCommit(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	var order []string
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		order = append(order, string([]byte(orderName(len(order)))))
		id := ""
		return record.Record{
			Envelope: record.Envelope{
				RelayID:       id,
				DispatchID:    "d",
				From:          "seat-a",
				Role:          "implementer",
				DeliveryState: record.Accepted,
				IntakeID:      cmd.IntakeID,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": id},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	reply1 := make(chan engine.Outcome, 1)
	reply2 := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "intake-1", Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`1`)}, ReplyCh: reply1}
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "intake-2", Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`2`)}, ReplyCh: reply2}

	out1 := <-reply1
	out2 := <-reply2
	if out1.State != record.Accepted || out2.State != record.Accepted {
		t.Fatalf("outcomes = %+v %+v", out1, out2)
	}
	if out1.RelayID == "" || out2.RelayID == "" {
		t.Fatalf("missing relay ids: %+v %+v", out1, out2)
	}
	if len(order) != 2 || order[0] != "1" || order[1] != "2" {
		t.Fatalf("order = %v, want [1 2]", order)
	}
}

func TestLoopCompletesObligationsOnSerializedTurn(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{
				RelayID:       "gate-from-loop",
				From:          "seat-a",
				Role:          "implementer",
				DeliveryState: record.Accepted,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes"},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{}`)}, ReplyCh: reply}
	out := <-reply
	if out.State != record.Accepted {
		t.Fatalf("outcome = %+v, want accepted", out)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var parked, outbox bool
	for _, rec := range records {
		if rec.Headers["parks_gate"] == "gate-from-loop" {
			parked = true
		}
		if rec.Envelope.RelayID == "outbox-gate-gate-from-loop" {
			outbox = true
		}
	}
	if !parked || !outbox {
		t.Fatalf("serialized obligation turn parked=%v outbox=%v records=%+v", parked, outbox, records)
	}
}

func TestLoopReplaysExistingOutcomeForDuplicateIntake(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	var calls int
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		calls++
		return record.Record{
			Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "duplicate"},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	cmd := intake.Cmd{IntakeID: "intake-duplicate", Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"same":true}`)}
	firstReply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: cmd, ReplyCh: firstReply}
	first := <-firstReply
	secondReply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: cmd, ReplyCh: secondReply}
	second := <-secondReply

	if calls != 1 {
		t.Fatalf("handler calls = %d, want one execution", calls)
	}
	if first != second {
		t.Fatalf("duplicate outcome = %+v, want original %+v", second, first)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var outcomes int
	for _, rec := range records {
		if rec.Envelope.IntakeID == cmd.IntakeID {
			outcomes++
		}
	}
	if outcomes != 1 {
		t.Fatalf("outcomes for %s = %d, want 1", cmd.IntakeID, outcomes)
	}
}

func TestLoopReplaysExistingOutcomeForDuplicateContentHash(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	originalID, err := journal.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"same":true}`), ContentHash: "same-hash"})
	if err != nil {
		t.Fatalf("Append original: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "original-outcome", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: originalID, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "original"},
	}, nil); err != nil {
		t.Fatalf("Commit original outcome: %v", err)
	}
	var calls int
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		calls++
		return record.Record{
			Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "should not execute"},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "intake-retry", Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"same":true}`), ContentHash: "same-hash"}, ReplyCh: reply}
	out := <-reply
	if calls != 0 {
		t.Fatalf("handler calls = %d, want replay without execution", calls)
	}
	if out.RelayID != "original-outcome" || out.IntakeID != originalID || out.State != record.Accepted {
		t.Fatalf("outcome = %+v, want original outcome", out)
	}
}

func TestLoopOutcomeDetailEqualsRecordedRejectionDetailPerClass(t *testing.T) {
	cases := []struct {
		name       string
		meta       seat.SeatMeta
		candidate  record.Record
		digest     string
		setupStore func(*testing.T, *store.Store)
	}{
		{
			name: "re-render",
			meta: seat.SeatMeta{Name: "seat-a", Role: "implementer"},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "stale digest",
			}},
			digest: "stale",
		},
		{
			name: "required",
			meta: seat.SeatMeta{Name: "seat-a", Role: "implementer"},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
			}},
		},
		{
			name: "enum",
			meta: seat.SeatMeta{Name: "seat-a", Role: "implementer"},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "not-authority",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "bad enum",
			}},
		},
		{
			name: "seat-scope",
			meta: seat.SeatMeta{Name: "seat-a", Role: "implementer"},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "bad scope",
				"record_kind":     "genesis",
			}},
		},
		{
			name: "canonical-encoding",
			meta: seat.SeatMeta{Name: "seat-a", Role: "implementer"},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "bad canonical",
				"SCOPE_DIFF":      `[{ "path":"README.md","status":"in"}]`,
			}},
		},
		{
			name: "lineage",
			meta: seat.SeatMeta{Name: "seat-a.planner", Role: "planner"},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":              "PLAN",
				"AUTHORITY":          "plan-only",
				"CEREMONY_TIER":      "medium",
				"EVIDENCE_TARGET":    "E1",
				"SUBJECT":            "bad design chain",
				"DESIGN_LOCK_ID":     "design-1",
				"DESIGN_RECORD_KIND": "design-doc",
			}},
		},
		{
			name: "layer-3",
			meta: seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "unknown owed",
				"record_kind":     "owed_disposition",
				"disposes_owed":   "owed-missing",
			}},
		},
		{
			name: "already-resolved",
			meta: seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true},
			setupStore: func(t *testing.T, st *store.Store) {
				t.Helper()
				if _, err := st.Commit(record.Record{
					Envelope: record.Envelope{RelayID: "owed-target", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
					Headers: map[string]string{
						"PHASE":            "SITREP",
						"SUBJECT":          "owed",
						"record_kind":      "owed_item",
						"owner":            "s6",
						"source":           "test",
						"target_surface":   "detail",
						"disposition_path": "done",
					},
				}, nil); err != nil {
					t.Fatalf("commit owed: %v", err)
				}
				if _, err := st.Commit(record.Record{
					Envelope: record.Envelope{RelayID: "owed-disposed", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
					Headers: map[string]string{
						"PHASE":         "SITREP",
						"SUBJECT":       "disposed",
						"record_kind":   "owed_disposition",
						"disposes_owed": "owed-target",
					},
				}, nil); err != nil {
					t.Fatalf("commit disposition: %v", err)
				}
			},
			candidate: record.Record{Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "duplicate disposition",
				"record_kind":     "owed_disposition",
				"disposes_owed":   "owed-target",
			}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			st, reg := submitDeps(t)
			if tc.setupStore != nil {
				tc.setupStore(t, st)
			}
			payload := submitPayload(t, reg, tc.meta, tc.candidate)
			if tc.digest != "" {
				var submit fieldspec.SubmitPayload
				if err := json.Unmarshal(payload, &submit); err != nil {
					t.Fatalf("decode submit payload: %v", err)
				}
				submit.FormDigest = tc.digest
				payload = mustJSON(t, submit)
			}
			loop := engine.New(st, engine.SubmitHandler(st, reg, tc.meta), engine.TestReady())
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			go loop.Run(ctx)

			reply := make(chan engine.Outcome, 1)
			loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "detail-" + tc.name, Seat: tc.meta.Name, Role: tc.meta.Role, IsOperator: tc.meta.IsOperator, Verb: "submit", Payload: payload}, ReplyCh: reply}
			out := <-reply
			if out.State != record.Rejected || out.RelayID == "" {
				t.Fatalf("outcome = %+v, want rejected relay", out)
			}
			rec, err := st.Read(out.RelayID)
			if err != nil {
				t.Fatalf("read rejected relay: %v", err)
			}
			if out.Detail != rec.Body {
				t.Fatalf("detail = %q, want committed body %q", out.Detail, rec.Body)
			}
		})
	}
}

func TestLoopProcessesQueuedQuarantineWhileIdle(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "bad-record", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "bad"},
	}, nil); err != nil {
		t.Fatalf("Commit: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "records", "bad-record.json"), []byte(`{"bad":true}`), 0o644); err != nil {
		t.Fatalf("corrupt record: %v", err)
	}
	loop := engine.New(st, nil, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	loop.EnqueueQuarantine("bad-record")
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(filepath.Join(root, "quarantine", "bad-record.json")); err == nil {
			if _, err := os.Stat(filepath.Join(root, "records", "incident-bad-record.json")); err == nil {
				return
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("queued quarantine was not processed while loop was idle")
}

func orderName(i int) string {
	if i == 0 {
		return "1"
	}
	return "2"
}
