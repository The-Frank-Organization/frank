// Package brokerclient owns the app-side CI-1 broker control protocol.
package brokerclient

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type FoldAction string

const (
	OpenAssign           FoldAction = "open_assign"
	AwaitEventOrDeadline FoldAction = "await_event_or_deadline"
	Repropose            FoldAction = "repropose"
	InvariantFault       FoldAction = "invariant_fault"
	Discard              FoldAction = "discard"
)

type FoldResult struct {
	Action     FoldAction
	Loud       bool
	Idempotent bool
}

type proposal struct {
	correlation string
	tuple       appipc.EpochStateBody
	opened      bool
	result      *appipc.StateProposalResultBody
}

type AssignGate struct {
	mu      sync.Mutex
	byRun   map[string]*proposal
	byCorr  map[string]string
	retired map[string]appipc.StateProposalResultBody
}

func NewAssignGate() *AssignGate {
	return &AssignGate{byRun: make(map[string]*proposal), byCorr: make(map[string]string), retired: make(map[string]appipc.StateProposalResultBody)}
}

func (gate *AssignGate) Propose(correlation string, tuple appipc.EpochStateBody) error {
	if gate == nil || correlation == "" || validateTuple(tuple) != nil {
		return fmt.Errorf("brokerclient: invalid proposal")
	}
	gate.mu.Lock()
	defer gate.mu.Unlock()
	if _, exists := gate.byRun[tuple.RunID]; exists {
		return fmt.Errorf("brokerclient: proposal already outstanding for run")
	}
	gate.byRun[tuple.RunID] = &proposal{correlation: correlation, tuple: tuple}
	gate.byCorr[correlation] = tuple.RunID
	return nil
}

func (gate *AssignGate) Fold(correlation string, body appipc.StateProposalResultBody) FoldResult {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	runID, ok := gate.byCorr[correlation]
	if !ok {
		if previous, duplicate := gate.retired[correlation]; duplicate && reflect.DeepEqual(previous, body) {
			return FoldResult{Action: Discard, Idempotent: true}
		}
		return FoldResult{Action: Discard}
	}
	pending := gate.byRun[runID]
	if body.ProposalCorrelation != correlation {
		return FoldResult{Action: InvariantFault, Loud: true}
	}
	pending.result = &body
	switch body.Disposition {
	case appipc.ProposalInstalled:
		if body.InstalledState == nil || !sameTuple(pending.tuple, *body.InstalledState) {
			return FoldResult{Action: InvariantFault, Loud: true}
		}
		pending.opened = true
		gate.retire(correlation, body, false)
		return FoldResult{Action: OpenAssign}
	case appipc.ProposalTransitionStarted, appipc.ProposalRejectedTransitionActive:
		return FoldResult{Action: AwaitEventOrDeadline}
	case appipc.ProposalRejectedStale, appipc.ProposalRejectedMalformed:
		gate.retire(correlation, body, true)
		return FoldResult{Action: InvariantFault, Loud: true}
	default:
		gate.retire(correlation, body, true)
		return FoldResult{Action: InvariantFault, Loud: true}
	}
}

func (gate *AssignGate) Deadline(runID string) FoldResult {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	pending, ok := gate.byRun[runID]
	if !ok || pending.opened {
		return FoldResult{Action: Discard}
	}
	delete(gate.byCorr, pending.correlation)
	delete(gate.byRun, runID)
	return FoldResult{Action: Repropose}
}

func (gate *AssignGate) InstallEvent(runID string, event appipc.EpochInstalledBody) bool {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	pending, ok := gate.byRun[runID]
	if !ok || pending.opened {
		return false
	}
	if pending.tuple.GenerationID != event.GenerationID || pending.tuple.TurnEpoch != event.TurnEpoch || pending.tuple.StateSeq != event.StateSeq {
		return false
	}
	pending.opened = true
	delete(gate.byCorr, pending.correlation)
	return true
}

func (gate *AssignGate) Open(tuple appipc.EpochStateBody) bool {
	gate.mu.Lock()
	defer gate.mu.Unlock()
	pending := gate.byRun[tuple.RunID]
	return pending != nil && pending.opened && sameTuple(pending.tuple, tuple)
}

func (gate *AssignGate) retire(correlation string, body appipc.StateProposalResultBody, remove bool) {
	runID := gate.byCorr[correlation]
	gate.retired[correlation] = body
	delete(gate.byCorr, correlation)
	if remove {
		delete(gate.byRun, runID)
	}
}

func sameTuple(left, right appipc.EpochStateBody) bool {
	return left == right
}

func validateTuple(tuple appipc.EpochStateBody) error {
	if tuple.RunID == "" || tuple.GenerationID == "" || (tuple.LeaseState != appipc.LeaseLeased && tuple.LeaseState != appipc.LeaseUnleased) {
		return fmt.Errorf("brokerclient: invalid tuple")
	}
	if _, err := appipc.ParseCounter(tuple.TurnEpoch); err != nil {
		return err
	}
	_, err := appipc.ParseCounter(tuple.StateSeq)
	return err
}
