package store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fsio"
	"github.com/jackli/frank/internal/record"
)

const (
	IntentIndex   = "index"
	IntentRender  = "render"
	IntentMailbox = "mailbox"
	IntentOutbox  = "outbox"
)

type Intent struct {
	Kind    string `json:"kind"`
	Path    string `json:"path"`
	Payload []byte `json:"payload"`
}

type Store struct {
	Root string
	mu   sync.Mutex
}

func Open(root string) (*Store, error) {
	for _, dir := range []string{
		"records",
		"staging",
		filepath.Join("journal"),
		filepath.Join("projections", "relays"),
		"mailboxes",
		"outbox",
		filepath.Join("binding"),
	} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			return nil, err
		}
	}
	return &Store{Root: root}, nil
}

func (s *Store) Commit(rec record.Record, intents []Intent) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	relayID := rec.Envelope.RelayID
	if relayID == "" {
		relayID = nextRelayID()
		rec.Envelope.RelayID = relayID
	}
	if rec.Envelope.SchemaVersion == 0 {
		rec.Envelope.SchemaVersion = 1
	}
	data, err := record.Seal(rec)
	if err != nil {
		return "", err
	}
	if err := s.appendRedo(relayID, intents); err != nil {
		return "", err
	}
	if err := fsio.WriteFileAtomic(s.Root, filepath.Join("records", relayID+".json"), data); err != nil {
		return "", err
	}
	if err := s.applyIntents(intents); err != nil {
		return "", err
	}
	return relayID, nil
}

func (s *Store) Records() ([]record.Record, error) {
	entries, err := os.ReadDir(filepath.Join(s.Root, "records"))
	if err != nil {
		return nil, err
	}
	records := make([]record.Record, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.Root, "records", entry.Name()))
		if err != nil {
			return nil, err
		}
		rec, err := record.Verify(data)
		if err != nil {
			return nil, err
		}
		records = append(records, rec)
	}
	return records, nil
}

func (s *Store) Read(relayID string) (record.Record, error) {
	data, err := os.ReadFile(filepath.Join(s.Root, "records", relayID+".json"))
	if err != nil {
		return record.Record{}, err
	}
	return record.Verify(data)
}

func (s *Store) Project(seat string) ([]string, error) {
	path := filepath.Join(s.Root, "mailboxes", seat+".jsonl")
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var relayIDs []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			relayIDs = append(relayIDs, line)
		}
	}
	return relayIDs, scanner.Err()
}

func (s *Store) appendRedo(relayID string, intents []Intent) error {
	path := filepath.Join(s.Root, "journal", "redo.jsonl")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	crashpoint.Hit("pre_redo_fsync")
	data, err := json.Marshal(redoEntry{RelayID: relayID, Intents: intents})
	if err != nil {
		return err
	}
	data = append(data, '\n')
	if err := fsio.AppendFsync(f, data); err != nil {
		return err
	}
	crashpoint.Hit("post_redo_fsync")
	return nil
}

func nextRelayID() string {
	return fmt.Sprintf("relay-%d", os.Getpid())
}

type redoEntry struct {
	RelayID string   `json:"relay_id"`
	Intents []Intent `json:"intents"`
}

func readRedo(root string) ([]redoEntry, error) {
	path := filepath.Join(root, "journal", "redo.jsonl")
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var entries []redoEntry
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var entry redoEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, scanner.Err()
}
