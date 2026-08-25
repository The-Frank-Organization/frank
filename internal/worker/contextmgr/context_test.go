package contextmgr

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPinnedTierSurvivesEveryEvictionAndSummary(t *testing.T) {
	pinned := []Item{{Kind: "system", Content: json.RawMessage(`"hard constraint"`)}, {Kind: "objective", Content: json.RawMessage(`"task"`)}}
	manager, _ := New(pinned)
	manager.AddEvictable(Item{Kind: "tool", Content: json.RawMessage(`"large"`)}, Item{Kind: "chat", Content: json.RawMessage(`"old"`)})
	manager.Summarize("derived", 1)
	manager.EvictOldest(99)
	assembled, _, err := manager.Assemble()
	if err != nil {
		t.Fatal(err)
	}
	if len(assembled) != 3 || assembled[0].Kind != "system" || assembled[1].Kind != "objective" {
		t.Fatalf("assembled=%+v", assembled)
	}
	if !strings.Contains(string(assembled[2].Content), SummarySentinel) {
		t.Fatalf("summary lacks sentinel: %s", assembled[2].Content)
	}
}

func TestPinnedIntegrityFailureIsFailClosed(t *testing.T) {
	manager, _ := New([]Item{{Kind: "system", Content: json.RawMessage(`"original"`)}})
	manager.pinned[0].Content = json.RawMessage(`"mutated"`)
	if _, _, err := manager.Assemble(); err != ErrPinnedIntegrity {
		t.Fatalf("err=%v", err)
	}
}

func TestLogicalSurfaceDigestStableNoOpAndMovesOnAnySurfaceChange(t *testing.T) {
	manager, _ := New([]Item{{Kind: "system", Content: json.RawMessage(`"pinned"`)}})
	_, first, _ := manager.Assemble()
	_, second, _ := manager.Assemble()
	if first != second {
		t.Fatalf("no-op digest moved %s %s", first, second)
	}
	manager.AddEvictable(Item{Kind: "input", Content: json.RawMessage(`"new"`)})
	_, third, _ := manager.Assemble()
	if third == first {
		t.Fatal("surface change did not move digest")
	}
}

func TestE0RedactionRejectsSecretShapedMembers(t *testing.T) {
	base := map[string]string{"attempt_id": "a", "turn_id": "t", "turn_epoch": "7", "phase": "completed"}
	if _, err := BuildE0(base); err != nil {
		t.Fatal(err)
	}
	for _, pair := range [][2]string{{"api_key", "value"}, {"note", "Bearer abc"}, {"credential_ref", "opaque"}, {"note", "token=abc"}} {
		copy := map[string]string{}
		for k, v := range base {
			copy[k] = v
		}
		copy[pair[0]] = pair[1]
		if _, err := BuildE0(copy); err == nil {
			t.Fatalf("secret-shaped pair passed: %#v", pair)
		}
	}
}

func TestProviderE0TotalTableAndNoEmissionCuts(t *testing.T) {
	cases := []struct {
		disposition string
		started     bool
		deny        string
		phases      []string
	}{
		{"completed", true, "", []string{"sent", "completed"}},
		{"transport_failed", true, "", []string{"sent", "failed"}},
		{"egress_denied", false, "policy_denied", []string{"denied"}},
		{"rejected_local", false, "", []string{"failed"}},
		{"cancelled_pre_transport", false, "", []string{"cancelled"}},
		{"cancelled_post_invocation", true, "", []string{"sent", "cancelled"}},
		{"stream_lost", true, "", []string{"sent", "unknown"}},
		{"STALE_EPOCH", false, "", nil},
		{"EPOCH_AHEAD", false, "", nil},
	}
	for _, testCase := range cases {
		events, err := ProviderE0Events(ProviderE0Input{AttemptID: "a", TurnID: "t", TurnEpoch: "7", Disposition: testCase.disposition, AttemptStarted: testCase.started, DenyReason: testCase.deny, WorkerLive: true})
		if err != nil {
			t.Fatalf("%s: %v", testCase.disposition, err)
		}
		if len(events) != len(testCase.phases) {
			t.Fatalf("%s emitted %d events, want %d", testCase.disposition, len(events), len(testCase.phases))
		}
		for index := range events {
			if events[index].Phase != testCase.phases[index] {
				t.Fatalf("%s phase[%d] = %q, want %q", testCase.disposition, index, events[index].Phase, testCase.phases[index])
			}
		}
	}
	for _, input := range []ProviderE0Input{
		{AttemptID: "a", TurnID: "t", TurnEpoch: "7", Disposition: "completed", AttemptStarted: false, WorkerLive: true},
		{AttemptID: "a", TurnID: "t", TurnEpoch: "7", Disposition: "egress_denied", AttemptStarted: false, WorkerLive: true},
	} {
		if _, err := ProviderE0Events(input); err == nil {
			t.Fatalf("contradictory row passed: %+v", input)
		}
	}
	if events, err := ProviderE0Events(ProviderE0Input{WorkerLive: false}); err != nil || len(events) != 0 {
		t.Fatalf("crashed worker emitted E0: %v %v", events, err)
	}
	if events, err := ProviderE0Events(ProviderE0Input{AttemptID: "a", TurnID: "t", TurnEpoch: "7", Disposition: "stream_lost", AttemptStarted: true, WorkerLive: true, RetirementWon: true}); err != nil || len(events) != 0 {
		t.Fatalf("retirement-won race emitted E0: %v %v", events, err)
	}
}

func TestDestroyClearsEveryTier(t *testing.T) {
	manager, _ := New([]Item{{Kind: "system", Content: json.RawMessage(`"p"`)}})
	manager.AddEvictable(Item{Kind: "input", Content: json.RawMessage(`"x"`)})
	manager.Summarize("s", 1)
	manager.Destroy()
	if manager.pinned != nil || manager.tierOne != nil || manager.tierTwo != nil || manager.pinnedDigest != "" {
		t.Fatalf("manager retained context: %+v", manager)
	}
}
