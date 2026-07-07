package fixtures_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestS6ProjectionPollutedArchivedMailboxFilteredAndRebuilt(t *testing.T) {
	st := s6ProjectionStore(t)
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-accepted", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "accepted"},
	})
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-rejected", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Rejected, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "rejected"},
		Body:     "rejected detail",
	})
	s6WriteMailbox(t, st, "seat-b", "relay-accepted\nrelay-rejected\n")

	got, err := st.Project("seat-b")
	if err != nil {
		t.Fatalf("Project: %v", err)
	}
	if want := []string{"relay-accepted"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("default project = %v, want %v", got, want)
	}
	if err := st.RebuildProjections(); err != nil {
		t.Fatalf("RebuildProjections: %v", err)
	}
	s6AssertFile(t, filepath.Join(st.Root, "mailboxes", "seat-b.jsonl"), "relay-accepted\n")
}

func TestS6ProjectionAuditViewReturnsOwnAttemptsAndOwnRejectReadable(t *testing.T) {
	st := s6ProjectionStore(t)
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-own-accepted", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "own accepted"},
	})
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-own-rejected", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Rejected, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "own rejected"},
		Body:     "rejected detail",
	})
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-other", From: "seat-c", To: "seat-a", Role: "planner", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "other"},
	})

	got, err := st.ProjectView("seat-a", "audit")
	if err != nil {
		t.Fatalf("ProjectView audit: %v", err)
	}
	if want := []string{"relay-own-accepted", "relay-own-rejected"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("audit project = %v, want %v", got, want)
	}
	rec, err := st.Read("relay-own-rejected")
	if err != nil {
		t.Fatalf("Read rejected own attempt: %v", err)
	}
	if rec.Body != "rejected detail" {
		t.Fatalf("rejected body = %q", rec.Body)
	}
}

func TestS6ProjectionAcceptedDeliveryUnaffected(t *testing.T) {
	st := s6ProjectionStore(t)
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-delivered", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "delivered"},
	})
	got, err := st.Project("seat-b")
	if err != nil {
		t.Fatalf("Project: %v", err)
	}
	if want := []string{"relay-delivered"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("default project = %v, want %v", got, want)
	}
}

func TestS6ProjectionRejectNeverWokenOnAnyInterleaving(t *testing.T) {
	for name, relays := range map[string]string{
		"reject first":  "relay-rejected\nrelay-accepted\n",
		"reject last":   "relay-accepted\nrelay-rejected\n",
		"reject middle": "relay-accepted\nrelay-rejected\nrelay-accepted\n",
	} {
		t.Run(name, func(t *testing.T) {
			st := s6ProjectionStore(t)
			s6CommitProjectionRecord(t, st, record.Record{
				Envelope: record.Envelope{RelayID: "relay-accepted", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "accepted"},
			})
			s6CommitProjectionRecord(t, st, record.Record{
				Envelope: record.Envelope{RelayID: "relay-rejected", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Rejected, SchemaVersion: 1},
				Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "rejected"},
			})
			s6WriteMailbox(t, st, "seat-b", relays)

			got, err := st.Project("seat-b")
			if err != nil {
				t.Fatalf("Project: %v", err)
			}
			if want := []string{"relay-accepted"}; !reflect.DeepEqual(got, want) {
				t.Fatalf("default project = %v, want %v", got, want)
			}
		})
	}
}

func TestS6ProjectionHeldOffSeatDefault(t *testing.T) {
	st := s6ProjectionStore(t)
	s6CommitProjectionRecord(t, st, record.Record{
		Envelope: record.Envelope{RelayID: "relay-held", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Held, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "held"},
	})
	s6WriteMailbox(t, st, "seat-b", "relay-held\n")

	got, err := st.Project("seat-b")
	if err != nil {
		t.Fatalf("Project: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("default project = %v, want empty", got)
	}
}

func s6ProjectionStore(t *testing.T) *store.Store {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	return st
}

func s6CommitProjectionRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func s6WriteMailbox(t *testing.T, st *store.Store, seatName, content string) {
	t.Helper()
	path := filepath.Join(st.Root, "mailboxes", seatName+".jsonl")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir mailbox: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write mailbox: %v", err)
	}
}

func s6AssertFile(t *testing.T, path, want string) {
	t.Helper()
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if string(got) != want {
		t.Fatalf("%s = %q, want %q", path, got, want)
	}
}
