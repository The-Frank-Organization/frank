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
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
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

func TestSubmitArgumentsRoundTripStructuredStringCarrier(t *testing.T) {
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
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	meta := fieldspec.SeatMeta{Name: "seat-a", Role: "implementer"}
	form, digest := reg.Render(fieldspec.RenderEnv{}, meta, "SITREP", "medium", fieldspec.ClosedGrantState)
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-schema-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return mustJSON(channel.DescriptionResponse{
					Tools:        []string{"submit", "project", "read"},
					SubmitSchema: &form,
					FormDigest:   digest,
				}), nil
			},
			Submit: func(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
				var payload fieldspec.SubmitPayload
				if err := json.Unmarshal(args, &payload); err != nil {
					return nil, err
				}
				violations := reg.Validate(payload.Record, meta, payload.FormDigest, fieldspec.RenderEnv{}, fieldspec.ClosedGrantState)
				if len(violations) > 0 {
					return mustJSON(map[string]any{"state": record.Rejected, "violations": violations}), nil
				}
				return mustJSON(map[string]any{"state": record.Accepted, "headers": payload.Headers, "to": payload.Envelope.To}), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = srv.Close() }()

	canonical := `[{"path":"README.md","status":"in"}]`
	acceptedArgs := map[string]any{
		"headers": map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "structured carrier",
			"SCOPE_DIFF":      canonical,
		},
		"to":          "seat-a",
		"dispatch_id": "schema-roundtrip",
		"body":        "body",
		"form_digest": digest,
	}
	acceptedInput := mcpCall("submit", 1, acceptedArgs)
	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, acceptedInput)
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	text := decodeRPCOutput(t, stdout)[0]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	var accepted struct {
		State   string            `json:"state"`
		Headers map[string]string `json:"headers"`
		To      string            `json:"to"`
	}
	if err := json.Unmarshal([]byte(text), &accepted); err != nil {
		t.Fatalf("decode accepted text %s: %v", text, err)
	}
	if accepted.State != record.Accepted || accepted.Headers["SCOPE_DIFF"] != canonical || accepted.To != "seat-a" {
		t.Fatalf("accepted submit = %+v", accepted)
	}

	badArgs := acceptedArgs
	badHeaders := map[string]string{
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "medium",
		"EVIDENCE_TARGET": "E1",
		"SUBJECT":         "structured carrier",
		"SCOPE_DIFF":      `[{"path": "README.md","status":"in"}]`,
	}
	badArgs["headers"] = badHeaders
	badInput := mcpCall("submit", 2, badArgs)
	stdout, _ = runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, badInput)
	text = decodeRPCOutput(t, stdout)[0]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(text, `"state":"rejected"`) ||
		!strings.Contains(text, `"Field":"SCOPE_DIFF"`) ||
		!strings.Contains(text, `"Class":"canonical-encoding"`) {
		t.Fatalf("non-canonical submit text = %s", text)
	}
}

func TestToolsListUsesRenderedSubmitSchemaWhenReachable(t *testing.T) {
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
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"AUTHORITY": {Type: "enum", Options: []string{"read-only", "report-only"}},
	}}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-list-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return mustJSON(channel.DescriptionResponse{
					Tools:        []string{"submit", "project", "read"},
					SubmitSchema: &form,
					FormDigest:   "render-digest",
				}), nil
			},
			Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"state":"accepted"}`), nil
			},
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
			Read:    func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`{}`), nil },
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = srv.Close() }()

	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, `{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}`+"\n")
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	data := stdout
	if !strings.Contains(data, `"const":"render-digest"`) || !strings.Contains(data, `"AUTHORITY"`) {
		t.Fatalf("tools/list did not include rendered submit schema: %s", data)
	}
}

func TestPhaseSwitchDriftLoop(t *testing.T) {
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
	describeFor := func(phase string) channel.DescriptionResponse {
		form := fieldspec.Form{Fields: map[string]fieldspec.Field{
			"PHASE":   {Type: "enum", Options: []string{phase}},
			"SUBJECT": {Type: "text"},
		}}
		if phase == "PLAN" {
			form.Fields["PLAN_ONLY"] = fieldspec.Field{Type: "text"}
		}
		return channel.DescriptionResponse{Tools: []string{"submit", "project", "read"}, SubmitSchema: &form, FormDigest: strings.ToLower(phase) + "-digest"}
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-drift-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(_ context.Context, payload json.RawMessage) (json.RawMessage, error) {
				var req channel.DescribeRequest
				_ = json.Unmarshal(payload, &req)
				if req.Phase == "" {
					req.Phase = "SITREP"
				}
				return mustJSON(describeFor(req.Phase)), nil
			},
			Submit: func(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
				var payload fieldspec.SubmitPayload
				if err := json.Unmarshal(args, &payload); err != nil {
					return nil, err
				}
				want := strings.ToLower(payload.Headers["PHASE"]) + "-digest"
				if payload.FormDigest != want {
					return mustJSON(map[string]any{
						"state": record.Rejected,
						"violations": []map[string]string{{
							"field": "form_digest",
							"class": "re-render",
						}},
					}), nil
				}
				return mustJSON(map[string]any{"state": record.Accepted}), nil
			},
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
			Read:    func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`{}`), nil },
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = srv.Close() }()

	stalePlan := map[string]any{
		"headers":     map[string]string{"PHASE": "PLAN", "SUBJECT": "phase switch"},
		"body":        "body",
		"form_digest": "sitrep-digest",
	}
	freshPlan := map[string]any{
		"headers":     map[string]string{"PHASE": "PLAN", "SUBJECT": "phase switch", "PLAN_ONLY": "yes"},
		"body":        "body",
		"form_digest": "plan-digest",
	}
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}`,
		strings.TrimSpace(mcpCall("submit", 2, stalePlan)),
		`{"jsonrpc":"2.0","id":3,"method":"tools/list","params":{}}`,
		strings.TrimSpace(mcpCall("submit", 4, freshPlan)),
		"",
	}, "\n")
	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, input)
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	lines := decodeRPCOutput(t, stdout)
	if !strings.Contains(mustMarshalString(t, lines[0]), `"const":"sitrep-digest"`) {
		t.Fatalf("initial tools/list did not serve SITREP digest: %s", stdout)
	}
	if lines[1]["method"] != "notifications/tools/list_changed" {
		t.Fatalf("line 1 = %#v, want list_changed notification; all output:\n%s", lines[1], stdout)
	}
	rejectText := lines[2]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(rejectText, `"class":"re-render"`) || !strings.Contains(rejectText, "form refreshed") {
		t.Fatalf("reject text = %s", rejectText)
	}
	if !strings.Contains(mustMarshalString(t, lines[3]), `"const":"plan-digest"`) || !strings.Contains(mustMarshalString(t, lines[3]), `"PLAN_ONLY"`) {
		t.Fatalf("refreshed tools/list did not serve PLAN schema: %s", stdout)
	}
	acceptedText := lines[4]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(acceptedText, `"state":"accepted"`) {
		t.Fatalf("fresh submit text = %s", acceptedText)
	}
}

func mcpCall(name string, id int, args map[string]any) string {
	payload, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0",
		"id":      id,
		"method":  "tools/call",
		"params": map[string]any{
			"name":      name,
			"arguments": args,
		},
	})
	return string(payload) + "\n"
}

func mustMarshalString(t *testing.T, v any) string {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return string(data)
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
