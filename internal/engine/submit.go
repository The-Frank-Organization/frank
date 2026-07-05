package engine

import (
	"context"
	"fmt"

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
		cand, formDigest, err := fieldspec.DecodeSubmitPayload(cmd.Payload)
		if err != nil {
			return rejected(cmd, meta, "bad submit payload"), nil, nil
		}
		cand = seat.Stamp(cand, meta)
		relayID, err := st.NewRelayID()
		if err != nil {
			return rejected(cmd, meta, safeReason("id-error")), nil, nil
		}
		cand.Envelope.RelayID = relayID
		cand.Envelope.IntakeID = cmd.IntakeID
		if cand.Envelope.SchemaVersion == 0 {
			cand.Envelope.SchemaVersion = 1
		}
		violations := reg.Validate(cand, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, formDigest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
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
		if violation := validateRecordKind(st, cand); violation != nil {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(*violation)
			return cand, nil, nil
		}
		if cand.Headers["resolves_gate"] != "" {
			cand = classifyVerdict(st, cand)
			if cand.Envelope.DeliveryState == record.Accepted {
				return cand, store.DefaultProjectionIntents(cand), nil
			}
			return cand, nil, nil
		}
		cand.Envelope.DeliveryState = record.Accepted
		intents := submitProjectionIntents(cand)
		if isOwedKind(cand) {
			intents = append(intents, store.OwedProjectionIntentsForCandidate(st, cand)...)
		}
		return cand, intents, nil
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

func submitProjectionIntents(rec record.Record) []store.Intent {
	intents := store.DefaultProjectionIntents(rec)
	if isGateCandidate(rec) {
		intents = append(intents, store.Intent{
			Kind:    store.IntentIndex,
			Path:    "INDEX.md",
			Payload: []byte(fmt.Sprintf("| %s | parked | %s |\n", rec.Envelope.RelayID, rec.Envelope.From)),
		})
	}
	return intents
}

func isGateCandidate(rec record.Record) bool {
	return rec.Headers["HUMAN_GATE_REQUIRED"] == "yes" || rec.Headers["gate_category"] != ""
}

func validateRecordKind(st *store.Store, cand record.Record) *fieldspec.Violation {
	switch cand.Headers["record_kind"] {
	case "":
		return nil
	case "owed_item":
		for _, field := range []string{"owner", "source", "target_surface", "disposition_path"} {
			if cand.Headers[field] == "" {
				return &fieldspec.Violation{Field: field, Class: "required", Reason: field + " required"}
			}
		}
		return nil
	case "owed_disposition":
		target := cand.Headers["disposes_owed"]
		if target == "" {
			return &fieldspec.Violation{Field: "disposes_owed", Class: "required", Reason: "disposes_owed required"}
		}
		records, err := st.Records()
		if err != nil {
			return &fieldspec.Violation{Field: "record_kind", Class: "store-read-error", Reason: "store read error"}
		}
		var found bool
		for _, rec := range records {
			if rec.Headers["record_kind"] == "owed_item" && rec.Envelope.RelayID == target && rec.Envelope.DeliveryState == record.Accepted {
				found = true
			}
			if rec.Headers["record_kind"] == "owed_disposition" && rec.Headers["disposes_owed"] == target && rec.Envelope.DeliveryState == record.Accepted {
				return &fieldspec.Violation{Field: "disposes_owed", Class: "already-resolved", Reason: "owed item already disposed"}
			}
		}
		if !found {
			return &fieldspec.Violation{Field: "disposes_owed", Class: lineage.ParentUnknownRecompose, Reason: "owed item unknown"}
		}
		return nil
	default:
		return &fieldspec.Violation{Field: "record_kind", Class: "unknown", Reason: "unknown record_kind"}
	}
}

func isOwedKind(rec record.Record) bool {
	return rec.Headers["record_kind"] == "owed_item" || rec.Headers["record_kind"] == "owed_disposition"
}

func classifyVerdict(st *store.Store, cand record.Record) record.Record {
	gateRef := cand.Headers["resolves_gate"]
	parent := cand.Headers["PARENT_DISPATCH_ID"]
	if parent != gateRef {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(lineage.Bounce{Edge: "PARENT_DISPATCH_ID", Kind: lineage.ParentInvalidDeadEdge})
		return cand
	}
	records, err := st.Records()
	if err != nil {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = safeReason("store-read-error")
		return cand
	}
	var gateFound bool
	var wakeSeat string
	for _, existing := range records {
		if existing.Envelope.RelayID == gateRef && existing.Envelope.DeliveryState == record.Accepted && isGateCandidate(existing) {
			gateFound = true
			wakeSeat = existing.Envelope.From
		}
		if existing.Headers["resolves_gate"] == gateRef && existing.Envelope.DeliveryState == record.Accepted {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(fieldspec.Violation{Field: "resolves_gate", Class: "already-resolved"})
			return cand
		}
	}
	if !gateFound {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(lineage.Bounce{Edge: "PARENT_DISPATCH_ID", Kind: lineage.ParentUnknownRecompose})
		return cand
	}
	cand.Envelope.To = wakeSeat
	cand.Envelope.DeliveryState = record.Accepted
	return cand
}

func anySlice[T any](values []T) []any {
	out := make([]any, len(values))
	for i, value := range values {
		out[i] = value
	}
	return out
}
