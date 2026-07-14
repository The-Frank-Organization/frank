package fixtures_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/gate"
	frankgc "github.com/jackli/frank/internal/gc"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	fixturedata "github.com/jackli/frank/test/fixtures"
)

func TestF11CrashpointRegistryCoversS1MutationBoundaries(t *testing.T) {
	names := crashpoint.Names()
	for _, want := range []string{
		"pre_record_fsync",
		"post_record_fsync",
		"pre_rename",
		"post_rename",
		"pre_dir_fsync",
		"post_dir_fsync",
		"pre_redo_fsync",
		"post_redo_fsync",
		"pre_projection_write",
		"post_projection_write",
		"pre_delivery_write",
		"post_delivery_write",
		"pre_outcome_reply",
	} {
		if !slices.Contains(names, want) {
			t.Fatalf("crashpoint registry missing %s in %v", want, names)
		}
	}
}

func TestF11CrashpointRegistryNamesAreLiveInSource(t *testing.T) {
	source := readSource(t,
		filepath.Join("..", "..", "internal", "fsio", "fsio.go"),
		filepath.Join("..", "..", "internal", "store", "store.go"),
		filepath.Join("..", "..", "internal", "store", "quarantine.go"),
		filepath.Join("..", "..", "internal", "store", "projections.go"),
		filepath.Join("..", "..", "internal", "intake", "journal.go"),
		filepath.Join("..", "..", "internal", "recover", "recover.go"),
		filepath.Join("..", "..", "internal", "obligation", "obligation.go"),
		filepath.Join("..", "..", "internal", "gc", "gc.go"),
		filepath.Join("..", "..", "internal", "engine", "loop.go"),
	)
	for _, name := range crashpoint.Names() {
		needle := `crashpoint.Hit("` + name + `")`
		if !strings.Contains(source, needle) {
			t.Fatalf("registered crashpoint %s has no live Hit site", name)
		}
	}
}

func TestF11CrashMatrixDrivesRealMutationsAndRecovery(t *testing.T) {
	if os.Getenv("FRANK_F11_CHILD") == "1" {
		runF11Child(t)
		return
	}
	cases := []struct {
		name       string
		mutation   string
		crashpoint string
	}{
		{name: "submit-accept-pre-rename", mutation: "submit-accept", crashpoint: "pre_rename"},
		{name: "submit-accept-post-rename", mutation: "submit-accept", crashpoint: "post_rename"},
		{name: "submit-accept-delivery", mutation: "submit-accept", crashpoint: "pre_delivery_write"},
		{name: "submit-reject-pre-rename", mutation: "submit-reject", crashpoint: "pre_rename"},
		{name: "config-change-pre-record-pivot", mutation: "config-change", crashpoint: "pre_rename"},
		{name: "config-change-post-record-pivot", mutation: "config-change", crashpoint: "post_rename"},
		{name: "config-change-pre-materialize", mutation: "config-change", crashpoint: "pre_projection_write"},
		{name: "config-change-post-materialize-index", mutation: "config-change", crashpoint: "post_projection_write"},
		{name: "config-change-pre-member-fsync", mutation: "config-change", crashpoint: "pre_record_fsync:2"},
		{name: "config-change-post-member-fsync", mutation: "config-change", crashpoint: "post_record_fsync:2"},
		{name: "config-change-pre-member-rename", mutation: "config-change", crashpoint: "pre_rename:2"},
		{name: "config-change-post-member-rename", mutation: "config-change", crashpoint: "post_rename:2"},
		{name: "config-change-pre-member-dir-fsync", mutation: "config-change", crashpoint: "pre_dir_fsync:2"},
		{name: "config-change-post-member-dir-fsync", mutation: "config-change", crashpoint: "post_dir_fsync:2"},
		{name: "held-pre-rename", mutation: "held", crashpoint: "pre_rename"},
		{name: "operator-verdict-pre-rename", mutation: "operator-verdict", crashpoint: "pre_rename"},
		{name: "outbox-enqueue-pre-projection", mutation: "outbox-enqueue", crashpoint: "pre_projection_write"},
		{name: "quarantine-pre-evict", mutation: "quarantine-disposition", crashpoint: "pre_quarantine_evict"},
		{name: "quarantine-post-evict", mutation: "quarantine-disposition", crashpoint: "post_quarantine_evict"},
		{name: "segment-rotate-pre", mutation: "segment-rotate", crashpoint: "pre_segment_rotate"},
		{name: "segment-rotate-post", mutation: "segment-rotate", crashpoint: "post_segment_rotate"},
		{name: "recovery-phase0", mutation: "genesis", crashpoint: "recovery_post_phase0"},
		{name: "recovery-phase1", mutation: "genesis", crashpoint: "recovery_post_phase1"},
		{name: "recovery-phase2", mutation: "genesis", crashpoint: "recovery_post_phase2"},
		{name: "recovery-phase3", mutation: "genesis", crashpoint: "recovery_post_phase3"},
		{name: "recovery-phase3-5", mutation: "genesis", crashpoint: "recovery_post_phase3_5"},
		{name: "recovery-phase3-6", mutation: "genesis", crashpoint: "recovery_post_phase3_6"},
		{name: "recovery-phase4", mutation: "genesis", crashpoint: "recovery_post_phase4"},
		{name: "gc-pre-marker", mutation: "gc-marker", crashpoint: "pre_gc_marker"},
		{name: "gc-post-marker", mutation: "gc-marker", crashpoint: "post_gc_marker"},
		{name: "gc-pre-unlink", mutation: "gc-marker", crashpoint: "pre_gc_unlink"},
		{name: "gc-post-unlink", mutation: "gc-marker", crashpoint: "post_gc_unlink"},
		{name: "owed-item-pre-rename", mutation: "owed-item", crashpoint: "pre_rename"},
		{name: "owed-disposition-pre-rename", mutation: "owed-disposition", crashpoint: "pre_rename"},
		{name: "seat-mint-pre-rename", mutation: "seat-mint", crashpoint: "pre_rename"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			counter := filepath.Join(root, "rename-counter.log")
			if err := os.WriteFile(counter, nil, 0o644); err != nil {
				t.Fatalf("create counter: %v", err)
			}
			prepareF11Root(t, root, tc.mutation)
			cmd := exec.Command(os.Args[0], "-test.run", "^TestF11CrashMatrixDrivesRealMutationsAndRecovery$")
			cmd.Env = append(os.Environ(),
				"FRANK_F11_CHILD=1",
				"FRANK_F11_ROOT="+root,
				"FRANK_F11_MUTATION="+tc.mutation,
				"FRANK_TEST_CRASHPOINT="+tc.crashpoint,
				"FRANK_TEST_RENAME_COUNTER="+counter,
			)
			err := cmd.Run()
			if err == nil {
				t.Fatalf("child completed without crash at %s", tc.crashpoint)
			}
			assertSIGKILL(t, err)
			result, err := frankrecover.Run(root, loadFixturePinned(t, root))
			if err != nil {
				t.Fatalf("recover after crash: %v", err)
			}
			if result.Ready == nil || result.Diag != nil {
				t.Fatalf("recover result = %+v, want Ready", result)
			}
			assertNoTornStaging(t, root)
			assertAllRecordsVerify(t, root)
			assertF11Converged(t, root, tc.mutation)
			if mutationHasSingleCanonicalRecord(tc.mutation) {
				assertAtMostOneCanonicalRename(t, counter, tc.mutation)
			}
		})
	}
}

func TestS2CleanCompletionClassesExecute(t *testing.T) {
	for _, mutation := range f11Classes() {
		t.Run(mutation, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			prepareF11Root(t, root, mutation)
			if err := runF11Mutation(root, mutation); err != nil {
				t.Fatalf("runF11Mutation(%s): %v", mutation, err)
			}
			assertF11Converged(t, root, mutation)
		})
	}
}

func TestS2ApplicabilityMapRowsMatchHitTrace(t *testing.T) {
	if os.Getenv("FRANK_F11_TRACE_CHILD") == "1" {
		runF11TraceChild(t)
		return
	}
	classMap := fixturedata.ApplicabilityMap()
	for _, mutation := range f11Classes() {
		t.Run(mutation, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			prepareF11Root(t, root, mutation)
			trace := filepath.Join(root, "hit-trace.log")
			cmd := exec.Command(os.Args[0], "-test.run", "^TestS2ApplicabilityMapRowsMatchHitTrace$")
			cmd.Env = append(os.Environ(),
				"FRANK_F11_TRACE_CHILD=1",
				"FRANK_F11_ROOT="+root,
				"FRANK_F11_MUTATION="+mutation,
				"FRANK_TEST_HIT_TRACE="+trace,
			)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("trace child failed: %v\n%s", err, out)
			}
			assertTraceMatchesApplicabilityRow(t, mutation, readHitTrace(t, trace), classMap[mutation])
		})
	}
}

func TestF9WholeAcrossRealCrashReenqueuesOnlyOutcomeLessIntake(t *testing.T) {
	if os.Getenv("FRANK_F9_CHILD") == "1" {
		runF9Child(t)
		return
	}
	root := t.TempDir()
	cmd := exec.Command(os.Args[0], "-test.run", "^TestF9WholeAcrossRealCrashReenqueuesOnlyOutcomeLessIntake$")
	cmd.Env = append(os.Environ(), "FRANK_F9_CHILD=1", "FRANK_F9_ROOT="+root, "FRANK_TEST_CRASHPOINT=post_intake_fsync:5")
	err := cmd.Run()
	if err == nil {
		t.Fatalf("F9 child completed without crash")
	}
	assertSIGKILL(t, err)

	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	unconsumed, err := intake.Unconsumed(context.Background(), journal, st)
	if err != nil {
		t.Fatalf("Unconsumed: %v", err)
	}
	if len(unconsumed) != 3 {
		t.Fatalf("unconsumed count = %d, want 3: %+v", len(unconsumed), unconsumed)
	}
	for i, cmd := range unconsumed {
		want := "intake-00000" + string(rune('3'+i))
		if cmd.IntakeID != want {
			t.Fatalf("unconsumed[%d] = %s, want %s", i, cmd.IntakeID, want)
		}
	}
}

func readSource(t *testing.T, paths ...string) string {
	t.Helper()
	var b strings.Builder
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		b.Write(data)
		b.WriteByte('\n')
	}
	return b.String()
}

func runF11Child(t *testing.T) {
	root := os.Getenv("FRANK_F11_ROOT")
	mutation := os.Getenv("FRANK_F11_MUTATION")
	if root == "" || mutation == "" {
		t.Fatalf("missing child env")
	}
	if err := runF11Mutation(root, mutation); err != nil {
		t.Fatalf("run mutation: %v", err)
	}
}

func runF11TraceChild(t *testing.T) {
	root := os.Getenv("FRANK_F11_ROOT")
	mutation := os.Getenv("FRANK_F11_MUTATION")
	if root == "" || mutation == "" || os.Getenv("FRANK_TEST_HIT_TRACE") == "" {
		t.Fatalf("missing trace child env")
	}
	if err := runF11Mutation(root, mutation); err != nil {
		t.Fatalf("run mutation: %v", err)
	}
	assertF11Converged(t, root, mutation)
}

func runF9Child(t *testing.T) {
	root := os.Getenv("FRANK_F9_ROOT")
	if root == "" {
		t.Fatalf("missing F9 root")
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	writer, err := intake.NewWriter[engine.Outcome](journal, config.EngineConfig{}, engine.TestReady())
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan intake.Job[engine.Outcome], 5)
	go writer.Run(ctx, out)
	for i := range 5 {
		payload, _ := json.Marshal(record.Record{Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "cmd"}})
		_, _, err := writer.Submit(ctx, intake.Cmd{Seat: "seat-" + string(rune('a'+i)), Role: "implementer", Verb: "submit", Payload: payload, ContentHash: "f9-" + string(rune('0'+i))})
		if err != nil {
			t.Fatalf("Submit: %v", err)
		}
		job := <-out
		if i < 2 {
			if _, err := st.Commit(record.Record{
				Envelope: record.Envelope{RelayID: "outcome-" + job.Cmd.IntakeID, From: job.Cmd.Seat, Role: "implementer", DeliveryState: record.Accepted, IntakeID: job.Cmd.IntakeID, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "outcome"},
			}, nil); err != nil {
				t.Fatalf("Commit outcome: %v", err)
			}
			job.ReplyCh <- engine.Outcome{State: record.Accepted, RelayID: "outcome-" + job.Cmd.IntakeID, IntakeID: job.Cmd.IntakeID}
		}
	}
}

func f11Classes() []string {
	return []string{
		"submit-accept",
		"submit-reject",
		"config-change",
		"held",
		"operator-verdict",
		"park",
		"outbox-enqueue",
		"genesis",
		"quarantine-disposition",
		"gc-marker",
		"owed-item",
		"owed-disposition",
		"seat-mint",
	}
}

func assertSIGKILL(t *testing.T, err error) {
	t.Helper()
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("child error = %T %v, want ExitError", err, err)
	}
	status, ok := exitErr.Sys().(syscall.WaitStatus)
	if !ok {
		t.Fatalf("child status = %T, want WaitStatus", exitErr.Sys())
	}
	if status.Signal() != syscall.SIGKILL {
		t.Fatalf("child signal = %s, want SIGKILL", status.Signal())
	}
}

func readHitTrace(t *testing.T, path string) map[string]bool {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read hit trace: %v", err)
	}
	out := map[string]bool{}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line != "" {
			out[line] = true
		}
	}
	return out
}

func assertTraceMatchesApplicabilityRow(t *testing.T, mutation string, fired map[string]bool, row map[string]string) {
	t.Helper()
	if row == nil {
		t.Fatalf("missing applicability row for %s", mutation)
	}
	var firedClean, expectedUnfired, unknown []string
	for name := range fired {
		if row[name] == "" {
			unknown = append(unknown, name)
		}
	}
	for _, name := range crashpoint.Names() {
		switch {
		case fired[name] && row[name] != "crash-expected":
			firedClean = append(firedClean, name)
		case !fired[name] && row[name] == "crash-expected":
			expectedUnfired = append(expectedUnfired, name)
		}
	}
	if len(firedClean) > 0 || len(expectedUnfired) > 0 || len(unknown) > 0 {
		slices.Sort(firedClean)
		slices.Sort(expectedUnfired)
		slices.Sort(unknown)
		t.Fatalf("%s applicability row mismatch:\nfired but labeled clean-completion: %v\ncrash-expected but not fired: %v\nunknown fired names: %v\nfired set: %v",
			mutation, firedClean, expectedUnfired, unknown, sortedSet(fired))
	}
}

func sortedSet(values map[string]bool) []string {
	out := make([]string, 0, len(values))
	for value := range values {
		out = append(out, value)
	}
	slices.Sort(out)
	return out
}

func prepareF11Root(t *testing.T, root, mutation string) {
	t.Helper()
	if mutation != "operator-verdict" && mutation != "outbox-enqueue" && mutation != "park" && mutation != "owed-disposition" {
		return
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	rec := record.Record{
		Envelope: record.Envelope{RelayID: "gate-base", From: "seat-a", To: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes"},
	}
	if mutation == "outbox-enqueue" {
		rec = record.Record{
			Envelope: record.Envelope{RelayID: "held-base", From: "system", Role: "system", DeliveryState: record.Held, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "held"},
		}
	}
	if mutation == "owed-disposition" {
		rec = record.Record{
			Envelope: record.Envelope{RelayID: "owed-base", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "owed", "record_kind": "owed_item", "owner": "s2", "source": "f11", "target_surface": "sweep", "disposition_path": "report"},
		}
	}
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("prepare commit: %v", err)
	}
}

func runF11Mutation(root, mutation string) error {
	st, err := store.Open(root)
	if err != nil {
		return err
	}
	switch mutation {
	case "submit-accept":
		rec := record.Record{
			Envelope: record.Envelope{RelayID: "submit-accept", From: "seat-a", To: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "accepted"},
		}
		_, err := st.Commit(rec, store.DefaultProjectionIntents(rec))
		return err
	case "submit-reject":
		rec := record.Record{
			Envelope: record.Envelope{RelayID: "submit-reject", From: "seat-a", Role: "implementer", DeliveryState: record.Rejected, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "rejected"},
			Body:     "PHASE:enum",
		}
		_, err := st.Commit(rec, nil)
		return err
	case "config-change":
		reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
		if err != nil {
			return err
		}
		meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
		handler := engine.SubmitHandler(st, reg, meta)
		cand, err := f11ConfigChangeRecord(root)
		if err != nil {
			return err
		}
		payload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, cand))
		rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "config-change-intake", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload})
		if err != nil {
			return err
		}
		if rec.Envelope.DeliveryState != record.Accepted {
			return errors.New("config-change rejected: " + rec.Body)
		}
		_, err = st.Commit(rec, intents)
		return err
	case "held":
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
			panic("validator fault")
		}, engine.TestReady())
		go loop.Run(ctx)
		payload, _ := json.Marshal(record.Record{Headers: map[string]string{"PHASE": "PLAN", "SUBJECT": "held"}})
		reply := make(chan engine.Outcome, 1)
		loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "held-intake", Seat: "s1.orchestrator-planner", Role: "orchestrator-planner", Payload: payload}, ReplyCh: reply}
		<-reply
		return nil
	case "operator-verdict":
		reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
		if err != nil {
			return err
		}
		meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
		handler := engine.SubmitHandler(st, reg, meta)
		payload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{
			Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "verdict", "parent_hint": "gate-base", "resolves_gate": "gate-base"},
		}))
		rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "verdict-intake", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload})
		if err != nil {
			return err
		}
		_, err = st.Commit(rec, intents)
		return err
	case "seat-mint":
		reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
		if err != nil {
			return err
		}
		meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
		handler := engine.SubmitHandler(st, reg, meta)
		payload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{
			Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "seat mint", "record_kind": "seat_mint"},
			Body:    `{"seat":"f11-seat.implementer","role":"implementer","is_operator":false}`,
		}))
		rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "seat-mint-intake", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload})
		if err != nil {
			return err
		}
		if rec.Envelope.DeliveryState != record.Accepted {
			return errors.New("seat-mint rejected: " + rec.Body)
		}
		_, err = st.Commit(rec, intents)
		return err
	case "outbox-enqueue":
		return gate.Complete(st)
	case "park":
		return gate.Complete(st)
	case "genesis":
		pinned, err := config.Load(store.StoreRootConfigPaths(root))
		if err != nil {
			return err
		}
		_, err = frankrecover.Run(root, pinned)
		return err
	case "quarantine-disposition":
		if _, err := st.Commit(record.Record{
			Envelope: record.Envelope{RelayID: "corrupt-me", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "corrupt"},
		}, nil); err != nil && !strings.Contains(err.Error(), "record already exists") {
			return err
		}
		recordPath := filepath.Join(root, "records", "corrupt-me.json")
		if err := os.WriteFile(recordPath, []byte(`{"bad":true}`), 0o644); err != nil {
			return err
		}
		_, err := st.ScanQuarantine()
		if err != nil {
			return err
		}
		return st.CompleteIncidents()
	case "segment-rotate":
		journal, err := intake.OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: 96})
		if err != nil {
			return err
		}
		for i := 0; i < 4; i++ {
			if _, err := journal.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"segment":true}`), ContentHash: "segment-" + string(rune('0'+i))}); err != nil {
				return err
			}
		}
		return nil
	case "gc-marker":
		journal, err := intake.OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: 96})
		if err != nil {
			return err
		}
		for i := 0; i < 4; i++ {
			if _, err := journal.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"gc":true}`), ContentHash: "gc-" + string(rune('0'+i))}); err != nil {
				return err
			}
		}
		segments, err := journal.Segments()
		if err != nil {
			return err
		}
		if len(segments) < 2 {
			return nil
		}
		for _, entry := range segments[0].Entries {
			if _, err := st.Commit(record.Record{
				Envelope: record.Envelope{RelayID: "outcome-" + entry.IntakeID, From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: entry.IntakeID, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "outcome"},
			}, nil); err != nil && !strings.Contains(err.Error(), "record already exists") {
				return err
			}
		}
		tables, err := obligation.BuildTables(st)
		if err != nil {
			return err
		}
		return frankgc.Pass(st, tables, config.EngineConfig{GCEnabled: true, SegmentRotateBytes: 96})
	case "owed-item":
		return commitOwedMutation(st, "owed-item", "")
	case "owed-disposition":
		return commitOwedMutation(st, "owed-disposition", "owed-base")
	default:
		return errors.New("unknown F11 mutation: " + mutation)
	}
}

const f11ConfigChangeMarker = "f11-config-change"

func f11ConfigChangeRecord(root string) (record.Record, error) {
	body, err := f11MutatedRegistryBody(root)
	if err != nil {
		return record.Record{}, err
	}
	digest, err := f11DigestWithMember(root, "fieldspec", body)
	if err != nil {
		return record.Record{}, err
	}
	return record.Record{
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"AUTHORITY":     "report-only",
			"CEREMONY_TIER": "medium",
			"SUBJECT":       "F11 registry config_change",
			"record_kind":   "config_change",
			"member":        "fieldspec",
			"new_digest":    digest,
		},
		Body: string(body),
	}, nil
}

func f11MutatedRegistryBody(root string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(root, "config", "fieldspec", "registry.json"))
	if err != nil {
		return nil, err
	}
	var doc map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	provenance, ok := doc["provenance"].(map[string]any)
	if !ok {
		provenance = map[string]any{}
		doc["provenance"] = provenance
	}
	provenance["s4_test_marker"] = f11ConfigChangeMarker
	return json.Marshal(doc)
}

func f11DigestWithMember(root, member string, body []byte) (string, error) {
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		return "", err
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return config.Digest(members), nil
}

func commitOwedMutation(st *store.Store, kind, parent string) error {
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		return err
	}
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	handler := engine.SubmitHandler(st, reg, meta)
	recordKind := strings.ReplaceAll(kind, "-", "_")
	headers := map[string]string{
		"PHASE":            "SITREP",
		"AUTHORITY":        "report-only",
		"SUBJECT":          kind,
		"record_kind":      recordKind,
		"owner":            "s2",
		"source":           "f11",
		"target_surface":   "sweep",
		"disposition_path": "report",
	}
	if kind == "owed-disposition" {
		headers = map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": kind, "record_kind": recordKind, "disposes_owed": parent}
	}
	payload, _ := json.Marshal(submitPayloadForRegistry(reg, meta, record.Record{Headers: headers}))
	rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: kind + "-intake", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload})
	if err != nil {
		return err
	}
	_, err = st.Commit(rec, intents)
	return err
}

func submitPayloadForRegistry(reg *fieldspec.Registry, meta seat.SeatMeta, rec record.Record) fieldspec.SubmitPayload {
	if rec.Headers != nil && rec.Headers["PHASE"] != "" && rec.Headers["EVIDENCE_TARGET"] == "" {
		rec.Headers["EVIDENCE_TARGET"] = "E1"
	}
	_, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	return fieldspec.SubmitPayload{Record: rec, FormDigest: digest}
}

func assertNoTornStaging(t *testing.T, root string) {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(root, "staging"))
	if err != nil {
		t.Fatalf("ReadDir staging: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("staging not clean after recovery: %v", entries)
	}
}

func assertAllRecordsVerify(t *testing.T, root string) {
	t.Helper()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records did not verify after recovery: %v", err)
	}
	seen := map[string]struct{}{}
	for _, rec := range records {
		if _, ok := seen[rec.Envelope.RelayID]; ok {
			t.Fatalf("duplicate relay id after recovery: %s", rec.Envelope.RelayID)
		}
		seen[rec.Envelope.RelayID] = struct{}{}
	}
}

func assertF11Converged(t *testing.T, root, mutation string) {
	t.Helper()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if mutation == "gc-marker" {
		tables, err := obligation.BuildTables(st)
		if err != nil {
			t.Fatalf("BuildTables: %v", err)
		}
		if err := frankgc.Pass(st, tables, config.EngineConfig{GCEnabled: true, SegmentRotateBytes: 96}); err != nil {
			t.Fatalf("GC convergence pass: %v", err)
		}
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records did not verify after %s: %v", mutation, err)
	}
	if mutation == "config-change" {
		assertF11ConfigChangeConverged(t, root, records)
	}
}

func assertF11ConfigChangeConverged(t *testing.T, root string, records []record.Record) {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read materialized registry: %v", err)
	}
	var cfg record.Record
	var found bool
	for _, rec := range records {
		if rec.Headers["record_kind"] == "config_change" && rec.Envelope.DeliveryState == record.Accepted {
			cfg = rec
			found = true
		}
	}
	if !found {
		if strings.Contains(string(body), f11ConfigChangeMarker) {
			t.Fatalf("config_change marker materialized without an accepted config_change record")
		}
		return
	}
	if string(body) != cfg.Body {
		t.Fatalf("materialized config does not match config_change record body")
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load materialized config: %v", err)
	}
	if pinned.Digest != cfg.Headers["new_digest"] {
		t.Fatalf("materialized digest = %s, want %s", pinned.Digest, cfg.Headers["new_digest"])
	}
}

func assertAtMostOneCanonicalRename(t *testing.T, counter, mutation string) {
	t.Helper()
	data, err := os.ReadFile(counter)
	if err != nil {
		t.Fatalf("rename counter missing for %s: %v", mutation, err)
	}
	var recordRenames int
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if strings.HasPrefix(line, "records/") {
			recordRenames++
		}
	}
	if recordRenames > 1 {
		t.Fatalf("%s canonical record renames = %d, want <= 1; counter:\n%s", mutation, recordRenames, data)
	}
}

func mutationHasSingleCanonicalRecord(mutation string) bool {
	switch mutation {
	case "submit-accept", "submit-reject", "config-change", "held", "operator-verdict", "outbox-enqueue", "owed-item", "owed-disposition", "seat-mint":
		return true
	default:
		return false
	}
}
