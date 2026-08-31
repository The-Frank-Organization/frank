package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/migrate"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/record"
	frankrecover "github.com/The-Frank-Organization/frank/internal/recover"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS11StaleChoiceRejectsAndCrashReplaysSameReplacementIdentity(t *testing.T) {
	if os.Getenv("FRANK_S11_STALE_REISSUE_CHILD") == "1" {
		st, err := store.Open(os.Getenv("FRANK_S11_STALE_REISSUE_ROOT"))
		if err != nil {
			panic(err)
		}
		if err := obligation.CompleteAuto(st); err != nil {
			panic(err)
		}
		panic("stale reissue crashpoint did not fire")
	}

	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	oldChoices := s11Choices(t,
		map[string]string{"label": "Approve", "value": "approve"},
		map[string]string{"label": "Reject", "value": "reject"},
	)
	gateID := "s11-8a-old-decision"
	s11CommitGateAndODB(t, st, gateID, oldChoices)
	odbPath := filepath.Join(root, "records", "odb-"+gateID+".json")
	before, err := os.ReadFile(odbPath)
	if err != nil {
		t.Fatalf("read source ODB: %v", err)
	}

	changedChoices := s11Choices(t,
		map[string]string{"label": "Ship", "value": "approve"},
		map[string]string{"label": "Reject", "value": "reject"},
	)
	migrations := migrate.New()
	migrations.Current = 2
	migrations.Register(1, func(rec record.Record) (record.Record, error) {
		if rec.Headers["record_kind"] == "odb" {
			rec.Headers["choices"] = changedChoices
		}
		return rec, nil
	})

	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	candidate := s11ResolutionCandidate(gateID, "approve")
	payload := s10ExitPayload(t, pinned.Registry, fieldspec.RenderEnv{}, meta, candidate)
	stale, intents, err := engine.SubmitHandlerWithMigration(st, pinned.Registry, migrations, meta)(context.Background(), intake.Cmd{
		IntakeID: "s11-stale-reply", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: payload,
	})
	if err != nil {
		t.Fatalf("SubmitHandlerWithMigration: %v", err)
	}
	if stale.Envelope.DeliveryState != record.Rejected || stale.Headers["failing_edge"] != "stale_choice_set" || len(intents) != 0 {
		t.Fatalf("stale candidate = %+v intents=%+v", stale, intents)
	}
	if state := engine.GateState(s10ExitTables(t, st), gateID); state == engine.GateResumed {
		t.Fatalf("stale reply woke old decision")
	}
	afterGuard, err := os.ReadFile(odbPath)
	if err != nil {
		t.Fatalf("read source ODB after guard: %v", err)
	}
	if !bytes.Equal(before, afterGuard) {
		t.Fatalf("in-place migrator changed source ODB bytes")
	}
	if source, err := st.Read("odb-" + gateID); err != nil || source.Headers["choices"] != oldChoices {
		t.Fatalf("source ODB view aliased: %+v, %v", source, err)
	}
	if _, err := st.Commit(stale, intents); err != nil {
		t.Fatalf("commit stale candidate: %v", err)
	}
	if bucketD, err := st.ProjectBucketD("operator"); err != nil || !s11Contains(bucketD, stale.Envelope.RelayID) {
		t.Fatalf("stale candidate absent from bucket D: %v, %v", bucketD, err)
	}
	var intent struct {
		ReplacementDecisionID string `json:"replacement_decision_id"`
	}
	if err := json.Unmarshal([]byte(stale.Body), &intent); err != nil || intent.ReplacementDecisionID == "" || intent.ReplacementDecisionID == gateID {
		t.Fatalf("durable reissue intent = %+v, %v", intent, err)
	}

	cmd := exec.Command(os.Args[0], "-test.run=^TestS11StaleChoiceRejectsAndCrashReplaysSameReplacementIdentity$")
	cmd.Env = append(os.Environ(),
		"FRANK_S11_STALE_REISSUE_CHILD=1",
		"FRANK_S11_STALE_REISSUE_ROOT="+root,
		"FRANK_TEST_CRASHPOINT=stale_reissue_after_held",
	)
	if output, err := cmd.CombinedOutput(); err == nil {
		t.Fatalf("child survived stale reissue crashpoint: %s", output)
	}
	heldID := "held-stale-schema-" + gateID
	held, err := st.Read(heldID)
	if err != nil {
		t.Fatalf("held stale_schema record not durable before crash: %v", err)
	}
	if held.Envelope.DeliveryState != record.Held || held.Envelope.To != "operator" || held.Headers["failing_edge"] != "stale_schema" {
		t.Fatalf("migration-fault signal = %+v", held)
	}
	if _, err := st.Read("outbox-held-" + heldID); err != nil {
		t.Fatalf("migration-fault signal was not operator-visible before crash: %v", err)
	}
	if _, err := st.Read(intent.ReplacementDecisionID); !os.IsNotExist(err) {
		t.Fatalf("replacement existed before recovery: %v", err)
	}

	if _, err := frankrecover.Run(root, pinned); err != nil {
		t.Fatalf("recover stale reissue: %v", err)
	}
	replacement, err := st.Read(intent.ReplacementDecisionID)
	if err != nil {
		t.Fatalf("replacement gate: %v", err)
	}
	if replacement.Envelope.From != "s11.implementer" || replacement.Envelope.DeliveryState != record.Accepted || replacement.Headers["choices"] != changedChoices {
		t.Fatalf("replacement gate = %+v", replacement)
	}
	if replacementODB, err := st.Read("odb-" + intent.ReplacementDecisionID); err != nil || replacementODB.Headers["choices"] != changedChoices {
		t.Fatalf("replacement ODB = %+v, %v", replacementODB, err)
	}
	if _, err := st.Read("park-" + intent.ReplacementDecisionID); err != nil {
		t.Fatalf("replacement gate not parked: %v", err)
	}
	if state := engine.GateState(s10ExitTables(t, st), gateID); state == engine.GateResumed {
		t.Fatalf("replacement replay resolved old decision")
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("idempotent CompleteAuto: %v", err)
	}
	assertS10RelayCount(t, st, intent.ReplacementDecisionID, 1)
	assertS10RelayCount(t, st, "held-stale-schema-"+gateID, 1)
}

func TestS11StructuralChoiceMigrationReordersAndAddsRepresentationThenResolves(t *testing.T) {
	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	oldChoices := s11Choices(t,
		map[string]string{"label": "Approve", "value": "approve"},
		map[string]string{"label": "Reject", "value": "reject"},
	)
	gateID := "s11-8a-structural"
	s11CommitGateAndODB(t, st, gateID, oldChoices)
	migrations := migrate.New()
	migrations.Current = 2
	migrations.Register(1, func(rec record.Record) (record.Record, error) {
		if rec.Headers["record_kind"] == "odb" {
			rec.Headers["choices"] = s11Choices(t,
				map[string]string{"label": "Reject", "value": "reject", "hint": "display-only"},
				map[string]string{"label": "Approve", "value": "approve", "hint": "display-only"},
			)
		}
		return rec, nil
	})
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	payload := s10ExitPayload(t, pinned.Registry, fieldspec.RenderEnv{}, meta, s11ResolutionCandidate(gateID, "approve"))
	resolved, intents, err := engine.SubmitHandlerWithMigration(st, pinned.Registry, migrations, meta)(context.Background(), intake.Cmd{
		IntakeID: "s11-structural-reply", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: payload,
	})
	if err != nil {
		t.Fatalf("SubmitHandlerWithMigration: %v", err)
	}
	if resolved.Envelope.DeliveryState != record.Accepted || resolved.Envelope.To != "s11.implementer" {
		t.Fatalf("structural migration resolution = %+v", resolved)
	}
	if _, err := st.Commit(resolved, intents); err != nil {
		t.Fatalf("commit resolution: %v", err)
	}
	if state := engine.GateState(s10ExitTables(t, st), gateID); state != engine.GateResumed {
		t.Fatalf("structural migration state = %q, want resumed", state)
	}
	if source, err := st.Read("odb-" + gateID); err != nil || source.Headers["choices"] != oldChoices {
		t.Fatalf("stored source ODB changed: %+v, %v", source, err)
	}
}

func s11CommitGateAndODB(t *testing.T, st *store.Store, gateID, choices string) {
	t.Helper()
	gate := s11BucketRecord(gateID, "authz_security")
	gate.Envelope.From = "s11.implementer"
	gate.Envelope.Role = "implementer"
	gate.Headers["HUMAN_GATE_REQUIRED"] = "yes"
	gate.Headers["choices"] = choices
	if _, err := st.Commit(gate, nil); err != nil {
		t.Fatalf("commit gate: %v", err)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto gate: %v", err)
	}
}

func s11ResolutionCandidate(gateID, choice string) record.Record {
	body, _ := json.Marshal(map[string]string{"choice": choice})
	return record.Record{
		Headers: map[string]string{
			"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
			"SUBJECT": "8a operator resolution", "record_kind": "gate_resolution", "parent_hint": gateID,
			"resolves_gate": gateID, "gate_category": "authz_security",
		},
		Body: string(body),
	}
}

func s11Choices(t *testing.T, rows ...map[string]string) string {
	t.Helper()
	encoded, err := fieldspec.CanonicalMarshal(rows)
	if err != nil {
		t.Fatalf("CanonicalMarshal choices: %v", err)
	}
	return encoded
}

func s11Contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
