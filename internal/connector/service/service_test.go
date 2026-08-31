package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/connector/control"
	"github.com/The-Frank-Organization/frank/internal/connector/credentials"
	"github.com/The-Frank-Organization/frank/internal/connector/frame"
	"github.com/The-Frank-Organization/frank/internal/connector/freeze"
	"github.com/The-Frank-Organization/frank/internal/connector/jcs"
	"github.com/The-Frank-Organization/frank/internal/connector/stream"
	"github.com/The-Frank-Organization/frank/internal/connector/transport"
)

const (
	testSecret  = "S14_SENTINEL_PROVIDER_SECRET"
	testPrompt  = "S14_SENTINEL_PROMPT"
	testCatalog = `{"lanes":[{"auth":{"auth_header_name":"x-openai-auth","auth_scheme":"bearer"},"compat_mode":"openai-responses","cost":{"effective_time":"2026-07-17T00:00:00Z","input":1,"output":10},"endpoint":"https://api.openai.com/v1/responses","lane_id":"lane-codex-1","limits":{"context":200000,"max_output":100000},"method":"POST","model_id":"gpt-5","observed_at":"2026-07-17T00:00:00Z","profile_facts":{"endpoint_kind":"coding","region":"us-east"},"provider_id":"openai","reasoning":{"effort_levels":["low","medium","high"],"replay_kind":"opaque_item","supported":true},"serving_profile_id":"codex-default","source":"seeded","tool_use":{"strict_schema":true,"supported":true},"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true}}],"schema":"m8.lane_catalog.v1"}`
	testPolicy  = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`
)

func TestFakeCounterpartAndProviderCompleteOneGovernedAttempt(t *testing.T) {
	artifacts := serviceArtifacts(t)
	ctrlConnector, ctrlPeer := net.Pipe()
	t.Cleanup(func() { _ = ctrlPeer.Close() })
	handshake := make(chan *control.Session, 1)
	handshakeErr := make(chan error, 1)
	go func() {
		session, err := control.Bootstrap(context.Background(), ctrlConnector, control.Config{Artifacts: artifacts, BuildInfo: "s14-test", PID: 42})
		if err != nil {
			handshakeErr <- err
			return
		}
		handshake <- session
	}()
	ctrlDecoder := frame.NewDecoder(map[string]frame.TypeSpec{"hello": {}, "connector_ready": {}, "attempt_result": {}})
	hello := readFrame(t, ctrlPeer, ctrlDecoder)
	if hello.Type != "hello" {
		t.Fatalf("first CTRL-C frame = %q", hello.Type)
	}
	assignment := testAssignment(t, artifacts)
	writeFrame(t, ctrlPeer, controlEnvelope(t, 1, "connector_assign", assignment))
	ready := readFrame(t, ctrlPeer, ctrlDecoder)
	if ready.Type != "connector_ready" {
		t.Fatalf("second CTRL-C frame = %q", ready.Type)
	}
	var session *control.Session
	select {
	case session = <-handshake:
	case err := <-handshakeErr:
		t.Fatal(err)
	case <-time.After(2 * time.Second):
		t.Fatal("bootstrap stalled")
	}
	t.Cleanup(session.Close)

	dataConnector, dataPeer := net.Pipe()
	t.Cleanup(func() { _ = dataPeer.Close() })
	provider := &fakeProvider{called: make(chan providerCapture, 1)}
	server, err := New(Config{Control: session, Data: dataConnector, Provider: provider, BuildInfo: "s14-test"})
	if err != nil {
		t.Fatal(err)
	}
	serveResult := make(chan error, 1)
	go func() { serveResult <- server.Serve(context.Background()) }()

	requestBody := canonical(t, `{"attempt_id":"attempt-1","input":[{"kind":"user_text","text":"`+testPrompt+`"}],"instructions":"system","provider_lane_id":"lane-codex-1","reasoning":{"effort":"high"},"run_id":"run-1","sampling":{"max_output_tokens":32,"temperature":0},"schema":"m8.llm_request.v1","tools":[],"turn_epoch":"7","turn_id":"turn-1"}`)
	epoch := frame.Counter(7)
	writeFrame(t, dataPeer, frame.Envelope{Version: 1, Channel: frame.ChannelDataProvider, Type: "llm_request", Seq: 1, RunID: "run-1", TurnEpoch: &epoch, Body: requestBody})

	dataDecoder := frame.NewDecoder(map[string]frame.TypeSpec{"provider_event": {}, "data_reply": {}})
	var dataFrames []frame.Envelope
	for {
		envelope := readFrame(t, dataPeer, dataDecoder)
		dataFrames = append(dataFrames, envelope)
		var event stream.Event
		if envelope.Type == "provider_event" && json.Unmarshal(envelope.Body, &event) == nil && event.Terminal() {
			if event.Kind != stream.Completed {
				t.Fatalf("terminal event = %+v", event)
			}
			break
		}
	}
	result := readFrame(t, ctrlPeer, ctrlDecoder)
	if result.Type != "attempt_result" || !strings.Contains(string(result.Body), `"disposition":"sent_completed"`) {
		t.Fatalf("attempt result = %s", result.Body)
	}
	capture := <-provider.called
	if capture.authorization != "Bearer "+testSecret || !strings.Contains(capture.body, testPrompt) {
		t.Fatalf("provider wire capture auth=%q body=%s", capture.authorization, capture.body)
	}
	for _, envelope := range append(dataFrames, hello, ready, result) {
		raw, _ := json.Marshal(envelope)
		if strings.Contains(string(raw), testSecret) || strings.Contains(string(raw), testPrompt) {
			t.Fatalf("sentinel leaked into connector frame: %s", raw)
		}
	}
	writeFrame(t, ctrlPeer, controlEnvelope(t, 2, "shutdown", assignment))
	select {
	case err := <-serveResult:
		if !errors.Is(err, control.ErrShutdown) {
			t.Fatalf("Serve() error = %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("service did not fail closed on shutdown")
	}
}

type providerCapture struct{ authorization, body string }

type fakeProvider struct{ called chan providerCapture }

func (provider *fakeProvider) Send(ctx context.Context, frozen *freeze.Request, wire *credentials.WireRequest) (*transport.Response, error) {
	request, err := credentials.PrepareHTTPRequest(ctx, frozen, wire)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	provider.called <- providerCapture{authorization: request.Header.Get("X-Openai-Auth"), body: string(body)}
	sse := strings.Join([]string{
		`data: {"type":"response.output_item.added","output_index":0,"item":{"id":"msg-1","type":"message"}}`,
		``,
		`data: {"type":"response.output_text.delta","output_index":0,"delta":"hello"}`,
		``,
		`data: {"type":"response.output_text.done","output_index":0}`,
		``,
		`data: {"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":1,"output_tokens":1}}}`,
		``, "",
	}, "\n")
	return &transport.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(sse))}, nil
}

func serviceArtifacts(t *testing.T) control.Artifacts {
	t.Helper()
	directory := t.TempDir()
	artifacts := control.Artifacts{
		CredentialPath: filepath.Join(directory, "credentials.json"),
		CatalogPath:    filepath.Join(directory, "catalog.json"),
		PolicyPath:     filepath.Join(directory, "policy.json"),
	}
	writeArtifact(t, artifacts.CredentialPath, `{"entries":{"cred-main":{"secret":"`+testSecret+`"}},"schema":"m8.credentials.v1"}`)
	writeArtifact(t, artifacts.CatalogPath, testCatalog)
	writeArtifact(t, artifacts.PolicyPath, testPolicy)
	return artifacts
}

func writeArtifact(t *testing.T, path, value string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(value), 0o600); err != nil {
		t.Fatal(err)
	}
}

func testAssignment(t *testing.T, artifacts control.Artifacts) control.Assign {
	t.Helper()
	catalogRaw, err := os.ReadFile(artifacts.CatalogPath)
	if err != nil {
		t.Fatal(err)
	}
	policyRaw, err := os.ReadFile(artifacts.PolicyPath)
	if err != nil {
		t.Fatal(err)
	}
	return control.Assign{
		RunID: "run-1", TurnEpoch: 7, RunManifestDigest: digest([]byte("manifest")),
		PolicyDigest: digest(policyRaw), ProviderLaneID: "lane-codex-1",
		LaneCatalogDigest: digest(catalogRaw), CredentialRef: "cred-main",
	}
}

func controlEnvelope(t *testing.T, sequence frame.Counter, kind string, assignment control.Assign) frame.Envelope {
	t.Helper()
	body := any(assignment)
	if kind == "shutdown" {
		body = struct{}{}
	}
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	epoch := assignment.TurnEpoch
	return frame.Envelope{Version: 1, Channel: frame.ChannelControlConnector, Type: kind, Seq: sequence, RunID: assignment.RunID, TurnEpoch: &epoch, Body: raw}
}

func writeFrame(t *testing.T, writer io.Writer, envelope frame.Envelope) {
	t.Helper()
	encoded, err := frame.Encode(envelope)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := writer.Write(encoded); err != nil {
		t.Fatal(err)
	}
}

func readFrame(t *testing.T, reader io.Reader, decoder *frame.Decoder) frame.Envelope {
	t.Helper()
	type result struct {
		envelope frame.Envelope
		err      error
	}
	completed := make(chan result, 1)
	go func() {
		envelope, err := decoder.Read(reader)
		completed <- result{envelope: envelope, err: err}
	}()
	select {
	case got := <-completed:
		if got.err != nil {
			t.Fatal(got.err)
		}
		return got.envelope
	case <-time.After(2 * time.Second):
		t.Fatal("frame read stalled")
		return frame.Envelope{}
	}
}

func canonical(t *testing.T, raw string) []byte {
	t.Helper()
	value, err := jcs.Canonicalize([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}
	return value
}

func digest(value []byte) string {
	sum := sha256.Sum256(value)
	return hex.EncodeToString(sum[:])
}
