package appipc

import (
	"errors"
	"reflect"
	"testing"
)

func TestProtocolRegistryContainsEveryFrozenMessageFamily(t *testing.T) {
	registry, err := NewProtocolRegistry()
	if err != nil {
		t.Fatalf("NewProtocolRegistry: %v", err)
	}
	want := map[Channel][]string{
		ChannelCtrlW: {
			"hello", "assign", "attach_result", "turn_open", "genesis_committed", "turn_terminal", "turn_cancel_ack",
			"turn_receipt", "turn_reject", "attempt_open", "attempt_open_ok", "attempt_open_reject",
			"attempt_stream_end", "app_event", "wake_forward", "authorize_tool_call", "ticket_granted",
			"authorize_reject", "denied_above_set", "duplicate_request", "stale_epoch", "identity_mismatch",
			"consume_ticket", "consume_ok", "duplicate_consume", "record_tool_outcome", "content_ready",
			"report_resume_disposition", "receipt", "disposition_conflict", "ping", "pong", "shutdown",
			"admission_refused",
		},
		ChannelCtrlC:  {"hello", "connector_assign", "connector_ready", "epoch_update", "attempt_result", "ping", "pong", "shutdown"},
		ChannelBroker: {"state_proposal", "state_proposal_result", "epoch_state", "boundary_cut", "epoch_installed", "broker_event_ack"},
	}
	for channel, names := range want {
		for _, name := range names {
			if !registry.Has(channel, name) {
				t.Errorf("protocol registry lacks %s/%s", channel, name)
			}
		}
	}
}

func TestProtocolRegistryTypedRoundTripsAcrossFamilies(t *testing.T) {
	registry, err := NewProtocolRegistry()
	if err != nil {
		t.Fatalf("NewProtocolRegistry: %v", err)
	}
	re := "4"
	runID := "run-1"
	epoch := "2"
	cases := []Envelope{
		{
			V: 1, Channel: ChannelCtrlW, Type: "turn_open", Seq: "5", RunID: &runID, TurnEpoch: &epoch,
			Body: &TurnOpenBody{
				TurnID:        "turn-1",
				AdmissionRef:  AdmissionRef{Kind: AdmissionOperatorInput, TaskInput: strptr("do work")},
				ParkedUnknown: []ParkedUnknown{}, SessionLogPath: "/runtime/session.log",
				RunDisposition: RunDispositionResume, CreateAuthID: "00000000000000000000000000000000",
			},
		},
		{
			V: 1, Channel: ChannelCtrlW, Type: "genesis_committed", Seq: "6", RunID: &runID, TurnEpoch: &epoch,
			Body: &GenesisCommittedBody{GenerationID: "generation-1"},
		},
		{
			V: 1, Channel: ChannelCtrlC, Type: "connector_assign", Seq: "5",
			Body: &ConnectorAssignBody{
				RunID: "run-1", TurnEpoch: "2", RunManifestDigest: digestA, PolicyDigest: digestA,
				ProviderLaneID: "lane-1", LaneCatalogDigest: digestA, CredentialRef: "cred-1",
			},
		},
		{
			V: 1, Channel: ChannelBroker, Type: "state_proposal_result", Seq: "5", Re: &re,
			Body: &StateProposalResultBody{
				ProposalCorrelation: "proposal-1", Disposition: ProposalInstalled,
				InstalledState: &EpochStateBody{RunID: "run-1", GenerationID: "gen-1", TurnEpoch: "2", LeaseState: LeaseUnleased, StateSeq: "9"},
			},
		},
	}
	for _, tc := range cases {
		t.Run(string(tc.Channel)+"/"+tc.Type, func(t *testing.T) {
			encoded, err := registry.Encode(tc)
			if err != nil {
				t.Fatalf("Encode: %v", err)
			}
			decoded, err := registry.Decode(encoded)
			if err != nil {
				t.Fatalf("Decode: %v", err)
			}
			if reflect.TypeOf(decoded.Body) != reflect.TypeOf(tc.Body) {
				t.Fatalf("decoded body type = %T, want %T", decoded.Body, tc.Body)
			}
			reencoded, err := registry.Encode(decoded)
			if err != nil {
				t.Fatalf("re-Encode: %v", err)
			}
			if string(reencoded) != string(encoded) {
				t.Fatalf("typed round trip changed bytes\nfirst:  %s\nsecond: %s", encoded, reencoded)
			}
		})
	}
}

func TestProtocolRegistryRejectsClosedEnumsAndDiscriminatedShapes(t *testing.T) {
	registry, err := NewProtocolRegistry()
	if err != nil {
		t.Fatalf("NewProtocolRegistry: %v", err)
	}
	re := "0"
	cases := []Envelope{
		{V: 1, Channel: ChannelCtrlW, Type: "attach_result", Seq: "1", Body: &AttachResultBody{GenerationID: "g", TurnEpoch: "1", Result: "future-token"}},
		{V: 1, Channel: ChannelCtrlW, Type: "attempt_open_reject", Seq: "1", Re: &re, Body: &AttemptOpenRejectBody{AttemptID: "a", Reason: "future-reason"}},
		{V: 1, Channel: ChannelCtrlW, Type: "record_tool_outcome", Seq: "1", Body: &RecordToolOutcomeBody{TicketID: "t", TurnEpoch: "1", Outcome: OutcomeExecuted}},
		{V: 1, Channel: ChannelBroker, Type: "state_proposal_result", Seq: "1", Re: &re, Body: &StateProposalResultBody{ProposalCorrelation: "p", Disposition: ProposalTransitionStarted, InstalledState: &EpochStateBody{RunID: "r", GenerationID: "g", TurnEpoch: "1", LeaseState: LeaseUnleased, StateSeq: "1"}}},
		{V: 1, Channel: ChannelCtrlW, Type: "turn_open", Seq: "1", RunID: strptr("run-1"), TurnEpoch: strptr("1"), Body: &TurnOpenBody{
			TurnID: "turn-2", AdmissionRef: AdmissionRef{Kind: AdmissionWakeRelay, RelayID: strptr("relay-1")},
			ParkedUnknown: []ParkedUnknown{}, RunDisposition: RunDispositionResume,
			CreateAuthID: "00000000000000000000000000000000", SessionLogPath: "/runtime/session.log",
			SettlementManifest: &SettlementManifest{Version: 1, RunID: "run-1", ProducedForTurnID: "turn-2", Entries: []SettlementEntry{{
				Kind: "future", Class: "uncertain", RunID: "run-1", TurnID: "turn-1", Terminal: "completed",
			}}},
			PredecessorTurnID: strptr("turn-1"),
		}},
	}
	for _, tc := range cases {
		if _, err := registry.Encode(tc); !errors.Is(err, ErrInvalidMessageBody) {
			t.Errorf("Encode(%s/%s) error = %v, want ErrInvalidMessageBody", tc.Channel, tc.Type, err)
		}
	}
}

func TestBrokerFamiliesAreClosedWhileControlFamiliesAreAdditive(t *testing.T) {
	registry, err := NewProtocolRegistry()
	if err != nil {
		t.Fatalf("NewProtocolRegistry: %v", err)
	}
	ctrl := []byte(`{"body":{"future":"ignored","relay_id":"r"},"chan":"ctrl-w","seq":"0","type":"wake_forward","v":1}`)
	if _, err := registry.Decode(ctrl); err != nil {
		t.Fatalf("Decode additive control message: %v", err)
	}
	broker := []byte(`{"body":{"crossing_count":"removed","epoch_transition_id":"e","generation_id":"g","state_seq":"2","turn_epoch":"1"},"chan":"broker","seq":"0","type":"epoch_installed","v":1}`)
	if _, err := registry.Decode(broker); !errors.Is(err, ErrUnknownField) {
		t.Fatalf("Decode broker message with removed field error = %v, want ErrUnknownField", err)
	}
}

func TestProtocolRegistryEnforcesMessageSpecificEnvelopeAndRejectIdentity(t *testing.T) {
	registry, err := NewProtocolRegistry()
	if err != nil {
		t.Fatalf("NewProtocolRegistry: %v", err)
	}
	body := &ContentReadyBody{
		TurnID: "turn-1", AttemptID: "attempt-1", RoundIdentity: digestA,
		SeqHWM: "4", GenerationID: "gen-1",
	}
	if _, err := registry.Encode(Envelope{V: 1, Channel: ChannelCtrlW, Type: "content_ready", Seq: "1", Body: body}); !errors.Is(err, ErrInvalidMessageBody) {
		t.Fatalf("content_ready without run/epoch error = %v, want ErrInvalidMessageBody", err)
	}

	re := "0"
	consumeReject := Envelope{
		V: 1, Channel: ChannelCtrlW, Type: "stale_epoch", Seq: "1", Re: &re,
		Body: &TypedRejectBody{TicketID: strptr("ticket-1")},
	}
	encoded, err := registry.Encode(consumeReject)
	if err != nil {
		t.Fatalf("Encode consume-side stale_epoch: %v", err)
	}
	decoded, err := registry.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode consume-side stale_epoch: %v", err)
	}
	if body, ok := decoded.Body.(*TypedRejectBody); !ok || body.TicketID == nil || *body.TicketID != "ticket-1" {
		t.Fatalf("decoded consume-side reject = %#v", decoded.Body)
	}

	bad := consumeReject
	bad.Body = &TypedRejectBody{ToolCallID: strptr("call-1"), TicketID: strptr("ticket-1")}
	if _, err := registry.Encode(bad); !errors.Is(err, ErrInvalidMessageBody) {
		t.Fatalf("two-identity typed reject error = %v, want ErrInvalidMessageBody", err)
	}
}

func strptr(value string) *string { return &value }

const digestA = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
