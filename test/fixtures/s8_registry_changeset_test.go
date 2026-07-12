package fixtures_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
)

func TestS8RegistryChangesetIsExactV6OptionB(t *testing.T) {
	reg := loadS5Registry(t)
	if reg.Version != "s8-fieldspec-v6" {
		t.Fatalf("version = %q, want s8-fieldspec-v6", reg.Version)
	}
	wantMembers := []string{"fieldspec", "engine", "catalog", "adoption"}
	if got := reg.NamedEnums["config_member"]; !reflect.DeepEqual(got, wantMembers) {
		t.Fatalf("config_member = %#v, want %#v", got, wantMembers)
	}
	member, ok := reg.ByID("member")
	if !ok {
		t.Fatal("member field missing")
	}
	if got := member.SeatScope["operator"]; !reflect.DeepEqual(got, wantMembers) {
		t.Fatalf("member operator scope = %#v, want %#v", got, wantMembers)
	}
	surface, ok := reg.ByID("surface_intent")
	if !ok {
		t.Fatal("surface_intent missing")
	}
	if len(surface.RequiredWhen.Raw) != 0 || len(surface.VisibleWhen.Raw) != 0 {
		t.Fatalf("surface_intent retains static predicates required=%s visible=%s", surface.RequiredWhen.Raw, surface.VisibleWhen.Raw)
	}
}

func TestS8RegistryTokensControlLiveOperatorFillGate(t *testing.T) {
	reg := loadS5Registry(t)
	operator := fieldspec.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	form, _ := reg.Render(fieldspec.RenderEnv{}, operator, "SITREP", "medium", fieldspec.ClosedGrantState)
	for _, token := range []string{"catalog", "adoption"} {
		if !form.OptionAllowed("member", token) {
			t.Fatalf("operator form omitted config member %q", token)
		}
	}

	data, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	for _, missing := range []string{"catalog", "adoption"} {
		mutated := bytes.ReplaceAll(data, []byte(`, "`+missing+`"`), nil)
		path := filepath.Join(t.TempDir(), "registry.json")
		if err := os.WriteFile(path, mutated, 0o644); err != nil {
			t.Fatalf("write registry without %s: %v", missing, err)
		}
		without, err := fieldspec.Load(path)
		if err != nil {
			t.Fatalf("load registry without %s: %v", missing, err)
		}
		withoutForm, _ := without.Render(fieldspec.RenderEnv{}, operator, "SITREP", "medium", fieldspec.ClosedGrantState)
		if withoutForm.OptionAllowed("member", missing) {
			t.Fatalf("member %q fill-passed without both enum and seat-scope tokens", missing)
		}
	}
}

func TestS8FXCFG10MemberTransitionRelation(t *testing.T) {
	currentEngine := []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`)
	valueOnly := []byte(`{"version":1,"gc_enabled":true,"segment_rotate_bytes":4194304,"present_layers":{"observe":true}}`)
	if err := config.ValidateMemberTransition("engine", currentEngine, valueOnly); err != nil {
		t.Fatalf("value-only transition: %v", err)
	}
	for name, candidate := range map[string][]byte{
		"added-key":    []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false},"new_key":true}`),
		"type-change":  []byte(`{"version":1,"gc_enabled":"no","segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`),
		"shape-change": []byte(`{"version":1,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":{"nested":true}}}`),
		"rollback":     []byte(`{"version":0,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`),
		"skip":         []byte(`{"version":3,"gc_enabled":false,"segment_rotate_bytes":4194304,"present_layers":{"observe":false}}`),
	} {
		t.Run(name, func(t *testing.T) {
			if err := config.ValidateMemberTransition("engine", currentEngine, candidate); err == nil {
				t.Fatalf("transition unexpectedly accepted: %s", candidate)
			}
		})
	}

	v5, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read registry: %v", err)
	}
	v5 = bytes.Replace(v5, []byte(`"version": "s8-fieldspec-v6"`), []byte(`"version": "s7a-fieldspec-v5"`), 1)
	v6, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v6 registry: %v", err)
	}
	if err := config.ValidateMemberTransition("fieldspec", v5, v6); err != nil {
		t.Fatalf("lawful v5-to-v6 transition: %v", err)
	}
}

func TestS8V5ToV6TransitionMakesCatalogAndAdoptionTokensLive(t *testing.T) {
	h := newS4ShimHarness(t)
	credential, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	h.start(t)
	operator := h.dial(t, credential)

	v6, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v6 registry: %v", err)
	}
	h.submit(t, operator, s8ConfigChangeRecord(t, h.root, "fieldspec", v6))
	_ = operator.Close()
	h.stop(t)
	h.start(t)
	operator = h.dial(t, credential)
	defer func() { _ = operator.Close() }()

	describe, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools after v6: %v", err)
	}
	for _, token := range []string{"catalog", "adoption"} {
		if !describe.SubmitSchema.OptionAllowed("member", token) {
			t.Fatalf("post-transition form omitted %q", token)
		}
	}

	catalog, err := os.ReadFile(filepath.Join(h.root, "config", "catalog", "catalog.json"))
	if err != nil {
		t.Fatalf("read catalog: %v", err)
	}
	h.submit(t, operator, s8ConfigChangeRecord(t, h.root, "catalog", catalog))

	adoption := s8ConfigChangeRecord(t, h.root, "adoption", []byte(`{"members":[]}`))
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: adoption, FormDigest: describe.FormDigest})
	if err != nil {
		t.Fatalf("marshal adoption: %v", err)
	}
	result, err := operator.Call(h.ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit adoption: %v", err)
	}
	var outcome struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode adoption outcome: %v", err)
	}
	if outcome.State != record.Rejected {
		t.Fatalf("live adoption state = %q, want rejected for offline bless", outcome.State)
	}
}

func s8ConfigChangeRecord(t *testing.T, root, member string, body []byte) record.Record {
	t.Helper()
	digest := "offline-bless-only"
	if member != "adoption" {
		digest = fixtureDigestWithMember(t, root, member, body)
	}
	return record.Record{Headers: map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium",
		"EVIDENCE_TARGET": "E1", "SUBJECT": "s8 config transition", "record_kind": "config_change",
		"member": member, "new_digest": digest,
	}, Body: string(body)}
}
