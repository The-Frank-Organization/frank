package f59

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/jackli/frank/internal/appctl/manifest"
	"github.com/jackli/frank/internal/appipc"
	"golang.org/x/text/unicode/norm"
)

type EffectDescriptor struct {
	Action              string  `json:"action"`
	CanonicalResource   *string `json:"canonical_resource,omitempty"`
	WorkspaceRootID     *string `json:"workspace_root_id,omitempty"`
	CWD                 *string `json:"cwd,omitempty"`
	BackendID           string  `json:"backend_id"`
	NetworkPolicyID     string  `json:"network_policy_id"`
	ToolImplRef         string  `json:"tool_impl_ref"`
	CanonicalArgsDigest string  `json:"canonical_args_digest"`
	OneShot             bool    `json:"one_shot"`
}

func OrderedTargetSetDigest(paths []string) (string, error) {
	if len(paths) == 0 {
		return "", ErrInvalidDescriptor
	}
	canonical := append([]string(nil), paths...)
	for _, path := range canonical {
		if !validRelativeForm(path) {
			return "", ErrInvalidDescriptor
		}
	}
	sort.Strings(canonical)
	for index := 1; index < len(canonical); index++ {
		if canonical[index] == canonical[index-1] {
			return "", ErrInvalidDescriptor
		}
	}
	encoded, err := appipc.MarshalJCS(canonical)
	if err != nil {
		return "", err
	}
	return digestBytes(encoded), nil
}

func BuildDescriptor(run manifest.Manifest, action, argsDigest string, operands Operands) (EffectDescriptor, []byte, error) {
	if !validDigest(argsDigest) {
		return EffectDescriptor{}, nil, ErrInvalidDescriptor
	}
	tool, ok := findTool(run.ToolSet, action)
	if !ok || tool.SchemaDigest == nil || tool.CatalogVersion == nil {
		return EffectDescriptor{}, nil, ErrInvalidDescriptor
	}
	descriptor := EffectDescriptor{Action: action, CanonicalArgsDigest: argsDigest, NetworkPolicyID: "none", OneShot: true}
	descriptor.ToolImplRef = toolReference(tool)
	local := action == "read" || action == "write" || action == "edit" || action == "apply_patch" || action == "bash"
	if local {
		descriptor.BackendID = "ambient"
		root := run.WorkspaceRootID
		descriptor.WorkspaceRootID = &root
		if operands.CWD == nil || !validRelativeForm(*operands.CWD) {
			return EffectDescriptor{}, nil, ErrInvalidDescriptor
		}
		descriptor.CWD = cloneString(operands.CWD)
	} else if strings.HasPrefix(action, "relay.") {
		descriptor.BackendID = "conductor-client"
		if operands.CWD != nil {
			return EffectDescriptor{}, nil, ErrInvalidDescriptor
		}
	} else {
		return EffectDescriptor{}, nil, ErrInvalidDescriptor
	}

	switch action {
	case "read", "write", "edit":
		if operands.CanonicalResource == nil || !validRelativeForm(*operands.CanonicalResource) {
			return EffectDescriptor{}, nil, ErrInvalidDescriptor
		}
		descriptor.CanonicalResource = cloneString(operands.CanonicalResource)
	case "apply_patch":
		if operands.CanonicalResource == nil || !validDigest(*operands.CanonicalResource) {
			return EffectDescriptor{}, nil, ErrInvalidDescriptor
		}
		descriptor.CanonicalResource = cloneString(operands.CanonicalResource)
	case "bash":
		if operands.CanonicalResource != nil {
			return EffectDescriptor{}, nil, ErrInvalidDescriptor
		}
	default:
		if operands.CanonicalResource == nil || !validRelayResource(action, *operands.CanonicalResource) {
			return EffectDescriptor{}, nil, ErrInvalidDescriptor
		}
		descriptor.CanonicalResource = cloneString(operands.CanonicalResource)
	}
	encoded, err := appipc.MarshalJCS(descriptor)
	if err != nil {
		return EffectDescriptor{}, nil, err
	}
	return descriptor, encoded, nil
}

func validRelativeForm(value string) bool {
	if value == "" || len([]byte(value)) > appipc.PathMaxM10 || !utf8.ValidString(value) || !norm.NFC.IsNormalString(value) || strings.HasPrefix(value, "/") || strings.HasSuffix(value, "/") || strings.Contains(value, "//") {
		return false
	}
	for _, segment := range strings.Split(value, "/") {
		if segment == "" || segment == ".." {
			return false
		}
	}
	for _, char := range value {
		if !(char >= 'A' && char <= 'Z') && !(char >= 'a' && char <= 'z') && !(char >= '0' && char <= '9') && !strings.ContainsRune("._:/-", char) {
			return false
		}
	}
	return true
}

func validRelayResource(action, value string) bool {
	if !strings.HasPrefix(value, action+":") || len(value) == len(action)+1 || len([]byte(value)) > appipc.PathMaxM10 || !norm.NFC.IsNormalString(value) {
		return false
	}
	for _, char := range value {
		if !(char >= 'A' && char <= 'Z') && !(char >= 'a' && char <= 'z') && !(char >= '0' && char <= '9') && !strings.ContainsRune("._:/-", char) {
			return false
		}
	}
	return true
}

func findTool(tools []manifest.ToolIdentity, name string) (manifest.ToolIdentity, bool) {
	for _, tool := range tools {
		if tool.Name == name {
			return tool, true
		}
	}
	return manifest.ToolIdentity{}, false
}

func toolReference(tool manifest.ToolIdentity) string {
	value := map[string]any{"name": tool.Name, "schema_digest": tool.SchemaDigest, "catalog_version": tool.CatalogVersion}
	if tool.MappingVersion != nil {
		value["mapping_version"] = tool.MappingVersion
	}
	encoded, err := appipc.MarshalJCS(value)
	if err != nil {
		panic(err)
	}
	return "manifest-tool:" + digestBytes(encoded)
}

func validDigest(value string) bool {
	if len(value) != 64 {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil && strings.ToLower(value) == value
}

func digestBytes(value []byte) string {
	hash := sha256.Sum256(value)
	return hex.EncodeToString(hash[:])
}

func cloneString(value *string) *string {
	if value == nil {
		return nil
	}
	copy := *value
	return &copy
}
