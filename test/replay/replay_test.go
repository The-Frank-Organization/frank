package replay

import (
	"os"
	"testing"
)

func TestClassifyAllDisposesEveryRepresentativeFixture(t *testing.T) {
	results := ClassifyAll()
	if len(results) == 0 {
		t.Fatalf("no replay results")
	}
	for _, result := range results {
		if result.Disposition == "" {
			t.Fatalf("missing disposition: %+v", result)
		}
	}
	assertDisposition(t, results, "bad-phase-token", "caught")
	assertDisposition(t, results, "bare-token-shape", "obsolete-by-construction")
	assertDisposition(t, results, "scope-diff-row-array", "uncovered-S3")
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
