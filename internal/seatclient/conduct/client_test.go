package conduct

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/jackli/frank/internal/channel"
)

func TestClientExposesOnlyCanonicalRelayVerbs(t *testing.T) {
	transport := &fakeTransport{}
	client, err := New(transport)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name string
		wire string
	}{
		{"relay.submit", "submit"},
		{"relay.project", "project"},
		{"relay.read", "read"},
	} {
		if _, err := client.Call(context.Background(), test.name, json.RawMessage(`{}`)); err != nil {
			t.Fatal(err)
		}
		if got := transport.calls[len(transport.calls)-1]; got != test.wire {
			t.Fatalf("wire verb = %q, want %q", got, test.wire)
		}
	}
	if _, err := client.Call(context.Background(), "submit", nil); !errors.Is(err, ErrUnknownVerb) {
		t.Fatalf("noncanonical verb error = %v", err)
	}
}

func TestDescribePushAndCloseDelegateWithoutStoreSurface(t *testing.T) {
	transport := &fakeTransport{description: channel.DescriptionResponse{Tools: []string{"submit"}}, push: []byte("wake")}
	client, err := New(transport)
	if err != nil {
		t.Fatal(err)
	}
	description, err := client.Describe(context.Background(), channel.DescribeRequest{Phase: "PLAN", Tier: "large"})
	if err != nil || !reflect.DeepEqual(description.Tools, []string{"submit"}) {
		t.Fatalf("description = %#v, %v", description, err)
	}
	push, err := client.NextPush(context.Background())
	if err != nil || string(push) != "wake" {
		t.Fatalf("push = %q, %v", push, err)
	}
	if err := client.Close(); err != nil || !transport.closed {
		t.Fatalf("close = %v, closed %v", err, transport.closed)
	}
}

type fakeTransport struct {
	calls       []string
	description channel.DescriptionResponse
	push        []byte
	closed      bool
}

func (transport *fakeTransport) InvokeWire(_ context.Context, name string, _ json.RawMessage) (json.RawMessage, error) {
	transport.calls = append(transport.calls, name)
	return json.RawMessage(`{"ok":true}`), nil
}

func (transport *fakeTransport) DescribeWire(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return transport.description, nil
}

func (transport *fakeTransport) ReceivePush(context.Context) ([]byte, error) {
	return append([]byte(nil), transport.push...), nil
}

func (transport *fakeTransport) Shutdown() error {
	transport.closed = true
	return nil
}
