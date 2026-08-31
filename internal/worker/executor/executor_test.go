package executor

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/worker/catalog"
)

func TestPreparedCallIsInertWithoutAuthorityPath(t *testing.T) {
	source := &mutableArguments{value: []byte(`{"path":"a"}`)}
	call, err := Prepare("run-1", "turn-1", "call-1", "read", "7", source)
	if err != nil {
		t.Fatal(err)
	}
	if call.FrozenIdentity().CanonicalArgsDigest == "" {
		t.Fatal("prepared call did not freeze identity")
	}
	if source.reads != 1 {
		t.Fatalf("argument snapshots = %d, want 1", source.reads)
	}
	// There is deliberately no invocation method on PreparedCall. Construction
	// has no authority or backend and therefore cannot produce an effect.
}

func TestMutationBeforeConsumeReturnsIdentityMismatchAndLeavesTicketIssued(t *testing.T) {
	source := &mutableArguments{value: []byte(`{"path":"a"}`)}
	authority := newFakeAuthority()
	authority.afterAuthorize = func() { source.set([]byte(`{"path":"b"}`)) }
	backend := &recordingInvoker{}
	executor := mustExecutor(t, authority, backend)
	call := mustPrepare(t, source)

	result, err := executor.Execute(context.Background(), call)
	assertCode(t, err, CodeIdentityMismatch)
	if result.TicketID == "" {
		t.Fatal("mismatch result omitted issued ticket id")
	}
	if got := authority.state(result.TicketID); got != TicketIssued {
		t.Fatalf("ticket state = %q, want %q", got, TicketIssued)
	}
	if backend.calls() != 0 {
		t.Fatalf("backend calls = %d, want 0", backend.calls())
	}
}

func TestMutationAfterConsumeRecordsIntegrityFaultAndNeverInvokes(t *testing.T) {
	source := &mutableArguments{value: []byte(`{"path":"a"}`)}
	authority := newFakeAuthority()
	authority.afterConsume = func() { source.set([]byte(`{"path":"b"}`)) }
	backend := &recordingInvoker{}
	executor := mustExecutor(t, authority, backend)

	result, err := executor.Execute(context.Background(), mustPrepare(t, source))
	if err != nil {
		t.Fatal(err)
	}
	if result.Outcome != OutcomeNotInvokedIntegrityFault {
		t.Fatalf("outcome = %q", result.Outcome)
	}
	if backend.calls() != 0 {
		t.Fatalf("backend calls = %d, want 0", backend.calls())
	}
	record := authority.outcome(result.TicketID)
	if record.Outcome != OutcomeNotInvokedIntegrityFault {
		t.Fatalf("record outcome = %q", record.Outcome)
	}
	if record.InvocationIdentity != nil {
		t.Fatal("integrity outcome carried invocation identity")
	}
	if record.IntegrityEvidence == nil || record.IntegrityEvidence.Expected == record.IntegrityEvidence.Observed {
		t.Fatal("integrity outcome lacks unequal expected and observed identities")
	}
}

func TestExecutedOutcomeCarriesActualInvocationIdentity(t *testing.T) {
	authority := newFakeAuthority()
	backend := &recordingInvoker{}
	executor := mustExecutor(t, authority, backend)

	result, err := executor.Execute(context.Background(), mustPrepare(t, staticArguments(`{"path":"a"}`)))
	if err != nil {
		t.Fatal(err)
	}
	if result.Outcome != OutcomeExecuted {
		t.Fatalf("outcome = %q", result.Outcome)
	}
	record := authority.outcome(result.TicketID)
	if record.InvocationIdentity == nil {
		t.Fatal("executed outcome omitted invocation identity")
	}
	if *record.InvocationIdentity != result.InvocationIdentity {
		t.Fatalf("record identity = %#v, result = %#v", *record.InvocationIdentity, result.InvocationIdentity)
	}
	if record.IntegrityEvidence != nil {
		t.Fatal("executed outcome carried integrity evidence")
	}
}

func TestOutcomeRecordDomainIsClosed(t *testing.T) {
	identity := Identity{CanonicalToolName: "read", CanonicalArgsDigest: digestOf(t, `{"path":"a"}`), TurnEpoch: "7"}
	tests := []struct {
		name   string
		record OutcomeRecord
		ok     bool
	}{
		{"executed", OutcomeRecord{TicketID: "ticket-1", TurnEpoch: "7", Outcome: OutcomeExecuted, InvocationIdentity: &identity}, true},
		{"integrity", OutcomeRecord{TicketID: "ticket-1", TurnEpoch: "7", Outcome: OutcomeNotInvokedIntegrityFault, IntegrityEvidence: &IntegrityEvidence{Expected: identity, Observed: Identity{CanonicalToolName: "read", CanonicalArgsDigest: digestOf(t, `{"path":"b"}`), TurnEpoch: "7"}}}, true},
		{"unknown member", OutcomeRecord{TicketID: "ticket-1", TurnEpoch: "7", Outcome: "partial"}, false},
		{"executed lacks identity", OutcomeRecord{TicketID: "ticket-1", TurnEpoch: "7", Outcome: OutcomeExecuted}, false},
		{"integrity equal", OutcomeRecord{TicketID: "ticket-1", TurnEpoch: "7", Outcome: OutcomeNotInvokedIntegrityFault, IntegrityEvidence: &IntegrityEvidence{Expected: identity, Observed: identity}}, false},
		{"integrity carries invocation", OutcomeRecord{TicketID: "ticket-1", TurnEpoch: "7", Outcome: OutcomeNotInvokedIntegrityFault, InvocationIdentity: &identity, IntegrityEvidence: &IntegrityEvidence{Expected: identity, Observed: Identity{CanonicalToolName: "write", CanonicalArgsDigest: identity.CanonicalArgsDigest, TurnEpoch: "7"}}}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.record.Validate() == nil; got != test.ok {
				t.Fatalf("Validate success = %v, want %v", got, test.ok)
			}
		})
	}
}

func TestAuthorizeRejectReasonDomainIsExactlyFourLifecycleTokens(t *testing.T) {
	want := []AuthorizeRejectReason{
		RejectRunNotAdmitted,
		RejectTurnInactive,
		RejectLeaseInvalid,
		RejectTurnBudgetExhausted,
	}
	got := LifecycleRejectReasons()
	if len(got) != len(want) {
		t.Fatalf("reason count = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("reason %d = %q, want %q", index, got[index], want[index])
		}
	}
	if (AuthorizeReply{Code: AuthorizeRejected, RejectReason: "expired"}).Validate() == nil {
		t.Fatal("accepted reason outside closed lifecycle domain")
	}
}

func TestAuthorizeRejectsAndEpochMismatchAreAttemptInert(t *testing.T) {
	for _, test := range []struct {
		name      string
		reply     AuthorizeReply
		wantCode  Code
		wantState TicketState
	}{
		{"run not admitted", AuthorizeReply{Code: AuthorizeRejected, RejectReason: RejectRunNotAdmitted}, CodeAuthorizeRejected, ""},
		{"turn inactive", AuthorizeReply{Code: AuthorizeRejected, RejectReason: RejectTurnInactive}, CodeAuthorizeRejected, ""},
		{"lease invalid", AuthorizeReply{Code: AuthorizeRejected, RejectReason: RejectLeaseInvalid}, CodeAuthorizeRejected, ""},
		{"budget", AuthorizeReply{Code: AuthorizeRejected, RejectReason: RejectTurnBudgetExhausted}, CodeAuthorizeRejected, ""},
		{"stale epoch", AuthorizeReply{Code: AuthorizeStaleEpoch}, CodeStaleEpoch, ""},
	} {
		t.Run(test.name, func(t *testing.T) {
			authority := newFakeAuthority()
			authority.authorizeOverride = &test.reply
			backend := &recordingInvoker{}
			_, err := mustExecutor(t, authority, backend).Execute(context.Background(), mustPrepare(t, staticArguments(`{"path":"a"}`)))
			assertCode(t, err, test.wantCode)
			if backend.calls() != 0 {
				t.Fatalf("backend calls = %d, want 0", backend.calls())
			}
		})
	}
}

func TestGrantedEffectDescriptorBindsInvocationBeforeConsume(t *testing.T) {
	for _, test := range []struct {
		name       string
		descriptor *EffectDescriptor
		wantCode   Code
	}{
		{name: "absent", wantCode: CodeProtocolFault},
		{name: "mismatched", descriptor: &EffectDescriptor{Action: "write", CanonicalArgsDigest: digestOf(t, `{"path":"a"}`), BackendID: "in-process", NetworkPolicyID: "none", ToolImplRef: "in-process:write", OneShot: true}, wantCode: CodeIdentityMismatch},
	} {
		t.Run(test.name, func(t *testing.T) {
			authority := newFakeAuthority()
			authority.authorizeOverride = &AuthorizeReply{Code: AuthorizeGranted, TicketID: "ticket", EffectDescriptor: test.descriptor}
			backend := &recordingInvoker{}
			_, err := mustExecutor(t, authority, backend).Execute(context.Background(), mustPrepare(t, staticArguments(`{"path":"a"}`)))
			assertCode(t, err, test.wantCode)
			if backend.calls() != 0 || authority.consumeCalls != 0 {
				t.Fatalf("forbidden path advanced: backend=%d consume=%d", backend.calls(), authority.consumeCalls)
			}
		})
	}
}

func TestEveryCatalogToolUsesUniformAuthorityPath(t *testing.T) {
	for _, tool := range catalog.ExpectedIdentities() {
		t.Run(tool.CanonicalName, func(t *testing.T) {
			authority := newFakeAuthority()
			backend := &recordingInvoker{}
			call, err := Prepare("run-1", "turn-1", "call-1", tool.CanonicalName, "7", staticArguments(`{}`))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := mustExecutor(t, authority, backend).Execute(context.Background(), call); err != nil {
				t.Fatal(err)
			}
			if authority.authorizeCalls != 1 || authority.consumeCalls != 1 || len(authority.outcomes) != 1 {
				t.Fatalf("path counts = authorize %d, consume %d, outcomes %d", authority.authorizeCalls, authority.consumeCalls, len(authority.outcomes))
			}
			if backend.calls() != 1 {
				t.Fatalf("backend calls = %d, want 1", backend.calls())
			}
		})
	}
}

func TestDoubleConsumeIsRejected(t *testing.T) {
	authority := newFakeAuthority()
	identity := FullIdentity{RunID: "run-1", TurnID: "turn-1", ToolCallID: "call-1", Identity: Identity{CanonicalToolName: "read", CanonicalArgsDigest: digestOf(t, `{"path":"a"}`), TurnEpoch: "7"}}
	reply, err := authority.Authorize(context.Background(), AuthorizeRequest{Identity: identity})
	if err != nil {
		t.Fatal(err)
	}
	request := ConsumeRequest{TicketID: reply.TicketID, TurnEpoch: "7", CanonicalToolName: "read", CanonicalArgsDigest: identity.CanonicalArgsDigest}
	if got, err := authority.Consume(context.Background(), request); err != nil || got.Code != ConsumeOK {
		t.Fatalf("first consume = %#v, %v", got, err)
	}
	if got, err := authority.Consume(context.Background(), request); err != nil || got.Code != ConsumeDuplicate {
		t.Fatalf("second consume = %#v, %v", got, err)
	}
}

func TestConsumeEpochMismatchIsAttemptInert(t *testing.T) {
	for _, test := range []struct {
		name      string
		epoch     string
		wantCode  ConsumeCode
		wantError bool
	}{
		{"stale", "6", ConsumeStaleEpoch, false},
		{"future", "8", "", true},
	} {
		t.Run(test.name, func(t *testing.T) {
			authority := newFakeAuthority()
			identity := FullIdentity{RunID: "run-1", TurnID: "turn-1", ToolCallID: "call-1", Identity: Identity{CanonicalToolName: "read", CanonicalArgsDigest: digestOf(t, `{}`), TurnEpoch: "7"}}
			reply, err := authority.Authorize(context.Background(), AuthorizeRequest{Identity: identity})
			if err != nil {
				t.Fatal(err)
			}
			got, err := authority.Consume(context.Background(), ConsumeRequest{TicketID: reply.TicketID, TurnEpoch: test.epoch, CanonicalToolName: "read", CanonicalArgsDigest: identity.CanonicalArgsDigest})
			if (err != nil) != test.wantError {
				t.Fatalf("error = %v, wantError %v", err, test.wantError)
			}
			if got.Code != test.wantCode {
				t.Fatalf("code = %q, want %q", got.Code, test.wantCode)
			}
			if state := authority.state(reply.TicketID); state != TicketIssued {
				t.Fatalf("state = %q, want ISSUED", state)
			}
		})
	}
}

func TestCancellationAndEOFBeforeConsumeFailClosed(t *testing.T) {
	for _, failure := range []error{context.Canceled, errors.New("EOF")} {
		t.Run(failure.Error(), func(t *testing.T) {
			authority := newFakeAuthority()
			authority.consumeErr = failure
			backend := &recordingInvoker{}
			result, err := mustExecutor(t, authority, backend).Execute(context.Background(), mustPrepare(t, staticArguments(`{"path":"a"}`)))
			if !errors.Is(err, failure) {
				t.Fatalf("error = %v, want %v", err, failure)
			}
			if result.Outcome != "" || backend.calls() != 0 {
				t.Fatalf("attempt was not inert: %#v, calls %d", result, backend.calls())
			}
		})
	}
}

func TestExecutorSerializesSettlement(t *testing.T) {
	authority := newFakeAuthority()
	backend := &recordingInvoker{block: make(chan struct{}), entered: make(chan struct{}, 2)}
	executor := mustExecutor(t, authority, backend)

	var wait sync.WaitGroup
	for index := 0; index < 2; index++ {
		wait.Add(1)
		go func(index int) {
			defer wait.Done()
			call, err := Prepare("run-1", "turn-1", string(rune('a'+index)), "read", "7", staticArguments(`{"path":"a"}`))
			if err != nil {
				t.Errorf("prepare: %v", err)
				return
			}
			if _, err := executor.Execute(context.Background(), call); err != nil {
				t.Errorf("execute: %v", err)
			}
		}(index)
	}
	select {
	case <-backend.entered:
	case <-time.After(time.Second):
		t.Fatal("first invocation did not start")
	}
	select {
	case <-backend.entered:
		t.Fatal("second invocation entered while first settlement was active")
	case <-time.After(50 * time.Millisecond):
	}
	close(backend.block)
	wait.Wait()
	if got := atomic.LoadInt32(&backend.maxActive); got != 1 {
		t.Fatalf("maximum concurrent invocations = %d, want 1", got)
	}
}

func TestAuthorityCrashWindowsAreTotal(t *testing.T) {
	t.Run("issued becomes void", func(t *testing.T) {
		authority := newFakeAuthority()
		identity := FullIdentity{RunID: "run-1", TurnID: "turn-1", ToolCallID: "call-1", Identity: Identity{CanonicalToolName: "read", CanonicalArgsDigest: digestOf(t, `{}`), TurnEpoch: "7"}}
		reply, err := authority.Authorize(context.Background(), AuthorizeRequest{Identity: identity})
		if err != nil {
			t.Fatal(err)
		}
		authority.retireEpoch("7")
		if got := authority.state(reply.TicketID); got != TicketVoid {
			t.Fatalf("state = %q, want VOID", got)
		}
	})
	t.Run("consumed without outcome becomes unknown", func(t *testing.T) {
		authority := newFakeAuthority()
		identity := FullIdentity{RunID: "run-1", TurnID: "turn-1", ToolCallID: "call-1", Identity: Identity{CanonicalToolName: "read", CanonicalArgsDigest: digestOf(t, `{}`), TurnEpoch: "7"}}
		reply, _ := authority.Authorize(context.Background(), AuthorizeRequest{Identity: identity})
		_, _ = authority.Consume(context.Background(), ConsumeRequest{TicketID: reply.TicketID, TurnEpoch: "7", CanonicalToolName: "read", CanonicalArgsDigest: identity.CanonicalArgsDigest})
		authority.retireEpoch("7")
		if got := authority.state(reply.TicketID); got != TicketUnknownToolOutcome {
			t.Fatalf("state = %q, want UNKNOWN_TOOL_OUTCOME", got)
		}
	})
	t.Run("outcome send lost before persistence becomes unknown", func(t *testing.T) {
		authority := newFakeAuthority()
		authority.recordErrBefore = errors.New("record send EOF")
		backend := &recordingInvoker{}
		result, err := mustExecutor(t, authority, backend).Execute(context.Background(), mustPrepare(t, staticArguments(`{}`)))
		if err == nil {
			t.Fatal("outcome send unexpectedly succeeded")
		}
		if state := authority.state(result.TicketID); state != TicketConsumed {
			t.Fatalf("pre-retirement state = %q, want CONSUMED", state)
		}
		authority.retireEpoch("7")
		if state := authority.state(result.TicketID); state != TicketUnknownToolOutcome {
			t.Fatalf("retired state = %q, want UNKNOWN_TOOL_OUTCOME", state)
		}
	})
	t.Run("partial send after persistence keeps terminal outcome", func(t *testing.T) {
		authority := newFakeAuthority()
		authority.recordErrAfter = errors.New("connection lost after persistence")
		result, err := mustExecutor(t, authority, &recordingInvoker{}).Execute(context.Background(), mustPrepare(t, staticArguments(`{}`)))
		if err == nil {
			t.Fatal("partial send unexpectedly looked acknowledged")
		}
		if state := authority.state(result.TicketID); state != TicketOutcomeRecorded {
			t.Fatalf("state = %q, want OUTCOME_RECORDED", state)
		}
		authority.retireEpoch("7")
		if state := authority.state(result.TicketID); state != TicketOutcomeRecorded {
			t.Fatalf("terminal state changed on retirement: %q", state)
		}
	})
}

type mutableArguments struct {
	mu    sync.Mutex
	value []byte
	reads int
}

func (source *mutableArguments) Snapshot() []byte {
	source.mu.Lock()
	defer source.mu.Unlock()
	source.reads++
	return append([]byte(nil), source.value...)
}

func (source *mutableArguments) set(value []byte) {
	source.mu.Lock()
	defer source.mu.Unlock()
	source.value = append(source.value[:0], value...)
}

type staticArguments string

func (source staticArguments) Snapshot() []byte { return []byte(source) }

type recordingInvoker struct {
	mu        sync.Mutex
	count     int
	active    int32
	maxActive int32
	block     chan struct{}
	entered   chan struct{}
}

func (backend *recordingInvoker) Invoke(_ context.Context, invocation Invocation) (any, error) {
	backend.mu.Lock()
	backend.count++
	backend.mu.Unlock()
	active := atomic.AddInt32(&backend.active, 1)
	for {
		maximum := atomic.LoadInt32(&backend.maxActive)
		if active <= maximum || atomic.CompareAndSwapInt32(&backend.maxActive, maximum, active) {
			break
		}
	}
	defer atomic.AddInt32(&backend.active, -1)
	if backend.entered != nil {
		backend.entered <- struct{}{}
	}
	if backend.block != nil {
		<-backend.block
	}
	return invocation.Identity, nil
}

func (backend *recordingInvoker) calls() int {
	backend.mu.Lock()
	defer backend.mu.Unlock()
	return backend.count
}

func mustPrepare(t *testing.T, source ArgumentSource) *PreparedCall {
	t.Helper()
	call, err := Prepare("run-1", "turn-1", "call-1", "read", "7", source)
	if err != nil {
		t.Fatal(err)
	}
	return call
}

func mustExecutor(t *testing.T, authority Authority, backend Invoker) *Executor {
	t.Helper()
	executor, err := New(authority, backend)
	if err != nil {
		t.Fatal(err)
	}
	return executor
}

func digestOf(t *testing.T, arguments string) string {
	t.Helper()
	return mustPrepare(t, staticArguments(arguments)).FrozenIdentity().CanonicalArgsDigest
}

func assertCode(t *testing.T, err error, want Code) {
	t.Helper()
	var executionError *Error
	if !errors.As(err, &executionError) || executionError.Code != want {
		t.Fatalf("error = %v, want code %q", err, want)
	}
}
