package fixtures_test

import (
	"crypto/sha256"
	"encoding/hex"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/observe"
	"github.com/The-Frank-Organization/frank/internal/record"
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

	for _, tc := range []struct {
		selection observe.Selection
		detail    string
	}{
		{selection: observe.Selection{CheckID: "read-file", ClaimRef: "abs", Params: map[string]string{"lane_ref": "lane-a", "path": filepath.Join(string(filepath.Separator), "secret"), "expect": "line:x"}}, detail: "check-params-refused"},
		{selection: observe.Selection{CheckID: "read-file", ClaimRef: "traversal", Params: map[string]string{"lane_ref": "lane-a", "path": "../secret", "expect": "line:x"}}, detail: "check-params-refused"},
		{selection: observe.Selection{CheckID: "read-file", ClaimRef: "symlink-escape", Params: map[string]string{"lane_ref": "lane-a", "path": "escape", "expect": "line:x"}}, detail: "not-regular-file"},
		{selection: observe.Selection{CheckID: "read-file", ClaimRef: "broken-symlink-escape", Params: map[string]string{"lane_ref": "lane-a", "path": "broken-escape", "expect": "line:x"}}, detail: "not-regular-file"},
		{selection: observe.Selection{CheckID: "run-suite", ClaimRef: "command", Params: map[string]string{"target": "go test ./...; leak", "expect_green": "true"}}, detail: "check-params-refused"},
	} {
		verdict := reg.Run(tc.selection)
		if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != tc.detail {
			t.Fatalf("selection %#v was not refused before execution: %#v", tc.selection, verdict)
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

func TestS8ReadFileRefusesNonRegularFilesWithoutReading(t *testing.T) {
	lane, err := os.MkdirTemp("/tmp", "s8-nonregular-")
	if err != nil {
		t.Fatalf("short non-regular lane: %v", err)
	}
	defer os.RemoveAll(lane)
	if err := syscall.Mkfifo(filepath.Join(lane, "blocking.fifo"), 0o600); err != nil {
		t.Fatalf("mkfifo fixture: %v", err)
	}
	if err := os.Symlink("blocking.fifo", filepath.Join(lane, "fifo-link")); err != nil {
		t.Fatalf("symlink-to-fifo fixture: %v", err)
	}
	socketPath := filepath.Join(lane, "blocking.sock")
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		t.Fatalf("listen unix fixture: %v", err)
	}
	defer listener.Close()

	reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"lane-a": lane}})
	for _, path := range []string{"blocking.fifo", "blocking.sock", "fifo-link"} {
		t.Run(path, func(t *testing.T) {
			verdict := reg.Run(s8ReadSelection("lane-a", path, "line:x"))
			if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "not-regular-file" {
				t.Fatalf("non-regular verdict = %#v", verdict)
			}
			s8AssertSymbolicDetail(t, verdict.FailingDetail)
		})
	}
}

func TestS8ReadFileDurableByteCeiling(t *testing.T) {
	const ceiling = 8 << 20
	lane := t.TempDir()
	if err := os.WriteFile(filepath.Join(lane, "exact.bin"), make([]byte, ceiling), 0o600); err != nil {
		t.Fatalf("write exact ceiling fixture: %v", err)
	}
	if err := os.WriteFile(filepath.Join(lane, "over.bin"), make([]byte, ceiling+1), 0o600); err != nil {
		t.Fatalf("write over ceiling fixture: %v", err)
	}

	reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"lane-a": lane}})
	exactSum := sha256.Sum256(make([]byte, ceiling))
	if verdict := reg.Run(s8ReadSelection("lane-a", "exact.bin", "hash:"+hex.EncodeToString(exactSum[:]))); verdict.Outcome != "pass" {
		t.Fatalf("exact ceiling verdict = %#v", verdict)
	}
	if verdict := reg.Run(s8ReadSelection("lane-a", "over.bin", "line:x")); verdict.Outcome != "unsafe" || verdict.FailingDetail != "read-size-exceeded" {
		t.Fatalf("ceiling+1 verdict = %#v", verdict)
	}

	growing := filepath.Join(lane, "growing.bin")
	if err := os.WriteFile(growing, []byte("start"), 0o600); err != nil {
		t.Fatalf("write growth fixture: %v", err)
	}
	readReached := make(chan struct{})
	releaseRead := make(chan struct{})
	var blocked atomic.Bool
	reg = observe.NewRegistry(observe.RegistryEnv{
		Lanes: map[string]string{"lane-a": lane},
		ReadFileStageHook: func(stage observe.ReadFileStage) {
			if stage == observe.ReadFileStageRead && blocked.CompareAndSwap(false, true) {
				close(readReached)
				<-releaseRead
			}
		},
	})
	done := make(chan observe.CheckVerdict, 1)
	go func() { done <- reg.Run(s8ReadSelection("lane-a", "growing.bin", "line:start")) }()
	<-readReached
	file, err := os.OpenFile(growing, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		t.Fatalf("open growth fixture: %v", err)
	}
	if _, err := file.Write(make([]byte, ceiling)); err != nil {
		file.Close()
		t.Fatalf("grow fixture: %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("close growth fixture: %v", err)
	}
	close(releaseRead)
	if verdict := <-done; verdict.Outcome != "unsafe" || verdict.FailingDetail != "read-size-exceeded" {
		t.Fatalf("growth-beyond verdict = %#v", verdict)
	}
}

func TestS8ReadFileConfinementUsesOpenedDescriptors(t *testing.T) {
	lane := t.TempDir()
	outside := t.TempDir()
	insideDir := filepath.Join(lane, "component")
	if err := os.Mkdir(insideDir, 0o700); err != nil {
		t.Fatalf("mkdir inside component: %v", err)
	}
	if err := os.WriteFile(filepath.Join(insideDir, "artifact"), []byte("inside\n"), 0o600); err != nil {
		t.Fatalf("write inside artifact: %v", err)
	}
	if err := os.WriteFile(filepath.Join(outside, "artifact"), []byte("outside\n"), 0o600); err != nil {
		t.Fatalf("write outside artifact: %v", err)
	}

	openReached := make(chan struct{})
	releaseOpen := make(chan struct{})
	var opens atomic.Int32
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes: map[string]string{"lane-a": lane},
		ReadFileStageHook: func(stage observe.ReadFileStage) {
			// root, intermediate component, then final file: swap only after the
			// intermediate descriptor has been opened.
			if stage == observe.ReadFileStageOpen && opens.Add(1) == 3 {
				close(openReached)
				<-releaseOpen
			}
		},
	})
	done := make(chan observe.CheckVerdict, 1)
	go func() { done <- reg.Run(s8ReadSelection("lane-a", "component/artifact", "line:inside")) }()
	<-openReached
	if err := os.Rename(insideDir, filepath.Join(lane, "component-old")); err != nil {
		t.Fatalf("rename checked component: %v", err)
	}
	if err := os.Symlink(outside, insideDir); err != nil {
		t.Fatalf("swap component toward outside: %v", err)
	}
	close(releaseOpen)
	verdict := <-done
	if verdict.Outcome != "pass" || verdict.Predicate != observe.Pass {
		t.Fatalf("component-swap verdict = %#v", verdict)
	}
}

func TestS8ReadFileRefusesInteriorDotDotBeforeWorkerLaunch(t *testing.T) {
	lane := t.TempDir()
	outside := t.TempDir()
	component := filepath.Join(lane, "component")
	if err := os.Mkdir(component, 0o700); err != nil {
		t.Fatalf("mkdir component: %v", err)
	}
	if err := os.WriteFile(filepath.Join(lane, "artifact"), []byte("inside\n"), 0o600); err != nil {
		t.Fatalf("write inside artifact: %v", err)
	}
	if err := os.WriteFile(filepath.Join(outside, "artifact"), []byte("outside\n"), 0o600); err != nil {
		t.Fatalf("write outside artifact: %v", err)
	}

	openReached := make(chan struct{})
	releaseOpen := make(chan struct{})
	var opens atomic.Int32
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes: map[string]string{"lane-a": lane},
		ReadFileStageHook: func(stage observe.ReadFileStage) {
			// Root and component are open before the third open stage attempts
			// component/.. . The vulnerable walk can then follow the moved
			// descriptor's new parent to the outside directory.
			if stage == observe.ReadFileStageOpen && opens.Add(1) == 3 {
				close(openReached)
				<-releaseOpen
			}
		},
	})
	done := make(chan observe.CheckVerdict, 1)
	go func() {
		done <- reg.Run(s8ReadSelection("lane-a", "component/../artifact", "line:outside"))
	}()

	var verdict observe.CheckVerdict
	select {
	case verdict = <-done:
		// Correct validation refuses before the worker/open seam is reached.
	case <-openReached:
		if err := os.Rename(component, filepath.Join(outside, "moved")); err != nil {
			t.Fatalf("move opened component outside: %v", err)
		}
		close(releaseOpen)
		verdict = <-done
	case <-time.After(time.Second):
		t.Fatal("interior-dotdot check neither refused nor reached the vulnerable walk")
	}

	if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "check-params-refused" {
		t.Fatalf("interior-dotdot verdict = %#v", verdict)
	}
	if got := opens.Load(); got != 0 {
		t.Fatalf("interior-dotdot launched filesystem worker: open stages = %d", got)
	}
}

func TestS8ReadFileDetachesEveryFilesystemBlockPoint(t *testing.T) {
	for _, stage := range []observe.ReadFileStage{
		observe.ReadFileStageBefore,
		observe.ReadFileStageOpen,
		observe.ReadFileStageMetadata,
		observe.ReadFileStageRead,
	} {
		t.Run(string(stage), func(t *testing.T) {
			lane := t.TempDir()
			if stage == observe.ReadFileStageBefore {
				// A missing root still reaches the worker seam and times out: the
				// serialized caller did not probe the filesystem before launch.
				lane = filepath.Join(lane, "missing-root")
			} else {
				if err := os.WriteFile(filepath.Join(lane, "artifact"), []byte("ok\n"), 0o600); err != nil {
					t.Fatalf("write block fixture: %v", err)
				}
			}
			reached := make(chan struct{})
			release := make(chan struct{})
			var blocked atomic.Bool
			reg := observe.NewRegistry(observe.RegistryEnv{
				Lanes: map[string]string{"lane-a": lane}, ReadTimeout: 25 * time.Millisecond,
				ReadFileStageHook: func(got observe.ReadFileStage) {
					if got == stage && blocked.CompareAndSwap(false, true) {
						close(reached)
						<-release
					}
				},
			})
			done := make(chan observe.CheckVerdict, 1)
			go func() { done <- reg.Run(s8ReadSelection("lane-a", "artifact", "line:ok")) }()
			<-reached
			select {
			case verdict := <-done:
				if verdict.Outcome != "skipped" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "check-machinery-read-file-timeout" || verdict.Timing != "timeout" {
					t.Fatalf("blocked %s verdict = %#v", stage, verdict)
				}
			case <-time.After(time.Second):
				t.Fatalf("control path did not return for blocked %s operation", stage)
			}
			close(release)
		})
	}
}

func TestS8ReadFileBreakerBoundsWorkersPerLaneAndTakesFaultEdge(t *testing.T) {
	laneA := t.TempDir()
	laneB := t.TempDir()
	for _, lane := range []string{laneA, laneB} {
		if err := os.WriteFile(filepath.Join(lane, "artifact"), []byte("ok\n"), 0o600); err != nil {
			t.Fatalf("write breaker fixture: %v", err)
		}
	}
	firstStarted := make(chan struct{})
	releaseFirst := make(chan struct{})
	var launches atomic.Int32
	reg := observe.NewRegistry(observe.RegistryEnv{
		Lanes: map[string]string{"lane-a": laneA, "lane-b": laneB}, ReadTimeout: 25 * time.Millisecond,
		ReadFileStageHook: func(stage observe.ReadFileStage) {
			if stage != observe.ReadFileStageBefore {
				return
			}
			if launches.Add(1) == 1 {
				close(firstStarted)
				<-releaseFirst
			}
		},
	})
	selectionA := s8ReadSelection("lane-a", "artifact", "line:ok")
	done := make(chan observe.CheckVerdict, 1)
	go func() { done <- reg.Run(selectionA) }()
	<-firstStarted
	first := <-done
	if first.FailingDetail != "check-machinery-read-file-timeout" {
		t.Fatalf("first timeout verdict = %#v", first)
	}
	second := reg.Run(selectionA)
	if second.FailingDetail != "check-machinery-read-file-breaker-open" || second.Predicate != observe.Blocked {
		t.Fatalf("same-lane breaker verdict = %#v", second)
	}
	if got := launches.Load(); got != 1 {
		t.Fatalf("same-lane breaker launched worker: launches = %d", got)
	}
	if other := reg.Run(s8ReadSelection("lane-b", "artifact", "line:ok")); other.Outcome != "pass" {
		t.Fatalf("different-lane verdict = %#v", other)
	}
	if got := launches.Load(); got != 2 {
		t.Fatalf("different lane did not launch independently: launches = %d", got)
	}

	for _, tc := range []struct {
		name           string
		authorityClass string
		wantTerminal   string
		wantEscalate   bool
	}{
		{name: "authority held", authorityClass: "yes", wantTerminal: record.Held, wantEscalate: true},
		{name: "non-authority rejected", authorityClass: "no", wantTerminal: record.Rejected},
	} {
		t.Run(tc.name, func(t *testing.T) {
			cand := record.Record{Headers: map[string]string{"authority_class": tc.authorityClass, "EVIDENCE_TARGET": "E1"}}
			result, terminal := observe.Gate(cand, "seat-a", "SITREP", "report-only", observe.Env{
				PresentLayers: map[string]bool{"observe": true}, Evaluate: reg.Evaluator(selectionA),
			})
			if terminal != tc.wantTerminal || result.FailureClass != "observe-machinery-fault" || result.FailingPredicate != selectionA.ClaimRef || result.Escalate != tc.wantEscalate {
				t.Fatalf("terminal = %q, result = %#v", terminal, result)
			}
		})
	}
	close(releaseFirst)
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
			detail:  "not-regular-file",
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

func s8ReadSelection(lane, path, expect string) observe.Selection {
	return observe.Selection{CheckID: "read-file", ClaimRef: "read", Params: map[string]string{
		"lane_ref": lane, "path": path, "expect": expect,
	}}
}

func s8AssertSymbolicDetail(t *testing.T, detail string) {
	t.Helper()
	if strings.Contains(detail, string(filepath.Separator)) || strings.Contains(detail, ".") {
		t.Fatalf("detail is not symbolic/path-free: %q", detail)
	}
}
