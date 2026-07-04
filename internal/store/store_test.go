package store_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestCommitShape(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	rec := record.Record{
		Envelope: record.Envelope{
			RelayID:       "relay-1",
			DispatchID:    "dispatch-1",
			From:          "seat-a.implementer",
			To:            "seat-b.planner",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "hello"},
		Body:    "hello body",
	}
	relayID, err := st.Commit(rec, []store.Intent{
		{Kind: store.IntentIndex, Path: "INDEX.md", Payload: []byte("| relay-1 | SITREP |\n")},
		{Kind: store.IntentRender, Path: "relays/dispatch-1/SITREP-implementer-relay-1.md", Payload: []byte("rendered relay-1\n")},
		{Kind: store.IntentMailbox, Path: "seat-b.planner.jsonl", Payload: []byte("relay-1\n")},
	})
	if err != nil {
		t.Fatalf("Commit: %v", err)
	}
	if relayID != "relay-1" {
		t.Fatalf("relayID = %q", relayID)
	}
	if _, err := record.Verify(mustRead(t, filepath.Join(root, "records", "relay-1.json"))); err != nil {
		t.Fatalf("record did not verify: %v", err)
	}
	assertFile(t, filepath.Join(root, "projections", "INDEX.md"), "| relay-1 | SITREP |\n")
	assertFile(t, filepath.Join(root, "projections", "relays", "dispatch-1", "SITREP-implementer-relay-1.md"), "rendered relay-1\n")
	assertFile(t, filepath.Join(root, "mailboxes", "seat-b.planner.jsonl"), "relay-1\n")
}

func TestCommitAutoRelayIDsDoNotCollide(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	first, err := st.Commit(record.Record{
		Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "first"},
	}, nil)
	if err != nil {
		t.Fatalf("first Commit: %v", err)
	}
	second, err := st.Commit(record.Record{
		Envelope: record.Envelope{From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "second"},
	}, nil)
	if err != nil {
		t.Fatalf("second Commit: %v", err)
	}
	if first == second {
		t.Fatalf("auto relay IDs collided: %q", first)
	}
	if _, err := os.Stat(filepath.Join(root, "records", first+".json")); err != nil {
		t.Fatalf("first record missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "records", second+".json")); err != nil {
		t.Fatalf("second record missing: %v", err)
	}
}

func TestRebuildProjectionsRestoresFromCanonicalAndRedo(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID:       "relay-2",
			DispatchID:    "dispatch-2",
			From:          "seat-a.implementer",
			To:            "seat-b.planner",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "hello"},
		Body:    "hello body",
	}
	intents := []store.Intent{
		{Kind: store.IntentIndex, Path: "INDEX.md", Payload: []byte("| relay-2 | SITREP |\n")},
		{Kind: store.IntentRender, Path: "relays/dispatch-2/SITREP-implementer-relay-2.md", Payload: []byte("rendered relay-2\n")},
		{Kind: store.IntentMailbox, Path: "seat-b.planner.jsonl", Payload: []byte("relay-2\n")},
	}
	if _, err := st.Commit(rec, intents); err != nil {
		t.Fatalf("Commit: %v", err)
	}

	indexPath := filepath.Join(root, "projections", "INDEX.md")
	renderPath := filepath.Join(root, "projections", "relays", "dispatch-2", "SITREP-implementer-relay-2.md")
	mailboxPath := filepath.Join(root, "mailboxes", "seat-b.planner.jsonl")
	if err := os.WriteFile(indexPath, []byte("corrupt row\n"), 0o644); err != nil {
		t.Fatalf("corrupt index: %v", err)
	}
	if err := os.Remove(renderPath); err != nil {
		t.Fatalf("remove render: %v", err)
	}
	if err := os.Remove(mailboxPath); err != nil {
		t.Fatalf("remove mailbox: %v", err)
	}

	if err := st.RebuildProjections(); err != nil {
		t.Fatalf("RebuildProjections: %v", err)
	}
	index := string(mustRead(t, indexPath))
	if !strings.Contains(index, "corrupt row\n") || !strings.Contains(index, "| relay-2 | SITREP |\n") {
		t.Fatalf("index did not preserve corrupt row and append correction: %q", index)
	}
	assertFile(t, renderPath, "rendered relay-2\n")
	assertFile(t, mailboxPath, "relay-2\n")
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return data
}

func assertFile(t *testing.T, path, want string) {
	t.Helper()
	got := string(mustRead(t, path))
	if got != want {
		t.Fatalf("%s = %q, want %q", path, got, want)
	}
}
