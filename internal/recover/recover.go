package recover

import (
	"context"
	"os"
	"path/filepath"

	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func Run(root string) error {
	return RunWithProcessor(root, nil)
}

func RunWithProcessor(root string, process func(intake.Cmd) error) error {
	if err := os.RemoveAll(filepath.Join(root, "staging")); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(root, "staging"), 0o755); err != nil {
		return err
	}
	st, err := store.Open(root)
	if err != nil {
		return err
	}
	if err := st.RebuildProjections(); err != nil {
		return err
	}
	if _, err := seat.Open(root); err != nil {
		return err
	}
	if process != nil {
		journal, err := intake.Open(root)
		if err != nil {
			return err
		}
		unconsumed, err := intake.Unconsumed(context.Background(), journal, st)
		if err != nil {
			return err
		}
		for _, cmd := range unconsumed {
			if err := process(cmd); err != nil {
				return err
			}
		}
	}
	if err := gate.Complete(st); err != nil {
		return err
	}
	return nil
}
