package fixtures_test

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
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
		filepath.Join("..", "..", "internal", "store", "projections.go"),
		filepath.Join("..", "..", "internal", "intake", "journal.go"),
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
		{name: "held-pre-rename", mutation: "held", crashpoint: "pre_rename"},
		{name: "operator-verdict-pre-rename", mutation: "operator-verdict", crashpoint: "pre_rename"},
		{name: "outbox-enqueue-pre-projection", mutation: "outbox-enqueue", crashpoint: "pre_projection_write"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
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
			if err := frankrecover.Run(root); err != nil {
				t.Fatalf("recover after crash: %v", err)
			}
			assertNoTornStaging(t, root)
			assertAllRecordsVerify(t, root)
			assertAtMostOneCanonicalRename(t, counter, tc.mutation)
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
	cmd.Env = append(os.Environ(), "FRANK_F9_CHILD=1", "FRANK_F9_ROOT="+root)
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
	for i := range 5 {
		payload, _ := json.Marshal(record.Record{Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "cmd"}})
		if _, err := journal.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: payload, ContentHash: "f9-" + string(rune('0'+i))}); err != nil {
			t.Fatalf("Append: %v", err)
		}
	}
	for i := 1; i <= 2; i++ {
		if _, err := st.Commit(record.Record{
			Envelope: record.Envelope{RelayID: "outcome-intake-00000" + string(rune('0'+i)), From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: "intake-00000" + string(rune('0'+i)), SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "outcome"},
		}, nil); err != nil {
			t.Fatalf("Commit outcome: %v", err)
		}
	}
	_ = syscall.Kill(os.Getpid(), syscall.SIGKILL)
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

func prepareF11Root(t *testing.T, root, mutation string) {
	t.Helper()
	if mutation != "operator-verdict" && mutation != "outbox-enqueue" {
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
	case "held":
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
			panic("validator fault")
		})
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
		handler := engine.SubmitHandler(st, reg, seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true})
		payload, _ := json.Marshal(record.Record{
			Headers: map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "verdict", "PARENT_DISPATCH_ID": "gate-base", "resolves_gate": "gate-base"},
		})
		rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "verdict-intake", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload})
		if err != nil {
			return err
		}
		_, err = st.Commit(rec, intents)
		return err
	case "outbox-enqueue":
		return gate.Complete(st)
	default:
		return nil
	}
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
