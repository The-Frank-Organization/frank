// Package authorize implements the pure, deterministic m-3 provider-request
// authorization predicate over an m-8 frozen request core.
package authorize

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/connector/catalog"
	"github.com/The-Frank-Organization/frank/internal/connector/endpoint"
	"github.com/The-Frank-Organization/frank/internal/connector/frame"
	"github.com/The-Frank-Organization/frank/internal/connector/policy"
)

type DenyReason string

const (
	PolicyUnavailable        DenyReason = "policy-unavailable"
	PolicyDigestMismatch     DenyReason = "policy-digest-mismatch"
	MalformedCore            DenyReason = "malformed-core"
	LaneMismatch             DenyReason = "lane-mismatch"
	LaneEndpointInvalid      DenyReason = "lane-endpoint-invalid"
	EndpointMismatch         DenyReason = "endpoint-mismatch"
	EndpointNotAllowlisted   DenyReason = "endpoint-not-allowlisted"
	MethodMismatch           DenyReason = "method-mismatch"
	ReservedAuthHeaderInCore DenyReason = "reserved-auth-header-in-core"
)

type Header struct {
	Name  string
	Value string
}

// Core carries only policy-relevant facts. BodyIsDesignated is the m-8
// producer assertion that the body is the pinned lane's designated encoding;
// authorization deliberately does not parse the body.
type Core struct {
	Method           string
	Endpoint         string
	Headers          []Header
	BodySHA256       string
	BodyLen          string
	BodyIsDesignated bool
}

type Input struct {
	Policy           *policy.Policy
	PolicyDigest     string
	ProviderLaneID   string
	PinnedLaneID     string
	FrozenCoreDigest string
	CredentialRef    string
	Lane             catalog.Lane
	Core             Core
}

type Verdict struct {
	Allowed    bool
	DenyReason DenyReason
	Envelope   *AuthorizedEnvelope
	binding    string
}

type AuthProfile struct {
	HeaderName string
	Scheme     string
}

type AuthorizedEnvelope struct {
	FrozenCoreDigest string
	AuthProfile      AuthProfile
	CredentialRef    string
}

// Authorized returns a copy only while the allowed verdict's envelope still
// matches the identity bound by Evaluate.
func (verdict Verdict) Authorized() (AuthorizedEnvelope, bool) {
	if !verdict.Allowed || verdict.DenyReason != "" || verdict.Envelope == nil || verdict.binding == "" {
		return AuthorizedEnvelope{}, false
	}
	if envelopeBinding(*verdict.Envelope) != verdict.binding {
		return AuthorizedEnvelope{}, false
	}
	return *verdict.Envelope, true
}

// Evaluate applies the closed nine-reason order and stops at the first
// failure. Epoch authority is intentionally absent from this predicate.
func Evaluate(input Input) Verdict {
	policyLane := input.Lane
	policyLane.LaneID = input.PinnedLaneID
	if input.Policy == nil || input.Policy.ValidateForLane(policyLane) != nil {
		return deny(PolicyUnavailable)
	}
	if input.Policy.VerifyDigest(input.PolicyDigest) != nil {
		return deny(PolicyDigestMismatch)
	}
	if !validCore(input.Core) || !validDigest(input.FrozenCoreDigest) {
		return deny(MalformedCore)
	}
	if input.ProviderLaneID != input.PinnedLaneID || input.Lane.LaneID != input.PinnedLaneID {
		return deny(LaneMismatch)
	}
	if endpoint.Validate(input.Lane.Endpoint) != nil {
		return deny(LaneEndpointInvalid)
	}
	if input.Core.Endpoint != input.Lane.Endpoint {
		return deny(EndpointMismatch)
	}
	if !input.Policy.AllowsEndpoint(input.Core.Endpoint) {
		return deny(EndpointNotAllowlisted)
	}
	if input.Core.Method != input.Lane.Method {
		return deny(MethodMismatch)
	}
	for _, header := range input.Core.Headers {
		if input.Policy.DeniesHeader(header.Name) {
			return deny(ReservedAuthHeaderInCore)
		}
	}
	envelope := AuthorizedEnvelope{
		FrozenCoreDigest: input.FrozenCoreDigest,
		AuthProfile: AuthProfile{
			HeaderName: input.Lane.Auth.HeaderName,
			Scheme:     input.Lane.Auth.Scheme,
		},
		CredentialRef: input.CredentialRef,
	}
	return Verdict{Allowed: true, Envelope: &envelope, binding: envelopeBinding(envelope)}
}

func validCore(core Core) bool {
	if core.Method == "" || endpoint.Validate(core.Endpoint) != nil || !core.BodyIsDesignated || !validDigest(core.BodySHA256) {
		return false
	}
	if _, err := frame.ParseCounter(core.BodyLen); err != nil {
		return false
	}
	for _, header := range core.Headers {
		if !validLowerHeaderName(header.Name) || !validHeaderValue(header.Value) {
			return false
		}
	}
	return true
}

func validLowerHeaderName(value string) bool {
	if value == "" || value != strings.ToLower(value) {
		return false
	}
	for index := range value {
		b := value[index]
		if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || strings.ContainsRune("!#$%&'*+-.^_`|~", rune(b)) {
			continue
		}
		return false
	}
	return true
}

func validHeaderValue(value string) bool {
	for index := range value {
		if value[index] == '\r' || value[index] == '\n' || value[index] == 0x00 {
			return false
		}
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

func deny(reason DenyReason) Verdict {
	return Verdict{DenyReason: reason}
}

func envelopeBinding(envelope AuthorizedEnvelope) string {
	sum := sha256.Sum256([]byte(envelope.FrozenCoreDigest + "\x00" + envelope.AuthProfile.HeaderName + "\x00" + envelope.AuthProfile.Scheme + "\x00" + envelope.CredentialRef))
	return hex.EncodeToString(sum[:])
}
