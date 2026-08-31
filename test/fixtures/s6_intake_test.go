package fixtures_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS6IntakeDuplicateReplaysOriginalOutcomeByteIdentical(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	writer, err := intake.NewWriter[engine.Outcome](journal, config.EngineConfig{}, engine.TestReady())
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	var executions int
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		executions++
		return record.Record{
			Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "duplicate"},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan intake.Job[engine.Outcome], 8)
	go writer.Run(ctx, out)
	loop.In = out
	go loop.Run(ctx)

	payload, _ := json.Marshal(record.Record{Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "duplicate"}})
	cmd := intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: payload}
	firstReply, _, err := writer.Submit(ctx, cmd)
	if err != nil {
		t.Fatalf("first Submit: %v", err)
	}
	first := s6WaitOutcome(t, firstReply)
	secondReply, _, err := writer.Submit(ctx, cmd)
	if err != nil {
		t.Fatalf("second Submit: %v", err)
	}
	second := s6WaitOutcome(t, secondReply)

	if executions != 1 {
		t.Fatalf("executions = %d, want one", executions)
	}
	if first != second {
		t.Fatalf("second outcome = %+v, want original %+v", second, first)
	}
}

func s6WaitOutcome(t *testing.T, reply <-chan engine.Outcome) engine.Outcome {
	t.Helper()
	select {
	case out := <-reply:
		return out
	case <-time.After(2 * time.Second):
		t.Fatalf("timed out waiting for outcome")
		return engine.Outcome{}
	}
}
