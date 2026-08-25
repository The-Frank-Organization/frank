//go:build !frank_test_reduced_limits

package appipc

const (
	FrameMax             = 4_194_304
	AdmissionRefEncMax   = 2_564_096
	ManifestMax          = 1_232_896
	ChainDepthMax        = 10
	EntryMax             = 768
	IDMax                = 64
	ParkedMax            = 327_680
	MaxParkedRowsPerRun  = 512
	MaxToolCallsPerTurn  = 128
	ParkedRowMax         = 640
	PathMaxM10           = 4_096
	OverheadMax          = 65_536
	AttemptAckMembersMax = 1_024

	turnOpenBoundAssertion   = FrameMax - (AdmissionRefEncMax + ManifestMax + ParkedMax + PathMaxM10 + OverheadMax)
	attemptAckBoundAssertion = FrameMax - (AttemptAckMembersMax + MaxParkedRowsPerRun*ParkedRowMax)
)

var (
	_ [turnOpenBoundAssertion]byte
	_ [attemptAckBoundAssertion]byte
)
