package lineage_test

import (
	"testing"

	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestAuthorityBearingPessimisticSuperset(t *testing.T) {
	meta := seat.SeatMeta{Name: "s1.orchestrator-planner", Role: "orchestrator-planner"}
	for name, rec := range map[string]record.Record{
		"grant":                    {Headers: map[string]string{"grant": "dispatch-impl"}},
		"hgr":                      {Headers: map[string]string{"HUMAN_GATE_REQUIRED": "yes"}},
		"gate a":                   {Headers: map[string]string{"gate_category": "authz_security"}},
		"plan phase":               {Headers: map[string]string{"PHASE": "PLAN"}},
		"implementation authority": {Headers: map[string]string{"AUTHORITY": "implementation"}},
		"orchestrator role":        {Headers: map[string]string{"PHASE": "PLAN"}},
	} {
		if !lineage.AuthorityBearing(rec, meta) {
			t.Fatalf("%s was not authority-bearing", name)
		}
	}
	if lineage.AuthorityBearing(record.Record{Headers: map[string]string{"PHASE": "SITREP"}}, seat.SeatMeta{Name: "s1-core.implementer", Role: "implementer"}) {
		t.Fatalf("plain pair SITREP was authority-bearing")
	}
}

func TestCheckParentEdges(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(record.Record{
		Envelope: record.Envelope{RelayID: "rejected-parent", DeliveryState: record.Rejected, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "bad"},
	}, nil); err != nil {
		t.Fatalf("commit parent: %v", err)
	}

	unknown := record.Record{Headers: map[string]string{"PHASE": "PLAN", "PARENT_DISPATCH_ID": "missing", "SUBJECT": "child"}}
	if bounce := lineage.Check(unknown, st); bounce == nil || bounce.Kind != lineage.ParentUnknownRecompose {
		t.Fatalf("unknown parent bounce = %+v", bounce)
	}

	dead := record.Record{Headers: map[string]string{"PHASE": "PLAN", "PARENT_DISPATCH_ID": "rejected-parent", "SUBJECT": "child"}}
	if bounce := lineage.Check(dead, st); bounce == nil || bounce.Kind != lineage.ParentInvalidDeadEdge {
		t.Fatalf("dead parent bounce = %+v", bounce)
	}
}
