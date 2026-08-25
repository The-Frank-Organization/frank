package appipc

import "fmt"

const (
	LeaseUnleased = "unleased"
	LeaseLeased   = "leased"

	ProposalInstalled                = "installed"
	ProposalTransitionStarted        = "transition-started"
	ProposalRejectedStale            = "rejected-stale"
	ProposalRejectedTransitionActive = "rejected-transition-active"
	ProposalRejectedMalformed        = "rejected-malformed"
)

type StateProposalBody struct {
	ProposalCorrelation string `json:"proposal_correlation"`
	RunID               string `json:"run_id"`
	GenerationID        string `json:"generation_id"`
	TurnEpoch           string `json:"turn_epoch"`
	LeaseState          string `json:"lease_state"`
	StateSeq            string `json:"state_seq"`
}

type EpochStateBody struct {
	RunID        string `json:"run_id"`
	GenerationID string `json:"generation_id"`
	TurnEpoch    string `json:"turn_epoch"`
	LeaseState   string `json:"lease_state"`
	StateSeq     string `json:"state_seq"`
}

type StateProposalResultBody struct {
	ProposalCorrelation string          `json:"proposal_correlation"`
	Disposition         string          `json:"disposition"`
	InstalledState      *EpochStateBody `json:"installed_state,omitempty"`
}

type BoundaryCutBody struct {
	EpochTransitionID string `json:"epoch_transition_id"`
	OperationID       string `json:"operation_id"`
	GenerationID      string `json:"generation_id"`
	AdmittedEpoch     string `json:"admitted_epoch"`
	Op                string `json:"op"`
	Disposition       string `json:"disposition"`
}

type EpochInstalledBody struct {
	EpochTransitionID string `json:"epoch_transition_id"`
	GenerationID      string `json:"generation_id"`
	TurnEpoch         string `json:"turn_epoch"`
	StateSeq          string `json:"state_seq"`
}

type BrokerEventAckBody struct {
	BrokerInstanceNonce string `json:"broker_instance_nonce"`
	EventSeq            string `json:"event_seq"`
}

func registerBroker(registry *Registry) error {
	closed := ClosedFamily
	leaseEnum := map[string][]string{"lease_state": {LeaseUnleased, LeaseLeased}}
	registrations := []func() error{
		func() error {
			return registerBody[StateProposalBody](registry, ChannelBroker, "state_proposal", false, closed, leaseEnum, validateStateProposal)
		},
		func() error {
			return registerBody[StateProposalResultBody](registry, ChannelBroker, "state_proposal_result", true, closed, map[string][]string{"disposition": {ProposalInstalled, ProposalTransitionStarted, ProposalRejectedStale, ProposalRejectedTransitionActive, ProposalRejectedMalformed}}, validateProposalResult)
		},
		func() error {
			return registerBody[EpochStateBody](registry, ChannelBroker, "epoch_state", false, closed, leaseEnum, validateEpochState)
		},
		func() error {
			return registerBody[BoundaryCutBody](registry, ChannelBroker, "boundary_cut", false, closed, map[string][]string{"disposition": {"unknown-outcome", "stale-cut"}}, validateBoundaryCut)
		},
		func() error {
			return registerBody[EpochInstalledBody](registry, ChannelBroker, "epoch_installed", false, closed, nil, validateEpochInstalled)
		},
		func() error {
			return registerBody[BrokerEventAckBody](registry, ChannelBroker, "broker_event_ack", true, closed, nil, validateBrokerAck)
		},
	}
	for _, register := range registrations {
		if err := register(); err != nil {
			return err
		}
	}
	return nil
}

func validateStateProposal(body *StateProposalBody) error {
	if err := requiredStrings(body.ProposalCorrelation, body.RunID, body.GenerationID); err != nil {
		return err
	}
	return validateCounterFields(body.TurnEpoch, body.StateSeq)
}

func validateEpochState(body *EpochStateBody) error {
	if err := requiredStrings(body.RunID, body.GenerationID); err != nil {
		return err
	}
	return validateCounterFields(body.TurnEpoch, body.StateSeq)
}

func validateProposalResult(body *StateProposalResultBody) error {
	if body.Disposition == ProposalInstalled {
		if body.InstalledState == nil {
			return fmt.Errorf("installed requires installed_state")
		}
		return validateEpochState(body.InstalledState)
	}
	if body.InstalledState != nil {
		return fmt.Errorf("non-installed disposition forbids installed_state")
	}
	return nil
}

func validateBoundaryCut(body *BoundaryCutBody) error {
	if err := requiredStrings(body.EpochTransitionID, body.OperationID, body.GenerationID, body.Op); err != nil {
		return err
	}
	return validateCounterFields(body.AdmittedEpoch)
}

func validateEpochInstalled(body *EpochInstalledBody) error {
	if err := requiredStrings(body.EpochTransitionID, body.GenerationID); err != nil {
		return err
	}
	return validateCounterFields(body.TurnEpoch, body.StateSeq)
}

func validateBrokerAck(body *BrokerEventAckBody) error {
	if err := requiredStrings(body.BrokerInstanceNonce); err != nil {
		return err
	}
	return validateCounterFields(body.EventSeq)
}
