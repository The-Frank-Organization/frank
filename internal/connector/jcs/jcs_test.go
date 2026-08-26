package jcs

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestCanonicalizeRejectsNonIntegerNumbers(t *testing.T) {
	t.Parallel()
	for _, input := range []string{"-0", "1.0", "1e2", "1E2", "0.5"} {
		if got, err := Canonicalize([]byte(input)); err == nil {
			t.Fatalf("Canonicalize(%q) = %s, want refusal", input, got)
		}
	}
}

func TestCanonicalizeFrozenPolicyVector(t *testing.T) {
	t.Parallel()

	policy := []byte(`{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`)
	if len(policy) != 255 {
		t.Fatalf("frozen policy length = %d, want 255", len(policy))
	}
	got, err := Canonicalize(policy)
	if err != nil {
		t.Fatalf("Canonicalize() error = %v", err)
	}
	if !bytes.Equal(got, policy) {
		t.Fatalf("Canonicalize() changed frozen policy bytes: %s", got)
	}
	digest := sha256.Sum256(got)
	if gotDigest := hex.EncodeToString(digest[:]); gotDigest != "ca364710764c3fb5fa6ca0f2faa6795e6aa6f49d3d55877ca4def10d540a3030" {
		t.Fatalf("policy digest = %s", gotDigest)
	}
}

func TestCanonicalizeCarriesArbitraryPrecisionIntegersVerbatim(t *testing.T) {
	t.Parallel()
	for _, input := range []string{"0", "-1", "9007199254740993", "1234567890123456789012345678901234567890"} {
		got, err := Canonicalize([]byte(input))
		if err != nil || string(got) != input {
			t.Fatalf("Canonicalize(%q) = %q, %v", input, got, err)
		}
	}
}

func TestCanonicalizeOrdersObjectNamesByUTF16CodeUnits(t *testing.T) {
	t.Parallel()

	in := []byte("{\"\\ufb33\":1,\"€\":2,\"😀\":3,\"ö\":4,\"1\":5,\"\\r\":6}")
	want := []byte("{\"\\r\":6,\"1\":5,\"ö\":4,\"€\":2,\"😀\":3,\"דּ\":1}")

	got, err := Canonicalize(in)
	if err != nil {
		t.Fatalf("Canonicalize() error = %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("Canonicalize() = %s, want %s", got, want)
	}
}

func TestCanonicalizePreservesStringCodePointsWithoutHTMLOrNFCEscaping(t *testing.T) {
	t.Parallel()

	in := []byte("{\"value\":\"e\u0301<>&/\\u000f\\n\"}")
	want := []byte("{\"value\":\"e\u0301<>&/\\u000f\\n\"}")

	got, err := Canonicalize(in)
	if err != nil {
		t.Fatalf("Canonicalize() error = %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("Canonicalize() = %q, want %q", got, want)
	}
}

func TestCanonicalizeRejectsDuplicateDecodedObjectNames(t *testing.T) {
	t.Parallel()

	if _, err := Canonicalize([]byte("{\"a\":1,\"\\u0061\":2}")); err == nil {
		t.Fatal("Canonicalize() accepted duplicate decoded object names")
	}
}

func TestCanonicalizeRejectsLoneSurrogates(t *testing.T) {
	t.Parallel()

	for _, in := range []string{"{\"value\":\"\\ud800\"}", "{\"value\":\"\\udc00\"}"} {
		if _, err := Canonicalize([]byte(in)); err == nil {
			t.Fatalf("Canonicalize(%q) accepted a lone surrogate", in)
		}
	}
}

func TestCanonicalizeRejectsTrailingJSON(t *testing.T) {
	t.Parallel()

	if _, err := Canonicalize([]byte(`{} {}`)); err == nil {
		t.Fatal("Canonicalize() accepted trailing JSON")
	}
}

func TestIsCanonicalRequiresExactJCSBytes(t *testing.T) {
	t.Parallel()

	if IsCanonical([]byte(`{ "a": 1 }`)) {
		t.Fatal("IsCanonical() accepted insignificant whitespace")
	}
	if !IsCanonical([]byte("{\"a\":1}")) {
		t.Fatal("IsCanonical() rejected canonical bytes")
	}
}
