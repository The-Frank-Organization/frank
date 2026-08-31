package engine

import (
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
	"github.com/The-Frank-Organization/frank/internal/store"
	"github.com/The-Frank-Organization/frank/internal/tables"
)

func stampParent(st *store.Store, tab *tables.T, cand record.Record, meta seat.SeatMeta, env fieldspec.RenderEnv) record.Record {
	if cand.Headers == nil {
		cand.Headers = map[string]string{}
	}
	delete(cand.Headers, "PARENT_DISPATCH_ID")
	delete(cand.Headers, "parent_hint_honored")
	delete(cand.Headers, "parent_provenance")
	delete(cand.Headers, "routing_ref_honored")

	if hint := cand.Headers["parent_hint"]; hint != "" {
		if parent, ok := provableParentHint(st, tab, meta, hint); ok {
			cand.Headers["PARENT_DISPATCH_ID"] = parent
			cand.Headers["parent_hint_honored"] = "yes"
			cand.Headers["parent_provenance"] = "hint"
			return cand
		}
		cand.Headers["parent_hint_honored"] = "no"
	}

	if parent, provenance, ok := computedParent(tab, cand, env.Turn); ok {
		cand.Headers["PARENT_DISPATCH_ID"] = parent
		cand.Headers["parent_provenance"] = provenance
	}
	return cand
}

func computedParent(tab *tables.T, cand record.Record, turn fieldspec.TurnContext) (string, string, bool) {
	if parent, ok := acceptedRelayRef(tab, turn.WokenOn); ok {
		return parent, "woken_on", true
	}
	if parent, ok := acceptedDispatchRoot(tab, turn.ActiveDispatch); ok {
		return parent, "active_dispatch", true
	}
	if parent, ok := acceptedDispatchRoot(tab, cand.Envelope.DispatchID); ok {
		return parent, "dispatch_root", true
	}
	return "", "", false
}

func provableParentHint(st *store.Store, tab *tables.T, meta seat.SeatMeta, hint string) (string, bool) {
	parent, ok := resolveParentRef(tab, hint)
	if !ok {
		return "", false
	}
	rec, ok := acceptedRecord(tab, parent)
	if !ok {
		return "", false
	}
	if operatorSeat(meta) || rec.Envelope.From == meta.Name || deliveredToSeat(st, tab, meta.Name, parent) {
		return parent, true
	}
	return "", false
}

func resolveParentRef(tab *tables.T, ref string) (string, bool) {
	if parent, ok := acceptedRelayRef(tab, ref); ok {
		return parent, true
	}
	return acceptedDispatchRoot(tab, ref)
}

func acceptedRelayRef(tab *tables.T, ref string) (string, bool) {
	rec, ok := acceptedRecord(tab, ref)
	if !ok {
		return "", false
	}
	return rec.Envelope.RelayID, true
}

func acceptedDispatchRoot(tab *tables.T, dispatchID string) (string, bool) {
	if tab == nil || dispatchID == "" {
		return "", false
	}
	for _, rec := range tab.ByDispatch[dispatchID] {
		if rec.Envelope.DeliveryState == record.Accepted && rec.Envelope.RelayID != "" {
			return rec.Envelope.RelayID, true
		}
	}
	return "", false
}

func acceptedRecord(tab *tables.T, relayID string) (record.Record, bool) {
	if tab == nil || relayID == "" {
		return record.Record{}, false
	}
	rec, ok := tab.ByRelay[relayID]
	if !ok || rec.Envelope.DeliveryState != record.Accepted {
		return record.Record{}, false
	}
	return rec, true
}

func deliveredToSeat(st *store.Store, tab *tables.T, seatName, relayID string) bool {
	if st == nil || seatName == "" || relayID == "" {
		return false
	}
	relays, err := st.Project(seatName)
	if err != nil {
		return false
	}
	for _, delivered := range relays {
		if delivered != relayID {
			continue
		}
		_, ok := acceptedRecord(tab, relayID)
		return ok
	}
	return false
}

func operatorSeat(meta seat.SeatMeta) bool {
	return meta.IsOperator || meta.Name == "operator" || meta.Role == "operator"
}
