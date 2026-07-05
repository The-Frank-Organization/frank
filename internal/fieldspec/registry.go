package fieldspec

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Registry struct {
	Version    string              `json:"version,omitempty"`
	Provenance map[string]string   `json:"provenance,omitempty"`
	NamedEnums map[string][]string `json:"named_enums,omitempty"`
	Fields     []FieldSpec         `json:"fields,omitempty"`

	Phase          []string            `json:"phase,omitempty"`
	Authority      []string            `json:"authority,omitempty"`
	CeremonyTier   []string            `json:"ceremony_tier,omitempty"`
	EvidenceTarget []string            `json:"evidence_target,omitempty"`
	GateCategory   map[string][]string `json:"gate_category,omitempty"`
	Grant          []string            `json:"grant,omitempty"`

	byID map[string]*FieldSpec
}

type FieldSpec struct {
	ID                string              `json:"id"`
	Layer             string              `json:"layer"`
	Owner             string              `json:"owner"`
	Type              string              `json:"type"`
	EnumSet           string              `json:"enum_set,omitempty"`
	Options           []string            `json:"options,omitempty"`
	GateReferenceable bool                `json:"gate_referenceable,omitempty"`
	ModelIdentity     bool                `json:"model_identity,omitempty"`
	SeatScope         map[string][]string `json:"seat_scope,omitempty"`
	RequiredWhen      Predicate           `json:"required_when,omitempty"`
	VisibleWhen       Predicate           `json:"visible_when,omitempty"`
	FillConstraints   string              `json:"fill_constraints,omitempty"`
	Consumers         []string            `json:"consumers,omitempty"`
	LineageRole       string              `json:"lineage_role,omitempty"`
}

type Predicate struct {
	Raw  json.RawMessage
	root *predicateExpr
}

func (p *Predicate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		p.Raw = nil
		return nil
	}
	p.Raw = append(p.Raw[:0], data...)
	return nil
}

func (p Predicate) MarshalJSON() ([]byte, error) {
	if len(p.Raw) == 0 {
		return []byte("null"), nil
	}
	return p.Raw, nil
}

func Load(path string) (*Registry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var reg Registry
	if err := json.Unmarshal(data, &reg); err != nil {
		return nil, err
	}
	if reg.NamedEnums == nil {
		reg.populateNamedEnumsFromLegacy()
	}
	reg.populateLegacyViews()
	if err := reg.validate(); err != nil {
		return nil, err
	}
	return &reg, nil
}

func (r *Registry) ByID(id string) (*FieldSpec, bool) {
	if r.byID == nil {
		r.index()
	}
	spec, ok := r.byID[id]
	return spec, ok
}

func (r *Registry) populateNamedEnumsFromLegacy() {
	r.NamedEnums = map[string][]string{}
	if len(r.Phase) > 0 {
		r.NamedEnums["PHASE"] = append([]string(nil), r.Phase...)
	}
	if len(r.Authority) > 0 {
		r.NamedEnums["AUTHORITY"] = append([]string(nil), r.Authority...)
	}
	if len(r.CeremonyTier) > 0 {
		r.NamedEnums["CEREMONY_TIER"] = append([]string(nil), r.CeremonyTier...)
	}
	if len(r.EvidenceTarget) > 0 {
		r.NamedEnums["EVIDENCE_TARGET"] = append([]string(nil), r.EvidenceTarget...)
	}
	if len(r.GateCategory) > 0 {
		r.NamedEnums["gate_category_A"] = append([]string(nil), r.GateCategory["A"]...)
		r.NamedEnums["gate_category_B"] = append([]string(nil), r.GateCategory["B"]...)
		r.NamedEnums["gate_category"] = append(append([]string(nil), r.GateCategory["A"]...), r.GateCategory["B"]...)
	}
	if len(r.Grant) > 0 {
		r.NamedEnums["grant"] = append([]string(nil), r.Grant...)
	}
}

func (r *Registry) populateLegacyViews() {
	r.Phase = append([]string(nil), r.NamedEnums["PHASE"]...)
	r.Authority = append([]string(nil), r.NamedEnums["AUTHORITY"]...)
	r.CeremonyTier = append([]string(nil), r.NamedEnums["CEREMONY_TIER"]...)
	r.EvidenceTarget = append([]string(nil), r.NamedEnums["EVIDENCE_TARGET"]...)
	r.GateCategory = map[string][]string{
		"A": append([]string(nil), r.NamedEnums["gate_category_A"]...),
		"B": append([]string(nil), r.NamedEnums["gate_category_B"]...),
	}
	r.Grant = append([]string(nil), r.NamedEnums["grant"]...)
	r.index()
}

func (r *Registry) index() {
	r.byID = make(map[string]*FieldSpec, len(r.Fields))
	for i := range r.Fields {
		r.byID[r.Fields[i].ID] = &r.Fields[i]
	}
}

func (r *Registry) validate() error {
	if len(r.Fields) == 0 {
		return nil
	}
	for i := range r.Fields {
		field := &r.Fields[i]
		if field.ID == "" {
			return fmt.Errorf("field <empty>: id required")
		}
		if !known(layers, field.Layer) {
			return fmt.Errorf("field %s: unknown layer %s", field.ID, field.Layer)
		}
		if !known(owners, field.Owner) {
			return fmt.Errorf("field %s: unknown owner %s", field.ID, field.Owner)
		}
		if !known(types, field.Type) {
			return fmt.Errorf("field %s: unknown type %s", field.ID, field.Type)
		}
		if field.FillConstraints != "" && !known(fillConstraints, field.FillConstraints) {
			return fmt.Errorf("field %s: unknown fill_constraints %s", field.ID, field.FillConstraints)
		}
		if field.LineageRole == "" {
			field.LineageRole = "none"
		}
		if !known(lineageRoles, field.LineageRole) {
			return fmt.Errorf("field %s: unknown lineage_role %s", field.ID, field.LineageRole)
		}
		if field.ModelIdentity && field.GateReferenceable {
			return fmt.Errorf("field %s: model_identity requires gate_referenceable false", field.ID)
		}
		if strings.HasPrefix(field.ID, "X-") || field.ID == "X-*" {
			if len(field.Consumers) > 0 || field.LineageRole != "none" {
				return fmt.Errorf("field %s: X-* rows must have no consumers and lineage_role none", field.ID)
			}
		}
		if field.EnumSet != "" {
			if _, ok := r.NamedEnums[field.EnumSet]; !ok {
				return fmt.Errorf("field %s: enum_set %s missing", field.ID, field.EnumSet)
			}
		}
		if err := r.compilePredicate(field.ID, &field.RequiredWhen); err != nil {
			return err
		}
		if err := r.compilePredicate(field.ID, &field.VisibleWhen); err != nil {
			return err
		}
	}
	return nil
}

func (r *Registry) compilePredicate(owner string, pred *Predicate) error {
	if pred == nil || len(pred.Raw) == 0 || string(pred.Raw) == "null" {
		return nil
	}
	parsed, err := parsePredicate(owner, pred.Raw, r)
	if err != nil {
		return err
	}
	pred.root = parsed.root
	return nil
}

func (r *Registry) validateGateReference(owner, ref string) error {
	if strings.HasPrefix(ref, "X-") || ref == "X-*" {
		return fmt.Errorf("field %s: X-* fields are not gate-referenceable", owner)
	}
	field, ok := r.ByID(ref)
	if !ok {
		return fmt.Errorf("field %s: predicate references unknown field %s", owner, ref)
	}
	if !field.GateReferenceable {
		return fmt.Errorf("field %s: predicate references non gate-referenceable field %s", owner, ref)
	}
	return nil
}

func known(set map[string]struct{}, value string) bool {
	_, ok := set[value]
	return ok
}

var layers = map[string]struct{}{
	"envelope": {},
	"header":   {},
	"body":     {},
	"system":   {},
}

var owners = map[string]struct{}{
	"system":           {},
	"courier":          {},
	"agent_enum_pick":  {},
	"seat_scoped_enum": {},
	"free_text":        {},
	"computed":         {},
	"body_text":        {},
}

var types = map[string]struct{}{
	"enum":         {},
	"bool":         {},
	"int":          {},
	"string":       {},
	"id_ref":       {},
	"evidence_ref": {},
	"text":         {},
	"row_array":    {},
	"address_list": {},
	"object":       {},
}

var fillConstraints = map[string]struct{}{
	"none":                {},
	"monotonic":           {},
	"system_only":         {},
	"seat_allowed_values": {},
	"computed_result":     {},
	"observed_value":      {},
	"parent_picker":       {},
	"recipient_picker":    {},
	"free_text":           {},
}

var lineageRoles = map[string]struct{}{
	"none":          {},
	"parent_edge":   {},
	"grant":         {},
	"verdict":       {},
	"lock_id":       {},
	"action_report": {},
	"merge_claim":   {},
	"routing_ref":   {},
}

var ops = map[string]struct{}{
	"==":      {},
	"in":      {},
	"present": {},
}
