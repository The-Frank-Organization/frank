package observe

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const readCheckTimeout = 5 * time.Second

func (r *Registry) runReadFile(selection Selection) CheckVerdict {
	root, full, ok := r.scopedPath(selection.Params["lane_ref"], selection.Params["path"])
	if !ok {
		return refusedVerdict(selection)
	}
	_ = root
	data, err := os.ReadFile(full)
	if err != nil {
		if os.IsNotExist(err) {
			return baseVerdict(selection, false, "read-file-absent")
		}
		return machineryVerdict(selection, "check-machinery-read-file", "")
	}
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

func (r *Registry) scopedPath(laneRef, relative string) (string, string, bool) {
	root := r.env.Lanes[laneRef]
	if root == "" || !validRelativePath(relative) {
		return "", "", false
	}
	resolvedRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return "", "", false
	}
	full := filepath.Join(resolvedRoot, filepath.Clean(relative))
	resolvedFull, err := filepath.EvalSymlinks(full)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", "", false
		}
		if _, statErr := os.Lstat(full); statErr == nil || !os.IsNotExist(statErr) {
			return "", "", false
		}
		resolvedParent, parentErr := filepath.EvalSymlinks(filepath.Dir(full))
		if parentErr != nil {
			return "", "", false
		}
		resolvedFull = filepath.Join(resolvedParent, filepath.Base(full))
	}
	rel, err := filepath.Rel(resolvedRoot, resolvedFull)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", "", false
	}
	return resolvedRoot, resolvedFull, true
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
