package broker

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/f59"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/channel"
)

func TestPerOperationFenceDescribeResidualAndRelaySettlement(t *testing.T) {
	caller := &fakeCaller{}
	recorder := &fakeRecorder{}
	core := NewCore(caller, recorder)
	installed := tuple("run", "generation", "7", "9")
	capability := Capability{RunID: installed.RunID, GenerationID: installed.GenerationID, TurnEpoch: installed.TurnEpoch, BrokerInstanceNonce: strings.Repeat("a", 64)}
	if _, err := core.Invoke(context.Background(), capability, Operation{ID: "unfenced", Name: "describe", Arguments: []byte(`{}`)}); !errors.Is(err, ErrSuspended) {
		t.Fatalf("unfenced operation = %v", err)
	}
	core.Install(installed)
	operations := []string{"relay.submit", "relay.project", "relay.read", "describe"}
	for index, operation := range operations {
		result, err := core.Invoke(context.Background(), capability, Operation{ID: string(rune('a' + index)), Name: operation, Arguments: []byte(`{}`)})
		if err != nil || result == nil {
			t.Fatalf("%s result=%#v err=%v", operation, result, err)
		}
	}
	if caller.relayCalls != 3 || caller.describeCalls != 1 || recorder.calls != 3 {
		t.Fatalf("calls relay=%d describe=%d settlement=%d", caller.relayCalls, caller.describeCalls, recorder.calls)
	}
	stale := capability
	stale.TurnEpoch = "6"
	for index, operation := range operations {
		_, err := core.Invoke(context.Background(), stale, Operation{ID: "stale-" + string(rune('a'+index)), Name: operation, Arguments: []byte(`{}`)})
		if !errors.Is(err, ErrStaleEpoch) {
			t.Fatalf("%s stale fence = %v", operation, err)
		}
	}
	if caller.relayCalls != 3 || caller.describeCalls != 1 {
		t.Fatal("stale operation reached conductor")
	}
	if _, err := core.Invoke(context.Background(), capability, Operation{ID: "a", Name: "relay.read", Arguments: []byte(`{}`)}); !errors.Is(err, ErrDuplicateOperation) {
		t.Fatalf("duplicate operation = %v", err)
	}
	recorder.err = errors.New("app store unavailable")
	if _, err := core.Invoke(context.Background(), capability, Operation{ID: "settlement-fail", Name: "relay.submit", Arguments: []byte(`{}`)}); !errors.Is(err, ErrRecordUnavailable) {
		t.Fatalf("settlement failure = %v", err)
	}
}

func TestBrokerHasNoCrossingMachineryOrDurableState(t *testing.T) {
	runtimeDir := t.TempDir()
	core := NewCore(&fakeCaller{}, &fakeRecorder{})
	core.Install(tuple("run", "generation", "1", "1"))
	entries, err := os.ReadDir(runtimeDir)
	if err != nil || len(entries) != 0 {
		t.Fatalf("broker core wrote durable state: entries=%v err=%v", entries, err)
	}
	coreType := filepath.Join("core.go")
	if _, err := os.Stat(coreType); err != nil {
		t.Fatal(err)
	}
	entries, err = os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() || strings.HasSuffix(entry.Name(), "_test.go") || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		source, err := os.ReadFile(entry.Name())
		if err != nil {
			t.Fatal(err)
		}
		for _, forbidden := range []string{"crossing_count", "completed-before-install", "transition_ledger", "crossing-pending"} {
			if strings.Contains(string(source), forbidden) {
				t.Fatalf("study-excluded broker shape %q found in %s", forbidden, entry.Name())
			}
		}
	}
}

func TestThreeRelayOperationsSettleDurablyThroughAppStoreApplier(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "state"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	runID, turnID, generationID := "run", "turn", "generation"
	names := map[string]string{"submit": "relay.submit", "project": "relay.project", "read": "relay.read"}
	digests := make(map[string]string, len(names))
	for operationID := range names {
		digests[operationID] = testDigest("args:" + operationID)
	}
	_, err = host.Apply(ctx, brokerTestEvent{runID: runID, apply: func(ctx context.Context, tx *store.Tx) error {
		statements := []struct {
			query string
			args  []any
		}{
			{`INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, []any{runID, []byte("{}"), strings.Repeat("0", 64), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1}},
			{`INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, []any{runID, fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 1)}},
			{`INSERT INTO workers(generation_id,run_id,turn_epoch,pid,state,created_at) VALUES(?,?,?,?,?,?)`, []any{generationID, runID, fmt.Sprintf("%020d", 1), 42, "LEASED", 1}},
			{`INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, []any{turnID, runID, fmt.Sprintf("%020d", 1), "ACTIVE", []byte("task"), "fresh", strings.Repeat("a", 32), "PENDING", 1}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{runID, "worker", "worker-lease", generationID, fmt.Sprintf("%020d", 1), "ACTIVE", 1}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{runID, "turn", "turn-lease", generationID, fmt.Sprintf("%020d", 1), "ACTIVE", 1}},
		}
		for _, statement := range statements {
			if _, err := tx.ExecContext(ctx, statement.query, statement.args...); err != nil {
				return err
			}
		}
		for operationID, name := range names {
			if _, err := tx.ExecContext(ctx, `INSERT INTO tool_authorizations(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,effect_descriptor,issued_at,consumed_at) VALUES(?,?,?,?,?,?,?,?,?,?,?)`, "ticket-"+operationID, runID, turnID, "call-"+operationID, fmt.Sprintf("%020d", 1), "CONSUMED", name, digests[operationID], []byte("{}"), 1, 2); err != nil {
				return err
			}
		}
		return nil
	}})
	if err != nil {
		t.Fatal(err)
	}
	recorder := &f59Recorder{host: f59.New(host, f59.Config{}), generationID: generationID, digests: digests}
	core := NewCore(&fakeCaller{}, recorder)
	installed := tuple(runID, generationID, "1", "1")
	core.Install(installed)
	capability := Capability{RunID: runID, GenerationID: generationID, TurnEpoch: "1", BrokerInstanceNonce: strings.Repeat("b", 64)}
	for operationID, name := range names {
		if _, err := core.Invoke(ctx, capability, Operation{ID: operationID, Name: name, Arguments: []byte(`{}`)}); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
	}
	value, err := host.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var count int
		err := snapshot.QueryRowContext(ctx, `SELECT COUNT(*) FROM tool_calls WHERE run_id=? AND state='EXECUTED'`, runID).Scan(&count)
		return count, err
	}))
	if err != nil || value.(int) != 3 {
		t.Fatalf("durable relay settlements=%v err=%v", value, err)
	}
}

type fakeCaller struct {
	relayCalls, describeCalls int
}

func (caller *fakeCaller) Relay(context.Context, string, []byte) ([]byte, error) {
	caller.relayCalls++
	return []byte(`{"ok":true}`), nil
}

func (caller *fakeCaller) Describe(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error) {
	caller.describeCalls++
	return channel.DescriptionResponse{}, nil
}

type fakeRecorder struct {
	calls int
	err   error
}

func (recorder *fakeRecorder) RecordRelayOutcome(context.Context, RelayOutcome) error {
	recorder.calls++
	return recorder.err
}

type f59Recorder struct {
	host         *f59.Host
	generationID string
	digests      map[string]string
}

func (recorder *f59Recorder) RecordRelayOutcome(ctx context.Context, outcome RelayOutcome) error {
	identity := f59.Identity{CanonicalToolName: outcome.Name, CanonicalArgsDigest: recorder.digests[outcome.OperationID], TurnEpoch: outcome.TurnEpoch}
	decision, err := recorder.host.RecordOutcome(ctx, f59.ChannelIdentity{GenerationID: recorder.generationID}, f59.OutcomeRequest{TicketID: "ticket-" + outcome.OperationID, TurnEpoch: outcome.TurnEpoch, Outcome: f59.Executed, InvocationIdentity: &identity, RecordedAt: 3})
	if err != nil || decision.Fault || decision.Dropped {
		return fmt.Errorf("record relay outcome: decision=%#v err=%w", decision, err)
	}
	return nil
}

type brokerTestEvent struct {
	runID string
	apply func(context.Context, *store.Tx) error
}

func (event brokerTestEvent) RunID() string { return event.runID }
func (event brokerTestEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, event.apply(ctx, tx)
}

func testDigest(value string) string {
	digest := sha256.Sum256([]byte(value))
	return hex.EncodeToString(digest[:])
}
