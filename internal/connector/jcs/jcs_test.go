package jcs

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"testing"
)

func TestCanonicalizeNumbersUseECMAScriptSpelling(t *testing.T) {
	t.Parallel()

	in := []byte(`[333333333.33333329,1E30,4.50,2e-3,1e-27,0.000001,1e-7,100000000000000000000,1e21,-0]`)
	want := []byte(`[333333333.3333333,1e+30,4.5,0.002,1e-27,0.000001,1e-7,100000000000000000000,1e+21,0]`)

	got, err := Canonicalize(in)
	if err != nil {
		t.Fatalf("Canonicalize() error = %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("Canonicalize() = %s, want %s", got, want)
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

func TestCanonicalNumberMatchesRFC8785AppendixB(t *testing.T) {
	t.Parallel()

	tests := []struct {
		bits uint64
		want string
	}{
		{0x0000000000000000, "0"},
		{0x8000000000000000, "0"},
		{0x0000000000000001, "5e-324"},
		{0x8000000000000001, "-5e-324"},
		{0x7fefffffffffffff, "1.7976931348623157e+308"},
		{0xffefffffffffffff, "-1.7976931348623157e+308"},
		{0x4340000000000000, "9007199254740992"},
		{0xc340000000000000, "-9007199254740992"},
		{0x4430000000000000, "295147905179352830000"},
		{0x44b52d02c7e14af5, "9.999999999999997e+22"},
		{0x44b52d02c7e14af6, "1e+23"},
		{0x44b52d02c7e14af7, "1.0000000000000001e+23"},
		{0x444b1ae4d6e2ef4e, "999999999999999700000"},
		{0x444b1ae4d6e2ef4f, "999999999999999900000"},
		{0x444b1ae4d6e2ef50, "1e+21"},
		{0x3eb0c6f7a0b5ed8c, "9.999999999999997e-7"},
		{0x3eb0c6f7a0b5ed8d, "0.000001"},
		{0x41b3de4355555553, "333333333.3333332"},
		{0x41b3de4355555554, "333333333.33333325"},
		{0x41b3de4355555555, "333333333.3333333"},
		{0x41b3de4355555556, "333333333.3333334"},
		{0x41b3de4355555557, "333333333.33333343"},
		{0xbecbf647612f3696, "-0.0000033333333333333333"},
		{0x43143ff3c1cb0959, "1424953923781206.2"},
	}

	for _, test := range tests {
		f := math.Float64frombits(test.bits)
		raw := strconv.FormatFloat(f, 'g', -1, 64)
		got, err := canonicalNumber(raw)
		if err != nil {
			t.Fatalf("canonicalNumber(%x) error = %v", test.bits, err)
		}
		if got != test.want {
			t.Errorf("canonicalNumber(%x) = %q, want %q", test.bits, got, test.want)
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
