package appipc

import (
	"errors"
	"testing"
)

func TestEnvelopeRegistryRejectsUnknownMessageWithPinnedFault(t *testing.T) {
	registry := NewRegistry()
	payload := []byte(`{"v":1,"chan":"ctrl-w","type":"not_registered","seq":"0","body":{}}`)
	if _, err := registry.Decode(payload); !errors.Is(err, ErrUnknownMessage) {
		t.Fatalf("Decode unknown type error = %v, want ErrUnknownMessage", err)
	} else if got := FaultCode(err); got != FaultUnknownMessage {
		t.Fatalf("Decode unknown type fault = %q, want %q", got, FaultUnknownMessage)
	}
}

func TestEnvelopeReplyClassRequiresCanonicalRe(t *testing.T) {
	registry := testEnvelopeRegistry(t, AdditiveFamily, true)
	withoutRe := []byte(`{"v":1,"chan":"ctrl-w","type":"sample","seq":"2","body":{"known":"ok"}}`)
	if _, err := registry.Decode(withoutRe); !errors.Is(err, ErrMalformedEnvelope) {
		t.Fatalf("Decode reply without re error = %v, want ErrMalformedEnvelope", err)
	} else if got := FaultCode(err); got != FaultUnknownMessage {
		t.Fatalf("Decode reply without re fault = %q, want %q", got, FaultUnknownMessage)
	}

	badRe := []byte(`{"v":1,"chan":"ctrl-w","type":"sample","seq":"2","re":"01","body":{"known":"ok"}}`)
	if _, err := registry.Decode(badRe); !errors.Is(err, ErrMalformedEnvelope) {
		t.Fatalf("Decode reply with noncanonical re error = %v, want ErrMalformedEnvelope", err)
	}

	valid := []byte(`{"v":1,"chan":"ctrl-w","type":"sample","seq":"2","re":"1","turn_epoch":"0","body":{"known":"ok"}}`)
	if _, err := registry.Decode(valid); err != nil {
		t.Fatalf("Decode valid reply: %v", err)
	}
}

func TestEnvelopeFamilyEvolutionIsAdditiveOrClosed(t *testing.T) {
	payload := []byte(`{"v":1,"chan":"ctrl-w","type":"sample","seq":"0","body":{"known":"kept","future":"ignored"}}`)

	additive := testEnvelopeRegistry(t, AdditiveFamily, false)
	envelope, err := additive.Decode(payload)
	if err != nil {
		t.Fatalf("additive Decode: %v", err)
	}
	body, ok := envelope.Body.(map[string]any)
	if !ok {
		t.Fatalf("decoded body type = %T, want map[string]any", envelope.Body)
	}
	if got := body["known"]; got != "kept" {
		t.Fatalf("known field = %#v, want kept", got)
	}
	if _, present := body["future"]; present {
		t.Fatalf("additive family retained ignored future field")
	}

	closed := testEnvelopeRegistry(t, ClosedFamily, false)
	if _, err := closed.Decode(payload); !errors.Is(err, ErrUnknownField) {
		t.Fatalf("closed Decode error = %v, want ErrUnknownField", err)
	}
}

func TestEnvelopeEncodeIsCanonicalAndRegistryScopedByChannel(t *testing.T) {
	registry := NewRegistry()
	for _, channel := range []Channel{ChannelCtrlW, ChannelCtrlC} {
		if err := registry.Register("hello", MessageSpec{
			Channel:    channel,
			Evolution:  AdditiveFamily,
			BodyFields: []string{"z", "a"},
		}); err != nil {
			t.Fatalf("Register hello on %s: %v", channel, err)
		}
	}
	runID := "run-1"
	payload, err := registry.Encode(Envelope{
		V:       1,
		Channel: ChannelCtrlC,
		Type:    "hello",
		Seq:     "0",
		RunID:   &runID,
		Body:    map[string]any{"z": "last", "a": "first"},
	})
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	want := `{"body":{"a":"first","z":"last"},"chan":"ctrl-c","run_id":"run-1","seq":"0","type":"hello","v":1}`
	if string(payload) != want {
		t.Fatalf("encoded envelope = %s, want %s", payload, want)
	}
	decoded, err := registry.Decode(payload)
	if err != nil {
		t.Fatalf("Decode(Encode): %v", err)
	}
	if decoded.Channel != ChannelCtrlC || decoded.Type != "hello" || decoded.Seq != "0" {
		t.Fatalf("decoded envelope = %#v", decoded)
	}
}

func TestEnvelopeRejectsInvalidVersionChannelAndCounter(t *testing.T) {
	registry := testEnvelopeRegistry(t, AdditiveFamily, false)
	for _, payload := range []string{
		`{"v":2,"chan":"ctrl-w","type":"sample","seq":"0","body":{"known":"ok"}}`,
		`{"v":1,"chan":"sideband","type":"sample","seq":"0","body":{"known":"ok"}}`,
		`{"v":1,"chan":"ctrl-w","type":"sample","seq":"00","body":{"known":"ok"}}`,
		`{"v":1,"chan":"ctrl-w","type":"sample","seq":"0","body":[]}`,
	} {
		if _, err := registry.Decode([]byte(payload)); !errors.Is(err, ErrMalformedEnvelope) {
			t.Fatalf("Decode(%s) error = %v, want ErrMalformedEnvelope", payload, err)
		}
	}
}

func testEnvelopeRegistry(t *testing.T, evolution FamilyEvolution, reply bool) *Registry {
	t.Helper()
	registry := NewRegistry()
	if err := registry.Register("sample", MessageSpec{
		Channel:    ChannelCtrlW,
		Reply:      reply,
		Evolution:  evolution,
		BodyFields: []string{"known"},
	}); err != nil {
		t.Fatalf("Register sample: %v", err)
	}
	return registry
}
