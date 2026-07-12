package fixtures_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jackli/frank/internal/executor"
	"github.com/jackli/frank/internal/observe"
)

func TestS8FXEXE1ExecutorProvidesOnlyRunScopedHandles(t *testing.T) {
	source := t.TempDir()
	s8WriteExecutable(t, source, "probe.sh", `#!/bin/sh
set -eu
[ "$HOME" = "$PWD" ]
[ "$TMPDIR" = "$PWD/.tmp" ]
[ "$GOCACHE" = "$PWD/.cache/go-build" ]
[ "$GOMODCACHE" = "$PWD/.cache/go-mod" ]
[ "$GOPATH" = "$PWD/.cache/gopath" ]
for key in FRANK_ROOT FRANK_SOCKET FRANK_CREDENTIAL SIGNING_KEY CONFIG_PATH OUTBOX_PATH; do
  eval "value=\${$key-}"
  [ -z "$value" ]
done
`)
	host := executor.New(executor.Config{
		TempRoot: t.TempDir(),
		Suites: map[string]executor.Suite{
			"isolation": {SourceDir: source, Command: "probe.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
		},
	})
	verdict := host.Spawn(observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}, observe.Selection{
		CheckID: "run-suite", ClaimRef: "isolation", Params: map[string]string{"target": "isolation", "expect_green": "true"},
	})
	if verdict.Outcome != "pass" || verdict.Predicate != observe.Pass || verdict.RungReached != "E2" {
		t.Fatalf("isolation verdict = %#v", verdict)
	}
	if strings.Contains(strings.ToLower(executor.AmbientResidual), "by construction") || !strings.Contains(executor.AmbientResidual, "same-uid") {
		t.Fatalf("ambient residual wording overclaims: %q", executor.AmbientResidual)
	}
}

func TestS8FXEXE2TimeoutReapsGroupBeforeCleanup(t *testing.T) {
	source := t.TempDir()
	s8WriteExecutable(t, source, "slow.sh", "#!/bin/sh\nsleep 30 &\nwait\n")
	tempRoot := t.TempDir()
	host := executor.New(executor.Config{
		TempRoot: tempRoot,
		Suites: map[string]executor.Suite{
			"slow": {SourceDir: source, Command: "slow.sh", TimeoutClass: "suite_bounded", Timeout: 80 * time.Millisecond},
		},
	})
	started := time.Now()
	verdict := host.Spawn(observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}, observe.Selection{
		CheckID: "run-suite", ClaimRef: "timeout", Params: map[string]string{"target": "slow", "expect_green": "true"},
	})
	if verdict.Outcome != "unsafe" || verdict.Predicate != observe.Blocked || verdict.FailingDetail != "executor-timeout" {
		t.Fatalf("timeout verdict = %#v", verdict)
	}
	if time.Since(started) > 3*time.Second {
		t.Fatalf("timeout did not bound the process group")
	}
	entries, err := os.ReadDir(tempRoot)
	if err != nil {
		t.Fatalf("read temp root: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("workdir survived exit-confirmed cleanup: %v", entries)
	}

	t.Run("bounded group verification fault preserves workdir", func(t *testing.T) {
		source := t.TempDir()
		s8WriteExecutable(t, source, "quick.sh", "#!/bin/sh\nexit 0\n")
		tempRoot := t.TempDir()
		host := executor.New(executor.Config{
			TempRoot: tempRoot, GroupVerifyBound: 30 * time.Millisecond,
			GroupGone: func(int) bool { return false },
			Suites: map[string]executor.Suite{
				"quick": {SourceDir: source, Command: "quick.sh", TimeoutClass: "suite_bounded", Timeout: time.Second},
			},
		})
		verdict := host.Spawn(observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}, observe.Selection{
			CheckID: "run-suite", ClaimRef: "survivor", Params: map[string]string{"target": "quick", "expect_green": "true"},
		})
		if verdict.Outcome != "unsafe" || verdict.FailingDetail != "executor-survivor" {
			t.Fatalf("survivor verdict = %#v", verdict)
		}
		entries, err := os.ReadDir(tempRoot)
		if err != nil {
			t.Fatalf("read preserved temp root: %v", err)
		}
		if len(entries) != 1 {
			t.Fatalf("survivor workdir not preserved: %v", entries)
		}
	})
}

func TestS8FXEXE3CoalescesAndReplaysOneManifestExecution(t *testing.T) {
	source := t.TempDir()
	counter := filepath.Join(t.TempDir(), "runs")
	s8WriteExecutable(t, source, "count.sh", fmt.Sprintf("#!/bin/sh\nprintf 'run\\n' >> %q\n", counter))
	host := executor.New(executor.Config{
		TempRoot: t.TempDir(),
		Suites: map[string]executor.Suite{
			"count": {SourceDir: source, Command: "count.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
		},
	})
	entry := observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}
	selection := observe.Selection{CheckID: "run-suite", ClaimRef: "coalesce", Params: map[string]string{"target": "count", "expect_green": "true"}}
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if verdict := host.Spawn(entry, selection); verdict.Outcome != "pass" {
				t.Errorf("concurrent verdict = %#v", verdict)
			}
		}()
	}
	wg.Wait()
	if verdict := host.Spawn(entry, selection); verdict.Outcome != "pass" {
		t.Fatalf("replay verdict = %#v", verdict)
	}
	data, err := os.ReadFile(counter)
	if err != nil {
		t.Fatalf("read counter: %v", err)
	}
	if got := strings.Count(string(data), "run\n"); got != 1 {
		t.Fatalf("manifest executed %d times, want 1", got)
	}
	selection.CandidateDigest = "different-candidate"
	if verdict := host.Spawn(entry, selection); verdict.Outcome != "pass" {
		t.Fatalf("different candidate verdict = %#v", verdict)
	}
	data, err = os.ReadFile(counter)
	if err != nil {
		t.Fatalf("read counter after distinct candidate: %v", err)
	}
	if got := strings.Count(string(data), "run\n"); got != 2 {
		t.Fatalf("distinct candidate manifest executed total %d times, want 2", got)
	}
}

func TestS8FXEXE4And6TruncationClosedVerdictAndSideEffectRefusal(t *testing.T) {
	source := t.TempDir()
	s8WriteExecutable(t, source, "loud.sh", "#!/bin/sh\ni=0; while [ $i -lt 200 ]; do printf 'abcdefghij'; i=$((i+1)); done\n")
	host := executor.New(executor.Config{
		TempRoot: t.TempDir(), OutputLimit: 64,
		Suites: map[string]executor.Suite{
			"loud": {SourceDir: source, Command: "loud.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
		},
	})
	selection := observe.Selection{CheckID: "run-suite", ClaimRef: "loud", Params: map[string]string{"target": "loud", "expect_green": "true"}}
	verdict := host.Spawn(observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}, selection)
	if verdict.Outcome != "pass" || verdict.FailingDetail != "output-truncated" {
		t.Fatalf("truncation verdict = %#v", verdict)
	}
	data, err := json.Marshal(verdict)
	if err != nil {
		t.Fatalf("marshal verdict: %v", err)
	}
	var shape map[string]any
	if err := json.Unmarshal(data, &shape); err != nil {
		t.Fatalf("unmarshal verdict: %v", err)
	}
	want := []string{"check_id", "claim_ref", "outcome", "rung_reached", "predicate", "timing", "failing_detail"}
	if len(shape) != len(want) {
		t.Fatalf("verdict boundary fields = %#v", shape)
	}
	for _, key := range want {
		if _, ok := shape[key]; !ok {
			t.Fatalf("verdict missing %s: %#v", key, shape)
		}
	}
	refused := host.Spawn(observe.CheckEntry{ID: "danger", Class: "side_effecting", ExecutorRequired: true}, selection)
	if refused.Outcome != "unsafe" || refused.FailingDetail != "executor-class-refused" {
		t.Fatalf("side-effecting verdict = %#v", refused)
	}
}

func TestS8FXEXE5RegistryRoutesNoVerdictFaultToOuterGate(t *testing.T) {
	source := t.TempDir()
	s8WriteExecutable(t, source, "fail.sh", "#!/bin/sh\nexit 2\n")
	host := executor.New(executor.Config{
		TempRoot: t.TempDir(),
		Suites: map[string]executor.Suite{
			"fail": {SourceDir: source, Command: "fail.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
		},
	})
	reg := observe.NewRegistry(observe.RegistryEnv{
		NamedSuites: map[string]bool{"fail": true}, Executor: host,
	})
	verdict := reg.Run(observe.Selection{CheckID: "run-suite", ClaimRef: "suite-green", Params: map[string]string{"target": "fail", "expect_green": "true"}})
	if verdict.Outcome != "fail" || verdict.Predicate != observe.Fail || verdict.FailingDetail != "suite-exit-mismatch" {
		t.Fatalf("suite failure verdict = %#v", verdict)
	}
}

func s8WriteExecutable(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o755); err != nil {
		t.Fatalf("write executable: %v", err)
	}
}
