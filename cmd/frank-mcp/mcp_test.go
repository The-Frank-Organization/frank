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

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/seatclient/conduct"
	"github.com/The-Frank-Organization/frank/internal/worker/executor"
	"github.com/The-Frank-Organization/frank/internal/worker/relaytool"
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

func TestShimReconnectsAndRetriesSingleCallAfterConnectionLoss(t *testing.T) {
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
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-reconnect-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	first, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`["before"]`), nil
			},
			Read: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"server":"old"}`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated first: %v", err)
	}
	server := NewMCPServer(Options{SocketPath: sock, Credential: cred.Value, Context: ctx})
	defer server.closeClient()
	var facadeCalls []facadeCall
	server.newFacade = recordingFacadeFactory(&facadeCalls)

	firstResult, _ := server.handleToolCall(toolCallParams("project", map[string]any{}))
	if firstResult.IsError || firstResult.Content[0].Text != `["before"]` {
		t.Fatalf("first result = %+v", firstResult)
	}
	if err := first.Close(); err != nil {
		t.Fatalf("close first server: %v", err)
	}

	second, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Read: func(_ context.Context, args json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"server":"new","args":` + string(args) + `}`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated second: %v", err)
	}
	defer func() { _ = second.Close() }()

	retried, _ := server.handleToolCall(toolCallParams("read", map[string]any{"relay_id": "relay-2"}))
	if retried.IsError {
		t.Fatalf("retried call errored: %+v", retried)
	}
	if !strings.Contains(retried.Content[0].Text, `"server":"new"`) || !strings.Contains(retried.Content[0].Text, `"relay_id":"relay-2"`) {
		t.Fatalf("retried text = %s", retried.Content[0].Text)
	}
	wantFacadeNames := []string{"relay.project", "relay.read", "relay.read"}
	gotFacadeNames := make([]string, 0, len(facadeCalls))
	for _, call := range facadeCalls {
		gotFacadeNames = append(gotFacadeNames, call.name)
	}
	if !reflect.DeepEqual(gotFacadeNames, wantFacadeNames) {
		t.Fatalf("facade calls = %#v, want close/re-auth/one-retry trace %#v", gotFacadeNames, wantFacadeNames)
	}
}

func TestShimReconnectRetrySecondFailureSurfacesTyped(t *testing.T) {
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
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-retry-fail-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	first, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`[]`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated first: %v", err)
	}
	server := NewMCPServer(Options{SocketPath: sock, Credential: cred.Value, Context: ctx})
	defer server.closeClient()
	if firstResult, _ := server.handleToolCall(toolCallParams("project", map[string]any{})); firstResult.IsError {
		t.Fatalf("first result = %+v", firstResult)
	}
	if err := first.Close(); err != nil {
		t.Fatalf("close first server: %v", err)
	}

	failed, _ := server.handleToolCall(toolCallParams("read", map[string]any{"relay_id": "relay-2"}))
	if !failed.IsError || !strings.Contains(failed.Content[0].Text, "shim:conductor-unreachable") {
		t.Fatalf("failed retry result = %+v", failed)
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
	text := lastToolResultText(t, decodeRPCOutput(t, stdout))
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
	text = lastToolResultText(t, decodeRPCOutput(t, stdout))
	if !strings.Contains(text, `"state":"rejected"`) ||
		!strings.Contains(text, `"Field":"SCOPE_DIFF"`) ||
		!strings.Contains(text, `"Class":"canonical-encoding"`) {
		t.Fatalf("non-canonical submit text = %s", text)
	}
}

func TestH16MCPRejectsForgedSystemHeadersBeforeConductorCall(t *testing.T) {
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("load registry: %v", err)
	}
	form, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}, "SITREP", "medium", fieldspec.ClosedGrantState)
	server := NewMCPServer(Options{})
	server.submitSchema = SchemaFromForm(form, digest)
	server.schemaPhase = "SITREP"
	server.schemaTier = "medium"

	for _, field := range []string{"hook_contract", "mint_predecessor", "admin_provenance"} {
		for _, value := range []string{"forged", ""} {
			name := field + "/non-empty"
			if value == "" {
				name = field + "/present-empty"
			}
			t.Run(name, func(t *testing.T) {
				result, changed := server.handleToolCall(toolCallParams("submit", map[string]any{
					"headers": map[string]string{
						"PHASE": "SITREP", "CEREMONY_TIER": "medium", "SUBJECT": "forged system header", field: value,
					},
					"form_digest": digest,
				}))
				if changed || !result.IsError || result.Content[0].Text != `{"error_class":"schema_invalid"}` {
					t.Fatalf("MCP result=%+v changed=%v, want typed schema_invalid", result, changed)
				}
				if server.client != nil {
					t.Fatal("schema-invalid request opened a conductor client")
				}
			})
		}
	}

	t.Run("phase-and-tier-mismatch-cannot-bypass-system-owned-gate", func(t *testing.T) {
		result, changed := server.handleToolCall(toolCallParams("submit", map[string]any{
			"headers": map[string]string{
				"PHASE": "PLAN", "CEREMONY_TIER": "high", "SUBJECT": "forged under drift", "admin_provenance": "ceremony",
			},
			"form_digest": "different-phase-digest",
		}))
		if changed || !result.IsError || result.Content[0].Text != `{"error_class":"schema_invalid"}` {
			t.Fatalf("MCP result=%+v changed=%v, want typed schema_invalid", result, changed)
		}
		if server.client != nil {
			t.Fatal("phase-mismatched forged request opened a conductor client")
		}
	})
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

func TestShimRefreshUsesTypedReRenderDetailWithoutReadback(t *testing.T) {
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
		"PHASE":   {Type: "enum", Options: []string{"PLAN"}},
		"SUBJECT": {Type: "text"},
	}}
	var readCalls int
	var describeCalls int
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-detail-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				describeCalls++
				digest := "old-digest"
				if describeCalls > 1 {
					digest = "plan-digest"
				}
				return mustJSON(channel.DescriptionResponse{Tools: []string{"submit", "project", "read"}, SubmitSchema: &form, FormDigest: digest}), nil
			},
			Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"state":"rejected","relay_id":"relay-stale","intake_id":"i-stale","detail":"form_digest:re-render"}`), nil
			},
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
			Read: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				readCalls++
				return json.RawMessage(`{"record":{"body":"form_digest:re-render"}}`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = srv.Close() }()

	args := map[string]any{"headers": map[string]string{"PHASE": "PLAN", "SUBJECT": "stale"}, "form_digest": "old-digest"}
	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, mcpCall("submit", 1, args))
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	if readCalls != 0 {
		t.Fatalf("shim issued %d read calls; refresh must key on typed detail", readCalls)
	}
	lines := decodeRPCOutput(t, stdout)
	if len(lines) != 2 || lines[0]["method"] != "notifications/tools/list_changed" {
		t.Fatalf("output = %s, want list_changed plus submit result", stdout)
	}
	text := lines[1]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(text, `"class":"re-render"`) || !strings.Contains(text, "form refreshed") {
		t.Fatalf("submit text = %s, want re-render refresh payload", text)
	}
}

func TestShimDoesNotUseReadbackToInferReRender(t *testing.T) {
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
	var readCalls int
	var describeCalls int
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"PHASE":   {Type: "enum", Options: []string{"PLAN"}},
		"SUBJECT": {Type: "text"},
	}}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-no-readback-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	srv, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				describeCalls++
				digest := "old-digest"
				if describeCalls > 1 {
					digest = "digest"
				}
				return mustJSON(channel.DescriptionResponse{Tools: []string{"submit", "project", "read"}, SubmitSchema: &form, FormDigest: digest}), nil
			},
			Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"state":"rejected","relay_id":"relay-without-detail","intake_id":"i-stale"}`), nil
			},
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
			Read: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				readCalls++
				return json.RawMessage(`{"record":{"body":"form_digest:re-render"}}`), nil
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = srv.Close() }()

	args := map[string]any{"headers": map[string]string{"PHASE": "PLAN", "SUBJECT": "stale"}, "form_digest": "old-digest"}
	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: cred.Value, Context: ctx}, mcpCall("submit", 1, args))
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	if readCalls != 0 {
		t.Fatalf("shim used rejected-relay readback %d times", readCalls)
	}
	lines := decodeRPCOutput(t, stdout)
	if len(lines) != 2 || lines[0]["method"] != "notifications/tools/list_changed" {
		t.Fatalf("output = %s, want F1/F2 list_changed plus submit result", stdout)
	}
	text := lines[1]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if strings.Contains(text, "form refreshed") {
		t.Fatalf("submit text = %s, refresh inferred without typed detail", text)
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
						"state":  record.Rejected,
						"detail": "form_digest:re-render",
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
	if rejectText != `{"error_class":"schema_invalid"}` {
		t.Fatalf("reject text = %s, want F1 fresh-schema rejection", rejectText)
	}
	if !strings.Contains(mustMarshalString(t, lines[3]), `"const":"plan-digest"`) || !strings.Contains(mustMarshalString(t, lines[3]), `"PLAN_ONLY"`) {
		t.Fatalf("refreshed tools/list did not serve PLAN schema: %s", stdout)
	}
	acceptedText := lines[4]["result"].(map[string]any)["content"].([]any)[0].(map[string]any)["text"].(string)
	if !strings.Contains(acceptedText, `"state":"accepted"`) {
		t.Fatalf("fresh submit text = %s", acceptedText)
	}
}

func TestMCPF1RefreshBeforeRejectHealsSameDigestFieldExpansion(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	mgr, err := seat.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	credential, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatal(err)
	}
	oldForm := fieldspec.Form{Fields: map[string]fieldspec.Field{"SUBJECT": {Type: "text"}}}
	newForm := fieldspec.Form{Fields: map[string]fieldspec.Field{
		"SUBJECT": {Type: "text"},
		"grant":   {Type: "text", ConductorVolatile: true, DigestExempt: true},
	}}
	var describeCalls, submitCalls int
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-f1-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				describeCalls++
				form := &oldForm
				if describeCalls > 1 {
					form = &newForm
				}
				return mustJSON(channel.DescriptionResponse{SubmitSchema: form, FormDigest: "same-digest"}), nil
			},
			Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				submitCalls++
				return json.RawMessage(`{"state":"accepted"}`), nil
			},
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = server.Close() }()

	input := strings.Join([]string{
		strings.TrimSpace(mcpCall("submit", 1, map[string]any{"headers": map[string]string{"SUBJECT": "prime"}, "form_digest": "same-digest"})),
		strings.TrimSpace(mcpCall("submit", 2, map[string]any{"headers": map[string]string{"SUBJECT": "expanded", "grant": "dispatch-impl"}, "form_digest": "same-digest"})),
		"",
	}, "\n")
	stdout, stderr := runMCPGolden(t, Options{SocketPath: sock, Credential: credential.Value, Context: ctx}, input)
	if stderr != "" {
		t.Fatalf("stderr = %s", stderr)
	}
	messages := decodeRPCOutput(t, stdout)
	var notifications int
	for _, message := range messages {
		if message["method"] == "notifications/tools/list_changed" {
			notifications++
		}
	}
	if describeCalls != 2 || submitCalls != 2 || notifications != 2 {
		t.Fatalf("describe=%d submit=%d notifications=%d output=%s", describeCalls, submitCalls, notifications, stdout)
	}
	if text := lastToolResultText(t, messages); !strings.Contains(text, `"state":"accepted"`) {
		t.Fatalf("expanded call did not proceed: %s", text)
	}
}

func TestNativeAndMCPUseSameConductFacadePayloadAndH16StillGates(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	mgr, err := seat.Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	credential, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatal(err)
	}
	form := fieldspec.Form{Fields: map[string]fieldspec.Field{"SUBJECT": {Type: "text"}}}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-mcp-parity-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.ServeAuthenticated(sock, mgr, func(seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return mustJSON(channel.DescriptionResponse{SubmitSchema: &form, FormDigest: "digest"}), nil
			},
			Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.RawMessage(`{"state":"accepted"}`), nil
			},
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = server.Close() }()

	arguments := map[string]any{
		"headers":     map[string]string{"SUBJECT": "same bytes"},
		"to":          "m-9.planner",
		"dispatch_id": "parity",
		"body":        "body",
		"form_digest": "digest",
	}
	mcpServer := NewMCPServer(Options{SocketPath: sock, Credential: credential.Value, Context: ctx})
	defer mcpServer.closeClient()
	var mcpCalls []facadeCall
	mcpServer.newFacade = recordingFacadeFactory(&mcpCalls)
	mcpResult, _ := mcpServer.handleToolCall(toolCallParams("submit", arguments))
	if mcpResult.IsError {
		t.Fatalf("MCP submit = %+v", mcpResult)
	}
	mcpServer.closeClient()

	channelClient, err := channel.DialAuthenticated(ctx, sock, credential.Value)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = channelClient.Close() }()
	nativeClient, err := conduct.FromAuthenticated(channelClient)
	if err != nil {
		t.Fatal(err)
	}
	var nativeCalls []facadeCall
	nativeFacade := &recordingConductorFacade{delegate: nativeClient, calls: &nativeCalls}
	nativeRegistry := relaytool.New(nativeFacade, nil)
	encodedArguments, _ := json.Marshal(arguments)
	value, err := nativeRegistry.Invoke(ctx, executor.Invocation{
		Identity:  executor.Identity{CanonicalToolName: "relay.submit"},
		Arguments: encodedArguments,
	})
	if err != nil || value.(relaytool.Result).Error != "" {
		t.Fatalf("native submit = %+v err=%v", value, err)
	}

	mcpCall := lastFacadeCall(t, mcpCalls, "relay.submit")
	nativeCall := lastFacadeCall(t, nativeCalls, "relay.submit")
	if mcpCall.name != nativeCall.name || mcpCall.args != nativeCall.args {
		t.Fatalf("frontend facade bytes differ:\nMCP:    %+v\nnative: %+v", mcpCall, nativeCall)
	}

	before := len(mcpCalls)
	forged := map[string]any{
		"headers":     map[string]string{"SUBJECT": "forged", "admin_provenance": "forged"},
		"form_digest": "digest",
	}
	guarded, changed := mcpServer.handleToolCall(toolCallParams("submit", forged))
	if changed || !guarded.IsError || guarded.Content[0].Text != `{"error_class":"schema_invalid"}` {
		t.Fatalf("H16 result=%+v changed=%v", guarded, changed)
	}
	if len(mcpCalls) != before {
		t.Fatalf("H16 guard was bypassed into refactored facade: before=%d after=%d", before, len(mcpCalls))
	}
}

type facadeCall struct {
	name string
	args string
}

type recordingConductorFacade struct {
	delegate conductorFacade
	calls    *[]facadeCall
}

func (facade *recordingConductorFacade) Relay(ctx context.Context, name string, args json.RawMessage) (json.RawMessage, error) {
	*facade.calls = append(*facade.calls, facadeCall{name: name, args: string(args)})
	return facade.delegate.Relay(ctx, name, args)
}

func (facade *recordingConductorFacade) Describe(ctx context.Context, request channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return facade.delegate.Describe(ctx, request)
}

func recordingFacadeFactory(calls *[]facadeCall) func(*channel.Client) conductorFacade {
	return func(authenticated *channel.Client) conductorFacade {
		client, err := conduct.FromAuthenticated(authenticated)
		if err != nil {
			panic(err)
		}
		return &recordingConductorFacade{delegate: client, calls: calls}
	}
}

func lastFacadeCall(t *testing.T, calls []facadeCall, name string) facadeCall {
	t.Helper()
	for index := len(calls) - 1; index >= 0; index-- {
		if calls[index].name == name {
			return calls[index]
		}
	}
	t.Fatalf("facade call %q not found in %#v", name, calls)
	return facadeCall{}
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

func toolCallParams(name string, args map[string]any) json.RawMessage {
	data, _ := json.Marshal(map[string]any{
		"name":      name,
		"arguments": args,
	})
	return data
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

func lastToolResultText(t *testing.T, messages []map[string]any) string {
	t.Helper()
	for index := len(messages) - 1; index >= 0; index-- {
		result, ok := messages[index]["result"].(map[string]any)
		if !ok {
			continue
		}
		content, ok := result["content"].([]any)
		if !ok || len(content) == 0 {
			continue
		}
		item, ok := content[0].(map[string]any)
		if ok {
			if text, ok := item["text"].(string); ok {
				return text
			}
		}
	}
	t.Fatalf("no tool result in messages: %#v", messages)
	return ""
}
