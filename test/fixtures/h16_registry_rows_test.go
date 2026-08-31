package fixtures_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/observe"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

const h16FieldspecVersion = "s12-fieldspec-v9"

func TestH16RegistryRowsAndRuledRecordKinds(t *testing.T) {
	data := h16RegistryBytes(t)
	var raw struct {
		Version    string                       `json:"version"`
		Provenance map[string]string            `json:"provenance"`
		NamedEnums map[string][]string          `json:"named_enums"`
		Fields     []map[string]json.RawMessage `json:"fields"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("decode registry bytes: %v", err)
	}
	if raw.Version != h16FieldspecVersion {
		t.Fatalf("version=%q, want %q", raw.Version, h16FieldspecVersion)
	}
	wantProvenance := map[string]string{
		"owner":         "m-2",
		"design_doc_id": "h16-outcome-split-design",
		"plan_lock_id":  "s12-h16-fix-plan",
		"supersedes":    "s10-fieldspec-v8",
		"note":          "s12 H-16-REG realization under the operator-opened step3-h16-h26-lane; three system-stamped header rows + the ruled record_kind tokens",
	}
	if !reflect.DeepEqual(raw.Provenance, wantProvenance) {
		t.Fatalf("provenance=%#v, want %#v", raw.Provenance, wantProvenance)
	}

	wantRow := func(id string) map[string]json.RawMessage {
		return map[string]json.RawMessage{
			"id":               json.RawMessage(`"` + id + `"`),
			"layer":            json.RawMessage(`"header"`),
			"owner":            json.RawMessage(`"system"`),
			"type":             json.RawMessage(`"string"`),
			"fill_constraints": json.RawMessage(`"system_only"`),
			"lineage_role":     json.RawMessage(`"none"`),
		}
	}
	for _, id := range []string{"hook_contract", "mint_predecessor", "admin_provenance"} {
		var got map[string]json.RawMessage
		for _, row := range raw.Fields {
			if bytes.Equal(row["id"], json.RawMessage(`"`+id+`"`)) {
				got = row
				break
			}
		}
		if !reflect.DeepEqual(got, wantRow(id)) {
			t.Fatalf("%s raw row=%#v, want exact %#v", id, got, wantRow(id))
		}
	}

	// Authority: m-2 DESIGN-planner-m2-20260817-201520.md section 1a/1b.
	wantKinds := []string{"genesis", "owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint", "odb", "resummon_command", "mint-chain-anchor", "attempt_resolution", "derived-work-attempt", "derived-work-transition"}
	if !reflect.DeepEqual(raw.NamedEnums["record_kind"], wantKinds) {
		t.Fatalf("record_kind=%q, want %q", raw.NamedEnums["record_kind"], wantKinds)
	}

	reg := loadH16Registry(t)
	recordKind, ok := reg.ByID("record_kind")
	if !ok {
		t.Fatal("registry missing record_kind")
	}
	wantScopes := map[string][]string{
		"operator": {"owed_item", "owed_disposition", "gate_resolution", "disposition", "diagnostics", "config_change", "waiver_retraction", "seat_mint", "mint-chain-anchor", "attempt_resolution"},
		"*":        {"diagnostics"},
	}
	if !reflect.DeepEqual(recordKind.SeatScope, wantScopes) {
		t.Fatalf("record_kind seat_scope=%#v, want %#v", recordKind.SeatScope, wantScopes)
	}
}

func TestH16FieldspecV9ReaderAndForwardTransition(t *testing.T) {
	data := h16RegistryBytes(t)
	v9 := h16RegistryAtVersion(t, data, h16FieldspecVersion)
	v8 := h16RegistryAtVersion(t, data, "s10-fieldspec-v8")
	v7 := h16RegistryAtVersion(t, data, "s8-fieldspec-v7")

	if err := config.ValidateFieldspecReaderMarker(v9, "s7a-fieldspec-v5", "s8-fieldspec-v6", "s8-fieldspec-v7", "s10-fieldspec-v8", h16FieldspecVersion); err != nil {
		t.Fatalf("v9 reader marker: %v", err)
	}
	if err := config.ValidateMemberTransition("fieldspec", v8, v9); err != nil {
		t.Fatalf("v8->v9 transition: %v", err)
	}
	if err := config.ValidateMemberTransition("fieldspec", v7, v9); !errors.Is(err, config.ErrConfigVersionTransition) {
		t.Fatalf("v7->v9 transition err=%v, want ErrConfigVersionTransition", err)
	}
}

func TestH16FreshGenesisIsBornAtV9AndPinsDigest(t *testing.T) {
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	wantFieldspec, err := os.ReadFile(sources["fieldspec"])
	if err != nil {
		t.Fatalf("read v9 source: %v", err)
	}
	if err := store.Init(root, sources); err != nil {
		t.Fatalf("Init fresh v9 store: %v", err)
	}

	gotFieldspec, err := os.ReadFile(filepath.Join(root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read materialized fieldspec: %v", err)
	}
	if !bytes.Equal(gotFieldspec, wantFieldspec) {
		t.Fatal("fresh genesis did not preserve v9 fieldspec source bytes")
	}
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load fresh v9 config: %v", err)
	}
	if pinned.Registry.Version != h16FieldspecVersion {
		t.Fatalf("fresh genesis fieldspec=%q, want %q", pinned.Registry.Version, h16FieldspecVersion)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open fresh v9 store: %v", err)
	}
	genesis, err := st.Genesis()
	if err != nil {
		t.Fatalf("read genesis: %v", err)
	}
	if genesis.Headers["config_digest"] != pinned.Digest {
		t.Fatalf("genesis digest=%q, want current v9 digest %q", genesis.Headers["config_digest"], pinned.Digest)
	}
}

func TestH16RegistryRowsJoinPresenceAndLaneSupplyGuards(t *testing.T) {
	reg := loadH16Registry(t)
	assertH16SystemHeaderPopulation(t, reg, 34)
	for _, id := range []string{"hook_contract", "mint_predecessor", "admin_provenance"} {
		t.Run(id, func(t *testing.T) {
			cand := record.Record{Headers: map[string]string{id: "forged"}}
			if got := observe.LaneSuppliedSystemField(reg, cand); got != id {
				t.Fatalf("LaneSuppliedSystemField=%q, want %q", got, id)
			}
		})
	}
}

func h16RegistryBytes(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	return data
}

func h16RegistryAtVersion(t *testing.T, data []byte, version string) []byte {
	t.Helper()
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("decode registry: %v", err)
	}
	raw["version"] = version
	result, err := json.Marshal(raw)
	if err != nil {
		t.Fatalf("marshal registry at %s: %v", version, err)
	}
	return result
}
