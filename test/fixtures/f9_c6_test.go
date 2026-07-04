package fixtures_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestF9RunWholeAndC6ReplayDedupe(t *testing.T) {
	root := t.TempDir()
	j, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}

	var ids []string
	for i := 0; i < 5; i++ {
		id, err := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage{byte('0' + i)}})
		if err != nil {
			t.Fatalf("Append %d: %v", i, err)
		}
		ids = append(ids, id)
	}
	duplicate, err := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage{'0'}})
	if err != nil {
		t.Fatalf("Append duplicate: %v", err)
	}
	if duplicate != ids[0] {
		t.Fatalf("duplicate id = %q, want %q", duplicate, ids[0])
	}
	for i := 0; i < 2; i++ {
		if _, err := st.Commit(record.Record{
			Envelope: record.Envelope{
				RelayID:       "outcome-" + ids[i],
				DispatchID:    "d",
				From:          "seat-a",
				Role:          "implementer",
				DeliveryState: record.Accepted,
				IntakeID:      ids[i],
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": ids[i]},
		}, nil); err != nil {
			t.Fatalf("commit outcome %d: %v", i, err)
		}
	}
	unconsumed, err := intake.Unconsumed(context.Background(), j, st)
	if err != nil {
		t.Fatalf("Unconsumed: %v", err)
	}
	if len(unconsumed) != 3 {
		t.Fatalf("unconsumed len = %d, want 3: %+v", len(unconsumed), unconsumed)
	}
	for i, cmd := range unconsumed {
		if cmd.IntakeID != ids[i+2] {
			t.Fatalf("unconsumed[%d] = %s, want %s", i, cmd.IntakeID, ids[i+2])
		}
	}
}
