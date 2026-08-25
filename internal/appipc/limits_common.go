package appipc

import "time"

const (
	FrameContentBound = 3_704_832

	SendDeadline           = 5 * time.Second
	HandshakeDeadline      = 10 * time.Second
	AttachDeadline         = 30 * time.Second
	HealthInterval         = 5 * time.Second
	BackoffBase            = time.Second
	ProposalResultDeadline = 35 * time.Second
)

type TurnOpenSizeDisposition string

const (
	TurnOpenSizeFits          TurnOpenSizeDisposition = ""
	TurnOpenTaskInputOverflow TurnOpenSizeDisposition = "task_input_frame_overflow"
	TurnOpenResumeOverflow    TurnOpenSizeDisposition = "resume_frame_overflow"
)

func ClassifyTurnOpenSize(encodedSize int, continuation bool) TurnOpenSizeDisposition {
	if encodedSize <= FrameMax {
		return TurnOpenSizeFits
	}
	if continuation {
		return TurnOpenResumeOverflow
	}
	return TurnOpenTaskInputOverflow
}
