package fixtures_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestS11FSMAddsBouncedRepairAndFixtureScopedEgressBlocked(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	bounced := s11BucketRecord("fsm-bounced", "")
	bounced.Envelope.DeliveryState = record.Rejected
	bounced.Headers["failing_edge"] = "observe-predicate"
	if _, err := st.Commit(bounced, nil); err != nil {
		t.Fatalf("commit bounced record: %v", err)
	}
	if state := engine.GateState(s10ExitTables(t, st), bounced.Envelope.RelayID); state != engine.GateBouncedRepair {
		t.Fatalf("bounced state = %q, want %q", state, engine.GateBouncedRepair)
	}

	egressBlocked := s11BucketRecord("fsm-egress-blocked", "authz_security")
	egressBlocked.Headers["egress_scan_result"] = "blocked"
	egressBlocked.Headers["failing_edge"] = "egress"
	if _, err := st.Commit(egressBlocked, nil); err != nil {
		t.Fatalf("commit egress-blocked gate: %v", err)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}
	if state := engine.GateState(s10ExitTables(t, st), egressBlocked.Envelope.RelayID); state != engine.GateEgressBlocked {
		t.Fatalf("egress state = %q, want %q", state, engine.GateEgressBlocked)
	}
	if park, err := st.Read("park-" + egressBlocked.Envelope.RelayID); err != nil || park.Headers["parks_gate"] != egressBlocked.Envelope.RelayID {
		t.Fatalf("egress-blocked gate not parked locally: %+v, %v", park, err)
	}

	input := engine.ResummonInput{
		Seat: "s11.planner", DecisionID: egressBlocked.Envelope.RelayID, CadenceSlot: "egress-blocked-1",
		Reason: engine.ResummonNoResponse, SummonChannel: engine.SummonLocal,
	}
	payload, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Marshal resummon: %v", err)
	}
	resummon, _, err := engine.ResummonHandler(nil)(context.Background(), intake.Cmd{
		Seat: "system", Role: "system", Verb: "emit-resummon", Payload: payload, ContentHash: engine.ResummonContentHash(input),
	})
	if err != nil {
		t.Fatalf("ResummonHandler: %v", err)
	}
	if resummon.Envelope.DeliveryState != record.Accepted || resummon.Headers["subject_ref"] != egressBlocked.Envelope.RelayID {
		t.Fatalf("local resummon = %+v", resummon)
	}
	for _, forbidden := range []string{"approve", "redact", "external_send"} {
		if strings.Contains(resummon.Body, forbidden) {
			t.Fatalf("fixture-scoped egress resummon contains %q: %s", forbidden, resummon.Body)
		}
	}
}
