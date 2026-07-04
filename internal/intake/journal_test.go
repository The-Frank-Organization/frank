package intake_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestAppendAssignsIDAndDedupesContentHash(t *testing.T) {
	j, err := intake.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	cmd := intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"subject":"one"}`)}
	first, err := j.Append(cmd)
	if err != nil {
		t.Fatalf("Append first: %v", err)
	}
	second, err := j.Append(cmd)
	if err != nil {
		t.Fatalf("Append duplicate: %v", err)
	}
	if first != second {
		t.Fatalf("duplicate content got new intake id %q, want %q", second, first)
	}
	entries, err := j.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(entries))
	}
	if entries[0].IntakeID == "" || entries[0].ContentHash == "" {
		t.Fatalf("entry missing intake id or content hash: %+v", entries[0])
	}
}

func TestUnconsumedSkipsExistingOutcomes(t *testing.T) {
	root := t.TempDir()
	j, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open journal: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	first, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":1}`)})
	second, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":2}`)})
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "outcome-1",
			DispatchID:    "d",
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			IntakeID:      first,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "one"},
	}, nil); err != nil {
		t.Fatalf("commit outcome: %v", err)
	}
	got, err := intake.Unconsumed(context.Background(), j, st)
	if err != nil {
		t.Fatalf("Unconsumed: %v", err)
	}
	if len(got) != 1 || got[0].IntakeID != second {
		t.Fatalf("unconsumed = %+v, want only %s", got, second)
	}
}

func TestAppendRotatesAndReadsSegmentsInOrder(t *testing.T) {
	root := t.TempDir()
	j, err := intake.OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: 96})
	if err != nil {
		t.Fatalf("OpenWithConfig: %v", err)
	}
	for i := 0; i < 5; i++ {
		if _, err := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":` + string(rune('0'+i)) + `}`)}); err != nil {
			t.Fatalf("Append %d: %v", i, err)
		}
	}
	segments, err := j.Segments()
	if err != nil {
		t.Fatalf("Segments: %v", err)
	}
	if len(segments) < 2 {
		t.Fatalf("segments = %d, want rotation", len(segments))
	}
	var gotIDs []string
	for _, segment := range segments {
		if filepath.Base(segment.Path) != formatSegmentName(segment.Seq) {
			t.Fatalf("segment path %q does not match seq %d", segment.Path, segment.Seq)
		}
		for _, entry := range segment.Entries {
			gotIDs = append(gotIDs, entry.IntakeID)
		}
	}
	wantIDs := []string{"intake-000001", "intake-000002", "intake-000003", "intake-000004", "intake-000005"}
	if strings.Join(gotIDs, ",") != strings.Join(wantIDs, ",") {
		t.Fatalf("ids across segments = %v, want %v", gotIDs, wantIDs)
	}
	all, err := j.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if len(all) != len(wantIDs) {
		t.Fatalf("ReadAll entries = %d, want %d", len(all), len(wantIDs))
	}
}

func TestOpenRejectsLegacySingleIntakeFile(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "journal"), 0o755); err != nil {
		t.Fatalf("mkdir journal: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "journal", "intake.jsonl"), []byte("{}\n"), 0o644); err != nil {
		t.Fatalf("write legacy intake: %v", err)
	}
	_, err := intake.Open(root)
	if !errors.Is(err, intake.ErrLegacyJournal) {
		t.Fatalf("Open err = %v, want ErrLegacyJournal", err)
	}
}

func TestUnconsumedPreservesArrivalOrderAcrossSegments(t *testing.T) {
	root := t.TempDir()
	j, err := intake.OpenWithConfig(root, config.EngineConfig{SegmentRotateBytes: 96})
	if err != nil {
		t.Fatalf("OpenWithConfig: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	first, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":1}`)})
	second, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":2}`)})
	third, _ := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":3}`)})
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "outcome-2", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: second, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "two"},
	}, nil); err != nil {
		t.Fatalf("commit outcome: %v", err)
	}
	got, err := intake.Unconsumed(context.Background(), j, st)
	if err != nil {
		t.Fatalf("Unconsumed: %v", err)
	}
	if len(got) != 2 || got[0].IntakeID != first || got[1].IntakeID != third {
		t.Fatalf("unconsumed = %+v, want %s then %s", got, first, third)
	}
}

func TestReadAllIgnoresTornTrailingLine(t *testing.T) {
	root := t.TempDir()
	j, err := intake.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	first, err := j.Append(intake.Cmd{Seat: "seat-a", Verb: "submit", Payload: json.RawMessage(`{"n":1}`)})
	if err != nil {
		t.Fatalf("Append: %v", err)
	}
	segmentPath := filepath.Join(root, "journal", "intake", "000001.jsonl")
	f, err := os.OpenFile(segmentPath, os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		t.Fatalf("open segment: %v", err)
	}
	if _, err := f.WriteString(`{"intake_id":"intake-torn"`); err != nil {
		_ = f.Close()
		t.Fatalf("write torn tail: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("close torn writer: %v", err)
	}
	entries, err := j.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if len(entries) != 1 || entries[0].IntakeID != first {
		t.Fatalf("entries = %+v, want only %s", entries, first)
	}
}

func formatSegmentName(seq int) string {
	return "00000" + string(rune('0'+seq)) + ".jsonl"
}
