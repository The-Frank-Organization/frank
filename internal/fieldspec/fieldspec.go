package fieldspec

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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
	Options []string `json:"options,omitempty"`
}

type Violation struct {
	Field  string
	Class  string
	Reason string
}

func (r *Registry) Render(seat SeatMeta, phase, tier string) (Form, string) {
	authority := append([]string(nil), r.Authority...)
	if !canGrant(seat) {
		authority = remove(authority, "merge-gated")
	}
	form := Form{Fields: map[string]Field{
		"PHASE":               {Options: r.Phase},
		"AUTHORITY":           {Options: authority},
		"CEREMONY_TIER":       {Options: r.CeremonyTier},
		"EVIDENCE_TARGET":     {Options: r.EvidenceTarget},
		"HUMAN_GATE_REQUIRED": {Options: []string{"no", "yes"}},
		"gate_category":       {Options: append(append([]string(nil), r.GateCategory["A"]...), r.GateCategory["B"]...)},
		"SUBJECT":             {},
		"body":                {},
		"X-*":                 {},
	}}
	if canGrant(seat) {
		grantOptions := []string{"dispatch-impl"}
		if phase == "MERGE-GATE" {
			grantOptions = append(grantOptions, "dispatch-merge")
		}
		form.Fields["grant"] = Field{Options: grantOptions}
	}
	digest := digestForm(form, seat, phase, tier)
	return form, digest
}

func (r *Registry) Validate(cand record.Record, seat SeatMeta, formDigest string) []Violation {
	var violations []Violation
	phase := cand.Headers["PHASE"]
	authority := cand.Headers["AUTHORITY"]
	subject := cand.Headers["SUBJECT"]
	_, digest := r.Render(seat, phase, cand.Headers["CEREMONY_TIER"])
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
	if authority == "merge-gated" && !canGrant(seat) {
		violations = append(violations, Violation{Field: "AUTHORITY", Class: "seat-scope", Reason: "merge-gated absent from pair forms"})
	}
	if cand.Headers["grant"] != "" && !canGrant(seat) {
		violations = append(violations, Violation{Field: "grant", Class: "seat-scope", Reason: "grant absent from pair forms"})
	}
	return violations
}

func canGrant(seat SeatMeta) bool {
	return seat.IsOperator || seat.Role == "operator" || seat.Role == "orchestrator-planner" || seat.Role == "orchestrator-reviewer"
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

func digestForm(form Form, seat SeatMeta, phase, tier string) string {
	data, err := json.Marshal(struct {
		Form  Form
		Seat  SeatMeta
		Phase string
		Tier  string
	}{form, seat, phase, tier})
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func remove(values []string, value string) []string {
	var out []string
	for _, candidate := range values {
		if candidate != value {
			out = append(out, candidate)
		}
	}
	return out
}
