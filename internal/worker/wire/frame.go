package wire

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/jackli/frank/internal/worker/jcs"
)

const (
	FrameVersion    = 1
	MaxFramePayload = 4 * 1024 * 1024
	ParkedRowMax    = 640
)

var (
	ErrFrameTooLarge     = errors.New("frame payload exceeds FRAME_MAX")
	ErrParkedRowTooLarge = errors.New("parked row exceeds PARKED_ROW_MAX")
)

// Channel is a governed app-side IPC channel identifier.
type Channel string

const (
	ChannelCTRLW  Channel = "ctrl-w"
	ChannelCTRLC  Channel = "ctrl-c"
	ChannelDATAP  Channel = "data-p"
	ChannelBroker Channel = "broker"
)

// Frame is the version-1 envelope shared by CTRL-W, CTRL-C, DATA-P, and the
// broker feed. Trust-bearing counters remain canonical decimal strings.
type Frame struct {
	Version   int             `json:"v"`
	Channel   Channel         `json:"chan"`
	Type      string          `json:"type"`
	Seq       string          `json:"seq"`
	ReplyTo   string          `json:"re,omitempty"`
	RunID     string          `json:"run_id,omitempty"`
	TurnEpoch string          `json:"turn_epoch,omitempty"`
	Body      json.RawMessage `json:"body"`
}

// MessageSpec pins a known message family. Additive body evolution is the
// default; ClosedBody restricts object members to BodyMembers.
type MessageSpec struct {
	Channel     Channel
	Type        string
	Reply       bool
	ClosedBody  bool
	BodyMembers []string
}

// Codec validates frames against a fixed message-family registry.
type Codec struct {
	specs map[messageKey]compiledSpec
}

type messageKey struct {
	channel Channel
	typeID  string
}

type compiledSpec struct {
	reply      bool
	closedBody bool
	bodyFields map[string]struct{}
}

// NewCodec creates a version-pinned codec. Duplicate or malformed specs fail.
func NewCodec(specs []MessageSpec) (*Codec, error) {
	codec := &Codec{specs: make(map[messageKey]compiledSpec, len(specs))}
	for _, spec := range specs {
		if !validChannel(spec.Channel) {
			return nil, fmt.Errorf("message type %q has unknown channel %q", spec.Type, spec.Channel)
		}
		if spec.Type == "" {
			return nil, errors.New("message type is empty")
		}
		key := messageKey{channel: spec.Channel, typeID: spec.Type}
		if _, exists := codec.specs[key]; exists {
			return nil, fmt.Errorf("duplicate message spec %q on %q", spec.Type, spec.Channel)
		}
		compiled := compiledSpec{reply: spec.Reply, closedBody: spec.ClosedBody}
		if spec.ClosedBody {
			compiled.bodyFields = make(map[string]struct{}, len(spec.BodyMembers))
			for _, member := range spec.BodyMembers {
				if member == "" {
					return nil, fmt.Errorf("message type %q has an empty body member", spec.Type)
				}
				if _, exists := compiled.bodyFields[member]; exists {
					return nil, fmt.Errorf("message type %q repeats body member %q", spec.Type, member)
				}
				compiled.bodyFields[member] = struct{}{}
			}
		}
		codec.specs[key] = compiled
	}
	return codec, nil
}

// Encode returns uint32-big-endian length-prefixed canonical JSON bytes.
func (codec *Codec) Encode(frame Frame) ([]byte, error) {
	if err := codec.validateFrame(frame); err != nil {
		return nil, err
	}
	encoded, err := json.Marshal(frame)
	if err != nil {
		return nil, fmt.Errorf("marshal frame: %w", err)
	}
	payload, err := jcs.Canonicalize(encoded)
	if err != nil {
		return nil, fmt.Errorf("canonicalize frame: %w", err)
	}
	if len(payload) > MaxFramePayload {
		return nil, ErrFrameTooLarge
	}
	result := make([]byte, 4+len(payload))
	binary.BigEndian.PutUint32(result[:4], uint32(len(payload)))
	copy(result[4:], payload)
	return result, nil
}

// Decode validates and decodes exactly one complete length-prefixed frame.
func (codec *Codec) Decode(input []byte) (Frame, error) {
	if len(input) < 4 {
		return Frame{}, errors.New("frame is shorter than length prefix")
	}
	length := binary.BigEndian.Uint32(input[:4])
	if length > MaxFramePayload {
		return Frame{}, ErrFrameTooLarge
	}
	if uint64(length)+4 != uint64(len(input)) {
		return Frame{}, fmt.Errorf("frame length prefix %d does not match %d payload bytes", length, len(input)-4)
	}
	payload := input[4:]
	canonical, err := jcs.Canonicalize(payload)
	if err != nil {
		return Frame{}, fmt.Errorf("invalid frame JSON: %w", err)
	}
	if !bytes.Equal(payload, canonical) {
		return Frame{}, errors.New("frame payload is not canonical JSON")
	}

	var frame Frame
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&frame); err != nil {
		return Frame{}, fmt.Errorf("decode frame envelope: %w", err)
	}
	if err := codec.validateFrame(frame); err != nil {
		return Frame{}, err
	}
	return frame, nil
}

// Read reads one frame from a stream without allocating an oversized payload.
func (codec *Codec) Read(reader io.Reader) (Frame, error) {
	prefix := make([]byte, 4)
	if _, err := io.ReadFull(reader, prefix); err != nil {
		return Frame{}, err
	}
	length := binary.BigEndian.Uint32(prefix)
	if length > MaxFramePayload {
		return Frame{}, ErrFrameTooLarge
	}
	input := make([]byte, 4+int(length))
	copy(input, prefix)
	if _, err := io.ReadFull(reader, input[4:]); err != nil {
		return Frame{}, err
	}
	return codec.Decode(input)
}

// Write writes one complete frame, retrying short writes through io.Copy.
func (codec *Codec) Write(writer io.Writer, frame Frame) error {
	encoded, err := codec.Encode(frame)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, bytes.NewReader(encoded))
	return err
}

func (codec *Codec) validateFrame(frame Frame) error {
	if frame.Version != FrameVersion {
		return fmt.Errorf("unsupported frame version %d", frame.Version)
	}
	if !validChannel(frame.Channel) {
		return fmt.Errorf("unknown frame channel %q", frame.Channel)
	}
	spec, exists := codec.specs[messageKey{channel: frame.Channel, typeID: frame.Type}]
	if !exists {
		return fmt.Errorf("unknown message type %q on %q", frame.Type, frame.Channel)
	}
	if _, err := ParseCounter(frame.Seq); err != nil {
		return fmt.Errorf("malformed seq: %w", err)
	}
	if err := validateReplyCorrelation(frame, spec); err != nil {
		return err
	}
	if frame.TurnEpoch != "" {
		if _, err := ParseCounter(frame.TurnEpoch); err != nil {
			return fmt.Errorf("malformed turn_epoch: %w", err)
		}
	}
	if len(frame.Body) == 0 {
		return errors.New("frame body is absent")
	}
	canonicalBody, err := jcs.Canonicalize(frame.Body)
	if err != nil {
		return fmt.Errorf("invalid frame body: %w", err)
	}
	if len(canonicalBody) == 0 || canonicalBody[0] != '{' {
		return errors.New("frame body must be an object")
	}
	if spec.closedBody {
		var body map[string]json.RawMessage
		if err := json.Unmarshal(canonicalBody, &body); err != nil {
			return fmt.Errorf("decode closed frame body: %w", err)
		}
		for member := range body {
			if _, allowed := spec.bodyFields[member]; !allowed {
				return fmt.Errorf("unknown member %q in closed body for %q", member, frame.Type)
			}
		}
	}
	return nil
}

func validateReplyCorrelation(frame Frame, spec compiledSpec) error {
	if spec.reply {
		if frame.ReplyTo == "" {
			return fmt.Errorf("reply-class message %q lacks re", frame.Type)
		}
		if _, err := ParseCounter(frame.ReplyTo); err != nil {
			return fmt.Errorf("malformed re: %w", err)
		}
		return nil
	}
	if frame.ReplyTo != "" {
		return fmt.Errorf("non-reply message %q carries re", frame.Type)
	}
	return nil
}

func validChannel(channel Channel) bool {
	switch channel {
	case ChannelCTRLW, ChannelCTRLC, ChannelDATAP, ChannelBroker:
		return true
	default:
		return false
	}
}

// ParkedUnknown is the closed six-member disclosure row carried by turn_open
// and attempt_open_ok.
type ParkedUnknown struct {
	TurnID              string `json:"turn_id"`
	ToolCallID          string `json:"tool_call_id"`
	TicketID            string `json:"ticket_id"`
	State               string `json:"state"`
	CanonicalToolName   string `json:"canonical_tool_name"`
	CanonicalArgsDigest string `json:"canonical_args_digest"`
}

const (
	ParkedUnknownToolOutcome = "UNKNOWN_TOOL_OUTCOME"
	ParkedPartialToolEffect  = "PARTIAL_TOOL_EFFECT"
)

// EncodeParkedUnknown canonicalizes a closed disclosure row and enforces its
// production 640-byte contribution bound.
func EncodeParkedUnknown(row ParkedUnknown) ([]byte, error) {
	if err := validateParkedUnknown(row); err != nil {
		return nil, err
	}
	raw, err := json.Marshal(row)
	if err != nil {
		return nil, fmt.Errorf("marshal parked row: %w", err)
	}
	encoded, err := jcs.Canonicalize(raw)
	if err != nil {
		return nil, fmt.Errorf("canonicalize parked row: %w", err)
	}
	if len(encoded) > ParkedRowMax {
		return nil, ErrParkedRowTooLarge
	}
	return encoded, nil
}

// DecodeParkedUnknown validates the row's canonical bytes and closed shape.
func DecodeParkedUnknown(input []byte) (ParkedUnknown, error) {
	if len(input) > ParkedRowMax {
		return ParkedUnknown{}, ErrParkedRowTooLarge
	}
	canonical, err := jcs.Canonicalize(input)
	if err != nil {
		return ParkedUnknown{}, fmt.Errorf("invalid parked row JSON: %w", err)
	}
	if !bytes.Equal(input, canonical) {
		return ParkedUnknown{}, errors.New("parked row is not canonical JSON")
	}
	var row ParkedUnknown
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&row); err != nil {
		return ParkedUnknown{}, fmt.Errorf("decode parked row: %w", err)
	}
	if err := validateParkedUnknown(row); err != nil {
		return ParkedUnknown{}, err
	}
	return row, nil
}

func validateParkedUnknown(row ParkedUnknown) error {
	if row.TurnID == "" || row.ToolCallID == "" || row.TicketID == "" || row.CanonicalToolName == "" {
		return errors.New("parked row identity member is empty")
	}
	if row.State != ParkedUnknownToolOutcome && row.State != ParkedPartialToolEffect {
		return fmt.Errorf("unknown parked state %q", row.State)
	}
	if !lowerHexSHA256(row.CanonicalArgsDigest) {
		return errors.New("parked row canonical_args_digest is not lowercase SHA-256")
	}
	return nil
}

func lowerHexSHA256(value string) bool {
	if len(value) != 64 {
		return false
	}
	for _, character := range value {
		if character >= '0' && character <= '9' {
			continue
		}
		if character < 'a' || character > 'f' {
			return false
		}
	}
	return true
}
