package fixtures_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/executor"
	"github.com/jackli/frank/internal/observe"
)

func TestS8FXSUP1And3EngineV2SupplyIsClosedAndForwardOnly(t *testing.T) {
	root := t.TempDir()
	root = s8CanonicalPath(t, root)
	s8WriteExecutable(t, root, "suite.sh", "#!/bin/sh\nexit 0\n")
	if err := os.WriteFile(filepath.Join(root, "not-executable.sh"), []byte("#!/bin/sh\nexit 0\n"), 0o644); err != nil {
		t.Fatalf("write non-executable command: %v", err)
	}
	if err := os.Mkdir(filepath.Join(root, "command-dir"), 0o755); err != nil {
		t.Fatalf("write command directory: %v", err)
	}
	if err := os.Symlink(filepath.Join(root, "suite.sh"), filepath.Join(root, "suite-link.sh")); err != nil {
		t.Fatalf("write command symlink: %v", err)
	}
	rootLink := filepath.Join(t.TempDir(), "root-link")
	if err := os.Symlink(root, rootLink); err != nil {
		t.Fatalf("write root symlink: %v", err)
	}
	valid := s8SupplyEngineBytes(t, root, map[string]string{"schema-a": strings.Repeat("a", 64)}, 120)
	path := filepath.Join(t.TempDir(), "engine.json")
	if err := os.WriteFile(path, valid, 0o644); err != nil {
		t.Fatalf("write engine v2: %v", err)
	}
	sources := s8ConfigSources(t, false)
	sources["engine"] = path
	pinned, err := config.Load(sources)
	if err != nil {
		t.Fatalf("load valid supply: %v", err)
	}
	if pinned.Engine.Version != 2 || pinned.Supply == nil || pinned.Supply.LaneRoots["repo"] != root {
		t.Fatalf("pinned supply = engine %d supply %#v", pinned.Engine.Version, pinned.Supply)
	}
	if got := pinned.Supply.Suites["dogfood-battery"]; got.Lane != "repo" || got.Command != "suite.sh" || got.TimeoutClass != "suite_bounded" || got.Timeout != 120*time.Second {
		t.Fatalf("dogfood descriptor = %#v", got)
	}
	pinned.Engine.Supply.LaneRoots["repo"] = "mutated"
	if pinned.Supply.LaneRoots["repo"] != root {
		t.Fatal("Pinned.Supply aliases mutable Engine.Supply")
	}
	t.Logf("engine_v1_to_v2_candidate_config_digest=%s", pinned.Digest)

	v1 := []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`)
	if err := config.ValidateMemberTransition("engine", v1, valid); err != nil {
		t.Fatalf("v1-to-v2 transition: %v", err)
	}
	v0 := []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`)
	if err := config.ValidateMemberTransition("engine", v0, valid); !errors.Is(err, config.ErrConfigVersionTransition) {
		t.Fatalf("v0-to-v2 skip = %v, want config-version-transition", err)
	}

	for _, tc := range []struct {
		name   string
		mutate func(map[string]any)
	}{
		{name: "missing supply", mutate: func(doc map[string]any) { delete(doc, "supply") }},
		{name: "dangling lane", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["lane"] = "absent"
		}},
		{name: "zero timeout", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["timeout_seconds"] = float64(0)
		}},
		{name: "over ceiling", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["timeout_seconds"] = float64(121)
		}},
		{name: "class mismatch", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["timeout_class"] = "read_short"
		}},
		{name: "malformed schema digest", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["schema_refs"].(map[string]any)["schema-a"] = "ABC"
		}},
		{name: "path arg", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["args"] = []any{"/config/engine.json"}
		}},
		{name: "effective class arg", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["args"] = []any{"suite_bounded"}
		}},
		{name: "effective timeout arg", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["args"] = []any{"120"}
		}},
		{name: "secret arg", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["args"] = []any{"credential-token"}
		}},
		{name: "escaping command", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["command"] = "../suite.sh"
		}},
		{name: "alternate separator command", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["command"] = `..\suite.sh`
		}},
		{name: "missing command", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["command"] = "missing.sh"
		}},
		{name: "non-executable command", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["command"] = "not-executable.sh"
		}},
		{name: "directory command", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["command"] = "command-dir"
		}},
		{name: "symlinked command", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["suites"].(map[string]any)["dogfood-battery"].(map[string]any)["command"] = "suite-link.sh"
		}},
		{name: "symlinked root", mutate: func(doc map[string]any) {
			doc["supply"].(map[string]any)["lane_roots"].(map[string]any)["repo"] = rootLink
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var doc map[string]any
			if err := json.Unmarshal(valid, &doc); err != nil {
				t.Fatalf("decode valid engine: %v", err)
			}
			tc.mutate(doc)
			candidate, err := json.Marshal(doc)
			if err != nil {
				t.Fatalf("marshal mutation: %v", err)
			}
			if err := os.WriteFile(path, candidate, 0o644); err != nil {
				t.Fatalf("write mutation: %v", err)
			}
			if _, err := config.Load(sources); !errors.Is(err, config.ErrConfigLoad) {
				t.Fatalf("load malformed supply = %v, want config-load", err)
			}
		})
	}
}

func TestS8FXSUP3V1ReaderRefusesV2BeforeSupplyInterpretation(t *testing.T) {
	planted := []byte(`{"version":2,"supply":"content-must-not-be-interpreted"}`)
	if err := config.ValidateEngineReaderMarker(planted, 1); !errors.Is(err, config.ErrConfigLoad) {
		t.Fatalf("v1 reader marker error = %v, want config-load", err)
	}
	if err := config.ValidateEngineReaderMarker(planted, 2); err != nil {
		t.Fatalf("v2 marker preflight interpreted planted supply: %v", err)
	}
}

func TestS8FXSUP3DuplicateAndTrailingMarkersRefuseBeforeInterpretation(t *testing.T) {
	for name, data := range map[string][]byte{
		"engine duplicate": []byte(`{"version":1,"content":"not-interpreted","version":2}`),
		"engine trailing":  []byte(`{"version":1} {"version":2}`),
	} {
		t.Run(name, func(t *testing.T) {
			if err := config.ValidateEngineReaderMarker(data, 2); !errors.Is(err, config.ErrConfigLoad) {
				t.Fatalf("marker error = %v, want config-load", err)
			}
		})
	}
	if err := config.ValidateFieldspecReaderMarker([]byte(`{"version":"s8-fieldspec-v6","version":"s8-fieldspec-v7"}`), "s8-fieldspec-v6", "s8-fieldspec-v7"); !errors.Is(err, config.ErrConfigLoad) {
		t.Fatalf("duplicate fieldspec marker = %v, want config-load", err)
	}
	sources := s8ConfigSources(t, false)
	catalog, err := os.ReadFile(sources["catalog"])
	if err != nil {
		t.Fatalf("read catalog: %v", err)
	}
	duplicate := bytes.Replace(catalog, []byte(`"version": "s8-v1"`), []byte(`"version": "s8-v1", "version": "s8-v1"`), 1)
	if err := os.WriteFile(sources["catalog"], duplicate, 0o644); err != nil {
		t.Fatalf("write duplicate catalog marker: %v", err)
	}
	if _, err := config.Load(sources); !errors.Is(err, config.ErrConfigLoad) {
		t.Fatalf("duplicate catalog marker load = %v, want config-load", err)
	}
}

func TestS8FXSUP3MemberTransitionRefusesDuplicateMarkers(t *testing.T) {
	sources := s8ConfigSources(t, false)
	for _, member := range []string{"engine", "fieldspec", "catalog"} {
		t.Run(member, func(t *testing.T) {
			current, err := os.ReadFile(sources[member])
			if err != nil {
				t.Fatalf("read %s: %v", member, err)
			}
			var candidate []byte
			switch member {
			case "engine":
				candidate = bytes.Replace(current, []byte(`"version":2`), []byte(`"version":2,"version":2`), 1)
			case "fieldspec":
				candidate = bytes.Replace(current, []byte(`"version": "s10-fieldspec-v8"`), []byte(`"version": "s10-fieldspec-v8", "version": "s10-fieldspec-v8"`), 1)
			case "catalog":
				candidate = bytes.Replace(current, []byte(`"version": "s8-v1"`), []byte(`"version": "s8-v1", "version": "s8-v1"`), 1)
			}
			if bytes.Equal(candidate, current) {
				t.Fatalf("%s marker mutation did not apply", member)
			}
			if err := config.ValidateMemberTransition(member, current, candidate); !errors.Is(err, config.ErrConfigVersionTransition) {
				t.Fatalf("duplicate %s transition = %v, want config-version-transition", member, err)
			}
		})
	}
}

func TestS8FXSUP5ProductionHasNoAmbientOrEmptySupplyFallback(t *testing.T) {
	mainBytes, err := os.ReadFile(filepath.Join("..", "..", "cmd", "frank", "main.go"))
	if err != nil {
		t.Fatalf("read production composition: %v", err)
	}
	text := string(mainBytes)
	for _, forbidden := range []string{"os.Getwd(", "Suites: map[string]executor.Suite{}", "NamedSuites: map[string]bool{}", "SchemaRefs: map[string]string{}"} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("production composition retains fallback %q", forbidden)
		}
	}
	for _, required := range []string{"pinned.Supply.LaneRoots", "pinned.Supply.SchemaRefs", "pinned.Supply.Suites"} {
		if !strings.Contains(text, required) {
			t.Fatalf("production composition omits governed source %q", required)
		}
	}
}

func TestS8FXSUP2ExecutorRejectsRuntimeTimeoutClassMismatch(t *testing.T) {
	root := t.TempDir()
	root = s8CanonicalPath(t, root)
	s8WriteExecutable(t, root, "suite.sh", "#!/bin/sh\nexit 0\n")
	host := executor.New(executor.Config{Suites: map[string]executor.Suite{
		"dogfood-battery": {SourceDir: root, Command: "suite.sh", TimeoutClass: "suite_bounded", Timeout: time.Second},
	}})
	selection := observe.Selection{CheckID: "run-suite", ClaimRef: "mismatch", Params: map[string]string{"target": "dogfood-battery", "expect_green": "true"}}
	verdict := host.Spawn(observe.CheckEntry{ID: "run-suite", Class: "suite", ExecutorRequired: true, TimeoutClass: "read_short"}, selection)
	if verdict.Outcome != "unsafe" || verdict.FailingDetail != "executor-timeout-class-mismatch" {
		t.Fatalf("timeout-class mismatch verdict = %#v", verdict)
	}
}

func TestS8FXSUP6SchemaRefThreeWaySplit(t *testing.T) {
	root := t.TempDir()
	content := []byte("governed schema\n")
	if err := os.WriteFile(filepath.Join(root, "schema.txt"), content, 0o644); err != nil {
		t.Fatalf("write schema: %v", err)
	}
	reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"repo": root}, SchemaRefs: map[string]string{
		"matching": s8SHA256(content), "different": strings.Repeat("0", 64),
	}})
	for _, tc := range []struct {
		name, ref, outcome, detail string
	}{
		{name: "match", ref: "matching", outcome: "pass"},
		{name: "content mismatch", ref: "different", outcome: "fail", detail: "read-file-mismatch"},
		{name: "unknown id", ref: "absent", outcome: "unsafe", detail: "schema-ref-unknown"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			verdict := reg.Run(observe.Selection{CheckID: "read-file", ClaimRef: tc.name, Params: map[string]string{
				"lane_ref": "repo", "path": "schema.txt", "expect": "schema_ref:" + tc.ref,
			}})
			if verdict.Outcome != tc.outcome || verdict.FailingDetail != tc.detail {
				t.Fatalf("schema-ref verdict = %#v", verdict)
			}
		})
	}
}

func s8SupplyEngineBytes(t *testing.T, root string, schemaRefs map[string]string, timeout int) []byte {
	t.Helper()
	raw, err := json.Marshal(map[string]any{
		"version": 2, "gc_enabled": false, "segment_rotate_bytes": 4194304,
		"present_layers": map[string]bool{"observe": false},
		"supply": map[string]any{
			"lane_roots": map[string]string{"repo": root}, "schema_refs": schemaRefs,
			"suites": map[string]any{"dogfood-battery": map[string]any{
				"lane": "repo", "command": "suite.sh", "args": []string{},
				"timeout_class": "suite_bounded", "timeout_seconds": timeout,
			}},
		},
	})
	if err != nil {
		t.Fatalf("marshal supply engine: %v", err)
	}
	return raw
}

func s8SHA256(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func s8CanonicalPath(t *testing.T, path string) string {
	t.Helper()
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		t.Fatalf("canonicalize %s: %v", path, err)
	}
	return resolved
}
