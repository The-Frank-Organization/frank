package tables_test

import (
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/tables"
)

func TestActivationRecordRefSetExactlyOnce(t *testing.T) {
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "mint-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "seat_mint"},
		Body:     `{"seat":"seat-a.implementer","role":"implementer","is_operator":false}`,
	})
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "boot-1", From: "seat-a.implementer", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "boot", "charter_loaded": "yes", "dispatch_status": "read"},
	})
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "ordinary-2", From: "seat-a.implementer", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "ordinary"},
	})

	state := tab.Lifecycle["seat-a.implementer"]
	if state.ActivationState != tables.LifecycleActive {
		t.Fatalf("ActivationState = %q, want active", state.ActivationState)
	}
	if state.ActivationRecordRef != "boot-1" {
		t.Fatalf("ActivationRecordRef = %q, want first accepted boot", state.ActivationRecordRef)
	}
	if state.LastAcceptedAt != "ordinary-2" {
		t.Fatalf("LastAcceptedAt = %q, want latest accepted record", state.LastAcceptedAt)
	}
}

func TestAlreadyActiveBootShapedRecordDoesNotMoveActivationEdge(t *testing.T) {
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "mint-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "seat_mint"},
		Body:     `{"seat":"seat-a.implementer","role":"implementer","is_operator":false}`,
	})
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "boot-1", From: "seat-a.implementer", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "boot", "charter_loaded": "yes", "dispatch_status": "read"},
	})
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "boot-shaped-2", From: "seat-a.implementer", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "ordinary boot-shaped", "charter_loaded": "yes", "dispatch_status": "read"},
	})

	state := tab.Lifecycle["seat-a.implementer"]
	if state.ActivationRecordRef != "boot-1" {
		t.Fatalf("ActivationRecordRef = %q, want first activation record", state.ActivationRecordRef)
	}
	if state.LastAcceptedAt != "boot-shaped-2" {
		t.Fatalf("LastAcceptedAt = %q, want latest accepted boot-shaped record", state.LastAcceptedAt)
	}
}
