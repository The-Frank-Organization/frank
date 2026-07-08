package fieldspec

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type TypedValue any

func CanonicalMarshal(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ParseTyped(spec *FieldSpec, raw string) (TypedValue, error) {
	if spec == nil {
		return nil, fmt.Errorf("field <nil>: spec required")
	}
	switch spec.Type {
	case "row_array":
		var rows []map[string]string
		if err := parseCanonical(spec, raw, &rows); err != nil {
			return nil, err
		}
		return rows, nil
	case "object":
		var obj map[string]string
		if err := parseCanonical(spec, raw, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "address_list":
		var list []string
		if err := parseCanonical(spec, raw, &list); err != nil {
			return nil, err
		}
		return list, nil
	case "bool":
		if raw != "yes" && raw != "no" {
			return nil, fmt.Errorf("field %s: bool must be yes or no", spec.ID)
		}
		return raw, nil
	case "int":
		value, err := strconv.Atoi(raw)
		if err != nil {
			return nil, fmt.Errorf("field %s: int parse: %w", spec.ID, err)
		}
		return value, nil
	case "enum", "string", "id_ref", "evidence_ref", "text":
		if raw == "" {
			return nil, fmt.Errorf("field %s: scalar value required", spec.ID)
		}
		return raw, nil
	default:
		return nil, fmt.Errorf("field %s: unknown type %s", spec.ID, spec.Type)
	}
}

func DecodeAddressList(raw string) ([]string, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	value, err := ParseTyped(&FieldSpec{ID: "address_list", Type: "address_list"}, raw)
	if err != nil {
		return nil, err
	}
	list, ok := value.([]string)
	if !ok {
		return nil, fmt.Errorf("field address_list: typed value is %T", value)
	}
	return list, nil
}

func EncodeAddressList(list []string) (string, error) {
	if list == nil {
		list = []string{}
	}
	return CanonicalMarshal(list)
}

func parseCanonical(spec *FieldSpec, raw string, out any) error {
	if err := json.Unmarshal([]byte(raw), out); err != nil {
		return fmt.Errorf("field %s: typed parse: %w", spec.ID, err)
	}
	canonical, err := CanonicalMarshal(out)
	if err != nil {
		return fmt.Errorf("field %s: canonical marshal: %w", spec.ID, err)
	}
	if canonical != raw {
		return fmt.Errorf("field %s: canonical-encoding", spec.ID)
	}
	return nil
}
