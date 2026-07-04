package config

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/jackli/frank/internal/fieldspec"
)

var ErrMissingEngine = errors.New("missing engine config member")

type Pinned struct {
	Members  map[string][]byte
	Engine   EngineConfig
	Registry *fieldspec.Registry
	Digest   string
}

type EngineConfig struct {
	GCEnabled          bool  `json:"gc_enabled"`
	SegmentRotateBytes int64 `json:"segment_rotate_bytes"`
}

func Load(members map[string]string) (*Pinned, error) {
	enginePath, ok := members["engine"]
	if !ok || enginePath == "" {
		return nil, ErrMissingEngine
	}

	loaded := make(map[string][]byte, len(members))
	for name, path := range members {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("load config member %s: %w", name, err)
		}
		loaded[name] = append([]byte(nil), data...)
	}

	var engine EngineConfig
	if err := json.Unmarshal(loaded["engine"], &engine); err != nil {
		return nil, fmt.Errorf("parse engine config: %w", err)
	}

	var registry *fieldspec.Registry
	if registryPath := members["fieldspec"]; registryPath != "" {
		parsed, err := fieldspec.Load(registryPath)
		if err != nil {
			return nil, fmt.Errorf("load fieldspec registry: %w", err)
		}
		registry = parsed
	}

	return &Pinned{
		Members:  loaded,
		Engine:   engine,
		Registry: registry,
		Digest:   Digest(loaded),
	}, nil
}

func Digest(members map[string][]byte) string {
	names := make([]string, 0, len(members))
	for name := range members {
		names = append(names, name)
	}
	sort.Strings(names)

	var manifest []byte
	for _, name := range names {
		memberSum := sha256.Sum256(members[name])
		manifest = append(manifest, name...)
		manifest = append(manifest, 0)
		manifest = append(manifest, hex.EncodeToString(memberSum[:])...)
		manifest = append(manifest, '\n')
	}
	sum := sha256.Sum256(manifest)
	return hex.EncodeToString(sum[:])
}
