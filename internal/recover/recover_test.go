package recover_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/store"
)

func TestRunReturnsDiagnosticsForMissingGenesis(t *testing.T) {
	root := t.TempDir()
	pinned := pinnedForRecoverTest(t)

	result, err := frankrecover.Run(root, pinned)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if result.Ready != nil || result.Diag == nil {
		t.Fatalf("result = %+v, want diagnostics only", result)
	}
	report := result.Diag.Report()
	if !strings.Contains(report, "genesis-missing") {
		t.Fatalf("diagnostics report = %q, want genesis-missing", report)
	}
	if strings.Contains(report, root) {
		t.Fatalf("diagnostics report leaked store path: %q", report)
	}
}

func TestRunCleansStagingRebuildsProjectionsAndCompletesDerivedWork(t *testing.T) {
	root := t.TempDir()
	pinned := initRecoverStore(t, root)
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

	result, err := frankrecover.Run(root, pinned)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if result.Ready == nil || result.Diag != nil {
		t.Fatalf("result = %+v, want Ready", result)
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
	pinned := initRecoverStore(t, root)
	j, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	if _, err := j.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: []byte(`{"n":1}`)}); err != nil {
		t.Fatalf("Append: %v", err)
	}

	var replayed []string
	result, err := frankrecover.RunWithProcessor(root, pinned, func(cmd intake.Cmd) error {
		replayed = append(replayed, cmd.IntakeID)
		return nil
	})
	if err != nil {
		t.Fatalf("RunWithProcessor: %v", err)
	}
	if result.Ready == nil || result.Diag != nil {
		t.Fatalf("result = %+v, want Ready", result)
	}
	if len(replayed) != 1 || replayed[0] != "intake-000001" {
		t.Fatalf("replayed = %v, want [intake-000001]", replayed)
	}
}

func TestRunCompletesMarkedGCSegments(t *testing.T) {
	root := t.TempDir()
	pinned := initRecoverStoreWithEngine(t, root, `{"gc_enabled":true,"segment_rotate_bytes":96}`)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	journal, err := intake.OpenWithConfig(root, pinned.Engine)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	for i := 0; i < 4; i++ {
		if _, err := journal.Append(intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: []byte(fmt.Sprintf(`{"n":%d}`, i))}); err != nil {
			t.Fatalf("Append %d: %v", i, err)
		}
	}
	segments, err := journal.Segments()
	if err != nil {
		t.Fatalf("Segments: %v", err)
	}
	if len(segments) < 2 {
		t.Fatalf("test setup did not rotate: %+v", segments)
	}
	for _, entry := range segments[0].Entries {
		if _, err := st.Commit(record.Record{
			Envelope: record.Envelope{RelayID: "outcome-" + entry.IntakeID, From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: entry.IntakeID, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "outcome"},
		}, nil); err != nil {
			t.Fatalf("Commit outcome: %v", err)
		}
	}
	body, err := json.Marshal(struct {
		Segments []string `json:"segments"`
	}{Segments: []string{filepath.Base(segments[0].Path)}})
	if err != nil {
		t.Fatalf("marshal marker: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: fmt.Sprintf("gc-%06d", segments[0].Seq), From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "gc_marker"},
		Body:     string(body),
	}, nil); err != nil {
		t.Fatalf("Commit marker: %v", err)
	}

	result, err := frankrecover.Run(root, pinned)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if result.Ready == nil || result.Diag != nil {
		t.Fatalf("result = %+v, want Ready", result)
	}
	if _, err := os.Stat(segments[0].Path); !os.IsNotExist(err) {
		t.Fatalf("marked drained segment still exists after recovery: %v", err)
	}
}

func pinnedForRecoverTest(t *testing.T) *config.Pinned {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(root, "engine.json")
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	if err := os.WriteFile(registryPath, []byte(`{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`), 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	pinned, err := config.Load(map[string]string{"engine": enginePath, "fieldspec": registryPath})
	if err != nil {
		t.Fatalf("load pinned config: %v", err)
	}
	return pinned
}

func initRecoverStore(t *testing.T, root string) *config.Pinned {
	t.Helper()
	sources := writeRecoverConfigSources(t, `{"gc_enabled":false,"segment_rotate_bytes":4194304}`)
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("store Init: %v", err)
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store pinned config: %v", err)
	}
	return pinned
}

func initRecoverStoreWithEngine(t *testing.T, root, engineJSON string) *config.Pinned {
	t.Helper()
	sources := writeRecoverConfigSources(t, engineJSON)
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("store Init: %v", err)
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store pinned config: %v", err)
	}
	return pinned
}

func writeRecoverConfigSources(t *testing.T, engineJSON string) map[string]string {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(root, "engine.json")
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(enginePath, []byte(engineJSON), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	if err := os.WriteFile(registryPath, []byte(`{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`), 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath}
}
