package fixtures_test

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
)

func TestS8CheckRegistryDescriptorsAndIPHValidation(t *testing.T) {
	lane := t.TempDir()
	outside := t.TempDir()
	if err := os.WriteFile(filepath.Join(outside, "secret"), []byte("x\n"), 0o644); err != nil {
		t.Fatalf("write outside fixture: %v", err)
	}
	if err := os.Symlink(filepath.Join(outside, "secret"), filepath.Join(lane, "escape")); err != nil {
		t.Fatalf("symlink outside fixture: %v", err)
	}
	if err := os.Symlink(filepath.Join(outside, "missing"), filepath.Join(lane, "broken-escape")); err != nil {
		t.Fatalf("broken symlink outside fixture: %v", err)
	}
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes:       map[string]string{"lane-a": lane},
		NamedSuites: map[string]bool{"all": true},
	})
	readFile, ok := reg.Entry("read-file")
	if !ok || readFile.Rung != "E1" || readFile.Class != "base" || readFile.ExecutorRequired || readFile.TimeoutClass != "read_short" {
		t.Fatalf("read-file descriptor = %#v, ok = %v", readFile, ok)
	}
	gitStatus, ok := reg.Entry("git-status")
	if !ok || gitStatus.Rung != "E1" || gitStatus.Class != "base" || gitStatus.ExecutorRequired {
		t.Fatalf("git-status descriptor = %#v, ok = %v", gitStatus, ok)
	}
	runSuite, ok := reg.Entry("run-suite")
	if !ok || runSuite.Rung != "E2" || runSuite.Class != "suite" || !runSuite.ExecutorRequired {
		t.Fatalf("run-suite descriptor = %#v, ok = %v", runSuite, ok)
	}

	for _, selection := range []observe.Selection{
		{CheckID: "read-file", ClaimRef: "abs", Params: map[string]string{"lane_ref": "lane-a", "path": filepath.Join(string(filepath.Separator), "secret"), "expect": "line:x"}},
		{CheckID: "read-file", ClaimRef: "traversal", Params: map[string]string{"lane_ref": "lane-a", "path": "../secret", "expect": "line:x"}},
		{CheckID: "read-file", ClaimRef: "symlink-escape", Params: map[string]string{"lane_ref": "lane-a", "path": "escape", "expect": "line:x"}},
		{CheckID: "read-file", ClaimRef: "broken-symlink-escape", Params: map[string]string{"lane_ref": "lane-a", "path": "broken-escape", "expect": "line:x"}},
		{CheckID: "run-suite", ClaimRef: "command", Params: map[string]string{"target": "go test ./...; leak", "expect_green": "true"}},
	} {
		verdict := reg.Run(selection)
		if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "check-params-refused" {
			t.Fatalf("selection %#v was not refused before execution: %#v", selection, verdict)
		}
		if strings.Contains(verdict.FailingDetail, "secret") || strings.Contains(verdict.FailingDetail, string(filepath.Separator)) {
			t.Fatalf("refusal leaked a path: %#v", verdict)
		}
	}
}

func TestS8ReadFileChecksLineHashAndSchemaRef(t *testing.T) {
	lane := t.TempDir()
	data := []byte("alpha\nbeta\n")
	if err := os.WriteFile(filepath.Join(lane, "artifact.txt"), data, 0o644); err != nil {
		t.Fatalf("write artifact: %v", err)
	}
	sum := sha256.Sum256(data)
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes:      map[string]string{"lane-a": lane},
		SchemaRefs: map[string]string{"artifact-schema": hex.EncodeToString(sum[:])},
	})
	for _, tc := range []struct {
		name   string
		expect string
	}{
		{name: "line", expect: "line:beta"},
		{name: "hash", expect: "hash:" + hex.EncodeToString(sum[:])},
		{name: "schema-ref", expect: "schema_ref:artifact-schema"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			verdict := reg.Run(observe.Selection{
				CheckID: "read-file", ClaimRef: tc.name,
				Params: map[string]string{"lane_ref": "lane-a", "path": "artifact.txt", "expect": tc.expect},
			})
			if verdict.Outcome != "pass" || verdict.Predicate != observe.Pass || verdict.RungReached != "E1" {
				t.Fatalf("verdict = %#v", verdict)
			}
		})
	}
}

func TestS8ReadFileRefusesFIFOWithoutBlocking(t *testing.T) {
	lane := t.TempDir()
	if err := syscall.Mkfifo(filepath.Join(lane, "blocking.fifo"), 0o600); err != nil {
		t.Fatalf("mkfifo fixture: %v", err)
	}
	reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"lane-a": lane}})
	done := make(chan observe.CheckVerdict, 1)
	go func() {
		done <- reg.Run(observe.Selection{CheckID: "read-file", ClaimRef: "fifo", Params: map[string]string{
			"lane_ref": "lane-a", "path": "blocking.fifo", "expect": "line:x",
		}})
	}()

	select {
	case verdict := <-done:
		if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "read-file-not-regular" {
			t.Fatalf("FIFO verdict = %#v", verdict)
		}
		if strings.Contains(verdict.FailingDetail, "blocking.fifo") || strings.Contains(verdict.FailingDetail, string(filepath.Separator)) {
			t.Fatalf("FIFO refusal leaked a path: %#v", verdict)
		}
	case <-time.After(250 * time.Millisecond):
		t.Fatal("read-file blocked on a FIFO")
	}
}

func TestS8ReadFileEnforcesDeadlineAndByteCeiling(t *testing.T) {
	lane := t.TempDir()
	if err := os.WriteFile(filepath.Join(lane, "small.txt"), []byte("alpha\n"), 0o644); err != nil {
		t.Fatalf("write deadline fixture: %v", err)
	}
	if err := os.WriteFile(filepath.Join(lane, "oversize.bin"), make([]byte, 2<<20), 0o644); err != nil {
		t.Fatalf("write byte-ceiling fixture: %v", err)
	}

	t.Run("deadline actually fires", func(t *testing.T) {
		reg := observe.NewRegistry(observe.RegistryEnv{
			Lanes: map[string]string{"lane-a": lane}, ReadTimeout: time.Nanosecond,
		})
		started := time.Now()
		verdict := reg.Run(observe.Selection{CheckID: "read-file", ClaimRef: "deadline", Params: map[string]string{
			"lane_ref": "lane-a", "path": "small.txt", "expect": "line:alpha",
		}})
		if verdict.Outcome != "skipped" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "check-machinery-read-file-timeout" || verdict.Timing != "timeout" {
			t.Fatalf("deadline verdict = %#v", verdict)
		}
		if elapsed := time.Since(started); elapsed > time.Second {
			t.Fatalf("deadline took %v", elapsed)
		}
	})

	t.Run("byte ceiling refuses oversized regular file", func(t *testing.T) {
		reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"lane-a": lane}})
		verdict := reg.Run(observe.Selection{CheckID: "read-file", ClaimRef: "oversize", Params: map[string]string{
			"lane_ref": "lane-a", "path": "oversize.bin", "expect": "line:x",
		}})
		if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "read-file-byte-ceiling" {
			t.Fatalf("byte-ceiling verdict = %#v", verdict)
		}
		if strings.Contains(verdict.FailingDetail, "oversize.bin") || strings.Contains(verdict.FailingDetail, string(filepath.Separator)) {
			t.Fatalf("byte-ceiling refusal leaked a path: %#v", verdict)
		}
	})
}

func TestS8GitStatusObservesCleanDirtyAndVetoesFalseCleanClaim(t *testing.T) {
	lane := t.TempDir()
	s8Git(t, lane, "init", "-q")
	s8Git(t, lane, "config", "user.email", "fixture@example.invalid")
	s8Git(t, lane, "config", "user.name", "Fixture")
	if err := os.WriteFile(filepath.Join(lane, "tracked.txt"), []byte("clean\n"), 0o644); err != nil {
		t.Fatalf("write tracked: %v", err)
	}
	s8Git(t, lane, "add", "tracked.txt")
	s8Git(t, lane, "commit", "-qm", "fixture")
	reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"lane-a": lane}})
	selection := observe.Selection{CheckID: "git-status", ClaimRef: "clean-tree", Params: map[string]string{"lane_ref": "lane-a", "expect": "clean"}}
	if verdict := reg.Run(selection); verdict.Outcome != "pass" || verdict.Predicate != observe.Pass {
		t.Fatalf("clean verdict = %#v", verdict)
	}
	if err := os.WriteFile(filepath.Join(lane, "tracked.txt"), []byte("dirty\n"), 0o644); err != nil {
		t.Fatalf("dirty tracked: %v", err)
	}
	if verdict := reg.Run(selection); verdict.Outcome != "fail" || verdict.Predicate != observe.Fail || verdict.FailingDetail != "git-status-mismatch" {
		t.Fatalf("dirty verdict = %#v", verdict)
	}

	cand := record.Record{Headers: map[string]string{"FINAL_GIT_STATUS_SHORT": "none - clean tree", "EVIDENCE_TARGET": "E1"}}
	result, terminal := observe.Gate(cand, "seat-a", "SITREP", "report-only", observe.Env{
		PresentLayers: map[string]bool{"observe": true},
		Evaluate:      reg.Evaluator(selection),
	})
	if terminal != record.Rejected || result.FailingPredicate != "clean-tree" {
		t.Fatalf("false clean claim terminal = %q, result = %#v", terminal, result)
	}
}

func TestS8E1DistinguishesFaultRefusalAndObservedAbsence(t *testing.T) {
	lane := t.TempDir()
	if err := os.Mkdir(filepath.Join(lane, "not-a-file"), 0o755); err != nil {
		t.Fatalf("mkdir read-file machinery fixture: %v", err)
	}
	git := s8TimeoutGit(t)
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes:         map[string]string{"lane-a": lane},
		GitExecutable: git,
		ReadTimeout:   20 * time.Millisecond,
	})

	for _, tc := range []struct {
		name      string
		selection observe.Selection
		outcome   string
		detail    string
	}{
		{
			name: "git status timeout is machinery",
			selection: observe.Selection{CheckID: "git-status", ClaimRef: "git-timeout", Params: map[string]string{
				"lane_ref": "lane-a", "expect": "clean",
			}},
			outcome: "skipped",
			detail:  "check-machinery-git-status-timeout",
		},
		{
			name: "non-regular read file is refused",
			selection: observe.Selection{CheckID: "read-file", ClaimRef: "read-error", Params: map[string]string{
				"lane_ref": "lane-a", "path": "not-a-file", "expect": "line:x",
			}},
			outcome: "unsafe",
			detail:  "read-file-not-regular",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			verdict := reg.Run(tc.selection)
			if verdict.Outcome != tc.outcome || verdict.Predicate != observe.Blocked || verdict.FailingDetail != tc.detail {
				t.Fatalf("verdict = %#v", verdict)
			}
		})
	}

	missing := reg.Run(observe.Selection{CheckID: "read-file", ClaimRef: "missing", Params: map[string]string{
		"lane_ref": "lane-a", "path": "missing.txt", "expect": "line:x",
	}})
	if missing.Outcome != "fail" || missing.Predicate != observe.Fail || missing.FailingDetail != "read-file-absent" {
		t.Fatalf("missing-file verdict = %#v", missing)
	}
}

func s8TimeoutGit(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "git")
	if err := os.WriteFile(path, []byte("#!/bin/sh\nsleep 1\n"), 0o755); err != nil {
		t.Fatalf("write timeout git: %v", err)
	}
	return path
}

func s8Git(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
}
