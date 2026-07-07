package tables

import (
	"encoding/json"

	"github.com/jackli/frank/internal/record"
)

const GenesisAuthGeneration = "genesis"

func CurrentAuthGeneration(t *T, seatName string) string {
	if seatName == "" {
		return ""
	}
	current := GenesisAuthGeneration
	if t == nil {
		return current
	}
	for _, rec := range t.Records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" {
			continue
		}
		var body struct {
			Seat string `json:"seat"`
		}
		if json.Unmarshal([]byte(rec.Body), &body) != nil || body.Seat != seatName {
			continue
		}
		if rec.Envelope.RelayID != "" {
			current = rec.Envelope.RelayID
		}
	}
	return current
}
