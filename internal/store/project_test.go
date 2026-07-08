package store_test

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestProjectScopesMailboxAndReadReturnsCommittedRecord(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-a", DispatchID: "d", From: "seat-a", To: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "a"},
	}, []store.Intent{{Kind: store.IntentMailbox, Path: "seat-a.jsonl", Payload: []byte("relay-a\n")}}); err != nil {
		t.Fatalf("commit a: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-b", DispatchID: "d", From: "seat-b", To: "seat-b", Role: "planner", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "b"},
	}, []store.Intent{{Kind: store.IntentMailbox, Path: "seat-b.jsonl", Payload: []byte("relay-b\n")}}); err != nil {
		t.Fatalf("commit b: %v", err)
	}

	project, err := st.Project("seat-a")
	if err != nil {
		t.Fatalf("Project: %v", err)
	}
	if len(project) != 1 || project[0] != "relay-a" {
		t.Fatalf("project seat-a = %v, want [relay-a]", project)
	}
	rec, err := st.Read("relay-b")
	if err != nil {
		t.Fatalf("Read relay-b: %v", err)
	}
	if rec.Envelope.RelayID != "relay-b" || rec.Headers["SUBJECT"] != "b" {
		t.Fatalf("read record = %+v", rec)
	}
}

func TestPendingDeliveryForUsesDurableRecipientMailboxes(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if pending, err := st.PendingDeliveryFor("seat-b"); err != nil || pending {
		t.Fatalf("PendingDeliveryFor empty = %v, %v; want false, nil", pending, err)
	}

	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "relay-recipients", DispatchID: "d", From: "seat-a", To: "seat-b", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "recipient", "CC": `["seat-c"]`},
	}, nil); err != nil {
		t.Fatalf("commit recipients: %v", err)
	}

	for _, seat := range []string{"seat-b", "seat-c"} {
		pending, err := st.PendingDeliveryFor(seat)
		if err != nil {
			t.Fatalf("PendingDeliveryFor %s: %v", seat, err)
		}
		if !pending {
			t.Fatalf("PendingDeliveryFor %s = false, want true", seat)
		}
		project, err := st.Project(seat)
		if err != nil {
			t.Fatalf("Project %s: %v", seat, err)
		}
		if len(project) != 1 || project[0] != "relay-recipients" {
			t.Fatalf("Project %s = %v, want [relay-recipients]", seat, project)
		}
	}

	pending, err := st.PendingDeliveryFor("seat-a")
	if err != nil {
		t.Fatalf("PendingDeliveryFor sender: %v", err)
	}
	if pending {
		t.Fatalf("sender has pending delivery")
	}
}

func TestDeliveryRecipientsRequireCanonicalAddressLists(t *testing.T) {
	canonical := record.Record{
		Envelope: record.Envelope{To: "seat-b"},
		Headers:  map[string]string{"TO": `["seat-d"]`, "CC": `["seat-c","seat-b"]`},
	}
	got, err := store.DeliveryRecipients(canonical)
	if err != nil {
		t.Fatalf("DeliveryRecipients canonical: %v", err)
	}
	if want := []string{"seat-d", "seat-c", "seat-b"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("canonical recipients = %v, want %v", got, want)
	}

	nonCanonical := record.Record{
		Envelope: record.Envelope{To: "seat-b"},
		Headers:  map[string]string{"TO": `["seat-d"]`, "CC": `["seat-c", "seat-e"]`},
	}
	if _, err := store.DeliveryRecipients(nonCanonical); err == nil {
		t.Fatalf("DeliveryRecipients accepted non-canonical address_list")
	}
}

func TestProjectionRecipientTruthUsesCanonicalHeaderLists(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID:       "relay-full-set",
			DispatchID:    "dispatch-1",
			From:          "seat-a",
			To:            "compat-envelope-only",
			Role:          "planner",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":   "SITREP",
			"SUBJECT": "full set",
			"TO":      `["seat-b","seat-d"]`,
			"CC":      `["seat-c","seat-b"]`,
		},
		Body: "body",
	}
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("Commit: %v", err)
	}

	for _, seatName := range []string{"seat-b", "seat-c", "seat-d"} {
		data, err := os.ReadFile(filepath.Join(st.Root, "mailboxes", seatName+".jsonl"))
		if err != nil {
			t.Fatalf("mailbox %s: %v", seatName, err)
		}
		if strings.TrimSpace(string(data)) != "relay-full-set" {
			t.Fatalf("mailbox %s = %q, want relay-full-set", seatName, data)
		}
	}
	if _, err := os.Stat(filepath.Join(st.Root, "mailboxes", "compat-envelope-only.jsonl")); !os.IsNotExist(err) {
		t.Fatalf("compat envelope recipient got a mailbox, err=%v", err)
	}

	render, err := os.ReadFile(filepath.Join(st.Root, "projections", "relays", "dispatch-1", "SITREP-planner-relay-full-set.md"))
	if err != nil {
		t.Fatalf("render read: %v", err)
	}
	if !bytes.Contains(render, []byte("TO: seat-b, seat-d")) || !bytes.Contains(render, []byte("CC: seat-c, seat-b")) {
		t.Fatalf("render = %s, want decoded TO and CC lists", render)
	}

	index, err := os.ReadFile(filepath.Join(st.Root, "projections", "INDEX.md"))
	if err != nil {
		t.Fatalf("index read: %v", err)
	}
	if !bytes.Contains(index, []byte("| relay-full-set | SITREP | seat-a | seat-b, seat-d | seat-c, seat-b | accepted |")) {
		t.Fatalf("index = %s, want decoded TO/CC sets", index)
	}
}
