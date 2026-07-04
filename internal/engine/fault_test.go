package engine_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestAuthorityBearingFaultCommitsHeldAndLoopContinues(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	var calls int
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		calls++
		if calls == 1 {
			panic("validator fault")
		}
		return record.Record{
			Envelope: record.Envelope{RelayID: "after-fault", DeliveryState: record.Accepted, IntakeID: "i2", SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "after"},
		}, nil, nil
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	candidate := record.Record{Headers: map[string]string{"PHASE": "PLAN", "SUBJECT": "authority"}}
	payload, _ := json.Marshal(candidate)
	reply1 := make(chan engine.Outcome, 1)
	reply2 := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "i1", Seat: "s1.orchestrator-planner", Payload: payload}, ReplyCh: reply1}
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "i2", Seat: "s1-core.implementer"}, ReplyCh: reply2}

	if out := <-reply1; out.State != record.Held {
		t.Fatalf("fault outcome = %+v, want held", out)
	}
	if out := <-reply2; out.State != record.Accepted || out.RelayID != "after-fault" {
		t.Fatalf("post-fault outcome = %+v", out)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var held int
	for _, rec := range records {
		if rec.Envelope.DeliveryState == record.Held && rec.Envelope.IntakeID == "i1" {
			held++
		}
	}
	if held != 1 {
		t.Fatalf("held records for i1 = %d, want 1", held)
	}
}

func TestFaultAuthorityClassificationUsesCommandSeatRole(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		panic("validator fault")
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	candidate := record.Record{
		Envelope: record.Envelope{Role: "orchestrator-planner"},
		Headers:  map[string]string{"PHASE": "AUDIT", "SUBJECT": "payload role must not classify"},
	}
	payload, _ := json.Marshal(candidate)
	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "i-role", Seat: "s1-core.implementer", Role: "implementer", Payload: payload}, ReplyCh: reply}

	if out := <-reply; out.State != record.Rejected {
		t.Fatalf("fault outcome = %+v, want rejected for implementer seat", out)
	}
}
