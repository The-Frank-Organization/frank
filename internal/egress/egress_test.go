package egress_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/egress"
	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestScanClassifiesRenderedFieldValue(t *testing.T) {
	item := egress.Item{
		Dest:  "operator",
		Field: egress.RenderedField{Name: "body", Value: "api_key=secret-value"},
	}

	verdict := egress.Scan(item, egress.DefaultRules())
	if !verdict.Blocked {
		t.Fatalf("verdict not blocked: %+v", verdict)
	}
	if len(verdict.Findings) != 1 || verdict.Findings[0].Field != "body" || verdict.Findings[0].Class != egress.ClassSafety {
		t.Fatalf("findings = %+v, want body:safety", verdict.Findings)
	}
	if got := verdict.Findings[0].String(); got != "body:safety" || strings.Contains(got, "secret-value") {
		t.Fatalf("finding string = %q, want Field:Class only", got)
	}
}

func TestScanAllowsOnlyConductorODBModelNameToOperator(t *testing.T) {
	tests := []struct {
		name    string
		item    egress.Item
		blocked bool
	}{
		{
			name: "ODB model name to operator passes",
			item: egress.Item{
				Dest:  "operator",
				Field: egress.RenderedField{Name: "model_name", Value: "gpt-5.5", Origin: egress.Origin{ConductorODB: true}},
			},
		},
		{
			name: "lane model name blocks",
			item: egress.Item{
				Dest:  "operator",
				Field: egress.RenderedField{Name: "model_name", Value: "gpt-5.5"},
			},
			blocked: true,
		},
		{
			name: "ODB model name to non-operator blocks",
			item: egress.Item{
				Dest:  "s5-b.implementer",
				Field: egress.RenderedField{Name: "model_name", Value: "gpt-5.5", Origin: egress.Origin{ConductorODB: true}},
			},
			blocked: true,
		},
		{
			name: "safety always blocks even from ODB",
			item: egress.Item{
				Dest:  "operator",
				Field: egress.RenderedField{Name: "body", Value: "password=hunter2", Origin: egress.Origin{ConductorODB: true}},
			},
			blocked: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			verdict := egress.Scan(tc.item, egress.DefaultRules())
			if verdict.Blocked != tc.blocked {
				t.Fatalf("blocked = %v, want %v; findings=%+v", verdict.Blocked, tc.blocked, verdict.Findings)
			}
		})
	}
}

func TestScanFailClosedOtherFindingBlocks(t *testing.T) {
	item := egress.Item{Field: egress.RenderedField{Name: "mystery", Value: "ambiguous"}}
	verdict := egress.Scan(item, egress.Rules{Classifiers: []egress.Classifier{
		func(egress.Item) []egress.Finding {
			return []egress.Finding{{Field: "mystery", Class: egress.ClassOther}}
		},
	}})

	if !verdict.Blocked {
		t.Fatalf("other finding did not fail closed: %+v", verdict)
	}
}

func TestDefaultRendererDerivesDestFromSourceThenItemAndKeepsLaneOrigin(t *testing.T) {
	source := record.Record{
		Envelope: record.Envelope{To: "operator"},
		Headers:  map[string]string{"SUBJECT": "hello", "model_name": "gpt-5.5"},
		Body:     "body text",
	}
	rendered, err := egress.DefaultRenderer(gate.OutboxItem{Seat: "fallback-seat"}, source)
	if err != nil {
		t.Fatalf("DefaultRenderer: %v", err)
	}
	if rendered.Dest != "operator" {
		t.Fatalf("Dest = %q, want source envelope To", rendered.Dest)
	}
	assertRenderedField(t, rendered, "SUBJECT", "hello", false)
	assertRenderedField(t, rendered, "model_name", "gpt-5.5", false)
	assertRenderedField(t, rendered, "body", "body text", false)

	source.Envelope.To = ""
	rendered, err = egress.DefaultRenderer(gate.OutboxItem{Seat: "fallback-seat"}, source)
	if err != nil {
		t.Fatalf("DefaultRenderer fallback: %v", err)
	}
	if rendered.Dest != "fallback-seat" {
		t.Fatalf("fallback Dest = %q, want item seat", rendered.Dest)
	}
}

func TestDrainReadsOutboxSourceRendersAndReportsBlocked(t *testing.T) {
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	source := record.Record{
		Envelope: record.Envelope{RelayID: "source-1", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"PHASE": "SITREP", "SUBJECT": "source"},
	}
	if _, err := st.Commit(source, nil); err != nil {
		t.Fatalf("Commit source: %v", err)
	}
	writeOutboxItem(t, st.Root, gate.OutboxItem{
		ItemID:          "gate-source-1",
		SourceKind:      "gate",
		SourceRecordRef: "source-1",
		Seat:            "seat-a",
		SchemaVersion:   1,
	})

	report, err := egress.Drain(st, egress.DefaultRules(), func(meta gate.OutboxItem, source record.Record) (egress.Rendered, error) {
		return egress.Rendered{
			Dest: "operator",
			Fields: []egress.RenderedField{{
				Name:  "body",
				Value: "auth_url=https://user:pass@example.test",
			}},
		}, nil
	})
	if err != nil {
		t.Fatalf("Drain: %v", err)
	}
	if len(report.Results) != 1 {
		t.Fatalf("results = %+v", report.Results)
	}
	result := report.Results[0]
	if result.ItemID != "gate-source-1" || result.SourceRecordRef != "source-1" || result.Disposition != egress.DispositionBlocked {
		t.Fatalf("result = %+v, want blocked source-1", result)
	}
	if len(result.Findings) != 1 || result.Findings[0].String() != "body:safety" {
		t.Fatalf("findings = %+v", result.Findings)
	}
}

func assertRenderedField(t *testing.T, rendered egress.Rendered, name, value string, conductorODB bool) {
	t.Helper()
	for _, field := range rendered.Fields {
		if field.Name == name && field.Value == value && field.Origin.ConductorODB == conductorODB {
			return
		}
	}
	t.Fatalf("missing rendered field %s=%q conductorODB=%v in %+v", name, value, conductorODB, rendered.Fields)
}

func writeOutboxItem(t *testing.T, root string, item gate.OutboxItem) {
	t.Helper()
	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("marshal outbox item: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "outbox", item.ItemID+".json"), data, 0o644); err != nil {
		t.Fatalf("write outbox item: %v", err)
	}
}
