package policy

import (
	"errors"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/connector/catalog"
	"github.com/The-Frank-Organization/frank/internal/connector/jcs"
)

const normativePolicy = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`

func TestLoadAcceptsNormativePolicyAndVerifiesExactDigest(t *testing.T) {
	t.Parallel()

	got, err := Load([]byte(normativePolicy), pinnedLane())
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	const wantDigest = "ca364710764c3fb5fa6ca0f2faa6795e6aa6f49d3d55877ca4def10d540a3030"
	if got.Digest != wantDigest {
		t.Fatalf("Load() digest = %q, want %q", got.Digest, wantDigest)
	}
	if err := got.VerifyDigest(wantDigest); err != nil {
		t.Fatalf("VerifyDigest() error = %v", err)
	}
	if err := got.VerifyDigest(strings.Repeat("0", 64)); !errors.Is(err, ErrPolicyDigestMismatch) {
		t.Fatalf("VerifyDigest(mismatch) error = %v", err)
	}
}

func TestLoadRejectsUnavailablePolicyForms(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  string
		lane catalog.Lane
	}{
		{name: "absent", raw: "", lane: pinnedLane()},
		{name: "terminal newline", raw: normativePolicy + "\n", lane: pinnedLane()},
		{name: "unknown method_by_lane", raw: strings.Replace(normativePolicy, `"pinned_lane"`, `"method_by_lane":{},"pinned_lane"`, 1), lane: pinnedLane()},
		{name: "unsorted allowlist", raw: strings.Replace(normativePolicy, `["https://api.openai.com/v1/responses"]`, `["https://z.example/v1","https://a.example/v1"]`, 1), lane: pinnedLane()},
		{name: "duplicate denied header", raw: strings.Replace(normativePolicy, `"cookie",`, `"authorization","cookie",`, 1), lane: pinnedLane()},
		{name: "uppercase denied header", raw: strings.Replace(normativePolicy, `"authorization"`, `"Authorization"`, 1), lane: pinnedLane()},
		{name: "invalid endpoint", raw: strings.Replace(normativePolicy, `https://api.openai.com/v1/responses`, `http://api.openai.com/v1/responses`, 1), lane: pinnedLane()},
		{name: "wrong pinned lane", raw: strings.Replace(normativePolicy, `lane-codex-1`, `lane-codex-2`, 1), lane: pinnedLane()},
		{name: "missing mandatory lane auth header", raw: strings.Replace(normativePolicy, `,"x-openai-auth"`, ``, 1), lane: pinnedLane()},
		{name: "non-NFC policy string", raw: strings.Replace(normativePolicy, `provider-request`, "provider-reque\u0301st", 1), lane: pinnedLane()},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if test.raw != "" && !strings.HasSuffix(test.raw, "\n") && !jcs.IsCanonical([]byte(test.raw)) {
				// Some cases deliberately test canonicality itself; semantic cases
				// below remain canonical so the named validator is exercised.
				switch test.name {
				case "unknown method_by_lane", "unsorted allowlist", "duplicate denied header", "uppercase denied header", "invalid endpoint", "wrong pinned lane", "missing mandatory lane auth header", "non-NFC policy string":
					t.Fatalf("semantic fixture %q is not canonical", test.name)
				}
			}
			if _, err := Load([]byte(test.raw), test.lane); !errors.Is(err, ErrPolicyUnavailable) {
				t.Fatalf("Load() error = %v, want ErrPolicyUnavailable", err)
			}
		})
	}
}

func pinnedLane() catalog.Lane {
	return catalog.Lane{
		LaneID:   "lane-codex-1",
		Endpoint: "https://api.openai.com/v1/responses",
		Auth: catalog.Auth{
			HeaderName: "x-openai-auth",
			Scheme:     "bearer",
		},
	}
}
