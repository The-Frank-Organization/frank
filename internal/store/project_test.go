package store_test

import (
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestProjectScopesMailboxAndReadReturnsCommittedRecord(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-a", DispatchID: "d", From: "seat-a", To: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "a"},
	}, []store.Intent{{Kind: store.IntentMailbox, Path: "seat-a.jsonl", Payload: []byte("relay-a\n")}}); err != nil {
		t.Fatalf("commit a: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-b", DispatchID: "d", From: "seat-b", To: "seat-b", Role: "planner", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "b"},
	}, []store.Intent{{Kind: store.IntentMailbox, Path: "seat-b.jsonl", Payload: []byte("relay-b\n")}}); err != nil {
		t.Fatalf("commit b: %v", err)
	}

	project, err := st.Project("seat-a")
	if err != nil {
		t.Fatalf("Project: %v", err)
	}
	if len(project) != 1 || project[0] != "relay-a" {
		t.Fatalf("project seat-a = %v, want [relay-a]", project)
	}
	rec, err := st.Read("relay-b")
	if err != nil {
		t.Fatalf("Read relay-b: %v", err)
	}
	if rec.Envelope.RelayID != "relay-b" || rec.Headers["SUBJECT"] != "b" {
		t.Fatalf("read record = %+v", rec)
	}
}
