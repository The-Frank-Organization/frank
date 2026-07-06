package engine_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestConfigChangeNonOperatorRejected(t *testing.T) {
	st, reg := configChangeDeps(t)
	meta := seat.SeatMeta{Name: "seat-a", Role: "implementer"}
	rec := configChangeRecord(t, st.Root, "fieldspec", []byte(`{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`))
	handler := engine.SubmitHandler(st, reg, meta)

	got, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "cfg-nonop", Seat: meta.Name, Role: meta.Role, Payload: submitPayload(t, reg, meta, rec)})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if got.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", got.Envelope.DeliveryState)
	}
	if len(intents) != 0 {
		t.Fatalf("non-operator produced intents: %#v", intents)
	}
	if !strings.Contains(got.Body, "record_kind") || strings.Contains(got.Body, st.Root) {
		t.Fatalf("non-operator body = %q, want typed path-free record_kind rejection", got.Body)
	}
}

func TestConfigChangeDigestMismatchRejected(t *testing.T) {
	st, reg := configChangeDeps(t)
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	rec := configChangeRecord(t, st.Root, "fieldspec", []byte(`{"phase":["SITREP"],"authority":[],"ceremony_tier":[],"evidence_target":[],"gate_category":{},"grant":[]}`))
	rec.Headers["new_digest"] = "wrong-digest"
	handler := engine.SubmitHandler(st, reg, meta)

	got, intents, err := handler(context.Background(), intake.Cmd{IntakeID: "cfg-digest", Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: submitPayload(t, reg, meta, rec)})
	if err != nil {
		t.Fatalf("handler: %v", err)
	}
	if got.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("state = %s, want rejected", got.Envelope.DeliveryState)
	}
	if len(intents) != 0 {
		t.Fatalf("digest mismatch produced intents: %#v", intents)
	}
	if !strings.Contains(got.Body, "new_digest") || strings.Contains(got.Body, st.Root) {
		t.Fatalf("digest mismatch body = %q, want typed path-free new_digest rejection", got.Body)
	}
}

func configChangeDeps(t *testing.T) (*store.Store, *fieldspec.Registry) {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(t.TempDir(), "engine.json")
	if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	registryPath := filepath.Join("..", "fieldspec", "registry.json")
	if err := store.Init(root, map[string]string{"engine": enginePath, "fieldspec": registryPath}); err != nil {
		t.Fatalf("Init: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	reg, err := fieldspec.Load(registryPath)
	if err != nil {
		t.Fatalf("Load registry: %v", err)
	}
	return st, reg
}

func configChangeRecord(t *testing.T, root, member string, body []byte) record.Record {
	t.Helper()
	return record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "config change",
			"record_kind":     "config_change",
			"member":          member,
			"new_digest":      engineDigestWithMember(t, root, member, body),
		},
		Body: string(body),
	}
}

func engineDigestWithMember(t *testing.T, root, member string, body []byte) string {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store config: %v", err)
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return config.Digest(members)
}
