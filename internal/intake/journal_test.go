package intake_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestAppendAssignsIDAndDedupesContentHash(t *testing.T) {
	j, err := intake.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	cmd := intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"subject":"one"}`)}
	first, err := j.Append(cmd)
	if err != nil {
		t.Fatalf("Append first: %v", err)
	}
	second, err := j.Append(cmd)
	if err != nil {
		t.Fatalf("Append duplicate: %v", err)
	}
	if first != second {
		t.Fatalf("duplicate content got new intake id %q, want %q", second, first)
	}
	entries, err := j.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(entries))
	}
	if entries[0].IntakeID == "" || entries[0].ContentHash == "" {
		t.Fatalf("entry missing intake id or content hash: %+v", entries[0])
	}
}

func TestUnconsumedSkipsExistingOutcomes(t *testing.T) {
	root := t.TempDir()
	j, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	first, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":1}`)})
	second, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":2}`)})
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "outcome-1",
			DispatchID:    "d",
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			IntakeID:      first,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "one"},
	}, nil); err != nil {
		t.Fatalf("commit outcome: %v", err)
	}
	got, err := intake.Unconsumed(context.Background(), j, st)
	if err != nil {
		t.Fatalf("Unconsumed: %v", err)
	}
	if len(got) != 1 || got[0].IntakeID != second {
		t.Fatalf("unconsumed = %+v, want only %s", got, second)
	}
}
