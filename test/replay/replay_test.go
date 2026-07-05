package replay

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestClassifyAllDisposesEveryRepresentativeFixture(t *testing.T) {
	results := ClassifyAll()
	if len(results) < 20 {
		t.Fatalf("ClassifyAll returned %d results, want real corpus enumeration", len(results))
	}
	for _, result := range results {
		if result.Disposition == "" {
			t.Fatalf("missing disposition: %+v", result)
		}
	}
	assertDisposition(t, results, "bad-phase-token", "caught")
	assertDisposition(t, results, "bare-token-shape", "obsolete-by-construction")
	assertDisposition(t, results, "scope-diff-row-array", "uncovered-S3")
	assertDisposition(t, results, "claude/B9-bad-enum.md", "caught")
	assertDisposition(t, results, "fold/FD1-fold-edit-no-foldscope.md", "uncovered-S3")
}

func TestReportArtifactMatchesGenerated(t *testing.T) {
	want := GenerateReport(ClassifyAll())
	got, err := os.ReadFile("report.md")
	if err != nil {
		t.Fatalf("read report.md: %v", err)
	}
	if string(got) != want {
		t.Fatalf("report.md does not match generated output\nwant:\n%s\ngot:\n%s", want, got)
	}
}

func TestDispositionArtifactsArePresentAndCovered(t *testing.T) {
	data, err := os.ReadFile("dispositions.json")
	if err != nil {
		t.Fatalf("read dispositions.json: %v", err)
	}
	var rows []struct {
		Anchor      string   `json:"anchor"`
		Disposition string   `json:"disposition"`
		Surface     string   `json:"surface"`
		Fixtures    []string `json:"fixtures"`
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatalf("decode dispositions.json: %v", err)
	}
	if len(rows) == 0 {
		t.Fatalf("empty disposition table")
	}
	allowed := map[string]bool{"dissolved-form": true, "dissolved-lineage": true, "retained": true, "obsolete": true}
	for _, row := range rows {
		if row.Anchor == "" || row.Surface == "" || len(row.Fixtures) == 0 {
			t.Fatalf("incomplete disposition row: %+v", row)
		}
		if !allowed[row.Disposition] {
			t.Fatalf("unexpected disposition %q in %+v", row.Disposition, row)
		}
	}
	table, err := os.ReadFile("../../docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md")
	if err != nil {
		t.Fatalf("read disposition-table.md: %v", err)
	}
	for _, row := range rows {
		if !strings.Contains(string(table), row.Anchor) {
			t.Fatalf("table missing anchor %s", row.Anchor)
		}
	}
}

func assertDisposition(t *testing.T, results []Result, name, disposition string) {
	t.Helper()
	for _, result := range results {
		if result.Name == name {
			if result.Disposition != disposition {
				t.Fatalf("%s disposition = %s, want %s", name, result.Disposition, disposition)
			}
			return
		}
	}
	t.Fatalf("missing result %s", name)
}
