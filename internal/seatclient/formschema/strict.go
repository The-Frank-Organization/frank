package formschema

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type ErrorID string

const (
	ErrorUnknownMember          ErrorID = "P-1.a"
	ErrorCaseVariantMember      ErrorID = "P-1.b"
	ErrorMemberType             ErrorID = "P-1.c"
	ErrorReservedSystemHeader   ErrorID = "P-2.a"
	ErrorReservedEnvelopeHeader ErrorID = "P-2.b"
	ErrorCCConflict             ErrorID = "P-3.c"
	ErrorDuplicateTopMember     ErrorID = "P-6.a"
	ErrorDuplicateNestedMember  ErrorID = "P-6.b"
	ErrorTrailingData           ErrorID = "P-6.c"
)

type MappingError struct {
	ID     ErrorID
	Detail string
}

func (err *MappingError) Error() string {
	if err.Detail == "" {
		return "formschema: " + string(err.ID)
	}
	return fmt.Sprintf("formschema: %s: %s", err.ID, err.Detail)
}

func AsMappingError(err error, target **MappingError) bool { return errors.As(err, target) }

var submitMemberNames = map[string]struct{}{
	"headers": {}, "to": {}, "cc": {}, "dispatch_id": {}, "body": {}, "form_digest": {},
}

func ParseSubmitArguments(encoded json.RawMessage) (SubmitArguments, error) {
	if len(encoded) == 0 {
		encoded = json.RawMessage(`{}`)
	}
	if err := scanUniqueJSON(encoded); err != nil {
		return SubmitArguments{}, err
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &raw); err != nil {
		return SubmitArguments{}, &MappingError{ID: ErrorMemberType, Detail: err.Error()}
	}
	for name := range raw {
		if _, known := submitMemberNames[name]; known {
			continue
		}
		for canonical := range submitMemberNames {
			if strings.EqualFold(name, canonical) {
				return SubmitArguments{}, &MappingError{ID: ErrorCaseVariantMember, Detail: name}
			}
		}
		return SubmitArguments{}, &MappingError{ID: ErrorUnknownMember, Detail: name}
	}
	arguments := SubmitArguments{}
	if value, present := raw["headers"]; present {
		arguments.headersPresent = true
		if !isJSONObject(value) || json.Unmarshal(value, &arguments.Headers) != nil {
			return SubmitArguments{}, &MappingError{ID: ErrorMemberType, Detail: "headers"}
		}
	}
	for name, target := range map[string]*string{
		"to": &arguments.To, "cc": &arguments.CC, "dispatch_id": &arguments.DispatchID,
		"body": &arguments.Body, "form_digest": &arguments.FormDigest,
	} {
		value, present := raw[name]
		if !present {
			continue
		}
		if err := decodeJSONString(value, target); err != nil {
			return SubmitArguments{}, &MappingError{ID: ErrorMemberType, Detail: name}
		}
		if name == "form_digest" {
			arguments.formDigestPresent = true
		}
	}
	for name := range arguments.Headers {
		switch strings.ToLower(name) {
		case "from", "role", "relay_id", "delivery_state":
			return SubmitArguments{}, &MappingError{ID: ErrorReservedSystemHeader, Detail: name}
		}
		switch name {
		case "TO", "DISPATCH_ID":
			return SubmitArguments{}, &MappingError{ID: ErrorReservedEnvelopeHeader, Detail: name}
		}
	}
	if arguments.CC != "" {
		if headerCC, present := arguments.Headers["CC"]; present && headerCC != arguments.CC {
			return SubmitArguments{}, &MappingError{ID: ErrorCCConflict, Detail: "CC"}
		}
	}
	return arguments, nil
}

func scanUniqueJSON(encoded []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.UseNumber()
	if err := scanValue(decoder, 0); err != nil {
		return err
	}
	if _, err := decoder.Token(); !errors.Is(err, io.EOF) {
		return &MappingError{ID: ErrorTrailingData, Detail: "trailing JSON"}
	}
	return nil
}

func scanValue(decoder *json.Decoder, depth int) error {
	token, err := decoder.Token()
	if err != nil {
		return &MappingError{ID: ErrorMemberType, Detail: err.Error()}
	}
	delimiter, compound := token.(json.Delim)
	if !compound {
		return nil
	}
	switch delimiter {
	case '{':
		seen := make(map[string]struct{})
		for decoder.More() {
			nameToken, err := decoder.Token()
			if err != nil {
				return &MappingError{ID: ErrorMemberType, Detail: err.Error()}
			}
			name, ok := nameToken.(string)
			if !ok {
				return &MappingError{ID: ErrorMemberType, Detail: "object member name"}
			}
			if _, duplicate := seen[name]; duplicate {
				id := ErrorDuplicateNestedMember
				if depth == 0 {
					id = ErrorDuplicateTopMember
				}
				return &MappingError{ID: id, Detail: name}
			}
			seen[name] = struct{}{}
			if err := scanValue(decoder, depth+1); err != nil {
				return err
			}
		}
		if _, err := decoder.Token(); err != nil {
			return &MappingError{ID: ErrorMemberType, Detail: err.Error()}
		}
	case '[':
		for decoder.More() {
			if err := scanValue(decoder, depth+1); err != nil {
				return err
			}
		}
		if _, err := decoder.Token(); err != nil {
			return &MappingError{ID: ErrorMemberType, Detail: err.Error()}
		}
	default:
		return &MappingError{ID: ErrorMemberType, Detail: "unexpected delimiter"}
	}
	return nil
}

func decodeJSONString(encoded []byte, target *string) error {
	trimmed := bytes.TrimSpace(encoded)
	if len(trimmed) == 0 || trimmed[0] != '"' {
		return errors.New("not a string")
	}
	return json.Unmarshal(trimmed, target)
}

func isJSONObject(encoded []byte) bool {
	trimmed := bytes.TrimSpace(encoded)
	return len(trimmed) > 0 && trimmed[0] == '{'
}
