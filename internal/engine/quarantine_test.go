package engine_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestLoopEnqueueQuarantineRunsDispositionOnLoopTurn(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "bad-live", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "bad"},
		Body:     "before",
	}, nil); err != nil {
		t.Fatalf("Commit bad-live: %v", err)
	}
	path := filepath.Join(root, "records", "bad-live.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read bad-live: %v", err)
	}
	data = []byte(string(data) + "\n")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("corrupt bad-live: %v", err)
	}

	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{RelayID: "tick", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: "tick", SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "tick"},
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	loop.EnqueueQuarantine("bad-live")
	reply := make(chan engine.Outcome, 1)
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "tick"}, ReplyCh: reply}
	<-reply

	if _, err := os.Stat(filepath.Join(root, "quarantine", "bad-live.json")); err != nil {
		t.Fatalf("quarantine file missing: %v", err)
	}
	if countEngineRelay(t, st, "incident-bad-live") != 1 {
		t.Fatalf("incident-bad-live not committed")
	}
}

func countEngineRelay(t *testing.T, st *store.Store, relayID string) int {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var count int
	for _, rec := range records {
		if rec.Envelope.RelayID == relayID {
			count++
		}
	}
	return count
}
