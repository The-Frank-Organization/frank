package terminal

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

func TestRegistryHasOnlyThreeMutatingVerbs(t *testing.T) {
	want := []string{"run cancel", "run start", "run stop"}
	if got := MutatingVerbs(); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("mutating verbs=%v", got)
	}
	for _, forbidden := range [][]string{{"parked", "clear"}, {"tickets", "void"}, {"disposition", "set"}, {"run", "resume"}} {
		fixture := newFixture(t, "registry-"+strings.Join(forbidden, "-"))
		before := fixture.stateSeq()
		stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
		code := fixture.runner.Execute(fixture.ctx, forbidden, stdout, stderr)
		if code == 0 || fixture.stateSeq() != before {
			t.Fatalf("forbidden %v code=%d seq=%s->%s", forbidden, code, before, fixture.stateSeq())
		}
	}
}

func TestPayloadFreeViewsAndPersistentFailureExit(t *testing.T) {
	fixture := newFixture(t, "views")
	fixture.seedRows()
	for _, command := range [][]string{{"status"}, {"attempts"}, {"tickets"}, {"parked"}, {"wakes"}} {
		stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
		code := fixture.runner.Execute(fixture.ctx, command, stdout, stderr)
		if command[0] == "status" && code != 1 {
			t.Fatalf("failed status exit=%d", code)
		}
		combined := stdout.String() + stderr.String()
		for _, secret := range []string{"credential-secret", "provider-payload", "/private/workspace", "/private/session"} {
			if strings.Contains(combined, secret) {
				t.Fatalf("%v leaked %q: %s", command, secret, combined)
			}
		}
	}
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	if code := fixture.runner.Execute(fixture.ctx, []string{"status"}, stdout, stderr); code != 1 || !strings.Contains(stderr.String(), "ALERT run-") {
		t.Fatalf("persistent alert code=%d stdout=%s stderr=%s", code, stdout, stderr)
	}
}

func TestStartStopCancelAreExplicit(t *testing.T) {
	fixture := newFixture(t, "mutations")
	starter := &fakeStarter{}
	fixture.runner = New(fixture.host, starter)
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	code := fixture.runner.Execute(fixture.ctx, []string{"run", "start", "--goal", "fix", "--lane", "lane", "--credential-ref", "opaque-ref", "--workspace-root", fixture.t.TempDir()}, stdout, stderr)
	if code != 0 || starter.calls != 1 {
		t.Fatalf("start code=%d calls=%d err=%s", code, starter.calls, stderr)
	}
	if code := fixture.runner.Execute(fixture.ctx, []string{"run", "stop", "--run-id", fixture.runID}, stdout, stderr); code != 0 {
		t.Fatalf("stop code=%d err=%s", code, stderr)
	}
	fixture.assertState("INTERRUPTED")
	fixture.resetActive()
	if code := fixture.runner.Execute(fixture.ctx, []string{"run", "cancel", "--run-id", fixture.runID, "--hard"}, stdout, stderr); code != 0 {
		t.Fatalf("cancel code=%d err=%s", code, stderr)
	}
	fixture.assertState("CANCELLED")
}

type fakeStarter struct{ calls int }

func (s *fakeStarter) Start(context.Context, StartRequest) error { s.calls++; return nil }

type fixture struct {
	t      *testing.T
	ctx    context.Context
	db     *store.Store
	host   *applier.Host
	runner *Runner
	runID  string
}

func newFixture(t *testing.T, suffix string) *fixture {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	f := &fixture{t: t, ctx: ctx, db: db, host: host, runID: "run-" + suffix}
	f.runner = New(host, nil)
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, f.runID, []byte(`{"credential_ref":"credential-secret","workspace_root_path":"/private/workspace"}`), strings.Repeat("0", 64), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, f.runID, fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0))
		return err
	})
	return f
}
func (f *fixture) seedRows() {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED',stop_reason='parked_unknown_capacity_exceeded' WHERE run_id=?`, f.runID); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, "turn", f.runID, fmt.Sprintf("%020d", 1), "INTERRUPTED", []byte("provider-payload"), "fresh", strings.Repeat("a", 32), "PENDING", 1); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO provider_attempts(attempt_id,run_id,turn_id,turn_epoch,state,logical_surface_digest,created_at) VALUES(?,?,?,?,?,?,?)`, "attempt", f.runID, "turn", fmt.Sprintf("%020d", 1), "UNKNOWN_PROVIDER_OUTCOME", strings.Repeat("a", 64), 1); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO tool_calls(tool_call_id,run_id,turn_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,created_at) VALUES(?,?,?,?,?,?,?,?)`, "call", f.runID, "turn", fmt.Sprintf("%020d", 1), "UNKNOWN_TOOL_OUTCOME", "read", strings.Repeat("b", 64), 1); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO tool_authorizations(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,canonical_tool_name,canonical_args_digest,effect_descriptor,issued_at) VALUES(?,?,?,?,?,?,?,?,?,?)`, "ticket", f.runID, "turn", "call", fmt.Sprintf("%020d", 1), "UNKNOWN_TOOL_OUTCOME", "read", strings.Repeat("b", 64), []byte("/private/session"), 1); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO wake_schedule(relay_id,run_id,disposition,received_at) VALUES(?,?,?,?)`, "relay", f.runID, "PENDING", 1)
		return err
	})
}
func (f *fixture) stateSeq() string {
	value := ""
	f.read(func(ctx context.Context, s *store.Snapshot) error {
		return s.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, f.runID).Scan(&value)
	})
	return value
}
func (f *fixture) assertState(want string) {
	f.t.Helper()
	got := ""
	f.read(func(ctx context.Context, s *store.Snapshot) error {
		return s.QueryRowContext(ctx, `SELECT state FROM runs WHERE run_id=?`, f.runID).Scan(&got)
	})
	if got != want {
		f.t.Fatalf("state=%s want=%s", got, want)
	}
}
func (f *fixture) resetActive() {
	f.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET state='ACTIVE',run_phase='established',stop_reason=NULL,resume_action=NULL WHERE run_id=?`, f.runID)
		return err
	})
}
func (f *fixture) mutate(fn func(context.Context, *store.Tx) error) {
	_, err := f.host.Apply(f.ctx, eventFunc{f.runID, fn})
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
func (f *fixture) read(fn func(context.Context, *store.Snapshot) error) {
	_, err := f.host.Read(f.ctx, applier.QueryFunc(func(ctx context.Context, s *store.Snapshot) (any, error) { return nil, fn(ctx, s) }))
	if err != nil {
		f.t.Fatal(err)
	}
}
