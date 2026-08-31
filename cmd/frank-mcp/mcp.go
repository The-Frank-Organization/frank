package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/seatclient/conduct"
	"github.com/The-Frank-Organization/frank/internal/seatclient/formschema"
)

const (
	mcpProtocolVersion = "2024-11-05"
	mcpServerVersion   = "0.4.0"
)

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
	newFacade    func(*channel.Client) conductorFacade
	submitSchema map[string]any
	submitForm   *fieldspec.Form
	formDigest   string
	schemaPhase  string
	schemaTier   string
	mu           sync.Mutex
}

type conductorFacade interface {
	Relay(context.Context, string, json.RawMessage) (json.RawMessage, error)
	Describe(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error)
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
	return &MCPServer{
		opts: opts,
		newFacade: func(authenticated *channel.Client) conductorFacade {
			client, err := conduct.FromAuthenticated(authenticated)
			if err != nil {
				panic(err)
			}
			return client
		},
	}
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
			"name":    "frank-mcp",
			"version": mcpServerVersion,
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
	listChanged := false
	if call.Name == "submit" {
		parsed, err := formschema.ParseSubmitArguments(args)
		if err != nil {
			return errorToolResult(errSchemaInvalid), false
		}
		submitArgs = parsed
		if containsH16SystemOwnedHeader(submitArgs.Headers) {
			return errorToolResult(errSchemaInvalid), false
		}
		client, refreshed, class := s.validateSubmitWithFreshness(submitArgs, args)
		listChanged = refreshed
		if class != "" {
			return errorToolResult(class), listChanged
		}
		payload, err := formschema.SubmitPayloadFromArguments(args)
		if err != nil {
			return errorToolResult(errSchemaInvalid), listChanged
		}
		args = payload
		if client != nil {
			result, client, class := s.callWithReconnect(client, call.Name, args)
			return s.finishSubmitCall(result, client, class, submitArgs, listChanged)
		}
	} else {
		var dispositions []formschema.Disposition
		if call.Name == "project" {
			dispositions = formschema.ValidateProjectArguments(args)
		} else {
			dispositions = formschema.ValidateReadArguments(args)
		}
		if len(dispositions) > 0 {
			return errorToolResult(errSchemaInvalid), false
		}
	}
	client, class := s.ensureClient()
	if class != "" {
		return errorToolResult(class), listChanged
	}
	result, client, class := s.callWithReconnect(client, call.Name, args)
	if call.Name == "submit" {
		return s.finishSubmitCall(result, client, class, submitArgs, listChanged)
	}
	if class != "" {
		return errorToolResult(class), listChanged
	}
	return textToolResult(string(result), false), listChanged
}

func containsH16SystemOwnedHeader(headers map[string]string) bool {
	for _, field := range []string{"hook_contract", "mint_predecessor", "admin_provenance"} {
		if _, present := headers[field]; present {
			return true
		}
	}
	return false
}

func (s *MCPServer) callWithReconnect(client *channel.Client, name string, args json.RawMessage) (json.RawMessage, *channel.Client, string) {
	result, err := s.newFacade(client).Relay(s.opts.Context, "relay."+name, args)
	if err == nil {
		return result, client, ""
	}
	s.closeClient()
	reconnected, class := s.ensureClient()
	if class != "" {
		return nil, nil, class
	}
	// Submit retry is safe because the conductor intake layer replays by content hash
	// instead of executing duplicate accepted commands.
	result, err = s.newFacade(reconnected).Relay(s.opts.Context, "relay."+name, args)
	if err != nil {
		s.closeClient()
		return nil, nil, scrubError(err)
	}
	return result, reconnected, ""
}

func (s *MCPServer) toolsListResult() map[string]any {
	submitSchema := s.cachedSubmitSchema()
	if submitSchema == nil {
		if client, class := s.ensureClient(); class == "" {
			_, _ = s.refreshSubmitSchema(client, "SITREP", "medium")
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

func (s *MCPServer) cachedSubmitFormFor(args submitArguments) (fieldspec.Form, string, bool) {
	phase, tier := formschema.DeclaredPhaseTier(args)
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.submitForm == nil {
		return fieldspec.Form{}, "", false
	}
	if s.schemaPhase != "" && s.schemaPhase != phase {
		return fieldspec.Form{}, "", false
	}
	if s.schemaTier != "" && s.schemaTier != tier {
		return fieldspec.Form{}, "", false
	}
	return *s.submitForm, s.formDigest, true
}

func (s *MCPServer) refreshSubmitSchema(client *channel.Client, phase, tier string) (bool, string) {
	if phase == "" {
		phase = "SITREP"
	}
	if tier == "" {
		tier = "medium"
	}
	describe, err := s.newFacade(client).Describe(s.opts.Context, channel.DescribeRequest{Phase: phase, Tier: tier})
	if err != nil {
		return false, errProtocol
	}
	if describe.SubmitSchema == nil {
		return false, errProtocol
	}
	s.mu.Lock()
	s.submitSchema = SchemaFromForm(*describe.SubmitSchema, describe.FormDigest)
	form := *describe.SubmitSchema
	s.submitForm = &form
	s.formDigest = describe.FormDigest
	s.schemaPhase = phase
	s.schemaTier = tier
	s.mu.Unlock()
	return true, ""
}

func (s *MCPServer) validateSubmitWithFreshness(args submitArguments, encoded json.RawMessage) (*channel.Client, bool, string) {
	form, digest, cached := s.cachedSubmitFormFor(args)
	if cached && len(formschema.ValidateSubmitArguments(form, digest, encoded)) == 0 {
		return nil, false, ""
	}
	client, class := s.ensureClient()
	if class != "" {
		return nil, false, class
	}
	phase, tier := formschema.DeclaredPhaseTier(args)
	refreshed, class := s.refreshSubmitSchema(client, phase, tier)
	if class != "" {
		return client, false, class
	}
	form, digest, cached = s.cachedSubmitFormFor(args)
	if !cached || len(formschema.ValidateSubmitArguments(form, digest, encoded)) > 0 {
		return client, refreshed, errSchemaInvalid
	}
	return client, refreshed, ""
}

func (s *MCPServer) finishSubmitCall(result json.RawMessage, client *channel.Client, class string, args submitArguments, listChanged bool) (mcpToolResult, bool) {
	if class != "" {
		return errorToolResult(class), listChanged
	}
	if !submitRejected(result) {
		return textToolResult(string(result), false), listChanged
	}
	phase, tier := formschema.DeclaredPhaseTier(args)
	refreshed, _ := s.refreshSubmitSchema(client, phase, tier)
	listChanged = listChanged || refreshed
	if submitNeedsReRender(result) {
		result = formschema.ReRenderResult(result)
	}
	return textToolResult(string(result), false), listChanged
}

func submitRejected(result json.RawMessage) bool {
	var outcome struct {
		State string `json:"state"`
	}
	return json.Unmarshal(result, &outcome) == nil && outcome.State == "rejected"
}

// Compatibility consumer retained for the executable H16 state-only census;
// the shared module remains the semantic owner of detail matching.
func submitNeedsReRender(result json.RawMessage) bool {
	var outcome struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil || outcome.State != "rejected" {
		return false
	}
	return containsReRender(result)
}

func containsReRender(data []byte) bool {
	return formschema.SubmitNeedsReRender(data)
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
			InputSchema: formschema.ProjectSchema(),
		},
		{
			Name:        "read",
			Description: "Reads a committed governance relay by relay_id (" + honesty + ").",
			InputSchema: formschema.ReadSchema(),
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
