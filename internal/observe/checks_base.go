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
	"syscall"
	"time"
)

const (
	readCheckTimeout = 5 * time.Second
	readByteCeiling  = 1 << 20
)

func (r *Registry) runReadFile(selection Selection) CheckVerdict {
	root, full, ok := r.scopedPath(selection.Params["lane_ref"], selection.Params["path"])
	if !ok {
		return refusedVerdict(selection)
	}
	_ = root
	timeout := r.env.ReadTimeout
	if timeout > readCheckTimeout {
		timeout = readCheckTimeout
	}
	data, detail, timing, err := readRegularFile(full, timeout)
	if err != nil {
		if detail == "read-file-not-regular" || detail == "read-file-byte-ceiling" {
			return refusedVerdictWithDetail(selection, detail)
		}
		if os.IsNotExist(err) {
			return baseVerdict(selection, false, "read-file-absent")
		}
		return machineryVerdict(selection, detail, timing)
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

func readRegularFile(path string, timeout time.Duration) ([]byte, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	file, err := os.OpenFile(path, os.O_RDONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, "check-machinery-read-file", "", err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return nil, "check-machinery-read-file", "", err
	}
	if !info.Mode().IsRegular() {
		return nil, "read-file-not-regular", "", syscall.EINVAL
	}
	if info.Size() > readByteCeiling {
		return nil, "read-file-byte-ceiling", "", syscall.EFBIG
	}
	if err := ctx.Err(); err != nil {
		return nil, "check-machinery-read-file-timeout", "timeout", err
	}

	data := make([]byte, 0, readByteCeiling+1)
	chunk := make([]byte, 32<<10)
	for {
		if err := ctx.Err(); err != nil {
			return nil, "check-machinery-read-file-timeout", "timeout", err
		}
		remaining := readByteCeiling + 1 - len(data)
		if remaining < len(chunk) {
			chunk = chunk[:remaining]
		}
		n, readErr := file.Read(chunk)
		data = append(data, chunk[:n]...)
		if len(data) > readByteCeiling {
			return nil, "read-file-byte-ceiling", "", syscall.EFBIG
		}
		if err := ctx.Err(); err != nil {
			return nil, "check-machinery-read-file-timeout", "timeout", err
		}
		if readErr == io.EOF {
			return data, "", "", nil
		}
		if readErr != nil {
			return nil, "check-machinery-read-file", "", readErr
		}
	}
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
