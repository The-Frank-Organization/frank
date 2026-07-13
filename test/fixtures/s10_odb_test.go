package fixtures_test

import (
	"testing"

	"github.com/jackli/frank/internal/engine"
)

func TestS10ODBRejectsChoiceOutsideFrozenAgentEnumPickSet(t *testing.T) {
	choices := []engine.ODBChoice{
		{Value: "approve", Label: "Approve"},
		{Value: "revise", Label: "Revise"},
	}
	odb, err := engine.RenderODB(engine.ODBInput{
		SubjectRef:          "decision-17",
		DispatchID:          "s10-build",
		ParentDispatchID:    "s10-plan",
		PlainLanguageChange: "Wake the parked lane after a reviewed operator decision.",
		WhyNow:              "The lane cannot proceed without human judgment.",
		CompletedProof:      "relay:proof-17",
		RecordIntegrity:     "observed",
		TradeoffsRisks:      "Proceeding may require a follow-up fold.",
		Recommendation:      "Approve the bounded wake.",
		Choices:             choices,
		ModelName:           "gpt-5.5",
	})
	if err != nil {
		t.Fatalf("RenderODB: %v", err)
	}
	choices[0].Value = "mutated-after-emit"

	state, violation := engine.ClassifyODBChoice(odb, "approve")
	if state != "accepted" || violation != nil {
		t.Fatalf("frozen legal choice = %q / %+v, want accepted", state, violation)
	}
	state, violation = engine.ClassifyODBChoice(odb, "mutated-after-emit")
	if state != "rejected" || violation == nil || violation.Field != "choices" || violation.Class != "enum" {
		t.Fatalf("out-of-set choice = %q / %+v, want rejected choices:enum", state, violation)
	}
	if odb.Envelope.DeliveryState != "accepted" || odb.Headers["record_kind"] != "odb" {
		t.Fatalf("ODB envelope/record kind = %+v / %q", odb.Envelope, odb.Headers["record_kind"])
	}
	for _, field := range []string{
		"plain_language_change", "why_now", "completed_proof", "record_integrity",
		"tradeoffs_risks", "recommendation", "choices",
	} {
		if odb.Headers[field] == "" {
			t.Fatalf("ODB omitted seven-field bundle member %q", field)
		}
	}
	if odb.Headers["model_name"] != "gpt-5.5" {
		t.Fatalf("typed model_name = %q", odb.Headers["model_name"])
	}
}
