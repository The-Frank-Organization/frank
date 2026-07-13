package fixtures_test

import (
	"context"
	"encoding/json"
	"sync/atomic"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestS10ExitLeg3FreshV8GateWakesExactlyOnceAfterLocalReobserve(t *testing.T) {
	root := t.TempDir()
	if err := store.Init(root, s8ConfigSources(t, false)); err != nil {
		t.Fatalf("Init fresh v8 store: %v", err)
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("Load pinned config: %v", err)
	}
	if pinned.Registry.Version != "s10-fieldspec-v8" {
		t.Fatalf("fresh store fieldspec = %q, want s10-fieldspec-v8", pinned.Registry.Version)
	}
	result, err := frankrecover.Run(root, pinned)
	if err != nil || result.Ready == nil {
		t.Fatalf("recover fresh store = %+v, %v", result, err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	journal, err := intake.OpenWithConfig(root, pinned.Engine)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var gateObservations atomic.Int32
	present := map[string]bool{"observe": true}
	observeEnv := observe.Env{
		PresentLayers: present,
		Evaluate: func(candidate observe.Candidate) observe.PredicateResult {
			if candidate.Record.Headers["HUMAN_GATE_REQUIRED"] == "yes" {
				gateObservations.Add(1)
			}
			return observe.PredicateResult{ID: "fresh-v8-gate", Predicate: observe.Pass}
		},
	}
	renderEnv := fieldspec.RenderEnv{ConfigDigest: pinned.Digest, PresentLayers: present}
	handler := func(handlerCtx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		meta := seat.SeatMeta{Name: cmd.Seat, Role: cmd.Role, IsOperator: cmd.IsOperator}
		tab, buildErr := tables.Build(st)
		if buildErr != nil {
			return record.Record{}, nil, buildErr
		}
		return engine.SubmitHandlerWithObservation(st, pinned.Registry, meta, renderEnv, observeEnv, tab)(handlerCtx, cmd)
	}
	loop := engine.New(st, handler, result.Ready)
	writer, err := intake.NewWriter[engine.Outcome](journal, pinned.Engine, result.Ready)
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	go loop.Run(ctx)
	go writer.Run(ctx, loop.In)

	seatMeta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}
	gatePayload := s10ExitPayload(t, pinned.Registry, renderEnv, seatMeta, record.Record{
		Headers: map[string]string{
			"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
			"SUBJECT": "fresh-v8 exit gate", "HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security",
			"ACTIONS_GIT_REF": "s10-exit-leg-3", "FINAL_GIT_STATUS_SHORT": "none - clean tree",
		},
	})
	gateReply, _, err := writer.Submit(ctx, intake.Cmd{Seat: seatMeta.Name, Role: seatMeta.Role, Verb: "submit", Payload: gatePayload})
	if err != nil {
		t.Fatalf("Submit gate: %v", err)
	}
	gateOutcome := <-gateReply
	if gateOutcome.State != record.Accepted || gateOutcome.RelayID == "" {
		t.Fatalf("gate outcome = %+v", gateOutcome)
	}
	gateID := gateOutcome.RelayID
	if odb := s10RecordByRelay(t, st, "odb-"+gateID); odb.Headers["record_kind"] != "odb" || odb.Headers["subject_ref"] != gateID {
		t.Fatalf("exit ODB = %+v", odb)
	}
	if park := s10RecordByRelay(t, st, "park-"+gateID); park.Headers["SUBJECT"] != engine.GateParkedWaitingHuman {
		t.Fatalf("exit park = %+v", park)
	}
	if state := engine.GateState(s10ExitTables(t, st), gateID); state != engine.GateParkedWaitingHuman {
		t.Fatalf("parked state = %q", state)
	}

	operatorMeta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	resolutionPayload := s10ExitPayload(t, pinned.Registry, renderEnv, operatorMeta, record.Record{
		Headers: map[string]string{
			"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
			"SUBJECT": "fresh-v8 exit verdict", "record_kind": "gate_resolution",
			"parent_hint": gateID, "resolves_gate": gateID, "gate_category": "authz_security",
			"ACTIONS_GIT_REF": "s10-exit-leg-3", "FINAL_GIT_STATUS_SHORT": "none - clean tree",
		},
		Body: `{"choice":"approve"}`,
	})
	wakeReply, _, err := writer.Submit(ctx, intake.Cmd{
		Seat: "operator", Role: "operator", IsOperator: true, Verb: "submit", Payload: resolutionPayload,
	})
	if err != nil {
		t.Fatalf("Submit wake: %v", err)
	}
	wakeOutcome := <-wakeReply
	if wakeOutcome.State != record.Accepted || wakeOutcome.RelayID == "" {
		t.Fatalf("wake outcome = %+v", wakeOutcome)
	}
	wake, err := st.Read(wakeOutcome.RelayID)
	if err != nil {
		t.Fatalf("Read wake: %v", err)
	}
	if wake.Envelope.From != "operator" || wake.Envelope.To != "seat-a" || wake.Headers["resolves_gate"] != gateID {
		t.Fatalf("wake record = %+v", wake)
	}
	if gateObservations.Load() != 2 {
		t.Fatalf("gate observe count = %d, want initial + local wake re-observe", gateObservations.Load())
	}
	tab := s10ExitTables(t, st)
	if state := engine.GateState(tab, gateID); state != engine.GateResumed {
		t.Fatalf("final gate state = %q, want resumed", state)
	}
	acceptedWakes := 0
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["resolves_gate"] == gateID {
			acceptedWakes++
		}
	}
	if acceptedWakes != 1 {
		t.Fatalf("accepted wake count = %d, want exactly one", acceptedWakes)
	}
	assertS10MailboxCount(t, root, "seat-a", wakeOutcome.RelayID, 1)
}

func s10ExitPayload(t *testing.T, reg *fieldspec.Registry, env fieldspec.RenderEnv, meta seat.SeatMeta, candidate record.Record) json.RawMessage {
	t.Helper()
	_, digest := reg.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, "SITREP", "medium", fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: candidate, FormDigest: digest})
	if err != nil {
		t.Fatalf("Marshal payload: %v", err)
	}
	return payload
}

func s10ExitTables(t *testing.T, st *store.Store) *tables.T {
	t.Helper()
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("Build tables: %v", err)
	}
	return tab
}
