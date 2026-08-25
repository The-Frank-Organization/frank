package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"testing"
	"time"
)

func TestCounterAcceptsOnlyCanonicalUint64Strings(t *testing.T) {
	t.Parallel()

	for _, text := range []string{"0", "1", "42", "18446744073709551615"} {
		counter, err := ParseCounter(text)
		if err != nil {
			t.Fatalf("ParseCounter(%q) error = %v", text, err)
		}
		if counter.String() != text {
			t.Fatalf("ParseCounter(%q).String() = %q", text, counter.String())
		}
	}
	for _, text := range []string{"", "00", "01", "+1", "-1", " 1", "1 ", "18446744073709551616"} {
		if _, err := ParseCounter(text); err == nil {
			t.Errorf("ParseCounter(%q) accepted a non-canonical counter", text)
		}
	}
}

func TestCodecWritesLengthPrefixedCanonicalEnvelope(t *testing.T) {
	t.Parallel()

	encoded, err := Encode(Envelope{
		Version: 1,
		Channel: ChannelControlConnector,
		Type:    "ping",
		Seq:     Counter(7),
		RunID:   "run-1",
		Body:    []byte(`{"z":2,"a":1}`),
	})
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	wantPayload := []byte(`{"body":{"a":1,"z":2},"chan":"ctrl-c","run_id":"run-1","seq":"7","type":"ping","v":1}`)
	if got := binary.BigEndian.Uint32(encoded[:4]); got != uint32(len(wantPayload)) {
		t.Fatalf("length prefix = %d, want %d", got, len(wantPayload))
	}
	if !bytes.Equal(encoded[4:], wantPayload) {
		t.Fatalf("payload = %s, want %s", encoded[4:], wantPayload)
	}
}

func TestDecoderRejectsOversizedFrameBeforeReadingPayload(t *testing.T) {
	t.Parallel()

	var prefix [4]byte
	binary.BigEndian.PutUint32(prefix[:], FrameMax+1)
	decoder := NewDecoder(map[string]TypeSpec{"ping": {}})
	if _, err := decoder.Read(bytes.NewReader(prefix[:])); !errors.Is(err, ErrFrameOverflow) {
		t.Fatalf("Read(oversized) error = %v, want ErrFrameOverflow", err)
	}
}

func TestDecoderRejectsMalformedCounterAndUnknownType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload string
		want    error
	}{
		{name: "leading-zero sequence", payload: `{"body":{},"chan":"ctrl-c","seq":"01","type":"ping","v":1}`, want: ErrMalformedFrame},
		{name: "numeric sequence", payload: `{"body":{},"chan":"ctrl-c","seq":1,"type":"ping","v":1}`, want: ErrMalformedFrame},
		{name: "unknown type", payload: `{"body":{},"chan":"ctrl-c","seq":"1","type":"surprise","v":1}`, want: ErrUnknownMessage},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			decoder := NewDecoder(map[string]TypeSpec{"ping": {}})
			if _, err := decoder.Read(bytes.NewReader(framed([]byte(test.payload)))); !errors.Is(err, test.want) {
				t.Fatalf("Read() error = %v, want %v", err, test.want)
			}
		})
	}
}

func TestDecoderIgnoresAdditiveFieldsOnKnownControlConnectorType(t *testing.T) {
	t.Parallel()

	payload := []byte(`{"body":{"future":true},"chan":"ctrl-c","future_envelope_field":17,"seq":"9","type":"ping","v":1}`)
	decoder := NewDecoder(map[string]TypeSpec{"ping": {}})
	got, err := decoder.Read(bytes.NewReader(framed(payload)))
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if got.Type != "ping" || got.Seq != Counter(9) || string(got.Body) != `{"future":true}` {
		t.Fatalf("Read() = %#v", got)
	}
}

func TestDecoderEnforcesReplyCorrelationAndMonotonicSequence(t *testing.T) {
	t.Parallel()

	decoder := NewDecoder(map[string]TypeSpec{"pong": {Reply: true}})
	missingReply := []byte(`{"body":{},"chan":"ctrl-c","seq":"1","type":"pong","v":1}`)
	if _, err := decoder.Read(bytes.NewReader(framed(missingReply))); !errors.Is(err, ErrUnknownMessage) {
		t.Fatalf("reply without re error = %v, want ErrUnknownMessage", err)
	}
	first := []byte(`{"body":{},"chan":"ctrl-c","re":"1","seq":"2","type":"pong","v":1}`)
	if _, err := decoder.Read(bytes.NewReader(framed(first))); err != nil {
		t.Fatalf("first Read() error = %v", err)
	}
	nonMonotonic := []byte(`{"body":{},"chan":"ctrl-c","re":"1","seq":"2","type":"pong","v":1}`)
	if _, err := decoder.Read(bytes.NewReader(framed(nonMonotonic))); !errors.Is(err, ErrNonMonotonicSequence) {
		t.Fatalf("duplicate sequence error = %v, want ErrNonMonotonicSequence", err)
	}
}

func TestSenderReturnsDeadlineFaultWhenBoundedQueueCannotDrain(t *testing.T) {
	t.Parallel()

	writer := &blockingWriter{started: make(chan struct{}), release: make(chan struct{})}
	sender := NewSender(writer, 1, 25*time.Millisecond)

	envelope := Envelope{Version: 1, Channel: ChannelControlConnector, Type: "ping", Seq: Counter(math.MaxUint64), Body: []byte(`{}`)}
	firstDone := make(chan error, 1)
	go func() { firstDone <- sender.Send(envelope) }()
	<-writer.started

	secondDone := make(chan error, 1)
	go func() { secondDone <- sender.Send(envelope) }()
	time.Sleep(5 * time.Millisecond)

	if err := sender.Send(envelope); !errors.Is(err, ErrSendDeadline) {
		t.Fatalf("Send() error = %v, want ErrSendDeadline", err)
	}
	close(writer.release)
	<-firstDone
	<-secondDone
	sender.Close()
}

func framed(payload []byte) []byte {
	frame := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(frame[:4], uint32(len(payload)))
	copy(frame[4:], payload)
	return frame
}

type blockingWriter struct {
	started chan struct{}
	release chan struct{}
}

func (w *blockingWriter) Write(p []byte) (int, error) {
	select {
	case <-w.started:
	default:
		close(w.started)
	}
	<-w.release
	return len(p), nil
}
