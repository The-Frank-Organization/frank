// Package formschema owns the m-2 form-to-tool-schema mapping shared by the
// native worker relay tools and the retained MCP frontend.
package formschema

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sort"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

const MappingVersion = "m2-mapping-v1"

type Disposition struct {
	Class  string `json:"class"`
	Member string `json:"member"`
}

type SubmitArguments struct {
	Headers    map[string]string
	To         string
	CC         string
	DispatchID string
	Body       string
	FormDigest string

	headersPresent    bool
	formDigestPresent bool
}

func SchemaFromForm(form fieldspec.Form, digest string) map[string]any {
	headerProperties := make(map[string]any)
	for name, field := range form.Fields {
		if skipFormField(name) {
			continue
		}
		headerProperties[name] = propertyForField(field)
	}
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"headers", "form_digest"},
		"properties": map[string]any{
			"headers": map[string]any{
				"type":                 "object",
				"additionalProperties": true,
				"properties":           headerProperties,
			},
			"to":          map[string]any{"type": "string"},
			"cc":          map[string]any{"type": "string", "description": structuredDescription("address_list")},
			"dispatch_id": map[string]any{"type": "string"},
			"body":        map[string]any{"type": "string"},
			"form_digest": map[string]any{"type": "string", "const": digest},
		},
	}
}

func ProjectSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"view": map[string]any{"type": "string", "enum": []string{"default", "audit", "roster"}},
		},
		"additionalProperties": false,
	}
}

func ReadSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"relay_id": map[string]any{"type": "string"},
		},
		"required":             []string{"relay_id"},
		"additionalProperties": false,
	}
}

// StaticSchemas returns the build-stable input-schema surfaces used by the F58
// identity vector. The live form properties and digest const are deliberately
// absent from the submit template.
func StaticSchemas() map[string]map[string]any {
	submit := SchemaFromForm(fieldspec.Form{Fields: map[string]fieldspec.Field{}}, "")
	submit["properties"].(map[string]any)["form_digest"] = map[string]any{"type": "string"}
	return map[string]map[string]any{
		"relay.submit":  submit,
		"relay.project": ProjectSchema(),
		"relay.read":    ReadSchema(),
	}
}

// StaticSchemaDigests is the release-binding recomputation hook for m-2's
// three F58 producers.
func StaticSchemaDigests() (map[string]string, error) {
	digests := make(map[string]string, 3)
	for name, schema := range StaticSchemas() {
		encoded, err := json.Marshal(schema)
		if err != nil {
			return nil, err
		}
		digest := sha256.Sum256(encoded)
		digests[name] = hex.EncodeToString(digest[:])
	}
	return digests, nil
}

func SubmitPayloadFromArguments(encoded json.RawMessage) (json.RawMessage, error) {
	arguments, err := ParseSubmitArguments(encoded)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string, len(arguments.Headers)+1)
	for name, value := range arguments.Headers {
		headers[name] = value
	}
	if arguments.CC != "" {
		headers["CC"] = arguments.CC
	}
	payload := fieldspec.SubmitPayload{
		Record: record.Record{
			Envelope: record.Envelope{To: arguments.To, DispatchID: arguments.DispatchID},
			Headers:  headers,
			Body:     arguments.Body,
		},
		FormDigest: arguments.FormDigest,
	}
	return json.Marshal(payload)
}

func ValidateSubmitArguments(form fieldspec.Form, digest string, encoded json.RawMessage) []Disposition {
	arguments, err := ParseSubmitArguments(encoded)
	if err != nil {
		var mappingError *MappingError
		if ok := AsMappingError(err, &mappingError); ok {
			return []Disposition{{Class: "mapping-error", Member: string(mappingError.ID)}}
		}
		return []Disposition{{Class: "mapping-error", Member: "unknown"}}
	}
	var dispositions []Disposition
	if !arguments.headersPresent {
		dispositions = append(dispositions, Disposition{Class: "required", Member: "headers"})
	}
	if !arguments.formDigestPresent {
		dispositions = append(dispositions, Disposition{Class: "required", Member: "form_digest"})
	} else if arguments.FormDigest != digest {
		dispositions = append(dispositions, Disposition{Class: "digest-mismatch", Member: "form_digest"})
	}
	names := make([]string, 0, len(arguments.Headers))
	for name := range arguments.Headers {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		field, present := form.Fields[name]
		if !present || skipFormField(name) {
			dispositions = append(dispositions, Disposition{Class: "unknown-header", Member: name})
			continue
		}
		if (field.ConductorVolatile || field.DigestExempt) || len(field.Options) == 0 {
			continue
		}
		if !stringContains(field.Options, arguments.Headers[name]) {
			dispositions = append(dispositions, Disposition{Class: "enum", Member: name})
		}
	}
	if len(dispositions) == 0 {
		return nil
	}
	return dispositions
}

func ValidateProjectArguments(encoded json.RawMessage) []Disposition {
	return validateStaticArguments(encoded, map[string]bool{"view": false}, map[string][]string{"view": {"default", "audit", "roster"}})
}

func ValidateReadArguments(encoded json.RawMessage) []Disposition {
	return validateStaticArguments(encoded, map[string]bool{"relay_id": true}, nil)
}

func SubmitNeedsReRender(result json.RawMessage) bool {
	var outcome struct {
		State  string `json:"state"`
		Detail string `json:"detail"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil || outcome.State != "rejected" {
		return false
	}
	return bytes.Contains([]byte(outcome.Detail), []byte("form_digest")) && bytes.Contains([]byte(outcome.Detail), []byte("re-render"))
}

func ReRenderResult(original json.RawMessage) json.RawMessage {
	var outcome struct {
		RelayID  string `json:"relay_id,omitempty"`
		IntakeID string `json:"intake_id,omitempty"`
	}
	_ = json.Unmarshal(original, &outcome)
	payload := map[string]any{
		"state": "rejected",
		"violations": []map[string]string{{
			"field": "form_digest",
			"class": "re-render",
			"hint":  "form refreshed - re-read the submit tool schema and re-submit",
		}},
	}
	if outcome.RelayID != "" {
		payload["relay_id"] = outcome.RelayID
	}
	if outcome.IntakeID != "" {
		payload["intake_id"] = outcome.IntakeID
	}
	encoded, _ := json.Marshal(payload)
	return encoded
}

func DeclaredPhaseTier(arguments SubmitArguments) (string, string) {
	phase := arguments.Headers["PHASE"]
	if phase == "" {
		phase = "SITREP"
	}
	tier := arguments.Headers["CEREMONY_TIER"]
	if tier == "" {
		tier = "medium"
	}
	return phase, tier
}

func propertyForField(field fieldspec.Field) map[string]any {
	property := map[string]any{"type": "string"}
	volatile := field.ConductorVolatile || field.DigestExempt
	if len(field.Options) > 0 {
		if volatile {
			options, _ := json.Marshal(field.Options)
			property["description"] = "live options (conductor-validated; may change without a form_digest change): " + string(options)
		} else {
			property["enum"] = append([]string(nil), field.Options...)
		}
	}
	if field.Default != "" {
		property["default"] = field.Default
	}
	if _, present := property["description"]; !present {
		if description := structuredDescription(field.Type); description != "" {
			property["description"] = description
		}
	}
	return property
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

func stringContains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func validateStaticArguments(encoded json.RawMessage, members map[string]bool, enums map[string][]string) []Disposition {
	if len(encoded) == 0 {
		encoded = json.RawMessage(`{}`)
	}
	var object map[string]any
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&object); err != nil {
		return []Disposition{{Class: "shape", Member: "arguments"}}
	}
	var dispositions []Disposition
	for name, required := range members {
		value, present := object[name]
		if required && !present {
			dispositions = append(dispositions, Disposition{Class: "required", Member: name})
			continue
		}
		if present {
			text, ok := value.(string)
			if !ok {
				dispositions = append(dispositions, Disposition{Class: "type", Member: name})
			} else if options := enums[name]; len(options) > 0 && !stringContains(options, text) {
				dispositions = append(dispositions, Disposition{Class: "enum", Member: name})
			}
		}
	}
	for name := range object {
		if _, known := members[name]; !known {
			dispositions = append(dispositions, Disposition{Class: "unknown-member", Member: name})
		}
	}
	sort.Slice(dispositions, func(i, j int) bool {
		if dispositions[i].Member == dispositions[j].Member {
			return dispositions[i].Class < dispositions[j].Class
		}
		return dispositions[i].Member < dispositions[j].Member
	})
	return dispositions
}
