package engine

import (
	"context"
	"encoding/json"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func SubmitHandler(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta) Handler {
	return func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		var cand record.Record
		if err := json.Unmarshal(cmd.Payload, &cand); err != nil {
			return rejected(cmd, meta, "bad submit payload"), nil, nil
		}
		cand = seat.Stamp(cand, meta)
		cand.Envelope.IntakeID = cmd.IntakeID
		if cand.Envelope.SchemaVersion == 0 {
			cand.Envelope.SchemaVersion = 1
		}
		violations := reg.Validate(cand, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, "")
		if len(violations) > 0 {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(anySlice(violations)...)
			return cand, nil, nil
		}
		if lineageBounce := lineage.Check(cand, st); lineageBounce != nil {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(lineageBounce)
			return cand, nil, nil
		}
		cand.Envelope.DeliveryState = record.Accepted
		return cand, nil, nil
	}
}

func rejected(cmd intake.Cmd, meta seat.SeatMeta, reason string) record.Record {
	return record.Record{
		Envelope: record.Envelope{
			RelayID:       "rejected-" + cmd.IntakeID,
			From:          meta.Name,
			Role:          meta.Role,
			DeliveryState: record.Rejected,
			IntakeID:      cmd.IntakeID,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "submit rejected"},
		Body:    reason,
	}
}

func anySlice[T any](values []T) []any {
	out := make([]any, len(values))
	for i, value := range values {
		out[i] = value
	}
	return out
}
