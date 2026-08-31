package fixtures_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/crashpoint"
	fixtures "github.com/The-Frank-Organization/frank/test/fixtures"
)

func TestS2ApplicabilityMapCoversEveryClassPointAndReport(t *testing.T) {
	classMap := fixtures.ApplicabilityMap()
	for _, class := range []string{
		"submit-accept",
		"submit-reject",
		"held",
		"operator-verdict",
		"park",
		"outbox-enqueue",
		"genesis",
		"quarantine-disposition",
		"gc-marker",
		"owed-item",
		"owed-disposition",
	} {
		row := classMap[class]
		if row == nil {
			t.Fatalf("missing class row %s", class)
		}
		for _, name := range crashpoint.Names() {
			got := row[name]
			if !slices.Contains([]string{"crash-expected", "clean-completion"}, got) {
				t.Fatalf("class %s crashpoint %s = %q", class, name, got)
			}
		}
	}
	reportPath := filepath.Join("..", "..", "docs", "sprints", "2026-07-03-s2-slice-2", "results", "f11-sweep-report.md")
	report, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("read sweep report: %v", err)
	}
	for _, want := range []string{"Executed crash-expected cells", "Hit-trace row equality", "FRANK_TEST_HIT_TRACE", "Clean-completion evidence", "submit-accept", "owed-item", "gc-marker", "pre_gc_marker", "recovery_post_phase4", "O3 fixture proof", "OI-S1-F11-SWEEP"} {
		if !strings.Contains(string(report), want) {
			t.Fatalf("report missing %q", want)
		}
	}
}
