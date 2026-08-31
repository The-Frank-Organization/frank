package gc_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/config"
	frankgc "github.com/The-Frank-Organization/frank/internal/gc"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
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

func TestPassResumesCommittedMarkerWithoutDuplicateRecord(t *testing.T) {
	root := t.TempDir()
	st, journal := gcStoreWithSegments(t, root)
	segments, err := journal.Segments()
	if err != nil {
		t.Fatalf("Segments: %v", err)
	}
	if len(segments) < 2 {
		t.Fatalf("test setup did not rotate: %+v", segments)
	}
	for _, entry := range segments[0].Entries {
		commitOutcome(t, st, entry.IntakeID)
	}
	markerID := fmt.Sprintf("gc-%06d", segments[0].Seq)
	body, err := json.Marshal(struct {
		Segments []string `json:"segments"`
	}{Segments: []string{filepath.Base(segments[0].Path)}})
	if err != nil {
		t.Fatalf("marshal marker body: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       markerID,
			From:          "system",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"record_kind": "gc_marker"},
		Body:    string(body),
	}, nil); err != nil {
		t.Fatalf("precommit marker: %v", err)
	}
	tables, err := obligation.BuildTables(st)
	if err != nil {
		t.Fatalf("BuildTables: %v", err)
	}

	if err := frankgc.Pass(st, tables, config.EngineConfig{GCEnabled: true, SegmentRotateBytes: 96}); err != nil {
		t.Fatalf("Pass after marker commit: %v", err)
	}
	if _, err := os.Stat(segments[0].Path); !os.IsNotExist(err) {
		t.Fatalf("marked segment still exists after resumed pass: %v", err)
	}
	if count := countGCMarkers(t, st); count != 1 {
		t.Fatalf("gc markers = %d, want 1", count)
	}
}

func TestNoIDReuseAfterGCAndRestart(t *testing.T) {
	root := t.TempDir()
	st, journal := gcStoreWithSegments(t, root)
	segments, err := journal.Segments()
	if err != nil {
		t.Fatalf("Segments: %v", err)
	}
	if len(segments) < 2 {
		t.Fatalf("test setup did not rotate: %+v", segments)
	}
	var lastID string
	for _, segment := range segments {
		for _, entry := range segment.Entries {
			if entry.IntakeID > lastID {
				lastID = entry.IntakeID
			}
		}
	}
	for _, entry := range segments[0].Entries {
		commitOutcome(t, st, entry.IntakeID)
	}
	tables, err := obligation.BuildTables(st)
	if err != nil {
		t.Fatalf("BuildTables: %v", err)
	}
	if err := frankgc.Pass(st, tables, config.EngineConfig{GCEnabled: true, SegmentRotateBytes: 96}); err != nil {
		t.Fatalf("Pass: %v", err)
	}
	if _, err := os.Stat(segments[0].Path); !os.IsNotExist(err) {
		t.Fatalf("first segment still exists after GC: %v", err)
	}
	reopened, err := intake.OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: 96})
	if err != nil {
		t.Fatalf("reopen intake: %v", err)
	}
	next, err := reopened.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"after_gc":true}`), ContentHash: "after-gc"})
	if err != nil {
		t.Fatalf("Append after GC: %v", err)
	}
	if next <= lastID {
		t.Fatalf("next id after GC = %s, want greater than historical max %s", next, lastID)
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
