package fsio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/fsio"
)

func TestWriteFileAtomicVisibleAfterCompletion(t *testing.T) {
	root := t.TempDir()
	if err := fsio.WriteFileAtomic(root, "records/one.json", []byte("hello")); err != nil {
		t.Fatalf("WriteFileAtomic: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(root, "records", "one.json"))
	if err != nil {
		t.Fatalf("read committed file: %v", err)
	}
	if string(data) != "hello" {
		t.Fatalf("committed data = %q", data)
	}
	if entries, err := os.ReadDir(filepath.Join(root, "staging")); err != nil {
		t.Fatalf("read staging: %v", err)
	} else if len(entries) != 0 {
		t.Fatalf("staging not empty after completed write: %v", entries)
	}
}

func TestAppendFsync(t *testing.T) {
	path := filepath.Join(t.TempDir(), "journal.jsonl")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		t.Fatalf("open journal: %v", err)
	}
	if err := fsio.AppendFsync(f, []byte("one\n")); err != nil {
		t.Fatalf("AppendFsync one: %v", err)
	}
	if err := fsio.AppendFsync(f, []byte("two\n")); err != nil {
		t.Fatalf("AppendFsync two: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read journal: %v", err)
	}
	if string(data) != "one\ntwo\n" {
		t.Fatalf("journal = %q", data)
	}
}
