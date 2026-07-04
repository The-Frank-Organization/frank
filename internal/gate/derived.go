package gate

import (
	"encoding/json"
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
		if sourceKind == "gate" {
			if err := completePark(st, rec); err != nil {
				return err
			}
		}
		if err := completeOutbox(st, rec, sourceKind); err != nil {
			return err
		}
	}
	return nil
}

func completePark(st *store.Store, gateRecord record.Record) error {
	parkID := "park-" + gateRecord.Envelope.RelayID
	records, err := st.Records()
	if err != nil {
		return err
	}
	for _, existing := range records {
		if existing.Envelope.RelayID == parkID || existing.Headers["parks_gate"] == gateRecord.Envelope.RelayID {
			return nil
		}
	}
	_, err = st.Commit(record.Record{
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

func completeOutbox(st *store.Store, rec record.Record, sourceKind string) error {
	itemID := sourceKind + "-" + rec.Envelope.RelayID
	outboxRecordID := "outbox-" + itemID
	records, err := st.Records()
	if err != nil {
		return err
	}
	for _, existing := range records {
		if existing.Envelope.RelayID == outboxRecordID {
			return nil
		}
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
			RelayID:       outboxRecordID,
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
