package scheduler

import (
	"context"
	"errors"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appipc"
)

// ReceiveEpochQuery resolves connector uncertainty from the serialized
// control-plane epoch row; it never accepts the connector's proposed epoch as
// truth.
func (scheduler *Scheduler) ReceiveEpochQuery(ctx context.Context, payload []byte) ([]byte, error) {
	if scheduler == nil || scheduler.applier == nil {
		return nil, errors.New("scheduler: epoch query host is unavailable")
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		return nil, err
	}
	envelope, err := registry.Decode(payload)
	if err != nil {
		return nil, err
	}
	body, ok := envelope.Body.(*appipc.EpochQueryBody)
	if envelope.Channel != appipc.ChannelCtrlC || envelope.Type != "epoch_query" || !ok || body.RunID == "" {
		return nil, errors.New("scheduler: invalid epoch query")
	}
	value, err := scheduler.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var stored string
		err := snapshot.QueryRowContext(ctx, `SELECT turn_epoch FROM epochs WHERE run_id=?`, body.RunID).Scan(&stored)
		return stored, err
	}))
	if err != nil {
		return nil, err
	}
	current, err := appipc.UnpadCounter(value.(string))
	if err != nil {
		return nil, err
	}
	replyBody := appipc.EpochUpdateBody{RunID: body.RunID, TurnEpoch: current}
	return registry.Encode(appipc.Envelope{
		V: 1, Channel: appipc.ChannelCtrlC, Type: "epoch_update", Seq: envelope.Seq,
		RunID: &replyBody.RunID, TurnEpoch: &replyBody.TurnEpoch, Body: replyBody,
	})
}
