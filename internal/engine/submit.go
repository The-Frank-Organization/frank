package engine

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/jackli/frank/internal/bounce"
	frankconfig "github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/derived"
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
	return submitHandlerWithObservation(st, reg, meta, env, observeEnv, nil, migrate.New(), existing...)
}

// SubmitHandlerWithMigration exposes the read-side migration registry to the
// live operator-verdict path. Production v1 callers continue through the
// default empty registry; the first breaking bump supplies its governed steps.
func SubmitHandlerWithMigration(st *store.Store, reg *fieldspec.Registry, migrations *migrate.Registry, meta seat.SeatMeta, existing ...*tables.T) Handler {
	return submitHandlerWithObservation(st, reg, meta, fieldspec.RenderEnv{}, observe.Env{}, nil, migrations, existing...)
}

func SubmitHandlerWithClaimRegistry(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, checks *observe.Registry, existing ...*tables.T) Handler {
	observeEnv := observe.Env{PresentLayers: env.PresentLayers}
	if checks != nil {
		observeEnv.Evaluate = func(candidate observe.Candidate) observe.PredicateResult {
			return checks.EvaluateCandidateClaims(candidate)
		}
	}
	return submitHandlerWithObservation(st, reg, meta, env, observeEnv, checks, migrate.New(), existing...)
}

func submitHandlerWithObservation(st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, env fieldspec.RenderEnv, observeEnv observe.Env, checks *observe.Registry, migrations *migrate.Registry, existing ...*tables.T) Handler {
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
				cand = rejectAtEdge(cand, "form-validation", bounce.Format(fieldspec.Violation{Field: field, Class: "lane-supplied-system-field"}))
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
			cand = rejectAtEdge(cand, "form-validation", bounce.Format(anySlice(violations)...))
			return cand, nil, nil
		}
		cand = stampParent(st, tab, cand, meta, env)
		if violation := validateWaiverRows(tab, cand, meta); violation != nil {
			cand = rejectAtEdge(cand, "form-validation", bounce.Format(*violation))
			return cand, nil, nil
		}
		if lineageBounce := (&lineage.Engine{Reg: reg, T: tab}).Check(cand, seat.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}); lineageBounce != nil {
			cand = rejectAtEdge(cand, "lineage", bounce.Format(lineageBounce))
			return cand, nil, nil
		}
		if violation := validateRecordKind(tab, cand, meta); violation != nil {
			edge := "form-validation"
			if violation.Class == "anchor-target-resolved" {
				edge = "anchor-target-resolved"
			}
			cand = rejectAtEdge(cand, edge, bounce.Format(*violation))
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
				cand = rejectAtEdge(cand, "declared-vs-observed", bounce.Format(*violation))
				return cand, nil, nil
			}
		}
		if terminal != record.Accepted {
			failureClass := observeResult.FailureClass
			if failureClass == "" {
				failureClass = "observed-false"
			}
			cand = failAtEdge(cand, terminal, "observe-predicate", bounce.Format(fieldspec.Violation{Field: observeResult.FailingPredicate, Class: failureClass}))
			return cand, nil, nil
		}
		if cand.Headers["record_kind"] == "seat_mint" {
			expected, violation := expectedMintPredecessor(tab.Records, cand)
			if violation != nil {
				cand = rejectAtEdge(cand, "mint-predecessor-mismatch", formatVerdictViolation(*violation))
				return cand, nil, nil
			}
			cand.Headers["mint_predecessor"] = expected
		}
		if cand.Headers["record_kind"] == "config_change" {
			cand = classifyConfigChange(st, cand, meta)
			if cand.Envelope.DeliveryState == record.Accepted {
				intents, err := store.ConfigChangeIntentsStrict(cand)
				return cand, intents, err
			}
			cand = failAtEdge(cand, cand.Envelope.DeliveryState, "form-validation", cand.Body)
			return cand, nil, nil
		}
		if cand.Headers["resolves_gate"] != "" {
			cand = classifyVerdict(tab, cand, meta, observeEnv, migrations)
			if cand.Envelope.DeliveryState == record.Accepted {
				intents, err := store.DefaultProjectionIntentsStrict(cand)
				return cand, intents, err
			}
			if cand.Headers["failing_edge"] == "" {
				cand = failAtEdge(cand, cand.Envelope.DeliveryState, "form-validation", cand.Body)
			} else {
				cand = clearGateRaiseHeaders(cand)
			}
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
		cand = rejectAtEdge(cand, "form-validation", bounce.Format(fieldspec.Violation{Field: "form_digest", Class: "re-render", Reason: "stale form digest"}))
		return cand, nil, true, nil
	}
	violations := bootAdmissionViolations(cand)
	if len(violations) > 0 {
		cand = rejectAtEdge(cand, "form-validation", bounce.Format(anySlice(violations)...))
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

// revalidateAtCommit closes the stale-snapshot window created when the
// serialized loop services a nested command while an outer handler is
// blocked. It reruns only uniqueness and pinned-config checks that can become
// false after handler evaluation; observations and external execution are not
// repeated.
func revalidateAtCommit(st *store.Store, cand record.Record, intents []store.Intent) (record.Record, []store.Intent, error) {
	if cand.Envelope.DeliveryState != record.Accepted {
		return cand, intents, nil
	}
	kind := cand.Headers["record_kind"]
	if cand.Headers["resolves_gate"] == "" && kind != "owed_disposition" && kind != "waiver_retraction" && kind != "config_change" && kind != "seat_mint" && kind != "mint-chain-anchor" && kind != "attempt_resolution" {
		return cand, intents, nil
	}
	tab, err := tables.Build(st)
	if err != nil {
		return cand, nil, err
	}
	meta := seat.SeatMeta{
		Name:       cand.Envelope.From,
		Role:       cand.Envelope.Role,
		IsOperator: cand.Envelope.From == "operator" || cand.Envelope.Role == "operator",
	}
	if cand.Headers["resolves_gate"] != "" {
		for _, existing := range tab.Records {
			if existing.Envelope.DeliveryState == record.Accepted && existing.Headers["resolves_gate"] == cand.Headers["resolves_gate"] {
				cand = clearGateRaiseHeaders(cand)
				cand.Envelope.DeliveryState = record.Rejected
				cand.Body = formatVerdictViolation(fieldspec.Violation{Field: "resolves_gate", Class: "already-resolved"})
				return cand, nil, nil
			}
		}
	}
	switch kind {
	case "owed_disposition", "waiver_retraction", "attempt_resolution", "mint-chain-anchor":
		if violation := validateRecordKind(tab, cand, meta); violation != nil {
			cand = clearGateRaiseHeaders(cand)
			cand.Envelope.DeliveryState = record.Rejected
			if violation.Class == "anchor-target-resolved" {
				cand.Headers["failing_edge"] = "anchor-target-resolved"
			}
			cand.Body = formatVerdictViolation(*violation)
			return cand, nil, nil
		}
	case "seat_mint":
		expected, violation := expectedMintPredecessor(tab.Records, cand)
		if violation != nil || (cand.Headers["mint_predecessor"] != "" && cand.Headers["mint_predecessor"] != expected) {
			if violation == nil {
				violation = &fieldspec.Violation{Field: "mint_predecessor", Class: "mint-predecessor-mismatch", Reason: "mint predecessor is not the current chain tip"}
			}
			cand.Envelope.DeliveryState = record.Rejected
			cand.Headers["failing_edge"] = "mint-predecessor-mismatch"
			cand.Body = formatVerdictViolation(*violation)
			return cand, nil, nil
		}
		cand.Headers["mint_predecessor"] = expected
		if violation := ValidateCeremonyRetryAuthority(tab.Records, cand); violation != nil {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Headers["failing_edge"] = "retry-authority-delta"
			cand.Body = formatVerdictViolation(*violation)
			return cand, nil, nil
		}
		if violation := validateRecordKind(tab, cand, meta); violation != nil {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = formatVerdictViolation(*violation)
			return cand, nil, nil
		}
	case "config_change":
		checked := classifyConfigChange(st, cand, meta)
		if checked.Envelope.DeliveryState != record.Accepted {
			return clearGateRaiseHeaders(checked), nil, nil
		}
	}
	return cand, intents, nil
}

// ValidateCeremonyRetryAuthority is the commit-time belt-and-braces check for
// the offline ceremony writer. A delivery retry is a generation rotation, not
// an authority edit: all authority-bearing body members must equal the
// predecessor tip byte-for-byte.
func ValidateCeremonyRetryAuthority(records []record.Record, cand record.Record) *fieldspec.Violation {
	if cand.Envelope.DeliveryState != record.Accepted || cand.Headers["record_kind"] != "seat_mint" || cand.Headers["admin_provenance"] != "ceremony" {
		return nil
	}
	predecessorID := cand.Headers["mint_predecessor"]
	var predecessor record.Record
	for _, rec := range records {
		if rec.Envelope.RelayID == predecessorID && rec.Envelope.DeliveryState == record.Accepted && rec.Headers["record_kind"] == "seat_mint" {
			predecessor = rec
			break
		}
	}
	if predecessor.Envelope.RelayID == "" {
		return nil
	}
	prior, priorViolation := ParseSeatMintBody(predecessor.Body, predecessor.Envelope.From)
	next, nextViolation := ParseSeatMintBody(cand.Body, cand.Envelope.From)
	if priorViolation != nil || nextViolation != nil {
		return nil
	}
	if prior.Seat != next.Seat || prior.Role != next.Role || prior.IsOperator != next.IsOperator {
		return &fieldspec.Violation{Field: "seat_mint", Class: "retry-authority-delta", Reason: "ceremony retry authority must equal predecessor"}
	}
	return nil
}

type MintChain struct {
	Tip        record.Record
	Conflicted bool
}

type mintChainAnchorBody struct {
	Seat    string `json:"seat"`
	Selects string `json:"selects"`
}

func MintChainAnchorRecord(seatName, selects string) record.Record {
	body, _ := json.Marshal(mintChainAnchorBody{Seat: seatName, Selects: selects})
	return record.Record{
		Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "legacy mint chain anchored", "record_kind": "mint-chain-anchor"},
		Body:     string(body),
	}
}

// BuildMintChains derives each seat's canonical pivot tip from set relations,
// never relay-ID or redo order. Multiple predecessor-less legacy pivots remain
// conflicted until the governed upgrade-anchor fold selects one.
func BuildMintChains(records []record.Record) (map[string]MintChain, error) {
	type pivot struct {
		rec  record.Record
		req  SeatMintRequest
		pred string
	}
	bySeat := map[string][]pivot{}
	anchors := map[string][]mintChainAnchorBody{}
	for _, rec := range records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["record_kind"] == "mint-chain-anchor" {
			var anchor mintChainAnchorBody
			if json.Unmarshal([]byte(rec.Body), &anchor) != nil || anchor.Seat == "" || anchor.Selects == "" {
				continue
			}
			anchors[anchor.Seat] = append(anchors[anchor.Seat], anchor)
			continue
		}
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" {
			continue
		}
		req, violation := ParseSeatMintBody(rec.Body, rec.Envelope.From)
		if violation != nil {
			return nil, fmt.Errorf("accepted seat_mint failed parse: %s:%s", violation.Field, violation.Class)
		}
		bySeat[req.Seat] = append(bySeat[req.Seat], pivot{rec: rec, req: req, pred: rec.Headers["mint_predecessor"]})
	}
	chains := make(map[string]MintChain, len(bySeat))
	for seatName, pivots := range bySeat {
		nodes := map[string]pivot{}
		var legacy []pivot
		for _, item := range pivots {
			nodes[item.rec.Envelope.RelayID] = item
			if item.pred == "" {
				legacy = append(legacy, item)
			}
		}
		seatAnchors := anchors[seatName]
		if len(seatAnchors) > 1 {
			chains[seatName] = MintChain{Conflicted: true}
			continue
		}
		var anchored *pivot
		if len(seatAnchors) == 1 {
			selected, ok := nodes[seatAnchors[0].Selects]
			if !ok || selected.pred != "" {
				chains[seatName] = MintChain{Conflicted: true}
				continue
			}
			anchored = &selected
		} else if len(legacy) > 1 {
			chains[seatName] = MintChain{Conflicted: true}
			continue
		}
		children := map[string][]string{}
		conflicted := false
		for _, item := range pivots {
			if item.pred == "" {
				continue
			}
			if item.pred != "genesis" {
				if parent, ok := nodes[item.pred]; !ok || parent.req.Seat != seatName {
					conflicted = true
				}
			} else if len(legacy) != 0 || anchored != nil {
				conflicted = true
			}
			children[item.pred] = append(children[item.pred], item.rec.Envelope.RelayID)
			if len(children[item.pred]) > 1 {
				conflicted = true
			}
		}
		start := "genesis"
		visited := map[string]bool{}
		var tip record.Record
		if anchored != nil {
			start = anchored.rec.Envelope.RelayID
			visited[start] = true
			tip = anchored.rec
		} else if len(legacy) == 1 {
			start = legacy[0].rec.Envelope.RelayID
			visited[start] = true
			tip = legacy[0].rec
		}
		current := start
		for len(children[current]) == 1 {
			next := children[current][0]
			if visited[next] {
				conflicted = true
				break
			}
			visited[next] = true
			tip = nodes[next].rec
			current = next
		}
		expectedVisited := len(pivots)
		if anchored != nil {
			expectedVisited = 1
			for _, item := range pivots {
				if item.pred != "" {
					expectedVisited++
				}
			}
		}
		if len(visited) != expectedVisited {
			conflicted = true
		}
		if conflicted {
			tip = record.Record{}
		}
		chains[seatName] = MintChain{Tip: tip, Conflicted: conflicted}
	}
	return chains, nil
}

func expectedMintPredecessor(records []record.Record, cand record.Record) (string, *fieldspec.Violation) {
	req, violation := ParseSeatMintBody(cand.Body, cand.Envelope.From)
	if violation != nil {
		return "", violation
	}
	chains, err := BuildMintChains(records)
	if err != nil {
		return "", &fieldspec.Violation{Field: "mint_predecessor", Class: "mint-predecessor-mismatch", Reason: err.Error()}
	}
	chain, exists := chains[req.Seat]
	if !exists {
		return "genesis", nil
	}
	if chain.Conflicted || chain.Tip.Envelope.RelayID == "" {
		return "", &fieldspec.Violation{Field: "mint_predecessor", Class: "mint-predecessor-mismatch", Reason: "mint chain has no unique tip"}
	}
	return chain.Tip.Envelope.RelayID, nil
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
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "submit rejected", "failing_edge": "form-validation"},
		Body:    reason,
	}
}

func rejectAtEdge(rec record.Record, edge, body string) record.Record {
	return failAtEdge(rec, record.Rejected, edge, body)
}

func failAtEdge(rec record.Record, state, edge, body string) record.Record {
	rec = clearGateRaiseHeaders(rec)
	if rec.Headers == nil {
		rec.Headers = map[string]string{}
	}
	rec.Headers["failing_edge"] = edge
	rec.Envelope.DeliveryState = state
	rec.Body = body
	return rec
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
	case "mint-chain-anchor":
		return validateMintChainAnchor(t, cand, meta)
	case "attempt_resolution":
		return validateAttemptResolution(t, cand, meta)
	default:
		return nil
	}
}

func validateMintChainAnchor(t *tables.T, cand record.Record, meta seat.SeatMeta) *fieldspec.Violation {
	if !operatorSeat(meta) {
		return &fieldspec.Violation{Field: "record_kind", Class: "seat-scope", Reason: "mint-chain-anchor requires operator"}
	}
	var body mintChainAnchorBody
	if err := json.Unmarshal([]byte(cand.Body), &body); err != nil {
		return &fieldspec.Violation{Field: "body", Class: "typed", Reason: "mint-chain-anchor body must be a JSON object"}
	}
	if body.Seat == "" {
		return &fieldspec.Violation{Field: "seat", Class: "required", Reason: "seat required"}
	}
	if body.Selects == "" {
		return &fieldspec.Violation{Field: "selects", Class: "required", Reason: "selects required"}
	}
	for _, rec := range t.Records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		if rec.Headers["record_kind"] == "mint-chain-anchor" {
			var existing mintChainAnchorBody
			if json.Unmarshal([]byte(rec.Body), &existing) == nil && existing.Seat == body.Seat {
				return &fieldspec.Violation{Field: "selects", Class: "anchor-target-resolved", Reason: "mint chain already has an anchor"}
			}
		}
	}
	chains, err := BuildMintChains(t.Records)
	if err != nil {
		return &fieldspec.Violation{Field: "selects", Class: "unknown-target", Reason: err.Error()}
	}
	chain, exists := chains[body.Seat]
	if !exists {
		return &fieldspec.Violation{Field: "selects", Class: "unknown-target", Reason: "mint chain unknown"}
	}
	if !chain.Conflicted {
		return &fieldspec.Violation{Field: "selects", Class: "anchor-target-resolved", Reason: "mint chain already resolved"}
	}
	legacyCount := 0
	selectionValid := false
	for _, rec := range t.Records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" || rec.Headers["mint_predecessor"] != "" {
			continue
		}
		req, violation := ParseSeatMintBody(rec.Body, rec.Envelope.From)
		if violation == nil && req.Seat == body.Seat {
			legacyCount++
			selectionValid = selectionValid || rec.Envelope.RelayID == body.Selects
		}
	}
	if legacyCount < 2 || !selectionValid {
		return &fieldspec.Violation{Field: "selects", Class: "unknown-target", Reason: "anchor must select a conflicted legacy pivot"}
	}
	trialCandidate := cand
	trialCandidate.Envelope.DeliveryState = record.Accepted
	trial := append(append([]record.Record(nil), t.Records...), trialCandidate)
	trialChains, err := BuildMintChains(trial)
	if err != nil {
		return &fieldspec.Violation{Field: "selects", Class: "unknown-target", Reason: err.Error()}
	}
	resolved, exists := trialChains[body.Seat]
	if !exists || resolved.Conflicted || resolved.Tip.Envelope.RelayID == "" {
		return &fieldspec.Violation{Field: "selects", Class: "unknown-target", Reason: "anchor selection does not resolve the mint chain"}
	}
	return nil
}

type attemptResolutionBody struct {
	Resolves    string `json:"resolves"`
	Disposition string `json:"disposition"`
	EvidenceRef string `json:"evidence_ref"`
}

func parseAttemptResolutionBody(body string) (attemptResolutionBody, map[string]json.RawMessage, *fieldspec.Violation) {
	var members map[string]json.RawMessage
	if err := json.Unmarshal([]byte(body), &members); err != nil || members == nil {
		return attemptResolutionBody{}, nil, &fieldspec.Violation{Field: "body", Class: "typed", Reason: "attempt_resolution body must be a JSON object"}
	}
	var parsed attemptResolutionBody
	if raw, ok := members["resolves"]; !ok || json.Unmarshal(raw, &parsed.Resolves) != nil || parsed.Resolves == "" {
		return attemptResolutionBody{}, members, &fieldspec.Violation{Field: "resolves", Class: "required", Reason: "resolves required"}
	}
	if raw, ok := members["disposition"]; ok && json.Unmarshal(raw, &parsed.Disposition) != nil {
		return attemptResolutionBody{}, members, &fieldspec.Violation{Field: "disposition", Class: "typed", Reason: "disposition must be text"}
	}
	if raw, ok := members["evidence_ref"]; ok && json.Unmarshal(raw, &parsed.EvidenceRef) != nil {
		return attemptResolutionBody{}, members, &fieldspec.Violation{Field: "evidence_ref", Class: "typed", Reason: "evidence_ref must be text"}
	}
	return parsed, members, nil
}

func validateAttemptResolution(t *tables.T, cand record.Record, meta seat.SeatMeta) *fieldspec.Violation {
	if !operatorSeat(meta) {
		return &fieldspec.Violation{Field: "record_kind", Class: "seat-scope", Reason: "attempt_resolution requires operator"}
	}
	body, members, violation := parseAttemptResolutionBody(cand.Body)
	if violation != nil {
		return violation
	}
	target, ok := t.ByRelay[body.Resolves]
	if !ok || target.Envelope.DeliveryState != record.Accepted {
		return &fieldspec.Violation{Field: "resolves", Class: "unknown-target", Reason: "resolution target unknown"}
	}
	targetType, sourceRelayID, hook := resolutionTarget(target)
	if targetType == "" {
		return &fieldspec.Violation{Field: "resolves", Class: "unknown-target", Reason: "resolution target unknown"}
	}
	for _, rec := range t.Records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "attempt_resolution" {
			continue
		}
		prior, _, issue := parseAttemptResolutionBody(rec.Body)
		if issue != nil || prior.Resolves != body.Resolves {
			continue
		}
		class := "duplicate-resolution"
		if targetType == "marker" && prior.Disposition != body.Disposition {
			class = "conflicting-resolution"
		}
		return &fieldspec.Violation{Field: "resolves", Class: class, Reason: "resolution target already resolved"}
	}
	switch targetType {
	case "marker":
		if len(members) != 3 || body.Disposition == "" || body.EvidenceRef == "" {
			return &fieldspec.Violation{Field: "body", Class: "typed", Reason: "marker resolution requires resolves, disposition, and evidence_ref"}
		}
		if body.Disposition != "effect-confirmed-realized" && body.Disposition != "effect-confirmed-unrealized" {
			return &fieldspec.Violation{Field: "disposition", Class: "typed", Reason: "unknown attempt disposition"}
		}
		status, present := derived.Fold(t.Records)[sourceRelayID]
		if !present || status.Status != "unknown" || !containsHook(status.Cursor, hook) {
			return &fieldspec.Violation{Field: "resolves", Class: "stale-resolution", Reason: "attempt is no longer current"}
		}
	case "park":
		if len(members) != 1 {
			return &fieldspec.Violation{Field: "body", Class: "typed", Reason: "park reopen requires resolves only"}
		}
		status, present := derived.Fold(t.Records)[sourceRelayID]
		if !present || status.Status != "failed" {
			return &fieldspec.Violation{Field: "resolves", Class: "stale-resolution", Reason: "park is no longer current"}
		}
	}
	return nil
}

func resolutionTarget(rec record.Record) (targetType, sourceRelayID, hook string) {
	var body struct {
		SourceRelayID string `json:"source_relay_id"`
		Hook          string `json:"hook"`
		Kind          string `json:"kind"`
	}
	if json.Unmarshal([]byte(rec.Body), &body) != nil || body.SourceRelayID == "" {
		return "", "", ""
	}
	switch rec.Headers["record_kind"] {
	case "derived-work-attempt":
		if body.Hook != "" {
			return "marker", body.SourceRelayID, body.Hook
		}
	case "derived-work-transition":
		if body.Kind == "parked" {
			return "park", body.SourceRelayID, ""
		}
	}
	return "", "", ""
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

func classifyVerdict(t *tables.T, cand record.Record, meta seat.SeatMeta, observeEnv observe.Env, migrationRegistry ...*migrate.Registry) record.Record {
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
		if violation != nil {
			cand.Envelope.DeliveryState = record.Rejected
			cand.Body = formatVerdictViolation(*violation)
			return cand
		}
		migrated, compatible := guardedMigratedODB(odb, firstMigrationRegistry(migrationRegistry))
		if !compatible {
			return staleChoiceCandidate(cand, gateRecord, odb, migrated)
		}
		violation = ValidateODBChoice(migrated, reply.Choice)
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

type staleChoiceIntent struct {
	Reason                string `json:"reason"`
	SourceGate            string `json:"source_gate"`
	ReplacementDecisionID string `json:"replacement_decision_id"`
	Choices               string `json:"choices"`
	TargetSchemaVersion   int    `json:"target_schema_version"`
}

func staleChoiceCandidate(cand, gate, sourceODB, migratedODB record.Record) record.Record {
	choices := ""
	for _, candidate := range []record.Record{migratedODB, sourceODB, gate} {
		if _, violation := decisionProjection(record.Record{Headers: map[string]string{
			"record_kind": "odb", "choices": candidate.Headers["choices"],
		}}); violation == nil {
			choices = candidate.Headers["choices"]
			break
		}
	}
	targetVersion := migratedODB.Envelope.SchemaVersion
	if targetVersion == 0 {
		targetVersion = sourceODB.Envelope.SchemaVersion
	}
	replacementID := replacementDecisionID(gate.Envelope.RelayID, choices, targetVersion)
	intent := staleChoiceIntent{
		Reason: "stale_choice_set", SourceGate: gate.Envelope.RelayID,
		ReplacementDecisionID: replacementID, Choices: choices, TargetSchemaVersion: targetVersion,
	}
	body, err := json.Marshal(intent)
	if err != nil {
		return rejectAtEdge(cand, "stale_choice_set", formatVerdictViolation(fieldspec.Violation{Field: "choices", Class: "typed-parse"}))
	}
	cand = rejectAtEdge(cand, "stale_choice_set", string(body))
	cand.Headers["subject_ref"] = gate.Envelope.RelayID
	return cand
}

func replacementDecisionID(gateID, choices string, targetVersion int) string {
	payload, _ := json.Marshal(struct {
		GateID        string `json:"gate_id"`
		Choices       string `json:"choices"`
		TargetVersion int    `json:"target_version"`
	}{gateID, choices, targetVersion})
	sum := sha256.Sum256(payload)
	return fmt.Sprintf("reissue-%x", sum[:12])
}

func firstMigrationRegistry(registries []*migrate.Registry) *migrate.Registry {
	if len(registries) == 0 || registries[0] == nil {
		return migrate.New()
	}
	return registries[0]
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
