package seat_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
)

func TestMintResolveAndRejectUnknownCredential(t *testing.T) {
	mgr, err := seat.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	cred, err := mgr.Mint("s1-core.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	meta, ok := mgr.Resolve(cred.Value)
	if !ok {
		t.Fatalf("credential did not resolve")
	}
	if meta.Name != "s1-core.implementer" || meta.Role != "implementer" || meta.IsOperator {
		t.Fatalf("meta = %+v", meta)
	}
	if _, ok := mgr.Resolve("missing"); ok {
		t.Fatalf("unknown credential resolved")
	}
}

func TestDuplicateMintRejectsWithoutSecondCredential(t *testing.T) {
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint first: %v", err)
	}
	before := mustRead(t, filepath.Join(root, "binding", "seats.json"))
	_, err = mgr.Mint("seat-a", "implementer", false)
	if !errors.Is(err, seat.ErrSeatAlreadyBound) {
		t.Fatalf("Mint duplicate err = %v, want ErrSeatAlreadyBound", err)
	}
	after := mustRead(t, filepath.Join(root, "binding", "seats.json"))
	if !bytes.Equal(before, after) {
		t.Fatalf("binding table changed after duplicate mint")
	}
	if _, ok := mgr.Resolve(cred.Value); !ok {
		t.Fatalf("original credential stopped resolving")
	}
	if got := mgr.CredentialsFor("seat-a"); got != 1 {
		t.Fatalf("credentials for seat-a = %d, want 1", got)
	}
}

func TestMintRejectsReservedSystemSeatWithoutBinding(t *testing.T) {
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	_, err = mgr.Mint("system", "system", false)
	if !errors.Is(err, seat.ErrReservedSeatName) {
		t.Fatalf("Mint system err = %v, want ErrReservedSeatName", err)
	}
	if got := mgr.CredentialsFor("system"); got != 0 {
		t.Fatalf("credentials for system = %d, want 0", got)
	}
	if _, err := os.Stat(filepath.Join(root, "binding", "seats.json")); !os.IsNotExist(err) {
		t.Fatalf("binding table created for rejected system mint: %v", err)
	}
}

func TestStampOverwritesPayloadIdentity(t *testing.T) {
	meta := seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}
	rec := record.Record{
		Envelope: record.Envelope{From: "victim.planner", Role: "planner"},
		Headers:  map[string]string{"SUBJECT": "hello"},
	}
	stamped := seat.Stamp(rec, meta)
	if stamped.Envelope.From != "s1-core.implementer" || stamped.Envelope.Role != "implementer" {
		t.Fatalf("stamped envelope = %+v", stamped.Envelope)
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return data
}
