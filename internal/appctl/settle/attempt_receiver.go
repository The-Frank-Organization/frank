package settle

import (
	"context"
	"errors"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

// ReceiveAttemptResult decodes the connector-owned disposition carrier, then
// resolves the durable attempt identity before applying it to the row.
func (host *Host) ReceiveAttemptResult(ctx context.Context, payload []byte) (CarriageDecision, error) {
	if host == nil || host.applier == nil {
		return "", errors.New("settle: attempt result host is unavailable")
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		return "", err
	}
	envelope, err := registry.Decode(payload)
	if err != nil {
		return "", err
	}
	body, ok := envelope.Body.(*appipc.AttemptResultBody)
	if envelope.Channel != appipc.ChannelCtrlC || envelope.Type != "attempt_result" || !ok {
		return "", errors.New("settle: invalid attempt result envelope")
	}
	type identity struct{ runID, turnID string }
	value, err := host.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var found identity
		err := snapshot.QueryRowContext(ctx, `SELECT run_id,turn_id FROM provider_attempts WHERE attempt_id=?`, body.AttemptID).Scan(&found.runID, &found.turnID)
		if store.IsNoRows(err) {
			return identity{}, nil
		}
		return found, err
	}))
	if err != nil {
		return "", err
	}
	resolved := value.(identity)
	if resolved.runID == "" {
		return CarriageUnknown, nil
	}
	disposition := body.Disposition
	if opening := strings.IndexByte(disposition, '('); opening >= 0 {
		disposition = disposition[:opening]
	}
	return host.RecordAttemptResult(ctx, AttemptResultRequest{
		RunID: resolved.runID, TurnID: resolved.turnID, AttemptID: body.AttemptID,
		TurnEpoch: body.TurnEpoch, Disposition: disposition, CancelPoint: body.CancelPoint, RefusalStage: body.RefusalStage,
		FrozenCoreDigest: body.FrozenCoreDigest, ProviderLoweredToolsDigest: body.ProviderLoweredToolsDigest,
	})
}
