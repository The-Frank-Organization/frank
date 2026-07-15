package config_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
)

func TestLaneVCSV3LoadAndClone(t *testing.T) {
	doc := laneVCSEngineDoc(t, 3, map[string]any{"repo": "git"}, true)
	pinned, err := loadLaneVCSDoc(t, doc)
	if err != nil {
		t.Fatalf("Load v3: %v", err)
	}
	if pinned.Supply.LaneVCS["repo"] != "git" || pinned.Engine.Supply.LaneVCS["repo"] != "git" {
		t.Fatalf("LaneVCS not retained: pinned=%#v engine=%#v", pinned.Supply.LaneVCS, pinned.Engine.Supply.LaneVCS)
	}
	pinned.Supply.LaneVCS["repo"] = "none"
	if pinned.Engine.Supply.LaneVCS["repo"] != "git" {
		t.Fatalf("LaneVCS maps alias: %#v %#v", pinned.Supply.LaneVCS, pinned.Engine.Supply.LaneVCS)
	}
}

func TestLaneVCSSchemaAndCompositionNegatives(t *testing.T) {
	for _, tc := range []struct {
		name    string
		version int
		vcs     any
		include bool
	}{
		{name: "v3 absent", version: 3},
		{name: "v3 invalid enum", version: 3, vcs: map[string]any{"repo": "svn"}, include: true},
		{name: "v3 extra lane", version: 3, vcs: map[string]any{"repo": "git", "other": "none"}, include: true},
		{name: "v3 missing lane", version: 3, vcs: map[string]any{}, include: true},
		{name: "v2 smuggles key", version: 2, vcs: map[string]any{"repo": "git"}, include: true},
		{name: "v3 array", version: 3, vcs: []any{"git"}, include: true},
		{name: "v3 numeric value", version: 3, vcs: map[string]any{"repo": 3}, include: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := loadLaneVCSDoc(t, laneVCSEngineDoc(t, tc.version, tc.vcs, tc.include)); !errors.Is(err, config.ErrConfigLoad) {
				t.Fatalf("err=%v, want ErrConfigLoad", err)
			}
		})
	}
}

func TestLaneVCSV2ResidencyAndTransitions(t *testing.T) {
	v2 := laneVCSEngineDoc(t, 2, nil, false)
	pinned, err := loadLaneVCSDoc(t, v2)
	if err != nil {
		t.Fatalf("v2 residency load: %v", err)
	}
	if pinned.Supply.LaneVCS != nil {
		t.Fatalf("v2 residency pinned=%#v err=%v", pinned.Supply, err)
	}
	v3 := laneVCSEngineDoc(t, 3, map[string]any{"repo": "git"}, true)
	if err := config.ValidateMemberTransition("engine", v2, v3); err != nil {
		t.Fatalf("v2->v3: %v", err)
	}
	if err := config.ValidateMemberTransition("engine", v3, v2); !errors.Is(err, config.ErrConfigVersionTransition) {
		t.Fatalf("v3->v2 err=%v", err)
	}
	v1 := []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`)
	if err := config.ValidateMemberTransition("engine", v1, v3); !errors.Is(err, config.ErrConfigVersionTransition) {
		t.Fatalf("v1->v3 err=%v", err)
	}
	v2Smuggled := laneVCSEngineDoc(t, 2, map[string]any{"repo": "git"}, true)
	if err := config.ValidateMemberTransition("engine", v2, v2Smuggled); !errors.Is(err, config.ErrConfigVersionTransition) {
		t.Fatalf("v2 smuggled candidate err=%v", err)
	}
}

func TestLaneVCSReaderCeilingRefusesV5BeforeSupply(t *testing.T) {
	doc := []byte(`{"version":5,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false},"supply":"deliberately-garbage"}`)
	_, err := loadLaneVCSDoc(t, doc)
	if !errors.Is(err, config.ErrConfigLoad) || !strings.Contains(err.Error(), "engine-marker") {
		t.Fatalf("err=%v, want phase-0 engine-marker ErrConfigLoad", err)
	}
}

func loadLaneVCSDoc(t *testing.T, doc []byte) (*config.Pinned, error) {
	t.Helper()
	root := t.TempDir()
	engine := filepath.Join(root, "engine.json")
	if err := os.WriteFile(engine, doc, 0o600); err != nil {
		t.Fatal(err)
	}
	return config.Load(map[string]string{"engine": engine, "fieldspec": filepath.Join("..", "fieldspec", "registry.json")})
}

func laneVCSEngineDoc(t *testing.T, version int, vcs any, include bool) []byte {
	t.Helper()
	root := t.TempDir()
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	command := filepath.Join(root, "suite.sh")
	if err := os.WriteFile(command, []byte("#!/bin/sh\nexit 0\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	supply := map[string]any{
		"lane_roots": map[string]any{"repo": root}, "schema_refs": map[string]any{},
		"suites": map[string]any{"suite": map[string]any{
			"lane": "repo", "command": "suite.sh", "args": []any{}, "timeout_class": "suite_bounded", "timeout_seconds": 1,
		}},
	}
	if include {
		supply["lane_vcs"] = vcs
	}
	doc := map[string]any{
		"version": version, "gc_enabled": false, "segment_rotate_bytes": 4194304,
		"present_layers": map[string]any{"observe": false}, "supply": supply,
	}
	raw, err := json.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}
