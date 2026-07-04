package recover

import (
	"os"
	"path/filepath"

	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func Run(root string) error {
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
	if err := gate.Complete(st); err != nil {
		return err
	}
	return nil
}
