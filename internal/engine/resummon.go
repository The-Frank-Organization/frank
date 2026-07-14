package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

const (
	ResummonNoResponse      = "no_response"
	ResummonAnsweredStalled = "answered_but_stalled"
	SummonLocal             = "local"
	SummonLouderLocal       = "louder_local"
)

type ResummonInput struct {
	Seat          string `json:"seat"`
	DecisionID    string `json:"decision_id"`
	CadenceSlot   string `json:"cadence_slot"`
	Reason        string `json:"reason"`
	SummonChannel string `json:"summon_channel"`
}

type ResummonResult struct {
	Record      record.Record
	ContentHash string
	Deduped     bool
}

type ResummonCadence struct {
	NoResponse      time.Duration
	AnsweredStalled time.Duration
}

var DefaultResummonCadence = ResummonCadence{
	NoResponse:      time.Hour,
	AnsweredStalled: time.Hour,
}

type resummonSubmitter interface {
	Submit(context.Context, intake.Cmd) (<-chan Outcome, string, error)
}

type ResummonScheduler struct {
	store     *store.Store
	submitter resummonSubmitter
	snapshot  func() *tables.T
	after     func(time.Duration) <-chan time.Time
	mu        sync.Mutex
	armed     map[string]bool
	seen      map[string]string
}

func NewResummonScheduler(st *store.Store, submitter resummonSubmitter, snapshot func() *tables.T) (*ResummonScheduler, error) {
	if st == nil {
		return nil, fmt.Errorf("resummon store required")
	}
	if submitter == nil {
		return nil, fmt.Errorf("resummon submitter required")
	}
	if snapshot == nil {
		return nil, fmt.Errorf("resummon table snapshot required")
	}
	return &ResummonScheduler{
		store: st, submitter: submitter, snapshot: snapshot,
		armed: map[string]bool{}, seen: map[string]string{},
	}, nil
}

func (s *ResummonScheduler) EmitAfter(ctx context.Context, delay time.Duration, input ResummonInput) (ResummonResult, error) {
	if delay < 0 {
		return ResummonResult{}, fmt.Errorf("resummon delay must not be negative")
	}
	if s.after != nil {
		select {
		case <-ctx.Done():
			return ResummonResult{}, ctx.Err()
		case <-s.after(delay):
			return s.Emit(ctx, input)
		}
	}
	timer := time.NewTimer(delay)
	defer func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	}()
	select {
	case <-ctx.Done():
		return ResummonResult{}, ctx.Err()
	case <-timer.C:
		return s.Emit(ctx, input)
	}
}

// ArmParked installs the two G4 timer classes for every durably parked gate.
// Each (gate, cadence-slot) is armed once per process; crash refires remain
// safe because Emit also deduplicates from the durable content hash.
func (s *ResummonScheduler) ArmParked(ctx context.Context, cadence ResummonCadence) error {
	if cadence.NoResponse < 0 || cadence.AnsweredStalled < 0 {
		return fmt.Errorf("resummon cadence must not be negative")
	}
	tab, err := tables.Build(s.store)
	if err != nil {
		return err
	}
	for gateID := range tab.ParkedLanes {
		if GateState(tab, gateID) == GateResumed || gateUnpreservable(tab, gateID) {
			continue
		}
		gate, ok := tab.ByRelay[gateID]
		if !ok {
			continue
		}
		s.arm(ctx, cadence.NoResponse, ResummonInput{
			Seat: gate.Envelope.From, DecisionID: gateID, CadenceSlot: "g4-no-response-1",
			Reason: ResummonNoResponse, SummonChannel: SummonLocal,
		})
		s.arm(ctx, cadence.AnsweredStalled, ResummonInput{
			Seat: gate.Envelope.From, DecisionID: gateID, CadenceSlot: "g4-answered-stalled-1",
			Reason: ResummonAnsweredStalled, SummonChannel: SummonLouderLocal,
		})
	}
	return nil
}

func (s *ResummonScheduler) arm(ctx context.Context, delay time.Duration, input ResummonInput) {
	key := ResummonContentHash(input)
	s.mu.Lock()
	if s.armed[key] {
		s.mu.Unlock()
		return
	}
	s.armed[key] = true
	s.mu.Unlock()
	go func() {
		if _, err := s.waitAndEmitIfDue(ctx, delay, input); err != nil {
			return
		}
	}()
}

func (s *ResummonScheduler) waitAndEmitIfDue(ctx context.Context, delay time.Duration, input ResummonInput) (ResummonResult, error) {
	if delay < 0 {
		return ResummonResult{}, fmt.Errorf("resummon delay must not be negative")
	}
	timer := time.NewTimer(delay)
	defer func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	}()
	select {
	case <-ctx.Done():
		return ResummonResult{}, ctx.Err()
	case <-timer.C:
	}
	due, err := s.due(input)
	if err != nil || !due {
		return ResummonResult{}, err
	}
	return s.Emit(ctx, input)
}

func (s *ResummonScheduler) due(input ResummonInput) (bool, error) {
	tab, err := tables.Build(s.store)
	if err != nil {
		return false, err
	}
	if GateState(tab, input.DecisionID) == GateResumed {
		return false, nil
	}
	if gateUnpreservable(tab, input.DecisionID) {
		return false, nil
	}
	answered := false
	for _, rec := range tab.Records {
		if rec.Headers["resolves_gate"] == input.DecisionID {
			answered = true
			break
		}
	}
	if input.Reason == ResummonAnsweredStalled {
		return answered, nil
	}
	return !answered, nil
}

func gateUnpreservable(tab *tables.T, gateID string) bool {
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState == record.Held && rec.Headers["failing_edge"] == "stale_schema" && rec.Headers["subject_ref"] == gateID {
			return true
		}
	}
	return false
}

func (s *ResummonScheduler) Emit(ctx context.Context, input ResummonInput) (ResummonResult, error) {
	if err := validateResummonInput(input); err != nil {
		return ResummonResult{}, err
	}
	body, err := json.Marshal(input)
	if err != nil {
		return ResummonResult{}, err
	}
	contentHash := ResummonContentHash(input)
	existingIntake := s.outcomeForContentHash(contentHash)
	knownBefore := existingIntake != ""
	reply, _, err := s.submitter.Submit(ctx, intake.Cmd{
		Seat: "system", Role: "system", Verb: "emit-resummon", Payload: body, ContentHash: contentHash,
	})
	if err != nil {
		return ResummonResult{}, err
	}
	select {
	case <-ctx.Done():
		return ResummonResult{}, ctx.Err()
	case outcome := <-reply:
		if outcome.State != record.Accepted || outcome.RelayID == "" {
			return ResummonResult{}, fmt.Errorf("resummon emit rejected: %s", outcome.Reason)
		}
		rec, err := s.store.Read(outcome.RelayID)
		if err != nil {
			return ResummonResult{}, err
		}
		s.mu.Lock()
		s.seen[contentHash] = rec.Envelope.IntakeID
		s.mu.Unlock()
		return ResummonResult{Record: rec, ContentHash: contentHash, Deduped: knownBefore}, nil
	}
}

func (s *ResummonScheduler) outcomeForContentHash(contentHash string) string {
	s.mu.Lock()
	known := s.seen[contentHash]
	s.mu.Unlock()
	if known != "" {
		return known
	}
	return s.snapshot().ContentHash[contentHash]
}

// ResummonHandler is an internal conductor arm. It accepts only the system
// command emitted by ResummonScheduler and returns the emit-only record to the
// serialized loop; seat submit validation is never bypassed or widened.
func ResummonHandler(next Handler) Handler {
	return func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		if cmd.Verb != "emit-resummon" {
			if next == nil {
				return record.Record{}, nil, fmt.Errorf("unsupported internal command %q", cmd.Verb)
			}
			return next(ctx, cmd)
		}
		if cmd.Seat != "system" || cmd.Role != "system" {
			return record.Record{}, nil, fmt.Errorf("resummon command requires system emitter")
		}
		var input ResummonInput
		if err := json.Unmarshal(cmd.Payload, &input); err != nil {
			return record.Record{}, nil, fmt.Errorf("decode resummon: %w", err)
		}
		if err := validateResummonInput(input); err != nil {
			return record.Record{}, nil, err
		}
		if cmd.ContentHash != ResummonContentHash(input) {
			return record.Record{}, nil, fmt.Errorf("resummon content hash mismatch")
		}
		to, err := fieldspec.EncodeAddressList([]string{"operator"})
		if err != nil {
			return record.Record{}, nil, err
		}
		return record.Record{
			Envelope: record.Envelope{
				From: "system", To: "operator", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1,
			},
			Headers: map[string]string{
				"PHASE":       "SITREP",
				"SUBJECT":     GateResummonDue,
				"TO":          to,
				"record_kind": "resummon_command",
				"subject_ref": input.DecisionID,
			},
			Body: string(cmd.Payload),
		}, nil, nil
	}
}

func ResummonContentHash(input ResummonInput) string {
	key, _ := json.Marshal(struct {
		Seat        string `json:"seat"`
		DecisionID  string `json:"decision_id"`
		CadenceSlot string `json:"cadence_slot"`
	}{Seat: input.Seat, DecisionID: input.DecisionID, CadenceSlot: input.CadenceSlot})
	sum := sha256.Sum256(key)
	return "resummon:" + hex.EncodeToString(sum[:])
}

func validateResummonInput(input ResummonInput) error {
	if input.Seat == "" || input.DecisionID == "" || input.CadenceSlot == "" {
		return fmt.Errorf("resummon seat, decision, and cadence slot required")
	}
	if input.Reason != ResummonNoResponse && input.Reason != ResummonAnsweredStalled {
		return fmt.Errorf("resummon reason invalid")
	}
	if input.SummonChannel != SummonLocal && input.SummonChannel != SummonLouderLocal {
		return fmt.Errorf("resummon channel invalid")
	}
	return nil
}
