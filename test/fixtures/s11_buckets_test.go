package fixtures_test

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestS11BucketBIsLiveNonInterruptingAndRaiseOnly(t *testing.T) {
	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	bRecord := s11BucketRecord("bucket-b-routing", "routing")
	if _, err := st.Commit(bRecord, nil); err != nil {
		t.Fatalf("commit B record: %v", err)
	}
	got, err := st.ProjectBucketB(pinned.Registry)
	if err != nil {
		t.Fatalf("ProjectBucketB: %v", err)
	}
	if want := []string{bRecord.Envelope.RelayID}; !reflect.DeepEqual(got, want) {
		t.Fatalf("bucket B = %v, want %v", got, want)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto B: %v", err)
	}
	if _, err := st.Read("odb-" + bRecord.Envelope.RelayID); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("B record entered operator decision queue: %v", err)
	}

	raised := s11BucketRecord("bucket-a-raised", "authz_security")
	raised.Headers["gate_category_raised"] = "yes"
	if _, err := st.Commit(raised, nil); err != nil {
		t.Fatalf("commit raised record: %v", err)
	}
	got, err = st.ProjectBucketB(pinned.Registry)
	if err != nil {
		t.Fatalf("ProjectBucketB after raise: %v", err)
	}
	if want := []string{bRecord.Envelope.RelayID}; !reflect.DeepEqual(got, want) {
		t.Fatalf("bucket B after raise = %v, want %v", got, want)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto raised: %v", err)
	}
	if odb, err := st.Read("odb-" + raised.Envelope.RelayID); err != nil || odb.Envelope.To != "operator" {
		t.Fatalf("raised record missing operator ODB: %+v, %v", odb, err)
	}
}

func s11BucketRecord(relayID, category string) record.Record {
	return record.Record{
		Envelope: record.Envelope{
			RelayID: relayID, DispatchID: "s11-buckets", From: "s11.planner", To: "s11.implementer", Role: "planner",
			DeliveryState: record.Accepted, SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE": "SITREP", "SUBJECT": relayID, "gate_category": category,
		},
	}
}
