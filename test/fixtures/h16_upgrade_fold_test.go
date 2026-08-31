package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16UpgradeAnchorOnlyWithCompleteRawRedo(t *testing.T) {
	for _, tc := range []struct {
		name       string
		mutate     func(*testing.T, string)
		wantAnchor bool
	}{
		{name: "complete", wantAnchor: true},
		{name: "absent", mutate: func(t *testing.T, root string) { t.Helper(); _ = os.RemoveAll(filepath.Join(root, "journal", "redo")) }},
		{name: "partial latest", mutate: func(t *testing.T, root string) { h16FilterRedo(t, root, "legacy-b", false) }},
		{name: "partial older", mutate: func(t *testing.T, root string) { h16FilterRedo(t, root, "legacy-a", false) }},
		{name: "duplicate", mutate: func(t *testing.T, root string) { h16FilterRedo(t, root, "legacy-b", true) }},
		{name: "segment gap", mutate: func(t *testing.T, root string) {
			t.Helper()
			dir := filepath.Join(root, "journal", "redo")
			if err := os.Rename(filepath.Join(dir, "000001.jsonl"), filepath.Join(dir, "000002.jsonl")); err != nil {
				t.Fatalf("gap redo: %v", err)
			}
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			st, err := store.Open(root)
			if err != nil {
				t.Fatalf("open: %v", err)
			}
			h16Commit(t, st, h16LegacyPivot("legacy-a", "upgrade.implementer"))
			h16Commit(t, st, h16LegacyPivot("legacy-b", "upgrade.implementer"))
			if tc.mutate != nil {
				tc.mutate(t, root)
			}
			h16StartAndStopFrank(t, buildFrank(t, context.Background()), root, "upgrade-"+tc.name)
			anchors := h16Anchors(t, st, "upgrade.implementer")
			if tc.wantAnchor {
				if len(anchors) != 1 || anchors[0] != "legacy-b" {
					t.Fatalf("anchors=%v, want legacy-b", anchors)
				}
				h16StartAndStopFrank(t, buildFrank(t, context.Background()), root, "upgrade-restart-"+tc.name)
				if restarted := h16Anchors(t, st, "upgrade.implementer"); len(restarted) != 1 {
					t.Fatalf("restart double-anchored: %v", restarted)
				}
			} else if len(anchors) != 0 {
				t.Fatalf("deficient redo authored anchors=%v", anchors)
			}
		})
	}
}

func TestH16RawRedoMalformedAndTornRules(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	st, _ := store.Open(root)
	h16Commit(t, st, h16LegacyPivot("raw-a", "raw.implementer"))
	path := filepath.Join(root, "journal", "redo", "000001.jsonl")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read redo: %v", err)
	}
	if err := os.WriteFile(path, append(data, []byte(`{"broken":`)...), 0o644); err != nil {
		t.Fatalf("torn final: %v", err)
	}
	if snapshot := st.RawRedo(); !snapshot.Complete {
		t.Fatal("torn final entry should be ignored")
	}
	if err := os.WriteFile(path, append(data, []byte("not-json\n")...), 0o644); err != nil {
		t.Fatalf("malformed: %v", err)
	}
	if snapshot := st.RawRedo(); snapshot.Complete {
		t.Fatal("complete malformed entry proved redo complete")
	}
}

func h16LegacyPivot(relayID, seatName string) record.Record {
	return record.Record{
		Envelope: record.Envelope{RelayID: relayID, From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": relayID, "record_kind": "seat_mint"},
		Body:     `{"seat":"` + seatName + `","role":"implementer","is_operator":false}`,
	}
}

func h16FilterRedo(t *testing.T, root, relayID string, duplicate bool) {
	t.Helper()
	path := filepath.Join(root, "journal", "redo", "000001.jsonl")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read redo: %v", err)
	}
	var out []byte
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		var row struct {
			RelayID string `json:"relay_id"`
		}
		if json.Unmarshal(line, &row) != nil {
			continue
		}
		if row.RelayID == relayID && !duplicate {
			continue
		}
		out = append(out, line...)
		out = append(out, '\n')
		if row.RelayID == relayID && duplicate {
			out = append(out, line...)
			out = append(out, '\n')
		}
	}
	if err := os.WriteFile(path, out, 0o644); err != nil {
		t.Fatalf("write redo: %v", err)
	}
}

func h16Anchors(t *testing.T, st *store.Store, seatName string) []string {
	t.Helper()
	var selected []string
	for _, rec := range h16Records(t, st) {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "mint-chain-anchor" {
			continue
		}
		var body struct {
			Seat    string `json:"seat"`
			Selects string `json:"selects"`
		}
		_ = json.Unmarshal([]byte(rec.Body), &body)
		if body.Seat == seatName {
			selected = append(selected, body.Selects)
		}
	}
	return selected
}
