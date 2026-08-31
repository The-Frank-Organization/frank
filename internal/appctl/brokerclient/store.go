package brokerclient

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type Client struct{ applier *applier.Host }

func New(host *applier.Host) *Client { return &Client{applier: host} }

type controlEvent struct {
	runID, token string
	at           int64
}

func (event controlEvent) RunID() string { return event.runID }
func (event controlEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	var stored string
	err := tx.QueryRowContext(ctx, `SELECT control_generation FROM broker_control WHERE singleton=1`).Scan(&stored)
	generation := uint64(1)
	if err == nil {
		wire, decodeErr := appipc.UnpadCounter(stored)
		if decodeErr != nil {
			return applier.Result{}, decodeErr
		}
		previous, decodeErr := appipc.ParseCounter(wire)
		if decodeErr != nil || previous == math.MaxUint64 {
			return applier.Result{}, errors.New("brokerclient: control generation exhausted")
		}
		generation = previous + 1
	} else if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	padded, _ := appipc.PadCounter(appipc.FormatCounter(generation))
	_, err = tx.ExecContext(ctx, `INSERT INTO broker_control(singleton,control_token,control_generation,minted_at) VALUES(1,?,?,?)
		ON CONFLICT(singleton) DO UPDATE SET control_token=excluded.control_token,control_generation=excluded.control_generation,minted_at=excluded.minted_at`, event.token, padded, event.at)
	return applier.Result{Value: appipc.FormatCounter(generation)}, err
}

func (client *Client) AdvanceControl(ctx context.Context, runID, token string, at int64) (string, error) {
	if client == nil || client.applier == nil || runID == "" || token == "" {
		return "", errors.New("brokerclient: invalid control advance")
	}
	result, err := client.applier.Apply(ctx, controlEvent{runID: runID, token: token, at: at})
	if err != nil {
		return "", err
	}
	return result.Value.(string), nil
}

type Event struct {
	BrokerInstanceNonce string
	EventSeq            string
	Type                string
	RunID               string
	TurnEpoch           string
	Body                []byte
	At                  int64
}

type eventResult struct {
	ack       []byte
	duplicate bool
}

type recordEvent struct{ event Event }

func (event recordEvent) RunID() string { return event.event.RunID }
func (event recordEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	in := event.event
	if in.Type != "boundary_cut" && in.Type != "epoch_installed" {
		return applier.Result{}, errors.New("brokerclient: unknown broker event")
	}
	if _, err := appipc.ParseCounter(in.EventSeq); err != nil {
		return applier.Result{}, err
	}
	if _, err := appipc.ParseCounter(in.TurnEpoch); err != nil {
		return applier.Result{}, err
	}
	var storedBody, storedAck []byte
	err := tx.QueryRowContext(ctx, `SELECT event_bytes,ack_bytes FROM broker_events WHERE broker_instance_nonce=? AND event_seq=?`, in.BrokerInstanceNonce, padCounter(in.EventSeq)).Scan(&storedBody, &storedAck)
	if err == nil {
		if string(storedBody) != string(in.Body) {
			return applier.Result{}, errors.New("brokerclient: duplicate broker event identity mismatch")
		}
		return applier.Result{Value: eventResult{ack: storedAck, duplicate: true}, NoMutation: true}, nil
	}
	if !store.IsNoRows(err) {
		return applier.Result{}, err
	}
	ack, err := appipc.MarshalJCS(map[string]any{"broker_instance_nonce": in.BrokerInstanceNonce, "event_seq": in.EventSeq})
	if err != nil {
		return applier.Result{}, err
	}
	_, err = tx.ExecContext(ctx, `INSERT INTO broker_events(broker_instance_nonce,event_seq,event_type,run_id,turn_epoch,event_bytes,ack_bytes,committed_at) VALUES(?,?,?,?,?,?,?,?)`, in.BrokerInstanceNonce, padCounter(in.EventSeq), in.Type, in.RunID, padCounter(in.TurnEpoch), in.Body, ack, in.At)
	return applier.Result{Value: eventResult{ack: ack}}, err
}

func (client *Client) RecordEvent(ctx context.Context, event Event) ([]byte, bool, error) {
	if client == nil || client.applier == nil || event.BrokerInstanceNonce == "" || event.RunID == "" || len(event.Body) == 0 {
		return nil, false, errors.New("brokerclient: invalid event")
	}
	result, err := client.applier.Apply(ctx, recordEvent{event: event})
	if err != nil {
		return nil, false, err
	}
	value, ok := result.Value.(eventResult)
	if !ok {
		return nil, false, fmt.Errorf("brokerclient: invalid event result")
	}
	return value.ack, value.duplicate, nil
}

func padCounter(wire string) string {
	padded, _ := appipc.PadCounter(wire)
	return padded
}
