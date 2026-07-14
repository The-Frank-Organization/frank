package observe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type CheckEntry struct {
	ID               string
	Rung             string
	Class            string
	ExecutorRequired bool
	ParamSchema      map[string]string
	Produces         []string
	TimeoutClass     string
}

type Selection struct {
	CheckID         string
	ClaimRef        string
	CandidateDigest string
	Seat            string
	Params          map[string]string
}

type ExpiryAction string

const (
	ExpiryKill   ExpiryAction = "kill"
	ExpiryExtend ExpiryAction = "extend"
)

type ExpiryRequest struct {
	Selection   Selection
	SoftExpiry  time.Duration
	HardCeiling time.Duration
}

type ExpiryDecision struct {
	Action ExpiryAction
}

type ExpiryDisposition func(context.Context, ExpiryRequest) ExpiryDecision

type ApprovalScope string

const (
	ApprovalOnce     ApprovalScope = "allow-once"
	ApprovalForEntry ApprovalScope = "allow-for-entry"
	ApprovalDenied   ApprovalScope = "deny"
)

type ApprovalRequest struct {
	Selection Selection
}

type ApprovalDecision struct {
	Allowed bool
	Scope   ApprovalScope
}

type ApprovalDisposition func(context.Context, ApprovalRequest) ApprovalDecision

type ClaimIssue struct {
	ClaimRef string
	Class    string
}

type CheckVerdict struct {
	CheckID       string `json:"check_id"`
	ClaimRef      string `json:"claim_ref"`
	Outcome       string `json:"outcome"`
	RungReached   string `json:"rung_reached"`
	Predicate     string `json:"predicate"`
	Timing        string `json:"timing"`
	FailingDetail string `json:"failing_detail"`
}

type SuiteExecutor interface {
	Spawn(CheckEntry, Selection) CheckVerdict
}

type RegistryEnv struct {
	Lanes                map[string]string
	SchemaRefs           map[string]string
	NamedSuites          map[string]bool
	Executor             SuiteExecutor
	GitExecutable        string
	ReadTimeout          time.Duration
	HardCeiling          time.Duration
	OnSoftExpiry         ExpiryDisposition
	Context              context.Context
	OnSideEffectApproval ApprovalDisposition
	// ReadFileStageHook is an injected blocking seam for proving that the
	// read-file control path detaches every filesystem-operation stage.
	ReadFileStageHook func(ReadFileStage)
	// FSStageHook is the generalized injected seam shared by every detachable
	// descriptor-rooted filesystem operation.
	FSStageHook   func(FSStage)
	findRefLimits *findRefLimits
}

type Registry struct {
	entries map[string]CheckEntry
	env     RegistryEnv
	fsMu    sync.Mutex
	fsLane  map[string]fsLaneState
}

func NewRegistry(env RegistryEnv) *Registry {
	gitExecutable := env.GitExecutable
	if gitExecutable == "" {
		gitExecutable = "git"
	}
	readTimeout := env.ReadTimeout
	if readTimeout <= 0 {
		readTimeout = readCheckTimeout
	}
	approvalContext := env.Context
	if approvalContext == nil {
		approvalContext = context.Background()
	}
	return &Registry{
		env: RegistryEnv{
			Lanes:                cloneMap(env.Lanes),
			SchemaRefs:           cloneMap(env.SchemaRefs),
			NamedSuites:          cloneBoolMap(env.NamedSuites),
			Executor:             env.Executor,
			GitExecutable:        gitExecutable,
			ReadTimeout:          readTimeout,
			HardCeiling:          env.HardCeiling,
			OnSoftExpiry:         env.OnSoftExpiry,
			Context:              approvalContext,
			OnSideEffectApproval: env.OnSideEffectApproval,
			ReadFileStageHook:    env.ReadFileStageHook,
			FSStageHook:          env.FSStageHook,
			findRefLimits:        env.findRefLimits,
		},
		fsLane: make(map[string]fsLaneState),
		entries: map[string]CheckEntry{
			"read-file": {
				ID: "read-file", Rung: "E1", Class: "base", TimeoutClass: "read_short",
				ParamSchema: map[string]string{"lane_ref": "registry-id", "path": "relative-path", "expect": "line|hash|schema_ref"},
				Produces:    []string{"achieved_evidence", "evidence_integrity"},
			},
			"git-status": {
				ID: "git-status", Rung: "E1", Class: "base", TimeoutClass: "read_short",
				ParamSchema: map[string]string{"lane_ref": "registry-id", "expect": "clean|dirty"},
				Produces:    []string{"achieved_evidence", "evidence_integrity"},
			},
			"find-references": {
				ID: "find-references", Rung: "E1", Class: "base", TimeoutClass: "read_short",
				ParamSchema: map[string]string{"symbol": "identifier-token", "lane_ref": "registry-id", "expect": "count:0"},
				Produces:    []string{"achieved_evidence"},
			},
			"run-suite": {
				ID: "run-suite", Rung: "E2", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded",
				ParamSchema: map[string]string{"target": "named-suite", "expect_green": "bool"},
				Produces:    []string{"achieved_evidence", "executable_claim_results"},
			},
			"run-suite-unbounded": {
				ID: "run-suite-unbounded", Rung: "E2", Class: "side_effecting",
				ExecutorRequired: true, TimeoutClass: "unbounded",
				ParamSchema: map[string]string{"target": "named-suite", "expect_green": "bool"},
				Produces:    []string{"achieved_evidence", "executable_claim_results"},
			},
		},
	}
}

func (r *Registry) Entry(id string) (CheckEntry, bool) {
	entry, ok := r.entries[id]
	if !ok {
		return CheckEntry{}, false
	}
	entry.ParamSchema = cloneMap(entry.ParamSchema)
	entry.Produces = append([]string(nil), entry.Produces...)
	return entry, true
}

func (r *Registry) Run(selection Selection) CheckVerdict {
	return r.runBound(selection).Verdict
}

func (r *Registry) runBound(selection Selection) boundVerdict {
	entry, ok := r.entries[selection.CheckID]
	if ok && entry.ID == "find-references" && r.env.Lanes[selection.Params["lane_ref"]] == "" {
		return boundVerdict{Verdict: refusedVerdictWithDetail(selection, "lane-ungoverned"), Origin: originConductorConfig, Entry: entry, Selected: selection}
	}
	if ok && entry.ID == "read-file" && strings.HasPrefix(selection.Params["expect"], "schema_ref:") {
		if r.env.SchemaRefs[strings.TrimPrefix(selection.Params["expect"], "schema_ref:")] == "" {
			return boundVerdict{Verdict: refusedVerdictWithDetail(selection, "schema-ref-unknown"), Origin: originConductorConfig, Entry: entry, Selected: selection}
		}
	}
	if !ok || !r.validParams(entry, selection.Params) {
		return boundVerdict{Verdict: refusedVerdict(selection), Origin: originConductorPolicy, Entry: entry, Selected: selection}
	}
	if entry.Class == "side_effecting" {
		if r.env.OnSideEffectApproval == nil {
			return boundVerdict{Verdict: refusedVerdictWithDetail(selection, "side-effecting-unapproved"), Origin: originConductorPolicy, Entry: entry, Selected: selection}
		}
		decision := r.env.OnSideEffectApproval(r.env.Context, ApprovalRequest{Selection: selection})
		if !decision.Allowed || decision.Scope != ApprovalOnce && decision.Scope != ApprovalForEntry {
			return boundVerdict{Verdict: refusedVerdictWithDetail(selection, "side-effecting-unapproved"), Origin: originConductorPolicy, Entry: entry, Selected: selection}
		}
		// Approval deliberately lifts only the prompt gate; real side-effecting execution remains a registered carry.
		return boundVerdict{Verdict: refusedVerdictWithDetail(selection, "side-effecting-execution-refused"), Origin: originConductorPolicy, Entry: entry, Selected: selection}
	}
	switch selection.CheckID {
	case "read-file":
		verdict := r.runReadFile(selection)
		origin := originBaseCheck
		if strings.HasPrefix(verdict.FailingDetail, "check-machinery-") {
			origin = originConductorMachinery
		}
		return boundVerdict{Verdict: verdict, Origin: origin, Entry: entry, Selected: selection}
	case "git-status":
		verdict := r.runGitStatus(selection)
		origin := originBaseCheck
		if strings.HasPrefix(verdict.FailingDetail, "check-machinery-") {
			origin = originConductorMachinery
		}
		return boundVerdict{Verdict: verdict, Origin: origin, Entry: entry, Selected: selection}
	case "find-references":
		result := r.executeFindReferences(selection)
		switch result.kind {
		case fsResultMachinery:
			return boundVerdict{Verdict: CheckVerdict{CheckID: selection.CheckID, ClaimRef: selection.ClaimRef, Outcome: "unsafe", RungReached: "none", Predicate: Blocked, Timing: result.timing, FailingDetail: result.detail}, Origin: originConductorMachinery, Entry: entry, Selected: selection}
		case fsResultDegraded:
			return boundVerdict{Verdict: CheckVerdict{CheckID: selection.CheckID, ClaimRef: selection.ClaimRef, Outcome: "skipped", RungReached: "none", Predicate: Degraded, FailingDetail: result.detail}, Origin: originBaseCheck, Entry: entry, Selected: selection}
		case fsResultData:
			return boundVerdict{Verdict: baseVerdict(selection, result.count == 0, "find-references-count-nonzero"), Origin: originBaseCheck, Entry: entry, Selected: selection}
		default:
			return boundVerdict{Verdict: CheckVerdict{CheckID: selection.CheckID, ClaimRef: selection.ClaimRef, Outcome: "unsafe", RungReached: "none", Predicate: Blocked, FailingDetail: "check-machinery-find-references-io"}, Origin: originConductorMachinery, Entry: entry, Selected: selection}
		}
	case "run-suite", "run-suite-unbounded":
		if r.env.Executor != nil {
			return boundVerdict{Verdict: r.env.Executor.Spawn(entry, selection), Origin: originExecutor, Entry: entry, Selected: selection}
		}
		fallthrough
	default:
		return boundVerdict{Verdict: CheckVerdict{
			CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
			Outcome: "unsafe", Predicate: Blocked, FailingDetail: "check-machinery-executor-required",
		}, Origin: originConductorMachinery, Entry: entry, Selected: selection}
	}
}

func (r *Registry) Evaluator(selection Selection) func(Candidate) PredicateResult {
	return func(candidate Candidate) PredicateResult {
		bound := selection
		bound.Seat = candidate.Seat
		if data, err := json.Marshal(candidate.Record); err == nil {
			sum := sha256.Sum256(data)
			bound.CandidateDigest = hex.EncodeToString(sum[:])
		}
		validated := validateBoundVerdict(r.runBound(bound))
		verdict := validated.Verdict
		id := bound.ClaimRef
		if id == "" {
			id = bound.CheckID
		}
		return PredicateResult{
			ID: id, Predicate: verdict.Predicate, MachineryFault: validated.MachineryFault,
			Verdicts: []CheckVerdict{verdict}, claimBindings: map[string]claimBinding{id: {RungReached: verdict.RungReached, SignalClass: validated.SignalClass, Integrity: validated.Integrity}},
		}
	}
}

func (r *Registry) ValidateClaims(raw string) ([]Selection, *ClaimIssue) {
	if raw == "" {
		return nil, &ClaimIssue{ClaimRef: "executable_claims", Class: "check-params-invalid"}
	}
	var rows []map[string]string
	if err := json.Unmarshal([]byte(raw), &rows); err != nil {
		return nil, &ClaimIssue{ClaimRef: "executable_claims", Class: "check-params-invalid"}
	}
	seen := map[string]bool{}
	selections := make([]Selection, 0, len(rows))
	for _, row := range rows {
		claimRef := row["claim_ref"]
		if len(row) != 3 || !validClaimRef(claimRef) {
			return nil, &ClaimIssue{ClaimRef: symbolicClaimRef(claimRef), Class: "check-params-invalid"}
		}
		if seen[claimRef] {
			return nil, &ClaimIssue{ClaimRef: claimRef, Class: "duplicate-claim-ref"}
		}
		seen[claimRef] = true
		entry, ok := r.entries[row["check_id"]]
		if !ok {
			return nil, &ClaimIssue{ClaimRef: claimRef, Class: "unknown-check"}
		}
		var params map[string]string
		if err := json.Unmarshal([]byte(row["params"]), &params); err != nil {
			return nil, &ClaimIssue{ClaimRef: claimRef, Class: "check-params-invalid"}
		}
		if entry.ID == "read-file" && strings.HasPrefix(params["expect"], "schema_ref:") && r.env.SchemaRefs[strings.TrimPrefix(params["expect"], "schema_ref:")] == "" {
			return nil, &ClaimIssue{ClaimRef: claimRef, Class: "schema-ref-unknown"}
		}
		canonical, err := json.Marshal(params)
		if err != nil || !bytes.Equal(canonical, []byte(row["params"])) || !r.validParams(entry, params) {
			return nil, &ClaimIssue{ClaimRef: claimRef, Class: "check-params-invalid"}
		}
		selections = append(selections, Selection{CheckID: row["check_id"], ClaimRef: claimRef, Params: params})
	}
	return selections, nil
}

func (r *Registry) EvaluateClaims(raw string, candidate Candidate) PredicateResult {
	selections, issue := r.ValidateClaims(raw)
	if issue != nil {
		return PredicateResult{ID: issue.ClaimRef, Predicate: Fail, FailureClass: issue.Class}
	}
	if len(selections) == 0 {
		return PredicateResult{ID: "observe-unavailable", Predicate: Blocked}
	}
	result := PredicateResult{ID: "executable-claims", Predicate: Pass, claimBindings: map[string]claimBinding{}}
	var firstFail, firstMachinery, firstNoVantage string
	for _, selection := range selections {
		row := r.Evaluator(selection)(candidate)
		result.Verdicts = append(result.Verdicts, row.Verdicts...)
		for claimRef, binding := range row.claimBindings {
			result.claimBindings[claimRef] = binding
		}
		switch {
		case row.Predicate == Fail && firstFail == "":
			firstFail = row.ID
		case row.MachineryFault && firstMachinery == "":
			firstMachinery = row.ID
		case row.Predicate == Blocked || row.Predicate == Degraded:
			if firstNoVantage == "" {
				firstNoVantage = row.ID
			}
		}
	}
	switch {
	case firstFail != "":
		result.ID, result.Predicate = firstFail, Fail
	case firstMachinery != "":
		result.ID, result.Predicate, result.MachineryFault = firstMachinery, Blocked, true
	case firstNoVantage != "":
		result.ID, result.Predicate = firstNoVantage, Blocked
	}
	return result
}

// EvaluateCandidateClaims preserves the binding distinction between an absent
// declaration and a present value. Only key absence enters the claimless floor;
// every present value follows closed declaration validation.
func (r *Registry) EvaluateCandidateClaims(candidate Candidate) PredicateResult {
	raw, present := candidate.Record.Headers["executable_claims"]
	if !present {
		return r.evaluateAbsenceFloor(candidate)
	}
	return r.EvaluateClaims(raw, candidate)
}

func validClaimRef(value string) bool {
	if value == "" || len(value) > 128 {
		return false
	}
	for _, char := range value {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || strings.ContainsRune("._:-", char) {
			continue
		}
		return false
	}
	return true
}

func symbolicClaimRef(value string) string {
	if validClaimRef(value) {
		return value
	}
	return "executable_claims"
}

func (r *Registry) validParams(entry CheckEntry, params map[string]string) bool {
	if len(params) != len(entry.ParamSchema) {
		return false
	}
	for key, value := range params {
		if _, ok := entry.ParamSchema[key]; !ok || value == "" {
			return false
		}
	}
	switch entry.ID {
	case "read-file":
		if !validRelativePath(params["path"]) || r.env.Lanes[params["lane_ref"]] == "" {
			return false
		}
		expect := params["expect"]
		switch {
		case strings.HasPrefix(expect, "line:") && len(expect) > len("line:"),
			strings.HasPrefix(expect, "hash:") && len(expect) > len("hash:"),
			strings.HasPrefix(expect, "schema_ref:") && r.env.SchemaRefs[strings.TrimPrefix(expect, "schema_ref:")] != "":
			return true
		default:
			return false
		}
	case "git-status":
		return r.env.Lanes[params["lane_ref"]] != "" && (params["expect"] == "clean" || params["expect"] == "dirty")
	case "find-references":
		return r.env.Lanes[params["lane_ref"]] != "" && params["expect"] == "count:0" && validFindRefSymbol(params["symbol"])
	case "run-suite", "run-suite-unbounded":
		return r.env.NamedSuites[params["target"]] && (params["expect_green"] == "true" || params["expect_green"] == "false")
	}
	return false
}

func validFindRefSymbol(value string) bool {
	if len(value) == 0 || len(value) > 128 || value[0] != '_' && (value[0] < 'a' || value[0] > 'z') && (value[0] < 'A' || value[0] > 'Z') {
		return false
	}
	for i := 1; i < len(value); i++ {
		if !findRefIdentifierByte(value[i]) {
			return false
		}
	}
	return true
}

func validRelativePath(path string) bool {
	if path == "" || filepath.IsAbs(path) {
		return false
	}
	clean := filepath.Clean(path)
	return path == clean && clean != "." && clean != ".." && !strings.HasPrefix(clean, ".."+string(filepath.Separator))
}

func refusedVerdict(selection Selection) CheckVerdict {
	return refusedVerdictWithDetail(selection, "check-params-refused")
}

func refusedVerdictWithDetail(selection Selection, detail string) CheckVerdict {
	return CheckVerdict{
		CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
		Outcome: "unsafe", Predicate: Blocked, FailingDetail: detail,
	}
}

func cloneBoolMap(in map[string]bool) map[string]bool {
	out := make(map[string]bool, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}
