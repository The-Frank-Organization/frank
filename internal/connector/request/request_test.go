package request

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/jcs"
)

const validRequestJSON = `{
  "schema":"m8.llm_request.v1",
  "run_id":"run-1",
  "turn_id":"turn-1",
  "attempt_id":"attempt-1",
  "turn_epoch":"7",
  "provider_lane_id":"lane-codex-1",
  "instructions":"system text",
  "input":[
    {"kind":"user_text","text":"hello"},
    {"kind":"assistant_text","text":"working"},
    {"kind":"assistant_tool_call","tool_call_id":"call-1","name":"read","arguments":"{\"path\":\"README.md\"}"},
    {"kind":"tool_result","tool_call_id":"call-1","content":"file text"},
    {"kind":"reasoning_replay","envelope":{"origin_provider_lane_id":"lane-codex-1","origin_turn_id":"turn-1","payload":"eyJlbmNyeXB0ZWRfY29udGVudCI6Im9wYXF1ZSIsImlkIjoicnNfMSIsInN1bW1hcnkiOltdLCJ0eXBlIjoicmVhc29uaW5nIn0="}}
  ],
  "tools":[{"name":"read","description":"","parameters":{"additionalProperties":false,"type":"object"}}],
  "sampling":{"max_output_tokens":1024,"temperature":0.2},
  "reasoning":{"effort":"high"}
}`

func TestParseAcceptsClosedRequestAndAllFiveInputKinds(t *testing.T) {
	t.Parallel()

	request, err := Parse(canonical(t, validRequestJSON), capableLane())
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if len(request.Input) != 5 {
		t.Fatalf("len(Input) = %d, want 5", len(request.Input))
	}
	if _, ok := request.Input[0].(UserText); !ok {
		t.Fatalf("Input[0] type = %T, want UserText", request.Input[0])
	}
	if _, ok := request.Input[1].(AssistantText); !ok {
		t.Fatalf("Input[1] type = %T, want AssistantText", request.Input[1])
	}
	if _, ok := request.Input[2].(AssistantToolCall); !ok {
		t.Fatalf("Input[2] type = %T, want AssistantToolCall", request.Input[2])
	}
	toolResult, ok := request.Input[3].(ToolResult)
	if !ok || toolResult.Content != "file text" {
		t.Fatalf("Input[3] = %#v, want string ToolResult", request.Input[3])
	}
	replay, ok := request.Input[4].(ReasoningReplay)
	if !ok || replay.Envelope.Payload != "eyJlbmNyeXB0ZWRfY29udGVudCI6Im9wYXF1ZSIsImlkIjoicnNfMSIsInN1bW1hcnkiOltdLCJ0eXBlIjoicmVhc29uaW5nIn0=" {
		t.Fatalf("Input[4] = %#v, want opaque ReasoningReplay", request.Input[4])
	}
}

func TestParsePreservesPresentEmptyTools(t *testing.T) {
	t.Parallel()

	raw := strings.Replace(validRequestJSON, `"tools":[{"name":"read","description":"","parameters":{"additionalProperties":false,"type":"object"}}]`, `"tools":[]`, 1)
	request, err := Parse(canonical(t, raw), capableLane())
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if request.Tools == nil || len(request.Tools) != 0 {
		t.Fatalf("Tools = %#v, want present empty slice", request.Tools)
	}
}

func TestParseRejectsMalformedClosedSchemaForms(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		raw  string
	}{
		{name: "unknown outer field", raw: strings.Replace(validRequestJSON, `"schema":`, `"api_key":"secret-shaped","schema":`, 1)},
		{name: "unknown item field", raw: strings.Replace(validRequestJSON, `"kind":"user_text","text":"hello"`, `"kind":"user_text","role":"user","text":"hello"`, 1)},
		{name: "unknown item kind", raw: strings.Replace(validRequestJSON, `"kind":"user_text"`, `"kind":"image"`, 1)},
		{name: "legacy replay payload", raw: strings.Replace(validRequestJSON, `"envelope":{"origin_provider_lane_id":"lane-codex-1","origin_turn_id":"turn-1","payload":"eyJlbmNyeXB0ZWRfY29udGVudCI6Im9wYXF1ZSIsImlkIjoicnNfMSIsInN1bW1hcnkiOltdLCJ0eXBlIjoicmVhc29uaW5nIn0="}`, `"replay_payload":"opaque-provider-bytes"`, 1)},
		{name: "non-base64 replay payload", raw: strings.Replace(validRequestJSON, `"payload":"eyJlbmNyeXB0ZWRfY29udGVudCI6Im9wYXF1ZSIsImlkIjoicnNfMSIsInN1bW1hcnkiOltdLCJ0eXBlIjoicmVhc29uaW5nIn0="`, `"payload":"not base64"`, 1)},
		{name: "structured tool result", raw: strings.Replace(validRequestJSON, `"content":"file text"`, `"content":{"text":"file text"}`, 1)},
		{name: "parameters not object", raw: strings.Replace(validRequestJSON, `"parameters":{"additionalProperties":false,"type":"object"}`, `"parameters":[]`, 1)},
		{name: "missing tools", raw: strings.Replace(validRequestJSON, `"tools":[{"name":"read","description":"","parameters":{"additionalProperties":false,"type":"object"}}],`, ``, 1)},
		{name: "invalid counter", raw: strings.Replace(validRequestJSON, `"turn_epoch":"7"`, `"turn_epoch":"07"`, 1)},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if _, err := Parse(canonical(t, test.raw), capableLane()); !errors.Is(err, MalformedRequest) {
				t.Fatalf("Parse() error = %v, want malformed_request", err)
			}
		})
	}
}

func TestParseRequiresExactJCSBytes(t *testing.T) {
	t.Parallel()

	if _, err := Parse([]byte(validRequestJSON), capableLane()); !errors.Is(err, MalformedRequest) {
		t.Fatalf("Parse(non-canonical) error = %v, want malformed_request", err)
	}
}

func TestParseRejectsReplayScopeBeforeTranslation(t *testing.T) {
	t.Parallel()

	for _, mutation := range []string{
		`"origin_provider_lane_id":"lane-other"`,
		`"origin_turn_id":"turn-other"`,
	} {
		raw := validRequestJSON
		if strings.Contains(mutation, "provider_lane") {
			raw = strings.Replace(raw, `"origin_provider_lane_id":"lane-codex-1"`, mutation, 1)
		} else {
			raw = strings.Replace(raw, `"origin_turn_id":"turn-1"`, mutation, 1)
		}
		if _, err := Parse(canonical(t, raw), capableLane()); !errors.Is(err, ReplayScopeViolation) {
			t.Fatalf("Parse(scope mutation %s) error = %v, want replay_scope_violation", mutation, err)
		}
	}
}

func TestParseRejectsLaneCapabilityMismatch(t *testing.T) {
	t.Parallel()

	notSupported := capableLane()
	notSupported.Reasoning.Supported = false
	if _, err := Parse(canonical(t, validRequestJSON), notSupported); !errors.Is(err, LaneCapabilityMismatch) {
		t.Fatalf("Parse(unsupported reasoning) error = %v, want lane_capability_mismatch", err)
	}

	unknownEffort := strings.Replace(validRequestJSON, `"effort":"high"`, `"effort":"ultra"`, 1)
	if _, err := Parse(canonical(t, unknownEffort), capableLane()); !errors.Is(err, LaneCapabilityMismatch) {
		t.Fatalf("Parse(unknown effort) error = %v, want lane_capability_mismatch", err)
	}
}

func TestParseChecksReplayScopeBeforeLaneCapability(t *testing.T) {
	t.Parallel()

	raw := strings.Replace(validRequestJSON, `"origin_turn_id":"turn-1"`, `"origin_turn_id":"turn-other"`, 1)
	lane := capableLane()
	lane.Reasoning.Supported = false
	if _, err := Parse(canonical(t, raw), lane); !errors.Is(err, ReplayScopeViolation) {
		t.Fatalf("Parse() error = %v, want replay_scope_violation before lane capability", err)
	}
}

func capableLane() catalog.Lane {
	return catalog.Lane{
		LaneID:     "lane-codex-1",
		CompatMode: "openai-responses",
		ToolUse:    catalog.ToolUse{Supported: true, StrictSchema: true},
		Reasoning: catalog.Reasoning{
			Supported:    true,
			EffortLevels: []string{"high", "low", "medium"},
			ReplayKind:   "opaque_item",
		},
	}
}

func canonical(t *testing.T, raw string) []byte {
	t.Helper()
	result, err := jcs.Canonicalize([]byte(raw))
	if err != nil {
		t.Fatalf("canonical fixture: %v", err)
	}
	if !json.Valid(result) {
		t.Fatal("canonical fixture is not JSON")
	}
	return result
}
