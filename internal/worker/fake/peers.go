// Package fake provides in-process m-10, broker, and m-8 counterparts for the
// worker's E2 tests. It performs no provider egress and owns no real authority.
package fake

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackli/frank/internal/worker/executor"
	"github.com/jackli/frank/internal/worker/provider"
	workerruntime "github.com/jackli/frank/internal/worker/runtime"
	"github.com/jackli/frank/internal/worker/turn"
)

type traceEvent struct {
	order uint64
	name  string
}

var traceSequence atomic.Uint64

type traceLog struct {
	mu     sync.Mutex
	events []traceEvent
}

func (log *traceLog) add(name string) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.events = append(log.events, traceEvent{order: traceSequence.Add(1), name: name})
}

func (log *traceLog) snapshot() []traceEvent {
	log.mu.Lock()
	defer log.mu.Unlock()
	return append([]traceEvent(nil), log.events...)
}

type ticket struct {
	identity executor.FullIdentity
	state    executor.TicketState
}

type M10 struct {
	mu             sync.Mutex
	Assignment     workerruntime.Assignment
	AuthorizeCalls int
	ConsumeCalls   int
	Outcomes       []executor.OutcomeRecord
	Wakes          []string
	Terminal       turn.Terminal
	tickets        map[string]*ticket
	nextTicket     int
	wakeSet        map[string]struct{}
	trace          traceLog
}

func NewM10(assignment workerruntime.Assignment) *M10 {
	return &M10{Assignment: assignment, tickets: make(map[string]*ticket), wakeSet: make(map[string]struct{})}
}

func (peer *M10) Hello(context.Context, workerruntime.Hello) (workerruntime.Assignment, error) {
	peer.trace.add("hello")
	return peer.Assignment, nil
}

func (peer *M10) ReportAttach(_ context.Context, _ string, _ string, result workerruntime.AttachResult) error {
	peer.trace.add("attach_result:" + string(result))
	return nil
}

func (peer *M10) WakeForward(_ context.Context, relayID string) error {
	peer.mu.Lock()
	defer peer.mu.Unlock()
	if _, exists := peer.wakeSet[relayID]; exists {
		return nil
	}
	peer.wakeSet[relayID] = struct{}{}
	peer.Wakes = append(peer.Wakes, relayID)
	peer.trace.add("wake_forward")
	return nil
}

func (peer *M10) TurnTerminal(_ context.Context, terminal turn.Terminal) error {
	peer.mu.Lock()
	peer.Terminal = terminal
	peer.mu.Unlock()
	peer.trace.add("turn_terminal:" + string(terminal))
	return nil
}

func (peer *M10) AttemptOpen(context.Context, provider.Request) error {
	peer.trace.add("attempt_open")
	return nil
}

func (peer *M10) RecordStreamEnd(_ context.Context, _ string, _ provider.StreamEnd) error {
	peer.trace.add("stream_end")
	return nil
}

func (peer *M10) Authorize(_ context.Context, request executor.AuthorizeRequest) (executor.AuthorizeReply, error) {
	peer.mu.Lock()
	defer peer.mu.Unlock()
	peer.AuthorizeCalls++
	peer.nextTicket++
	id := "ticket-" + strconv.Itoa(peer.nextTicket)
	peer.tickets[id] = &ticket{identity: request.Identity, state: executor.TicketIssued}
	peer.trace.add("authorize")
	return executor.AuthorizeReply{Code: executor.AuthorizeGranted, TicketID: id}, nil
}

func (peer *M10) Consume(_ context.Context, request executor.ConsumeRequest) (executor.ConsumeReply, error) {
	peer.mu.Lock()
	defer peer.mu.Unlock()
	peer.ConsumeCalls++
	ticket := peer.tickets[request.TicketID]
	if ticket == nil {
		return executor.ConsumeReply{}, errors.New("fake m10: unknown ticket")
	}
	if ticket.state != executor.TicketIssued {
		return executor.ConsumeReply{Code: executor.ConsumeDuplicate}, nil
	}
	if request.TurnEpoch != ticket.identity.TurnEpoch || request.CanonicalToolName != ticket.identity.CanonicalToolName || request.CanonicalArgsDigest != ticket.identity.CanonicalArgsDigest {
		return executor.ConsumeReply{Code: executor.ConsumeIdentityMismatch}, nil
	}
	ticket.state = executor.TicketConsumed
	peer.trace.add("consume")
	return executor.ConsumeReply{Code: executor.ConsumeOK}, nil
}

func (peer *M10) RecordOutcome(_ context.Context, outcome executor.OutcomeRecord) error {
	if err := outcome.Validate(); err != nil {
		return err
	}
	peer.mu.Lock()
	defer peer.mu.Unlock()
	ticket := peer.tickets[outcome.TicketID]
	if ticket == nil || ticket.state != executor.TicketConsumed {
		return errors.New("fake m10: outcome without consumed ticket")
	}
	if outcome.Outcome == executor.OutcomeExecuted && *outcome.InvocationIdentity != ticket.identity.Identity {
		return errors.New("fake m10: invocation identity mismatch")
	}
	ticket.state = executor.TicketOutcomeRecorded
	peer.Outcomes = append(peer.Outcomes, outcome)
	peer.trace.add("record_outcome")
	return nil
}

type Broker struct {
	Result     workerruntime.AttachResult
	Capability string
	Wakes      []string
	Err        error
	trace      traceLog
}

func (broker *Broker) Attach(context.Context, string, workerruntime.AttachTuple) (workerruntime.AttachResult, string, error) {
	broker.trace.add("attach")
	return broker.Result, broker.Capability, broker.Err
}

func (broker *Broker) Rediscover(context.Context, string) ([]string, error) {
	broker.trace.add("rediscover")
	return append([]string(nil), broker.Wakes...), nil
}

type M8 struct {
	Disposition  provider.Disposition
	OpaqueItems  [][]byte
	ScriptedTool workerruntime.ToolCall
	Err          error
	Attempts     int
	trace        traceLog
}

func (peer *M8) Attempt(context.Context, provider.Request) (provider.Disposition, []json.RawMessage, error) {
	peer.Attempts++
	peer.trace.add("provider_attempt")
	items := make([]json.RawMessage, len(peer.OpaqueItems))
	for index := range peer.OpaqueItems {
		items[index] = append(json.RawMessage(nil), peer.OpaqueItems[index]...)
	}
	return peer.Disposition, items, peer.Err
}

func (peer *M8) Cancel(context.Context, string, uint64) (provider.Disposition, error) {
	return provider.CancelledPre, nil
}

func (peer *M8) NextToolCall() (workerruntime.ToolCall, error) {
	if peer.ScriptedTool.ID == "" {
		return workerruntime.ToolCall{}, errors.New("fake m8: no scripted tool call")
	}
	call := peer.ScriptedTool
	call.Arguments = append([]byte(nil), call.Arguments...)
	return call, nil
}

// Backend is an in-memory local-tool backend with effect counters.
type Backend struct {
	mu     sync.Mutex
	Files  map[string]string
	Writes int
}

func NewBackend() *Backend { return &Backend{Files: make(map[string]string)} }

func (backend *Backend) Read(_ context.Context, path string, _, _ int64) (string, error) {
	backend.mu.Lock()
	defer backend.mu.Unlock()
	value, exists := backend.Files[path]
	if !exists {
		return "", errors.New("fake backend: file absent")
	}
	return value, nil
}

func (backend *Backend) Write(_ context.Context, path, content string) error {
	backend.mu.Lock()
	defer backend.mu.Unlock()
	backend.Files[path] = content
	backend.Writes++
	return nil
}

func (backend *Backend) Edit(_ context.Context, path, oldText, newText string, _ bool) error {
	backend.mu.Lock()
	defer backend.mu.Unlock()
	if backend.Files[path] != oldText {
		return errors.New("fake backend: edit target mismatch")
	}
	backend.Files[path] = newText
	return nil
}

func (backend *Backend) ApplyPatch(context.Context, string) error { return nil }

func (backend *Backend) Bash(context.Context, string, time.Duration) (string, error) {
	return "", nil
}

func (backend *Backend) BoundOutput(value string) string { return value }

type ObjectiveSource struct {
	Relays map[string]string
	Reads  int
}

func (source *ObjectiveSource) ResolveObjective(_ context.Context, relayID string) (string, error) {
	source.Reads++
	value, exists := source.Relays[relayID]
	if !exists {
		return "", errors.New("fake seat objective: relay absent")
	}
	return value, nil
}

func JoinTrace(control *M10, broker *Broker, providerPeer *M8) []string {
	events := append(control.trace.snapshot(), broker.trace.snapshot()...)
	events = append(events, providerPeer.trace.snapshot()...)
	sort.Slice(events, func(left, right int) bool { return events[left].order < events[right].order })
	result := make([]string, len(events))
	for index := range events {
		result[index] = events[index].name
	}
	return result
}
