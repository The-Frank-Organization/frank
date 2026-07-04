package obligation

import (
	"encoding/json"
	"time"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

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

type Tables struct {
	OutcomeByIntake map[string]record.Record
	ContentHash     map[string]string
	CompletionIndex map[string]map[string]bool
	ParkedLanes     map[string]bool
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
	records, err := st.Records()
	if err != nil {
		return nil, err
	}
	tables := &Tables{
		OutcomeByIntake: map[string]record.Record{},
		ContentHash:     map[string]string{},
		CompletionIndex: map[string]map[string]bool{
			"park":       {},
			"outbox":     {},
			"incident":   {},
			"quarantine": {},
			"owed":       {},
		},
		ParkedLanes: map[string]bool{},
	}
	for _, rec := range records {
		if rec.Envelope.IntakeID != "" {
			tables.OutcomeByIntake[rec.Envelope.IntakeID] = rec
		}
		if gate := rec.Headers["parks_gate"]; gate != "" {
			tables.CompletionIndex["park"][gate] = true
			tables.ParkedLanes[gate] = true
		}
		if rec.Headers["quarantined_ref"] != "" {
			tables.CompletionIndex["incident"][rec.Headers["quarantined_ref"]] = true
		}
		if rec.Headers["disposes_owed"] != "" {
			tables.CompletionIndex["owed"][rec.Headers["disposes_owed"]] = true
		}
		var outbox struct {
			SourceRecordRef string `json:"source_record_ref"`
		}
		if rec.Body != "" && json.Unmarshal([]byte(rec.Body), &outbox) == nil && outbox.SourceRecordRef != "" {
			tables.CompletionIndex["outbox"][outbox.SourceRecordRef] = true
		}
	}
	return tables, nil
}

func CompleteAuto(st *store.Store) error {
	records, err := st.Records()
	if err != nil {
		return err
	}
	for _, rec := range records {
		sourceKind := ""
		switch {
		case rec.Envelope.DeliveryState == record.Accepted && isGateRecord(rec):
			sourceKind = "gate"
		case rec.Envelope.DeliveryState == record.Held:
			sourceKind = "held"
		default:
			continue
		}
		if sourceKind == "gate" {
			if err := completePark(st, records, rec); err != nil {
				return err
			}
		}
		if err := completeOutbox(st, records, rec, sourceKind); err != nil {
			return err
		}
	}
	return st.CompleteIncidents()
}

func completePark(st *store.Store, records []record.Record, gateRecord record.Record) error {
	parkID := "park-" + gateRecord.Envelope.RelayID
	for _, existing := range records {
		if existing.Envelope.RelayID == parkID || existing.Headers["parks_gate"] == gateRecord.Envelope.RelayID {
			return nil
		}
	}
	_, err := st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       parkID,
			From:          "system",
			Role:          "system",
			To:            gateRecord.Envelope.From,
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":      "SITREP",
			"SUBJECT":    "parked gate",
			"parks_gate": gateRecord.Envelope.RelayID,
		},
	}, nil)
	return err
}

func completeOutbox(st *store.Store, records []record.Record, rec record.Record, sourceKind string) error {
	itemID := sourceKind + "-" + rec.Envelope.RelayID
	outboxRecordID := "outbox-" + itemID
	for _, existing := range records {
		if existing.Envelope.RelayID == outboxRecordID {
			return nil
		}
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
	_, err = st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       outboxRecordID,
			From:          "system",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "derived outbox item"},
		Body:    string(data),
	}, []store.Intent{{Kind: store.IntentOutbox, Path: itemID + ".json", Payload: data}})
	return err
}

func isGateRecord(rec record.Record) bool {
	return rec.Headers["HUMAN_GATE_REQUIRED"] == "yes" || rec.Headers["gate_category"] != ""
}
