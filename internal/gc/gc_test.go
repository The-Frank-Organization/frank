package gc_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	frankgc "github.com/jackli/frank/internal/gc"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestPassIsOffByDefault(t *testing.T) {
	root := t.TempDir()
	st, journal := gcStoreWithSegments(t, root)
	tables, err := obligation.BuildTables(st)
	if err != nil {
		t.Fatalf("BuildTables: %v", err)
	}
	before := existingSegments(t, journal)
	if err := frankgc.Pass(st, tables, config.EngineConfig{GCEnabled: false, SegmentRotateBytes: 96}); err != nil {
		t.Fatalf("Pass: %v", err)
	}
	after := existingSegments(t, journal)
	if strings.Join(before, ",") != strings.Join(after, ",") {
		t.Fatalf("segments changed with GC disabled: before %v after %v", before, after)
	}
	if countGCMarkers(t, st) != 0 {
		t.Fatalf("gc marker committed while disabled")
	}
}

func TestPassCollectsOnlyFullyDrainedNonActiveSegments(t *testing.T) {
	root := t.TempDir()
	st, journal := gcStoreWithSegments(t, root)
	segments, err := journal.Segments()
	if err != nil {
		t.Fatalf("Segments: %v", err)
	}
	if len(segments) < 2 || len(segments[0].Entries) == 0 {
		t.Fatalf("test setup did not rotate with entries: %+v", segments)
	}
	for _, entry := range segments[0].Entries {
		commitOutcome(t, st, entry.IntakeID)
	}
	tables, err := obligation.BuildTables(st)
	if err != nil {
		t.Fatalf("BuildTables: %v", err)
	}
	drained := frankgc.Drained(segments, tables)
	if len(drained) != 1 || drained[0].Seq != segments[0].Seq {
		t.Fatalf("Drained = %+v, want first segment only", drained)
	}

	if err := frankgc.Pass(st, tables, config.EngineConfig{GCEnabled: true, SegmentRotateBytes: 96}); err != nil {
		t.Fatalf("Pass: %v", err)
	}
	if _, err := os.Stat(segments[0].Path); !os.IsNotExist(err) {
		t.Fatalf("drained segment still exists: %v", err)
	}
	active := segments[len(segments)-1]
	if _, err := os.Stat(active.Path); err != nil {
		t.Fatalf("active segment removed or missing: %v", err)
	}
	marker := readGCMarker(t, st)
	if !strings.Contains(marker.Body, filepath.Base(segments[0].Path)) {
		t.Fatalf("marker body %q does not name collected segment %s", marker.Body, segments[0].Path)
	}
}

func gcStoreWithSegments(t *testing.T, root string) (*store.Store, *intake.Journal) {
	t.Helper()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	journal, err := intake.OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: 96})
	if err != nil {
		t.Fatalf("OpenWithConfig: %v", err)
	}
	for i := 0; i < 4; i++ {
		payload, _ := json.Marshal(map[string]int{"n": i})
		if _, err := journal.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: payload}); err != nil {
			t.Fatalf("Append %d: %v", i, err)
		}
	}
	return st, journal
}

func commitOutcome(t *testing.T, st *store.Store, intakeID string) {
	t.Helper()
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "outcome-" + intakeID, From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: intakeID, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "outcome"},
	}, nil); err != nil {
		t.Fatalf("Commit outcome: %v", err)
	}
}

func existingSegments(t *testing.T, journal *intake.Journal) []string {
	t.Helper()
	segments, err := journal.Segments()
	if err != nil {
		t.Fatalf("Segments: %v", err)
	}
	out := make([]string, 0, len(segments))
	for _, segment := range segments {
		if _, err := os.Stat(segment.Path); err == nil {
			out = append(out, filepath.Base(segment.Path))
		}
	}
	return out
}

func countGCMarkers(t *testing.T, st *store.Store) int {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var count int
	for _, rec := range records {
		if rec.Headers["record_kind"] == "gc_marker" {
			count++
		}
	}
	return count
}

func readGCMarker(t *testing.T, st *store.Store) record.Record {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	for _, rec := range records {
		if rec.Headers["record_kind"] == "gc_marker" {
			return rec
		}
	}
	t.Fatalf("gc marker not found")
	return record.Record{}
}
