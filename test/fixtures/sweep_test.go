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

func hasExclusivityClaim(text string) bool {
	for _, token := range []string{"only writer", "sole", "no lane can", "non-lane-writable", "unbypassable"} {
		if strings.Contains(text, token) {
			return true
		}
	}
	return false
}
