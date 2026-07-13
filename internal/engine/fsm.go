package engine

import (
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/tables"
)

const (
	GateActive                   = "active"
	GateParkedWaitingHuman       = "parked_waiting_human"
	GateResummonDue              = "resummon_due"
	GateRepliedPendingValidation = "replied_pending_validation"
	GateResumed                  = "resumed"
)

// GateState derives the wake state from durable records. A valid accepted
// resolution dominates the earlier park marker, so recovery never strands a
// lane in parked_waiting_human after its exactly-once wake was committed.
func GateState(tab *tables.T, gateRef string) string {
	if tab == nil || gateRef == "" {
		return GateActive
	}
	state := GateActive
	if tab.ParkedLanes[gateRef] {
		state = GateParkedWaitingHuman
	}
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["resolves_gate"] == gateRef {
			return GateResumed
		}
	}
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["record_kind"] == "resummon_command" && rec.Headers["subject_ref"] == gateRef {
			return GateResummonDue
		}
	}
	return state
}
