package appipc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ConnectorBuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	BuiltAt string `json:"built_at"`
}

func (info *ConnectorBuildInfo) UnmarshalJSON(encoded []byte) error {
	type wire ConnectorBuildInfo
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.DisallowUnknownFields()
	var decoded wire
	if err := decoder.Decode(&decoded); err != nil {
		return err
	}
	if err := requireJSONEOF(decoder); err != nil {
		return err
	}
	*info = ConnectorBuildInfo(decoded)
	return nil
}

type ConnectorHelloBody struct {
	PID       int64              `json:"pid"`
	BuildInfo ConnectorBuildInfo `json:"build_info"`
}

type ConnectorAssignBody struct {
	RunID             string `json:"run_id"`
	TurnEpoch         string `json:"turn_epoch"`
	RunManifestDigest string `json:"run_manifest_digest"`
	PolicyDigest      string `json:"policy_digest"`
	ProviderLaneID    string `json:"provider_lane_id"`
	LaneCatalogDigest string `json:"lane_catalog_digest"`
	CredentialRef     string `json:"credential_ref"`
}

type ConnectorReadyBody struct {
	RunID     string `json:"run_id"`
	TurnEpoch string `json:"turn_epoch"`
}

type EpochUpdateBody struct {
	RunID     string `json:"run_id"`
	TurnEpoch string `json:"turn_epoch"`
}

type EpochQueryBody struct {
	RunID     string `json:"run_id"`
	TurnEpoch string `json:"turn_epoch"`
}

type AttemptResultBody struct {
	AttemptID                  string  `json:"attempt_id"`
	TurnEpoch                  string  `json:"turn_epoch"`
	Disposition                string  `json:"disposition"`
	DenyReason                 *string `json:"deny_reason,omitempty"`
	RejectReason               *string `json:"reject_reason,omitempty"`
	RefusalStage               *string `json:"refusal_stage,omitempty"`
	CancelPoint                *string `json:"cancel_point,omitempty"`
	FrozenCoreDigest           *string `json:"frozen_core_digest,omitempty"`
	ProviderLoweredToolsDigest *string `json:"provider_lowered_tools_digest,omitempty"`
}

func registerCtrlC(registry *Registry) error {
	add := AdditiveFamily
	registrations := []func() error{
		func() error {
			return registerBody[ConnectorHelloBody](registry, ChannelCtrlC, "hello", false, add, nil, validateConnectorHello)
		},
		func() error {
			return registerBody[ConnectorAssignBody](registry, ChannelCtrlC, "connector_assign", false, add, nil, validateConnectorAssign)
		},
		func() error {
			return registerBody[ConnectorReadyBody](registry, ChannelCtrlC, "connector_ready", false, add, nil, validateConnectorReady)
		},
		func() error {
			return registerBody[EpochQueryBody](registry, ChannelCtrlC, "epoch_query", false, add, nil, validateEpochQuery)
		},
		func() error {
			return registerBody[EpochUpdateBody](registry, ChannelCtrlC, "epoch_update", false, add, nil, validateEpochUpdate)
		},
		func() error {
			return registerBody[AttemptResultBody](registry, ChannelCtrlC, "attempt_result", false, add, nil, validateAttemptResult)
		},
		func() error { return registerBody[EmptyBody](registry, ChannelCtrlC, "ping", false, add, nil, nil) },
		func() error { return registerBody[EmptyBody](registry, ChannelCtrlC, "pong", true, add, nil, nil) },
		func() error { return registerBody[EmptyBody](registry, ChannelCtrlC, "shutdown", false, add, nil, nil) },
	}
	for _, register := range registrations {
		if err := register(); err != nil {
			return err
		}
	}
	for _, messageType := range []string{"epoch_query", "epoch_update"} {
		if err := registry.requireEnvelope(ChannelCtrlC, messageType, true, true); err != nil {
			return err
		}
	}
	return nil
}

func validateConnectorHello(body *ConnectorHelloBody) error {
	if body.PID <= 0 {
		return fmt.Errorf("pid must be positive")
	}
	for _, member := range []string{body.BuildInfo.Version, body.BuildInfo.Commit, body.BuildInfo.BuiltAt} {
		if member == "" || len(member) > 64 {
			return fmt.Errorf("build_info members must be 1..64 UTF-8 bytes")
		}
	}
	encoded, err := MarshalJCS(body.BuildInfo)
	if err != nil {
		return err
	}
	if len(encoded) > 256 {
		return fmt.Errorf("build_info exceeds 256 encoded bytes")
	}
	return nil
}

func validateConnectorAssign(body *ConnectorAssignBody) error {
	if err := requiredStrings(body.RunID, body.ProviderLaneID, body.CredentialRef); err != nil {
		return err
	}
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	for _, digest := range []string{body.RunManifestDigest, body.PolicyDigest, body.LaneCatalogDigest} {
		if err := validateDigest(digest); err != nil {
			return err
		}
	}
	return nil
}

func validateConnectorReady(body *ConnectorReadyBody) error {
	return validateCounterFields(body.TurnEpoch)
}

func validateEpochUpdate(body *EpochUpdateBody) error {
	return validateCounterFields(body.TurnEpoch)
}

func validateEpochQuery(body *EpochQueryBody) error {
	return validateCounterFields(body.TurnEpoch)
}

func validateAttemptResult(body *AttemptResultBody) error {
	if err := validateCounterFields(body.TurnEpoch); err != nil {
		return err
	}
	if body.FrozenCoreDigest != nil {
		if err := validateDigest(*body.FrozenCoreDigest); err != nil {
			return err
		}
	}
	if body.ProviderLoweredToolsDigest != nil {
		if err := validateDigest(*body.ProviderLoweredToolsDigest); err != nil {
			return err
		}
	}
	disposition, parameter, parameterized := splitDisposition(body.Disposition)
	switch disposition {
	case "denied":
		if body.DenyReason == nil || body.RejectReason != nil || body.CancelPoint != nil {
			return fmt.Errorf("denied requires only deny_reason")
		}
		if parameterized && parameter != *body.DenyReason {
			return fmt.Errorf("denied reason differs from disposition")
		}
	case "rejected_local":
		if body.RejectReason == nil || body.DenyReason != nil || body.CancelPoint != nil {
			return fmt.Errorf("rejected_local requires only reject_reason")
		}
		if body.RefusalStage == nil || (*body.RefusalStage != "pre_freeze" && *body.RefusalStage != "post_freeze") {
			return fmt.Errorf("rejected_local requires a closed refusal_stage")
		}
		if parameterized && parameter != *body.RejectReason {
			return fmt.Errorf("reject reason differs from disposition")
		}
	case "cancelled":
		if body.CancelPoint == nil || (*body.CancelPoint != "pre_transport" && *body.CancelPoint != "post_invocation") || body.DenyReason != nil || body.RejectReason != nil {
			return fmt.Errorf("cancelled requires only a closed cancel_point")
		}
		if parameterized && parameter != *body.CancelPoint {
			return fmt.Errorf("cancel point differs from disposition")
		}
	case "sent_completed", "transport_failed", "unknown":
		if parameterized {
			return fmt.Errorf("plain attempt disposition forbids a parameter")
		}
		if body.DenyReason != nil || body.RejectReason != nil || body.CancelPoint != nil {
			return fmt.Errorf("plain attempt disposition forbids discriminant members")
		}
	default:
		return fmt.Errorf("unknown disposition token %q", body.Disposition)
	}
	if disposition != "rejected_local" && body.RefusalStage != nil {
		return fmt.Errorf("refusal_stage is only valid for rejected_local")
	}
	return nil
}

func splitDisposition(value string) (string, string, bool) {
	for _, prefix := range []string{"denied", "rejected_local", "cancelled"} {
		opening := prefix + "("
		if len(value) > len(opening) && value[:len(opening)] == opening && value[len(value)-1] == ')' {
			return prefix, value[len(opening) : len(value)-1], true
		}
	}
	return value, "", false
}
