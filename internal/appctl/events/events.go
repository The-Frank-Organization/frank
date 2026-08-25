// Package events validates and durably queues worker-carried E0 events.
package events

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

var digestPattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

type Event struct {
	Schema               string  `json:"schema"`
	EventKind            string  `json:"event_kind"`
	Phase                string  `json:"phase"`
	Scope                string  `json:"scope"`
	RunID                string  `json:"run_id"`
	TurnID               string  `json:"turn_id"`
	AttemptID            string  `json:"attempt_id"`
	TurnEpoch            string  `json:"turn_epoch"`
	ProviderLaneID       string  `json:"provider_lane_id"`
	RunManifestDigest    string  `json:"run_manifest_digest"`
	PolicyDigest         string  `json:"policy_digest"`
	DenyReason           *string `json:"deny_reason,omitempty"`
	FrozenCoreDigest     *string `json:"frozen_core_digest,omitempty"`
	LogicalSurfaceDigest *string `json:"logical_surface_digest,omitempty"`
	EventEvidence        string  `json:"event_evidence"`
	EventIntegrity       string  `json:"event_integrity"`
	ReportedBy           string  `json:"reported_by"`
	EventTS              string  `json:"event_ts"`
}

type Request struct {
	EventID, RunID, TurnID, ReportedBy string
	Raw                                []byte
	At                                 int64
}

type Host struct{ applier *applier.Host }

func New(host *applier.Host) *Host { return &Host{applier: host} }

type persistEvent struct{ request Request }

func (event persistEvent) RunID() string { return event.request.RunID }

func (host *Host) Persist(ctx context.Context, request Request) (bool, error) {
	if host == nil || host.applier == nil || request.EventID == "" || request.RunID == "" || request.TurnID == "" || request.ReportedBy == "" || len(request.Raw) == 0 {
		return false, errors.New("events: invalid persist request")
	}
	parsed, err := Validate(request.Raw)
	if err != nil {
		return false, err
	}
	if parsed.RunID != request.RunID || parsed.TurnID != request.TurnID || parsed.ReportedBy != request.ReportedBy {
		return false, errors.New("events: outer identity mismatch")
	}
	result, err := host.applier.Apply(ctx, persistEvent{request: request})
	if err != nil {
		return false, err
	}
	return result.Value.(bool), nil
}

func (event persistEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	var raw []byte
	var runID, turnID, reportedBy string
	err := tx.QueryRowContext(ctx, `SELECT event_bytes,run_id,turn_id,reported_by FROM pending_app_events WHERE event_id=?`, request.EventID).Scan(&raw, &runID, &turnID, &reportedBy)
	if err == nil {
		if !bytes.Equal(raw, request.Raw) || runID != request.RunID || turnID != request.TurnID || reportedBy != request.ReportedBy {
			return applier.Result{}, errors.New("events: event identity conflict")
		}
		return applier.Result{Value: true, NoMutation: true}, nil
	}
	if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO pending_app_events(event_id,run_id,turn_id,event_bytes,reported_by,created_at) VALUES(?,?,?,?,?,?)`, request.EventID, request.RunID, request.TurnID, request.Raw, request.ReportedBy, request.At)
	return applier.Result{Value: false}, err
}

func Validate(raw []byte) (Event, error) {
	if err := rejectTopLevelDuplicates(raw); err != nil {
		return Event{}, err
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var event Event
	if err := decoder.Decode(&event); err != nil {
		return Event{}, fmt.Errorf("events: malformed event: %w", err)
	}
	if err := requireEOF(decoder); err != nil {
		return Event{}, err
	}
	if event.Schema != "m3.app_event.v1" && event.Schema != "m3.app_event.v2" {
		return Event{}, errors.New("events: unknown schema")
	}
	if event.EventKind != "provider_attempt" || event.Scope != "attempt" || event.EventEvidence != "E0" || event.EventIntegrity != "self_reported" {
		return Event{}, errors.New("events: invalid fixed field")
	}
	if event.RunID == "" || event.TurnID == "" || event.AttemptID == "" || event.ProviderLaneID == "" || event.ReportedBy == "" {
		return Event{}, errors.New("events: missing identity")
	}
	if _, err := appipc.ParseCounter(event.TurnEpoch); err != nil {
		return Event{}, err
	}
	if !digestPattern.MatchString(event.RunManifestDigest) || !digestPattern.MatchString(event.PolicyDigest) {
		return Event{}, errors.New("events: invalid required digest")
	}
	if event.FrozenCoreDigest != nil && !digestPattern.MatchString(*event.FrozenCoreDigest) {
		return Event{}, errors.New("events: invalid frozen core digest")
	}
	if event.Schema == "m3.app_event.v1" && event.LogicalSurfaceDigest != nil {
		return Event{}, errors.New("events: v1 forbids logical surface digest")
	}
	if event.Schema == "m3.app_event.v2" && (event.LogicalSurfaceDigest == nil || !digestPattern.MatchString(*event.LogicalSurfaceDigest)) {
		return Event{}, errors.New("events: v2 requires logical surface digest")
	}
	if !validPhase(event.Phase) {
		return Event{}, errors.New("events: invalid phase")
	}
	if event.Phase == "denied" {
		if event.DenyReason == nil || !validDeny(*event.DenyReason) {
			return Event{}, errors.New("events: denied requires closed reason")
		}
	} else if event.DenyReason != nil {
		return Event{}, errors.New("events: deny reason forbidden")
	}
	if _, err := time.Parse(time.RFC3339, event.EventTS); err != nil {
		return Event{}, errors.New("events: invalid event timestamp")
	}
	return event, nil
}

func rejectTopLevelDuplicates(raw []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	token, err := decoder.Token()
	if err != nil {
		return err
	}
	delimiter, ok := token.(json.Delim)
	if !ok || delimiter != '{' {
		return errors.New("events: event is not an object")
	}
	seen := map[string]struct{}{}
	for decoder.More() {
		token, err = decoder.Token()
		if err != nil {
			return err
		}
		name, ok := token.(string)
		if !ok {
			return errors.New("events: invalid member name")
		}
		if _, duplicate := seen[name]; duplicate {
			return errors.New("events: duplicate member")
		}
		seen[name] = struct{}{}
		var value json.RawMessage
		if err := decoder.Decode(&value); err != nil {
			return err
		}
	}
	if _, err := decoder.Token(); err != nil {
		return err
	}
	return requireEOF(decoder)
}

func requireEOF(decoder *json.Decoder) error {
	var extra any
	err := decoder.Decode(&extra)
	if err == io.EOF {
		return nil
	}
	if err == nil {
		return errors.New("events: trailing JSON value")
	}
	return err
}

func validPhase(value string) bool {
	switch value {
	case "denied", "sent", "completed", "failed", "cancelled", "unknown":
		return true
	}
	return false
}
func validDeny(value string) bool {
	switch value {
	case "policy-unavailable", "policy-digest-mismatch", "malformed-core", "lane-mismatch", "lane-endpoint-invalid", "endpoint-mismatch", "endpoint-not-allowlisted", "method-mismatch", "reserved-auth-header-in-core":
		return true
	}
	return false
}
