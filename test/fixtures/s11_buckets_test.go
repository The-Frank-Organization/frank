package fixtures_test

import (
	"context"
	"errors"
	"os"
	"reflect"
	"slices"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
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

func TestS11BucketCIsOperatorCCFYIWithoutDecisionObligation(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	ccOnly := s11BucketRecord("bucket-c-cc-only", "")
	ccOnly.Headers["TO"] = `["s11.implementer"]`
	ccOnly.Headers["CC"] = `["operator"]`
	if _, err := st.Commit(ccOnly, nil); err != nil {
		t.Fatalf("commit CC-only record: %v", err)
	}
	operatorTo := s11BucketRecord("bucket-not-c-operator-to", "")
	operatorTo.Envelope.To = "operator"
	operatorTo.Headers["TO"] = `["operator"]`
	operatorTo.Headers["CC"] = `["operator"]`
	if _, err := st.Commit(operatorTo, nil); err != nil {
		t.Fatalf("commit operator-TO record: %v", err)
	}

	got, err := st.ProjectBucketC()
	if err != nil {
		t.Fatalf("ProjectBucketC: %v", err)
	}
	if want := []string{ccOnly.Envelope.RelayID}; !reflect.DeepEqual(got, want) {
		t.Fatalf("bucket C = %v, want %v", got, want)
	}
	project, err := st.Project("operator")
	if err != nil {
		t.Fatalf("Project operator: %v", err)
	}
	if want := []string{ccOnly.Envelope.RelayID, operatorTo.Envelope.RelayID}; !reflect.DeepEqual(project, want) {
		t.Fatalf("operator FYI projection = %v, want %v", project, want)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto C: %v", err)
	}
	if _, err := st.Read("odb-" + ccOnly.Envelope.RelayID); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("CC-FYI created a decision obligation: %v", err)
	}
}

func TestS11BucketDIsAuthorFacingAndEgressBlockedStaysA(t *testing.T) {
	root := t.TempDir()
	pinned := initFixtureStore(t, root)
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	rejected := s11BucketRecord("bucket-d-observe", "")
	rejected.Envelope.DeliveryState = record.Rejected
	rejected.Headers["failing_edge"] = "observe-predicate"
	if _, err := st.Commit(rejected, nil); err != nil {
		t.Fatalf("commit D record: %v", err)
	}
	egressBlocked := s11BucketRecord("bucket-a-egress", "authz_security")
	egressBlocked.Headers["failing_edge"] = "egress"
	egressBlocked.Headers["egress_scan_result"] = "blocked"
	if _, err := st.Commit(egressBlocked, nil); err != nil {
		t.Fatalf("commit egress-blocked record: %v", err)
	}

	got, err := st.ProjectBucketD("s11.planner")
	if err != nil {
		t.Fatalf("ProjectBucketD author: %v", err)
	}
	if want := []string{rejected.Envelope.RelayID}; !reflect.DeepEqual(got, want) {
		t.Fatalf("bucket D = %v, want %v", got, want)
	}
	if got, err := st.ProjectBucketD("operator"); err != nil || len(got) != 0 {
		t.Fatalf("operator bucket D = %v, %v; want empty", got, err)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto egress blocked: %v", err)
	}
	if odb, err := st.Read("odb-" + egressBlocked.Envelope.RelayID); err != nil || odb.Envelope.To != "operator" {
		t.Fatalf("egress-blocked A missing local operator ODB: %+v, %v", odb, err)
	}

	meta := seat.SeatMeta{Name: "s11.implementer", Role: "implementer"}
	renderEnv := fieldspec.RenderEnv{ConfigDigest: pinned.Digest}
	payload := s10ExitPayload(t, pinned.Registry, renderEnv, meta, record.Record{Headers: map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
	}})
	formRejected, _, err := engine.SubmitHandlerWithRender(st, pinned.Registry, meta, renderEnv)(context.Background(), intake.Cmd{
		IntakeID: "s11-form-rejected", Seat: meta.Name, Role: meta.Role, Payload: payload,
	})
	if err != nil {
		t.Fatalf("SubmitHandler form rejection: %v", err)
	}
	if formRejected.Envelope.DeliveryState != record.Rejected || formRejected.Headers["failing_edge"] != "form-validation" {
		t.Fatalf("form rejection = %s/%q, want rejected/form-validation", formRejected.Envelope.DeliveryState, formRejected.Headers["failing_edge"])
	}
}

func TestS11BucketTerminalAndFailingEdgeMatrix(t *testing.T) {
	cases := []struct {
		name  string
		state string
		edge  string
		setup func(*record.Record)
		wantA bool
		wantB bool
		wantC bool
		wantD bool
	}{
		{name: "A accepted authority gate", state: record.Accepted, setup: func(rec *record.Record) {
			rec.Headers["gate_category"] = "authz_security"
		}, wantA: true},
		{name: "A accepted egress block never D", state: record.Accepted, edge: "egress", setup: func(rec *record.Record) {
			rec.Headers["gate_category"] = "authz_security"
			rec.Headers["egress_scan_result"] = "blocked"
		}, wantA: true},
		{name: "A held fault never D", state: record.Held, edge: "stale_schema", wantA: true},
		{name: "B accepted category", state: record.Accepted, setup: func(rec *record.Record) {
			rec.Headers["gate_category"] = "routing"
		}, wantB: true},
		{name: "B category rejected at acceptance becomes D", state: record.Rejected, edge: "form-validation", setup: func(rec *record.Record) {
			rec.Headers["gate_category"] = "routing"
		}, wantD: true},
		{name: "C accepted operator CC only", state: record.Accepted, setup: func(rec *record.Record) {
			rec.Headers["TO"] = `["s11.implementer"]`
			rec.Headers["CC"] = `["operator"]`
		}, wantC: true},
		{name: "C addressing rejected at acceptance becomes D", state: record.Rejected, edge: "lineage", setup: func(rec *record.Record) {
			rec.Headers["TO"] = `["s11.implementer"]`
			rec.Headers["CC"] = `["operator"]`
		}, wantD: true},
		{name: "D form validation", state: record.Rejected, edge: "form-validation", wantD: true},
		{name: "D lineage", state: record.Rejected, edge: "lineage", wantD: true},
		{name: "D observe predicate", state: record.Rejected, edge: "observe-predicate", wantD: true},
		{name: "D declared versus observed", state: record.Rejected, edge: "declared-vs-observed", wantD: true},
		{name: "D requires rejected terminal", state: record.Accepted, edge: "observe-predicate"},
		{name: "held never falls through to D", state: record.Held, edge: "observe-predicate", wantA: true},
		{name: "egress edge without accepted A stage is not D", state: record.Rejected, edge: "egress"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			pinned := initFixtureStore(t, root)
			st, err := store.Open(root)
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			rec := s11BucketRecord("matrix-record", "")
			rec.Envelope.DeliveryState = tc.state
			rec.Headers["failing_edge"] = tc.edge
			if tc.setup != nil {
				tc.setup(&rec)
			}
			if _, err := st.Commit(rec, nil); err != nil {
				t.Fatalf("Commit: %v", err)
			}
			if err := obligation.CompleteAuto(st); err != nil {
				t.Fatalf("CompleteAuto: %v", err)
			}

			bucketB, err := st.ProjectBucketB(pinned.Registry)
			if err != nil {
				t.Fatalf("ProjectBucketB: %v", err)
			}
			bucketC, err := st.ProjectBucketC()
			if err != nil {
				t.Fatalf("ProjectBucketC: %v", err)
			}
			bucketD, err := st.ProjectBucketD(rec.Envelope.From)
			if err != nil {
				t.Fatalf("ProjectBucketD: %v", err)
			}
			if got := slices.Contains(bucketB, rec.Envelope.RelayID); got != tc.wantB {
				t.Fatalf("bucket B membership = %v, want %v; %v", got, tc.wantB, bucketB)
			}
			if got := slices.Contains(bucketC, rec.Envelope.RelayID); got != tc.wantC {
				t.Fatalf("bucket C membership = %v, want %v; %v", got, tc.wantC, bucketC)
			}
			if got := slices.Contains(bucketD, rec.Envelope.RelayID); got != tc.wantD {
				t.Fatalf("bucket D membership = %v, want %v; %v", got, tc.wantD, bucketD)
			}
			hasA := s11HasBucketAArtifact(st, rec)
			if hasA != tc.wantA {
				t.Fatalf("bucket A derived artifact = %v, want %v", hasA, tc.wantA)
			}
		})
	}
}

func TestS11KnownABPickIsRaisedRecordedAndNeverAbsorbed(t *testing.T) {
	st, reg, meta := s5GateRaiseDeps(t)
	rec := s5GateRaiseCandidate()
	rec.Headers["gate_category"] = "routing"
	committed := s5SubmitAndCommit(t, st, reg, meta, s5AFloorEnv(reg, "authz_security"), rec)
	if committed.Envelope.DeliveryState != record.Accepted ||
		committed.Headers["gate_category"] != "authz_security" ||
		committed.Headers["gate_category_raised"] != "yes" ||
		committed.Headers["gate_category_pick"] != "routing" {
		t.Fatalf("known-A raise = %+v", committed)
	}
	if bucketB, err := st.ProjectBucketB(reg); err != nil || slices.Contains(bucketB, committed.Envelope.RelayID) {
		t.Fatalf("known-A record silently absorbed in B: %v, %v", bucketB, err)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}
	if _, err := st.Read("odb-" + committed.Envelope.RelayID); err != nil {
		t.Fatalf("known-A record did not reach A decision surface: %v", err)
	}
}

func s11HasBucketAArtifact(st *store.Store, rec record.Record) bool {
	if _, err := st.Read("odb-" + rec.Envelope.RelayID); err == nil {
		return true
	}
	if _, err := st.Read("outbox-held-" + rec.Envelope.RelayID); err == nil {
		return true
	}
	return false
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
