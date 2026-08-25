// Package request decodes and validates the closed app-internal
// m8.llm_request.v1 request shape before translation or freeze.
package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strconv"

	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/jcs"
)

type RejectReason string

const (
	MalformedRequest       RejectReason = "malformed_request"
	LaneCapabilityMismatch RejectReason = "lane_capability_mismatch"
	ReplayScopeViolation   RejectReason = "replay_scope_violation"
	InternalIntegrityFault RejectReason = "internal_integrity_fault"
)

func (reason RejectReason) Error() string { return string(reason) }

type Request struct {
	Schema         string
	RunID          string
	TurnID         string
	AttemptID      string
	TurnEpoch      frame.Counter
	ProviderLaneID string
	Instructions   string
	Input          []InputItem
	Tools          []Tool
	Sampling       Sampling
	Reasoning      Reasoning
}

type InputItem interface {
	inputItem()
}

type UserText struct{ Text string }
type AssistantText struct{ Text string }
type AssistantToolCall struct {
	ToolCallID string
	Name       string
	Arguments  string
}
type ToolResult struct {
	ToolCallID string
	Content    string
}
type ReasoningReplay struct{ Envelope ReplayEnvelope }

func (UserText) inputItem()          {}
func (AssistantText) inputItem()     {}
func (AssistantToolCall) inputItem() {}
func (ToolResult) inputItem()        {}
func (ReasoningReplay) inputItem()   {}

type ReplayEnvelope struct {
	OriginProviderLaneID string
	OriginTurnID         string
	Payload              string
}

type Tool struct {
	Name        string
	Description string
	Parameters  json.RawMessage
}

type Sampling struct {
	MaxOutputTokens int64
	Temperature     *json.Number
}

type Reasoning struct {
	Effort *string
}

type requestWire struct {
	Schema         *string           `json:"schema"`
	RunID          *string           `json:"run_id"`
	TurnID         *string           `json:"turn_id"`
	AttemptID      *string           `json:"attempt_id"`
	TurnEpoch      *frame.Counter    `json:"turn_epoch"`
	ProviderLaneID *string           `json:"provider_lane_id"`
	Instructions   *string           `json:"instructions"`
	Input          []json.RawMessage `json:"input"`
	Tools          []toolWire        `json:"tools"`
	Sampling       *samplingWire     `json:"sampling"`
	Reasoning      *reasoningWire    `json:"reasoning"`
}

type toolWire struct {
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	Parameters  json.RawMessage `json:"parameters"`
}

type samplingWire struct {
	MaxOutputTokens *int64       `json:"max_output_tokens"`
	Temperature     *json.Number `json:"temperature,omitempty"`
}

type reasoningWire struct {
	Effort *string `json:"effort,omitempty"`
}

// Parse validates shape first, replay provenance second, and lane capability
// last. This preserves the contract's pre-translate replay-scope boundary.
func Parse(raw []byte, lane catalog.Lane) (*Request, error) {
	if !jcs.IsCanonical(raw) {
		return nil, malformed("request bytes are not exact JCS")
	}
	var wire requestWire
	if err := decodeClosed(raw, &wire); err != nil {
		return nil, malformed("decode closed request: %v", err)
	}
	if wire.Schema == nil || *wire.Schema != "m8.llm_request.v1" ||
		wire.RunID == nil || *wire.RunID == "" || wire.TurnID == nil || *wire.TurnID == "" ||
		wire.AttemptID == nil || *wire.AttemptID == "" || wire.TurnEpoch == nil ||
		wire.ProviderLaneID == nil || *wire.ProviderLaneID == "" || wire.Instructions == nil ||
		wire.Input == nil || wire.Tools == nil || wire.Sampling == nil || wire.Reasoning == nil {
		return nil, malformed("missing or invalid required request member")
	}

	result := &Request{
		Schema:         *wire.Schema,
		RunID:          *wire.RunID,
		TurnID:         *wire.TurnID,
		AttemptID:      *wire.AttemptID,
		TurnEpoch:      *wire.TurnEpoch,
		ProviderLaneID: *wire.ProviderLaneID,
		Instructions:   *wire.Instructions,
		Input:          make([]InputItem, 0, len(wire.Input)),
		Tools:          make([]Tool, 0, len(wire.Tools)),
	}
	for index, itemRaw := range wire.Input {
		item, err := parseItem(itemRaw)
		if err != nil {
			return nil, malformed("input item %d: %v", index, err)
		}
		result.Input = append(result.Input, item)
	}
	for index, tool := range wire.Tools {
		parsed, err := parseTool(tool)
		if err != nil {
			return nil, malformed("tool %d: %v", index, err)
		}
		result.Tools = append(result.Tools, parsed)
	}
	if wire.Sampling.MaxOutputTokens == nil || *wire.Sampling.MaxOutputTokens <= 0 || !validOptionalNumber(wire.Sampling.Temperature) {
		return nil, malformed("invalid sampling")
	}
	result.Sampling = Sampling{MaxOutputTokens: *wire.Sampling.MaxOutputTokens, Temperature: wire.Sampling.Temperature}
	result.Reasoning = Reasoning{Effort: wire.Reasoning.Effort}

	for _, item := range result.Input {
		replay, ok := item.(ReasoningReplay)
		if !ok {
			continue
		}
		if replay.Envelope.OriginProviderLaneID != result.ProviderLaneID || replay.Envelope.OriginTurnID != result.TurnID {
			return nil, ReplayScopeViolation
		}
	}
	if result.Reasoning.Effort != nil && (!lane.Reasoning.Supported || !contains(lane.Reasoning.EffortLevels, *result.Reasoning.Effort)) {
		return nil, LaneCapabilityMismatch
	}
	return result, nil
}

func parseItem(raw json.RawMessage) (InputItem, error) {
	var discriminator struct {
		Kind *string `json:"kind"`
	}
	if err := json.Unmarshal(raw, &discriminator); err != nil || discriminator.Kind == nil {
		return nil, fmt.Errorf("missing kind")
	}
	switch *discriminator.Kind {
	case "user_text", "assistant_text":
		var wire struct {
			Kind *string `json:"kind"`
			Text *string `json:"text"`
		}
		if err := decodeClosed(raw, &wire); err != nil || wire.Text == nil {
			return nil, fmt.Errorf("invalid %s", *discriminator.Kind)
		}
		if *discriminator.Kind == "user_text" {
			return UserText{Text: *wire.Text}, nil
		}
		return AssistantText{Text: *wire.Text}, nil
	case "assistant_tool_call":
		var wire struct {
			Kind       *string `json:"kind"`
			ToolCallID *string `json:"tool_call_id"`
			Name       *string `json:"name"`
			Arguments  *string `json:"arguments"`
		}
		if err := decodeClosed(raw, &wire); err != nil || wire.ToolCallID == nil || *wire.ToolCallID == "" || wire.Name == nil || *wire.Name == "" || wire.Arguments == nil || !validJSONObjectString(*wire.Arguments) {
			return nil, fmt.Errorf("invalid assistant_tool_call")
		}
		return AssistantToolCall{ToolCallID: *wire.ToolCallID, Name: *wire.Name, Arguments: *wire.Arguments}, nil
	case "tool_result":
		var wire struct {
			Kind       *string `json:"kind"`
			ToolCallID *string `json:"tool_call_id"`
			Content    *string `json:"content"`
		}
		if err := decodeClosed(raw, &wire); err != nil || wire.ToolCallID == nil || *wire.ToolCallID == "" || wire.Content == nil {
			return nil, fmt.Errorf("invalid tool_result")
		}
		return ToolResult{ToolCallID: *wire.ToolCallID, Content: *wire.Content}, nil
	case "reasoning_replay":
		var wire struct {
			Kind     *string `json:"kind"`
			Envelope *struct {
				OriginProviderLaneID *string `json:"origin_provider_lane_id"`
				OriginTurnID         *string `json:"origin_turn_id"`
				Payload              *string `json:"payload"`
			} `json:"envelope"`
		}
		if err := decodeClosed(raw, &wire); err != nil || wire.Envelope == nil || wire.Envelope.OriginProviderLaneID == nil || *wire.Envelope.OriginProviderLaneID == "" || wire.Envelope.OriginTurnID == nil || *wire.Envelope.OriginTurnID == "" || wire.Envelope.Payload == nil || !validOpaquePayload(*wire.Envelope.Payload) {
			return nil, fmt.Errorf("invalid reasoning_replay")
		}
		return ReasoningReplay{Envelope: ReplayEnvelope{
			OriginProviderLaneID: *wire.Envelope.OriginProviderLaneID,
			OriginTurnID:         *wire.Envelope.OriginTurnID,
			Payload:              *wire.Envelope.Payload,
		}}, nil
	default:
		return nil, fmt.Errorf("unknown kind %q", *discriminator.Kind)
	}
}

func parseTool(wire toolWire) (Tool, error) {
	if wire.Name == nil || *wire.Name == "" || wire.Description == nil || len(wire.Parameters) == 0 {
		return Tool{}, fmt.Errorf("missing required member")
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(wire.Parameters, &object); err != nil || object == nil {
		return Tool{}, fmt.Errorf("parameters must be an object")
	}
	return Tool{Name: *wire.Name, Description: *wire.Description, Parameters: append(json.RawMessage(nil), wire.Parameters...)}, nil
}

func decodeClosed(raw []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	decoder.UseNumber()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return fmt.Errorf("trailing JSON value")
	}
	return nil
}

func validJSONObjectString(raw string) bool {
	var object map[string]json.RawMessage
	return json.Unmarshal([]byte(raw), &object) == nil && object != nil
}

func validOptionalNumber(number *json.Number) bool {
	if number == nil {
		return true
	}
	parsed, err := strconv.ParseFloat(number.String(), 64)
	return err == nil && !math.IsInf(parsed, 0) && !math.IsNaN(parsed)
}

func validOpaquePayload(payload string) bool {
	if payload == "" {
		return false
	}
	_, err := base64.StdEncoding.Strict().DecodeString(payload)
	return err == nil
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func malformed(format string, args ...any) error {
	return fmt.Errorf("%w: %s", MalformedRequest, fmt.Sprintf(format, args...))
}
