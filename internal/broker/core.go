package broker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"unicode/utf8"

	"github.com/jackli/frank/internal/appipc"
	"github.com/jackli/frank/internal/channel"
)

var (
	ErrSuspended          = errors.New("broker:suspended")
	ErrStaleEpoch         = errors.New("broker:stale-epoch")
	ErrDuplicateOperation = errors.New("broker:duplicate-operation")
	ErrRecordUnavailable  = errors.New("broker:record-unavailable")
	ErrInvalidOperation   = errors.New("broker:invalid-operation")
)

type Conductor interface {
	Relay(context.Context, string, []byte) ([]byte, error)
	Describe(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error)
}

type OutcomeRecorder interface {
	RecordRelayOutcome(context.Context, RelayOutcome) error
}

type Capability struct {
	RunID, GenerationID, TurnEpoch, BrokerInstanceNonce string
}

type Operation struct {
	ID        string
	Name      string
	Arguments []byte
}

type RelayOutcome struct {
	OperationID string
	Name        string
	TurnEpoch   string
	Succeeded   bool
}

type Core struct {
	caller    Conductor
	recorder  OutcomeRecorder
	mu        sync.Mutex
	installed *appipc.EpochStateBody
	control   bool
	preparing bool
	seen      map[string]struct{}
}

func NewCore(caller Conductor, recorder OutcomeRecorder) *Core {
	return &Core{caller: caller, recorder: recorder, seen: make(map[string]struct{})}
}

func (core *Core) Install(tuple appipc.EpochStateBody) {
	if core == nil || !validTuple(tuple) {
		return
	}
	core.mu.Lock()
	defer core.mu.Unlock()
	core.installed = cloneTuple(tuple)
	core.control = true
	core.preparing = false
}

func (core *Core) Suspend() {
	if core == nil {
		return
	}
	core.mu.Lock()
	defer core.mu.Unlock()
	core.control = false
}

func (core *Core) SetPreparing(preparing bool) {
	if core == nil {
		return
	}
	core.mu.Lock()
	defer core.mu.Unlock()
	core.preparing = preparing
}

func (core *Core) Invoke(ctx context.Context, capability Capability, operation Operation) (any, error) {
	if core == nil || core.caller == nil || operation.ID == "" || !validOperation(operation.Name) || !validCapability(capability) {
		return nil, ErrInvalidOperation
	}
	core.mu.Lock()
	if !core.control || core.preparing || core.installed == nil {
		core.mu.Unlock()
		return nil, ErrSuspended
	}
	installed := *core.installed
	if capability.RunID != installed.RunID || capability.GenerationID != installed.GenerationID || capability.TurnEpoch != installed.TurnEpoch {
		core.mu.Unlock()
		return nil, ErrStaleEpoch
	}
	if _, duplicate := core.seen[operation.ID]; duplicate {
		core.mu.Unlock()
		return nil, ErrDuplicateOperation
	}
	core.seen[operation.ID] = struct{}{}
	core.mu.Unlock()
	if operation.Name == "describe" {
		request, err := decodeDescribe(operation.Arguments)
		if err != nil {
			return nil, err
		}
		return core.caller.Describe(ctx, request)
	}
	response, callErr := core.caller.Relay(ctx, operation.Name, append([]byte(nil), operation.Arguments...))
	if core.recorder == nil || core.recorder.RecordRelayOutcome(ctx, RelayOutcome{OperationID: operation.ID, Name: operation.Name, TurnEpoch: capability.TurnEpoch, Succeeded: callErr == nil}) != nil {
		return nil, ErrRecordUnavailable
	}
	return response, callErr
}

func validOperation(name string) bool {
	switch name {
	case "relay.submit", "relay.project", "relay.read", "describe":
		return true
	default:
		return false
	}
}

func validCapability(capability Capability) bool {
	if capability.RunID == "" || capability.GenerationID == "" || !utf8.ValidString(capability.RunID) || !utf8.ValidString(capability.GenerationID) {
		return false
	}
	if _, err := appipc.ParseCounter(capability.TurnEpoch); err != nil {
		return false
	}
	if len(capability.BrokerInstanceNonce) != 64 {
		return false
	}
	for _, character := range capability.BrokerInstanceNonce {
		if character < '0' || character > '9' && character < 'a' || character > 'f' {
			return false
		}
	}
	return true
}

func decodeDescribe(arguments []byte) (channel.DescribeRequest, error) {
	decoder := json.NewDecoder(bytes.NewReader(arguments))
	decoder.DisallowUnknownFields()
	var request channel.DescribeRequest
	if err := decoder.Decode(&request); err != nil || requireEOF(decoder) != nil {
		return channel.DescribeRequest{}, fmt.Errorf("%w: describe arguments", ErrInvalidOperation)
	}
	return request, nil
}
