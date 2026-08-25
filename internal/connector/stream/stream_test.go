package stream

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/transport"
)

func TestParseEachReleasesEventsBeforeProviderTerminal(t *testing.T) {
	t.Parallel()

	reader, writer := io.Pipe()
	t.Cleanup(func() { _ = reader.Close() })
	t.Cleanup(func() { _ = writer.Close() })
	events := make(chan Event, 8)
	completed := make(chan error, 1)
	go func() {
		completed <- ParseEach(reader, testMeta(), func(event Event) error {
			events <- event
			return nil
		})
	}()
	if got := receiveEvent(t, events); got.Kind != AttemptStarted {
		t.Fatalf("first event = %s", got.Kind)
	}
	if _, err := io.WriteString(writer, "data: {\"type\":\"response.output_item.added\",\"output_index\":0,\"item\":{\"id\":\"msg-1\",\"type\":\"message\"}}\n\n"); err != nil {
		t.Fatal(err)
	}
	if got := receiveEvent(t, events); got.Kind != TextStart {
		t.Fatalf("pre-terminal event = %s", got.Kind)
	}
	select {
	case err := <-completed:
		t.Fatalf("parser completed before provider terminal: %v", err)
	default:
	}
	if _, err := io.WriteString(writer, "data: {\"type\":\"response.output_item.done\",\"output_index\":0,\"item\":{\"id\":\"msg-1\",\"type\":\"message\"}}\n\ndata: {\"type\":\"response.completed\",\"response\":{\"status\":\"completed\",\"usage\":{\"input_tokens\":1,\"output_tokens\":1}}}\n\n"); err != nil {
		t.Fatal(err)
	}
	_ = writer.Close()
	if got := receiveEvent(t, events); got.Kind != TextEnd {
		t.Fatalf("block end = %s", got.Kind)
	}
	if got := receiveEvent(t, events); got.Kind != Completed {
		t.Fatalf("terminal = %s", got.Kind)
	}
	select {
	case err := <-completed:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(time.Second):
		t.Fatal("streaming parser did not complete")
	}
}

func TestParseEachStopsImmediatelyOnEmitterFailure(t *testing.T) {
	t.Parallel()

	want := errors.New("injected DATA-P failure")
	calls := 0
	err := ParseEach(strings.NewReader(sse(`{"type":"response.created"}`)), testMeta(), func(Event) error {
		calls++
		return want
	})
	if !errors.Is(err, want) || calls != 1 {
		t.Fatalf("ParseEach() error=%v calls=%d, want emitter failure after attempt_started", err, calls)
	}
}

func receiveEvent(t *testing.T, events <-chan Event) Event {
	t.Helper()
	select {
	case event := <-events:
		return event
	case <-time.After(time.Second):
		t.Fatal("normalized event was not emitted incrementally")
		return Event{}
	}
}

const (
	testB = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	testE = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)

func TestParseTextCorpusAndAuthoritativeTerminalUsage(t *testing.T) {
	t.Parallel()

	events := Parse(strings.NewReader(sse(
		`{"type":"response.output_item.added","output_index":0,"item":{"id":"msg-1","type":"message"}}`,
		`{"type":"response.output_text.delta","output_index":0,"delta":"hello"}`,
		`{"type":"response.output_text.done","output_index":0}`,
		`{"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":11,"output_tokens":7,"output_tokens_details":{"reasoning_tokens":3},"input_tokens_details":{"cached_tokens":2}}}}`,
	)), testMeta())

	wantKinds := []Kind{AttemptStarted, TextStart, TextDelta, TextEnd, Completed}
	assertKinds(t, events, wantKinds)
	if events[2].Text != "hello" || events[4].FinishReason != "stop" {
		t.Fatalf("normalized text/finish = %+v", events)
	}
	if events[4].Usage == nil || events[4].Usage.InputTokens != 11 || events[4].Usage.OutputTokens != 7 || events[4].Usage.ReasoningTokens == nil || *events[4].Usage.ReasoningTokens != 3 || events[4].Usage.CacheReadTokens == nil || *events[4].Usage.CacheReadTokens != 2 {
		t.Fatalf("terminal usage = %+v", events[4].Usage)
	}
	assertTerminalDigests(t, events[4])
}

func TestToolFragmentsStayInertUntilComplete(t *testing.T) {
	t.Parallel()

	events := Parse(strings.NewReader(sse(
		`{"type":"response.output_item.added","output_index":0,"item":{"id":"item-1","type":"function_call","call_id":"call-1","name":"bash"}}`,
		`{"type":"response.function_call_arguments.delta","output_index":0,"delta":"{\"command\":"}`,
		`{"type":"response.function_call_arguments.delta","output_index":0,"delta":"\"pwd\"}"}`,
		`{"type":"response.function_call_arguments.done","output_index":0,"arguments":"{\"command\":\"pwd\"}"}`,
		`{"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":1,"output_tokens":1}}}`,
	)), testMeta())

	assertKinds(t, events, []Kind{AttemptStarted, ToolCallStart, ToolCallDelta, ToolCallDelta, ToolCallEnd, Completed})
	for _, event := range events[2:4] {
		if event.ToolCall != nil || event.Text != "" {
			t.Fatalf("fragment became active: %+v", event)
		}
	}
	call := events[4].ToolCall
	if call == nil || call.ToolCallID != "call-1" || call.Name != "bash" || string(call.Arguments) != `{"command":"pwd"}` {
		t.Fatalf("assembled tool call = %+v", call)
	}
	if events[5].FinishReason != "tool_calls" {
		t.Fatalf("finish reason = %q", events[5].FinishReason)
	}
}

func TestReasoningReplayPayloadIsOpaqueAndBytePreserved(t *testing.T) {
	t.Parallel()

	const item = `{"id":"reason-1","type":"reasoning","encrypted_content":"opaque+/==","summary":[]}`
	events := Parse(strings.NewReader(sse(
		`{"type":"response.output_item.added","output_index":0,"item":{"id":"reason-1","type":"reasoning"}}`,
		`{"type":"response.reasoning_summary_text.delta","output_index":0,"delta":"summary"}`,
		`{"type":"response.output_item.done","output_index":0,"item":`+item+`}`,
		`{"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":1,"output_tokens":1}}}`,
	)), testMeta())

	assertKinds(t, events, []Kind{AttemptStarted, ReasoningStart, ReasoningDelta, ReasoningEnd, Completed})
	replay := events[3].ReplayEnvelope
	if replay == nil || replay.OriginProviderLaneID != "lane-1" || replay.OriginTurnID != "turn-1" || string(replay.Payload) != item {
		t.Fatalf("replay envelope = %+v payload=%s", replay, replay.Payload)
	}
}

func TestSynthesizesStableBlockAndToolCallIDs(t *testing.T) {
	t.Parallel()

	events := Parse(strings.NewReader(sse(
		`{"type":"response.output_item.added","output_index":4,"item":{"type":"function_call","name":"read"}}`,
		`{"type":"response.function_call_arguments.done","output_index":4,"arguments":"{}"}`,
		`{"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":0,"output_tokens":0}}}`,
	)), testMeta())
	if events[1].BlockID != "m8-block-4" || events[2].ToolCall == nil || events[2].ToolCall.ToolCallID != "m8-call-4" {
		t.Fatalf("synthesized ids = %+v", events)
	}
}

func TestUnknownFinishAndMalformedArgumentsFailInStreamExactlyOnce(t *testing.T) {
	t.Parallel()

	corpora := []string{
		sse(`{"type":"response.incomplete","response":{"status":"incomplete","incomplete_details":{"reason":"future_reason"},"usage":{"input_tokens":1,"output_tokens":1}}}`),
		sse(
			`{"type":"response.output_item.added","output_index":0,"item":{"type":"function_call","name":"bash"}}`,
			`{"type":"response.function_call_arguments.done","output_index":0,"arguments":"[]"}`,
		),
		sse(
			`{"type":"response.output_item.added","output_index":0,"item":{"type":"message"}}`,
			`{"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":1,"output_tokens":1}}}`,
		),
	}
	for _, corpus := range corpora {
		events := Parse(strings.NewReader(corpus), testMeta())
		terminals := 0
		for _, event := range events {
			if event.Terminal() {
				terminals++
				if event.Kind != Failed || event.ErrorClass != "protocol" {
					t.Fatalf("terminal = %+v", event)
				}
				assertTerminalDigests(t, event)
			}
		}
		if terminals != 1 {
			t.Fatalf("terminal count = %d in %+v", terminals, events)
		}
	}
}

func TestReadFailuresBecomeTypedInStreamTerminals(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		err       error
		wantClass string
	}{
		{err: transport.ErrStallDeadline, wantClass: "timeout_stall"},
		{err: io.ErrUnexpectedEOF, wantClass: "transport"},
	} {
		events := Parse(errorReader{err: test.err}, testMeta())
		terminal := events[len(events)-1]
		if terminal.Kind != Failed || terminal.ErrorClass != test.wantClass {
			t.Fatalf("Parse(%v) terminal = %+v", test.err, terminal)
		}
	}
}

func TestOverflowBecomesFrameOverflowTerminal(t *testing.T) {
	t.Parallel()

	oversized := `data: {"type":"response.output_text.delta","delta":"` + strings.Repeat("x", int(frame.FrameMax)) + `"}\n\n`
	events := Parse(strings.NewReader(oversized), testMeta())
	if terminal := events[len(events)-1]; terminal.Kind != Failed || terminal.ErrorClass != "frame_overflow" {
		t.Fatalf("oversized terminal = %+v", terminal)
	}

	events = parseWithLimits(strings.NewReader(sse(
		`{"type":"response.output_item.added","output_index":0,"item":{"id":"msg-1","type":"message"}}`,
		`{"type":"response.output_text.delta","output_index":0,"delta":"hello"}`,
	)), testMeta(), limits{frameMax: frame.FrameMax, attemptMax: 32})
	if terminal := events[len(events)-1]; terminal.Kind != Failed || terminal.ErrorClass != "frame_overflow" {
		t.Fatalf("attempt overflow terminal = %+v", terminal)
	}
}

func TestMultilineSSEEventCannotEvadeFrameLimit(t *testing.T) {
	t.Parallel()

	padding := strings.Repeat("x", int(frame.FrameMax)/2)
	corpus := "data: {\"type\":\"response.output_text.delta\",\"padding_a\":\"" + padding + "\",\n" +
		"data: \"padding_b\":\"" + padding + "\",\"output_index\":0,\"delta\":\"x\"}\n\n"
	events := Parse(strings.NewReader(corpus), testMeta())
	if terminal := events[len(events)-1]; terminal.Kind != Failed || terminal.ErrorClass != "frame_overflow" {
		t.Fatalf("multiline oversized terminal = %+v", terminal)
	}
}

func TestEveryNormalizedEventFitsFrameMax(t *testing.T) {
	t.Parallel()

	events := Parse(strings.NewReader(sse(
		`{"type":"response.output_item.added","output_index":0,"item":{"id":"msg-1","type":"message"}}`,
		`{"type":"response.output_text.delta","output_index":0,"delta":"`+strings.Repeat("x", 1024*1024)+`"}`,
		`{"type":"response.output_text.done","output_index":0}`,
		`{"type":"response.completed","response":{"status":"completed","usage":{"input_tokens":1,"output_tokens":1}}}`,
	)), testMeta())
	for _, event := range events {
		raw, err := json.Marshal(event)
		if err != nil || len(raw) > int(frame.FrameMax) {
			t.Fatalf("event encoded length/error = %d/%v", len(raw), err)
		}
	}
}

func testMeta() Meta {
	return Meta{AttemptID: "attempt-1", ProviderLaneID: "lane-1", TurnID: "turn-1", FrozenCoreDigest: testB, ProviderLoweredToolsDigest: testE}
}

func sse(events ...string) string {
	var builder strings.Builder
	for _, event := range events {
		builder.WriteString("data: ")
		builder.WriteString(event)
		builder.WriteString("\n\n")
	}
	return builder.String()
}

func assertKinds(t *testing.T, events []Event, want []Kind) {
	t.Helper()
	if len(events) != len(want) {
		t.Fatalf("event count = %d, want %d: %+v", len(events), len(want), events)
	}
	for index := range want {
		if events[index].Kind != want[index] || events[index].Schema != SchemaV2 {
			t.Fatalf("event[%d] = %+v, want kind %s/v2", index, events[index], want[index])
		}
	}
}

func assertTerminalDigests(t *testing.T, event Event) {
	t.Helper()
	if !event.Terminal() || event.FrozenCoreDigest != testB || event.ProviderLoweredToolsDigest != testE {
		t.Fatalf("terminal digest carriage = %+v", event)
	}
}

type errorReader struct{ err error }

func (reader errorReader) Read([]byte) (int, error) {
	return 0, reader.err
}
