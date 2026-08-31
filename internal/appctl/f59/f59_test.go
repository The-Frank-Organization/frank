package f59

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/manifest"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestEffectDescriptorExactApplicabilityAndStablePatchDigest(t *testing.T) {
	fixture := newFixture(t, "descriptor")
	paths := []string{"src/z.go", "src/a.go"}
	patchDigest, err := OrderedTargetSetDigest(paths)
	if err != nil {
		t.Fatal(err)
	}
	reversed, err := OrderedTargetSetDigest([]string{"src/a.go", "src/z.go"})
	if err != nil || reversed != patchDigest {
		t.Fatalf("ordered target digest drift: first=%q second=%q err=%v", patchDigest, reversed, err)
	}

	resource := "src/main.go"
	cwd := "."
	patch := patchDigest
	relay := "relay.read:relay-1"
	tests := []struct {
		name     string
		operands Operands
		wantRoot bool
		wantRes  bool
		wantCWD  bool
		backend  string
	}{
		{name: "read", operands: Operands{CanonicalResource: &resource, CWD: &cwd}, wantRoot: true, wantRes: true, wantCWD: true, backend: "ambient"},
		{name: "write", operands: Operands{CanonicalResource: &resource, CWD: &cwd}, wantRoot: true, wantRes: true, wantCWD: true, backend: "ambient"},
		{name: "edit", operands: Operands{CanonicalResource: &resource, CWD: &cwd}, wantRoot: true, wantRes: true, wantCWD: true, backend: "ambient"},
		{name: "apply_patch", operands: Operands{CanonicalResource: &patch, CWD: &cwd}, wantRoot: true, wantRes: true, wantCWD: true, backend: "ambient"},
		{name: "bash", operands: Operands{CWD: &cwd}, wantRoot: true, wantCWD: true, backend: "ambient"},
		{name: "relay.read", operands: Operands{CanonicalResource: &relay}, wantRes: true, backend: "conductor-client"},
		{name: "relay.project", operands: Operands{CanonicalResource: stringPointer("relay.project:inbox")}, wantRes: true, backend: "conductor-client"},
		{name: "relay.submit", operands: Operands{CanonicalResource: stringPointer("relay.submit:" + digest("target"))}, wantRes: true, backend: "conductor-client"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			descriptor, encoded, err := BuildDescriptor(fixture.frozen.Manifest, test.name, digest("args"), test.operands)
			if err != nil {
				t.Fatalf("BuildDescriptor: %v", err)
			}
			if (descriptor.WorkspaceRootID != nil) != test.wantRoot || (descriptor.CanonicalResource != nil) != test.wantRes || (descriptor.CWD != nil) != test.wantCWD {
				t.Fatalf("descriptor applicability = %#v", descriptor)
			}
			if descriptor.BackendID != test.backend || descriptor.NetworkPolicyID != "none" || !descriptor.OneShot || descriptor.ToolImplRef == "" || len(encoded) == 0 {
				t.Fatalf("descriptor constants = %#v bytes=%q", descriptor, encoded)
			}
		})
	}

	for name, operands := range map[string]Operands{
		"absent required resource":           {CWD: &cwd},
		"present inapplicable bash resource": {CanonicalResource: &resource, CWD: &cwd},
		"present inapplicable relay cwd":     {CanonicalResource: &relay, CWD: &cwd},
		"parent segment":                     {CanonicalResource: stringPointer("src/../secret"), CWD: &cwd},
		"non canonical cwd":                  {CanonicalResource: &resource, CWD: stringPointer("work/")},
	} {
		t.Run(name, func(t *testing.T) {
			action := "read"
			if name == "present inapplicable bash resource" {
				action = "bash"
			}
			if name == "present inapplicable relay cwd" {
				action = "relay.read"
			}
			if _, _, err := BuildDescriptor(fixture.frozen.Manifest, action, digest("args"), operands); !errors.Is(err, ErrInvalidDescriptor) {
				t.Fatalf("descriptor error = %v, want ErrInvalidDescriptor", err)
			}
		})
	}
}

func TestIssueReplayPrecedesMutableStateAndOperandsJoinIdentity(t *testing.T) {
	fixture := newFixture(t, "issue-replay")
	request := fixture.issue("call-1", "read")
	result, err := fixture.f59.Issue(fixture.ctx, request)
	if err != nil || result.Kind != TicketGranted || result.TicketID == "" {
		t.Fatalf("first issue = %#v err=%v", result, err)
	}
	firstSeq := fixture.stateSeq()

	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED' WHERE turn_id=?`, fixture.turnID); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `UPDATE epochs SET turn_epoch=? WHERE run_id=?`, storeCounter(2), fixture.runID)
		return err
	})
	replay, err := fixture.f59.Issue(fixture.ctx, request)
	if err != nil || replay.Kind != TicketGranted || replay.TicketID != result.TicketID {
		t.Fatalf("replay after mutable drift = %#v err=%v", replay, err)
	}
	if got := fixture.stateSeq(); got != firstSeq+1 { // only the explicit fixture mutation advanced it
		t.Fatalf("replay mutated state_seq: got=%d want=%d", got, firstSeq+1)
	}

	changed := request
	changed.Operands.CanonicalResource = stringPointer("src/other.go")
	mismatch, err := fixture.f59.Issue(fixture.ctx, changed)
	if err != nil || mismatch.Kind != IdentityMismatch {
		t.Fatalf("operand mismatch replay = %#v err=%v", mismatch, err)
	}
	if got := fixture.authorizationCount(); got != 1 {
		t.Fatalf("authorization count = %d, want 1", got)
	}
}

func TestIssueOrderedDenialsAndSharedCounter(t *testing.T) {
	fixture := newFixture(t, "issue-order")
	unknown := fixture.issue("unknown-run", "read")
	unknown.RunID = "missing"
	if got, err := fixture.f59.Issue(fixture.ctx, unknown); err != nil || got.Kind != AuthorizeReject || got.Reason != RunNotAdmitted || fixture.authorizationCount() != 0 {
		t.Fatalf("unknown run = %#v err=%v count=%d", got, err, fixture.authorizationCount())
	}

	stale := fixture.issue("stale", "read")
	stale.TurnEpoch = "0"
	if got, err := fixture.f59.Issue(fixture.ctx, stale); err != nil || got.Kind != StaleEpoch || fixture.authorizationCount() != 0 {
		t.Fatalf("stale = %#v err=%v count=%d", got, err, fixture.authorizationCount())
	}

	denied := fixture.issue("denied", "network.fetch")
	denied.Operands = Operands{}
	if got, err := fixture.f59.Issue(fixture.ctx, denied); err != nil || got.Kind != DeniedAboveSet {
		t.Fatalf("above set = %#v err=%v", got, err)
	}
	if row := fixture.authorization("denied"); row.State != "VOID" || row.VoidReason != "denied_above_set" {
		t.Fatalf("denial row = %#v", row)
	}
	denialSeq := fixture.stateSeq()
	if got, err := fixture.f59.Issue(fixture.ctx, denied); err != nil || got.Kind != DeniedAboveSet || fixture.authorizationCount() != 1 || fixture.stateSeq() != denialSeq {
		t.Fatalf("denial replay = %#v err=%v count=%d", got, err, fixture.authorizationCount())
	}

	granted := fixture.issue("granted", "bash")
	granted.Operands = Operands{CWD: stringPointer(".")}
	if got, err := fixture.f59.Issue(fixture.ctx, granted); err != nil || got.Kind != TicketGranted {
		t.Fatalf("grant = %#v err=%v", got, err)
	}
	if fixture.authorizationCount() != 2 {
		t.Fatalf("ISSUED and VOID did not share counter: %d", fixture.authorizationCount())
	}
}

func TestIssueLifecycleDenialsCeilingAndLeaseFaultAtomicity(t *testing.T) {
	terminal := newFixture(t, "terminal")
	terminal.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE runs SET state='FAILED' WHERE run_id=?`, terminal.runID)
		return err
	})
	if got, err := terminal.f59.Issue(terminal.ctx, terminal.issue("terminal-call", "read")); err != nil || got.Kind != AuthorizeReject || got.Reason != RunNotAdmitted || terminal.authorization("terminal-call").VoidReason != "run_not_admitted" {
		t.Fatalf("terminal run denial = %#v err=%v", got, err)
	}

	inactive := newFixture(t, "inactive")
	inactive.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED' WHERE turn_id=?`, inactive.turnID)
		return err
	})
	if got, err := inactive.f59.Issue(inactive.ctx, inactive.issue("inactive-call", "read")); err != nil || got.Kind != AuthorizeReject || got.Reason != TurnInactive || inactive.authorization("inactive-call").VoidReason != "turn_inactive" {
		t.Fatalf("inactive turn denial = %#v err=%v", got, err)
	}

	unknownTurn := newFixture(t, "unknown-turn")
	request := unknownTurn.issue("unknown-turn-call", "read")
	request.TurnID = "missing-turn"
	before := unknownTurn.stateSeq()
	if got, err := unknownTurn.f59.Issue(unknownTurn.ctx, request); err != nil || got.Kind != AuthorizeReject || got.Reason != TurnInactive || unknownTurn.authorizationCount() != 0 || unknownTurn.stateSeq() != before {
		t.Fatalf("unknown turn denial = %#v err=%v", got, err)
	}

	lease := newFixture(t, "lease")
	lease.mutate(func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `DELETE FROM leases WHERE run_id=? AND lease_kind='turn'`, lease.runID)
		return err
	})
	if got, err := lease.f59.Issue(lease.ctx, lease.issue("lease-call", "read")); err != nil || got.Kind != AuthorizeReject || got.Reason != LeaseInvalid {
		t.Fatalf("lease denial = %#v err=%v", got, err)
	}
	if row := lease.authorization("lease-call"); row.State != "VOID" || row.VoidReason != "lease_invalid" {
		t.Fatalf("lease fault row = %#v", row)
	}
	if epoch := lease.currentEpoch(); epoch != "2" || lease.turnState() != "INTERRUPTED" || lease.workerState() != "FAILED" {
		t.Fatalf("lease fault not atomic: epoch=%s turn=%s worker=%s", epoch, lease.turnState(), lease.workerState())
	}

	ceiling := newFixture(t, "ceiling")
	ceiling.seedDenied(appipc.MaxToolCallsPerTurn)
	aboveSet := ceiling.issue("over-ceiling", "network.fetch")
	aboveSet.Operands = Operands{}
	before = ceiling.stateSeq()
	if got, err := ceiling.f59.Issue(ceiling.ctx, aboveSet); err != nil || got.Kind != AuthorizeReject || got.Reason != TurnBudgetExhausted || ceiling.authorizationCount() != appipc.MaxToolCallsPerTurn || ceiling.stateSeq() != before {
		t.Fatalf("ceiling precedes above-set = %#v err=%v count=%d", got, err, ceiling.authorizationCount())
	}
}

func TestConsumeTotalOrderAndOneShot(t *testing.T) {
	fixture := newFixture(t, "consume")
	request := fixture.issue("call", "read")
	issued, err := fixture.f59.Issue(fixture.ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	mismatch := ConsumeRequest{TicketID: issued.TicketID, TurnEpoch: "1", CanonicalToolName: "read", CanonicalArgsDigest: digest("changed")}
	if got, err := fixture.consume(mismatch); err != nil || got.Kind != IdentityMismatch || fixture.authorization("call").State != "ISSUED" {
		t.Fatalf("mismatched consume = %#v err=%v row=%#v", got, err, fixture.authorization("call"))
	}
	matched := mismatch
	matched.CanonicalArgsDigest = request.CanonicalArgsDigest
	if got, err := fixture.consume(matched); err != nil || got.Kind != ConsumeOK {
		t.Fatalf("consume = %#v err=%v", got, err)
	}
	if got := fixture.toolCallCount(); got != 0 {
		t.Fatalf("tool_calls materialized at consume: %d", got)
	}
	if got, err := fixture.consume(matched); err != nil || got.Kind != DuplicateConsume {
		t.Fatalf("duplicate consume = %#v err=%v", got, err)
	}
	unknown := matched
	unknown.TicketID = "never-minted"
	if got, err := fixture.consume(unknown); err != nil || got.Kind != NoReply || !got.Fault {
		t.Fatalf("unknown consume = %#v err=%v", got, err)
	}
}

func TestExpireAndOutcomeRecordsKeepCrashTruth(t *testing.T) {
	fixture := newFixture(t, "outcome")
	issuedRequest := fixture.issue("issued", "read")
	issued, _ := fixture.f59.Issue(fixture.ctx, issuedRequest)
	consumedRequest := fixture.issue("consumed", "bash")
	consumedRequest.Operands = Operands{CWD: stringPointer(".")}
	consumed, _ := fixture.f59.Issue(fixture.ctx, consumedRequest)
	consume := ConsumeRequest{TicketID: consumed.TicketID, TurnEpoch: "1", CanonicalToolName: "bash", CanonicalArgsDigest: consumedRequest.CanonicalArgsDigest}
	if _, err := fixture.consume(consume); err != nil {
		t.Fatal(err)
	}
	if _, err := fixture.f59.Expire(fixture.ctx, fixture.runID, fixture.turnID, 99); err != nil {
		t.Fatal(err)
	}
	if row := fixture.authorization("issued"); row.State != "VOID" || row.VoidReason != "expired" {
		t.Fatalf("issued expiry = %#v ticket=%s", row, issued.TicketID)
	}
	if row := fixture.authorization("consumed"); row.State != "UNKNOWN_TOOL_OUTCOME" {
		t.Fatalf("consumed expiry = %#v", row)
	}

	executedFixture := newFixture(t, "executed")
	req := executedFixture.issue("executed", "read")
	ticket, _ := executedFixture.f59.Issue(executedFixture.ctx, req)
	_, _ = executedFixture.consume(ConsumeRequest{TicketID: ticket.TicketID, TurnEpoch: "1", CanonicalToolName: "read", CanonicalArgsDigest: req.CanonicalArgsDigest})
	identity := Identity{CanonicalToolName: "read", CanonicalArgsDigest: req.CanonicalArgsDigest, TurnEpoch: "1"}
	result, err := executedFixture.recordOutcome(OutcomeRequest{TicketID: ticket.TicketID, TurnEpoch: "1", Outcome: Executed, InvocationIdentity: &identity})
	if err != nil || result.Kind != NoReply || result.Fault {
		t.Fatalf("executed outcome = %#v err=%v", result, err)
	}
	if row := executedFixture.authorization("executed"); row.State != "OUTCOME_RECORDED" {
		t.Fatalf("outcome row = %#v", row)
	}
	duplicate, err := executedFixture.recordOutcome(OutcomeRequest{TicketID: ticket.TicketID, TurnEpoch: "1", Outcome: Executed, InvocationIdentity: &identity})
	if err != nil || duplicate.Fault || !duplicate.Idempotent {
		t.Fatalf("outcome duplicate = %#v err=%v", duplicate, err)
	}

	faultFixture := newFixture(t, "integrity")
	faultReq := faultFixture.issue("integrity", "edit")
	faultTicket, _ := faultFixture.f59.Issue(faultFixture.ctx, faultReq)
	_, _ = faultFixture.consume(ConsumeRequest{TicketID: faultTicket.TicketID, TurnEpoch: "1", CanonicalToolName: "edit", CanonicalArgsDigest: faultReq.CanonicalArgsDigest})
	expected := Identity{CanonicalToolName: "edit", CanonicalArgsDigest: faultReq.CanonicalArgsDigest, TurnEpoch: "1"}
	observed := expected
	observed.CanonicalArgsDigest = digest("post-consume-mutation")
	result, err = faultFixture.recordOutcome(OutcomeRequest{TicketID: faultTicket.TicketID, TurnEpoch: "1", Outcome: NotInvokedIntegrityFault, IntegrityEvidence: &IntegrityEvidence{Expected: expected, Observed: observed}})
	if err != nil || result.Fault {
		t.Fatalf("integrity outcome = %#v err=%v", result, err)
	}
	if state := faultFixture.toolCallState("integrity"); state != "NOT_INVOKED_INTEGRITY_FAULT" {
		t.Fatalf("tool call state = %q", state)
	}
}

type fixture struct {
	t             *testing.T
	ctx           context.Context
	db            *store.Store
	applier       *applier.Host
	f59           *Host
	frozen        manifest.Frozen
	gate          manifest.Gate
	runID         string
	turnID        string
	generationID  string
	ticketCounter int
}

func (fixture *fixture) consume(request ConsumeRequest) (Decision, error) {
	return fixture.f59.Consume(fixture.ctx, ChannelIdentity{GenerationID: fixture.generationID}, request)
}

func (fixture *fixture) recordOutcome(request OutcomeRequest) (Decision, error) {
	return fixture.f59.RecordOutcome(fixture.ctx, ChannelIdentity{GenerationID: fixture.generationID}, request)
}

func newFixture(t *testing.T, suffix string) *fixture {
	t.Helper()
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	root := filepath.Join(t.TempDir(), "workspace")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	tools := lockedTools()
	policy := []byte(`{"pinned_lane":"fixture"}`)
	catalog := digest("catalog")
	input := manifest.BuildInput{
		RunID: "run-" + suffix, PolicySourceRef: "ratification", PolicyDigest: digestBytes(policy), PolicyBytes: policy,
		PolicyPinnedLane: manifest.LaneID{ModelID: "model", ProviderID: "provider", ServingProfileID: "profile", CompatMode: "native"},
		ToolSet:          tools, ToolCatalogDigest: &catalog,
		ProviderLane:   manifest.ProviderLane{LaneID: manifest.LaneID{ModelID: "model", ProviderID: "provider", ServingProfileID: "profile", CompatMode: "native"}, LaneCatalogDigest: digest("lanes"), CredentialRef: "opaque"},
		WorkspaceRoot:  root,
		ReleaseBinding: &manifest.ReleaseBinding{BoundAtRef: "release", ReleaseDigest: stringPointer(digest("release"))},
	}
	frozen, err := manifest.Build(input)
	if err != nil {
		t.Fatal(err)
	}
	gate := manifest.Gate{LockedTools: tools, ShippedToolCatalogDigest: catalog, PolicyBytes: policy, LaneCatalogDigest: digest("lanes")}
	if _, err := host.Apply(ctx, manifest.FreezeEvent{Frozen: frozen, Gate: gate, SessionLogPath: "/session.log", CreatedAt: 1}); err != nil {
		t.Fatal(err)
	}
	result := &fixture{t: t, ctx: ctx, db: db, applier: host, frozen: frozen, gate: gate, runID: input.RunID, turnID: "turn-" + suffix, generationID: "generation-" + suffix}
	result.mutate(func(ctx context.Context, tx *store.Tx) error {
		statements := []struct {
			query string
			args  []any
		}{
			{`UPDATE runs SET state='ACTIVE',run_phase='established' WHERE run_id=?`, []any{result.runID}},
			{`INSERT INTO workers(generation_id,run_id,turn_epoch,pid,state,created_at) VALUES(?,?,?,?,?,?)`, []any{result.generationID, result.runID, storeCounter(1), 42, "LEASED", 2}},
			{`INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, []any{result.turnID, result.runID, storeCounter(1), "ACTIVE", []byte("task"), "fresh", strings.Repeat("a", 32), "PENDING", 2}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{result.runID, "worker", "worker-lease", result.generationID, storeCounter(1), "ACTIVE", 2}},
			{`INSERT INTO leases(run_id,lease_kind,lease_id,generation_id,turn_epoch,state,granted_at) VALUES(?,?,?,?,?,?,?)`, []any{result.runID, "turn", "turn-lease", result.generationID, storeCounter(1), "ACTIVE", 2}},
		}
		for _, statement := range statements {
			if _, err := tx.ExecContext(ctx, statement.query, statement.args...); err != nil {
				return err
			}
		}
		return nil
	})
	result.f59 = New(host, Config{
		Gate: gate,
		NewTicketID: func() string {
			result.ticketCounter++
			return fmt.Sprintf("ticket-%s-%d", suffix, result.ticketCounter)
		},
		LeaseInvalid: func(ctx context.Context, tx *store.Tx, fault LeaseFault) error {
			if _, err := tx.ExecContext(ctx, `UPDATE workers SET state='FAILED',updated_at=? WHERE generation_id=?`, fault.At, fault.GenerationID); err != nil {
				return err
			}
			if _, err := tx.ExecContext(ctx, `UPDATE leases SET state='RELEASED',released_at=? WHERE run_id=? AND state='ACTIVE'`, fault.At, fault.RunID); err != nil {
				return err
			}
			if _, err := tx.ExecContext(ctx, `UPDATE turns SET state='INTERRUPTED',updated_at=? WHERE turn_id=?`, fault.At, fault.TurnID); err != nil {
				return err
			}
			if _, err := ParkOpen(ctx, tx, fault.RunID, fault.TurnID, fault.At); err != nil {
				return err
			}
			_, err := tx.ExecContext(ctx, `UPDATE epochs SET turn_epoch=? WHERE run_id=?`, storeCounter(2), fault.RunID)
			return err
		},
	})
	return result
}

func (fixture *fixture) issue(callID, tool string) IssueRequest {
	resource := "src/main.go"
	cwd := "."
	return IssueRequest{RunID: fixture.runID, TurnID: fixture.turnID, TurnEpoch: "1", ToolCallID: callID,
		CanonicalToolName: tool, CanonicalArgsDigest: digest("args:" + callID), GenerationID: fixture.generationID,
		Operands: Operands{CanonicalResource: &resource, CWD: &cwd}, IssuedAt: 10}
}

func (fixture *fixture) mutate(apply func(context.Context, *store.Tx) error) {
	fixture.t.Helper()
	_, err := fixture.applier.Apply(fixture.ctx, testEvent{runID: fixture.runID, apply: apply})
	if err != nil {
		fixture.t.Fatal(err)
	}
}

type testEvent struct {
	runID string
	apply func(context.Context, *store.Tx) error
}

func (event testEvent) RunID() string { return event.runID }
func (event testEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, event.apply(ctx, tx)
}

type authorizationRow struct{ State, VoidReason string }

func (fixture *fixture) authorization(callID string) authorizationRow {
	fixture.t.Helper()
	value, err := fixture.applier.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var state string
		var reason *string
		if err := snapshot.QueryRowContext(ctx, `SELECT state,void_reason FROM tool_authorizations WHERE tool_call_id=?`, callID).Scan(&state, &reason); err != nil {
			return nil, err
		}
		row := authorizationRow{State: state}
		if reason != nil {
			row.VoidReason = *reason
		}
		return row, nil
	}))
	if err != nil {
		fixture.t.Fatal(err)
	}
	return value.(authorizationRow)
}

func (fixture *fixture) authorizationCount() int {
	fixture.t.Helper()
	value, err := fixture.applier.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var count int
		err := snapshot.QueryRowContext(ctx, `SELECT COUNT(*) FROM tool_authorizations WHERE run_id=?`, fixture.runID).Scan(&count)
		return count, err
	}))
	if err != nil {
		fixture.t.Fatal(err)
	}
	return value.(int)
}

func (fixture *fixture) stateSeq() uint64 {
	fixture.t.Helper()
	value, err := fixture.applier.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var counter string
		err := snapshot.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, fixture.runID).Scan(&counter)
		if err != nil {
			return uint64(0), err
		}
		wire, err := appipc.UnpadCounter(counter)
		if err != nil {
			return uint64(0), err
		}
		return appipc.ParseCounter(wire)
	}))
	if err != nil {
		fixture.t.Fatal(err)
	}
	return value.(uint64)
}

func (fixture *fixture) toolCallState(callID string) string {
	fixture.t.Helper()
	value, err := fixture.applier.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var state string
		err := snapshot.QueryRowContext(ctx, `SELECT state FROM tool_calls WHERE tool_call_id=?`, callID).Scan(&state)
		return state, err
	}))
	if err != nil {
		fixture.t.Fatal(err)
	}
	return value.(string)
}

func (fixture *fixture) currentEpoch() string {
	return fixture.readString(`SELECT turn_epoch FROM epochs WHERE run_id=?`, fixture.runID, true)
}

func (fixture *fixture) turnState() string {
	return fixture.readString(`SELECT state FROM turns WHERE turn_id=?`, fixture.turnID, false)
}

func (fixture *fixture) workerState() string {
	return fixture.readString(`SELECT state FROM workers WHERE generation_id=?`, fixture.generationID, false)
}

func (fixture *fixture) readString(query string, arg any, counter bool) string {
	fixture.t.Helper()
	value, err := fixture.applier.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var value string
		err := snapshot.QueryRowContext(ctx, query, arg).Scan(&value)
		return value, err
	}))
	if err != nil {
		fixture.t.Fatal(err)
	}
	result := value.(string)
	if counter {
		wire, err := appipc.UnpadCounter(result)
		if err != nil {
			fixture.t.Fatal(err)
		}
		return wire
	}
	return result
}

func (fixture *fixture) seedDenied(count int) {
	fixture.t.Helper()
	fixture.mutate(func(ctx context.Context, tx *store.Tx) error {
		for index := 0; index < count; index++ {
			callID := fmt.Sprintf("seed-call-%03d", index)
			ticketID := fmt.Sprintf("seed-ticket-%03d", index)
			argsDigest := digest(callID)
			if _, err := tx.ExecContext(ctx, `INSERT INTO tool_authorizations(ticket_id,run_id,turn_id,tool_call_id,turn_epoch,state,void_reason,canonical_tool_name,canonical_args_digest,effect_descriptor,issued_at) VALUES(?,?,?,?,?,?,?,?,?,?,?)`, ticketID, fixture.runID, fixture.turnID, callID, storeCounter(1), "VOID", "denied_above_set", "network.fetch", argsDigest, []byte("{}"), 3); err != nil {
				return err
			}
		}
		return nil
	})
}

func (fixture *fixture) toolCallCount() int {
	fixture.t.Helper()
	value, err := fixture.applier.Read(fixture.ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var count int
		err := snapshot.QueryRowContext(ctx, `SELECT COUNT(*) FROM tool_calls WHERE run_id=?`, fixture.runID).Scan(&count)
		return count, err
	}))
	if err != nil {
		fixture.t.Fatal(err)
	}
	return value.(int)
}

func lockedTools() []manifest.ToolIdentity {
	tools := manifest.StagingToolSet()
	for index := range tools {
		schema := digest("schema:" + tools[index].Name)
		catalog := "catalog-v1"
		tools[index].SchemaDigest = &schema
		tools[index].CatalogVersion = &catalog
		if len(tools[index].Name) > 6 && tools[index].Name[:6] == "relay." {
			mapping := "mapping-v1"
			tools[index].MappingVersion = &mapping
		}
	}
	return tools
}

func digest(value string) string         { return digestBytes([]byte(value)) }
func stringPointer(value string) *string { return &value }
func storeCounter(value uint64) string   { return fmt.Sprintf("%020d", value) }
