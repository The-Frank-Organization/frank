package fixtures_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/store"
)

func TestS10V7ReaderRefusesV8MarkerBeforeContent(t *testing.T) {
	v8, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v8 registry: %v", err)
	}
	var planted map[string]any
	if err := json.Unmarshal(v8, &planted); err != nil {
		t.Fatalf("decode v8 registry: %v", err)
	}
	if got := planted["version"]; got != "s10-fieldspec-v8" {
		t.Fatalf("registry version = %q, want s10-fieldspec-v8", got)
	}
	planted["fields"] = "content-must-not-be-interpreted"
	plantedBytes, err := json.Marshal(planted)
	if err != nil {
		t.Fatalf("marshal planted registry: %v", err)
	}
	if _, err := fieldspec.Parse(plantedBytes); err == nil {
		t.Fatal("planted fieldspec content unexpectedly valid")
	}
	if err := config.ValidateFieldspecReaderMarker(
		plantedBytes,
		"s7a-fieldspec-v5",
		"s8-fieldspec-v6",
		"s8-fieldspec-v7",
	); !errors.Is(err, config.ErrConfigLoad) {
		t.Fatalf("v7 reader marker error = %v, want config-load", err)
	}
	if err := config.ValidateFieldspecReaderMarker(
		plantedBytes,
		"s7a-fieldspec-v5",
		"s8-fieldspec-v6",
		"s8-fieldspec-v7",
		"s10-fieldspec-v8",
	); err != nil {
		t.Fatalf("v8 marker preflight interpreted planted content: %v", err)
	}
}

func TestS10FreshGenesisIsBornAtV8WithoutPreV8Records(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	wantFieldspec, err := os.ReadFile(sources["fieldspec"])
	if err != nil {
		t.Fatalf("read v8 source: %v", err)
	}
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init fresh v8 store: %v", err)
	}

	gotFieldspec, err := os.ReadFile(filepath.Join(root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read materialized fieldspec: %v", err)
	}
	if !bytes.Equal(gotFieldspec, wantFieldspec) {
		t.Fatal("fresh genesis did not write the v8 fieldspec bytes directly")
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load fresh v8 config: %v", err)
	}
	if pinned.Registry.Version != "s10-fieldspec-v8" {
		t.Fatalf("fresh genesis fieldspec = %q, want s10-fieldspec-v8", pinned.Registry.Version)
	}

	entries, err := os.ReadDir(filepath.Join(root, "records"))
	if err != nil {
		t.Fatalf("read records: %v", err)
	}
	if len(entries) != 1 || entries[0].Name() != "genesis.json" {
		t.Fatalf("fresh v8 records = %v, want only genesis.json", entries)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open fresh v8 store: %v", err)
	}
	genesis, err := st.Genesis()
	if err != nil {
		t.Fatalf("read genesis: %v", err)
	}
	if genesis.Headers["config_digest"] != pinned.Digest {
		t.Fatalf("genesis digest = %q, want current v8 digest %q", genesis.Headers["config_digest"], pinned.Digest)
	}
}

func TestS10FreshGenesisRejectsV8WithCorruptedLockedPredecessorByte(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	v8, err := os.ReadFile(sources["fieldspec"])
	if err != nil {
		t.Fatalf("read v8 source: %v", err)
	}
	old := []byte("non-gate-bearing records only")
	next := []byte("non-gate-bearing recordz only")
	if bytes.Count(v8, old) != 1 {
		t.Fatalf("locked predecessor byte fixture count = %d, want 1", bytes.Count(v8, old))
	}
	corrupt := bytes.Replace(v8, old, next, 1)
	corruptPath := filepath.Join(t.TempDir(), "registry-corrupt-v8.json")
	if err := os.WriteFile(corruptPath, corrupt, 0o644); err != nil {
		t.Fatalf("write corrupt v8 source: %v", err)
	}
	sources["fieldspec"] = corruptPath
	if err := store.Init(root, sources); err == nil || !strings.Contains(err.Error(), "fieldspec genesis predecessor") {
		t.Fatalf("Init corrupt v8 err = %v, want predecessor integrity failure", err)
	}
}
