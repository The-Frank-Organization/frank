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
	"github.com/jackli/frank/internal/seat"
)

func TestS4IPHGrepRejectsPlantedLeak(t *testing.T) {
	if os.Getenv("FRANK_S4_IPH_PLANTED_LEAK") != "1" {
		t.Skip("set FRANK_S4_IPH_PLANTED_LEAK=1 for the deliberate red leg")
	}
	assertNoS4IPHLeaks(t, s4IPHForbidden{StoreRoot: "/tmp/frank-store", SocketPath: "/tmp/frank.sock", Credential: "secret"}, s4IPHCapture{Name: "planted", Bytes: "leak /tmp/frank-store binding/seats.json secret"})
}

func TestS4IPHBridgeSurfaceMatrix(t *testing.T) {
	h := newS4ShimHarness(t)
	cred := h.mint(t, "seat-a", "implementer")
	h.start(t)
	mcpBin := buildFrankMCP(t, h.ctx)
	forbidden := s4IPHForbidden{StoreRoot: h.root, SocketPath: h.sock, Credential: cred.Value}

	initialOut, initialErr := runFrankMCP(t, h.ctx, mcpBin, h.sock, cred.Value, `{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}`+"\n")
	if initialErr != "" {
		t.Fatalf("initial mcp stderr = %s", initialErr)
	}
	digest := submitDigestFromToolsList(t, initialOut)

	rejected := mcpSubmitInput(2, map[string]any{
		"headers": map[string]string{
			"PHASE":         "SITREP",
			"AUTHORITY":     "merge-gated",
			"CEREMONY_TIER": "medium",
			"SUBJECT":       "rejected surface",
		},
		"to":          "seat-a",
		"dispatch_id": "iph-reject",
		"body":        "body",
		"form_digest": digest,
	})
	rejectOut, rejectErr := runFrankMCP(t, h.ctx, mcpBin, h.sock, cred.Value, rejected)
	if rejectErr != "" {
		t.Fatalf("reject mcp stderr = %s", rejectErr)
	}

	stalePlan := mcpSubmitInput(3, map[string]any{
		"headers": map[string]string{
			"PHASE":         "PLAN",
			"AUTHORITY":     "report-only",
			"CEREMONY_TIER": "medium",
			"SUBJECT":       "drift surface",
		},
		"to":          "seat-a",
		"dispatch_id": "iph-drift",
		"body":        "body",
		"form_digest": digest,
	})
	driftOut, driftErr := runFrankMCP(t, h.ctx, mcpBin, h.sock, cred.Value, strings.Join([]string{
		strings.TrimSpace(stalePlan),
		`{"jsonrpc":"2.0","id":4,"method":"tools/list","params":{}}`,
		"",
	}, "\n"))
	if driftErr != "" {
		t.Fatalf("drift mcp stderr = %s", driftErr)
	}
	if !strings.Contains(driftOut, "notifications/tools/list_changed") {
		t.Fatalf("drift output missing list_changed notification:\n%s", driftOut)
	}

	badCredOut, badCredErr := runFrankMCP(t, h.ctx, mcpBin, h.sock, "bad-"+cred.Value, `{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"name":"project","arguments":{}}}`+"\n")
	if strings.Contains(badCredErr, cred.Value) {
		t.Fatalf("bad credential stderr leaked credential: %s", badCredErr)
	}
	if !strings.Contains(badCredOut, "shim:auth-failed") {
		t.Fatalf("bad credential output = %s", badCredOut)
	}

	missingSock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-missing-%d.sock", time.Now().UnixNano()))
	diagOut, diagErr := runFrankMCP(t, h.ctx, mcpBin, missingSock, cred.Value, `{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"read","arguments":{"relay_id":"missing"}}}`+"\n")
	if strings.Contains(diagErr, cred.Value) {
		t.Fatalf("diagnostic stderr leaked credential: %s", diagErr)
	}
	if !strings.Contains(diagOut, "shim:conductor-unreachable") {
		t.Fatalf("diagnostic output = %s", diagOut)
	}

	assertNoS4IPHLeaks(t, forbidden,
		s4IPHCapture{Name: "tools-list-descriptions", Bytes: initialOut},
		s4IPHCapture{Name: "input-schemas-across-phases", Bytes: initialOut + "\n" + driftOut},
		s4IPHCapture{Name: "tool-call-results", Bytes: rejectOut},
		s4IPHCapture{Name: "notifications-poll-hints", Bytes: driftOut},
		s4IPHCapture{Name: "credential-failure-errors", Bytes: badCredOut},
		s4IPHCapture{Name: "shim-diagnostics", Bytes: diagOut},
	)
}

func TestCarveOutExactlyOneValue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("seat.Open: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-frame-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`"` + strings.Repeat("x", 1024) + `"`), nil
			},
		}
	}, channel.WithFrameLimit(256))
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = server.Close() }()

	mcpBin := buildFrankMCP(t, ctx)
	out, errOut := runFrankMCP(t, ctx, mcpBin, sock, cred.Value, `{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"project","arguments":{}}}`+"\n")
	if errOut != "" {
		t.Fatalf("mcp stderr = %s", errOut)
	}
	if !strings.Contains(out, "shim:frame-too-large") {
		t.Fatalf("frame refusal output = %s", out)
	}
	for _, forbidden := range []string{"max_frame_bytes", "segment_rotate_bytes", "engine.json", "256"} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("frame refusal leaked config value %q in %s", forbidden, out)
		}
	}
}

func TestS4IPHReconnectErrorPathClean(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("seat.Open: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-reconnect-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`[]`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}

	mcpBin := buildFrankMCP(t, ctx)
	cmd := exec.CommandContext(ctx, mcpBin)
	cmd.Env = append(os.Environ(), "FRANK_SOCKET="+sock, "FRANK_CREDENTIAL="+cred.Value)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("StdinPipe: %v", err)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start frank-mcp: %v", err)
	}
	waitForCredentialActive(t, ctx, sock, cred.Value)
	if err := server.Close(); err != nil {
		t.Fatalf("close server: %v", err)
	}
	if _, err := fmt.Fprintln(stdin, `{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"project","arguments":{}}}`); err != nil {
		t.Fatalf("write mcp stdin: %v", err)
	}
	_ = stdin.Close()
	if err := cmd.Wait(); err != nil {
		t.Fatalf("wait frank-mcp: %v\nstdout=%s\nstderr=%s", err, stdout.String(), stderr.String())
	}
	if !strings.Contains(stdout.String(), "shim:connection-lost") {
		t.Fatalf("reconnect error output = %s", stdout.String())
	}
	if strings.Contains(stderr.String(), cred.Value) {
		t.Fatalf("stderr leaked credential: %s", stderr.String())
	}
	assertNoS4IPHLeaks(t, s4IPHForbidden{StoreRoot: root, SocketPath: sock, Credential: cred.Value}, s4IPHCapture{Name: "reconnect-errors", Bytes: stdout.String()})
}

type s4IPHForbidden struct {
	StoreRoot  string
	SocketPath string
	Credential string
}

func waitForCredentialActive(t *testing.T, ctx context.Context, sock, credential string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		client, err := channel.DialAuthenticated(ctx, sock, credential)
		if err != nil {
			if strings.Contains(err.Error(), "auth:channel-active") {
				return
			}
		} else {
			_ = client.Close()
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("credential never became active")
}

type s4IPHCapture struct {
	Name  string
	Bytes string
}

func assertNoS4IPHLeaks(t *testing.T, forbidden s4IPHForbidden, captures ...s4IPHCapture) {
	t.Helper()
	for _, capture := range captures {
		for _, value := range []string{forbidden.StoreRoot, forbidden.SocketPath, forbidden.Credential, "config/", "binding", "seats.json"} {
			if value == "" {
				continue
			}
			if strings.Contains(capture.Bytes, value) {
				t.Fatalf("%s leaked %q:\n%s", capture.Name, value, capture.Bytes)
			}
		}
	}
}

func buildFrankMCP(t *testing.T, ctx context.Context) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "frank-mcp")
	build := exec.CommandContext(ctx, "go", "build", "-o", bin, filepath.Join("..", "..", "cmd", "frank-mcp"))
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build frank-mcp: %v\n%s", err, out)
	}
	return bin
}

func runFrankMCP(t *testing.T, ctx context.Context, bin, sock, credential, input string) (string, string) {
	t.Helper()
	cmd := exec.CommandContext(ctx, bin)
	cmd.Env = append(os.Environ(), "FRANK_SOCKET="+sock, "FRANK_CREDENTIAL="+credential)
	cmd.Stdin = strings.NewReader(input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("run frank-mcp: %v\nstdout=%s\nstderr=%s", err, stdout.String(), stderr.String())
	}
	return stdout.String(), stderr.String()
}

func submitDigestFromToolsList(t *testing.T, output string) string {
	t.Helper()
	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		var msg struct {
			Result struct {
				Tools []struct {
					Name        string         `json:"name"`
					InputSchema map[string]any `json:"inputSchema"`
				} `json:"tools"`
			} `json:"result"`
		}
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			t.Fatalf("decode tools/list line %s: %v", line, err)
		}
		for _, tool := range msg.Result.Tools {
			if tool.Name != "submit" {
				continue
			}
			props, ok := tool.InputSchema["properties"].(map[string]any)
			if !ok {
				t.Fatalf("submit schema missing properties: %s", line)
			}
			formDigest, ok := props["form_digest"].(map[string]any)
			if !ok {
				t.Fatalf("submit schema missing form_digest: %s", line)
			}
			digest, ok := formDigest["const"].(string)
			if !ok || digest == "" {
				t.Fatalf("submit schema digest = %#v", formDigest["const"])
			}
			return digest
		}
	}
	t.Fatalf("no submit tool in output:\n%s", output)
	return ""
}

func mcpSubmitInput(id int, args map[string]any) string {
	payload, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0",
		"id":      id,
		"method":  "tools/call",
		"params": map[string]any{
			"name":      "submit",
			"arguments": args,
		},
	})
	return string(payload) + "\n"
}
