package fieldspec

import (
	"encoding/json"
	"fmt"
	"strings"
)

type EvalContext struct {
	Seat          SeatMeta
	Phase         string
	Tier          string
	Slot          string
	PresentLayers map[string]bool
}

func DefaultLayers() map[string]bool {
	return map[string]bool{
		"store":   true,
		"form":    true,
		"lineage": true,
	}
}

type predicateKind string

const (
	predicateAllOf       predicateKind = "all_of"
	predicateAnyOf       predicateKind = "any_of"
	predicateNot         predicateKind = "not"
	predicatePhaseIn     predicateKind = "phase_in"
	predicateTierGTE     predicateKind = "ceremony_tier_gte"
	predicateAuthorityIn predicateKind = "authority_in"
	predicateSeatIs      predicateKind = "seat_is"
	predicateRoleIn      predicateKind = "role_in"
	predicateSlotIn      predicateKind = "slot_in"
	predicateRecordKind  predicateKind = "record_kind_in"
	predicateScanResult  predicateKind = "scan_result_in"
	predicateClaims      predicateKind = "claims_action"
	predicateField       predicateKind = "field"
	predicateAnyRow      predicateKind = "any_row"
	predicateLayer       predicateKind = "layer_present"
)

type predicateExpr struct {
	kind     predicateKind
	children []*predicateExpr
	child    *predicateExpr
	field    string
	array    string
	rowField string
	op       string
	values   []string
}

func ParsePredicate(raw json.RawMessage, reg *Registry) (*Predicate, error) {
	return parsePredicate("predicate", raw, reg)
}

func parsePredicate(owner string, raw json.RawMessage, reg *Registry) (*Predicate, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return &Predicate{}, nil
	}
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		return nil, fmt.Errorf("field %s: predicate parse: %w", owner, err)
	}
	root, err := parsePredicateValue(owner, value, reg)
	if err != nil {
		return nil, err
	}
	copied := append(json.RawMessage(nil), raw...)
	return &Predicate{Raw: copied, root: root}, nil
}

func parsePredicateValue(owner string, value any, reg *Registry) (*predicateExpr, error) {
	obj, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("field %s: predicate must be object", owner)
	}

	var expr *predicateExpr
	setExpr := func(next *predicateExpr) error {
		if expr != nil {
			return fmt.Errorf("field %s: predicate must have one atom", owner)
		}
		expr = next
		return nil
	}

	for key, raw := range obj {
		switch key {
		case "all_of", "any_of":
			items, ok := raw.([]any)
			if !ok {
				return nil, fmt.Errorf("field %s: %s must be array", owner, key)
			}
			children := make([]*predicateExpr, 0, len(items))
			for _, item := range items {
				child, err := parsePredicateValue(owner, item, reg)
				if err != nil {
					return nil, err
				}
				children = append(children, child)
			}
			kind := predicateAllOf
			if key == "any_of" {
				kind = predicateAnyOf
			}
			if err := setExpr(&predicateExpr{kind: kind, children: children}); err != nil {
				return nil, err
			}
		case "not":
			child, err := parsePredicateValue(owner, raw, reg)
			if err != nil {
				return nil, err
			}
			if err := setExpr(&predicateExpr{kind: predicateNot, child: child}); err != nil {
				return nil, err
			}
		case "field":
			ref, ok := raw.(string)
			if !ok {
				return nil, fmt.Errorf("field %s: field atom must name a column", owner)
			}
			if err := validateGateReference(reg, owner, ref); err != nil {
				return nil, err
			}
			op, values, err := parsePredicateOp(owner, obj)
			if err != nil {
				return nil, err
			}
			if err := setExpr(&predicateExpr{kind: predicateField, field: ref, op: op, values: values}); err != nil {
				return nil, err
			}
		case "any_row":
			ref, ok := raw.(string)
			if !ok {
				return nil, fmt.Errorf("field %s: any_row atom must name a row field", owner)
			}
			array, rowField, ok := strings.Cut(ref, ".")
			if !ok || array == "" || rowField == "" {
				return nil, fmt.Errorf("field %s: any_row atom must use ARRAY.field", owner)
			}
			if err := validateGateReference(reg, owner, array); err != nil {
				return nil, err
			}
			op, values, err := parsePredicateOp(owner, obj)
			if err != nil {
				return nil, err
			}
			if err := setExpr(&predicateExpr{kind: predicateAnyRow, array: array, rowField: rowField, op: op, values: values}); err != nil {
				return nil, err
			}
		case "slot_in":
			values, err := stringList(raw)
			if err != nil {
				return nil, fmt.Errorf("field %s: slot_in must be string array", owner)
			}
			if len(values) > 0 {
				return nil, fmt.Errorf("field %s: slot_in is reserved for Step-1", owner)
			}
			if err := setExpr(&predicateExpr{kind: predicateSlotIn, values: values}); err != nil {
				return nil, err
			}
		case "phase_in", "authority_in", "seat_is", "role_in", "record_kind_in", "scan_result_in":
			values, err := stringList(raw)
			if err != nil {
				return nil, fmt.Errorf("field %s: %s must be string array", owner, key)
			}
			if err := setExpr(&predicateExpr{kind: predicateKind(key), values: values}); err != nil {
				return nil, err
			}
		case "ceremony_tier_gte":
			value, ok := raw.(string)
			if !ok {
				return nil, fmt.Errorf("field %s: ceremony_tier_gte must be string", owner)
			}
			if _, ok := tierOrder[value]; !ok {
				return nil, fmt.Errorf("field %s: ceremony_tier_gte unknown tier %s", owner, value)
			}
			if err := setExpr(&predicateExpr{kind: predicateTierGTE, values: []string{value}}); err != nil {
				return nil, err
			}
		case "claims_action":
			if rawBool, ok := raw.(bool); ok && !rawBool {
				return nil, fmt.Errorf("field %s: claims_action false is not a predicate atom", owner)
			}
			if err := setExpr(&predicateExpr{kind: predicateClaims}); err != nil {
				return nil, err
			}
		case "layer_present":
			value, ok := raw.(string)
			if !ok {
				return nil, fmt.Errorf("field %s: layer_present must be string", owner)
			}
			if !known(predicateLayers, value) {
				return nil, fmt.Errorf("field %s: unknown layer_present %s", owner, value)
			}
			if err := setExpr(&predicateExpr{kind: predicateLayer, values: []string{value}}); err != nil {
				return nil, err
			}
		case "op", "value":
			continue
		default:
			return nil, fmt.Errorf("field %s: unknown predicate atom %s", owner, key)
		}
	}
	if expr == nil {
		return nil, fmt.Errorf("field %s: unknown predicate atom", owner)
	}
	return expr, nil
}

func parsePredicateOp(owner string, obj map[string]any) (string, []string, error) {
	raw, ok := obj["op"]
	if !ok {
		return "", nil, fmt.Errorf("field %s: predicate op required", owner)
	}
	op, ok := raw.(string)
	if !ok || !known(ops, op) {
		return "", nil, fmt.Errorf("field %s: unknown predicate op %v", owner, raw)
	}
	if op == "present" {
		return op, nil, nil
	}
	rawValue, ok := obj["value"]
	if !ok {
		return "", nil, fmt.Errorf("field %s: predicate value required", owner)
	}
	values, err := stringList(rawValue)
	if err != nil {
		return "", nil, fmt.Errorf("field %s: predicate value must be string or string array", owner)
	}
	if op == "==" && len(values) != 1 {
		return "", nil, fmt.Errorf("field %s: predicate == requires one value", owner)
	}
	return op, values, nil
}

func validateGateReference(reg *Registry, owner, ref string) error {
	if reg == nil {
		return nil
	}
	return reg.validateGateReference(owner, ref)
}

func stringList(value any) ([]string, error) {
	switch typed := value.(type) {
	case string:
		if typed == "" {
			return nil, nil
		}
		return []string{typed}, nil
	case []any:
		values := make([]string, 0, len(typed))
		for _, item := range typed {
			itemString, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("non-string item")
			}
			values = append(values, itemString)
		}
		return values, nil
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("not a string list")
	}
}

func (p *Predicate) Eval(fields map[string]string, rows func(array, field string) []string, ctx EvalContext) bool {
	if p == nil {
		return false
	}
	root := p.root
	if root == nil {
		parsed, err := ParsePredicate(p.Raw, nil)
		if err != nil {
			return false
		}
		root = parsed.root
	}
	if root == nil {
		return false
	}
	if fields == nil {
		fields = map[string]string{}
	}
	return root.eval(fields, rows, ctx)
}

func (e *predicateExpr) eval(fields map[string]string, rows func(array, field string) []string, ctx EvalContext) bool {
	switch e.kind {
	case predicateAllOf:
		for _, child := range e.children {
			if !child.eval(fields, rows, ctx) {
				return false
			}
		}
		return true
	case predicateAnyOf:
		for _, child := range e.children {
			if child.eval(fields, rows, ctx) {
				return true
			}
		}
		return false
	case predicateNot:
		return e.child != nil && !e.child.eval(fields, rows, ctx)
	case predicatePhaseIn:
		value := ctx.Phase
		if value == "" {
			value = fields["PHASE"]
		}
		return containsString(e.values, value)
	case predicateTierGTE:
		return tierAtLeast(ctx.Tier, e.values[0])
	case predicateAuthorityIn:
		return containsString(e.values, fields["AUTHORITY"])
	case predicateSeatIs:
		return containsString(e.values, ctx.Seat.Name)
	case predicateRoleIn:
		return containsString(e.values, ctx.Seat.Role)
	case predicateSlotIn:
		return containsString(e.values, ctx.Slot)
	case predicateRecordKind:
		return containsString(e.values, fields["record_kind"])
	case predicateScanResult:
		for _, field := range scanResultFields {
			if containsString(e.values, fields[field]) {
				return true
			}
		}
		return false
	case predicateClaims:
		return fields["ACTIONS_GIT_REF"] != "" || fields["FINAL_GIT_STATUS_SHORT"] != ""
	case predicateField:
		return evalOp(fields[e.field], e.op, e.values)
	case predicateAnyRow:
		if rows == nil {
			return false
		}
		for _, value := range rows(e.array, e.rowField) {
			if evalOp(value, e.op, e.values) {
				return true
			}
		}
		return false
	case predicateLayer:
		return ctx.PresentLayers[e.values[0]]
	default:
		return false
	}
}

func evalOp(candidate, op string, values []string) bool {
	switch op {
	case "present":
		return candidate != ""
	case "==":
		return len(values) == 1 && candidate == values[0]
	case "in":
		return containsString(values, candidate)
	default:
		return false
	}
}

func tierAtLeast(got, want string) bool {
	gotRank, gotOK := tierOrder[got]
	wantRank, wantOK := tierOrder[want]
	return gotOK && wantOK && gotRank >= wantRank
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

var predicateLayers = map[string]struct{}{
	"store":   {},
	"form":    {},
	"lineage": {},
	"observe": {},
}

var tierOrder = map[string]int{
	"tiny":            0,
	"small":           1,
	"medium":          2,
	"large":           3,
	"production-risk": 4,
}

var scanResultFields = []string{
	"scan_result",
	"SCOPE_DIFF_RESULT",
	"FOLD_SCOPE_RESULT",
	"ESCALATION_SCAN_RESULT",
}
