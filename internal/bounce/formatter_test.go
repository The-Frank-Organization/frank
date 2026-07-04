package bounce_test

import (
	"strings"
	"testing"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/lineage"
)

func TestFormatContainsNoStorePaths(t *testing.T) {
	text := bounce.Format(
		fieldspec.Violation{Field: "PHASE", Class: "enum", Reason: "/tmp/store/records leaked"},
		lineage.Bounce{Edge: "PARENT_DISPATCH_ID", Kind: lineage.ParentUnknownRecompose},
	)
	if strings.Contains(text, "/tmp/store") || strings.Contains(text, "records/") {
		t.Fatalf("formatter leaked path text: %q", text)
	}
	if !strings.Contains(text, "PHASE") || !strings.Contains(text, "parent-unknown") {
		t.Fatalf("formatter lost field/edge context: %q", text)
	}
}
