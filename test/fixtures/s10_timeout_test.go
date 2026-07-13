package fixtures_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/executor"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestS10E1LongReadUsesTheSameParkPromptDisposition(t *testing.T) {
	lane := t.TempDir()
	if err := os.WriteFile(filepath.Join(lane, "evidence.txt"), []byte("done\n"), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	release := make(chan struct{})
	prompted := make(chan struct{})
	var once sync.Once
	var promptOnce sync.Once
	var prompts atomic.Int32
	registry := observe.NewRegistry(observe.RegistryEnv{
		Lanes: map[string]string{"repo": lane}, ReadTimeout: 20 * time.Millisecond, HardCeiling: time.Second,
		OnSoftExpiry: func(context.Context, observe.ExpiryRequest) observe.ExpiryDecision {
			prompts.Add(1)
			promptOnce.Do(func() { close(prompted) })
			return observe.ExpiryDecision{Action: observe.ExpiryExtend}
		},
		ReadFileStageHook: func(stage observe.ReadFileStage) {
			if stage == observe.ReadFileStageRead {
				once.Do(func() { <-release })
			}
		},
	})
	selection := observe.Selection{CheckID: "read-file", ClaimRef: "e1-long", Params: map[string]string{
		"lane_ref": "repo", "path": "evidence.txt", "expect": "line:done",
	}}
	verdictCh := make(chan observe.CheckVerdict, 1)
	go func() { verdictCh <- registry.Run(selection) }()
	select {
	case <-prompted:
	case <-time.After(time.Second):
		t.Fatal("E1 soft expiry did not invoke prompt disposition")
	}
	close(release)
	verdict := <-verdictCh
	if verdict.Outcome != "pass" || prompts.Load() != 1 {
		t.Fatalf("extended E1 verdict/prompts = %#v/%d", verdict, prompts.Load())
	}
}

func TestS10SuiteSoftExpiryPromptsBeforeAnyKillAndHardCeilingOnlyBlocks(t *testing.T) {
	entry := observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}
	selection := observe.Selection{CheckID: "run-suite", ClaimRef: "long-run", Params: map[string]string{"target": "slow", "expect_green": "true"}}
	newHost := func(disposition executor.ExpiryDisposition, hardCeiling time.Duration) *executor.Host {
		source := t.TempDir()
		s8WriteExecutable(t, source, "slow.sh", "#!/bin/sh\nsleep 1\n")
		return executor.New(executor.Config{
			TempRoot:     t.TempDir(),
			HardCeiling:  hardCeiling,
			OnSoftExpiry: disposition,
			Suites: map[string]executor.Suite{
				"slow": {SourceDir: source, Command: "slow.sh", TimeoutClass: "suite_bounded", Timeout: 20 * time.Millisecond},
			},
		})
	}

	var prompts atomic.Int32
	extended := newHost(func(context.Context, executor.ExpiryRequest) executor.ExpiryDecision {
		prompts.Add(1)
		return executor.ExpiryDecision{Action: executor.ExpiryExtend}
	}, 2*time.Second)
	verdict := extended.Spawn(entry, selection)
	if prompts.Load() != 1 || verdict.Outcome != "pass" || verdict.Predicate != observe.Pass || verdict.Timing != "extended" {
		t.Fatalf("extended verdict/prompts = %#v/%d", verdict, prompts.Load())
	}

	killed := newHost(func(context.Context, executor.ExpiryRequest) executor.ExpiryDecision {
		return executor.ExpiryDecision{Action: executor.ExpiryKill}
	}, 2*time.Second).Spawn(entry, selection)
	if killed.Outcome != "unsafe" || killed.Predicate != observe.Blocked || killed.FailingDetail != "executor-timeout" {
		t.Fatalf("operator kill verdict = %#v", killed)
	}

	ceiling := newHost(func(ctx context.Context, _ executor.ExpiryRequest) executor.ExpiryDecision {
		<-ctx.Done()
		return executor.ExpiryDecision{Action: executor.ExpiryExtend}
	}, 45*time.Millisecond).Spawn(entry, selection)
	if ceiling.Outcome != "unsafe" || ceiling.Predicate != observe.Blocked || ceiling.FailingDetail != "executor-timeout" {
		t.Fatalf("hard-ceiling verdict = %#v", ceiling)
	}
}

func TestS10LongRunningCheckParksAlertsAndAppliesOperatorExtend(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var host *executor.Host
	finalVerdict := make(chan observe.CheckVerdict, 1)
	base := func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		if cmd.Verb == "run" {
			entry := observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}
			selection := observe.Selection{CheckID: "run-suite", ClaimRef: "park-me", Seat: "seat-a", CandidateDigest: "candidate-a", Params: map[string]string{"target": "slow", "expect_green": "true"}}
			verdict := host.Spawn(entry, selection)
			finalVerdict <- verdict
			return record.Record{
				Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "long check completed"},
			}, nil, nil
		}
		var gateID string
		if err := json.Unmarshal(cmd.Payload, &gateID); err != nil {
			return record.Record{}, nil, err
		}
		return record.Record{
			Envelope: record.Envelope{From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "expiry verdict", "resolves_gate": gateID},
			Body:     `{"choice":"extend"}`,
		}, nil, nil
	}
	loop := engine.New(st, engine.ExpiryHandler(base), engine.TestReady())
	loop.ServiceWhileBlocked = true
	journal, err := intake.Open(st.Root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	writer, err := intake.NewWriter[engine.Outcome](journal, config.EngineConfig{}, engine.TestReady())
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	prompter, err := engine.NewExpiryPrompter(st, writer)
	if err != nil {
		t.Fatalf("NewExpiryPrompter: %v", err)
	}
	loop.AfterGateResolution = prompter.Apply
	go loop.Run(ctx)
	go writer.Run(ctx, loop.In)

	source := t.TempDir()
	s8WriteExecutable(t, source, "slow.sh", "#!/bin/sh\nsleep 1\n")
	host = executor.New(executor.Config{
		TempRoot: t.TempDir(), HardCeiling: 2 * time.Second, OnSoftExpiry: prompter.Prompt,
		Suites: map[string]executor.Suite{
			"slow": {SourceDir: source, Command: "slow.sh", TimeoutClass: "suite_bounded", Timeout: 20 * time.Millisecond},
		},
	})
	runReply, _, err := writer.Submit(ctx, intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "run", Payload: []byte(`{}`)})
	if err != nil {
		t.Fatalf("Submit run: %v", err)
	}

	var gateID string
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) && gateID == "" {
		records, readErr := st.Records()
		if readErr != nil {
			t.Fatalf("Records: %v", readErr)
		}
		for _, rec := range records {
			if rec.Headers["expiry_check_id"] == "run-suite" {
				gateID = rec.Envelope.RelayID
			}
		}
		if gateID == "" {
			time.Sleep(10 * time.Millisecond)
		}
	}
	if gateID == "" {
		t.Fatal("soft expiry did not emit an accepted gate")
	}
	for _, relayID := range []string{"odb-" + gateID, "park-" + gateID, "outbox-gate-" + gateID} {
		var rec record.Record
		var readErr error
		derivedDeadline := time.Now().Add(time.Second)
		for time.Now().Before(derivedDeadline) {
			rec, readErr = st.Read(relayID)
			if readErr == nil {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		if readErr != nil || rec.Envelope.DeliveryState != record.Accepted {
			t.Fatalf("derived alert %s = %#v, %v", relayID, rec, readErr)
		}
		if strings.HasPrefix(relayID, "odb-") && (!strings.Contains(rec.Headers["choices"], `"value":"kill"`) || !strings.Contains(rec.Headers["choices"], `"value":"extend"`)) {
			t.Fatalf("expiry ODB choices = %q", rec.Headers["choices"])
		}
	}
	gatePayload, _ := json.Marshal(gateID)
	reply, _, err := writer.Submit(ctx, intake.Cmd{Seat: "operator", Role: "operator", Verb: "resolve", Payload: gatePayload})
	if err != nil {
		t.Fatalf("Submit resolution: %v", err)
	}
	if outcome := <-reply; outcome.State != record.Accepted {
		t.Fatalf("resolution outcome = %+v", outcome)
	}
	if outcome := <-runReply; outcome.State != record.Accepted {
		t.Fatalf("run outcome = %+v", outcome)
	}
	if verdict := <-finalVerdict; verdict.Outcome != "pass" || verdict.Timing != "extended" {
		t.Fatalf("extended final verdict = %#v", verdict)
	}
}

func TestS10ExpiryPrompterRejectsNonOperatorExtendAcrossReplayAndLivePaths(t *testing.T) {
	request := observe.ExpiryRequest{
		Selection: observe.Selection{
			CheckID: "run-suite", ClaimRef: "non-operator-extend", Seat: "seat-a", CandidateDigest: "candidate-a",
		},
		SoftExpiry: 10 * time.Millisecond, HardCeiling: 50 * time.Millisecond,
	}
	nonOperatorExtend := func(gateID string) record.Record {
		return record.Record{
			Envelope: record.Envelope{
				RelayID: "non-operator-extend", From: "seat-a", Role: "implementer",
				DeliveryState: record.Accepted, SchemaVersion: 1,
			},
			Headers: map[string]string{"resolves_gate": gateID},
			Body:    `{"choice":"extend"}`,
		}
	}

	t.Run("replay lookup", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		submitter := &s10ExpirySubmitter{store: st}
		submitter.afterGate = func(gateID string) error {
			_, err := st.Commit(nonOperatorExtend(gateID), []store.Intent{})
			return err
		}
		prompter, err := engine.NewExpiryPrompter(st, submitter)
		if err != nil {
			t.Fatalf("NewExpiryPrompter: %v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		if got := prompter.Prompt(ctx, request); got.Action != observe.ExpiryKill {
			t.Fatalf("non-operator replay decision = %q, want kill", got.Action)
		}
	})

	t.Run("live apply", func(t *testing.T) {
		st, err := store.Open(t.TempDir())
		if err != nil {
			t.Fatalf("Open: %v", err)
		}
		gateReady := make(chan string, 1)
		submitter := &s10ExpirySubmitter{store: st, gateReady: gateReady}
		prompter, err := engine.NewExpiryPrompter(st, submitter)
		if err != nil {
			t.Fatalf("NewExpiryPrompter: %v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		decision := make(chan observe.ExpiryDecision, 1)
		go func() { decision <- prompter.Prompt(ctx, request) }()
		gateID := <-gateReady
		if err := prompter.Apply(nonOperatorExtend(gateID)); err != nil {
			t.Fatalf("Apply: %v", err)
		}
		if got := <-decision; got.Action != observe.ExpiryKill {
			t.Fatalf("non-operator live decision = %q, want kill", got.Action)
		}
	})
}

type s10ExpirySubmitter struct {
	store     *store.Store
	afterGate func(string) error
	gateReady chan<- string
}

func (s *s10ExpirySubmitter) Submit(_ context.Context, cmd intake.Cmd) (<-chan engine.Outcome, string, error) {
	var input engine.ExpiryPromptInput
	if err := json.Unmarshal(cmd.Payload, &input); err != nil {
		return nil, "", err
	}
	gateID := engine.ExpiryGateID(input)
	gate := record.Record{
		Envelope: record.Envelope{
			RelayID: gateID, From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1,
		},
		Headers: map[string]string{"expiry_check_id": input.CheckID},
	}
	if _, err := s.store.Commit(gate, []store.Intent{}); err != nil {
		return nil, "", err
	}
	if s.afterGate != nil {
		if err := s.afterGate(gateID); err != nil {
			return nil, "", err
		}
	}
	if s.gateReady != nil {
		s.gateReady <- gateID
	}
	reply := make(chan engine.Outcome, 1)
	reply <- engine.Outcome{State: record.Accepted, RelayID: gateID}
	return reply, "expiry-test-intake", nil
}
