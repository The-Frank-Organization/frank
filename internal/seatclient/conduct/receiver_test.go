package conduct

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
)

type attachSequence struct {
	replies []AttachReply
	calls   int
}

func (sequence *attachSequence) Attach(context.Context, AttachTuple) (AttachReply, error) {
	sequence.calls++
	if len(sequence.replies) == 0 {
		return AttachReply{}, errors.New("empty attach fixture")
	}
	reply := sequence.replies[0]
	sequence.replies = sequence.replies[1:]
	return reply, nil
}

type controlRecorder struct {
	results []string
	wakes   []string
}

func (control *controlRecorder) AttachResult(_ context.Context, _ AttachTuple, result string) error {
	control.results = append(control.results, result)
	return nil
}

func (control *controlRecorder) WakeForward(_ context.Context, relayID string) error {
	control.wakes = append(control.wakes, relayID)
	return nil
}

type receiverTransport struct {
	calls  []string
	pushes [][]byte
}

func (transport *receiverTransport) InvokeWire(_ context.Context, name string, args json.RawMessage) (json.RawMessage, error) {
	transport.calls = append(transport.calls, name+":"+string(args))
	switch name {
	case "project":
		return json.RawMessage(`["relay-a","relay-b"]`), nil
	case "read":
		return json.RawMessage(`{"record":` + string(args) + `}`), nil
	default:
		return nil, errors.New("unexpected wire call")
	}
}

func (*receiverTransport) DescribeWire(context.Context, channel.DescribeRequest) (channel.DescriptionResponse, error) {
	return channel.DescriptionResponse{}, nil
}

func (transport *receiverTransport) ReceivePush(context.Context) ([]byte, error) {
	if len(transport.pushes) == 0 {
		return nil, errors.New("no push")
	}
	push := transport.pushes[0]
	transport.pushes = transport.pushes[1:]
	return push, nil
}

func (*receiverTransport) Shutdown() error { return nil }

func TestReceiverSuspendedBacksOffThenAttachesAndReportsByteExact(t *testing.T) {
	transport := &receiverTransport{}
	attacher := &attachSequence{replies: []AttachReply{
		{Result: AttachSuspended},
		{Result: AttachOK, Transport: transport},
	}}
	control := &controlRecorder{}
	receiver, err := NewReceiver(attacher, control, 100*time.Millisecond, time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	client, err := receiver.Attach(context.Background(), AttachTuple{RunID: "run", GenerationID: "gen", TurnEpoch: 7})
	if err != nil || client == nil || attacher.calls != 2 {
		t.Fatalf("client=%v calls=%d err=%v", client, attacher.calls, err)
	}
	if !reflect.DeepEqual(control.results, []string{AttachSuspended, AttachOK}) {
		t.Fatalf("reported results = %#v", control.results)
	}
}

func TestReceiverTupleMismatchIsTerminalWithoutRetry(t *testing.T) {
	attacher := &attachSequence{replies: []AttachReply{{Result: AttachTupleMismatch}}}
	control := &controlRecorder{}
	receiver, _ := NewReceiver(attacher, control, time.Second, time.Millisecond)
	_, err := receiver.Attach(context.Background(), AttachTuple{})
	if !errors.Is(err, ErrGenerationFenced) || attacher.calls != 1 {
		t.Fatalf("err=%v calls=%d", err, attacher.calls)
	}
}

func TestReattachPerformsDurableProjectReadRediscovery(t *testing.T) {
	transport := &receiverTransport{}
	attacher := &attachSequence{replies: []AttachReply{{Result: AttachOK, Transport: transport}}}
	receiver, _ := NewReceiver(attacher, &controlRecorder{}, time.Second, time.Millisecond)
	records, err := receiver.Reattach(context.Background(), AttachTuple{})
	if err != nil || len(records) != 2 {
		t.Fatalf("records=%s err=%v", records, err)
	}
	want := []string{`project:{}`, `read:{"relay_id":"relay-a"}`, `read:{"relay_id":"relay-b"}`}
	if !reflect.DeepEqual(transport.calls, want) {
		t.Fatalf("calls=%#v want=%#v", transport.calls, want)
	}
}

func TestPushIsAdvisoryAndDuplicatesForwardWithoutLocalLedger(t *testing.T) {
	transport := &receiverTransport{pushes: [][]byte{
		[]byte(`{"kind":"delivery-nudge","relay_id":"relay-a"}`),
		[]byte(`{"kind":"delivery-nudge","relay_id":"relay-a"}`),
	}}
	control := &controlRecorder{}
	receiver, _ := NewReceiver(&attachSequence{replies: []AttachReply{{Result: AttachOK, Transport: transport}}}, control, time.Second, time.Millisecond)
	if _, err := receiver.Attach(context.Background(), AttachTuple{}); err != nil {
		t.Fatal(err)
	}
	if err := receiver.ForwardNextPush(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := receiver.ForwardNextPush(context.Background()); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(control.wakes, []string{"relay-a", "relay-a"}) {
		t.Fatalf("wakes=%#v", control.wakes)
	}
}

func TestBrokerErrorDispositionTable(t *testing.T) {
	tests := []struct {
		message string
		submit  bool
		want    ErrorDisposition
	}{
		{"broker:stale-epoch", false, DispositionFence},
		{"broker:suspended", false, DispositionHold},
		{"broker:preparing", false, DispositionHold},
		{"broker:record-unavailable", false, DispositionReinvoke},
		{"broker:record-unavailable", true, DispositionRediscover},
		{"broker:unknown-outcome", true, DispositionRediscover},
		{"shim:connection-lost", false, DispositionReconnect},
		{"application rejected", true, DispositionTerminal},
	}
	for _, test := range tests {
		if got := DispositionForError(errors.New(test.message), test.submit); got != test.want {
			t.Errorf("%s submit=%v: got %s want %s", test.message, test.submit, got, test.want)
		}
	}
}
