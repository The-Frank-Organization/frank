package observe

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"path/filepath"
	"strings"
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
	Params          map[string]string
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
	Lanes         map[string]string
	SchemaRefs    map[string]string
	NamedSuites   map[string]bool
	Executor      SuiteExecutor
	GitExecutable string
	ReadTimeout   time.Duration
}

type Registry struct {
	entries map[string]CheckEntry
	env     RegistryEnv
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
	return &Registry{
		env: RegistryEnv{
			Lanes:         cloneMap(env.Lanes),
			SchemaRefs:    cloneMap(env.SchemaRefs),
			NamedSuites:   cloneBoolMap(env.NamedSuites),
			Executor:      env.Executor,
			GitExecutable: gitExecutable,
			ReadTimeout:   readTimeout,
		},
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
			"run-suite": {
				ID: "run-suite", Rung: "E2", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded",
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
	entry, ok := r.entries[selection.CheckID]
	if !ok || !r.validParams(entry, selection.Params) {
		return refusedVerdict(selection)
	}
	switch selection.CheckID {
	case "read-file":
		return r.runReadFile(selection)
	case "git-status":
		return r.runGitStatus(selection)
	case "run-suite":
		if r.env.Executor != nil {
			return r.env.Executor.Spawn(entry, selection)
		}
		fallthrough
	default:
		return CheckVerdict{
			CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
			Outcome: "unsafe", Predicate: Blocked, FailingDetail: "executor-required",
		}
	}
}

func (r *Registry) Evaluator(selection Selection) func(Candidate) PredicateResult {
	return func(candidate Candidate) PredicateResult {
		bound := selection
		if data, err := json.Marshal(candidate.Record); err == nil {
			sum := sha256.Sum256(data)
			bound.CandidateDigest = hex.EncodeToString(sum[:])
		}
		verdict := r.Run(bound)
		id := verdict.ClaimRef
		if id == "" {
			id = verdict.CheckID
		}
		return PredicateResult{
			ID: id, Predicate: verdict.Predicate, MachineryFault: machineryFaultDetail(verdict.FailingDetail),
			Verdicts: []CheckVerdict{verdict},
		}
	}
}

func machineryFaultDetail(detail string) bool {
	return strings.HasPrefix(detail, "executor-") || strings.HasPrefix(detail, "check-machinery-")
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
	case "run-suite":
		return r.env.NamedSuites[params["target"]] && (params["expect_green"] == "true" || params["expect_green"] == "false")
	}
	return false
}

func validRelativePath(path string) bool {
	if path == "" || filepath.IsAbs(path) {
		return false
	}
	clean := filepath.Clean(path)
	return clean != "." && clean != ".." && !strings.HasPrefix(clean, ".."+string(filepath.Separator))
}

func refusedVerdict(selection Selection) CheckVerdict {
	return CheckVerdict{
		CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
		Outcome: "unsafe", Predicate: Blocked, FailingDetail: "check-params-refused",
	}
}

func cloneBoolMap(in map[string]bool) map[string]bool {
	out := make(map[string]bool, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}
