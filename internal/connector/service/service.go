// Package service composes the connector pipeline behind CTRL-C and DATA-P.
package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/The-Frank-Organization/frank/internal/connector/attempt"
	"github.com/The-Frank-Organization/frank/internal/connector/authorize"
	"github.com/The-Frank-Organization/frank/internal/connector/control"
	"github.com/The-Frank-Organization/frank/internal/connector/credentials"
	"github.com/The-Frank-Organization/frank/internal/connector/frame"
	"github.com/The-Frank-Organization/frank/internal/connector/freeze"
	"github.com/The-Frank-Organization/frank/internal/connector/outcome"
	"github.com/The-Frank-Organization/frank/internal/connector/request"
	"github.com/The-Frank-Organization/frank/internal/connector/stream"
	"github.com/The-Frank-Organization/frank/internal/connector/translate"
	"github.com/The-Frank-Organization/frank/internal/connector/transport"
)

var ErrDataClosed = errors.New("connector: DATA-P closed")

type Provider interface {
	Send(context.Context, *freeze.Request, *credentials.WireRequest) (*transport.Response, error)
}

type invocationGatedProvider interface {
	SendGated(context.Context, *freeze.Request, *credentials.WireRequest, func() bool) (*transport.Response, error)
}

type Config struct {
	Control   *control.Session
	Data      io.ReadWriteCloser
	Provider  Provider
	BuildInfo string
}

type Service struct {
	control  *control.Session
	data     io.ReadWriteCloser
	sender   *frame.Sender
	decoder  *frame.Decoder
	provider Provider
	attempts *attempt.Manager
	build    string

	sequenceMu sync.Mutex
	sequence   frame.Counter
	eventMu    sync.Mutex
	activeMu   sync.Mutex
	activeID   string
	workers    sync.WaitGroup
	faults     chan error
}

func New(config Config) (*Service, error) {
	if config.Control == nil || config.Data == nil || config.BuildInfo == "" || !config.Control.Ready() {
		return nil, errors.New("connector: invalid service configuration")
	}
	if config.Provider == nil {
		config.Provider = transport.NewClient()
	}
	return &Service{
		control: config.Control, data: config.Data, provider: config.Provider, build: config.BuildInfo,
		sender:   frame.NewSender(config.Data, frame.SendQueueDepth, frame.SendDeadline),
		decoder:  frame.NewDecoder(map[string]frame.TypeSpec{"llm_request": {}, "cancel_attempt": {}}),
		attempts: attempt.New(), sequence: 1, faults: make(chan error, 1),
	}, nil
}

func (service *Service) Serve(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	results := make(chan error, 2)
	go func() { results <- service.controlLoop(ctx) }()
	go func() { results <- service.dataLoop(ctx) }()
	var err error
	select {
	case err = <-results:
	case err = <-service.faults:
	}
	cancel()
	service.attempts.AbortAllForLoss()
	_ = service.data.Close()
	service.workers.Wait()
	service.sender.Close()
	return err
}

func (service *Service) controlLoop(ctx context.Context) error {
	for {
		if err := service.control.HandleControl(ctx); err != nil {
			return err
		}
	}
}

func (service *Service) dataLoop(ctx context.Context) error {
	for {
		envelope, err := service.decoder.Read(service.data)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return ErrDataClosed
			}
			return err
		}
		if envelope.Channel != frame.ChannelDataProvider || envelope.RunID != service.control.RunID() || envelope.TurnEpoch == nil {
			return errors.New("connector: malformed DATA-P envelope")
		}
		switch envelope.Type {
		case "llm_request":
			if err := service.startRequest(ctx, envelope); err != nil {
				return err
			}
		case "cancel_attempt":
			if err := service.cancelRequest(ctx, envelope); err != nil && !errors.Is(err, attempt.ErrStaleEpoch) && !errors.Is(err, attempt.ErrEpochAhead) && !errors.Is(err, attempt.ErrUnknownAttempt) {
				return err
			}
		}
	}
}

func (service *Service) startRequest(ctx context.Context, envelope frame.Envelope) error {
	service.activeMu.Lock()
	if service.activeID != "" {
		service.activeMu.Unlock()
		return errors.New("connector: concurrent attempt refused")
	}
	metadata, err := requestMetadata(envelope.Body)
	if err != nil {
		service.activeMu.Unlock()
		return err
	}
	service.activeID = metadata.AttemptID
	service.activeMu.Unlock()
	service.workers.Add(1)
	go func() {
		defer service.workers.Done()
		defer func() {
			service.activeMu.Lock()
			service.activeID = ""
			service.activeMu.Unlock()
		}()
		if err := service.runRequest(ctx, envelope, metadata); err != nil {
			select {
			case service.faults <- err:
			default:
			}
		}
	}()
	return nil
}

type metadata struct {
	AttemptID string        `json:"attempt_id"`
	RunID     string        `json:"run_id"`
	TurnEpoch frame.Counter `json:"turn_epoch"`
}

func requestMetadata(raw []byte) (metadata, error) {
	var value metadata
	if json.Unmarshal(raw, &value) != nil || value.AttemptID == "" || value.RunID == "" {
		return metadata{}, errors.New("connector: malformed request identity")
	}
	return value, nil
}

func (service *Service) runRequest(ctx context.Context, envelope frame.Envelope, meta metadata) error {
	if meta.RunID != service.control.RunID() || meta.TurnEpoch != *envelope.TurnEpoch {
		return errors.New("connector: request identity mismatch")
	}
	fence, err := service.control.FenceDataEpoch(ctx, meta.TurnEpoch)
	if err != nil {
		return err
	}
	if fence != control.EpochAllowed {
		reply, _ := outcome.EpochDataReply(meta.AttemptID, meta.TurnEpoch, string(fence))
		return service.sendData(ctx, "data_reply", meta.TurnEpoch, &envelope.Seq, reply)
	}
	parsed, err := request.Parse(envelope.Body, service.control.Lane())
	if err != nil {
		reason := request.MalformedRequest
		var reject request.RejectReason
		if errors.As(err, &reject) {
			reason = reject
		}
		return service.emitLocalReject(ctx, envelope.Seq, meta, outcome.Outcome{Kind: outcome.RejectedLocal, RejectReason: reason, RefusalStage: outcome.PreFreeze})
	}
	translated, err := translate.Translate(parsed, service.control.Lane())
	if err != nil {
		return service.emitLocalReject(ctx, envelope.Seq, meta, outcome.Outcome{Kind: outcome.RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: outcome.PreFreeze})
	}
	frozen, err := freeze.Freeze(translated, service.control.Lane(), service.build)
	if err != nil {
		return service.emitLocalReject(ctx, envelope.Seq, meta, outcome.Outcome{Kind: outcome.RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: outcome.PreFreeze})
	}
	digests := &outcome.Digests{FrozenCore: frozen.CoreDigest(), ProviderLoweredTools: frozen.LoweredToolsDigest()}
	active, err := service.attempts.Begin(ctx, meta.AttemptID, meta.TurnEpoch, digests.FrozenCore, digests.ProviderLoweredTools)
	if err != nil {
		return err
	}
	defer service.attempts.Finish(meta.AttemptID)
	assignment := service.control.Assignment()
	verdict := authorize.Evaluate(authorize.Input{
		Policy: service.control.Policy(), PolicyDigest: assignment.PolicyDigest,
		ProviderLaneID: parsed.ProviderLaneID, PinnedLaneID: assignment.ProviderLaneID,
		FrozenCoreDigest: frozen.CoreDigest(), CredentialRef: assignment.CredentialRef,
		Lane: service.control.Lane(), Core: frozen.Core(),
	})
	if !verdict.Allowed {
		denied := outcome.Outcome{Kind: outcome.Denied, DenyReason: verdict.DenyReason, Digests: digests}
		return service.emitResultThenReply(ctx, envelope.Seq, meta, denied)
	}
	wire, err := credentials.Attach(frozen, verdict, service.control.Credentials())
	if err != nil {
		return service.emitLocalReject(ctx, envelope.Seq, meta, outcome.Outcome{Kind: outcome.RejectedLocal, RejectReason: request.InternalIntegrityFault, RefusalStage: outcome.PostFreeze, Digests: digests})
	}
	var response *transport.Response
	if gated, ok := service.provider.(invocationGatedProvider); ok {
		response, err = gated.SendGated(active.Context(), frozen, wire, active.TryMarkInvoked)
	} else {
		if !active.TryMarkInvoked() {
			return nil
		}
		response, err = service.provider.Send(active.Context(), frozen, wire)
	}
	if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
		return nil
	}
	if err != nil {
		if errors.Is(err, freeze.ErrFrozenMutation) {
			return service.emitLocalReject(ctx, envelope.Seq, meta, outcome.Outcome{
				Kind: outcome.RejectedLocal, RejectReason: request.InternalIntegrityFault,
				RefusalStage: outcome.PostFreeze, Digests: digests,
			})
		}
		class := "transport"
		if errors.Is(err, transport.ErrTTFBDeadline) {
			class = "timeout_ttfb"
		} else if errors.Is(err, transport.ErrRedirect) {
			class = "protocol"
		}
		event := terminalFailure(meta.AttemptID, digests, class, "provider request failed")
		if err := service.emitFailedStream(ctx, meta, event); err != nil {
			return err
		}
		if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
			return nil
		}
		result, _ := outcome.AttemptResult(meta.AttemptID, meta.TurnEpoch, outcome.Outcome{Kind: outcome.TransportFailed, Digests: digests})
		return service.control.SendAttemptResult(ctx, result)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		detail := strconv.Itoa(response.StatusCode)
		if response.RetryAfter != "" {
			detail += " retry-after=" + response.RetryAfter
		}
		event := terminalFailure(meta.AttemptID, digests, "provider_http", detail)
		if err := service.emitFailedStream(ctx, meta, event); err != nil {
			return err
		}
		if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
			return nil
		}
		result, _ := outcome.AttemptResult(meta.AttemptID, meta.TurnEpoch, outcome.Outcome{Kind: outcome.TransportFailed, Digests: digests})
		return service.control.SendAttemptResult(ctx, result)
	}
	terminalKind := stream.Failed
	err = stream.ParseEach(response.Body, stream.Meta{
		AttemptID: meta.AttemptID, ProviderLaneID: parsed.ProviderLaneID, TurnID: parsed.TurnID,
		FrozenCoreDigest: digests.FrozenCore, ProviderLoweredToolsDigest: digests.ProviderLoweredTools,
	}, func(event stream.Event) error {
		service.eventMu.Lock()
		defer service.eventMu.Unlock()
		if event.Terminal() {
			// Finish is the terminal-vs-cancel arbitration point. Whichever
			// removes the active attempt first owns the single terminal.
			service.attempts.Finish(meta.AttemptID)
		}
		if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
			return context.Canceled
		}
		switch event.Kind {
		case stream.TextStart, stream.TextDelta:
			if err := active.SetPartial("text"); err != nil {
				return err
			}
		case stream.ToolCallStart, stream.ToolCallDelta:
			if err := active.SetPartial("tool_call_incomplete"); err != nil {
				return err
			}
		case stream.ToolCallEnd:
			if err := active.SetPartial("none"); err != nil {
				return err
			}
		}
		if event.Terminal() {
			terminalKind = event.Kind
		}
		if err := service.sendData(ctx, "provider_event", meta.TurnEpoch, nil, event); err != nil {
			return err
		}
		return nil
	})
	if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
		return nil
	}
	if err != nil {
		return err
	}
	kind := outcome.TransportFailed
	if terminalKind == stream.Completed {
		kind = outcome.SentCompleted
	}
	result, _ := outcome.AttemptResult(meta.AttemptID, meta.TurnEpoch, outcome.Outcome{Kind: kind, Digests: digests})
	return service.control.SendAttemptResult(ctx, result)
}

func (service *Service) cancelRequest(ctx context.Context, envelope frame.Envelope) error {
	intent, err := attempt.ParseIntent(envelope.Body)
	if err != nil || intent.TurnEpoch != *envelope.TurnEpoch {
		return attempt.ErrInvalidIntent
	}
	cancellation, err := service.attempts.Cancel(intent, service.control.Epoch())
	if err != nil {
		return err
	}
	if cancellation.Duplicate {
		return nil
	}
	// Cancel above aborts the HTTP context immediately. The output lock below
	// only orders its terminal behind any event whose write was already live.
	service.eventMu.Lock()
	defer service.eventMu.Unlock()
	result, err := outcome.AttemptResult(intent.AttemptID, intent.TurnEpoch, cancellation.Outcome)
	if err != nil {
		return err
	}
	if err := service.control.SendAttemptResult(ctx, result); err != nil {
		return err
	}
	return service.sendData(ctx, "provider_event", intent.TurnEpoch, nil, cancellation.Event)
}

type localRejectEmitter struct {
	service *Service
	ctx     context.Context
	re      frame.Counter
	epoch   frame.Counter
}

func (emitter localRejectEmitter) SendAttemptResult(result outcome.AttemptResultV2) error {
	return emitter.service.control.SendAttemptResult(emitter.ctx, result)
}

func (emitter localRejectEmitter) SendDataReply(reply outcome.DataReplyV2) error {
	return emitter.service.sendData(emitter.ctx, "data_reply", emitter.epoch, &emitter.re, reply)
}

func (service *Service) emitLocalReject(ctx context.Context, re frame.Counter, meta metadata, value outcome.Outcome) error {
	service.eventMu.Lock()
	defer service.eventMu.Unlock()
	service.attempts.Finish(meta.AttemptID)
	if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
		return nil
	}
	emitter := localRejectEmitter{service: service, ctx: ctx, re: re, epoch: meta.TurnEpoch}
	if err := outcome.EmitLocalReject(emitter, meta.AttemptID, meta.TurnEpoch, value); err != nil {
		return err
	}
	return nil
}

func (service *Service) emitResultThenReply(ctx context.Context, re frame.Counter, meta metadata, value outcome.Outcome) error {
	service.eventMu.Lock()
	defer service.eventMu.Unlock()
	service.attempts.Finish(meta.AttemptID)
	if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
		return nil
	}
	result, err := outcome.AttemptResult(meta.AttemptID, meta.TurnEpoch, value)
	if err != nil {
		return err
	}
	reply, err := outcome.DataReply(meta.AttemptID, meta.TurnEpoch, value)
	if err != nil {
		return err
	}
	if err := service.control.SendAttemptResult(ctx, result); err != nil {
		return err
	}
	if err := service.sendData(ctx, "data_reply", meta.TurnEpoch, &re, reply); err != nil {
		return err
	}
	return nil
}

func (service *Service) sendData(ctx context.Context, kind string, epoch frame.Counter, re *frame.Counter, body any) error {
	raw, err := json.Marshal(body)
	if err != nil {
		return err
	}
	service.sequenceMu.Lock()
	sequence := service.sequence
	service.sequence++
	service.sequenceMu.Unlock()
	return service.sender.SendContext(ctx, frame.Envelope{
		Version: 1, Channel: frame.ChannelDataProvider, Type: kind, Seq: sequence, Re: re,
		RunID: service.control.RunID(), TurnEpoch: &epoch, Body: raw,
	})
}

func terminalFailure(attemptID string, digests *outcome.Digests, class, detail string) stream.Event {
	return stream.Event{
		Schema: stream.SchemaV2, Kind: stream.Failed, AttemptID: attemptID,
		ErrorClass: class, Detail: detail, FrozenCoreDigest: digests.FrozenCore,
		ProviderLoweredToolsDigest: digests.ProviderLoweredTools,
	}
}

func (service *Service) emitFailedStream(ctx context.Context, meta metadata, terminal stream.Event) error {
	service.eventMu.Lock()
	defer service.eventMu.Unlock()
	service.attempts.Finish(meta.AttemptID)
	if _, cancelled := service.attempts.Cancellation(meta.AttemptID); cancelled {
		return nil
	}
	started := stream.Event{Schema: stream.SchemaV2, Kind: stream.AttemptStarted, AttemptID: meta.AttemptID}
	if err := service.sendData(ctx, "provider_event", meta.TurnEpoch, nil, started); err != nil {
		return err
	}
	if err := service.sendData(ctx, "provider_event", meta.TurnEpoch, nil, terminal); err != nil {
		return err
	}
	return nil
}

func (service *Service) String() string {
	return fmt.Sprintf("connector-service{%s}", service.control.RunID())
}
