// Package catalog defines the m-9-owned F58 tool identity catalog.
package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/worker/jcs"
)

const (
	CatalogVersion = "m9-catalog-v1"
	MappingVersion = "m2-mapping-v1"

	RelaySubmitSchemaDigest  = "6bb7bbf46d8bf5d210cee410fbd0fa59106145425878c065adf0d54b05ace08e"
	RelayProjectSchemaDigest = "be5c41ec848bd7f6a7afd16af5acc56c65cf39bc113041941bb6747153bd582a"
	RelayReadSchemaDigest    = "a84645cb3f57ea1172661ddcc42e8a710f5a320ee3ed6c944f5e469026b3036e"

	ExpectedDigest = "151a7e614abd6b25e643062b26cb9c6af60b0eadedf8e03c1f12b1850458913d"
)

// Identity is the complete manifest-tool identity vector compared by the m-10
// serve gate. MappingVersion is absent from JSON for local tools.
type Identity struct {
	CanonicalName            string `json:"name"`
	FormSchemaMappingVersion string `json:"mapping_version,omitempty"`
	ToolImplCatalogVersion   string `json:"catalog_version"`
	ToolSchemaDigest         string `json:"schema_digest"`
}

var localSchemas = map[string]json.RawMessage{
	"apply_patch": json.RawMessage(`{"additionalProperties":false,"properties":{"patch":{"type":"string"}},"required":["patch"],"type":"object"}`),
	"bash":        json.RawMessage(`{"additionalProperties":false,"properties":{"command":{"type":"string"},"timeout_ms":{"type":"integer"}},"required":["command"],"type":"object"}`),
	"edit":        json.RawMessage(`{"additionalProperties":false,"properties":{"new_string":{"type":"string"},"old_string":{"type":"string"},"path":{"type":"string"},"replace_all":{"type":"boolean"}},"required":["path","old_string","new_string"],"type":"object"}`),
	"read":        json.RawMessage(`{"additionalProperties":false,"properties":{"limit":{"type":"integer"},"offset":{"type":"integer"},"path":{"type":"string"}},"required":["path"],"type":"object"}`),
	"write":       json.RawMessage(`{"additionalProperties":false,"properties":{"content":{"type":"string"},"path":{"type":"string"}},"required":["path","content"],"type":"object"}`),
}

var expectedIdentities = []Identity{
	{CanonicalName: "apply_patch", ToolSchemaDigest: "f6594e18aaedfb029106fa669c557027854ec5f86cce436fcec1723791743cd7", ToolImplCatalogVersion: CatalogVersion},
	{CanonicalName: "bash", ToolSchemaDigest: "ddd9efb1ca815ee8a71ffda66eefbd627f1b4be32c9955a23eb874ea9cba5271", ToolImplCatalogVersion: CatalogVersion},
	{CanonicalName: "edit", ToolSchemaDigest: "396e58a8583dc71e366be6cb516258efcc1f666810bcdebee81bcb1fdb9d6f4d", ToolImplCatalogVersion: CatalogVersion},
	{CanonicalName: "read", ToolSchemaDigest: "4dc4e27079b79febaa3ff2b1a91c90df7cbabd864868d4c052cee2d3b2903356", ToolImplCatalogVersion: CatalogVersion},
	{CanonicalName: "relay.project", ToolSchemaDigest: RelayProjectSchemaDigest, ToolImplCatalogVersion: CatalogVersion, FormSchemaMappingVersion: MappingVersion},
	{CanonicalName: "relay.read", ToolSchemaDigest: RelayReadSchemaDigest, ToolImplCatalogVersion: CatalogVersion, FormSchemaMappingVersion: MappingVersion},
	{CanonicalName: "relay.submit", ToolSchemaDigest: RelaySubmitSchemaDigest, ToolImplCatalogVersion: CatalogVersion, FormSchemaMappingVersion: MappingVersion},
	{CanonicalName: "write", ToolSchemaDigest: "0863ca49d0670920e458e08df8a8547017f99b2f09b02375d0bdf10f53e24a7b", ToolImplCatalogVersion: CatalogVersion},
}

// ExpectedIdentities returns a copy of the locked, canonical-name-sorted vector.
func ExpectedIdentities() []Identity {
	return append([]Identity(nil), expectedIdentities...)
}

// LocalSchemas returns copies of the five design-fixed local-tool schemas.
func LocalSchemas() map[string]json.RawMessage {
	result := make(map[string]json.RawMessage, len(localSchemas))
	for name, schema := range localSchemas {
		result[name] = append(json.RawMessage(nil), schema...)
	}
	return result
}

// SchemaDigest computes the canonical JSON SHA-256 identity of an input schema.
func SchemaDigest(schema []byte) (string, error) {
	return jcs.Digest(schema)
}

// Digest sorts a complete identity vector by canonical name and hashes its JCS
// encoding. Argument schemas are never inspected here.
func Digest(identities []Identity) (string, error) {
	canonical, err := normalized(identities)
	if err != nil {
		return "", err
	}
	encoded, err := json.Marshal(canonical)
	if err != nil {
		return "", fmt.Errorf("marshal catalog: %w", err)
	}
	return jcs.Digest(encoded)
}

// EqualIdentitySet applies exact canonical-set equality to both vectors.
func EqualIdentitySet(left, right []Identity) error {
	normalizedLeft, err := normalized(left)
	if err != nil {
		return fmt.Errorf("left identity set: %w", err)
	}
	normalizedRight, err := normalized(right)
	if err != nil {
		return fmt.Errorf("right identity set: %w", err)
	}
	if len(normalizedLeft) != len(normalizedRight) {
		return fmt.Errorf("identity count mismatch: %d != %d", len(normalizedLeft), len(normalizedRight))
	}
	for index := range normalizedLeft {
		if normalizedLeft[index] != normalizedRight[index] {
			return fmt.Errorf("identity mismatch at %q", normalizedLeft[index].CanonicalName)
		}
	}
	return nil
}

func normalized(identities []Identity) ([]Identity, error) {
	if identities == nil {
		return nil, errors.New("identity set is absent")
	}
	result := append([]Identity(nil), identities...)
	sort.Slice(result, func(i, j int) bool {
		return result[i].CanonicalName < result[j].CanonicalName
	})
	for index, identity := range result {
		if err := validateIdentity(identity); err != nil {
			return nil, fmt.Errorf("identity %d: %w", index, err)
		}
		if index > 0 && result[index-1].CanonicalName == identity.CanonicalName {
			return nil, fmt.Errorf("duplicate canonical name %q", identity.CanonicalName)
		}
	}
	return result, nil
}

func validateIdentity(identity Identity) error {
	if identity.CanonicalName == "" {
		return errors.New("canonical name is empty")
	}
	if !isLowerHexDigest(identity.ToolSchemaDigest) {
		return fmt.Errorf("tool schema digest for %q is not lowercase SHA-256", identity.CanonicalName)
	}
	if identity.ToolImplCatalogVersion == "" {
		return fmt.Errorf("tool catalog version for %q is empty", identity.CanonicalName)
	}
	if strings.HasPrefix(identity.CanonicalName, "relay.") {
		if identity.FormSchemaMappingVersion == "" {
			return fmt.Errorf("relay identity %q lacks mapping version", identity.CanonicalName)
		}
	} else if identity.FormSchemaMappingVersion != "" {
		return fmt.Errorf("local identity %q carries a mapping version", identity.CanonicalName)
	}
	return nil
}

func isLowerHexDigest(value string) bool {
	if len(value) != 64 {
		return false
	}
	for _, character := range value {
		if character < '0' || character > '9' {
			if character < 'a' || character > 'f' {
				return false
			}
		}
	}
	return true
}
