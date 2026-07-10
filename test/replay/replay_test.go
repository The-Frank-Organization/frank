package replay

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

// requireOracle skips oracle-replay tests when the external relay-lint
// tooling (prior-art reference material, not vendored here) is absent.
func requireOracle(t *testing.T) {
	t.Helper()
	if !OracleAvailable() {
		t.Skip("relay-lint oracle not present; set FRANK_RELAY_LINT_TOOLS_DIR to its tools/ directory to run the oracle replay")
	}
}

func TestFullOracleReplayBothLegsGreen(t *testing.T) {
	requireOracle(t)
	results := ReplayAll()
	if len(results) != 146 {
		t.Fatalf("ReplayAll returned %d results, want frozen oracle size 146", len(results))
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
	contexts := map[string]bool{"live": true, "reconstructed-observe": true}
	for _, row := range rows {
		if row.Anchor == "" || row.Surface == "" {
			t.Fatalf("incomplete disposition row: %+v", row)
		}
		if !allowed[row.Disposition] {
			t.Fatalf("unexpected disposition %q in %+v", row.Disposition, row)
		}
		if !contexts[row.Context] {
			t.Fatalf("unexpected context %q in %+v", row.Context, row)
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

func TestDispositionTableMatchesRelayLintInventory(t *testing.T) {
	rows := loadDispositionRows(t)
	rowAnchors := map[string]bool{}
	symbolic := regexp.MustCompile(`^relay-lint\.py:[0-9]+$`)
	for _, row := range rows {
		if !symbolic.MatchString(row.Anchor) {
			t.Fatalf("non-code anchor %s", row.Anchor)
		}
		rowAnchors[row.Anchor] = true
	}
	requireOracle(t)
	inventory, err := RelayLintInventory()
	if err != nil {
		t.Fatalf("RelayLintInventory: %v", err)
	}
	var missing []string
	for _, anchor := range inventory {
		if !rowAnchors[anchor] {
			missing = append(missing, anchor)
		}
	}
	if len(missing) > 0 {
		t.Fatalf("disposition table missing inventory anchors: %v", missing)
	}
	for _, anchor := range []string{"relay-lint.py:844", "relay-lint.py:846", "relay-lint.py:848", "relay-lint.py:855", "relay-lint.py:857", "relay-lint.py:859", "relay-lint.py:861", "relay-lint.py:863", "relay-lint.py:868", "relay-lint.py:873"} {
		if !rowAnchors[anchor] {
			t.Fatalf("explicit addressed-token anchor missing: %s", anchor)
		}
	}
}

func TestObserveContextLabelAndTallyAreStructural(t *testing.T) {
	rows := loadDispositionRows(t)
	var live, reconstructed int
	for _, row := range rows {
		switch row.Context {
		case "live":
			live++
		case "reconstructed-observe":
			reconstructed++
			if row.Check != "caught — surviving §10b rule, fires under the reconstructed `layer_present:observe` context" {
				t.Fatalf("observe row label = %q", row.Check)
			}
		}
	}
	if reconstructed == 0 {
		t.Fatalf("no reconstructed-observe rows")
	}
	requireOracle(t)
	results := ReplayAll()
	if got := LiveCaughtCount(results); got >= live+reconstructed {
		t.Fatalf("live caught tally %d appears to include reconstructed rows; live rows=%d reconstructed=%d", got, live, reconstructed)
	}
}

func TestObsoleteGroundVocabularyIsClosed(t *testing.T) {
	allowed := map[string]bool{
		"vanished-markdown-channel":   true,
		"strict-submit-api":           true,
		"one-channel-role-from-stamp": true,
	}
	for _, row := range loadDispositionRows(t) {
		if row.Disposition != "obsolete" {
			continue
		}
		if !allowed[row.ObsoleteGround] {
			t.Fatalf("obsolete ground %q not in closed vocabulary for %+v", row.ObsoleteGround, row)
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
	Anchor         string   `json:"anchor"`
	Check          string   `json:"check"`
	Class10        string   `json:"class_10"`
	Disposition    string   `json:"disposition"`
	Surface        string   `json:"surface"`
	Context        string   `json:"context"`
	ObsoleteGround string   `json:"obsolete_ground"`
	Fixtures       []string `json:"fixtures"`
}

func generateDispositionTableForTest(rows []dispositionRow) string {
	var b strings.Builder
	b.WriteString("# S3 Disposition Table\n\n")
	b.WriteString("Generated pair: `test/replay/dispositions.json` is the machine-readable source; this table is the human sprint artifact. Fresh-store qualifier: the S3 registry rides `store.Init`; registry evolution on an existing store awaits the §7 config-change record.\n\n")
	b.WriteString("| anchor | check | class | disposition | surface | context | obsolete_ground | fixtures |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- |\n")
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
		b.WriteString(row.Context)
		b.WriteString(" | ")
		b.WriteString(row.ObsoleteGround)
		b.WriteString(" | ")
		b.WriteString(strings.Join(row.Fixtures, "; "))
		b.WriteString(" |\n")
	}
	return b.String()
}

func loadDispositionRows(t *testing.T) []dispositionRow {
	t.Helper()
	data, err := os.ReadFile("dispositions.json")
	if err != nil {
		t.Fatalf("read dispositions.json: %v", err)
	}
	var rows []dispositionRow
	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatalf("decode dispositions.json: %v", err)
	}
	if reflect.ValueOf(rows).Len() == 0 {
		t.Fatalf("empty disposition rows")
	}
	return rows
}
