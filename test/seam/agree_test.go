//go:build seam

package seam_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"testing"
	"time"

	"github.com/jackli/frank/internal/appctl/f59"
	"github.com/jackli/frank/internal/appctl/manifest"
	"github.com/jackli/frank/internal/appctl/settle"
	"github.com/jackli/frank/internal/appctl/supervisor"
	"github.com/jackli/frank/internal/appipc"
	connectorframe "github.com/jackli/frank/internal/connector/frame"
	connectorjcs "github.com/jackli/frank/internal/connector/jcs"
	connectoroutcome "github.com/jackli/frank/internal/connector/outcome"
	"github.com/jackli/frank/internal/connector/stream"
	"github.com/jackli/frank/internal/worker/catalog"
	"github.com/jackli/frank/internal/worker/executor"
	"github.com/jackli/frank/internal/worker/fake"
	workerjcs "github.com/jackli/frank/internal/worker/jcs"
	"github.com/jackli/frank/internal/worker/provider"
	"github.com/jackli/frank/internal/worker/resume"
	workerruntime "github.com/jackli/frank/internal/worker/runtime"
	"github.com/jackli/frank/internal/worker/turn"
	"github.com/jackli/frank/internal/worker/wire"
)

// TestCT_G01 binds CT-G01; selectors r9 G01 and S-M10 D§2/A§1/C§10; flags AGREE.
func TestCT_G01(t *testing.T) {
	contract(t, appipc.FrameMax == 4*1024*1024 && connectorframe.FrameMax == 4*1024*1024 && wire.MaxFramePayload == 4*1024*1024, "all three frame maxima must be 4 MiB")
	payload := bytes.Repeat([]byte{'x'}, appipc.FrameMax)
	var framed bytes.Buffer
	if err := appipc.WriteFrame(&framed, payload); err != nil {
		t.Fatal(err)
	}
	decoded, err := appipc.ReadFrame(&framed)
	contract(t, err == nil && bytes.Equal(decoded, payload), "the boundary payload must round-trip")
}

// TestCT_G02 binds CT-G02; selectors r9 G02 and S-M10 D§1/A§2; flags AGREE.
func TestCT_G02(t *testing.T) {
	valid := []string{"0", "1", "9", "10", "18446744073709551615"}
	invalid := []string{"", "00", "01", "+1", "-1", " 1", "18446744073709551616"}
	for _, text := range valid {
		_, a := appipc.ParseCounter(text)
		_, c := connectorframe.ParseCounter(text)
		_, w := wire.ParseCounter(text)
		contract(t, a == nil && c == nil && w == nil, "valid counter rejected: "+text)
	}
	for _, text := range invalid {
		_, a := appipc.ParseCounter(text)
		_, c := connectorframe.ParseCounter(text)
		_, w := wire.ParseCounter(text)
		contract(t, a != nil && c != nil && w != nil, "invalid counter accepted: "+text)
	}
}

// TestCT_G03 binds CT-G03; selectors r9 G03 and S-M10 A§3; flags AGREE.
func TestCT_G03(t *testing.T) {
	contract(t, equalStrings(
		[]string{string(appipc.ChannelCtrlW), string(appipc.ChannelCtrlC), string(appipc.ChannelDataP), string(appipc.ChannelBroker)},
		[]string{string(wire.ChannelCTRLW), string(wire.ChannelCTRLC), string(wire.ChannelDATAP), string(wire.ChannelBroker)},
	), "appipc and worker channel enums differ")
}

// TestCT_G04 binds CT-G04; selectors r9 G04 and S-M10 A§4; flags AGREE.
func TestCT_G04(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	body := map[string]any{"pid": 7, "build_info": map[string]string{"version": "dev", "commit": "unknown", "built_at": "unknown"}}
	payload, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "hello", Seq: "1", Body: body})
	if err != nil {
		t.Fatal(err)
	}
	codec, err := wire.NewCodec([]wire.MessageSpec{{Channel: wire.ChannelCTRLW, Type: "hello"}})
	if err != nil {
		t.Fatal(err)
	}
	workerFrame, err := codec.Decode(framed(payload))
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := codec.Encode(workerFrame)
	if err != nil {
		t.Fatal(err)
	}
	_, err = registry.Decode(encoded[4:])
	contract(t, err == nil, "CTRL-W envelope must decode in both directions")
}

// TestCT_G05 binds CT-G05; selectors r9 G05 and S-M10 D§4; flags AGREE.
func TestCT_G05(t *testing.T) {
	corpus := [][]byte{[]byte(`null`), []byte(`{"b":2,"a":1}`), []byte(`[true,"x",{"z":0}]`), []byte(`{"\u20ac":"line\n"}`)}
	for _, input := range corpus {
		a, ea := appCanonical(input)
		c, ec := connectorjcs.Canonicalize(input)
		w, ew := workerjcs.Canonicalize(input)
		contract(t, ea == nil && ec == nil && ew == nil && canonicalEqual(a, c, w), "JCS bytes differ for "+string(input))
	}
}

// TestCT_G06 binds CT-G06; selectors r9 G06 and ratified eight-name constant; flags AGREE.
func TestCT_G06(t *testing.T) {
	names := make([]string, 0, len(catalog.ExpectedIdentities()))
	for _, identity := range catalog.ExpectedIdentities() {
		names = append(names, identity.CanonicalName)
	}
	contract(t, equalStrings(manifest.RatifiedToolNames, names), "manifest and worker tool catalogs differ")
}

// TestCT_G07 binds CT-G07; selectors r9 G07 and S-M10 B§7-8; flags AGREE.
func TestCT_G07(t *testing.T) {
	contract(t, string(executor.OutcomeExecuted) == appipc.OutcomeExecuted && string(executor.OutcomeNotInvokedIntegrityFault) == appipc.OutcomeNotInvokedIntegrityFault, "executor and F59 outcome identities differ")
}

// TestCT_G08 binds CT-G08; selectors r9 G08 and S-M10 B§9; flags AGREE.
func TestCT_G08(t *testing.T) {
	identity := executor.FullIdentity{RunID: "run", TurnID: "turn", ToolCallID: "call", Identity: executor.Identity{CanonicalToolName: "read", CanonicalArgsDigest: digest('a'), TurnEpoch: "1"}}
	peer := fake.NewM10(workerruntime.Assignment{})
	granted, grantErr := peer.Authorize(context.Background(), executor.AuthorizeRequest{Identity: identity})
	request := executor.ConsumeRequest{TicketID: granted.TicketID, TurnEpoch: identity.TurnEpoch, CanonicalToolName: identity.CanonicalToolName, CanonicalArgsDigest: identity.CanonicalArgsDigest}
	first, firstErr := peer.Consume(context.Background(), request)
	replay, replayErr := peer.Consume(context.Background(), request)
	contract(t, grantErr == nil && granted.Code == executor.AuthorizeGranted && firstErr == nil && first.Code == executor.ConsumeOK && replayErr == nil && replay.Code == executor.ConsumeDuplicate,
		explain("one-shot/replay matrix differs: grant=%+v/%v first=%+v/%v replay=%+v/%v", granted, grantErr, first, firstErr, replay, replayErr))
	contract(t, structHasField(f59.ConsumeRequest{}, "TicketID") && structHasField(f59.OutcomeRequest{}, "TicketID"), "consume and outcome must retain the one-shot ticket identity")
}

// TestCT_G09 binds CT-G09; selectors r9 G09 and S-M10 B§10/A§12; flags AGREE.
func TestCT_G09(t *testing.T) {
	reasons := executor.LifecycleRejectReasons()
	reasonNames := make([]string, len(reasons))
	for i, reason := range reasons {
		reasonNames[i] = string(reason)
	}
	contract(t, equalStrings(reasonNames, []string{"run_not_admitted", "turn_inactive", "lease_invalid", "turn_budget_exhausted"}), "authorize reject_reason enum differs")
	contract(t, equalStrings([]string{string(provider.StreamCompleted), string(provider.StreamFailed), string(provider.StreamCancelled), string(provider.StreamLostEnd)}, []string{"stream_completed", "stream_failed", "stream_cancelled", "stream_lost"}), "stream-end enum differs")
	contract(t, equalStrings([]string{string(turn.TurnCompleted), string(turn.TurnRefused), string(turn.TurnDenied), string(turn.TurnCancelled), string(turn.TurnFailed), string(turn.TurnExhausted)}, []string{"turn_completed", "turn_refused", "turn_denied", "turn_cancelled", "turn_failed", "turn_exhausted"}), "turn-terminal enum differs")
	contract(t, equalStrings([]string{string(workerruntime.AttachOK), string(workerruntime.AttachSuspended), string(workerruntime.AttachTupleMismatch)}, []string{appipc.AttachOK, appipc.AttachSuspended, appipc.AttachTupleMismatch}), "attach-result enum differs")
}

// TestCT_G10 binds CT-G10; selectors r9 G10 and S-M10 A-add; flags AGREE.
func TestCT_G10(t *testing.T) {
	now := testTime()
	frame, err := resume.ContentReady(resume.ContentReadyInput{Seq: "1", RunID: "run", TurnEpoch: "2", TurnID: "turn", AttemptID: "attempt", RoundIdentity: digest('a'), SeqHWM: "3", GenerationID: "gen", ContentFsyncAt: now, MarkerFsyncAt: now, EmitAt: now})
	if err != nil {
		t.Fatal(err)
	}
	var body appipc.ContentReadyBody
	err = json.Unmarshal(frame.Body, &body)
	contract(t, err == nil && body.TurnID == "turn" && body.AttemptID == "attempt" && body.RoundIdentity == digest('a') && body.SeqHWM == "3" && body.GenerationID == "gen", "content_ready body shapes differ")
}

// TestCT_G11 binds CT-G11; selectors r9 G11 and S-M10 A-add; flags AGREE.
func TestCT_G11(t *testing.T) {
	contract(t, equalStrings(jsonFields(appipc.ResumeDispositionBody{}), []string{"turn_id", "disposition", "resume_action"}), "resume-disposition body shape changed")
	gate := resume.NewWorkGate("run", "turn", "1")
	frame, err := gate.Report("1", resume.DispositionDegraded)
	contract(t, err == nil && frame.Type == "report_resume_disposition", "worker cannot report the registered resume disposition")
}

// TestCT_G12 binds CT-G12; selectors r9 G12 and S-M10 A-add; flags AGREE.
func TestCT_G12(t *testing.T) {
	contract(t, reflect.DeepEqual(jsonFields(appipc.ParkedUnknown{}), jsonFields(wire.ParkedUnknown{})) && reflect.DeepEqual(jsonFields(wire.ParkedUnknown{}), jsonFields(turn.ParkedUnknown{})), "parked-unknown shapes differ")
}

// TestCT_G13 binds CT-G13; selectors r9 G13 and S-M10 A§9; flags AGREE shape-only.
func TestCT_G13(t *testing.T) {
	contract(t, reflect.DeepEqual(jsonFields(appipc.SettlementManifest{}), jsonFields(resume.SettlementManifest{})), "settlement-manifest top-level shapes differ")
	contract(t, reflect.DeepEqual(jsonFields(appipc.SettlementEntry{}), jsonFields(resume.ManifestEntryWire{})), "settlement entry wire shapes differ")
}

type opaqueGate struct{}

func (opaqueGate) AttemptOpen(context.Context, provider.Request) error               { return nil }
func (opaqueGate) RecordStreamEnd(context.Context, string, provider.StreamEnd) error { return nil }

type opaqueConnector struct{ item []byte }

func (connector opaqueConnector) Attempt(context.Context, provider.Request) (provider.Disposition, []json.RawMessage, error) {
	return provider.Completed, []json.RawMessage{append([]byte(nil), connector.item...)}, nil
}
func (opaqueConnector) Cancel(context.Context, string, string) (provider.Disposition, error) {
	return provider.CancelledPre, nil
}

// TestCT_G14 binds CT-G14; selectors r9 G14 and S-M8v1 C§11; flags AGREE.
func TestCT_G14(t *testing.T) {
	want := []byte{0xff, 0x00, '{', 'x', '}'}
	cycle, err := provider.New(opaqueGate{}, opaqueConnector{item: want}, "1")
	if err != nil {
		t.Fatal(err)
	}
	got, err := cycle.Run(context.Background(), provider.Request{AttemptID: "a", TurnID: "t", TurnEpoch: "1", ProviderLane: "p"})
	contract(t, err == nil && len(got.Events) == 1 && bytes.Equal(got.Events[0].Opaque, want), "opaque provider item was parsed or rewritten")
}

// TestCT_G15 binds CT-G15; selectors r9 G15 and S-M8v1 C§8; flags AGREE.
func TestCT_G15(t *testing.T) {
	contract(t, string(provider.EpochAhead) == "EPOCH_AHEAD" && string(provider.StaleEpoch) == "STALE_EPOCH", "epoch disposition facts differ")
}

// TestCT_G16 binds CT-G16; selectors r9 G16 and S-M8v1 C-lineup; flags AGREE facts.
func TestCT_G16(t *testing.T) {
	wantKinds := []stream.Kind{stream.AttemptStarted, stream.TextStart, stream.TextDelta, stream.TextEnd, stream.ReasoningStart, stream.ReasoningDelta, stream.ReasoningEnd, stream.ToolCallStart, stream.ToolCallDelta, stream.ToolCallEnd, stream.UsageProgress, stream.Completed, stream.Failed, stream.Cancelled}
	contract(t, len(wantKinds) == 14 && stream.SchemaV2 == "m8.provider_event.v2", "normalized provider factual lineup differs")
	contract(t, string(provider.StreamCompleted) == "stream_completed", "worker cannot represent connector terminal facts")
}

// TestCT_G17 binds CT-G17; selectors r9 G17 and A-M1-SECRET; flags AGREE negative.
func TestCT_G17(t *testing.T) {
	providerLane := reflect.TypeOf(manifest.ProviderLane{})
	for i := 0; i < providerLane.NumField(); i++ {
		contract(t, providerLane.Field(i).Type.Kind() != reflect.Slice, "m-10 manifest exposes credential bytes")
	}
	contract(t, structHasField(manifest.ProviderLane{}, "CredentialRef") && structHasField(appipc.ConnectorAssignBody{}, "CredentialRef"), "credential orchestration must remain opaque-reference only")
}

// TestCT_G18 binds CT-G18; selectors r9 G18 and S-M10 D§14; flags AGREE.
func TestCT_G18(t *testing.T) {
	pipe, err := supervisor.NewDeathPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer pipe.Child.Close()
	if err := pipe.Parent.Close(); err != nil {
		t.Fatal(err)
	}
	buffer := make([]byte, 1)
	_, err = pipe.Child.Read(buffer)
	contract(t, errors.Is(err, io.EOF), "connector death pipe does not signal parent death as EOF")
}

// TestCT_G19 binds CT-G19; selectors r9 G19 and S-M10 A§1; flags AGREE.
func TestCT_G19(t *testing.T) {
	contract(t, appipc.ParkedRowMax == 640 && wire.ParkedRowMax == 640, "PARKED_ROW_MAX differs")
}

// TestCT_G20 binds CT-G20; selectors G20 addendum and CTRL-C attempt_result; flags AGREE regression pin.
func TestCT_G20(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	fixture := newAppFixture(t)
	fixture.seedAttempt(t, "present")
	fixture.seedAttempt(t, "absent")
	host := settle.New(fixture.applier)

	produced, err := connectoroutcome.AttemptResult("present", 1, connectoroutcome.Outcome{Kind: connectoroutcome.TransportFailed, Digests: &connectoroutcome.Digests{FrozenCore: digest('a'), ProviderLoweredTools: digest('b')}})
	if err != nil {
		t.Fatal(err)
	}
	presentBody := appipc.AttemptResultBody{AttemptID: produced.AttemptID, TurnEpoch: produced.TurnEpoch.String(), Disposition: produced.Disposition, FrozenCoreDigest: pointer(produced.FrozenCoreDigest), ProviderLoweredToolsDigest: pointer(produced.ProviderLoweredToolsDigest)}
	payload, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "attempt_result", Seq: "1", Body: presentBody})
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := registry.Decode(payload)
	if err != nil {
		t.Fatal(err)
	}
	present := decoded.Body.(*appipc.AttemptResultBody)
	decision, err := host.RecordAttemptResult(fixture.ctx, settle.AttemptResultRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: present.AttemptID, TurnEpoch: present.TurnEpoch, Disposition: present.Disposition, FrozenCoreDigest: present.FrozenCoreDigest, ProviderLoweredToolsDigest: present.ProviderLoweredToolsDigest, At: 2})
	if err != nil || decision != settle.CarriageRecorded {
		t.Fatalf("G20 present carriage: decision=%q err=%v", decision, err)
	}
	row, err := host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "present")
	contract(t, err == nil && row.FrozenCoreDigest != nil && *row.FrozenCoreDigest == digest('a') && row.ProviderLoweredToolsDigest != nil && *row.ProviderLoweredToolsDigest == digest('b'), "produced B/E digests were not stored byte-verbatim")

	reason := "malformed_request"
	stage := "pre_freeze"
	absentBody := appipc.AttemptResultBody{AttemptID: "absent", TurnEpoch: "1", Disposition: "rejected_local", RejectReason: &reason, RefusalStage: &stage}
	absentPayload, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "attempt_result", Seq: "2", Body: absentBody})
	if err != nil {
		t.Fatal(err)
	}
	absentEnvelope, err := registry.Decode(absentPayload)
	if err != nil {
		t.Fatal(err)
	}
	absent := absentEnvelope.Body.(*appipc.AttemptResultBody)
	decision, err = host.RecordAttemptResult(fixture.ctx, settle.AttemptResultRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: absent.AttemptID, TurnEpoch: absent.TurnEpoch, Disposition: absent.Disposition, RefusalStage: absent.RefusalStage, At: 3})
	if err != nil || decision != settle.CarriageRecorded {
		t.Fatalf("G20 absent carriage: decision=%q err=%v", decision, err)
	}
	row, err = host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "absent")
	contract(t, err == nil && row.FrozenCoreDigest == nil && row.ProviderLoweredToolsDigest == nil, "absent B/E fields did not remain SQL NULL")

	invalid, err := appipc.MarshalJCS(map[string]any{"v": 1, "chan": "ctrl-c", "type": "attempt_result", "seq": "3", "body": map[string]any{"attempt_id": "bad", "turn_epoch": "1", "disposition": "sent_completed", "frozen_core_digest": "bad"}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = registry.Decode(invalid)
	contract(t, err != nil, "invalid-width digest was accepted at decode")
}

func testTime() (now time.Time) { return time.Unix(1, 0).UTC() }
