package migrate_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestApplyWalksRegisteredMigratorsOnCopies(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	rec := record.Record{Envelope: record.Envelope{RelayID: "r1", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"old": "value"}}
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("Commit: %v", err)
	}
	path := filepath.Join(st.Root, "records", "r1.json")
	before, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read before: %v", err)
	}
	reg := migrate.New()
	reg.Current = 2
	reg.Register(1, func(rec record.Record) (record.Record, error) {
		rec.Headers["new"] = rec.Headers["old"]
		delete(rec.Headers, "old")
		return rec, nil
	})
	view, err := (migrate.Reader{Store: st, Registry: reg}).Read("r1")
	if err != nil {
		t.Fatalf("Read view: %v", err)
	}
	if view.SourceVersion != 1 || view.Record.Envelope.SchemaVersion != 2 || view.Record.Headers["new"] != "value" {
		t.Fatalf("view = %+v", view)
	}
	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read after: %v", err)
	}
	if string(before) != string(after) {
		t.Fatalf("stored bytes mutated")
	}
}

func TestApplyRefusalLegs(t *testing.T) {
	reg := migrate.New()
	reg.Current = 2
	for _, tc := range []struct {
		name string
		rec  record.Record
		want error
	}{
		{"future", record.Record{Envelope: record.Envelope{SchemaVersion: 3}}, migrate.ErrFutureVersion},
		{"unversioned", record.Record{}, migrate.ErrUnversioned},
		{"gap", record.Record{Envelope: record.Envelope{SchemaVersion: 1}}, migrate.ErrMigrationGap},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := migrate.Apply(reg, tc.rec); !errors.Is(err, tc.want) {
				t.Fatalf("Apply error = %v, want %v", err, tc.want)
			}
		})
	}
}

func TestReaderStoreErrorsWinBeforeMigration(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if err := os.WriteFile(filepath.Join(st.Root, "records", "bad.json"), []byte(`{"bad":true}`), 0o644); err != nil {
		t.Fatalf("write corrupt: %v", err)
	}
	reg := migrate.New()
	reg.Current = 2
	_, err = (migrate.Reader{Store: st, Registry: reg}).Read("bad")
	var checksum store.ErrChecksum
	if !errors.As(err, &checksum) {
		t.Fatalf("Read error = %v, want checksum", err)
	}
}
