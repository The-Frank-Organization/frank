package fieldspec

import (
	"slices"

	"github.com/jackli/frank/internal/record"
)

type SeatMeta struct {
	Name       string
	Role       string
	IsOperator bool
}

type Form struct {
	Fields map[string]Field `json:"fields"`
}

type Field struct {
	Type         string   `json:"type,omitempty"`
	Options      []string `json:"options,omitempty"`
	Default      string   `json:"default,omitempty"`
	DigestExempt bool     `json:"digest_exempt,omitempty"`
}

type Violation struct {
	Field  string
	Class  string
	Reason string
}

func (r *Registry) Validate(cand record.Record, seat SeatMeta, formDigest string) []Violation {
	var violations []Violation
	phase := cand.Headers["PHASE"]
	authority := cand.Headers["AUTHORITY"]
	subject := cand.Headers["SUBJECT"]
	_, digest := r.Render(RenderEnv{}, seat, phase, cand.Headers["CEREMONY_TIER"], ClosedGrantState)
	if formDigest != "" && formDigest != digest {
		violations = append(violations, Violation{Field: "form_digest", Class: "re-render", Reason: "stale form digest"})
	}
	if subject == "" {
		violations = append(violations, Violation{Field: "SUBJECT", Class: "required", Reason: "SUBJECT required"})
	}
	if !slices.Contains(r.Phase, phase) {
		violations = append(violations, Violation{Field: "PHASE", Class: "enum", Reason: "unknown PHASE"})
	}
	if authority != "" && !slices.Contains(r.Authority, authority) {
		violations = append(violations, Violation{Field: "AUTHORITY", Class: "enum", Reason: "unknown AUTHORITY"})
	}
	if authority == "merge-gated" && !r.optionAllowedForSeat("AUTHORITY", seat, phase, authority, ClosedGrantState) {
		violations = append(violations, Violation{Field: "AUTHORITY", Class: "seat-scope", Reason: "merge-gated absent from pair forms"})
	}
	if cand.Headers["grant"] != "" && !r.optionAllowedForSeat("grant", seat, phase, cand.Headers["grant"], ClosedGrantState) {
		violations = append(violations, Violation{Field: "grant", Class: "seat-scope", Reason: "grant absent from pair forms"})
	}
	return violations
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

func (f Form) HasField(name string) bool {
	_, ok := f.Fields[name]
	return ok
}

func (f Form) OptionAllowed(field, option string) bool {
	spec, ok := f.Fields[field]
	if !ok {
		return false
	}
	return slices.Contains(spec.Options, option)
}
