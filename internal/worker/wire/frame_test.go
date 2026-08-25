package wire

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestFrameRoundTripIsByteExact(t *testing.T) {
	codec := testCodec(t)
	epoch := "9007199254740993"
	frame := Frame{
		Version:   FrameVersion,
		Channel:   ChannelCTRLW,
		Type:      "attempt_open",
		Seq:       "18446744073709551615",
		RunID:     "run-1",
		TurnEpoch: epoch,
		Body:      json.RawMessage(`{"attempt_id":"attempt-1","provider_lane_id":"lane-1","turn_epoch":"9007199254740993","turn_id":"turn-1"}`),
	}

	wireBytes, err := codec.Encode(frame)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	decoded, err := codec.Decode(wireBytes)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	reencoded, err := codec.Encode(decoded)
	if err != nil {
		t.Fatalf("re-Encode: %v", err)
	}
	if !bytes.Equal(reencoded, wireBytes) {
		t.Fatalf("round trip changed bytes:\n got: %q\nwant: %q", reencoded, wireBytes)
	}
	if binary.BigEndian.Uint32(wireBytes[:4]) != uint32(len(wireBytes)-4) {
		t.Fatalf("length prefix = %d, payload length = %d", binary.BigEndian.Uint32(wireBytes[:4]), len(wireBytes)-4)
	}
}

func TestFrameRejectsMalformedEnvelopeClasses(t *testing.T) {
	codec := testCodec(t)
	cases := []struct {
		name    string
		payload string
	}{
		{name: "bad version", payload: `{"body":{},"chan":"ctrl-w","seq":"1","type":"attempt_open","v":2}`},
		{name: "unknown channel", payload: `{"body":{},"chan":"sideband","seq":"1","type":"attempt_open","v":1}`},
		{name: "unknown message type", payload: `{"body":{},"chan":"ctrl-w","seq":"1","type":"not_a_message","v":1}`},
		{name: "non-canonical counter", payload: `{"body":{},"chan":"ctrl-w","seq":"01","type":"attempt_open","v":1}`},
		{name: "numeric counter", payload: `{"body":{},"chan":"ctrl-w","seq":1,"type":"attempt_open","v":1}`},
		{name: "unknown envelope member", payload: `{"body":{},"chan":"ctrl-w","extra":true,"seq":"1","type":"attempt_open","v":1}`},
		{name: "non-canonical json", payload: `{ "body":{},"chan":"ctrl-w","seq":"1","type":"attempt_open","v":1}`},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if _, err := codec.Decode(frameBytes([]byte(testCase.payload))); err == nil {
				t.Fatal("Decode accepted malformed frame")
			}
		})
	}
}

func TestFrameReplyCorrelationAndClosedBody(t *testing.T) {
	codec := testCodec(t)

	missingReplyCorrelation := `{"body":{"attempt_id":"attempt-1"},"chan":"ctrl-w","seq":"2","type":"attempt_open_ok","v":1}`
	if _, err := codec.Decode(frameBytes([]byte(missingReplyCorrelation))); err == nil {
		t.Fatal("Decode accepted a reply-class frame without re")
	}

	requestWithReplyCorrelation := `{"body":{"attempt_id":"attempt-1","provider_lane_id":"lane-1","turn_epoch":"1","turn_id":"turn-1"},"chan":"ctrl-w","re":"1","seq":"2","type":"attempt_open","v":1}`
	if _, err := codec.Decode(frameBytes([]byte(requestWithReplyCorrelation))); err == nil {
		t.Fatal("Decode accepted re on a non-reply frame")
	}

	unknownClosedMember := `{"body":{"attempt_id":"attempt-1","extra":true},"chan":"ctrl-w","re":"1","seq":"2","type":"attempt_open_ok","v":1}`
	if _, err := codec.Decode(frameBytes([]byte(unknownClosedMember))); err == nil {
		t.Fatal("Decode accepted an unknown member in a closed message body")
	}
}

func TestFrameRejectsOversizeBeforePayloadAllocation(t *testing.T) {
	codec := testCodec(t)
	onlyPrefix := make([]byte, 4)
	binary.BigEndian.PutUint32(onlyPrefix, MaxFramePayload+1)
	_, err := codec.Decode(onlyPrefix)
	if !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("Decode oversize error = %v, want ErrFrameTooLarge", err)
	}
}

func TestFrameEncodeRejectsOversizePayload(t *testing.T) {
	codec := testCodec(t)
	frame := Frame{
		Version: FrameVersion,
		Channel: ChannelDATAP,
		Type:    "provider_chunk",
		Seq:     "1",
		Body:    json.RawMessage(`{"chunk":"` + strings.Repeat("x", MaxFramePayload) + `"}`),
	}
	_, err := codec.Encode(frame)
	if !errors.Is(err, ErrFrameTooLarge) {
		t.Fatalf("Encode oversize error = %v, want ErrFrameTooLarge", err)
	}
}

func TestParkedUnknownRowBoundAndClosedShape(t *testing.T) {
	row := ParkedUnknown{
		TurnID:              "turn-1",
		ToolCallID:          "call-1",
		TicketID:            "ticket-1",
		State:               ParkedUnknownToolOutcome,
		CanonicalToolName:   "bash",
		CanonicalArgsDigest: strings.Repeat("a", 64),
	}
	encoded, err := EncodeParkedUnknown(row)
	if err != nil {
		t.Fatalf("EncodeParkedUnknown: %v", err)
	}
	if len(encoded) > ParkedRowMax {
		t.Fatalf("parked row length = %d, max %d", len(encoded), ParkedRowMax)
	}
	decoded, err := DecodeParkedUnknown(encoded)
	if err != nil {
		t.Fatalf("DecodeParkedUnknown: %v", err)
	}
	if decoded != row {
		t.Fatalf("parked row = %#v, want %#v", decoded, row)
	}

	row.ToolCallID = strings.Repeat("x", ParkedRowMax)
	if _, err := EncodeParkedUnknown(row); !errors.Is(err, ErrParkedRowTooLarge) {
		t.Fatalf("oversize parked row error = %v, want ErrParkedRowTooLarge", err)
	}

	unknownMember := append(encoded[:len(encoded)-1], []byte(`,"extra":true}`)...)
	if _, err := DecodeParkedUnknown(unknownMember); err == nil {
		t.Fatal("DecodeParkedUnknown accepted an unknown member")
	}
}

func testCodec(t *testing.T) *Codec {
	t.Helper()
	codec, err := NewCodec([]MessageSpec{
		{Channel: ChannelCTRLW, Type: "attempt_open", ClosedBody: true, BodyMembers: []string{"attempt_id", "provider_lane_id", "turn_epoch", "turn_id"}},
		{Channel: ChannelCTRLW, Type: "attempt_open_ok", Reply: true, ClosedBody: true, BodyMembers: []string{"attempt_id", "parked_unknown"}},
		{Channel: ChannelDATAP, Type: "provider_chunk"},
	})
	if err != nil {
		t.Fatalf("NewCodec: %v", err)
	}
	return codec
}

func frameBytes(payload []byte) []byte {
	result := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(result[:4], uint32(len(payload)))
	copy(result[4:], payload)
	return result
}
