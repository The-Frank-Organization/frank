package observe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var claimlessReportCeilings = map[string]bool{
	"report-only": true,
	"read-only":   true,
	"design-only": true,
	"review-only": true,
	"plan-only":   true,
}

var claimlessReportPhases = map[string]bool{
	"AUDIT":         true,
	"DESIGN":        true,
	"DESIGN-REVIEW": true,
	"RECONCILE":     true,
	"PLAN":          true,
	"PLAN-REVIEW":   true,
	"SITREP":        true,
}

var claimlessDeferredCeilings = map[string]bool{
	"implementation": true,
	"fold-in-only":   true,
	"merge-gated":    true,
	"live-verify":    true,
}

var claimlessDeferredPhases = map[string]bool{
	"IMPL":        true,
	"REVIEW-FOLD": true,
	"MERGE-GATE":  true,
	"LIVE-VERIFY": true,
}

func (r *Registry) evaluateAbsenceFloor(candidate Candidate) PredicateResult {
	if claimlessDeferredCeilings[candidate.Authority] || claimlessDeferredPhases[candidate.Phase] {
		return claimlessResult("phase-done", Fail, "unsafe", "phase-predicate-deferred")
	}
	if !claimlessReportCeilings[candidate.Authority] || !claimlessReportPhases[candidate.Phase] {
		return claimlessResult("phase-done", Fail, "unsafe", "phase-predicate-unsupported")
	}

	root := r.env.Lanes["repo"]
	resolved, err := filepath.EvalSymlinks(root)
	if root == "" || err != nil {
		return claimlessResult("phase-done", Degraded, "skipped", "observation-unavailable")
	}
	dir, err := os.Open(resolved)
	if err != nil {
		return claimlessResult("phase-done", Degraded, "skipped", "observation-unavailable")
	}
	info, statErr := dir.Stat()
	_, readErr := dir.Readdirnames(1)
	closeErr := dir.Close()
	if statErr != nil || (readErr != nil && readErr != io.EOF) || closeErr != nil || !info.IsDir() {
		return claimlessResult("phase-done", Degraded, "skipped", "observation-unavailable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), r.env.ReadTimeout)
	defer cancel()
	status, err := exec.CommandContext(ctx, r.env.GitExecutable, "-C", resolved, "status", "--porcelain=v1").Output()
	if ctx.Err() != nil {
		result := claimlessResult("phase-done", Blocked, "skipped", "check-machinery-git-status-timeout")
		result.MachineryFault = true
		return result
	}
	if err != nil {
		result := claimlessResult("phase-done", Blocked, "skipped", "check-machinery-git-status")
		result.MachineryFault = true
		return result
	}
	observed, valid := canonicalPorcelainV1(status)
	if !valid {
		return claimlessResult("FINAL_GIT_STATUS_SHORT", Fail, "fail", "git-status-porcelain-invalid")
	}
	if candidate.Record.Headers["FINAL_GIT_STATUS_SHORT"] != observed {
		return claimlessResult("FINAL_GIT_STATUS_SHORT", Fail, "fail", "observed-false")
	}
	return claimlessResult("phase-done", Degraded, "skipped", "turn-attribution-unavailable")
}

func canonicalPorcelainV1(raw []byte) (string, bool) {
	if len(raw) == 0 {
		return "none - clean tree", true
	}
	withoutTerminator := bytes.TrimSuffix(raw, []byte{'\n'})
	for _, entry := range bytes.Split(withoutTerminator, []byte{'\n'}) {
		if len(entry) == 0 {
			continue
		}
		if len(entry) < 4 || entry[2] != ' ' {
			return "", false
		}
	}
	return string(withoutTerminator), true
}

func claimlessResult(id, predicate, outcome, detail string) PredicateResult {
	return PredicateResult{
		ID: id, Predicate: predicate, FailureClass: detail,
		Verdicts: []CheckVerdict{{
			CheckID: "claimless-floor", ClaimRef: id, Outcome: outcome,
			RungReached: "none", Predicate: predicate, FailingDetail: detail,
		}},
	}
}

func (r *Registry) runReadFile(selection Selection) CheckVerdict {
	// validParams requires the supplied path to equal this canonical form. Pass
	// only the canonical component sequence to the descriptor-relative worker.
	relative := filepath.Clean(selection.Params["path"])
	result := r.executeReadFile(selection, selection.Params["lane_ref"], relative)
	switch result.kind {
	case readFileResultAbsent:
		return baseVerdict(selection, false, result.detail)
	case readFileResultRefused:
		return refusedVerdictWithDetail(selection, result.detail)
	case readFileResultMachinery:
		return machineryVerdict(selection, result.detail, result.timing)
	}
	data := result.data
	expect := selection.Params["expect"]
	matched := false
	switch {
	case strings.HasPrefix(expect, "line:"):
		needle := []byte(strings.TrimPrefix(expect, "line:"))
		for _, line := range bytes.Split(data, []byte{'\n'}) {
			if bytes.Equal(line, needle) {
				matched = true
				break
			}
		}
	case strings.HasPrefix(expect, "hash:"):
		sum := sha256.Sum256(data)
		matched = hex.EncodeToString(sum[:]) == strings.TrimPrefix(expect, "hash:")
	case strings.HasPrefix(expect, "schema_ref:"):
		sum := sha256.Sum256(data)
		matched = hex.EncodeToString(sum[:]) == r.env.SchemaRefs[strings.TrimPrefix(expect, "schema_ref:")]
	}
	return baseVerdict(selection, matched, "read-file-mismatch")
}

func (r *Registry) runGitStatus(selection Selection) CheckVerdict {
	root := r.env.Lanes[selection.Params["lane_ref"]]
	ctx, cancel := context.WithTimeout(context.Background(), r.env.ReadTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, r.env.GitExecutable, "-C", root, "status", "--porcelain", "--untracked-files=normal")
	out, err := cmd.Output()
	if ctx.Err() != nil {
		return machineryVerdict(selection, "check-machinery-git-status-timeout", "timeout")
	}
	if err != nil {
		return machineryVerdict(selection, "check-machinery-git-status", "")
	}
	isClean := len(bytes.TrimSpace(out)) == 0
	matched := selection.Params["expect"] == "clean" && isClean || selection.Params["expect"] == "dirty" && !isClean
	return baseVerdict(selection, matched, "git-status-mismatch")
}

func machineryVerdict(selection Selection, detail, timing string) CheckVerdict {
	return CheckVerdict{
		CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
		Outcome: "skipped", RungReached: "none", Predicate: Blocked,
		Timing: timing, FailingDetail: detail,
	}
}

func baseVerdict(selection Selection, pass bool, failingDetail string) CheckVerdict {
	verdict := CheckVerdict{
		CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
		Outcome: "pass", RungReached: "E1", Predicate: Pass,
	}
	if !pass {
		verdict.Outcome = "fail"
		verdict.RungReached = "none"
		verdict.Predicate = Fail
		verdict.FailingDetail = failingDetail
	}
	return verdict
}
