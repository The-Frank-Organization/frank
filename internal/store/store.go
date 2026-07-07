package store

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fsio"
	"github.com/jackli/frank/internal/record"
)

const defaultSegmentRotateBytes int64 = 4 * 1024 * 1024

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

const (
	ProjectDefault = "default"
	ProjectAudit   = "audit"
)

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
	var err error
	if relayID == "" {
		relayID, err = s.nextRelayIDLocked()
		if err != nil {
			return "", err
		}
		rec.Envelope.RelayID = relayID
	} else if _, err := os.Stat(filepath.Join(s.Root, "records", relayID+".json")); err == nil {
		return "", fmt.Errorf("record already exists")
	} else if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	if rec.Envelope.SchemaVersion == 0 {
		rec.Envelope.SchemaVersion = 1
	}
	if intents == nil && rec.Envelope.DeliveryState == record.Accepted && rec.Headers["PHASE"] != "" {
		intents, err = DefaultProjectionIntentsStrict(rec)
		if err != nil {
			return "", err
		}
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

func (s *Store) NewRelayID() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.nextRelayIDLocked()
}

func (s *Store) Records() ([]record.Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.recordsLocked()
}

func (s *Store) recordsLocked() ([]record.Record, error) {
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
			return nil, ErrChecksum{RelayID: strings.TrimSuffix(entry.Name(), ".json")}
		}
		records = append(records, rec)
	}
	return records, nil
}

func (s *Store) Read(relayID string) (record.Record, error) {
	data, err := os.ReadFile(filepath.Join(s.Root, "records", relayID+".json"))
	if err != nil {
		if os.IsNotExist(err) {
			if qerr, ok := s.quarantinedError(relayID); ok {
				return record.Record{}, qerr
			}
		}
		return record.Record{}, err
	}
	rec, err := record.Verify(data)
	if err != nil {
		return record.Record{}, ErrChecksum{RelayID: relayID}
	}
	return rec, nil
}

func (s *Store) Project(seat string) ([]string, error) {
	return s.ProjectView(seat, ProjectDefault)
}

func (s *Store) ProjectView(seat, view string) ([]string, error) {
	if view == "" {
		view = ProjectDefault
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	records, err := s.recordsLocked()
	if err != nil {
		return nil, err
	}
	switch view {
	case ProjectDefault:
		return s.projectDefaultLocked(seat, records)
	case ProjectAudit:
		return projectAudit(seat, records), nil
	default:
		return nil, fmt.Errorf("unknown project view: %s", view)
	}
}

func (s *Store) projectDefaultLocked(seat string, records []record.Record) ([]string, error) {
	path := filepath.Join(s.Root, "mailboxes", safeMailbox(seat)+".jsonl")
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var relayIDs []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			relayIDs = append(relayIDs, line)
		}
	}
	if err := scanner.Err(); err != nil {
		_ = f.Close()
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	return filterDefaultProject(seat, relayIDs, records), nil
}

func filterDefaultProject(seat string, relayIDs []string, records []record.Record) []string {
	byRelay := make(map[string]record.Record, len(records))
	for _, rec := range records {
		byRelay[rec.Envelope.RelayID] = rec
	}
	seen := map[string]bool{}
	var out []string
	for _, relayID := range relayIDs {
		rec, ok := byRelay[relayID]
		if !ok || rec.Envelope.DeliveryState != record.Accepted || seen[relayID] {
			continue
		}
		if !recordDeliveredTo(rec, seat) {
			continue
		}
		seen[relayID] = true
		out = append(out, relayID)
	}
	return out
}

func projectAudit(seat string, records []record.Record) []string {
	var out []string
	for _, rec := range records {
		if rec.Envelope.From == seat && rec.Envelope.RelayID != "" {
			out = append(out, rec.Envelope.RelayID)
		}
	}
	return out
}

func recordDeliveredTo(rec record.Record, seatName string) bool {
	recipients, err := DeliveryRecipients(rec)
	if err != nil {
		return false
	}
	for _, recipient := range recipients {
		if recipient == seatName {
			return true
		}
	}
	return false
}

func (s *Store) PendingDeliveryFor(seat string) (bool, error) {
	data, err := os.ReadFile(filepath.Join(s.Root, "mailboxes", safeMailbox(seat)+".jsonl"))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(data)) != "", nil
}

func (s *Store) PendingDeliverySeats() ([]string, error) {
	dir := filepath.Join(s.Root, "mailboxes")
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var seats []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".jsonl") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(string(data)) == "" {
			continue
		}
		seats = append(seats, strings.TrimSuffix(entry.Name(), ".jsonl"))
	}
	return seats, nil
}

func (s *Store) appendRedo(relayID string, intents []Intent) error {
	dir := filepath.Join(s.Root, "journal", "redo")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	seq, path, err := activeRedoSegment(dir)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	crashpoint.Hit("pre_redo_fsync")
	data, err := json.Marshal(redoEntry{RelayID: relayID, Intents: intents})
	if err != nil {
		_ = f.Close()
		return err
	}
	data = append(data, '\n')
	if err := fsio.AppendFsync(f, data); err != nil {
		_ = f.Close()
		return err
	}
	crashpoint.Hit("post_redo_fsync")
	if err := f.Close(); err != nil {
		return err
	}
	return rotateRedoAfterAppend(dir, seq, path)
}

func (s *Store) nextRelayIDLocked() (string, error) {
	for range 16 {
		buf := make([]byte, 12)
		if _, err := rand.Read(buf); err != nil {
			return "", err
		}
		relayID := "relay-" + hex.EncodeToString(buf)
		if _, err := os.Stat(filepath.Join(s.Root, "records", relayID+".json")); os.IsNotExist(err) {
			return relayID, nil
		} else if err != nil {
			return "", err
		}
	}
	return "", errors.New("relay id collision")
}

type redoEntry struct {
	RelayID string   `json:"relay_id"`
	Intents []Intent `json:"intents"`
}

func readRedo(root string) ([]redoEntry, error) {
	dir := filepath.Join(root, "journal", "redo")
	seqs, err := redoSegmentSeqs(dir)
	if err != nil {
		return nil, err
	}
	var entries []redoEntry
	for _, seq := range seqs {
		path := filepath.Join(dir, fmt.Sprintf("%06d.jsonl", seq))
		data, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		lines := bytes.Split(data, []byte("\n"))
		completeTail := bytes.HasSuffix(data, []byte("\n"))
		for i, line := range lines {
			if len(line) == 0 {
				continue
			}
			var entry redoEntry
			if err := json.Unmarshal(line, &entry); err != nil {
				if i == len(lines)-1 && !completeTail {
					break
				}
				return nil, err
			}
			entries = append(entries, entry)
		}
	}
	return entries, nil
}

func activeRedoSegment(dir string) (int, string, error) {
	seqs, err := redoSegmentSeqs(dir)
	if err != nil {
		return 0, "", err
	}
	if len(seqs) == 0 {
		return 1, filepath.Join(dir, "000001.jsonl"), nil
	}
	seq := seqs[len(seqs)-1]
	path := filepath.Join(dir, fmt.Sprintf("%06d.jsonl", seq))
	info, err := os.Stat(path)
	if err != nil {
		return 0, "", err
	}
	if info.Size() > defaultSegmentRotateBytes {
		next := seq + 1
		nextPath := filepath.Join(dir, fmt.Sprintf("%06d.jsonl", next))
		if err := createEmptyRedoSegment(nextPath); err != nil {
			return 0, "", err
		}
		return next, nextPath, nil
	}
	return seq, path, nil
}

func rotateRedoAfterAppend(dir string, seq int, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.Size() <= defaultSegmentRotateBytes {
		return nil
	}
	nextPath := filepath.Join(dir, fmt.Sprintf("%06d.jsonl", seq+1))
	if _, err := os.Stat(nextPath); err == nil {
		return nil
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}
	crashpoint.Hit("pre_segment_rotate")
	if err := createEmptyRedoSegment(nextPath); err != nil {
		return err
	}
	crashpoint.Hit("post_segment_rotate")
	return nil
}

func redoSegmentSeqs(dir string) ([]int, error) {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var seqs []int
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".jsonl") {
			continue
		}
		seq, err := strconv.Atoi(strings.TrimSuffix(entry.Name(), ".jsonl"))
		if err != nil {
			continue
		}
		seqs = append(seqs, seq)
	}
	sort.Ints(seqs)
	return seqs, nil
}

func createEmptyRedoSegment(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	return f.Close()
}
