package fixtures_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

const s8FieldspecV5SHA256 = "1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485"

func TestS8FXCFG7GenesisComposesThreePinnedMembers(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init: %v", err)
	}

	paths := store.StoreRootConfigPaths(root)
	for _, member := range []string{"catalog", "engine", "fieldspec"} {
		if paths[member] == "" {
			t.Fatalf("StoreRootConfigPaths missing %s: %#v", member, paths)
		}
	}
	pinned, err := config.Load(paths)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if pinned.Engine.Version != 1 {
		t.Fatalf("engine.version = %d, want 1", pinned.Engine.Version)
	}
	if config.PresentLayers(pinned)["observe"] {
		t.Fatalf("observe active at genesis")
	}
	fieldspecSum := sha256.Sum256(pinned.Members["fieldspec"])
	if got := hex.EncodeToString(fieldspecSum[:]); got != s8FieldspecV5SHA256 {
		t.Fatalf("fieldspec hash = %s, want %s", got, s8FieldspecV5SHA256)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	genesis, err := st.Genesis()
	if err != nil {
		t.Fatalf("Genesis: %v", err)
	}
	if got := genesis.Headers["config_digest"]; got != pinned.Digest {
		t.Fatalf("genesis digest = %s, loaded digest = %s", got, pinned.Digest)
	}
}

func TestS8FXCFG1And2ConfigDerivedPresentLayersAreGenerationStable(t *testing.T) {
	off := s8LoadPinnedSources(t, s8ConfigSources(t, false))
	on := s8LoadPinnedSources(t, s8ConfigSources(t, true))
	if config.PresentLayers(off)["observe"] {
		t.Fatalf("observe present while knob is false")
	}
	layers := config.PresentLayers(on)
	if !layers["store"] || !layers["form"] || !layers["lineage"] || !layers["observe"] {
		t.Fatalf("activated layers = %#v", layers)
	}

	reg := on.Registry
	seat := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	env := fieldspec.RenderEnv{ConfigDigest: on.Digest, PresentLayers: layers}
	form, digest := reg.Render(env, seat, "SITREP", "medium", fieldspec.ClosedGrantState)
	if !form.HasField("ACTIONS_GIT_REF") || !form.HasField("FINAL_GIT_STATUS_SHORT") {
		t.Fatalf("observe-active form omitted observed refs: %#v", form.Fields)
	}
	for i := 0; i < 8; i++ {
		_, again := reg.Render(env, seat, "SITREP", "medium", fieldspec.ClosedGrantState)
		if again != digest {
			t.Fatalf("generation digest bounced at iteration %d: %s != %s", i, again, digest)
		}
	}

	cand := record.Record{Headers: map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium",
		"EVIDENCE_TARGET": "E1", "SUBJECT": "activation", "ACTIONS_GIT_REF": "none — no edits",
		"FINAL_GIT_STATUS_SHORT": "none — clean tree",
	}}
	violations := reg.Validate(cand, seat, digest, env, fieldspec.ClosedGrantState)
	for _, violation := range violations {
		if violation.Field == "form_digest" {
			t.Fatalf("render/validate split semantics: %#v", violations)
		}
	}
}

func TestS8ProductionInitPinsCatalog(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, buildFrank(t, ctx),
		"-root", root,
		"-registry", sources["fieldspec"],
		"-engine-config", sources["engine"],
		"-catalog", sources["catalog"],
		"-init",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("frank -init: %v\n%s", err, out)
	}
	if _, err := os.Stat(filepath.Join(root, "config", "catalog", "catalog.json")); err != nil {
		t.Fatalf("catalog not pinned by production init: %v", err)
	}
}

func TestS8ServeRejectsLegacyStoreWithBlessInstruction(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	delete(sources, "catalog")
	if err := os.WriteFile(sources["engine"], []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write legacy engine: %v", err)
	}
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("legacy Init: %v", err)
	}

	legacyPaths := store.LegacyStoreRootConfigPaths(root)
	if len(legacyPaths) != 2 || legacyPaths["catalog"] != "" {
		t.Fatalf("legacy paths = %#v, want explicit two-member expectation", legacyPaths)
	}
	if _, err := config.Load(legacyPaths); err != nil {
		t.Fatalf("explicit legacy load: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, buildFrank(t, ctx), "-root", root)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("legacy store unexpectedly served")
	}
	if ctx.Err() != nil {
		t.Fatalf("legacy store served until timeout; output=%s", out)
	}
	if !bytes.Contains(out, []byte("store-not-adopted: run frank -bless")) {
		t.Fatalf("legacy serve output = %q, want typed bless instruction", out)
	}
}

func s8ConfigSources(t *testing.T, observe bool) map[string]string {
	t.Helper()
	root := t.TempDir()
	engine, err := json.Marshal(map[string]any{
		"version":              1,
		"gc_enabled":           false,
		"segment_rotate_bytes": 4194304,
		"present_layers":       map[string]bool{"observe": observe},
	})
	if err != nil {
		t.Fatalf("marshal engine: %v", err)
	}
	enginePath := filepath.Join(root, "engine.json")
	if err := os.WriteFile(enginePath, engine, 0o644); err != nil {
		t.Fatalf("write engine: %v", err)
	}
	registryPath := filepath.Join(root, "registry.json")
	registry, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	if err := os.WriteFile(registryPath, registry, 0o644); err != nil {
		t.Fatalf("write registry: %v", err)
	}
	catalogPath := filepath.Join(root, "catalog.json")
	catalog, err := os.ReadFile(filepath.Join("..", "invariants", "catalog.v1.json"))
	if err != nil {
		t.Fatalf("read catalog: %v", err)
	}
	if err := os.WriteFile(catalogPath, catalog, 0o644); err != nil {
		t.Fatalf("write catalog: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath, "catalog": catalogPath}
}

func s8LoadPinnedSources(t *testing.T, sources map[string]string) *config.Pinned {
	t.Helper()
	pinned, err := config.Load(sources)
	if err != nil {
		t.Fatalf("Load sources: %v", err)
	}
	return pinned
}
