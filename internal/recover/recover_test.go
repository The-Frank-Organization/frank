package recover_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/store"
)

func TestRunCleansStagingRebuildsProjectionsAndCompletesDerivedWork(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "staging", "torn.tmp"), []byte("torn"), 0o644); err != nil {
		t.Fatalf("write torn staging: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "gate-recover", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes"},
	}, []store.Intent{{Kind: store.IntentIndex, Path: "INDEX.md", Payload: []byte("| gate-recover |\n")}}); err != nil {
		t.Fatalf("commit gate: %v", err)
	}
	if err := os.Remove(filepath.Join(root, "projections", "INDEX.md")); err != nil {
		t.Fatalf("remove index: %v", err)
	}

	if err := frankrecover.Run(root); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "staging", "torn.tmp")); !os.IsNotExist(err) {
		t.Fatalf("torn staging still exists or stat failed differently: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "projections", "INDEX.md")); err != nil {
		t.Fatalf("index not rebuilt: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "outbox", "gate-gate-recover.json")); err != nil {
		t.Fatalf("derived outbox not completed: %v", err)
	}
}

func TestRunWithProcessorReplaysUnconsumedIntakeBeforeReturn(t *testing.T) {
	root := t.TempDir()
	j, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	if _, err := j.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: []byte(`{"n":1}`)}); err != nil {
		t.Fatalf("Append: %v", err)
	}

	var replayed []string
	if err := frankrecover.RunWithProcessor(root, func(cmd intake.Cmd) error {
		replayed = append(replayed, cmd.IntakeID)
		return nil
	}); err != nil {
		t.Fatalf("RunWithProcessor: %v", err)
	}
	if len(replayed) != 1 || replayed[0] != "intake-000001" {
		t.Fatalf("replayed = %v, want [intake-000001]", replayed)
	}
}
