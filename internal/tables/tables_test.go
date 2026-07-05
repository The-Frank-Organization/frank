package tables_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestIncrementalMatchesRebuild(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	inc, err := tables.Build(st)
	if err != nil {
		t.Fatalf("initial Build: %v", err)
	}
	for _, rec := range []record.Record{
		{Envelope: record.Envelope{RelayID: "r1", DispatchID: "d1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, IntakeID: "i1", SchemaVersion: 1}, Headers: map[string]string{"PHASE": "SITREP", "HUMAN_GATE_REQUIRED": "yes"}},
		{Envelope: record.Envelope{RelayID: "park-r1", From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"parks_gate": "r1"}},
		{Envelope: record.Envelope{RelayID: "owed", From: "operator", Role: "operator", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"record_kind": "owed_disposition", "disposes_owed": "owed-base"}},
	} {
		if _, err := st.Commit(rec, nil); err != nil {
			t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
		}
		inc.OnCommit(rec)
	}
	rebuilt, err := tables.Build(st)
	if err != nil {
		t.Fatalf("rebuilt Build: %v", err)
	}
	if !reflect.DeepEqual(inc, rebuilt) {
		t.Fatalf("incremental tables differ from rebuild\nincremental=%#v\nrebuilt=%#v", inc, rebuilt)
	}
}

func TestLiveSnapshotIsImmutableAcrossPublish(t *testing.T) {
	initial := tables.New()
	initial.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "r1", DispatchID: "d1", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP"},
	})
	live := tables.NewLive(initial)

	snapshot := live.Snapshot()
	next := snapshot.Clone()
	next.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "r2", DispatchID: "d1", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP"},
	})
	live.Publish(next)

	if _, ok := snapshot.ByRelay["r2"]; ok {
		t.Fatalf("previous snapshot observed later publish")
	}
	current := live.Snapshot()
	if _, ok := current.ByRelay["r2"]; !ok {
		t.Fatalf("current snapshot missing published record")
	}
}

func TestLiveSnapshotConcurrentPublishAndRead(t *testing.T) {
	live := tables.NewLive(tables.New())
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			next := live.Snapshot().Clone()
			next.OnCommit(record.Record{
				Envelope: record.Envelope{RelayID: "r", DispatchID: "d", DeliveryState: record.Accepted, SchemaVersion: 1},
				Headers:  map[string]string{"SUBJECT": "publish"},
			})
			live.Publish(next)
		}()
		go func() {
			defer wg.Done()
			snapshot := live.Snapshot()
			_ = snapshot.Records
			_ = snapshot.ByDispatch["d"]
		}()
	}
	wg.Wait()
}
