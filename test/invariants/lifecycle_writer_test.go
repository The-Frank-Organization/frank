package invariants_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

func TestLawDerivedOnlyActivation(t *testing.T) {
	_, law := requireCatalogLaw(t, "TestLawDerivedOnlyActivation")
	for _, required := range []string{
		"minted and active derive from committed records in commit order",
		"bound_now is runtime-only",
		"no activation or bound marker persists",
	} {
		if !strings.Contains(law.Claim, required) {
			t.Fatalf("row-4 claim %q missing %q", law.Claim, required)
		}
	}

	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open lifecycle store: %v", err)
	}
	const seatName = "life-seat.implementer"
	mintBody, err := json.Marshal(map[string]any{
		"seat":        seatName,
		"role":        "implementer",
		"is_operator": false,
	})
	if err != nil {
		t.Fatalf("marshal seat mint: %v", err)
	}
	commitRecord(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "mint-life-seat",
			From:          "operator",
			Role:          "operator",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":       "SITREP",
			"SUBJECT":     "mint life seat",
			"record_kind": "seat_mint",
		},
		Body: string(mintBody),
	})

	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("build minted lifecycle: %v", err)
	}
	minted := tab.Lifecycle[seatName]
	if minted.ActivationState != tables.LifecycleMinted ||
		minted.MintedAt != "mint-life-seat" ||
		minted.ActivationRecordRef != "" {
		t.Fatalf("minted lifecycle = %+v", minted)
	}

	manager, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open seat manager: %v", err)
	}
	credential, err := manager.MintOrReplace(seatName, "implementer", false, "mint-life-seat")
	if err != nil {
		t.Fatalf("realize seat mint: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s7-life-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	tools := channel.ToolSet{
		Submit:  staticToolResult(`{"state":"accepted"}`),
		Project: staticToolResult(`[]`),
		Read:    staticToolResult(`{"relay_id":"mint-life-seat"}`),
	}
	server, err := channel.ServeAuthenticated(sock, manager, func(seat.SeatMeta) channel.ToolSet {
		return channel.FullSurface(engine.TestReady(), tools)
	})
	if err != nil {
		t.Fatalf("serve lifecycle channel: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := channel.DialAuthenticated(ctx, sock, credential.Value)
	if err != nil {
		_ = server.Close()
		t.Fatalf("bind life seat: %v", err)
	}
	if _, err := client.ListTools(ctx); err != nil {
		_ = server.Close()
		t.Fatalf("list bound-seat tools: %v", err)
	}

	mintedRoster := rosterFor(t, tab, manager, server.ActiveSeats(), seatName)
	if mintedRoster.ActivationState != tables.LifecycleMinted || !mintedRoster.BoundNow {
		t.Fatalf("minted+bound roster = %+v", mintedRoster)
	}

	reg := loadRegistry(t)
	meta := seat.SeatMeta{Name: seatName, Role: "implementer"}
	bootCandidate := record.Record{Headers: map[string]string{
		"PHASE":           "SITREP",
		"CEREMONY_TIER":   "small",
		"SUBJECT":         "boot life seat",
		"charter_loaded":  "yes",
		"dispatch_status": "read",
	}}
	bootEnv := fieldspec.RenderEnv{PreActive: true}
	_, bootDigest := reg.Render(
		bootEnv,
		fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role},
		bootCandidate.Headers["PHASE"],
		bootCandidate.Headers["CEREMONY_TIER"],
		fieldspec.ClosedGrantState,
	)
	bootPayload, err := json.Marshal(fieldspec.SubmitPayload{Record: bootCandidate, FormDigest: bootDigest})
	if err != nil {
		t.Fatalf("marshal boot candidate: %v", err)
	}
	bootRecord, bootIntents, err := engine.SubmitHandlerWithRender(st, reg, meta, bootEnv, tab)(
		context.Background(),
		intake.Cmd{
			IntakeID: "intake-life-boot",
			Seat:     seatName,
			Role:     "implementer",
			Verb:     "submit",
			Payload:  bootPayload,
		},
	)
	if err != nil {
		t.Fatalf("run pre-active boot admission: %v", err)
	}
	if bootRecord.Envelope.DeliveryState != record.Accepted || bootRecord.Envelope.RelayID == "" {
		t.Fatalf("boot admission record = %+v", bootRecord)
	}
	bootRelayID := bootRecord.Envelope.RelayID
	if _, err := st.Commit(bootRecord, bootIntents); err != nil {
		t.Fatalf("commit admitted boot %s: %v", bootRelayID, err)
	}
	commitRecord(t, st, record.Record{
		Envelope: record.Envelope{
			RelayID:       "after-life-seat",
			From:          seatName,
			Role:          "implementer",
			DeliveryState: record.Accepted,
			IntakeID:      "intake-life-after",
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "after activation"},
	})
	tab, err = tables.Build(st)
	if err != nil {
		t.Fatalf("rebuild active lifecycle: %v", err)
	}
	active := tab.Lifecycle[seatName]
	if active.ActivationState != tables.LifecycleActive ||
		active.ActivationRecordRef != bootRelayID ||
		active.LastAcceptedAt != "after-life-seat" {
		t.Fatalf("active lifecycle = %+v", active)
	}
	activeRoster := rosterFor(t, tab, manager, server.ActiveSeats(), seatName)
	if !activeRoster.BoundNow || activeRoster.ActivationState != tables.LifecycleActive {
		t.Fatalf("active+bound roster = %+v", activeRoster)
	}

	if err := server.Close(); err != nil {
		t.Fatalf("close lifecycle server: %v", err)
	}
	recovered, err := tables.Build(st)
	if err != nil {
		t.Fatalf("recover lifecycle tables: %v", err)
	}
	disconnected := rosterFor(t, recovered, manager, map[string]bool{}, seatName)
	if disconnected.ActivationState != tables.LifecycleActive || disconnected.BoundNow {
		t.Fatalf("active-but-disconnected roster = %+v", disconnected)
	}
	if disconnected.ActivationRecordRef != bootRelayID {
		t.Fatalf("recovered activation ref = %q, want %s", disconnected.ActivationRecordRef, bootRelayID)
	}

	assertNoLifecycleMarker(t, st, reg)
}

func TestLawSoleGovernedWriter(t *testing.T) {
	_, law := requireCatalogLaw(t, "TestLawSoleGovernedWriter")
	for _, required := range []string{"sole governed write path", "D5 same-uid direct-store"} {
		if !strings.Contains(law.Claim, required) {
			t.Fatalf("row-5 claim %q missing %q", law.Claim, required)
		}
	}

	lockRoot := t.TempDir()
	first, err := store.AcquireRoot(lockRoot)
	if err != nil {
		t.Fatalf("acquire first root owner: %v", err)
	}
	defer func() { _ = first.Release() }()
	_, err = store.AcquireRoot(lockRoot)
	var locked store.ErrRootLocked
	if !errors.As(err, &locked) || locked.ErrorClass != "root-lock-held" {
		t.Fatalf("second root owner error = %#v, want root-lock-held", err)
	}

	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open governed writer store: %v", err)
	}
	journal, err := intake.Open(root)
	if err != nil {
		t.Fatalf("open governed writer journal: %v", err)
	}
	ready := engine.TestReady()
	writer, err := intake.NewWriter[engine.Outcome](journal, config.EngineConfig{}, ready)
	if err != nil {
		t.Fatalf("new governed writer: %v", err)
	}
	loop := engine.New(st, func(_ context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{
			Envelope: record.Envelope{
				From:          cmd.Seat,
				Role:          cmd.Role,
				DeliveryState: record.Accepted,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "serialized mutation"},
		}, nil, nil
	}, ready)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)
	go writer.Run(ctx, loop.In)

	reply, intakeID, err := writer.Submit(ctx, intake.Cmd{
		Seat:    "s7.implementer",
		Role:    "implementer",
		Verb:    "submit",
		Payload: json.RawMessage(`{"subject":"serialized mutation"}`),
	})
	if err != nil {
		t.Fatalf("submit governed mutation: %v", err)
	}
	out := waitOutcome(t, reply)
	if out.State != record.Accepted || out.IntakeID != intakeID || out.RelayID == "" {
		t.Fatalf("governed mutation outcome = %+v, intake=%s", out, intakeID)
	}
	records, err := st.Records()
	if err != nil {
		t.Fatalf("read governed records: %v", err)
	}
	if len(records) != 1 ||
		records[0].Envelope.IntakeID != intakeID ||
		records[0].Envelope.RelayID != out.RelayID {
		t.Fatalf("governed records = %+v, outcome=%+v", records, out)
	}
}

func staticToolResult(value string) channel.ToolFunc {
	return func(context.Context, json.RawMessage) (json.RawMessage, error) {
		return json.RawMessage(value), nil
	}
}

func commitRecord(t *testing.T, st *store.Store, rec record.Record) {
	t.Helper()
	if _, err := st.Commit(rec, []store.Intent{}); err != nil {
		t.Fatalf("commit %s: %v", rec.Envelope.RelayID, err)
	}
}

func rosterFor(t *testing.T, tab *tables.T, manager *seat.Manager, bound map[string]bool, seatName string) tables.RosterRow {
	t.Helper()
	for _, row := range tables.Roster(tab, manager.Seats(), bound) {
		if row.Seat == seatName {
			return row
		}
	}
	t.Fatalf("roster missing %s", seatName)
	return tables.RosterRow{}
}

func assertNoLifecycleMarker(t *testing.T, st *store.Store, reg *fieldspec.Registry) {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("read lifecycle records: %v", err)
	}
	for _, rec := range records {
		if kind := strings.ToLower(rec.Headers["record_kind"]); strings.Contains(kind, "activation") || strings.Contains(kind, "bound") {
			t.Fatalf("persisted lifecycle marker record %s kind=%q", rec.Envelope.RelayID, kind)
		}
		for key := range rec.Headers {
			lower := strings.ToLower(key)
			if strings.Contains(lower, "activation_marker") || lower == "bound" || lower == "bound_now" {
				t.Fatalf("persisted lifecycle marker header %s on %s", key, rec.Envelope.RelayID)
			}
		}
		body := strings.ToLower(rec.Body)
		if strings.Contains(body, `"bound"`) ||
			strings.Contains(body, `"bound_now"`) ||
			strings.Contains(body, "activation_marker") {
			t.Fatalf("persisted lifecycle marker body on %s: %s", rec.Envelope.RelayID, rec.Body)
		}
	}
	for _, field := range reg.Fields {
		lower := strings.ToLower(field.ID)
		if strings.Contains(lower, "activation_marker") || lower == "bound" || lower == "bound_now" {
			t.Fatalf("persisted lifecycle marker registry field %s", field.ID)
		}
	}
}

func waitOutcome(t *testing.T, reply <-chan engine.Outcome) engine.Outcome {
	t.Helper()
	select {
	case out := <-reply:
		return out
	case <-time.After(3 * time.Second):
		t.Fatal("timed out waiting for outcome")
		return engine.Outcome{}
	}
}
