// Package catalog loads and validates the factual provider lane catalog.
package catalog

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/The-Frank-Organization/frank/internal/connector/endpoint"
	"github.com/The-Frank-Organization/frank/internal/connector/jcs"
	"github.com/The-Frank-Organization/frank/internal/connector/nfccheck"
)

var ErrInvalidCatalog = errors.New("connector: invalid lane catalog")

type Catalog struct {
	Lanes  []Lane `json:"lanes"`
	Schema string `json:"schema"`
	Digest string `json:"-"`
}

type Lane struct {
	Auth             Auth         `json:"auth"`
	CompatMode       string       `json:"compat_mode"`
	Cost             Cost         `json:"cost"`
	Endpoint         string       `json:"endpoint"`
	LaneID           string       `json:"lane_id"`
	Limits           Limits       `json:"limits"`
	Method           string       `json:"method"`
	ModelID          string       `json:"model_id"`
	ObservedAt       string       `json:"observed_at"`
	ProfileFacts     ProfileFacts `json:"profile_facts"`
	ProviderID       string       `json:"provider_id"`
	Reasoning        Reasoning    `json:"reasoning"`
	ServingProfileID string       `json:"serving_profile_id"`
	Source           string       `json:"source"`
	ToolUse          ToolUse      `json:"tool_use"`
	Wire             Wire         `json:"wire"`
}

type Auth struct {
	HeaderName string `json:"auth_header_name"`
	Scheme     string `json:"auth_scheme"`
}

type Wire struct {
	MaxOutputTokensField string `json:"max_output_tokens_field"`
	ServerRetention      bool   `json:"server_retention"`
	Streaming            bool   `json:"streaming"`
	UsageInStreaming     bool   `json:"usage_in_streaming"`
}

type Reasoning struct {
	EffortLevels []string `json:"effort_levels"`
	ReplayKind   string   `json:"replay_kind"`
	Supported    bool     `json:"supported"`
}

type ToolUse struct {
	StrictSchema bool `json:"strict_schema"`
	Supported    bool `json:"supported"`
}

type Limits struct {
	Context   int64 `json:"context"`
	MaxOutput int64 `json:"max_output"`
}

type Cost struct {
	CacheRead     *json.Number `json:"cache_read,omitempty"`
	EffectiveTime string       `json:"effective_time"`
	Input         json.Number  `json:"input"`
	Output        json.Number  `json:"output"`
}

type ProfileFacts struct {
	EndpointKind string  `json:"endpoint_kind"`
	Precision    *string `json:"precision,omitempty"`
	Region       string  `json:"region"`
}

// Load requires exact canonical stored bytes and a closed JSON schema before
// applying the catalog's semantic composition rules.
func Load(raw []byte) (*Catalog, error) {
	if len(raw) == 0 || !jcs.IsCanonical(raw) {
		return nil, invalid("stored bytes are not exact JCS")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	decoder.UseNumber()
	var catalog Catalog
	if err := decoder.Decode(&catalog); err != nil {
		return nil, invalid("decode closed schema: %v", err)
	}
	if err := requireEOF(decoder); err != nil {
		return nil, err
	}
	if err := catalog.Validate(); err != nil {
		return nil, err
	}
	digest := sha256.Sum256(raw)
	catalog.Digest = hex.EncodeToString(digest[:])
	return &catalog, nil
}

func (c *Catalog) Validate() error {
	if c == nil || c.Schema != "m8.lane_catalog.v1" || len(c.Lanes) == 0 {
		return invalid("schema or lanes")
	}
	seenAxes := make(map[string]struct{}, len(c.Lanes))
	previousID := ""
	for index := range c.Lanes {
		lane := &c.Lanes[index]
		if !validID(lane.LaneID) || !validID(lane.ServingProfileID) {
			return invalid("lane %d has invalid opaque id", index)
		}
		if index > 0 && previousID >= lane.LaneID {
			return invalid("lanes are not sorted and unique")
		}
		previousID = lane.LaneID
		axes := strings.Join([]string{lane.ModelID, lane.ProviderID, lane.ServingProfileID, lane.CompatMode}, "\x00")
		if _, duplicate := seenAxes[axes]; duplicate {
			return invalid("duplicate four-axis lane key")
		}
		seenAxes[axes] = struct{}{}
		if err := validateLane(lane); err != nil {
			return invalid("lane %q: %v", lane.LaneID, err)
		}
	}
	return nil
}

// ValidateDeniedHeaders performs the catalog side of the P0 membership
// invariant after the policy artifact has been loaded.
func (c *Catalog) ValidateDeniedHeaders(deniedHeaderNames []string) error {
	denied := make(map[string]struct{}, len(deniedHeaderNames))
	for _, name := range deniedHeaderNames {
		denied[name] = struct{}{}
	}
	for _, lane := range c.Lanes {
		if _, ok := denied[lane.Auth.HeaderName]; !ok {
			return invalid("lane %q auth header is not denied by policy", lane.LaneID)
		}
	}
	return nil
}

func validateLane(lane *Lane) error {
	stringsToCheck := []string{
		lane.LaneID, lane.ModelID, lane.ProviderID, lane.ServingProfileID,
		lane.CompatMode, lane.Endpoint, lane.Method, lane.Auth.HeaderName,
		lane.Auth.Scheme, lane.Wire.MaxOutputTokensField, lane.Reasoning.ReplayKind,
		lane.Cost.EffectiveTime, lane.ProfileFacts.Region, lane.ProfileFacts.EndpointKind,
		lane.Source, lane.ObservedAt,
	}
	stringsToCheck = append(stringsToCheck, lane.Reasoning.EffortLevels...)
	if lane.ProfileFacts.Precision != nil {
		stringsToCheck = append(stringsToCheck, *lane.ProfileFacts.Precision)
	}
	for _, value := range stringsToCheck {
		if value == "" || !nfccheck.IsNormalString(value) {
			return fmt.Errorf("empty or non-NFC string")
		}
	}
	if err := endpoint.Validate(lane.Endpoint); err != nil {
		return err
	}
	if lane.Method != "POST" || lane.CompatMode != "openai-responses" || lane.Auth.Scheme != "bearer" {
		return fmt.Errorf("unsupported method, compatibility mode, or auth scheme")
	}
	if !validLowerHeader(lane.Auth.HeaderName) {
		return fmt.Errorf("invalid lowercase auth header")
	}
	if !lane.Wire.Streaming || !lane.Wire.UsageInStreaming || lane.Wire.MaxOutputTokensField != "max_output_tokens" || lane.Wire.ServerRetention {
		return fmt.Errorf("unsupported wire facts")
	}
	if !lane.ToolUse.Supported {
		return fmt.Errorf("tool use must be supported for the MVP lane")
	}
	if lane.Reasoning.ReplayKind != "opaque_item" && lane.Reasoning.ReplayKind != "none" {
		return fmt.Errorf("invalid replay kind")
	}
	if lane.Limits.Context <= 0 || lane.Limits.MaxOutput <= 0 || lane.Limits.MaxOutput > lane.Limits.Context {
		return fmt.Errorf("invalid limits")
	}
	if err := validateCost(lane.Cost); err != nil {
		return err
	}
	if lane.ProfileFacts.EndpointKind != "general" && lane.ProfileFacts.EndpointKind != "coding" {
		return fmt.Errorf("invalid endpoint kind")
	}
	if lane.Source != "seeded" {
		return fmt.Errorf("invalid source")
	}
	if _, err := time.Parse(time.RFC3339, lane.ObservedAt); err != nil {
		return fmt.Errorf("invalid observed_at")
	}
	return nil
}

func validateCost(cost Cost) error {
	if _, err := time.Parse(time.RFC3339, cost.EffectiveTime); err != nil {
		return fmt.Errorf("invalid effective_time")
	}
	values := []*json.Number{&cost.Input, &cost.Output}
	if cost.CacheRead != nil {
		values = append(values, cost.CacheRead)
	}
	for _, value := range values {
		parsed, err := strconv.ParseFloat(value.String(), 64)
		if err != nil || parsed < 0 {
			return fmt.Errorf("invalid cost")
		}
	}
	return nil
}

func validID(value string) bool {
	if len(value) == 0 || len(value) > 64 || !isLowerAlnum(value[0]) {
		return false
	}
	for i := 1; i < len(value); i++ {
		if !isLowerAlnum(value[i]) && value[i] != '-' {
			return false
		}
	}
	return true
}

func validLowerHeader(value string) bool {
	if value == "" || value != strings.ToLower(value) {
		return false
	}
	for i := range value {
		if !isLowerAlnum(value[i]) && value[i] != '-' {
			return false
		}
	}
	return true
}

func isLowerAlnum(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= '0' && value <= '9'
}

func requireEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return invalid("trailing JSON value")
		}
		return invalid("trailing bytes: %v", err)
	}
	return nil
}

func invalid(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrInvalidCatalog, fmt.Sprintf(format, args...))
}
