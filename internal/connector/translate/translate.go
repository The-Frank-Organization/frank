// Package translate implements the compiled openai-responses lowering profile.
// It has no endpoint, credential, authorization, or transport inputs.
package translate

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/The-Frank-Organization/frank/internal/connector/catalog"
	"github.com/The-Frank-Organization/frank/internal/connector/jcs"
	"github.com/The-Frank-Organization/frank/internal/connector/request"
)

const OpenAIResponsesProfileVersion = "openai-responses.v1"

var (
	ErrUnsupportedProfile = errors.New("translate: unsupported compatibility profile")
	ErrOpaqueReplay       = errors.New("translate: invalid opaque replay carrier")
	ErrMissingTools       = errors.New("translate: observed body is missing tools")
	ErrInvalidTools       = errors.New("translate: observed tools violate locked field set")
)

type Result struct {
	Body           []byte
	LoweredTools   []byte
	ProfileVersion string
}

// Translate is a deterministic function of the validated app request, pinned
// lane facts, and the compiled profile. Opaque replay bytes are mechanically
// unwrapped and spliced without interpreting their provider-owned contents.
func Translate(input *request.Request, lane catalog.Lane) (Result, error) {
	if input == nil || lane.CompatMode != "openai-responses" {
		return Result{}, ErrUnsupportedProfile
	}
	items := make([][]byte, 0, len(input.Input))
	for _, item := range input.Input {
		lowered, err := lowerItem(item)
		if err != nil {
			return Result{}, err
		}
		items = append(items, lowered)
	}
	loweredTools, err := lowerTools(input.Tools, lane.ToolUse.StrictSchema)
	if err != nil {
		return Result{}, err
	}

	fields := map[string][]byte{
		"input":             encodeArray(items),
		"instructions":      mustMarshal(input.Instructions),
		"max_output_tokens": []byte(strconv.FormatInt(input.Sampling.MaxOutputTokens, 10)),
		"model":             mustMarshal(lane.ModelID),
		"store":             []byte("false"),
		"stream":            []byte(strconv.FormatBool(lane.Wire.Streaming)),
		"tools":             loweredTools,
	}
	if lane.Reasoning.ReplayKind == "opaque_item" {
		fields["include"] = mustMarshal([]string{"reasoning.encrypted_content"})
	}
	if input.Reasoning.Effort != nil {
		fields["reasoning"] = mustCanonical(struct {
			Effort string `json:"effort"`
		}{Effort: *input.Reasoning.Effort})
	}
	if input.Sampling.Temperature != nil {
		fields["temperature"] = []byte(input.Sampling.Temperature.String())
	}
	body := encodeObject(fields)
	if !json.Valid(body) {
		return Result{}, ErrOpaqueReplay
	}
	return Result{
		Body:           body,
		LoweredTools:   append([]byte(nil), loweredTools...),
		ProfileVersion: OpenAIResponsesProfileVersion,
	}, nil
}

// DigestObservedTools enforces presence and the locked five-member tool shape
// before deriving E from a captured wire body.
func DigestObservedTools(body []byte) (string, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(body, &object); err != nil || object == nil {
		return "", fmt.Errorf("%w: body is not an object", ErrInvalidTools)
	}
	toolsRaw, ok := object["tools"]
	if !ok {
		return "", ErrMissingTools
	}
	var tools []json.RawMessage
	if err := json.Unmarshal(toolsRaw, &tools); err != nil || tools == nil {
		return "", fmt.Errorf("%w: tools is not an array", ErrInvalidTools)
	}
	for _, toolRaw := range tools {
		if err := validateObservedTool(toolRaw); err != nil {
			return "", err
		}
	}
	canonical, err := jcs.Canonicalize(toolsRaw)
	if err != nil {
		return "", fmt.Errorf("%w: canonicalize tools", ErrInvalidTools)
	}
	digest := sha256.Sum256(canonical)
	return hex.EncodeToString(digest[:]), nil
}

func lowerItem(item request.InputItem) ([]byte, error) {
	switch item := item.(type) {
	case request.UserText:
		return mustCanonical(struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		}{Content: item.Text, Role: "user"}), nil
	case request.AssistantText:
		return mustCanonical(struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		}{Content: item.Text, Role: "assistant"}), nil
	case request.AssistantToolCall:
		return mustCanonical(struct {
			Arguments string `json:"arguments"`
			CallID    string `json:"call_id"`
			Name      string `json:"name"`
			Type      string `json:"type"`
		}{Arguments: item.Arguments, CallID: item.ToolCallID, Name: item.Name, Type: "function_call"}), nil
	case request.ToolResult:
		return mustCanonical(struct {
			CallID string `json:"call_id"`
			Output string `json:"output"`
			Type   string `json:"type"`
		}{CallID: item.ToolCallID, Output: item.Content, Type: "function_call_output"}), nil
	case request.ReasoningReplay:
		decoded, err := base64.StdEncoding.Strict().DecodeString(item.Envelope.Payload)
		if err != nil || len(decoded) < 2 || decoded[0] != '{' || decoded[len(decoded)-1] != '}' || !json.Valid(decoded) {
			return nil, ErrOpaqueReplay
		}
		return append([]byte(nil), decoded...), nil
	default:
		return nil, fmt.Errorf("translate: unknown validated input item %T", item)
	}
}

func lowerTools(tools []request.Tool, strict bool) ([]byte, error) {
	lowered := make([][]byte, 0, len(tools))
	for _, tool := range tools {
		canonicalParameters, err := jcs.Canonicalize(tool.Parameters)
		if err != nil {
			return nil, fmt.Errorf("translate: invalid tool parameters")
		}
		lowered = append(lowered, mustCanonical(struct {
			Description string          `json:"description"`
			Name        string          `json:"name"`
			Parameters  json.RawMessage `json:"parameters"`
			Strict      bool            `json:"strict"`
			Type        string          `json:"type"`
		}{
			Description: tool.Description,
			Name:        tool.Name,
			Parameters:  canonicalParameters,
			Strict:      strict,
			Type:        "function",
		}))
	}
	return encodeArray(lowered), nil
}

func validateObservedTool(raw []byte) error {
	canonical, err := jcs.Canonicalize(raw)
	if err != nil {
		return fmt.Errorf("%w: malformed tool", ErrInvalidTools)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(canonical, &fields); err != nil || len(fields) != 5 {
		return fmt.Errorf("%w: tool member count", ErrInvalidTools)
	}
	for _, name := range []string{"description", "name", "parameters", "strict", "type"} {
		if _, ok := fields[name]; !ok {
			return fmt.Errorf("%w: missing %s", ErrInvalidTools, name)
		}
	}
	var kind, name, description string
	var strict bool
	if json.Unmarshal(fields["type"], &kind) != nil || kind != "function" ||
		json.Unmarshal(fields["name"], &name) != nil || name == "" ||
		json.Unmarshal(fields["description"], &description) != nil ||
		json.Unmarshal(fields["strict"], &strict) != nil {
		return fmt.Errorf("%w: invalid member type", ErrInvalidTools)
	}
	var parameters map[string]json.RawMessage
	if json.Unmarshal(fields["parameters"], &parameters) != nil || parameters == nil {
		return fmt.Errorf("%w: parameters is not an object", ErrInvalidTools)
	}
	return nil
}

func encodeArray(items [][]byte) []byte {
	result := []byte{'['}
	for index, item := range items {
		if index > 0 {
			result = append(result, ',')
		}
		result = append(result, item...)
	}
	return append(result, ']')
}

func encodeObject(fields map[string][]byte) []byte {
	names := make([]string, 0, len(fields))
	for name := range fields {
		names = append(names, name)
	}
	sort.Strings(names)
	result := []byte{'{'}
	for index, name := range names {
		if index > 0 {
			result = append(result, ',')
		}
		result = append(result, mustMarshal(name)...)
		result = append(result, ':')
		result = append(result, fields[name]...)
	}
	return append(result, '}')
}

func mustCanonical(value any) []byte {
	raw := mustMarshal(value)
	canonical, err := jcs.Canonicalize(raw)
	if err != nil {
		panic(err)
	}
	return canonical
}

func mustMarshal(value any) []byte {
	raw, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return raw
}
