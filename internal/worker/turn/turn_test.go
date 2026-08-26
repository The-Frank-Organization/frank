package turn

import (
	"errors"
	"testing"
	"time"
)

func TestOneActiveTurnAndStaleEpochAreInert(t *testing.T) {
	machine := New()
	if err := machine.Admit(validOpen(nil), time.Now()); err != nil {
		t.Fatal(err)
	}
	if err := machine.Admit(validOpen(nil), time.Now()); !errors.Is(err, ErrTurnActive) {
		t.Fatalf("second admit = %v", err)
	}
	if err := machine.BeginAssembly(6); !errors.Is(err, ErrStaleEpoch) {
		t.Fatalf("stale transition = %v", err)
	}
	state, _, attempts, tools := machine.Snapshot()
	if state != StateAdmitted || attempts != 0 || tools != 0 {
		t.Fatalf("stale changed state: %s %d %d", state, attempts, tools)
	}
}

func TestParkedUnknownComparatorTotalAndGateOrdering(t *testing.T) {
	a := parked("a", "UNKNOWN_TOOL_OUTCOME")
	b := parked("b", "PARTIAL_TOOL_EFFECT")
	changed := a
	changed.State = "PARTIAL_TOOL_EFFECT"
	tests := []struct {
		name        string
		left, right []ParkedUnknown
		want        ParkedRelation
	}{
		{"empty", []ParkedUnknown{}, []ParkedUnknown{}, ParkedEqual},
		{"equal", []ParkedUnknown{a}, []ParkedUnknown{a}, ParkedEqual},
		{"added", []ParkedUnknown{a}, []ParkedUnknown{a, b}, ParkedAddedOrChanged},
		{"changed", []ParkedUnknown{a}, []ParkedUnknown{changed}, ParkedAddedOrChanged},
		{"removed", []ParkedUnknown{a, b}, []ParkedUnknown{a}, ParkedRemovedOnly},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := CompareParked(test.left, test.right)
			if err != nil || got != test.want {
				t.Fatalf("got=%s err=%v want=%s", got, err, test.want)
			}
		})
	}
	machine := New()
	if err := machine.Admit(validOpen([]ParkedUnknown{a}), time.Now()); err != nil {
		t.Fatal(err)
	}
	if err := machine.BeginAssembly(7); err != nil {
		t.Fatal(err)
	}
	if err := machine.AttemptOpenOK(7, []ParkedUnknown{a, b}); !errors.Is(err, ErrReassemblyDue) {
		t.Fatalf("gate two=%v", err)
	}
	state, _, attempts, _ := machine.Snapshot()
	if state != StateAssembling || attempts != 0 {
		t.Fatalf("provider work crossed changed disclosure: %s attempts=%d", state, attempts)
	}
}

func TestMalformedAndDuplicateParkedUnknownFailClosed(t *testing.T) {
	a := parked("a", "UNKNOWN_TOOL_OUTCOME")
	for _, items := range [][]ParkedUnknown{{{}}, {a, a}} {
		if _, err := CompareParked(nil, items); !errors.Is(err, ErrProtocol) {
			t.Fatalf("items=%#v err=%v", items, err)
		}
	}
}

func TestEveryCompiledBoundTerminatesExhausted(t *testing.T) {
	t.Run("attempts", func(t *testing.T) {
		m := New()
		_ = m.Admit(validOpen(nil), time.Now())
		for i := 0; i <= MaxProviderAttempts; i++ {
			if i == 0 {
				_ = m.BeginAssembly(7)
			}
			_ = m.AttemptOpenOK(7, nil)
			state, term, _, _ := m.Snapshot()
			if state == StateTerminal {
				if term != TurnExhausted {
					t.Fatal(term)
				}
				return
			}
			_ = m.Observe(7)
			_ = m.ToolRound(7, "call-"+string(rune(i)), false)
			_ = m.Reassemble(7)
		}
		t.Fatal("attempt bound did not trip")
	})
	t.Run("tool calls and denials have no worker-local ceiling", func(t *testing.T) {
		m := New()
		_ = m.Admit(validOpen(nil), time.Now())
		_ = m.BeginAssembly(7)
		_ = m.AttemptOpenOK(7, nil)
		_ = m.Observe(7)
		for i := 0; i <= 64; i++ {
			_ = m.ToolRound(7, "unique-"+string(rune(i)), true)
			state, term, _, _ := m.Snapshot()
			if state == StateTerminal {
				t.Fatalf("worker imposed a local tool ceiling: %s", term)
			}
			m.state = StateObserving
		}
	})
	t.Run("doom loop", func(t *testing.T) {
		m := New()
		_ = m.Admit(validOpen(nil), time.Now())
		_ = m.BeginAssembly(7)
		_ = m.AttemptOpenOK(7, nil)
		_ = m.Observe(7)
		for i := 0; i < DoomLoopThreshold; i++ {
			_ = m.ToolRound(7, "same", false)
			state, term, _, _ := m.Snapshot()
			if state == StateTerminal {
				if term != TurnExhausted {
					t.Fatal(term)
				}
				return
			}
			m.state = StateObserving
		}
		t.Fatal("doom loop did not trip")
	})
	t.Run("wall clock", func(t *testing.T) {
		m := New()
		_ = m.Admit(validOpen(nil), time.Now().Add(-MaxWallClock-time.Second))
		_ = m.BeginAssembly(7)
		_ = m.AttemptOpenOK(7, nil)
		state, term, _, _ := m.Snapshot()
		if state != StateTerminal || term != TurnExhausted {
			t.Fatalf("state=%s terminal=%s", state, term)
		}
	})
}

func TestCancelAckExactlyOnceAndPrecedesCancelledTerminal(t *testing.T) {
	m := New()
	_ = m.Admit(validOpen(nil), time.Now())
	_ = m.BeginAssembly(7)
	if err := m.Finish(7, TurnCancelled); !errors.Is(err, ErrProtocol) {
		t.Fatalf("cancel without ack=%v", err)
	}
	if err := m.CancelAck(7); err != nil {
		t.Fatal(err)
	}
	if err := m.CancelAck(7); !errors.Is(err, ErrProtocol) {
		t.Fatalf("duplicate ack=%v", err)
	}
	if err := m.Finish(7, TurnCancelled); err != nil {
		t.Fatal(err)
	}
}

func validOpen(items []ParkedUnknown) Open {
	return Open{RunID: "run", TurnID: "turn", TurnEpoch: "7", AdmissionRef: AdmissionRef{Kind: "operator_input", TaskInput: "task"}, ParkedUnknown: items}
}
func parked(id, state string) ParkedUnknown {
	return ParkedUnknown{TurnID: "turn-" + id, ToolCallID: "call-" + id, TicketID: "ticket-" + id, State: state, CanonicalToolName: "bash", CanonicalArgsDigest: "digest-" + id}
}
