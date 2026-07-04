package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

type Job = intake.Job[Outcome]

type Outcome struct {
	State    string `json:"state"`
	RelayID  string `json:"relay_id,omitempty"`
	IntakeID string `json:"intake_id,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

type Handler func(context.Context, intake.Cmd) (record.Record, []store.Intent, error)

type Loop struct {
	In          chan Job
	Store       *store.Store
	Handler     Handler
	Timeout     time.Duration
	AfterCommit func(*store.Store) error
	quarantine  chan string
}

func New(st *store.Store, handler Handler, ready *Ready) *Loop {
	if ready == nil {
		panic("engine.New requires Ready")
	}
	return &Loop{
		In:         make(chan Job, 32),
		Store:      st,
		Handler:    handler,
		Timeout:    5 * time.Second,
		quarantine: make(chan string, 32),
	}
}

func (l *Loop) EnqueueQuarantine(relayID string) {
	if relayID == "" {
		return
	}
	select {
	case l.quarantine <- relayID:
	default:
	}
}

func (l *Loop) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case relayID := <-l.quarantine:
			l.processQuarantine(relayID)
			l.drainQuarantine()
		case job := <-l.In:
			l.drainQuarantine()
			out := l.process(ctx, job.Cmd)
			l.drainQuarantine()
			timeout := l.Timeout
			if timeout <= 0 {
				timeout = 5 * time.Second
			}
			crashpoint.Hit("pre_outcome_reply")
			select {
			case job.ReplyCh <- out:
			case <-ctx.Done():
				return
			case <-time.After(timeout):
			}
		}
	}
}

func (l *Loop) drainQuarantine() {
	for {
		select {
		case relayID := <-l.quarantine:
			l.processQuarantine(relayID)
		default:
			return
		}
	}
}

func (l *Loop) processQuarantine(relayID string) {
	_, _ = l.Store.QuarantineOne(relayID)
	_ = l.completeTurn()
}

func (l *Loop) process(ctx context.Context, cmd intake.Cmd) (out Outcome) {
	defer func() {
		if recovered := recover(); recovered != nil {
			out = l.faultOutcome(cmd, fmt.Sprint(recovered))
		}
	}()
	if l.Handler == nil {
		return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: "no handler"}
	}
	rec, intents, err := l.Handler(ctx, cmd)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: safeReason("internal-error")}
	}
	if rec.Envelope.IntakeID == "" {
		rec.Envelope.IntakeID = cmd.IntakeID
	}
	relayID, err := l.Store.Commit(rec, intents)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: rec.Envelope.IntakeID, Reason: safeReason("commit-error")}
	}
	if err := l.completeTurn(); err != nil {
		return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: rec.Envelope.IntakeID, Reason: safeReason("obligation-error")}
	}
	return Outcome{
		State:    rec.Envelope.DeliveryState,
		RelayID:  relayID,
		IntakeID: rec.Envelope.IntakeID,
	}
}

func (l *Loop) completeTurn() error {
	if err := obligation.CompleteAuto(l.Store); err != nil {
		return err
	}
	if l.AfterCommit != nil {
		return l.AfterCommit(l.Store)
	}
	return nil
}

func (l *Loop) faultOutcome(cmd intake.Cmd, reason string) Outcome {
	var cand record.Record
	_ = json.Unmarshal(cmd.Payload, &cand)
	meta := seat.SeatMeta{Name: cmd.Seat, Role: commandRole(cmd), IsOperator: cmd.IsOperator}
	if lineage.AuthorityBearing(cand, meta) {
		held := record.Record{
			Envelope: record.Envelope{
				RelayID:       "held-" + cmd.IntakeID,
				From:          "system",
				Role:          "system",
				DeliveryState: record.Held,
				IntakeID:      cmd.IntakeID,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "authority-bearing candidate held after internal fault"},
			Body:    string(cmd.Payload),
		}
		relayID, err := l.Store.Commit(held, nil)
		if err != nil {
			return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: safeReason("commit-error")}
		}
		return Outcome{State: record.Held, RelayID: relayID, IntakeID: cmd.IntakeID, Reason: safeReason("internal-fault")}
	}
	rejected := record.Record{
		Envelope: record.Envelope{
			RelayID:       "rejected-" + cmd.IntakeID,
			From:          "system",
			Role:          "system",
			DeliveryState: record.Rejected,
			IntakeID:      cmd.IntakeID,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "candidate rejected after internal fault"},
		Body:    safeReason("internal-fault"),
	}
	relayID, err := l.Store.Commit(rejected, nil)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: safeReason("commit-error")}
	}
	return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: cmd.IntakeID, Reason: safeReason("internal-fault")}
}

func safeReason(class string) string {
	return bounce.Format(fieldspec.Violation{Field: "system", Class: class})
}

func commandRole(cmd intake.Cmd) string {
	if cmd.Role != "" {
		return cmd.Role
	}
	if i := strings.LastIndex(cmd.Seat, "."); i >= 0 && i+1 < len(cmd.Seat) {
		return cmd.Seat[i+1:]
	}
	return ""
}
