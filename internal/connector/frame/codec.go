// Package frame implements the version-one app IPC frame shared by the
// connector's CTRL-C and DATA-P channels.
package frame

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/jackli/frank/internal/connector/jcs"
)

const (
	// FrameMax is the maximum payload size, excluding the uint32 length prefix.
	FrameMax uint32 = 4 * 1024 * 1024
)

var (
	ErrFrameOverflow        = errors.New("frame: payload exceeds FRAME_MAX")
	ErrMalformedFrame       = errors.New("frame: malformed frame")
	ErrUnknownMessage       = errors.New("FAULT_UNKNOWN_MESSAGE")
	ErrNonMonotonicSequence = errors.New("frame: non-monotonic sequence")
)

type Channel string

const (
	ChannelControlWorker    Channel = "ctrl-w"
	ChannelControlConnector Channel = "ctrl-c"
	ChannelDataProvider     Channel = "data-p"
	ChannelBroker           Channel = "broker"
)

// Envelope is the version-one wire envelope. Optional counters remain
// pointers so absence cannot be confused with the valid value zero.
type Envelope struct {
	Version   int             `json:"v"`
	Channel   Channel         `json:"chan"`
	Type      string          `json:"type"`
	Seq       Counter         `json:"seq"`
	Re        *Counter        `json:"re,omitempty"`
	RunID     string          `json:"run_id,omitempty"`
	TurnEpoch *Counter        `json:"turn_epoch,omitempty"`
	Body      json.RawMessage `json:"body"`
}

// TypeSpec records the envelope-level rules for a known message type.
type TypeSpec struct {
	Reply bool
}

// Encode returns uint32_be(length) followed by the canonical JSON envelope.
func Encode(envelope Envelope) ([]byte, error) {
	if err := validateEnvelope(envelope); err != nil {
		return nil, err
	}
	raw, err := json.Marshal(envelope)
	if err != nil {
		return nil, fmt.Errorf("%w: encode envelope: %v", ErrMalformedFrame, err)
	}
	payload, err := jcs.Canonicalize(raw)
	if err != nil {
		return nil, fmt.Errorf("%w: canonicalize envelope: %v", ErrMalformedFrame, err)
	}
	if len(payload) > int(FrameMax) {
		return nil, ErrFrameOverflow
	}
	encoded := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(encoded[:4], uint32(len(payload)))
	copy(encoded[4:], payload)
	return encoded, nil
}

// Decoder validates one sender's envelope stream and tracks sequence order per
// channel. Sequence gaps are accepted; duplicates and regressions are not.
type Decoder struct {
	known map[string]TypeSpec
	last  map[Channel]Counter
	seen  map[Channel]bool
}

func NewDecoder(known map[string]TypeSpec) *Decoder {
	copyOfKnown := make(map[string]TypeSpec, len(known))
	for name, spec := range known {
		copyOfKnown[name] = spec
	}
	return &Decoder{
		known: copyOfKnown,
		last:  make(map[Channel]Counter),
		seen:  make(map[Channel]bool),
	}
}

// Read reads one complete frame. An oversize prefix is rejected before any
// payload allocation or read.
func (d *Decoder) Read(reader io.Reader) (Envelope, error) {
	var prefix [4]byte
	if _, err := io.ReadFull(reader, prefix[:]); err != nil {
		return Envelope{}, fmt.Errorf("%w: read length: %w", ErrMalformedFrame, err)
	}
	length := binary.BigEndian.Uint32(prefix[:])
	if length > FrameMax {
		return Envelope{}, ErrFrameOverflow
	}
	payload := make([]byte, int(length))
	if _, err := io.ReadFull(reader, payload); err != nil {
		return Envelope{}, fmt.Errorf("%w: read payload: %w", ErrMalformedFrame, err)
	}
	return d.Decode(payload)
}

// Decode validates one payload. Unknown envelope fields are deliberately
// ignored for the additive CTRL-C/CTRL-W/DATA-P families.
func (d *Decoder) Decode(payload []byte) (Envelope, error) {
	type wireEnvelope struct {
		Version   *int            `json:"v"`
		Channel   *Channel        `json:"chan"`
		Type      *string         `json:"type"`
		Seq       *Counter        `json:"seq"`
		Re        *Counter        `json:"re"`
		RunID     string          `json:"run_id"`
		TurnEpoch *Counter        `json:"turn_epoch"`
		Body      json.RawMessage `json:"body"`
	}
	var wire wireEnvelope
	if err := json.Unmarshal(payload, &wire); err != nil {
		return Envelope{}, fmt.Errorf("%w: decode envelope: %v", ErrMalformedFrame, err)
	}
	if wire.Version == nil || wire.Channel == nil || wire.Type == nil || wire.Seq == nil || wire.Body == nil {
		return Envelope{}, fmt.Errorf("%w: missing required envelope member", ErrMalformedFrame)
	}
	envelope := Envelope{
		Version:   *wire.Version,
		Channel:   *wire.Channel,
		Type:      *wire.Type,
		Seq:       *wire.Seq,
		Re:        wire.Re,
		RunID:     wire.RunID,
		TurnEpoch: wire.TurnEpoch,
		Body:      wire.Body,
	}
	if err := validateEnvelope(envelope); err != nil {
		return Envelope{}, err
	}
	spec, ok := d.known[envelope.Type]
	if !ok {
		return Envelope{}, fmt.Errorf("%w: %s", ErrUnknownMessage, envelope.Type)
	}
	if spec.Reply && envelope.Re == nil {
		return Envelope{}, fmt.Errorf("%w: reply %s lacks re", ErrUnknownMessage, envelope.Type)
	}
	if d.seen[envelope.Channel] && envelope.Seq <= d.last[envelope.Channel] {
		return Envelope{}, fmt.Errorf("%w: channel %s got %s after %s", ErrNonMonotonicSequence, envelope.Channel, envelope.Seq, d.last[envelope.Channel])
	}
	d.seen[envelope.Channel] = true
	d.last[envelope.Channel] = envelope.Seq
	return envelope, nil
}

func validateEnvelope(envelope Envelope) error {
	if envelope.Version != 1 {
		return fmt.Errorf("%w: version must be 1", ErrMalformedFrame)
	}
	switch envelope.Channel {
	case ChannelControlWorker, ChannelControlConnector, ChannelDataProvider, ChannelBroker:
	default:
		return fmt.Errorf("%w: unknown channel %q", ErrMalformedFrame, envelope.Channel)
	}
	if envelope.Type == "" {
		return fmt.Errorf("%w: empty message type", ErrMalformedFrame)
	}
	if len(envelope.Body) == 0 || !json.Valid(envelope.Body) {
		return fmt.Errorf("%w: body is not valid JSON", ErrMalformedFrame)
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(envelope.Body, &object); err != nil || object == nil {
		return fmt.Errorf("%w: body must be a JSON object", ErrMalformedFrame)
	}
	return nil
}
