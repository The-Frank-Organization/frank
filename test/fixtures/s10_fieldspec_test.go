package fixtures_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fieldspec"
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
