package engine

import (
	"context"
	"time"

	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
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

func (l *Loop) process(ctx context.Context, cmd intake.Cmd) Outcome {
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
