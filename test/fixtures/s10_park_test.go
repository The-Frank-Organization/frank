package fixtures_test

import (
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestS10OnlyAcceptedAGateEmitsODBAndParksWaitingForHuman(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitS10ParkRecord(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "non-a-gate",
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"SUBJECT":       "bounded sequencing decision",
			"gate_category": "sequencing",
		},
	})
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto non-A: %v", err)
	}
	assertS10RelayCount(t, st, "park-non-a-gate", 0)
	assertS10RelayCount(t, st, "odb-non-a-gate", 0)

	commitS10ParkRecord(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "a-gate",
			DispatchID:    "s10-build",
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":               "SITREP",
			"SUBJECT":             "operator authorization required",
			"HUMAN_GATE_REQUIRED": "yes",
			"gate_category":       "authz_security",
		},
	})
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto A: %v", err)
	}

	odb := s10RecordByRelay(t, st, "odb-a-gate")
	if odb.Envelope.DeliveryState != record.Accepted || odb.Headers["record_kind"] != "odb" || odb.Headers["subject_ref"] != "a-gate" {
		t.Fatalf("ODB = %+v, want accepted ODB for a-gate", odb)
	}
	if violation := engine.ValidateODBChoice(odb, "approve"); violation != nil {
		t.Fatalf("emitted ODB approve choice: %+v", violation)
	}
	if violation := engine.ValidateODBChoice(odb, "outside-frozen-set"); violation == nil {
		t.Fatalf("emitted ODB accepted a choice outside its frozen set")
	}
	park := s10RecordByRelay(t, st, "park-a-gate")
	if park.Envelope.DeliveryState != record.Accepted || park.Headers["SUBJECT"] != "parked_waiting_human" || park.Headers["parks_gate"] != "a-gate" {
		t.Fatalf("park = %+v, want accepted parked_waiting_human state", park)
	}
	tables, err := obligation.BuildTables(st)
	if err != nil {
		t.Fatalf("BuildTables: %v", err)
	}
	if !tables.ParkedLanes["a-gate"] || tables.ParkedLanes["non-a-gate"] {
		t.Fatalf("parked lanes = %+v, want only a-gate", tables.ParkedLanes)
	}
}

func commitS10ParkRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func s10RecordByRelay(t *testing.T, st *store.Store, relayID string) record.Record {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	for _, rec := range records {
		if rec.Envelope.RelayID == relayID {
			return rec
		}
	}
	t.Fatalf("record %s missing", relayID)
	return record.Record{}
}

func assertS10RelayCount(t *testing.T, st *store.Store, relayID string, want int) {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	got := 0
	for _, rec := range records {
		if rec.Envelope.RelayID == relayID {
			got++
		}
	}
	if got != want {
		t.Fatalf("record %s count = %d, want %d", relayID, got, want)
	}
}
