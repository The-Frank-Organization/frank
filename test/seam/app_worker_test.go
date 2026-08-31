//go:build seam

package seam_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/appctl/f59"
	"github.com/The-Frank-Organization/frank/internal/appctl/manifest"
	"github.com/The-Frank-Organization/frank/internal/appipc"
	"github.com/The-Frank-Organization/frank/internal/worker/catalog"
	"github.com/The-Frank-Organization/frank/internal/worker/executor"
	"github.com/The-Frank-Organization/frank/internal/worker/journal"
	"github.com/The-Frank-Organization/frank/internal/worker/provider"
	workerruntime "github.com/The-Frank-Organization/frank/internal/worker/runtime"
	"github.com/The-Frank-Organization/frank/internal/worker/turn"
)

// TestCT_A01 binds CT-A01; selectors r9 A01 and S-M10/S-D1 §4b/§0C-11; flags UNWIRED, W, star.
func TestCT_A01(t *testing.T) {
	probe := newWorkerProbe(t)
	_, err := probe.run(context.Background())
	contract(t, err == nil && len(probe.control.genesisBodies) == 1 && probe.control.genesisBodies[0].GenerationID == probe.control.assignment.GenerationID,
		explain("live worker genesis notification count=%d err=%v", len(probe.control.genesisBodies), err))
}

// TestCT_A02 binds CT-A02; selectors r9 A02 and S-M10/S-D1 §4b; flags MISMATCH, W, star.
func TestCT_A02(t *testing.T) {
	valid := newWorkerProbe(t)
	_, validErr := valid.run(context.Background())
	invalid := newWorkerProbe(t)
	invalid.control.assignment.CreateAuthID = strings.Repeat("z", 32)
	_, invalidErr := invalid.run(context.Background())
	contract(t, validErr == nil && invalidErr != nil && invalid.broker.attachCalls == 0,
		explain("pre-journal admission witness: valid=%v invalid=%v invalid_attach_calls=%d", validErr, invalidErr, invalid.broker.attachCalls))
}

// TestCT_A03 binds CT-A03; selectors r9 A03 and S-M10 A§7; flags MISMATCH, W, star.
func TestCT_A03(t *testing.T) {
	probe := newWorkerProbe(t)
	chosenDir := filepath.Join(t.TempDir(), "app-chosen")
	if err := os.Mkdir(chosenDir, 0o700); err != nil {
		t.Fatal(err)
	}
	chosen := filepath.Join(chosenDir, journal.SessionLogName)
	probe.control.assignment = assignmentWithJSON(t, probe.control.assignment, map[string]any{"session_log_path": chosen})
	_, err := probe.run(context.Background())
	_, chosenErr := os.Stat(chosen)
	_, fallbackErr := os.Stat(filepath.Join(probe.runtimeDir, journal.SessionLogName))
	contract(t, err == nil && chosenErr == nil && os.IsNotExist(fallbackErr),
		explain("app-chosen journal effect: run=%v chosen=%v fallback=%v", err, chosenErr, fallbackErr))
}

// TestCT_A04 binds CT-A04; selectors r9 A04 and S-M10/S-D1 §0C-11; flags MISMATCH, W, star.
func TestCT_A04(t *testing.T) {
	fresh := newWorkerProbe(t)
	if _, err := fresh.run(context.Background()); err != nil {
		t.Fatal(err)
	}
	selected := filepath.Join(fresh.runtimeDir, journal.SessionLogName)
	resumeProbe := newWorkerProbe(t)
	resumeProbe.control.assignment.TurnID = "turn-resume"
	resumeProbe.control.assignment = assignmentWithJSON(t, resumeProbe.control.assignment, map[string]any{"session_log_path": selected})
	resumeProbe.config.RunDisposition = appipc.RunDispositionResume
	_, err := resumeProbe.run(context.Background())
	contract(t, err == nil, explain("resume failed to consume the app-selected existing log: %v", err))
}

// TestCT_A05 binds CT-A05; selectors r9 A05 and S-M10 A§10/B§2; flags MISMATCH, W, star.
func TestCT_A05(t *testing.T) {
	kind, present := fieldKind(provider.Request{}, "TurnEpoch")
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	runID, epoch := "run", "1"
	body := appipc.AuthorizeToolCallBody{RunID: runID, TurnID: "turn", TurnEpoch: epoch, ToolCallID: "call", CanonicalToolName: "read", CanonicalArgsDigest: digest('a')}
	encoded, roundErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "authorize_tool_call", Seq: "1", RunID: &runID, TurnEpoch: &epoch, Body: body})
	_, decodeErr := registry.Decode(encoded)
	badEpoch := "01"
	_, leadingZeroErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "authorize_tool_call", Seq: "1", RunID: &runID, TurnEpoch: &badEpoch, Body: appipc.AuthorizeToolCallBody{RunID: runID, TurnID: "turn", TurnEpoch: badEpoch, ToolCallID: "call", CanonicalToolName: "read", CanonicalArgsDigest: digest('a')}})
	contract(t, present && kind == reflect.String && roundErr == nil && decodeErr == nil && leadingZeroErr != nil,
		explain("canonical-decimal CTRL-W round trip: provider_kind=%v encode=%v decode=%v leading_zero=%v", kind, roundErr, decodeErr, leadingZeroErr))
}

// TestCT_A06 binds CT-A06; selectors r9 A06 and S-M10 A§11/B§6; flags MISMATCH, W, star.
func TestCT_A06(t *testing.T) {
	probe := newWorkerProbe(t)
	_, runErr := probe.run(context.Background())
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	var encodeErr error
	if len(probe.control.outcomeRecords) == 1 {
		_, encodeErr = registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "record_tool_outcome", Seq: "1", Body: probe.control.outcomeRecords[0]})
	}
	contract(t, runErr == nil && len(probe.control.outcomeRecords) == 1 && encodeErr == nil,
		explain("live outcome CTRL-W carriage: run=%v records=%d encode=%v", runErr, len(probe.control.outcomeRecords), encodeErr))
}

// TestCT_A07 binds CT-A07; selectors r9 A07 and S-M10 A§11/B§3; flags MISMATCH, W, star.
func TestCT_A07(t *testing.T) {
	probe := newWorkerProbe(t)
	_, runErr := probe.run(context.Background())
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	var encodeErr error
	if len(probe.control.authorizeRequests) == 1 {
		_, encodeErr = registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "authorize_tool_call", Seq: "1", Body: probe.control.authorizeRequests[0]})
	}
	contract(t, runErr == nil && len(probe.control.authorizeRequests) == 1 && encodeErr == nil,
		explain("live authorization CTRL-W carriage: run=%v requests=%d encode=%v", runErr, len(probe.control.authorizeRequests), encodeErr))
}

// TestCT_A08 binds CT-A08; selectors r9 A08 and S-M10/S-CP B§5; flags MISMATCH, CP, star.
func TestCT_A08(t *testing.T) {
	contract(t, !structHasField(f59.ConsumeRequest{}, "GenerationID") && !structHasField(f59.OutcomeRequest{}, "GenerationID"), "m-10 consume/outcome requests still trust a wire generation_id instead of channel-assigned identity")
}

// TestCT_A09 binds CT-A09; selectors r9 A09 and S-M10/S-CP B§4; flags MISMATCH, W, latent.
func TestCT_A09(t *testing.T) {
	probe := newWorkerProbe(t)
	probe.control.authorizeReply = executor.AuthorizeReply{Code: executor.AuthorizeGranted, TicketID: "ticket-without-effect-descriptor"}
	_, err := probe.run(context.Background())
	grantedCarriesDescriptor := structHasField(executor.AuthorizeReply{}, "EffectDescriptor")
	var typed *executor.Error
	typedDescriptorReason := errors.As(err, &typed) && strings.Contains(strings.ToLower(typed.Error()), "descriptor")
	contract(t, grantedCarriesDescriptor && typedDescriptorReason && probe.backend.Writes == 0,
		explain("granted absent/mismatched descriptor was not rejected with a typed descriptor reason: descriptor_field=%v typed_descriptor_reason=%v writes=%d err=%v", grantedCarriesDescriptor, typedDescriptorReason, probe.backend.Writes, err))
}

// TestCT_A10 binds CT-A10; selectors r9 A10 and S-M10 A§12/B§10; flags MISMATCH, W, latent.
func TestCT_A10(t *testing.T) {
	failures := exerciseReplyDemux(t)
	contract(t, len(failures) == 0, explain("executor reply branches not reached at lowercase wire tokens: %v", failures))
}

// TestCT_A11 binds CT-A11; selectors r9 A11 and S-M10 B§11; flags MISMATCH, W, latent.
func TestCT_A11(t *testing.T) {
	prepared, err := executor.Prepare("run", "turn", "call", "read", "1", staticArguments(`{"path":"x"}`))
	if err != nil {
		t.Fatal(err)
	}
	authority := &replyAuthority{authorizeErr: io.EOF}
	worker, err := executor.New(authority, &countingInvoker{})
	if err != nil {
		t.Fatal(err)
	}
	_, runErr := worker.Execute(context.Background(), prepared)
	var typed *executor.Error
	contract(t, errors.As(runErr, &typed) && string(typed.Code) == "no_reply", explain("withheld reply error is not typed no_reply: %T %v", runErr, runErr))
}

// TestCT_A12 binds CT-A12; selectors r9 A12 and ratified eight-name constant; flags UNWIRED, W, star.
func TestCT_A12(t *testing.T) {
	probe := newWorkerProbe(t)
	probe.provider.tool = workerruntime.ToolCall{ID: "call", CanonicalName: "relay.read", Arguments: []byte(`{"relay_id":"relay-1"}`)}
	_, err := probe.run(context.Background())
	contract(t, err == nil && probe.backend.relayReads == 1, explain("native relay registry effect: reads=%d err=%v", probe.backend.relayReads, err))
}

// TestCT_A13 binds CT-A13; selectors r9 A13 and S-M10 A§9; flags UNWIRED, W/CP, latent.
func TestCT_A13(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	reasons := []string{"run_not_admitted", "turn_inactive", "lease_invalid", "denied_above_set", "expired"}
	encodeReason := func(reason string) error {
		runID, epoch := "run", "1"
		body := appipc.TurnOpenBody{TurnID: "turn", AdmissionRef: appipc.AdmissionRef{Kind: appipc.AdmissionOperatorInput, TaskInput: pointer("x")}, ParkedUnknown: []appipc.ParkedUnknown{}, RunDisposition: appipc.RunDispositionResume, CreateAuthID: strings.Repeat("b", 32), SessionLogPath: "/session.log", PredecessorTurnID: pointer("previous"), SettlementManifest: &appipc.SettlementManifest{Version: 1, RunID: "run", ProducedForTurnID: "turn", Entries: []appipc.SettlementEntry{{Kind: "tool", Class: "determinate_no_resume", RunID: "run", TurnID: "previous", ToolCallID: pointer("call"), ArgsDigest: pointer(digest('a')), Terminal: "VOID", VoidReason: pointer(reason)}}}}
		_, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "turn_open", Seq: "1", RunID: &runID, TurnEpoch: &epoch, Body: body})
		return err
	}
	var membershipFailures []string
	for _, reason := range reasons {
		if err := encodeReason(reason); err != nil {
			membershipFailures = append(membershipFailures, reason+":"+err.Error())
		}
	}
	unknownErr := encodeReason("future_reason")
	probe := newWorkerProbe(t)
	probe.control.assignment = assignmentWithJSON(t, probe.control.assignment, map[string]any{"predecessor_turn_id": "previous", "settlement_manifest": map[string]any{"settlement_manifest_v": 1, "run_id": "run", "produced_for_turn_id": "turn", "entries": []any{map[string]any{"kind": "tool", "class": "determinate_no_resume", "run_id": "run", "turn_id": "previous", "tool_call_id": "call", "args_digest": digest('a'), "terminal": "VOID", "void_reason": "future_reason"}}}})
	_, workerErr := probe.run(context.Background())
	contract(t, len(membershipFailures) == 0 && unknownErr != nil && workerErr != nil && probe.provider.attempts == 0,
		explain("settlement void_reason totality: members=%v unknown=%v worker=%v attempts=%d", membershipFailures, unknownErr, workerErr, probe.provider.attempts))
}

// TestCT_A14 binds CT-A14; selectors r9 A14 and S-M10 B bonus; flags MISMATCH, CP/W, latent.
func TestCT_A14(t *testing.T) {
	expectedKeys := []string{"name", "schema_digest", "catalog_version", "mapping_version"}
	schema, catalogVersion, mapping := digest('1'), "catalog-v1", "mapping-v1"
	tool := manifest.ToolIdentity{Name: "relay.read", SchemaDigest: &schema, CatalogVersion: &catalogVersion, MappingVersion: &mapping}
	resource := "relay.read:relay-1"
	descriptor, _, err := f59.BuildDescriptor(manifest.Manifest{ToolSet: []manifest.ToolIdentity{tool}}, "relay.read", digest('2'), f59.Operands{CanonicalResource: &resource})
	encoded, encodeErr := appipc.MarshalJCS(map[string]any{"name": tool.Name, "schema_digest": tool.SchemaDigest, "catalog_version": tool.CatalogVersion, "mapping_version": tool.MappingVersion})
	hash := sha256.Sum256(encoded)
	wantRef := "manifest-tool:" + hex.EncodeToString(hash[:])
	contract(t, equalStrings(jsonFields(catalog.Identity{}), expectedKeys) && err == nil && encodeErr == nil && descriptor.ToolImplRef == wantRef,
		explain("manifest identity/recomputability mismatch: keys=%v descriptor=%v encode=%v got_ref=%q want_ref=%q", jsonFields(catalog.Identity{}), err, encodeErr, descriptor.ToolImplRef, wantRef))
}

// TestCT_A15 binds CT-A15; selectors r9 A15 and S-M10 A-add; flags MISMATCH, CP, latent.
func TestCT_A15(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	nonCanonical := []byte(`{"v":1, "chan":"ctrl-w","type":"ping","seq":"1","body":{}}`)
	_, err = registry.Decode(nonCanonical)
	contract(t, err != nil, "m-10 receive accepts non-canonical envelope bytes that the worker rejects")
}

// TestCT_A16 binds CT-A16; selectors r9 A16 and S-M10 A-add; flags MISMATCH, CP/W, latent.
func TestCT_A16(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	_, encodeErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlW, Type: "ping", Seq: "1", Re: pointer("0"), Body: appipc.EmptyBody{}})
	decodeBytes := []byte(`{"v":1,"chan":"ctrl-w","type":"ping","seq":"1","re":"0","body":{}}`)
	_, decodeErr := registry.Decode(decodeBytes)
	contract(t, encodeErr != nil && decodeErr != nil, explain("reply-only iff differs: encode=%v decode=%v", encodeErr, decodeErr))
}

// TestCT_A17 binds CT-A17; selectors r9 A17 and S-M10 A-add; flags MISMATCH, CP, latent.
func TestCT_A17(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	replies := []string{"content_ready_recorded", "content_ready_duplicate", "content_ready_conflict", "content_ready_stale_epoch", "content_ready_unknown_attempt", "content_ready_future_epoch_fault"}
	var missing []string
	for _, reply := range replies {
		if !registry.Has(appipc.ChannelCtrlW, reply) {
			missing = append(missing, reply)
		}
	}
	contract(t, len(missing) == 0, explain("receipt decisions lack registered reply types: %v", missing))
}

// TestCT_A18 binds CT-A18; selectors r9 A18 and S-M10 A-add; flags MISMATCH, W, latent.
func TestCT_A18(t *testing.T) {
	var absent, empty turn.Open
	if err := json.Unmarshal([]byte(`{"run_id":"r","turn_id":"t","turn_epoch":"1","admission_ref":{"kind":"operator_input","task_input":"x"}}`), &absent); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(`{"run_id":"r","turn_id":"t","turn_epoch":"1","admission_ref":{"kind":"operator_input","task_input":"x"},"parked_unknown":[]}`), &empty); err != nil {
		t.Fatal(err)
	}
	contract(t, !reflect.DeepEqual(absent, empty), "worker collapses absent parked_unknown and present empty array")
}

// TestCT_A19 binds CT-A19; selectors r9 A19 and A-M10-LIMITS; flags MISMATCH, W, benign/latent.
func TestCT_A19(t *testing.T) {
	workerCarriesNoBound := !declaresIdentifier(t, "internal/worker/turn/turn.go", "MaxToolCalls")
	prepared, err := executor.Prepare("run", "turn", "call", "read", "1", staticArguments(`{"path":"x"}`))
	if err != nil {
		t.Fatal(err)
	}
	authority := &replyAuthority{authorizeReply: executor.AuthorizeReply{Code: executor.AuthorizeRejected, RejectReason: executor.RejectTurnBudgetExhausted}}
	worker, err := executor.New(authority, &countingInvoker{})
	if err != nil {
		t.Fatal(err)
	}
	_, runErr := worker.Execute(context.Background(), prepared)
	var typed *executor.Error
	consumesCeiling := errors.As(runErr, &typed) && typed.Code == executor.CodeAuthorizeRejected && typed.Detail == string(executor.RejectTurnBudgetExhausted)
	contract(t, workerCarriesNoBound && consumesCeiling, explain("worker-local bound=%v ceiling-consumption=%v err=%v", !workerCarriesNoBound, consumesCeiling, runErr))
}

type staticArguments string

func (arguments staticArguments) Snapshot() []byte { return []byte(arguments) }

type countingInvoker struct{ calls int }

func (invoker *countingInvoker) Invoke(context.Context, executor.Invocation) (any, error) {
	invoker.calls++
	return "ok", nil
}

type replyAuthority struct {
	authorizeReply executor.AuthorizeReply
	authorizeErr   error
	consumeReply   executor.ConsumeReply
}

func (authority *replyAuthority) Authorize(_ context.Context, request executor.AuthorizeRequest) (executor.AuthorizeReply, error) {
	if authority.authorizeErr != nil {
		return executor.AuthorizeReply{}, authority.authorizeErr
	}
	if authority.authorizeReply.Code == "" {
		return executor.AuthorizeReply{Code: executor.AuthorizeGranted, TicketID: "ticket", EffectDescriptor: executor.DescriptorForIdentity(request.FrozenIdentity())}, nil
	}
	reply := authority.authorizeReply
	if reply.Code == executor.AuthorizeGranted && reply.EffectDescriptor == nil {
		reply.EffectDescriptor = executor.DescriptorForIdentity(request.FrozenIdentity())
	}
	return reply, nil
}

func (authority *replyAuthority) Consume(context.Context, executor.ConsumeRequest) (executor.ConsumeReply, error) {
	if authority.consumeReply.Code == "" {
		return executor.ConsumeReply{Code: executor.ConsumeOK}, nil
	}
	return authority.consumeReply, nil
}

func (*replyAuthority) RecordOutcome(context.Context, executor.OutcomeRecord) error { return nil }

func exerciseReplyDemux(t *testing.T) []string {
	t.Helper()
	authorizeCases := []struct {
		wire   string
		reason executor.AuthorizeRejectReason
		code   executor.Code
	}{
		{"ticket_granted", "", ""},
		{"authorize_reject", executor.RejectTurnInactive, executor.CodeAuthorizeRejected},
		{"stale_epoch", "", executor.CodeStaleEpoch},
		{"denied_above_set", "", executor.CodeDeniedAboveSet},
		{"duplicate_request", "", executor.CodeDuplicateRequest},
		{"identity_mismatch", "", executor.CodeIdentityMismatch},
	}
	consumeCases := []struct {
		wire string
		code executor.Code
	}{
		{"consume_ok", ""},
		{"stale_epoch", executor.CodeStaleEpoch},
		{"duplicate_consume", executor.CodeDuplicateConsume},
		{"identity_mismatch", executor.CodeIdentityMismatch},
	}
	var failures []string
	for _, test := range authorizeCases {
		invoker := &countingInvoker{}
		reply := executor.AuthorizeReply{Code: executor.AuthorizeCode(test.wire), RejectReason: test.reason}
		if test.wire == "ticket_granted" {
			reply.TicketID = "ticket"
		}
		worker, _ := executor.New(&replyAuthority{authorizeReply: reply}, invoker)
		prepared, _ := executor.Prepare("run", "turn", "call", "read", "1", staticArguments(`{"path":"x"}`))
		_, err := worker.Execute(context.Background(), prepared)
		if !replyBranchMatches(err, test.code) {
			failures = append(failures, test.wire+":"+explain("%v", err))
		}
	}
	for _, test := range consumeCases {
		invoker := &countingInvoker{}
		worker, _ := executor.New(&replyAuthority{consumeReply: executor.ConsumeReply{Code: executor.ConsumeCode(test.wire)}}, invoker)
		prepared, _ := executor.Prepare("run", "turn", "call", "read", "1", staticArguments(`{"path":"x"}`))
		_, err := worker.Execute(context.Background(), prepared)
		if !replyBranchMatches(err, test.code) {
			failures = append(failures, test.wire+":"+explain("%v", err))
		}
	}
	return failures
}

func replyBranchMatches(err error, want executor.Code) bool {
	if want == "" {
		return err == nil
	}
	var typed *executor.Error
	return errors.As(err, &typed) && typed.Code == want
}

func declaresIdentifier(t *testing.T, relative, name string) bool {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), filepath.Join(repoRoot(t), filepath.FromSlash(relative)), nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, declaration := range file.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, specification := range general.Specs {
			value, ok := specification.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for _, identifier := range value.Names {
				if identifier.Name == name {
					return true
				}
			}
		}
	}
	return false
}
