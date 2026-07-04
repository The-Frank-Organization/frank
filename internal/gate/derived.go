package gate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

type OutboxItem struct {
	ItemID          string `json:"item_id"`
	SourceKind      string `json:"source_kind"`
	SourceRecordRef string `json:"source_record_ref"`
	Seat            string `json:"seat"`
	GateCategory    string `json:"gate_category,omitempty"`
	CreatedTS       string `json:"created_ts"`
	SchemaVersion   int    `json:"schema_version"`
}

var resolveMu sync.Mutex

func Complete(st *store.Store) error {
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
		if err := completeOutbox(st, rec, sourceKind); err != nil {
			return err
		}
	}
	return nil
}

func Resolve(st *store.Store, gateRecordRef, verdictID string) (string, error) {
	resolveMu.Lock()
	defer resolveMu.Unlock()

	records, err := st.Records()
	if err != nil {
		return "", err
	}
	state := record.Accepted
	for _, rec := range records {
		if rec.Headers["resolves_gate"] == gateRecordRef && rec.Envelope.DeliveryState == record.Accepted {
			state = record.Rejected
			break
		}
	}
	_, err = st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "verdict-" + verdictID,
			From:          "operator",
			Role:          "operator",
			DeliveryState: state,
			SchemaVersion: 1,
		},
		Headers: map[string]string{
			"PHASE":         "SITREP",
			"SUBJECT":       "operator verdict",
			"resolves_gate": gateRecordRef,
		},
	}, nil)
	if err != nil {
		return "", err
	}
	return state, nil
}

func completeOutbox(st *store.Store, rec record.Record, sourceKind string) error {
	itemID := sourceKind + "-" + rec.Envelope.RelayID
	if _, err := os.Stat(filepath.Join(st.Root, "outbox", itemID+".json")); err == nil {
		return nil
	}
	item := OutboxItem{
		ItemID:          itemID,
		SourceKind:      sourceKind,
		SourceRecordRef: rec.Envelope.RelayID,
		Seat:            rec.Envelope.From,
		GateCategory:    rec.Headers["gate_category"],
		CreatedTS:       time.Now().UTC().Format(time.RFC3339Nano),
		SchemaVersion:   1,
	}
	if sourceKind == "held" && item.GateCategory == "" {
		item.GateCategory = ""
	}
	data, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}
	_, err = st.Commit(record.Record{
		Envelope: record.Envelope{
			RelayID:       "outbox-" + itemID,
			From:          "system",
			Role:          "system",
			DeliveryState: record.Accepted,
			SchemaVersion: 1,
		},
		Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "derived outbox item"},
	}, []store.Intent{{Kind: store.IntentOutbox, Path: itemID + ".json", Payload: data}})
	return err
}

func isGateRecord(rec record.Record) bool {
	return rec.Headers["HUMAN_GATE_REQUIRED"] == "yes" || rec.Headers["gate_category"] != ""
}
