package tables_test

import (
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestIncrementalMatchesRebuild(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	inc, err := tables.Build(st)
	if err != nil {
		t.Fatalf("initial Build: %v", err)
	}
	for _, rec := range []record.Record{
		{Envelope: record.Envelope{RelayID: "r1", DispatchID: "d1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: "i1", SchemaVersion: 1}, Headers: map[string]string{"PHASE": "SITREP", "HUMAN_GATE_REQUIRED": "yes"}},
		{Envelope: record.Envelope{RelayID: "park-r1", From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"parks_gate": "r1"}},
		{Envelope: record.Envelope{RelayID: "owed", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"record_kind": "owed_disposition", "disposes_owed": "owed-base"}},
	} {
		if _, err := st.Commit(rec, nil); err != nil {
			t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
		}
		inc.OnCommit(rec)
	}
	rebuilt, err := tables.Build(st)
	if err != nil {
		t.Fatalf("rebuilt Build: %v", err)
	}
	if !reflect.DeepEqual(inc, rebuilt) {
		t.Fatalf("incremental tables differ from rebuild\nincremental=%#v\nrebuilt=%#v", inc, rebuilt)
	}
}
