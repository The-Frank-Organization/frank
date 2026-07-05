package fixtures_test

import (
	"os"
	"strings"
	"testing"
)

func TestSweepReadmeClaimHonesty(t *testing.T) {
	data, err := os.ReadFile("../../README.md")
	if err != nil {
		t.Fatalf("README missing: %v", err)
	}
	readme := string(data)
	for _, want := range []string{
		"S1 = provenance + transport, not verified work",
		"self_reported",
		"tool-mediated confusion-resistance",
		"D5 residual",
		"pair-Planner grant rendering lands in S3",
		"the S3 registry rides `store.Init`",
		"registry evolution on an existing store awaits the §7 config-change record",
	} {
		if !strings.Contains(readme, want) {
			t.Fatalf("README missing honesty phrase %q", want)
		}
	}
	for _, forbidden := range []string{"bounced", "submitted"} {
		if strings.Contains(readme, forbidden) {
			t.Fatalf("README contains forbidden value token %q", forbidden)
		}
	}
	for _, paragraph := range strings.Split(readme, "\n\n") {
		if hasExclusivityClaim(paragraph) && !strings.Contains(paragraph, "D5 residual") && !strings.Contains(paragraph, "governance-surface") {
			t.Fatalf("claim lacks D5/governance qualifier: %q", paragraph)
		}
	}
}

func TestS3SweepGateCategoryAuthorityUsesRegistry(t *testing.T) {
	data, err := os.ReadFile("../../internal/lineage/lineage.go")
	if err != nil {
		t.Fatalf("read lineage.go: %v", err)
	}
	helperName := "isA" + "GateCategory"
	if strings.Contains(string(data), helperName) {
		t.Fatalf("lineage still carries the S1 hard-coded gate-category helper")
	}
}

func hasExclusivityClaim(text string) bool {
	for _, token := range []string{"only writer", "sole", "no lane can", "non-lane-writable", "unbypassable"} {
		if strings.Contains(text, token) {
			return true
		}
	}
	return false
}
