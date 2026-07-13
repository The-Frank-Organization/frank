package engine

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackli/frank/internal/bounce"
	frankconfig "github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func SubmitHandler(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, existing ...*tables.T) Handler {
	return SubmitHandlerWithRender(st, reg, meta, fieldspec.RenderEnv{}, existing...)
}

func SubmitHandlerWithRender(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, existing ...*tables.T) Handler {
	return SubmitHandlerWithObservation(st, reg, meta, env, observe.Env{PresentLayers: env.PresentLayers}, existing...)
}

func SubmitHandlerWithObservation(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, observeEnv observe.Env, existing ...*tables.T) Handler {
	return submitHandlerWithObservation(st, reg, meta, env, observeEnv, nil, existing...)
}

func SubmitHandlerWithClaimRegistry(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, checks *observe.Registry, existing ...*tables.T) Handler {
	observeEnv := observe.Env{PresentLayers: env.PresentLayers}
	if checks != nil {
		observeEnv.Evaluate = func(candidate observe.Candidate) observe.PredicateResult {
			return checks.EvaluateCandidateClaims(candidate)
		}
	}
	return submitHandlerWithObservation(st, reg, meta, env, observeEnv, checks, existing...)
}

func submitHandlerWithObservation(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, observeEnv observe.Env, checks *observe.Registry, existing ...*tables.T) Handler {
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
		if observeEnv.PresentLayers["observe"] {
			if field := observe.LaneSuppliedSystemField(reg, cand); field != "" {
				cand.Envelope.DeliveryState = record.Rejected
				cand.Body = bounce.Format(fieldspec.Violation{Field: field, Class: "lane-supplied-system-field"})
				return cand, nil, nil
			}
		}
		if env.PreActive {
			if rec, intents, handled, err := classifyBootAdmission(reg, cand, meta, formDigest, env); handled {
				return rec, intents, err
			}
		}
		violations := reg.Validate(cand, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, formDigest, env, lineage.RealGrantState(tab))
		if rawClaims, claimsPresent := cand.Headers["executable_claims"]; len(violations) == 0 && observeEnv.PresentLayers["observe"] && checks != nil && claimsPresent {
			if _, issue := checks.ValidateClaims(rawClaims); issue != nil {
				violations = append(violations, fieldspec.Violation{Field: issue.ClaimRef, Class: issue.Class})
			}
		}
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
		if violation := validateRecordKind(tab, cand, meta); violation != nil {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = bounce.Format(*violation)
			return cand, nil, nil
		}
		if observeEnv.PresentLayers["observe"] {
			if cand.Headers == nil {
				cand.Headers = map[string]string{}
			}
			cand.Headers["authority_class"] = AuthorityClass(reg, cand, meta)
		}
		observeResult, terminal := observe.Gate(cand, meta.Name, cand.Headers["PHASE"], cand.Headers["AUTHORITY"], observeEnv)
		if observeEnv.PresentLayers["observe"] {
			var violation *fieldspec.Violation
			cand, violation = CompleteObserved(cand, observeResult.ObservedFields)
			if violation != nil {
				cand = clearGateRaiseHeaders(cand)
				cand.Envelope.DeliveryState = record.Rejected
				cand.Body = bounce.Format(*violation)
				return cand, nil, nil
			}
		}
		if terminal != record.Accepted {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = terminal
			failureClass := observeResult.FailureClass
			if failureClass == "" {
				failureClass = "observed-false"
			}
			cand.Body = bounce.Format(fieldspec.Violation{Field: observeResult.FailingPredicate, Class: failureClass})
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
			cand = classifyVerdict(tab, cand, meta, observeEnv)
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

func classifyBootAdmission(reg *fieldspec.Registry, cand record.Record, meta seat.SeatMeta, formDigest string, env fieldspec.RenderEnv) (record.Record, []store.Intent, bool, error) {
	bootEnv := env
	bootEnv.PreActive = env.PreActive
	_, currentDigest := reg.Render(bootEnv, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, cand.Headers["PHASE"], cand.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	if formDigest == "" || formDigest != currentDigest {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(fieldspec.Violation{Field: "form_digest", Class: "re-render", Reason: "stale form digest"})
		return cand, nil, true, nil
	}
	violations := bootAdmissionViolations(cand)
	if len(violations) > 0 {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = bounce.Format(anySlice(violations)...)
		return cand, nil, true, nil
	}
	cand.Envelope.DeliveryState = record.Accepted
	intents, err := submitProjectionIntents(cand)
	return cand, intents, true, err
}

func bootAdmissionViolations(cand record.Record) []fieldspec.Violation {
	allowed := map[string]bool{
		"PHASE":           true,
		"CEREMONY_TIER":   true,
		"SUBJECT":         true,
		"charter_loaded":  true,
		"dispatch_status": true,
	}
	var violations []fieldspec.Violation
	for key := range cand.Headers {
		if !allowed[key] {
			violations = append(violations, fieldspec.Violation{Field: key, Class: "non-boot-before-active"})
		}
	}
	for _, key := range []string{"PHASE", "CEREMONY_TIER", "SUBJECT", "charter_loaded", "dispatch_status"} {
		if cand.Headers[key] == "" {
			violations = append(violations, fieldspec.Violation{Field: key, Class: "non-boot-before-active"})
		}
	}
	if cand.Headers["PHASE"] != "" && cand.Headers["PHASE"] != "SITREP" {
		violations = append(violations, fieldspec.Violation{Field: "PHASE", Class: "non-boot-before-active"})
	}
	if value := cand.Headers["charter_loaded"]; value != "" && value != "yes" && value != "no" {
		violations = append(violations, fieldspec.Violation{Field: "charter_loaded", Class: "non-boot-before-active"})
	}
	if value := cand.Headers["dispatch_status"]; value != "" && value != "read" && value != "awaiting" {
		violations = append(violations, fieldspec.Violation{Field: "dispatch_status", Class: "non-boot-before-active"})
	}
	return violations
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

func validateRecordKind(t *tables.T, cand record.Record, meta seat.SeatMeta) *fieldspec.Violation {
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
	case "seat_mint":
		_, violation := ParseSeatMintBody(cand.Body, meta.Name)
		return violation
	default:
		return nil
	}
}

type SeatMintRequest struct {
	Seat       string
	Role       string
	IsOperator bool
}

func ParseSeatMintBody(body, requester string) (SeatMintRequest, *fieldspec.Violation) {
	var raw struct {
		Seat       string `json:"seat"`
		Role       string `json:"role"`
		IsOperator *bool  `json:"is_operator"`
	}
	if err := json.Unmarshal([]byte(body), &raw); err != nil {
		return SeatMintRequest{}, &fieldspec.Violation{Field: "body", Class: "typed", Reason: "seat_mint body must be JSON"}
	}
	if raw.Seat == "" {
		return SeatMintRequest{}, &fieldspec.Violation{Field: "seat", Class: "required", Reason: "seat required"}
	}
	if raw.Role == "" {
		return SeatMintRequest{}, &fieldspec.Violation{Field: "role", Class: "required", Reason: "role required"}
	}
	if raw.IsOperator == nil {
		return SeatMintRequest{}, &fieldspec.Violation{Field: "is_operator", Class: "required", Reason: "is_operator required"}
	}
	if raw.Seat == "system" {
		return SeatMintRequest{}, &fieldspec.Violation{Field: "seat", Class: "reserved", Reason: "system seat is reserved"}
	}
	if requester != "" && raw.Seat == requester {
		return SeatMintRequest{}, &fieldspec.Violation{Field: "seat", Class: "self-remint", Reason: "seat_mint cannot re-mint the submitting seat"}
	}
	return SeatMintRequest{Seat: raw.Seat, Role: raw.Role, IsOperator: *raw.IsOperator}, nil
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
	if member == "adoption" {
		return reject("member", "offline-bless-only", "adoption requires offline bless")
	}
	if member != "fieldspec" && member != "engine" && member != "catalog" {
		return reject("member", "enum", "unknown config member")
	}
	if cand.Headers["new_digest"] == "" {
		return reject("new_digest", "required", "new_digest required")
	}
	pinned, err := frankconfig.Load(store.StoreRootConfigPaths(st.Root))
	if err != nil {
		return reject("member", "config-read-error", "could not load pinned config")
	}
	current, ok := pinned.Members[member]
	if !ok {
		return reject("member", "not-adopted", "config member requires adoption")
	}
	digest, err := configDigestWithPinnedMember(pinned, member, []byte(cand.Body))
	if err != nil {
		return reject("new_digest", "config-read-error", "could not recompute config digest")
	}
	if cand.Headers["new_digest"] != digest {
		return reject("new_digest", "digest-mismatch", "new_digest does not match recomputed config digest")
	}
	if err := frankconfig.ValidateMemberTransition(member, current, []byte(cand.Body)); err != nil {
		return reject("member", "config-version-transition", "config member transition refused")
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
	return configDigestWithPinnedMember(pinned, member, body)
}

func configDigestWithPinnedMember(pinned *frankconfig.Pinned, member string, body []byte) (string, error) {
	if pinned == nil {
		return "", fmt.Errorf("pinned config required")
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

func classifyVerdict(t *tables.T, cand record.Record, meta seat.SeatMeta, observeEnv observe.Env) record.Record {
	if !(meta.IsOperator || meta.Name == "operator" || meta.Role == "operator") {
		cand.Envelope.DeliveryState = record.Rejected
		cand.Body = formatVerdictViolation(fieldspec.Violation{Field: "record_kind", Class: "seat-scope", Reason: "gate_resolution requires operator"})
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
	var gateRecord record.Record
	var odb record.Record
	for _, existing := range t.Records {
		if existing.Envelope.RelayID == gateRef && existing.Envelope.DeliveryState == record.Accepted && isGateCandidate(existing) {
			gateFound = true
			wakeSeat = existing.Envelope.From
			gateRecord = existing
		}
		if existing.Envelope.RelayID == "odb-"+gateRef && existing.Envelope.DeliveryState == record.Accepted && existing.Headers["record_kind"] == "odb" {
			odb = existing
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
	if odb.Envelope.RelayID != "" {
		reply, violation := ParseODBReply(cand.Body)
		if violation == nil {
			violation = ValidateODBChoice(odb, reply.Choice)
		}
		if violation != nil {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = formatVerdictViolation(*violation)
			return cand
		}
	}
	if observeEnv.PresentLayers["observe"] {
		observeResult, terminal := observe.Gate(
			gateRecord,
			gateRecord.Envelope.From,
			gateRecord.Headers["PHASE"],
			gateRecord.Headers["AUTHORITY"],
			observeEnv,
		)
		if terminal != record.Accepted {
			failureClass := observeResult.FailureClass
			if failureClass == "" {
				failureClass = "observed-false"
			}
			cand.Envelope.DeliveryState = terminal
			cand.Body = formatVerdictViolation(fieldspec.Violation{Field: observeResult.FailingPredicate, Class: failureClass})
			return cand
		}
	}
	cand.Envelope.To = wakeSeat
	cand.Envelope.DeliveryState = record.Accepted
	return cand
}

func formatVerdictViolation(violation fieldspec.Violation) string {
	return bounce.Format(violation)
}

func anySlice[T any](values []T) []any {
	out := make([]any, len(values))
	for i, value := range values {
		out[i] = value
	}
	return out
}
