// Package contextmgr implements the governed three-tier context assembly.
package contextmgr

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
)

const SummarySentinel = "[GOVERNED_SUMMARY: derived context; Tier-0 constraints remain authoritative]"

var ErrPinnedIntegrity = errors.New("contextmgr: pinned tier integrity failure")

type Item struct {
	Kind    string          `json:"kind"`
	Content json.RawMessage `json:"content"`
}

type Manager struct {
	pinned       []Item
	pinnedDigest string
	tierOne      []Item
	tierTwo      []Item
}

func New(pinned []Item) (*Manager, error) {
	copyPinned := cloneItems(pinned)
	digest, err := digestItems(copyPinned)
	if err != nil {
		return nil, err
	}
	return &Manager{pinned: copyPinned, pinnedDigest: digest}, nil
}

func (manager *Manager) AddEvictable(items ...Item) {
	manager.tierOne = append(manager.tierOne, cloneItems(items)...)
}

func (manager *Manager) EvictOldest(count int) {
	if count <= 0 {
		return
	}
	if count >= len(manager.tierOne) {
		manager.tierOne = nil
		return
	}
	manager.tierOne = append([]Item(nil), manager.tierOne[count:]...)
}

func (manager *Manager) Summarize(summary string, affected int) {
	if affected < 0 {
		affected = 0
	}
	if affected > len(manager.tierOne) {
		affected = len(manager.tierOne)
	}
	content, _ := json.Marshal(SummarySentinel + "\n" + summary)
	manager.tierTwo = append(manager.tierTwo, Item{Kind: "summary", Content: content})
	manager.EvictOldest(affected)
}

func (manager *Manager) Assemble() ([]Item, string, error) {
	digest, err := digestItems(manager.pinned)
	if err != nil || digest != manager.pinnedDigest {
		return nil, "", ErrPinnedIntegrity
	}
	assembled := append(cloneItems(manager.pinned), cloneItems(manager.tierTwo)...)
	assembled = append(assembled, cloneItems(manager.tierOne)...)
	surfaceDigest, err := digestItems(assembled)
	if err != nil {
		return nil, "", err
	}
	return assembled, surfaceDigest, nil
}

func (manager *Manager) Destroy() {
	manager.pinned = nil
	manager.tierOne = nil
	manager.tierTwo = nil
	manager.pinnedDigest = ""
}

// CorruptPinnedForTest is intentionally absent: integrity negatives mutate a
// package-local fixture directly, so production has no corruption handle.

type E0 struct {
	AttemptID  string `json:"attempt_id"`
	TurnID     string `json:"turn_id"`
	TurnEpoch  string `json:"turn_epoch"`
	Phase      string `json:"phase"`
	DenyReason string `json:"deny_reason,omitempty"`
}

type ProviderE0Input struct {
	AttemptID      string
	TurnID         string
	TurnEpoch      string
	Disposition    string
	AttemptStarted bool
	DenyReason     string
	WorkerLive     bool
	RetirementWon  bool
}

// ProviderE0Events realizes the total m-8 disposition to m-3 E0 table. An
// absent live populator and epoch-inert cuts emit nothing; sent means only
// that the normalized attempt_started event was observed.
func ProviderE0Events(input ProviderE0Input) ([]E0, error) {
	if !input.WorkerLive {
		return nil, nil
	}
	base := map[string]string{"attempt_id": input.AttemptID, "turn_id": input.TurnID, "turn_epoch": input.TurnEpoch}
	if _, err := BuildE0(withPhase(base, "sent", "")); err != nil {
		return nil, err
	}
	if input.Disposition == "STALE_EPOCH" || input.Disposition == "EPOCH_AHEAD" {
		if input.AttemptStarted || input.DenyReason != "" {
			return nil, errors.New("contextmgr: epoch-inert disposition carried live evidence")
		}
		return nil, nil
	}
	if input.Disposition == "stream_lost" && input.RetirementWon {
		return nil, nil
	}

	wantStarted := false
	phase := ""
	switch input.Disposition {
	case "completed":
		wantStarted, phase = true, "completed"
	case "transport_failed":
		wantStarted, phase = true, "failed"
	case "egress_denied":
		phase = "denied"
		if input.DenyReason == "" {
			return nil, errors.New("contextmgr: egress denial lacks deny_reason")
		}
	case "rejected_local":
		phase = "failed"
	case "cancelled_pre_transport":
		phase = "cancelled"
	case "cancelled_post_invocation":
		wantStarted, phase = true, "cancelled"
	case "stream_lost":
		wantStarted, phase = true, "unknown"
	default:
		return nil, errors.New("contextmgr: unknown provider disposition")
	}
	if input.AttemptStarted != wantStarted {
		return nil, errors.New("contextmgr: attempt_started observation contradicts disposition")
	}
	if phase != "denied" && input.DenyReason != "" {
		return nil, errors.New("contextmgr: deny_reason on non-denial")
	}
	var events []E0
	if wantStarted {
		sent, err := BuildE0(withPhase(base, "sent", ""))
		if err != nil {
			return nil, err
		}
		events = append(events, sent)
	}
	terminal, err := BuildE0(withPhase(base, phase, input.DenyReason))
	if err != nil {
		return nil, err
	}
	return append(events, terminal), nil
}

func withPhase(base map[string]string, phase, denyReason string) map[string]string {
	result := make(map[string]string, len(base)+2)
	for key, value := range base {
		result[key] = value
	}
	result["phase"] = phase
	if denyReason != "" {
		result["deny_reason"] = denyReason
	}
	return result
}

func BuildE0(input map[string]string) (E0, error) {
	for key, value := range input {
		if secretShaped(key) || secretShaped(value) {
			return E0{}, errors.New("contextmgr: secret-shaped E0 member refused")
		}
	}
	event := E0{AttemptID: input["attempt_id"], TurnID: input["turn_id"], TurnEpoch: input["turn_epoch"], Phase: input["phase"], DenyReason: input["deny_reason"]}
	if event.AttemptID == "" || event.TurnID == "" || event.TurnEpoch == "" {
		return E0{}, errors.New("contextmgr: incomplete E0 identity")
	}
	switch event.Phase {
	case "denied", "sent", "completed", "failed", "cancelled", "unknown":
	default:
		return E0{}, errors.New("contextmgr: invalid E0 phase")
	}
	if event.Phase != "denied" && event.DenyReason != "" {
		return E0{}, errors.New("contextmgr: deny_reason is policy-only")
	}
	return event, nil
}

func secretShaped(value string) bool {
	lower := strings.ToLower(value)
	for _, part := range []string{"credential", "api_key", "apikey", "authorization", "bearer ", "secret", "token="} {
		if strings.Contains(lower, part) {
			return true
		}
	}
	return false
}

func cloneItems(items []Item) []Item {
	out := make([]Item, len(items))
	for i, item := range items {
		out[i] = Item{Kind: item.Kind, Content: append(json.RawMessage(nil), item.Content...)}
	}
	return out
}
func digestItems(items []Item) (string, error) {
	encoded, err := json.Marshal(cloneItems(items))
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(encoded)
	return hex.EncodeToString(sum[:]), nil
}
