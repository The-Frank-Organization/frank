//go:build frank_test_reduced_limits

package appipc

import _ "embed"

const (
	FrameMax             = 4_096
	AdmissionRefEncMax   = 8_192
	ManifestMax          = 2_048
	ChainDepthMax        = 4
	EntryMax             = 512
	IDMax                = 64
	ParkedMax            = 1_024
	MaxParkedRowsPerRun  = 4
	MaxToolCallsPerTurn  = 4
	ParkedRowMax         = 256
	PathMaxM10           = 512
	OverheadMax          = 1_024
	AttemptAckMembersMax = 512

	turnOpenOverflowAssertion = AdmissionRefEncMax + ManifestMax + ParkedMax + PathMaxM10 + OverheadMax - FrameMax - 1
	attemptAckBoundAssertion  = FrameMax - (AttemptAckMembersMax + MaxParkedRowsPerRun*ParkedRowMax)
)

var (
	_ [turnOpenOverflowAssertion]byte
	_ [attemptAckBoundAssertion]byte
)

//go:embed testdata/limits-table-testonly.json
var reducedLimitsArtifact []byte

func ReducedLimitsArtifact() []byte {
	artifact := reducedLimitsArtifact
	// Source files conventionally carry a terminal newline; the frozen
	// content-addressed artifact does not. Return the exact governed bytes.
	if len(artifact) != 0 && artifact[len(artifact)-1] == '\n' {
		artifact = artifact[:len(artifact)-1]
	}
	return append([]byte(nil), artifact...)
}
