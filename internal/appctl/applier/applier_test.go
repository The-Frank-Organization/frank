package applier

import (
	"context"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

func TestApplyCommitsAndAdvancesStateBeforeEmissionAndReply(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer db.Close()

	entered := make(chan struct{})
	release := make(chan struct{})
	emitted := make(chan Emission, 1)
	host := New(db, Config{Emitter: EmitterFunc(func(ctx context.Context, emission Emission) error {
		var state, stateSeq string
		if err := db.Read(ctx, func(snapshot *store.Snapshot) error {
			if err := snapshot.QueryRowContext(ctx, `SELECT state FROM runs WHERE run_id=?`, "run-1").Scan(&state); err != nil {
				return err
			}
			return snapshot.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, "run-1").Scan(&stateSeq)
		}); err != nil {
			return fmt.Errorf("read committed state from emitter: %w", err)
		}
		if state != "ADMITTED" || stateSeq != "00000000000000000001" {
			return fmt.Errorf("emitter observed state=%q seq=%q", state, stateSeq)
		}
		emitted <- emission
		return nil
	})})
	defer host.Close()

	done := make(chan Response, 1)
	go func() {
		result, err := host.Apply(ctx, testEvent{
			runID: "run-1",
			apply: func(ctx context.Context, tx *store.Tx) (Result, error) {
				if _, err := tx.ExecContext(ctx, `INSERT INTO runs
					(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at)
					VALUES(?, '{}', ?, 'ADMITTED', 'created', '00000000000000000000', 1)`, "run-1", strings.Repeat("0", 64)); err != nil {
					return Result{}, err
				}
				if _, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq)
					VALUES(?, '00000000000000000001', '00000000000000000000')`, "run-1"); err != nil {
					return Result{}, err
				}
				close(entered)
				<-release
				return Result{Value: "durable", Emissions: []Emission{{Kind: "turn_open", Value: "run-1"}}}, nil
			},
		})
		done <- Response{Result: result, Err: err}
	}()

	<-entered
	select {
	case got := <-emitted:
		t.Fatalf("emission visible before commit: %#v", got)
	default:
	}
	select {
	case got := <-done:
		t.Fatalf("reply visible before commit: %#v", got)
	default:
	}
	close(release)

	response := <-done
	if response.Err != nil {
		t.Fatalf("Apply: %v", response.Err)
	}
	if response.Result.Value != "durable" || response.Result.StateSeq != "1" {
		t.Fatalf("result = %#v", response.Result)
	}
	if got := <-emitted; got.Kind != "turn_open" || got.Value != "run-1" {
		t.Fatalf("emission = %#v", got)
	}
}

func TestConcurrentEventsAndSnapshotsUseOneSerializedLoop(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer db.Close()
	host := New(db, Config{})
	defer host.Close()

	if _, err := host.Apply(ctx, seedRunEvent("run-serial")); err != nil {
		t.Fatalf("seed run: %v", err)
	}

	var active atomic.Int32
	var maximum atomic.Int32
	const workers = 24
	var wait sync.WaitGroup
	for i := 0; i < workers; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			_, err := host.Apply(ctx, testEvent{runID: "run-serial", apply: func(ctx context.Context, tx *store.Tx) (Result, error) {
				now := active.Add(1)
				defer active.Add(-1)
				for previous := maximum.Load(); now > previous && !maximum.CompareAndSwap(previous, now); previous = maximum.Load() {
				}
				if _, err := tx.ExecContext(ctx, `UPDATE runs SET updated_at=coalesce(updated_at,0)+1 WHERE run_id=?`, "run-serial"); err != nil {
					return Result{}, err
				}
				return Result{}, nil
			}})
			if err != nil {
				t.Errorf("Apply: %v", err)
			}
		}()
	}
	wait.Wait()
	if got := maximum.Load(); got != 1 {
		t.Fatalf("maximum concurrent event applications = %d, want 1", got)
	}

	value, err := host.Read(ctx, QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var updates int
		var seq string
		if err := snapshot.QueryRowContext(ctx, `SELECT updated_at FROM runs WHERE run_id=?`, "run-serial").Scan(&updates); err != nil {
			return nil, err
		}
		if err := snapshot.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, "run-serial").Scan(&seq); err != nil {
			return nil, err
		}
		return fmt.Sprintf("%d/%s", updates, seq), nil
	}))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want := "24/00000000000000000025"; value != want {
		t.Fatalf("snapshot = %q, want %q", value, want)
	}
}

func TestTimerReentersTheApplierAsAnEvent(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer db.Close()
	host := New(db, Config{})
	defer host.Close()
	if _, err := host.Apply(ctx, seedRunEvent("run-timer")); err != nil {
		t.Fatalf("seed run: %v", err)
	}

	future := host.After(ctx, 10*time.Millisecond, testEvent{runID: "run-timer", apply: func(ctx context.Context, tx *store.Tx) (Result, error) {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET updated_at=7 WHERE run_id=?`, "run-timer")
		return Result{Value: "timer-fired"}, err
	}})
	select {
	case response := <-future:
		if response.Err != nil || response.Result.Value != "timer-fired" || response.Result.StateSeq != "2" {
			t.Fatalf("timer response = %#v", response)
		}
	case <-time.After(time.Second):
		t.Fatal("timer event did not re-enter applier")
	}
}

func TestFailedEventRollsBackStateSequenceAndEmitsNothing(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer db.Close()
	var emissions atomic.Int32
	host := New(db, Config{Emitter: EmitterFunc(func(context.Context, Emission) error {
		emissions.Add(1)
		return nil
	})})
	defer host.Close()
	if _, err := host.Apply(ctx, seedRunEvent("run-rollback")); err != nil {
		t.Fatalf("seed run: %v", err)
	}

	wantErr := errors.New("refuse transition")
	_, err = host.Apply(ctx, testEvent{runID: "run-rollback", apply: func(ctx context.Context, tx *store.Tx) (Result, error) {
		if _, err := tx.ExecContext(ctx, `UPDATE runs SET updated_at=99 WHERE run_id=?`, "run-rollback"); err != nil {
			return Result{}, err
		}
		return Result{Emissions: []Emission{{Kind: "must-not-emit"}}}, wantErr
	}})
	if !errors.Is(err, wantErr) {
		t.Fatalf("Apply error = %v, want %v", err, wantErr)
	}
	if got := emissions.Load(); got != 0 {
		t.Fatalf("emissions after rollback = %d, want 0", got)
	}

	value, err := host.Read(ctx, QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var updatedAt *int
		var seq string
		if err := snapshot.QueryRowContext(ctx, `SELECT updated_at FROM runs WHERE run_id=?`, "run-rollback").Scan(&updatedAt); err != nil {
			return nil, err
		}
		if err := snapshot.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, "run-rollback").Scan(&seq); err != nil {
			return nil, err
		}
		return fmt.Sprintf("%v/%s", updatedAt, seq), nil
	}))
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want := "<nil>/00000000000000000001"; value != want {
		t.Fatalf("rolled-back snapshot = %q, want %q", value, want)
	}
}

func TestAppControlTreeHasOneDatabaseWriterImport(t *testing.T) {
	root := filepath.Clean("..")
	fset := token.NewFileSet()
	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return err
		}
		for _, imported := range file.Imports {
			if imported.Path.Value == `"database/sql"` && filepath.Base(filepath.Dir(path)) != "store" {
				t.Errorf("database/sql imported outside store boundary: %s", path)
			}
		}
		if filepath.Base(path) == "applier.go" || filepath.Base(path) == "store_test.go" {
			return nil
		}
		full, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			return err
		}
		ast.Inspect(full, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			selector, ok := call.Fun.(*ast.SelectorExpr)
			if ok && selector.Sel.Name == "Update" {
				t.Errorf("store update call outside applier chokepoint: %s", path)
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("walk appctl tree: %v", err)
	}
}

type testEvent struct {
	runID string
	apply func(context.Context, *store.Tx) (Result, error)
}

func (event testEvent) RunID() string { return event.runID }

func (event testEvent) Apply(ctx context.Context, tx *store.Tx) (Result, error) {
	return event.apply(ctx, tx)
}

func seedRunEvent(runID string) Event {
	return testEvent{runID: runID, apply: func(ctx context.Context, tx *store.Tx) (Result, error) {
		if _, err := tx.ExecContext(ctx, `INSERT INTO runs
			(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at)
			VALUES(?, '{}', ?, 'ADMITTED', 'created', '00000000000000000000', 1)`, runID, strings.Repeat("0", 64)); err != nil {
			return Result{}, err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq)
			VALUES(?, '00000000000000000001', '00000000000000000000')`, runID)
		return Result{}, err
	}}
}
