package fixtures_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/egress"
	"github.com/The-Frank-Organization/frank/internal/engine"
	"github.com/The-Frank-Organization/frank/internal/gate"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestS9Decision5FixtureScopedConductorODBModelNameIsLabeledAndScanned(t *testing.T) {
	odb, err := engine.RenderODB(engine.ODBInput{
		SubjectRef:          "s9-egress-fixture",
		DispatchID:          "s9-build",
		ParentDispatchID:    "s9-plan",
		PlainLanguageChange: "Exercise the dormant egress scanner without an external call.",
		WhyNow:              "Decision 5 requires an executable exit fixture.",
		CompletedProof:      "fixture:s9-egress",
		RecordIntegrity:     "observed",
		TradeoffsRisks:      "The live away bridge remains out of scope.",
		Recommendation:      "Keep the scanner fixture-scoped.",
		Choices:             []engine.ODBChoice{{Value: "keep-local", Label: "Keep local"}},
		ModelName:           "gpt-5.5",
	})
	if err != nil {
		t.Fatalf("RenderODB: %v", err)
	}
	st := s9StoreWithFixtureOutbox(t, odb)

	report, err := egress.Drain(st, egress.DefaultRules(), s9FixtureODBRenderer)
	if err != nil {
		t.Fatalf("Drain: %v", err)
	}
	result := s9OnlyEgressResult(t, report)
	if result.Disposition != egress.DispositionReady {
		t.Fatalf("Disposition = %q, want %q; result=%+v", result.Disposition, egress.DispositionReady, result)
	}
	if len(result.Findings) != 1 || result.Findings[0].String() != "model_name:confidentiality" {
		t.Fatalf("Findings = %+v, want scanned model_name:confidentiality", result.Findings)
	}
}

func TestS9Decision5OutOfFenceModelNameIsRefused(t *testing.T) {
	laneSend := record.Record{
		Envelope: record.Envelope{
			RelayID:       "s9-egress-out-of-fence",
			From:          "s9.implementer",
			To:            "operator",
			Role:          "implementer",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":      "SITREP",
			"SUBJECT":    "out-of-fence model identity",
			"model_name": "gpt-5.5",
		},
	}
	st := s9StoreWithFixtureOutbox(t, laneSend)

	report, err := egress.Drain(st, egress.DefaultRules(), s9FixtureODBRenderer)
	if err != nil {
		t.Fatalf("Drain: %v", err)
	}
	result := s9OnlyEgressResult(t, report)
	if result.Disposition != egress.DispositionBlocked {
		t.Fatalf("Disposition = %q, want %q; result=%+v", result.Disposition, egress.DispositionBlocked, result)
	}
	if len(result.Findings) != 1 || result.Findings[0].String() != "model_name:confidentiality" {
		t.Fatalf("Findings = %+v, want blocked model_name:confidentiality", result.Findings)
	}
}

func s9FixtureODBRenderer(meta gate.OutboxItem, source record.Record) (egress.Rendered, error) {
	rendered, err := egress.DefaultRenderer(meta, source)
	if err != nil {
		return egress.Rendered{}, err
	}
	if source.Envelope.From != "system" ||
		source.Envelope.Role != "system" ||
		source.Envelope.To != "operator" ||
		source.Headers["record_kind"] != "odb" {
		return rendered, nil
	}
	for i := range rendered.Fields {
		if rendered.Fields[i].Name == "model_name" {
			rendered.Fields[i].Origin.ConductorODB = true
		}
	}
	return rendered, nil
}

func s9StoreWithFixtureOutbox(t *testing.T, source record.Record) *store.Store {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	if _, err := st.Commit(source, nil); err != nil {
		t.Fatalf("Commit source: %v", err)
	}
	item := gate.OutboxItem{
		ItemID:          "fixture-" + source.Envelope.RelayID,
		SourceKind:      "fixture-egress",
		SourceRecordRef: source.Envelope.RelayID,
		Seat:            source.Envelope.To,
		SchemaVersion:   1,
	}
	data, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("marshal outbox item: %v", err)
	}
	if err := os.WriteFile(filepath.Join(st.Root, "outbox", item.ItemID+".json"), data, 0o600); err != nil {
		t.Fatalf("write fixture outbox item: %v", err)
	}
	return st
}

func s9OnlyEgressResult(t *testing.T, report egress.Report) egress.Result {
	t.Helper()
	if len(report.Results) != 1 {
		t.Fatalf("results = %+v, want exactly one", report.Results)
	}
	return report.Results[0]
}
