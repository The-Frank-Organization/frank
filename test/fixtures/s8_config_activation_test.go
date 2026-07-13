package fixtures_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
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
const s8CatalogSHA256 = "943f07bb51da3414cf45a16d4bfa00bcee28cc538533fcb7fcd3e8a64b5e209d"

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
	if pinned.Engine.Version != 2 {
		t.Fatalf("engine.version = %d, want 2", pinned.Engine.Version)
	}
	t.Logf("fresh_v2_genesis_config_digest=%s", pinned.Digest)
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

func TestS8ProductionInitRejectsMissingCatalog(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, buildFrank(t, ctx),
		"-root", root,
		"-registry", sources["fieldspec"],
		"-engine-config", sources["engine"],
		"-init",
	)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("frank -init without catalog unexpectedly succeeded: %s", out)
	}
	if !bytes.Contains(out, []byte("catalog required for init")) {
		t.Fatalf("frank -init output = %q, want catalog-required refusal", out)
	}
}

func TestS8ProductionInitRejectsObserveTrueEngine(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, true)
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, buildFrank(t, ctx),
		"-root", root,
		"-registry", sources["fieldspec"],
		"-engine-config", sources["engine"],
		"-catalog", sources["catalog"],
		"-init",
	)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("frank -init with observe:true unexpectedly succeeded: %s", out)
	}
	if !bytes.Contains(out, []byte("observe must be false at genesis")) {
		t.Fatalf("frank -init output = %q, want observe-at-genesis refusal", out)
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

func TestS8FXCFG3CatalogLoadShapeIsClosed(t *testing.T) {
	sources := s8ConfigSources(t, false)
	data, err := os.ReadFile(sources["catalog"])
	if err != nil {
		t.Fatalf("read catalog: %v", err)
	}
	sum := sha256.Sum256(data)
	if got := hex.EncodeToString(sum[:]); got != s8CatalogSHA256 {
		t.Fatalf("catalog hash = %s, want amended owner bytes %s", got, s8CatalogSHA256)
	}
	if _, err := config.Load(sources); err != nil {
		t.Fatalf("load valid s8 catalog: %v", err)
	}
	var valid map[string]any
	if err := json.Unmarshal(data, &valid); err != nil {
		t.Fatalf("decode valid catalog: %v", err)
	}
	claim := valid["change_convention"].(map[string]any)["section_7_claim"]
	if claim != "pinned-at-s8: load enforces digest and member/provenance shape; owner review remains a relay/design-review gate" {
		t.Fatalf("load-claim boundary = %q", claim)
	}

	for _, tc := range []struct {
		name   string
		mutate func(map[string]any)
	}{
		{name: "duplicate law id", mutate: func(doc map[string]any) {
			laws := doc["laws"].([]any)
			laws[1].(map[string]any)["id"] = laws[0].(map[string]any)["id"]
		}},
		{name: "empty owners", mutate: func(doc map[string]any) {
			doc["laws"].([]any)[0].(map[string]any)["owners"] = []any{}
		}},
		{name: "broken census section", mutate: func(doc map[string]any) {
			delete(doc["discovery"].(map[string]any), "sink_patterns")
		}},
		{name: "unknown top level", mutate: func(doc map[string]any) {
			doc["unexpected"] = true
		}},
		{name: "missing canonical directory", mutate: func(doc map[string]any) {
			rows := doc["discovery"].(map[string]any)["canonical_path_families"].(map[string]any)["rows"].([]any)
			delete(rows[0].(map[string]any), "directory")
		}},
		{name: "missing fidelity review", mutate: func(doc map[string]any) {
			delete(doc["laws"].([]any)[0].(map[string]any), "fidelity_review")
		}},
		{name: "missing channel ident calls", mutate: func(doc map[string]any) {
			recognizers := doc["discovery"].(map[string]any)["recognizers"].(map[string]any)
			delete(recognizers["channel_context"].(map[string]any), "ident_calls")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var doc map[string]any
			if err := json.Unmarshal(data, &doc); err != nil {
				t.Fatalf("decode catalog: %v", err)
			}
			tc.mutate(doc)
			candidate, err := json.Marshal(doc)
			if err != nil {
				t.Fatalf("marshal catalog mutation: %v", err)
			}
			path := filepath.Join(t.TempDir(), "catalog.json")
			if err := os.WriteFile(path, candidate, 0o644); err != nil {
				t.Fatalf("write catalog mutation: %v", err)
			}
			mutated := map[string]string{"engine": sources["engine"], "fieldspec": sources["fieldspec"], "catalog": path}
			if _, err := config.Load(mutated); !errors.Is(err, config.ErrConfigLoad) {
				t.Fatalf("malformed catalog error = %v, want config-load", err)
			}
		})
	}
}

func TestS8FXCFG5GenesisCatalogMatchesSourceByteExact(t *testing.T) {
	sources := s8ConfigSources(t, false)
	root := t.TempDir()
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init: %v", err)
	}
	source, err := os.ReadFile(sources["catalog"])
	if err != nil {
		t.Fatalf("read source catalog: %v", err)
	}
	runtime, err := os.ReadFile(filepath.Join(root, "config", "catalog", "catalog.json"))
	if err != nil {
		t.Fatalf("read runtime catalog: %v", err)
	}
	if !bytes.Equal(runtime, source) {
		t.Fatal("genesis catalog differs from source artifact bytes")
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("Load runtime config: %v", err)
	}
	want, err := st.ExpectedConfigDigest()
	if err != nil {
		t.Fatalf("ExpectedConfigDigest: %v", err)
	}
	if pinned.Digest != want {
		t.Fatalf("runtime digest = %s, history expects %s", pinned.Digest, want)
	}
}

func s8ConfigSources(t *testing.T, observe bool) map[string]string {
	t.Helper()
	root := t.TempDir()
	engine := fixtureEngineConfig(t, observe)
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
