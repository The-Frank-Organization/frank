package recover

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/crashpoint"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/gate"
	frankgc "github.com/The-Frank-Organization/frank/internal/gc"
	"github.com/The-Frank-Organization/frank/internal/intake"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
)

type Result struct {
	Ready *engine.Ready
	Diag  *engine.Diagnostics
}

func Run(root string, pinned *config.Pinned) (*Result, error) {
	return RunWithProcessor(root, pinned, nil)
}

func RunWithProcessor(root string, pinned *config.Pinned, process func(intake.Cmd) error) (*Result, error) {
	probe := &store.Store{Root: root}
	if err := probe.ValidateGenesis(pinned); err != nil {
		if errors.Is(err, store.ErrGenesisMissing) {
			return &Result{Diag: engine.NewDiagnostics("genesis-missing", "", "")}, nil
		}
		var mismatch store.ErrDigestMismatch
		if errors.As(err, &mismatch) {
			return &Result{Diag: engine.NewDiagnostics("digest-mismatch", mismatch.Want, mismatch.Got)}, nil
		}
		return nil, err
	}
	crashpoint.Hit("recovery_post_phase0")

	if err := os.RemoveAll(filepath.Join(root, "staging")); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Join(root, "staging"), 0o755); err != nil {
		return nil, err
	}
	st, err := store.Open(root)
	if err != nil {
		return nil, err
	}
	if _, err := st.ScanQuarantine(); err != nil {
		return nil, err
	}
	if err := st.CompleteIncidents(); err != nil {
		return nil, err
	}
	crashpoint.Hit("recovery_post_phase1")

	if err := st.RebuildProjections(); err != nil {
		return nil, err
	}
	crashpoint.Hit("recovery_post_phase2")

	if _, err := seat.Open(root); err != nil {
		return nil, err
	}
	crashpoint.Hit("recovery_post_phase3")

	if process != nil {
		journal, err := intake.Open(root)
		if err != nil {
			return nil, err
		}
		unconsumed, err := intake.Unconsumed(context.Background(), journal, st)
		if err != nil {
			return nil, err
		}
		for _, cmd := range unconsumed {
			if err := process(cmd); err != nil {
				return nil, err
			}
		}
	}
	crashpoint.Hit("recovery_post_phase3_5")

	if err := gate.Complete(st); err != nil {
		return nil, err
	}
	tables, err := obligation.BuildTables(st)
	if err != nil {
		return nil, err
	}
	if err := frankgc.Pass(st, tables, pinned.Engine); err != nil {
		return nil, err
	}
	crashpoint.Hit("recovery_post_phase3_6")

	ready := engine.NewReady()
	crashpoint.Hit("recovery_post_phase4")
	return &Result{Ready: ready}, nil
}
