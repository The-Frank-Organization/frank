package fixtures_test

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/record"
	frankrecover "github.com/The-Frank-Organization/frank/internal/recover"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS10CrashSafeParkAndWakeConvergeExactlyOnce(t *testing.T) {
	if os.Getenv("FRANK_S10_T7_CHILD") == "1" {
		runS10T7Child(t)
		return
	}
	for _, mutation := range []string{"park", "wake"} {
		t.Run(mutation, func(t *testing.T) {
			root := t.TempDir()
			initFixtureStore(t, root)
			st, err := store.Open(root)
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			commitS10ParkRecord(t, st, record.Record{
				Envelope: record.Envelope{
					RelayID:       "t7-gate",
					From:          "seat-a",
					Role:          "implementer",
					DeliveryState: record.Accepted,
					SchemaVersion: 1,
				},
				Headers: map[string]string{
					"PHASE":               "SITREP",
					"AUTHORITY":           "report-only",
					"SUBJECT":             "crash-safe gate",
					"HUMAN_GATE_REQUIRED": "yes",
				},
			})
			if mutation == "wake" {
				if err := obligation.CompleteAuto(st); err != nil {
					t.Fatalf("prepare park: %v", err)
				}
			}

			cmd := exec.Command(os.Args[0], "-test.run", "^TestS10CrashSafeParkAndWakeConvergeExactlyOnce$")
			cmd.Env = append(os.Environ(),
				"FRANK_S10_T7_CHILD=1",
				"FRANK_S10_T7_ROOT="+root,
				"FRANK_S10_T7_MUTATION="+mutation,
				"FRANK_TEST_CRASHPOINT=pre_delivery_write",
			)
			err = cmd.Run()
			if err == nil {
				t.Fatalf("%s child completed without crashing during durable delivery", mutation)
			}
			assertSIGKILL(t, err)

			result, err := frankrecover.Run(root, loadFixturePinned(t, root))
			if err != nil {
				t.Fatalf("recover %s: %v", mutation, err)
			}
			if result.Ready == nil || result.Diag != nil {
				t.Fatalf("recover %s result = %+v, want ready", mutation, result)
			}
			st, err = store.Open(root)
			if err != nil {
				t.Fatalf("reopen: %v", err)
			}
			if err := obligation.CompleteAuto(st); err != nil {
				t.Fatalf("complete after recovery: %v", err)
			}
			assertS10RelayCount(t, st, "odb-t7-gate", 1)
			assertS10RelayCount(t, st, "park-t7-gate", 1)
			assertS10MailboxCount(t, root, "operator", "odb-t7-gate", 1)
			tab, err := obligation.BuildTables(st)
			if err != nil {
				t.Fatalf("BuildTables: %v", err)
			}

			if mutation == "wake" {
				wake := s10RecordByIntake(t, st, "t7-wake-intake")
				if wake.Envelope.DeliveryState != record.Accepted || wake.Envelope.To != "seat-a" {
					t.Fatalf("recovered wake = %+v", wake)
				}
				assertS10MailboxCount(t, root, "seat-a", wake.Envelope.RelayID, 1)
				if state := engine.GateState(tab, "t7-gate"); state != "resumed" {
					t.Fatalf("recovered wake state = %q, want resumed", state)
				}
			} else if state := engine.GateState(tab, "t7-gate"); state != "parked_waiting_human" {
				t.Fatalf("recovered park state = %q, want parked_waiting_human", state)
			}
		})
	}
}

func runS10T7Child(t *testing.T) {
	root := os.Getenv("FRANK_S10_T7_ROOT")
	mutation := os.Getenv("FRANK_S10_T7_MUTATION")
	if root == "" || mutation == "" {
		t.Fatalf("missing T7 child environment")
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	switch mutation {
	case "park":
		if err := obligation.CompleteAuto(st); err != nil {
			t.Fatalf("CompleteAuto: %v", err)
		}
	case "wake":
		reg, err := fieldspec.Load(filepath.Join(root, "config", "fieldspec", "registry.json"))
		if err != nil {
			t.Fatalf("Load registry: %v", err)
		}
		meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
		candidate := record.Record{
			Headers: map[string]string{
				"PHASE":           "SITREP",
				"AUTHORITY":       "report-only",
				"CEREMONY_TIER":   "medium",
				"EVIDENCE_TARGET": "E1",
				"SUBJECT":         "crash-safe operator reply",
				"record_kind":     "gate_resolution",
				"parent_hint":     "t7-gate",
				"resolves_gate":   "t7-gate",
			},
			Body: `{"choice":"approve"}`,
		}
		_, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: true}, "SITREP", "medium", fieldspec.ClosedGrantState)
		payload, err := json.Marshal(fieldspec.SubmitPayload{Record: candidate, FormDigest: digest})
		if err != nil {
			t.Fatalf("Marshal payload: %v", err)
		}
		rec, intents, err := engine.SubmitHandler(st, reg, meta)(context.Background(), intake.Cmd{
			IntakeID: "t7-wake-intake", Seat: "operator", Role: "operator", IsOperator: true, Payload: payload,
		})
		if err != nil {
			t.Fatalf("SubmitHandler: %v", err)
		}
		if rec.Envelope.DeliveryState != record.Accepted {
			t.Fatalf("reply rejected: %s", rec.Body)
		}
		if _, err := st.Commit(rec, intents); err != nil {
			t.Fatalf("Commit wake: %v", err)
		}
	default:
		t.Fatalf("unknown T7 mutation %q", mutation)
	}
}

func assertS10MailboxCount(t *testing.T, root, seatName, relayID string, want int) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, "mailboxes", seatName+".jsonl"))
	if err != nil {
		t.Fatalf("read %s mailbox: %v", seatName, err)
	}
	got := 0
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if strings.TrimSpace(line) == relayID {
			got++
		}
	}
	if got != want {
		t.Fatalf("%s mailbox %s count = %d, want %d; mailbox=%q", seatName, relayID, got, want, data)
	}
}

func s10RecordByIntake(t *testing.T, st *store.Store, intakeID string) record.Record {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	for _, rec := range records {
		if rec.Envelope.IntakeID == intakeID {
			return rec
		}
	}
	t.Fatalf("record for intake %s missing", intakeID)
	return record.Record{}
}
