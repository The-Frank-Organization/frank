package lineage_test

import (
	"testing"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
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

func TestAuthorityBearingGateCategoryUsesRegistry(t *testing.T) {
	meta := seat.SeatMeta{Name: "s3-form.implementer", Role: "implementer"}
	eng := lineage.Engine{Reg: &fieldspec.Registry{GateCategory: map[string][]string{
		"A": {"registry-only-a"},
		"B": {"registry-only-b"},
	}}}

	for _, token := range []string{"registry-only-a", "other"} {
		if !eng.AuthorityBearing(record.Record{Headers: map[string]string{"gate_category": token}}, meta) {
			t.Fatalf("%s was not authority-bearing", token)
		}
	}
	for _, token := range []string{"registry-only-b", "authz_security", "unknown"} {
		if eng.AuthorityBearing(record.Record{Headers: map[string]string{"gate_category": token}}, meta) {
			t.Fatalf("%s unexpectedly authority-bearing", token)
		}
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
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("Build tables: %v", err)
	}
	if bounce := lineage.Check(unknown, tab); bounce == nil || bounce.Kind != lineage.ParentUnknownRecompose {
		t.Fatalf("unknown parent bounce = %+v", bounce)
	}

	dead := record.Record{Headers: map[string]string{"PHASE": "PLAN", "PARENT_DISPATCH_ID": "rejected-parent", "SUBJECT": "child"}}
	tab, err = tables.Build(st)
	if err != nil {
		t.Fatalf("Build tables: %v", err)
	}
	if bounce := lineage.Check(dead, tab); bounce == nil || bounce.Kind != lineage.ParentInvalidDeadEdge {
		t.Fatalf("dead parent bounce = %+v", bounce)
	}
}

func TestEngineIMPLActionMustParentAddressingDispatch(t *testing.T) {
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "dispatch-1", DispatchID: "dispatch-1", To: "seat-a", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"TO": "seat-a"},
	})
	eng := lineage.Engine{T: tab}
	cand := record.Record{Headers: map[string]string{"PHASE": "IMPL", "PARENT_DISPATCH_ID": "dispatch-1", "ACTIONS_GIT_REF": "branch@sha"}}
	if bounce := eng.Check(cand, seat.SeatMeta{Name: "seat-a", Role: "implementer"}); bounce != nil {
		t.Fatalf("addressed IMPL bounced: %+v", bounce)
	}
	if bounce := eng.Check(cand, seat.SeatMeta{Name: "seat-b", Role: "implementer"}); bounce == nil || bounce.Kind != lineage.ParentNotAddressed {
		t.Fatalf("non-addressee bounce = %+v", bounce)
	}
}

func TestEngineMergeClaimRequiresPriorMergeGateGrant(t *testing.T) {
	tab := tables.New()
	eng := lineage.Engine{T: tab}
	cand := record.Record{Envelope: record.Envelope{DispatchID: "dispatch-1"}, Headers: map[string]string{"PHASE": "MERGE-GATE", "grant": "dispatch-merge"}}
	if bounce := eng.Check(cand, seat.SeatMeta{Name: "s1-core.planner", Role: "planner"}); bounce == nil || bounce.Kind != lineage.MergeGrantMissing {
		t.Fatalf("missing merge grant bounce = %+v", bounce)
	}
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "merge-grant", DispatchID: "dispatch-1", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "MERGE-GATE", "grant": "dispatch-merge"},
	})
	if bounce := eng.Check(cand, seat.SeatMeta{Name: "s1-core.planner", Role: "planner"}); bounce != nil {
		t.Fatalf("merge claim with prior grant bounced: %+v", bounce)
	}
}

func TestEngineScopeFlipDrift(t *testing.T) {
	tab := tables.New()
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "scope-old", DispatchID: "dispatch-1", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"SCOPE_DIFF": `[{"path":"README.md","state":"OUT"}]`},
	})
	eng := lineage.Engine{T: tab}
	cand := record.Record{
		Envelope: record.Envelope{DispatchID: "dispatch-1"},
		Headers:  map[string]string{"PHASE": "PLAN", "ROW_TRUTH_CHECK": "required", "SCOPE_DIFF": `[{"path":"README.md","state":"IN"}]`},
	}
	if bounce := eng.Check(cand, seat.SeatMeta{Name: "s1-core.planner", Role: "planner"}); bounce == nil || bounce.Kind != lineage.ScopeFlipDrift {
		t.Fatalf("scope flip bounce = %+v", bounce)
	}
}

func TestRealGrantStateAndActiveLineageCandidates(t *testing.T) {
	tab := tables.New()
	grantState := lineage.RealGrantState(tab)
	if grantState(fieldspec.SeatMeta{Name: "s1-core.planner", Role: "planner"}) {
		t.Fatalf("grant state true before approving review")
	}
	tab.OnCommit(record.Record{
		Envelope: record.Envelope{RelayID: "review-1", DispatchID: "dispatch-1", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "PLAN-REVIEW", "PLAN_REVIEW_VERDICT": "approve"},
	})
	if !grantState(fieldspec.SeatMeta{Name: "s1-core.planner", Role: "planner"}) {
		t.Fatalf("grant state false after approving review")
	}
	candidates, dflt := lineage.ActiveLineageCandidates(tab, lineage.TurnContext{WokenOn: "wake-1", ActiveDispatch: "dispatch-1"})(fieldspec.SeatMeta{})
	if dflt != "wake-1" || len(candidates) != 3 || candidates[1] != "dispatch-1" || candidates[2] != "review-1" {
		t.Fatalf("candidates=%v default=%s", candidates, dflt)
	}
}
