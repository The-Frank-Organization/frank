package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestMintLockHeldRootRefusesWithoutBindingWrite(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	bindingPath := filepath.Join(root, "binding", "seats.json")
	before := readOptionalFile(t, bindingPath)
	held, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot: %v", err)
	}
	defer held.Release()

	result := runMintProcess(t, root, "held-seat", "")
	if result.err == nil || !strings.Contains(string(result.output), "root-lock-held") {
		t.Fatalf("mint under held root err=%v output=%q, want root-lock-held", result.err, result.output)
	}
	if after := readOptionalFile(t, bindingPath); !bytes.Equal(after, before) {
		t.Fatalf("binding changed after held-root refusal\nbefore=%s\nafter=%s", before, after)
	}
}

func TestMintLockPrecedesLiveSocketDiagnostic(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	socket := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mint-decoy-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(socket) })
	listener, err := net.Listen("unix", socket)
	if err != nil {
		t.Fatalf("listen decoy socket: %v", err)
	}
	defer listener.Close()
	held, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot: %v", err)
	}
	defer held.Release()

	result := runMintProcess(t, root, "precedence-seat", socket)
	text := string(result.output)
	if result.err == nil || !strings.Contains(text, "root-lock-held") || strings.Contains(text, "conductor is serving") {
		t.Fatalf("mint precedence err=%v output=%q, want root-lock-held before socket diagnostic", result.err, result.output)
	}
}

func TestMintLockStaleSocketRemainsDiagnosticOnly(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	socket := filepath.Join(t.TempDir(), "stale.sock")
	if err := os.WriteFile(socket, []byte("stale"), 0o600); err != nil {
		t.Fatalf("write stale socket marker: %v", err)
	}

	result := runMintProcess(t, root, "stale-seat", socket)
	if result.err != nil || !strings.Contains(string(result.output), "credential=") {
		t.Fatalf("mint with stale socket err=%v output=%q, want credential", result.err, result.output)
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	if got := mgr.CredentialsFor("stale-seat"); got != 1 {
		t.Fatalf("stale-seat credentials=%d, want 1", got)
	}
}

func TestMintLockConcurrentMintHasOneTypedLoser(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	writeLargeBindingFixture(t, root, 40_000)
	results := raceMintProcesses(t, root, "race-seat")
	successes := 0
	for _, result := range results {
		if result.err == nil {
			successes++
			continue
		}
		if !strings.Contains(string(result.output), "root-lock-held") {
			t.Fatalf("mint loser err=%v output=%q, want root-lock-held", result.err, result.output)
		}
	}
	if successes != 1 {
		t.Fatalf("mint successes=%d results=%+v, want exactly one", successes, results)
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	if got := mgr.CredentialsFor("race-seat"); got != 1 {
		t.Fatalf("race-seat credentials=%d, want exactly one", got)
	}
}

type mintProcessResult struct {
	output []byte
	err    error
}

func runMintProcess(t *testing.T, root, target, socket string) mintProcessResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	return runMintBinary(ctx, buildFrank(t, ctx), root, target, socket)
}

func runMintBinary(ctx context.Context, bin, root, target, socket string) mintProcessResult {
	args := []string{"-root", root, "-mint", target, "-role", "implementer"}
	if socket != "" {
		args = append(args, "-socket", socket)
	}
	output, err := exec.CommandContext(ctx, bin, args...).CombinedOutput()
	return mintProcessResult{output: output, err: err}
}

func raceMintProcesses(t *testing.T, root, target string) []mintProcessResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	start := make(chan struct{})
	results := make([]mintProcessResult, 2)
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start
			results[i] = runMintBinary(ctx, bin, root, target, "")
		}(i)
	}
	close(start)
	wg.Wait()
	return results
}

func writeLargeBindingFixture(t *testing.T, root string, count int) {
	t.Helper()
	type binding struct {
		Credential string        `json:"credential"`
		Meta       seat.SeatMeta `json:"meta"`
	}
	table := struct {
		Seats map[string]binding `json:"seats"`
	}{Seats: make(map[string]binding, count)}
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("existing-%05d", i)
		table.Seats[name] = binding{
			Credential: fmt.Sprintf("credential-%05d", i),
			Meta:       seat.SeatMeta{Name: name, Role: "implementer"},
		}
	}
	data, err := json.Marshal(table)
	if err != nil {
		t.Fatalf("marshal binding fixture: %v", err)
	}
	dir := filepath.Join(root, "binding")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatalf("create binding dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "seats.json"), data, 0o600); err != nil {
		t.Fatalf("write binding fixture: %v", err)
	}
}
