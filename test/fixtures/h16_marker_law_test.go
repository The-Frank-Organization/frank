package fixtures_test

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

func TestH16MintMarkerCommitFailureNeverEntersEffectAndCallerReplayHeals(t *testing.T) {
	loop, st, state, cancel := h16MarkerLoop(t)
	defer cancel()
	failMarker := true
	loop.DerivedCommit = func(rec record.Record) (string, error) {
		if failMarker && rec.Headers["record_kind"] == "derived-work-attempt" {
			failMarker = false
			return "", errors.New("injected marker commit failure")
		}
		return st.Commit(rec, nil)
	}
	cmd := h16MintCmd("marker-failure")
	first := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "pending", false)
	if state.effectCalls != 0 {
		t.Fatalf("effect calls=%d before durable marker", state.effectCalls)
	}
	second := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, second, record.Accepted, "complete", true)
	if state.effectCalls != 1 || second.Credential != "credential-once" {
		t.Fatalf("healing outcome=%+v effectCalls=%d", second, state.effectCalls)
	}
}

func TestH16MintDurableMarkerWithoutEvidenceResumesOnlyForCaller(t *testing.T) {
	loop, _, state, cancel := h16MarkerLoop(t)
	defer cancel()
	classGCalls := 0
	loop.AfterCommit = func(*store.Store) error {
		classGCalls++
		if classGCalls == 2 {
			return errors.New("injected post-marker failure")
		}
		return nil
	}
	cmd := h16MintCmd("marker-before-effect")
	first := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "pending", false)
	if state.effectCalls != 0 || state.realizedRef != "" {
		t.Fatalf("effect crossed failed post-marker hook: calls=%d realized=%q", state.effectCalls, state.realizedRef)
	}

	healed := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, healed, record.Accepted, "complete", true)
	if healed.Credential != "credential-once" || state.effectCalls != 1 {
		t.Fatalf("caller did not resume unrealized marker: outcome=%+v calls=%d", healed, state.effectCalls)
	}
}

func TestH16MintRealizedBeforeAdvanceIsFailedUndeliveredAndNeverReruns(t *testing.T) {
	loop, st, state, cancel := h16MarkerLoop(t)
	defer cancel()
	failAdvance := true
	loop.DerivedCommit = func(rec record.Record) (string, error) {
		if failAdvance && rec.Headers["record_kind"] == "derived-work-transition" {
			failAdvance = false
			return "", errors.New("injected advance commit failure")
		}
		return st.Commit(rec, nil)
	}
	cmd := h16MintCmd("advance-failure")
	first := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "failed", false)
	if first.Credential != "" || state.effectCalls != 1 {
		t.Fatalf("post-effect failure leaked extras or missed effect: %+v calls=%d", first, state.effectCalls)
	}
	assertRealizedUndeliveredTransition(t, st, first.RelayID, 1)
	second := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, second, record.Accepted, "failed", false)
	if second.Credential != "" || state.effectCalls != 1 {
		t.Fatalf("failed replay reran or leaked extras: %+v calls=%d", second, state.effectCalls)
	}
	assertRealizedUndeliveredTransition(t, st, first.RelayID, 1)
	if status := engineDerivedStatus(t, st, first.RelayID); status != "failed" {
		t.Fatalf("canonical restart fold=%q, want failed", status)
	}
}

func TestH16MintTerminalReplyCarriesExtrasOnceAndReplayDoesNot(t *testing.T) {
	loop, _, state, cancel := h16MarkerLoop(t)
	defer cancel()
	cmd := h16MintCmd("terminal-reply")
	first := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "complete", true)
	if first.Credential != "credential-once" || state.effectCalls != 1 {
		t.Fatalf("terminal outcome=%+v calls=%d", first, state.effectCalls)
	}
	second := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, second, record.Accepted, "complete", true)
	if second.Credential != "" || state.effectCalls != 1 {
		t.Fatalf("terminal replay leaked extras or reran: %+v calls=%d", second, state.effectCalls)
	}
}

func TestH16EarlierClassGFailureAndUnrelatedCommandNeverMint(t *testing.T) {
	loop, _, state, cancel := h16MarkerLoop(t)
	defer cancel()
	failClassG := true
	loop.AfterCommit = func(*store.Store) error {
		if failClassG {
			failClassG = false
			return errors.New("injected earlier Class-G failure")
		}
		return nil
	}
	mint := h16MintCmd("caller-only")
	first := h16SubmitLoop(t, loop, mint)
	assertH16PostCommitState(t, first, record.Accepted, "pending", false)
	if state.effectCalls != 0 {
		t.Fatalf("mint ran after earlier hook failure: %d", state.effectCalls)
	}
	unrelated := intake.Cmd{IntakeID: "unrelated-command", Seat: "operator", Role: "operator", Verb: "submit"}
	_ = h16SubmitLoop(t, loop, unrelated)
	if state.effectCalls != 0 {
		t.Fatalf("unrelated command minted: %d", state.effectCalls)
	}
	healed := h16SubmitLoop(t, loop, mint)
	assertH16PostCommitState(t, healed, record.Accepted, "complete", true)
	if healed.Credential != "credential-once" || state.effectCalls != 1 {
		t.Fatalf("caller replay=%+v calls=%d", healed, state.effectCalls)
	}
}

type h16MarkerState struct {
	realizedRef string
	effectCalls int
}

func h16MarkerLoop(t *testing.T) (*engine.Loop, *store.Store, *h16MarkerState, context.CancelFunc) {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		headers := map[string]string{"PHASE": "SITREP", "SUBJECT": cmd.IntakeID}
		if strings.HasPrefix(cmd.IntakeID, "mint-") {
			headers["record_kind"] = "seat_mint"
		}
		return record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  headers,
			Body:     `{"seat":"s12-marker.implementer","role":"implementer","is_operator":false}`,
		}, nil, nil
	}, engine.TestReady())
	state := &h16MarkerState{}
	loop.MintRealized = func(rec record.Record) bool { return state.realizedRef == rec.Envelope.RelayID }
	loop.AfterAccepted = func(rec record.Record) (engine.OutcomeExtras, error) {
		if rec.Headers["record_kind"] != "seat_mint" {
			return engine.OutcomeExtras{}, nil
		}
		state.effectCalls++
		state.realizedRef = rec.Envelope.RelayID
		return engine.OutcomeExtras{Credential: "credential-once", Endpoint: "local"}, nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	go loop.Run(ctx)
	return loop, st, state, cancel
}

func h16MintCmd(name string) intake.Cmd {
	return intake.Cmd{IntakeID: "mint-" + name, Seat: "operator", Role: "operator", Verb: "submit"}
}

func assertRealizedUndeliveredTransition(t *testing.T, st *store.Store, sourceRelayID string, want int) {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("records: %v", err)
	}
	got := 0
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "derived-work-transition" {
			continue
		}
		var body struct {
			SourceRelayID string `json:"source_relay_id"`
			Kind          string `json:"kind"`
		}
		if json.Unmarshal([]byte(rec.Body), &body) == nil && body.SourceRelayID == sourceRelayID && body.Kind == "realized-undelivered" {
			got++
		}
	}
	if got != want {
		t.Fatalf("realized-undelivered transitions=%d, want %d", got, want)
	}
}

func engineDerivedStatus(t *testing.T, st *store.Store, sourceRelayID string) string {
	t.Helper()
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("tables.Build: %v", err)
	}
	return tab.DerivedWork[sourceRelayID].Status
}
