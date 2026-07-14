package engine

import (
	"encoding/json"
	"maps"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/obligation"
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
	choices := make([]obligation.ODBChoice, 0, len(input.Choices))
	for _, choice := range input.Choices {
		choices = append(choices, obligation.ODBChoice{Value: choice.Value, Label: choice.Label})
	}
	return obligation.RenderODB(obligation.ODBRenderInput{
		SubjectRef: input.SubjectRef, DispatchID: input.DispatchID, ParentDispatchID: input.ParentDispatchID,
		PlainLanguageChange: input.PlainLanguageChange, WhyNow: input.WhyNow,
		CompletedProof: input.CompletedProof, RecordIntegrity: input.RecordIntegrity,
		TradeoffsRisks: input.TradeoffsRisks, Recommendation: input.Recommendation,
		Choices: choices, ModelName: input.ModelName, IncludeDispatchField: true,
	})
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
