package obligation

import (
	"encoding/json"
	"time"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
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

type Tables = tables.T

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
	records := append([]record.Record(nil), t.Records...)
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
			if err := completePark(st, t, rec); err != nil {
				return err
			}
		}
		if err := completeOutbox(st, t, rec, sourceKind); err != nil {
			return err
		}
	}
	return st.CompleteIncidents()
}

func firstTable(existing []*tables.T) *tables.T {
	if len(existing) == 0 {
		return nil
	}
	return existing[0]
}

func completePark(st *store.Store, t *tables.T, gateRecord record.Record) error {
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
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":      "SITREP",
			"SUBJECT":    "parked gate",
			"parks_gate": gateRecord.Envelope.RelayID,
		},
	}
	_, err := st.Commit(rec, nil)
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

func isGateRecord(rec record.Record) bool {
	return rec.Headers["HUMAN_GATE_REQUIRED"] == "yes" || rec.Headers["gate_category"] != ""
}
