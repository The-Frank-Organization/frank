package fixtures_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

// Step-1 ③ detection claims exactly S1 + S2 + S3 + fail-safe.
// It does not claim to catch every content mis-pick: if a B pick is invisible
// to those sources and is not the fail-safe "other" token, it commits as B.
func TestS5GateRaiseRewritesDetectorHitAndCompletesGateConsequences(t *testing.T) {
	for _, tc := range []struct {
		name     string
		pick     string
		wantPick string
	}{
		{name: "B-pick", pick: "routing", wantPick: "routing"},
		{name: "B-absorb"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			st, reg, meta := s5GateRaiseDeps(t)
			rec := s5GateRaiseCandidate()
			if tc.pick != "" {
				rec.Headers["gate_category"] = tc.pick
			}
			committed := s5SubmitAndCommit(t, st, reg, meta, s5AFloorEnv(reg, "authz_security"), rec)
			if committed.Envelope.DeliveryState != record.Accepted {
				t.Fatalf("state = %s, body=%s", committed.Envelope.DeliveryState, committed.Body)
			}
			if committed.Headers["gate_category"] != "authz_security" {
				t.Fatalf("gate_category = %q, want authz_security", committed.Headers["gate_category"])
			}
			if committed.Headers["gate_category_raised"] != "yes" {
				t.Fatalf("gate_category_raised = %q, want yes", committed.Headers["gate_category_raised"])
			}
			if got := committed.Headers["gate_category_pick"]; got != tc.wantPick {
				t.Fatalf("gate_category_pick = %q, want %q", got, tc.wantPick)
			}
			if err := obligation.CompleteAuto(st); err != nil {
				t.Fatalf("CompleteAuto: %v", err)
			}
			if _, err := st.Read("park-" + committed.Envelope.RelayID); err != nil {
				t.Fatalf("missing park record: %v", err)
			}
			items := s5ReadOutboxItems(t, st.Root)
			if len(items) != 1 || items[0].SourceRecordRef != committed.Envelope.RelayID || items[0].GateCategory != "authz_security" {
				t.Fatalf("outbox items = %+v", items)
			}
		})
	}
}

func TestS5GateRaiseDoesNotLowerAPick(t *testing.T) {
	st, reg, meta := s5GateRaiseDeps(t)
	rec := s5GateRaiseCandidate()
	rec.Headers["gate_category"] = "authz_security"

	committed := s5SubmitAndCommit(t, st, reg, meta, s5AFloorEnv(reg, "product_semantics"), rec)
	if committed.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", committed.Envelope.DeliveryState, committed.Body)
	}
	if committed.Headers["gate_category"] != "authz_security" {
		t.Fatalf("gate_category = %q, want unchanged A pick", committed.Headers["gate_category"])
	}
	if committed.Headers["gate_category_raised"] != "" || committed.Headers["gate_category_pick"] != "" {
		t.Fatalf("A pick unexpectedly stamped: %+v", committed.Headers)
	}
}

func TestS5GateRaiseOtherUsesYesByteAndRegistryToken(t *testing.T) {
	st, reg, meta := s5GateRaiseDeps(t)
	rec := s5GateRaiseCandidate()
	rec.Headers["gate_category"] = "other"

	committed := s5SubmitAndCommit(t, st, reg, meta, fieldspec.RenderEnv{}, rec)
	if committed.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", committed.Envelope.DeliveryState, committed.Body)
	}
	if committed.Headers["gate_category"] != "other" || committed.Headers["gate_category_raised"] != "yes" {
		t.Fatalf("other pick headers = %+v", committed.Headers)
	}
	if class, raised := reg.ClassifyGateCategory(committed.Headers["gate_category"], false); class != "A" || !raised {
		t.Fatalf("committed other class=(%s,%v), want A raised", class, raised)
	}
	if !committedBoolTokenAllowed(t, reg, committed.Headers["gate_category_raised"]) {
		t.Fatalf("gate_category_raised byte %q is not in registry bool enum", committed.Headers["gate_category_raised"])
	}
}

func TestS5GateRaiseClaimBoundaryInvisibleBPickCommitsAsB(t *testing.T) {
	st, reg, meta := s5GateRaiseDeps(t)
	rec := s5GateRaiseCandidate()
	rec.Headers["gate_category"] = "routing"

	committed := s5SubmitAndCommit(t, st, reg, meta, fieldspec.RenderEnv{}, rec)
	if committed.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", committed.Envelope.DeliveryState, committed.Body)
	}
	if committed.Headers["gate_category"] != "routing" {
		t.Fatalf("gate_category = %q, want boundary-preserved B pick", committed.Headers["gate_category"])
	}
	if committed.Headers["gate_category_raised"] != "" || committed.Headers["gate_category_pick"] != "" {
		t.Fatalf("invisible B pick unexpectedly stamped: %+v", committed.Headers)
	}
}

func s5GateRaiseDeps(t *testing.T) (*store.Store, *fieldspec.Registry, seat.SeatMeta) {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return st, reg, seat.SeatMeta{Name: "s5-b.implementer", Role: "implementer"}
}

func s5GateRaiseCandidate() record.Record {
	return record.Record{Headers: map[string]string{
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "medium",
		"EVIDENCE_TARGET": "E1",
		"SUBJECT":         "gate raise",
	}}
}

func s5AFloorEnv(reg *fieldspec.Registry, member string) fieldspec.RenderEnv {
	return fieldspec.RenderEnv{KnownA: engine.KnownADetector(reg, nil, engine.DetectorConfig{
		AFloor: map[engine.DetectorKey]string{{Phase: "SITREP"}: member},
	})}
}

func s5SubmitAndCommit(t *testing.T, st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, rec record.Record) record.Record {
	t.Helper()
	_, digest := reg.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	payload := mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	got, intents, err := engine.SubmitHandlerWithRender(st, reg, meta, env)(context.Background(), intake.Cmd{
		IntakeID: "s5-gate-raise",
		Seat:     meta.Name,
		Role:     meta.Role,
		Payload:  payload,
	})
	if err != nil {
		t.Fatalf("SubmitHandlerWithRender: %v", err)
	}
	if _, err := st.Commit(got, intents); err != nil {
		t.Fatalf("Commit: %v", err)
	}
	return got
}

func s5ReadOutboxItems(t *testing.T, root string) []gate.OutboxItem {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(root, "outbox"))
	if err != nil {
		t.Fatalf("ReadDir outbox: %v", err)
	}
	items := make([]gate.OutboxItem, 0, len(entries))
	for _, entry := range entries {
		data, err := os.ReadFile(filepath.Join(root, "outbox", entry.Name()))
		if err != nil {
			t.Fatalf("read outbox item: %v", err)
		}
		var item gate.OutboxItem
		if err := json.Unmarshal(data, &item); err != nil {
			t.Fatalf("decode outbox item: %v", err)
		}
		items = append(items, item)
	}
	return items
}

func committedBoolTokenAllowed(t *testing.T, reg *fieldspec.Registry, value string) bool {
	t.Helper()
	spec, ok := reg.ByID("gate_category_raised")
	if !ok {
		t.Fatalf("gate_category_raised spec missing")
	}
	form, _ := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "SITREP", "medium", fieldspec.ClosedGrantState)
	if form.HasField("gate_category_raised") {
		t.Fatalf("computed gate_category_raised unexpectedly renderable")
	}
	for _, token := range reg.NamedEnums[spec.EnumSet] {
		if token == value {
			return true
		}
	}
	return false
}
