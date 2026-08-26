package broker

import (
	"testing"

	"github.com/jackli/frank/internal/appctl/brokerclient"
	"github.com/jackli/frank/internal/appipc"
)

func TestProposalCorrelationBoundaryAndOrderedTable(t *testing.T) {
	missing := DecodeProposal([]byte(`{"run_id":"run","generation_id":"g1","turn_epoch":"1","lease_state":"unleased","state_seq":"1"}`))
	invalid := DecodeProposal([]byte(`{"proposal_correlation":"bad\ncorrelation","run_id":"run","generation_id":"g1","turn_epoch":"1","lease_state":"unleased","state_seq":"1"}`))
	unknown := DecodeProposal([]byte(`{"proposal_correlation":"c1","run_id":"run","generation_id":"g1","turn_epoch":"1","lease_state":"unleased","state_seq":"1","foreign":true}`))
	if missing.Kind != ProposalFrameFault || invalid.Kind != ProposalFrameFault || unknown.Kind != ProposalCorrelatedMalformed || unknown.Correlation != "c1" {
		t.Fatalf("correlation staging missing=%#v invalid=%#v unknown=%#v", missing, invalid, unknown)
	}

	engine := NewProposalEngine()
	if got := engine.Propose(missing); got != nil {
		t.Fatalf("uncorrelated malformed proposal emitted a result: %#v", got)
	}
	malformed := engine.Propose(unknown)
	if malformed == nil || malformed.Disposition != appipc.ProposalRejectedMalformed || malformed.ProposalCorrelation != "c1" {
		t.Fatalf("correlated malformed = %#v", malformed)
	}
	bootstrap := DecodeProposal(proposalJSON("bootstrap", tuple("run", "g1", "1", "1")))
	installed := engine.Propose(bootstrap)
	if installed == nil || installed.Disposition != appipc.ProposalInstalled || installed.InstalledState == nil || *installed.InstalledState != bootstrap.Tuple {
		t.Fatalf("bootstrap = %#v", installed)
	}
	transition := DecodeProposal(proposalJSON("transition", tuple("run", "g2", "2", "2")))
	started := engine.Propose(transition)
	joined := engine.Propose(DecodeProposal(proposalJSON("join", transition.Tuple)))
	conflict := engine.Propose(DecodeProposal(proposalJSON("conflict", tuple("run", "g3", "3", "3"))))
	staleDuring := engine.Propose(DecodeProposal(proposalJSON("stale-during", tuple("run", "old", "0", "0"))))
	if started.Disposition != appipc.ProposalTransitionStarted || joined.Disposition != appipc.ProposalTransitionStarted || conflict.Disposition != appipc.ProposalRejectedTransitionActive || staleDuring.Disposition != appipc.ProposalRejectedTransitionActive {
		t.Fatalf("PREPARING table started=%#v joined=%#v conflict=%#v stale=%#v", started, joined, conflict, staleDuring)
	}
	if !engine.CompleteTransition() {
		t.Fatal("active transition did not install")
	}
	equal := engine.Propose(DecodeProposal(proposalJSON("equal", transition.Tuple)))
	sameEpochNewer := engine.Propose(DecodeProposal(proposalJSON("washout", tuple("run", "g3", "2", "3"))))
	stale := engine.Propose(DecodeProposal(proposalJSON("stale", tuple("run", "g1", "1", "1"))))
	if equal.Disposition != appipc.ProposalInstalled || sameEpochNewer.Disposition != appipc.ProposalInstalled || stale.Disposition != appipc.ProposalRejectedStale {
		t.Fatalf("installed table equal=%#v newer=%#v stale=%#v", equal, sameEpochNewer, stale)
	}
	faulted := NewProposalEngine()
	outOfTable := tuple("run", "generation", "1", "1")
	outOfTable.LeaseState = "out-of-table"
	faulted.installed = &outOfTable
	if got := faulted.Propose(bootstrap); got == nil || got.Disposition != appipc.ProposalRejectedMalformed {
		t.Fatalf("out-of-table state did not fail closed: %#v", got)
	}
}

func TestTwoFormAssignGateUsesCanonicalTuple(t *testing.T) {
	trusted := tuple("run", "g2", "2", "4")
	responseGate := brokerclient.NewAssignGate()
	if err := responseGate.Propose("response", trusted); err != nil {
		t.Fatal(err)
	}
	result := responseGate.Fold("response", appipc.StateProposalResultBody{ProposalCorrelation: "response", Disposition: appipc.ProposalInstalled, InstalledState: &trusted})
	if result.Action != brokerclient.OpenAssign {
		t.Fatalf("response proof = %#v", result)
	}
	eventGate := brokerclient.NewAssignGate()
	if err := eventGate.Propose("event", trusted); err != nil {
		t.Fatal(err)
	}
	event := appipc.EpochInstalledBody{EpochTransitionID: "telemetry", GenerationID: trusted.GenerationID, TurnEpoch: trusted.TurnEpoch, StateSeq: trusted.StateSeq}
	if !eventGate.InstallEvent(trusted.RunID, event) {
		t.Fatal("event proof did not open assign")
	}
	wrong := event
	wrong.StateSeq = "5"
	negative := brokerclient.NewAssignGate()
	_ = negative.Propose("wrong", trusted)
	if negative.InstallEvent(trusted.RunID, wrong) {
		t.Fatal("wrong trusted tuple opened assign")
	}
}

func tuple(runID, generationID, epoch, stateSeq string) appipc.EpochStateBody {
	return appipc.EpochStateBody{RunID: runID, GenerationID: generationID, TurnEpoch: epoch, LeaseState: appipc.LeaseUnleased, StateSeq: stateSeq}
}

func proposalJSON(correlation string, value appipc.EpochStateBody) []byte {
	return []byte(`{"proposal_correlation":"` + correlation + `","run_id":"` + value.RunID + `","generation_id":"` + value.GenerationID + `","turn_epoch":"` + value.TurnEpoch + `","lease_state":"` + value.LeaseState + `","state_seq":"` + value.StateSeq + `"}`)
}
