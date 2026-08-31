package settle

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestProducerTotalClassesAndTelemetryNeverInput(t *testing.T) {
	fixture := newFixture(t)
	fixture.seedRows()
	var manifest appipc.SettlementManifest
	var encoded []byte
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		var err error
		manifest, encoded, err = fixture.producer.Produce(ctx, tx, fixture.runID, fixture.turnID, "next")
		return err
	})
	if len(manifest.Entries) != 7 || len(encoded) == 0 {
		t.Fatalf("manifest = %#v bytes=%d", manifest, len(encoded))
	}
	want := []string{
		"tool/call-executed/settled_with_content",
		"tool/call-fault/determinate_no_resume",
		"tool/call-unknown/uncertain",
		"tool/orphan-local/determinate_no_resume",
		"tool/orphan-relay/uncertain",
		"provider/attempt-complete/settled_with_content",
		"provider/attempt-unknown/uncertain",
	}
	for i, entry := range manifest.Entries {
		id := ""
		if entry.ToolCallID != nil {
			id = *entry.ToolCallID
		}
		if entry.AttemptID != nil {
			id = *entry.AttemptID
		}
		got := entry.Kind + "/" + id + "/" + entry.Class
		if got != want[i] {
			t.Fatalf("entry %d = %s want %s", i, got, want[i])
		}
	}
	if strings.Contains(string(encoded), "boundary_cut") {
		t.Fatal("broker telemetry entered the manifest")
	}
}

func TestReceiptDispositionAndCarriageFirstCommittedWins(t *testing.T) {
	fixture := newFixture(t)
	fixture.seedAttempt("attempt", "OPEN")
	result, err := fixture.host.RecordContentReady(fixture.ctx, ContentReadyRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "attempt", RoundIdentity: "round", SeqHWM: "4", GenerationID: fixture.generationID, TurnEpoch: "1", At: 10})
	if err != nil || result != ReceiptRecorded {
		t.Fatalf("receipt=%q err=%v", result, err)
	}
	result, err = fixture.host.RecordContentReady(fixture.ctx, ContentReadyRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "attempt", RoundIdentity: "round", SeqHWM: "4", GenerationID: "retired", TurnEpoch: "0", At: 11})
	if err != nil || result != ReceiptDuplicate {
		t.Fatalf("stale duplicate=%q err=%v", result, err)
	}
	result, err = fixture.host.RecordContentReady(fixture.ctx, ContentReadyRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "attempt", RoundIdentity: "different", SeqHWM: "4", GenerationID: fixture.generationID, TurnEpoch: "1", At: 12})
	if err != nil || result != ReceiptConflict {
		t.Fatalf("receipt conflict=%q err=%v", result, err)
	}

	resumeAction := "re_derive"
	decision, pair, err := fixture.host.ReportDisposition(fixture.ctx, DispositionRequest{RunID: fixture.runID, TurnID: fixture.turnID, TurnEpoch: "1", GenerationID: fixture.generationID, Disposition: "degraded", ResumeAction: &resumeAction, At: 13})
	if err != nil || decision != DispositionRecorded || pair.Disposition != "degraded" {
		t.Fatalf("disposition=%q pair=%#v err=%v", decision, pair, err)
	}
	decision, pair, err = fixture.host.ReportDisposition(fixture.ctx, DispositionRequest{RunID: fixture.runID, TurnID: fixture.turnID, TurnEpoch: "1", GenerationID: fixture.generationID, Disposition: "resumable", At: 14})
	if err != nil || decision != DispositionConflict || pair.Disposition != "degraded" {
		t.Fatalf("disposition conflict=%q pair=%#v err=%v", decision, pair, err)
	}

	d1, d2, d3 := strings.Repeat("1", 64), strings.Repeat("2", 64), strings.Repeat("3", 64)
	if decision, err := fixture.host.RecordAttemptResult(fixture.ctx, AttemptResultRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "attempt", TurnEpoch: "1", Disposition: "sent_completed", FrozenCoreDigest: &d2, ProviderLoweredToolsDigest: &d3, At: 15}); err != nil || decision != CarriageRecorded {
		t.Fatalf("carriage=%q err=%v", decision, err)
	}
	row, err := fixture.host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "attempt")
	if err != nil || row.State != RowPresent || *row.FrozenCoreDigest != d2 || row.LogicalSurfaceDigest != d1 {
		t.Fatalf("row=%#v err=%v", row, err)
	}
	missing, err := fixture.host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "absent")
	if err != nil || missing.State != RowNotFound || missing.LogicalSurfaceDigest != "" {
		t.Fatalf("missing=%#v err=%v", missing, err)
	}
	bad := strings.Repeat("4", 64)
	if decision, err := fixture.host.RecordAttemptResult(fixture.ctx, AttemptResultRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "attempt", TurnEpoch: "1", Disposition: "sent_completed", FrozenCoreDigest: &bad, ProviderLoweredToolsDigest: &d3, At: 16}); err != nil || decision != CarriageConflict {
		t.Fatalf("carriage conflict=%q err=%v", decision, err)
	}
}

type fixture struct {
	t                           *testing.T
	ctx                         context.Context
	db                          *store.Store
	applier                     *applier.Host
	host                        *Host
	producer                    Producer
	runID, turnID, generationID string
}

func newFixture(t *testing.T) *fixture {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	a := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = a.Close(); _ = db.Close() })
	f := &fixture{t: t, ctx: ctx, db: db, applier: a, runID: "run", turnID: "turn", generationID: "generation"}
	f.host = New(a)
	f.producer = Producer{}
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		rows := []struct {
			q string
			a []any
		}{{`INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, []any{f.runID, []byte("{}"), strings.Repeat("0", 64), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1}}, {`INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, []any{f.runID, fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0)}}, {`INSERT INTO workers(generation_id,run_id,turn_epoch,state,attach_result,created_at) VALUES(?,?,?,?,?,?)`, []any{f.generationID, f.runID, fmt.Sprintf("%020d", 1), "LEASED", "attach-ok", 1}}, {`INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, []any{f.turnID, f.runID, fmt.Sprintf("%020d", 1), "INTERRUPTED", []byte("task"), "fresh", strings.Repeat("a", 32), "PENDING", 1}}, {`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{f.runID, "worker", "wl", f.generationID, fmt.Sprintf("%020d", 1), "ACTIVE", 1}}, {`INSERT INTO broker_events(broker_instance_nonce,event_seq,event_type,run_id,turn_epoch,event_bytes,ack_bytes,committed_at) VALUES(?,?,?,?,?,?,?,?)`, []any{"nonce", fmt.Sprintf("%020d", 1), "boundary_cut", f.runID, fmt.Sprintf("%020d", 1), []byte("boundary_cut"), []byte("ack"), 1}}}
		for _, r := range rows {
			if _, err := tx.ExecContext(ctx, r.q, r.a...); err != nil {
				return err
			}
		}
		return nil
	})
	return f
}
func (f *fixture) seedAttempt(id, state string) {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `INSERT INTO provider_attempts(attempt_id,run_id,turn_id,turn_epoch,state,logical_surface_digest,created_at) VALUES(?,?,?,?,?,?,?)`, id, f.runID, f.turnID, fmt.Sprintf("%020d", 1), state, strings.Repeat("1", 64), 1)
		return err
	})
}
func (f *fixture) seedRows() {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		toolRows := []struct{ id, state string }{{"call-executed", "EXECUTED"}, {"call-fault", "NOT_INVOKED_INTEGRITY_FAULT"}, {"call-unknown", "UNKNOWN_TOOL_OUTCOME"}}
		for _, r := range toolRows {
			if _, err := tx.ExecContext(ctx, `INSERT INTO tool_calls(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,invocation_tool_name,invocation_args_digest,expected_tool_name,expected_args_digest,expected_turn_epoch,observed_tool_name,observed_args_digest,observed_turn_epoch,created_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, r.id, f.runID, f.turnID, fmt.Sprintf("%020d", 1), r.state, "read", strings.Repeat("a", 64), nullable(r.state == "EXECUTED", "read"), nullable(r.state == "EXECUTED", strings.Repeat("a", 64)), nullable(r.state == "NOT_INVOKED_INTEGRITY_FAULT", "read"), nullable(r.state == "NOT_INVOKED_INTEGRITY_FAULT", strings.Repeat("a", 64)), nullable(r.state == "NOT_INVOKED_INTEGRITY_FAULT", fmt.Sprintf("%020d", 1)), nullable(r.state == "NOT_INVOKED_INTEGRITY_FAULT", "write"), nullable(r.state == "NOT_INVOKED_INTEGRITY_FAULT", strings.Repeat("b", 64)), nullable(r.state == "NOT_INVOKED_INTEGRITY_FAULT", fmt.Sprintf("%020d", 1)), 1); err != nil {
				return err
			}
		}
		for _, r := range []struct{ id, name string }{{"orphan-local", "read"}, {"orphan-relay", "relay.submit"}} {
			if _, err := tx.ExecContext(ctx, `INSERT INTO tool_authorizations(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,void_reason,canonical_tool_name,canonical_args_digest,effect_descriptor,issued_at) VALUES(?,?,?,?,?,'VOID','expired',?,?,?,?)`, "ticket-"+r.id, f.runID, f.turnID, r.id, fmt.Sprintf("%020d", 1), r.name, strings.Repeat("c", 64), []byte("{}"), 1); err != nil {
				return err
			}
		}
		if err := f.seedAttemptTx(ctx, tx, "attempt-complete", "COMPLETED"); err != nil {
			return err
		}
		if err := f.seedAttemptTx(ctx, tx, "attempt-unknown", "UNKNOWN_PROVIDER_OUTCOME"); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO content_ready_receipts(run_id,turn_id,attempt_id,round_identity,seq_hwm,generation_id,committed_at) VALUES(?,?,?,?,?,?,?)`, f.runID, f.turnID, "attempt-complete", "round", fmt.Sprintf("%020d", 3), f.generationID, 1)
		return err
	})
}
func (f *fixture) seedAttemptTx(ctx context.Context, tx *store.Tx, id, state string) error {
	_, err := tx.ExecContext(ctx, `INSERT INTO provider_attempts(attempt_id,run_id,turn_id,turn_epoch,state,logical_surface_digest,created_at) VALUES(?,?,?,?,?,?,?)`, id, f.runID, f.turnID, fmt.Sprintf("%020d", 1), state, strings.Repeat("1", 64), 1)
	return err
}
func nullable(ok bool, value string) any {
	if ok {
		return value
	}
	return nil
}
func (f *fixture) mutate(fn func(context.Context, *store.Tx) error) {
	_, err := f.applier.Apply(f.ctx, eventFunc{f.runID, fn})
	if err != nil {
		f.t.Fatal(err)
	}
}

type eventFunc struct {
	run string
	fn  func(context.Context, *store.Tx) error
}

func (e eventFunc) RunID() string { return e.run }
func (e eventFunc) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, e.fn(ctx, tx)
}
