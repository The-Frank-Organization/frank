package observe

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestFindReferencesCompleteScanCountZeroPasses(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRef(t, root, "a.txt", "alpha beta\n")
	verdict := runFindRef(t, root, "missing_token")
	if verdict.Outcome != "pass" || verdict.Predicate != Pass || verdict.RungReached != "E1" {
		t.Fatalf("verdict = %#v", verdict)
	}
}

func TestFindReferencesCountNonZeroFails(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRef(t, root, "a.txt", "alpha target beta\n")
	verdict := runFindRef(t, root, "target")
	if verdict.Outcome != "fail" || verdict.Predicate != Fail || verdict.FailingDetail != "find-references-count-nonzero" {
		t.Fatalf("verdict = %#v", verdict)
	}
}

func TestFindReferencesTokenBoundary(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRef(t, root, "tokens.txt", "foobar foo.bar foo foo_bar (foo)\n")
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": root}})
	plain := reg.executeFindReferences(findRefSelection("foo"))
	if plain.count != 2 {
		t.Fatalf("foo count = %d, want standalone and parenthesized only", plain.count)
	}
	dotted := reg.executeFindReferences(findRefSelection("foo.bar"))
	if dotted.count != 1 {
		t.Fatalf("foo.bar count = %d, want exact dotted token", dotted.count)
	}
}

func TestFindReferencesBinaryOutOfDomain(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRefBytes(t, root, "binary.bin", append([]byte{'x', 0}, []byte(" target")...))
	verdict := runFindRef(t, root, "target")
	if verdict.Outcome != "pass" {
		t.Fatalf("binary exclusion verdict = %#v", verdict)
	}
}

func TestFindReferencesIncompleteScanNeverPasses(t *testing.T) {
	t.Run("undecodable", func(t *testing.T) {
		root := t.TempDir()
		mustWriteFindRefBytes(t, root, "bad.txt", []byte{0xff, 'x'})
		verdict := runFindRef(t, root, "target")
		if verdict.Outcome != "skipped" || verdict.Predicate != Degraded || verdict.FailingDetail != "scan-domain-incomplete-undecodable" {
			t.Fatalf("verdict = %#v", verdict)
		}
	})

	tests := []struct {
		name   string
		limits findRefLimits
		build  func(*testing.T, string)
		detail string
	}{
		{name: "depth", limits: findRefLimits{depth: 0, files: 10, perFileBytes: 10, totalBytes: 20, matches: 10}, build: func(t *testing.T, root string) {
			if err := os.Mkdir(filepath.Join(root, "nested"), 0o700); err != nil {
				t.Fatal(err)
			}
			mustWriteFindRef(t, root, "nested/a.txt", "x")
		}, detail: "check-machinery-find-references-depth-ceiling"},
		{name: "file", limits: findRefLimits{depth: 2, files: 1, perFileBytes: 10, totalBytes: 20, matches: 10}, build: func(t *testing.T, root string) {
			mustWriteFindRef(t, root, "a.txt", "x")
			mustWriteFindRef(t, root, "b.txt", "x")
		}, detail: "check-machinery-find-references-file-ceiling"},
		{name: "per-file", limits: findRefLimits{depth: 2, files: 2, perFileBytes: 1, totalBytes: 20, matches: 10}, build: func(t *testing.T, root string) {
			mustWriteFindRef(t, root, "a.txt", "xx")
		}, detail: "check-machinery-find-references-per-file-ceiling"},
		{name: "total-byte", limits: findRefLimits{depth: 2, files: 2, perFileBytes: 10, totalBytes: 1, matches: 10}, build: func(t *testing.T, root string) {
			mustWriteFindRef(t, root, "a.txt", "xx")
		}, detail: "check-machinery-find-references-total-byte-ceiling"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			tc.build(t, root)
			reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": root}, findRefLimits: &tc.limits})
			verdict := reg.Run(findRefSelection("target"))
			if verdict.Outcome == "pass" || verdict.Predicate != Blocked || verdict.FailingDetail != tc.detail {
				t.Fatalf("verdict = %#v", verdict)
			}
		})
	}
}

func TestFindReferencesCeilingsAllowExactlyLimit(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRef(t, root, "a.txt", "xx")
	limits := findRefLimits{depth: 0, files: 1, perFileBytes: 2, totalBytes: 2, matches: 1}
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": root}, findRefLimits: &limits})
	verdict := reg.Run(findRefSelection("target"))
	if verdict.Outcome != "pass" {
		t.Fatalf("exact ceilings verdict = %#v", verdict)
	}
}

func TestFindReferencesSymlinkNotFollowed(t *testing.T) {
	root := t.TempDir()
	outside := t.TempDir()
	mustWriteFindRef(t, outside, "secret.txt", "target")
	if err := os.Symlink(filepath.Join(outside, "secret.txt"), filepath.Join(root, "link.txt")); err != nil {
		t.Fatal(err)
	}
	if verdict := runFindRef(t, root, "target"); verdict.Outcome != "pass" {
		t.Fatalf("symlink verdict = %#v", verdict)
	}
}

func TestFindReferencesCountSaturation(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRef(t, root, "many.txt", "target target target")
	limits := findRefLimits{depth: 1, files: 2, perFileBytes: 100, totalBytes: 100, matches: 2}
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": root}, findRefLimits: &limits})
	result := reg.executeFindReferences(findRefSelection("target"))
	if result.count != 2 || !result.saturated {
		t.Fatalf("result = %#v, want count saturated at 2", result)
	}
}

func TestFindReferencesTimeoutAndBreaker(t *testing.T) {
	root := t.TempDir()
	mustWriteFindRef(t, root, "a.txt", "target")
	blocked := make(chan struct{})
	reg := NewRegistry(RegistryEnv{
		Lanes: map[string]string{"repo": root}, ReadTimeout: 10 * time.Millisecond,
		FSStageHook: func(stage FSStage) {
			if stage == FSStageDirectoryRead {
				<-blocked
			}
		},
	})
	first := reg.Run(findRefSelection("target"))
	if first.FailingDetail != "check-machinery-find-references-timeout" || first.Predicate != Blocked {
		t.Fatalf("timeout verdict = %#v", first)
	}
	second := reg.Run(findRefSelection("target"))
	if second.FailingDetail != "check-machinery-find-references-breaker-open" || second.Predicate != Blocked {
		t.Fatalf("breaker verdict = %#v", second)
	}
	close(blocked)
}

func TestFindReferencesIPH(t *testing.T) {
	root := t.TempDir()
	name := "secret-filename.txt"
	mustWriteFindRefBytes(t, root, name, []byte{0xff})
	verdict := runFindRef(t, root, "target")
	raw, err := json.Marshal(verdict)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), root) || strings.Contains(string(raw), name) {
		t.Fatalf("verdict leaks path: %s", raw)
	}
}

func TestFindReferencesUngovernedLaneRefused(t *testing.T) {
	reg := NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": t.TempDir()}})
	selection := findRefSelection("target")
	selection.Params["lane_ref"] = "absent"
	verdict := reg.Run(selection)
	if verdict.Outcome != "unsafe" || verdict.Predicate != Blocked || verdict.FailingDetail != "lane-ungoverned" {
		t.Fatalf("verdict = %#v", verdict)
	}
}

func runFindRef(t *testing.T, root, symbol string) CheckVerdict {
	t.Helper()
	return NewRegistry(RegistryEnv{Lanes: map[string]string{"repo": root}}).Run(findRefSelection(symbol))
}

func findRefSelection(symbol string) Selection {
	return Selection{CheckID: "find-references", ClaimRef: "refs", Params: map[string]string{
		"symbol": symbol, "lane_ref": "repo", "expect": "count:0",
	}}
}

func mustWriteFindRef(t *testing.T, root, relative, value string) {
	t.Helper()
	mustWriteFindRefBytes(t, root, relative, []byte(value))
}

func mustWriteFindRefBytes(t *testing.T, root, relative string, value []byte) {
	t.Helper()
	path := filepath.Join(root, relative)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, value, 0o600); err != nil {
		t.Fatal(err)
	}
}
