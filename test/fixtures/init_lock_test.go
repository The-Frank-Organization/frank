package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	frankconfig "github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/store"
)

func TestInitLockRefusesHeldRootBeforeAnyInitWrite(t *testing.T) {
	root := filepath.Join(t.TempDir(), "root")
	held, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot: %v", err)
	}
	defer held.Release()

	result := runInitProcess(t, root, writeFixtureConfigSources(t))
	if result.err == nil || !strings.Contains(string(result.output), "root-lock-held") {
		t.Fatalf("init under held root err=%v output=%q, want root-lock-held", result.err, result.output)
	}
	assertNoInitWrites(t, root)
}

func TestInitLockConcurrentDoubleInitKeepsOneCompleteWinner(t *testing.T) {
	root := filepath.Join(t.TempDir(), "root")
	sources := widenedInitSources(t, writeFixtureConfigSources(t), 0)
	results := raceInitProcesses(t, root, sources, sources)
	assertOneInitWinner(t, root, results, []map[string]string{sources})
}

func TestInitLockDifferingConfigRaceKeepsWinnerInternallyConsistent(t *testing.T) {
	root := filepath.Join(t.TempDir(), "root")
	first := widenedInitSources(t, writeFixtureConfigSources(t), 11)
	second := widenedInitSources(t, writeFixtureConfigSources(t), 29)
	results := raceInitProcesses(t, root, first, second)
	assertOneInitWinner(t, root, results, []map[string]string{first, second})
}

func TestInitLockSymlinkAliasRefusesHeldRoot(t *testing.T) {
	realRoot := filepath.Join(t.TempDir(), "root")
	held, err := store.AcquireRoot(realRoot)
	if err != nil {
		t.Fatalf("AcquireRoot: %v", err)
	}
	defer held.Release()
	aliasRoot := filepath.Join(t.TempDir(), "alias")
	if err := os.Symlink(realRoot, aliasRoot); err != nil {
		t.Fatalf("Symlink: %v", err)
	}

	result := runInitProcess(t, aliasRoot, writeFixtureConfigSources(t))
	if result.err == nil || !strings.Contains(string(result.output), "root-lock-held") {
		t.Fatalf("alias init under held root err=%v output=%q, want root-lock-held", result.err, result.output)
	}
	assertNoInitWrites(t, realRoot)
}

type initProcessResult struct {
	output []byte
	err    error
}

func runInitProcess(t *testing.T, root string, sources map[string]string) initProcessResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, buildFrank(t, ctx),
		"-root", root,
		"-registry", sources["fieldspec"],
		"-engine-config", sources["engine"],
		"-catalog", sources["catalog"],
		"-init",
	)
	output, err := cmd.CombinedOutput()
	return initProcessResult{output: output, err: err}
}

func raceInitProcesses(t *testing.T, root string, first, second map[string]string) []initProcessResult {
	t.Helper()
	start := make(chan struct{})
	results := make([]initProcessResult, 2)
	sources := []map[string]string{first, second}
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start
			results[i] = runInitProcess(t, root, sources[i])
		}(i)
	}
	close(start)
	wg.Wait()
	return results
}

func assertOneInitWinner(t *testing.T, root string, results []initProcessResult, candidates []map[string]string) {
	t.Helper()
	successes := 0
	for _, result := range results {
		if result.err == nil {
			successes++
			continue
		}
		text := string(result.output)
		if !strings.Contains(text, "root-lock-held") && !strings.Contains(text, store.ErrGenesisExists.Error()) {
			t.Fatalf("init loser err=%v output=%q, want root-lock-held or post-release genesis-exists", result.err, result.output)
		}
	}
	if successes != 1 {
		t.Fatalf("init successes=%d results=%+v, want exactly one", successes, results)
	}

	pinned, err := frankconfig.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load winning config: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open winning store: %v", err)
	}
	genesis, err := st.Genesis()
	if err != nil {
		t.Fatalf("read winning genesis: %v", err)
	}
	if got := genesis.Headers["config_digest"]; got != pinned.Digest {
		t.Fatalf("genesis config_digest=%q, on-disk digest=%q", got, pinned.Digest)
	}

	for _, candidate := range candidates {
		want, err := frankconfig.Load(candidate)
		if err != nil {
			t.Fatalf("load candidate config: %v", err)
		}
		if pinned.Digest == want.Digest && equalConfigMembers(pinned.Members, want.Members) {
			return
		}
	}
	t.Fatalf("winning config digest %s is not one complete candidate", pinned.Digest)
}

func widenedInitSources(t *testing.T, sources map[string]string, rotateDelta int64) map[string]string {
	t.Helper()
	engine, err := os.ReadFile(sources["engine"])
	if err != nil {
		t.Fatalf("read engine source: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(engine, &raw); err != nil {
		t.Fatalf("decode engine source: %v", err)
	}
	raw["segment_rotate_bytes"] = int64(4_194_304) + rotateDelta
	engine, err = json.Marshal(raw)
	if err != nil {
		t.Fatalf("encode engine source: %v", err)
	}
	engine = append(engine, bytes.Repeat([]byte{' '}, 4<<20)...)
	if err := os.WriteFile(sources["engine"], engine, 0o644); err != nil {
		t.Fatalf("write widened engine source: %v", err)
	}
	return sources
}

func equalConfigMembers(left, right map[string][]byte) bool {
	if len(left) != len(right) {
		return false
	}
	for name, value := range left {
		if !bytes.Equal(value, right[name]) {
			return false
		}
	}
	return true
}

func assertNoInitWrites(t *testing.T, root string) {
	t.Helper()
	for _, name := range []string{"config", "records", "indexes", "redo"} {
		if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
			t.Fatalf("init wrote %s while root lock was held: %v", name, err)
		}
	}
}
