package fixtures_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16DirtyClassGBlocksBlindClassDUntilSuccessfulPass(t *testing.T) {
	st := mustH16Store(t)
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers: map[string]string{
				"PHASE":         "SITREP",
				"SUBJECT":       "H16 Class-G barrier",
				"resolves_gate": "h16-gate-target",
			},
		}, nil, nil
	}, engine.TestReady())
	classGHealthy := false
	var events []string
	loop.AfterCommit = func(*store.Store) error {
		if !classGHealthy {
			events = append(events, "class-g-failed")
			return errors.New("injected persistent Class-G failure")
		}
		events = append(events, "class-g-succeeded")
		return nil
	}
	gateCalls, approvalCalls := 0, 0
	loop.AfterGateResolution = func(record.Record) error {
		gateCalls++
		events = append(events, "gate")
		return nil
	}
	loop.AfterApprovalResolution = func(record.Record) error {
		approvalCalls++
		events = append(events, "approval")
		return nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	cmd := intake.Cmd{IntakeID: "h16-classg-barrier", Seat: "operator", Role: "operator", Verb: "submit"}
	first := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "pending", false)
	beforeBlockedReplay, err := st.Records()
	if err != nil {
		t.Fatalf("records before blocked replay: %v", err)
	}
	blocked := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, blocked, record.Accepted, "pending", false)
	afterBlockedReplay, err := st.Records()
	if err != nil {
		t.Fatalf("records after blocked replay: %v", err)
	}
	if gateCalls != 0 || approvalCalls != 0 || len(afterBlockedReplay) != len(beforeBlockedReplay) {
		t.Fatalf("dirty Class-G allowed Class-D work: gate=%d approval=%d records=%d->%d events=%v", gateCalls, approvalCalls, len(beforeBlockedReplay), len(afterBlockedReplay), events)
	}

	classGHealthy = true
	healed := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, healed, record.Accepted, "complete", true)
	if gateCalls != 1 || approvalCalls != 1 {
		t.Fatalf("healing hooks gate=%d approval=%d, want one each; events=%v", gateCalls, approvalCalls, events)
	}
	wantTail := []string{"class-g-succeeded", "gate", "class-g-succeeded", "approval", "class-g-succeeded"}
	if len(events) < len(wantTail) || !reflect.DeepEqual(events[len(events)-len(wantTail):], wantTail) {
		t.Fatalf("healing order=%v, want tail %v", events, wantTail)
	}
	beforeIdempotentReplay, err := st.Records()
	if err != nil {
		t.Fatalf("records before idempotent replay: %v", err)
	}
	replayed := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, replayed, record.Accepted, "complete", true)
	afterIdempotentReplay, err := st.Records()
	if err != nil {
		t.Fatalf("records after idempotent replay: %v", err)
	}
	if gateCalls != 1 || approvalCalls != 1 || len(afterIdempotentReplay) != len(beforeIdempotentReplay) {
		t.Fatalf("idempotent replay changed work: gate=%d approval=%d records=%d->%d events=%v", gateCalls, approvalCalls, len(beforeIdempotentReplay), len(afterIdempotentReplay), events)
	}
}
