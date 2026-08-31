package gate

import (
	"github.com/The-Frank-Organization/frank/internal/obligation"
	"github.com/The-Frank-Organization/frank/internal/store"
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
	return obligation.CompleteAuto(st)
}
