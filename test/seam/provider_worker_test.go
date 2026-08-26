//go:build seam

package seam_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"testing"
	"time"

	connectorcatalog "github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/outcome"
	"github.com/jackli/frank/internal/connector/request"
	"github.com/jackli/frank/internal/worker/executor"
	"github.com/jackli/frank/internal/worker/journal"
	"github.com/jackli/frank/internal/worker/provider"
	"github.com/jackli/frank/internal/worker/turn"
)

// TestCT_B01 binds CT-B01; selectors r9 B01, S-M8v1 §1.2 and A-M8-BE v2; flags MISMATCH, W/MASTER, star.
func TestCT_B01(t *testing.T) {
	probe := newWorkerProbe(t)
	probe.provider.items = completeV2Events()
	probe.provider.nextErr = errors.New("tool call must be consumed from the v2 stream")
	_, err := probe.run(context.Background())
	contract(t, err == nil && probe.backend.Writes == 1, explain("complete v2 event stream was not consumed behaviorally: writes=%d err=%v", probe.backend.Writes, err))
}

// TestCT_B02 binds CT-B02; selectors r9 B02 and A-M8-BE; flags MISMATCH, MASTER/W, latent.
func TestCT_B02(t *testing.T) {
	v2 := newWorkerProbe(t)
	v2.provider.items = completeV2Events()
	v2.provider.nextErr = errors.New("tool call must be consumed from the v2 stream")
	_, v2Err := v2.run(context.Background())
	v1 := newWorkerProbe(t)
	v1.provider.items = [][]byte{normalizedEvent("m8.provider_event.v1", "tool_call_end", map[string]any{"tool_call_id": "call", "name": "write", "arguments": map[string]any{"path": "answer.txt", "content": "x"}})}
	v1.provider.nextErr = errors.New("v1 must be rejected before tool extraction")
	_, v1Err := v1.run(context.Background())
	contract(t, v2Err == nil && v1Err != nil && v1.backend.Writes == 0,
		explain("worker version gate: v2=%v v1=%v v1_writes=%d", v2Err, v1Err, v1.backend.Writes))
}

// TestCT_B03 binds CT-B03; selectors r9 B03 and S-M8v1 C§3; flags MISMATCH, W, latent.
func TestCT_B03(t *testing.T) {
	probe := newWorkerProbe(t)
	probe.provider.items = [][]byte{normalizedEvent("m8.provider_event.v2", "attempt_started", map[string]any{"attempt_id": "attempt-1"})}
	_, err := probe.run(context.Background())
	contract(t, err == nil && len(probe.control.e0Events) == 1 && len(probe.control.e0Events[0]) != 0,
		explain("observed provider events did not reach E0: batches=%d err=%v", len(probe.control.e0Events), err))
}

// TestCT_B04 binds CT-B04; selectors r9 B04 and S-M8v1 C§4; flags MISMATCH, W, star.
func TestCT_B04(t *testing.T) {
	cases := []struct {
		wire string
		want provider.Disposition
	}{
		{"sent_completed", provider.Completed},
		{"denied(policy-unavailable)", provider.EgressDenied},
		{"rejected_local(malformed_request)", provider.RejectedLocal},
		{"transport_failed", provider.TransportFailed},
		{"unknown", provider.StreamLost},
		{"cancelled(pre_transport)", provider.CancelledPre},
	}
	var failures []string
	for _, test := range cases {
		cycle, err := provider.New(opaqueGate{}, dispositionConnector{disposition: provider.Disposition(test.wire)}, "1")
		if err != nil {
			t.Fatal(err)
		}
		got, runErr := cycle.Run(context.Background(), provider.Request{AttemptID: "a", TurnID: "t", TurnEpoch: "1", ProviderLane: "p"})
		if runErr != nil || got.Disposition != test.want {
			failures = append(failures, explain("%s=>%s/%v", test.wire, got.Disposition, runErr))
		}
	}
	lossCycle, err := provider.New(opaqueGate{}, dispositionConnector{err: errors.New("fresh dial failed")}, "1")
	if err != nil {
		t.Fatal(err)
	}
	loss, lossErr := lossCycle.Run(context.Background(), provider.Request{AttemptID: "loss", TurnID: "t", TurnEpoch: "1", ProviderLane: "p"})
	if lossErr == nil || loss.Disposition != "" || len(loss.Events) != 0 {
		failures = append(failures, explain("fresh-dial cut minted carrier disposition=%q events=%d err=%v", loss.Disposition, len(loss.Events), lossErr))
	}
	contract(t, len(failures) == 0, explain("six-way/conditional disposition translation mismatch: %v", failures))
}

type dispositionConnector struct {
	disposition provider.Disposition
	err         error
}

func (connector dispositionConnector) Attempt(context.Context, provider.Request) (provider.Disposition, []json.RawMessage, error) {
	return connector.disposition, nil, connector.err
}
func (dispositionConnector) Cancel(context.Context, string, string) (provider.Disposition, error) {
	return provider.CancelledPre, nil
}

// TestCT_B05 binds CT-B05; selectors r9 B05 and S-M8v1 §1.1/C§12; flags MISMATCH, W, star.
func TestCT_B05(t *testing.T) {
	probe := newWorkerProbe(t)
	_, runErr := probe.run(context.Background())
	parsed, parseErr := request.Parse(probe.provider.request.OpaqueRequest, connectorcatalog.Lane{})
	contract(t, runErr == nil && parseErr == nil && parsed != nil && parsed.Schema == "m8.llm_request.v1" && parsed.ProviderLaneID != "" && parsed.Sampling.MaxOutputTokens > 0,
		explain("live worker request rejected by connector parser: run=%v parse=%v", runErr, parseErr))
}

// TestCT_B06 binds CT-B06; selectors r9 B06 and S-M8v1 C§12b; flags MISMATCH, W, star.
func TestCT_B06(t *testing.T) {
	probe := newWorkerProbe(t)
	probe.provider.items = [][]byte{normalizedEvent("m8.provider_event.v2", "tool_call_end", map[string]any{"tool_call_id": "call", "name": "WRITE", "arguments": map[string]any{"path": "answer.txt", "content": "x"}})}
	probe.provider.nextErr = errors.New("tool call must originate at tool_call_end")
	_, err := probe.run(context.Background())
	contract(t, err == nil && probe.backend.Files["answer.txt"] == "x", explain("tool_call_end lowering/effect: file=%q err=%v", probe.backend.Files["answer.txt"], err))
}

// TestCT_B07 binds CT-B07; selectors r9 B07 and S-M8v1/S-M8BE C§6; flags MISMATCH, W, latent.
func TestCT_B07(t *testing.T) {
	probe := newWorkerProbe(t)
	probe.provider.items = [][]byte{normalizedEvent("m8.provider_event.v2", "usage", map[string]any{"input_tokens": 7, "output_tokens": 3})}
	result, err := probe.run(context.Background())
	recorded := bytes.Contains(result.PersistedTranscript, []byte(`"usage":{"input_tokens":7,"output_tokens":3}`))
	contract(t, err == nil && recorded, explain("terminal usage not recorded structurally: recorded=%v err=%v", recorded, err))
}

// TestCT_B08 binds CT-B08; selectors r9 B08 and S-M8v1 C§7; flags MISMATCH, CN/W, latent.
func TestCT_B08(t *testing.T) {
	path := []string{"internal/worker/provider/provider.go"}
	dispositions := sourcesContain(t, path, "sent_completed", "denied", "rejected_local", "transport_failed", "unknown", "cancelled")
	reasons := sourcesContain(t, path, "malformed_request", "lane_capability_mismatch", "replay_scope_violation", "internal_integrity_fault")
	known := []provider.Disposition{"sent_completed", "denied", "rejected_local", "transport_failed", "unknown", "cancelled"}
	var consumptionFailures []string
	for _, disposition := range known {
		cycle, err := provider.New(opaqueGate{}, dispositionConnector{disposition: disposition}, "1")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := cycle.Run(context.Background(), provider.Request{AttemptID: "a", TurnID: "t", TurnEpoch: "1", ProviderLane: "p"}); err != nil {
			consumptionFailures = append(consumptionFailures, string(disposition)+":"+err.Error())
		}
	}
	unknownCycle, err := provider.New(opaqueGate{}, dispositionConnector{disposition: "future_disposition"}, "1")
	if err != nil {
		t.Fatal(err)
	}
	_, unknownErr := unknownCycle.Run(context.Background(), provider.Request{AttemptID: "a", TurnID: "t", TurnEpoch: "1", ProviderLane: "p"})
	contract(t, dispositions && reasons && len(consumptionFailures) == 0 && unknownErr != nil,
		explain("worker enum membership/consumption differs: six=%v four=%v known=%v unknown=%v", dispositions, reasons, consumptionFailures, unknownErr))
}

// TestCT_B09 binds CT-B09; selectors r9 B09 and S-M8v1 C§8; flags MISMATCH, W, latent.
func TestCT_B09(t *testing.T) {
	cycle, err := provider.New(opaqueGate{}, opaqueConnector{}, "1")
	if err != nil {
		t.Fatal(err)
	}
	result, err := cycle.Run(context.Background(), provider.Request{AttemptID: "a", TurnID: "t", TurnEpoch: "2", ProviderLane: "p"})
	malformed, malformedErr := cycle.Run(context.Background(), provider.Request{TurnID: "t", TurnEpoch: "1", ProviderLane: "p"})
	contract(t, err == nil && result.Disposition == provider.EpochAhead && malformedErr != nil && malformed.Disposition != provider.StaleEpoch,
		explain("epoch/malformed classification: ahead=%q/%v malformed=%q/%v", result.Disposition, err, malformed.Disposition, malformedErr))
}

// TestCT_B10 binds CT-B10; selectors r9 B10 and S-M8v1/S-M8BE C§11; flags MISMATCH, W, star.
func TestCT_B10(t *testing.T) {
	roundTrip := func(item []byte) ([]byte, string, error) {
		probe := newWorkerProbe(t)
		probe.provider.items = [][]byte{item}
		result, err := probe.run(context.Background())
		if err != nil {
			return nil, "", err
		}
		for _, line := range bytes.Split(result.PersistedTranscript, []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			record, decodeErr := journal.DecodeRecord(line)
			if decodeErr != nil {
				return nil, "", decodeErr
			}
			if record.Kind != journal.KindProviderOutput {
				continue
			}
			verbatim, hasVerbatim := record.Fields["verbatim"]
			rawBase64, hasRawBase64 := record.Fields["raw_b64"]
			if hasVerbatim == hasRawBase64 {
				return nil, "", errors.New("provider item carrier must select exactly one branch")
			}
			if hasVerbatim {
				return append([]byte(nil), verbatim...), "verbatim", nil
			}
			var encoded string
			if err := json.Unmarshal(rawBase64, &encoded); err != nil {
				return nil, "", err
			}
			decoded, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil || base64.StdEncoding.EncodeToString(decoded) != encoded {
				return nil, "", errors.New("raw_b64 is not canonical RFC 4648 standard base64")
			}
			return decoded, "raw_b64", nil
		}
		return nil, "", errors.New("provider_output record not found")
	}
	invalidUTF8 := []byte{0xff, 0xfe, 0x00, 'x'}
	canonicalJCS := []byte(`{"provider_native":{"encrypted":"opaque"},"type":"reasoning"}`)
	invalidRoundTrip, invalidBranch, invalidErr := roundTrip(invalidUTF8)
	canonicalRoundTrip, canonicalBranch, canonicalErr := roundTrip(canonicalJCS)
	contract(t, invalidErr == nil && canonicalErr == nil && invalidBranch == "raw_b64" && canonicalBranch == "verbatim" && bytes.Equal(invalidRoundTrip, invalidUTF8) && bytes.Equal(canonicalRoundTrip, canonicalJCS),
		explain("worker closed opaque-item carrier mismatch: invalid=%s/%v canonical=%s/%v", invalidBranch, invalidErr, canonicalBranch, canonicalErr))
}

// TestCT_B11 binds CT-B11; selectors r9 B11 and A-M9-LIFECYCLE §2.4; flags MISMATCH, W, latent.
func TestCT_B11(t *testing.T) {
	attemptMachine := turn.New()
	if err := attemptMachine.Admit(turn.Open{RunID: "run", TurnID: "turn", TurnEpoch: "1", AdmissionRef: turn.AdmissionRef{Kind: "operator_input", TaskInput: "x"}}, time.Now()); err != nil {
		t.Fatal(err)
	}
	if err := attemptMachine.BeginAssembly(1); err != nil {
		t.Fatal(err)
	}
	for index := 0; index < 17; index++ {
		if err := attemptMachine.AttemptOpenOK(1, nil); err != nil {
			t.Fatal(err)
		}
		state, _, _, _ := attemptMachine.Snapshot()
		if state == turn.StateTerminal {
			break
		}
		if err := attemptMachine.Observe(1); err != nil {
			t.Fatal(err)
		}
		if err := attemptMachine.ToolRound(1, explain("call-%d", index), false); err != nil {
			t.Fatal(err)
		}
		if err := attemptMachine.Reassemble(1); err != nil {
			t.Fatal(err)
		}
	}
	_, attemptTerminal, attempts, _ := attemptMachine.Snapshot()

	wallMachine := turn.New()
	_ = wallMachine.Admit(turn.Open{RunID: "run", TurnID: "turn", TurnEpoch: "1", AdmissionRef: turn.AdmissionRef{Kind: "operator_input", TaskInput: "x"}}, time.Now().Add(-31*time.Minute))
	_ = wallMachine.BeginAssembly(1)
	_ = wallMachine.AttemptOpenOK(1, nil)
	_, wallTerminal, _, _ := wallMachine.Snapshot()

	probe := newWorkerProbe(t)
	probe.backend.bashOutput = string(bytes.Repeat([]byte{'x'}, 2<<20))
	probe.provider.tool.CanonicalName = "bash"
	probe.provider.tool.Arguments = []byte(`{"command":"probe","timeout_ms":1}`)
	result, runErr := probe.run(context.Background())
	honestTruncation := len(result.PersistedTranscript) < 2<<20 && bytes.Contains(result.PersistedTranscript, []byte(`"truncated":true`))

	prepared, _ := executor.Prepare("run", "turn", "call", "read", "1", staticArguments(`{"path":"x"}`))
	invoker := &countingInvoker{}
	ceilingWorker, _ := executor.New(&replyAuthority{authorizeReply: executor.AuthorizeReply{Code: executor.AuthorizeRejected, RejectReason: executor.RejectTurnBudgetExhausted}}, invoker)
	_, ceilingErr := ceilingWorker.Execute(context.Background(), prepared)
	contract(t, attemptTerminal == turn.TurnExhausted && attempts == 17 && wallTerminal == turn.TurnExhausted && runErr == nil && honestTruncation && ceilingErr != nil && invoker.calls == 0,
		explain("lifecycle bounds: attempts=%d/%s wall=%s output_bytes=%d truncated=%v run=%v ceiling=%v invokes=%d", attempts, attemptTerminal, wallTerminal, len(result.PersistedTranscript), honestTruncation, runErr, ceilingErr, invoker.calls))
}

func normalizedEvent(schema, kind string, body map[string]any) []byte {
	value := map[string]any{"schema": schema, "kind": kind, "body": body}
	raw, _ := json.Marshal(value)
	return raw
}

func completeV2Events() [][]byte {
	return [][]byte{
		normalizedEvent("m8.provider_event.v2", "attempt_started", map[string]any{"attempt_id": "attempt-1"}),
		normalizedEvent("m8.provider_event.v2", "tool_call_end", map[string]any{"tool_call_id": "call", "name": "WRITE", "arguments": map[string]any{"path": "answer.txt", "content": "x"}}),
		normalizedEvent("m8.provider_event.v2", "usage", map[string]any{"input_tokens": 7, "output_tokens": 3}),
	}
}

var _ = outcome.AttemptResultSchemaV2
var _ = request.MalformedRequest
