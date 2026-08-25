// Package transport owns the connector's pinned, single-attempt HTTP/1.1
// provider client and its below-encoder attempt counters.
package transport

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackli/frank/internal/connector/credentials"
	"github.com/jackli/frank/internal/connector/freeze"
)

const (
	TTFB_DEADLINE  = 30 * time.Second
	STALL_DEADLINE = 30 * time.Second
)

var (
	ErrRedirect      = errors.New("transport: redirect refused")
	ErrTTFBDeadline  = errors.New("transport: TTFB_DEADLINE exceeded")
	ErrStallDeadline = errors.New("transport: STALL_DEADLINE exceeded")
)

type Counters struct {
	DialAttempts           uint64
	ConnectionsEstablished uint64
	RequestWriteStarted    uint64
	RequestWriteCompleted  uint64
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
	RetryAfter string
}

type Client struct {
	transport  *http.Transport
	httpClient *http.Client
	stall      time.Duration
	counters   counterState
	postTLS    func(net.Conn) net.Conn
}

type counterState struct {
	dialAttempts           atomic.Uint64
	connectionsEstablished atomic.Uint64
	requestWriteStarted    atomic.Uint64
	requestWriteCompleted  atomic.Uint64
}

type clientConfig struct {
	dialContext func(context.Context, string, string) (net.Conn, error)
	tlsConfig   *tls.Config
	postTLS     func(net.Conn) net.Conn
	ttfb        time.Duration
	stall       time.Duration
}

func NewClient() *Client {
	return newClient(clientConfig{})
}

func newClient(config clientConfig) *Client {
	if config.dialContext == nil {
		dialer := &net.Dialer{}
		config.dialContext = dialer.DialContext
	}
	if config.tlsConfig == nil {
		config.tlsConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	if config.ttfb <= 0 {
		config.ttfb = TTFB_DEADLINE
	}
	if config.stall <= 0 {
		config.stall = STALL_DEADLINE
	}
	client := &Client{stall: config.stall, postTLS: config.postTLS}
	client.transport = &http.Transport{
		Proxy:                 nil,
		DialContext:           client.dialContext(config.dialContext),
		DialTLSContext:        client.dialTLSContext(config.dialContext, config.tlsConfig),
		ForceAttemptHTTP2:     false,
		TLSNextProto:          make(map[string]func(string, *tls.Conn) http.RoundTripper),
		DisableKeepAlives:     true,
		DisableCompression:    true,
		ResponseHeaderTimeout: config.ttfb,
	}
	client.httpClient = &http.Client{
		Transport: client.transport,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return ErrRedirect
		},
	}
	return client
}

func (client *Client) Send(ctx context.Context, frozen *freeze.Request, wire *credentials.WireRequest) (*Response, error) {
	return client.SendGated(ctx, frozen, wire, nil)
}

// SendGated runs beforeInvoke at the last pre-transport boundary, after the
// mutation guard has built the request but before http.Client can dial or
// write. A false result is a typed cancellation race win with zero wire.
func (client *Client) SendGated(ctx context.Context, frozen *freeze.Request, wire *credentials.WireRequest, beforeInvoke func() bool) (*Response, error) {
	request, err := credentials.PrepareHTTPRequest(ctx, frozen, wire)
	if err != nil {
		return nil, err
	}
	if beforeInvoke != nil && !beforeInvoke() {
		return nil, context.Canceled
	}
	var connection *countingConn
	var connectionMu sync.Mutex
	var responseStarted atomic.Bool
	wroteRequest := make(chan error, 1)
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			if counted, ok := info.Conn.(*countingConn); ok {
				counted.responseStarted = &responseStarted
				connectionMu.Lock()
				connection = counted
				connectionMu.Unlock()
			}
		},
		GotFirstResponseByte: func() {
			responseStarted.Store(true)
		},
		WroteRequest: func(info httptrace.WroteRequestInfo) {
			wroteRequest <- info.Err
		},
	}
	request = request.WithContext(httptrace.WithClientTrace(request.Context(), trace))
	response, err := client.httpClient.Do(request)
	connectionMu.Lock()
	counted := connection
	connectionMu.Unlock()
	if counted != nil && counted.writeStarted.Load() {
		<-wroteRequest
		if response != nil && !counted.failedBeforeResponse.Load() {
			counted.complete()
		}
	}
	if err != nil {
		if errors.Is(err, ErrRedirect) {
			return nil, ErrRedirect
		}
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		var networkError net.Error
		if errors.As(err, &networkError) && networkError.Timeout() {
			return nil, fmt.Errorf("%w", ErrTTFBDeadline)
		}
		return nil, err
	}
	if response.StatusCode >= http.StatusMultipleChoices && response.StatusCode < http.StatusBadRequest {
		_ = response.Body.Close()
		return nil, ErrRedirect
	}
	body := io.ReadCloser(response.Body)
	if counted != nil {
		body = &deadlineBody{body: response.Body, connection: counted, deadline: client.stall, ctx: ctx}
	}
	return &Response{
		StatusCode: response.StatusCode,
		Header:     response.Header.Clone(),
		Body:       body,
		RetryAfter: response.Header.Get("Retry-After"),
	}, nil
}

func (client *Client) Counters() Counters {
	return Counters{
		DialAttempts:           client.counters.dialAttempts.Load(),
		ConnectionsEstablished: client.counters.connectionsEstablished.Load(),
		RequestWriteStarted:    client.counters.requestWriteStarted.Load(),
		RequestWriteCompleted:  client.counters.requestWriteCompleted.Load(),
	}
}

func (client *Client) dialContext(dial func(context.Context, string, string) (net.Conn, error)) func(context.Context, string, string) (net.Conn, error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		return client.dial(ctx, dial, network, address)
	}
}

func (client *Client) dialTLSContext(dial func(context.Context, string, string) (net.Conn, error), baseTLS *tls.Config) func(context.Context, string, string) (net.Conn, error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		raw, err := client.dial(ctx, dial, network, address)
		if err != nil {
			return nil, err
		}
		host, _, err := net.SplitHostPort(address)
		if err != nil {
			raw.Close()
			return nil, err
		}
		config := baseTLS.Clone()
		config.ServerName = host
		config.NextProtos = []string{"http/1.1"}
		secured := tls.Client(raw, config)
		if err := secured.HandshakeContext(ctx); err != nil {
			raw.Close()
			return nil, err
		}
		var application net.Conn = secured
		if client.postTLS != nil {
			application = client.postTLS(application)
		}
		return &countingConn{
			Conn: application, started: &client.counters.requestWriteStarted,
			completed: &client.counters.requestWriteCompleted, contentLength: -1,
		}, nil
	}
}

func (client *Client) dial(ctx context.Context, dial func(context.Context, string, string) (net.Conn, error), network, address string) (net.Conn, error) {
	client.counters.dialAttempts.Add(1)
	connection, err := dial(ctx, network, address)
	if err != nil {
		return nil, err
	}
	client.counters.connectionsEstablished.Add(1)
	return connection, nil
}

type countingConn struct {
	net.Conn
	started              *atomic.Uint64
	completed            *atomic.Uint64
	startOnce            sync.Once
	completeOnce         sync.Once
	writeMu              sync.Mutex
	header               []byte
	contentLength        int64
	bodyWritten          int64
	writeStarted         atomic.Bool
	failedBeforeResponse atomic.Bool
	responseStarted      *atomic.Bool
}

func (connection *countingConn) Write(value []byte) (int, error) {
	connection.writeStarted.Store(true)
	connection.startOnce.Do(func() { connection.started.Add(1) })
	written, err := connection.Conn.Write(value)
	if written > 0 {
		connection.observeWritten(value[:written])
	}
	if err != nil || written != len(value) {
		if connection.responseStarted == nil || !connection.responseStarted.Load() {
			connection.failedBeforeResponse.Store(true)
		}
	}
	return written, err
}

func (connection *countingConn) observeWritten(value []byte) {
	connection.writeMu.Lock()
	defer connection.writeMu.Unlock()
	if connection.contentLength == -2 {
		return
	}
	if connection.contentLength >= 0 {
		connection.bodyWritten += int64(len(value))
		connection.markComplete()
		return
	}
	connection.header = append(connection.header, value...)
	headerEnd := bytes.Index(connection.header, []byte("\r\n\r\n"))
	if headerEnd < 0 {
		return
	}
	request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(connection.header[:headerEnd+4])))
	if err != nil || request.ContentLength < 0 {
		connection.header = nil
		connection.contentLength = -2
		return
	}
	_ = request.Body.Close()
	connection.contentLength = request.ContentLength
	connection.bodyWritten = int64(len(connection.header) - headerEnd - 4)
	connection.header = nil
	connection.markComplete()
}

func (connection *countingConn) markComplete() {
	if connection.bodyWritten >= connection.contentLength {
		connection.complete()
	}
}

func (connection *countingConn) complete() {
	connection.completeOnce.Do(func() { connection.completed.Add(1) })
}

type deadlineBody struct {
	body       io.ReadCloser
	connection net.Conn
	deadline   time.Duration
	ctx        context.Context
}

func (body *deadlineBody) Read(destination []byte) (int, error) {
	deadline := time.Now().Add(body.deadline)
	if contextDeadline, ok := body.ctx.Deadline(); ok && contextDeadline.Before(deadline) {
		deadline = contextDeadline
	}
	if err := body.connection.SetReadDeadline(deadline); err != nil {
		return 0, err
	}
	read, err := body.body.Read(destination)
	if err != nil {
		var networkError net.Error
		if errors.As(err, &networkError) && networkError.Timeout() && body.ctx.Err() == nil {
			return read, ErrStallDeadline
		}
	}
	return read, err
}

func (body *deadlineBody) Close() error {
	return body.body.Close()
}
