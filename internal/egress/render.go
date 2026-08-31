package egress

import (
	"sort"

	"github.com/The-Frank-Organization/frank/internal/gate"
	"github.com/The-Frank-Organization/frank/internal/record"
)

type RenderedField struct {
	Name   string
	Value  string
	Origin Origin
}

type Rendered struct {
	Dest   string
	Fields []RenderedField
}

type Renderer func(meta gate.OutboxItem, source record.Record) (Rendered, error)

func DefaultRenderer(meta gate.OutboxItem, source record.Record) (Rendered, error) {
	dest := source.Envelope.To
	if dest == "" {
		dest = meta.Seat
	}
	fields := make([]RenderedField, 0, len(source.Headers)+1)
	names := make([]string, 0, len(source.Headers))
	for name := range source.Headers {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fields = append(fields, RenderedField{Name: name, Value: source.Headers[name]})
	}
	if source.Body != "" {
		fields = append(fields, RenderedField{Name: "body", Value: source.Body})
	}
	return Rendered{Dest: dest, Fields: fields}, nil
}
