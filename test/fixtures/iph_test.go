package fixtures_test

import (
	"strings"
	"testing"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/fieldspec"
)

func TestP1NoPathFamiliesInSeatDeliverableStrings(t *testing.T) {
	outputs := []string{
		bounce.Format(fieldspec.Violation{Field: "PHASE", Class: "enum", Reason: "/store/records/leak"}),
		`{"state":"accepted"}`,
		`["submit","project","read"]`,
	}
	for _, output := range outputs {
		for _, family := range []string{"/records", "/staging", "/outbox", "/binding", "operator-socket"} {
			if strings.Contains(output, family) {
				t.Fatalf("seat-deliverable output leaked %s in %q", family, output)
			}
		}
	}
}
