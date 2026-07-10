package invariants_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func TestLawTerminalEnumByteExact(t *testing.T) {
	_, _ = requireCatalogLaw(t, "TestLawTerminalEnumByteExact")
	reg := loadRegistry(t)
	want := []string{"accepted", "rejected", "held"}
	if got := reg.NamedEnums["delivery_state"]; !reflect.DeepEqual(got, want) {
		t.Fatalf("terminal enum bytes = %q, want literal %q", got, want)
	}

	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("open table store: %v", err)
	}
	for index, state := range want {
		rec := record.Record{
			Envelope: record.Envelope{
				RelayID:       "terminal-" + state,
				From:          "system",
				Role:          "system",
				DeliveryState: state,
				SchemaVersion: 1,
			},
			Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "terminal state"},
		}
		if _, err := st.Commit(rec, []store.Intent{}); err != nil {
			t.Fatalf("commit terminal state %d/%s: %v", index, state, err)
		}
	}
	tab, err := tables.Build(st)
	if err != nil {
		t.Fatalf("build engine tables: %v", err)
	}
	var tableStates []string
	for _, rec := range tab.Records {
		tableStates = append(tableStates, rec.Envelope.DeliveryState)
	}
	sort.Strings(tableStates)
	sortedWant := append([]string(nil), want...)
	sort.Strings(sortedWant)
	if !reflect.DeepEqual(tableStates, sortedWant) {
		t.Fatalf("engine table terminal states = %q, want literal %q", tableStates, sortedWant)
	}

	meta := seat.SeatMeta{Name: "s7.implementer", Role: "implementer"}
	candidate := record.Record{Headers: map[string]string{
		"PHASE":           "SITREP",
		"AUTHORITY":       "report-only",
		"CEREMONY_TIER":   "small",
		"EVIDENCE_TARGET": "E1",
		"SUBJECT":         "forged fourth terminal",
		"delivery_state":  "forwarded",
	}}
	_, digest := reg.Render(
		fieldspec.RenderEnv{},
		fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role},
		candidate.Headers["PHASE"],
		candidate.Headers["CEREMONY_TIER"],
		fieldspec.ClosedGrantState,
	)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: candidate, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal forged submit: %v", err)
	}
	got, _, err := engine.SubmitHandler(st, reg, meta)(context.Background(), intake.Cmd{
		IntakeID: "intake-terminal-forgery",
		Seat:     meta.Name,
		Role:     meta.Role,
		Verb:     "submit",
		Payload:  payload,
	})
	if err != nil {
		t.Fatalf("submit forged fourth terminal: %v", err)
	}
	if got.Envelope.DeliveryState != record.Rejected {
		t.Fatalf("forged fourth terminal state = %q, want rejected", got.Envelope.DeliveryState)
	}
	if !strings.Contains(got.Body, "delivery_state:system-owned") {
		t.Fatalf("forged fourth terminal detail = %q, want delivery_state:system-owned", got.Body)
	}
}

func TestLawThreeVerbSurface(t *testing.T) {
	_, _ = requireCatalogLaw(t, "TestLawThreeVerbSurface")
	want := []string{"submit", "project", "read"}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-s7-verbs-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })

	result := func(value string) channel.ToolFunc {
		return func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(fmt.Sprintf(`{"tool":%q}`, value)), nil
		}
	}
	toolSet := channel.ToolSet{
		Submit:  result("submit"),
		Project: result("project"),
		Read:    result("read"),
		Describe: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.Marshal(channel.DescriptionResponse{
				Tools: want,
				Descriptions: map[string]string{
					"submit": "submit", "project": "project", "read": "read",
				},
			})
		},
	}
	var toolSetFields []string
	for i, typ := 0, reflect.TypeOf(toolSet); i < typ.NumField(); i++ {
		toolSetFields = append(toolSetFields, typ.Field(i).Name)
	}
	if wantFields := []string{"Submit", "Project", "Read", "Describe"}; !reflect.DeepEqual(toolSetFields, wantFields) {
		t.Fatalf("ToolSet fields = %q, want literal census %q", toolSetFields, wantFields)
	}
	server, err := channel.Serve(sock, channel.FullSurface(engine.TestReady(), toolSet))
	if err != nil {
		t.Fatalf("serve three-verb surface: %v", err)
	}
	defer func() { _ = server.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := channel.Dial(ctx, sock)
	if err != nil {
		t.Fatalf("dial three-verb surface: %v", err)
	}

	got, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("list seat-visible tools: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("seat-visible verb bytes = %q, want literal %q", got, want)
	}
	descriptions, err := client.ToolDescriptions(ctx)
	if err != nil {
		t.Fatalf("render tool descriptions: %v", err)
	}
	var described []string
	for name := range descriptions {
		described = append(described, name)
	}
	sort.Strings(described)
	sortedWant := append([]string(nil), want...)
	sort.Strings(sortedWant)
	if !reflect.DeepEqual(described, sortedWant) {
		t.Fatalf("described verb bytes = %q, want literal %q", described, sortedWant)
	}
	for _, name := range want {
		raw, err := client.Call(ctx, name, nil)
		if err != nil {
			t.Fatalf("call %s: %v", name, err)
		}
		if !bytes.Contains(raw, []byte(name)) {
			t.Fatalf("%s result = %s, want named tool result", name, raw)
		}
	}
	for _, name := range []string{"describe", "arbitrary"} {
		if _, err := client.Call(ctx, name, nil); err == nil || !strings.Contains(err.Error(), "unknown tool") {
			t.Fatalf("call %s error = %v, want unknown tool", name, err)
		}
	}
}

func TestLawR2NoModelPredicate(t *testing.T) {
	_, _ = requireCatalogLaw(t, "TestLawR2NoModelPredicate")
	reg := loadRegistry(t)
	var modelIDs []string
	for i := range reg.Fields {
		field := &reg.Fields[i]
		if !field.ModelIdentity {
			continue
		}
		modelIDs = append(modelIDs, field.ID)
		if field.GateReferenceable {
			t.Fatalf("model identity %s is gate-referenceable", field.ID)
		}
	}
	if len(modelIDs) == 0 {
		t.Fatal("live registry has no model-identity field to guard")
	}
	for _, field := range reg.Fields {
		for kind, raw := range map[string]json.RawMessage{
			"required_when": field.RequiredWhen.Raw,
			"visible_when":  field.VisibleWhen.Raw,
		} {
			for _, modelID := range modelIDs {
				if bytes.Contains(raw, []byte(`"`+modelID+`"`)) {
					t.Fatalf("%s.%s references model identity %s: %s", field.ID, kind, modelID, raw)
				}
			}
		}
	}

	for _, predicateKind := range []string{"required_when", "visible_when"} {
		t.Run("synthetic "+predicateKind, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(repoRoot(t), "internal", "fieldspec", "registry.json"))
			if err != nil {
				t.Fatalf("read registry fixture: %v", err)
			}
			var fixture map[string]any
			if err := json.Unmarshal(data, &fixture); err != nil {
				t.Fatalf("decode registry fixture: %v", err)
			}
			fields, ok := fixture["fields"].([]any)
			if !ok || len(fields) == 0 {
				t.Fatalf("registry fields shape = %T", fixture["fields"])
			}
			target, ok := fields[0].(map[string]any)
			if !ok {
				t.Fatalf("registry first field shape = %T", fields[0])
			}
			target[predicateKind] = map[string]any{"field": modelIDs[0], "op": "present"}
			mutated, err := json.Marshal(fixture)
			if err != nil {
				t.Fatalf("marshal model-keyed registry: %v", err)
			}
			path := filepath.Join(t.TempDir(), "registry.json")
			if err := os.WriteFile(path, mutated, 0o644); err != nil {
				t.Fatalf("write model-keyed registry: %v", err)
			}
			_, err = fieldspec.Load(path)
			if err == nil {
				t.Fatalf("model-keyed %s loaded successfully", predicateKind)
			}
			if !strings.Contains(err.Error(), "gate-referenceable") {
				t.Fatalf("model-keyed %s error = %q, want gate-referenceable", predicateKind, err)
			}
		})
	}
}

func loadRegistry(t *testing.T) *fieldspec.Registry {
	t.Helper()
	reg, err := fieldspec.Load(filepath.Join(repoRoot(t), "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("load live registry: %v", err)
	}
	return reg
}
