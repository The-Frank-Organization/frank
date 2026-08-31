package fixtures_test

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS5SubmitGuardRejectsSeatSuppliedSystemHeaders(t *testing.T) {
	for _, meta := range []seat.SeatMeta{
		{Name: "s5-b.implementer", Role: "implementer"},
		{Name: "operator", Role: "operator", IsOperator: true},
	} {
		for _, field := range []string{"failing_edge", "gate_category_raised", "delivery_state"} {
			t.Run(meta.Name+"/"+field, func(t *testing.T) {
				st, err := store.Open(t.TempDir())
				if err != nil {
					t.Fatalf("Open store: %v", err)
				}
				reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
				if err != nil {
					t.Fatalf("Load registry: %v", err)
				}
				rec := record.Record{Headers: map[string]string{
					"PHASE":           "SITREP",
					"AUTHORITY":       "report-only",
					"CEREMONY_TIER":   "medium",
					"EVIDENCE_TARGET": "E1",
					"SUBJECT":         "forged " + field,
				}}
				rec.Headers[field] = forgedSystemHeaderValue(field)
				payload := mustJSONBytes(t, submitPayloadForRegistry(reg, meta, rec))
				got, _, err := engine.SubmitHandler(st, reg, meta)(context.Background(), intake.Cmd{
					IntakeID:   "s5-guard-" + field,
					Seat:       meta.Name,
					Role:       meta.Role,
					IsOperator: meta.IsOperator,
					Payload:    payload,
				})
				if err != nil {
					t.Fatalf("SubmitHandler: %v", err)
				}
				if got.Envelope.DeliveryState != record.Rejected {
					t.Fatalf("state = %s, want rejected", got.Envelope.DeliveryState)
				}
				want := field + ":system-owned"
				if !strings.Contains(got.Body, want) {
					t.Fatalf("bounce body = %q, want %q", got.Body, want)
				}
			})
		}
	}
}

func TestS5SubmitGuardDoesNotBlockDerivedSystemWriters(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "gate-derived",
			From:          "s5-b.implementer",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":               "SITREP",
			"SUBJECT":             "derived gate",
			"HUMAN_GATE_REQUIRED": "yes",
			"gate_category":       "authz_security",
		},
	}, nil); err != nil {
		t.Fatalf("Commit gate: %v", err)
	}

	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}
	if _, err := st.Read("park-gate-derived"); err != nil {
		t.Fatalf("Read derived park record: %v", err)
	}
	if _, err := st.Read("outbox-gate-gate-derived"); err != nil {
		t.Fatalf("Read derived outbox record: %v", err)
	}
}

func forgedSystemHeaderValue(field string) string {
	switch field {
	case "gate_category_raised":
		return "yes"
	case "delivery_state":
		return record.Accepted
	default:
		return "forged"
	}
}
