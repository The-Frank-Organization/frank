package fixtures_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

const (
	s5ARegistryCommit = "dd7d0b5"
	// SHA-256 of internal/fieldspec/registry.json from s5-a-registry @ dd7d0b5.
	s5ARegistrySHA256 = "827d24dafd0c1bc47e0968c9596aeae2f1575ad4b6e8c2f46a483b4187f1a9db"
)

func TestS5ConfigChangeOperatorAcceptsLandedRegistryShape(t *testing.T) {
	st, reg := s5ConfigChangeDeps(t)
	body := s5ALandedRegistryBytes(t)
	meta := s5OperatorMeta()
	rec := s5ConfigChangeRelay(t, st.Root, body)

	got, intents := s5SubmitConfigChange(t, st, reg, meta, rec)
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", got.Envelope.DeliveryState, got.Body)
	}
	if got.Envelope.From != "operator" || got.Envelope.Role != "operator" {
		t.Fatalf("operator shape envelope = %+v", got.Envelope)
	}
	if got.Headers["record_kind"] != "config_change" || got.Headers["member"] != "fieldspec" || got.Headers["new_digest"] == "" {
		t.Fatalf("operator shape headers = %+v", got.Headers)
	}
	if got.Body != string(body) {
		t.Fatalf("config_change body differs from s5-a registry bytes")
	}
	if len(intents) == 0 {
		t.Fatalf("accepted config_change returned no intents")
	}
}

func TestS5ConfigChangeMovesOldToNewDigestOverLandedRegistry(t *testing.T) {
	st, reg := s5ConfigChangeDeps(t)
	body := s5ALandedRegistryBytes(t)
	oldPinned := s5LoadPinned(t, st.Root)
	rec := s5ConfigChangeRelay(t, st.Root, body)

	got, intents := s5SubmitConfigChange(t, st, reg, s5OperatorMeta(), rec)
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", got.Envelope.DeliveryState, got.Body)
	}
	if _, err := st.Commit(got, intents); err != nil {
		t.Fatalf("Commit config_change: %v", err)
	}
	newPinned := s5LoadPinned(t, st.Root)
	if oldPinned.Digest == newPinned.Digest {
		t.Fatalf("digest did not move: %s", oldPinned.Digest)
	}
	if newPinned.Digest != got.Headers["new_digest"] {
		t.Fatalf("new digest = %s, want %s", newPinned.Digest, got.Headers["new_digest"])
	}
	if got.Headers["new_digest"] != fixtureDigestWithMember(t, st.Root, "fieldspec", body) {
		t.Fatalf("config_change digest does not match landed registry bytes")
	}
	if materialized := mustReadFile(t, filepath.Join(st.Root, "config", "fieldspec", "registry.json")); !bytes.Equal(materialized, body) {
		t.Fatalf("materialized registry differs from s5-a %s bytes", s5ARegistryCommit)
	}
}

func TestS5ConfigChangeDoesNotReGenesis(t *testing.T) {
	st, reg := s5ConfigChangeDeps(t)
	rec := s5ConfigChangeRelay(t, st.Root, s5ALandedRegistryBytes(t))
	got, intents := s5SubmitConfigChange(t, st, reg, s5OperatorMeta(), rec)
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", got.Envelope.DeliveryState, got.Body)
	}
	if _, err := st.Commit(got, intents); err != nil {
		t.Fatalf("Commit config_change: %v", err)
	}
	if count := s5GenesisRecordCount(t, st); count != 1 {
		t.Fatalf("genesis count = %d, want 1", count)
	}
}

func TestS5ConfigChangePhase0ValidateGenesisWalksAcceptedChain(t *testing.T) {
	st, reg := s5ConfigChangeDeps(t)
	rec := s5ConfigChangeRelay(t, st.Root, s5ALandedRegistryBytes(t))
	got, intents := s5SubmitConfigChange(t, st, reg, s5OperatorMeta(), rec)
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("state = %s, body=%s", got.Envelope.DeliveryState, got.Body)
	}
	if _, err := st.Commit(got, intents); err != nil {
		t.Fatalf("Commit config_change: %v", err)
	}

	reopened, err := store.Open(st.Root)
	if err != nil {
		t.Fatalf("reopen store: %v", err)
	}
	if err := reopened.ValidateGenesis(s5LoadPinned(t, st.Root)); err != nil {
		t.Fatalf("ValidateGenesis after config_change chain: %v", err)
	}
}

func TestS5ConfigChangeStaleFormRerenderShowsRoutingEscalation(t *testing.T) {
	st, oldReg := s5ConfigChangeDeps(t)
	meta := s5OperatorMeta()
	stale := record.Record{Headers: map[string]string{
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "medium",
		"EVIDENCE_TARGET": "E1",
		"SUBJECT":         "stale form",
	}}
	_, staleDigest := oldReg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, "SITREP", "medium", fieldspec.ClosedGrantState)

	change := s5ConfigChangeRelay(t, st.Root, s5ALandedRegistryBytes(t))
	accepted, intents := s5SubmitConfigChange(t, st, oldReg, meta, change)
	if accepted.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("config_change state = %s, body=%s", accepted.Envelope.DeliveryState, accepted.Body)
	}
	if _, err := st.Commit(accepted, intents); err != nil {
		t.Fatalf("Commit config_change: %v", err)
	}

	newReg, err := fieldspec.Load(filepath.Join(st.Root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("Load materialized registry: %v", err)
	}
	handler := engine.SubmitHandler(st, newReg, meta)
	got, _, err := handler(context.Background(), intake.Cmd{
		IntakeID:   "s5-stale-form",
		Seat:       meta.Name,
		Role:       meta.Role,
		IsOperator: meta.IsOperator,
		Payload:    mustJSONBytes(t, fieldspec.SubmitPayload{Record: stale, FormDigest: staleDigest}),
	})
	if err != nil {
		t.Fatalf("Submit stale form: %v", err)
	}
	if got.Envelope.DeliveryState != record.Rejected || !bytes.Contains([]byte(got.Body), []byte("form_digest:re-render")) {
		t.Fatalf("stale state/body = %s/%q, want form_digest:re-render", got.Envelope.DeliveryState, got.Body)
	}

	form, freshDigest := newReg.Render(fieldspec.RenderEnv{}, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, "SITREP", "medium", fieldspec.ClosedGrantState)
	if !form.OptionAllowed("gate_category", "routing_escalation") {
		t.Fatalf("fresh gate_category options omit routing_escalation: %+v", form.Fields["gate_category"])
	}
	fresh := record.Record{Headers: map[string]string{
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "medium",
		"EVIDENCE_TARGET": "E1",
		"SUBJECT":         "fresh form",
		"gate_category":   "routing_escalation",
	}}
	got, _, err = handler(context.Background(), intake.Cmd{
		IntakeID:   "s5-fresh-form",
		Seat:       meta.Name,
		Role:       meta.Role,
		IsOperator: meta.IsOperator,
		Payload:    mustJSONBytes(t, fieldspec.SubmitPayload{Record: fresh, FormDigest: freshDigest}),
	})
	if err != nil {
		t.Fatalf("Submit fresh form: %v", err)
	}
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("fresh state = %s, body=%s", got.Envelope.DeliveryState, got.Body)
	}
}

func s5ConfigChangeDeps(t *testing.T) (*store.Store, *fieldspec.Registry) {
	t.Helper()
	root := t.TempDir()
	enginePath := filepath.Join(t.TempDir(), "engine.json")
	if err := os.WriteFile(enginePath, []byte(`{"gc_enabled":false,"segment_rotate_bytes":4194304}`), 0o644); err != nil {
		t.Fatalf("write engine config: %v", err)
	}
	registryPath := filepath.Join("..", "..", "internal", "fieldspec", "registry.json")
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

func s5ConfigChangeRelay(t *testing.T, root string, body []byte) record.Record {
	t.Helper()
	return record.Record{
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"AUTHORITY":     "report-only",
			"CEREMONY_TIER": "medium",
			"SUBJECT":       "s5-a registry config_change",
			"record_kind":   "config_change",
			"member":        "fieldspec",
			"new_digest":    fixtureDigestWithMember(t, root, "fieldspec", body),
		},
		Body: string(body),
	}
}

func s5SubmitConfigChange(t *testing.T, st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, rec record.Record) (record.Record, []store.Intent) {
	t.Helper()
	handler := engine.SubmitHandler(st, reg, meta)
	got, intents, err := handler(context.Background(), intake.Cmd{
		IntakeID:   "s5-config-change",
		Seat:       meta.Name,
		Role:       meta.Role,
		IsOperator: meta.IsOperator,
		Payload:    mustJSONBytes(t, submitPayloadForRegistry(reg, meta, rec)),
	})
	if err != nil {
		t.Fatalf("SubmitHandler config_change: %v", err)
	}
	return got, intents
}

func s5ALandedRegistryBytes(t *testing.T) []byte {
	t.Helper()
	var mismatches []string
	for _, path := range s5ALandedRegistryPaths() {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		sum := sha256.Sum256(data)
		got := hex.EncodeToString(sum[:])
		if got != s5ARegistrySHA256 {
			mismatches = append(mismatches, path+"="+got)
			continue
		}
		return data
	}
	if len(mismatches) > 0 {
		t.Fatalf("no s5-a registry candidate matched SHA256 %s from %s; mismatches=%v", s5ARegistrySHA256, s5ARegistryCommit, mismatches)
	}
	t.Skipf("s5-a registry @ %s not found; set FRANK_S5_A_REGISTRY", s5ARegistryCommit)
	return nil
}

func TestS5ALandedRegistryBytesSkipsMismatchedCandidateWhenLaterCandidateMatches(t *testing.T) {
	mismatch := filepath.Join(t.TempDir(), "registry.json")
	if err := os.WriteFile(mismatch, []byte(`{"version":"wrong"}`), 0o644); err != nil {
		t.Fatalf("write mismatch registry: %v", err)
	}
	t.Setenv("FRANK_S5_A_REGISTRY", mismatch)

	data := s5ALandedRegistryBytes(t)
	sum := sha256.Sum256(data)
	if got := hex.EncodeToString(sum[:]); got != s5ARegistrySHA256 {
		t.Fatalf("registry SHA256 = %s, want %s", got, s5ARegistrySHA256)
	}
}

func s5ALandedRegistry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	path := filepath.Join(t.TempDir(), "registry.json")
	if err := os.WriteFile(path, s5ALandedRegistryBytes(t), 0o644); err != nil {
		t.Fatalf("write s5-a registry fixture: %v", err)
	}
	reg, err := fieldspec.Load(path)
	if err != nil {
		t.Fatalf("Load s5-a registry fixture: %v", err)
	}
	return reg
}

func s5ALandedRegistryPaths() []string {
	var paths []string
	if path := os.Getenv("FRANK_S5_A_REGISTRY"); path != "" {
		paths = append(paths, path)
	}
	paths = append(paths,
		filepath.Clean(filepath.Join("..", "..", "internal", "fieldspec", "registry.json")),
		filepath.Clean(filepath.Join("..", "..", "..", "s5-a", "internal", "fieldspec", "registry.json")),
	)
	return paths
}

func s5OperatorMeta() seat.SeatMeta {
	return seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
}

func s5LoadPinned(t *testing.T, root string) *config.Pinned {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load pinned config: %v", err)
	}
	return pinned
}

func s5GenesisRecordCount(t *testing.T, st *store.Store) int {
	t.Helper()
	records, err := st.Records()
	if err != nil {
		t.Fatalf("Records: %v", err)
	}
	var count int
	for _, rec := range records {
		if rec.Headers["record_kind"] == "genesis" {
			count++
		}
	}
	return count
}
