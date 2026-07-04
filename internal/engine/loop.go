package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

type Job struct {
	Cmd     intake.Cmd
	ReplyCh chan Outcome
}

type Outcome struct {
	State    string
	RelayID  string
	IntakeID string
	Reason   string
}

type Handler func(context.Context, intake.Cmd) (record.Record, []store.Intent, error)

type Loop struct {
	In      chan Job
	Store   *store.Store
	Handler Handler
	Timeout time.Duration
}

func New(st *store.Store, handler Handler) *Loop {
	return &Loop{
		In:      make(chan Job, 32),
		Store:   st,
		Handler: handler,
		Timeout: 5 * time.Second,
	}
}

func (l *Loop) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-l.In:
			job.ReplyCh <- l.process(ctx, job.Cmd)
		}
	}
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
		return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: err.Error()}
	}
	if rec.Envelope.IntakeID == "" {
		rec.Envelope.IntakeID = cmd.IntakeID
	}
	relayID, err := l.Store.Commit(rec, intents)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: rec.Envelope.IntakeID, Reason: err.Error()}
	}
	return Outcome{
		State:    rec.Envelope.DeliveryState,
		RelayID:  relayID,
		IntakeID: rec.Envelope.IntakeID,
	}
}

func (l *Loop) faultOutcome(cmd intake.Cmd, reason string) Outcome {
	var cand record.Record
	_ = json.Unmarshal(cmd.Payload, &cand)
	meta := seat.SeatMeta{Name: cmd.Seat, Role: cand.Envelope.Role}
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
			return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: err.Error()}
		}
		return Outcome{State: record.Held, RelayID: relayID, IntakeID: cmd.IntakeID, Reason: reason}
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
		Body:    reason,
	}
	relayID, err := l.Store.Commit(rejected, nil)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: err.Error()}
	}
	return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: cmd.IntakeID, Reason: reason}
}
