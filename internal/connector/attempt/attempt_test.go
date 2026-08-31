package attempt

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/connector/outcome"
	"github.com/The-Frank-Organization/frank/internal/connector/stream"
)

const (
	testB = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	testE = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)

func TestTypedCancelBeforeTransportIsZeroWireAndImmediate(t *testing.T) {
	manager := New()
	active, err := manager.Begin(context.Background(), "attempt-1", 7, testB, testE)
	if err != nil {
		t.Fatal(err)
	}
	cancellation, err := manager.Cancel(Intent{AttemptID: "attempt-1", TurnEpoch: 7}, 7)
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-active.Context().Done():
	default:
		t.Fatal("HTTP context was not aborted synchronously")
	}
	if cancellation.Duplicate || cancellation.Event.Kind != stream.Cancelled || cancellation.Event.Partial != "none" || cancellation.Outcome.Kind != outcome.Cancelled || cancellation.Outcome.CancelPoint != outcome.PreTransport {
		t.Fatalf("pre-transport cancellation = %+v", cancellation)
	}
	if active.Invoked() {
		t.Fatal("pre-transport cancellation crossed the invocation boundary")
	}
}

func TestTypedCancelAfterInvocationPreservesPartialTruth(t *testing.T) {
	manager := New()
	active, err := manager.Begin(context.Background(), "attempt-2", 7, testB, testE)
	if err != nil {
		t.Fatal(err)
	}
	active.MarkInvoked()
	if err := active.SetPartial("tool_call_incomplete"); err != nil {
		t.Fatal(err)
	}
	cancellation, err := manager.Cancel(Intent{AttemptID: "attempt-2", TurnEpoch: 7}, 7)
	if err != nil {
		t.Fatal(err)
	}
	if !active.Invoked() || cancellation.Event.Partial != "tool_call_incomplete" || cancellation.Outcome.CancelPoint != outcome.PostInvocation {
		t.Fatalf("post-invocation cancellation = %+v invoked=%v", cancellation, active.Invoked())
	}
	if cancellation.Event.FrozenCoreDigest != testB || cancellation.Event.ProviderLoweredToolsDigest != testE {
		t.Fatalf("terminal digest carriage = %+v", cancellation.Event)
	}
}

func TestRawClosureNeverMintsCancellation(t *testing.T) {
	manager := New()
	active, err := manager.Begin(context.Background(), "attempt-loss", 7, testB, testE)
	if err != nil {
		t.Fatal(err)
	}
	active.MarkInvoked()
	if err := manager.AbortForLoss("attempt-loss"); err != nil {
		t.Fatal(err)
	}
	select {
	case <-active.Context().Done():
	default:
		t.Fatal("loss did not abort live work")
	}
	if _, ok := manager.Cancellation("attempt-loss"); ok {
		t.Fatal("raw closure was upgraded into a cancellation fact")
	}
}

func TestTryMarkInvokedRefusesAnAlreadyCancelledAttempt(t *testing.T) {
	manager := New()
	active, err := manager.Begin(context.Background(), "attempt-race", 7, testB, testE)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := manager.Cancel(Intent{AttemptID: "attempt-race", TurnEpoch: 7}, 7); err != nil {
		t.Fatal(err)
	}
	if active.TryMarkInvoked() {
		t.Fatal("cancelled attempt crossed the provider invocation boundary")
	}
}

func TestCancelNeverClassifiesPreTransportWhenInvocationGateSucceeds(t *testing.T) {
	for iteration := range 64 {
		manager := New()
		attemptID := fmt.Sprintf("attempt-race-%d", iteration)
		active, err := manager.Begin(context.Background(), attemptID, 7, testB, testE)
		if err != nil {
			t.Fatal(err)
		}

		originalCancel := active.cancel
		cancelEntered := make(chan struct{})
		releaseCancel := make(chan struct{})
		active.cancel = func() {
			close(cancelEntered)
			<-releaseCancel
			originalCancel()
		}
		type cancelOutcome struct {
			cancellation Cancellation
			err          error
		}
		cancelled := make(chan cancelOutcome, 1)
		go func() {
			cancellation, err := manager.Cancel(Intent{AttemptID: attemptID, TurnEpoch: 7}, 7)
			cancelled <- cancelOutcome{cancellation: cancellation, err: err}
		}()

		select {
		case <-cancelEntered:
		case <-time.After(time.Second):
			t.Fatalf("iteration %d: Cancel did not reach the cancellation boundary", iteration)
		}
		invoked := active.TryMarkInvoked()
		close(releaseCancel)
		got := <-cancelled
		if got.err != nil {
			t.Fatalf("iteration %d: Cancel() error = %v", iteration, got.err)
		}
		if invoked && got.cancellation.Outcome.CancelPoint == outcome.PreTransport {
			t.Fatalf("iteration %d: successful invocation gate recorded as pre_transport", iteration)
		}
	}
}

func TestDuplicateIntentIsEquivalentAndEpochConflictsAreInert(t *testing.T) {
	manager := New()
	if _, err := manager.Begin(context.Background(), "attempt-dup", 9, testB, testE); err != nil {
		t.Fatal(err)
	}
	first, err := manager.Cancel(Intent{AttemptID: "attempt-dup", TurnEpoch: 9}, 9)
	if err != nil {
		t.Fatal(err)
	}
	duplicate, err := manager.Cancel(Intent{AttemptID: "attempt-dup", TurnEpoch: 9}, 9)
	if err != nil || !duplicate.Duplicate || !reflect.DeepEqual(first.Event, duplicate.Event) || !reflect.DeepEqual(first.Outcome, duplicate.Outcome) {
		t.Fatalf("duplicate = %+v, %v; first=%+v", duplicate, err, first)
	}
	if _, err := manager.Cancel(Intent{AttemptID: "attempt-dup", TurnEpoch: 8}, 9); !errors.Is(err, ErrStaleEpoch) {
		t.Fatalf("stale duplicate error = %v", err)
	}
	if _, err := manager.Cancel(Intent{AttemptID: "attempt-dup", TurnEpoch: 10}, 9); !errors.Is(err, ErrEpochAhead) {
		t.Fatalf("ahead duplicate error = %v", err)
	}
}

func TestCancelIntentIsClosedTypedShape(t *testing.T) {
	intent, err := ParseIntent([]byte(`{"attempt_id":"attempt-1","turn_epoch":"7","type":"cancel_attempt"}`))
	if err != nil || intent.AttemptID != "attempt-1" || intent.TurnEpoch != 7 {
		t.Fatalf("ParseIntent() = %+v, %v", intent, err)
	}
	for _, raw := range []string{
		`{"attempt_id":"attempt-1","turn_epoch":"7"}`,
		`{"attempt_id":"attempt-1","turn_epoch":"7","type":"cancel_attempt","unknown":true}`,
		`{"attempt_id":"attempt-1","turn_epoch":7,"type":"cancel_attempt"}`,
	} {
		if _, err := ParseIntent([]byte(raw)); err == nil {
			t.Fatalf("malformed cancel intent accepted: %s", raw)
		}
	}
}
