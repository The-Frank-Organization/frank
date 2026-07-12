package observe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os/exec"
	"path/filepath"
	"strings"
)

func (r *Registry) runReadFile(selection Selection) CheckVerdict {
	// validParams requires the supplied path to equal this canonical form. Pass
	// only the canonical component sequence to the descriptor-relative worker.
	relative := filepath.Clean(selection.Params["path"])
	result := r.executeReadFile(selection.Params["lane_ref"], relative)
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
