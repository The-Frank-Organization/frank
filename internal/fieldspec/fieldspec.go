package fieldspec

import "slices"

type SeatMeta struct {
	Name       string
	Role       string
	IsOperator bool
}

type Form struct {
	Fields map[string]Field `json:"fields"`
}

type Field struct {
	Type              string   `json:"type,omitempty"`
	Options           []string `json:"options,omitempty"`
	Default           string   `json:"default,omitempty"`
	DigestExempt      bool     `json:"digest_exempt,omitempty"`
	ConductorVolatile bool     `json:"conductor_volatile,omitempty"`
}

type Violation struct {
	Field  string
	Class  string
	Reason string
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
