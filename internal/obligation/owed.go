package obligation

import (
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func OpenOwed(st *store.Store) ([]Fact, error) {
	records, err := st.Records()
	if err != nil {
		return nil, err
	}
	disposed := map[string]bool{}
	for _, rec := range records {
		if rec.Headers["record_kind"] == "owed_disposition" && rec.Headers["disposes_owed"] != "" {
			disposed[rec.Headers["disposes_owed"]] = true
		}
	}
	var open []Fact
	for _, rec := range records {
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
