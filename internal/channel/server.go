package channel

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/seat"
)

type ToolFunc func(context.Context, json.RawMessage) (json.RawMessage, error)

type ToolSet struct {
	Submit   ToolFunc
	Project  ToolFunc
	Read     ToolFunc
	Describe ToolFunc
}

type ToolFactory func(seat.SeatMeta) ToolSet

func FullSurface(ready *engine.Ready, tools ToolSet) ToolSet {
	if ready == nil {
		return ToolSet{}
	}
	return tools
}

func ReadOnlySurface(_ *engine.Diagnostics, tools ToolSet) ToolSet {
	return ToolSet{Project: tools.Project, Read: tools.Read}
}

type DescribeRequest struct {
	Phase string `json:"phase,omitempty"`
	Tier  string `json:"tier,omitempty"`
}

type DescriptionResponse struct {
	Tools        []string          `json:"tools"`
	Descriptions map[string]string `json:"descriptions,omitempty"`
	SubmitSchema *fieldspec.Form   `json:"submit_schema,omitempty"`
	FormDigest   string            `json:"form_digest,omitempty"`
}

type Server struct {
	ln      net.Listener
	tools   ToolSet
	auth    *seat.Manager
	factory ToolFactory
	done    chan struct{}
	clients map[*serverConn]struct{}
	pending [][]byte
	mu      sync.Mutex
}

func Serve(sockPath string, tools ToolSet) (*Server, error) {
	if err := os.MkdirAll(filepath.Dir(sockPath), 0o755); err != nil {
		return nil, err
	}
	if err := os.Remove(sockPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	ln, err := net.Listen("unix", sockPath)
	if err != nil {
		return nil, err
	}
	s := &Server{
		ln:      ln,
		tools:   tools,
		done:    make(chan struct{}),
		clients: map[*serverConn]struct{}{},
	}
	go s.accept()
	return s, nil
}

func ServeAuthenticated(sockPath string, manager *seat.Manager, factory ToolFactory) (*Server, error) {
	s, err := Serve(sockPath, ToolSet{})
	if err != nil {
		return nil, err
	}
	s.auth = manager
	s.factory = factory
	return s, nil
}

func (s *Server) accept() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			select {
			case <-s.done:
				return
			default:
				continue
			}
		}
		sc := &serverConn{server: s, conn: conn, enc: json.NewEncoder(conn)}
		s.mu.Lock()
		s.clients[sc] = struct{}{}
		s.mu.Unlock()
		go sc.run()
	}
}

func (s *Server) Push(frame []byte) error {
	return s.broadcast(frame)
}

func (s *Server) QueuePush(frame []byte) error {
	if !json.Valid(frame) {
		return fmt.Errorf("invalid push frame")
	}
	s.mu.Lock()
	s.pending = append(s.pending, append([]byte(nil), frame...))
	clients := make([]*serverConn, 0, len(s.clients))
	for c := range s.clients {
		clients = append(clients, c)
	}
	s.mu.Unlock()
	return writePushes(clients, frame)
}

func (s *Server) broadcast(frame []byte) error {
	if !json.Valid(frame) {
		return fmt.Errorf("invalid push frame")
	}
	s.mu.Lock()
	clients := make([]*serverConn, 0, len(s.clients))
	for c := range s.clients {
		clients = append(clients, c)
	}
	s.mu.Unlock()
	return writePushes(clients, frame)
}

func writePushes(clients []*serverConn, frame []byte) error {
	for _, c := range clients {
		if err := c.write(rpcMessage{Method: "notifications/nudge", Params: frame}); err != nil {
			continue
		}
	}
	return nil
}

func (s *Server) Close() error {
	select {
	case <-s.done:
	default:
		close(s.done)
	}
	err := s.ln.Close()
	s.mu.Lock()
	for c := range s.clients {
		_ = c.conn.Close()
	}
	s.clients = map[*serverConn]struct{}{}
	s.mu.Unlock()
	return err
}

type serverConn struct {
	server *Server
	conn   net.Conn
	enc    *json.Encoder
	mu     sync.Mutex
	tools  ToolSet
	authed bool
}

func (c *serverConn) run() {
	defer func() {
		c.server.mu.Lock()
		delete(c.server.clients, c)
		c.server.mu.Unlock()
		_ = c.conn.Close()
	}()
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		var req rpcMessage
		if err := json.Unmarshal(scanner.Bytes(), &req); err != nil {
			continue
		}
		result, errText := c.handle(req)
		resp := rpcMessage{ID: req.ID, Result: result}
		if errText != "" {
			resp.Error = errText
		}
		_ = c.write(resp)
	}
}

func (c *serverConn) handle(req rpcMessage) (json.RawMessage, string) {
	if c.server.auth != nil && !c.authed {
		if req.Method != "session/connect" {
			return nil, "auth:required"
		}
		var connect connectRequest
		if err := json.Unmarshal(req.Params, &connect); err != nil {
			return nil, "auth:bad-request"
		}
		meta, ok := c.server.auth.Resolve(connect.Credential)
		if !ok {
			return nil, "auth:invalid-credential"
		}
		c.tools = c.server.factory(meta)
		c.authed = true
		c.flushPending()
		return mustJSON(connectResponse{Seat: meta.Name, Role: meta.Role}), ""
	}
	switch req.Method {
	case "tools/list":
		return mustJSON(c.activeTools().names()), ""
	case "tools/descriptions":
		tools := c.activeTools()
		if tools.Describe != nil {
			result, err := tools.Describe(context.Background(), req.Params)
			if err != nil {
				return nil, safeError("tool-error")
			}
			return result, ""
		}
		return mustJSON(DescriptionResponse{Tools: tools.names(), Descriptions: toolDescriptions()}), ""
	case "tools/call":
		var call toolCall
		if err := json.Unmarshal(req.Params, &call); err != nil {
			return nil, "bad tool call"
		}
		tool, ok := c.activeTools().byName(call.Name)
		if !ok {
			return nil, "unknown tool"
		}
		result, err := tool(context.Background(), call.Args)
		if err != nil {
			return nil, safeError("tool-error")
		}
		return result, ""
	default:
		return nil, "unknown method"
	}
}

func (c *serverConn) activeTools() ToolSet {
	if c.server.auth != nil {
		return c.tools
	}
	return c.server.tools
}

func (c *serverConn) flushPending() {
	c.server.mu.Lock()
	pending := make([][]byte, len(c.server.pending))
	for i, frame := range c.server.pending {
		pending[i] = append([]byte(nil), frame...)
	}
	c.server.mu.Unlock()
	for _, frame := range pending {
		_ = c.write(rpcMessage{Method: "notifications/nudge", Params: frame})
	}
}

func (c *serverConn) write(msg rpcMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.enc.Encode(msg)
}

func (t ToolSet) byName(name string) (ToolFunc, bool) {
	switch name {
	case "submit":
		return t.Submit, t.Submit != nil
	case "project":
		return t.Project, t.Project != nil
	case "read":
		return t.Read, t.Read != nil
	default:
		return nil, false
	}
}

func (t ToolSet) names() []string {
	var names []string
	if t.Submit != nil {
		names = append(names, "submit")
	}
	if t.Project != nil {
		names = append(names, "project")
	}
	if t.Read != nil {
		names = append(names, "read")
	}
	return names
}

type Client struct {
	conn    net.Conn
	enc     *json.Encoder
	mu      sync.Mutex
	nextID  atomic.Int64
	pending map[int64]chan rpcMessage
	pushes  chan []byte
	done    chan struct{}
}

func Dial(ctx context.Context, sockPath string) (*Client, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "unix", sockPath)
	if err != nil {
		return nil, err
	}
	c := &Client{
		conn:    conn,
		enc:     json.NewEncoder(conn),
		pending: map[int64]chan rpcMessage{},
		pushes:  make(chan []byte, 16),
		done:    make(chan struct{}),
	}
	go c.readLoop()
	return c, nil
}

func DialAuthenticated(ctx context.Context, sockPath, credential string) (*Client, error) {
	c, err := Dial(ctx, sockPath)
	if err != nil {
		return nil, err
	}
	if _, err := c.request(ctx, "session/connect", mustJSON(connectRequest{Credential: credential})); err != nil {
		_ = c.Close()
		return nil, err
	}
	return c, nil
}

func (c *Client) ListTools(ctx context.Context) ([]string, error) {
	resp, err := c.request(ctx, "tools/list", nil)
	if err != nil {
		return nil, err
	}
	var tools []string
	if err := json.Unmarshal(resp.Result, &tools); err != nil {
		return nil, err
	}
	return tools, nil
}

func (c *Client) ToolDescriptions(ctx context.Context) (map[string]string, error) {
	describe, err := c.DescribeTools(ctx, DescribeRequest{})
	if err != nil {
		return nil, err
	}
	return describe.Descriptions, nil
}

func (c *Client) DescribeTools(ctx context.Context, req DescribeRequest) (DescriptionResponse, error) {
	resp, err := c.request(ctx, "tools/descriptions", mustJSON(req))
	if err != nil {
		return DescriptionResponse{}, err
	}
	var description DescriptionResponse
	if err := json.Unmarshal(resp.Result, &description); err != nil {
		return DescriptionResponse{}, err
	}
	return description, nil
}

func (c *Client) Call(ctx context.Context, name string, args json.RawMessage) (json.RawMessage, error) {
	params := mustJSON(toolCall{Name: name, Args: args})
	resp, err := c.request(ctx, "tools/call", params)
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (c *Client) CallExpectNoTool(ctx context.Context, name string) error {
	_, err := c.Call(ctx, name, nil)
	if err == nil {
		return nil
	}
	return err
}

func (c *Client) NextPush(ctx context.Context) ([]byte, error) {
	select {
	case frame := <-c.pushes:
		return frame, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c.done:
		return nil, net.ErrClosed
	}
}

func (c *Client) Close() error {
	select {
	case <-c.done:
	default:
		close(c.done)
	}
	return c.conn.Close()
}

func (c *Client) request(ctx context.Context, method string, params json.RawMessage) (rpcMessage, error) {
	id := c.nextID.Add(1)
	ch := make(chan rpcMessage, 1)
	c.mu.Lock()
	c.pending[id] = ch
	if err := c.enc.Encode(rpcMessage{ID: id, Method: method, Params: params}); err != nil {
		delete(c.pending, id)
		c.mu.Unlock()
		return rpcMessage{}, err
	}
	c.mu.Unlock()
	select {
	case resp := <-ch:
		if resp.Error != "" {
			return rpcMessage{}, errors.New(resp.Error)
		}
		return resp, nil
	case <-ctx.Done():
		c.mu.Lock()
		delete(c.pending, id)
		c.mu.Unlock()
		return rpcMessage{}, ctx.Err()
	case <-c.done:
		return rpcMessage{}, net.ErrClosed
	}
}

func (c *Client) readLoop() {
	defer func() {
		select {
		case <-c.done:
		default:
			close(c.done)
		}
	}()
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		var msg rpcMessage
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			continue
		}
		if msg.ID != 0 {
			c.mu.Lock()
			ch := c.pending[msg.ID]
			delete(c.pending, msg.ID)
			c.mu.Unlock()
			if ch != nil {
				ch <- msg
			}
			continue
		}
		if msg.Method == "notifications/nudge" {
			select {
			case c.pushes <- msg.Params:
			default:
			}
		}
	}
}

type toolCall struct {
	Name string          `json:"name"`
	Args json.RawMessage `json:"args,omitempty"`
}

type connectRequest struct {
	Credential string `json:"credential"`
}

type connectResponse struct {
	Seat string `json:"seat"`
	Role string `json:"role"`
}

type rpcMessage struct {
	ID     int64           `json:"id,omitempty"`
	Method string          `json:"method,omitempty"`
	Params json.RawMessage `json:"params,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  string          `json:"error,omitempty"`
}

func mustJSON(v any) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

func safeError(class string) string {
	return bounce.Format(fieldspec.Violation{Field: "system", Class: class})
}

func toolDescriptions() map[string]string {
	return map[string]string{
		"submit":  "Submit a stamped governance record through the serialized loop.",
		"project": "List committed relay IDs currently visible to this seat mailbox.",
		"read":    "Read an immutable committed relay record by relay ID.",
	}
}
