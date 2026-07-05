package store

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fsio"
	"github.com/jackli/frank/internal/record"
)

var (
	ErrGenesisExists  = errors.New("genesis already exists")
	ErrGenesisMissing = errors.New("genesis missing")
)

type ErrDigestMismatch struct {
	Want string
	Got  string
}

func (e ErrDigestMismatch) Error() string {
	return fmt.Sprintf("genesis config digest mismatch: want %s got %s", e.Want, e.Got)
}

func Init(root string, sources map[string]string) error {
	if _, err := os.Stat(filepath.Join(root, "records", "genesis.json")); err == nil {
		return ErrGenesisExists
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	for name, source := range sources {
		target, err := configTarget(name)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(source)
		if err != nil {
			return fmt.Errorf("read config source %s: %w", name, err)
		}
		if err := fsio.WriteFileAtomic(root, target, data); err != nil {
			return err
		}
	}

	pinned, err := config.Load(StoreRootConfigPaths(root))
	if err != nil {
		return err
	}
	st, err := Open(root)
	if err != nil {
		return err
	}
	seed, err := addressSpaceSeed()
	if err != nil {
		return err
	}
	_, err = st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "genesis",
			DispatchID:    "genesis",
			From:          "system",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"record_kind":        "genesis",
			"config_digest":      pinned.Digest,
			"address_space_seed": seed,
			"created_ts":         time.Now().UTC().Format(time.RFC3339Nano),
		},
	}, nil)
	if err != nil && err.Error() == "record already exists" {
		return ErrGenesisExists
	}
	return err
}

func StoreRootConfigPaths(root string) map[string]string {
	return map[string]string{
		"engine":    filepath.Join(root, "config", "engine.json"),
		"fieldspec": filepath.Join(root, "config", "fieldspec", "registry.json"),
	}
}

func (s *Store) Genesis() (record.Record, error) {
	rec, err := s.Read("genesis")
	if err == nil {
		return rec, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return record.Record{}, ErrGenesisMissing
	}
	return record.Record{}, err
}

func (s *Store) ValidateGenesis(pinned *config.Pinned) error {
	expected, err := s.ExpectedConfigDigest()
	if err != nil {
		return err
	}
	got, loadErr := s.currentConfigDigest()
	if loadErr == nil && expected != "" && got == expected {
		return nil
	}
	if err := s.rematerializeLatestConfigChange(); err == nil {
		got, loadErr = s.currentConfigDigest()
		if loadErr == nil && expected != "" && got == expected {
			return nil
		}
	}
	if loadErr != nil {
		got = "unreadable"
	}
	return ErrDigestMismatch{Want: expected, Got: got}
}

func (s *Store) ExpectedConfigDigest() (string, error) {
	genesis, err := s.Genesis()
	if err != nil {
		return "", err
	}
	expected := genesis.Headers["config_digest"]
	chain, err := s.acceptedConfigChanges()
	if err != nil {
		return "", err
	}
	for _, rec := range chain {
		if digest := rec.Headers["new_digest"]; digest != "" {
			expected = digest
		}
	}
	return expected, nil
}

func (s *Store) currentConfigDigest() (string, error) {
	pinned, err := config.Load(StoreRootConfigPaths(s.Root))
	if err != nil {
		return "", err
	}
	return pinned.Digest, nil
}

func (s *Store) rematerializeLatestConfigChange() error {
	chain, err := s.acceptedConfigChanges()
	if err != nil {
		return err
	}
	if len(chain) == 0 {
		return errors.New("no config_change chain")
	}
	latest := map[string]record.Record{}
	for _, rec := range chain {
		member := rec.Headers["member"]
		if _, err := configTarget(member); err != nil {
			return err
		}
		latest[member] = rec
	}
	var intents []Intent
	for _, member := range []string{"engine", "fieldspec"} {
		rec, ok := latest[member]
		if !ok {
			continue
		}
		target, err := configTarget(member)
		if err != nil {
			return err
		}
		intents = append(intents, Intent{Kind: IntentConfig, Path: target, Payload: []byte(rec.Body)})
	}
	if len(intents) == 0 {
		return errors.New("no config_change members")
	}
	return s.applyIntents(intents)
}

func (s *Store) acceptedConfigChanges() ([]record.Record, error) {
	entries, err := readRedo(s.Root)
	if err != nil {
		return nil, err
	}
	var chain []record.Record
	for _, entry := range entries {
		rec, err := s.Read(entry.RelayID)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			var checksum ErrChecksum
			var quarantined ErrQuarantined
			if errors.As(err, &checksum) || errors.As(err, &quarantined) {
				continue
			}
			return nil, err
		}
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["record_kind"] == "config_change" {
			chain = append(chain, rec)
		}
	}
	return chain, nil
}

func configTarget(name string) (string, error) {
	switch name {
	case "engine":
		return filepath.Join("config", "engine.json"), nil
	case "fieldspec":
		return filepath.Join("config", "fieldspec", "registry.json"), nil
	default:
		return "", fmt.Errorf("unknown config member %s", name)
	}
}

func addressSpaceSeed() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf) + ":operator", nil
}
