package fieldspec

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/jackli/frank/internal/record"
)

type RenderEnv struct {
	ConfigDigest        string
	KnownA              KnownADetector
	ParentCandidates    ParentCandidates
	RecipientCandidates RecipientCandidates
	MonotonicFloors     map[string]string
	Turn                TurnContext
	PreActive           bool
}

type TurnContext struct {
	WokenOn        string
	ActiveDispatch string
}

type GrantState func(seat SeatMeta) bool

func ClosedGrantState(SeatMeta) bool { return false }

type ParentCandidates func(seat SeatMeta) (candidates []string, dflt string)

func EmptyParentCandidates(SeatMeta) ([]string, string) { return nil, "" }

type RecipientCandidates func(seat SeatMeta) []string

func EmptyRecipientCandidates(SeatMeta) []string { return nil }

type KnownADetector func(cand record.Record, fields map[string]string) (member string, hit bool)

func (r *Registry) Render(env RenderEnv, seat SeatMeta, phase, tier string, grants GrantState) (Form, string) {
	if grants == nil {
		grants = ClosedGrantState
	}
	if env.ParentCandidates == nil {
		env.ParentCandidates = EmptyParentCandidates
	}
	if env.RecipientCandidates == nil {
		env.RecipientCandidates = EmptyRecipientCandidates
	}
	fields := map[string]string{
		"PHASE":         phase,
		"CEREMONY_TIER": tier,
	}
	ctx := EvalContext{Seat: seat, Phase: phase, Tier: tier, PresentLayers: DefaultLayers()}
	form := Form{Fields: map[string]Field{}}
	for i := range r.Fields {
		spec := &r.Fields[i]
		if !r.renderable(spec, fields, ctx, env) {
			continue
		}
		field, ok := r.renderField(env, spec, seat, phase, grants)
		if !ok {
			continue
		}
		form.Fields[spec.ID] = field
	}
	if env.PreActive {
		form = bootForm(form)
	}
	return form, r.digestRenderedForm(form, env, seat, phase, tier)
}

func bootForm(form Form) Form {
	out := Form{Fields: map[string]Field{}}
	for _, id := range []string{"PHASE", "CEREMONY_TIER", "SUBJECT", "charter_loaded", "dispatch_status"} {
		field, ok := form.Fields[id]
		if !ok {
			continue
		}
		if id == "PHASE" {
			field.Options = []string{"SITREP"}
		}
		out.Fields[id] = field
	}
	return out
}

func (r *Registry) renderable(spec *FieldSpec, fields map[string]string, ctx EvalContext, env RenderEnv) bool {
	if spec.Owner == "system" || spec.Owner == "computed" || spec.FillConstraints == "system_only" || spec.FillConstraints == "computed_result" {
		return false
	}
	if !env.PreActive && (spec.ID == "charter_loaded" || spec.ID == "dispatch_status") {
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
		markConductorVolatile(&field)
	case "recipient_picker":
		field.Options = append([]string(nil), env.RecipientCandidates(seat)...)
		markConductorVolatile(&field)
	case "seat_allowed_values":
		field.Options = r.optionsForSeat(spec, seat, phase, grants)
		if spec.ID == "grant" {
			markConductorVolatile(&field)
		}
	case "monotonic":
		field.Options = r.monotonicOptions(spec, env.MonotonicFloors[spec.ID])
		markConductorVolatile(&field)
	default:
		if spec.Owner == "seat_scoped_enum" {
			field.Options = r.optionsForSeat(spec, seat, phase, grants)
			if spec.ID == "grant" {
				markConductorVolatile(&field)
			}
		} else {
			field.Options = r.baseOptions(spec)
		}
	}
	if spec.ID == "grant" && len(field.Options) == 0 {
		return Field{}, false
	}
	return field, true
}

func markConductorVolatile(field *Field) {
	field.ConductorVolatile = true
	field.DigestExempt = true
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
		Form:         r.formForDigest(form, seat, phase),
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

func (r *Registry) formForDigest(form Form, seat SeatMeta, phase string) Form {
	out := Form{Fields: map[string]Field{}}
	for name, field := range form.Fields {
		if field.DigestExempt || field.ConductorVolatile {
			field.Options = nil
			field.Default = ""
		}
		out.Fields[name] = field
	}
	r.addGrantDigestShape(out.Fields, seat, phase)
	return out
}

func (r *Registry) addGrantDigestShape(fields map[string]Field, seat SeatMeta, phase string) {
	if _, ok := fields["grant"]; ok {
		return
	}
	spec, ok := r.ByID("grant")
	if !ok || !r.renderable(spec, map[string]string{"PHASE": phase}, EvalContext{Seat: seat, Phase: phase, PresentLayers: DefaultLayers()}, RenderEnv{}) {
		return
	}
	options := r.scopeOptions(spec, seat)
	if len(options) == 0 {
		return
	}
	if phase != "MERGE-GATE" {
		options = removeString(options, "dispatch-merge")
	}
	if len(options) == 0 {
		return
	}
	field := Field{Type: spec.Type}
	markConductorVolatile(&field)
	fields["grant"] = field
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
