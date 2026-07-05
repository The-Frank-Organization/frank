package obligation

import (
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

func OpenOwed(st *store.Store, existing ...*tables.T) ([]Fact, error) {
	t := firstTable(existing)
	if t == nil {
		var err error
		t, err = tables.Build(st)
		if err != nil {
			return nil, err
		}
	}
	disposed := map[string]bool{}
	for _, rec := range t.Records {
		if rec.Headers["record_kind"] == "owed_disposition" && rec.Headers["disposes_owed"] != "" {
			disposed[rec.Headers["disposes_owed"]] = true
		}
	}
	var open []Fact
	for _, rec := range t.Records {
		if rec.Headers["record_kind"] != "owed_item" || rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		if disposed[rec.Envelope.RelayID] {
			continue
		}
		open = append(open, Fact{Kind: "owed_item", Key: rec.Envelope.RelayID})
	}
	return open, nil
}
