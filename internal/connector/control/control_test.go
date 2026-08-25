package control

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/credentials"
	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/policy"
)

const (
	testCatalog = `{"lanes":[{"auth":{"auth_header_name":"x-openai-auth","auth_scheme":"bearer"},"compat_mode":"openai-responses","cost":{"effective_time":"2026-07-17T00:00:00Z","input":1.25,"output":10},"endpoint":"https://api.openai.com/v1/responses","lane_id":"lane-codex-1","limits":{"context":200000,"max_output":100000},"method":"POST","model_id":"gpt-5","observed_at":"2026-07-17T00:00:00Z","profile_facts":{"endpoint_kind":"coding","region":"us-east"},"provider_id":"openai","reasoning":{"effort_levels":["low","medium","high"],"replay_kind":"opaque_item","supported":true},"serving_profile_id":"codex-default","source":"seeded","tool_use":{"strict_schema":true,"supported":true},"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true}}],"schema":"m8.lane_catalog.v1"}`
	testPolicy  = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`
)

func TestBootstrapLoadsInOrderAndEmitsHelloThenReady(t *testing.T) {
	paths := validArtifacts(t)
	control := newScriptedControl(assignFrame(t, 1, testAssign(paths)))
	session, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "s14-test", PID: 42})
	if err != nil {
		t.Fatalf("Bootstrap() error = %v", err)
	}
	t.Cleanup(session.Close)
	if !session.Ready() || session.Epoch() != 7 || session.RunID() != "run-1" {
		t.Fatalf("boot state ready=%v epoch=%s run=%q", session.Ready(), session.Epoch(), session.RunID())
	}
	if session.Lane().LaneID != "lane-codex-1" || session.Policy() == nil || session.Credentials() == nil {
		t.Fatalf("loaded artifacts lane=%+v policy=%p credentials=%p", session.Lane(), session.Policy(), session.Credentials())
	}

	out := decodeOutbound(t, control.Output(), map[string]frame.TypeSpec{"hello": {}, "connector_ready": {}})
	if len(out) != 2 || out[0].Type != "hello" || out[1].Type != "connector_ready" {
		t.Fatalf("outbound handshake = %+v", out)
	}
	assertBodyFields(t, out[0].Body, map[string]any{"pid": float64(42), "build_info": "s14-test"})
	assertBodyFields(t, out[1].Body, map[string]any{"run_id": "run-1", "turn_epoch": "7"})
}

func TestBootstrapFailsClosedInCredentialCatalogPolicyOrder(t *testing.T) {
	paths := validArtifacts(t)
	if err := os.WriteFile(paths.CredentialPath, []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(paths.CatalogPath, []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	control := newScriptedControl(nil)
	if _, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1}); !errors.Is(err, credentials.ErrCredentialFileInvalid) {
		t.Fatalf("credential-first error = %v", err)
	}
	if len(control.Output()) != 0 {
		t.Fatal("hello emitted before credential load succeeded")
	}

	paths = validArtifacts(t)
	if err := os.WriteFile(paths.CatalogPath, []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	control = newScriptedControl(nil)
	if _, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1}); !errors.Is(err, catalog.ErrInvalidCatalog) {
		t.Fatalf("catalog-second error = %v", err)
	}
	if len(control.Output()) != 0 {
		t.Fatal("hello emitted before catalog load succeeded")
	}

	paths = validArtifacts(t)
	if err := os.Remove(paths.PolicyPath); err != nil {
		t.Fatal(err)
	}
	control = newScriptedControl(nil)
	if _, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1}); !errors.Is(err, policy.ErrPolicyUnavailable) {
		t.Fatalf("policy-read-third error = %v", err)
	}
	if len(control.Output()) != 0 {
		t.Fatal("hello emitted before policy bytes loaded")
	}

	paths = validArtifacts(t)
	if err := os.WriteFile(paths.PolicyPath, []byte("{}"), 0o600); err != nil {
		t.Fatal(err)
	}
	control = newScriptedControl(assignFrame(t, 1, testAssign(paths)))
	if _, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1}); !errors.Is(err, policy.ErrPolicyUnavailable) {
		t.Fatalf("policy-third error = %v", err)
	}
	out := decodeOutbound(t, control.Output(), map[string]frame.TypeSpec{"hello": {}})
	if len(out) != 1 || out[0].Type != "hello" {
		t.Fatalf("policy failure output = %+v", out)
	}
}

func TestBootstrapWithholdsReadyOnAnyAssignmentMismatch(t *testing.T) {
	paths := validArtifacts(t)
	base := testAssign(paths)
	mutations := []func(*Assign){
		func(value *Assign) { value.RunManifestDigest = "bad" },
		func(value *Assign) { value.PolicyDigest = digest([]byte("wrong")) },
		func(value *Assign) { value.ProviderLaneID = "lane-absent" },
		func(value *Assign) { value.LaneCatalogDigest = digest([]byte("wrong")) },
		func(value *Assign) { value.CredentialRef = "missing-ref" },
	}
	for index, mutate := range mutations {
		assignment := base
		mutate(&assignment)
		control := newScriptedControl(assignFrame(t, 1, assignment))
		if session, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: index + 1}); err == nil || session != nil {
			t.Fatalf("mismatch %d accepted", index)
		}
		out := decodeOutbound(t, control.Output(), map[string]frame.TypeSpec{"hello": {}})
		if len(out) != 1 || out[0].Type != "hello" {
			t.Fatalf("mismatch %d output = %+v", index, out)
		}
	}
}

func TestReadyGateAndM10OnlyEpochAuthority(t *testing.T) {
	paths := validArtifacts(t)
	input := append(assignFrame(t, 1, testAssign(paths)), controlFrame(t, 2, "epoch_update", 9, map[string]any{"run_id": "run-1", "turn_epoch": "9"})...)
	input = append(input, controlFrame(t, 3, "ping", 9, map[string]any{})...)
	input = append(input, controlFrame(t, 4, "shutdown", 9, map[string]any{})...)
	control := newScriptedControl(input)
	session, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(session.Close)

	if got, err := session.FenceDataEpoch(context.Background(), 6); err != nil || got != StaleEpoch {
		t.Fatalf("stale fence = %q, %v", got, err)
	}
	if got, err := session.FenceDataEpoch(context.Background(), 7); err != nil || got != EpochAllowed {
		t.Fatalf("current fence = %q, %v", got, err)
	}
	if got, err := session.FenceDataEpoch(context.Background(), 9); err != nil || got != EpochAhead {
		t.Fatalf("ahead fence = %q, %v", got, err)
	}
	if session.Epoch() != 7 {
		t.Fatalf("peer-presented epoch advanced authority to %s", session.Epoch())
	}
	if err := session.HandleControl(context.Background()); err != nil || session.Epoch() != 9 {
		t.Fatalf("m10 epoch update error=%v epoch=%s", err, session.Epoch())
	}
	if got, _ := session.FenceDataEpoch(context.Background(), 7); got != StaleEpoch {
		t.Fatalf("post-update old epoch = %q", got)
	}
	if err := session.HandleControl(context.Background()); err != nil {
		t.Fatalf("ping: %v", err)
	}
	if err := session.HandleControl(context.Background()); !errors.Is(err, ErrShutdown) || session.Ready() {
		t.Fatalf("shutdown error=%v ready=%v", err, session.Ready())
	}

	out := decodeOutbound(t, control.Output(), map[string]frame.TypeSpec{
		"hello": {}, "connector_ready": {}, "epoch_query": {}, "pong": {Reply: true},
	})
	if len(out) != 4 || out[2].Type != "epoch_query" || out[3].Type != "pong" || out[3].Re == nil || *out[3].Re != 3 {
		t.Fatalf("control responses = %+v", out)
	}
}

func TestFenceDataEpochReevaluatesUpdateDuringBoundedHold(t *testing.T) {
	paths := validArtifacts(t)
	input := append(assignFrame(t, 1, testAssign(paths)), controlFrame(t, 2, "epoch_update", 9, map[string]any{"run_id": "run-1", "turn_epoch": "9"})...)
	control := newScriptedControl(input)
	session, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(session.Close)
	control.waitWrites(t, 2)

	type fenceOutcome struct {
		result FenceResult
		err    error
	}
	outcome := make(chan fenceOutcome, 1)
	go func() {
		result, err := session.FenceDataEpoch(context.Background(), 9)
		outcome <- fenceOutcome{result: result, err: err}
	}()
	control.waitWrites(t, 1) // epoch_query is on the wire before m-10 answers.
	if err := session.HandleControl(context.Background()); err != nil {
		t.Fatalf("epoch update: %v", err)
	}

	select {
	case got := <-outcome:
		if got.err != nil || got.result != EpochAllowed {
			t.Fatalf("re-evaluated fence = %q, %v; want %q", got.result, got.err, EpochAllowed)
		}
	case <-time.After(time.Second):
		t.Fatal("ahead epoch fence did not finish after the query answer")
	}
}

func TestControlEOFFailsClosedAndPreReadyGateRejects(t *testing.T) {
	var zero Session
	if got, err := zero.FenceDataEpoch(context.Background(), 0); !errors.Is(err, ErrNotReady) || got != NotReady {
		t.Fatalf("pre-ready fence = %q, %v", got, err)
	}

	paths := validArtifacts(t)
	control := newScriptedControl(assignFrame(t, 1, testAssign(paths)))
	session, err := Bootstrap(context.Background(), control, Config{Artifacts: paths, BuildInfo: "test", PID: 1})
	if err != nil {
		t.Fatal(err)
	}
	if err := session.HandleControl(context.Background()); !errors.Is(err, ErrControlEOF) || session.Ready() {
		t.Fatalf("EOF error=%v ready=%v", err, session.Ready())
	}
	session.Close()
}

func validArtifacts(t *testing.T) Artifacts {
	t.Helper()
	dir := t.TempDir()
	paths := Artifacts{
		CredentialPath: filepath.Join(dir, "credentials.json"),
		CatalogPath:    filepath.Join(dir, "catalog.json"),
		PolicyPath:     filepath.Join(dir, "policy.json"),
	}
	if err := os.WriteFile(paths.CredentialPath, []byte(`{"entries":{"cred-main":{"secret":"SENTINEL_SECRET"}},"schema":"m8.credentials.v1"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(paths.CatalogPath, []byte(testCatalog), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(paths.PolicyPath, []byte(testPolicy), 0o600); err != nil {
		t.Fatal(err)
	}
	return paths
}

func testAssign(paths Artifacts) Assign {
	catalogRaw, _ := os.ReadFile(paths.CatalogPath)
	policyRaw, _ := os.ReadFile(paths.PolicyPath)
	return Assign{
		RunID: "run-1", TurnEpoch: 7, RunManifestDigest: digest([]byte("manifest")),
		PolicyDigest: digest(policyRaw), ProviderLaneID: "lane-codex-1",
		LaneCatalogDigest: digest(catalogRaw), CredentialRef: "cred-main",
	}
}

func assignFrame(t *testing.T, seq frame.Counter, assignment Assign) []byte {
	t.Helper()
	return controlFrame(t, seq, "connector_assign", assignment.TurnEpoch, assignment)
}

func controlFrame(t *testing.T, seq frame.Counter, kind string, epoch frame.Counter, body any) []byte {
	t.Helper()
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := frame.Encode(frame.Envelope{Version: 1, Channel: frame.ChannelControlConnector, Type: kind, Seq: seq, RunID: "run-1", TurnEpoch: &epoch, Body: raw})
	if err != nil {
		t.Fatal(err)
	}
	return encoded
}

func decodeOutbound(t *testing.T, raw []byte, known map[string]frame.TypeSpec) []frame.Envelope {
	t.Helper()
	reader := bytes.NewReader(raw)
	decoder := frame.NewDecoder(known)
	var envelopes []frame.Envelope
	for reader.Len() > 0 {
		envelope, err := decoder.Read(reader)
		if err != nil {
			t.Fatalf("decode outbound: %v", err)
		}
		envelopes = append(envelopes, envelope)
	}
	return envelopes
}

func assertBodyFields(t *testing.T, raw []byte, want map[string]any) {
	t.Helper()
	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	for key, value := range want {
		if got[key] != value {
			t.Fatalf("body[%s] = %#v, want %#v in %s", key, got[key], value, raw)
		}
	}
}

func digest(raw []byte) string {
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:])
}

type scriptedControl struct {
	reader *bytes.Reader
	mu     sync.Mutex
	output bytes.Buffer
	writes chan struct{}
}

func newScriptedControl(input []byte) *scriptedControl {
	return &scriptedControl{reader: bytes.NewReader(input), writes: make(chan struct{}, 16)}
}

func (control *scriptedControl) Read(destination []byte) (int, error) {
	return control.reader.Read(destination)
}

func (control *scriptedControl) Write(source []byte) (int, error) {
	control.mu.Lock()
	defer control.mu.Unlock()
	written, err := control.output.Write(source)
	select {
	case control.writes <- struct{}{}:
	default:
	}
	return written, err
}

func (control *scriptedControl) Close() error { return nil }

func (control *scriptedControl) Output() []byte {
	control.mu.Lock()
	defer control.mu.Unlock()
	return append([]byte(nil), control.output.Bytes()...)
}

func (control *scriptedControl) waitWrites(t *testing.T, count int) {
	t.Helper()
	for range count {
		select {
		case <-control.writes:
		case <-time.After(time.Second):
			t.Fatalf("timed out waiting for control write %d of %d", count, count)
		}
	}
}

var _ io.ReadWriteCloser = (*scriptedControl)(nil)
