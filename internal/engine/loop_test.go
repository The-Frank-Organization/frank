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

func TestLoopProcessesFIFOAndRepliesAfterCommit(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	var order []string
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		order = append(order, string([]byte(orderName(len(order)))))
		id := "relay-" + orderName(len(order))
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
	})
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

func orderName(i int) string {
	if i == 0 {
		return "1"
	}
	return "2"
}
