package fsio

import (
	"os"
	"path/filepath"

	"github.com/jackli/frank/internal/crashpoint"
)

func WriteFileAtomic(root, rel string, data []byte) error {
	stage := filepath.Join(root, "staging", rel+".tmp")
	final := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(stage), 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(final), 0o755); err != nil {
		return err
	}
	if err := os.Remove(stage); err != nil && !os.IsNotExist(err) {
		return err
	}
	f, err := os.OpenFile(stage, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		_ = f.Close()
		return err
	}
	crashpoint.Hit("pre_record_fsync")
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return err
	}
	crashpoint.Hit("post_record_fsync")
	if err := f.Close(); err != nil {
		return err
	}
	crashpoint.Hit("pre_rename")
	if err := os.Rename(stage, final); err != nil {
		return err
	}
	crashpoint.Hit("post_rename")
	crashpoint.Hit("pre_dir_fsync")
	if err := syncDir(filepath.Dir(final)); err != nil {
		return err
	}
	crashpoint.Hit("post_dir_fsync")
	_ = removeEmptyParents(filepath.Dir(stage), filepath.Join(root, "staging"))
	return nil
}

func AppendFsync(f *os.File, line []byte) error {
	if _, err := f.Write(line); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return syncDir(filepath.Dir(f.Name()))
}

func syncDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	if err := d.Sync(); err != nil {
		_ = d.Close()
		return err
	}
	return d.Close()
}

func removeEmptyParents(dir, stop string) error {
	for {
		if dir == stop || dir == "." || dir == string(filepath.Separator) {
			return nil
		}
		if err := os.Remove(dir); err != nil {
			return err
		}
		dir = filepath.Dir(dir)
	}
}
