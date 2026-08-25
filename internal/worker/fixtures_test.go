package worker_test

import (
	"os"
	"strings"
	"testing"
)

func TestFixtureTraceabilityTableNamesEveryBuildFamilyAndOpenGate(t *testing.T) {
	contents, err := os.ReadFile("FIXTURES.md")
	if err != nil {
		t.Fatal(err)
	}
	text := string(contents)
	for _, required := range []string{
		"T1 JCS + counters", "T2 F58 catalog", "T3 frame codec", "T4 one-file journal",
		"T5 F59 executor", "T6 local tools", "T7 mapping + conductor facade", "T8 turn machine",
		"T9 provider cycle", "T10 context + E0", "T11 resume", "T12 governed turn",
		"SUPERSEDED", "PARTIAL", "HELD", "OWNER-EXTERNAL", "Branch A is operative",
		"master/exit-fixtures/**", "d4580c52", "RLBS-1",
	} {
		if !strings.Contains(text, required) {
			t.Fatalf("FIXTURES.md lacks required traceability token %q", required)
		}
	}
	for _, planted := range []string{
		"TestPreparedCallIsInertWithoutAuthorityPath",
		"TestDescriptorBatteryRejectsSymlinkModeHardlinkAndReplacement",
		"TestHandshakeFailuresFailClosedBeforeProviderOrToolWork",
		"TestTrustWindowViolationsNeverTrustContent",
		"TestProviderE0TotalTableAndNoEmissionCuts",
	} {
		if !strings.Contains(text, planted) {
			t.Fatalf("FIXTURES.md silently dropped anti-vacuity proof %q", planted)
		}
	}
}
