package manifest

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

func TestWorkspaceRootOrderedRefusalTable(t *testing.T) {
	valid := "/workspace"
	tests := []struct {
		name     string
		input    string
		resolved string
		resolve  error
		want     WorkspaceRootReason
	}{
		{name: "absent", want: RootNotAbsolute},
		{name: "relative ignores crafted cwd", input: "crafted-relative-root", resolved: valid, want: RootNotAbsolute},
		{name: "filesystem root", input: "/", resolved: "/", want: RootFilesystemRoot},
		{name: "symlink to root", input: "/root-link", resolved: "/", want: RootFilesystemRoot},
		{name: "nonexistent", input: "/missing", resolve: os.ErrNotExist, want: RootUnresolvable},
		{name: "loop", input: "/loop", resolve: errors.New("too many symbolic links"), want: RootUnresolvable},
		{name: "os path length", input: "/long", resolve: errors.New("file name too long"), want: RootUnresolvable},
		{name: "grammar before length", input: "/bad", resolved: "/" + strings.Repeat("a", PathMaxM10) + "\n", want: RootOutOfGrammar},
		{name: "out of grammar", input: "/bad", resolved: "/bad\nname", want: RootOutOfGrammar},
		{name: "too long", input: "/long", resolved: "/" + strings.Repeat("a", PathMaxM10), want: RootTooLong},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calls := 0
			_, err := resolveWorkspaceRoot(test.input, func(string) (string, error) {
				calls++
				return test.resolved, test.resolve
			})
			var typed *WorkspaceRootError
			if !errors.As(err, &typed) || typed.Reason != test.want {
				t.Fatalf("error = %v, want reason %q", err, test.want)
			}
			if test.want == RootNotAbsolute && calls != 0 {
				t.Fatalf("realpath calls = %d, want zero", calls)
			}
			if test.input != "" && test.input != "/" && strings.Contains(err.Error(), test.input) {
				t.Fatalf("typed refusal leaked path: %q", err)
			}
		})
	}
}

func TestWorkspaceRootConvergesAndHashesNormalizedRealpath(t *testing.T) {
	realRoot := filepath.Join(t.TempDir(), "café")
	if err := os.Mkdir(realRoot, 0o700); err != nil {
		t.Fatalf("mkdir root: %v", err)
	}
	link := filepath.Join(t.TempDir(), "workspace-link")
	if err := os.Symlink(realRoot, link); err != nil {
		t.Fatalf("symlink: %v", err)
	}

	direct, err := ResolveWorkspaceRoot(realRoot + string(filepath.Separator))
	if err != nil {
		t.Fatalf("resolve direct: %v", err)
	}
	linked, err := ResolveWorkspaceRoot(link)
	if err != nil {
		t.Fatalf("resolve link: %v", err)
	}
	if direct != linked {
		t.Fatalf("root convergence failed: direct=%#v linked=%#v", direct, linked)
	}
	want := digestString(direct.Path)
	if direct.ID != want {
		t.Fatalf("workspace_root_id = %q, want sha256(path) %q", direct.ID, want)
	}

	nfd := "/tmp/cafe\u0301"
	normalized, err := resolveWorkspaceRoot(nfd, func(string) (string, error) { return nfd, nil })
	if err != nil {
		t.Fatalf("resolve NFD fixture: %v", err)
	}
	if normalized.Path != "/tmp/café" || normalized.ID != digestString(normalized.Path) {
		t.Fatalf("NFC root = %#v", normalized)
	}
}

func TestManifestConstructionUsesExactEightNamesAndFrozenCanonicalBytes(t *testing.T) {
	root := filepath.Join(t.TempDir(), "workspace")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	input := validBuildInput(root, "run-manifest")
	frozen, err := Build(input)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	wantNames := []string{"apply_patch", "bash", "edit", "read", "relay.project", "relay.read", "relay.submit", "write"}
	if got := ToolNames(frozen.Manifest.ToolSet); !reflect.DeepEqual(got, wantNames) {
		t.Fatalf("tool names = %v, want %v", got, wantNames)
	}
	if got := hexDigest(frozen.Bytes); got != frozen.Digest {
		t.Fatalf("digest = %q, want digest(stored bytes) %q", frozen.Digest, got)
	}
	if !strings.Contains(string(frozen.Bytes), `"mapping_version":"mapping-v1"`) || !strings.Contains(string(frozen.Bytes), `"workspace_root_path"`) {
		t.Fatalf("canonical manifest omitted required locked/pair members: %s", frozen.Bytes)
	}

	stagingInput := validBuildInput(root, "run-staging-manifest")
	stagingInput.ToolSet = StagingToolSet()
	stagingInput.ToolCatalogDigest = nil
	stagingInput.ReleaseBinding = nil
	staged, err := Build(stagingInput)
	if err != nil {
		t.Fatalf("Build staging: %v", err)
	}
	stagedBytes := string(staged.Bytes)
	for _, marker := range []string{`"schema_digest":null`, `"catalog_version":null`, `"mapping_version":null`, `"tool_catalog_digest":null`} {
		if !strings.Contains(stagedBytes, marker) {
			t.Fatalf("canonical staging manifest omitted %s: %s", marker, staged.Bytes)
		}
	}
	if strings.Contains(stagedBytes, `"release_binding"`) {
		t.Fatalf("canonical staging manifest unexpectedly bound a release: %s", staged.Bytes)
	}

	input.ToolSet[0].Name = "mutated-after-build"
	if strings.Contains(string(frozen.Bytes), "mutated-after-build") {
		t.Fatal("frozen bytes aliased mutable input")
	}
}

func TestServeGateIsFailClosedUntilFullIdentityAndReleaseBindingMatch(t *testing.T) {
	root := filepath.Join(t.TempDir(), "workspace")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	input := validBuildInput(root, "run-gate")
	gate := Gate{
		LockedTools: input.ToolSet, ShippedToolCatalogDigest: digestString("catalog"),
		PolicyBytes: input.PolicyBytes, LaneCatalogDigest: input.ProviderLane.LaneCatalogDigest,
	}

	staging := input
	staging.ToolSet = StagingToolSet()
	staging.ToolCatalogDigest = nil
	staging.ReleaseBinding = nil
	staged, err := Build(staging)
	if err != nil {
		t.Fatalf("Build staging: %v", err)
	}
	if err := gate.Validate(staged); !errors.Is(err, ErrIdentityUnlocked) {
		t.Fatalf("staging gate error = %v, want ErrIdentityUnlocked", err)
	}

	frozen, err := Build(input)
	if err != nil {
		t.Fatalf("Build locked: %v", err)
	}
	if err := gate.Validate(frozen); err != nil {
		t.Fatalf("Validate locked: %v", err)
	}
	tampered := frozen
	tampered.Bytes = append([]byte(nil), frozen.Bytes...)
	tampered.Bytes[len(tampered.Bytes)-1] = ' '
	if err := gate.Validate(tampered); !errors.Is(err, ErrManifestIntegrity) {
		t.Fatalf("tampered frozen bytes = %v, want ErrManifestIntegrity", err)
	}
	for _, name := range RatifiedToolNames {
		if err := gate.Authorize(frozen, name); err != nil {
			t.Fatalf("Authorize %q: %v", name, err)
		}
	}
	if err := gate.Authorize(frozen, "network.fetch"); !errors.Is(err, ErrDeniedAboveSet) {
		t.Fatalf("above-set authorization = %v", err)
	}

	for name, mutate := range map[string]func(*BuildInput){
		"missing tool": func(candidate *BuildInput) { candidate.ToolSet = candidate.ToolSet[1:] },
		"widened tool": func(candidate *BuildInput) { candidate.ToolSet[0].Name = "network.fetch" },
		"identity mismatch": func(candidate *BuildInput) {
			value := digestString("drift")
			candidate.ToolSet[0].SchemaDigest = &value
		},
		"catalog mismatch": func(candidate *BuildInput) { value := digestString("other"); candidate.ToolCatalogDigest = &value },
		"lane mismatch": func(candidate *BuildInput) {
			candidate.ProviderLane.LaneCatalogDigest = digestString("other lane catalog")
		},
		"release both forms": func(candidate *BuildInput) {
			value := digestString("release")
			candidate.ReleaseBinding.ReleaseDigest = &value
		},
		"release neither form": func(candidate *BuildInput) { candidate.ReleaseBinding.BuildDigests = nil },
	} {
		t.Run(name, func(t *testing.T) {
			candidate := validBuildInput(root, "run-"+strings.ReplaceAll(name, " ", "-"))
			mutate(&candidate)
			built, buildErr := Build(candidate)
			if buildErr == nil {
				buildErr = gate.Validate(built)
			}
			if buildErr == nil {
				t.Fatal("drifted manifest passed serve gate")
			}
		})
	}
}

func TestFreezeEventCommitsCanonicalManifestOnce(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	host := applier.New(db, applier.Config{})
	defer host.Close()
	root := filepath.Join(t.TempDir(), "workspace")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	input := validBuildInput(root, "run-freeze")
	frozen, err := Build(input)
	if err != nil {
		t.Fatal(err)
	}
	gate := Gate{LockedTools: input.ToolSet, ShippedToolCatalogDigest: *input.ToolCatalogDigest, PolicyBytes: input.PolicyBytes, LaneCatalogDigest: input.ProviderLane.LaneCatalogDigest}
	if _, err := host.Apply(ctx, FreezeEvent{Frozen: frozen, Gate: gate, SessionLogPath: "/session.log", CreatedAt: 7}); err != nil {
		t.Fatalf("freeze: %v", err)
	}
	loaded, err := Load(ctx, host, "run-freeze", gate)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if !reflect.DeepEqual(loaded.Bytes, frozen.Bytes) || loaded.Digest != frozen.Digest {
		t.Fatalf("loaded manifest drift: got=%#v want=%#v", loaded, frozen)
	}

	changed := validBuildInput(root, "run-freeze")
	changed.PolicySourceRef = "different"
	other, err := Build(changed)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := host.Apply(ctx, FreezeEvent{Frozen: other, Gate: gate, SessionLogPath: "/session.log", CreatedAt: 8}); err == nil {
		t.Fatal("same run_id manifest mutation committed")
	}
	again, err := Load(ctx, host, "run-freeze", gate)
	if err != nil || !reflect.DeepEqual(again.Bytes, frozen.Bytes) {
		t.Fatalf("immutable row changed after conflicting freeze: loaded=%#v err=%v", again, err)
	}

	driftedGate := gate
	driftedGate.ShippedToolCatalogDigest = digestString("different shipped catalog")
	if _, err := Load(ctx, host, "run-freeze", driftedGate); !errors.Is(err, ErrServeGate) {
		t.Fatalf("reload with F63 mismatch = %v, want ErrServeGate", err)
	}
}

func validBuildInput(root, runID string) BuildInput {
	policy := []byte(`{"pinned_lane":{"model_id":"model","provider_id":"provider","serving_profile_id":"profile","compat_mode":"native"}}`)
	lane := LaneID{ModelID: "model", ProviderID: "provider", ServingProfileID: "profile", CompatMode: "native"}
	catalog := digestString("catalog")
	return BuildInput{
		RunID: runID, PolicySourceRef: "operator-ratification/ref", PolicyDigest: hexDigest(policy), PolicyBytes: policy,
		PolicyPinnedLane: lane, ToolSet: lockedToolSet(), ToolCatalogDigest: &catalog,
		ProviderLane:  ProviderLane{LaneID: lane, LaneCatalogDigest: digestString("lane catalog"), CredentialRef: "opaque-ref"},
		WorkspaceRoot: root,
		ReleaseBinding: &ReleaseBinding{BoundAtRef: "release-binding/ref", BuildDigests: &BuildDigests{
			AppMainBuildDigest: digestString("app"), M9WorkerBuildDigest: digestString("worker"), M8BuildDigest: digestString("connector"),
		}},
	}
}

func lockedToolSet() []ToolIdentity {
	tools := StagingToolSet()
	for i := range tools {
		schema := digestString("schema:" + tools[i].Name)
		catalog := "catalog-v1"
		tools[i].SchemaDigest = &schema
		tools[i].CatalogVersion = &catalog
		if strings.HasPrefix(tools[i].Name, "relay.") {
			mapping := "mapping-v1"
			tools[i].MappingVersion = &mapping
		}
	}
	return tools
}

func hexDigest(value []byte) string {
	digest := sha256.Sum256(value)
	return hex.EncodeToString(digest[:])
}

func digestString(value string) string { return hexDigest([]byte(value)) }
