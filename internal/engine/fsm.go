package engine

import (
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

const (
	GateActive                   = "active"
	GateParkedWaitingHuman       = "parked_waiting_human"
	GateResummonDue              = "resummon_due"
	GateRepliedPendingValidation = "replied_pending_validation"
	GateResumed                  = "resumed"
	GateBouncedRepair            = "bounced_repair"
	GateEgressBlocked            = "egress_blocked"
)

// GateState derives the wake state from durable records. A valid accepted
// resolution dominates the earlier park marker, so recovery never strands a
// lane in parked_waiting_human after its exactly-once wake was committed.
func GateState(tab *tables.T, gateRef string) string {
	if tab == nil || gateRef == "" {
		return GateActive
	}
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["resolves_gate"] == gateRef {
			return GateResumed
		}
	}
	gate := tab.ByRelay[gateRef]
	if gate.Envelope.DeliveryState == record.Rejected && store.AcceptanceBounceEdge(gate.Headers["failing_edge"]) {
		return GateBouncedRepair
	}
	if gate.Envelope.DeliveryState == record.Accepted && gate.Headers["egress_scan_result"] == "blocked" && gate.Headers["failing_edge"] == "egress" {
		return GateEgressBlocked
	}
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Headers["record_kind"] == "resummon_command" && rec.Headers["subject_ref"] == gateRef {
			return GateResummonDue
		}
	}
	if tab.ParkedLanes[gateRef] {
		return GateParkedWaitingHuman
	}
	return GateActive
}
