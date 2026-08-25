// Package policy loads the frozen m-3 provider-request policy artifact.
package policy

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/endpoint"
	"github.com/jackli/frank/internal/connector/jcs"
	"github.com/jackli/frank/internal/connector/nfccheck"
)

var (
	ErrPolicyUnavailable    = errors.New("policy-unavailable")
	ErrPolicyDigestMismatch = errors.New("policy-digest-mismatch")
)

type Policy struct {
	DeniedHeaderNames []string `json:"denied_header_names"`
	EgressClass       string   `json:"egress_class"`
	EndpointAllowlist []string `json:"endpoint_allowlist"`
	PinnedLane        string   `json:"pinned_lane"`
	Schema            string   `json:"schema"`
	Digest            string   `json:"-"`
}

// Load verifies exact stored-byte canonicality, the closed schema, semantic
// set rules, and the catalog-side lane facts needed by P0.
func Load(raw []byte, lane catalog.Lane) (*Policy, error) {
	if len(raw) == 0 || !jcs.IsCanonical(raw) {
		return nil, unavailable("stored bytes are absent or non-canonical")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var policy Policy
	if err := decoder.Decode(&policy); err != nil {
		return nil, unavailable("decode closed schema: %v", err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, unavailable("trailing JSON value")
	}
	if err := policy.validate(lane); err != nil {
		return nil, err
	}
	digest := sha256.Sum256(raw)
	policy.Digest = hex.EncodeToString(digest[:])
	return &policy, nil
}

func (p *Policy) VerifyDigest(expected string) error {
	if p == nil || !validDigest(expected) || p.Digest != expected {
		return ErrPolicyDigestMismatch
	}
	return nil
}

// ValidateForLane rechecks the loaded policy's P0 semantic invariants against
// the immutable pinned-lane facts presented to authorization.
func (p *Policy) ValidateForLane(lane catalog.Lane) error {
	if p == nil {
		return unavailable("policy is absent")
	}
	return p.validate(lane)
}

func (p *Policy) AllowsEndpoint(value string) bool {
	if p == nil {
		return false
	}
	for _, allowed := range p.EndpointAllowlist {
		if value == allowed {
			return true
		}
	}
	return false
}

func (p *Policy) DeniesHeader(name string) bool {
	if p == nil {
		return false
	}
	for _, denied := range p.DeniedHeaderNames {
		if name == denied {
			return true
		}
	}
	return false
}

func (p *Policy) validate(lane catalog.Lane) error {
	if p.Schema != "m3.egress_policy.v1" || p.EgressClass != "provider-request" || p.PinnedLane != lane.LaneID {
		return unavailable("schema, egress class, or pinned lane mismatch")
	}
	allStrings := []string{p.Schema, p.EgressClass, p.PinnedLane}
	allStrings = append(allStrings, p.EndpointAllowlist...)
	allStrings = append(allStrings, p.DeniedHeaderNames...)
	for _, value := range allStrings {
		if value == "" || !nfccheck.IsNormalString(value) {
			return unavailable("empty or non-NFC string")
		}
	}
	if len(p.EndpointAllowlist) == 0 || !strictlySorted(p.EndpointAllowlist) || !strictlySorted(p.DeniedHeaderNames) {
		return unavailable("set arrays must be sorted and duplicate-free")
	}
	for _, value := range p.EndpointAllowlist {
		if err := endpoint.Validate(value); err != nil {
			return unavailable("invalid allowlist endpoint: %v", err)
		}
	}
	for _, name := range p.DeniedHeaderNames {
		if !validLowerHeader(name) {
			return unavailable("invalid denied header name")
		}
	}
	required := []string{"authorization", "cookie", "proxy-authorization", "x-api-key", strings.ToLower(lane.Auth.HeaderName)}
	for _, name := range required {
		if !contains(p.DeniedHeaderNames, name) {
			return unavailable("mandatory denied header %q absent", name)
		}
	}
	return nil
}

func strictlySorted(values []string) bool {
	for index := 1; index < len(values); index++ {
		if values[index-1] >= values[index] {
			return false
		}
	}
	return true
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func validLowerHeader(value string) bool {
	if value == "" || value != strings.ToLower(value) {
		return false
	}
	for i := range value {
		if value[i] >= 'a' && value[i] <= 'z' || value[i] >= '0' && value[i] <= '9' || value[i] == '-' {
			continue
		}
		return false
	}
	return true
}

func validDigest(value string) bool {
	if len(value) != sha256.Size*2 || value != strings.ToLower(value) {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}

func unavailable(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrPolicyUnavailable, fmt.Sprintf(format, args...))
}
