package main

import (
	"encoding/json"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func SchemaFromForm(form fieldspec.Form, digest string) map[string]any {
	headerProps := map[string]any{}
	for name, field := range form.Fields {
		if skipFormField(name) {
			continue
		}
		headerProps[name] = propertyForField(field)
	}
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"headers", "form_digest"},
		"properties": map[string]any{
			"headers": map[string]any{
				"type":                 "object",
				"additionalProperties": false,
				"properties":           headerProps,
			},
			"to": map[string]any{
				"type": "string",
			},
			"cc": map[string]any{
				"type":        "string",
				"description": structuredDescription("address_list"),
			},
			"dispatch_id": map[string]any{
				"type": "string",
			},
			"body": map[string]any{
				"type": "string",
			},
			"form_digest": map[string]any{
				"type":  "string",
				"const": digest,
			},
		},
	}
}

func propertyForField(field fieldspec.Field) map[string]any {
	prop := map[string]any{"type": "string"}
	if len(field.Options) > 0 {
		prop["enum"] = append([]string(nil), field.Options...)
	}
	if field.Default != "" {
		prop["default"] = field.Default
	}
	if description := structuredDescription(field.Type); description != "" {
		prop["description"] = description
	}
	return prop
}

func structuredDescription(fieldType string) string {
	switch fieldType {
	case "row_array":
		return "canonical JSON string - array of row objects"
	case "object":
		return "canonical JSON string - object"
	case "address_list":
		return "canonical JSON string - array of address strings"
	default:
		return ""
	}
}

func skipFormField(name string) bool {
	switch strings.ToLower(name) {
	case "from", "role", "relay_id", "delivery_state":
		return true
	}
	switch name {
	case "TO", "CC", "DISPATCH_ID":
		return true
	default:
		return false
	}
}

func SubmitPayloadFromArguments(args json.RawMessage) (json.RawMessage, error) {
	var in submitArguments
	if len(args) > 0 {
		if err := json.Unmarshal(args, &in); err != nil {
			return nil, err
		}
	}
	headers := map[string]string{}
	for key, value := range in.Headers {
		headers[key] = value
	}
	if in.CC != "" {
		headers["CC"] = in.CC
	}
	payload := fieldspec.SubmitPayload{
		Record: record.Record{
			Envelope: record.Envelope{
				To:         in.To,
				DispatchID: in.DispatchID,
			},
			Headers: headers,
			Body:    in.Body,
		},
		FormDigest: in.FormDigest,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type submitArguments struct {
	Headers    map[string]string `json:"headers"`
	To         string            `json:"to"`
	CC         string            `json:"cc"`
	DispatchID string            `json:"dispatch_id"`
	Body       string            `json:"body"`
	FormDigest string            `json:"form_digest"`
}
