package engine_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestLoopProcessesFIFOAndRepliesAfterCommit(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	var order []string
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		order = append(order, string([]byte(orderName(len(order)))))
		id := ""
		return record.Record{
			Envelope: record.Envelope{
				RelayID:       id,
				DispatchID:    "d",
				From:          "seat-a",
				Role:          "implementer",
				DeliveryState: record.Accepted,
				IntakeID:      "intake-" + orderName(len(order)),
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
	loop.In <- engine.Job{Cmd: intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`1`)}, ReplyCh: reply1}
	loop.In <- engine.Job{Cmd: intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`2`)}, ReplyCh: reply2}

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
