package broker

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/jackli/frank/internal/appipc"
)

type ProposalKind uint8

const (
	ProposalFrameFault ProposalKind = iota
	ProposalCorrelatedMalformed
	ProposalValid
)

type DecodedProposal struct {
	Kind        ProposalKind
	Correlation string
	Tuple       appipc.EpochStateBody
}

func DecodeProposal(raw []byte) DecodedProposal {
	var staged map[string]json.RawMessage
	if json.Unmarshal(raw, &staged) != nil {
		return DecodedProposal{Kind: ProposalFrameFault}
	}
	var correlation string
	encodedCorrelation, exists := staged["proposal_correlation"]
	if !exists || json.Unmarshal(encodedCorrelation, &correlation) != nil || !validCorrelation(correlation) {
		return DecodedProposal{Kind: ProposalFrameFault}
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var body appipc.StateProposalBody
	if decoder.Decode(&body) != nil || requireEOF(decoder) != nil || !validProposalBody(body) {
		return DecodedProposal{Kind: ProposalCorrelatedMalformed, Correlation: correlation}
	}
	return DecodedProposal{Kind: ProposalValid, Correlation: correlation, Tuple: appipc.EpochStateBody{
		RunID: body.RunID, GenerationID: body.GenerationID, TurnEpoch: body.TurnEpoch, LeaseState: body.LeaseState, StateSeq: body.StateSeq,
	}}
}

func validCorrelation(value string) bool {
	if value == "" || len(value) > 256 || !utf8.ValidString(value) {
		return false
	}
	for _, character := range value {
		if character < 0x20 || character == 0x7f {
			return false
		}
	}
	return true
}

func validProposalBody(body appipc.StateProposalBody) bool {
	if body.ProposalCorrelation == "" || body.RunID == "" || body.GenerationID == "" || !utf8.ValidString(body.RunID) || !utf8.ValidString(body.GenerationID) {
		return false
	}
	if body.LeaseState != appipc.LeaseLeased && body.LeaseState != appipc.LeaseUnleased {
		return false
	}
	if _, err := appipc.ParseCounter(body.TurnEpoch); err != nil {
		return false
	}
	_, err := appipc.ParseCounter(body.StateSeq)
	return err == nil
}

func requireEOF(decoder *json.Decoder) error {
	var extra any
	err := decoder.Decode(&extra)
	if err == io.EOF {
		return nil
	}
	if err == nil {
		return errors.New("broker: trailing proposal value")
	}
	return err
}

type ProposalEngine struct {
	mu        sync.Mutex
	installed *appipc.EpochStateBody
	active    *appipc.EpochStateBody
}

func NewProposalEngine() *ProposalEngine { return &ProposalEngine{} }

func (engine *ProposalEngine) Propose(proposal DecodedProposal) *appipc.StateProposalResultBody {
	if proposal.Kind == ProposalFrameFault {
		return nil
	}
	if proposal.Kind == ProposalCorrelatedMalformed {
		return result(proposal.Correlation, appipc.ProposalRejectedMalformed, nil)
	}
	engine.mu.Lock()
	defer engine.mu.Unlock()
	if !validTuple(proposal.Tuple) {
		return result(proposal.Correlation, appipc.ProposalRejectedMalformed, nil)
	}
	if engine.installed != nil && !validTuple(*engine.installed) || engine.active != nil && !validTuple(*engine.active) {
		return result(proposal.Correlation, appipc.ProposalRejectedMalformed, nil)
	}
	if engine.installed == nil && engine.active == nil {
		engine.installed = cloneTuple(proposal.Tuple)
		return result(proposal.Correlation, appipc.ProposalInstalled, engine.installed)
	}
	if engine.active != nil {
		if proposal.Tuple == *engine.active {
			return result(proposal.Correlation, appipc.ProposalTransitionStarted, nil)
		}
		return result(proposal.Correlation, appipc.ProposalRejectedTransitionActive, nil)
	}
	comparison := compareTuple(*engine.installed, proposal.Tuple)
	switch comparison {
	case tupleEqual:
		return result(proposal.Correlation, appipc.ProposalInstalled, engine.installed)
	case tupleSameEpochNewer:
		engine.installed = cloneTuple(proposal.Tuple)
		return result(proposal.Correlation, appipc.ProposalInstalled, engine.installed)
	case tupleNewerEpoch:
		engine.active = cloneTuple(proposal.Tuple)
		return result(proposal.Correlation, appipc.ProposalTransitionStarted, nil)
	default:
		return result(proposal.Correlation, appipc.ProposalRejectedStale, nil)
	}
}

func (engine *ProposalEngine) CompleteTransition() bool {
	engine.mu.Lock()
	defer engine.mu.Unlock()
	if engine.active == nil {
		return false
	}
	engine.installed = cloneTuple(*engine.active)
	engine.active = nil
	return true
}

type tupleOrder uint8

const (
	tupleStale tupleOrder = iota
	tupleEqual
	tupleSameEpochNewer
	tupleNewerEpoch
)

func compareTuple(installed, proposed appipc.EpochStateBody) tupleOrder {
	if installed == proposed {
		return tupleEqual
	}
	if installed.RunID != proposed.RunID {
		return tupleStale
	}
	installedEpoch, installedEpochErr := appipc.ParseCounter(installed.TurnEpoch)
	proposedEpoch, proposedEpochErr := appipc.ParseCounter(proposed.TurnEpoch)
	installedSeq, installedSeqErr := appipc.ParseCounter(installed.StateSeq)
	proposedSeq, proposedSeqErr := appipc.ParseCounter(proposed.StateSeq)
	if installedEpochErr != nil || proposedEpochErr != nil || installedSeqErr != nil || proposedSeqErr != nil {
		return tupleStale
	}
	if proposedEpoch > installedEpoch {
		return tupleNewerEpoch
	}
	if proposedEpoch == installedEpoch && proposedSeq > installedSeq {
		return tupleSameEpochNewer
	}
	return tupleStale
}

func validTuple(tuple appipc.EpochStateBody) bool {
	body := appipc.StateProposalBody{
		ProposalCorrelation: "validation", RunID: tuple.RunID, GenerationID: tuple.GenerationID,
		TurnEpoch: tuple.TurnEpoch, LeaseState: tuple.LeaseState, StateSeq: tuple.StateSeq,
	}
	return validProposalBody(body) && !strings.ContainsRune(tuple.RunID, '\x00') && !strings.ContainsRune(tuple.GenerationID, '\x00')
}

func cloneTuple(tuple appipc.EpochStateBody) *appipc.EpochStateBody {
	copy := tuple
	return &copy
}

func result(correlation, disposition string, installed *appipc.EpochStateBody) *appipc.StateProposalResultBody {
	return &appipc.StateProposalResultBody{ProposalCorrelation: correlation, Disposition: disposition, InstalledState: installed}
}
