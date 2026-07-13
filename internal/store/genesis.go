package store

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
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
	ErrGenesisExists       = errors.New("genesis already exists")
	ErrGenesisMissing      = errors.New("genesis missing")
	ErrStoreNotAdopted     = errors.New("store-not-adopted: run frank -bless")
	ErrStoreAlreadyAdopted = errors.New("store already adopted")
)

const genesisFieldspecV5SHA256 = "1ef6abab4d496b11017f57ca400e8296d63824994ffce8311e4533f70cc92485"

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
		data, err = genesisMemberBytes(name, data)
		if err != nil {
			return err
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

func genesisMemberBytes(name string, source []byte) ([]byte, error) {
	if name != "fieldspec" {
		return source, nil
	}
	predecessor := append([]byte(nil), source...)
	isV8 := bytes.Contains(predecessor, []byte(`"version": "s10-fieldspec-v8"`))
	if isV8 {
		v8Marker := []byte(`"version": "s10-fieldspec-v8"`)
		v8Kinds := []byte(`"record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint", "odb", "resummon_command"]`)
		v7Kinds := []byte(`"record_kind": ["genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint"]`)
		if bytes.Count(predecessor, v8Marker) != 1 || bytes.Count(predecessor, v8Kinds) != 1 {
			return nil, fmt.Errorf("fieldspec v8 predecessor delta mismatch")
		}
		predecessor = bytes.Replace(predecessor, v8Marker, []byte(`"version": "s8-fieldspec-v7"`), 1)
		predecessor = bytes.Replace(predecessor, v8Kinds, v7Kinds, 1)
	}
	if bytes.Contains(predecessor, []byte(`"version": "s8-fieldspec-v7"`)) {
		claimRow := []byte("    " + executableClaimsV7Row + ",\n")
		if bytes.Count(predecessor, claimRow) != 1 {
			return nil, fmt.Errorf("fieldspec v7 claim row mismatch")
		}
		predecessor = bytes.Replace(predecessor, []byte(`"version": "s8-fieldspec-v7"`), []byte(`"version": "s8-fieldspec-v6"`), 1)
		predecessor = bytes.Replace(predecessor, claimRow, nil, 1)
	}
	if !bytes.Contains(predecessor, []byte(`"version": "s8-fieldspec-v6"`)) {
		return source, nil
	}
	replacements := [][2]string{
		{`"version": "s8-fieldspec-v6"`, `"version": "s7a-fieldspec-v5"`},
		{`"config_member": ["fieldspec", "engine", "catalog", "adoption"]`, `"config_member": ["fieldspec", "engine"]`},
		{`"seat_scope": {"operator": ["fieldspec", "engine", "catalog", "adoption"]}`, `"seat_scope": {"operator": ["fieldspec", "engine"]}`},
		{`"annotation": "non-gate-bearing records only; static required_when/visible_when REMOVED per Option B (s8 v6 review-driven amendment) - applicability = the m-5/m-6 producer/profile manifest at step-4.5 completeness, NOT a registry predicate; conductor-derived, never lane-authored (R2). suppliability guard = the s5-b (h) typed-REJECT validator rule (channel-keyed, envelope-asymmetry preserved); until it lands, dormancy is render-absence, not submit-rejection - no non-lane-writability claim."`, `"required_when": {"all_of": [{"layer_present": "observe"}]}, "visible_when": {"all_of": [{"layer_present": "observe"}]}, "annotation": "non-gate-bearing records only; Step-2+ semantics, annotation only. suppliability guard = the s5-b (h) typed-REJECT validator rule (channel-keyed, envelope-asymmetry preserved); until it lands, dormancy is render-absence, not submit-rejection - no non-lane-writability claim."`},
	}
	for _, replacement := range replacements {
		old, next := []byte(replacement[0]), []byte(replacement[1])
		if bytes.Count(predecessor, old) != 1 {
			return nil, fmt.Errorf("fieldspec genesis predecessor mismatch")
		}
		predecessor = bytes.Replace(predecessor, old, next, 1)
	}
	sum := sha256.Sum256(predecessor)
	if hex.EncodeToString(sum[:]) != genesisFieldspecV5SHA256 {
		return nil, fmt.Errorf("fieldspec genesis predecessor hash mismatch")
	}
	if isV8 {
		return source, nil
	}
	return predecessor, nil
}

const executableClaimsV7Row = `{"id": "executable_claims", "layer": "header", "owner": "agent_enum_pick", "type": "row_array", "gate_referenceable": false, "fill_constraints": "free_text", "lineage_role": "none", "visible_when": {"all_of": [{"layer_present": "observe"}]}, "annotation": "seat-declared executable-claim DECLARATION input (counterpart to system-owned executable_claim_results). Columns: claim_ref (non-empty/bounded/symbolic/UNIQUE - duplicate = typed reject), check_id, params. Nested check_id/params validation is registry-schema-aware at the m-3 fill-time + authoritative observe-time seam (typed rejects: unknown-check / check-params-invalid) - NOT a top-level enum_set (row_array columns carry no FieldSpec enum). R2: non-gate-referenceable (selection = lane intent; disposition rides system-owned executable_claim_results). Rail A (s8-claim-input v7, ratified absence-open/present-closed): ABSENCE = honest Evaluate:nil no-vantage degrade; a PRESENT declaration = CLOSED/fail-closed via the governed v6->v7 transition + m-7 marker-first v7 capability ceiling (a non-v7 reader refuses at phase-0, never silently ignores). suppliability guard = the s5-b (h) typed-REJECT validator rule (channel-keyed, envelope-asymmetry preserved); until it lands, dormancy is render-absence, not submit-rejection - no non-lane-writability claim."}`

func StoreRootConfigPaths(root string) map[string]string {
	paths := LegacyStoreRootConfigPaths(root)
	catalog := filepath.Join(root, "config", "catalog", "catalog.json")
	if _, err := os.Stat(catalog); err == nil {
		paths["catalog"] = catalog
	}
	return paths
}

// LegacyStoreRootConfigPaths is the explicit pre-adoption two-member
// expectation used only by bless and its migration fixtures.
func LegacyStoreRootConfigPaths(root string) map[string]string {
	return map[string]string{
		"engine":    filepath.Join(root, "config", "engine.json"),
		"fieldspec": filepath.Join(root, "config", "fieldspec", "registry.json"),
	}

}

// RequireAdoptedConfig prevents the normal serve path from silently using the
// legacy two-member expectation. Only offline bless may load that state.
func RequireAdoptedConfig(root string) error {
	_, err := os.Stat(filepath.Join(root, "config", "catalog", "catalog.json"))
	if errors.Is(err, os.ErrNotExist) {
		return ErrStoreNotAdopted
	}
	return err
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

// BlessS8 performs the one offline two-member to three-member adoption. The
// root lock is acquired here so callers cannot accidentally run it alongside
// the serving process.
func BlessS8(root string, candidateSources map[string]string) error {
	rootLock, err := AcquireRoot(root)
	if err != nil {
		return err
	}
	defer func() { _ = rootLock.Release() }()

	st, err := Open(root)
	if err != nil {
		return err
	}
	if adopted, err := st.hasAdoptionRecord(); err != nil {
		return err
	} else if adopted {
		if err := st.CompleteAdoptionConfig(); err != nil {
			return err
		}
		pinned, err := config.Load(StoreRootConfigPaths(root))
		if err != nil {
			return err
		}
		if err := st.ValidateGenesis(pinned); err != nil {
			return err
		}
		return ErrStoreAlreadyAdopted
	}
	if _, err := os.Stat(filepath.Join(root, "config", "catalog", "catalog.json")); err == nil {
		return errors.New("catalog member exists without adoption record")
	} else if !os.IsNotExist(err) {
		return err
	}

	legacy, err := config.Load(LegacyStoreRootConfigPaths(root))
	if err != nil {
		return err
	}
	if err := st.ValidateGenesis(legacy); err != nil {
		return err
	}
	engineSource, catalogSource := candidateSources["engine"], candidateSources["catalog"]
	if engineSource == "" || catalogSource == "" || len(candidateSources) != 2 {
		return errors.New("bless requires exactly engine and catalog candidates")
	}
	candidatePaths := map[string]string{
		"fieldspec": LegacyStoreRootConfigPaths(root)["fieldspec"],
		"engine":    engineSource,
		"catalog":   catalogSource,
	}
	candidate, err := config.Load(candidatePaths)
	if err != nil {
		return err
	}
	if candidate.Engine.Version != 2 || candidate.Supply == nil {
		return errors.New("current engine v2 with governed supply required")
	}
	for name, present := range candidate.Engine.PresentLayers {
		if present {
			return fmt.Errorf("bless engine candidate optional layer %s must be false", name)
		}
	}
	body := adoptionBody{Members: make([]adoptionMember, 0, 2)}
	for _, name := range []string{"catalog", "engine"} {
		data := candidate.Members[name]
		body.Members = append(body.Members, adoptionMember{Name: name, BytesB64: base64.StdEncoding.EncodeToString(data)})
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID: "s8-adoption", DispatchID: "s8-adoption", From: "operator", Role: "operator",
			DeliveryState: record.Accepted, SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE": "IMPL", "SUBJECT": "offline s8 store adoption", "record_kind": "config_change",
			"member": "adoption", "new_digest": candidate.Digest, "channel": "bless",
		},
		Body: string(bodyBytes),
	}
	intents, err := ConfigChangeIntentsStrict(rec)
	if err != nil {
		return err
	}
	_, err = st.Commit(rec, intents)
	return err
}

// CompleteAdoptionConfig replays only the derived member files of a committed
// adoption record. It is safe before the full phase-0 load and is idempotent.
func (s *Store) CompleteAdoptionConfig() error {
	chain, err := s.acceptedConfigChanges()
	if err != nil {
		return err
	}
	adoptionAt := -1
	for i, rec := range chain {
		if rec.Headers["member"] == "adoption" {
			adoptionAt = i
			break
		}
	}
	if adoptionAt < 0 {
		return nil
	}
	var configIntents []Intent
	for _, rec := range chain[adoptionAt:] {
		intents, err := ConfigChangeIntentsStrict(rec)
		if err != nil {
			return err
		}
		configIntents = append(configIntents, configOnlyIntents(intents)...)
	}
	if len(configIntents) == 0 {
		return errors.New("adoption record has no config intents")
	}
	return s.applyIntents(configIntents)
}

func (s *Store) hasAdoptionRecord() (bool, error) {
	chain, err := s.acceptedConfigChanges()
	if err != nil {
		return false, err
	}
	for _, rec := range chain {
		if rec.Headers["member"] == "adoption" {
			return true, nil
		}
	}
	return false, nil
}

func (s *Store) rematerializeLatestConfigChange() error {
	chain, err := s.acceptedConfigChanges()
	if err != nil {
		return err
	}
	if len(chain) == 0 {
		return errors.New("no config_change chain")
	}
	var intents []Intent
	for _, rec := range chain {
		recIntents, err := ConfigChangeIntentsStrict(rec)
		if err != nil {
			return err
		}
		intents = append(intents, configOnlyIntents(recIntents)...)
	}
	if len(intents) == 0 {
		return errors.New("no config_change members")
	}
	return s.applyIntents(intents)
}

func configOnlyIntents(intents []Intent) []Intent {
	out := make([]Intent, 0, len(intents))
	for _, intent := range intents {
		if intent.Kind == IntentConfig {
			out = append(out, intent)
		}
	}
	return out
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
	case "catalog":
		return filepath.Join("config", "catalog", "catalog.json"), nil
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
