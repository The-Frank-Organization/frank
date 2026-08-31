package authorize

import (
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/connector/catalog"
	"github.com/The-Frank-Organization/frank/internal/connector/policy"
)

const policyBytes = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`

func TestEvaluateReturnsFirstFailureInNormativeOrder(t *testing.T) {
	t.Parallel()

	valid := validInput(t)
	tests := []struct {
		name   string
		mutate func(*Input)
		want   DenyReason
	}{
		{name: "policy unavailable", mutate: func(input *Input) { input.Policy = nil }, want: PolicyUnavailable},
		{name: "policy digest mismatch", mutate: func(input *Input) { input.PolicyDigest = strings.Repeat("0", 64) }, want: PolicyDigestMismatch},
		{name: "malformed core", mutate: func(input *Input) { input.Core.BodyIsDesignated = false }, want: MalformedCore},
		{name: "lane mismatch", mutate: func(input *Input) { input.ProviderLaneID = "lane-other" }, want: LaneMismatch},
		{name: "lane endpoint invalid", mutate: func(input *Input) { input.Lane.Endpoint = "http://api.openai.com/v1/responses" }, want: LaneEndpointInvalid},
		{name: "endpoint mismatch", mutate: func(input *Input) { input.Lane.Endpoint = "https://other.example/v1/responses" }, want: EndpointMismatch},
		{name: "endpoint not allowlisted", mutate: func(input *Input) {
			input.Core.Endpoint = "https://other.example/v1/responses"
			input.Lane.Endpoint = input.Core.Endpoint
		}, want: EndpointNotAllowlisted},
		{name: "method mismatch", mutate: func(input *Input) { input.Core.Method = "GET" }, want: MethodMismatch},
		{name: "reserved auth header in core", mutate: func(input *Input) {
			input.Core.Headers = append(input.Core.Headers, Header{Name: "authorization", Value: "name-only-check"})
		}, want: ReservedAuthHeaderInCore},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			input := valid
			input.Core.Headers = append([]Header(nil), valid.Core.Headers...)
			test.mutate(&input)
			for iteration := 0; iteration < 10; iteration++ {
				verdict := Evaluate(input)
				if verdict.Allowed || verdict.DenyReason != test.want {
					t.Fatalf("Evaluate() = %+v, want denied %q", verdict, test.want)
				}
			}
		})
	}
}

func TestEvaluateNormativeMultiFailureStopsAtLaneMismatch(t *testing.T) {
	t.Parallel()

	input := validInput(t)
	input.ProviderLaneID = "lane-other"
	input.Core.Endpoint = "https://not-allowed.example/v1/responses"
	input.Core.Headers = append(input.Core.Headers, Header{Name: "authorization", Value: "not-inspected"})

	verdict := Evaluate(input)
	if verdict.Allowed || verdict.DenyReason != LaneMismatch {
		t.Fatalf("Evaluate() = %+v, want denied %q", verdict, LaneMismatch)
	}
}

func TestEvaluateLaneFactIDMismatchIsLaneMismatchNotPolicyUnavailable(t *testing.T) {
	t.Parallel()

	input := validInput(t)
	input.Lane.LaneID = "lane-other"
	verdict := Evaluate(input)
	if verdict.Allowed || verdict.DenyReason != LaneMismatch {
		t.Fatalf("Evaluate() = %+v, want denied %q", verdict, LaneMismatch)
	}
}

func TestEvaluateAllowsValidFrozenCore(t *testing.T) {
	t.Parallel()

	verdict := Evaluate(validInput(t))
	if !verdict.Allowed || verdict.DenyReason != "" || verdict.Envelope == nil {
		t.Fatalf("Evaluate() = %+v, want allowed with no deny reason", verdict)
	}
	if verdict.Envelope.FrozenCoreDigest != strings.Repeat("b", 64) || verdict.Envelope.CredentialRef != "provider-main" || verdict.Envelope.AuthProfile.HeaderName != "x-openai-auth" || verdict.Envelope.AuthProfile.Scheme != "bearer" {
		t.Fatalf("authorized envelope = %+v", verdict.Envelope)
	}
}

func TestEvaluateTreatsMalformedHeaderAsMalformedCoreBeforePolicyNames(t *testing.T) {
	t.Parallel()

	input := validInput(t)
	input.Core.Headers = []Header{{Name: "Authorization", Value: "not-canonical"}}
	verdict := Evaluate(input)
	if verdict.Allowed || verdict.DenyReason != MalformedCore {
		t.Fatalf("Evaluate() = %+v, want denied %q", verdict, MalformedCore)
	}
}

func TestEvaluateRejectsNonCanonicalBodyLength(t *testing.T) {
	t.Parallel()

	input := validInput(t)
	input.Core.BodyLen = "02"
	verdict := Evaluate(input)
	if verdict.Allowed || verdict.DenyReason != MalformedCore {
		t.Fatalf("Evaluate() = %+v, want denied %q", verdict, MalformedCore)
	}
}

func validInput(t *testing.T) Input {
	t.Helper()
	lane := catalog.Lane{
		LaneID:   "lane-codex-1",
		Endpoint: "https://api.openai.com/v1/responses",
		Method:   "POST",
		Auth: catalog.Auth{
			HeaderName: "x-openai-auth",
			Scheme:     "bearer",
		},
	}
	loaded, err := policy.Load([]byte(policyBytes), lane)
	if err != nil {
		t.Fatalf("policy.Load() error = %v", err)
	}
	return Input{
		Policy:           loaded,
		PolicyDigest:     loaded.Digest,
		ProviderLaneID:   lane.LaneID,
		PinnedLaneID:     lane.LaneID,
		FrozenCoreDigest: strings.Repeat("b", 64),
		CredentialRef:    "provider-main",
		Lane:             lane,
		Core: Core{
			Method:           "POST",
			Endpoint:         lane.Endpoint,
			Headers:          []Header{{Name: "content-type", Value: "application/json"}},
			BodySHA256:       strings.Repeat("a", 64),
			BodyLen:          "2",
			BodyIsDesignated: true,
		},
	}
}
