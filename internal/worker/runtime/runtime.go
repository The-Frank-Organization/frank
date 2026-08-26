// Package runtime wires the m-9 worker's bounded governed-turn components.
// App-side transports are injected; this package does not own provider wire,
// credentials, authority policy, or conductor storage.
package runtime

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jackli/frank/internal/appipc"
	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/worker/catalog"
	"github.com/jackli/frank/internal/worker/contextmgr"
	"github.com/jackli/frank/internal/worker/executor"
	workerjcs "github.com/jackli/frank/internal/worker/jcs"
	"github.com/jackli/frank/internal/worker/journal"
	"github.com/jackli/frank/internal/worker/provider"
	"github.com/jackli/frank/internal/worker/relaytool"
	"github.com/jackli/frank/internal/worker/tools"
	"github.com/jackli/frank/internal/worker/turn"
	"github.com/jackli/frank/internal/worker/wire"
)

type AttachResult string

const (
	AttachSuspended     AttachResult = "broker:attach-suspended"
	AttachOK            AttachResult = "attach-ok"
	AttachTupleMismatch AttachResult = "broker:attach-tuple-mismatch"
)

type Hello struct {
	PID       int
	BuildInfo string
}

type Assignment struct {
	RunID              string
	TurnID             string
	TurnEpoch          string
	ManifestDigest     string
	GenerationID       string
	CreateAuthID       string
	BrokerEndpoint     string
	SessionLogPath     string `json:"session_log_path"`
	AdmissionRef       turn.AdmissionRef
	ParkedUnknown      []turn.ParkedUnknown
	SettlementManifest *appipc.SettlementManifest `json:"settlement_manifest"`
	PredecessorTurnID  *string                    `json:"predecessor_turn_id"`
}

type AttachTuple struct {
	RunID        string
	GenerationID string
	TurnEpoch    string
}

type Control interface {
	executor.Authority
	provider.Gate
	Hello(context.Context, Hello) (Assignment, error)
	GenesisCommitted(context.Context, appipc.GenesisCommittedBody) error
	ReportAttach(context.Context, string, string, AttachResult) error
	EvaluateProviderE0(context.Context, []provider.Event) error
	WakeForward(context.Context, string) error
	TurnTerminal(context.Context, turn.Terminal) error
}

type Broker interface {
	Attach(context.Context, string, AttachTuple) (AttachResult, string, error)
	Rediscover(context.Context, string) ([]string, error)
}

type Provider interface {
	provider.Connector
	NextToolCall() (ToolCall, error)
}

type ObjectiveResolver interface {
	ResolveObjective(context.Context, string) (string, error)
}

type ToolCall struct {
	ID            string
	CanonicalName string
	Arguments     []byte
}

type Config struct {
	PID            int
	BuildInfo      string
	RuntimeDir     string
	RunDisposition string
	AttachDeadline time.Duration
	AttachBackoff  time.Duration
}

type Result struct {
	Terminal            turn.Terminal
	PersistedTranscript []byte
	ReplayedTranscript  []byte
}

type Runner struct {
	Control   Control
	Broker    Broker
	Provider  Provider
	Backend   tools.Backend
	Conductor relaytool.Conductor
	Objective ObjectiveResolver
}

func (runner Runner) Run(ctx context.Context, config Config) (Result, error) {
	if runner.Control == nil || runner.Broker == nil || runner.Provider == nil || runner.Backend == nil {
		return Result{}, errors.New("worker runtime: incomplete dependency set")
	}
	if config.PID <= 0 || config.BuildInfo == "" || config.RuntimeDir == "" || config.AttachDeadline <= 0 || config.AttachBackoff <= 0 {
		return Result{}, errors.New("worker runtime: invalid configuration")
	}
	if digest, err := catalog.Digest(catalog.ExpectedIdentities()); err != nil || digest != catalog.ExpectedDigest {
		return Result{}, errors.New("worker runtime: tool catalog integrity failure")
	}

	assignment, err := runner.Control.Hello(ctx, Hello{PID: config.PID, BuildInfo: config.BuildInfo})
	if err != nil {
		return Result{}, fmt.Errorf("worker hello: %w", err)
	}
	epoch, err := validateAssignment(assignment)
	if err != nil {
		return Result{}, err
	}
	capability, err := runner.attach(ctx, assignment, config)
	if err != nil {
		return Result{}, err
	}
	if err := runner.rediscover(ctx, capability); err != nil {
		return Result{}, err
	}

	machine := turn.New()
	if err := machine.Admit(turn.Open{RunID: assignment.RunID, TurnID: assignment.TurnID, TurnEpoch: assignment.TurnEpoch, AdmissionRef: assignment.AdmissionRef, ParkedUnknown: assignment.ParkedUnknown}, time.Now()); err != nil {
		return Result{}, err
	}
	runtimeDir := config.RuntimeDir
	if assignment.SessionLogPath != "" {
		if filepath.Base(assignment.SessionLogPath) != journal.SessionLogName {
			return Result{}, errors.New("worker runtime: invalid session log path")
		}
		runtimeDir = filepath.Dir(assignment.SessionLogPath)
	}
	writer, err := journal.Open(journal.Config{
		RuntimeDir: runtimeDir, RunDisposition: config.RunDisposition,
		Identity:     journal.Identity{RunID: assignment.RunID, RunManifestDigest: assignment.ManifestDigest, CreateAuthID: assignment.CreateAuthID},
		GenerationID: assignment.GenerationID, TurnEpoch: assignment.TurnEpoch,
		OnGenesisDurable: func(commit journal.GenesisCommit) error {
			return runner.Control.GenesisCommitted(ctx, appipc.GenesisCommittedBody{GenerationID: commit.GenerationID})
		},
	})
	if err != nil {
		return Result{}, err
	}
	closed := false
	defer func() {
		if !closed {
			_ = writer.Close()
		}
	}()

	if err := machine.BeginAssembly(epoch); err != nil {
		return Result{}, err
	}
	objective, err := taskText(ctx, runner.Objective, assignment.AdmissionRef)
	if err != nil {
		return Result{}, err
	}
	manager, err := contextmgr.New([]contextmgr.Item{
		{Kind: "system", Content: json.RawMessage(`"You are a governed coding worker."`)},
		{Kind: "objective", Content: mustJSON(objective)},
	})
	if err != nil {
		return Result{}, err
	}
	if _, _, err := manager.Assemble(); err != nil {
		return Result{}, err
	}
	attemptID := "attempt-1"
	opaqueRequest, err := marshalProviderRequest(assignment, attemptID, objective)
	if err != nil {
		return Result{}, err
	}
	if err := machine.AttemptOpenOK(epoch, assignment.ParkedUnknown); err != nil {
		return Result{}, err
	}
	cycle, err := provider.New(runner.Control, runner.Provider, assignment.TurnEpoch)
	if err != nil {
		return Result{}, err
	}
	providerResult, err := cycle.Run(ctx, provider.Request{AttemptID: attemptID, TurnID: assignment.TurnID, TurnEpoch: assignment.TurnEpoch, ProviderLane: "assigned", OpaqueRequest: opaqueRequest})
	if err != nil {
		return Result{}, errors.Join(errors.New("worker runtime: provider attempt failed"), err)
	}
	if err := runner.Control.EvaluateProviderE0(ctx, providerResult.Events); err != nil {
		return Result{}, fmt.Errorf("worker runtime: provider E0 evaluation: %w", err)
	}
	if providerResult.Disposition != provider.Completed {
		return Result{}, errors.Join(errors.New("worker runtime: provider attempt did not complete"), err)
	}
	if err := machine.Observe(epoch); err != nil {
		return Result{}, err
	}
	providerMembers := []journal.Record{{Kind: journal.KindInputItem, Fields: map[string]json.RawMessage{
		"role": rawString("user"), "item_index": rawString("0"), "content": mustJSON(objective),
	}}}
	for index, event := range providerResult.Events {
		// The normalized response-item interior remains opaque: exact JCS rides
		// verbatim and every other byte sequence rides the closed base64 branch.
		fields := map[string]json.RawMessage{
			"attempt_id": rawString(attemptID), "item_index": rawString(wire.FormatCounter(uint64(index))),
		}
		for member, raw := range journal.ProviderItemCarrier(event.Opaque) {
			fields[member] = raw
		}
		if usage := providerUsage(event); usage != nil {
			fields["usage"] = usage
		}
		providerMembers = append(providerMembers, journal.Record{Kind: journal.KindProviderOutput, Fields: fields})
	}
	if _, err := writer.AppendRound(assignment.TurnID, "0", providerMembers); err != nil {
		return Result{}, err
	}

	var call ToolCall
	normalizedCall, err := provider.ToolCallFromEvents(providerResult.Events)
	if err != nil {
		return Result{}, err
	}
	if normalizedCall != nil {
		call = ToolCall{ID: normalizedCall.ID, CanonicalName: normalizedCall.CanonicalName, Arguments: append([]byte(nil), normalizedCall.Arguments...)}
	} else {
		call, err = runner.Provider.NextToolCall()
		if err != nil {
			return Result{}, err
		}
	}
	if call.ID == "" || call.CanonicalName == "" || len(call.Arguments) == 0 {
		return Result{}, errors.New("worker runtime: malformed normalized tool request")
	}
	if err := machine.ToolRound(epoch, call.CanonicalName+"\x00"+string(call.Arguments), false); err != nil {
		return Result{}, err
	}
	source := &argumentBytes{value: append([]byte(nil), call.Arguments...)}
	prepared, err := executor.Prepare(assignment.RunID, assignment.TurnID, call.ID, call.CanonicalName, assignment.TurnEpoch, source)
	if err != nil {
		return Result{}, err
	}
	invoker := &durableInvoker{delegate: newToolInvoker(runner.Backend, runner.Conductor), writer: writer, turnID: assignment.TurnID, roundIndex: "1", call: call}
	toolExecutor, err := executor.New(runner.Control, invoker)
	if err != nil {
		return Result{}, err
	}
	if _, err := toolExecutor.Execute(ctx, prepared); err != nil {
		return Result{}, err
	}
	if err := machine.Finish(epoch, turn.TurnCompleted); err != nil {
		return Result{}, err
	}
	if err := runner.Control.TurnTerminal(ctx, turn.TurnCompleted); err != nil {
		return Result{}, err
	}
	if err := writer.Close(); err != nil {
		return Result{}, err
	}
	closed = true
	persisted, err := os.ReadFile(filepath.Join(runtimeDir, journal.SessionLogName))
	if err != nil {
		return Result{}, err
	}
	recovery := journal.Recover(persisted, journal.Identity{RunID: assignment.RunID, RunManifestDigest: assignment.ManifestDigest, CreateAuthID: assignment.CreateAuthID})
	if recovery.Disposition != journal.DispositionResumable || recovery.Boundary.Offset != int64(len(persisted)) {
		return Result{}, errors.New("worker runtime: persisted transcript is not fully replayable")
	}
	replayed := append([]byte(nil), persisted[:recovery.Boundary.Offset]...)
	return Result{Terminal: turn.TurnCompleted, PersistedTranscript: persisted, ReplayedTranscript: replayed}, nil
}

func providerUsage(event provider.Event) json.RawMessage {
	var envelope struct {
		Kind string          `json:"kind"`
		Type string          `json:"type"`
		Body json.RawMessage `json:"body"`
	}
	if err := json.Unmarshal(event.Opaque, &envelope); err != nil {
		return nil
	}
	if envelope.Kind == "" {
		envelope.Kind = envelope.Type
	}
	if envelope.Kind != "usage" || len(envelope.Body) == 0 {
		return nil
	}
	return append(json.RawMessage(nil), envelope.Body...)
}

func marshalProviderRequest(assignment Assignment, attemptID, objective string) ([]byte, error) {
	request := struct {
		Schema         string           `json:"schema"`
		RunID          string           `json:"run_id"`
		TurnID         string           `json:"turn_id"`
		AttemptID      string           `json:"attempt_id"`
		TurnEpoch      string           `json:"turn_epoch"`
		ProviderLaneID string           `json:"provider_lane_id"`
		Instructions   string           `json:"instructions"`
		Input          []map[string]any `json:"input"`
		Tools          []map[string]any `json:"tools"`
		Sampling       struct {
			MaxOutputTokens int64 `json:"max_output_tokens"`
		} `json:"sampling"`
		Reasoning map[string]any `json:"reasoning"`
	}{
		Schema: "m8.llm_request.v1", RunID: assignment.RunID, TurnID: assignment.TurnID,
		AttemptID: attemptID, TurnEpoch: assignment.TurnEpoch, ProviderLaneID: "assigned",
		Instructions: "You are a governed coding worker.",
		Input:        []map[string]any{{"kind": "user_text", "text": objective}},
		Tools:        []map[string]any{}, Reasoning: map[string]any{},
	}
	request.Sampling.MaxOutputTokens = 4096
	encoded, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return workerjcs.Canonicalize(encoded)
}

func (runner Runner) attach(ctx context.Context, assignment Assignment, config Config) (string, error) {
	deadline := time.Now().Add(config.AttachDeadline)
	tuple := AttachTuple{RunID: assignment.RunID, GenerationID: assignment.GenerationID, TurnEpoch: assignment.TurnEpoch}
	for {
		result, capability, err := runner.Broker.Attach(ctx, assignment.BrokerEndpoint, tuple)
		if err != nil {
			return "", fmt.Errorf("worker broker attach: %w", err)
		}
		switch result {
		case AttachSuspended:
			if time.Now().Add(config.AttachBackoff).After(deadline) {
				return "", errors.New("worker broker attach deadline exceeded")
			}
			timer := time.NewTimer(config.AttachBackoff)
			select {
			case <-ctx.Done():
				timer.Stop()
				return "", ctx.Err()
			case <-timer.C:
			}
			continue
		case AttachTupleMismatch:
			_ = runner.Control.ReportAttach(ctx, assignment.GenerationID, assignment.TurnEpoch, result)
			return "", errors.New("worker broker attach tuple mismatch")
		case AttachOK:
			if capability == "" {
				return "", errors.New("worker broker returned an empty capability")
			}
			if err := runner.Control.ReportAttach(ctx, assignment.GenerationID, assignment.TurnEpoch, result); err != nil {
				return "", err
			}
			return capability, nil
		default:
			return "", errors.New("worker broker returned an unknown attach result")
		}
	}
}

func (runner Runner) rediscover(ctx context.Context, capability string) error {
	wakes, err := runner.Broker.Rediscover(ctx, capability)
	if err != nil {
		return fmt.Errorf("worker durable rediscovery: %w", err)
	}
	for _, relayID := range wakes {
		if relayID == "" {
			return errors.New("worker durable rediscovery returned an empty relay id")
		}
		if err := runner.Control.WakeForward(ctx, relayID); err != nil {
			return err
		}
	}
	return nil
}

func validateAssignment(assignment Assignment) (uint64, error) {
	epoch, err := wire.ParseCounter(assignment.TurnEpoch)
	if err != nil || assignment.RunID == "" || assignment.TurnID == "" || assignment.GenerationID == "" || assignment.BrokerEndpoint == "" ||
		!lowerHex(assignment.ManifestDigest, 64) || !lowerHex(assignment.CreateAuthID, 32) || assignment.AdmissionRef.Validate() != nil {
		return 0, errors.New("worker runtime: invalid assignment")
	}
	if (assignment.SettlementManifest == nil) != (assignment.PredecessorTurnID == nil) {
		return 0, errors.New("worker runtime: incomplete settlement continuation")
	}
	if assignment.SettlementManifest != nil {
		if *assignment.PredecessorTurnID == "" || assignment.SettlementManifest.RunID != assignment.RunID || assignment.SettlementManifest.ProducedForTurnID != assignment.TurnID {
			return 0, errors.New("worker runtime: settlement identity mismatch")
		}
		if err := assignment.SettlementManifest.Validate(); err != nil {
			return 0, fmt.Errorf("worker runtime: settlement manifest: %w", err)
		}
	}
	return epoch, nil
}

func lowerHex(value string, length int) bool {
	if len(value) != length {
		return false
	}
	for index := range value {
		if (value[index] < '0' || value[index] > '9') && (value[index] < 'a' || value[index] > 'f') {
			return false
		}
	}
	return true
}

type argumentBytes struct{ value []byte }

func (source *argumentBytes) Snapshot() []byte { return append([]byte(nil), source.value...) }

type relayReadPeer interface {
	RelayRead(context.Context, string) ([]byte, error)
}

type relayReadConductor struct{ peer relayReadPeer }

func (conductor relayReadConductor) Relay(ctx context.Context, name string, arguments json.RawMessage) (json.RawMessage, error) {
	if name != "relay.read" {
		return nil, fmt.Errorf("worker runtime: conductor operation %q is unavailable", name)
	}
	var request struct {
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(arguments, &request); err != nil {
		return nil, err
	}
	result, err := conductor.peer.RelayRead(ctx, request.RelayID)
	return append(json.RawMessage(nil), result...), err
}

func (relayReadConductor) Describe(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return channel.DescriptionResponse{}, errors.New("worker runtime: conductor Describe is unavailable")
}

type composedToolInvoker struct {
	local executor.Invoker
	relay executor.Invoker
}

func newToolInvoker(backend tools.Backend, conductor relaytool.Conductor) executor.Invoker {
	if conductor == nil {
		if peer, ok := backend.(relayReadPeer); ok {
			conductor = relayReadConductor{peer: peer}
		}
	}
	return composedToolInvoker{local: tools.NewRegistry(backend), relay: relaytool.New(conductor, nil)}
}

func (invoker composedToolInvoker) Invoke(ctx context.Context, invocation executor.Invocation) (any, error) {
	switch invocation.Identity.CanonicalToolName {
	case "relay.submit", "relay.project", "relay.read":
		return invoker.relay.Invoke(ctx, invocation)
	default:
		return invoker.local.Invoke(ctx, invocation)
	}
}

type durableInvoker struct {
	delegate   executor.Invoker
	writer     *journal.Writer
	turnID     string
	roundIndex string
	call       ToolCall
}

func (invoker *durableInvoker) Invoke(ctx context.Context, invocation executor.Invocation) (any, error) {
	value, err := invoker.delegate.Invoke(ctx, invocation)
	if err != nil {
		return nil, err
	}
	persistedValue, truncated := truncateToolResult(value, turn.MaxCapturedOutput)
	toolCall := journal.Record{Kind: journal.KindToolCall, Fields: map[string]json.RawMessage{
		"tool_call_id": rawString(invoker.call.ID), "canonical_tool_name": rawString(invoker.call.CanonicalName),
		"canonical_args_digest": rawString(invocation.Identity.CanonicalArgsDigest), "args": append(json.RawMessage(nil), invocation.Arguments...),
	}}
	toolResult := journal.Record{Kind: journal.KindToolResult, Fields: map[string]json.RawMessage{
		"tool_call_id": rawString(invoker.call.ID), "content": mustJSON(persistedValue), "truncated": json.RawMessage(strconv.FormatBool(truncated)),
	}}
	commit, appendErr := invoker.writer.AppendRound(invoker.turnID, invoker.roundIndex, []journal.Record{toolCall, toolResult})
	if appendErr != nil {
		return nil, appendErr
	}
	if err := journal.RequireToolOutcome(commit, invoker.call.ID); err != nil {
		return nil, err
	}
	return value, nil
}

// MaxCapturedOutput is a local evidence bound, not an authority bound. The
// trusted control plane remains the only tool-call ceiling; this only makes a
// completed tool result honest about bytes omitted from the durable journal.
func truncateToolResult(value any, limit int) (any, bool) {
	switch typed := value.(type) {
	case string:
		if len(typed) > limit {
			return typed[:limit], true
		}
	case []byte:
		if len(typed) > limit {
			return append([]byte(nil), typed[:limit]...), true
		}
	}
	return value, false
}

func taskText(ctx context.Context, resolver ObjectiveResolver, reference turn.AdmissionRef) (string, error) {
	if reference.Kind == "operator_input" {
		return reference.TaskInput, nil
	}
	if resolver == nil {
		return "", errors.New("worker runtime: wake relay objective resolver is absent")
	}
	objective, err := resolver.ResolveObjective(ctx, reference.RelayID)
	if err != nil || objective == "" {
		return "", errors.Join(errors.New("worker runtime: wake relay objective unavailable"), err)
	}
	return objective, nil
}

func rawString(value string) json.RawMessage {
	raw, _ := json.Marshal(value)
	return raw
}

func mustJSON(value any) json.RawMessage {
	raw, _ := json.Marshal(value)
	return raw
}
