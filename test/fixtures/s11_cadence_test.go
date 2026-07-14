package fixtures_test

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

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestS11G4CadenceComesFromOperatorConfigAndNeverApproves(t *testing.T) {
	for _, tc := range []struct {
		name                   string
		noResponseSeconds      int64
		answeredStalledSeconds int64
		answered               bool
		wantReason             string
	}{
		{name: "no response", noResponseSeconds: 0, answeredStalledSeconds: 3600, wantReason: "no_response"},
		{name: "answered but stalled", noResponseSeconds: 3600, answeredStalledSeconds: 0, answered: true, wantReason: "answered_but_stalled"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			sources := writeFixtureConfigSources(t)
			writeS11CadenceEngine(t, sources["engine"], tc.noResponseSeconds, tc.answeredStalledSeconds, false)
			if err := store.Init(root, sources); err != nil {
				t.Fatalf("Init with G4 cadence: %v", err)
			}
			st, err := store.Open(root)
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			gateID := "s11-g4-" + strings.ReplaceAll(tc.name, " ", "-")
			for _, rec := range []record.Record{
				{
					Envelope: record.Envelope{RelayID: gateID, From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
					Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "G4 configured cadence", "HUMAN_GATE_REQUIRED": "yes"},
				},
				{
					Envelope: record.Envelope{RelayID: "park-" + gateID, From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
					Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "parked_waiting_human", "parks_gate": gateID},
				},
			} {
				if _, err := st.Commit(rec, nil); err != nil {
					t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
				}
			}
			if tc.answered {
				if _, err := st.Commit(record.Record{
					Envelope: record.Envelope{RelayID: "rejected-answer-" + gateID, From: "operator", Role: "operator", DeliveryState: record.Rejected, SchemaVersion: 1},
					Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "answer awaiting repair", "resolves_gate": gateID},
				}, nil); err != nil {
					t.Fatalf("Commit rejected answer: %v", err)
				}
			}

			ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
			defer cancel()
			sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s11-g4-%d.sock", time.Now().UnixNano()))
			t.Cleanup(func() { _ = os.Remove(sock) })
			cmd, stderr := startFrank(t, ctx, buildFrank(t, ctx), root, sock)
			t.Cleanup(func() {
				cancel()
				if cmd.Process != nil {
					_ = cmd.Process.Kill()
				}
				_ = cmd.Wait()
			})
			socketDeadline := time.Now().Add(3 * time.Second)
			for time.Now().Before(socketDeadline) {
				if _, err := os.Stat(sock); err == nil {
					break
				}
				time.Sleep(10 * time.Millisecond)
			}
			if _, err := os.Stat(sock); err != nil {
				t.Fatalf("configured-cadence frank did not start: %v; stderr=%s", err, stderr.String())
			}

			deadline := time.Now().Add(2 * time.Second)
			for time.Now().Before(deadline) {
				records, err := st.Records()
				if err != nil {
					t.Fatalf("Records: %v", err)
				}
				for _, rec := range records {
					if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["resolves_gate"] == gateID {
						t.Fatalf("G4 cadence auto-approved gate: %#v", rec)
					}
					if rec.Headers["record_kind"] == "resummon_command" && rec.Headers["subject_ref"] == gateID {
						if !strings.Contains(rec.Body, `"reason":"`+tc.wantReason+`"`) {
							t.Fatalf("resummon body = %q, want reason %q", rec.Body, tc.wantReason)
						}
						return
					}
				}
				time.Sleep(10 * time.Millisecond)
			}
			t.Fatalf("configured %s cadence did not fire; stderr=%s", tc.name, stderr.String())
		})
	}
}

func TestS11G4CadenceConfigCannotEncodeAutoApprove(t *testing.T) {
	sources := writeFixtureConfigSources(t)
	writeS11CadenceEngine(t, sources["engine"], 1, 2, true)
	if _, err := config.Load(sources); !errors.Is(err, config.ErrConfigLoad) {
		t.Fatalf("config with auto_approve = %v, want ErrConfigLoad", err)
	}
}

func TestS11G4CadenceOmissionKeepsShippedOneHourDefaults(t *testing.T) {
	noResponse, answeredButStalled := (config.EngineConfig{}).ResummonCadenceDelays()
	if noResponse != time.Hour || answeredButStalled != time.Hour {
		t.Fatalf("omitted cadence = (%s, %s), want one hour per class", noResponse, answeredButStalled)
	}
}

func writeS11CadenceEngine(t *testing.T, path string, noResponseSeconds, answeredStalledSeconds int64, autoApprove bool) {
	t.Helper()
	var engine map[string]any
	if err := json.Unmarshal(fixtureEngineConfig(t, false), &engine); err != nil {
		t.Fatalf("decode engine fixture: %v", err)
	}
	cadence := map[string]any{
		"no_response_seconds":          noResponseSeconds,
		"answered_but_stalled_seconds": answeredStalledSeconds,
	}
	if autoApprove {
		cadence["auto_approve"] = true
	}
	engine["resummon_cadence"] = cadence
	data, err := json.Marshal(engine)
	if err != nil {
		t.Fatalf("marshal cadence engine: %v", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("write cadence engine: %v", err)
	}
}
