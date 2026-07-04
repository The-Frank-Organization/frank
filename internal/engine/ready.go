package engine

import (
	"fmt"
	"testing"
)

type Ready struct{}

type Diagnostics struct {
	class string
	want  string
	got   string
}

func NewReady() *Ready {
	return &Ready{}
}

func TestReady() *Ready {
	if !testing.Testing() {
		panic("engine.TestReady called outside go test")
	}
	return NewReady()
}

func NewDiagnostics(class, want, got string) *Diagnostics {
	return &Diagnostics{class: class, want: want, got: got}
}

func (d *Diagnostics) Report() string {
	if d == nil {
		return ""
	}
	if d.want != "" || d.got != "" {
		return fmt.Sprintf("diagnostics:%s want=%s got=%s", d.class, d.want, d.got)
	}
	return "diagnostics:" + d.class
}
