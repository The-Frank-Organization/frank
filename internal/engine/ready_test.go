package engine_test

import (
	"context"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestNewPanicsWithoutReady(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	defer func() {
		if recovered := recover(); recovered == nil {
			t.Fatalf("New did not panic without Ready")
		}
	}()
	_ = engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{}, nil, nil
	}, nil)
}
