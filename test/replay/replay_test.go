package replay

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
)

func TestFullOracleReplayBothLegsGreen(t *testing.T) {
	results := ClassifyAll()
	if len(results) != 146 {
		t.Fatalf("ClassifyAll returned %d results, want frozen oracle size 146", len(results))
	}
	var failSide, passSide int
	for _, result := range results {
		if result.Disposition == "" {
			t.Fatalf("missing disposition: %+v", result)
		}
		if strings.Contains(result.Disposition, "uncovered") {
			t.Fatalf("uncovered replay bucket survived: %+v", result)
		}
		switch result.Disposition {
		case "accepted":
			passSide++
		case "caught", "genuinely-obsolete":
			failSide++
		default:
			t.Fatalf("%s has unexpected disposition %s", result.Name, result.Disposition)
		}
	}
	if failSide != 96 || passSide != 50 {
		t.Fatalf("oracle split fail/pass = %d/%d, want 96/50", failSide, passSide)
	}
	assertOutcome(t, results, "claude/B9-bad-enum.md", "caught")
	assertOutcome(t, results, "fold/FD1-fold-edit-no-foldscope.md", "caught")
	assertOutcome(t, results, "addressing/T2-token-no-to.md", "genuinely-obsolete")
	assertOutcome(t, results, "addressing/G1-casefold-lineage", "accepted")
}

func TestStaleRepresentativeReplayArtifactsAreRemoved(t *testing.T) {
	for _, path := range []string{"classmap.go", "report.md"} {
		if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("%s still exists or stat failed unexpectedly: %v", path, err)
		}
	}
}

func TestDispositionArtifactsArePresentAndCovered(t *testing.T) {
	data, err := os.ReadFile("dispositions.json")
	if err != nil {
		t.Fatalf("read dispositions.json: %v", err)
	}
	var rows []dispositionRow
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
	wantTable := generateDispositionTableForTest(rows)
	table, err := os.ReadFile("../../docs/sprints/2026-07-04-s3-slice-3/results/disposition-table.md")
	if err != nil {
		t.Fatalf("read disposition-table.md: %v", err)
	}
	if string(table) != wantTable {
		t.Fatalf("disposition table does not match generated output\nwant:\n%s\ngot:\n%s", wantTable, table)
	}
	for _, row := range rows {
		if !strings.Contains(string(table), row.Anchor) {
			t.Fatalf("table missing anchor %s", row.Anchor)
		}
	}
}

func assertOutcome(t *testing.T, results []Result, name, outcome string) {
	t.Helper()
	for _, result := range results {
		if result.Name == name {
			if result.Disposition != outcome {
				t.Fatalf("%s outcome = %s, want %s", name, result.Disposition, outcome)
			}
			return
		}
	}
	t.Fatalf("missing result %s", name)
}

type dispositionRow struct {
	Anchor      string   `json:"anchor"`
	Check       string   `json:"check"`
	Class10     string   `json:"class_10"`
	Disposition string   `json:"disposition"`
	Surface     string   `json:"surface"`
	Fixtures    []string `json:"fixtures"`
}

func generateDispositionTableForTest(rows []dispositionRow) string {
	var b strings.Builder
	b.WriteString("# S3 Disposition Table\n\n")
	b.WriteString("Generated pair: `test/replay/dispositions.json` is the machine-readable source; this table is the human sprint artifact.\n\n")
	b.WriteString("| anchor | check | class | disposition | surface | fixtures |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- |\n")
	for _, row := range rows {
		b.WriteString("| ")
		b.WriteString(row.Anchor)
		b.WriteString(" | ")
		b.WriteString(row.Check)
		b.WriteString(" | ")
		b.WriteString(row.Class10)
		b.WriteString(" | ")
		b.WriteString(row.Disposition)
		b.WriteString(" | ")
		b.WriteString(row.Surface)
		b.WriteString(" | ")
		b.WriteString(strings.Join(row.Fixtures, "; "))
		b.WriteString(" |\n")
	}
	return b.String()
}
