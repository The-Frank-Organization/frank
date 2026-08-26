package appipc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
	"unicode/utf8"
)

type Channel string

const (
	ChannelCtrlW  Channel = "ctrl-w"
	ChannelCtrlC  Channel = "ctrl-c"
	ChannelDataP  Channel = "data-p"
	ChannelBroker Channel = "broker"
)

type FamilyEvolution uint8

const (
	AdditiveFamily FamilyEvolution = iota
	ClosedFamily
)

type MessageSpec struct {
	Channel            Channel
	Reply              bool
	Evolution          FamilyEvolution
	BodyFields         []string
	RequiredBodyFields []string
	Enums              map[string][]string
	NewBody            func() any
	ValidateBody       func(map[string]any) error
}

type Envelope struct {
	V         int
	Channel   Channel
	Type      string
	Seq       string
	Re        *string
	RunID     *string
	TurnEpoch *string
	Body      any
}

var (
	ErrUnknownMessage     = errors.New("appipc: unknown message type")
	ErrUnknownField       = errors.New("appipc: unknown field in closed message family")
	ErrMalformedEnvelope  = errors.New("appipc: malformed envelope")
	ErrDuplicateMessage   = errors.New("appipc: duplicate message registration")
	ErrInvalidMessageBody = errors.New("appipc: invalid message body")
)

type Fault string

const FaultUnknownMessage Fault = "FAULT_UNKNOWN_MESSAGE"

type ProtocolError struct {
	Fault Fault
	Cause error
	Text  string
}

func (err *ProtocolError) Error() string {
	if err.Text == "" {
		return err.Cause.Error()
	}
	return err.Text + ": " + err.Cause.Error()
}

func (err *ProtocolError) Unwrap() error { return err.Cause }

func FaultCode(err error) Fault {
	var protocolError *ProtocolError
	if errors.As(err, &protocolError) {
		return protocolError.Fault
	}
	return ""
}

type messageKey struct {
	channel     Channel
	messageType string
}

type registeredMessage struct {
	reply        bool
	evolution    FamilyEvolution
	bodyFields   map[string]struct{}
	required     map[string]struct{}
	enums        map[string]map[string]struct{}
	newBody      func() any
	validateBody func(map[string]any) error
	requireRunID bool
	requireEpoch bool
}

type Registry struct {
	mu       sync.RWMutex
	messages map[messageKey]registeredMessage
}

func NewRegistry() *Registry {
	return &Registry{messages: make(map[messageKey]registeredMessage)}
}

func (registry *Registry) Register(messageType string, spec MessageSpec) error {
	if registry == nil || messageType == "" || !validChannel(spec.Channel) {
		return fmt.Errorf("%w: invalid registration", ErrMalformedEnvelope)
	}
	if spec.Evolution != AdditiveFamily && spec.Evolution != ClosedFamily {
		return fmt.Errorf("%w: invalid family evolution", ErrMalformedEnvelope)
	}
	fields := make(map[string]struct{}, len(spec.BodyFields))
	for _, field := range spec.BodyFields {
		if field == "" || !utf8.ValidString(field) {
			return fmt.Errorf("%w: invalid body field", ErrMalformedEnvelope)
		}
		if _, duplicate := fields[field]; duplicate {
			return fmt.Errorf("%w: duplicate body field %q", ErrMalformedEnvelope, field)
		}
		fields[field] = struct{}{}
	}
	required := make(map[string]struct{}, len(spec.RequiredBodyFields))
	for _, field := range spec.RequiredBodyFields {
		if _, known := fields[field]; !known {
			return fmt.Errorf("%w: required unknown body field %q", ErrMalformedEnvelope, field)
		}
		required[field] = struct{}{}
	}
	enums := make(map[string]map[string]struct{}, len(spec.Enums))
	for field, values := range spec.Enums {
		if _, known := fields[field]; !known || len(values) == 0 {
			return fmt.Errorf("%w: enum for invalid body field %q", ErrMalformedEnvelope, field)
		}
		set := make(map[string]struct{}, len(values))
		for _, value := range values {
			if value == "" {
				return fmt.Errorf("%w: empty enum token for %q", ErrMalformedEnvelope, field)
			}
			set[value] = struct{}{}
		}
		enums[field] = set
	}
	key := messageKey{channel: spec.Channel, messageType: messageType}
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if _, duplicate := registry.messages[key]; duplicate {
		return fmt.Errorf("%w: %s/%s", ErrDuplicateMessage, spec.Channel, messageType)
	}
	registry.messages[key] = registeredMessage{
		reply:        spec.Reply,
		evolution:    spec.Evolution,
		bodyFields:   fields,
		required:     required,
		enums:        enums,
		newBody:      spec.NewBody,
		validateBody: spec.ValidateBody,
	}
	return nil
}

func (registry *Registry) Has(channel Channel, messageType string) bool {
	if registry == nil {
		return false
	}
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	_, ok := registry.messages[messageKey{channel: channel, messageType: messageType}]
	return ok
}

func (registry *Registry) requireEnvelope(channel Channel, messageType string, runID, turnEpoch bool) error {
	key := messageKey{channel: channel, messageType: messageType}
	registry.mu.Lock()
	defer registry.mu.Unlock()
	spec, ok := registry.messages[key]
	if !ok {
		return fmt.Errorf("%w: %s/%s", ErrUnknownMessage, channel, messageType)
	}
	spec.requireRunID = runID
	spec.requireEpoch = turnEpoch
	registry.messages[key] = spec
	return nil
}

func (registry *Registry) Encode(envelope Envelope) ([]byte, error) {
	spec, err := registry.validateEnvelope(envelope)
	if err != nil {
		return nil, err
	}
	body, err := normalizeBody(envelope.Body)
	if err != nil {
		return nil, err
	}
	body, err = applyEvolution(body, spec)
	if err != nil {
		return nil, err
	}
	if err := validateRegisteredBody(body, spec); err != nil {
		return nil, err
	}
	if err := validateEnvelopeBodyIdentity(envelope, body); err != nil {
		return nil, err
	}
	wire := map[string]any{
		"v":    envelope.V,
		"chan": string(envelope.Channel),
		"type": envelope.Type,
		"seq":  envelope.Seq,
		"body": body,
	}
	if envelope.Re != nil {
		wire["re"] = *envelope.Re
	}
	if envelope.RunID != nil {
		wire["run_id"] = *envelope.RunID
	}
	if envelope.TurnEpoch != nil {
		wire["turn_epoch"] = *envelope.TurnEpoch
	}
	return MarshalJCS(wire)
}

func (registry *Registry) Decode(payload []byte) (Envelope, error) {
	if !utf8.Valid(payload) {
		return Envelope{}, fmt.Errorf("%w: payload is not UTF-8", ErrMalformedEnvelope)
	}
	if err := requireCanonicalEnvelope(payload); err != nil {
		return Envelope{}, err
	}
	type rawEnvelope struct {
		V         int             `json:"v"`
		Channel   Channel         `json:"chan"`
		Type      string          `json:"type"`
		Seq       string          `json:"seq"`
		Re        *string         `json:"re"`
		RunID     *string         `json:"run_id"`
		TurnEpoch *string         `json:"turn_epoch"`
		Body      json.RawMessage `json:"body"`
	}
	var raw rawEnvelope
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&raw); err != nil {
		return Envelope{}, fmt.Errorf("%w: %v", ErrMalformedEnvelope, err)
	}
	if err := requireJSONEOF(decoder); err != nil {
		return Envelope{}, fmt.Errorf("%w: %v", ErrMalformedEnvelope, err)
	}
	envelope := Envelope{
		V:         raw.V,
		Channel:   raw.Channel,
		Type:      raw.Type,
		Seq:       raw.Seq,
		Re:        raw.Re,
		RunID:     raw.RunID,
		TurnEpoch: raw.TurnEpoch,
	}
	spec, err := registry.validateEnvelope(envelope)
	if err != nil {
		return Envelope{}, err
	}
	body, err := decodeBody(raw.Body)
	if err != nil {
		return Envelope{}, err
	}
	body, err = applyEvolution(body, spec)
	if err != nil {
		return Envelope{}, err
	}
	if err := validateRegisteredBody(body, spec); err != nil {
		return Envelope{}, err
	}
	if err := validateEnvelopeBodyIdentity(envelope, body); err != nil {
		return Envelope{}, err
	}
	if spec.newBody != nil {
		typed := spec.newBody()
		if typed == nil {
			return Envelope{}, fmt.Errorf("%w: nil body factory", ErrInvalidMessageBody)
		}
		encoded, err := json.Marshal(body)
		if err != nil {
			return Envelope{}, fmt.Errorf("%w: %v", ErrInvalidMessageBody, err)
		}
		if err := json.Unmarshal(encoded, typed); err != nil {
			return Envelope{}, fmt.Errorf("%w: %v", ErrInvalidMessageBody, err)
		}
		envelope.Body = typed
	} else {
		envelope.Body = body
	}
	return envelope, nil
}

func validateEnvelopeBodyIdentity(envelope Envelope, body map[string]any) error {
	for _, field := range []struct {
		name     string
		envelope *string
	}{
		{name: "run_id", envelope: envelope.RunID},
		{name: "turn_epoch", envelope: envelope.TurnEpoch},
	} {
		if field.envelope == nil {
			continue
		}
		value, present := body[field.name]
		if !present {
			continue
		}
		bodyValue, ok := value.(string)
		if !ok || bodyValue != *field.envelope {
			return fmt.Errorf("%w: envelope/body %s mismatch", ErrInvalidMessageBody, field.name)
		}
	}
	return nil
}

func requireCanonicalEnvelope(payload []byte) error {
	var value any
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.UseNumber()
	if err := decoder.Decode(&value); err != nil {
		return fmt.Errorf("%w: %v", ErrMalformedEnvelope, err)
	}
	if err := requireJSONEOF(decoder); err != nil {
		return fmt.Errorf("%w: %v", ErrMalformedEnvelope, err)
	}
	canonical, err := MarshalJCS(value)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrMalformedEnvelope, err)
	}
	if !bytes.Equal(payload, canonical) {
		return fmt.Errorf("%w: payload is not canonical JSON", ErrMalformedEnvelope)
	}
	return nil
}

func (registry *Registry) validateEnvelope(envelope Envelope) (registeredMessage, error) {
	if envelope.V != 1 || !validChannel(envelope.Channel) || envelope.Type == "" {
		return registeredMessage{}, ErrMalformedEnvelope
	}
	if _, err := ParseCounter(envelope.Seq); err != nil {
		return registeredMessage{}, fmt.Errorf("%w: seq: %v", ErrMalformedEnvelope, err)
	}
	if envelope.Re != nil {
		if _, err := ParseCounter(*envelope.Re); err != nil {
			return registeredMessage{}, fmt.Errorf("%w: re: %v", ErrMalformedEnvelope, err)
		}
	}
	if envelope.TurnEpoch != nil {
		if _, err := ParseCounter(*envelope.TurnEpoch); err != nil {
			return registeredMessage{}, fmt.Errorf("%w: turn_epoch: %v", ErrMalformedEnvelope, err)
		}
	}
	if registry == nil {
		return registeredMessage{}, ErrUnknownMessage
	}
	key := messageKey{channel: envelope.Channel, messageType: envelope.Type}
	registry.mu.RLock()
	spec, ok := registry.messages[key]
	registry.mu.RUnlock()
	if !ok {
		return registeredMessage{}, &ProtocolError{
			Fault: FaultUnknownMessage,
			Cause: ErrUnknownMessage,
			Text:  fmt.Sprintf("appipc: %s/%s", envelope.Channel, envelope.Type),
		}
	}
	if spec.reply && envelope.Re == nil {
		return registeredMessage{}, &ProtocolError{
			Fault: FaultUnknownMessage,
			Cause: ErrMalformedEnvelope,
			Text:  "appipc: reply-class message lacks re",
		}
	}
	if !spec.reply && envelope.Re != nil {
		return registeredMessage{}, &ProtocolError{
			Fault: FaultUnknownMessage,
			Cause: ErrMalformedEnvelope,
			Text:  "appipc: non-reply message carries re",
		}
	}
	if spec.requireRunID && (envelope.RunID == nil || *envelope.RunID == "") {
		return registeredMessage{}, fmt.Errorf("%w: run_id is required for %s", ErrInvalidMessageBody, envelope.Type)
	}
	if spec.requireEpoch && envelope.TurnEpoch == nil {
		return registeredMessage{}, fmt.Errorf("%w: turn_epoch is required for %s", ErrInvalidMessageBody, envelope.Type)
	}
	return spec, nil
}

func normalizeBody(body any) (map[string]any, error) {
	if body == nil {
		return nil, fmt.Errorf("%w: body must be an object", ErrMalformedEnvelope)
	}
	encoded, err := MarshalJCS(body)
	if err != nil {
		return nil, fmt.Errorf("%w: body: %v", ErrMalformedEnvelope, err)
	}
	return decodeBody(encoded)
}

func decodeBody(raw []byte) (map[string]any, error) {
	if len(raw) == 0 {
		return nil, fmt.Errorf("%w: body is required", ErrMalformedEnvelope)
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, fmt.Errorf("%w: body: %v", ErrMalformedEnvelope, err)
	}
	if err := requireJSONEOF(decoder); err != nil {
		return nil, fmt.Errorf("%w: body: %v", ErrMalformedEnvelope, err)
	}
	body, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("%w: body must be an object", ErrMalformedEnvelope)
	}
	return body, nil
}

func applyEvolution(body map[string]any, spec registeredMessage) (map[string]any, error) {
	for field := range body {
		if _, known := spec.bodyFields[field]; known {
			continue
		}
		if spec.evolution == ClosedFamily {
			return nil, &ProtocolError{
				Fault: FaultUnknownMessage,
				Cause: ErrUnknownField,
				Text:  fmt.Sprintf("appipc: closed family field %q", field),
			}
		}
		delete(body, field)
	}
	return body, nil
}

func validateRegisteredBody(body map[string]any, spec registeredMessage) error {
	for field := range spec.required {
		value, present := body[field]
		if !present || value == nil {
			return fmt.Errorf("%w: required field %q is absent", ErrInvalidMessageBody, field)
		}
		if text, ok := value.(string); ok && text == "" {
			return fmt.Errorf("%w: required field %q is empty", ErrInvalidMessageBody, field)
		}
	}
	for field, allowed := range spec.enums {
		value, present := body[field]
		if !present {
			continue
		}
		text, ok := value.(string)
		if !ok {
			return fmt.Errorf("%w: enum field %q is not a string", ErrInvalidMessageBody, field)
		}
		if _, ok := allowed[text]; !ok {
			return fmt.Errorf("%w: unknown %s token %q", ErrInvalidMessageBody, field, text)
		}
	}
	if spec.validateBody != nil {
		if err := spec.validateBody(body); err != nil {
			return fmt.Errorf("%w: %v", ErrInvalidMessageBody, err)
		}
	}
	return nil
}

func requireJSONEOF(decoder *json.Decoder) error {
	var extra any
	err := decoder.Decode(&extra)
	if errors.Is(err, io.EOF) {
		return nil
	}
	if err == nil {
		return errors.New("multiple JSON values")
	}
	return err
}

func validChannel(channel Channel) bool {
	switch channel {
	case ChannelCtrlW, ChannelCtrlC, ChannelDataP, ChannelBroker:
		return true
	default:
		return false
	}
}
