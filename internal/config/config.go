package config

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jackli/frank/internal/fieldspec"
)

var ErrMissingEngine = errors.New("missing engine config member")
var ErrConfigVersionTransition = errors.New("config-version-transition")
var ErrConfigLoad = errors.New("config-load")

type Pinned struct {
	Members  map[string][]byte
	Engine   EngineConfig
	Registry *fieldspec.Registry
	Supply   *Supply
	Digest   string
}

type EngineConfig struct {
	Version            int             `json:"version,omitempty"`
	GCEnabled          bool            `json:"gc_enabled"`
	SegmentRotateBytes int64           `json:"segment_rotate_bytes"`
	MaxFrameBytes      int             `json:"max_frame_bytes,omitempty"`
	PresentLayers      map[string]bool `json:"present_layers,omitempty"`
	Detector           json.RawMessage `json:"detector,omitempty"`
	Supply             *Supply         `json:"supply,omitempty"`
}

type Supply struct {
	LaneRoots  map[string]string `json:"lane_roots"`
	SchemaRefs map[string]string `json:"schema_refs"`
	Suites     map[string]Suite  `json:"suites"`
}

type Suite struct {
	Lane           string        `json:"lane"`
	Command        string        `json:"command"`
	Args           []string      `json:"args"`
	TimeoutClass   string        `json:"timeout_class"`
	TimeoutSeconds int           `json:"timeout_seconds"`
	SourceDir      string        `json:"-"`
	Timeout        time.Duration `json:"-"`
}

type catalogMember struct {
	Version          string                  `json:"version"`
	Status           string                  `json:"status"`
	ChangeConvention catalogChangeConvention `json:"change_convention"`
	Discovery        catalogDiscovery        `json:"discovery"`
	Laws             []catalogLaw            `json:"laws"`
}

type catalogChangeConvention struct {
	SingleWriter  string `json:"single_writer"`
	OwnerFidelity string `json:"owner_fidelity"`
	Section7Claim string `json:"section_7_claim"`
	Section7Carry string `json:"section_7_carry"`
}

type catalogDiscovery struct {
	Scan                  catalogScan                  `json:"scan"`
	BoundaryFiles         []string                     `json:"boundary_files"`
	Recognizers           catalogRecognizers           `json:"recognizers"`
	SiteCensus            []string                     `json:"site_census"`
	OutputFamilies        []catalogOutputFamily        `json:"output_families"`
	SinkPatterns          []catalogSinkPattern         `json:"sink_patterns"`
	CanonicalPathFamilies catalogCanonicalPathFamilies `json:"canonical_path_families"`
}

type catalogScan struct {
	Root                string   `json:"root"`
	ExcludeDirPrefixes  []string `json:"exclude_dir_prefixes"`
	ExcludeFileSuffixes []string `json:"exclude_file_suffixes"`
	IncludeOnlySuffixes []string `json:"include_only_suffixes"`
}

type catalogChannelRecognizer struct {
	PathPrefix        string   `json:"path_prefix"`
	CallSelectors     []string `json:"call_selectors"`
	IdentCalls        []string `json:"ident_calls"`
	ConnReceiverCalls []string `json:"conn_receiver_calls"`
}

type catalogMCPRecognizer struct {
	PathPrefix    string   `json:"path_prefix"`
	CallSelectors []string `json:"call_selectors"`
}

type catalogRecognizers struct {
	ChannelContext            catalogChannelRecognizer `json:"channel_context"`
	MCPContext                catalogMCPRecognizer     `json:"mcp_context"`
	ProtocolSwitchTagSelector string                   `json:"protocol_switch_tag_selector"`
	TreeWideIdioms            []string                 `json:"tree_wide_idioms"`
}

type catalogOutputFamily struct {
	ID        string   `json:"id"`
	CarveOuts []string `json:"carve_outs"`
}

type catalogSinkPattern struct {
	ID            string `json:"id"`
	Family        string `json:"family"`
	Symbol        string `json:"symbol"`
	ExpectedSites int    `json:"expected_sites"`
}

type catalogCanonicalPathFamilies struct {
	Rows []catalogCanonicalPath `json:"rows"`
}

type catalogCanonicalPath struct {
	ID        string `json:"id"`
	Relative  string `json:"relative"`
	Forbidden string `json:"forbidden"`
	Directory *bool  `json:"directory"`
}

type catalogLaw struct {
	ID             string   `json:"id"`
	Test           string   `json:"test"`
	Owners         []string `json:"owners"`
	FidelityReview []string `json:"fidelity_review"`
	Claim          string   `json:"claim"`
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
	if err := preflightMemberMarkers(loaded); err != nil {
		return nil, err
	}
	if catalog, ok := loaded["catalog"]; ok {
		if err := validateCatalogSchema(catalog); err != nil {
			return nil, fmt.Errorf("%w: catalog", ErrConfigLoad)
		}
	}
	engineVersion, err := validateEngineSchema(loaded["engine"])
	if err != nil {
		return nil, fmt.Errorf("%w: engine", ErrConfigLoad)
	}

	var engine EngineConfig
	if err := json.Unmarshal(loaded["engine"], &engine); err != nil {
		return nil, fmt.Errorf("parse engine config: %w", err)
	}
	var supply *Supply
	if engineVersion == 2 {
		var err error
		supply, err = composeSupply(engine.Supply)
		if err != nil {
			return nil, fmt.Errorf("%w: supply", ErrConfigLoad)
		}
		engine.Supply = cloneSupply(supply)
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
		Supply:   supply,
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
	if err := validateCandidateMarker(member, candidate); err != nil {
		return ErrConfigVersionTransition
	}
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
		if candidateVersion == currentVersion || candidateVersion == currentVersion+1 && candidateVersion <= 2 {
			return nil
		}
		return ErrConfigVersionTransition
	case "fieldspec":
		if _, err := fieldspec.Parse(candidate); err != nil {
			return ErrConfigVersionTransition
		}
		return validateFieldspecMarkerTransition(current, candidate)
	case "catalog":
		if err := validateCatalogSchema(candidate); err != nil {
			return ErrConfigVersionTransition
		}
		return validateStringMarkerTransition(current, candidate, "s7-v1", "s8-v1")
	default:
		return ErrConfigVersionTransition
	}
}

func validateCandidateMarker(member string, candidate []byte) error {
	switch member {
	case "engine":
		_, present, err := rawIntMarker(candidate)
		if err != nil || !present {
			return ErrConfigVersionTransition
		}
	case "fieldspec", "catalog":
		if _, err := rawStringMarker(candidate); err != nil {
			return ErrConfigVersionTransition
		}
	default:
		return ErrConfigVersionTransition
	}
	return nil
}

func preflightMemberMarkers(loaded map[string][]byte) error {
	if data, ok := loaded["catalog"]; ok {
		marker, err := rawStringMarker(data)
		if err != nil || marker != "s8-v1" {
			return fmt.Errorf("%w: catalog-marker", ErrConfigLoad)
		}
	}
	if data, ok := loaded["fieldspec"]; ok {
		if err := ValidateFieldspecReaderMarker(data, "s7a-fieldspec-v5", "s8-fieldspec-v6", "s8-fieldspec-v7"); err != nil {
			return err
		}
	}
	if data, ok := loaded["engine"]; ok {
		_, present, err := rawIntMarker(data)
		if err != nil || present && ValidateEngineReaderMarker(data, 2) != nil || !present && loaded["catalog"] != nil {
			return fmt.Errorf("%w: engine-marker", ErrConfigLoad)
		}
	}
	return nil
}

func ValidateEngineReaderMarker(data []byte, maxVersion int) error {
	version, present, err := rawIntMarker(data)
	if err != nil || !present || version < 1 || version > maxVersion {
		return fmt.Errorf("%w: engine-marker", ErrConfigLoad)
	}
	return nil
}

func ValidateFieldspecReaderMarker(data []byte, supported ...string) error {
	marker, err := rawStringMarker(data)
	if err != nil || !stringMember(marker, supported...) {
		return fmt.Errorf("%w: fieldspec-marker", ErrConfigLoad)
	}
	return nil
}

func rawStringMarker(data []byte) (string, error) {
	raw, present, err := rawObjectField(data, "version")
	if err != nil || !present {
		return "", ErrConfigLoad
	}
	var value string
	if err := json.Unmarshal(raw, &value); err != nil || value == "" {
		return "", ErrConfigLoad
	}
	return value, nil
}

func rawIntMarker(data []byte) (int, bool, error) {
	raw, present, err := rawObjectField(data, "version")
	if err != nil || !present {
		return 0, present, err
	}
	var value int
	if err := json.Unmarshal(raw, &value); err != nil || value < 0 {
		return 0, true, ErrConfigLoad
	}
	return value, true, nil
}

func rawObjectField(data []byte, field string) (json.RawMessage, bool, error) {
	decoder := json.NewDecoder(bytes.NewReader(data))
	token, err := decoder.Token()
	if err != nil || token != json.Delim('{') {
		return nil, false, ErrConfigLoad
	}
	var found json.RawMessage
	for decoder.More() {
		keyToken, err := decoder.Token()
		if err != nil {
			return nil, false, ErrConfigLoad
		}
		key, ok := keyToken.(string)
		if !ok {
			return nil, false, ErrConfigLoad
		}
		var raw json.RawMessage
		if err := decoder.Decode(&raw); err != nil {
			return nil, false, ErrConfigLoad
		}
		if key == field {
			if found != nil {
				return nil, false, ErrConfigLoad
			}
			found = append(json.RawMessage(nil), raw...)
		}
	}
	closing, err := decoder.Token()
	if err != nil || closing != json.Delim('}') {
		return nil, false, ErrConfigLoad
	}
	if _, err := decoder.Token(); err != io.EOF {
		return nil, false, ErrConfigLoad
	}
	return found, found != nil, nil
}

func stringMember(value string, values ...string) bool {
	for _, candidate := range values {
		if value == candidate {
			return true
		}
	}
	return false
}

func validateFieldspecMarkerTransition(current, candidate []byte) error {
	from, err := stringMarker(current)
	if err != nil {
		return ErrConfigVersionTransition
	}
	to, err := stringMarker(candidate)
	if err != nil {
		return ErrConfigVersionTransition
	}
	if from == to || from == "s7a-fieldspec-v5" && to == "s8-fieldspec-v6" || from == "s8-fieldspec-v6" && to == "s8-fieldspec-v7" {
		return nil
	}
	return ErrConfigVersionTransition
}

func validateCatalogSchema(data []byte) error {
	var marker struct {
		Version string `json:"version"`
	}
	if err := json.Unmarshal(data, &marker); err != nil || marker.Version != "s8-v1" {
		return ErrConfigLoad
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var member catalogMember
	if err := decoder.Decode(&member); err != nil {
		return ErrConfigLoad
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return ErrConfigLoad
	}
	if member.Status != "section7-pinned" || member.ChangeConvention.SingleWriter == "" || member.ChangeConvention.OwnerFidelity == "" || member.ChangeConvention.Section7Claim == "" || member.ChangeConvention.Section7Carry == "" {
		return ErrConfigLoad
	}
	if member.Discovery.Scan.Root == "" || len(member.Discovery.Scan.ExcludeDirPrefixes) == 0 || len(member.Discovery.Scan.ExcludeFileSuffixes) == 0 || len(member.Discovery.Scan.IncludeOnlySuffixes) == 0 || len(member.Discovery.BoundaryFiles) == 0 || len(member.Discovery.SiteCensus) == 0 {
		return ErrConfigLoad
	}
	if member.Discovery.Recognizers.ChannelContext.PathPrefix == "" || len(member.Discovery.Recognizers.ChannelContext.CallSelectors) == 0 || len(member.Discovery.Recognizers.ChannelContext.IdentCalls) == 0 || len(member.Discovery.Recognizers.ChannelContext.ConnReceiverCalls) == 0 || member.Discovery.Recognizers.MCPContext.PathPrefix == "" || len(member.Discovery.Recognizers.MCPContext.CallSelectors) == 0 || member.Discovery.Recognizers.ProtocolSwitchTagSelector == "" || len(member.Discovery.Recognizers.TreeWideIdioms) == 0 {
		return ErrConfigLoad
	}
	families := map[string]bool{}
	for _, family := range member.Discovery.OutputFamilies {
		if family.ID == "" || family.CarveOuts == nil || families[family.ID] {
			return ErrConfigLoad
		}
		families[family.ID] = true
	}
	patterns := map[string]bool{}
	for _, pattern := range member.Discovery.SinkPatterns {
		if pattern.ID == "" || pattern.Family == "" || pattern.Symbol == "" || pattern.ExpectedSites < 0 || patterns[pattern.ID] || !families[pattern.Family] {
			return ErrConfigLoad
		}
		patterns[pattern.ID] = true
	}
	paths := map[string]bool{}
	for _, row := range member.Discovery.CanonicalPathFamilies.Rows {
		if row.ID == "" || row.Relative == "" || row.Forbidden == "" || row.Directory == nil || paths[row.ID] {
			return ErrConfigLoad
		}
		paths[row.ID] = true
	}
	lawIDs, tests := map[string]bool{}, map[string]bool{}
	for _, law := range member.Laws {
		if law.ID == "" || law.Test == "" || law.Claim == "" || len(law.Owners) == 0 || law.FidelityReview == nil || lawIDs[law.ID] || tests[law.Test] {
			return ErrConfigLoad
		}
		lawIDs[law.ID], tests[law.Test] = true, true
	}
	if len(families) == 0 || len(patterns) == 0 || len(paths) == 0 || len(lawIDs) == 0 {
		return ErrConfigLoad
	}
	return nil
}

const suiteBoundedCeiling = 120 * time.Second

func composeSupply(raw *Supply) (*Supply, error) {
	if raw == nil || len(raw.LaneRoots) == 0 || len(raw.Suites) == 0 {
		return nil, ErrConfigLoad
	}
	out := &Supply{LaneRoots: map[string]string{}, SchemaRefs: map[string]string{}, Suites: map[string]Suite{}}
	for id, root := range raw.LaneRoots {
		if !validSupplyID(id) || !filepath.IsAbs(root) || filepath.Clean(root) != root {
			return nil, ErrConfigLoad
		}
		resolved, err := filepath.EvalSymlinks(root)
		if err != nil || resolved != root {
			return nil, ErrConfigLoad
		}
		info, err := os.Stat(resolved)
		if err != nil || !info.IsDir() {
			return nil, ErrConfigLoad
		}
		out.LaneRoots[id] = resolved
	}
	for id, digest := range raw.SchemaRefs {
		if !validSupplyID(id) || len(digest) != 64 || strings.ToLower(digest) != digest {
			return nil, ErrConfigLoad
		}
		if _, err := hex.DecodeString(digest); err != nil {
			return nil, ErrConfigLoad
		}
		out.SchemaRefs[id] = digest
	}
	for target, suite := range raw.Suites {
		root := out.LaneRoots[suite.Lane]
		if !validSupplyID(target) || root == "" || suite.TimeoutClass != "suite_bounded" || suite.TimeoutSeconds < 1 || time.Duration(suite.TimeoutSeconds)*time.Second > suiteBoundedCeiling {
			return nil, ErrConfigLoad
		}
		if !validSupplyCommand(root, suite.Command) || !validSupplyArgs(suite.Args) {
			return nil, ErrConfigLoad
		}
		suite.Args = append([]string(nil), suite.Args...)
		suite.SourceDir = root
		suite.Timeout = time.Duration(suite.TimeoutSeconds) * time.Second
		out.Suites[target] = suite
	}
	return out, nil
}

func validSupplyID(value string) bool {
	if value == "" || len(value) > 128 {
		return false
	}
	for _, char := range value {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || strings.ContainsRune("._:-", char) {
			continue
		}
		return false
	}
	return true
}

func validSupplyCommand(root, command string) bool {
	if command == "" || filepath.IsAbs(command) || filepath.Clean(command) != command || command == "." || command == ".." || strings.HasPrefix(command, ".."+string(filepath.Separator)) {
		return false
	}
	joined := filepath.Join(root, command)
	resolved, err := filepath.EvalSymlinks(joined)
	if err != nil || resolved != joined {
		return false
	}
	rel, err := filepath.Rel(root, resolved)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return false
	}
	info, err := os.Stat(resolved)
	return err == nil && info.Mode().IsRegular() && info.Mode().Perm()&0o111 != 0
}

func validSupplyArgs(args []string) bool {
	return args != nil && len(args) == 0
}

func cloneSupply(in *Supply) *Supply {
	if in == nil {
		return nil
	}
	out := &Supply{LaneRoots: map[string]string{}, SchemaRefs: map[string]string{}, Suites: map[string]Suite{}}
	for id, root := range in.LaneRoots {
		out.LaneRoots[id] = root
	}
	for id, digest := range in.SchemaRefs {
		out.SchemaRefs[id] = digest
	}
	for target, suite := range in.Suites {
		suite.Args = append([]string(nil), suite.Args...)
		out.Suites[target] = suite
	}
	return out
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
	if err != nil || version < 0 || version > 2 {
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
	} else if version == 2 {
		allowed["version"] = "number"
		allowed["present_layers"] = "layers"
		allowed["supply"] = "object"
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
	if version >= 1 {
		if _, ok := raw["version"]; !ok {
			return 0, ErrConfigVersionTransition
		}
		if _, ok := raw["present_layers"]; !ok {
			return 0, ErrConfigVersionTransition
		}
	}
	if version == 2 {
		supply, ok := raw["supply"].(map[string]any)
		if !ok || !validSupplyShape(supply) {
			return 0, ErrConfigVersionTransition
		}
	}
	return version, nil
}

func validSupplyShape(raw map[string]any) bool {
	if !exactKeys(raw, "lane_roots", "schema_refs", "suites") {
		return false
	}
	lanes, ok := raw["lane_roots"].(map[string]any)
	if !ok || len(lanes) == 0 || !stringMapShape(lanes) {
		return false
	}
	refs, ok := raw["schema_refs"].(map[string]any)
	if !ok || !stringMapShape(refs) {
		return false
	}
	suites, ok := raw["suites"].(map[string]any)
	if !ok || len(suites) == 0 {
		return false
	}
	for _, value := range suites {
		suite, ok := value.(map[string]any)
		if !ok || !exactKeys(suite, "lane", "command", "args", "timeout_class", "timeout_seconds") {
			return false
		}
		if _, ok := suite["lane"].(string); !ok {
			return false
		}
		if _, ok := suite["command"].(string); !ok {
			return false
		}
		args, ok := suite["args"].([]any)
		if !ok {
			return false
		}
		for _, arg := range args {
			if _, ok := arg.(string); !ok {
				return false
			}
		}
		if _, ok := suite["timeout_class"].(string); !ok {
			return false
		}
		seconds, ok := suite["timeout_seconds"].(float64)
		if !ok || seconds != float64(int(seconds)) {
			return false
		}
	}
	return true
}

func exactKeys(raw map[string]any, keys ...string) bool {
	if len(raw) != len(keys) {
		return false
	}
	for _, key := range keys {
		if _, ok := raw[key]; !ok {
			return false
		}
	}
	return true
}

func stringMapShape(raw map[string]any) bool {
	for _, value := range raw {
		if _, ok := value.(string); !ok {
			return false
		}
	}
	return true
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
