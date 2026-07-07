package store_test

import (
	"errors"
	"testing"

	"github.com/jackli/frank/internal/store"
)

func TestAcquireRootExcludesSecondHolderAndReleases(t *testing.T) {
	root := t.TempDir()
	first, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot first: %v", err)
	}
	defer first.Release()

	_, err = store.AcquireRoot(root)
	var locked store.ErrRootLocked
	if !errors.As(err, &locked) {
		t.Fatalf("second AcquireRoot err = %v, want ErrRootLocked", err)
	}
	if locked.HolderPID == 0 {
		t.Fatalf("locked error missing holder pid: %+v", locked)
	}

	if err := first.Release(); err != nil {
		t.Fatalf("Release: %v", err)
	}
	second, err := store.AcquireRoot(root)
	if err != nil {
		t.Fatalf("AcquireRoot after release: %v", err)
	}
	defer second.Release()
}
