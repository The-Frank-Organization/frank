package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

type cachedFixtureBinary struct {
	path string
	err  error
}

var (
	frankFixtureBinary    cachedFixtureBinary
	frankMCPFixtureBinary cachedFixtureBinary
)

func TestMain(m *testing.M) {
	if fixtureChildMode() {
		buildErr := fmt.Errorf("child-mode: no cached binary")
		frankFixtureBinary.err = buildErr
		frankMCPFixtureBinary.err = buildErr
		os.Exit(m.Run())
	}

	dir, err := os.MkdirTemp("", "frank-fixtures-")
	if err != nil {
		buildErr := fmt.Errorf("create fixture binary directory: %w", err)
		frankFixtureBinary.err = buildErr
		frankMCPFixtureBinary.err = buildErr
	} else {
		frankFixtureBinary = buildCachedFixtureBinary(dir, "frank", filepath.Join("..", "..", "cmd", "frank"))
		frankMCPFixtureBinary = buildCachedFixtureBinary(dir, "frank-mcp", filepath.Join("..", "..", "cmd", "frank-mcp"))
	}

	code := m.Run()
	if dir != "" {
		if err := os.RemoveAll(dir); err != nil && code == 0 {
			fmt.Fprintf(os.Stderr, "remove fixture binary directory: %v\n", err)
			code = 1
		}
	}
	os.Exit(code)
}

func fixtureChildMode() bool {
	for _, marker := range []string{
		"FRANK_F10_CHILD",
		"FRANK_F11_CHILD",
		"FRANK_F11_TRACE_CHILD",
		"FRANK_F9_CHILD",
	} {
		if os.Getenv(marker) == "1" {
			return true
		}
	}
	return false
}

func buildCachedFixtureBinary(dir, name, target string) cachedFixtureBinary {
	path := filepath.Join(dir, name)
	build := exec.Command("go", "build", "-o", path, target)
	if out, err := build.CombinedOutput(); err != nil {
		return cachedFixtureBinary{err: fmt.Errorf("build %s: %w\n%s", name, err, out)}
	}
	return cachedFixtureBinary{path: path}
}

func TestFrankBinaryAssemblesAuthenticatedSubmitProjectRead(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-main-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	cmd := exec.CommandContext(ctx, bin, "-root", root, "-socket", sock, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start frank: %v", err)
	}
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() {
		if client != nil {
			_ = client.Close()
		}
	}()

	describe, err := client.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools stderr=%s: %v", stderr.String(), err)
	}
	if describe.FormDigest == "" || describe.SubmitSchema == nil {
		t.Fatalf("describe missing schema/digest: %+v", describe)
	}
	if describe.SubmitSchema.HasField("grant") || describe.SubmitSchema.OptionAllowed("AUTHORITY", "merge-gated") {
		t.Fatalf("pair describe exposed grant authority: %+v", describe.SubmitSchema.Fields)
	}
	payload := mustJSONBytes(t, fieldspec.SubmitPayload{Record: record.Record{
		Envelope: record.Envelope{To: "seat-a", DispatchID: "dispatch-main"},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1", "SUBJECT": "binary path"},
		Body:     "hello",
	}, FormDigest: describe.FormDigest})
	result, err := client.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stderr=%s: %v", stderr.String(), err)
	}
	var outcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode outcome %s: %v", result, err)
	}
	if outcome.State != record.Accepted || outcome.RelayID == "" {
		t.Fatalf("outcome = %+v", outcome)
	}

	project, err := client.Call(ctx, "project", nil)
	if err != nil {
		t.Fatalf("project: %v", err)
	}
	var relayIDs []string
	if err := json.Unmarshal(project, &relayIDs); err != nil {
		t.Fatalf("decode project %s: %v", project, err)
	}
	if len(relayIDs) != 1 || relayIDs[0] != outcome.RelayID {
		t.Fatalf("project = %v, want [%s]", relayIDs, outcome.RelayID)
	}

	readArgs, _ := json.Marshal(map[string]string{"relay_id": outcome.RelayID})
	readResult, err := client.Call(ctx, "read", readArgs)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	var read struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(readResult, &read); err != nil {
		t.Fatalf("decode read %s: %v", readResult, err)
	}
	if read.Record.Envelope.From != "seat-a" || read.Record.Envelope.Role != "implementer" {
		t.Fatalf("read record not stamped: %+v", read.Record.Envelope)
	}

	payloadFile := filepath.Join(t.TempDir(), "operator-submit.json")
	if err := os.WriteFile(payloadFile, payload, 0o644); err != nil {
		t.Fatalf("write operator payload: %v", err)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("close authenticated client before operator-submit: %v", err)
	}
	client = nil
	operatorOut := runOperatorSubmit(t, ctx, bin, sock, payloadFile, cred.Value)
	var operatorOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(bytes.TrimSpace(operatorOut), &operatorOutcome); err != nil {
		t.Fatalf("decode operator-submit output %s: %v", operatorOut, err)
	}
	if operatorOutcome.State != record.Accepted || operatorOutcome.RelayID == "" {
		t.Fatalf("operator outcome = %+v", operatorOutcome)
	}
}

func TestFrankBinaryOperatorChannelO3OwedSweepOpenAndDisposition(t *testing.T) {
	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	reg := loadAssemblyRegistry(t)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	reportPath := filepath.Join("..", "..", "docs", "sprints", "2026-07-03-s2-slice-2", "results", "f11-sweep-report.md")
	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read F11 sweep report: %v", err)
	}
	if !bytes.Contains(report, []byte("S2 F11 Sweep Report")) {
		t.Fatalf("unexpected F11 sweep report contents")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-o3-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = client.Close() }()
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	submit := func(rec record.Record) struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	} {
		t.Helper()
		payload := submitPayloadBytes(t, reg, pinned.Digest, meta, rec)
		result, err := client.Call(ctx, "submit", payload)
		if err != nil {
			t.Fatalf("submit stderr=%s: %v", stderr.String(), err)
		}
		var outcome struct {
			State   string `json:"state"`
			RelayID string `json:"relay_id"`
		}
		if err := json.Unmarshal(result, &outcome); err != nil {
			t.Fatalf("decode submit outcome %s: %v", result, err)
		}
		if outcome.State != record.Accepted || outcome.RelayID == "" {
			t.Fatalf("submit outcome = %+v", outcome)
		}
		return outcome
	}

	source := "docs/sprints/2026-07-03-s1-slice-1/RECONCILE.md:160-161 + .relays/s1/s1-exit-gate/SITREP-planner-20260703-200827.md"
	owed := submit(record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "oi-s1-f11-sweep"},
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"AUTHORITY":        "report-only",
			"SUBJECT":          "OI-S1-F11-SWEEP",
			"record_kind":      "owed_item",
			"owner":            "s1 (S2 slice)",
			"source":           source,
			"target_surface":   "F11 full class×point sweep",
			"disposition_path": "S2 exit gate",
		},
	})
	projectResult, err := client.Call(ctx, "project", nil)
	if err != nil {
		t.Fatalf("project: %v", err)
	}
	var relayIDs []string
	if err := json.Unmarshal(projectResult, &relayIDs); err != nil {
		t.Fatalf("decode project %s: %v", projectResult, err)
	}
	if !containsString(relayIDs, owed.RelayID) {
		t.Fatalf("project = %v, missing owed relay %s", relayIDs, owed.RelayID)
	}
	readArgs, _ := json.Marshal(map[string]string{"relay_id": owed.RelayID})
	readResult, err := client.Call(ctx, "read", readArgs)
	if err != nil {
		t.Fatalf("read owed: %v", err)
	}
	var read struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(readResult, &read); err != nil {
		t.Fatalf("decode read %s: %v", readResult, err)
	}
	if read.Record.Envelope.From != "operator" || read.Record.Envelope.Role != "operator" || read.Record.Headers["record_kind"] != "owed_item" {
		t.Fatalf("read owed record = %+v", read.Record)
	}
	for _, want := range []string{"s1 (S2 slice)", source, "F11 full class×point sweep", "S2 exit gate"} {
		if !strings.Contains(read.Record.Headers["owner"]+read.Record.Headers["source"]+read.Record.Headers["target_surface"]+read.Record.Headers["disposition_path"], want) {
			t.Fatalf("read owed headers missing %q: %+v", want, read.Record.Headers)
		}
	}
	openPath := filepath.Join(root, "projections", "owed", "OPEN.md")
	open := string(mustReadFile(t, openPath))
	for _, want := range []string{owed.RelayID, "s1 (S2 slice)", source, "F11 full class×point sweep", "S2 exit gate"} {
		if !strings.Contains(open, want) {
			t.Fatalf("OPEN.md missing %q:\n%s", want, open)
		}
	}

	disposition := submit(record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "oi-s1-f11-sweep"},
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"AUTHORITY":     "report-only",
			"SUBJECT":       "OI-S1-F11-SWEEP disposition",
			"record_kind":   "owed_disposition",
			"disposes_owed": owed.RelayID,
			"source":        reportPath,
		},
		Body: "Disposes OI-S1-F11-SWEEP after docs/sprints/2026-07-03-s2-slice-2/results/f11-sweep-report.md.",
	})
	if disposition.RelayID == owed.RelayID {
		t.Fatalf("disposition reused owed relay id %s", owed.RelayID)
	}
	open = string(mustReadFile(t, openPath))
	if strings.Contains(open, owed.RelayID) {
		t.Fatalf("OPEN.md still contains disposed owed id %s:\n%s", owed.RelayID, open)
	}
}

func TestFrankBinaryMintSeatAdminTimeCredential(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	run := func(args ...string) ([]byte, []byte, error) {
		t.Helper()
		cmd := exec.CommandContext(ctx, bin, args...)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		return stdout.Bytes(), stderr.Bytes(), err
	}

	bindingPath := filepath.Join(root, "binding", "seats.json")
	beforeSystem := readOptionalFile(t, bindingPath)
	stdout, stderr, err := run("-root", root, "-mint", "system")
	if err == nil {
		t.Fatalf("-mint system unexpectedly succeeded with stdout=%s", stdout)
	}
	if len(stdout) != 0 || !bytes.Contains(stderr, []byte(seat.ErrReservedSeatName.Error())) {
		t.Fatalf("-mint system stdout=%q stderr=%q, want typed reserved-seat error only on stderr", stdout, stderr)
	}
	if after := readOptionalFile(t, bindingPath); !bytes.Equal(after, beforeSystem) {
		t.Fatalf("binding table changed after rejected system mint\nbefore=%s\nafter=%s", beforeSystem, after)
	}

	stdout, stderr, err = run("-root", root, "-mint", "operator", "-role", "operator", "-operator")
	if err != nil {
		t.Fatalf("-mint operator failed: %v\nstdout=%s\nstderr=%s", err, stdout, stderr)
	}
	credential := parseCredentialStdout(t, stdout, stderr)
	assertCredentialNotInSeatDeliverables(t, root, credential)
	afterOperator := mustReadFile(t, bindingPath)

	stdout, stderr, err = run("-root", root, "-mint", "operator", "-role", "operator", "-operator")
	if err == nil {
		t.Fatalf("duplicate -mint unexpectedly succeeded with stdout=%s", stdout)
	}
	if len(stdout) != 0 || !bytes.Contains(stderr, []byte(seat.ErrSeatAlreadyBound.Error())) {
		t.Fatalf("duplicate -mint stdout=%q stderr=%q, want typed duplicate-seat error only on stderr", stdout, stderr)
	}
	if after := mustReadFile(t, bindingPath); !bytes.Equal(after, afterOperator) {
		t.Fatalf("binding table changed after duplicate mint\nbefore=%s\nafter=%s", afterOperator, after)
	}

	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mint-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, serverStderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)
	client, err := channel.DialAuthenticated(ctx, sock, credential)
	if err != nil {
		t.Fatalf("DialAuthenticated with minted credential stderr=%s: %v", serverStderr.String(), err)
	}
	defer func() { _ = client.Close() }()
	tools, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("ListTools: %v", err)
	}
	if got, want := strings.Join(tools, ","), "submit,project,read"; got != want {
		t.Fatalf("full-surface tools = %v, want %s", tools, want)
	}
	descriptions, err := client.ToolDescriptions(ctx)
	if err != nil {
		t.Fatalf("ToolDescriptions: %v", err)
	}
	if _, ok := descriptions["mint"]; ok || strings.Contains(strings.Join(tools, ","), "mint") {
		t.Fatalf("mint appeared in seat-facing registry: tools=%v descriptions=%v", tools, descriptions)
	}

	beforeLiveMint := mustReadFile(t, bindingPath)
	stdout, stderr, err = run("-root", root, "-socket", sock, "-mint", "late-seat", "-role", "implementer")
	if err == nil {
		t.Fatalf("live -mint unexpectedly succeeded with stdout=%s", stdout)
	}
	if len(stdout) != 0 || !bytes.Contains(stderr, []byte("conductor is serving")) {
		t.Fatalf("live -mint stdout=%q stderr=%q, want admin-time serving error only on stderr", stdout, stderr)
	}
	if after := mustReadFile(t, bindingPath); !bytes.Equal(after, beforeLiveMint) {
		t.Fatalf("binding table changed after live mint rejection\nbefore=%s\nafter=%s", beforeLiveMint, after)
	}
}

func TestFrankInitTwiceRejectsExistingGenesis(t *testing.T) {
	root := t.TempDir()
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sources := writeFixtureConfigSources(t)

	first := exec.CommandContext(ctx, bin, "-root", root, "-registry", sources["fieldspec"], "-engine-config", sources["engine"], "-init")
	if out, err := first.CombinedOutput(); err != nil {
		t.Fatalf("first init: %v\n%s", err, out)
	}
	second := exec.CommandContext(ctx, bin, "-root", root, "-registry", sources["fieldspec"], "-engine-config", sources["engine"], "-init")
	out, err := second.CombinedOutput()
	if err == nil {
		t.Fatalf("second init unexpectedly succeeded")
	}
	if !strings.Contains(string(out), "genesis") {
		t.Fatalf("second init output = %s, want genesis error", out)
	}
}

func TestFrankBinarySocketPathPreflight(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	longSock := filepath.Join(os.TempDir(), strings.Repeat("s", 120)+".sock")
	cmd := exec.CommandContext(ctx, bin, "-root", root, "-socket", longSock, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("frank unexpectedly started with long socket path")
	}
	text := string(out)
	if !strings.Contains(text, "socket path too long") || !strings.Contains(text, "darwin") {
		t.Fatalf("socket preflight output = %s, want typed darwin length error", text)
	}
	if strings.Contains(text, "bind: invalid argument") {
		t.Fatalf("socket preflight leaked raw bind error: %s", text)
	}
}

func TestFrankBinaryReissuesRecoveryWakeForExistingMailbox(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "preexisting", From: "seat-b", To: "seat-a", Role: "planner", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "preexisting delivery"},
	}, nil); err != nil {
		t.Fatalf("Commit preexisting delivery: %v", err)
	}

	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-wake-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = client.Close() }()
	frame, err := client.NextPush(ctx)
	if err != nil {
		t.Fatalf("recovery wake was not reissued stderr=%s: %v", stderr.String(), err)
	}
	if !bytes.Contains(frame, []byte("recovery")) {
		t.Fatalf("recovery wake frame = %s", frame)
	}
}

func TestFrankBinaryServesReadOnlyDiagnosticsOnDigestMismatch(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "config", "engine.json"), []byte(`{"gc_enabled":true,"segment_rotate_bytes":96}`), 0o644); err != nil {
		t.Fatalf("mutate engine config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-diag-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = client.Close() }()
	tools, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("ListTools stderr=%s: %v", stderr.String(), err)
	}
	want := []string{"project", "read"}
	if strings.Join(tools, ",") != strings.Join(want, ",") {
		t.Fatalf("tools = %v, want %v; stderr=%s", tools, want, stderr.String())
	}
	describe, err := client.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools stderr=%s: %v", stderr.String(), err)
	}
	if describe.SubmitSchema != nil || describe.FormDigest != "" {
		t.Fatalf("diagnostics describe exposed submit schema: %+v", describe)
	}
	if _, err := client.Call(ctx, "submit", json.RawMessage(`{}`)); err == nil {
		t.Fatalf("submit unexpectedly available in diagnostics mode")
	}
}

func TestFrankBinaryReadCorruptionQueuesLiveQuarantine(t *testing.T) {
	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	operatorCred, err := mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-k2-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)
	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = client.Close() }()
	reg := loadAssemblyRegistry(t)
	meta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}
	payload := submitPayloadBytes(t, reg, pinned.Digest, meta, record.Record{
		Envelope: record.Envelope{To: "seat-a", DispatchID: "dispatch-k2"},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "corrupt me"},
	})
	result, err := client.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stderr=%s: %v", stderr.String(), err)
	}
	var out struct {
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &out); err != nil {
		t.Fatalf("decode submit %s: %v", result, err)
	}
	if err := os.WriteFile(filepath.Join(root, "records", out.RelayID+".json"), []byte(`{"bad":true}`), 0o644); err != nil {
		t.Fatalf("corrupt record: %v", err)
	}
	readArgs, _ := json.Marshal(map[string]string{"relay_id": out.RelayID})
	first, err := client.Call(ctx, "read", readArgs)
	if err != nil {
		t.Fatalf("first read: %v", err)
	}
	if !bytes.Contains(first, []byte("checksum-mismatch")) {
		t.Fatalf("first read = %s, want checksum-mismatch", first)
	}
	deadline := time.Now().Add(2 * time.Second)
	quarantined := false
	for time.Now().Before(deadline) {
		second, err := client.Call(ctx, "read", readArgs)
		if err != nil {
			t.Fatalf("second read: %v", err)
		}
		if bytes.Contains(second, []byte("record-quarantined")) {
			quarantined = true
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
	if !quarantined {
		t.Fatalf("read never returned record-quarantined for %s", out.RelayID)
	}

	operatorClient, err := channel.DialAuthenticated(ctx, sock, operatorCred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated operator stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = operatorClient.Close() }()
	operatorMeta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	owed := submitLiveRecord(t, ctx, operatorClient, stderr, reg, pinned.Digest, operatorMeta, record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "post-quarantine-owed"},
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"AUTHORITY":        "report-only",
			"SUBJECT":          "owed after quarantine",
			"record_kind":      "owed_item",
			"owner":            "s3",
			"source":           "live quarantine",
			"target_surface":   "fresh table validation",
			"disposition_path": "fold report",
		},
	})
	disposition := submitLiveRecord(t, ctx, operatorClient, stderr, reg, pinned.Digest, operatorMeta, record.Record{
		Envelope: record.Envelope{To: "operator", DispatchID: "post-quarantine-owed"},
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"AUTHORITY":     "report-only",
			"SUBJECT":       "owed disposition after quarantine",
			"record_kind":   "owed_disposition",
			"disposes_owed": owed.RelayID,
		},
	})
	if disposition.State != record.Accepted {
		t.Fatalf("post-quarantine disposition = %+v, want accepted", disposition)
	}
}

func TestFrankBinaryQuarantinesCorruptStoreBeforeServing(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "corrupt-at-start", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "corrupt at startup"},
	}, nil); err != nil {
		t.Fatalf("Commit corrupt-at-start: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "records", "corrupt-at-start.json"), []byte(`{"bad":true}`), 0o644); err != nil {
		t.Fatalf("corrupt startup record: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-startup-quarantine-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)
	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = client.Close() }()
	readArgs, _ := json.Marshal(map[string]string{"relay_id": "corrupt-at-start"})
	read, err := client.Call(ctx, "read", readArgs)
	if err != nil {
		t.Fatalf("read corrupt-at-start stderr=%s: %v", stderr.String(), err)
	}
	if !bytes.Contains(read, []byte("record-quarantined")) {
		t.Fatalf("read corrupt-at-start = %s, want record-quarantined", read)
	}
}

func waitForSocket(t *testing.T, sock string) {
	t.Helper()
	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(sock); err == nil {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatalf("socket %s was not created", sock)
}

func buildFrank(t *testing.T, ctx context.Context) string {
	t.Helper()
	_ = ctx
	if frankFixtureBinary.err != nil {
		t.Fatalf("cached frank binary: %v", frankFixtureBinary.err)
	}
	return frankFixtureBinary.path
}

func startFrank(t *testing.T, ctx context.Context, bin, root, sock string) (*exec.Cmd, *bytes.Buffer) {
	t.Helper()
	cmd := exec.CommandContext(ctx, bin, "-root", root, "-socket", sock, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start frank: %v", err)
	}
	return cmd, &stderr
}

func runOperatorSubmit(t *testing.T, ctx context.Context, bin, sock, payloadFile, credential string) []byte {
	t.Helper()
	deadline := time.Now().Add(3 * time.Second)
	for {
		operator := exec.CommandContext(ctx, bin, "-socket", sock, "-operator-submit", payloadFile, "-credential", credential)
		out, err := operator.CombinedOutput()
		if err == nil {
			return out
		}
		if !bytes.Contains(out, []byte("auth:channel-active")) || time.Now().After(deadline) {
			t.Fatalf("operator-submit failed: %v\n%s", err, out)
		}
		time.Sleep(25 * time.Millisecond)
	}
}

type liveSubmitOutcome struct {
	State   string `json:"state"`
	RelayID string `json:"relay_id"`
}

func submitLiveRecord(t *testing.T, ctx context.Context, client *channel.Client, stderr *bytes.Buffer, reg *fieldspec.Registry, configDigest string, meta seat.SeatMeta, rec record.Record) liveSubmitOutcome {
	t.Helper()
	payload := submitPayloadBytes(t, reg, configDigest, meta, rec)
	result, err := client.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stderr=%s: %v", stderr.String(), err)
	}
	var outcome liveSubmitOutcome
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode submit outcome %s: %v", result, err)
	}
	if outcome.RelayID == "" {
		t.Fatalf("submit outcome missing relay id: %+v", outcome)
	}
	return outcome
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func mustReadFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return data
}

func loadAssemblyRegistry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	root := t.TempDir()
	if err := store.Init(root, writeFixtureConfigSources(t)); err != nil {
		t.Fatalf("init assembly registry generation: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join(root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("load assembly registry: %v", err)
	}
	return reg
}

func submitPayloadBytes(t *testing.T, reg *fieldspec.Registry, configDigest string, meta seat.SeatMeta, rec record.Record) []byte {
	t.Helper()
	if rec.Headers != nil && rec.Headers["PHASE"] != "" && rec.Headers["EVIDENCE_TARGET"] == "" {
		rec.Headers["EVIDENCE_TARGET"] = "E1"
	}
	_, digest := reg.Render(fieldspec.RenderEnv{ConfigDigest: configDigest}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal submit payload: %v", err)
	}
	return payload
}

func mustJSONBytes(t *testing.T, v any) []byte {
	t.Helper()
	payload, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal payload: %v", err)
	}
	return payload
}

func readOptionalFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return data
}

func parseCredentialStdout(t *testing.T, stdout, stderr []byte) string {
	t.Helper()
	if len(stderr) != 0 {
		t.Fatalf("mint wrote stderr on success: %s", stderr)
	}
	text := string(stdout)
	if strings.Count(text, "credential=") != 1 || !strings.HasSuffix(text, "\n") {
		t.Fatalf("mint stdout = %q, want one credential line", text)
	}
	credential := strings.TrimPrefix(strings.TrimSpace(text), "credential=")
	if len(credential) != 64 {
		t.Fatalf("credential length = %d, want 64", len(credential))
	}
	for _, ch := range credential {
		if !strings.ContainsRune("0123456789abcdef", ch) {
			t.Fatalf("credential contains non-hex rune %q", ch)
		}
	}
	return credential
}

func assertCredentialNotInSeatDeliverables(t *testing.T, root, credential string) {
	t.Helper()
	for _, dir := range []string{"records", "mailboxes", "projections", "outbox"} {
		base := filepath.Join(root, dir)
		err := filepath.WalkDir(base, func(path string, entry os.DirEntry, err error) error {
			if err != nil || entry.IsDir() {
				return err
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if bytes.Contains(data, []byte(credential)) {
				return fmt.Errorf("credential leaked into %s", path)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("scan seat deliverables: %v", err)
		}
	}
}
