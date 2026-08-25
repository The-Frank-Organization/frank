package appipc

import "fmt"

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

type AttemptResultBody struct {
	AttemptID                  string  `json:"attempt_id"`
	TurnEpoch                  string  `json:"turn_epoch"`
	Disposition                string  `json:"disposition"`
	DenyReason                 *string `json:"deny_reason,omitempty"`
	RejectReason               *string `json:"reject_reason,omitempty"`
	CancelPoint                *string `json:"cancel_point,omitempty"`
	FrozenCoreDigest           *string `json:"frozen_core_digest,omitempty"`
	ProviderLoweredToolsDigest *string `json:"provider_lowered_tools_digest,omitempty"`
}

func registerCtrlC(registry *Registry) error {
	add := AdditiveFamily
	registrations := []func() error{
		func() error {
			return registerBody[HelloBody](registry, ChannelCtrlC, "hello", false, add, nil, validateHello)
		},
		func() error {
			return registerBody[ConnectorAssignBody](registry, ChannelCtrlC, "connector_assign", false, add, nil, validateConnectorAssign)
		},
		func() error {
			return registerBody[ConnectorReadyBody](registry, ChannelCtrlC, "connector_ready", false, add, nil, validateConnectorReady)
		},
		func() error {
			return registerBody[EpochUpdateBody](registry, ChannelCtrlC, "epoch_update", false, add, nil, validateEpochUpdate)
		},
		func() error {
			return registerBody[AttemptResultBody](registry, ChannelCtrlC, "attempt_result", false, add, map[string][]string{"disposition": {"sent_completed", "denied", "transport_failed", "unknown", "rejected_local", "cancelled"}}, validateAttemptResult)
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
	switch body.Disposition {
	case "denied":
		if body.DenyReason == nil || body.RejectReason != nil || body.CancelPoint != nil {
			return fmt.Errorf("denied requires only deny_reason")
		}
	case "rejected_local":
		if body.RejectReason == nil || body.DenyReason != nil || body.CancelPoint != nil {
			return fmt.Errorf("rejected_local requires only reject_reason")
		}
	case "cancelled":
		if body.CancelPoint == nil || (*body.CancelPoint != "pre_transport" && *body.CancelPoint != "post_invocation") || body.DenyReason != nil || body.RejectReason != nil {
			return fmt.Errorf("cancelled requires only a closed cancel_point")
		}
	default:
		if body.DenyReason != nil || body.RejectReason != nil || body.CancelPoint != nil {
			return fmt.Errorf("plain attempt disposition forbids discriminant members")
		}
	}
	return nil
}
