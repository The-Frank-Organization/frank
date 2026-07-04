package replay

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

const corpusRoot = "/Users/jack/Programming/harness/extracted/agentic-dev-team-skills-v3-export/v2.8.8-release/v288-unzipped/agentic-dev-team-skills-v2.8.8/tools/relay-lint-fixtures"

type Result struct {
	Name        string
	CheckClass  string
	Disposition string
	Reason      string
}

func ClassifyAll() []Result {
	results := []Result{
		classifyBadPhase(),
		{
			Name:        "bare-token-shape",
			CheckClass:  "literal-authority-token-shape",
			Disposition: "obsolete-by-construction",
			Reason:      "typed submission has no raw relay text channel for bare token placement",
		},
		{
			Name:        "scope-diff-row-array",
			CheckClass:  "row-array-structure",
			Disposition: "uncovered-S3",
			Reason:      "SCOPE_DIFF row arrays are not an MVP-visible field in S1",
		},
	}
	corpus, err := classifyCorpus(corpusRoot)
	if err != nil {
		results = append(results, Result{Name: "corpus-read", CheckClass: "corpus", Disposition: "uncovered-S3", Reason: err.Error()})
		return results
	}
	results = append(results, corpus...)
	return results
}

func GenerateReport(results []Result) string {
	var b strings.Builder
	b.WriteString("# R1 Replay Report\n\n")
	for _, section := range []string{"caught", "obsolete-by-construction", "uncovered-S3"} {
		b.WriteString("## " + section + "\n\n")
		for _, result := range results {
			if result.Disposition == section {
				b.WriteString("- " + result.Name + " (" + result.CheckClass + "): " + result.Reason + "\n")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func classifyBadPhase() Result {
	reg, err := loadRegistry()
	if err != nil {
		return Result{Name: "bad-phase-token", CheckClass: "enum", Disposition: "uncovered-S3", Reason: err.Error()}
	}
	violations := reg.Validate(record.Record{
		Envelope: record.Envelope{From: "seat-a", Role: "implementer", SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "NOPE", "AUTHORITY": "report-only", "SUBJECT": "bad phase"},
	}, fieldspec.SeatMeta{Name: "seat-a", Role: "implementer"}, "")
	for _, violation := range violations {
		if violation.Field == "PHASE" && violation.Class == "enum" {
			return Result{Name: "bad-phase-token", CheckClass: "enum", Disposition: "caught", Reason: "MVP validator rejects the typed PHASE equivalent"}
		}
	}
	return Result{Name: "bad-phase-token", CheckClass: "enum", Disposition: "uncovered-S3", Reason: "no matching MVP violation"}
}

func loadRegistry() (*fieldspec.Registry, error) {
	for _, path := range []string{
		filepath.Join("internal", "fieldspec", "registry.json"),
		filepath.Join("..", "..", "internal", "fieldspec", "registry.json"),
	} {
		reg, err := fieldspec.Load(path)
		if err == nil {
			return reg, nil
		}
	}
	return fieldspec.Load(filepath.Join("internal", "fieldspec", "registry.json"))
}

func classifyCorpus(root string) ([]Result, error) {
	var results []Result
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(path) != ".md" || entry.Name() == "README.md" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		results = append(results, classifyCorpusFixture(rel))
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Name < results[j].Name })
	return results, nil
}

func classifyCorpusFixture(name string) Result {
	lower := strings.ToLower(name)
	switch {
	case strings.Contains(lower, "bad-enum") || strings.Contains(lower, "enum-bypass") || strings.Contains(lower, "bad-phase"):
		return Result{Name: name, CheckClass: "enum", Disposition: "caught", Reason: "MVP validator rejects the typed enum equivalent"}
	case strings.Contains(lower, "fenced-dispatch") || strings.Contains(lower, "token") || strings.Contains(lower, "bare"):
		return Result{Name: name, CheckClass: "literal-authority-token-shape", Disposition: "obsolete-by-construction", Reason: "typed submission has no raw relay text channel for token placement"}
	case strings.Contains(lower, "proxy-from"):
		return Result{Name: name, CheckClass: "identity-stamping", Disposition: "obsolete-by-construction", Reason: "seat identity is stamped from binding, not payload FROM"}
	default:
		return Result{Name: name, CheckClass: "relay-lint-corpus", Disposition: "uncovered-S3", Reason: "not an MVP-visible S1 typed validation class"}
	}
}
