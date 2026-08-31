package manifest

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

type FreezeEvent struct {
	Frozen         Frozen
	Gate           Gate
	SessionLogPath string
	CreatedAt      int64
}

func (event FreezeEvent) RunID() string { return event.Frozen.Manifest.RunID }

func (event FreezeEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	if err := event.Gate.Validate(event.Frozen); err != nil {
		return applier.Result{}, err
	}
	manifest := event.Frozen.Manifest
	if event.SessionLogPath == "" {
		return applier.Result{}, errors.New("freeze run manifest: session log path is required")
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO runs
			(run_id,manifest_bytes,run_manifest_digest,state,run_phase,session_log_path,consecutive_failures,created_at)
			VALUES(?,?,?,?,?,?,?,?)`, manifest.RunID, event.Frozen.Bytes, event.Frozen.Digest,
		"ADMITTED", "created", event.SessionLogPath, "00000000000000000000", event.CreatedAt); err != nil {
		return applier.Result{}, fmt.Errorf("freeze run manifest: %w", err)
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`,
		manifest.RunID, "00000000000000000001", "00000000000000000000"); err != nil {
		return applier.Result{}, fmt.Errorf("freeze run epoch: %w", err)
	}
	return applier.Result{Value: event.Frozen.Digest}, nil
}

func Load(ctx context.Context, host *applier.Host, runID string, gate Gate) (Frozen, error) {
	value, err := host.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var encoded []byte
		var digest string
		if err := snapshot.QueryRowContext(ctx, `SELECT manifest_bytes,run_manifest_digest FROM runs WHERE run_id=?`, runID).Scan(&encoded, &digest); err != nil {
			return nil, err
		}
		return Frozen{Bytes: append([]byte(nil), encoded...), Digest: digest}, nil
	}))
	if err != nil {
		return Frozen{}, err
	}
	frozen, ok := value.(Frozen)
	if !ok {
		return Frozen{}, errors.New("app control manifest: invalid snapshot type")
	}
	return DecodeFrozen(frozen.Bytes, frozen.Digest, gate)
}

// DecodeFrozen reconstructs a manifest exclusively from its immutable stored
// bytes and reruns the complete serve gate before returning it.
func DecodeFrozen(encoded []byte, digest string, gate Gate) (Frozen, error) {
	frozen := Frozen{Bytes: append([]byte(nil), encoded...), Digest: digest}
	if digestBytes(frozen.Bytes) != frozen.Digest {
		return Frozen{}, ErrManifestIntegrity
	}
	decoded, err := parse(frozen.Bytes)
	if err != nil {
		return Frozen{}, err
	}
	frozen.Manifest = decoded
	canonical, err := appCanonical(decoded)
	if err != nil || !bytes.Equal(canonical, frozen.Bytes) {
		return Frozen{}, ErrManifestIntegrity
	}
	if err := gate.Validate(frozen); err != nil {
		return Frozen{}, err
	}
	return frozen, nil
}
