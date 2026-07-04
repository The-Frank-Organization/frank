package intake_test

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/intake"
)

type writerOutcome struct {
	IntakeID string
}

func TestWriterConcurrentSubmitsAreGapFreeAndEmittedInJournalOrder(t *testing.T) {
	j, err := intake.OpenWithConfig(t.TempDir(), config.EngineConfig{SegmentRotateBytes: 256})
	if err != nil {
		t.Fatalf("OpenWithConfig: %v", err)
	}
	w, err := intake.NewWriter[writerOutcome](j, config.EngineConfig{SegmentRotateBytes: 256})
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan intake.Job[writerOutcome], 200)
	go w.Run(ctx, out)

	const workers = 8
	const perWorker = 25
	var wg sync.WaitGroup
	ids := make(chan string, workers*perWorker)
	replies := make(chan (<-chan writerOutcome), workers*perWorker)
	for worker := 0; worker < workers; worker++ {
		worker := worker
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := 0; n < perWorker; n++ {
				payload := json.RawMessage(fmt.Sprintf(`{"worker":%d,"n":%d}`, worker, n))
				reply, intakeID, err := w.Submit(ctx, intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: payload})
				if err != nil {
					t.Errorf("Submit: %v", err)
					return
				}
				ids <- intakeID
				replies <- reply
			}
		}()
	}
	wg.Wait()
	close(ids)
	close(replies)

	var submitted []string
	for id := range ids {
		submitted = append(submitted, id)
	}
	sort.Strings(submitted)
	for i, id := range submitted {
		want := fmt.Sprintf("intake-%06d", i+1)
		if id != want {
			t.Fatalf("submitted id[%d] = %s, want %s", i, id, want)
		}
	}

	var emitted []string
	for i := 0; i < workers*perWorker; i++ {
		select {
		case job := <-out:
			emitted = append(emitted, job.Cmd.IntakeID)
			job.ReplyCh <- writerOutcome{IntakeID: job.Cmd.IntakeID}
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out waiting for emitted job %d", i)
		}
	}
	for i, id := range emitted {
		want := fmt.Sprintf("intake-%06d", i+1)
		if id != want {
			t.Fatalf("emitted id[%d] = %s, want %s", i, id, want)
		}
	}
	for reply := range replies {
		select {
		case out := <-reply:
			if out.IntakeID == "" {
				t.Fatalf("empty reply outcome: %+v", out)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out waiting for writer reply")
		}
	}

	entries, err := j.ReadAll()
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if len(entries) != workers*perWorker {
		t.Fatalf("journal entries = %d, want %d", len(entries), workers*perWorker)
	}
	for i, entry := range entries {
		if entry.IntakeID != emitted[i] {
			t.Fatalf("journal[%d] = %s, emitted %s", i, entry.IntakeID, emitted[i])
		}
	}
}
