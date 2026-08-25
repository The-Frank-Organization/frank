package appipc

import "fmt"

const (
	AttachOK            = "attach-ok"
	AttachSuspended     = "broker:attach-suspended"
	AttachTupleMismatch = "broker:attach-tuple-mismatch"

	AdmissionWakeRelay     = "wake_relay"
	AdmissionOperatorInput = "operator_input"

	RunDispositionFresh  = "fresh"
	RunDispositionResume = "resume"

	ParkedUnknownOutcome = "UNKNOWN_TOOL_OUTCOME"
	ParkedPartialEffect  = "PARTIAL_TOOL_EFFECT"

	OutcomeExecuted                 = "executed"
	OutcomeNotInvokedIntegrityFault = "not_invoked_integrity_fault"
)

type HelloBody struct {
	PID       int64             `json:"pid"`
	BuildInfo map[string]string `json:"build_info"`
}

type AssignBody struct {
	RunID                string `json:"run_id"`
	TurnEpoch            string `json:"turn_epoch"`
	ManifestDigest       string `json:"manifest_digest"`
	GenerationID         string `json:"generation_id"`
	BrokerWorkerEndpoint string `json:"broker_worker_endpoint"`
	WorkspaceRootPath    string `json:"workspace_root_path,omitempty"`
	WorkspaceRootID      string `json:"workspace_root_id,omitempty"`
}

type AttachResultBody struct {
	GenerationID string `json:"generation_id"`
	TurnEpoch    string `json:"turn_epoch"`
	Result       string `json:"result"`
}

type AdmissionRef struct {
	Kind      string  `json:"kind"`
	RelayID   *string `json:"relay_id,omitempty"`
	TaskInput *string `json:"task_input,omitempty"`
}

type ParkedUnknown struct {
	TurnID              string `json:"turn_id"`
	ToolCallID          string `json:"tool_call_id"`
	TicketID            string `json:"ticket_id"`
	State               string `json:"state"`
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
}

type SettlementEntry struct {
	Kind        string  `json:"kind"`
	Class       string  `json:"class"`
	RunID       string  `json:"run_id"`
	TurnID      string  `json:"turn_id"`
	AttemptID   *string `json:"attempt_id,omitempty"`
	ToolCallID  *string `json:"tool_call_id,omitempty"`
	ArgsDigest  *string `json:"args_digest,omitempty"`
	Terminal    string  `json:"terminal"`
	VoidReason  *string `json:"void_reason,omitempty"`
	CancelPoint *string `json:"cancel_point,omitempty"`
}

type SettlementManifest struct {
	Version           int               `json:"settlement_manifest_v"`
	RunID             string            `json:"run_id"`
	ProducedForTurnID string            `json:"produced_for_turn_id"`
	Entries           []SettlementEntry `json:"entries"`
}

type TurnOpenBody struct {
	TurnID             string              `json:"turn_id"`
	AdmissionRef       AdmissionRef        `json:"admission_ref"`
	ParkedUnknown      []ParkedUnknown     `json:"parked_unknown"`
	RunDisposition     string              `json:"run_disposition"`
	CreateAuthID       string              `json:"create_auth_id"`
	SessionLogPath     string              `json:"session_log_path"`
	SettlementManifest *SettlementManifest `json:"settlement_manifest,omitempty"`
	PredecessorTurnID  *string             `json:"predecessor_turn_id,omitempty"`
}

type GenesisCommittedBody struct {
	GenerationID string `json:"generation_id"`
}

type TurnTerminalBody struct {
	RunID     string `json:"run_id"`
	TurnID    string `json:"turn_id"`
	TurnEpoch string `json:"turn_epoch"`
	Terminal  string `json:"terminal"`
}

type TurnCancelAckBody struct {
	RunID              string `json:"run_id"`
	TurnID             string `json:"turn_id"`
	TurnEpoch          string `json:"turn_epoch"`
	PartialDisposition string `json:"partial_disposition"`
}

type TurnReceiptBody struct {
	RunID     string `json:"run_id"`
	TurnID    string `json:"turn_id"`
	TurnEpoch string `json:"turn_epoch"`
	Receipt   string `json:"receipt"`
}

type TurnRejectBody struct {
	RunID  string `json:"run_id"`
	TurnID string `json:"turn_id"`
	Reason string `json:"reason"`
}

type AttemptOpenBody struct {
	AttemptID            string `json:"attempt_id"`
	TurnID               string `json:"turn_id"`
	TurnEpoch            string `json:"turn_epoch"`
	ProviderLaneID       string `json:"provider_lane_id"`
	LogicalSurfaceDigest string `json:"logical_surface_digest"`
}

type AttemptOpenOKBody struct {
	AttemptID     string          `json:"attempt_id"`
	ParkedUnknown []ParkedUnknown `json:"parked_unknown"`
}

type AttemptOpenRejectBody struct {
	AttemptID string `json:"attempt_id"`
	Reason    string `json:"reason"`
}

type AttemptStreamEndBody struct {
	AttemptID   string `json:"attempt_id"`
	Disposition string `json:"disposition"`
}

type AppEventBody struct {
	EventID string         `json:"event_id"`
	Event   map[string]any `json:"event"`
}

type WakeForwardBody struct {
	RelayID string `json:"relay_id"`
}

type AuthorizeToolCallBody struct {
	RunID               string  `json:"run_id"`
	TurnID              string  `json:"turn_id"`
	TurnEpoch           string  `json:"turn_epoch"`
	ToolCallID          string  `json:"tool_call_id"`
	CanonicalToolName   string  `json:"canonical_tool_name"`
	CanonicalArgsDigest string  `json:"canonical_args_digest"`
	CanonicalResource   *string `json:"canonical_resource,omitempty"`
	CWD                 *string `json:"cwd,omitempty"`
}

type TicketGrantedBody struct {
	TicketID string `json:"ticket_id"`
}

type AuthorizeRejectBody struct {
	ToolCallID string `json:"tool_call_id"`
	Reason     string `json:"reason"`
}

type ToolCallRejectBody struct {
	ToolCallID string `json:"tool_call_id"`
}

type TypedRejectBody struct {
	ToolCallID *string `json:"tool_call_id,omitempty"`
	TicketID   *string `json:"ticket_id,omitempty"`
}

type ConsumeTicketBody struct {
	TicketID            string `json:"ticket_id"`
	TurnEpoch           string `json:"turn_epoch"`
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
}

type ConsumeOKBody struct {
	TicketID string `json:"ticket_id"`
}

type TicketRejectBody struct {
	TicketID string `json:"ticket_id"`
}

type InvocationIdentity struct {
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
	TurnEpoch           string `json:"turn_epoch"`
}

type IntegrityEvidence struct {
	ExpectedIdentity InvocationIdentity `json:"expected_identity"`
	ObservedIdentity InvocationIdentity `json:"observed_identity"`
}

type RecordToolOutcomeBody struct {
	TicketID           string              `json:"ticket_id"`
	TurnEpoch          string              `json:"turn_epoch"`
	Outcome            string              `json:"outcome"`
	InvocationIdentity *InvocationIdentity `json:"invocation_identity,omitempty"`
	IntegrityEvidence  *IntegrityEvidence  `json:"integrity_evidence,omitempty"`
}

type ContentReadyBody struct {
	TurnID        string `json:"turn_id"`
	AttemptID     string `json:"attempt_id"`
	RoundIdentity string `json:"round_identity"`
	SeqHWM        string `json:"seq_hwm"`
	GenerationID  string `json:"generation_id"`
}

type ResumeDispositionBody struct {
	TurnID       string  `json:"turn_id"`
	Disposition  string  `json:"disposition"`
	ResumeAction *string `json:"resume_action,omitempty"`
}

type AdmissionRefusedBody struct {
	Reason string `json:"reason"`
}

func registerCtrlW(registry *Registry) error {
	add := AdditiveFamily
	registrations := []func() error{
		func() error {
			return registerBody[HelloBody](registry, ChannelCtrlW, "hello", false, add, nil, validateHello)
		},
		func() error {
			return registerBody[AssignBody](registry, ChannelCtrlW, "assign", false, add, nil, validateAssign)
		},
		func() error {
			return registerBody[AttachResultBody](registry, ChannelCtrlW, "attach_result", false, add, map[string][]string{"result": {AttachOK, AttachSuspended, AttachTupleMismatch}}, validateAttachResult)
		},
		func() error {
			return registerBody[TurnOpenBody](registry, ChannelCtrlW, "turn_open", false, add, map[string][]string{"run_disposition": {RunDispositionFresh, RunDispositionResume}}, validateTurnOpen)
		},
		func() error {
			return registerBody[GenesisCommittedBody](registry, ChannelCtrlW, "genesis_committed", false, add, nil, validateGenesisCommitted)
		},
		func() error {
			return registerBody[TurnTerminalBody](registry, ChannelCtrlW, "turn_terminal", false, add, map[string][]string{"terminal": {"turn_completed", "turn_refused", "turn_denied", "turn_cancelled", "turn_failed", "turn_exhausted"}}, validateTurnTerminal)
		},
		func() error {
			return registerBody[TurnCancelAckBody](registry, ChannelCtrlW, "turn_cancel_ack", false, add, map[string][]string{"partial_disposition": {"none", "partials_committed_labeled"}}, validateTurnCancelAck)
		},
		func() error {
			return registerBody[TurnReceiptBody](registry, ChannelCtrlW, "turn_receipt", true, add, map[string][]string{"receipt": {"terminal_recorded", "cancel_recorded"}}, validateTurnReceipt)
		},
		func() error {
			return registerBody[TurnRejectBody](registry, ChannelCtrlW, "turn_reject", true, add, map[string][]string{"reason": {"stale_epoch", "unknown_turn", "conflicting_report"}}, nil)
		},
		func() error {
			return registerBody[AttemptOpenBody](registry, ChannelCtrlW, "attempt_open", false, add, nil, validateAttemptOpen)
		},
		func() error {
			return registerBody[AttemptOpenOKBody](registry, ChannelCtrlW, "attempt_open_ok", true, add, nil, validateAttemptOpenOK)
		},
		func() error {
			return registerBody[AttemptOpenRejectBody](registry, ChannelCtrlW, "attempt_open_reject", true, add, map[string][]string{"reason": {"stale_epoch", "invalid_turn", "invalid_lease"}}, nil)
		},
		func() error {
			return registerBody[AttemptStreamEndBody](registry, ChannelCtrlW, "attempt_stream_end", false, add, map[string][]string{"disposition": {"stream_completed", "stream_failed", "stream_cancelled", "stream_lost"}}, nil)
		},
		func() error {
			return registerBody[AppEventBody](registry, ChannelCtrlW, "app_event", false, add, nil, nil)
		},
		func() error {
			return registerBody[WakeForwardBody](registry, ChannelCtrlW, "wake_forward", false, add, nil, nil)
		},
		func() error {
			return registerBody[AuthorizeToolCallBody](registry, ChannelCtrlW, "authorize_tool_call", false, add, nil, validateAuthorize)
		},
		func() error {
			return registerBody[TicketGrantedBody](registry, ChannelCtrlW, "ticket_granted", true, add, nil, nil)
		},
		func() error {
			return registerBody[AuthorizeRejectBody](registry, ChannelCtrlW, "authorize_reject", true, add, map[string][]string{"reason": {"run_not_admitted", "turn_inactive", "lease_invalid", "turn_budget_exhausted"}}, nil)
		},
		func() error {
			return registerBody[ToolCallRejectBody](registry, ChannelCtrlW, "denied_above_set", true, add, nil, nil)
		},
		func() error {
			return registerBody[ToolCallRejectBody](registry, ChannelCtrlW, "duplicate_request", true, add, nil, nil)
		},
		func() error {
			return registerBody[TypedRejectBody](registry, ChannelCtrlW, "stale_epoch", true, add, nil, validateTypedReject)
		},
		func() error {
			return registerBody[TypedRejectBody](registry, ChannelCtrlW, "identity_mismatch", true, add, nil, validateTypedReject)
		},
		func() error {
			return registerBody[ConsumeTicketBody](registry, ChannelCtrlW, "consume_ticket", false, add, nil, validateConsumeTicket)
		},
		func() error {
			return registerBody[ConsumeOKBody](registry, ChannelCtrlW, "consume_ok", true, add, nil, nil)
		},
		func() error {
			return registerBody[TicketRejectBody](registry, ChannelCtrlW, "duplicate_consume", true, add, nil, nil)
		},
		func() error {
			return registerBody[RecordToolOutcomeBody](registry, ChannelCtrlW, "record_tool_outcome", false, add, map[string][]string{"outcome": {OutcomeExecuted, OutcomeNotInvokedIntegrityFault}}, validateRecordOutcome)
		},
		func() error {
			return registerBody[ContentReadyBody](registry, ChannelCtrlW, "content_ready", false, add, nil, validateContentReady)
		},
		func() error {
			return registerBody[ResumeDispositionBody](registry, ChannelCtrlW, "report_resume_disposition", false, add, map[string][]string{"disposition": {"resumable", "degraded"}}, validateResumeDisposition)
		},
		func() error {
			return registerBody[ResumeDispositionBody](registry, ChannelCtrlW, "receipt", true, add, map[string][]string{"disposition": {"resumable", "degraded"}}, validateResumeDisposition)
		},
		func() error {
			return registerBody[ResumeDispositionBody](registry, ChannelCtrlW, "disposition_conflict", true, add, map[string][]string{"disposition": {"resumable", "degraded"}}, validateResumeDisposition)
		},
		func() error { return registerBody[EmptyBody](registry, ChannelCtrlW, "ping", false, add, nil, nil) },
		func() error { return registerBody[EmptyBody](registry, ChannelCtrlW, "pong", true, add, nil, nil) },
		func() error { return registerBody[EmptyBody](registry, ChannelCtrlW, "shutdown", false, add, nil, nil) },
		func() error {
			return registerBody[AdmissionRefusedBody](registry, ChannelCtrlW, "admission_refused", false, add, map[string][]string{"reason": {"task_input_frame_overflow"}}, nil)
		},
	}
	for _, register := range registrations {
		if err := register(); err != nil {
			return err
		}
	}
	for _, messageType := range []string{"turn_open", "genesis_committed", "content_ready", "report_resume_disposition", "receipt", "disposition_conflict"} {
		if err := registry.requireEnvelope(ChannelCtrlW, messageType, true, true); err != nil {
			return err
		}
	}
	return nil
}

func validateHello(body *HelloBody) error {
	if body.PID <= 0 || body.BuildInfo == nil {
		return fmt.Errorf("pid must be positive and build_info present")
	}
	return nil
}

func validateAssign(body *AssignBody) error {
	if err := requiredStrings(body.RunID, body.GenerationID, body.BrokerWorkerEndpoint); err != nil {
		return err
	}
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	return validateDigest(body.ManifestDigest)
}

func validateAttachResult(body *AttachResultBody) error { return validateCounterFields(body.TurnEpoch) }

func validateTurnOpen(body *TurnOpenBody) error {
	if err := requiredStrings(body.TurnID, body.SessionLogPath); err != nil {
		return err
	}
	if err := validateHex32(body.CreateAuthID); err != nil {
		return err
	}
	if err := validateAdmissionRef(body.AdmissionRef); err != nil {
		return err
	}
	if (body.SettlementManifest == nil) != (body.PredecessorTurnID == nil) {
		return fmt.Errorf("continuation members must appear together")
	}
	if body.SettlementManifest != nil {
		if err := validateSettlementManifest(body.SettlementManifest); err != nil {
			return err
		}
	}
	for _, parked := range body.ParkedUnknown {
		if parked.State != ParkedUnknownOutcome && parked.State != ParkedPartialEffect {
			return fmt.Errorf("unknown parked state %q", parked.State)
		}
		if err := validateDigest(parked.CanonicalArgsDigest); err != nil {
			return err
		}
	}
	return nil
}

func validateGenesisCommitted(body *GenesisCommittedBody) error {
	return requiredStrings(body.GenerationID)
}

func validateSettlementManifest(manifest *SettlementManifest) error {
	if manifest.Version != 1 || manifest.Entries == nil {
		return fmt.Errorf("settlement manifest version or entries invalid")
	}
	if err := requiredStrings(manifest.RunID, manifest.ProducedForTurnID); err != nil {
		return err
	}
	for _, entry := range manifest.Entries {
		if err := requiredStrings(entry.Kind, entry.Class, entry.RunID, entry.TurnID, entry.Terminal); err != nil {
			return err
		}
		if entry.Class != "settled_with_content" && entry.Class != "determinate_no_resume" && entry.Class != "uncertain" {
			return fmt.Errorf("unknown settlement class %q", entry.Class)
		}
		switch entry.Kind {
		case "tool":
			if entry.ToolCallID == nil || entry.ArgsDigest == nil || entry.AttemptID != nil || entry.CancelPoint != nil {
				return fmt.Errorf("tool settlement entry has invalid discriminated shape")
			}
			if err := validateDigest(*entry.ArgsDigest); err != nil {
				return err
			}
			if entry.Terminal == "VOID" {
				if entry.VoidReason == nil {
					return fmt.Errorf("orphan VOID entry requires void_reason")
				}
			} else if entry.VoidReason != nil {
				return fmt.Errorf("non-VOID tool entry forbids void_reason")
			}
		case "provider":
			if entry.AttemptID == nil || entry.ToolCallID != nil || entry.ArgsDigest != nil || entry.VoidReason != nil {
				return fmt.Errorf("provider settlement entry has invalid discriminated shape")
			}
			if entry.Terminal == "CANCELLED" {
				if entry.CancelPoint == nil || (*entry.CancelPoint != "pre_transport" && *entry.CancelPoint != "post_invocation") {
					return fmt.Errorf("cancelled provider entry requires closed cancel_point")
				}
			} else if entry.CancelPoint != nil {
				return fmt.Errorf("non-cancelled provider entry forbids cancel_point")
			}
		default:
			return fmt.Errorf("unknown settlement kind %q", entry.Kind)
		}
	}
	return nil
}

func validateAdmissionRef(ref AdmissionRef) error {
	switch ref.Kind {
	case AdmissionWakeRelay:
		if ref.RelayID == nil || *ref.RelayID == "" || ref.TaskInput != nil {
			return fmt.Errorf("wake_relay requires only relay_id")
		}
	case AdmissionOperatorInput:
		if ref.TaskInput == nil || ref.RelayID != nil {
			return fmt.Errorf("operator_input requires only task_input")
		}
	default:
		return fmt.Errorf("unknown admission_ref kind %q", ref.Kind)
	}
	return nil
}

func validateTurnTerminal(body *TurnTerminalBody) error { return validateCounterFields(body.TurnEpoch) }
func validateTurnCancelAck(body *TurnCancelAckBody) error {
	return validateCounterFields(body.TurnEpoch)
}
func validateTurnReceipt(body *TurnReceiptBody) error { return validateCounterFields(body.TurnEpoch) }

func validateAttemptOpen(body *AttemptOpenBody) error {
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	return validateDigest(body.LogicalSurfaceDigest)
}

func validateAttemptOpenOK(body *AttemptOpenOKBody) error {
	if body.ParkedUnknown == nil {
		return fmt.Errorf("parked_unknown must be present, empty array when none")
	}
	return nil
}

func validateAuthorize(body *AuthorizeToolCallBody) error {
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	return validateDigest(body.CanonicalArgsDigest)
}

func validateConsumeTicket(body *ConsumeTicketBody) error {
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	return validateDigest(body.CanonicalArgsDigest)
}

func validateTypedReject(body *TypedRejectBody) error {
	if (body.ToolCallID == nil) == (body.TicketID == nil) {
		return fmt.Errorf("typed rejection requires exactly one request identity")
	}
	if body.ToolCallID != nil && *body.ToolCallID == "" {
		return fmt.Errorf("tool_call_id is empty")
	}
	if body.TicketID != nil && *body.TicketID == "" {
		return fmt.Errorf("ticket_id is empty")
	}
	return nil
}

func validateIdentity(identity InvocationIdentity) error {
	if err := requiredStrings(identity.CanonicalToolName); err != nil {
		return err
	}
	if err := validateDigest(identity.CanonicalArgsDigest); err != nil {
		return err
	}
	return validateCounterFields(identity.TurnEpoch)
}

func validateRecordOutcome(body *RecordToolOutcomeBody) error {
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	switch body.Outcome {
	case OutcomeExecuted:
		if body.InvocationIdentity == nil || body.IntegrityEvidence != nil {
			return fmt.Errorf("executed requires only invocation_identity")
		}
		return validateIdentity(*body.InvocationIdentity)
	case OutcomeNotInvokedIntegrityFault:
		if body.InvocationIdentity != nil || body.IntegrityEvidence == nil {
			return fmt.Errorf("not_invoked_integrity_fault requires only integrity_evidence")
		}
		if err := validateIdentity(body.IntegrityEvidence.ExpectedIdentity); err != nil {
			return err
		}
		if err := validateIdentity(body.IntegrityEvidence.ObservedIdentity); err != nil {
			return err
		}
		if body.IntegrityEvidence.ExpectedIdentity == body.IntegrityEvidence.ObservedIdentity {
			return fmt.Errorf("integrity evidence must describe a mismatch")
		}
	}
	return nil
}

func validateContentReady(body *ContentReadyBody) error {
	if err := validateDigest(body.RoundIdentity); err != nil {
		return err
	}
	return validateCounterFields(body.SeqHWM)
}

func validateResumeDisposition(body *ResumeDispositionBody) error {
	if body.Disposition == "degraded" {
		if body.ResumeAction == nil || *body.ResumeAction != "re_derive" {
			return fmt.Errorf("degraded requires resume_action re_derive")
		}
	} else if body.ResumeAction != nil {
		return fmt.Errorf("resumable forbids resume_action")
	}
	return nil
}
