package fieldspec

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/jackli/frank/internal/record"
)

func (r *Registry) Validate(cand record.Record, seat SeatMeta, formDigest string, env RenderEnv, grants GrantState) []Violation {
	if grants == nil {
		grants = ClosedGrantState
	}
	if cand.Headers == nil {
		cand.Headers = map[string]string{}
	}
	var violations []Violation
	phase := cand.Headers["PHASE"]
	tier := cand.Headers["CEREMONY_TIER"]
	_, currentDigest := r.Render(env, seat, phase, tier, grants)
	if formDigest == "" || formDigest != currentDigest {
		violations = append(violations, Violation{Field: "form_digest", Class: "re-render", Reason: "stale form digest"})
	}

	fields := r.evalFields(cand)
	rows := r.rowValues(cand)
	ctx := EvalContext{Seat: seat, Phase: phase, Tier: tier, PresentLayers: DefaultLayers()}

	for i := range r.Fields {
		spec := &r.Fields[i]
		raw, present := valueForSpec(cand, spec)
		if present && raw != "" && r.systemOwnedHeader(spec) {
			violations = append(violations, Violation{Field: spec.ID, Class: "system-owned", Reason: spec.ID + " is system-owned"})
			continue
		}
		if r.ignorePayloadField(spec) {
			continue
		}
		if len(spec.RequiredWhen.Raw) > 0 && spec.RequiredWhen.Eval(fields, rows, ctx) && raw == "" {
			violations = append(violations, Violation{Field: spec.ID, Class: "required", Reason: spec.ID + " required"})
		}
		if spec.ID == "SUBJECT" && raw == "" {
			violations = append(violations, Violation{Field: "SUBJECT", Class: "required", Reason: "SUBJECT required"})
		}
		if !present || raw == "" {
			continue
		}
		if err := validateTyped(spec, raw); err != nil {
			violations = append(violations, Violation{Field: spec.ID, Class: violationClass(err), Reason: err.Error()})
			continue
		}
		if spec.Type == "enum" || spec.Type == "bool" {
			if !containsString(r.enumTokens(spec), raw) {
				violations = append(violations, Violation{Field: spec.ID, Class: "enum", Reason: "unknown " + spec.ID})
			}
		}
		if len(spec.SeatScope) > 0 && !r.optionAllowedForSeat(spec.ID, seat, phase, raw, grants) {
			violations = append(violations, Violation{Field: spec.ID, Class: "seat-scope", Reason: spec.ID + " absent from seat form"})
		}
		if spec.FillConstraints == "monotonic" && r.belowMonotonicFloor(spec, raw, env.MonotonicFloors[spec.ID]) {
			violations = append(violations, Violation{Field: spec.ID, Class: "monotonic-floor", Reason: spec.ID + " below floor"})
		}
		if spec.ID == "gate_category" {
			_, raised := r.ClassifyGateCategory(raw, false)
			if raised {
				cand.Headers["gate_category_raised"] = "true"
			}
		}
	}
	if parent := cand.Headers["PARENT_DISPATCH_ID"]; parent != "" && env.ParentCandidates != nil {
		candidates, _ := env.ParentCandidates(seat)
		if len(candidates) == 0 {
			violations = append(violations, Violation{Field: "PARENT_DISPATCH_ID", Class: "active-lineage-empty", Reason: "no active lineage candidates"})
		} else if !containsString(candidates, parent) {
			violations = append(violations, Violation{Field: "PARENT_DISPATCH_ID", Class: "outside-active-lineage", Reason: "parent outside active lineage candidate set"})
		}
	}
	return violations
}

func (r *Registry) belowMonotonicFloor(spec *FieldSpec, raw, floor string) bool {
	if floor == "" || raw == "" {
		return false
	}
	options := r.baseOptions(spec)
	rawIndex, floorIndex := -1, -1
	for i, option := range options {
		if option == raw {
			rawIndex = i
		}
		if option == floor {
			floorIndex = i
		}
	}
	return rawIndex >= 0 && floorIndex >= 0 && rawIndex < floorIndex
}

func validateTyped(spec *FieldSpec, raw string) error {
	switch spec.Type {
	case "row_array", "object", "address_list", "int":
		_, err := ParseTyped(spec, raw)
		return err
	case "enum", "bool", "string", "id_ref", "evidence_ref", "text":
		return nil
	default:
		return fmt.Errorf("field %s: unknown type %s", spec.ID, spec.Type)
	}
}

func violationClass(err error) string {
	if strings.Contains(err.Error(), "canonical-encoding") {
		return "canonical-encoding"
	}
	return "typed"
}

func (r *Registry) ignorePayloadField(spec *FieldSpec) bool {
	if spec.ID == "record_kind" {
		return false
	}
	return spec.Owner == "system" || spec.FillConstraints == "system_only"
}

func (r *Registry) systemOwnedHeader(spec *FieldSpec) bool {
	if spec.Layer != "header" {
		return false
	}
	return spec.Owner == "system" ||
		spec.Owner == "computed" ||
		spec.FillConstraints == "system_only" ||
		spec.FillConstraints == "computed_result"
}

func (r *Registry) enumTokens(spec *FieldSpec) []string {
	if spec.EnumSet != "" {
		return r.NamedEnums[spec.EnumSet]
	}
	return spec.Options
}

func (r *Registry) evalFields(cand record.Record) map[string]string {
	fields := make(map[string]string, len(cand.Headers)+1)
	for key, value := range cand.Headers {
		fields[key] = value
	}
	fields["body"] = cand.Body
	return fields
}

func (r *Registry) rowValues(cand record.Record) func(array, field string) []string {
	return func(array, field string) []string {
		spec, ok := r.ByID(array)
		if !ok {
			return nil
		}
		raw := cand.Headers[array]
		if raw == "" {
			return nil
		}
		value, err := ParseTyped(spec, raw)
		if err != nil {
			return nil
		}
		rows, ok := value.([]map[string]string)
		if !ok {
			return nil
		}
		values := make([]string, 0, len(rows))
		for _, row := range rows {
			if value := row[field]; value != "" {
				values = append(values, value)
			}
		}
		return values
	}
}

func valueForSpec(cand record.Record, spec *FieldSpec) (string, bool) {
	switch spec.ID {
	case "body":
		return cand.Body, cand.Body != ""
	default:
		value, ok := cand.Headers[spec.ID]
		return value, ok
	}
}

func (r *Registry) ClassifyGateCategory(token string, knownA bool) (string, bool) {
	if knownA {
		return "A", true
	}
	if slices.Contains(r.GateCategory["A"], token) {
		return "A", false
	}
	if slices.Contains(r.GateCategory["B"], token) {
		return "B", false
	}
	return "A", true
}

func (r *Registry) GateCategoryAuthorityBearing(token string) bool {
	if token == "" {
		return false
	}
	if token == "other" {
		return true
	}
	return slices.Contains(r.GateCategory["A"], token)
}

type SubmitPayload struct {
	record.Record
	FormDigest string `json:"form_digest,omitempty"`
}

func DecodeSubmitPayload(data []byte) (record.Record, string, error) {
	var payload SubmitPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return record.Record{}, "", err
	}
	return payload.Record, payload.FormDigest, nil
}
