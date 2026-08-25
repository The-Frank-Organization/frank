//go:build !frank_test_reduced_limits

package appipc

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestProductionLimitsAndAssertions(t *testing.T) {
	if FrameMax != 4_194_304 || AdmissionRefEncMax != 2_564_096 || ManifestMax != 1_232_896 {
		t.Fatalf("production primary limits = (%d, %d, %d)", FrameMax, AdmissionRefEncMax, ManifestMax)
	}
	if EntryMax != 768 || ParkedMax != 327_680 || MaxParkedRowsPerRun != 512 || ParkedRowMax != 640 {
		t.Fatalf("production nested limits = (%d, %d, %d, %d)", EntryMax, ParkedMax, MaxParkedRowsPerRun, ParkedRowMax)
	}
	if PathMaxM10 != 4_096 || OverheadMax != 65_536 || AttemptAckMembersMax != 1_024 || ChainDepthMax != 10 || IDMax != 64 {
		t.Fatalf("production remaining limits do not match the frozen table")
	}
	if got := AdmissionRefEncMax + ManifestMax + ParkedMax + PathMaxM10 + OverheadMax; got != FrameMax {
		t.Fatalf("P1 sum = %d, want FRAME_MAX %d", got, FrameMax)
	}
	if got := AttemptAckMembersMax + MaxParkedRowsPerRun*ParkedRowMax; got > FrameMax {
		t.Fatalf("P2 sum = %d, exceeds FRAME_MAX %d", got, FrameMax)
	}
	if FrameContentBound != 3_704_832 {
		t.Fatalf("FRAME_CONTENT_BOUND = %d", FrameContentBound)
	}
}

func TestTurnOpenSizeClassification(t *testing.T) {
	for _, test := range []struct {
		name         string
		encodedSize  int
		continuation bool
		want         TurnOpenSizeDisposition
	}{
		{name: "exact fit", encodedSize: FrameMax, continuation: true, want: TurnOpenSizeFits},
		{name: "initial overflow", encodedSize: FrameMax + 1, want: TurnOpenTaskInputOverflow},
		{name: "resume overflow", encodedSize: FrameMax + 1, continuation: true, want: TurnOpenResumeOverflow},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := ClassifyTurnOpenSize(test.encodedSize, test.continuation); got != test.want {
				t.Fatalf("ClassifyTurnOpenSize(%d, %t) = %q, want %q", test.encodedSize, test.continuation, got, test.want)
			}
		})
	}
}

func TestLimitsCompileMatrix(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller did not report the test file")
	}
	moduleRoot := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", ".."))

	for _, args := range [][]string{
		{"build", "./..."},
		{"build", "-tags", "frank_test_reduced_limits", "./internal/appipc"},
	} {
		command := exec.Command("go", args...)
		command.Dir = moduleRoot
		if output, err := command.CombinedOutput(); err != nil {
			t.Fatalf("go %s failed: %v\n%s", strings.Join(args, " "), err, output)
		}
	}

	command := exec.Command("go", "build", "-tags", "frank_limits_compile_negative", "./internal/appipc/limitscompileneg")
	command.Dir = moduleRoot
	output, err := command.CombinedOutput()
	if err == nil {
		t.Fatal("P2-violating compile-negative fixture unexpectedly built")
	}
	if !bytes.Contains(output, []byte("AttemptAckBoundAssertion")) {
		t.Fatalf("compile-negative stderr does not name the assertion declaration:\n%s", output)
	}
}
