package fixtures_test

import (
	"context"
	"encoding/json"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

const s10LiveApprovalClaimBoundary = "the live-prompt approval channel is proven to the gate; EXECUTION of side-effecting classes remains fail-closed refused at the executor — a future design round when a real side-effecting operation exists"

type s10SpawnSpy struct {
	calls atomic.Int32
}

func (s *s10SpawnSpy) Spawn(observe.CheckEntry, observe.Selection) observe.CheckVerdict {
	s.calls.Add(1)
	return observe.CheckVerdict{Outcome: "pass", Predicate: observe.Pass}
}

func TestS10SideEffectingGateDefaultsDenyAndApprovedGateStillNeverSpawns(t *testing.T) {
	spy := &s10SpawnSpy{}
	selection := observe.Selection{CheckID: "run-suite-unbounded", ClaimRef: "live-gate", Seat: "seat-a", CandidateDigest: "candidate-a", Params: map[string]string{
		"target": "slow", "expect_green": "true",
	}}
	deniedRegistry := observe.NewRegistry(observe.RegistryEnv{NamedSuites: map[string]bool{"slow": true}, Executor: spy})
	entry, ok := deniedRegistry.Entry("run-suite-unbounded")
	if !ok || entry.Class != "side_effecting" || !entry.ExecutorRequired || entry.TimeoutClass != "unbounded" {
		t.Fatalf("owner-named representative entry = %#v, present=%v", entry, ok)
	}
	denied := deniedRegistry.Run(selection)
	if denied.Outcome != "unsafe" || denied.Predicate != observe.Blocked || denied.FailingDetail != "side-effecting-unapproved" {
		t.Fatalf("default-DENY pre-spawn verdict = %#v", denied)
	}
	approvedRegistry := observe.NewRegistry(observe.RegistryEnv{
		NamedSuites: map[string]bool{"slow": true}, Executor: spy,
		OnSideEffectApproval: func(context.Context, observe.ApprovalRequest) observe.ApprovalDecision {
			return observe.ApprovalDecision{Allowed: true, Scope: observe.ApprovalOnce}
		},
	})
	approved := approvedRegistry.Run(selection)
	if approved.FailingDetail != "side-effecting-execution-refused" || spy.calls.Load() != 0 {
		t.Fatalf("%s: verdict=%#v spawn_calls=%d", s10LiveApprovalClaimBoundary, approved, spy.calls.Load())
	}
}

func TestS10LiveOperatorVerdictLiftsApprovalGateThroughODBWithoutExecution(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	spy := &s10SpawnSpy{}
	var registry *observe.Registry
	result := make(chan observe.CheckVerdict, 1)
	base := func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		if cmd.Verb == "run-side-effecting" {
			result <- registry.Run(observe.Selection{
				CheckID: "run-suite-unbounded", ClaimRef: "live-gate", Seat: "seat-a", CandidateDigest: "candidate-a",
				Params: map[string]string{"target": "slow", "expect_green": "true"},
			})
			return record.Record{
				Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "approval gate evaluated"},
			}, nil, nil
		}
		var gateID string
		if err := json.Unmarshal(cmd.Payload, &gateID); err != nil {
			return record.Record{}, nil, err
		}
		return record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers: map[string]string{
				"PHASE": "SITREP", "SUBJECT": "approval verdict", "resolves_gate": gateID,
			},
			Body: `{"choice":"allow-once"}`,
		}, nil, nil
	}
	loop := engine.New(st, engine.ApprovalHandler(base), engine.TestReady())
	loop.ServiceWhileBlocked = true
	journal, err := intake.Open(st.Root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	writer, err := intake.NewWriter[engine.Outcome](journal, config.EngineConfig{}, engine.TestReady())
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	prompter, err := engine.NewApprovalPrompter(st, writer)
	if err != nil {
		t.Fatalf("NewApprovalPrompter: %v", err)
	}
	loop.AfterApprovalResolution = prompter.Apply
	registry = observe.NewRegistry(observe.RegistryEnv{
		NamedSuites: map[string]bool{"slow": true}, Executor: spy, OnSideEffectApproval: prompter.Prompt,
	})
	go loop.Run(ctx)
	go writer.Run(ctx, loop.In)

	runReply, _, err := writer.Submit(ctx, intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "run-side-effecting", Payload: []byte(`{}`)})
	if err != nil {
		t.Fatalf("Submit run: %v", err)
	}
	gateID := s10AwaitApprovalGate(t, st)
	for _, relayID := range []string{"odb-" + gateID, "park-" + gateID, "outbox-gate-" + gateID} {
		rec := s10AwaitRecord(t, st, relayID)
		if rec.Envelope.DeliveryState != record.Accepted {
			t.Fatalf("derived approval alert %s = %#v", relayID, rec)
		}
		if strings.HasPrefix(relayID, "odb-") && (!strings.Contains(rec.Headers["choices"], `"value":"allow-once"`) || !strings.Contains(rec.Headers["choices"], `"value":"allow-for-entry"`)) {
			t.Fatalf("approval ODB choices = %q", rec.Headers["choices"])
		}
	}
	payload, _ := json.Marshal(gateID)
	reply, _, err := writer.Submit(ctx, intake.Cmd{Seat: "operator", Role: "operator", Verb: "resolve", Payload: payload})
	if err != nil {
		t.Fatalf("Submit resolution: %v", err)
	}
	if outcome := <-reply; outcome.State != record.Accepted {
		t.Fatalf("resolution outcome = %+v", outcome)
	}
	if outcome := <-runReply; outcome.State != record.Accepted {
		t.Fatalf("run outcome = %+v", outcome)
	}
	verdict := <-result
	if verdict.FailingDetail != "side-effecting-execution-refused" || spy.calls.Load() != 0 {
		t.Fatalf("%s: verdict=%#v spawn_calls=%d", s10LiveApprovalClaimBoundary, verdict, spy.calls.Load())
	}
}

func TestS10ApprovalPrompterSharesDuplicateGateWaiters(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	gateReady := make(chan string, 1)
	submitter := &s10ApprovalSubmitter{store: st, gateReady: gateReady}
	prompter, err := engine.NewApprovalPrompter(st, submitter)
	if err != nil {
		t.Fatalf("NewApprovalPrompter: %v", err)
	}
	request := observe.ApprovalRequest{Selection: observe.Selection{
		CheckID: "run-suite-unbounded", ClaimRef: "same-gate", Seat: "seat-a", CandidateDigest: "same-candidate",
	}}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	decisions := make(chan observe.ApprovalDecision, 2)
	go func() { decisions <- prompter.Prompt(ctx, request) }()
	gateID := <-gateReady
	go func() { decisions <- prompter.Prompt(ctx, request) }()
	time.Sleep(30 * time.Millisecond)
	resolution := record.Record{
		Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"resolves_gate": gateID},
		Body:     `{"choice":"allow-once"}`,
	}
	if err := prompter.Apply(resolution); err != nil {
		t.Fatalf("Apply: %v", err)
	}
	for i := 0; i < 2; i++ {
		if got := <-decisions; !got.Allowed || got.Scope != observe.ApprovalOnce {
			t.Fatalf("duplicate waiter %d decision = %#v, want allow-once", i, got)
		}
	}
	if calls := submitter.calls.Load(); calls != 1 {
		t.Fatalf("duplicate gate submit calls = %d, want one", calls)
	}
}

type s10ApprovalSubmitter struct {
	store     *store.Store
	gateReady chan<- string
	calls     atomic.Int32
}

func (s *s10ApprovalSubmitter) Submit(_ context.Context, cmd intake.Cmd) (<-chan engine.Outcome, string, error) {
	s.calls.Add(1)
	var input engine.ApprovalPromptInput
	if err := json.Unmarshal(cmd.Payload, &input); err != nil {
		return nil, "", err
	}
	gateID := engine.ApprovalGateID(input)
	if _, err := s.store.Commit(record.Record{
		Envelope: record.Envelope{RelayID: gateID, From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"approval_entry_id": input.EntryID},
	}, []store.Intent{}); err != nil {
		return nil, "", err
	}
	if s.gateReady != nil {
		s.gateReady <- gateID
	}
	reply := make(chan engine.Outcome, 1)
	reply <- engine.Outcome{State: record.Accepted, RelayID: gateID}
	return reply, "approval-test-intake", nil
}

func s10AwaitApprovalGate(t *testing.T, st *store.Store) string {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		records, err := st.Records()
		if err != nil {
			t.Fatalf("Records: %v", err)
		}
		for _, rec := range records {
			if rec.Headers["approval_entry_id"] == "run-suite-unbounded" {
				return rec.Envelope.RelayID
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("live approval gate not emitted")
	return ""
}

func s10AwaitRecord(t *testing.T, st *store.Store, relayID string) record.Record {
	t.Helper()
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		rec, err := st.Read(relayID)
		if err == nil {
			return rec
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("record %s not committed", relayID)
	return record.Record{}
}
