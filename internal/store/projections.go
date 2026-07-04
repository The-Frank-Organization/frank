package store

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fsio"
)

func (s *Store) AppendIndex(row string) error {
	return appendUnique(filepath.Join(s.Root, "projections", "INDEX.md"), []byte(row))
}

func (s *Store) RebuildProjections() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries, err := readRedo(s.Root)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		recordPath := filepath.Join(s.Root, "records", entry.RelayID+".json")
		if _, err := os.Stat(recordPath); err != nil {
			continue
		}
		if err := s.applyIntents(entry.Intents); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) applyIntents(intents []Intent) error {
	for _, intent := range intents {
		switch intent.Kind {
		case IntentIndex:
			crashpoint.Hit("pre_projection_write")
			if err := appendUnique(filepath.Join(s.Root, "projections", intent.Path), intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		case IntentRender:
			crashpoint.Hit("pre_projection_write")
			if err := fsio.WriteFileAtomic(filepath.Join(s.Root, "projections"), intent.Path, intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		case IntentMailbox:
			crashpoint.Hit("pre_projection_write")
			if err := appendUnique(filepath.Join(s.Root, "mailboxes", intent.Path), intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		case IntentOutbox:
			crashpoint.Hit("pre_projection_write")
			if err := fsio.WriteFileAtomic(filepath.Join(s.Root, "outbox"), intent.Path, intent.Payload); err != nil {
				return err
			}
			crashpoint.Hit("post_projection_write")
		}
	}
	return nil
}

func appendUnique(path string, payload []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	current, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if strings.Contains(string(current), string(payload)) {
		return nil
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return fsio.AppendFsync(f, payload)
}
