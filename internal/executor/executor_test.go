package executor

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/observe"
)

func TestSpawnPassesHostModuleCacheWithOfflineReadonlyFlags(t *testing.T) {
	source := t.TempDir()
	copyRootModuleClosure(t, source)
	writeExecutorTestFile(t, source, "probe/offline_test.go", `package probe

import (
	"testing"
)

func TestOffline(t *testing.T) {
	t.Log("root closure resolved offline")
}
`, 0o600)
	hostCache, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatalf("resolve host module cache: %v", err)
	}
	t.Setenv("GOMODCACHE", hostCache)
	writeExecutorTestFile(t, source, "run.sh", `#!/bin/sh
set -eu
[ "$GOPROXY" = "off" ]
[ "$GOSUMDB" = "off" ]
[ "$GOFLAGS" = "-mod=readonly" ]
[ "$GOTOOLCHAIN" = "local" ]
[ "$GOWORK" = "off" ]
[ "$GOMODCACHE" = "$1" ]
[ "$GOCACHE" = "$PWD/.cache/go-build" ]
[ "$GOPATH" = "$PWD/.cache/gopath" ]
exec go test -mod=readonly ./probe
`, 0o700)
	host := New(Config{
		TempRoot: t.TempDir(),
		Suites: map[string]Suite{
			"offline": {SourceDir: source, Command: "run.sh", Args: []string{hostCache}, TimeoutClass: "suite_bounded", Timeout: 30 * time.Second},
		},
	})
	verdict := host.Spawn(executorTestEntry(), observe.Selection{
		CheckID: "run-suite", ClaimRef: "offline", Params: map[string]string{"target": "offline", "expect_green": "true"},
	})
	if verdict.Outcome != "pass" || verdict.Predicate != observe.Pass || verdict.RungReached != "E2" {
		t.Fatalf("offline verdict = %#v", verdict)
	}
}

func TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork(t *testing.T) {
	source := t.TempDir()
	writeExecutorTestFile(t, source, "go.mod", `module example.invalid/frank-offline-probe

go 1.25.0

require example.invalid/frank-missing-module v0.0.0
`, 0o600)
	writeExecutorTestFile(t, source, "go.sum", `example.invalid/frank-missing-module v0.0.0 h1:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=
`, 0o600)
	writeExecutorTestFile(t, source, "probe/offline_test.go", `package probe

import (
	"testing"

	_ "example.invalid/frank-missing-module/probe"
)

func TestOffline(t *testing.T) {
	t.Log("unreachable while the synthetic module is absent")
}
`, 0o600)
	writeExecutorTestFile(t, source, "run.sh", `#!/bin/sh
set -eu
exec go test -mod=readonly ./probe
`, 0o700)
	t.Setenv("GOMODCACHE", t.TempDir())
	tempRoot := t.TempDir()
	host := New(Config{
		TempRoot: tempRoot,
		Suites: map[string]Suite{
			"missing": {SourceDir: source, Command: "run.sh", TimeoutClass: "suite_bounded", Timeout: 30 * time.Second},
		},
	})
	started := time.Now()
	verdict := host.Spawn(executorTestEntry(), observe.Selection{
		CheckID: "run-suite", ClaimRef: "missing", Params: map[string]string{"target": "missing", "expect_green": "true"},
	})
	if verdict.Outcome != "fail" || verdict.Predicate != observe.Fail || verdict.RungReached != "none" || verdict.FailingDetail != "suite-exit-mismatch" {
		t.Fatalf("missing-module verdict = %#v", verdict)
	}
	if elapsed := time.Since(started); elapsed > 2*time.Second {
		t.Fatalf("offline module miss took %s, want immediate local refusal", elapsed)
	}
	if diagnostic := executorTestDiagnostic(t, tempRoot); !strings.Contains(diagnostic, "example.invalid/frank-missing-module") {
		t.Fatalf("missing-module diagnostic = %q, want module name", diagnostic)
	}
}

func TestGoModuleCachePathHonorsExplicitRunCache(t *testing.T) {
	want := t.TempDir()
	t.Setenv("GOMODCACHE", want)
	want, err := filepath.EvalSymlinks(want)
	if err != nil {
		t.Fatalf("resolve expected GOMODCACHE: %v", err)
	}
	got, err := goModuleCachePath()
	if err != nil {
		t.Fatalf("resolve explicit GOMODCACHE: %v", err)
	}
	if got != want {
		t.Fatalf("resolved GOMODCACHE = %q, want %q", got, want)
	}
}

func TestSpawnNestedExecutorPackageFromZeroExternalClosure(t *testing.T) {
	source := t.TempDir()
	writeExecutorTestFile(t, source, "go.mod", `module github.com/The-Frank-Organization/frank

go 1.25.0
`, 0o600)
	copyExecutorPackageFixture(t, source)
	writeExecutorTestFile(t, source, "run.sh", `#!/bin/sh
set -eu
exec go test -mod=readonly ./internal/executor -run '^(TestSpawnPassesHostModuleCacheWithOfflineReadonlyFlags|TestSpawnNamesMissingOfflineModuleWithoutWaitingForNetwork|TestGoModuleCachePathHonorsExplicitRunCache|TestSpawnRetainsPrivateCappedTailOutsideBareFailureVerdict)$' -count=1
	`, 0o700)
	t.Setenv("GOMODCACHE", t.TempDir())
	tempRoot := t.TempDir()
	host := New(Config{
		TempRoot: tempRoot,
		Suites: map[string]Suite{
			"nested": {SourceDir: source, Command: "run.sh", TimeoutClass: "suite_bounded", Timeout: 30 * time.Second},
		},
	})
	verdict := host.Spawn(executorTestEntry(), observe.Selection{
		CheckID: "run-suite", ClaimRef: "zero-external-closure", Params: map[string]string{"target": "nested", "expect_green": "true"},
	})
	if verdict.Outcome != "pass" || verdict.Predicate != observe.Pass || verdict.RungReached != "E2" {
		t.Fatalf("nested zero-external-closure verdict = %#v; diagnostic = %q", verdict, executorTestDiagnostic(t, tempRoot))
	}
}

func TestSpawnRetainsPrivateCappedTailOutsideBareFailureVerdict(t *testing.T) {
	source := t.TempDir()
	writeExecutorTestFile(t, source, "fail.sh", `#!/bin/sh
set -eu
printf 'discard-this-prefix-abcdefghijklmnopqrstuvwxyz'
printf 'retained-tail-marker\n'
exit 7
`, 0o700)
	tempRoot := t.TempDir()
	host := New(Config{
		TempRoot:    tempRoot,
		OutputLimit: 24,
		Suites: map[string]Suite{
			"fail": {SourceDir: source, Command: "fail.sh", TimeoutClass: "suite_bounded", Timeout: 2 * time.Second},
		},
	})
	selection := observe.Selection{
		CheckID: "run-suite", ClaimRef: "retained", Params: map[string]string{"target": "fail", "expect_green": "true"},
	}
	verdict := host.Spawn(executorTestEntry(), selection)
	if verdict.Outcome != "fail" || verdict.Predicate != observe.Fail || verdict.RungReached != "none" || verdict.FailingDetail != "suite-exit-mismatch" {
		t.Fatalf("failure verdict = %#v", verdict)
	}

	entries, err := os.ReadDir(tempRoot)
	if err != nil {
		t.Fatalf("read temp root: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("retained entries = %v, want one diagnostic and no workdir", entries)
	}
	entry := entries[0]
	if entry.IsDir() || !regexp.MustCompile(`^frank-executor-diagnostic-[0-9a-f]{64}$`).MatchString(entry.Name()) {
		t.Fatalf("diagnostic entry = %q, want run-key-named regular file", entry.Name())
	}
	info, err := entry.Info()
	if err != nil {
		t.Fatalf("diagnostic info: %v", err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Fatalf("diagnostic permissions = %o, want 600", got)
	}
	data, err := os.ReadFile(filepath.Join(tempRoot, entry.Name()))
	if err != nil {
		t.Fatalf("read retained diagnostic: %v", err)
	}
	if len(data) > 24 || !strings.Contains(string(data), "retained-tail-marker") || strings.Contains(string(data), "discard-this-prefix") {
		t.Fatalf("retained diagnostic = %q (%d bytes), want bounded last bytes", data, len(data))
	}

	if replay := host.Spawn(executorTestEntry(), selection); replay != verdict {
		t.Fatalf("replayed verdict = %#v, want %#v", replay, verdict)
	}
	entries, err = os.ReadDir(tempRoot)
	if err != nil {
		t.Fatalf("read temp root after replay: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("replay multiplied diagnostic artifacts: %v", entries)
	}
}

func executorTestEntry() observe.CheckEntry {
	return observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "suite_bounded"}
}

func writeExecutorTestFile(t *testing.T, dir, name, body string, mode os.FileMode) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("create parent for %s: %v", name, err)
	}
	if err := os.WriteFile(path, []byte(body), mode); err != nil {
		t.Fatalf("write %s: %v", name, err)
	}
}

func copyExecutorPackageFixture(t *testing.T, target string) {
	t.Helper()
	root := executorRepositoryRoot(t)
	for _, packagePath := range []string{"internal"} {
		source := filepath.Join(root, packagePath)
		if err := filepath.Walk(source, func(path string, info os.FileInfo, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if info.IsDir() || filepath.Ext(path) != ".go" {
				return nil
			}
			body, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			relative, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			writeExecutorTestFile(t, target, relative, string(body), 0o600)
			return nil
		}); err != nil {
			t.Fatalf("copy %s fixture: %v", packagePath, err)
		}
	}
}

func copyRootModuleClosure(t *testing.T, target string) {
	t.Helper()
	root := executorRepositoryRoot(t)
	for _, name := range []string{"go.mod", "go.sum"} {
		body, err := os.ReadFile(filepath.Join(root, name))
		if errors.Is(err, os.ErrNotExist) && name == "go.sum" {
			continue
		}
		if err != nil {
			t.Fatalf("read root %s: %v", name, err)
		}
		writeExecutorTestFile(t, target, name, string(body), 0o600)
	}
}

func executorRepositoryRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("resolve fixture source root: %v", err)
	}
	return root
}

func executorTestDiagnostic(t *testing.T, tempRoot string) string {
	t.Helper()
	entries, err := os.ReadDir(tempRoot)
	if err != nil {
		t.Fatalf("read executor temp root: %v", err)
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), diagnosticPrefix) {
			body, err := os.ReadFile(filepath.Join(tempRoot, entry.Name()))
			if err != nil {
				t.Fatalf("read executor diagnostic: %v", err)
			}
			return string(body)
		}
	}
	return ""
}
