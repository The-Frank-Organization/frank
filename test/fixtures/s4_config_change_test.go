package fixtures_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestRunningConfigUnchangedUntilRestart(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	h.start(t)

	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	before, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools before: %v", err)
	}

	body := mutatedRegistryBody(t, h.root)
	rec := record.Record{
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"AUTHORITY":     "report-only",
			"CEREMONY_TIER": "medium",
			"SUBJECT":       "registry config_change",
			"record_kind":   "config_change",
			"member":        "fieldspec",
			"new_digest":    fixtureDigestWithMember(t, h.root, "fieldspec", body),
		},
		Body: string(body),
	}
	h.submit(t, operator, rec)

	materialized, err := os.ReadFile(filepath.Join(h.root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read materialized registry: %v", err)
	}
	if string(materialized) != string(body) {
		t.Fatalf("materialized registry differs from submitted body")
	}
	after, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools after: %v", err)
	}
	if after.FormDigest != before.FormDigest {
		t.Fatalf("running registry hot-reloaded: before digest %s after %s", before.FormDigest, after.FormDigest)
	}
}

func mutatedRegistryBody(t *testing.T, root string) []byte {
	t.Helper()
	data := mustReadFile(t, filepath.Join(root, "config", "fieldspec", "registry.json"))
	var doc map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("decode registry: %v", err)
	}
	provenance, ok := doc["provenance"].(map[string]any)
	if !ok {
		provenance = map[string]any{}
		doc["provenance"] = provenance
	}
	provenance["s4_test_marker"] = "running-config-unchanged"
	body, err := json.Marshal(doc)
	if err != nil {
		t.Fatalf("marshal mutated registry: %v", err)
	}
	return body
}

func fixtureDigestWithMember(t *testing.T, root, member string, body []byte) string {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store config: %v", err)
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return config.Digest(members)
}
