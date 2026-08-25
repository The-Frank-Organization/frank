// Package turn owns the worker's one-active-turn state and bounded lifecycle.
package turn

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type State string

const (
	StateIdle       State = "IDLE"
	StateAdmitted   State = "ADMITTED"
	StateAssembling State = "ASSEMBLING"
	StateAttempting State = "ATTEMPTING"
	StateObserving  State = "OBSERVING"
	StateToolRound  State = "TOOL_ROUND"
	StateTerminal   State = "TERMINAL"
)

type Terminal string

const (
	TurnCompleted Terminal = "turn_completed"
	TurnRefused   Terminal = "turn_refused"
	TurnDenied    Terminal = "turn_denied"
	TurnCancelled Terminal = "turn_cancelled"
	TurnFailed    Terminal = "turn_failed"
	TurnExhausted Terminal = "turn_exhausted"
)

const (
	MaxProviderAttempts = 16
	MaxToolCalls        = 64
	MaxWallClock        = 30 * time.Minute
	MaxStreamedOutput   = 8 << 20
	MaxCapturedOutput   = 1 << 20
	DoomLoopThreshold   = 4
)

var (
	ErrTurnActive    = errors.New("turn: another logical turn is active")
	ErrStaleEpoch    = errors.New("turn: stale epoch")
	ErrProtocol      = errors.New("turn: protocol fault")
	ErrNotActive     = errors.New("turn: no active turn")
	ErrReassemblyDue = errors.New("turn: parked-unknown disclosure changed; reassembly required")
)

type AdmissionRef struct {
	Kind      string `json:"kind"`
	RelayID   string `json:"relay_id,omitempty"`
	TaskInput string `json:"task_input,omitempty"`
}

func (reference AdmissionRef) Validate() error {
	switch reference.Kind {
	case "wake_relay":
		if reference.RelayID == "" || reference.TaskInput != "" {
			return ErrProtocol
		}
	case "operator_input":
		if reference.TaskInput == "" || reference.RelayID != "" {
			return ErrProtocol
		}
	default:
		return ErrProtocol
	}
	return nil
}

type ParkedUnknown struct {
	TurnID              string `json:"turn_id"`
	ToolCallID          string `json:"tool_call_id"`
	TicketID            string `json:"ticket_id"`
	State               string `json:"state"`
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
}

func (item ParkedUnknown) validate() error {
	if item.TurnID == "" || item.ToolCallID == "" || item.TicketID == "" || item.CanonicalToolName == "" || item.CanonicalArgsDigest == "" {
		return ErrProtocol
	}
	if item.State != "UNKNOWN_TOOL_OUTCOME" && item.State != "PARTIAL_TOOL_EFFECT" {
		return ErrProtocol
	}
	return nil
}

type Open struct {
	RunID         string
	TurnID        string
	TurnEpoch     string
	AdmissionRef  AdmissionRef
	ParkedUnknown []ParkedUnknown
}

type Machine struct {
	mu sync.Mutex

	state       State
	runID       string
	turnID      string
	epoch       uint64
	started     time.Time
	attempts    int
	toolCalls   int
	repeated    int
	lastCallKey string
	gateOne     []ParkedUnknown
	terminal    Terminal
	cancelAcked bool
}

func New() *Machine { return &Machine{state: StateIdle} }

func (machine *Machine) Admit(open Open, now time.Time) error {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	if machine.state != StateIdle && machine.state != StateTerminal {
		return ErrTurnActive
	}
	epoch, err := parseCounter(open.TurnEpoch)
	if err != nil || open.RunID == "" || open.TurnID == "" || open.AdmissionRef.Validate() != nil || validateParked(open.ParkedUnknown) != nil {
		return ErrProtocol
	}
	machine.state = StateAdmitted
	machine.runID, machine.turnID, machine.epoch = open.RunID, open.TurnID, epoch
	machine.started = now
	machine.attempts, machine.toolCalls, machine.repeated = 0, 0, 0
	machine.lastCallKey, machine.terminal, machine.cancelAcked = "", "", false
	machine.gateOne = append([]ParkedUnknown(nil), open.ParkedUnknown...)
	return nil
}

func (machine *Machine) BeginAssembly(epoch uint64) error {
	return machine.advance(epoch, StateAdmitted, StateAssembling)
}

func (machine *Machine) AttemptOpenOK(epoch uint64, parked []ParkedUnknown) error {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	if err := machine.check(epoch, StateAssembling); err != nil {
		return err
	}
	relation, err := CompareParked(machine.gateOne, parked)
	if err != nil {
		return err
	}
	if relation == ParkedAddedOrChanged {
		machine.gateOne = append([]ParkedUnknown(nil), parked...)
		return ErrReassemblyDue
	}
	machine.attempts++
	if machine.attempts > MaxProviderAttempts || time.Since(machine.started) > MaxWallClock {
		machine.state, machine.terminal = StateTerminal, TurnExhausted
		return nil
	}
	machine.state = StateAttempting
	return nil
}

func (machine *Machine) Observe(epoch uint64) error {
	return machine.advance(epoch, StateAttempting, StateObserving)
}

func (machine *Machine) ToolRound(epoch uint64, callKey string, denied bool) error {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	if err := machine.check(epoch, StateObserving); err != nil {
		return err
	}
	machine.toolCalls++
	if callKey != "" && callKey == machine.lastCallKey {
		machine.repeated++
	} else {
		machine.lastCallKey, machine.repeated = callKey, 1
	}
	if machine.toolCalls > MaxToolCalls || machine.repeated >= DoomLoopThreshold {
		machine.state, machine.terminal = StateTerminal, TurnExhausted
		return nil
	}
	_ = denied // denials count identically and continue until the bound trips.
	machine.state = StateToolRound
	return nil
}

func (machine *Machine) Reassemble(epoch uint64) error {
	return machine.advance(epoch, StateToolRound, StateAssembling)
}

func (machine *Machine) Finish(epoch uint64, terminal Terminal) error {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	if epoch != machine.epoch {
		return ErrStaleEpoch
	}
	if machine.state == StateIdle || machine.state == StateTerminal || !validTerminal(terminal) {
		return ErrProtocol
	}
	if terminal == TurnCancelled && !machine.cancelAcked {
		return ErrProtocol
	}
	machine.state, machine.terminal = StateTerminal, terminal
	return nil
}

func (machine *Machine) CancelAck(epoch uint64) error {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	if epoch != machine.epoch {
		return ErrStaleEpoch
	}
	if machine.state == StateIdle || machine.state == StateTerminal || machine.cancelAcked {
		return ErrProtocol
	}
	machine.cancelAcked = true
	return nil
}

func (machine *Machine) Snapshot() (State, Terminal, int, int) {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	return machine.state, machine.terminal, machine.attempts, machine.toolCalls
}

func (machine *Machine) advance(epoch uint64, from, to State) error {
	machine.mu.Lock()
	defer machine.mu.Unlock()
	if err := machine.check(epoch, from); err != nil {
		return err
	}
	machine.state = to
	return nil
}

func (machine *Machine) check(epoch uint64, state State) error {
	if epoch != machine.epoch {
		return ErrStaleEpoch
	}
	if machine.state != state {
		return fmt.Errorf("%w: state %s, want %s", ErrProtocol, machine.state, state)
	}
	return nil
}

type ParkedRelation string

const (
	ParkedEqual          ParkedRelation = "equal"
	ParkedAddedOrChanged ParkedRelation = "added-or-changed"
	ParkedRemovedOnly    ParkedRelation = "removed-only"
)

func CompareParked(gateOne, gateTwo []ParkedUnknown) (ParkedRelation, error) {
	left, err := parkedMap(gateOne)
	if err != nil {
		return "", err
	}
	right, err := parkedMap(gateTwo)
	if err != nil {
		return "", err
	}
	for key, second := range right {
		first, present := left[key]
		if !present || first != second {
			return ParkedAddedOrChanged, nil
		}
	}
	if len(right) < len(left) {
		return ParkedRemovedOnly, nil
	}
	return ParkedEqual, nil
}

func parkedMap(items []ParkedUnknown) (map[string]ParkedUnknown, error) {
	result := make(map[string]ParkedUnknown, len(items))
	for _, item := range items {
		if err := item.validate(); err != nil {
			return nil, err
		}
		key := item.TurnID + "\x00" + item.ToolCallID + "\x00" + item.TicketID
		if _, duplicate := result[key]; duplicate {
			return nil, ErrProtocol
		}
		result[key] = item
	}
	return result, nil
}

func validateParked(items []ParkedUnknown) error { _, err := parkedMap(items); return err }

func parseCounter(encoded string) (uint64, error) {
	if encoded == "" || (len(encoded) > 1 && encoded[0] == '0') {
		return 0, ErrProtocol
	}
	return strconv.ParseUint(encoded, 10, 64)
}

func validTerminal(terminal Terminal) bool {
	switch terminal {
	case TurnCompleted, TurnRefused, TurnDenied, TurnCancelled, TurnFailed, TurnExhausted:
		return true
	default:
		return false
	}
}
