package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/jackli/frank/internal/channel"
)

const mcpProtocolVersion = "2024-11-05"

type Options struct {
	SocketPath string
	Credential string
	Context    context.Context
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
}

type MCPServer struct {
	opts         Options
	client       *channel.Client
	submitSchema map[string]any
	schemaPhase  string
	schemaTier   string
	mu           sync.Mutex
}

func NewMCPServer(opts Options) *MCPServer {
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	if opts.Stdin == nil {
		opts.Stdin = strings.NewReader("")
	}
	if opts.Stdout == nil {
		opts.Stdout = io.Discard
	}
	if opts.Stderr == nil {
		opts.Stderr = io.Discard
	}
	return &MCPServer{opts: opts}
}

func (s *MCPServer) Serve() error {
	_, _ = s.ensureClient()
	defer s.closeClient()

	reader := bufio.NewReaderSize(s.opts.Stdin, 64*1024)
	encoder := json.NewEncoder(s.opts.Stdout)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && len(line) == 0 {
			if err == io.EOF {
				return nil
			}
			return err
		}
		line = []byte(strings.TrimSpace(string(line)))
		if len(line) == 0 {
			if err == io.EOF {
				return nil
			}
			continue
		}
		responses := s.handle(line)
		for _, resp := range responses {
			if err := encoder.Encode(resp); err != nil {
				return err
			}
		}
		if err == io.EOF {
			return nil
		}
	}
}

func (s *MCPServer) handle(frame []byte) []mcpResponse {
	var req mcpRequest
	if err := json.Unmarshal(frame, &req); err != nil {
		return []mcpResponse{{JSONRPC: "2.0", Error: &rpcError{Code: -32700, Message: "parse error"}}}
	}
	switch req.Method {
	case "notifications/initialized":
		return nil
	case "initialize":
		return []mcpResponse{{JSONRPC: "2.0", ID: req.ID, Result: mustJSON(initializeResult())}}
	case "ping":
		return []mcpResponse{{JSONRPC: "2.0", ID: req.ID, Result: json.RawMessage(`{}`)}}
	case "tools/list":
		return []mcpResponse{{JSONRPC: "2.0", ID: req.ID, Result: mustJSON(s.toolsListResult())}}
	case "tools/call":
		result, listChanged := s.handleToolCall(req.Params)
		responses := make([]mcpResponse, 0, 2)
		if listChanged {
			responses = append(responses, mcpResponse{JSONRPC: "2.0", Method: "notifications/tools/list_changed", Params: json.RawMessage(`{}`)})
		}
		responses = append(responses, mcpResponse{JSONRPC: "2.0", ID: req.ID, Result: mustJSON(result)})
		return responses
	default:
		return []mcpResponse{{JSONRPC: "2.0", ID: req.ID, Error: &rpcError{Code: -32601, Message: "method not found"}}}
	}
}

func initializeResult() map[string]any {
	return map[string]any{
		"protocolVersion": mcpProtocolVersion,
		"capabilities": map[string]any{
			"tools": map[string]any{"listChanged": true},
		},
		"serverInfo": map[string]any{
			"name": "frank-mcp",
		},
	}
}

func (s *MCPServer) handleToolCall(params json.RawMessage) (mcpToolResult, bool) {
	var call mcpToolCall
	if err := json.Unmarshal(params, &call); err != nil {
		return errorToolResult(errProtocol), false
	}
	if call.Name != "submit" && call.Name != "project" && call.Name != "read" {
		return errorToolResult(errProtocol), false
	}
	args := call.Arguments
	if len(args) == 0 {
		args = call.Args
	}
	if len(args) == 0 || string(args) == "null" {
		args = nil
	}
	var submitArgs submitArguments
	if call.Name == "submit" {
		if len(args) > 0 {
			if err := json.Unmarshal(args, &submitArgs); err != nil {
				return errorToolResult(errProtocol), false
			}
		}
		payload, err := SubmitPayloadFromArguments(args)
		if err != nil {
			return errorToolResult(errProtocol), false
		}
		args = payload
	}
	client, class := s.ensureClient()
	if class != "" {
		return errorToolResult(class), false
	}
	result, err := client.Call(s.opts.Context, call.Name, args)
	if err != nil {
		s.closeClient()
		return errorToolResult(scrubError(err)), false
	}
	if call.Name == "submit" && s.submitNeedsReRender(client, result) {
		phase, tier := declaredPhaseTier(submitArgs)
		return textToolResult(string(reRenderResult(result)), false), s.refreshSubmitSchema(client, phase, tier)
	}
	return textToolResult(string(result), false), false
}

func (s *MCPServer) toolsListResult() map[string]any {
	submitSchema := s.cachedSubmitSchema()
	if submitSchema == nil {
		if client, class := s.ensureClient(); class == "" {
			_ = s.refreshSubmitSchema(client, "SITREP", "medium")
			submitSchema = s.cachedSubmitSchema()
		}
	}
	return map[string]any{"tools": mcpTools(submitSchema)}
}

func (s *MCPServer) cachedSubmitSchema() map[string]any {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.submitSchema
}

func (s *MCPServer) refreshSubmitSchema(client *channel.Client, phase, tier string) bool {
	if phase == "" {
		phase = "SITREP"
	}
	if tier == "" {
		tier = "medium"
	}
	describe, err := client.DescribeTools(s.opts.Context, channel.DescribeRequest{Phase: phase, Tier: tier})
	if err != nil || describe.SubmitSchema == nil {
		return false
	}
	s.mu.Lock()
	s.submitSchema = SchemaFromForm(*describe.SubmitSchema, describe.FormDigest)
	s.schemaPhase = phase
	s.schemaTier = tier
	s.mu.Unlock()
	return true
}

func declaredPhaseTier(args submitArguments) (string, string) {
	phase := args.Headers["PHASE"]
	if phase == "" {
		phase = "SITREP"
	}
	tier := args.Headers["CEREMONY_TIER"]
	if tier == "" {
		tier = "medium"
	}
	return phase, tier
}

func (s *MCPServer) submitNeedsReRender(client *channel.Client, result json.RawMessage) bool {
	if containsReRender(result) {
		return true
	}
	var outcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil || outcome.State != "rejected" || outcome.RelayID == "" {
		return false
	}
	read, err := client.Call(s.opts.Context, "read", mustJSON(map[string]string{"relay_id": outcome.RelayID}))
	if err != nil {
		return false
	}
	return containsReRender(read)
}

func containsReRender(data []byte) bool {
	return bytes.Contains(data, []byte("form_digest")) && bytes.Contains(data, []byte("re-render"))
}

func reRenderResult(original json.RawMessage) json.RawMessage {
	var outcome struct {
		RelayID  string `json:"relay_id,omitempty"`
		IntakeID string `json:"intake_id,omitempty"`
	}
	_ = json.Unmarshal(original, &outcome)
	payload := map[string]any{
		"state": "rejected",
		"violations": []map[string]string{{
			"field": "form_digest",
			"class": "re-render",
			"hint":  "form refreshed - re-read the submit tool schema and re-submit",
		}},
	}
	if outcome.RelayID != "" {
		payload["relay_id"] = outcome.RelayID
	}
	if outcome.IntakeID != "" {
		payload["intake_id"] = outcome.IntakeID
	}
	return mustJSON(payload)
}

func (s *MCPServer) ensureClient() (*channel.Client, string) {
	s.mu.Lock()
	if s.client != nil {
		client := s.client
		s.mu.Unlock()
		return client, ""
	}
	s.mu.Unlock()
	if s.opts.SocketPath == "" || s.opts.Credential == "" {
		return nil, errAuthFailed
	}
	client, err := channel.DialAuthenticated(s.opts.Context, s.opts.SocketPath, s.opts.Credential)
	if err != nil {
		fmt.Fprintf(s.opts.Stderr, "frank-mcp conductor connection failed: %v\n", err)
		return nil, scrubError(err)
	}
	s.mu.Lock()
	s.client = client
	s.mu.Unlock()
	return client, ""
}

func (s *MCPServer) closeClient() {
	s.mu.Lock()
	client := s.client
	s.client = nil
	s.mu.Unlock()
	if client != nil {
		_ = client.Close()
	}
}

func textToolResult(text string, isError bool) mcpToolResult {
	return mcpToolResult{
		Content: []mcpContent{{Type: "text", Text: text}},
		IsError: isError,
	}
}

func errorToolResult(class string) mcpToolResult {
	if class == "" {
		class = errProtocol
	}
	payload, _ := json.Marshal(map[string]string{"error_class": class})
	return textToolResult(string(payload), true)
}

func mcpTools(submitSchema map[string]any) []mcpTool {
	honesty := "transport/provenance only; done-state and record_integrity remain self_reported until Step-2 observe; content claims are not checked by this tool"
	if submitSchema == nil {
		submitSchema = map[string]any{
			"type":                 "object",
			"additionalProperties": true,
		}
	}
	return []mcpTool{
		{
			Name:        "submit",
			Description: "Files a governance relay (" + honesty + ").",
			InputSchema: submitSchema,
		},
		{
			Name:        "project",
			Description: "Lists visible governance relay IDs (" + honesty + ").",
			InputSchema: map[string]any{
				"type":                 "object",
				"properties":           map[string]any{},
				"additionalProperties": false,
			},
		},
		{
			Name:        "read",
			Description: "Reads a committed governance relay by relay_id (" + honesty + ").",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"relay_id": map[string]any{"type": "string"},
				},
				"required":             []string{"relay_id"},
				"additionalProperties": false,
			},
		},
	}
}

func mustJSON(v any) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

type mcpRequest struct {
	JSONRPC string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type mcpResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type mcpTool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

type mcpToolCall struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments,omitempty"`
	Args      json.RawMessage `json:"args,omitempty"`
}

type mcpToolResult struct {
	Content []mcpContent `json:"content"`
	IsError bool         `json:"isError"`
}

type mcpContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

var _ = net.ErrClosed
