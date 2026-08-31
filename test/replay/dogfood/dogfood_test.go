package dogfood_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

// The recorded dogfood pattern completes without livelock; the race class is proven by TestConcurrentAcceptNoParentClassBounce (G-3).
func TestS6DogfoodArchiveTrafficReDrive(t *testing.T) {
	archiveRoot := os.Getenv("FRANK_S6_DOGFOOD_STORE")
	if archiveRoot == "" {
		t.Skip("FRANK_S6_DOGFOOD_STORE unset")
	}
	archive, err := store.Open(archiveRoot)
	if err != nil {
		t.Fatalf("Open archive: %v", err)
	}
	records, err := archiveRecordsInCommitOrder(archive)
	if err != nil {
		t.Fatalf("archive records: %v", err)
	}
	freshRoot := t.TempDir()
	if err := store.Init(freshRoot, dogfoodConfigSources(t)); err != nil {
		t.Fatalf("Init fresh store: %v", err)
	}
	fresh, err := store.Open(freshRoot)
	if err != nil {
		t.Fatalf("Open fresh: %v", err)
	}
	reg, err := fieldspec.Load(filepath.Join("..", "..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	var accepted int
	for _, archived := range records {
		if !replayableDogfoodRecord(archived) {
			continue
		}
		meta := seat.SeatMeta{
			Name:       archived.Envelope.From,
			Role:       archived.Envelope.Role,
			IsOperator: archived.Envelope.From == "operator" || archived.Envelope.Role == "operator",
		}
		candidate := dogfoodCandidate(archived)
		handler := engine.SubmitHandler(fresh, reg, meta)
		payload := dogfoodPayload(reg, meta, candidate)
		rec, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "dogfood-" + archived.Envelope.RelayID, Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, Payload: payload})
		if err != nil {
			t.Fatalf("%s handler: %v", archived.Envelope.RelayID, err)
		}
		if rec.Envelope.DeliveryState != record.Accepted {
			if strings.Contains(rec.Body, "PARENT_DISPATCH_ID") || strings.Contains(rec.Body, "parent") || strings.Contains(rec.Body, "form_digest:re-render") {
				t.Fatalf("%s replay bounced with forbidden class: %s", archived.Envelope.RelayID, rec.Body)
			}
			t.Fatalf("%s replay rejected: %s", archived.Envelope.RelayID, rec.Body)
		}
		if _, err := fresh.Commit(rec, intents); err != nil {
			t.Fatalf("%s commit: %v", archived.Envelope.RelayID, err)
		}
		accepted++
	}
	if accepted == 0 {
		t.Fatalf("archive contained no replayable accepted traffic")
	}
}

func archiveRecordsInCommitOrder(st *store.Store) ([]record.Record, error) {
	records, err := st.Records()
	if err != nil {
		return nil, err
	}
	order, err := st.CommitOrder()
	if err != nil {
		return nil, err
	}
	byRelay := map[string]record.Record{}
	for _, rec := range records {
		byRelay[rec.Envelope.RelayID] = rec
	}
	var out []record.Record
	seen := map[string]bool{}
	for _, relayID := range order {
		rec, ok := byRelay[relayID]
		if !ok {
			continue
		}
		out = append(out, rec)
		seen[relayID] = true
	}
	for _, rec := range records {
		if !seen[rec.Envelope.RelayID] {
			out = append(out, rec)
		}
	}
	return out, nil
}

func replayableDogfoodRecord(rec record.Record) bool {
	if rec.Envelope.DeliveryState != record.Accepted {
		return false
	}
	if rec.Envelope.From == "" || rec.Envelope.From == "system" || rec.Headers["record_kind"] == "genesis" {
		return false
	}
	if rec.Headers["record_kind"] == "config_change" {
		return false
	}
	return true
}

func dogfoodCandidate(rec record.Record) record.Record {
	rec.Envelope.RelayID = ""
	rec.Envelope.From = ""
	rec.Envelope.Role = ""
	rec.Envelope.DeliveryState = ""
	rec.Envelope.IntakeID = ""
	rec.Envelope.SchemaVersion = 0
	delete(rec.Headers, "PARENT_DISPATCH_ID")
	delete(rec.Headers, "parent_hint_honored")
	delete(rec.Headers, "parent_provenance")
	delete(rec.Headers, "routing_ref_honored")
	delete(rec.Headers, "gate_category_raised")
	delete(rec.Headers, "gate_category_pick")
	return rec
}

func dogfoodPayload(reg *fieldspec.Registry, meta seat.SeatMeta, rec record.Record) []byte {
	_, digest := reg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	data, _ := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	return data
}

func dogfoodConfigSources(t *testing.T) map[string]string {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(root, "engine.json")
	registryPath := filepath.Join(root, "registry.json")
	if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	registryBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	if err := os.WriteFile(registryPath, registryBytes, 0o644); err != nil {
		t.Fatalf("write registry config: %v", err)
	}
	return map[string]string{"engine": enginePath, "fieldspec": registryPath}
}
