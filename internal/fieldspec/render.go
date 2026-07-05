package fieldspec

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type RenderEnv struct {
	ConfigDigest     string
	ParentCandidates ParentCandidates
	MonotonicFloors  map[string]string
}

type GrantState func(seat SeatMeta) bool

func ClosedGrantState(SeatMeta) bool { return false }

type ParentCandidates func(seat SeatMeta) (candidates []string, dflt string)

func EmptyParentCandidates(SeatMeta) ([]string, string) { return nil, "" }

func (r *Registry) Render(env RenderEnv, seat SeatMeta, phase, tier string, grants GrantState) (Form, string) {
	if grants == nil {
		grants = ClosedGrantState
	}
	if env.ParentCandidates == nil {
		env.ParentCandidates = EmptyParentCandidates
	}
	fields := map[string]string{
		"PHASE":         phase,
		"CEREMONY_TIER": tier,
	}
	ctx := EvalContext{Seat: seat, Phase: phase, Tier: tier, PresentLayers: DefaultLayers()}
	form := Form{Fields: map[string]Field{}}
	for i := range r.Fields {
		spec := &r.Fields[i]
		if !r.renderable(spec, fields, ctx) {
			continue
		}
		field, ok := r.renderField(env, spec, seat, phase, grants)
		if !ok {
			continue
		}
		form.Fields[spec.ID] = field
	}
	return form, r.digestRenderedForm(form, env, seat, phase, tier)
}

func (r *Registry) renderable(spec *FieldSpec, fields map[string]string, ctx EvalContext) bool {
	if spec.Owner == "system" || spec.Owner == "computed" || spec.FillConstraints == "system_only" || spec.FillConstraints == "computed_result" {
		return false
	}
	if len(spec.VisibleWhen.Raw) > 0 && !spec.VisibleWhen.Eval(fields, nil, ctx) {
		return false
	}
	return true
}

func (r *Registry) renderField(env RenderEnv, spec *FieldSpec, seat SeatMeta, phase string, grants GrantState) (Field, bool) {
	field := Field{Type: spec.Type}
	switch spec.FillConstraints {
	case "parent_picker":
		candidates, dflt := env.ParentCandidates(seat)
		field.Options = append([]string(nil), candidates...)
		field.Default = dflt
		field.DigestExempt = true
	case "seat_allowed_values":
		field.Options = r.optionsForSeat(spec, seat, phase, grants)
	case "monotonic":
		field.Options = r.monotonicOptions(spec, env.MonotonicFloors[spec.ID])
	default:
		if spec.Owner == "seat_scoped_enum" {
			field.Options = r.optionsForSeat(spec, seat, phase, grants)
		} else {
			field.Options = r.baseOptions(spec)
		}
	}
	if spec.ID == "grant" && len(field.Options) == 0 {
		return Field{}, false
	}
	return field, true
}

func (r *Registry) monotonicOptions(spec *FieldSpec, floor string) []string {
	options := r.baseOptions(spec)
	if floor == "" {
		return options
	}
	for i, option := range options {
		if option == floor {
			return append([]string(nil), options[i:]...)
		}
	}
	return options
}

func (r *Registry) baseOptions(spec *FieldSpec) []string {
	if len(spec.Options) > 0 {
		return append([]string(nil), spec.Options...)
	}
	if spec.EnumSet != "" {
		return append([]string(nil), r.NamedEnums[spec.EnumSet]...)
	}
	return nil
}

func (r *Registry) optionsForSeat(spec *FieldSpec, seat SeatMeta, phase string, grants GrantState) []string {
	options := r.scopeOptions(spec, seat)
	if spec.ID == "grant" {
		if len(options) == 0 {
			return nil
		}
		if isPairPlanner(seat) && !grants(seat) {
			return nil
		}
		if phase != "MERGE-GATE" {
			options = removeString(options, "dispatch-merge")
		}
	}
	return options
}

func (r *Registry) scopeOptions(spec *FieldSpec, seat SeatMeta) []string {
	if len(spec.SeatScope) == 0 {
		return r.baseOptions(spec)
	}
	for _, key := range []string{seat.Name, seat.Role, "operator"} {
		if key == "operator" && !(seat.IsOperator || seat.Role == "operator" || seat.Name == "operator") {
			continue
		}
		if values, ok := spec.SeatScope[key]; ok {
			return append([]string(nil), values...)
		}
	}
	for pattern, values := range spec.SeatScope {
		if pattern == "*" || !strings.HasPrefix(pattern, "*.") {
			continue
		}
		suffix := strings.TrimPrefix(pattern, "*")
		role := strings.TrimPrefix(pattern, "*.")
		if strings.HasSuffix(seat.Name, suffix) || seat.Role == role {
			return append([]string(nil), values...)
		}
	}
	if values, ok := spec.SeatScope["*"]; ok {
		return append([]string(nil), values...)
	}
	return nil
}

func isPairPlanner(seat SeatMeta) bool {
	return seat.Role == "planner" || strings.HasSuffix(seat.Name, ".planner")
}

func (r *Registry) optionAllowedForSeat(fieldID string, seat SeatMeta, phase, option string, grants GrantState) bool {
	spec, ok := r.ByID(fieldID)
	if !ok {
		return false
	}
	return containsString(r.optionsForSeat(spec, seat, phase, grants), option)
}

func (r *Registry) digestRenderedForm(form Form, env RenderEnv, seat SeatMeta, phase, tier string) string {
	canonical, err := CanonicalMarshal(struct {
		Form         Form   `json:"form"`
		ConfigDigest string `json:"config_digest"`
		SeatPattern  string `json:"seat_pattern"`
		Phase        string `json:"phase"`
		Tier         string `json:"tier"`
	}{
		Form:         formForDigest(form),
		ConfigDigest: env.ConfigDigest,
		SeatPattern:  seatDigestKey(seat),
		Phase:        phase,
		Tier:         tier,
	})
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(sum[:])
}

func formForDigest(form Form) Form {
	out := Form{Fields: map[string]Field{}}
	for name, field := range form.Fields {
		if field.DigestExempt {
			field.Options = nil
			field.Default = ""
		}
		out.Fields[name] = field
	}
	return out
}

func seatDigestKey(seat SeatMeta) string {
	if seat.IsOperator {
		return "operator:" + seat.Name
	}
	return seat.Role + ":" + seat.Name
}

func removeString(values []string, value string) []string {
	var out []string
	for _, candidate := range values {
		if candidate != value {
			out = append(out, candidate)
		}
	}
	return out
}
