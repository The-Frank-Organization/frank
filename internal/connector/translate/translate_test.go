package translate

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"testing"

	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/request"
)

func TestOpenAIResponsesTranslationExactBytes(t *testing.T) {
	t.Parallel()

	result, err := Translate(completeRequest(), capableLane(true))
	if err != nil {
		t.Fatalf("Translate() error = %v", err)
	}
	want := `{"include":["reasoning.encrypted_content"],"input":[{"content":"hello","role":"user"},{"content":"working","role":"assistant"},{"arguments":"{\"path\":\"README.md\"}","call_id":"call-1","name":"read","type":"function_call"},{"call_id":"call-1","output":"file text","type":"function_call_output"},{"encrypted_content":"opaque","id":"rs_1","summary":[],"type":"reasoning"}],"instructions":"system text","max_output_tokens":1024,"model":"gpt-5.1-codex","reasoning":{"effort":"high"},"store":false,"stream":true,"temperature":0.2,"tools":[{"description":"","name":"read","parameters":{"additionalProperties":false,"type":"object"},"strict":true,"type":"function"}]}`
	if string(result.Body) != want {
		t.Fatalf("Translate() body\n got: %s\nwant: %s", result.Body, want)
	}
	if result.ProfileVersion != "openai-responses.v1" {
		t.Fatalf("ProfileVersion = %q", result.ProfileVersion)
	}
	digest, err := DigestObservedTools(result.Body)
	if err != nil || digest != "95ba4f43e9afa69295bc04a1477533d662e9d91b7a1ef22a52eaaf5a2058abb1" {
		t.Fatalf("DigestObservedTools() = %q, %v", digest, err)
	}
}

func TestTranslationUnwrapsOnlyOpaqueReplayPayload(t *testing.T) {
	t.Parallel()

	result, err := Translate(completeRequest(), capableLane(true))
	if err != nil {
		t.Fatalf("Translate() error = %v", err)
	}
	for _, forbidden := range [][]byte{[]byte("origin_provider_lane_id"), []byte("origin_turn_id"), []byte("turn-1"), []byte("lane-codex-1")} {
		if bytes.Contains(result.Body, forbidden) {
			t.Fatalf("translated body contains replay-wrapper provenance %q", forbidden)
		}
	}
	payload := []byte(`{"encrypted_content":"opaque","id":"rs_1","summary":[],"type":"reasoning"}`)
	if bytes.Count(result.Body, payload) != 1 {
		t.Fatalf("opaque replay payload occurrence count = %d, want 1", bytes.Count(result.Body, payload))
	}
}

func TestOpaqueReplayValidationNeverTrimsProviderBytes(t *testing.T) {
	t.Parallel()

	requestValue := completeRequest()
	payload := []byte(" {\"type\":\"reasoning\"} ")
	requestValue.Input = []request.InputItem{request.ReasoningReplay{Envelope: request.ReplayEnvelope{Payload: base64.StdEncoding.EncodeToString(payload)}}}
	if _, err := Translate(requestValue, capableLane(true)); !errors.Is(err, ErrOpaqueReplay) {
		t.Fatalf("Translate(whitespace-wrapped replay) error = %v, want ErrOpaqueReplay rather than byte rewriting", err)
	}
}

func TestZeroToolsIsPresentEmptyAndHasFrozenDigest(t *testing.T) {
	t.Parallel()

	requestValue := completeRequest()
	requestValue.Tools = []request.Tool{}
	result, err := Translate(requestValue, capableLane(true))
	if err != nil {
		t.Fatalf("Translate() error = %v", err)
	}
	if !bytes.Contains(result.Body, []byte(`"tools":[]`)) {
		t.Fatalf("translated body lacks present-empty tools: %s", result.Body)
	}
	const want = "4f53cda18c2baa0c0354bb5f9a3ecbe5ed12ab4d8e11ba873c2f11161202b945"
	observed, err := DigestObservedTools(result.Body)
	if err != nil || observed != want {
		t.Fatalf("DigestObservedTools() = %q, %v", observed, err)
	}

	withoutTools := bytes.Replace(result.Body, []byte(`,"tools":[]`), nil, 1)
	if _, err := DigestObservedTools(withoutTools); !errors.Is(err, ErrMissingTools) {
		t.Fatalf("DigestObservedTools(missing) error = %v, want ErrMissingTools", err)
	}
}

func TestLoweredToolsDigestChangesWithOrderAndLaneStrictness(t *testing.T) {
	t.Parallel()

	requestValue := completeRequest()
	requestValue.Tools = append(requestValue.Tools, request.Tool{
		Name:        "write",
		Description: "write file",
		Parameters:  json.RawMessage(`{"type":"object"}`),
	})
	first, err := Translate(requestValue, capableLane(true))
	if err != nil {
		t.Fatalf("Translate(first) error = %v", err)
	}
	requestValue.Tools[0], requestValue.Tools[1] = requestValue.Tools[1], requestValue.Tools[0]
	reordered, err := Translate(requestValue, capableLane(true))
	if err != nil {
		t.Fatalf("Translate(reordered) error = %v", err)
	}
	strictFalse, err := Translate(requestValue, capableLane(false))
	if err != nil {
		t.Fatalf("Translate(strict false) error = %v", err)
	}
	firstDigest, _ := DigestObservedTools(first.Body)
	reorderedDigest, _ := DigestObservedTools(reordered.Body)
	strictFalseDigest, _ := DigestObservedTools(strictFalse.Body)
	if firstDigest == reorderedDigest || reorderedDigest == strictFalseDigest {
		t.Fatalf("digest did not bind order/strict: %q %q %q", firstDigest, reorderedDigest, strictFalseDigest)
	}
}

func TestObservedToolsRequiresExactFiveMemberFieldSet(t *testing.T) {
	t.Parallel()

	result, err := Translate(completeRequest(), capableLane(true))
	if err != nil {
		t.Fatalf("Translate() error = %v", err)
	}
	withoutDescription := bytes.Replace(result.Body, []byte(`"description":"",`), nil, 1)
	if _, err := DigestObservedTools(withoutDescription); !errors.Is(err, ErrInvalidTools) {
		t.Fatalf("DigestObservedTools(omitted description) error = %v, want ErrInvalidTools", err)
	}
}

func TestTranslationIsDeterministic(t *testing.T) {
	t.Parallel()

	requestValue := completeRequest()
	lane := capableLane(true)
	first, err := Translate(requestValue, lane)
	if err != nil {
		t.Fatalf("Translate(first) error = %v", err)
	}
	for iteration := 0; iteration < 100; iteration++ {
		next, err := Translate(requestValue, lane)
		if err != nil {
			t.Fatalf("Translate(%d) error = %v", iteration, err)
		}
		if !bytes.Equal(next.Body, first.Body) {
			t.Fatalf("Translate(%d) was non-deterministic", iteration)
		}
	}
}

func completeRequest() *request.Request {
	effort := "high"
	temperature := json.Number("0.2")
	return &request.Request{
		Instructions: "system text",
		Input: []request.InputItem{
			request.UserText{Text: "hello"},
			request.AssistantText{Text: "working"},
			request.AssistantToolCall{ToolCallID: "call-1", Name: "read", Arguments: `{"path":"README.md"}`},
			request.ToolResult{ToolCallID: "call-1", Content: "file text"},
			request.ReasoningReplay{Envelope: request.ReplayEnvelope{
				OriginProviderLaneID: "lane-codex-1",
				OriginTurnID:         "turn-1",
				Payload:              "eyJlbmNyeXB0ZWRfY29udGVudCI6Im9wYXF1ZSIsImlkIjoicnNfMSIsInN1bW1hcnkiOltdLCJ0eXBlIjoicmVhc29uaW5nIn0=",
			}},
		},
		Tools: []request.Tool{{
			Name:        "read",
			Description: "",
			Parameters:  json.RawMessage(`{"additionalProperties":false,"type":"object"}`),
		}},
		Sampling:  request.Sampling{MaxOutputTokens: 1024, Temperature: &temperature},
		Reasoning: request.Reasoning{Effort: &effort},
	}
}

func capableLane(strict bool) catalog.Lane {
	return catalog.Lane{
		ModelID:    "gpt-5.1-codex",
		CompatMode: "openai-responses",
		Wire: catalog.Wire{
			Streaming:        true,
			UsageInStreaming: true,
			ServerRetention:  false,
		},
		Reasoning: catalog.Reasoning{Supported: true, ReplayKind: "opaque_item"},
		ToolUse:   catalog.ToolUse{Supported: true, StrictSchema: strict},
	}
}
