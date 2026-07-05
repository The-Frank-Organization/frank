package store_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestConfigChangeCommitMaterializesMember(t *testing.T) {
	root := t.TempDir()
	if err := store.Init(root, writeConfigSources(t)); err != nil {
		t.Fatalf("Init: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	before, err := os.ReadDir(filepath.Join(root, "records"))
	if err != nil {
		t.Fatalf("ReadDir before records: %v", err)
	}

	body := `{"phase":["SITREP","PLAN"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`
	rec := record.Record{
		Envelope: record.Envelope{RelayID: "config-change-1", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":       "SITREP",
			"SUBJECT":     "registry update",
			"record_kind": "config_change",
			"member":      "fieldspec",
			"new_digest":  digestWithMember(t, root, "fieldspec", []byte(body)),
		},
		Body: body,
	}
	if _, err := st.Commit(rec, store.ConfigChangeIntents(rec)); err != nil {
		t.Fatalf("Commit config_change: %v", err)
	}

	after, err := os.ReadDir(filepath.Join(root, "records"))
	if err != nil {
		t.Fatalf("ReadDir after records: %v", err)
	}
	if len(after) != len(before)+1 {
		t.Fatalf("records before=%d after=%d, want exactly one new canonical record", len(before), len(after))
	}
	got, err := os.ReadFile(filepath.Join(root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read materialized registry: %v", err)
	}
	if string(got) != body {
		t.Fatalf("materialized registry = %q, want body %q", got, body)
	}
	index := string(mustRead(t, filepath.Join(root, "projections", "INDEX.md")))
	if !strings.Contains(index, "config-change-1") {
		t.Fatalf("INDEX missing config_change row:\n%s", index)
	}
}

func TestChainWalkLatestWins(t *testing.T) {
	root := t.TempDir()
	if err := store.Init(root, writeConfigSources(t)); err != nil {
		t.Fatalf("Init: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	first := `{"phase":["SITREP","PLAN"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`
	firstDigest := digestWithMember(t, root, "fieldspec", []byte(first))
	firstRec := configChangeRecord("config-change-first", first, firstDigest)
	if _, err := st.Commit(firstRec, store.ConfigChangeIntents(firstRec)); err != nil {
		t.Fatalf("commit first config_change: %v", err)
	}
	second := `{"phase":["SITREP","PLAN","IMPL"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`
	secondDigest := digestWithMember(t, root, "fieldspec", []byte(second))
	secondRec := configChangeRecord("config-change-second", second, secondDigest)
	if _, err := st.Commit(secondRec, store.ConfigChangeIntents(secondRec)); err != nil {
		t.Fatalf("commit second config_change: %v", err)
	}

	got, err := st.ExpectedConfigDigest()
	if err != nil {
		t.Fatalf("ExpectedConfigDigest: %v", err)
	}
	if got != secondDigest {
		t.Fatalf("ExpectedConfigDigest = %s, want latest %s", got, secondDigest)
	}
}

func configChangeRecord(relayID, body, digest string) record.Record {
	return record.Record{
		Envelope: record.Envelope{RelayID: relayID, From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers: map[string]string{
			"PHASE":       "SITREP",
			"SUBJECT":     "registry update",
			"record_kind": "config_change",
			"member":      "fieldspec",
			"new_digest":  digest,
		},
		Body: body,
	}
}

func digestWithMember(t *testing.T, root, member string, body []byte) string {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store config: %v", err)
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return config.Digest(members)
}
