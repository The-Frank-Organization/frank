package manifest

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/appipc"
)

const ManifestSchema = 1

var RatifiedToolNames = []string{
	"apply_patch", "bash", "edit", "read", "relay.project", "relay.read", "relay.submit", "write",
}

var (
	ErrMalformedManifest    = errors.New("app control manifest: malformed")
	ErrPolicyLaneMismatch   = errors.New("app control manifest: policy lane mismatch")
	ErrPolicyDigestMismatch = errors.New("app control manifest: policy digest mismatch")
)

type LaneID struct {
	ModelID          string `json:"model_id"`
	ProviderID       string `json:"provider_id"`
	ServingProfileID string `json:"serving_profile_id"`
	CompatMode       string `json:"compat_mode"`
}

type ProviderLane struct {
	LaneID            LaneID `json:"lane_id"`
	LaneCatalogDigest string `json:"lane_catalog_digest"`
	CredentialRef     string `json:"credential_ref"`
}

type ToolIdentity struct {
	Name           string
	SchemaDigest   *string
	CatalogVersion *string
	MappingVersion *string
}

type BuildDigests struct {
	AppMainBuildDigest  string `json:"app_main_build_digest"`
	M9WorkerBuildDigest string `json:"m9_worker_build_digest"`
	M8BuildDigest       string `json:"m8_build_digest"`
}

type ReleaseBinding struct {
	BoundAtRef    string
	BuildDigests  *BuildDigests
	ReleaseDigest *string
}

type Manifest struct {
	ManifestSchema    int
	RunID             string
	PolicySourceRef   string
	PolicyDigest      string
	ToolSet           []ToolIdentity
	ToolCatalogDigest *string
	ProviderLane      ProviderLane
	WorkspaceRootPath string
	WorkspaceRootID   string
	ReleaseBinding    *ReleaseBinding
}

type BuildInput struct {
	RunID             string
	PolicySourceRef   string
	PolicyDigest      string
	PolicyBytes       []byte
	PolicyPinnedLane  LaneID
	ToolSet           []ToolIdentity
	ToolCatalogDigest *string
	ProviderLane      ProviderLane
	WorkspaceRoot     string
	ReleaseBinding    *ReleaseBinding
}

type Frozen struct {
	Manifest Manifest
	Bytes    []byte
	Digest   string
}

// ConnectorAssign carries the frozen digest comparands byte-for-byte. The
// remaining lifecycle fields are populated by the supervisor when it binds a
// connector incarnation to an epoch.
func (frozen Frozen) ConnectorAssign() (appipc.ConnectorAssignBody, error) {
	if frozen.Manifest.RunID == "" || digestBytes(frozen.Bytes) != frozen.Digest || !validDigest(frozen.Digest) || !validDigest(frozen.Manifest.PolicyDigest) || !validDigest(frozen.Manifest.ProviderLane.LaneCatalogDigest) || !validCredentialReference(frozen.Manifest.ProviderLane.CredentialRef) {
		return appipc.ConnectorAssignBody{}, ErrMalformedManifest
	}
	return appipc.ConnectorAssignBody{
		RunID:             frozen.Manifest.RunID,
		RunManifestDigest: frozen.Digest,
		PolicyDigest:      frozen.Manifest.PolicyDigest,
		LaneCatalogDigest: frozen.Manifest.ProviderLane.LaneCatalogDigest,
		CredentialRef:     frozen.Manifest.ProviderLane.CredentialRef,
	}, nil
}

func StagingToolSet() []ToolIdentity {
	tools := make([]ToolIdentity, len(RatifiedToolNames))
	for i, name := range RatifiedToolNames {
		tools[i] = ToolIdentity{Name: name}
	}
	return tools
}

func ToolNames(tools []ToolIdentity) []string {
	names := make([]string, len(tools))
	for i := range tools {
		names[i] = tools[i].Name
	}
	sort.Strings(names)
	return names
}

func Build(input BuildInput) (Frozen, error) {
	if input.RunID == "" || input.PolicySourceRef == "" || !validCredentialReference(input.ProviderLane.CredentialRef) {
		return Frozen{}, ErrMalformedManifest
	}
	if digestBytes(input.PolicyBytes) != input.PolicyDigest {
		return Frozen{}, ErrPolicyDigestMismatch
	}
	if input.PolicyPinnedLane != input.ProviderLane.LaneID {
		return Frozen{}, ErrPolicyLaneMismatch
	}
	if !validDigest(input.PolicyDigest) || !validDigest(input.ProviderLane.LaneCatalogDigest) {
		return Frozen{}, ErrMalformedManifest
	}
	root, err := ResolveWorkspaceRoot(input.WorkspaceRoot)
	if err != nil {
		return Frozen{}, err
	}
	tools := cloneTools(input.ToolSet)
	sort.Slice(tools, func(i, j int) bool { return tools[i].Name < tools[j].Name })
	if !equalStrings(ToolNames(tools), RatifiedToolNames) || !validToolStaging(tools) {
		return Frozen{}, ErrMalformedManifest
	}
	if input.ToolCatalogDigest != nil && !validDigest(*input.ToolCatalogDigest) {
		return Frozen{}, ErrMalformedManifest
	}
	if input.ReleaseBinding != nil {
		if err := validateRelease(*input.ReleaseBinding); err != nil {
			return Frozen{}, err
		}
	}
	manifest := Manifest{
		ManifestSchema: ManifestSchema, RunID: input.RunID, PolicySourceRef: input.PolicySourceRef,
		PolicyDigest: input.PolicyDigest, ToolSet: tools, ToolCatalogDigest: cloneString(input.ToolCatalogDigest),
		ProviderLane: input.ProviderLane, WorkspaceRootPath: root.Path, WorkspaceRootID: root.ID,
		ReleaseBinding: cloneRelease(input.ReleaseBinding),
	}
	encoded, err := appipc.MarshalJCS(canonicalValue(manifest))
	if err != nil {
		return Frozen{}, err
	}
	return Frozen{Manifest: manifest, Bytes: append([]byte(nil), encoded...), Digest: digestBytes(encoded)}, nil
}

func canonicalValue(manifest Manifest) map[string]any {
	tools := make([]any, 0, len(manifest.ToolSet))
	for _, tool := range manifest.ToolSet {
		value := map[string]any{
			"name": tool.Name, "schema_digest": tool.SchemaDigest, "catalog_version": tool.CatalogVersion,
		}
		if strings.HasPrefix(tool.Name, "relay.") {
			value["mapping_version"] = tool.MappingVersion
		}
		tools = append(tools, value)
	}
	value := map[string]any{
		"manifest_schema":     manifest.ManifestSchema,
		"run_id":              manifest.RunID,
		"policy_source_ref":   manifest.PolicySourceRef,
		"policy_digest":       manifest.PolicyDigest,
		"tool_set":            tools,
		"tool_catalog_digest": manifest.ToolCatalogDigest,
		"provider_lane": map[string]any{
			"lane_id": map[string]any{
				"model_id":           manifest.ProviderLane.LaneID.ModelID,
				"provider_id":        manifest.ProviderLane.LaneID.ProviderID,
				"serving_profile_id": manifest.ProviderLane.LaneID.ServingProfileID,
				"compat_mode":        manifest.ProviderLane.LaneID.CompatMode,
			},
			"lane_catalog_digest": manifest.ProviderLane.LaneCatalogDigest,
			"credential_ref":      manifest.ProviderLane.CredentialRef,
		},
		"workspace_root_path": manifest.WorkspaceRootPath,
		"workspace_root_id":   manifest.WorkspaceRootID,
	}
	if manifest.ReleaseBinding != nil {
		release := map[string]any{"bound_at_ref": manifest.ReleaseBinding.BoundAtRef}
		if manifest.ReleaseBinding.BuildDigests != nil {
			release["build_digests"] = *manifest.ReleaseBinding.BuildDigests
		}
		if manifest.ReleaseBinding.ReleaseDigest != nil {
			release["release_digest"] = *manifest.ReleaseBinding.ReleaseDigest
		}
		value["release_binding"] = release
	}
	return value
}

func parse(bytesValue []byte) (Manifest, error) {
	type wireTool struct {
		Name           string  `json:"name"`
		SchemaDigest   *string `json:"schema_digest"`
		CatalogVersion *string `json:"catalog_version"`
		MappingVersion *string `json:"mapping_version,omitempty"`
	}
	type wireRelease struct {
		BoundAtRef    string        `json:"bound_at_ref"`
		BuildDigests  *BuildDigests `json:"build_digests,omitempty"`
		ReleaseDigest *string       `json:"release_digest,omitempty"`
	}
	type wireManifest struct {
		ManifestSchema    int          `json:"manifest_schema"`
		RunID             string       `json:"run_id"`
		PolicySourceRef   string       `json:"policy_source_ref"`
		PolicyDigest      string       `json:"policy_digest"`
		ToolSet           []wireTool   `json:"tool_set"`
		ToolCatalogDigest *string      `json:"tool_catalog_digest"`
		ProviderLane      ProviderLane `json:"provider_lane"`
		WorkspaceRootPath string       `json:"workspace_root_path"`
		WorkspaceRootID   string       `json:"workspace_root_id"`
		ReleaseBinding    *wireRelease `json:"release_binding,omitempty"`
	}
	decoder := json.NewDecoder(bytes.NewReader(bytesValue))
	decoder.DisallowUnknownFields()
	var wire wireManifest
	if err := decoder.Decode(&wire); err != nil {
		return Manifest{}, fmt.Errorf("%w: %v", ErrMalformedManifest, err)
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return Manifest{}, ErrMalformedManifest
	}
	manifest := Manifest{
		ManifestSchema: wire.ManifestSchema, RunID: wire.RunID, PolicySourceRef: wire.PolicySourceRef,
		PolicyDigest: wire.PolicyDigest, ToolCatalogDigest: wire.ToolCatalogDigest, ProviderLane: wire.ProviderLane,
		WorkspaceRootPath: wire.WorkspaceRootPath, WorkspaceRootID: wire.WorkspaceRootID,
	}
	for _, tool := range wire.ToolSet {
		manifest.ToolSet = append(manifest.ToolSet, ToolIdentity{
			Name: tool.Name, SchemaDigest: tool.SchemaDigest, CatalogVersion: tool.CatalogVersion, MappingVersion: tool.MappingVersion,
		})
	}
	if wire.ReleaseBinding != nil {
		manifest.ReleaseBinding = &ReleaseBinding{BoundAtRef: wire.ReleaseBinding.BoundAtRef, BuildDigests: wire.ReleaseBinding.BuildDigests, ReleaseDigest: wire.ReleaseBinding.ReleaseDigest}
	}
	return manifest, nil
}

func validateRelease(release ReleaseBinding) error {
	if release.BoundAtRef == "" || (release.BuildDigests == nil) == (release.ReleaseDigest == nil) {
		return ErrMalformedManifest
	}
	if release.BuildDigests != nil {
		if !validDigest(release.BuildDigests.AppMainBuildDigest) || !validDigest(release.BuildDigests.M9WorkerBuildDigest) || !validDigest(release.BuildDigests.M8BuildDigest) {
			return ErrMalformedManifest
		}
	}
	if release.ReleaseDigest != nil && !validDigest(*release.ReleaseDigest) {
		return ErrMalformedManifest
	}
	return nil
}

func validToolStaging(tools []ToolIdentity) bool {
	for _, tool := range tools {
		for _, digest := range []*string{tool.SchemaDigest} {
			if digest != nil && !validDigest(*digest) {
				return false
			}
		}
		if !strings.HasPrefix(tool.Name, "relay.") && tool.MappingVersion != nil {
			return false
		}
	}
	return true
}

func validDigest(value string) bool {
	if len(value) != 64 {
		return false
	}
	for _, char := range value {
		if (char < '0' || char > '9') && (char < 'a' || char > 'f') {
			return false
		}
	}
	return true
}

// validCredentialReference mirrors the connector-owned opaque-reference
// grammar at the app-control admission boundary. It deliberately validates
// only the reference, never credential bytes.
func validCredentialReference(value string) bool {
	if len(value) == 0 || len(value) > 64 || !lowerAlnum(value[0]) {
		return false
	}
	for index := 1; index < len(value); index++ {
		if !lowerAlnum(value[index]) && value[index] != '-' {
			return false
		}
	}
	return true
}

func lowerAlnum(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= '0' && value <= '9'
}

func digestBytes(value []byte) string {
	digest := sha256.Sum256(value)
	return hex.EncodeToString(digest[:])
}

func cloneTools(input []ToolIdentity) []ToolIdentity {
	tools := make([]ToolIdentity, len(input))
	for i := range input {
		tools[i] = input[i]
		tools[i].SchemaDigest = cloneString(input[i].SchemaDigest)
		tools[i].CatalogVersion = cloneString(input[i].CatalogVersion)
		tools[i].MappingVersion = cloneString(input[i].MappingVersion)
	}
	return tools
}

func cloneRelease(input *ReleaseBinding) *ReleaseBinding {
	if input == nil {
		return nil
	}
	result := *input
	result.ReleaseDigest = cloneString(input.ReleaseDigest)
	if input.BuildDigests != nil {
		copy := *input.BuildDigests
		result.BuildDigests = &copy
	}
	return &result
}

func cloneString(value *string) *string {
	if value == nil {
		return nil
	}
	copy := *value
	return &copy
}

func equalStrings(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
