package obligation

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

var ErrGateRegistry = errors.New("gate registry unavailable")

type Fact struct {
	Kind string
	Key  string
}

type Class struct {
	Name      string
	Sources   func(*store.Store) ([]Fact, error)
	Completed func(Fact, *Tables) bool
	Complete  func(Fact, *store.Store) error
}

type Tables = tables.T

type ODBChoice struct {
	Value string
	Label string
}

type ODBRenderInput struct {
	SubjectRef           string
	DispatchID           string
	ParentDispatchID     string
	PlainLanguageChange  string
	WhyNow               string
	CompletedProof       string
	RecordIntegrity      string
	TradeoffsRisks       string
	Recommendation       string
	Choices              []ODBChoice
	ChoiceRows           []map[string]string
	ModelName            string
	Phase                string
	SchemaVersion        int
	IncludeDispatchField bool
}

type SystemOperatorInput struct {
	RelayID       string
	DispatchID    string
	DeliveryState string
	SchemaVersion int
	Headers       map[string]string
	Body          string
}

func SystemOperatorRecord(input SystemOperatorInput) (record.Record, error) {
	to, err := fieldspec.EncodeAddressList([]string{"operator"})
	if err != nil {
		return record.Record{}, err
	}
	headers := cloneStrings(input.Headers)
	if headers == nil {
		headers = map[string]string{}
	}
	headers["TO"] = to
	schemaVersion := input.SchemaVersion
	if schemaVersion == 0 {
		schemaVersion = 1
	}
	return record.Record{
		Envelope: record.Envelope{
			RelayID: input.RelayID, DispatchID: input.DispatchID,
			From: "system", To: "operator", Role: "system",
			DeliveryState: input.DeliveryState, SchemaVersion: schemaVersion,
		},
		Headers: headers,
		Body:    input.Body,
	}, nil
}

func RenderODB(input ODBRenderInput) (record.Record, error) {
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
	rows := input.ChoiceRows
	if len(rows) == 0 {
		rows = make([]map[string]string, 0, len(input.Choices))
		for _, choice := range input.Choices {
			rows = append(rows, map[string]string{"label": choice.Label, "value": choice.Value})
		}
	}
	seen := make(map[string]bool, len(rows))
	for _, row := range rows {
		if strings.TrimSpace(row["value"]) == "" || strings.TrimSpace(row["label"]) == "" {
			return record.Record{}, fmt.Errorf("ODB choice value and label required")
		}
		if seen[row["value"]] {
			return record.Record{}, fmt.Errorf("ODB choice %q duplicated", row["value"])
		}
		seen[row["value"]] = true
	}
	if len(rows) == 0 {
		return record.Record{}, fmt.Errorf("ODB choices required")
	}
	choices, err := fieldspec.CanonicalMarshal(rows)
	if err != nil {
		return record.Record{}, fmt.Errorf("ODB choices: %w", err)
	}
	headers := map[string]string{
		"SUBJECT":               "Owner Decision Brief: " + input.SubjectRef,
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
	if input.Phase != "" {
		headers["PHASE"] = input.Phase
	}
	if input.IncludeDispatchField && input.DispatchID != "" {
		headers["DISPATCH_ID"] = input.DispatchID
	}
	if input.ParentDispatchID != "" {
		headers["PARENT_DISPATCH_ID"] = input.ParentDispatchID
	}
	if input.ModelName != "" {
		headers["model_name"] = input.ModelName
	}
	rec, err := SystemOperatorRecord(SystemOperatorInput{
		RelayID: "odb-" + input.SubjectRef, DispatchID: input.DispatchID,
		DeliveryState: record.Accepted, SchemaVersion: input.SchemaVersion, Headers: headers,
	})
	if err != nil {
		return record.Record{}, fmt.Errorf("ODB TO: %w", err)
	}
	return rec, nil
}

type Engine struct {
	classes map[string]Class
	tables  *Tables
	open    map[string][]Fact
}

func New() *Engine {
	return &Engine{classes: map[string]Class{}, open: map[string][]Fact{}}
}

func (e *Engine) Register(class Class) {
	e.classes[class.Name] = class
}

func (e *Engine) Open(name string) []Fact {
	return append([]Fact(nil), e.open[name]...)
}

func (e *Engine) BuildTables(st *store.Store) error {
	tables, err := BuildTables(st)
	if err != nil {
		return err
	}
	for name := range e.classes {
		if tables.CompletionIndex[name] == nil {
			tables.CompletionIndex[name] = map[string]bool{}
		}
	}
	e.tables = tables
	e.open = map[string][]Fact{}
	for name, class := range e.classes {
		sources, err := class.Sources(st)
		if err != nil {
			return err
		}
		for _, fact := range sources {
			if class.Completed != nil && class.Completed(fact, tables) {
				continue
			}
			e.open[name] = append(e.open[name], fact)
		}
	}
	return nil
}

func (e *Engine) OnCommit(record.Record) {}

func BuildTables(st *store.Store) (*Tables, error) {
	return tables.Build(st)
}

func CompleteAuto(st *store.Store, existing ...*tables.T) error {
	t := firstTable(existing)
	if t == nil {
		var err error
		t, err = tables.Build(st)
		if err != nil {
			return err
		}
	}
	gateRegistry, useDefaultCategories, registryErr := loadGateRegistry(st)
	records := append([]record.Record(nil), t.Records...)
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Rejected || rec.Headers["failing_edge"] != "stale_choice_set" {
			continue
		}
		if err := completeStaleChoiceReissue(st, t, rec); err != nil {
			return err
		}
	}
	for _, rec := range records {
		sourceKind := ""
		switch rec.Envelope.DeliveryState {
		case record.Accepted:
			if !isAGateRecord(gateRegistry, useDefaultCategories, rec) {
				continue
			}
			sourceKind = "gate"
		case record.Held:
			sourceKind = "held"
		default:
			continue
		}
		if sourceKind == "gate" {
			if err := completePark(st, t, rec); err != nil {
				return err
			}
		}
		if err := completeOutbox(st, t, rec, sourceKind); err != nil {
			return err
		}
	}
	if err := st.CompleteIncidents(); err != nil {
		return err
	}
	if registryErr != nil {
		recorded, err := completeGateRegistryFault(st, t)
		if err != nil {
			return err
		}
		if recorded {
			return registryErr
		}
	}
	return nil
}

type staleChoiceIntent struct {
	Reason                string `json:"reason"`
	SourceGate            string `json:"source_gate"`
	ReplacementDecisionID string `json:"replacement_decision_id"`
	Choices               string `json:"choices"`
	TargetSchemaVersion   int    `json:"target_schema_version"`
}

func completeStaleChoiceReissue(st *store.Store, t *tables.T, stale record.Record) error {
	var intent staleChoiceIntent
	if err := json.Unmarshal([]byte(stale.Body), &intent); err != nil {
		return fmt.Errorf("decode stale choice reissue: %w", err)
	}
	if intent.Reason != "stale_choice_set" || intent.SourceGate == "" || intent.ReplacementDecisionID == "" || intent.TargetSchemaVersion < 1 {
		return errors.New("stale choice reissue intent invalid")
	}
	source, ok := t.ByRelay[intent.SourceGate]
	if !ok || source.Envelope.DeliveryState != record.Accepted {
		return errors.New("stale choice source gate unavailable")
	}

	heldID := "held-stale-schema-" + intent.SourceGate
	held, heldExists := t.ByRelay[heldID]
	if !heldExists {
		var err error
		held, err = SystemOperatorRecord(SystemOperatorInput{
			RelayID: heldID, DispatchID: source.Envelope.DispatchID,
			DeliveryState: record.Held, SchemaVersion: intent.TargetSchemaVersion,
			Headers: map[string]string{
				"PHASE": "SITREP", "SUBJECT": "stale schema requires replacement decision",
				"failing_edge": "stale_schema", "subject_ref": intent.SourceGate,
			},
			Body: "stale_schema",
		})
		if err != nil {
			return err
		}
		if _, err := st.Commit(held, nil); err != nil {
			return err
		}
		t.OnCommit(held)
		if err := completeOutbox(st, t, held, "held"); err != nil {
			return err
		}
		crashpoint.Hit("stale_reissue_after_held")
	}
	if _, err := fieldspec.ParseTyped(&fieldspec.FieldSpec{ID: "choices", Type: "row_array"}, intent.Choices); err != nil {
		return fmt.Errorf("stale choice reissue choices: %w", err)
	}

	if _, ok := t.ByRelay[intent.ReplacementDecisionID]; !ok {
		replacement := cloneRecord(source)
		replacement.Checksum = ""
		replacement.Envelope.RelayID = intent.ReplacementDecisionID
		replacement.Envelope.IntakeID = ""
		replacement.Envelope.SchemaVersion = intent.TargetSchemaVersion
		replacement.Headers["choices"] = intent.Choices
		replacement.Headers["reissues_gate"] = intent.SourceGate
		replacement.Headers["SUBJECT"] = source.Headers["SUBJECT"] + " (schema replacement)"
		if _, err := st.Commit(replacement, nil); err != nil {
			return err
		}
		t.OnCommit(replacement)
	}
	replacement := t.ByRelay[intent.ReplacementDecisionID]
	if err := completePark(st, t, replacement); err != nil {
		return err
	}
	return completeOutbox(st, t, replacement, "gate")
}

func cloneRecord(rec record.Record) record.Record {
	rec.Headers = cloneStrings(rec.Headers)
	rec.XFields = cloneStrings(rec.XFields)
	return rec
}

func cloneStrings(values map[string]string) map[string]string {
	if values == nil {
		return nil
	}
	out := make(map[string]string, len(values))
	for key, value := range values {
		out[key] = value
	}
	return out
}

func firstTable(existing []*tables.T) *tables.T {
	if len(existing) == 0 {
		return nil
	}
	return existing[0]
}

func completePark(st *store.Store, t *tables.T, gateRecord record.Record) error {
	if err := completeODB(st, t, gateRecord); err != nil {
		return err
	}
	parkID := "park-" + gateRecord.Envelope.RelayID
	if _, ok := t.ByRelay[parkID]; ok || t.CompletionIndex["park"][gateRecord.Envelope.RelayID] {
		return nil
	}
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID:       parkID,
			From:          "system",
			Role:          "system",
			To:            gateRecord.Envelope.From,
			DeliveryState: record.Accepted,
			SchemaVersion: gateRecord.Envelope.SchemaVersion,
		},
		Headers: map[string]string{
			"PHASE":      "SITREP",
			"SUBJECT":    "parked_waiting_human",
			"parks_gate": gateRecord.Envelope.RelayID,
		},
	}
	_, err := st.Commit(rec, nil)
	if err == nil {
		t.OnCommit(rec)
	}
	return err
}

func completeODB(st *store.Store, t *tables.T, gateRecord record.Record) error {
	odbID := "odb-" + gateRecord.Envelope.RelayID
	if _, ok := t.ByRelay[odbID]; ok {
		return nil
	}
	choiceRows := []map[string]string{
		{"label": "Approve", "value": "approve"},
		{"label": "Reject", "value": "reject"},
	}
	if raw := gateRecord.Headers["choices"]; raw != "" {
		typed, parseErr := fieldspec.ParseTyped(&fieldspec.FieldSpec{ID: "choices", Type: "row_array"}, raw)
		var ok bool
		choiceRows, ok = typed.([]map[string]string)
		if parseErr != nil || !ok || len(choiceRows) == 0 {
			return errors.New("gate choices invalid")
		}
	}
	subject := gateRecord.Headers["SUBJECT"]
	if subject == "" {
		subject = gateRecord.Envelope.RelayID
	}
	completedProof := gateRecord.Headers["evidence_ref"]
	if completedProof == "" {
		completedProof = "accepted:" + gateRecord.Envelope.RelayID
	}
	rec, err := RenderODB(ODBRenderInput{
		SubjectRef: gateRecord.Envelope.RelayID, DispatchID: gateRecord.Envelope.DispatchID,
		PlainLanguageChange: subject,
		WhyNow:              "The accepted A-gate requires an operator verdict before the lane can continue.",
		CompletedProof:      completedProof, RecordIntegrity: "observed",
		TradeoffsRisks: "The lane remains parked until a bounded operator choice is validated.",
		Recommendation: "Choose approve only when the accepted gate may safely resume.",
		ChoiceRows:     choiceRows, Phase: "SITREP", SchemaVersion: gateRecord.Envelope.SchemaVersion,
	})
	if err != nil {
		return err
	}
	_, err = st.Commit(rec, nil)
	if err == nil {
		t.OnCommit(rec)
	}
	return err
}

func completeOutbox(st *store.Store, t *tables.T, rec record.Record, sourceKind string) error {
	itemID := sourceKind + "-" + rec.Envelope.RelayID
	outboxRecordID := "outbox-" + itemID
	if _, ok := t.ByRelay[outboxRecordID]; ok || t.CompletionIndex["outbox"][rec.Envelope.RelayID] {
		return nil
	}
	item := struct {
		ItemID          string `json:"item_id"`
		SourceKind      string `json:"source_kind"`
		SourceRecordRef string `json:"source_record_ref"`
		Seat            string `json:"seat"`
		GateCategory    string `json:"gate_category,omitempty"`
		CreatedTS       string `json:"created_ts"`
		SchemaVersion   int    `json:"schema_version"`
	}{
		ItemID:          itemID,
		SourceKind:      sourceKind,
		SourceRecordRef: rec.Envelope.RelayID,
		Seat:            rec.Envelope.From,
		GateCategory:    rec.Headers["gate_category"],
		CreatedTS:       time.Now().UTC().Format(time.RFC3339Nano),
		SchemaVersion:   1,
	}
	data, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}
	outboxRec := record.Record{
		Envelope: record.Envelope{
			RelayID:       outboxRecordID,
			From:          "system",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "derived outbox item"},
		Body:    string(data),
	}
	_, err = st.Commit(outboxRec, []store.Intent{{Kind: store.IntentOutbox, Path: itemID + ".json", Payload: data}})
	if err == nil {
		t.OnCommit(outboxRec)
	}
	return err
}

func loadGateRegistry(st *store.Store) (*fieldspec.Registry, bool, error) {
	reg, err := fieldspec.Load(filepath.Join(st.Root, "config", "fieldspec", "registry.json"))
	if err == nil {
		return reg, false, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return nil, true, nil
	}
	return nil, false, fmt.Errorf("%w: %v", ErrGateRegistry, err)
}

func completeGateRegistryFault(st *store.Store, t *tables.T) (bool, error) {
	const relayID = "held-gate-registry-unavailable"
	if _, ok := t.ByRelay[relayID]; ok {
		return false, nil
	}
	rec := record.Record{
		Envelope: record.Envelope{
			RelayID: relayID, From: "system", Role: "system",
			DeliveryState: record.Held, SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE": "SITREP", "SUBJECT": "gate registry unavailable", "failure_class": "gate-registry-unavailable",
		},
		Body: "gate_category:registry-unavailable",
	}
	if _, err := st.Commit(rec, nil); err != nil {
		return false, err
	}
	t.OnCommit(rec)
	return true, nil
}

func isAGateRecord(reg *fieldspec.Registry, useDefaultCategories bool, rec record.Record) bool {
	if rec.Headers["record_kind"] == "gate_resolution" {
		return false
	}
	if rec.Headers["HUMAN_GATE_REQUIRED"] == "yes" || rec.Headers["egress_scan_result"] == "blocked" {
		return true
	}
	category := rec.Headers["gate_category"]
	if category == "" {
		return false
	}
	if reg != nil {
		return reg.GateCategoryAuthorityBearing(category)
	}
	if useDefaultCategories {
		return defaultAGateCategories[category]
	}
	return false
}

var defaultAGateCategories = map[string]bool{
	"merge_to_protected":       true,
	"irreversible_write":       true,
	"residual_risk_acceptance": true,
	"live_verify_skip":         true,
	"ceremony_downgrade":       true,
	"authz_security":           true,
	"product_semantics":        true,
	"scope_expansion":          true,
	"routing_escalation":       true,
	"other":                    true,
}
