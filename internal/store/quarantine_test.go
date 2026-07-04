package store_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestScanQuarantineEvictsCorruptRecordAndCompletesIncident(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitRecord(t, st, "clean", "clean body")
	commitRecord(t, st, "corrupt", "before corruption")
	corruptRecordBody(t, root, "corrupt", "before corruption", "after corruption")

	if _, err := st.Records(); !isChecksumFor(err, "corrupt") {
		t.Fatalf("Records err = %v, want ErrChecksum for corrupt", err)
	}
	if _, err := st.Read("corrupt"); !isChecksumFor(err, "corrupt") {
		t.Fatalf("Read corrupt err = %v, want ErrChecksum", err)
	}

	evicted, err := st.ScanQuarantine()
	if err != nil {
		t.Fatalf("ScanQuarantine: %v", err)
	}
	if len(evicted) != 1 || evicted[0] != "corrupt" {
		t.Fatalf("evicted = %v, want [corrupt]", evicted)
	}
	if _, err := os.Stat(filepath.Join(root, "records", "corrupt.json")); !os.IsNotExist(err) {
		t.Fatalf("corrupt record still in records/: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "quarantine", "corrupt.json")); err != nil {
		t.Fatalf("quarantined record missing: %v", err)
	}
	if err := st.CompleteIncidents(); err != nil {
		t.Fatalf("CompleteIncidents: %v", err)
	}
	if err := st.CompleteIncidents(); err != nil {
		t.Fatalf("CompleteIncidents idempotent run: %v", err)
	}
	if got := countIncidents(t, st, "corrupt"); got != 1 {
		t.Fatalf("incident records for corrupt = %d, want 1", got)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records after quarantine: %v", err)
	}
	if got := relayIDs(records); strings.Join(got, ",") != "clean,incident-corrupt" {
		t.Fatalf("records after quarantine = %v, want clean + incident-corrupt", got)
	}

	incident, err := st.Read("incident-corrupt")
	if err != nil {
		t.Fatalf("Read incident-corrupt: %v", err)
	}
	if incident.Envelope.DeliveryState != record.Held ||
		incident.Headers["record_kind"] != "incident" ||
		incident.Headers["quarantined_ref"] != "corrupt" ||
		incident.Headers["failure_class"] != "checksum-mismatch" {
		t.Fatalf("incident shape = envelope %+v headers %#v", incident.Envelope, incident.Headers)
	}
	if strings.Contains(incident.Body, root) {
		t.Fatalf("incident body leaked store path: %q", incident.Body)
	}

	_, err = st.Read("corrupt")
	var quarantined store.ErrQuarantined
	if !errors.As(err, &quarantined) {
		t.Fatalf("Read quarantined err = %v, want ErrQuarantined", err)
	}
	if quarantined.RelayID != "corrupt" ||
		quarantined.IncidentID != "incident-corrupt" ||
		quarantined.FailureClass != "checksum-mismatch" {
		t.Fatalf("ErrQuarantined = %+v", quarantined)
	}
}

func TestQuarantineOneIsIdempotent(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	commitRecord(t, st, "target", "body")

	evicted, err := st.QuarantineOne("target")
	if err != nil {
		t.Fatalf("QuarantineOne first: %v", err)
	}
	if !evicted {
		t.Fatalf("QuarantineOne first evicted = false, want true")
	}
	evicted, err = st.QuarantineOne("target")
	if err != nil {
		t.Fatalf("QuarantineOne second: %v", err)
	}
	if evicted {
		t.Fatalf("QuarantineOne second evicted = true, want false")
	}
}

func commitRecord(t *testing.T, st *store.Store, relayID, body string) {
	t.Helper()
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       relayID,
			From:          "seat-a",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": relayID},
		Body:    body,
	}, nil); err != nil {
		t.Fatalf("Commit %s: %v", relayID, err)
	}
}

func corruptRecordBody(t *testing.T, root, relayID, old, replacement string) {
	t.Helper()
	path := filepath.Join(root, "records", relayID+".json")
	data := string(mustRead(t, path))
	if !strings.Contains(data, old) {
		t.Fatalf("record %s does not contain %q", relayID, old)
	}
	data = strings.Replace(data, old, replacement, 1)
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatalf("write corrupt record: %v", err)
	}
}

func isChecksumFor(err error, relayID string) bool {
	var checksum store.ErrChecksum
	return errors.As(err, &checksum) && checksum.RelayID == relayID
}

func countIncidents(t *testing.T, st *store.Store, relayID string) int {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var count int
	for _, rec := range records {
		if rec.Headers["record_kind"] == "incident" && rec.Headers["quarantined_ref"] == relayID {
			count++
		}
	}
	return count
}

func relayIDs(records []record.Record) []string {
	ids := make([]string, 0, len(records))
	for _, rec := range records {
		ids = append(ids, rec.Envelope.RelayID)
	}
	return ids
}
