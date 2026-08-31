//go:build seam

package seam_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/manifest"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
	workercatalog "github.com/The-Frank-Organization/frank/internal/worker/catalog"
	"github.com/The-Frank-Organization/frank/internal/worker/executor"
	"github.com/The-Frank-Organization/frank/internal/worker/fake"
	"github.com/The-Frank-Organization/frank/internal/worker/provider"
	workerruntime "github.com/The-Frank-Organization/frank/internal/worker/runtime"
	"github.com/The-Frank-Organization/frank/internal/worker/turn"
)

const seamTagSentinel = "S16A_SEAM_TAG_ACTIVE"

func TestMain(m *testing.M) {
	fmt.Println(seamTagSentinel)
	os.Exit(m.Run())
}

func contract(t *testing.T, satisfied bool, evidence string) {
	t.Helper()
	if !satisfied {
		row := strings.TrimPrefix(t.Name(), "Test")
		t.Fatalf("%s contract not yet satisfied: %s", row, evidence)
	}
}

func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate seam test source")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "../.."))
}

func source(t *testing.T, relative string) string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(repoRoot(t), filepath.FromSlash(relative)))
	if err != nil {
		t.Fatalf("read %s: %v", relative, err)
	}
	return string(raw)
}

func sourcesContain(t *testing.T, paths []string, needles ...string) bool {
	t.Helper()
	var joined strings.Builder
	for _, path := range paths {
		joined.WriteString(source(t, path))
		joined.WriteByte('\n')
	}
	text := joined.String()
	for _, needle := range needles {
		if !strings.Contains(text, needle) {
			return false
		}
	}
	return true
}

func structHasField(value any, name string) bool {
	typeOf := reflect.TypeOf(value)
	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
	}
	_, ok := typeOf.FieldByName(name)
	return ok
}

func jsonFields(value any) []string {
	typeOf := reflect.TypeOf(value)
	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
	}
	fields := make([]string, 0, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name == "" {
			name = field.Name
		}
		if name != "-" {
			fields = append(fields, name)
		}
	}
	sort.Strings(fields)
	return fields
}

func equalStrings(left, right []string) bool {
	left = append([]string(nil), left...)
	right = append([]string(nil), right...)
	sort.Strings(left)
	sort.Strings(right)
	return reflect.DeepEqual(left, right)
}

func pointer(value string) *string { return &value }

func digest(char byte) string { return strings.Repeat(string(char), 64) }

func bytesDigest(value []byte) string {
	hash := sha256.Sum256(value)
	return hex.EncodeToString(hash[:])
}

func framed(payload []byte) []byte {
	result := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(result[:4], uint32(len(payload)))
	copy(result[4:], payload)
	return result
}

func mustJSON(t *testing.T, value any) []byte {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func appCanonical(input []byte) ([]byte, error) {
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	return appipc.MarshalJCS(value)
}

func canonicalEqual(values ...[]byte) bool {
	if len(values) < 2 {
		return true
	}
	for _, value := range values[1:] {
		if !bytes.Equal(values[0], value) {
			return false
		}
	}
	return true
}

func fieldKind(value any, name string) (reflect.Kind, bool) {
	typeOf := reflect.TypeOf(value)
	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
	}
	field, ok := typeOf.FieldByName(name)
	if !ok {
		return reflect.Invalid, false
	}
	return field.Type.Kind(), true
}

func explain(format string, args ...any) string { return fmt.Sprintf(format, args...) }

type seamEvent struct {
	runID string
	apply func(context.Context, *store.Tx) error
}

func (event seamEvent) RunID() string { return event.runID }
func (event seamEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, event.apply(ctx, tx)
}

type appFixture struct {
	ctx     context.Context
	db      *store.Store
	applier *applier.Host
	runID   string
	turnID  string
}

func newAppFixture(t *testing.T) *appFixture {
	t.Helper()
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "state"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	fixture := &appFixture{ctx: ctx, db: db, applier: host, runID: "run", turnID: "turn"}
	fixture.mutate(t, func(ctx context.Context, tx *store.Tx) error {
		rows := []struct {
			query string
			args  []any
		}{
			{`INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, []any{fixture.runID, []byte("{}"), digest('0'), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1}},
			{`INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, []any{fixture.runID, fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0)}},
			{`INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, []any{fixture.turnID, fixture.runID, fmt.Sprintf("%020d", 1), "ACTIVE", []byte("task"), "fresh", strings.Repeat("a", 32), "PENDING", 1}},
		}
		for _, row := range rows {
			if _, err := tx.ExecContext(ctx, row.query, row.args...); err != nil {
				return err
			}
		}
		return nil
	})
	return fixture
}

func (fixture *appFixture) mutate(t *testing.T, apply func(context.Context, *store.Tx) error) {
	t.Helper()
	if _, err := fixture.applier.Apply(fixture.ctx, seamEvent{runID: fixture.runID, apply: apply}); err != nil {
		t.Fatal(err)
	}
}

func (fixture *appFixture) seedAttempt(t *testing.T, attemptID string) {
	t.Helper()
	fixture.mutate(t, func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `INSERT INTO provider_attempts(attempt_id,run_id,turn_id,turn_epoch,state,logical_surface_digest,created_at) VALUES(?,?,?,?,?,?,?)`, attemptID, fixture.runID, fixture.turnID, fmt.Sprintf("%020d", 1), "OPEN", digest('1'), 1)
		return err
	})
}

type probeControl struct {
	assignment        workerruntime.Assignment
	helloCalls        int
	genesisBodies     []appipc.GenesisCommittedBody
	attachResults     []workerruntime.AttachResult
	attemptRequests   []provider.Request
	streamEnds        []provider.StreamEnd
	e0Events          [][]provider.Event
	authorizeRequests []executor.AuthorizeRequest
	consumeRequests   []executor.ConsumeRequest
	outcomeRecords    []executor.OutcomeRecord
	authorizeReply    executor.AuthorizeReply
	authorizeErr      error
	consumeReply      executor.ConsumeReply
	consumeErr        error
	outcomeErr        error
	terminal          turn.Terminal
}

func (control *probeControl) Hello(context.Context, workerruntime.Hello) (workerruntime.Assignment, error) {
	control.helloCalls++
	return control.assignment, nil
}

// GenesisCommitted is deliberately present before the production Control
// interface grows the notification. A live worker call is therefore observable
// without a source-text proxy.
func (control *probeControl) GenesisCommitted(_ context.Context, body appipc.GenesisCommittedBody) error {
	control.genesisBodies = append(control.genesisBodies, body)
	return nil
}

func (control *probeControl) ReportAttach(_ context.Context, _ string, _ string, result workerruntime.AttachResult) error {
	control.attachResults = append(control.attachResults, result)
	return nil
}

func (*probeControl) WakeForward(context.Context, string) error { return nil }

func (control *probeControl) TurnTerminal(_ context.Context, terminal turn.Terminal) error {
	control.terminal = terminal
	return nil
}

func (control *probeControl) AttemptOpen(_ context.Context, request provider.Request) error {
	control.attemptRequests = append(control.attemptRequests, request)
	return nil
}

func (control *probeControl) RecordStreamEnd(_ context.Context, _ string, end provider.StreamEnd) error {
	control.streamEnds = append(control.streamEnds, end)
	return nil
}

func (control *probeControl) EvaluateProviderE0(_ context.Context, events []provider.Event) error {
	control.e0Events = append(control.e0Events, append([]provider.Event(nil), events...))
	return nil
}

func (control *probeControl) Authorize(_ context.Context, request executor.AuthorizeRequest) (executor.AuthorizeReply, error) {
	control.authorizeRequests = append(control.authorizeRequests, request)
	if control.authorizeErr != nil {
		return executor.AuthorizeReply{}, control.authorizeErr
	}
	if control.authorizeReply.Code != "" {
		return control.authorizeReply, nil
	}
	return executor.AuthorizeReply{Code: executor.AuthorizeGranted, TicketID: "ticket-1", EffectDescriptor: executor.DescriptorForIdentity(request.FrozenIdentity())}, nil
}

func (control *probeControl) Consume(_ context.Context, request executor.ConsumeRequest) (executor.ConsumeReply, error) {
	control.consumeRequests = append(control.consumeRequests, request)
	if control.consumeErr != nil {
		return executor.ConsumeReply{}, control.consumeErr
	}
	if control.consumeReply.Code != "" {
		return control.consumeReply, nil
	}
	return executor.ConsumeReply{Code: executor.ConsumeOK}, nil
}

func (control *probeControl) RecordOutcome(_ context.Context, record executor.OutcomeRecord) error {
	control.outcomeRecords = append(control.outcomeRecords, record)
	return control.outcomeErr
}

type probeBroker struct {
	attachCalls int
	result      workerruntime.AttachResult
	capability  string
}

func (broker *probeBroker) Attach(context.Context, string, workerruntime.AttachTuple) (workerruntime.AttachResult, string, error) {
	broker.attachCalls++
	return broker.result, broker.capability, nil
}

func (*probeBroker) Rediscover(context.Context, string) ([]string, error) { return nil, nil }

type probeProvider struct {
	disposition provider.Disposition
	items       [][]byte
	request     provider.Request
	attempts    int
	tool        workerruntime.ToolCall
	nextErr     error
}

func (peer *probeProvider) Attempt(_ context.Context, request provider.Request) (provider.Disposition, []json.RawMessage, error) {
	peer.attempts++
	peer.request = request
	items := make([]json.RawMessage, len(peer.items))
	for index := range peer.items {
		items[index] = append(json.RawMessage(nil), peer.items[index]...)
	}
	return peer.disposition, items, nil
}

func (*probeProvider) Cancel(context.Context, string, string) (provider.Disposition, error) {
	return provider.CancelledPre, nil
}

func (peer *probeProvider) NextToolCall() (workerruntime.ToolCall, error) {
	if peer.nextErr != nil {
		return workerruntime.ToolCall{}, peer.nextErr
	}
	return peer.tool, nil
}

type relayProbeBackend struct {
	*fake.Backend
	relayReads int
	bashOutput string
}

func (backend *relayProbeBackend) RelayRead(context.Context, string) ([]byte, error) {
	backend.relayReads++
	return []byte(`{"relay_id":"relay-1"}`), nil
}

func (backend *relayProbeBackend) Bash(ctx context.Context, command string, timeout time.Duration) (string, error) {
	if backend.bashOutput != "" {
		return backend.bashOutput, nil
	}
	return backend.Backend.Bash(ctx, command, timeout)
}

type workerProbe struct {
	control    *probeControl
	broker     *probeBroker
	provider   *probeProvider
	backend    *relayProbeBackend
	runtimeDir string
	config     workerruntime.Config
}

func newWorkerProbe(t *testing.T) *workerProbe {
	t.Helper()
	runtimeDir := filepath.Join(t.TempDir(), "runtime")
	if err := os.Mkdir(runtimeDir, 0o700); err != nil {
		t.Fatal(err)
	}
	assignment := workerruntime.Assignment{
		RunID: "run", TurnID: "turn", TurnEpoch: "1", ManifestDigest: digest('a'),
		GenerationID: "generation", CreateAuthID: strings.Repeat("b", 32), BrokerEndpoint: "broker.sock",
		AdmissionRef: turn.AdmissionRef{Kind: "operator_input", TaskInput: "probe tool"},
	}
	return &workerProbe{
		control:    &probeControl{assignment: assignment},
		broker:     &probeBroker{result: workerruntime.AttachOK, capability: "capability"},
		provider:   &probeProvider{disposition: provider.Completed, tool: workerruntime.ToolCall{ID: "call", CanonicalName: "write", Arguments: []byte(`{"path":"answer.txt","content":"x"}`)}},
		backend:    &relayProbeBackend{Backend: fake.NewBackend()},
		runtimeDir: runtimeDir,
		config:     workerruntime.Config{PID: 1, BuildInfo: "test", RuntimeDir: runtimeDir, RunDisposition: "fresh", AttachDeadline: time.Second, AttachBackoff: time.Millisecond},
	}
}

func (probe *workerProbe) run(ctx context.Context) (workerruntime.Result, error) {
	runner := workerruntime.Runner{Control: probe.control, Broker: probe.broker, Provider: probe.provider, Backend: probe.backend}
	return runner.Run(ctx, probe.config)
}

func assignmentWithJSON(t *testing.T, base workerruntime.Assignment, fields map[string]any) workerruntime.Assignment {
	t.Helper()
	raw, err := json.Marshal(base)
	if err != nil {
		t.Fatal(err)
	}
	var value map[string]any
	if err := json.Unmarshal(raw, &value); err != nil {
		t.Fatal(err)
	}
	for key, field := range fields {
		value[key] = field
	}
	raw, err = json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var result workerruntime.Assignment
	if err := json.Unmarshal(raw, &result); err != nil {
		t.Fatal(err)
	}
	return result
}

func lockedManifestFixture(t *testing.T, credentialRef string) (manifest.Frozen, manifest.Gate, error) {
	t.Helper()
	identities := workercatalog.ExpectedIdentities()
	tools := make([]manifest.ToolIdentity, 0, len(identities))
	for _, identity := range identities {
		schema := identity.ToolSchemaDigest
		catalogVersion := identity.ToolImplCatalogVersion
		tool := manifest.ToolIdentity{Name: identity.CanonicalName, SchemaDigest: &schema, CatalogVersion: &catalogVersion}
		if identity.FormSchemaMappingVersion != "" {
			mapping := identity.FormSchemaMappingVersion
			tool.MappingVersion = &mapping
		}
		tools = append(tools, tool)
	}
	policy := []byte(`{"policy":"fixed"}`)
	policyDigest := bytesDigest(policy)
	lane := manifest.LaneID{ModelID: "model", ProviderID: "provider", ServingProfileID: "profile", CompatMode: "native"}
	catalogDigest := workercatalog.ExpectedDigest
	frozen, err := manifest.Build(manifest.BuildInput{
		RunID: "run", PolicySourceRef: "policy", PolicyDigest: policyDigest, PolicyBytes: policy,
		PolicyPinnedLane: lane, ToolSet: tools, ToolCatalogDigest: &catalogDigest,
		ProviderLane:  manifest.ProviderLane{LaneID: lane, LaneCatalogDigest: digest('c'), CredentialRef: credentialRef},
		WorkspaceRoot: t.TempDir(), ReleaseBinding: &manifest.ReleaseBinding{BoundAtRef: "ref", BuildDigests: &manifest.BuildDigests{AppMainBuildDigest: digest('a'), M9WorkerBuildDigest: digest('b'), M8BuildDigest: digest('c')}},
	})
	gate := manifest.Gate{LockedTools: tools, ShippedToolCatalogDigest: catalogDigest, PolicyBytes: policy, LaneCatalogDigest: digest('c')}
	return frozen, gate, err
}
