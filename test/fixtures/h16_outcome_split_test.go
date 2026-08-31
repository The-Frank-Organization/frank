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
)

func TestH16OutcomeSplitFivePostCommitFailureSites(t *testing.T) {
	tests := []struct {
		name         string
		configure    func(*engine.Loop)
		cmd          intake.Cmd
		wantDecision string
		wantReason   string
		wantDetail   string
	}{
		{
			name: "completeTurn",
			configure: func(loop *engine.Loop) {
				loop.AfterCommit = func(*store.Store) error { return errors.New("injected completeTurn failure") }
			},
			wantDecision: record.Accepted,
			wantReason:   "system:obligation-error",
		},
		{
			name: "AfterGateResolution",
			configure: func(loop *engine.Loop) {
				loop.AfterGateResolution = func(record.Record) error { return errors.New("injected gate-derived failure") }
			},
			cmd:          intake.Cmd{Verb: "submit", Payload: json.RawMessage(`{"resolves_gate":true}`)},
			wantDecision: record.Accepted,
			wantReason:   "system:derived-work-error",
		},
		{
			name: "AfterApprovalResolution",
			configure: func(loop *engine.Loop) {
				loop.AfterApprovalResolution = func(record.Record) error { return errors.New("injected approval-derived failure") }
			},
			cmd:          intake.Cmd{Verb: "submit", Payload: json.RawMessage(`{"resolves_gate":true}`)},
			wantDecision: record.Accepted,
			wantReason:   "system:derived-work-error",
		},
		{
			name: "AfterAccepted",
			configure: func(loop *engine.Loop) {
				loop.AfterAccepted = func(record.Record) (engine.OutcomeExtras, error) {
					return engine.OutcomeExtras{}, errors.New("injected accepted-derived failure")
				}
			},
			wantDecision: record.Accepted,
			wantReason:   "system:derived-work-error",
		},
		{
			name: "supersededCredentialCompleteTurn",
			configure: func(loop *engine.Loop) {
				loop.CurrentAuthGeneration = func(string) string { return "pivot-current" }
				loop.AfterCommit = func(*store.Store) error { return errors.New("injected superseded completeTurn failure") }
			},
			cmd: intake.Cmd{
				Verb:           "submit",
				AuthGeneration: "pivot-stale",
				Payload:        json.RawMessage(`{"stale":true}`),
			},
			wantDecision: record.Rejected,
			wantReason:   "system:obligation-error",
			wantDetail:   "auth_generation:credential-superseded",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := h16RunOutcome(t, tc.configure, tc.cmd)
			encoded, err := json.Marshal(out)
			if err != nil {
				t.Fatalf("marshal outcome: %v", err)
			}
			var got map[string]any
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("decode outcome: %v", err)
			}
			if _, present := got["state"]; present {
				t.Fatalf("pending outcome exposed legacy state: %s", encoded)
			}
			if got["decision_state"] != tc.wantDecision || got["post_commit_state"] != "pending" {
				t.Fatalf("split outcome=%s, want decision=%q post_commit=pending", encoded, tc.wantDecision)
			}
			if got["reason"] != tc.wantReason {
				t.Fatalf("reason=%q, want %q; outcome=%s", got["reason"], tc.wantReason, encoded)
			}
			if tc.wantDetail != "" && !strings.Contains(out.Detail, tc.wantDetail) {
				t.Fatalf("detail=%q, want preserved %q", out.Detail, tc.wantDetail)
			}
		})
	}
}

func TestH16PendingOutcomeFailsClosedForLegacyStateOnlyDecode(t *testing.T) {
	out := h16RunOutcome(t, func(loop *engine.Loop) {
		loop.AfterAccepted = func(record.Record) (engine.OutcomeExtras, error) {
			return engine.OutcomeExtras{}, errors.New("injected pending state")
		}
	}, intake.Cmd{})
	encoded, err := json.Marshal(out)
	if err != nil {
		t.Fatalf("marshal outcome: %v", err)
	}

	var legacy struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(encoded, &legacy); err != nil {
		t.Fatalf("legacy decode: %v", err)
	}
	if legacy.State == record.Accepted || legacy.State == record.Rejected || legacy.State == record.Held {
		t.Fatalf("legacy state-only decoder matched pending outcome as %q: %s", legacy.State, encoded)
	}

	var split struct {
		DecisionState   string `json:"decision_state"`
		PostCommitState string `json:"post_commit_state"`
	}
	if err := json.Unmarshal(encoded, &split); err != nil {
		t.Fatalf("split decode: %v", err)
	}
	if split.DecisionState != record.Accepted || split.PostCommitState != "pending" {
		t.Fatalf("split decode=%+v, want accepted/pending; bytes=%s", split, encoded)
	}
}

func TestH16OutcomeSplitCompleteEmittersCarryBothDimensions(t *testing.T) {
	t.Run("committed accepted", func(t *testing.T) {
		out := h16RunOutcome(t, func(*engine.Loop) {}, intake.Cmd{})
		assertH16CompleteOutcome(t, out, record.Accepted)
	})
	t.Run("precommit rejected", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("open store: %v", err)
		}
		loop := engine.New(st, nil, engine.TestReady())
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go loop.Run(ctx)
		reply := make(chan engine.Outcome, 1)
		loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "h16-precommit-reject"}, ReplyCh: reply}
		assertH16CompleteOutcome(t, <-reply, record.Rejected)
	})
}

func assertH16CompleteOutcome(t *testing.T, out engine.Outcome, wantDecision string) {
	t.Helper()
	encoded, err := json.Marshal(out)
	if err != nil {
		t.Fatalf("marshal outcome: %v", err)
	}
	var got map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatalf("decode outcome: %v", err)
	}
	if got["state"] != wantDecision || got["decision_state"] != wantDecision || got["post_commit_state"] != "complete" {
		t.Fatalf("complete outcome=%s, want state=decision=%q post_commit=complete", encoded, wantDecision)
	}
}

func h16RunOutcome(t *testing.T, configure func(*engine.Loop), cmd intake.Cmd) engine.Outcome {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		headers := map[string]string{"PHASE": "SITREP", "SUBJECT": "h16 post-commit failure"}
		if strings.Contains(string(cmd.Payload), "resolves_gate") {
			headers["resolves_gate"] = "gate-1"
		}
		return record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  headers,
		}, nil, nil
	}, engine.TestReady())
	configure(loop)

	if cmd.IntakeID == "" {
		cmd.IntakeID = "h16-outcome-" + strings.ReplaceAll(t.Name(), "/", "-")
	}
	if cmd.Seat == "" {
		cmd.Seat = "s12.implementer"
	}
	if cmd.Role == "" {
		cmd.Role = "implementer"
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)
	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: cmd, ReplyCh: reply}
	return <-reply
}
