package migrate

import (
	"errors"
	"fmt"

	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

const Current = 1

var (
	ErrFutureVersion = errors.New("future schema_version")
	ErrUnversioned   = errors.New("unversioned record")
	ErrMigrationGap  = errors.New("migration gap")
)

type Func func(record.Record) (record.Record, error)

type Registry struct {
	Current int
	steps   map[int]Func
}

type Reader struct {
	Store    *store.Store
	Registry *Registry
}

type View struct {
	Record        record.Record
	SourceVersion int
}

func New() *Registry {
	return &Registry{Current: Current, steps: map[int]Func{}}
}

func (r *Registry) Register(from int, fn Func) {
	if r.steps == nil {
		r.steps = map[int]Func{}
	}
	r.steps[from] = fn
}

func Apply(reg *Registry, rec record.Record) (record.Record, error) {
	if reg == nil {
		reg = New()
	}
	current := reg.Current
	if current == 0 {
		current = Current
	}
	version := rec.Envelope.SchemaVersion
	switch {
	case version == 0:
		return record.Record{}, ErrUnversioned
	case version > current:
		return record.Record{}, ErrFutureVersion
	case version == current:
		return cloneRecord(rec), nil
	}
	out := cloneRecord(rec)
	for from := version; from < current; from++ {
		step, ok := reg.steps[from]
		if !ok {
			return record.Record{}, fmt.Errorf("%w: %d->%d", ErrMigrationGap, from, from+1)
		}
		next, err := step(cloneRecord(out))
		if err != nil {
			return record.Record{}, err
		}
		next.Envelope.SchemaVersion = from + 1
		out = next
	}
	return out, nil
}

func cloneRecord(rec record.Record) record.Record {
	rec.Headers = cloneStrings(rec.Headers)
	rec.XFields = cloneStrings(rec.XFields)
	return rec
}

func cloneStrings(values map[string]string) map[string]string {
	if values == nil {
		return nil
	}
	out := make(map[string]string, len(values))
	for key, value := range values {
		out[key] = value
	}
	return out
}

func (r Reader) Read(relayID string) (View, error) {
	rec, err := r.Store.Read(relayID)
	if err != nil {
		return View{}, err
	}
	source := rec.Envelope.SchemaVersion
	view, err := Apply(r.Registry, rec)
	if err != nil {
		return View{}, err
	}
	return View{Record: view, SourceVersion: source}, nil
}
