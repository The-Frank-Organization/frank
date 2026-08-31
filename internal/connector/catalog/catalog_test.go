package catalog

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/connector/jcs"
)

func TestLoadAcceptsCanonicalClosedCatalogAndReturnsExactDigest(t *testing.T) {
	t.Parallel()

	raw := validCatalogBytes()
	got, err := Load(raw)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if len(got.Lanes) != 1 || got.Lanes[0].LaneID != "lane-codex-1" {
		t.Fatalf("Load() lanes = %#v", got.Lanes)
	}
	wantDigest := sha256.Sum256(raw)
	if got.Digest != hex.EncodeToString(wantDigest[:]) {
		t.Fatalf("Load() digest = %q, want %x", got.Digest, wantDigest)
	}
	if err := got.ValidateDeniedHeaders([]string{"authorization", "cookie", "proxy-authorization", "x-api-key", "x-openai-auth"}); err != nil {
		t.Fatalf("ValidateDeniedHeaders() error = %v", err)
	}
}

func TestLoadRejectsNonCanonicalUnknownAndNonNFCBytes(t *testing.T) {
	t.Parallel()

	nonCanonical := []byte(`{"schema":"m8.lane_catalog.v1","lanes":[]}`)
	if _, err := Load(nonCanonical); !errors.Is(err, ErrInvalidCatalog) {
		t.Fatalf("Load(non-canonical) error = %v", err)
	}

	unknown := strings.Replace(string(validCatalogBytes()), `,"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true}}]`, `,"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true},"zz_unknown":"SENTINEL"}]`, 1)
	if !jcs.IsCanonical([]byte(unknown)) {
		t.Fatal("unknown-field fixture is not canonical")
	}
	if _, err := Load([]byte(unknown)); !errors.Is(err, ErrInvalidCatalog) {
		t.Fatalf("Load(unknown field) error = %v", err)
	}

	nonNFC := strings.Replace(string(validCatalogBytes()), `"region":"us-east"`, `"region":"éu"`, 1)
	if !jcs.IsCanonical([]byte(nonNFC)) {
		t.Fatal("non-NFC fixture is not canonical")
	}
	if _, err := Load([]byte(nonNFC)); !errors.Is(err, ErrInvalidCatalog) {
		t.Fatalf("Load(non-NFC) error = %v", err)
	}

	composed := strings.Replace(string(validCatalogBytes()), `"region":"us-east"`, `"region":"éu"`, 1)
	if !jcs.IsCanonical([]byte(composed)) {
		t.Fatal("composed fixture is not canonical")
	}
	if _, err := Load([]byte(composed)); err != nil {
		t.Fatalf("Load(composed NFC) error = %v", err)
	}
}

func TestValidateRejectsSemanticCatalogCollisionsAndMalformedFacts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		mutate func(*Catalog)
	}{
		{name: "unsorted lane ids", mutate: func(c *Catalog) {
			second := c.Lanes[0]
			second.LaneID = "lane-a"
			second.ModelID = "model-2"
			c.Lanes = append(c.Lanes, second)
		}},
		{name: "duplicate four-axis key", mutate: func(c *Catalog) {
			second := c.Lanes[0]
			second.LaneID = "lane-codex-2"
			c.Lanes = append(c.Lanes, second)
		}},
		{name: "invalid lane id", mutate: func(c *Catalog) { c.Lanes[0].LaneID = "Lane_1" }},
		{name: "invalid serving profile id", mutate: func(c *Catalog) { c.Lanes[0].ServingProfileID = "profile/1" }},
		{name: "invalid endpoint", mutate: func(c *Catalog) { c.Lanes[0].Endpoint = "http://api.openai.com/v1" }},
		{name: "uppercase auth header", mutate: func(c *Catalog) { c.Lanes[0].Auth.HeaderName = "Authorization" }},
		{name: "unsupported method", mutate: func(c *Catalog) { c.Lanes[0].Method = "GET" }},
		{name: "bad effective time", mutate: func(c *Catalog) { c.Lanes[0].Cost.EffectiveTime = "yesterday" }},
		{name: "bad observed at", mutate: func(c *Catalog) { c.Lanes[0].ObservedAt = "soon" }},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			catalog := validCatalog()
			test.mutate(catalog)
			if err := catalog.Validate(); !errors.Is(err, ErrInvalidCatalog) {
				t.Fatalf("Validate() error = %v", err)
			}
		})
	}
}

func TestValidateDeniedHeadersRequiresLaneAuthHeaderMembership(t *testing.T) {
	t.Parallel()

	catalog := validCatalog()
	if err := catalog.ValidateDeniedHeaders([]string{"authorization", "cookie", "proxy-authorization", "x-api-key"}); !errors.Is(err, ErrInvalidCatalog) {
		t.Fatalf("ValidateDeniedHeaders() error = %v", err)
	}
}

func validCatalog() *Catalog {
	catalog, err := Load(validCatalogBytes())
	if err != nil {
		panic(err)
	}
	return catalog
}

func validCatalogBytes() []byte {
	return []byte(`{"lanes":[{"auth":{"auth_header_name":"x-openai-auth","auth_scheme":"bearer"},"compat_mode":"openai-responses","cost":{"effective_time":"2026-07-17T00:00:00Z","input":1,"output":10},"endpoint":"https://api.openai.com/v1/responses","lane_id":"lane-codex-1","limits":{"context":200000,"max_output":100000},"method":"POST","model_id":"gpt-5","observed_at":"2026-07-17T00:00:00Z","profile_facts":{"endpoint_kind":"coding","region":"us-east"},"provider_id":"openai","reasoning":{"effort_levels":["low","medium","high"],"replay_kind":"opaque_item","supported":true},"serving_profile_id":"codex-default","source":"seeded","tool_use":{"strict_schema":true,"supported":true},"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true}}],"schema":"m8.lane_catalog.v1"}`)
}
