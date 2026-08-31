package zeroloss

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/migrate"
	"github.com/The-Frank-Organization/frank/internal/store"
)

type Report struct {
	Total       int
	Read        int
	Lost        int
	Quarantined int
	Views       map[string]migrate.View
}

func Replay(root string, reg *migrate.Registry) (Report, error) {
	st, err := store.Open(root)
	if err != nil {
		return Report{}, err
	}
	ids, err := replayIDs(root)
	if err != nil {
		return Report{}, err
	}
	reader := migrate.Reader{Store: st, Registry: reg}
	report := Report{Total: len(ids), Views: map[string]migrate.View{}}
	for _, relayID := range ids {
		view, err := reader.Read(relayID)
		if err == nil {
			report.Read++
			report.Views[relayID] = view
			continue
		}
		var checksum store.ErrChecksum
		var quarantined store.ErrQuarantined
		switch {
		case errors.As(err, &checksum), errors.As(err, &quarantined):
			report.Quarantined++
		case errors.Is(err, os.ErrNotExist):
			report.Lost++
		default:
			return Report{}, err
		}
	}
	return report, nil
}

func replayIDs(root string) ([]string, error) {
	seen := map[string]bool{}
	add := func(relayID string) {
		if relayID != "" {
			seen[relayID] = true
		}
	}
	redoIDs, err := readRedoIDs(root)
	if err != nil {
		return nil, err
	}
	for _, relayID := range redoIDs {
		add(relayID)
	}
	recordIDs, err := readRecordIDs(root)
	if err != nil {
		return nil, err
	}
	for _, relayID := range recordIDs {
		add(relayID)
	}
	ids := make([]string, 0, len(seen))
	for relayID := range seen {
		ids = append(ids, relayID)
	}
	sort.Strings(ids)
	return ids, nil
}

type redoEntry struct {
	RelayID string `json:"relay_id"`
}

func readRedoIDs(root string) ([]string, error) {
	dir := filepath.Join(root, "journal", "redo")
	entries, err := os.ReadDir(dir)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".jsonl" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
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
			ids = append(ids, entry.RelayID)
		}
	}
	return ids, nil
}

func readRecordIDs(root string) ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(root, "records"))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		ids = append(ids, strings.TrimSuffix(entry.Name(), ".json"))
	}
	return ids, nil
}
