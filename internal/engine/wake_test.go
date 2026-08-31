package engine

import (
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/observe"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

func TestS10WakeReobservesParkedGateBeforeResuming(t *testing.T) {
	tab, gate := s10WakeTable(t)
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	candidate := record.Record{
		Headers: map[string]string{
			"PARENT_DISPATCH_ID": "wake-gate",
			"resolves_gate":      "wake-gate",
		},
		Body: `{"choice":"approve"}`,
	}

	var observedRelay string
	passing := observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(candidate observe.Candidate) observe.PredicateResult {
			observedRelay = candidate.Record.Envelope.RelayID
			return observe.PredicateResult{ID: "fresh-done", Predicate: observe.Pass}
		},
	}
	resumed := classifyVerdict(tab, candidate, meta, passing)
	if observedRelay != gate.Envelope.RelayID {
		t.Fatalf("re-observed relay = %q, want parked gate %q", observedRelay, gate.Envelope.RelayID)
	}
	if resumed.Envelope.DeliveryState != record.Accepted || resumed.Envelope.To != "seat-a" {
		t.Fatalf("fresh wake = %+v, want accepted resume to seat-a", resumed)
	}

	stale := observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate: func(candidate observe.Candidate) observe.PredicateResult {
			return observe.PredicateResult{ID: "fresh-done", Predicate: observe.Fail}
		},
	}
	blocked := classifyVerdict(tab, candidate, meta, stale)
	if blocked.Envelope.DeliveryState != record.Rejected || blocked.Envelope.To != "" || !strings.Contains(blocked.Body, "fresh-done:observed-false") {
		t.Fatalf("stale wake = %+v, want rejected without delivery", blocked)
	}
	if !tab.ParkedLanes[gate.Envelope.RelayID] {
		t.Fatalf("stale wake removed durable park: %+v", tab.ParkedLanes)
	}
}

func s10WakeTable(t *testing.T) (*tables.T, record.Record) {
	t.Helper()
	gate := record.Record{
		Envelope: record.Envelope{
			RelayID:       "wake-gate",
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":               "SITREP",
			"AUTHORITY":           "report-only",
			"HUMAN_GATE_REQUIRED": "yes",
		},
	}
	odb, err := RenderODB(ODBInput{
		SubjectRef:          gate.Envelope.RelayID,
		PlainLanguageChange: "Resume the parked lane.",
		WhyNow:              "An operator reply is available.",
		CompletedProof:      "accepted:wake-gate",
		RecordIntegrity:     "observed",
		TradeoffsRisks:      "The ground truth may have changed while parked.",
		Recommendation:      "Resume only after a fresh observe pass.",
		Choices:             []ODBChoice{{Value: "approve", Label: "Approve"}, {Value: "reject", Label: "Reject"}},
	})
	if err != nil {
		t.Fatalf("RenderODB: %v", err)
	}
	tab := tables.New()
	tab.OnCommit(gate)
	tab.OnCommit(odb)
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "park-wake-gate", From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"parks_gate": gate.Envelope.RelayID},
	})
	return tab, gate
}
