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
var ErrConfigVersionTransition = errors.New("config-version-transition")

type Pinned struct {
	Members  map[string][]byte
	Engine   EngineConfig
	Registry *fieldspec.Registry
	Digest   string
}

type EngineConfig struct {
	Version            int             `json:"version,omitempty"`
	GCEnabled          bool            `json:"gc_enabled"`
	SegmentRotateBytes int64           `json:"segment_rotate_bytes"`
	MaxFrameBytes      int             `json:"max_frame_bytes,omitempty"`
	PresentLayers      map[string]bool `json:"present_layers,omitempty"`
}

func (c EngineConfig) FrameBytes() int {
	if c.MaxFrameBytes > 0 {
		return c.MaxFrameBytes
	}
	return 1 << 20
}

// PresentLayers returns the immutable-by-convention predicate context for one
// pinned config generation. Core layers are always present; optional layers
// are copied from the engine member so callers cannot mutate pinned config.
func PresentLayers(pinned *Pinned) map[string]bool {
	layers := map[string]bool{
		"store":   true,
		"form":    true,
		"lineage": true,
	}
	if pinned == nil {
		return layers
	}
	for name, present := range pinned.Engine.PresentLayers {
		if present {
			layers[name] = true
		}
	}
	return layers
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

// ValidateMemberTransition is the acceptance-time, forward-only member gate.
// It validates the candidate against the schema named by its marker before
// checking the owner-supplied successor relation. Error text is deliberately
// symbolic and path-free for the submit bounce surface.
func ValidateMemberTransition(member string, current, candidate []byte) error {
	switch member {
	case "engine":
		currentVersion, err := engineVersion(current)
		if err != nil {
			return ErrConfigVersionTransition
		}
		candidateVersion, err := validateEngineSchema(candidate)
		if err != nil {
			return ErrConfigVersionTransition
		}
		if candidateVersion == currentVersion || candidateVersion == currentVersion+1 && candidateVersion == 1 {
			return nil
		}
		return ErrConfigVersionTransition
	case "fieldspec":
		return validateStringMarkerTransition(current, candidate, "s7a-fieldspec-v5", "s8-fieldspec-v6")
	case "catalog":
		return validateStringMarkerTransition(current, candidate, "s7-v1", "s8-v1")
	default:
		return ErrConfigVersionTransition
	}
}

func engineVersion(data []byte) (int, error) {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return 0, err
	}
	value, ok := raw["version"]
	if !ok {
		return 0, nil
	}
	number, ok := value.(float64)
	if !ok || number < 0 || number != float64(int(number)) {
		return 0, ErrConfigVersionTransition
	}
	return int(number), nil
}

func validateEngineSchema(data []byte) (int, error) {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return 0, err
	}
	version, err := engineVersion(data)
	if err != nil || version < 0 || version > 1 {
		return 0, ErrConfigVersionTransition
	}
	allowed := map[string]string{
		"gc_enabled":           "bool",
		"segment_rotate_bytes": "number",
		"max_frame_bytes":      "number",
		"detector":             "object",
	}
	if version == 1 {
		allowed["version"] = "number"
		allowed["present_layers"] = "layers"
	}
	for key, value := range raw {
		kind, ok := allowed[key]
		if !ok || !matchesConfigKind(kind, value) {
			return 0, ErrConfigVersionTransition
		}
	}
	for _, required := range []string{"gc_enabled", "segment_rotate_bytes"} {
		if _, ok := raw[required]; !ok {
			return 0, ErrConfigVersionTransition
		}
	}
	if version == 1 {
		if _, ok := raw["version"]; !ok {
			return 0, ErrConfigVersionTransition
		}
		if _, ok := raw["present_layers"]; !ok {
			return 0, ErrConfigVersionTransition
		}
	}
	return version, nil
}

func matchesConfigKind(kind string, value any) bool {
	switch kind {
	case "bool":
		_, ok := value.(bool)
		return ok
	case "number":
		number, ok := value.(float64)
		return ok && number == float64(int64(number))
	case "object":
		_, ok := value.(map[string]any)
		return ok
	case "layers":
		layers, ok := value.(map[string]any)
		if !ok {
			return false
		}
		for name, present := range layers {
			if name != "observe" {
				return false
			}
			if _, ok := present.(bool); !ok {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func validateStringMarkerTransition(current, candidate []byte, from, to string) error {
	currentMarker, err := stringMarker(current)
	if err != nil {
		return ErrConfigVersionTransition
	}
	candidateMarker, err := stringMarker(candidate)
	if err != nil {
		return ErrConfigVersionTransition
	}
	if candidateMarker == currentMarker || currentMarker == from && candidateMarker == to {
		return nil
	}
	return ErrConfigVersionTransition
}

func stringMarker(data []byte) (string, error) {
	var raw struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &raw); err != nil || raw.Version == "" {
		return "", ErrConfigVersionTransition
	}
	return raw.Version, nil
}
