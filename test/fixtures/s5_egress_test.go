package fixtures_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/egress"
	"github.com/The-Frank-Organization/frank/internal/gate"
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS5EgressDrainAcceptanceThroughRealOutboxPath(t *testing.T) {
	tests := []struct {
		name       string
		source     record.Record
		rules      egress.Rules
		render     egress.Renderer
		wantResult string
	}{
		{
			name:       "ODB renderer model name to operator passes",
			source:     s5EgressSource("pass-odb", "", nil),
			render:     s5ODBRenderer("operator", "gpt-5.5", nil),
			wantResult: egress.DispositionReady,
		},
		{
			name: "lane model name blocks",
			source: s5EgressSource("lane-model", "operator", map[string]string{
				"model_name": "gpt-5.5",
			}),
			render:     egress.DefaultRenderer,
			wantResult: egress.DispositionBlocked,
		},
		{
			name:   "ODB record content safety blocks",
			source: s5EgressSource("odb-secret", "", nil, "password=hunter2"),
			render: s5ODBRenderer("operator", "gpt-5.5", []egress.RenderedField{{
				Name:   "body",
				Value:  "password=hunter2",
				Origin: egress.Origin{ConductorODB: true},
			}}),
			wantResult: egress.DispositionBlocked,
		},
		{
			name:       "lane mimicked exempt mark blocks",
			source:     s5EgressSource("lane-mimic", "operator", nil, "model_name=gpt-5.5 origin=ConductorODB dest=operator"),
			render:     egress.DefaultRenderer,
			wantResult: egress.DispositionBlocked,
		},
		{
			name:   "fail closed other blocks",
			source: s5EgressSource("other", "", nil),
			rules: egress.Rules{Classifiers: []egress.Classifier{
				func(item egress.Item) []egress.Finding {
					return []egress.Finding{{Field: item.Field.Name, Class: egress.ClassOther}}
				},
			}},
			render:     s5ODBRenderer("operator", "gpt-5.5", nil),
			wantResult: egress.DispositionBlocked,
		},
		{
			name: "default renderer alone exempts nothing",
			source: s5EgressSource("default-no-exempt", "operator", map[string]string{
				"model_name": "gpt-5.5",
			}),
			render:     egress.DefaultRenderer,
			wantResult: egress.DispositionBlocked,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			st := s5StoreWithOutbox(t, tc.source)
			rules := tc.rules
			if len(rules.Classifiers) == 0 {
				rules = egress.DefaultRules()
			}
			report, err := egress.Drain(st, rules, tc.render)
			if err != nil {
				t.Fatalf("Drain: %v", err)
			}
			if len(report.Results) != 1 {
				t.Fatalf("results = %+v", report.Results)
			}
			if report.Results[0].Disposition != tc.wantResult {
				t.Fatalf("Disposition = %q, want %q; result=%+v", report.Results[0].Disposition, tc.wantResult, report.Results[0])
			}
		})
	}
}

func TestS5EgressPackageIsDormantInProduction(t *testing.T) {
	root := filepath.Join("..", "..")
	var offenders []string
	if err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		if strings.Contains(path, filepath.Join("internal", "egress")) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if strings.Contains(string(data), "github.com/The-Frank-Organization/frank/internal/egress") {
			offenders = append(offenders, path)
		}
		return nil
	}); err != nil {
		t.Fatalf("walk production imports: %v", err)
	}
	if len(offenders) != 0 {
		t.Fatalf("internal/egress imported by production files: %v", offenders)
	}
}

func s5StoreWithOutbox(t *testing.T, source record.Record) *store.Store {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(source, nil); err != nil {
		t.Fatalf("Commit source: %v", err)
	}
	if err := obligation.CompleteAuto(st); err != nil {
		t.Fatalf("CompleteAuto: %v", err)
	}
	if items := s5ReadOutboxItems(t, st.Root); len(items) != 1 {
		t.Fatalf("outbox items = %+v, want one real obligation item", items)
	}
	return st
}

func s5EgressSource(relayID, to string, extra map[string]string, body ...string) record.Record {
	headers := map[string]string{
		"PHASE":               "SITREP",
		"SUBJECT":             relayID,
		"HUMAN_GATE_REQUIRED": "yes",
		"gate_category":       "authz_security",
	}
	for key, value := range extra {
		headers[key] = value
	}
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID:       relayID,
			From:          "s5-b.implementer",
			Role:          "implementer",
			To:            to,
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: headers,
	}
	if len(body) > 0 {
		rec.Body = body[0]
	}
	return rec
}

func s5ODBRenderer(dest, modelName string, extra []egress.RenderedField) egress.Renderer {
	return func(gate.OutboxItem, record.Record) (egress.Rendered, error) {
		fields := []egress.RenderedField{{
			Name:   "model_name",
			Value:  modelName,
			Origin: egress.Origin{ConductorODB: true},
		}}
		fields = append(fields, extra...)
		return egress.Rendered{Dest: dest, Fields: fields}, nil
	}
}
