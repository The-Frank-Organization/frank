// Package outcome builds the payload-free v2 DATA-P replies and CTRL-C
// attempt results at each pipeline outcome's own terminal boundary.
package outcome

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackli/frank/internal/connector/authorize"
	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/jcs"
	"github.com/jackli/frank/internal/connector/request"
)

const (
	AttemptResultSchemaV2 = "m8.attempt_result.v2"
	DataReplySchemaV2     = "m8.dataP_reply.v2"
)

var ErrInvalidOutcome = errors.New("connector: invalid outcome carrier")

type Kind string

const (
	RejectedLocal   Kind = "rejected_local"
	Denied          Kind = "denied"
	SentCompleted   Kind = "sent_completed"
	TransportFailed Kind = "transport_failed"
	Unknown         Kind = "unknown"
	Cancelled       Kind = "cancelled"
)

type RefusalStage string

const (
	PreFreeze  RefusalStage = "pre_freeze"
	PostFreeze RefusalStage = "post_freeze"
)

type RejectCut string

const (
	RejectPreFreeze  RejectCut = "reject_pre_freeze"
	RejectPostFreeze RejectCut = "reject_post_freeze"
)

type CancelPoint string

const (
	PreTransport   CancelPoint = "pre_transport"
	PostInvocation CancelPoint = "post_invocation"
)

type Digests struct {
	FrozenCore           string
	ProviderLoweredTools string
}

type Outcome struct {
	Kind         Kind
	RejectReason request.RejectReason
	DenyReason   authorize.DenyReason
	RefusalStage RefusalStage
	CancelPoint  CancelPoint
	Digests      *Digests
}

type AttemptResultV2 struct {
	Schema                     string        `json:"schema"`
	AttemptID                  string        `json:"attempt_id"`
	TurnEpoch                  frame.Counter `json:"turn_epoch"`
	Disposition                string        `json:"disposition"`
	RefusalStage               RefusalStage  `json:"refusal_stage,omitempty"`
	FrozenCoreDigest           string        `json:"frozen_core_digest,omitempty"`
	ProviderLoweredToolsDigest string        `json:"provider_lowered_tools_digest,omitempty"`
}

type DataReplyV2 struct {
	Schema                     string               `json:"schema"`
	Kind                       string               `json:"type"`
	AttemptID                  string               `json:"attempt_id"`
	TurnEpoch                  frame.Counter        `json:"turn_epoch"`
	RejectReason               request.RejectReason `json:"reject_reason,omitempty"`
	DenyReason                 authorize.DenyReason `json:"deny_reason,omitempty"`
	RefusalStage               RefusalStage         `json:"refusal_stage,omitempty"`
	FrozenCoreDigest           string               `json:"frozen_core_digest,omitempty"`
	ProviderLoweredToolsDigest string               `json:"provider_lowered_tools_digest,omitempty"`
}

type LocalRejectWriter interface {
	SendAttemptResult(AttemptResultV2) error
	SendDataReply(DataReplyV2) error
}

func AttemptResult(attemptID string, epoch frame.Counter, outcome Outcome) (AttemptResultV2, error) {
	if attemptID == "" {
		return AttemptResultV2{}, invalid("empty attempt id")
	}
	if err := validate(outcome); err != nil {
		return AttemptResultV2{}, err
	}
	result := AttemptResultV2{Schema: AttemptResultSchemaV2, AttemptID: attemptID, TurnEpoch: epoch}
	switch outcome.Kind {
	case RejectedLocal:
		result.Disposition = fmt.Sprintf("rejected_local(%s)", outcome.RejectReason)
		result.RefusalStage = outcome.RefusalStage
	case Denied:
		result.Disposition = fmt.Sprintf("denied(%s)", outcome.DenyReason)
	case SentCompleted:
		result.Disposition = string(SentCompleted)
	case TransportFailed:
		result.Disposition = string(TransportFailed)
	case Unknown:
		result.Disposition = string(Unknown)
	case Cancelled:
		result.Disposition = fmt.Sprintf("cancelled(%s)", outcome.CancelPoint)
	}
	applyDigests(&result.FrozenCoreDigest, &result.ProviderLoweredToolsDigest, outcome.Digests)
	return result, nil
}

func DataReply(attemptID string, epoch frame.Counter, outcome Outcome) (DataReplyV2, error) {
	if attemptID == "" {
		return DataReplyV2{}, invalid("empty attempt id")
	}
	if err := validate(outcome); err != nil {
		return DataReplyV2{}, err
	}
	reply := DataReplyV2{Schema: DataReplySchemaV2, AttemptID: attemptID, TurnEpoch: epoch}
	switch outcome.Kind {
	case RejectedLocal:
		reply.Kind = "rejected_local"
		if outcome.RejectReason == request.InternalIntegrityFault {
			reply.Kind = "internal_integrity_reject"
		}
		reply.RejectReason = outcome.RejectReason
		reply.RefusalStage = outcome.RefusalStage
	case Denied:
		reply.Kind = "egress_denied"
		reply.DenyReason = outcome.DenyReason
	default:
		return DataReplyV2{}, invalid("outcome %s has no typed DATA-P reply", outcome.Kind)
	}
	applyDigests(&reply.FrozenCoreDigest, &reply.ProviderLoweredToolsDigest, outcome.Digests)
	return reply, nil
}

func EpochDataReply(attemptID string, epoch frame.Counter, disposition string) (DataReplyV2, *AttemptResultV2) {
	if disposition != "STALE_EPOCH" && disposition != "EPOCH_AHEAD" {
		disposition = "EPOCH_AHEAD"
	}
	return DataReplyV2{Schema: DataReplySchemaV2, Kind: disposition, AttemptID: attemptID, TurnEpoch: epoch}, nil
}

func ClassifyReject(reason request.RejectReason, stage RefusalStage) (RejectCut, error) {
	switch stage {
	case PreFreeze:
		switch reason {
		case request.MalformedRequest, request.LaneCapabilityMismatch, request.ReplayScopeViolation, request.InternalIntegrityFault:
			return RejectPreFreeze, nil
		}
	case PostFreeze:
		if reason == request.InternalIntegrityFault {
			return RejectPostFreeze, nil
		}
	}
	return "", invalid("invalid reject reason/stage pair")
}

func EmitLocalReject(writer LocalRejectWriter, attemptID string, epoch frame.Counter, outcome Outcome) error {
	if writer == nil || outcome.Kind != RejectedLocal {
		return invalid("local reject writer or outcome")
	}
	result, err := AttemptResult(attemptID, epoch, outcome)
	if err != nil {
		return err
	}
	reply, err := DataReply(attemptID, epoch, outcome)
	if err != nil {
		return err
	}
	if err := writer.SendAttemptResult(result); err != nil {
		return err
	}
	return writer.SendDataReply(reply)
}

func (result AttemptResultV2) CanonicalBytes() []byte {
	raw, _ := json.Marshal(result)
	canonical, _ := jcs.Canonicalize(raw)
	return canonical
}

func validate(outcome Outcome) error {
	switch outcome.Kind {
	case RejectedLocal:
		cut, err := ClassifyReject(outcome.RejectReason, outcome.RefusalStage)
		if err != nil {
			return err
		}
		if cut == RejectPostFreeze {
			if !validDigests(outcome.Digests) {
				return invalid("post-freeze reject lacks both digests")
			}
		} else if outcome.Digests != nil {
			return invalid("pre-freeze reject carries digests")
		}
		if outcome.DenyReason != "" || outcome.CancelPoint != "" {
			return invalid("reject carries foreign outcome fields")
		}
	case Denied:
		if !validDenyReason(outcome.DenyReason) || !validDigests(outcome.Digests) {
			return invalid("denial token or digests")
		}
	case SentCompleted, TransportFailed, Unknown:
		if !validDigests(outcome.Digests) {
			return invalid("post-freeze terminal lacks both digests")
		}
	case Cancelled:
		if (outcome.CancelPoint != PreTransport && outcome.CancelPoint != PostInvocation) || !validDigests(outcome.Digests) {
			return invalid("cancellation point or digests")
		}
	default:
		return invalid("unknown outcome kind")
	}
	return nil
}

func validDenyReason(reason authorize.DenyReason) bool {
	switch reason {
	case authorize.PolicyUnavailable, authorize.PolicyDigestMismatch, authorize.MalformedCore,
		authorize.LaneMismatch, authorize.LaneEndpointInvalid, authorize.EndpointMismatch,
		authorize.EndpointNotAllowlisted, authorize.MethodMismatch, authorize.ReservedAuthHeaderInCore:
		return true
	default:
		return false
	}
}

func validDigests(digests *Digests) bool {
	return digests != nil && validDigest(digests.FrozenCore) && validDigest(digests.ProviderLoweredTools)
}

func validDigest(value string) bool {
	if len(value) != sha256.Size*2 {
		return false
	}
	decoded, err := hex.DecodeString(value)
	return err == nil && hex.EncodeToString(decoded) == value
}

func applyDigests(frozenCore, loweredTools *string, digests *Digests) {
	if digests == nil {
		return
	}
	*frozenCore = digests.FrozenCore
	*loweredTools = digests.ProviderLoweredTools
}

func invalid(format string, args ...any) error {
	return fmt.Errorf("%w: %s", ErrInvalidOutcome, fmt.Sprintf(format, args...))
}
