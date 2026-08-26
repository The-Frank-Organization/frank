package supervisor

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
)

type BrokerFailureClass string

const (
	BrokerSpawnFail        BrokerFailureClass = "spawn-fail"
	BrokerNoReady          BrokerFailureClass = "no-ready"
	BrokerMalformedReady   BrokerFailureClass = "malformed-ready"
	BrokerReadyCrash       BrokerFailureClass = "ready-crash"
	BrokerReattachDeadline BrokerFailureClass = "reattach-deadline"
)

type BrokerFailureRequest struct {
	RunID, InstanceID string
	Class             BrokerFailureClass
	At                int64
}

type BrokerFailureResult struct {
	Failures     uint64
	BackoffUntil *int64
	Terminal     bool
	Idempotent   bool
}

type brokerFailureEvent struct{ request BrokerFailureRequest }

func (event brokerFailureEvent) RunID() string { return event.request.RunID }

func (controller *Controller) RecordBrokerFailure(ctx context.Context, request BrokerFailureRequest) (BrokerFailureResult, error) {
	if controller == nil || controller.applier == nil || request.RunID == "" || request.InstanceID == "" || !validBrokerFailureClass(request.Class) {
		return BrokerFailureResult{}, errors.New("supervisor: invalid broker failure")
	}
	result, err := controller.applier.Apply(ctx, brokerFailureEvent{request: request})
	if err != nil {
		return BrokerFailureResult{}, err
	}
	return result.Value.(BrokerFailureResult), nil
}

func (event brokerFailureEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	request := event.request
	eventID := brokerFailureEventID(request)
	body := []byte(request.Class)
	var stored []byte
	err := tx.QueryRowContext(ctx, `SELECT event_bytes FROM pending_app_events WHERE event_id=?`, eventID).Scan(&stored)
	if err == nil {
		if string(stored) != string(body) {
			return applier.Result{}, errors.New("supervisor: broker failure identity mismatch")
		}
		var failuresStored string
		if err := tx.QueryRowContext(ctx, `SELECT consecutive_failures FROM runs WHERE run_id=?`, request.RunID).Scan(&failuresStored); err != nil {
			return applier.Result{}, err
		}
		failures, err := parseStored(failuresStored)
		if err != nil {
			return applier.Result{}, err
		}
		return applier.Result{Value: BrokerFailureResult{Failures: failures, Idempotent: true}, NoMutation: true}, nil
	}
	if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	var failuresStored string
	if err := tx.QueryRowContext(ctx, `SELECT consecutive_failures FROM runs WHERE run_id=?`, request.RunID).Scan(&failuresStored); err != nil {
		return applier.Result{}, err
	}
	failures, err := parseStored(failuresStored)
	if err != nil || failures == math.MaxUint64 {
		return applier.Result{}, errors.New("supervisor: failure counter exhausted")
	}
	failures++
	if _, err := tx.ExecContext(ctx, `INSERT INTO pending_app_events(event_id,run_id,event_bytes,reported_by,created_at) VALUES(?,?,?,?,?)`, eventID, request.RunID, body, "broker-supervisor", request.At); err != nil {
		return applier.Result{}, err
	}
	if failures >= MaxConsecutiveFailures {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED',consecutive_failures=?,backoff_until=NULL,updated_at=? WHERE run_id=?`, pad(failures), request.At, request.RunID); err != nil {
			return applier.Result{}, err
		}
		return applier.Result{Value: BrokerFailureResult{Failures: failures, Terminal: true}}, nil
	}
	backoff := backoffDeadline(request.At, failures)
	if _, err := tx.ExecContext(ctx, `UPDATE runs SET consecutive_failures=?,backoff_until=?,updated_at=? WHERE run_id=?`, pad(failures), backoff, request.At, request.RunID); err != nil {
		return applier.Result{}, err
	}
	return applier.Result{Value: BrokerFailureResult{Failures: failures, BackoffUntil: backoff}}, nil
}

func validBrokerFailureClass(class BrokerFailureClass) bool {
	switch class {
	case BrokerSpawnFail, BrokerNoReady, BrokerMalformedReady, BrokerReadyCrash, BrokerReattachDeadline:
		return true
	default:
		return false
	}
}

func brokerFailureEventID(request BrokerFailureRequest) string {
	digest := sha256.Sum256([]byte(request.RunID + "\x00" + request.InstanceID + "\x00" + string(request.Class)))
	return "broker-failure-" + hex.EncodeToString(digest[:])
}
