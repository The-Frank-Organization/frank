// Package stream parses the pinned provider SSE dialect into the uniform
// m8.provider_event.v2 event grammar.
package stream

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/jcs"
	"github.com/jackli/frank/internal/connector/transport"
)

const (
	SchemaV2           = "m8.provider_event.v2"
	ATTEMPT_STREAM_MAX = 64 * 1024 * 1024
)

type Kind string

const (
	AttemptStarted Kind = "attempt_started"
	TextStart      Kind = "text_start"
	TextDelta      Kind = "text_delta"
	TextEnd        Kind = "text_end"
	ReasoningStart Kind = "reasoning_start"
	ReasoningDelta Kind = "reasoning_delta"
	ReasoningEnd   Kind = "reasoning_end"
	ToolCallStart  Kind = "tool_call_start"
	ToolCallDelta  Kind = "tool_call_delta"
	ToolCallEnd    Kind = "tool_call_end"
	UsageProgress  Kind = "usage"
	Completed      Kind = "completed"
	Failed         Kind = "failed"
	Cancelled      Kind = "cancelled"
)

type Meta struct {
	AttemptID                  string
	ProviderLaneID             string
	TurnID                     string
	FrozenCoreDigest           string
	ProviderLoweredToolsDigest string
}

type Usage struct {
	InputTokens     uint64  `json:"input_tokens"`
	OutputTokens    uint64  `json:"output_tokens"`
	ReasoningTokens *uint64 `json:"reasoning_tokens,omitempty"`
	CacheReadTokens *uint64 `json:"cache_read_tokens,omitempty"`
}

type ReplayEnvelope struct {
	OriginProviderLaneID string `json:"origin_provider_lane_id"`
	OriginTurnID         string `json:"origin_turn_id"`
	Payload              []byte `json:"payload"`
}

type ToolCall struct {
	ToolCallID string          `json:"tool_call_id"`
	Name       string          `json:"name"`
	Arguments  json.RawMessage `json:"arguments"`
}

type Event struct {
	Schema                     string          `json:"schema"`
	Kind                       Kind            `json:"type"`
	AttemptID                  string          `json:"attempt_id"`
	BlockID                    string          `json:"block_id,omitempty"`
	Text                       string          `json:"text,omitempty"`
	ToolCall                   *ToolCall       `json:"tool_call,omitempty"`
	Usage                      *Usage          `json:"usage,omitempty"`
	FinishReason               string          `json:"finish_reason,omitempty"`
	ErrorClass                 string          `json:"error_class,omitempty"`
	Detail                     string          `json:"detail,omitempty"`
	Partial                    string          `json:"partial,omitempty"`
	ReplayEnvelope             *ReplayEnvelope `json:"replay_envelope,omitempty"`
	FrozenCoreDigest           string          `json:"frozen_core_digest,omitempty"`
	ProviderLoweredToolsDigest string          `json:"provider_lowered_tools_digest,omitempty"`
}

func (event Event) Terminal() bool {
	return event.Kind == Completed || event.Kind == Failed || event.Kind == Cancelled
}

type limits struct {
	frameMax   uint32
	attemptMax uint64
}

func Parse(reader io.Reader, meta Meta) []Event {
	return parseWithLimits(reader, meta, limits{frameMax: frame.FrameMax, attemptMax: ATTEMPT_STREAM_MAX})
}

// ParseEach emits each normalized event as soon as its provider event is
// parsed. It never waits for the terminal event before releasing earlier
// blocks, and an emitter failure stops provider consumption immediately.
func ParseEach(reader io.Reader, meta Meta, emit func(Event) error) error {
	if emit == nil {
		return errors.New("stream: nil event emitter")
	}
	_, err := parseWithLimitsEmit(reader, meta, limits{frameMax: frame.FrameMax, attemptMax: ATTEMPT_STREAM_MAX}, emit)
	return err
}

func parseWithLimits(reader io.Reader, meta Meta, bounds limits) []Event {
	events, _ := parseWithLimitsEmit(reader, meta, bounds, nil)
	return events
}

func parseWithLimitsEmit(reader io.Reader, meta Meta, bounds limits, emit func(Event) error) ([]Event, error) {
	parser := parserState{
		meta:   meta,
		bounds: bounds,
		blocks: make(map[int]*block),
		emit:   emit,
	}
	parser.append(Event{Kind: AttemptStarted})
	if parser.emitErr != nil {
		return parser.events, parser.emitErr
	}
	buffered := bufio.NewReaderSize(reader, 64*1024)
	var data bytes.Buffer
	for {
		line, consumed, err := readBoundedLine(buffered, int(bounds.frameMax)+1)
		parser.total += uint64(consumed)
		if parser.total > bounds.attemptMax {
			parser.fail("frame_overflow", "attempt stream limit exceeded")
			return parser.events, parser.emitErr
		}
		if errors.Is(err, errEventOverflow) {
			parser.fail("frame_overflow", "provider event exceeds frame limit")
			return parser.events, parser.emitErr
		}
		if len(line) > 0 {
			trimmed := strings.TrimSuffix(strings.TrimSuffix(string(line), "\n"), "\r")
			switch {
			case trimmed == "":
				if data.Len() > 0 && parser.consume(data.Bytes()) {
					return parser.events, parser.emitErr
				}
				data.Reset()
			case strings.HasPrefix(trimmed, "data:"):
				value := strings.TrimPrefix(trimmed, "data:")
				value = strings.TrimPrefix(value, " ")
				separatorBytes := 0
				if data.Len() > 0 {
					separatorBytes = 1
				}
				if data.Len()+separatorBytes+len(value) > int(bounds.frameMax) {
					parser.fail("frame_overflow", "provider event exceeds frame limit")
					return parser.events, parser.emitErr
				}
				if separatorBytes != 0 {
					data.WriteByte('\n')
				}
				data.WriteString(value)
			case strings.HasPrefix(trimmed, ":"), strings.HasPrefix(trimmed, "event:"), strings.HasPrefix(trimmed, "id:"):
				// SSE metadata is non-semantic for the compiled Responses profile.
			default:
				parser.fail("protocol", "invalid SSE field")
				return parser.events, parser.emitErr
			}
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				if data.Len() > 0 && parser.consume(data.Bytes()) {
					return parser.events, parser.emitErr
				}
				if !parser.terminal {
					parser.fail("protocol", "stream ended without terminal")
				}
				return parser.events, parser.emitErr
			}
			if errors.Is(err, transport.ErrStallDeadline) {
				parser.fail("timeout_stall", "provider stream stalled")
			} else {
				parser.fail("transport", "provider stream read failed")
			}
			return parser.events, parser.emitErr
		}
	}
}

var errEventOverflow = errors.New("provider event overflow")

func readBoundedLine(reader *bufio.Reader, maximum int) ([]byte, int, error) {
	var line []byte
	for {
		fragment, err := reader.ReadSlice('\n')
		line = append(line, fragment...)
		if len(line) > maximum {
			return nil, len(line), errEventOverflow
		}
		if !errors.Is(err, bufio.ErrBufferFull) {
			return line, len(line), err
		}
	}
}

type parserState struct {
	meta        Meta
	bounds      limits
	events      []Event
	blocks      map[int]*block
	total       uint64
	terminal    bool
	sawToolCall bool
	emit        func(Event) error
	emitErr     error
}

type block struct {
	kind       Kind
	blockID    string
	toolCallID string
	name       string
	arguments  []byte
	ended      bool
}

type providerEvent struct {
	Type        string          `json:"type"`
	OutputIndex int             `json:"output_index"`
	Delta       json.RawMessage `json:"delta"`
	Arguments   string          `json:"arguments"`
	Item        json.RawMessage `json:"item"`
	Response    *wireResponse   `json:"response"`
	Usage       *wireUsage      `json:"usage"`
}

type wireItem struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	CallID    string `json:"call_id"`
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type wireResponse struct {
	Status            string             `json:"status"`
	IncompleteDetails *incompleteDetails `json:"incomplete_details"`
	Usage             *wireUsage         `json:"usage"`
}

type incompleteDetails struct {
	Reason string `json:"reason"`
}

type wireUsage struct {
	InputTokens         uint64              `json:"input_tokens"`
	OutputTokens        uint64              `json:"output_tokens"`
	InputTokensDetails  *inputTokenDetails  `json:"input_tokens_details"`
	OutputTokensDetails *outputTokenDetails `json:"output_tokens_details"`
}

type inputTokenDetails struct {
	CachedTokens uint64 `json:"cached_tokens"`
}

type outputTokenDetails struct {
	ReasoningTokens uint64 `json:"reasoning_tokens"`
}

func (parser *parserState) consume(data []byte) bool {
	if bytes.Equal(data, []byte("[DONE]")) {
		if !parser.terminal {
			parser.fail("protocol", "DONE before terminal")
		}
		return parser.terminal
	}
	var wire providerEvent
	if err := json.Unmarshal(data, &wire); err != nil || wire.Type == "" {
		parser.fail("protocol", "malformed provider event")
		return true
	}
	switch wire.Type {
	case "response.created", "response.in_progress", "response.content_part.added", "response.content_part.done", "response.output_text.annotation.added", "response.reasoning_summary_part.added", "response.reasoning_summary_part.done":
		return false
	case "response.output_item.added":
		parser.startBlock(wire.OutputIndex, wire.Item)
	case "response.output_text.delta":
		text, ok := decodeString(wire.Delta)
		if !ok {
			parser.fail("protocol", "invalid text delta")
			return true
		}
		current := parser.ensureBlock(wire.OutputIndex, TextStart)
		parser.appendText(current.blockID, TextDelta, text)
	case "response.output_text.done":
		parser.endSimple(wire.OutputIndex, TextEnd)
	case "response.reasoning_summary_text.delta":
		text, ok := decodeString(wire.Delta)
		if !ok {
			parser.fail("protocol", "invalid reasoning delta")
			return true
		}
		current := parser.ensureBlock(wire.OutputIndex, ReasoningStart)
		parser.appendText(current.blockID, ReasoningDelta, text)
	case "response.reasoning_summary_text.done":
		// The opaque replay item arrives at output_item.done; do not end early.
	case "response.function_call_arguments.delta":
		fragment, ok := decodeString(wire.Delta)
		if !ok {
			parser.fail("protocol", "invalid tool arguments delta")
			return true
		}
		current := parser.ensureBlock(wire.OutputIndex, ToolCallStart)
		current.arguments = append(current.arguments, fragment...)
		parser.append(Event{Kind: ToolCallDelta, BlockID: current.blockID})
	case "response.function_call_arguments.done":
		if !parser.endTool(wire.OutputIndex, wire.Arguments) {
			return true
		}
	case "response.output_item.done":
		if !parser.endItem(wire.OutputIndex, wire.Item) {
			return true
		}
	case "response.usage":
		usage, ok := normalizeUsage(wire.Usage)
		if !ok {
			parser.fail("protocol", "invalid usage")
			return true
		}
		parser.append(Event{Kind: UsageProgress, Usage: usage})
	case "response.completed":
		if !parser.complete(wire.Response, false) {
			return true
		}
	case "response.incomplete":
		if !parser.complete(wire.Response, true) {
			return true
		}
	case "response.failed", "error":
		parser.fail("protocol", "provider reported failure")
	default:
		parser.fail("protocol", "unknown provider event")
	}
	return parser.terminal
}

func (parser *parserState) startBlock(index int, raw json.RawMessage) {
	if parser.blocks[index] != nil {
		parser.fail("protocol", "duplicate output item")
		return
	}
	var item wireItem
	if len(raw) == 0 || json.Unmarshal(raw, &item) != nil {
		parser.fail("protocol", "invalid output item")
		return
	}
	blockID := item.ID
	if blockID == "" {
		blockID = "m8-block-" + strconv.Itoa(index)
	}
	current := &block{blockID: blockID, name: item.Name, toolCallID: item.CallID}
	switch item.Type {
	case "message":
		current.kind = TextStart
	case "reasoning":
		current.kind = ReasoningStart
	case "function_call":
		current.kind = ToolCallStart
		parser.sawToolCall = true
		if current.toolCallID == "" {
			current.toolCallID = "m8-call-" + strconv.Itoa(index)
		}
		current.arguments = append(current.arguments, item.Arguments...)
	default:
		parser.fail("protocol", "unknown output item")
		return
	}
	parser.blocks[index] = current
	parser.append(Event{Kind: current.kind, BlockID: current.blockID})
}

func (parser *parserState) ensureBlock(index int, kind Kind) *block {
	if current := parser.blocks[index]; current != nil {
		if current.kind != kind {
			parser.fail("protocol", "block kind mismatch")
		}
		return current
	}
	current := &block{kind: kind, blockID: "m8-block-" + strconv.Itoa(index)}
	if kind == ToolCallStart {
		current.toolCallID = "m8-call-" + strconv.Itoa(index)
		parser.sawToolCall = true
	}
	parser.blocks[index] = current
	parser.append(Event{Kind: kind, BlockID: current.blockID})
	return current
}

func (parser *parserState) endSimple(index int, kind Kind) {
	current := parser.blocks[index]
	wantStart := TextStart
	if kind == ReasoningEnd {
		wantStart = ReasoningStart
	}
	if current == nil || current.ended || current.kind != wantStart {
		parser.fail("protocol", "invalid block end")
		return
	}
	current.ended = true
	parser.append(Event{Kind: kind, BlockID: current.blockID})
}

func (parser *parserState) endTool(index int, supplied string) bool {
	current := parser.ensureBlock(index, ToolCallStart)
	if current.ended {
		parser.fail("protocol", "duplicate tool end")
		return false
	}
	arguments := current.arguments
	if supplied != "" {
		if len(arguments) > 0 && string(arguments) != supplied {
			parser.fail("protocol", "tool argument assembly mismatch")
			return false
		}
		arguments = []byte(supplied)
	}
	canonical, err := jcs.Canonicalize(arguments)
	if err != nil || len(canonical) == 0 || canonical[0] != '{' || current.name == "" {
		parser.fail("protocol", "tool arguments are not an object")
		return false
	}
	current.ended = true
	parser.append(Event{Kind: ToolCallEnd, BlockID: current.blockID, ToolCall: &ToolCall{ToolCallID: current.toolCallID, Name: current.name, Arguments: canonical}})
	return !parser.terminal
}

func (parser *parserState) endItem(index int, raw json.RawMessage) bool {
	current := parser.blocks[index]
	if current == nil {
		parser.fail("protocol", "output item ended before start")
		return false
	}
	if current.kind == ToolCallStart {
		if current.ended {
			return true
		}
		var item wireItem
		if json.Unmarshal(raw, &item) != nil {
			parser.fail("protocol", "invalid tool output item")
			return false
		}
		return parser.endTool(index, item.Arguments)
	}
	if current.kind == ReasoningStart {
		if current.ended || len(raw) == 0 || raw[0] != '{' {
			parser.fail("protocol", "invalid reasoning output item")
			return false
		}
		current.ended = true
		payload := append([]byte(nil), raw...)
		parser.append(Event{Kind: ReasoningEnd, BlockID: current.blockID, ReplayEnvelope: &ReplayEnvelope{OriginProviderLaneID: parser.meta.ProviderLaneID, OriginTurnID: parser.meta.TurnID, Payload: payload}})
		return !parser.terminal
	}
	if !current.ended {
		current.ended = true
		parser.append(Event{Kind: TextEnd, BlockID: current.blockID})
	}
	return !parser.terminal
}

func (parser *parserState) complete(response *wireResponse, incomplete bool) bool {
	if response == nil {
		parser.fail("protocol", "terminal response missing")
		return false
	}
	usage, ok := normalizeUsage(response.Usage)
	if !ok {
		parser.fail("protocol", "terminal usage missing")
		return false
	}
	for _, current := range parser.blocks {
		if !current.ended {
			parser.fail("protocol", "terminal arrived before block end")
			return false
		}
	}
	finish := "stop"
	if parser.sawToolCall {
		finish = "tool_calls"
	}
	if incomplete || response.Status == "incomplete" {
		if response.IncompleteDetails == nil || response.IncompleteDetails.Reason != "max_output_tokens" {
			parser.fail("protocol", "unknown finish reason")
			return false
		}
		finish = "length"
	} else if response.Status != "completed" {
		parser.fail("protocol", "unknown finish status")
		return false
	}
	parser.append(Event{Kind: Completed, FinishReason: finish, Usage: usage})
	return !parser.terminal
}

func normalizeUsage(wire *wireUsage) (*Usage, bool) {
	if wire == nil {
		return nil, false
	}
	usage := &Usage{InputTokens: wire.InputTokens, OutputTokens: wire.OutputTokens}
	if wire.OutputTokensDetails != nil {
		value := wire.OutputTokensDetails.ReasoningTokens
		usage.ReasoningTokens = &value
	}
	if wire.InputTokensDetails != nil {
		value := wire.InputTokensDetails.CachedTokens
		usage.CacheReadTokens = &value
	}
	return usage, true
}

func decodeString(raw json.RawMessage) (string, bool) {
	var value string
	if len(raw) == 0 || json.Unmarshal(raw, &value) != nil || !utf8.ValidString(value) {
		return "", false
	}
	return value, true
}

func (parser *parserState) appendText(blockID string, kind Kind, value string) {
	const safeChunk = int(frame.FrameMax) / 8
	for len(value) > safeChunk {
		cut := safeChunk
		for cut > 0 && !utf8.RuneStart(value[cut]) {
			cut--
		}
		parser.append(Event{Kind: kind, BlockID: blockID, Text: value[:cut]})
		if parser.terminal {
			return
		}
		value = value[cut:]
	}
	parser.append(Event{Kind: kind, BlockID: blockID, Text: value})
}

func (parser *parserState) append(event Event) {
	if parser.terminal {
		return
	}
	event.Schema = SchemaV2
	event.AttemptID = parser.meta.AttemptID
	if event.Terminal() {
		event.FrozenCoreDigest = parser.meta.FrozenCoreDigest
		event.ProviderLoweredToolsDigest = parser.meta.ProviderLoweredToolsDigest
	}
	raw, err := json.Marshal(event)
	if err != nil || len(raw) > int(parser.bounds.frameMax) {
		parser.fail("frame_overflow", "normalized event exceeds frame limit")
		return
	}
	if parser.emit != nil {
		if err := parser.emit(event); err != nil {
			parser.emitErr = err
			parser.terminal = true
			return
		}
	} else {
		parser.events = append(parser.events, event)
	}
	parser.terminal = event.Terminal()
}

func (parser *parserState) fail(class, detail string) {
	if parser.terminal {
		return
	}
	parser.append(Event{Kind: Failed, ErrorClass: class, Detail: detail})
}

func (event Event) String() string {
	return fmt.Sprintf("%s{%s}", event.Schema, event.Kind)
}
