package intake_test

import (
	"context"
	"encoding/json"
	"errors"
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
	w, err := intake.NewWriter[writerOutcome](j, config.EngineConfig{SegmentRotateBytes: 256}, struct{}{})
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

func TestNewWriterRequiresReadyToken(t *testing.T) {
	j, err := intake.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	_, err = intake.NewWriter[writerOutcome](j, config.EngineConfig{}, nil)
	if !errors.Is(err, intake.ErrWriterNotReady) {
		t.Fatalf("NewWriter err = %v, want ErrWriterNotReady", err)
	}
}

func TestWriterInFlightCoalescesSingleExecution(t *testing.T) {
	j, err := intake.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	w, err := intake.NewWriter[writerOutcome](j, config.EngineConfig{}, struct{}{})
	if err != nil {
		t.Fatalf("NewWriter: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan intake.Job[writerOutcome], 2)
	go w.Run(ctx, out)

	cmd := intake.Cmd{Seat: "seat-a", Role: "implementer", Verb: "submit", Payload: json.RawMessage(`{"same":true}`)}
	firstReply, firstID, err := w.Submit(ctx, cmd)
	if err != nil {
		t.Fatalf("first Submit: %v", err)
	}
	secondReply, secondID, err := w.Submit(ctx, cmd)
	if err != nil {
		t.Fatalf("second Submit: %v", err)
	}
	if firstID != secondID {
		t.Fatalf("ids = %s/%s, want same", firstID, secondID)
	}
	job := <-out
	select {
	case extra := <-out:
		t.Fatalf("duplicate emitted extra job %+v", extra.Cmd)
	case <-time.After(100 * time.Millisecond):
	}

	want := writerOutcome{IntakeID: job.Cmd.IntakeID}
	job.ReplyCh <- want
	for name, reply := range map[string]<-chan writerOutcome{"first": firstReply, "second": secondReply} {
		select {
		case got := <-reply:
			if got != want {
				t.Fatalf("%s reply = %+v, want %+v", name, got, want)
			}
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out waiting for %s coalesced reply", name)
		}
	}
}
