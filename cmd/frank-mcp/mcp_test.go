package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/seat"
)

func TestInitializeHandshakeGolden(t *testing.T) {
	stdout, _ := runMCPGolden(t, Options{}, `{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`+"\n")
	lines := decodeRPCOutput(t, stdout)
	result := lines[0]["result"].(map[string]any)
	if result["protocolVersion"] == "" {
		t.Fatalf("initialize missing protocolVersion: %#v", result)
	}
	serverInfo := result["serverInfo"].(map[string]any)
	if serverInfo["name"] != "frank-mcp" {
		t.Fatalf("serverInfo = %#v", serverInfo)
	}
	capabilities := result["capabilities"].(map[string]any)
	tools := capabilities["tools"].(map[string]any)
	if tools["listChanged"] != true {
		t.Fatalf("tools capabilities = %#v", tools)
	}
}

func TestToolsListGolden(t *testing.T) {
	stdout, _ := runMCPGolden(t, Options{}, `{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`+"\n")
	lines := decodeRPCOutput(t, stdout)
	result := lines[0]["result"].(map[string]any)
	tools := result["tools"].([]any)
	var names []string
	for _, tool := range tools {
		item := tool.(map[string]any)
		names = append(names, item["name"].(string))
		description := item["description"].(string)
		if !strings.Contains(description, "transport/provenance only") {
			t.Fatalf("%s description missing honesty line: %q", item["name"], description)
		}
		if item["inputSchema"] == nil {
			t.Fatalf("%s missing inputSchema", item["name"])
		}
	}
	if !reflect.DeepEqual(names, []string{"submit", "project", "read"}) {
		t.Fatalf("tool names = %v", names)
	}
}

func TestUnknownMethodReturnsJSONRPCError(t *testing.T) {
	stdout, _ := runMCPGolden(t, Options{}, `{"jsonrpc":"2.0","id":3,"method":"frank/unknown","params":{}}`+"\n")
	lines := decodeRPCOutput(t, stdout)
	errObj := lines[0]["error"].(map[string]any)
	if errObj["code"] != float64(-32601) || !strings.Contains(errObj["message"].(string), "method not found") {
		t.Fatalf("unknown-method error = %#v", errObj)
	}
}

func TestDialFailureToolCallIsScrubbed(t *testing.T) {
	socketPath := filepath.Join(t.TempDir(), "missing.sock")
	credential := "secret-credential-value"
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"read","arguments":{"relay_id":"r1"}}}`,
		"",
	}, "\n")
	stdout, stderr := runMCPGolden(t, Options{SocketPath: socketPath, Credential: credential}, input)
	for _, forbidden := range []string{socketPath, credential} {
		if strings.Contains(stdout, forbidden) {
			t.Fatalf("stdout leaked %q:\n%s", forbidden, stdout)
		}
	}
	if strings.Contains(stderr, credential) {
		t.Fatalf("stderr leaked credential:\n%s", stderr)
	}
	lines := decodeRPCOutput(t, stdout)
	result := lines[1]["result"].(map[string]any)
	if result["isError"] != true {
		t.Fatalf("tool result = %#v, want isError", result)
	}
	text := result["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(text, "shim:conductor-unreachable") {
		t.Fatalf("tool error text = %q", text)
	}
}

func TestToolsCallProxiesRealAuthenticatedChannel(t *testing.T) {
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
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`["relay-1"]`), nil
			},
			Read: func(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"echo":` + string(args) + `}`), nil
			},
			Submit: func(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"submitted":` + string(args) + `}`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = srv.Close() }()

	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"project","arguments":{}}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"read","arguments":{"relay_id":"relay-1"}}}`,
		"",
	}, "\n")
	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, input)
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	lines := decodeRPCOutput(t, stdout)
	projectText := lines[0]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if projectText != `["relay-1"]` {
		t.Fatalf("project text = %s", projectText)
	}
	readText := lines[1]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(readText, `"relay_id":"relay-1"`) {
		t.Fatalf("read text = %s", readText)
	}
}

func runMCPGolden(t *testing.T, opts Options, stdin string) (string, string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	opts.Stdin = strings.NewReader(stdin)
	opts.Stdout = &stdout
	opts.Stderr = &stderr
	if opts.Context == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		opts.Context = ctx
	}
	if err := NewMCPServer(opts).Serve(); err != nil {
		t.Fatalf("Serve: %v", err)
	}
	return stdout.String(), stderr.String()
}

func decodeRPCOutput(t *testing.T, output string) []map[string]any {
	t.Helper()
	var lines []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var msg map[string]any
		if err := json.Unmarshal([]byte(line), &msg); err != nil {
			t.Fatalf("decode output line %q: %v\nall output:\n%s", line, err, output)
		}
		lines = append(lines, msg)
	}
	if len(lines) == 0 {
		t.Fatalf("no output")
	}
	return lines
}
