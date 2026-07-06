package zeroloss_test

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/test/replay/zeroloss"
)

func TestReplayConstructedStoreZeroLossIdentityAndCanonicalWins(t *testing.T) {
	if migrate.Current != 1 {
		t.Fatalf("migrate.Current = %d, want 1", migrate.Current)
	}
	root := constructedReplayStore(t)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	outboxProjectionPath := filepath.Join(root, "outbox", "gate-gate-1.json")
	corruptProjection(t, outboxProjectionPath)

	report, err := zeroloss.Replay(root, migrate.New())
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	wantTotal := len(recordFiles(t, root))
	if report.Total != wantTotal || report.Read != wantTotal || report.Lost != 0 || report.Quarantined != 0 {
		t.Fatalf("report = %+v, want all %d readable and zero lost/quarantined", report, wantTotal)
	}
	for relayID, view := range report.Views {
		if view.SourceVersion != migrate.Current {
			t.Fatalf("%s source version = %d, want current", relayID, view.SourceVersion)
		}
		stored, err := st.Read(relayID)
		if err != nil {
			t.Fatalf("Read stored %s: %v", relayID, err)
		}
		if !reflect.DeepEqual(view.Record, stored) {
			t.Fatalf("%s view != stored\nview=%+v\nstored=%+v", relayID, view.Record, stored)
		}
	}
	outboxView, ok := report.Views["outbox-gate-gate-1"]
	if !ok {
		t.Fatalf("missing canonical outbox record in replay views")
	}
	if outboxView.Record.Body == "corrupted projection" || !strings.Contains(outboxView.Record.Body, `"source_record_ref": "gate-1"`) {
		t.Fatalf("canonical outbox view body = %q", outboxView.Record.Body)
	}
	if got := string(mustReadFile(t, outboxProjectionPath)); got != "corrupted projection" {
		t.Fatalf("projection bytes = %q, want corrupted projection still on disk", got)
	}
	if _, ok := report.Views["config-change-1"]; !ok {
		t.Fatalf("config_change record missing from replay")
	}
}

func TestReplayOptionalArchiveStore(t *testing.T) {
	root := os.Getenv("FRANK_S5_REPLAY_STORE")
	if root == "" {
		t.Skip("FRANK_S5_REPLAY_STORE unset")
	}
	report, err := zeroloss.Replay(root, migrate.New())
	if err != nil {
		t.Fatalf("Replay archive: %v", err)
	}
	if report.Total == 0 || report.Read+report.Quarantined != report.Total || report.Lost != 0 {
		t.Fatalf("archive report = %+v, want self-derived readable/quarantined accounting", report)
	}
}

func TestReaderReadRefusalLegsThroughStore(t *testing.T) {
	for _, tc := range []struct {
		name string
		rec  record.Record
		reg  *migrate.Registry
		want error
	}{
		{
			name: "unversioned",
			rec:  record.Record{Envelope: record.Envelope{RelayID: "unversioned"}},
			reg:  migrate.New(),
			want: migrate.ErrUnversioned,
		},
		{
			name: "future",
			rec:  record.Record{Envelope: record.Envelope{RelayID: "future", SchemaVersion: migrate.Current + 1}},
			reg:  migrate.New(),
			want: migrate.ErrFutureVersion,
		},
		{
			name: "gap",
			rec:  record.Record{Envelope: record.Envelope{RelayID: "gap", SchemaVersion: 1}},
			reg: func() *migrate.Registry {
				reg := migrate.New()
				reg.Current = 2
				return reg
			}(),
			want: migrate.ErrMigrationGap,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			st, err := store.Open(t.TempDir())
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			plantSealedRecord(t, st.Root, tc.rec)
			_, err = (migrate.Reader{Store: st, Registry: tc.reg}).Read(tc.rec.Envelope.RelayID)
			if !errors.Is(err, tc.want) {
				t.Fatalf("Read error = %v, want %v", err, tc.want)
			}
		})
	}
}

func constructedReplayStore(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	if err := store.Init(root, replayConfigSources(t)); err != nil {
		t.Fatalf("Init: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitReplayRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "accepted-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "accepted"},
	})
	commitReplayRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "rejected-1", From: "operator", Role: "operator", DeliveryState: record.Rejected, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "rejected"},
		Body:     "PHASE:enum",
	})
	commitReplayRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "held-1", From: "system", Role: "system", DeliveryState: record.Held, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "held"},
	})
	commitReplayRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "gate-1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "gate", "HUMAN_GATE_REQUIRED": "yes", "gate_category": "authz_security"},
	})
	commitReplayRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "owed-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":            "SITREP",
			"AUTHORITY":        "report-only",
			"SUBJECT":          "owed",
			"record_kind":      "owed_item",
			"owner":            "s5",
			"source":           "replay",
			"target_surface":   "zeroloss",
			"disposition_path": "done",
		},
	})
	commitReplayRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "owed-disposition-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "disposed", "record_kind": "owed_disposition", "disposes_owed": "owed-1"},
	})
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}
	body := mustReadFile(t, filepath.Join(root, "config", "fieldspec", "registry.json"))
	rec := record.Record{
		Envelope: record.Envelope{RelayID: "config-change-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":       "SITREP",
			"SUBJECT":     "registry update",
			"record_kind": "config_change",
			"member":      "fieldspec",
			"new_digest":  digestWithReplayMember(t, root, "fieldspec", body),
		},
		Body: string(body),
	}
	if _, err := st.Commit(rec, store.ConfigChangeIntents(rec)); err != nil {
		t.Fatalf("Commit config_change: %v", err)
	}
	return root
}

func replayConfigSources(t *testing.T) map[string]string {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(root, "engine.json")
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	registryBytes := mustReadFile(t, filepath.Join("..", "..", "..", "internal", "fieldspec", "registry.json"))
	if err := os.WriteFile(registryPath, registryBytes, 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath}
}

func commitReplayRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func digestWithReplayMember(t *testing.T, root, member string, body []byte) string {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("Load config: %v", err)
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return config.Digest(members)
}

func plantSealedRecord(t *testing.T, root string, rec record.Record) {
	t.Helper()
	if rec.Headers == nil {
		rec.Headers = map[string]string{}
	}
	data, err := record.Seal(rec)
	if err != nil {
		t.Fatalf("Seal %s: %v", rec.Envelope.RelayID, err)
	}
	if err := os.WriteFile(filepath.Join(root, "records", rec.Envelope.RelayID+".json"), data, 0o644); err != nil {
		t.Fatalf("write planted record: %v", err)
	}
}

func recordFiles(t *testing.T, root string) []string {
	t.Helper()
	entries, err := os.ReadDir(filepath.Join(root, "records"))
	if err != nil {
		t.Fatalf("ReadDir records: %v", err)
	}
	var ids []string
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		ids = append(ids, strings.TrimSuffix(entry.Name(), ".json"))
	}
	return ids
}

func corruptProjection(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte("corrupted projection"), 0o644); err != nil {
		t.Fatalf("corrupt projection: %v", err)
	}
}

func mustReadFile(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return data
}
