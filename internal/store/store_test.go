package store_test

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
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
	if _, err := os.Stat(filepath.Join(root, "journal", "redo", "000001.jsonl")); err != nil {
		t.Fatalf("segmented redo missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "journal", "redo.jsonl")); !os.IsNotExist(err) {
		t.Fatalf("legacy redo file exists: %v", err)
	}
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
	if err := os.RemoveAll(filepath.Join(root, "journal", "redo")); err != nil {
		t.Fatalf("remove redo: %v", err)
	}

	if err := st.RebuildProjections(); err != nil {
		t.Fatalf("RebuildProjections: %v", err)
	}
	index := string(mustRead(t, indexPath))
	wantIndex := "| relay-2 | SITREP | seat-a.implementer | seat-b.planner | accepted |\n"
	if !strings.Contains(index, "corrupt row\n") || !strings.Contains(index, wantIndex) {
		t.Fatalf("index did not preserve corrupt row and append canonical correction: %q", index)
	}
	assertFile(t, renderPath, "## relay-2\n\nFROM: seat-a.implementer\nTO: seat-b.planner\nSUBJECT: hello\n\nhello body\n")
	assertFile(t, mailboxPath, "relay-2\n")
}

func TestInitWritesStoreRootConfigAndGenesisRecord(t *testing.T) {
	root := t.TempDir()
	sources := writeConfigSources(t)
	counter := filepath.Join(t.TempDir(), "renames.log")
	t.Setenv("FRANK_TEST_RENAME_COUNTER", counter)

	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init: %v", err)
	}

	assertFile(t, filepath.Join(root, "config", "engine.json"), defaultEngineConfig)
	assertFile(t, filepath.Join(root, "config", "fieldspec", "registry.json"), testRegistryConfig)
	records, err := os.ReadDir(filepath.Join(root, "records"))
	if err != nil {
		t.Fatalf("read records: %v", err)
	}
	if len(records) != 1 || records[0].Name() != "genesis.json" {
		t.Fatalf("records = %v, want only genesis.json", records)
	}
	rec, err := record.Verify(mustRead(t, filepath.Join(root, "records", "genesis.json")))
	if err != nil {
		t.Fatalf("verify genesis: %v", err)
	}
	if rec.Envelope.RelayID != "genesis" ||
		rec.Envelope.DispatchID != "genesis" ||
		rec.Envelope.From != "system" ||
		rec.Envelope.Role != "system" ||
		rec.Envelope.DeliveryState != record.Accepted ||
		rec.Envelope.SchemaVersion != 1 {
		t.Fatalf("genesis envelope = %+v", rec.Envelope)
	}
	if rec.Headers["record_kind"] != "genesis" {
		t.Fatalf("record_kind = %q, want genesis", rec.Headers["record_kind"])
	}
	if _, ok := rec.Headers["schema_version"]; ok {
		t.Fatalf("schema_version header present; schema_version belongs in envelope only")
	}
	if rec.Headers["config_digest"] == "" ||
		rec.Headers["address_space_seed"] == "" ||
		rec.Headers["created_ts"] == "" {
		t.Fatalf("genesis headers incomplete: %#v", rec.Headers)
	}
	if !strings.Contains(rec.Headers["address_space_seed"], "operator") {
		t.Fatalf("address_space_seed does not carry literal operator address: %q", rec.Headers["address_space_seed"])
	}
	renames := string(mustRead(t, counter))
	if got := strings.Count(renames, "records/genesis.json\n"); got != 1 {
		t.Fatalf("genesis canonical rename count = %d, renames:\n%s", got, renames)
	}
}

func TestInitRejectsExistingGenesisWithoutChangingStoreBytes(t *testing.T) {
	root := t.TempDir()
	sources := writeConfigSources(t)
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init first: %v", err)
	}
	beforeRecords := hashTree(t, filepath.Join(root, "records"))
	beforeConfig := hashTree(t, filepath.Join(root, "config"))

	_, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open existing store: %v", err)
	}
	err = store.Init(root, sources)
	if !errors.Is(err, store.ErrGenesisExists) {
		t.Fatalf("Init second err = %v, want ErrGenesisExists", err)
	}
	if after := hashTree(t, filepath.Join(root, "records")); after != beforeRecords {
		t.Fatalf("records changed after rejected re-init:\nbefore %s\nafter  %s", beforeRecords, after)
	}
	if after := hashTree(t, filepath.Join(root, "config")); after != beforeConfig {
		t.Fatalf("config changed after rejected re-init:\nbefore %s\nafter  %s", beforeConfig, after)
	}
}

func TestValidateGenesisUsesStoreRootConfigDigest(t *testing.T) {
	root := t.TempDir()
	sources := writeConfigSources(t)
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	pinned := loadStorePinned(t, root)
	if err := st.ValidateGenesis(pinned); err != nil {
		t.Fatalf("ValidateGenesis: %v", err)
	}

	if err := os.WriteFile(sources["engine"], []byte(`{"gc_enabled":true,"segment_rotate_bytes":1}`), 0o644); err != nil {
		t.Fatalf("mutate outside source: %v", err)
	}
	if err := st.ValidateGenesis(loadStorePinned(t, root)); err != nil {
		t.Fatalf("ValidateGenesis changed after outside-source mutation: %v", err)
	}

	storeEngine := filepath.Join(root, "config", "engine.json")
	if err := os.WriteFile(storeEngine, []byte(`{"gc_enabled":true,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("mutate store-root config: %v", err)
	}
	err = st.ValidateGenesis(loadStorePinned(t, root))
	var mismatch store.ErrDigestMismatch
	if !errors.As(err, &mismatch) {
		t.Fatalf("ValidateGenesis err = %v, want ErrDigestMismatch", err)
	}
	if mismatch.Want == "" || mismatch.Got == "" || mismatch.Want == mismatch.Got {
		t.Fatalf("mismatch digests = %+v", mismatch)
	}
	for _, path := range []string{root, storeEngine, sources["engine"]} {
		if strings.Contains(err.Error(), path) {
			t.Fatalf("digest mismatch error leaked path %q in %q", path, err.Error())
		}
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

func assertFile(t *testing.T, path, want string) {
	t.Helper()
	got := string(mustRead(t, path))
	if got != want {
		t.Fatalf("%s = %q, want %q", path, got, want)
	}
}

const defaultEngineConfig = `{"gc_enabled":false,"segment_rotate_bytes":4194304}`

const testRegistryConfig = `{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`

func writeConfigSources(t *testing.T) map[string]string {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(root, "engine.json")
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(enginePath, []byte(defaultEngineConfig), 0o644); err != nil {
		t.Fatalf("write engine source: %v", err)
	}
	if err := os.WriteFile(registryPath, []byte(testRegistryConfig), 0o644); err != nil {
		t.Fatalf("write registry source: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath}
}

func loadStorePinned(t *testing.T, root string) *config.Pinned {
	t.Helper()
	pinned, err := config.Load(map[string]string{
		"engine":    filepath.Join(root, "config", "engine.json"),
		"fieldspec": filepath.Join(root, "config", "fieldspec", "registry.json"),
	})
	if err != nil {
		t.Fatalf("load store pinned config: %v", err)
	}
	return pinned
}

func hashTree(t *testing.T, root string) string {
	t.Helper()
	var parts []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(data)
		parts = append(parts, rel+"\x00"+hex.EncodeToString(sum[:]))
		return nil
	})
	if err != nil {
		t.Fatalf("hash tree %s: %v", root, err)
	}
	sort.Strings(parts)
	sum := sha256.Sum256([]byte(strings.Join(parts, "\n")))
	return hex.EncodeToString(sum[:])
}
