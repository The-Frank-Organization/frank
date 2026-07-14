package engine

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/record"
)

type ODBChoice struct {
	Value string
	Label string
}

type ODBReply struct {
	Choice string `json:"choice"`
	Note   string `json:"note,omitempty"`
}

type ODBInput struct {
	SubjectRef          string
	DispatchID          string
	ParentDispatchID    string
	PlainLanguageChange string
	WhyNow              string
	CompletedProof      string
	RecordIntegrity     string
	TradeoffsRisks      string
	Recommendation      string
	Choices             []ODBChoice
	ModelName           string
}

func RenderODB(input ODBInput) (record.Record, error) {
	required := map[string]string{
		"subject_ref":           input.SubjectRef,
		"plain_language_change": input.PlainLanguageChange,
		"why_now":               input.WhyNow,
		"completed_proof":       input.CompletedProof,
		"record_integrity":      input.RecordIntegrity,
		"tradeoffs_risks":       input.TradeoffsRisks,
		"recommendation":        input.Recommendation,
	}
	for field, value := range required {
		if strings.TrimSpace(value) == "" {
			return record.Record{}, fmt.Errorf("ODB %s required", field)
		}
	}
	if input.RecordIntegrity != "observed" && input.RecordIntegrity != "self_reported" && input.RecordIntegrity != "mixed" {
		return record.Record{}, fmt.Errorf("ODB record_integrity invalid")
	}
	rows := make([]map[string]string, 0, len(input.Choices))
	seen := make(map[string]bool, len(input.Choices))
	for _, choice := range input.Choices {
		if strings.TrimSpace(choice.Value) == "" || strings.TrimSpace(choice.Label) == "" {
			return record.Record{}, fmt.Errorf("ODB choice value and label required")
		}
		if seen[choice.Value] {
			return record.Record{}, fmt.Errorf("ODB choice %q duplicated", choice.Value)
		}
		seen[choice.Value] = true
		rows = append(rows, map[string]string{"label": choice.Label, "value": choice.Value})
	}
	if len(rows) == 0 {
		return record.Record{}, fmt.Errorf("ODB choices required")
	}
	choices, err := fieldspec.CanonicalMarshal(rows)
	if err != nil {
		return record.Record{}, fmt.Errorf("ODB choices: %w", err)
	}
	to, err := fieldspec.EncodeAddressList([]string{"operator"})
	if err != nil {
		return record.Record{}, fmt.Errorf("ODB TO: %w", err)
	}
	headers := map[string]string{
		"SUBJECT":               "Owner Decision Brief: " + input.SubjectRef,
		"TO":                    to,
		"record_kind":           "odb",
		"subject_ref":           input.SubjectRef,
		"plain_language_change": input.PlainLanguageChange,
		"why_now":               input.WhyNow,
		"completed_proof":       input.CompletedProof,
		"record_integrity":      input.RecordIntegrity,
		"tradeoffs_risks":       input.TradeoffsRisks,
		"recommendation":        input.Recommendation,
		"choices":               choices,
	}
	if input.DispatchID != "" {
		headers["DISPATCH_ID"] = input.DispatchID
	}
	if input.ParentDispatchID != "" {
		headers["PARENT_DISPATCH_ID"] = input.ParentDispatchID
	}
	if input.ModelName != "" {
		headers["model_name"] = input.ModelName
	}
	return record.Record{
		Envelope: record.Envelope{
			RelayID:       "odb-" + input.SubjectRef,
			DispatchID:    input.DispatchID,
			From:          "system",
			To:            "operator",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: headers,
	}, nil
}

func ValidateODBChoice(odb record.Record, picked string) *fieldspec.Violation {
	projection, violation := decisionProjection(odb)
	if violation != nil {
		return violation
	}
	if projection[picked] == "" {
		return &fieldspec.Violation{Field: "choices", Class: "enum"}
	}
	return nil
}

func decisionProjection(odb record.Record) (map[string]string, *fieldspec.Violation) {
	if odb.Headers["record_kind"] != "odb" {
		return nil, &fieldspec.Violation{Field: "record_kind", Class: "enum"}
	}
	typed, err := fieldspec.ParseTyped(&fieldspec.FieldSpec{ID: "choices", Type: "row_array"}, odb.Headers["choices"])
	if err != nil {
		return nil, &fieldspec.Violation{Field: "choices", Class: "typed-parse"}
	}
	rows, ok := typed.([]map[string]string)
	if !ok || len(rows) == 0 {
		return nil, &fieldspec.Violation{Field: "choices", Class: "typed-parse"}
	}
	projection := make(map[string]string, len(rows))
	for _, row := range rows {
		if row["value"] == "" || row["label"] == "" || projection[row["value"]] != "" {
			return nil, &fieldspec.Violation{Field: "choices", Class: "typed-parse"}
		}
		projection[row["value"]] = row["label"]
	}
	return projection, nil
}

func guardedMigratedODB(source record.Record, migrations *migrate.Registry) (record.Record, bool) {
	sourceProjection, violation := decisionProjection(source)
	migrated, err := migrate.Apply(migrations, source)
	if err != nil {
		return source, false
	}
	migratedProjection, violation := decisionProjection(migrated)
	if sourceProjection == nil || violation != nil || !maps.Equal(sourceProjection, migratedProjection) {
		return migrated, false
	}
	return migrated, true
}

func ClassifyODBChoice(odb record.Record, picked string) (string, *fieldspec.Violation) {
	if violation := ValidateODBChoice(odb, picked); violation != nil {
		return record.Rejected, violation
	}
	return record.Accepted, nil
}

func ParseODBReply(body string) (ODBReply, *fieldspec.Violation) {
	var reply ODBReply
	if err := json.Unmarshal([]byte(body), &reply); err != nil || strings.TrimSpace(reply.Choice) == "" {
		return ODBReply{}, &fieldspec.Violation{Field: "choices", Class: "typed-parse"}
	}
	return reply, nil
}
