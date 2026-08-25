package store

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestOpenGenesisCreatesPrivateDurableSchema(t *testing.T) {
	ctx := context.Background()
	runtimeDir := filepath.Join(t.TempDir(), "runtime")
	store, err := Open(ctx, runtimeDir)
	if err != nil {
		t.Fatalf("Open genesis: %v", err)
	}
	defer store.Close()

	if got := store.SchemaVersion(ctx); got != CurrentSchemaVersion {
		t.Fatalf("schema version = %d, want %d", got, CurrentSchemaVersion)
	}
	if got := store.JournalMode(ctx); got != "wal" {
		t.Fatalf("journal_mode = %q, want wal", got)
	}
	if got := store.Synchronous(ctx); got != 2 {
		t.Fatalf("synchronous = %d, want FULL(2)", got)
	}
	if got := store.Integrity(ctx); got != "ok" {
		t.Fatalf("integrity_check = %q, want ok", got)
	}

	dirInfo, err := os.Stat(runtimeDir)
	if err != nil {
		t.Fatalf("stat runtime dir: %v", err)
	}
	if got := dirInfo.Mode().Perm(); got != 0o700 {
		t.Fatalf("runtime dir mode = %o, want 700", got)
	}
	fileInfo, err := os.Stat(store.Path())
	if err != nil {
		t.Fatalf("stat store file: %v", err)
	}
	if got := fileInfo.Mode().Perm(); got != 0o600 {
		t.Fatalf("store file mode = %o, want 600", got)
	}

	wantTables := []string{
		"broker_control", "broker_events", "cancellations", "content_ready_receipts", "epochs", "leases",
		"pending_app_events", "provider_attempts", "runs", "tool_authorizations", "tool_calls", "turn_disclosures",
		"turns", "wake_schedule", "workers",
	}
	if got := store.Tables(ctx); !slices.Equal(got, wantTables) {
		t.Fatalf("tables = %v, want %v", got, wantTables)
	}
	for _, forbidden := range []string{"epoch_transitions", "crossing_ops", "events"} {
		if slices.Contains(store.Tables(ctx), forbidden) {
			t.Fatalf("deferred/superseded table %q exists", forbidden)
		}
	}
}

func TestOpenRefusesHigherVersionAndCorruption(t *testing.T) {
	ctx := context.Background()
	runtimeDir := filepath.Join(t.TempDir(), "higher")
	store, err := Open(ctx, runtimeDir)
	if err != nil {
		t.Fatalf("Open genesis: %v", err)
	}
	if _, err := store.db.ExecContext(ctx, "PRAGMA user_version = 2"); err != nil {
		t.Fatalf("set higher version: %v", err)
	}
	if err := store.Close(); err != nil {
		t.Fatalf("close higher fixture: %v", err)
	}
	if _, err := Open(ctx, runtimeDir); !errors.Is(err, ErrSchemaTooNew) {
		t.Fatalf("Open higher version error = %v, want ErrSchemaTooNew", err)
	}

	corruptDir := filepath.Join(t.TempDir(), "corrupt")
	if err := os.Mkdir(corruptDir, 0o700); err != nil {
		t.Fatalf("mkdir corrupt fixture: %v", err)
	}
	if err := os.WriteFile(filepath.Join(corruptDir, StoreFilename), []byte("not sqlite"), 0o600); err != nil {
		t.Fatalf("write corrupt fixture: %v", err)
	}
	if _, err := Open(ctx, corruptDir); !errors.Is(err, ErrIntegrity) {
		t.Fatalf("Open corrupt error = %v, want ErrIntegrity", err)
	}
}

func TestOpenMigratesLowerVersionForward(t *testing.T) {
	ctx := context.Background()
	runtimeDir := filepath.Join(t.TempDir(), "lower")
	store, err := Open(ctx, runtimeDir)
	if err != nil {
		t.Fatalf("Open genesis: %v", err)
	}
	if _, err := store.db.ExecContext(ctx, "DROP TABLE turn_disclosures; PRAGMA user_version = 0"); err != nil {
		t.Fatalf("prepare lower fixture: %v", err)
	}
	if err := store.Close(); err != nil {
		t.Fatalf("close lower fixture: %v", err)
	}

	migrated, err := Open(ctx, runtimeDir)
	if err != nil {
		t.Fatalf("Open lower version: %v", err)
	}
	defer migrated.Close()
	if migrated.SchemaVersion(ctx) != CurrentSchemaVersion || !slices.Contains(migrated.Tables(ctx), "turn_disclosures") {
		t.Fatalf("lower schema did not migrate forward: version=%d tables=%v", migrated.SchemaVersion(ctx), migrated.Tables(ctx))
	}
}

func TestUpdateCommitsOrRollsBackAsOneTransaction(t *testing.T) {
	ctx := context.Background()
	store, err := Open(ctx, filepath.Join(t.TempDir(), "tx"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer store.Close()

	insert := func(tx *Tx, runID string) error {
		_, err := tx.ExecContext(ctx, `INSERT INTO runs
			(run_id, manifest_bytes, run_manifest_digest, state, run_phase, consecutive_failures, created_at)
			VALUES (?, '{}', ?, 'ADMITTED', 'created', '00000000000000000000', 1)`, runID, strings.Repeat("0", 64))
		return err
	}
	if err := store.Update(ctx, func(tx *Tx) error { return insert(tx, "committed") }); err != nil {
		t.Fatalf("committing update: %v", err)
	}
	wantRollback := errors.New("rollback")
	if err := store.Update(ctx, func(tx *Tx) error {
		if err := insert(tx, "rolled-back"); err != nil {
			return err
		}
		return wantRollback
	}); !errors.Is(err, wantRollback) {
		t.Fatalf("rollback update error = %v", err)
	}
	if got := store.countRows(ctx, "runs"); got != 1 {
		t.Fatalf("runs after commit+rollback = %d, want 1", got)
	}
}

func (store *Store) countRows(ctx context.Context, table string) int {
	var count int
	if err := store.db.QueryRowContext(ctx, "SELECT count(*) FROM "+table).Scan(&count); err != nil {
		return -1
	}
	return count
}

func TestSchemaRejectsInvalidClosedStatesCountersAndUniqueness(t *testing.T) {
	ctx := context.Background()
	store, err := Open(ctx, filepath.Join(t.TempDir(), "checks"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer store.Close()

	for name, statement := range map[string]string{
		"run state": `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at)
			VALUES('r','{}','0000000000000000000000000000000000000000000000000000000000000000','BOGUS','created','00000000000000000000',1)`,
		"counter": `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES('r','01','00000000000000000000')`,
		"broker singleton": `INSERT INTO broker_control(singleton,control_token,control_generation,minted_at)
			VALUES(2,'token','00000000000000000001',1)`,
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := store.db.ExecContext(ctx, statement); err == nil {
				t.Fatal("invalid row unexpectedly committed")
			}
		})
	}
}

func TestSchemaPinsLifecyclePhasesAndCreateAuthorization(t *testing.T) {
	ctx := context.Background()
	db, err := Open(ctx, filepath.Join(t.TempDir(), "lifecycle-checks"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer db.Close()

	digest := strings.Repeat("0", 64)
	for _, phase := range []string{"GENESIS", "RUNNING", "TERMINAL", "Created"} {
		if _, err := db.db.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, "bad-"+phase, []byte("{}"), digest, "ACTIVE", phase, "00000000000000000000", 1); err == nil {
			t.Fatalf("invalid lifecycle phase %q committed", phase)
		}
	}
	if _, err := db.db.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, "run", []byte("{}"), digest, "ACTIVE", "established", "00000000000000000000", 1); err != nil {
		t.Fatal(err)
	}
	for name, authID := range map[string]any{
		"missing":   nil,
		"short":     strings.Repeat("a", 31),
		"uppercase": strings.Repeat("A", 32),
	} {
		t.Run(name, func(t *testing.T) {
			_, err := db.db.ExecContext(ctx, `INSERT INTO turns(turn_id,run_id,turn_epoch,state,admission_ref,run_disposition,create_auth_id,resume_disposition,created_at) VALUES(?,?,?,?,?,?,?,?,?)`, "turn-"+name, "run", "00000000000000000001", "ACTIVE", []byte("{}"), "fresh", authID, "PENDING", 1)
			if err == nil {
				t.Fatal("invalid create authorization committed")
			}
		})
	}
}
