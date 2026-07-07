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
	"github.com/jackli/frank/internal/tables"
)

type Job = intake.Job[Outcome]

type Outcome struct {
	State      string `json:"state"`
	RelayID    string `json:"relay_id,omitempty"`
	IntakeID   string `json:"intake_id,omitempty"`
	Reason     string `json:"reason,omitempty"`
	Detail     string `json:"detail,omitempty"`
	Credential string `json:"credential,omitempty"`
	Endpoint   string `json:"endpoint,omitempty"`
}

type Handler func(context.Context, intake.Cmd) (record.Record, []store.Intent, error)

type OutcomeExtras struct {
	Credential string
	Endpoint   string
}

type Loop struct {
	In                    chan Job
	Store                 *store.Store
	Handler               Handler
	Tables                *tables.T
	Timeout               time.Duration
	AfterCommit           func(*store.Store) error
	AfterAccepted         func(record.Record) (OutcomeExtras, error)
	CurrentAuthGeneration func(seatName string) string
	quarantine            chan string
}

func New(st *store.Store, handler Handler, ready *Ready) *Loop {
	if ready == nil {
		panic("engine.New requires Ready")
	}
	tab, _ := tables.Build(st)
	if tab == nil {
		tab = tables.New()
	}
	return &Loop{
		In:         make(chan Job, 32),
		Store:      st,
		Handler:    handler,
		Tables:     tab,
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
	if tab, err := tables.Build(l.Store); err == nil {
		l.Tables = tab
	}
	_ = l.completeTurn()
}

func (l *Loop) process(ctx context.Context, cmd intake.Cmd) (out Outcome) {
	defer func() {
		if recovered := recover(); recovered != nil {
			out = l.faultOutcome(cmd, fmt.Sprint(recovered))
		}
	}()
	if rec, ok := l.existingOutcomeForCommand(cmd); ok {
		return outcomeFromRecord(rec)
	}
	if out, superseded := l.supersededCredentialOutcome(cmd); superseded {
		return out
	}
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
	cmd.IntakeID = rec.Envelope.IntakeID
	if rec, ok := l.existingOutcomeForCommand(cmd); ok {
		return outcomeFromRecord(rec)
	}
	relayID, err := l.Store.Commit(rec, intents)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: rec.Envelope.IntakeID, Reason: safeReason("commit-error")}
	}
	rec.Envelope.RelayID = relayID
	if l.Tables != nil {
		l.Tables.OnCommit(rec)
	}
	if err := l.completeTurn(); err != nil {
		return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: rec.Envelope.IntakeID, Reason: safeReason("obligation-error")}
	}
	var extras OutcomeExtras
	if rec.Envelope.DeliveryState == record.Accepted && l.AfterAccepted != nil {
		extras, err = l.AfterAccepted(rec)
		if err != nil {
			return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: rec.Envelope.IntakeID, Reason: safeReason("derived-work-error")}
		}
	}
	out = Outcome{
		State:      rec.Envelope.DeliveryState,
		RelayID:    relayID,
		IntakeID:   rec.Envelope.IntakeID,
		Credential: extras.Credential,
		Endpoint:   extras.Endpoint,
	}
	if rec.Envelope.DeliveryState == record.Rejected {
		out.Detail = rec.Body
	}
	return out
}

func (l *Loop) supersededCredentialOutcome(cmd intake.Cmd) (Outcome, bool) {
	if cmd.AuthGeneration == "" || l.CurrentAuthGeneration == nil {
		return Outcome{}, false
	}
	current := l.CurrentAuthGeneration(cmd.Seat)
	if current == "" || current == cmd.AuthGeneration {
		return Outcome{}, false
	}
	rec := CredentialSupersededRecord(cmd)
	relayID, err := l.Store.Commit(rec, nil)
	if err != nil {
		return Outcome{State: record.Rejected, IntakeID: cmd.IntakeID, Reason: safeReason("commit-error")}, true
	}
	rec.Envelope.RelayID = relayID
	if l.Tables != nil {
		l.Tables.OnCommit(rec)
	}
	if err := l.completeTurn(); err != nil {
		return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: cmd.IntakeID, Reason: safeReason("obligation-error")}, true
	}
	return outcomeFromRecord(rec), true
}

func CredentialSupersededRecord(cmd intake.Cmd) record.Record {
	return record.Record{
		Envelope: record.Envelope{
			From:          cmd.Seat,
			Role:          commandRole(cmd),
			DeliveryState: record.Rejected,
			IntakeID:      cmd.IntakeID,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "candidate rejected after credential superseded"},
		Body:    bounce.Format(fieldspec.Violation{Field: "auth_generation", Class: "credential-superseded"}),
	}
}

func (l *Loop) existingOutcomeForCommand(cmd intake.Cmd) (record.Record, bool) {
	if rec, ok := l.existingOutcome(cmd.IntakeID); ok {
		return rec, true
	}
	if cmd.ContentHash == "" {
		return record.Record{}, false
	}
	tab := l.ensureTables()
	if tab == nil {
		return record.Record{}, false
	}
	intakeID := tab.ContentHash[cmd.ContentHash]
	if intakeID == "" {
		return record.Record{}, false
	}
	return l.existingOutcome(intakeID)
}

func (l *Loop) existingOutcome(intakeID string) (record.Record, bool) {
	if intakeID == "" {
		return record.Record{}, false
	}
	if l.Tables != nil {
		if rec, ok := l.Tables.OutcomeByIntake[intakeID]; ok {
			return rec, true
		}
	}
	tab := l.ensureTables()
	if tab == nil {
		return record.Record{}, false
	}
	rec, ok := tab.OutcomeByIntake[intakeID]
	return rec, ok
}

func (l *Loop) ensureTables() *tables.T {
	if l.Tables != nil {
		return l.Tables
	}
	tab, err := tables.Build(l.Store)
	if err != nil {
		return nil
	}
	l.Tables = tab
	return tab
}

func outcomeFromRecord(rec record.Record) Outcome {
	out := Outcome{
		State:    rec.Envelope.DeliveryState,
		RelayID:  rec.Envelope.RelayID,
		IntakeID: rec.Envelope.IntakeID,
	}
	switch rec.Envelope.DeliveryState {
	case record.Rejected:
		out.Detail = rec.Body
		if rec.Envelope.From == "system" && rec.Headers["SUBJECT"] == "candidate rejected after internal fault" {
			out.Reason = rec.Body
		}
	case record.Held:
		if rec.Envelope.From == "system" && rec.Headers["SUBJECT"] == "authority-bearing candidate held after internal fault" {
			out.Reason = safeReason("internal-fault")
		}
	}
	return out
}

func (l *Loop) completeTurn() error {
	if err := obligation.CompleteAuto(l.Store, l.Tables); err != nil {
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
		if l.Tables != nil {
			l.Tables.OnCommit(held)
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
	if l.Tables != nil {
		l.Tables.OnCommit(rejected)
	}
	detail := safeReason("internal-fault")
	return Outcome{State: record.Rejected, RelayID: relayID, IntakeID: cmd.IntakeID, Reason: detail, Detail: detail}
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
