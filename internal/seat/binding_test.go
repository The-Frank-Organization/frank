package seat_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
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

func TestMintOrReplaceReplacesCredentialAndKeepsSingleBinding(t *testing.T) {
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint first: %v", err)
	}

	replacement, err := mgr.MintOrReplace("seat-a", "planner", true, "pivot-1")
	if err != nil {
		t.Fatalf("MintOrReplace: %v", err)
	}
	if replacement.Value == "" || replacement.Value == cred.Value {
		t.Fatalf("replacement credential = %q, original = %q", replacement.Value, cred.Value)
	}
	if _, ok := mgr.Resolve(cred.Value); ok {
		t.Fatalf("old credential still resolves after replacement")
	}
	meta, ok := mgr.Resolve(replacement.Value)
	if !ok {
		t.Fatalf("replacement credential did not resolve")
	}
	if meta.Name != "seat-a" || meta.Role != "planner" || !meta.IsOperator {
		t.Fatalf("replacement meta = %+v", meta)
	}
	if got := mgr.CredentialsFor("seat-a"); got != 1 {
		t.Fatalf("credentials for seat-a = %d, want 1", got)
	}
	if got, ok := mgr.RealizedMintRef("seat-a"); !ok || got != "pivot-1" {
		t.Fatalf("RealizedMintRef = %q, %v; want pivot-1, true", got, ok)
	}
	var table struct {
		Seats map[string]struct {
			Credential      string `json:"credential"`
			RealizedMintRef string `json:"realized_mint_ref"`
		} `json:"seats"`
	}
	if err := json.Unmarshal(mustRead(t, filepath.Join(root, "binding", "seats.json")), &table); err != nil {
		t.Fatalf("decode binding table: %v", err)
	}
	row := table.Seats["seat-a"]
	if row.Credential != replacement.Value || row.RealizedMintRef != "pivot-1" {
		t.Fatalf("binding row = %+v, want replacement credential and pivot", row)
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
