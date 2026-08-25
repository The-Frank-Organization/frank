package main

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/worker/executor"
	"github.com/jackli/frank/internal/worker/fake"
	workerruntime "github.com/jackli/frank/internal/worker/runtime"
	"github.com/jackli/frank/internal/worker/tools"
	"github.com/jackli/frank/internal/worker/turn"
)

func TestHandshakeFailuresFailClosedBeforeProviderOrToolWork(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*fake.M10, *fake.Broker)
	}{
		{"bad capability", func(_ *fake.M10, broker *fake.Broker) { broker.Capability = "" }},
		{"wrong epoch", func(_ *fake.M10, broker *fake.Broker) { broker.Result = workerruntime.AttachTupleMismatch }},
		{"dead broker", func(_ *fake.M10, broker *fake.Broker) { broker.Err = context.DeadlineExceeded }},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			control, broker, providerPeer, backend, config := fixture(t)
			testCase.mutate(control, broker)
			runner := workerruntime.Runner{Control: control, Broker: broker, Provider: providerPeer, Backend: backend}
			if _, err := runner.Run(context.Background(), config); err == nil {
				t.Fatal("handshake failure advanced")
			}
			if providerPeer.Attempts != 0 || control.AuthorizeCalls != 0 || control.ConsumeCalls != 0 || backend.Writes != 0 {
				t.Fatalf("failure was not inert: provider=%d authorize=%d consume=%d writes=%d", providerPeer.Attempts, control.AuthorizeCalls, control.ConsumeCalls, backend.Writes)
			}
		})
	}
}

func TestOneHonestGovernedTurnE2E(t *testing.T) {
	control, broker, providerPeer, backend, config := fixture(t)
	broker.Wakes = []string{"relay-1", "relay-1"}
	runner := workerruntime.Runner{Control: control, Broker: broker, Provider: providerPeer, Backend: backend}
	result, err := runner.Run(context.Background(), config)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if result.Terminal != turn.TurnCompleted || backend.Writes != 1 || backend.Files["answer.txt"] != "governed\n" {
		t.Fatalf("turn result=%+v files=%v writes=%d", result, backend.Files, backend.Writes)
	}
	if control.AuthorizeCalls != 1 || control.ConsumeCalls != 1 || len(control.Outcomes) != 1 || control.Outcomes[0].Outcome != executor.OutcomeExecuted {
		t.Fatalf("F59 trace authorize=%d consume=%d outcomes=%+v", control.AuthorizeCalls, control.ConsumeCalls, control.Outcomes)
	}
	if !bytes.Equal(result.PersistedTranscript, result.ReplayedTranscript) {
		t.Fatal("journal replay-fidelity readback changed bytes")
	}
	if len(control.Wakes) != 1 || control.Wakes[0] != "relay-1" {
		t.Fatalf("duplicate wake was not no-op at m-10: %v", control.Wakes)
	}
	wantOrder := []string{"hello", "attach", "attach_result:attach-ok", "attempt_open", "provider_attempt", "stream_end", "authorize", "consume", "record_outcome", "turn_terminal:turn_completed"}
	joined := strings.Join(fake.JoinTrace(control, broker, providerPeer), "|")
	last := -1
	for _, event := range wantOrder {
		index := strings.Index(joined, event)
		if index < 0 || index <= last {
			t.Fatalf("trace lacks governed order %q after %d: %s", event, last, joined)
		}
		last = index
	}
	if _, err := os.Stat(filepath.Join(config.RuntimeDir, "session.log")); err != nil {
		t.Fatalf("session log: %v", err)
	}
}

func TestWakeRelayObjectiveUsesWorkerSeatResolver(t *testing.T) {
	control, broker, providerPeer, backend, config := fixture(t)
	control.Assignment.AdmissionRef = turn.AdmissionRef{Kind: "wake_relay", RelayID: "relay-task"}
	source := &fake.ObjectiveSource{Relays: map[string]string{"relay-task": "durable governed objective"}}
	runner := workerruntime.Runner{Control: control, Broker: broker, Provider: providerPeer, Backend: backend, Objective: source}
	if _, err := runner.Run(context.Background(), config); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if source.Reads != 1 {
		t.Fatalf("objective reads = %d, want 1", source.Reads)
	}

	control2, broker2, provider2, backend2, config2 := fixture(t)
	control2.Assignment.AdmissionRef = turn.AdmissionRef{Kind: "wake_relay", RelayID: "relay-task"}
	runner = workerruntime.Runner{Control: control2, Broker: broker2, Provider: provider2, Backend: backend2}
	if _, err := runner.Run(context.Background(), config2); err == nil || provider2.Attempts != 0 || backend2.Writes != 0 {
		t.Fatalf("missing resolver advanced: provider=%d writes=%d err=%v", provider2.Attempts, backend2.Writes, err)
	}
}

func fixture(t *testing.T) (*fake.M10, *fake.Broker, *fake.M8, *fake.Backend, workerruntime.Config) {
	t.Helper()
	runtimeDir := filepath.Join(t.TempDir(), "runtime")
	if err := os.Mkdir(runtimeDir, 0o700); err != nil {
		t.Fatal(err)
	}
	assignment := workerruntime.Assignment{
		RunID: "run-1", TurnID: "turn-1", TurnEpoch: "7", ManifestDigest: strings.Repeat("a", 64),
		GenerationID: "generation-1", CreateAuthID: strings.Repeat("b", 32), BrokerEndpoint: "broker.sock",
		AdmissionRef: turn.AdmissionRef{Kind: "operator_input", TaskInput: "write the answer"},
	}
	control := fake.NewM10(assignment)
	broker := &fake.Broker{Result: workerruntime.AttachOK, Capability: "cap-worker-1"}
	providerPeer := &fake.M8{
		Disposition:  "completed",
		OpaqueItems:  [][]byte{[]byte(`{"type":"reasoning","opaque":"ciphertext"}`)},
		ScriptedTool: workerruntime.ToolCall{ID: "call-1", CanonicalName: "write", Arguments: []byte(`{"content":"governed\n","path":"answer.txt"}`)},
	}
	backend := fake.NewBackend()
	return control, broker, providerPeer, backend, workerruntime.Config{PID: 123, BuildInfo: "test-build", RuntimeDir: runtimeDir, RunDisposition: "fresh", AttachDeadline: time.Second, AttachBackoff: time.Millisecond}
}

var _ tools.Backend = (*fake.Backend)(nil)
