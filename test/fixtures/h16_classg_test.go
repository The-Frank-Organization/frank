package fixtures_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestH16ClassGDirtyRetryBeforeReplay(t *testing.T) {
	loop, cancel := h16ClassGLoop(t)
	defer cancel()
	var attempts int
	loop.ClassGGC = func(*store.Store, *tables.T) error {
		attempts++
		if attempts <= 2 {
			return errors.New("injected gc failure")
		}
		return nil
	}

	cmd := intake.Cmd{IntakeID: "h16-classg-replay", Seat: "s12.implementer", Role: "implementer", Verb: "submit"}
	first := h16SubmitLoop(t, loop, cmd)
	second := h16SubmitLoop(t, loop, cmd)
	third := h16SubmitLoop(t, loop, cmd)
	assertH16PostCommitState(t, first, record.Accepted, "pending", false)
	assertH16PostCommitState(t, second, record.Accepted, "pending", false)
	assertH16PostCommitState(t, third, record.Accepted, "complete", true)
	if loop.ClassGDirty() || loop.Tables.ClassGDirty {
		t.Fatalf("Class-G remained dirty after successful retry: loop=%v tables=%v", loop.ClassGDirty(), loop.Tables.ClassGDirty)
	}
}

func TestH16ClassGAllFourSubstepsExposeDirtyAndPreServeHeal(t *testing.T) {
	tests := []struct {
		name      string
		configure func(*engine.Loop, *bool)
	}{
		{
			name: "CompleteAuto",
			configure: func(loop *engine.Loop, fail *bool) {
				loop.ClassGCompleteAuto = func(st *store.Store, tab *tables.T) error {
					if *fail {
						return errors.New("injected CompleteAuto failure")
					}
					return obligation.CompleteAuto(st, tab)
				}
			},
		},
		{
			name: "GC",
			configure: func(loop *engine.Loop, fail *bool) {
				loop.ClassGGC = func(*store.Store, *tables.T) error {
					if *fail {
						return errors.New("injected GC failure")
					}
					return nil
				}
			},
		},
		{
			name: "tables-build-publication",
			configure: func(loop *engine.Loop, fail *bool) {
				loop.ClassGBuildTables = func(st *store.Store) (*tables.T, error) {
					if *fail {
						return nil, errors.New("injected tables failure")
					}
					return tables.Build(st)
				}
			},
		},
		{
			name: "scheduler-arming",
			configure: func(loop *engine.Loop, fail *bool) {
				loop.ClassGArmScheduler = func() error {
					if *fail {
						return errors.New("injected scheduler failure")
					}
					return nil
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			loop, cancel := h16ClassGLoop(t)
			fail := true
			tc.configure(loop, &fail)
			out := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "h16-classg-" + tc.name, Seat: "s12.implementer", Role: "implementer", Verb: "submit"})
			assertH16PostCommitState(t, out, record.Accepted, "pending", false)
			if !loop.ClassGDirty() || !loop.Tables.ClassGDirty {
				t.Fatalf("dirty diagnostic absent: loop=%v tables=%v", loop.ClassGDirty(), loop.Tables.ClassGDirty)
			}
			cancel()

			fail = false
			restarted := engine.New(loop.Store, nil, engine.TestReady())
			tc.configure(restarted, &fail)
			if err := restarted.DrainClassG(); err != nil {
				t.Fatalf("pre-serve drain: %v", err)
			}
			if restarted.ClassGDirty() || restarted.Tables.ClassGDirty {
				t.Fatalf("pre-serve drain left dirty diagnostic set")
			}
		})
	}
}

func TestH16ProcessQuarantineFailureIsObservableAndNextEventHeals(t *testing.T) {
	loop, cancel := h16ClassGLoop(t)
	defer cancel()
	fail := true
	loop.ClassGGC = func(*store.Store, *tables.T) error {
		if fail {
			return errors.New("injected quarantine Class-G failure")
		}
		return nil
	}
	loop.EnqueueQuarantine("missing-relay")
	deadline := time.Now().Add(2 * time.Second)
	for !loop.ClassGDirty() && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if !loop.ClassGDirty() || !loop.Tables.ClassGDirty {
		t.Fatalf("quarantine Class-G failure was silent: loop=%v tables=%v", loop.ClassGDirty(), loop.Tables.ClassGDirty)
	}

	fail = false
	out := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "h16-after-quarantine", Seat: "s12.implementer", Role: "implementer", Verb: "submit"})
	assertH16PostCommitState(t, out, record.Accepted, "complete", true)
}

func TestH16PreServeDrainGatesRealHostAndConvertsPanicToNoServeError(t *testing.T) {
	loop := engine.New(mustH16Store(t), nil, engine.TestReady())
	loop.ClassGGC = func(*store.Store, *tables.T) error { panic("injected drain panic") }
	if err := loop.DrainClassG(); err == nil || !strings.Contains(err.Error(), "class-g-drain-panic") {
		t.Fatalf("DrainClassG panic err=%v, want class-g-drain-panic", err)
	}
	if !loop.ClassGDirty() || !loop.Tables.ClassGDirty {
		t.Fatal("drain panic did not publish dirty diagnostic")
	}

	mainBytes, err := os.ReadFile(filepath.Join("..", "..", "cmd", "frank", "main.go"))
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	drainAt := bytes.Index(mainBytes, []byte("loop.DrainClassG()"))
	runAt := bytes.Index(mainBytes, []byte("go loop.Run(ctx)"))
	serveAt := bytes.Index(mainBytes, []byte("channel.ServeAuthenticated("))
	if drainAt < 0 || runAt < 0 || serveAt < 0 || !(drainAt < runAt && drainAt < serveAt) {
		t.Fatalf("pre-serve ordering drain=%d loop=%d serve=%d", drainAt, runAt, serveAt)
	}
}

func TestH16NudgeRemainsMechanicalOnCompleteAcceptedStateOnly(t *testing.T) {
	mainBytes, err := os.ReadFile(filepath.Join("..", "..", "cmd", "frank", "main.go"))
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	if !bytes.Contains(mainBytes, []byte("out.State != record.Accepted")) {
		t.Fatal("delivery nudge no longer gates mechanically on complete accepted state")
	}
	if got := bytes.Count(mainBytes, []byte("deliveryNudgeFrame(out.RelayID)")); got != 1 {
		t.Fatalf("delivery nudge emission sites=%d, want one ordinary-reply site and no synthetic heal", got)
	}
}

func h16ClassGLoop(t *testing.T) (*engine.Loop, context.CancelFunc) {
	t.Helper()
	loop := engine.New(mustH16Store(t), func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "h16 Class-G"},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	go loop.Run(ctx)
	return loop, cancel
}

func mustH16Store(t *testing.T) *store.Store {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	return st
}

func h16SubmitLoop(t *testing.T, loop *engine.Loop, cmd intake.Cmd) engine.Outcome {
	t.Helper()
	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: cmd, ReplyCh: reply}
	return <-reply
}

func assertH16PostCommitState(t *testing.T, out engine.Outcome, decision, post string, wantLegacy bool) {
	t.Helper()
	if out.DecisionState != decision || out.PostCommitState != post {
		t.Fatalf("outcome=%+v, want decision=%s post=%s", out, decision, post)
	}
	if (out.State != "") != wantLegacy {
		t.Fatalf("outcome state=%q, want presence=%v; full=%+v", out.State, wantLegacy, out)
	}
}
