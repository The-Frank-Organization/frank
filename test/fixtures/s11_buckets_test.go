package fixtures_test

import (
	"context"
	"errors"
	"os"
	"reflect"
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
