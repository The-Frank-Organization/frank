package fixtures_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/derived"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestH16RetryCeilingParksAndReopenGetsFreshCeiling(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		rec := record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": cmd.IntakeID},
		}
		if cmd.IntakeID == "ceiling-source" {
			rec.Headers["resolves_gate"] = "ceiling-gate"
		}
		if strings.HasPrefix(cmd.IntakeID, "reopen-") {
			rec.Headers["record_kind"] = "attempt_resolution"
			rec.Body = fmt.Sprintf(`{"resolves":%q}`, strings.TrimPrefix(cmd.IntakeID, "reopen-"))
		}
		return rec, nil, nil
	}, engine.TestReady())
	gateCalls := 0
	loop.AfterGateResolution = func(record.Record) error {
		gateCalls++
		return errors.New("injected retry failure")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	sourceCmd := intake.Cmd{IntakeID: "ceiling-source", Seat: "operator", Role: "operator", Verb: "submit"}
	var sourceOut engine.Outcome
	for attempt := 1; attempt <= engine.DerivedRetryCeiling; attempt++ {
		sourceOut = h16SubmitLoop(t, loop, sourceCmd)
	}
	assertH16PostCommitState(t, sourceOut, record.Accepted, "failed", false)
	firstPark := h16OnlyUnresolvedPark(t, st, sourceOut.RelayID)
	if gateCalls != engine.DerivedRetryCeiling {
		t.Fatalf("gate calls=%d, want ceiling=%d", gateCalls, engine.DerivedRetryCeiling)
	}

	reopen := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "reopen-" + firstPark, Seat: "operator", Role: "operator", IsOperator: true, Verb: "submit"})
	assertH16PostCommitState(t, reopen, record.Accepted, "complete", true)
	if got := derived.Fold(h16Records(t, st))[sourceOut.RelayID].Status; got != "pending" {
		t.Fatalf("status after reopen=%q, want pending", got)
	}

	for attempt := 1; attempt <= engine.DerivedRetryCeiling; attempt++ {
		sourceOut = h16SubmitLoop(t, loop, sourceCmd)
	}
	assertH16PostCommitState(t, sourceOut, record.Accepted, "failed", false)
	secondPark := h16OnlyUnresolvedPark(t, st, sourceOut.RelayID)
	if secondPark == firstPark {
		t.Fatalf("reopened work reused park instance %q", firstPark)
	}
	if gateCalls != 2*engine.DerivedRetryCeiling {
		t.Fatalf("gate calls after reopen=%d, want %d", gateCalls, 2*engine.DerivedRetryCeiling)
	}
}

func TestH16AttemptResolutionOneShotClassesCommitAsCanonicalAnomalies(t *testing.T) {
	t.Run("duplicate and conflicting", func(t *testing.T) {
		st := h16TransitionStore(t, false)
		loop, cancel := h16TransitionLoop(t, st)
		defer cancel()
		first := h16TransitionSubmit(t, loop, "resolution-first", `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized","evidence_ref":"E1"}`)
		assertH16PostCommitState(t, first, record.Accepted, "complete", true)
		duplicate := h16TransitionSubmit(t, loop, "resolution-duplicate", `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized","evidence_ref":"E1"}`)
		h16AssertRejectedClass(t, duplicate, "resolves:duplicate-resolution")
		conflict := h16TransitionSubmit(t, loop, "resolution-conflict", `{"resolves":"marker-a","disposition":"effect-confirmed-realized","evidence_ref":"E2"}`)
		h16AssertRejectedClass(t, conflict, "resolves:conflicting-resolution")
		if !(conflict.RelayID < first.RelayID) {
			t.Fatalf("later conflicting relay %q must sort before first %q for adversarial-order proof", conflict.RelayID, first.RelayID)
		}
		records := h16Records(t, st)
		if got := derived.Fold(records)["transition-source"].Status; got != "pending" {
			t.Fatalf("rejected anomalies changed fold: %q", got)
		}
		reverse := append([]record.Record(nil), records...)
		for left, right := 0, len(reverse)-1; left < right; left, right = left+1, right-1 {
			reverse[left], reverse[right] = reverse[right], reverse[left]
		}
		if !reflect.DeepEqual(derived.Fold(records), derived.Fold(reverse)) {
			t.Fatal("shuffled rebuild changed fold with canonical anomaly rows")
		}
	})

	t.Run("target-specific body shape", func(t *testing.T) {
		for _, tc := range []struct {
			name string
			body string
		}{
			{name: "marker missing evidence", body: `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized"}`},
			{name: "marker extra member", body: `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized","evidence_ref":"E1","extra":"no"}`},
		} {
			t.Run(tc.name, func(t *testing.T) {
				st := h16TransitionStore(t, false)
				loop, cancel := h16TransitionLoop(t, st)
				defer cancel()
				out := h16TransitionSubmit(t, loop, "shape-"+strings.ReplaceAll(tc.name, " ", "-"), tc.body)
				h16AssertRejectedClass(t, out, "body:typed")
			})
		}
	})

	t.Run("stale and unknown", func(t *testing.T) {
		st := h16TransitionStore(t, true)
		loop, cancel := h16TransitionLoop(t, st)
		defer cancel()
		stale := h16TransitionSubmit(t, loop, "resolution-stale", `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized","evidence_ref":"E1"}`)
		h16AssertRejectedClass(t, stale, "resolves:stale-resolution")
		unknown := h16TransitionSubmit(t, loop, "resolution-unknown", `{"resolves":"missing","disposition":"effect-confirmed-unrealized","evidence_ref":"E1"}`)
		h16AssertRejectedClass(t, unknown, "resolves:unknown-target")
	})
}

func TestH16FreshAttemptInstanceSurvivesShuffledRebuild(t *testing.T) {
	source := h16DerivedSource("source", "1", "seat_mint")
	markerA := h16DerivedRecord("marker-a", "derived-work-attempt", map[string]any{
		"source_relay_id": "source", "hook": "mint", "state": "running_or_unknown", "predecessor": "none",
	})
	resolutionA := h16DerivedRecord("z-resolution-a", "attempt_resolution", map[string]any{
		"resolves": "marker-a", "disposition": "effect-confirmed-unrealized", "evidence_ref": "E1",
	})
	if got := derived.Fold([]record.Record{source, markerA, resolutionA})["source"].Status; got != "pending" {
		t.Fatalf("resolved attempt status=%q, want pending", got)
	}
	markerB := h16DerivedRecord("a-marker-b", "derived-work-attempt", map[string]any{
		"source_relay_id": "source", "hook": "mint", "state": "running_or_unknown", "predecessor": "z-resolution-a",
	})
	forward := []record.Record{source, markerA, resolutionA, markerB}
	reverse := []record.Record{markerB, resolutionA, markerA, source}
	want := derived.WorkStatus{Cursor: []string{"mint"}, Status: "unknown"}
	if got := derived.Fold(forward)["source"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("forward fold=%+v, want %+v", got, want)
	}
	if got := derived.Fold(reverse)["source"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("reverse fold=%+v, want %+v", got, want)
	}
	conflictingResolution := h16DerivedRecord("b-conflicting-resolution", "attempt_resolution", map[string]any{
		"resolves": "marker-a", "disposition": "effect-confirmed-realized", "evidence_ref": "E2",
	})
	conflicted := []record.Record{source, markerA, resolutionA, conflictingResolution}
	if got := derived.Fold(conflicted)["source"].Status; got != "unknown" {
		t.Fatalf("belt-and-suspenders conflicting accepted resolutions fold=%q, want unknown", got)
	}
	for left, right := 0, len(conflicted)-1; left < right; left, right = left+1, right-1 {
		conflicted[left], conflicted[right] = conflicted[right], conflicted[left]
	}
	if got := derived.Fold(conflicted)["source"].Status; got != "unknown" {
		t.Fatalf("reversed conflict fold=%q, want unknown", got)
	}
}

func TestH16AttemptResolutionAuthorityUsesStampedSeatNotClaimedRole(t *testing.T) {
	for _, target := range []struct {
		name string
		body string
	}{
		{name: "park", body: `{"resolves":"park-a"}`},
		{name: "marker", body: `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized","evidence_ref":"E-authority"}`},
	} {
		for _, claimedRole := range []string{"implementer", "operator"} {
			t.Run(target.name+"-claimed-"+claimedRole, func(t *testing.T) {
				st := h16AuthorityTransitionStore(t)
				worker := seat.SeatMeta{Name: "worker.implementer", Role: "implementer"}
				out := h16GovernedResolution(t, st, worker, claimedRole, "worker-"+target.name+"-"+claimedRole, target.body)
				h16AssertRejectedClass(t, out, "record_kind:seat-scope")
				stored, err := st.Read(out.RelayID)
				if err != nil || stored.Headers["SUBJECT"] != "authority-evidence" {
					t.Fatalf("rejected evidence not preserved: rec=%+v err=%v", stored, err)
				}
			})
		}
	}

	for _, target := range []struct {
		name   string
		body   string
		source string
		want   string
	}{
		{name: "park", body: `{"resolves":"park-a"}`, source: "park-source", want: "pending"},
		{name: "marker", body: `{"resolves":"marker-a","disposition":"effect-confirmed-unrealized","evidence_ref":"E-authority"}`, source: "marker-source", want: "pending"},
	} {
		t.Run("operator-"+target.name, func(t *testing.T) {
			st := h16AuthorityTransitionStore(t)
			operator := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
			out := h16GovernedResolution(t, st, operator, "operator", "operator-"+target.name, target.body)
			assertH16PostCommitState(t, out, record.Accepted, "complete", true)
			if got := derived.Fold(h16Records(t, st))[target.source].Status; got != target.want {
				t.Fatalf("operator %s fold=%q, want %q", target.name, got, target.want)
			}
		})
	}
}

func TestH16PostCommitPanicReportsDurableFoldAndPreCommitPanicStaysRejected(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		if cmd.IntakeID == "pre-panic" {
			panic("pre-commit")
		}
		return record.Record{
			Envelope: record.Envelope{From: cmd.Seat, Role: cmd.Role, DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "panic source", "resolves_gate": "panic-gate"},
		}, nil, nil
	}, engine.TestReady())
	postPanicked := false
	loop.AfterGateResolution = func(record.Record) error {
		if !postPanicked {
			postPanicked = true
			panic("post-commit")
		}
		return nil
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	post := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "post-panic", Seat: "operator", Role: "operator"})
	assertH16PostCommitState(t, post, record.Accepted, "pending", false)
	pre := h16SubmitLoop(t, loop, intake.Cmd{IntakeID: "pre-panic", Seat: "operator", Role: "operator"})
	assertH16PostCommitState(t, pre, record.Rejected, "complete", true)
}

func h16TransitionStore(t *testing.T, terminal bool) *store.Store {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	h16Commit(t, st, h16DerivedSource("transition-source", "1", "seat_mint"))
	h16Commit(t, st, h16DerivedRecord("marker-a", "derived-work-attempt", map[string]any{
		"source_relay_id": "transition-source", "hook": "mint", "state": "running_or_unknown", "predecessor": "none",
	}))
	if terminal {
		h16Commit(t, st, h16DerivedRecord("advance-a", "derived-work-transition", map[string]any{
			"source_relay_id": "transition-source", "kind": "cursor_advance", "completed_hooks": []string{"mint"},
		}))
	}
	return st
}

func h16TransitionLoop(t *testing.T, st *store.Store) (*engine.Loop, context.CancelFunc) {
	t.Helper()
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{RelayID: cmd.IntakeID, From: "operator", Role: "operator", DeliveryState: record.Accepted, IntakeID: cmd.IntakeID, SchemaVersion: 1},
			Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": cmd.IntakeID, "record_kind": "attempt_resolution"},
			Body:     string(cmd.Payload),
		}, nil, nil
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	go loop.Run(ctx)
	return loop, cancel
}

func h16TransitionSubmit(t *testing.T, loop *engine.Loop, intakeID, body string) engine.Outcome {
	t.Helper()
	return h16SubmitLoop(t, loop, intake.Cmd{IntakeID: intakeID, Seat: "operator", Role: "operator", IsOperator: true, Payload: []byte(body)})
}

func h16AuthorityTransitionStore(t *testing.T) *store.Store {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	parkSource := h16DerivedSource("park-source", "1", "gate")
	markerSource := h16DerivedSource("marker-source", "1", "seat_mint")
	h16Commit(t, st, parkSource)
	h16Commit(t, st, markerSource)
	h16Commit(t, st, h16DerivedRecord("park-a", "derived-work-transition", map[string]any{
		"source_relay_id": "park-source", "kind": "parked", "reason": "retry-ceiling",
	}))
	h16Commit(t, st, h16DerivedRecord("marker-a", "derived-work-attempt", map[string]any{
		"source_relay_id": "marker-source", "hook": "mint", "state": "running_or_unknown", "predecessor": "none",
	}))
	return st
}

func h16GovernedResolution(t *testing.T, st *store.Store, meta seat.SeatMeta, claimedRole, intakeID, body string) engine.Outcome {
	t.Helper()
	reg := loadH16Registry(t)
	rec := h16PresenceCandidate()
	rec.Headers["SUBJECT"] = "authority-evidence"
	rec.Headers["record_kind"] = "attempt_resolution"
	rec.Body = body
	payload := mustJSONBytes(t, submitPayloadForRegistry(reg, meta, rec))
	loop := engine.New(st, engine.SubmitHandler(st, reg, meta), engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)
	return h16SubmitLoop(t, loop, intake.Cmd{IntakeID: intakeID, Seat: meta.Name, Role: claimedRole, IsOperator: claimedRole == "operator", Verb: "submit", Payload: payload})
}

func h16OnlyUnresolvedPark(t *testing.T, st *store.Store, sourceRelayID string) string {
	t.Helper()
	records := h16Records(t, st)
	resolved := map[string]bool{}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "attempt_resolution" {
			continue
		}
		var body struct {
			Resolves string `json:"resolves"`
		}
		_ = json.Unmarshal([]byte(rec.Body), &body)
		resolved[body.Resolves] = true
	}
	parks := []string{}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "derived-work-transition" || resolved[rec.Envelope.RelayID] {
			continue
		}
		var body struct {
			SourceRelayID string `json:"source_relay_id"`
			Kind          string `json:"kind"`
		}
		_ = json.Unmarshal([]byte(rec.Body), &body)
		if body.SourceRelayID == sourceRelayID && body.Kind == "parked" {
			parks = append(parks, rec.Envelope.RelayID)
		}
	}
	if len(parks) != 1 {
		t.Fatalf("unresolved parks=%v, want exactly one", parks)
	}
	return parks[0]
}

func h16Records(t *testing.T, st *store.Store) []record.Record {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("records: %v", err)
	}
	return records
}

func h16Commit(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, nil); err != nil {
		t.Fatalf("commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func h16AssertRejectedClass(t *testing.T, out engine.Outcome, class string) {
	t.Helper()
	assertH16PostCommitState(t, out, record.Rejected, "complete", true)
	if !strings.Contains(out.Detail, class) {
		t.Fatalf("rejection detail=%q, want %q", out.Detail, class)
	}
}
