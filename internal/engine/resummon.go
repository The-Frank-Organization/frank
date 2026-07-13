package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

type resummonSubmitter interface {
	Submit(context.Context, intake.Cmd) (<-chan Outcome, string, error)
}

type ResummonScheduler struct {
	store     *store.Store
	submitter resummonSubmitter
	after     func(time.Duration) <-chan time.Time
}

func NewResummonScheduler(st *store.Store, submitter resummonSubmitter) (*ResummonScheduler, error) {
	if st == nil {
		return nil, fmt.Errorf("resummon store required")
	}
	if submitter == nil {
		return nil, fmt.Errorf("resummon submitter required")
	}
	return &ResummonScheduler{store: st, submitter: submitter, after: time.After}, nil
}

func (s *ResummonScheduler) EmitAfter(ctx context.Context, delay time.Duration, input ResummonInput) (ResummonResult, error) {
	if delay < 0 {
		return ResummonResult{}, fmt.Errorf("resummon delay must not be negative")
	}
	select {
	case <-ctx.Done():
		return ResummonResult{}, ctx.Err()
	case <-s.after(delay):
		return s.Emit(ctx, input)
	}
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
	existingIntake, err := s.outcomeForContentHash(contentHash)
	if err != nil {
		return ResummonResult{}, err
	}
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
		return ResummonResult{Record: rec, ContentHash: contentHash, Deduped: knownBefore}, nil
	}
}

func (s *ResummonScheduler) outcomeForContentHash(contentHash string) (string, error) {
	tab, err := tables.Build(s.store)
	if err != nil {
		return "", err
	}
	return tab.ContentHash[contentHash], nil
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
