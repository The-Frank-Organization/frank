package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/derived"
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
	State           string `json:"state,omitempty"`
	DecisionState   string `json:"decision_state"`
	PostCommitState string `json:"post_commit_state"`
	RelayID         string `json:"relay_id,omitempty"`
	IntakeID        string `json:"intake_id,omitempty"`
	Reason          string `json:"reason,omitempty"`
	Detail          string `json:"detail,omitempty"`
	Credential      string `json:"credential,omitempty"`
	Endpoint        string `json:"endpoint,omitempty"`
}

type Handler func(context.Context, intake.Cmd) (record.Record, []store.Intent, error)

type OutcomeExtras struct {
	Credential string
	Endpoint   string
}

const DerivedRetryCeiling = 8

type Loop struct {
	In                      chan Job
	Store                   *store.Store
	Handler                 Handler
	Tables                  *tables.T
	Timeout                 time.Duration
	AfterCommit             func(*store.Store) error
	AfterAccepted           func(record.Record) (OutcomeExtras, error)
	DerivedCommit           func(record.Record) (string, error)
	MintRealized            func(record.Record) bool
	AfterGateResolution     func(record.Record) error
	AfterApprovalResolution func(record.Record) error
	ServiceWhileBlocked     bool
	CurrentAuthGeneration   func(seatName string) string
	ClassGCompleteAuto      func(*store.Store, *tables.T) error
	ClassGGC                func(*store.Store, *tables.T) error
	ClassGBuildTables       func(*store.Store) (*tables.T, error)
	ClassGPublishTables     func(*tables.T)
	ClassGArmScheduler      func() error
	quarantine              chan string
	classGMu                sync.RWMutex
	classGDirty             bool
	derivedAttempts         map[string]int
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
		In:              make(chan Job, 32),
		Store:           st,
		Handler:         handler,
		Tables:          tab,
		Timeout:         5 * time.Second,
		quarantine:      make(chan string, 32),
		derivedAttempts: map[string]int{},
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
	if _, err := l.Store.QuarantineOne(relayID); err != nil {
		l.setClassGDirty(true)
		return
	}
	if tab, err := tables.Build(l.Store); err == nil {
		l.Tables = tab
	} else {
		l.setClassGDirty(true)
		return
	}
	_ = l.completeTurn()
}

func (l *Loop) process(ctx context.Context, cmd intake.Cmd) (out Outcome) {
	defer func() {
		if recovered := recover(); recovered != nil {
			out = l.faultOutcome(cmd, fmt.Sprint(recovered))
		}
		out = l.classGOutcome(out)
	}()
	if l.retryClassGIfDirty() {
		l.retryAllBlind()
	}
	if rec, ok := l.existingOutcomeForCommand(cmd); ok {
		credential, endpoint, err := l.resumeCallerPresent(rec)
		out = l.outcomeFromRecord(rec)
		if err != nil {
			out.Reason = safeReason("derived-work-error")
			return out
		}
		out.Credential = credential
		out.Endpoint = endpoint
		return out
	}
	if out, superseded := l.supersededCredentialOutcome(cmd); superseded {
		return out
	}
	if l.Handler == nil {
		return outcomeFor(record.Rejected, "complete", "", cmd.IntakeID, "no handler")
	}
	rec, intents, err := l.callHandler(ctx, cmd)
	if err != nil {
		return outcomeFor(record.Rejected, "complete", "", cmd.IntakeID, safeReason("internal-error"))
	}
	if rec.Envelope.IntakeID == "" {
		rec.Envelope.IntakeID = cmd.IntakeID
	}
	cmd.IntakeID = rec.Envelope.IntakeID
	if rec, ok := l.existingOutcomeForCommand(cmd); ok {
		return l.outcomeFromRecord(rec)
	}
	rec, intents, err = revalidateAtCommit(l.Store, rec, intents)
	if err != nil {
		return outcomeFor(record.Rejected, "complete", "", rec.Envelope.IntakeID, safeReason("commit-validation-error"))
	}
	derived.Stamp(&rec)
	relayID, err := l.Store.Commit(rec, intents)
	if err != nil {
		return outcomeFor(record.Rejected, "complete", "", rec.Envelope.IntakeID, safeReason("commit-error"))
	}
	rec.Envelope.RelayID = relayID
	if l.Tables != nil {
		l.Tables.OnCommit(rec)
	}
	if err := l.completeTurn(); err != nil {
		return outcomeFor(rec.Envelope.DeliveryState, "pending", relayID, rec.Envelope.IntakeID, safeReason("obligation-error"))
	}
	if err = l.runBlindForRecord(rec); err != nil {
		out = l.outcomeFromRecord(rec)
		out.Reason = safeReason("derived-work-error")
		return out
	}
	var extras OutcomeExtras
	if rec.Envelope.DeliveryState == record.Accepted {
		if containsHook(derived.Cursor(rec), "mint") {
			extras.Credential, extras.Endpoint, err = l.resumeCallerPresent(rec)
		} else if l.AfterAccepted != nil {
			extras, err = l.AfterAccepted(rec)
		}
		if err != nil {
			out = l.outcomeFromRecord(rec)
			if out.PostCommitState == "complete" {
				out = outcomeFor(rec.Envelope.DeliveryState, "pending", rec.Envelope.RelayID, rec.Envelope.IntakeID, "")
			}
			out.Reason = safeReason("derived-work-error")
			return out
		}
	}
	out = l.outcomeFromRecord(rec)
	out.Credential = extras.Credential
	out.Endpoint = extras.Endpoint
	if rec.Envelope.DeliveryState == record.Rejected {
		out.Detail = rec.Body
	}
	return out
}

type handlerResult struct {
	record  record.Record
	intents []store.Intent
	err     error
}

// callHandler keeps the commit loop as the sole serialized writer while a
// parked handler awaits an operator disposition. Only handler computation is
// detached; every nested command still returns here for ordered commit.
func (l *Loop) callHandler(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
	if !l.ServiceWhileBlocked {
		return l.Handler(ctx, cmd)
	}
	done := make(chan handlerResult, 1)
	go func() {
		rec, intents, err := l.Handler(ctx, cmd)
		done <- handlerResult{record: rec, intents: intents, err: err}
	}()
	for {
		select {
		case result := <-done:
			return result.record, result.intents, result.err
		case relayID := <-l.quarantine:
			l.processQuarantine(relayID)
		case nested := <-l.In:
			out := l.process(ctx, nested.Cmd)
			timer := time.NewTimer(l.replyTimeout())
			select {
			case nested.ReplyCh <- out:
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
			case <-ctx.Done():
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				return record.Record{}, nil, ctx.Err()
			case <-timer.C:
			}
		case <-ctx.Done():
			return record.Record{}, nil, ctx.Err()
		}
	}
}

func (l *Loop) replyTimeout() time.Duration {
	if l.Timeout > 0 {
		return l.Timeout
	}
	return 5 * time.Second
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
	if existing, ok := l.existingOutcome(rec.Envelope.IntakeID); ok {
		return l.outcomeFromRecord(existing), true
	}
	relayID, err := l.Store.Commit(rec, nil)
	if err != nil {
		return outcomeFor(record.Rejected, "complete", "", cmd.IntakeID, safeReason("commit-error")), true
	}
	rec.Envelope.RelayID = relayID
	if l.Tables != nil {
		l.Tables.OnCommit(rec)
	}
	if err := l.completeTurn(); err != nil {
		out := outcomeFor(rec.Envelope.DeliveryState, "pending", relayID, cmd.IntakeID, safeReason("obligation-error"))
		out.Detail = rec.Body
		return out, true
	}
	return l.outcomeFromRecord(rec), true
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
	out := outcomeFor(rec.Envelope.DeliveryState, "complete", rec.Envelope.RelayID, rec.Envelope.IntakeID, "")
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

func (l *Loop) outcomeFromRecord(rec record.Record) Outcome {
	out := outcomeFromRecord(rec)
	tab := l.ensureTables()
	if tab == nil {
		return out
	}
	status, ok := tab.DerivedWork[rec.Envelope.RelayID]
	if !ok || status.Status == "" {
		return out
	}
	if status.Status == "unknown" && containsHook(status.Cursor, "mint") && derived.ResumableAttempt(tab.Records, rec.Envelope.RelayID, "mint") {
		if l.MintRealized != nil && l.MintRealized(rec) {
			status.Status = "failed"
		} else {
			status.Status = "pending"
		}
	}
	out.State = ""
	out.PostCommitState = status.Status
	return out
}

// resumeCallerPresent is the only delivery path for the non-idempotent mint
// hook. A durable attempt marker precedes the effect; a replay with realization
// evidence never invokes the effect again.
func (l *Loop) resumeCallerPresent(rec record.Record) (string, string, error) {
	if rec.Envelope.DeliveryState != record.Accepted || !containsHook(derived.Cursor(rec), "mint") || l.AfterAccepted == nil {
		return "", "", nil
	}
	if l.ClassGDirty() {
		return "", "", nil
	}
	tab := l.ensureTables()
	if tab == nil {
		return "", "", errors.New("derived tables unavailable")
	}
	status, ok := tab.DerivedWork[rec.Envelope.RelayID]
	if !ok || !containsHook(status.Cursor, "mint") || status.Status == "failed" || status.Status == "" {
		return "", "", nil
	}
	openAttempt, predecessor, validAttempts := derived.AttemptState(tab.Records, rec.Envelope.RelayID, "mint")
	if !validAttempts {
		return "", "", nil
	}
	if openAttempt && l.MintRealized != nil && l.MintRealized(rec) {
		if err := l.commitDerived(derived.RealizedUndeliveredRecord(rec.Envelope.RelayID)); err != nil {
			return "", "", err
		}
		return "", "", nil
	}
	if !openAttempt {
		if err := l.commitDerived(derived.AttemptRecord(rec.Envelope.RelayID, "mint", predecessor)); err != nil {
			return "", "", err
		}
	}
	extras, err := l.AfterAccepted(rec)
	if err != nil {
		return "", "", err
	}
	if err := l.commitCursorAdvance(rec.Envelope.RelayID, "mint"); err != nil {
		if l.MintRealized != nil && l.MintRealized(rec) {
			_ = l.commitDerived(derived.RealizedUndeliveredRecord(rec.Envelope.RelayID))
		}
		return "", "", err
	}
	return extras.Credential, extras.Endpoint, nil
}

func (l *Loop) retryAllBlind() {
	for _, relayID := range derived.OpenRelayIDs(l.Tables.DerivedWork) {
		rec, ok := l.Tables.ByRelay[relayID]
		if !ok {
			continue
		}
		_ = l.runBlindForRecord(rec)
	}
}

func (l *Loop) runBlindForRecord(rec record.Record) error {
	status, ok := l.Tables.DerivedWork[rec.Envelope.RelayID]
	if !ok || status.Status != "pending" {
		return nil
	}
	for _, hook := range []string{"gate", "approval"} {
		if !containsHook(status.Cursor, hook) {
			continue
		}
		epoch := derived.RetryEpoch(l.Tables.Records, rec.Envelope.RelayID)
		attemptKey := rec.Envelope.RelayID + "\x00" + hook + "\x00" + epoch
		l.derivedAttempts[attemptKey]++
		var err error
		switch hook {
		case "gate":
			if l.AfterGateResolution != nil {
				err = l.AfterGateResolution(rec)
			}
		case "approval":
			if l.AfterApprovalResolution != nil {
				err = l.AfterApprovalResolution(rec)
			}
		}
		if err != nil {
			if l.derivedAttempts[attemptKey] >= DerivedRetryCeiling {
				if parkErr := l.commitDerived(derived.ParkRecord(rec.Envelope.RelayID, "retry-ceiling")); parkErr != nil {
					return parkErr
				}
			}
			return err
		}
		if err := l.commitCursorAdvance(rec.Envelope.RelayID, hook); err != nil {
			return err
		}
		status = l.Tables.DerivedWork[rec.Envelope.RelayID]
	}
	return nil
}

func (l *Loop) commitCursorAdvance(sourceRelayID, hook string) error {
	return l.commitDerived(derived.CursorAdvanceRecord(sourceRelayID, []string{hook}))
}

func (l *Loop) commitDerived(rec record.Record) error {
	if violation := validateAutomaticTransition(l.Tables, rec); violation != nil {
		rec.Envelope.DeliveryState = record.Rejected
		rec.Body = formatVerdictViolation(*violation)
		if rec.Headers == nil {
			rec.Headers = map[string]string{}
		}
		rec.Headers["failing_edge"] = "derived-transition"
		if err := l.commitDerivedRecord(rec); err != nil {
			return err
		}
		return errors.New("derived transition rejected: " + violation.Class)
	}
	return l.commitDerivedRecord(rec)
}

// CommitAutomaticDerived applies the same live-fold validation used by the
// service loop for an exclusive offline writer such as the recovery ceremony.
func CommitAutomaticDerived(st *store.Store, rec record.Record) (string, error) {
	if st == nil {
		return "", errors.New("derived transition store unavailable")
	}
	tab, err := tables.Build(st)
	if err != nil {
		return "", err
	}
	if violation := validateAutomaticTransition(tab, rec); violation != nil {
		rec.Envelope.DeliveryState = record.Rejected
		rec.Body = formatVerdictViolation(*violation)
		if rec.Headers == nil {
			rec.Headers = map[string]string{}
		}
		rec.Headers["failing_edge"] = "derived-transition"
		if _, err := st.Commit(rec, nil); err != nil {
			return "", err
		}
		return "", errors.New("derived transition rejected: " + violation.Class)
	}
	return st.Commit(rec, nil)
}

func (l *Loop) commitDerivedRecord(rec record.Record) error {
	commit := l.DerivedCommit
	if commit == nil {
		commit = func(rec record.Record) (string, error) { return l.Store.Commit(rec, nil) }
	}
	relayID, err := commit(rec)
	if err != nil {
		return err
	}
	rec.Envelope.RelayID = relayID
	if l.Tables != nil {
		l.Tables.OnCommit(rec)
	}
	return l.completeTurn()
}

func validateAutomaticTransition(tab *tables.T, rec record.Record) *fieldspec.Violation {
	if rec.Envelope.DeliveryState != record.Accepted {
		return nil
	}
	kind := rec.Headers["record_kind"]
	if kind != "derived-work-attempt" && kind != "derived-work-transition" {
		return nil
	}
	if tab == nil {
		return &fieldspec.Violation{Field: "source_relay_id", Class: "unknown-target"}
	}
	var body struct {
		SourceRelayID  string   `json:"source_relay_id"`
		Hook           string   `json:"hook"`
		State          string   `json:"state"`
		Predecessor    string   `json:"predecessor"`
		Kind           string   `json:"kind"`
		CompletedHooks []string `json:"completed_hooks"`
	}
	if json.Unmarshal([]byte(rec.Body), &body) != nil || body.SourceRelayID == "" {
		return &fieldspec.Violation{Field: "body", Class: "typed"}
	}
	status, exists := tab.DerivedWork[body.SourceRelayID]
	if !exists {
		return &fieldspec.Violation{Field: "source_relay_id", Class: "unknown-target"}
	}
	if kind == "derived-work-attempt" {
		if status.Status != "pending" || body.State != "running_or_unknown" || body.Hook == "" || !containsHook(status.Cursor, body.Hook) {
			return &fieldspec.Violation{Field: "hook", Class: "stale-resolution"}
		}
		open, predecessor, valid := derived.AttemptState(tab.Records, body.SourceRelayID, body.Hook)
		if !valid {
			return &fieldspec.Violation{Field: "predecessor", Class: "conflicting-resolution"}
		}
		if open {
			return &fieldspec.Violation{Field: "predecessor", Class: "conflicting-resolution"}
		}
		if body.Predecessor != predecessor {
			return &fieldspec.Violation{Field: "predecessor", Class: "stale-resolution"}
		}
		return nil
	}
	switch body.Kind {
	case "cursor_advance":
		if status.Status != "pending" && status.Status != "unknown" {
			return &fieldspec.Violation{Field: "completed_hooks", Class: "stale-resolution"}
		}
		if len(body.CompletedHooks) == 0 {
			return &fieldspec.Violation{Field: "completed_hooks", Class: "required"}
		}
		for _, hook := range body.CompletedHooks {
			if !containsHook(status.Cursor, hook) {
				return &fieldspec.Violation{Field: "completed_hooks", Class: "duplicate-resolution"}
			}
		}
	case "parked":
		if status.Status == "failed" {
			return &fieldspec.Violation{Field: "source_relay_id", Class: "duplicate-resolution"}
		}
		if status.Status != "pending" {
			return &fieldspec.Violation{Field: "source_relay_id", Class: "stale-resolution"}
		}
	case "realized-undelivered":
		if status.Status == "failed" {
			return &fieldspec.Violation{Field: "source_relay_id", Class: "duplicate-resolution"}
		}
		if status.Status != "unknown" || !containsHook(status.Cursor, "mint") {
			return &fieldspec.Violation{Field: "source_relay_id", Class: "stale-resolution"}
		}
	default:
		return &fieldspec.Violation{Field: "kind", Class: "typed"}
	}
	return nil
}

func containsHook(hooks []string, want string) bool {
	for _, hook := range hooks {
		if hook == want {
			return true
		}
	}
	return false
}

func (l *Loop) completeTurn() (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			l.setClassGDirty(true)
			panic(recovered)
		}
		l.setClassGDirty(err != nil)
	}()
	return l.runClassG()
}

func (l *Loop) runClassG() error {
	if l.ClassGCompleteAuto != nil {
		if err := l.ClassGCompleteAuto(l.Store, l.Tables); err != nil {
			return err
		}
	} else if err := obligation.CompleteAuto(l.Store, l.Tables); err != nil {
		return err
	}
	if l.ClassGGC != nil {
		if err := l.ClassGGC(l.Store, l.Tables); err != nil {
			return err
		}
	}
	if l.ClassGBuildTables != nil {
		tab, err := l.ClassGBuildTables(l.Store)
		if err != nil {
			return err
		}
		l.Tables = tab
	}
	if l.ClassGArmScheduler != nil {
		if err := l.ClassGArmScheduler(); err != nil {
			return err
		}
	}
	if l.AfterCommit != nil {
		return l.AfterCommit(l.Store)
	}
	return nil
}

func (l *Loop) retryClassGIfDirty() bool {
	if l.ClassGDirty() {
		_ = l.completeTurn()
	}
	return !l.ClassGDirty()
}

func (l *Loop) classGOutcome(out Outcome) Outcome {
	if !l.ClassGDirty() || out.DecisionState == "" {
		return out
	}
	out.State = ""
	out.PostCommitState = "pending"
	if out.Reason == "" {
		out.Reason = safeReason("obligation-error")
	}
	return out
}

func (l *Loop) setClassGDirty(dirty bool) {
	l.classGMu.Lock()
	l.classGDirty = dirty
	if l.Tables == nil {
		l.Tables = tables.New()
	}
	l.Tables.ClassGDirty = dirty
	tab := l.Tables
	l.classGMu.Unlock()
	if l.ClassGPublishTables != nil {
		l.ClassGPublishTables(tab)
	}
}

func (l *Loop) ClassGDirty() bool {
	l.classGMu.RLock()
	defer l.classGMu.RUnlock()
	return l.classGDirty
}

// DrainClassG is the synchronous startup barrier. A panic becomes an error so
// the host returns before opening its socket while retaining the diagnostic.
func (l *Loop) DrainClassG() (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			l.setClassGDirty(true)
			err = fmt.Errorf("class-g-drain-panic: %v", recovered)
		}
	}()
	if err := l.completeTurn(); err != nil {
		return err
	}
	if err := l.foldRealizedMintEvidence(); err != nil {
		return err
	}
	l.retryAllBlind()
	return nil
}

func (l *Loop) foldRealizedMintEvidence() error {
	if l.Tables == nil || l.MintRealized == nil {
		return nil
	}
	var relayIDs []string
	for relayID, status := range l.Tables.DerivedWork {
		if status.Status == "unknown" && containsHook(status.Cursor, "mint") && derived.ResumableAttempt(l.Tables.Records, relayID, "mint") {
			relayIDs = append(relayIDs, relayID)
		}
	}
	sort.Strings(relayIDs)
	for _, relayID := range relayIDs {
		rec, ok := l.Tables.ByRelay[relayID]
		if !ok || !l.MintRealized(rec) {
			continue
		}
		if err := l.commitDerived(derived.RealizedUndeliveredRecord(relayID)); err != nil {
			return err
		}
	}
	return nil
}

func (l *Loop) faultOutcome(cmd intake.Cmd, reason string) Outcome {
	if existing, ok := l.existingOutcome(cmd.IntakeID); ok {
		return l.outcomeFromRecord(existing)
	}
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
		if existing, ok := l.existingOutcome(held.Envelope.IntakeID); ok {
			return l.outcomeFromRecord(existing)
		}
		relayID, err := l.Store.Commit(held, nil)
		if err != nil {
			return outcomeFor(record.Rejected, "complete", "", cmd.IntakeID, safeReason("commit-error"))
		}
		if l.Tables != nil {
			l.Tables.OnCommit(held)
		}
		return outcomeFor(record.Held, "complete", relayID, cmd.IntakeID, safeReason("internal-fault"))
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
	if existing, ok := l.existingOutcome(rejected.Envelope.IntakeID); ok {
		return l.outcomeFromRecord(existing)
	}
	relayID, err := l.Store.Commit(rejected, nil)
	if err != nil {
		return outcomeFor(record.Rejected, "complete", "", cmd.IntakeID, safeReason("commit-error"))
	}
	if l.Tables != nil {
		l.Tables.OnCommit(rejected)
	}
	detail := safeReason("internal-fault")
	out := outcomeFor(record.Rejected, "complete", relayID, cmd.IntakeID, detail)
	out.Detail = detail
	return out
}

func outcomeFor(decision, postCommit, relayID, intakeID, reason string) Outcome {
	out := Outcome{
		DecisionState:   decision,
		PostCommitState: postCommit,
		RelayID:         relayID,
		IntakeID:        intakeID,
		Reason:          reason,
	}
	if postCommit == "complete" {
		out.State = decision
	}
	return out
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
