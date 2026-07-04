package config_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
)

func TestDigestIsStableByMemberNameAndBytes(t *testing.T) {
	membersA := map[string][]byte{
		"engine":    []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`),
		"fieldspec": []byte(`{"phase":["SITREP"]}`),
	}
	membersB := map[string][]byte{
		"fieldspec": []byte(`{"phase":["SITREP"]}`),
		"engine":    []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`),
	}

	digestA := config.Digest(membersA)
	digestB := config.Digest(membersB)
	if digestA != digestB {
		t.Fatalf("Digest changed with map order: %q != %q", digestA, digestB)
	}

	membersB["fieldspec"] = []byte(`{"phase":["PLAN"]}`)
	if got := config.Digest(membersB); got == digestA {
		t.Fatalf("Digest did not change after member byte flip: %q", got)
	}
}

func TestLoadParsesEngineConfigAndStoresMemberBytes(t *testing.T) {
	root := t.TempDir()
	enginePath := writeFile(t, root, "engine.json", `{"gc_enabled":false,"segment_rotate_bytes":4194304}`)
	registryPath := writeFile(t, root, "registry.json", `{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`)

	pinned, err := config.Load(map[string]string{
		"engine":    enginePath,
		"fieldspec": registryPath,
	})
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if pinned.Engine.GCEnabled {
		t.Fatalf("GCEnabled = true, want false")
	}
	if pinned.Engine.SegmentRotateBytes != 4194304 {
		t.Fatalf("SegmentRotateBytes = %d, want 4194304", pinned.Engine.SegmentRotateBytes)
	}
	if string(pinned.Members["engine"]) != `{"gc_enabled":false,"segment_rotate_bytes":4194304}` {
		t.Fatalf("engine member bytes not preserved: %q", string(pinned.Members["engine"]))
	}
	if pinned.Digest == "" {
		t.Fatalf("Digest empty")
	}
	for _, path := range []string{enginePath, registryPath} {
		if strings.Contains(pinned.Digest, path) {
			t.Fatalf("Digest leaked member path %q in %q", path, pinned.Digest)
		}
	}
}

func TestLoadRequiresEngineMember(t *testing.T) {
	root := t.TempDir()
	registryPath := writeFile(t, root, "registry.json", `{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`)

	_, err := config.Load(map[string]string{"fieldspec": registryPath})
	if !errors.Is(err, config.ErrMissingEngine) {
		t.Fatalf("Load err = %v, want ErrMissingEngine", err)
	}
}

func writeFile(t *testing.T, root, name, body string) string {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
	return path
}
