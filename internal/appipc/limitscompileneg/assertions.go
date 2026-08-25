//go:build frank_limits_compile_negative

package limitscompileneg

const (
	FrameMax                 = 4_096
	AdmissionRefEncMax       = 1_024
	ManifestMax              = 1_024
	ParkedMax                = 512
	PathMaxM10               = 256
	OverheadMax              = 256
	AttemptAckMembersMax     = 3_073
	MaxParkedRowsPerRun      = 4
	ParkedRowMax             = 256
	TurnOpenBoundAssertion   = FrameMax - (AdmissionRefEncMax + ManifestMax + ParkedMax + PathMaxM10 + OverheadMax)
	AttemptAckBoundAssertion = FrameMax - (AttemptAckMembersMax + MaxParkedRowsPerRun*ParkedRowMax)
)

var (
	_ [TurnOpenBoundAssertion]byte
	_ [AttemptAckBoundAssertion]byte
)
