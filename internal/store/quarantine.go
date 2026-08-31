package store

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/crashpoint"
	"github.com/The-Frank-Organization/frank/internal/record"
)

const checksumMismatch = "checksum-mismatch"

type ErrChecksum struct {
	RelayID string
}

func (e ErrChecksum) Error() string {
	return fmt.Sprintf("checksum-mismatch: relay_id=%s", e.RelayID)
}

type ErrQuarantined struct {
	RelayID      string
	IncidentID   string
	FailureClass string
}

func (e ErrQuarantined) Error() string {
	if e.IncidentID == "" {
		return fmt.Sprintf("record-quarantined: relay_id=%s failure_class=%s", e.RelayID, e.FailureClass)
	}
	return fmt.Sprintf("record-quarantined: relay_id=%s incident_id=%s failure_class=%s", e.RelayID, e.IncidentID, e.FailureClass)
}

func (s *Store) ScanQuarantine() ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(s.Root, "records"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var evicted []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		relayID := strings.TrimSuffix(entry.Name(), ".json")
		data, err := os.ReadFile(filepath.Join(s.Root, "records", entry.Name()))
		if err != nil {
			return nil, err
		}
		if _, err := record.Verify(data); err == nil {
			continue
		}
		ok, err := s.QuarantineOne(relayID)
		if err != nil {
			return nil, err
		}
		if ok {
			evicted = append(evicted, relayID)
		}
	}
	return evicted, nil
}

func (s *Store) QuarantineOne(relayID string) (bool, error) {
	recordPath := filepath.Join(s.Root, "records", relayID+".json")
	quarantinePath := filepath.Join(s.Root, "quarantine", relayID+".json")
	if _, err := os.Stat(recordPath); errors.Is(err, os.ErrNotExist) {
		if _, qerr := os.Stat(quarantinePath); qerr == nil || errors.Is(qerr, os.ErrNotExist) {
			return false, nil
		} else {
			return false, qerr
		}
	} else if err != nil {
		return false, err
	}
	if err := os.MkdirAll(filepath.Dir(quarantinePath), 0o755); err != nil {
		return false, err
	}
	crashpoint.Hit("pre_quarantine_evict")
	if err := os.Rename(recordPath, quarantinePath); err != nil {
		return false, err
	}
	crashpoint.Hit("post_quarantine_evict")
	return true, nil
}

func (s *Store) CompleteIncidents() error {
	entries, err := os.ReadDir(filepath.Join(s.Root, "quarantine"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		relayID := strings.TrimSuffix(entry.Name(), ".json")
		incidentID := incidentRelayID(relayID)
		if _, err := os.Stat(filepath.Join(s.Root, "records", incidentID+".json")); err == nil {
			continue
		} else if err != nil && !os.IsNotExist(err) {
			return err
		}
		_, err := s.Commit(record.Record{
			Envelope: record.Envelope{
				RelayID:       incidentID,
				DispatchID:    "quarantine",
				From:          "system",
				Role:          "system",
				DeliveryState: record.Held,
				SchemaVersion: 1,
			},
			Headers: map[string]string{
				"record_kind":     "incident",
				"quarantined_ref": relayID,
				"failure_class":   checksumMismatch,
			},
			Body: "record:" + checksumMismatch,
		}, nil)
		if err != nil && err.Error() == "record already exists" {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) quarantinedError(relayID string) (error, bool) {
	if _, err := os.Stat(filepath.Join(s.Root, "quarantine", relayID+".json")); errors.Is(err, os.ErrNotExist) {
		return nil, false
	} else if err != nil {
		return err, true
	}
	out := ErrQuarantined{RelayID: relayID, FailureClass: checksumMismatch}
	incidentID := incidentRelayID(relayID)
	data, err := os.ReadFile(filepath.Join(s.Root, "records", incidentID+".json"))
	if err != nil {
		return out, true
	}
	incident, err := record.Verify(data)
	if err != nil {
		return out, true
	}
	out.IncidentID = incident.Envelope.RelayID
	if failure := incident.Headers["failure_class"]; failure != "" {
		out.FailureClass = failure
	}
	return out, true
}

func incidentRelayID(relayID string) string {
	return "incident-" + relayID
}
