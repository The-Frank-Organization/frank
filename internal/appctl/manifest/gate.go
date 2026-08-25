package manifest

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/jackli/frank/internal/appipc"
)

var (
	ErrIdentityUnlocked  = errors.New("app control manifest: identity not locked")
	ErrServeGate         = errors.New("app control manifest: serve gate refused")
	ErrDeniedAboveSet    = errors.New("app control manifest: denied above set")
	ErrManifestIntegrity = errors.New("app control manifest: stored bytes mismatch")
)

type Gate struct {
	LockedTools              []ToolIdentity
	ShippedToolCatalogDigest string
	PolicyBytes              []byte
	LaneCatalogDigest        string
}

func (gate Gate) Validate(frozen Frozen) error {
	if digestBytes(frozen.Bytes) != frozen.Digest {
		return ErrManifestIntegrity
	}
	encoded, err := appCanonical(frozen.Manifest)
	if err != nil || !bytes.Equal(encoded, frozen.Bytes) {
		return ErrManifestIntegrity
	}
	manifest := frozen.Manifest
	if manifest.ManifestSchema != ManifestSchema || manifest.RunID == "" || !equalStrings(ToolNames(manifest.ToolSet), RatifiedToolNames) {
		return ErrServeGate
	}
	if manifest.ToolCatalogDigest == nil || manifest.ReleaseBinding == nil {
		return ErrIdentityUnlocked
	}
	for _, tool := range manifest.ToolSet {
		if tool.SchemaDigest == nil || tool.CatalogVersion == nil || (isRelayTool(tool.Name) && tool.MappingVersion == nil) {
			return ErrIdentityUnlocked
		}
		if !isRelayTool(tool.Name) && tool.MappingVersion != nil {
			return ErrServeGate
		}
	}
	locked := cloneTools(gate.LockedTools)
	manifestTools := cloneTools(manifest.ToolSet)
	sortTools(locked)
	sortTools(manifestTools)
	if !reflect.DeepEqual(manifestTools, locked) {
		return ErrServeGate
	}
	if *manifest.ToolCatalogDigest != gate.ShippedToolCatalogDigest || !validDigest(gate.ShippedToolCatalogDigest) {
		return ErrServeGate
	}
	if digestBytes(gate.PolicyBytes) != manifest.PolicyDigest {
		return ErrServeGate
	}
	if manifest.ProviderLane.LaneCatalogDigest != gate.LaneCatalogDigest {
		return ErrServeGate
	}
	if digestBytes([]byte(manifest.WorkspaceRootPath)) != manifest.WorkspaceRootID {
		return ErrServeGate
	}
	if err := validateRelease(*manifest.ReleaseBinding); err != nil {
		return ErrServeGate
	}
	return nil
}

func (gate Gate) Authorize(frozen Frozen, canonicalName string) error {
	if err := gate.Validate(frozen); err != nil {
		return err
	}
	for _, name := range RatifiedToolNames {
		if canonicalName == name {
			return nil
		}
	}
	return ErrDeniedAboveSet
}

func appCanonical(manifest Manifest) ([]byte, error) {
	return appipcMarshal(canonicalValue(manifest))
}

func isRelayTool(name string) bool {
	return len(name) > len("relay.") && name[:len("relay.")] == "relay."
}

func sortTools(tools []ToolIdentity) {
	sort.Slice(tools, func(i, j int) bool { return tools[i].Name < tools[j].Name })
}

// Kept as a narrow seam so gate validation and construction share exactly the
// same canonical encoder.
func appipcMarshal(value any) ([]byte, error) {
	encoded, err := appipc.MarshalJCS(value)
	if err != nil {
		return nil, fmt.Errorf("manifest canonical encoding: %w", err)
	}
	return encoded, nil
}
