package fixtures_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/bounce"
	"github.com/The-Frank-Organization/frank/internal/egress"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/migrate"
	"github.com/The-Frank-Organization/frank/test/replay/zeroloss"
)

func TestS5IPHNewSurfacesAreFieldClassOnlyAndPathClean(t *testing.T) {
	secretFinding := egress.Finding{Field: "body", Class: egress.ClassSafety}
	if got := secretFinding.String(); got != "body:safety" || strings.Contains(got, "secret") {
		t.Fatalf("finding string = %q, want Field:Class only", got)
	}

	report := egress.Report{Results: []egress.Result{{
		ItemID:          "gate-1",
		SourceRecordRef: "source-1",
		Dest:            "operator",
		Disposition:     egress.DispositionBlocked,
		Findings:        []egress.Finding{secretFinding},
	}}}
	reportBytes, err := json.Marshal(report)
	if err != nil {
		t.Fatalf("marshal report: %v", err)
	}
	if strings.Contains(string(reportBytes), "secret-value") || strings.Contains(string(reportBytes), "gpt-5.5") {
		t.Fatalf("egress report leaked scanned value: %s", reportBytes)
	}

	systemOwned := bounce.Format(fieldspec.Violation{
		Field:  "delivery_state",
		Class:  "system-owned",
		Reason: "/tmp/frank-store/records/leak.json",
	})
	if systemOwned != "delivery_state:system-owned" {
		t.Fatalf("formatter output changed: %q", systemOwned)
	}

	replayStatus := fmt.Sprintf("zeroloss:read=%d lost=%d quarantined=%d", 12, 0, 0)
	refusals := []string{
		migrate.ErrFutureVersion.Error(),
		migrate.ErrUnversioned.Error(),
		fmt.Errorf("%w: 1->2", migrate.ErrMigrationGap).Error(),
	}
	drainSentinels := []string{
		"outbox-read-error",
		"outbox-item-read-error",
		"outbox-item-decode-error",
		"source-record-read-error",
		"egress-render-error",
		"outbox-item-id-empty",
	}
	reportStatus := fmt.Sprintf("zeroloss:read=%d lost=%d quarantined=%d", zeroloss.Report{Read: 1}.Read, zeroloss.Report{}.Lost, zeroloss.Report{}.Quarantined)

	outputs := []string{
		secretFinding.String(),
		string(reportBytes),
		systemOwned,
		replayStatus,
		reportStatus,
	}
	outputs = append(outputs, refusals...)
	outputs = append(outputs, drainSentinels...)
	assertNoPathFamilies(t, outputs...)
	assertNoS5IPHForbiddenTokens(t, outputs...)
}

func TestS5IPHFormatterValveUntouched(t *testing.T) {
	got := bounce.Format(
		fieldspec.Violation{Field: "body", Class: "safety", Reason: "secret-value /outbox/item.json"},
		fieldspec.Violation{Field: "delivery_state", Class: "system-owned", Reason: "internal path"},
	)
	want := "body:safety; delivery_state:system-owned"
	if got != want {
		t.Fatalf("bounce.Format = %q, want %q", got, want)
	}
}

func assertNoS5IPHForbiddenTokens(t *testing.T, outputs ...string) {
	t.Helper()
	for _, output := range outputs {
		for _, forbidden := range []string{
			"secret-value",
			"gpt-5.5",
			"registry.json",
			"engine.json",
			"StoreRoot",
			"config/",
			"/tmp/frank-store",
		} {
			if strings.Contains(output, forbidden) {
				t.Fatalf("S5 I-PH output leaked %q in %q", forbidden, output)
			}
		}
	}
}
