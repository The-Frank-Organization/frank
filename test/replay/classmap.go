package replay

import (
	"path/filepath"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

type Result struct {
	Name        string
	CheckClass  string
	Disposition string
	Reason      string
}

func ClassifyAll() []Result {
	return []Result{
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
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
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
