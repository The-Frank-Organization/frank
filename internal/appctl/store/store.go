package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

const (
	CurrentSchemaVersion = 1
	StoreFilename        = "app-control.sqlite3"
)

var (
	ErrSchemaTooNew = errors.New("app control store: schema version is newer than this binary")
	ErrIntegrity    = errors.New("app control store: integrity check failed")
)

type Store struct {
	db   *sql.DB
	path string
}

type Tx struct {
	tx *sql.Tx
}

type Snapshot struct {
	tx *sql.Tx
}

func Open(ctx context.Context, runtimeDir string) (*Store, error) {
	if runtimeDir == "" {
		return nil, fmt.Errorf("app control store: empty runtime directory")
	}
	if err := os.MkdirAll(runtimeDir, 0o700); err != nil {
		return nil, fmt.Errorf("app control store: create runtime directory: %w", err)
	}
	if err := os.Chmod(runtimeDir, 0o700); err != nil {
		return nil, fmt.Errorf("app control store: secure runtime directory: %w", err)
	}

	path := filepath.Join(runtimeDir, StoreFilename)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return nil, fmt.Errorf("app control store: create private database: %w", err)
	}
	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("app control store: close database bootstrap handle: %w", err)
	}
	if err := os.Chmod(path, 0o600); err != nil {
		return nil, fmt.Errorf("app control store: secure database: %w", err)
	}

	dsnURL := &url.URL{Scheme: "file", Path: path}
	query := dsnURL.Query()
	query.Set("_pragma", "busy_timeout(5000)")
	query.Set("_txlock", "immediate")
	dsnURL.RawQuery = query.Encode()
	db, err := sql.Open("sqlite", dsnURL.String())
	if err != nil {
		return nil, fmt.Errorf("app control store: open sqlite: %w", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	store := &Store{db: db, path: path}
	if err := store.initialize(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}
	return store, nil
}

func (store *Store) initialize(ctx context.Context) error {
	if err := store.db.PingContext(ctx); err != nil {
		return fmt.Errorf("%w: open database: %v", ErrIntegrity, err)
	}
	var journalMode string
	if err := store.db.QueryRowContext(ctx, "PRAGMA journal_mode=WAL").Scan(&journalMode); err != nil {
		return fmt.Errorf("%w: enable WAL: %v", ErrIntegrity, err)
	}
	if !strings.EqualFold(journalMode, "wal") {
		return fmt.Errorf("app control store: journal mode %q, want WAL", journalMode)
	}
	for _, statement := range []string{
		"PRAGMA synchronous=FULL",
		"PRAGMA foreign_keys=ON",
	} {
		if _, err := store.db.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("app control store: %s: %w", statement, err)
		}
	}
	if got := store.Integrity(ctx); got != "ok" {
		return fmt.Errorf("%w: %s", ErrIntegrity, got)
	}

	version, err := store.schemaVersion(ctx)
	if err != nil {
		return fmt.Errorf("%w: read schema version: %v", ErrIntegrity, err)
	}
	if version > CurrentSchemaVersion {
		return fmt.Errorf("%w: database=%d binary=%d", ErrSchemaTooNew, version, CurrentSchemaVersion)
	}
	if version < CurrentSchemaVersion {
		if err := store.migrate(ctx, version); err != nil {
			return fmt.Errorf("app control store: migrate v%d to v%d: %w", version, CurrentSchemaVersion, err)
		}
	}
	return nil
}

func (store *Store) migrate(ctx context.Context, from int) error {
	if from != 0 {
		return fmt.Errorf("no forward migration from version %d", from)
	}
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, statement := range schemaV1 {
		if _, err := tx.ExecContext(ctx, statement); err != nil {
			return err
		}
	}
	if _, err := tx.ExecContext(ctx, "PRAGMA user_version = 1"); err != nil {
		return err
	}
	return tx.Commit()
}

func (store *Store) Close() error { return store.db.Close() }

func (store *Store) Path() string { return store.path }

func (store *Store) Update(ctx context.Context, apply func(*Tx) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	wrapper := &Tx{tx: tx}
	if err := apply(wrapper); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("app control store: commit: %w", err)
	}
	return nil
}

func (store *Store) Read(ctx context.Context, read func(*Snapshot) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	wrapper := &Snapshot{tx: tx}
	if err := read(wrapper); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, query, args...)
}

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return tx.tx.QueryRowContext(ctx, query, args...)
}

func (snapshot *Snapshot) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return snapshot.tx.QueryContext(ctx, query, args...)
}

func (snapshot *Snapshot) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return snapshot.tx.QueryRowContext(ctx, query, args...)
}

func IsNoRows(err error) bool { return errors.Is(err, sql.ErrNoRows) }

func (store *Store) schemaVersion(ctx context.Context) (int, error) {
	var version int
	err := store.db.QueryRowContext(ctx, "PRAGMA user_version").Scan(&version)
	return version, err
}

func (store *Store) SchemaVersion(ctx context.Context) int {
	version, err := store.schemaVersion(ctx)
	if err != nil {
		return -1
	}
	return version
}

func (store *Store) JournalMode(ctx context.Context) string {
	var mode string
	if err := store.db.QueryRowContext(ctx, "PRAGMA journal_mode").Scan(&mode); err != nil {
		return ""
	}
	return strings.ToLower(mode)
}

func (store *Store) Synchronous(ctx context.Context) int {
	var synchronous int
	if err := store.db.QueryRowContext(ctx, "PRAGMA synchronous").Scan(&synchronous); err != nil {
		return -1
	}
	return synchronous
}

func (store *Store) Integrity(ctx context.Context) string {
	var result string
	if err := store.db.QueryRowContext(ctx, "PRAGMA integrity_check").Scan(&result); err != nil {
		return err.Error()
	}
	return result
}

func (store *Store) Tables(ctx context.Context) []string {
	rows, err := store.db.QueryContext(ctx, `SELECT name FROM sqlite_schema
		WHERE type='table' AND name NOT LIKE 'sqlite_%' ORDER BY name`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var tables []string
	for rows.Next() {
		var table string
		if rows.Scan(&table) == nil {
			tables = append(tables, table)
		}
	}
	return tables
}
