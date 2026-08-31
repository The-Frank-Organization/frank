package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/worker/journal"
)

type fakeGate struct {
	opens, ends int
	lastEnd     StreamEnd
}

func (gate *fakeGate) AttemptOpen(context.Context, Request) error { gate.opens++; return nil }
func (gate *fakeGate) RecordStreamEnd(_ context.Context, _ string, end StreamEnd) error {
	gate.ends++
	gate.lastEnd = end
	return nil
}

type fakeConnector struct {
	disposition       Disposition
	items             []json.RawMessage
	attempts, cancels int
	err               error
}

func (connector *fakeConnector) Attempt(context.Context, Request) (Disposition, []json.RawMessage, error) {
	connector.attempts++
	return connector.disposition, connector.items, connector.err
}
func (connector *fakeConnector) Cancel(context.Context, string, string) (Disposition, error) {
	connector.cancels++
	return connector.disposition, connector.err
}

func TestAttemptOrderingTotalTerminalMappingAndNoRetry(t *testing.T) {
	for _, test := range []struct {
		d    Disposition
		end  *StreamEnd
		wire bool
	}{
		{Completed, ptr(StreamCompleted), true}, {TransportFailed, ptr(StreamFailed), true}, {CancelledPost, ptr(StreamCancelled), true}, {StreamLost, ptr(StreamLostEnd), true},
		{EgressDenied, nil, false}, {RejectedLocal, nil, false}, {CancelledPre, nil, false}, {StaleEpoch, nil, false}, {EpochAhead, nil, false},
	} {
		t.Run(string(test.d), func(t *testing.T) {
			g := &fakeGate{}
			c := &fakeConnector{disposition: test.d}
			cycle, _ := New(g, c, "7")
			out, err := cycle.Run(context.Background(), request("7"))
			if err != nil {
				t.Fatal(err)
			}
			if g.opens != 1 || c.attempts != 1 {
				t.Fatalf("ordering/count opens=%d attempts=%d", g.opens, c.attempts)
			}
			if out.WireStarted != test.wire || !equalEnd(out.StreamEnd, test.end) {
				t.Fatalf("out=%+v", out)
			}
			if test.end == nil && g.ends != 0 || test.end != nil && g.ends != 1 {
				t.Fatalf("end calls=%d", g.ends)
			}
		})
	}
}

func TestEpochRejectIsAttemptInert(t *testing.T) {
	g := &fakeGate{}
	c := &fakeConnector{disposition: Completed}
	cycle, _ := New(g, c, "8")
	out, err := cycle.Run(context.Background(), request("7"))
	if err != nil || out.Disposition != StaleEpoch || g.opens != 0 || c.attempts != 0 {
		t.Fatalf("out=%+v opens=%d attempts=%d err=%v", out, g.opens, c.attempts, err)
	}
}

func TestBothCancellationCuts(t *testing.T) {
	for _, d := range []Disposition{CancelledPre, CancelledPost} {
		g := &fakeGate{}
		c := &fakeConnector{disposition: d}
		cycle, _ := New(g, c, "7")
		out, err := cycle.Cancel(context.Background(), "attempt", "7")
		if err != nil {
			t.Fatal(err)
		}
		if d == CancelledPre && (out.WireStarted || g.ends != 0) {
			t.Fatalf("pre=%+v ends=%d", out, g.ends)
		}
		if d == CancelledPost && (!out.WireStarted || g.lastEnd != StreamCancelled) {
			t.Fatalf("post=%+v end=%s", out, g.lastEnd)
		}
	}
}

func TestOpaqueResponseItemRoundTripsThroughJournalBytePreserved(t *testing.T) {
	opaque := json.RawMessage(`{"type":"reasoning","provider_native":{"encrypted":"opaque_item_passthrough"}}`)
	g := &fakeGate{}
	c := &fakeConnector{disposition: Completed, items: []json.RawMessage{opaque}}
	cycle, _ := New(g, c, "7")
	out, err := cycle.Run(context.Background(), request("7"))
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Events) != 1 || !bytes.Equal(out.Events[0].Opaque, opaque) {
		t.Fatalf("event=%+v", out.Events)
	}
	record := journal.Record{Seq: "1", Kind: journal.KindProviderOutput, GenerationID: "gen", TurnID: "turn", RoundIndex: "0", TSMonotonic: "1", Fields: map[string]json.RawMessage{"attempt_id": json.RawMessage(`"attempt"`), "item_index": json.RawMessage(`"0"`)}}
	for member, raw := range journal.ProviderItemCarrier(out.Events[0].Opaque) {
		record.Fields[member] = raw
	}
	finalized, err := journal.FinalizeRecord(record)
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := journal.MarshalRecord(finalized)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := journal.DecodeRecord(encoded)
	if err != nil {
		t.Fatal(err)
	}
	restored, err := journal.ProviderItemBytes(decoded)
	if err != nil || !bytes.Equal(restored, opaque) {
		t.Fatalf("opaque changed: %s err=%v", restored, err)
	}
}

func TestConnectorFaultDoesNotAutoRetry(t *testing.T) {
	g := &fakeGate{}
	c := &fakeConnector{err: errors.New("lost")}
	cycle, _ := New(g, c, "7")
	out, err := cycle.Run(context.Background(), request("7"))
	if err == nil || out.Disposition != "" || len(out.Events) != 0 || c.attempts != 1 {
		t.Fatalf("out=%+v attempts=%d err=%v", out, c.attempts, err)
	}
}

func request(epoch string) Request {
	return Request{AttemptID: "attempt", TurnID: "turn", TurnEpoch: epoch, ProviderLane: "lane", OpaqueRequest: json.RawMessage(`{"opaque":true}`)}
}
func ptr(value StreamEnd) *StreamEnd { return &value }
func equalEnd(left, right *StreamEnd) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}
