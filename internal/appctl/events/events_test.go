package events

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

func TestValidatePersistAndNoSilentDrop(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	seedRun(t, ctx, host)
	events := New(host)
	raw := []byte(`{"schema":"m3.app_event.v2","event_kind":"provider_attempt","phase":"completed","scope":"attempt","run_id":"run","turn_id":"turn","attempt_id":"attempt","turn_epoch":"1","provider_lane_id":"lane","run_manifest_digest":"` + strings.Repeat("a", 64) + `","policy_digest":"` + strings.Repeat("b", 64) + `","frozen_core_digest":"` + strings.Repeat("c", 64) + `","logical_surface_digest":"` + strings.Repeat("d", 64) + `","event_evidence":"E0","event_integrity":"self_reported","reported_by":"m-9-worker/g1","event_ts":"2026-08-22T04:00:00Z"}`)
	request := Request{EventID: "event", RunID: "run", TurnID: "turn", ReportedBy: "m-9-worker/g1", Raw: raw, At: 10}
	duplicate, err := events.Persist(ctx, request)
	if err != nil || duplicate {
		t.Fatalf("Persist duplicate=%v err=%v", duplicate, err)
	}
	duplicate, err = events.Persist(ctx, request)
	if err != nil || !duplicate {
		t.Fatalf("replay duplicate=%v err=%v", duplicate, err)
	}
	stored := []byte(nil)
	_, err = host.Read(ctx, applier.QueryFunc(func(ctx context.Context, s *store.Snapshot) (any, error) {
		return nil, s.QueryRowContext(ctx, `SELECT event_bytes FROM pending_app_events WHERE event_id='event'`).Scan(&stored)
	}))
	if err != nil || string(stored) != string(raw) {
		t.Fatalf("stored=%s err=%v", stored, err)
	}

	bad := [][]byte{
		[]byte(`{"schema":"m3.app_event.v2"}`),
		[]byte(strings.Replace(string(raw), `"phase":"completed"`, `"phase":"invented"`, 1)),
		[]byte(strings.Replace(string(raw), `"event_ts":"2026-08-22T04:00:00Z"`, `"event_ts":"no"`, 1)),
		[]byte(strings.Replace(string(raw), `"schema":"m3.app_event.v2"`, `"schema":"m3.app_event.v2","schema":"m3.app_event.v2"`, 1)),
		[]byte(strings.Replace(string(raw), `"event_ts":`, `"workspace_root_path":"/secret","event_ts":`, 1)),
		[]byte(strings.Replace(string(raw), `"event_ts":`, `"session_log_path":"/secret","event_ts":`, 1)),
	}
	for i, candidate := range bad {
		_, err := events.Persist(ctx, Request{EventID: fmt.Sprintf("bad-%d", i), RunID: "run", TurnID: "turn", ReportedBy: "m-9-worker/g1", Raw: candidate, At: 11})
		if err == nil {
			t.Fatalf("bad event %d accepted", i)
		}
	}
	count := 0
	_, err = host.Read(ctx, applier.QueryFunc(func(ctx context.Context, s *store.Snapshot) (any, error) {
		return nil, s.QueryRowContext(ctx, `SELECT COUNT(*) FROM pending_app_events`).Scan(&count)
	}))
	if err != nil || count != 1 {
		t.Fatalf("count=%d err=%v", count, err)
	}
}

func seedRun(t *testing.T, ctx context.Context, host *applier.Host) {
	t.Helper()
	_, err := host.Apply(ctx, eventFunc{func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, "run", []byte("{}"), strings.Repeat("0", 64), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, "run", fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0)); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, "turn", "run", fmt.Sprintf("%020d", 1), "ACTIVE", []byte("{}"), "fresh", strings.Repeat("a", 32), "PENDING", 1)
		return err
	}})
	if err != nil {
		t.Fatal(err)
	}
}

type eventFunc struct {
	fn func(context.Context, *store.Tx) error
}

func (eventFunc) RunID() string { return "run" }
func (e eventFunc) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, e.fn(ctx, tx)
}
