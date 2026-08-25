package derived

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/jackli/frank/internal/record"
)

const HookContractV1 = "1"

type WorkStatus struct {
	Cursor []string
	Status string
}

type workBody struct {
	SourceRelayID  string   `json:"source_relay_id"`
	Hook           string   `json:"hook"`
	State          string   `json:"state"`
	Predecessor    string   `json:"predecessor"`
	Kind           string   `json:"kind"`
	CompletedHooks []string `json:"completed_hooks"`
	Reason         string   `json:"reason"`
	Resolves       string   `json:"resolves"`
	Disposition    string   `json:"disposition"`
	EvidenceRef    string   `json:"evidence_ref"`
}

func Cursor(rec record.Record) []string {
	if rec.Headers["resolves_gate"] != "" {
		return []string{"gate", "approval"}
	}
	if rec.Headers["record_kind"] == "seat_mint" {
		return []string{"mint"}
	}
	return nil
}

func Stamp(rec *record.Record) {
	if rec == nil || rec.Envelope.DeliveryState != record.Accepted {
		return
	}
	if rec.Headers == nil {
		rec.Headers = map[string]string{}
	}
	if len(Cursor(*rec)) != 0 {
		rec.Headers["hook_contract"] = HookContractV1
		return
	}
	delete(rec.Headers, "hook_contract")
}

func CursorAdvanceRecord(sourceRelayID string, completedHooks []string) record.Record {
	body, _ := json.Marshal(workBody{
		SourceRelayID:  sourceRelayID,
		Kind:           "cursor_advance",
		CompletedHooks: append([]string(nil), completedHooks...),
	})
	return record.Record{
		Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "derived-work-transition", "PHASE": "SITREP", "SUBJECT": "derived cursor advance"},
		Body:     string(body),
	}
}

func AttemptRecord(sourceRelayID, hook, predecessor string) record.Record {
	if predecessor == "" {
		predecessor = "none"
	}
	body, _ := json.Marshal(workBody{
		SourceRelayID: sourceRelayID,
		Hook:          hook,
		State:         "running_or_unknown",
		Predecessor:   predecessor,
	})
	return record.Record{
		Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "derived-work-attempt", "PHASE": "SITREP", "SUBJECT": "derived work attempt"},
		Body:     string(body),
	}
}

func ParkRecord(sourceRelayID, reason string) record.Record {
	body, _ := json.Marshal(workBody{SourceRelayID: sourceRelayID, Kind: "parked", Reason: reason})
	return record.Record{
		Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "derived-work-transition", "PHASE": "SITREP", "SUBJECT": "derived work parked"},
		Body:     string(body),
	}
}

func RealizedUndeliveredRecord(sourceRelayID string) record.Record {
	body, _ := json.Marshal(workBody{SourceRelayID: sourceRelayID, Kind: "realized-undelivered"})
	return record.Record{
		Envelope: record.Envelope{From: "system", Role: "system", DeliveryState: record.Accepted, SchemaVersion: 1},
		Headers:  map[string]string{"record_kind": "derived-work-transition", "PHASE": "SITREP", "SUBJECT": "mint realized without delivery"},
		Body:     string(body),
	}
}

// AttemptState returns the unresolved marker, if any, and the predecessor for
// the next legal fresh marker. The predecessor is the terminal resolution of
// the prior instance, or "none" for the first instance.
func AttemptState(records []record.Record, sourceRelayID, hook string) (open bool, predecessor string, valid bool) {
	type marker struct {
		id   string
		body workBody
	}
	markers := map[string]marker{}
	resolutions := map[string][]record.Record{}
	resolutionBodies := map[string]workBody{}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		var body workBody
		_ = json.Unmarshal([]byte(rec.Body), &body)
		switch rec.Headers["record_kind"] {
		case "derived-work-attempt":
			if body.SourceRelayID == sourceRelayID && body.Hook == hook {
				markers[rec.Envelope.RelayID] = marker{id: rec.Envelope.RelayID, body: body}
			}
		case "attempt_resolution":
			resolutions[body.Resolves] = append(resolutions[body.Resolves], rec)
			resolutionBodies[rec.Envelope.RelayID] = body
		}
	}
	resolvedIDs := map[string]bool{}
	unresolved := 0
	for markerID := range markers {
		switch len(resolutions[markerID]) {
		case 0:
			unresolved++
		case 1:
			resolvedIDs[resolutions[markerID][0].Envelope.RelayID] = true
		default:
			return true, "", false
		}
	}
	if unresolved > 1 {
		return true, "", false
	}
	roots := 0
	usedPredecessors := map[string]bool{}
	for _, item := range markers {
		if item.body.Predecessor == "" || item.body.Predecessor == "none" {
			roots++
			continue
		}
		prior, ok := resolutionBodies[item.body.Predecessor]
		if !ok || prior.Disposition != "effect-confirmed-unrealized" || usedPredecessors[item.body.Predecessor] {
			return unresolved == 1, "", false
		}
		priorMarker := markers[prior.Resolves]
		if priorMarker.id == "" || priorMarker.body.SourceRelayID != sourceRelayID || priorMarker.body.Hook != hook {
			return unresolved == 1, "", false
		}
		usedPredecessors[item.body.Predecessor] = true
		delete(resolvedIDs, item.body.Predecessor)
	}
	if len(markers) > 0 && roots != 1 {
		return unresolved == 1, "", false
	}
	if unresolved == 1 {
		return true, "", true
	}
	if len(resolvedIDs) == 1 {
		for resolutionID := range resolvedIDs {
			return false, resolutionID, true
		}
	}
	if len(markers) == 0 {
		return false, "none", true
	}
	return false, "", false
}

func ResumableAttempt(records []record.Record, sourceRelayID, hook string) bool {
	open, _, valid := AttemptState(records, sourceRelayID, hook)
	return open && valid
}

func RetryEpoch(records []record.Record, sourceRelayID string) string {
	parks := map[string]bool{}
	var reopenIDs []string
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		var body workBody
		_ = json.Unmarshal([]byte(rec.Body), &body)
		if rec.Headers["record_kind"] == "derived-work-transition" && body.SourceRelayID == sourceRelayID && body.Kind == "parked" {
			parks[rec.Envelope.RelayID] = true
		}
	}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "attempt_resolution" {
			continue
		}
		var body workBody
		_ = json.Unmarshal([]byte(rec.Body), &body)
		if parks[body.Resolves] {
			reopenIDs = append(reopenIDs, rec.Envelope.RelayID)
		}
	}
	sort.Strings(reopenIDs)
	return strings.Join(reopenIDs, ",")
}

func Fold(records []record.Record) map[string]WorkStatus {
	type aggregate struct {
		cursor    []string
		completed map[string]bool
		future    bool
		failed    bool
		conflict  bool
	}
	work := map[string]*aggregate{}
	bodyFor := func(rec record.Record) workBody {
		var body workBody
		_ = json.Unmarshal([]byte(rec.Body), &body)
		return body
	}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		contract, present := rec.Headers["hook_contract"]
		if !present {
			continue
		}
		cursor := Cursor(rec)
		if contract != "1" {
			work[rec.Envelope.RelayID] = &aggregate{cursor: cursor, completed: map[string]bool{}, future: true}
			continue
		}
		if len(cursor) != 0 {
			work[rec.Envelope.RelayID] = &aggregate{cursor: cursor, completed: map[string]bool{}}
		}
	}
	attempts := map[string]workBody{}
	parks := map[string]workBody{}
	resolutions := map[string][]workBody{}
	for _, rec := range records {
		if rec.Envelope.DeliveryState != record.Accepted {
			continue
		}
		kind := rec.Headers["record_kind"]
		body := bodyFor(rec)
		switch kind {
		case "derived-work-attempt":
			if body.State == "running_or_unknown" {
				attempts[rec.Envelope.RelayID] = body
			}
		case "attempt_resolution":
			resolutions[body.Resolves] = append(resolutions[body.Resolves], body)
		case "derived-work-transition":
			target := work[body.SourceRelayID]
			if target == nil {
				continue
			}
			if body.Kind == "cursor_advance" {
				for _, hook := range body.CompletedHooks {
					target.completed[hook] = true
				}
			} else if body.Kind == "parked" {
				parks[rec.Envelope.RelayID] = body
			} else if body.Kind == "realized-undelivered" {
				target.failed = true
			}
		}
	}
	unknownHooks := map[string]map[string]int{}
	attemptPairs := map[string]map[string]bool{}
	for markerID, body := range attempts {
		target := work[body.SourceRelayID]
		if target == nil || target.completed[body.Hook] {
			continue
		}
		if attemptPairs[body.SourceRelayID] == nil {
			attemptPairs[body.SourceRelayID] = map[string]bool{}
		}
		attemptPairs[body.SourceRelayID][body.Hook] = true
		resolved := resolutions[markerID]
		switch len(resolved) {
		case 0:
			if unknownHooks[body.SourceRelayID] == nil {
				unknownHooks[body.SourceRelayID] = map[string]int{}
			}
			unknownHooks[body.SourceRelayID][body.Hook]++
		case 1:
			switch resolved[0].Disposition {
			case "effect-confirmed-unrealized":
			case "effect-confirmed-realized":
				target.failed = true
			default:
				target.conflict = true
			}
		default:
			target.conflict = true
		}
	}
	for sourceRelayID, hooks := range attemptPairs {
		for hook := range hooks {
			_, _, valid := AttemptState(records, sourceRelayID, hook)
			if !valid {
				work[sourceRelayID].conflict = true
			}
		}
	}
	for parkID, body := range parks {
		target := work[body.SourceRelayID]
		if target == nil {
			continue
		}
		switch len(resolutions[parkID]) {
		case 0:
			target.failed = true
		case 1:
			if resolutions[parkID][0].Disposition != "" || resolutions[parkID][0].EvidenceRef != "" {
				target.conflict = true
			}
		default:
			target.conflict = true
		}
	}

	result := make(map[string]WorkStatus, len(work))
	for relayID, item := range work {
		cursor := make([]string, 0, len(item.cursor))
		for _, hook := range item.cursor {
			if !item.completed[hook] {
				cursor = append(cursor, hook)
			}
		}
		unresolvedAttempt := false
		for hook, count := range unknownHooks[relayID] {
			if !item.completed[hook] && count > 0 {
				unresolvedAttempt = true
				if count > 1 {
					item.conflict = true
				}
			}
		}
		status := "pending"
		switch {
		case item.conflict || item.future:
			status = "unknown"
		case item.failed:
			status = "failed"
		case unresolvedAttempt:
			status = "unknown"
		case len(cursor) == 0:
			status = ""
		}
		result[relayID] = WorkStatus{Cursor: cursor, Status: status}
	}
	return result
}

func OpenRelayIDs(work map[string]WorkStatus) []string {
	ids := make([]string, 0, len(work))
	for relayID, status := range work {
		if status.Status == "pending" {
			ids = append(ids, relayID)
		}
	}
	sort.Strings(ids)
	return ids
}
