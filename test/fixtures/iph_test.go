package fixtures_test

import (
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestP1NoPathFamiliesInSeatDeliverableStrings(t *testing.T) {
	outputs := []string{
		bounce.Format(fieldspec.Violation{Field: "PHASE", Class: "enum", Reason: "/store/records/leak"}),
		`{"state":"accepted"}`,
		`["submit","project","read"]`,
	}
	for _, output := range outputs {
		for _, family := range []string{"/records", "/staging", "/outbox", "/binding", "operator-socket"} {
			if strings.Contains(output, family) {
				t.Fatalf("seat-deliverable output leaked %s in %q", family, output)
			}
		}
	}
}

func TestP1LoopOutcomeDoesNotLeakStorePaths(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{}, nil, errors.New(filepath.Join(root, "records", "leak.json"))
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	reply := make(chan engine.Outcome, 1)
	payload, _ := json.Marshal(record.Record{Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "leak"}})
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "i-leak", Seat: "seat-a", Role: "implementer", Payload: payload}, ReplyCh: reply}
	out := <-reply
	if strings.Contains(out.Reason, root) || strings.Contains(out.Reason, "/records/") {
		t.Fatalf("loop outcome leaked path: %+v", out)
	}
}
