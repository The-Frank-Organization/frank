package engine

import (
	"context"
	"fmt"

	"github.com/jackli/frank/internal/bounce"
	frankconfig "github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func SubmitHandler(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, existing ...*tables.T) Handler {
	return SubmitHandlerWithRender(st, reg, meta, fieldspec.RenderEnv{}, existing...)
}

func SubmitHandlerWithRender(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, existing ...*tables.T) Handler {
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
		cand.Envelope.SchemaVersion = migrate.Current
		tab := firstSubmitTable(existing)
		if tab == nil {
			var err error
			tab, err = tables.Build(st)
			if err != nil {
				cand.Envelope.DeliveryState = record.Rejected
				cand.Body = bounce.Format(fieldspec.Violation{Field: "system", Class: "store-read-error"})
				return cand, nil, nil
			}
		}
		violations := reg.Validate(cand, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, formDigest, env, lineage.RealGrantState(tab))
		if len(violations) > 0 {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(anySlice(violations)...)
			return cand, nil, nil
		}
		cand = stampParent(st, tab, cand, meta, env)
		if violation := validateWaiverRows(tab, cand, meta); violation != nil {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(*violation)
			return cand, nil, nil
		}
		if lineageBounce := (&lineage.Engine{Reg: reg, T: tab}).Check(cand, seat.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}); lineageBounce != nil {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(lineageBounce)
			return cand, nil, nil
		}
		if violation := validateRecordKind(tab, cand); violation != nil {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(*violation)
			return cand, nil, nil
		}
		if cand.Headers["record_kind"] == "config_change" {
			cand = classifyConfigChange(st, cand, meta)
			if cand.Envelope.DeliveryState == record.Accepted {
				intents, err := store.ConfigChangeIntentsStrict(cand)
				return cand, intents, err
			}
			cand = clearGateRaiseHeaders(cand)
			return cand, nil, nil
		}
		if cand.Headers["resolves_gate"] != "" {
			cand = classifyVerdict(tab, cand, meta)
			if cand.Envelope.DeliveryState == record.Accepted {
				intents, err := store.DefaultProjectionIntentsStrict(cand)
				return cand, intents, err
			}
			cand = clearGateRaiseHeaders(cand)
			return cand, nil, nil
		}
		cand.Envelope.DeliveryState = record.Accepted
		intents, err := submitProjectionIntents(cand)
		if err != nil {
			return cand, nil, err
		}
		if isOwedKind(cand) {
			intents = append(intents, owedProjectionIntentsFromTable(tab, cand)...)
		}
		return cand, intents, nil
	}
}

func firstSubmitTable(existing []*tables.T) *tables.T {
	if len(existing) == 0 {
		return nil
	}
	return existing[0]
}

func clearGateRaiseHeaders(rec record.Record) record.Record {
	if rec.Headers == nil {
		return rec
	}
	if pick := rec.Headers["gate_category_pick"]; pick != "" {
		rec.Headers["gate_category"] = pick
	} else if rec.Headers["gate_category_raised"] != "" {
		delete(rec.Headers, "gate_category")
	}
	delete(rec.Headers, "gate_category_raised")
	delete(rec.Headers, "gate_category_pick")
	return rec
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

func submitProjectionIntents(rec record.Record) ([]store.Intent, error) {
	intents, err := store.DefaultProjectionIntentsStrict(rec)
	if err != nil {
		return nil, err
	}
	if isGateCandidate(rec) {
		intents = append(intents, store.Intent{
			Kind:    store.IntentIndex,
			Path:    "INDEX.md",
			Payload: []byte(fmt.Sprintf("| %s | parked | %s |\n", rec.Envelope.RelayID, rec.Envelope.From)),
		})
	}
	return intents, nil
}

func isGateCandidate(rec record.Record) bool {
	return rec.Headers["HUMAN_GATE_REQUIRED"] == "yes" || rec.Headers["gate_category"] != ""
}

func validateRecordKind(t *tables.T, cand record.Record) *fieldspec.Violation {
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
		var found bool
		for _, rec := range t.Records {
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
	case "config_change":
		return nil
	case "waiver_retraction":
		target := cand.Headers["retracts"]
		if target == "" {
			return &fieldspec.Violation{Field: "retracts", Class: "required", Reason: "retracts required"}
		}
		var found bool
		for _, rec := range t.Records {
			if rec.Envelope.RelayID == target && acceptedWaiverRecord(rec) {
				found = true
			}
			if rec.Headers["record_kind"] == "waiver_retraction" && rec.Headers["retracts"] == target && rec.Envelope.DeliveryState == record.Accepted {
				return &fieldspec.Violation{Field: "retracts", Class: "already-resolved", Reason: "waiver already retracted"}
			}
		}
		if !found {
			return &fieldspec.Violation{Field: "retracts", Class: lineage.ParentUnknownRecompose, Reason: "waiver unknown"}
		}
		return nil
	default:
		return nil
	}
}

func validateWaiverRows(t *tables.T, cand record.Record, meta seat.SeatMeta) *fieldspec.Violation {
	if operatorSeat(meta) {
		return nil
	}
	for _, field := range []string{"waiver_scope", "rationale", "retracts"} {
		if cand.Headers[field] != "" {
			return &fieldspec.Violation{Field: field, Class: "seat-scope", Reason: field + " requires operator"}
		}
	}
	return nil
}

func acceptedWaiverRecord(rec record.Record) bool {
	if rec.Envelope.DeliveryState != record.Accepted {
		return false
	}
	if !(rec.Envelope.From == "operator" || rec.Envelope.Role == "operator") {
		return false
	}
	return rec.Headers["waiver_scope"] != "" || rec.Headers["ORCH_REVIEW_WAIVER"] != ""
}

func isOwedKind(rec record.Record) bool {
	return rec.Headers["record_kind"] == "owed_item" || rec.Headers["record_kind"] == "owed_disposition"
}

func classifyConfigChange(st *store.Store, cand record.Record, meta seat.SeatMeta) record.Record {
	reject := func(field, class, reason string) record.Record {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(fieldspec.Violation{Field: field, Class: class, Reason: reason})
		return cand
	}
	if !(meta.IsOperator || meta.Name == "operator" || meta.Role == "operator") {
		return reject("record_kind", "seat-scope", "config_change requires operator")
	}
	member := cand.Headers["member"]
	if member != "fieldspec" && member != "engine" {
		return reject("member", "enum", "member must be fieldspec or engine")
	}
	if cand.Headers["new_digest"] == "" {
		return reject("new_digest", "required", "new_digest required")
	}
	digest, err := configDigestWithMember(st, member, []byte(cand.Body))
	if err != nil {
		return reject("new_digest", "config-read-error", "could not recompute config digest")
	}
	if cand.Headers["new_digest"] != digest {
		return reject("new_digest", "digest-mismatch", "new_digest does not match recomputed config digest")
	}
	cand.Envelope.DeliveryState = record.Accepted
	return cand
}

func configDigestWithMember(st *store.Store, member string, body []byte) (string, error) {
	if st == nil {
		return "", fmt.Errorf("store required")
	}
	pinned, err := frankconfig.Load(store.StoreRootConfigPaths(st.Root))
	if err != nil {
		return "", err
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return frankconfig.Digest(members), nil
}

func owedProjectionIntentsFromTable(t *tables.T, cand record.Record) []store.Intent {
	var records []record.Record
	if t != nil {
		records = append(records, t.Records...)
	}
	records = append(records, cand)
	return []store.Intent{store.OwedOpenProjectionIntent(records)}
}

func classifyVerdict(t *tables.T, cand record.Record, meta seat.SeatMeta) record.Record {
	if !(meta.IsOperator || meta.Name == "operator" || meta.Role == "operator") {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(fieldspec.Violation{Field: "record_kind", Class: "seat-scope", Reason: "gate_resolution requires operator"})
		return cand
	}
	gateRef := cand.Headers["resolves_gate"]
	parent := cand.Headers["PARENT_DISPATCH_ID"]
	if parent != gateRef {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(lineage.Bounce{Edge: "PARENT_DISPATCH_ID", Kind: lineage.ParentInvalidDeadEdge})
		return cand
	}
	var gateFound bool
	var wakeSeat string
	for _, existing := range t.Records {
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
