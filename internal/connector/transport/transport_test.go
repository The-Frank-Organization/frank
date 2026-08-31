package transport

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/connector/authorize"
	"github.com/The-Frank-Organization/frank/internal/connector/catalog"
	"github.com/The-Frank-Organization/frank/internal/connector/credentials"
	"github.com/The-Frank-Organization/frank/internal/connector/freeze"
	"github.com/The-Frank-Organization/frank/internal/connector/policy"
	"github.com/The-Frank-Organization/frank/internal/connector/translate"
)

func TestOnWireCaptureEqualsFrozenEnvelopePlusOneAttachedCredential(t *testing.T) {
	t.Parallel()

	type capture struct {
		method, path, host, body string
		headers                  http.Header
		contentLength            int64
		close                    bool
	}
	captured := make(chan capture, 1)
	server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		body, _ := io.ReadAll(request.Body)
		captured <- capture{
			method: request.Method, path: request.URL.RequestURI(), host: request.Host,
			body: string(body), headers: request.Header.Clone(),
			contentLength: request.ContentLength, close: request.Close,
		}
		writer.WriteHeader(http.StatusNoContent)
	}))
	frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
	for _, header := range frozen.Core().Headers {
		if header.Name == "x-openai-auth" || strings.Contains(header.Value, "sentinel-provider-secret") {
			t.Fatalf("credential present before Attach: %+v", frozen.Core().Headers)
		}
	}
	client := testClient(t, server, nil)
	response, err := client.Send(context.Background(), frozen, wire)
	if err != nil {
		t.Fatal(err)
	}
	_ = response.Body.Close()
	got := <-captured
	if got.method != frozen.Core().Method || got.path != "/v1/responses" || got.body != string(frozen.Body()) {
		t.Fatalf("wire request differs from frozen envelope: %+v", got)
	}
	if got.headers.Get("Content-Type") != "application/json" || got.headers.Get("User-Agent") != "frank-connector/build-1" || got.headers.Get("X-Openai-Auth") != "Bearer sentinel-provider-secret" {
		t.Fatalf("wire headers = %#v", got.headers)
	}
	if len(got.headers) != 5 || got.headers.Get("Connection") != "close" || got.headers.Get("Content-Length") != strconv.Itoa(len(frozen.Body())) || got.headers.Get("Accept-Encoding") != "" {
		t.Fatalf("uncensused wire headers = %#v", got.headers)
	}
	if got.host == "" || got.contentLength != int64(len(frozen.Body())) || !got.close {
		t.Fatalf("derived wire fields host=%q content-length=%d close=%v", got.host, got.contentLength, got.close)
	}
}

func TestPinnedTransportConfiguration(t *testing.T) {
	t.Parallel()

	client := NewClient()
	if client.transport.ForceAttemptHTTP2 || !client.transport.DisableKeepAlives || !client.transport.DisableCompression || client.transport.Proxy != nil {
		t.Fatalf("unpinned transport: %+v", client.transport)
	}
	if client.transport.TLSNextProto == nil || len(client.transport.TLSNextProto) != 0 {
		t.Fatalf("TLSNextProto = %#v, want non-nil empty map", client.transport.TLSNextProto)
	}
	if client.httpClient.CheckRedirect == nil {
		t.Fatal("redirect policy is absent")
	}
}

func TestFreshDialFailureHasOneAttemptAndNoWrite(t *testing.T) {
	t.Parallel()

	frozen, wire := wireRequest(t, "https://provider.test/v1/responses")
	client := newClient(clientConfig{
		dialContext: func(context.Context, string, string) (net.Conn, error) {
			return nil, errors.New("dial refused")
		},
	})
	if _, err := client.Send(context.Background(), frozen, wire); err == nil {
		t.Fatal("Send() error = nil")
	}
	assertCounters(t, client.Counters(), Counters{DialAttempts: 1})
}

func TestPreInvocationCancellationGateHasZeroWire(t *testing.T) {
	t.Parallel()

	frozen, wire := wireRequest(t, "https://provider.test/v1/responses")
	client := newClient(clientConfig{
		dialContext: func(context.Context, string, string) (net.Conn, error) {
			t.Fatal("pre-invocation cancellation reached dial")
			return nil, errors.New("unreachable")
		},
	})
	if _, err := client.SendGated(context.Background(), frozen, wire, func() bool { return false }); !errors.Is(err, context.Canceled) {
		t.Fatalf("SendGated() error = %v, want context.Canceled", err)
	}
	assertCounters(t, client.Counters(), Counters{})
}

func TestPostConnectWriteFailureIsNotRetried(t *testing.T) {
	t.Parallel()

	server := newTLSServer(t, http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
	client := testClient(t, server, func(connection net.Conn) net.Conn {
		return &failWriteConn{Conn: connection}
	})
	if _, err := client.Send(context.Background(), frozen, wire); err == nil {
		t.Fatal("Send() error = nil")
	}
	assertCounters(t, client.Counters(), Counters{DialAttempts: 1, ConnectionsEstablished: 1, RequestWriteStarted: 1})
}

func TestMidBodyWriteFailureDoesNotCompleteRequest(t *testing.T) {
	t.Parallel()

	server := newTLSServer(t, http.HandlerFunc(func(_ http.ResponseWriter, request *http.Request) {
		_, _ = io.Copy(io.Discard, request.Body)
	}))
	frozen, wire := wireRequestWithBody(t, providerEndpoint(server.URL, "/v1/responses"), []byte(strings.Repeat("x", 128*1024)))
	client := testClient(t, server, func(connection net.Conn) net.Conn {
		return &failAfterFirstWriteConn{Conn: connection}
	})
	if _, err := client.Send(context.Background(), frozen, wire); err == nil {
		t.Fatal("Send() error = nil")
	}
	assertCounters(t, client.Counters(), Counters{DialAttempts: 1, ConnectionsEstablished: 1, RequestWriteStarted: 1})
}

func TestHeadersAndMidStreamCutsNeverRetry(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name    string
		partial string
	}{
		{name: "headers received"},
		{name: "mid stream", partial: "data: partial\n\n"},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Type", "text/event-stream")
				writer.Header().Set("Content-Length", "100")
				writer.WriteHeader(http.StatusOK)
				_, _ = io.WriteString(writer, test.partial)
				if flusher, ok := writer.(http.Flusher); ok {
					flusher.Flush()
				}
				panic(http.ErrAbortHandler)
			}))
			frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
			client := testClient(t, server, nil)
			response, err := client.Send(context.Background(), frozen, wire)
			if err != nil {
				t.Fatalf("Send() error = %v", err)
			}
			_, readErr := io.ReadAll(response.Body)
			_ = response.Body.Close()
			if readErr == nil {
				t.Fatal("cut response unexpectedly reached EOF cleanly")
			}
			assertCounters(t, client.Counters(), Counters{DialAttempts: 1, ConnectionsEstablished: 1, RequestWriteStarted: 1, RequestWriteCompleted: 1})
		})
	}
}

func TestNoPoolNoRedirectFollowAndNoH2Offer(t *testing.T) {
	t.Parallel()

	var protocols []string
	server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.Copy(io.Discard, request.Body)
		protocols = append(protocols, request.TLS.NegotiatedProtocol)
		if request.URL.Path == "/redirect" {
			http.Redirect(writer, request, "/elsewhere", http.StatusTemporaryRedirect)
			return
		}
		writer.WriteHeader(http.StatusNoContent)
	}))
	client := testClient(t, server, nil)
	for range 2 {
		frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/ok"))
		response, err := client.Send(context.Background(), frozen, wire)
		if err != nil {
			t.Fatalf("Send() error = %v", err)
		}
		_ = response.Body.Close()
	}
	frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/redirect"))
	if _, err := client.Send(context.Background(), frozen, wire); !errors.Is(err, ErrRedirect) {
		t.Fatalf("redirect error = %v, want ErrRedirect", err)
	}
	assertCounters(t, client.Counters(), Counters{DialAttempts: 3, ConnectionsEstablished: 3, RequestWriteStarted: 3, RequestWriteCompleted: 3})
	for _, protocol := range protocols {
		if protocol == "h2" {
			t.Fatal("client offered/negotiated h2")
		}
	}
}

func TestRetryAfterIsRecordedAndNeverActedOn(t *testing.T) {
	t.Parallel()

	server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Set("Retry-After", "120")
		writer.WriteHeader(http.StatusTooManyRequests)
	}))
	frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
	client := testClient(t, server, nil)
	response, err := client.Send(context.Background(), frozen, wire)
	if err != nil {
		t.Fatalf("Send() error = %v", err)
	}
	defer response.Body.Close()
	if response.RetryAfter != "120" {
		t.Fatalf("RetryAfter = %q", response.RetryAfter)
	}
	assertCounters(t, client.Counters(), Counters{DialAttempts: 1, ConnectionsEstablished: 1, RequestWriteStarted: 1, RequestWriteCompleted: 1})
}

func TestTTFBAndStallDeadlinesComposeWithCallerContext(t *testing.T) {
	t.Parallel()

	t.Run("ttfb", func(t *testing.T) {
		server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			writer.WriteHeader(http.StatusNoContent)
		}))
		frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
		client := testClient(t, server, nil)
		client.transport.ResponseHeaderTimeout = 20 * time.Millisecond
		if _, err := client.Send(context.Background(), frozen, wire); !errors.Is(err, ErrTTFBDeadline) {
			t.Fatalf("Send() error = %v, want ErrTTFBDeadline", err)
		}
		assertCounters(t, client.Counters(), Counters{DialAttempts: 1, ConnectionsEstablished: 1, RequestWriteStarted: 1, RequestWriteCompleted: 1})
	})

	t.Run("stall", func(t *testing.T) {
		server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
			writer.Header().Set("Content-Type", "text/event-stream")
			writer.WriteHeader(http.StatusOK)
			if flusher, ok := writer.(http.Flusher); ok {
				flusher.Flush()
			}
			time.Sleep(100 * time.Millisecond)
		}))
		frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
		client := testClient(t, server, nil)
		client.stall = 20 * time.Millisecond
		response, err := client.Send(context.Background(), frozen, wire)
		if err != nil {
			t.Fatalf("Send() error = %v", err)
		}
		defer response.Body.Close()
		buffer := make([]byte, 1)
		if _, err := response.Body.Read(buffer); !errors.Is(err, ErrStallDeadline) {
			t.Fatalf("Body.Read() error = %v, want ErrStallDeadline", err)
		}
	})

	t.Run("caller cancellation", func(t *testing.T) {
		server := newTLSServer(t, http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
			time.Sleep(100 * time.Millisecond)
			writer.WriteHeader(http.StatusNoContent)
		}))
		frozen, wire := wireRequest(t, providerEndpoint(server.URL, "/v1/responses"))
		client := testClient(t, server, nil)
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
		defer cancel()
		if _, err := client.Send(ctx, frozen, wire); !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("Send() error = %v, want caller deadline", err)
		}
	})
}

func wireRequest(t *testing.T, endpoint string) (*freeze.Request, *credentials.WireRequest) {
	t.Helper()
	return wireRequestWithBody(t, endpoint, []byte(`{}`))
}

func wireRequestWithBody(t *testing.T, endpoint string, body []byte) (*freeze.Request, *credentials.WireRequest) {
	t.Helper()
	lane := catalog.Lane{LaneID: "lane-1", Method: "POST", Endpoint: endpoint, Auth: catalog.Auth{HeaderName: "x-openai-auth", Scheme: "bearer"}}
	frozen, err := freeze.Freeze(translate.Result{Body: body, LoweredTools: []byte(`[]`)}, lane, "build-1")
	if err != nil {
		t.Fatalf("Freeze() error = %v", err)
	}
	const rawPolicyTemplate = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":[%q],"pinned_lane":"lane-1","schema":"m3.egress_policy.v1"}`
	rawPolicy := []byte(strings.ReplaceAll(rawPolicyTemplate, "%q", strconv.Quote(endpoint)))
	loaded, err := policy.Load(rawPolicy, lane)
	if err != nil {
		t.Fatalf("policy.Load() error = %v", err)
	}
	verdict := authorize.Evaluate(authorize.Input{Policy: loaded, PolicyDigest: loaded.Digest, ProviderLaneID: lane.LaneID, PinnedLaneID: lane.LaneID, FrozenCoreDigest: frozen.CoreDigest(), CredentialRef: "provider-main", Lane: lane, Core: frozen.Core()})
	directory := t.TempDir()
	path := filepath.Join(directory, "credentials.json")
	if err := os.WriteFile(path, []byte(`{"entries":{"provider-main":{"secret":"sentinel-provider-secret"}},"schema":"m8.credentials.v1"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	store, err := credentials.Load(path)
	if err != nil {
		t.Fatalf("credentials.Load() error = %v", err)
	}
	wire, err := credentials.Attach(frozen, verdict, store)
	if err != nil {
		t.Fatalf("credentials.Attach() error = %v", err)
	}
	return frozen, wire
}

func newTLSServer(t *testing.T, handler http.Handler) *httptest.Server {
	t.Helper()
	server := httptest.NewUnstartedServer(handler)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	return server
}

func providerEndpoint(serverURL, path string) string {
	_, port, _ := net.SplitHostPort(strings.TrimPrefix(serverURL, "https://"))
	return "https://provider.test:" + port + path
}

func testClient(t *testing.T, server *httptest.Server, postTLS func(net.Conn) net.Conn) *Client {
	t.Helper()
	return newClient(clientConfig{
		dialContext: func(ctx context.Context, network, _ string) (net.Conn, error) {
			var dialer net.Dialer
			return dialer.DialContext(ctx, network, server.Listener.Addr().String())
		},
		tlsConfig: &tls.Config{InsecureSkipVerify: true}, // test server only
		postTLS:   postTLS,
		ttfb:      2 * time.Second,
		stall:     2 * time.Second,
	})
}

func assertCounters(t *testing.T, got, want Counters) {
	t.Helper()
	if got != want {
		t.Fatalf("Counters() = %+v, want %+v", got, want)
	}
}

type failWriteConn struct{ net.Conn }

func (connection *failWriteConn) Write([]byte) (int, error) {
	return 0, errors.New("injected nothing-written failure")
}

type failAfterFirstWriteConn struct {
	net.Conn
	writes int
}

func (connection *failAfterFirstWriteConn) Write(value []byte) (int, error) {
	connection.writes++
	if connection.writes > 1 {
		return 0, errors.New("injected mid-body write failure")
	}
	return connection.Conn.Write(value)
}
