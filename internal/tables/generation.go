package tables

import (
	"encoding/json"
	"sort"

	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/seat"
)

const GenesisAuthGeneration = "genesis"

const (
	LifecycleMinted = "minted"
	LifecycleActive = "active"
)

type Lifecycle struct {
	Seat                string
	ActivationState     string
	Role                string
	MintedAt            string
	ActivationRecordRef string
	LastAcceptedAt      string
}

type RosterRow struct {
	Seat                string `json:"seat"`
	ActivationState     string `json:"activation_state"`
	BoundNow            bool   `json:"bound_now"`
	Role                string `json:"role"`
	MintedAt            string `json:"minted_at"`
	ActivationRecordRef string `json:"activation_record_ref"`
	LastAcceptedAt      string `json:"last_accepted_at"`
}

func CurrentAuthGeneration(t *T, seatName string) string {
	if seatName == "" {
		return ""
	}
	if t == nil {
		return GenesisAuthGeneration
	}
	if state, ok := t.Lifecycle[seatName]; ok && state.MintedAt != "" {
		return state.MintedAt
	}
	return GenesisAuthGeneration
}

func SeatActive(t *T, meta seat.SeatMeta) bool {
	if meta.IsOperator || meta.Name == "operator" || meta.Role == "operator" {
		return true
	}
	if t == nil {
		return true
	}
	state, ok := t.Lifecycle[meta.Name]
	if !ok {
		return true
	}
	return state.ActivationState == LifecycleActive
}

func (t *T) updateLifecycle(rec record.Record) {
	if rec.Envelope.DeliveryState != record.Accepted {
		return
	}
	if rec.Headers["record_kind"] == "seat_mint" {
		mint, ok := parseSeatMint(rec)
		if !ok {
			return
		}
		t.Lifecycle[mint.Seat] = Lifecycle{
			Seat:            mint.Seat,
			ActivationState: LifecycleMinted,
			Role:            mint.Role,
			MintedAt:        rec.Envelope.RelayID,
		}
		return
	}
	if rec.Envelope.From == "" || rec.Envelope.From == "operator" || rec.Envelope.Role == "operator" || rec.Envelope.From == "system" {
		return
	}
	state, ok := t.Lifecycle[rec.Envelope.From]
	if !ok {
		return
	}
	if state.ActivationState != LifecycleActive {
		state.ActivationState = LifecycleActive
		state.ActivationRecordRef = rec.Envelope.RelayID
	}
	state.LastAcceptedAt = rec.Envelope.RelayID
	if state.Role == "" {
		state.Role = rec.Envelope.Role
	}
	t.Lifecycle[rec.Envelope.From] = state
}

type seatMintBody struct {
	Seat       string `json:"seat"`
	Role       string `json:"role"`
	IsOperator bool   `json:"is_operator"`
}

func parseSeatMint(rec record.Record) (seatMintBody, bool) {
	var body seatMintBody
	if json.Unmarshal([]byte(rec.Body), &body) != nil || body.Seat == "" {
		return seatMintBody{}, false
	}
	return body, true
}

func Roster(t *T, seats []seat.SeatMeta, bound map[string]bool) []RosterRow {
	if t == nil {
		t = New()
	}
	seen := map[string]bool{}
	var rows []RosterRow
	for _, meta := range seats {
		if meta.Name == "" || meta.Name == "operator" || meta.Role == "operator" || meta.IsOperator {
			continue
		}
		state, ok := t.Lifecycle[meta.Name]
		if !ok {
			continue
		}
		if state.Role == "" {
			state.Role = meta.Role
		}
		rows = append(rows, RosterRow{
			Seat:                meta.Name,
			ActivationState:     state.ActivationState,
			BoundNow:            bound[meta.Name],
			Role:                state.Role,
			MintedAt:            state.MintedAt,
			ActivationRecordRef: state.ActivationRecordRef,
			LastAcceptedAt:      state.LastAcceptedAt,
		})
		seen[meta.Name] = true
	}
	for seatName, state := range t.Lifecycle {
		if seen[seatName] || seatName == "operator" {
			continue
		}
		rows = append(rows, RosterRow{
			Seat:                seatName,
			ActivationState:     state.ActivationState,
			BoundNow:            bound[seatName],
			Role:                state.Role,
			MintedAt:            state.MintedAt,
			ActivationRecordRef: state.ActivationRecordRef,
			LastAcceptedAt:      state.LastAcceptedAt,
		})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Seat < rows[j].Seat })
	return rows
}
