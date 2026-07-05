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
	opts   Options
	client *channel.Client
	mu     sync.Mutex
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
		resp, ok := s.handle(line)
		if ok {
			if err := encoder.Encode(resp); err != nil {
				return err
			}
		}
		if err == io.EOF {
			return nil
		}
	}
}

func (s *MCPServer) handle(frame []byte) (mcpResponse, bool) {
	var req mcpRequest
	if err := json.Unmarshal(frame, &req); err != nil {
		return mcpResponse{JSONRPC: "2.0", Error: &rpcError{Code: -32700, Message: "parse error"}}, true
	}
	switch req.Method {
	case "notifications/initialized":
		return mcpResponse{}, false
	case "initialize":
		return mcpResponse{JSONRPC: "2.0", ID: req.ID, Result: mustJSON(initializeResult())}, true
	case "ping":
		return mcpResponse{JSONRPC: "2.0", ID: req.ID, Result: json.RawMessage(`{}`)}, true
	case "tools/list":
		return mcpResponse{JSONRPC: "2.0", ID: req.ID, Result: mustJSON(map[string]any{"tools": mcpTools()})}, true
	case "tools/call":
		result := s.handleToolCall(req.Params)
		return mcpResponse{JSONRPC: "2.0", ID: req.ID, Result: mustJSON(result)}, true
	default:
		return mcpResponse{JSONRPC: "2.0", ID: req.ID, Error: &rpcError{Code: -32601, Message: "method not found"}}, true
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

func (s *MCPServer) handleToolCall(params json.RawMessage) mcpToolResult {
	var call mcpToolCall
	if err := json.Unmarshal(params, &call); err != nil {
		return errorToolResult(errProtocol)
	}
	if call.Name != "submit" && call.Name != "project" && call.Name != "read" {
		return errorToolResult(errProtocol)
	}
	args := call.Arguments
	if len(args) == 0 {
		args = call.Args
	}
	if len(args) == 0 || string(args) == "null" {
		args = nil
	}
	client, class := s.ensureClient()
	if class != "" {
		return errorToolResult(class)
	}
	result, err := client.Call(s.opts.Context, call.Name, args)
	if err != nil {
		s.closeClient()
		return errorToolResult(scrubError(err))
	}
	return textToolResult(string(result), false)
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

func mcpTools() []mcpTool {
	honesty := "transport/provenance only; content claims are not verified by this tool"
	return []mcpTool{
		{
			Name:        "submit",
			Description: "Files a governance relay (" + honesty + ").",
			InputSchema: map[string]any{
				"type":                 "object",
				"additionalProperties": true,
			},
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
