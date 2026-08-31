//go:build seam

package seam_test

import (
	"bytes"
	"context"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/brokerclient"
	"github.com/The-Frank-Organization/frank/internal/appctl/manifest"
	"github.com/The-Frank-Organization/frank/internal/appctl/scheduler"
	"github.com/The-Frank-Organization/frank/internal/appctl/settle"
	"github.com/The-Frank-Organization/frank/internal/appctl/supervisor"
	"github.com/The-Frank-Organization/frank/internal/appctl/testutil"
	"github.com/The-Frank-Organization/frank/internal/appipc"
	"github.com/The-Frank-Organization/frank/internal/connector/authorize"
	"github.com/The-Frank-Organization/frank/internal/connector/credentials"
	"github.com/The-Frank-Organization/frank/internal/connector/outcome"
	"github.com/The-Frank-Organization/frank/internal/connector/request"
)

// TestCT_C01 binds CT-C01; selectors r9 C01 and registered A-M10-SEAM amendment; flags MISMATCH, CN/CP, star.
func TestCT_C01(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	encode := func(buildInfo any) error {
		_, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "hello", Seq: "1", Body: map[string]any{"pid": 1, "build_info": buildInfo}})
		return err
	}
	valid := map[string]any{"version": "dev", "commit": "unknown", "built_at": "unknown"}
	missing := map[string]any{"version": "dev", "commit": "unknown"}
	extra := map[string]any{"version": "dev", "commit": "unknown", "built_at": "unknown", "fourth": "x"}
	oversizeMember := map[string]any{"version": string(make([]byte, 65)), "commit": "unknown", "built_at": "unknown"}
	validErr, missingErr, extraErr, stringErr, oversizeErr := encode(valid), encode(missing), encode(extra), encode("dev"), encode(oversizeMember)
	contract(t, validErr == nil && missingErr != nil && extraErr != nil && stringErr != nil && oversizeErr != nil, explain("hello build_info bounds are not closed: valid=%v missing=%v extra=%v string=%v oversize=%v", validErr, missingErr, extraErr, stringErr, oversizeErr))
}

// TestCT_C02 binds CT-C02; selectors r9 C02 and S-M8v1 plus A-M10 additive registration; flags MISMATCH, CP, star.
func TestCT_C02(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	hasQuery := registry.Has(appipc.ChannelCtrlC, "epoch_query")
	hasUpdate := registry.Has(appipc.ChannelCtrlC, "epoch_update")
	unknown := []byte(`{"v":1,"chan":"ctrl-c","type":"future_epoch_query","seq":"1","body":{}}`)
	_, unknownErr := registry.Decode(unknown)
	fixture := newAppFixture(t)
	query, err := appipc.MarshalJCS(map[string]any{"v": 1, "chan": "ctrl-c", "type": "epoch_query", "seq": "1", "run_id": fixture.runID, "turn_epoch": "9", "body": map[string]any{"run_id": fixture.runID, "turn_epoch": "9"}})
	if err != nil {
		t.Fatal(err)
	}
	var effectOK bool
	var effectErr error
	if receiver, ok := any(scheduler.New(fixture.applier)).(epochQueryReceiver); ok {
		var reply []byte
		reply, effectErr = receiver.ReceiveEpochQuery(fixture.ctx, query)
		if effectErr == nil {
			decoded, decodeErr := registry.Decode(reply)
			body, bodyOK := decoded.Body.(*appipc.EpochUpdateBody)
			effectOK = decodeErr == nil && bodyOK && body.RunID == fixture.runID && body.TurnEpoch == "1"
		}
	}
	contract(t, hasQuery && hasUpdate && unknownErr != nil && effectOK,
		explain("CTRL-C epoch resolution: query=%v update=%v unknown=%v effect=%v/%v", hasQuery, hasUpdate, unknownErr, effectOK, effectErr))
}

// TestCT_C03 binds CT-C03; selectors r9 C03 and S-M8v1 disposition forms; flags MISMATCH, CP/CN, star.
func TestCT_C03(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	cases := []outcome.Outcome{
		{Kind: outcome.SentCompleted, Digests: &outcome.Digests{FrozenCore: digest('a'), ProviderLoweredTools: digest('b')}},
		{Kind: outcome.TransportFailed, Digests: &outcome.Digests{FrozenCore: digest('a'), ProviderLoweredTools: digest('b')}},
		{Kind: outcome.Unknown, Digests: &outcome.Digests{FrozenCore: digest('a'), ProviderLoweredTools: digest('b')}},
		{Kind: outcome.Denied, DenyReason: authorize.PolicyUnavailable, Digests: &outcome.Digests{FrozenCore: digest('a'), ProviderLoweredTools: digest('b')}},
		{Kind: outcome.RejectedLocal, RejectReason: request.MalformedRequest, RefusalStage: outcome.PreFreeze},
		{Kind: outcome.Cancelled, CancelPoint: outcome.PreTransport, Digests: &outcome.Digests{FrozenCore: digest('a'), ProviderLoweredTools: digest('b')}},
	}
	var failures []string
	for index, candidate := range cases {
		produced, buildErr := outcome.AttemptResult("attempt", 1, candidate)
		if buildErr != nil {
			failures = append(failures, explain("case %d producer: %v", index, buildErr))
			continue
		}
		body := map[string]any{"attempt_id": produced.AttemptID, "turn_epoch": produced.TurnEpoch.String(), "disposition": produced.Disposition}
		if produced.RefusalStage != "" {
			body["refusal_stage"] = produced.RefusalStage
		}
		if candidate.RejectReason != "" {
			body["reject_reason"] = candidate.RejectReason
		}
		if candidate.DenyReason != "" {
			body["deny_reason"] = candidate.DenyReason
		}
		if candidate.CancelPoint != "" {
			body["cancel_point"] = candidate.CancelPoint
		}
		encoded, encodeErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "attempt_result", Seq: "1", Body: body})
		if encodeErr != nil {
			failures = append(failures, explain("%q: %v", produced.Disposition, encodeErr))
		} else if _, decodeErr := registry.Decode(encoded); decodeErr != nil {
			failures = append(failures, explain("%q decode: %v", produced.Disposition, decodeErr))
		}
	}
	_, unknownErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "attempt_result", Seq: "1", Body: map[string]any{"attempt_id": "attempt", "turn_epoch": "1", "disposition": "future_disposition"}})
	fixture := newAppFixture(t)
	fixture.seedAttempt(t, "attempt")
	encoded, _ := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "attempt_result", Seq: "1", Body: map[string]any{"attempt_id": "attempt", "turn_epoch": "1", "disposition": "sent_completed", "frozen_core_digest": digest('a'), "provider_lowered_tools_digest": digest('b')}})
	var effectOK bool
	var effectErr error
	host := settle.New(fixture.applier)
	if receiver, ok := any(host).(attemptResultReceiver); ok {
		_, effectErr = receiver.ReceiveAttemptResult(fixture.ctx, encoded)
		if effectErr == nil {
			row, queryErr := host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "attempt")
			effectOK = queryErr == nil && row.State == settle.RowPresent && row.FrozenCoreDigest != nil && *row.FrozenCoreDigest == digest('a')
		}
	}
	contract(t, len(failures) == 0 && unknownErr != nil && effectOK,
		explain("connector disposition decode-to-row: forms=%v unknown=%v effect=%v/%v", failures, unknownErr, effectOK, effectErr))
}

// TestCT_C04 binds CT-C04; selectors r9 C04 and A-M8-BE refusal_stage amendment; flags MISMATCH, MASTER/CP/CN, latent.
func TestCT_C04(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	encode := func(disposition string, stage any) error {
		body := map[string]any{"attempt_id": "attempt", "turn_epoch": "1", "disposition": disposition}
		if disposition == "rejected_local" {
			body["reject_reason"] = "malformed_request"
		}
		if stage != nil {
			body["refusal_stage"] = stage
		}
		_, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "attempt_result", Seq: "1", Body: body})
		return err
	}
	pre, post := encode("rejected_local", "pre_freeze"), encode("rejected_local", "post_freeze")
	missing, foreign, invalid := encode("rejected_local", nil), encode("sent_completed", "pre_freeze"), encode("rejected_local", "elsewhere")
	fixture := newAppFixture(t)
	fixture.seedAttempt(t, "rejected")
	fixture.seedAttempt(t, "plain")
	host := settle.New(fixture.applier)
	rejectedRequest := settle.AttemptResultRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "rejected", TurnEpoch: "1", Disposition: "rejected_local", At: 2}
	requestCarried := setStringField(&rejectedRequest, "RefusalStage", "pre_freeze")
	decision, recordErr := host.RecordAttemptResult(fixture.ctx, rejectedRequest)
	rejectedRow, rejectedErr := host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "rejected")
	rowStage, rowCarried := optionalStringField(rejectedRow, "RefusalStage")
	plainDecision, plainErr := host.RecordAttemptResult(fixture.ctx, settle.AttemptResultRequest{RunID: fixture.runID, TurnID: fixture.turnID, AttemptID: "plain", TurnEpoch: "1", Disposition: "sent_completed", At: 3})
	plainRow, plainQueryErr := host.QueryAttempt(fixture.ctx, fixture.runID, fixture.turnID, "plain")
	plainStage, plainCarried := optionalStringField(plainRow, "RefusalStage")
	persisted := requestCarried && rowCarried && rowStage != nil && *rowStage == "pre_freeze" && plainCarried && plainStage == nil
	contract(t, structHasField(appipc.AttemptResultBody{}, "RefusalStage") && persisted && decision == settle.CarriageRecorded && plainDecision == settle.CarriageRecorded && recordErr == nil && rejectedErr == nil && plainErr == nil && plainQueryErr == nil && pre == nil && post == nil && missing != nil && foreign != nil && invalid != nil,
		explain("refusal_stage iff/round-trip mismatch: persisted=%v decision=%s/%v plain=%s/%v query=%v/%v pre=%v post=%v missing=%v foreign=%v invalid=%v", persisted, decision, recordErr, plainDecision, plainErr, rejectedErr, plainQueryErr, pre, post, missing, foreign, invalid))
}

// TestCT_C05 binds CT-C05; selectors r9 C05 and S-M10 bootstrap mirror; flags MISMATCH, CP, latent.
func TestCT_C05(t *testing.T) {
	valid, gate, validErr := lockedManifestFixture(t, "provider-main")
	_, validDecodeErr := manifest.DecodeFrozen(valid.Bytes, valid.Digest, gate)
	_, _, absentErr := lockedManifestFixture(t, "")
	badRef := "Provider_Main"
	_, _, badErr := lockedManifestFixture(t, badRef)
	duplicated := bytes.Replace(valid.Bytes, []byte(`"credential_ref":"provider-main"`), []byte(`"credential_ref":"provider-main","credential_ref":"provider-main"`), 1)
	_, duplicateErr := manifest.DecodeFrozen(duplicated, bytesDigest(duplicated), gate)
	contract(t, validErr == nil && validDecodeErr == nil && credentials.ValidReference("provider-main") && absentErr != nil && !credentials.ValidReference(badRef) && badErr != nil && duplicateErr != nil,
		explain("credential-ref admission mirror: valid=%v/%v absent=%v bad=%v duplicate=%v", validErr, validDecodeErr, absentErr, badErr, duplicateErr))
}

// TestCT_C06 binds CT-C06; selectors r9 C06 and S-M10 D§13; flags MISMATCH, CP, latent.
func TestCT_C06(t *testing.T) {
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	run, epoch := "envelope-run", "1"
	_, runErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "epoch_update", Seq: "1", RunID: &run, TurnEpoch: &epoch, Body: appipc.EpochUpdateBody{RunID: "different-run", TurnEpoch: epoch}})
	_, epochErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "epoch_update", Seq: "1", RunID: &run, TurnEpoch: &epoch, Body: appipc.EpochUpdateBody{RunID: run, TurnEpoch: "2"}})
	contract(t, runErr != nil && epochErr != nil, explain("CTRL-C envelope/body identity coherence: run=%v epoch=%v", runErr, epochErr))
}

// TestCT_C07 binds CT-C07; selectors r9 C07 and S-M10/S-CP connector spawn; flags UNWIRED, CP/CN, star.
func TestCT_C07(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	runtimeDir := t.TempDir()
	connectorDir := filepath.Join(runtimeDir, "connector")
	if err := os.Mkdir(connectorDir, 0o700); err != nil {
		t.Fatal(err)
	}
	connectorBin := filepath.Join(runtimeDir, "frank-connector")
	build := exec.CommandContext(ctx, "go", "build", "-o", connectorBin, "./cmd/frank-connector")
	build.Dir = repoRoot(t)
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build real connector: %v: %s", err, output)
	}
	credentialPath := filepath.Join(connectorDir, "credentials.json")
	catalogPath := filepath.Join(connectorDir, "catalog.json")
	policyPath := filepath.Join(connectorDir, "policy.json")
	artifacts := map[string]string{
		credentialPath: `{"entries":{"credential-ref":{"secret":"seam-secret"}},"schema":"m8.credentials.v1"}`,
		catalogPath:    `{"lanes":[{"auth":{"auth_header_name":"x-openai-auth","auth_scheme":"bearer"},"compat_mode":"openai-responses","cost":{"effective_time":"2026-07-17T00:00:00Z","input":1,"output":10},"endpoint":"https://api.openai.com/v1/responses","lane_id":"lane-codex-1","limits":{"context":200000,"max_output":100000},"method":"POST","model_id":"gpt-5","observed_at":"2026-07-17T00:00:00Z","profile_facts":{"endpoint_kind":"coding","region":"us-east"},"provider_id":"openai","reasoning":{"effort_levels":["low","medium","high"],"replay_kind":"opaque_item","supported":true},"serving_profile_id":"codex-default","source":"seeded","tool_use":{"strict_schema":true,"supported":true},"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true}}],"schema":"m8.lane_catalog.v1"}`,
		policyPath:     `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`,
	}
	for path, value := range artifacts {
		if err := os.WriteFile(path, []byte(value), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	command := exec.CommandContext(ctx, "go", "run", "./cmd/frank-app", "--state-dir", runtimeDir, "run", "start", "--goal", "probe", "--lane", "lane-codex-1", "--credential-ref", "credential-ref", "--workspace-root", repoRoot(t))
	command.Dir = repoRoot(t)
	command.Env = append(os.Environ(),
		"FRANK_CONNECTOR_BIN="+connectorBin,
		"FRANK_CONNECTOR_CREDENTIAL="+credentialPath,
		"FRANK_CONNECTOR_CATALOG="+catalogPath,
		"FRANK_CONNECTOR_POLICY="+policyPath,
	)
	output, err := command.CombinedOutput()
	contract(t, err == nil && strings.Contains(string(output), "CONNECTOR_READY\n"), explain("frank-app run start did not observe the real connector READY token: %v: %s", err, output))
}

// TestCT_C08 binds CT-C08; selectors r9 C08 and S-M10 assign digests; flags MISMATCH, CP/CN, latent.
func TestCT_C08(t *testing.T) {
	frozen, _, err := lockedManifestFixture(t, "provider-main")
	if err != nil {
		t.Fatal(err)
	}
	mutated, _, mutatedErr := lockedManifestFixture(t, "provider-backup")
	if mutatedErr != nil {
		t.Fatal(mutatedErr)
	}
	var body appipc.ConnectorAssignBody
	var buildErr error
	var built bool
	if builder, ok := any(frozen).(connectorAssignBuilder); ok {
		body, buildErr = builder.ConnectorAssign()
		built = true
	} else if builder, ok := any(&frozen).(connectorAssignBuilder); ok {
		body, buildErr = builder.ConnectorAssign()
		built = true
	}
	verbatim := built && buildErr == nil && body.RunManifestDigest == frozen.Digest && body.PolicyDigest == frozen.Manifest.PolicyDigest && body.LaneCatalogDigest == frozen.Manifest.ProviderLane.LaneCatalogDigest
	mutatedControl := mutated.Digest != frozen.Digest && mutated.Manifest.PolicyDigest == frozen.Manifest.PolicyDigest && mutated.Manifest.ProviderLane.LaneCatalogDigest == frozen.Manifest.ProviderLane.LaneCatalogDigest
	contract(t, verbatim && mutatedControl, explain("connector_assign digest copy: built=%v err=%v body=%+v mutated_control=%v", built, buildErr, body, mutatedControl))
}

// TestCT_C09 binds CT-C09; selectors r9 C09 and A-M7-BROKER CI-1; flags UNWIRED, MASTER/m-7, star.
func TestCT_C09(t *testing.T) {
	fixture := newAppFixture(t)
	runtimeDir := t.TempDir()
	configHome := filepath.Join(t.TempDir(), "broker-config")
	if err := os.Mkdir(configHome, 0o700); err != nil {
		t.Fatal(err)
	}
	credentialPath := filepath.Join(configHome, "credential")
	if err := os.WriteFile(credentialPath, []byte("credential-sentinel\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configHome, "broker.json"), []byte(`{"credential_file":"credential"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	brokerBin := filepath.Join(t.TempDir(), "frank-broker")
	build := exec.Command("go", "build", "-o", brokerBin, "./cmd/frank-broker")
	build.Dir = repoRoot(t)
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build frank-broker: %v: %s", err, output)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	process, err := supervisor.LaunchBroker(ctx, supervisor.BrokerLaunch{
		BinaryPath: brokerBin, RuntimeDir: runtimeDir, ConfigHome: configHome,
		RunID: fixture.runID, ControlToken: "token", At: 1, Client: brokerclient.New(fixture.applier),
		Controller: supervisor.New(fixture.applier), InstanceID: "ct-c09-broker",
	})
	if process != nil {
		t.Cleanup(func() {
			closeCtx, closeCancel := context.WithTimeout(context.Background(), time.Second)
			defer closeCancel()
			_ = process.Close(closeCtx)
		})
	}
	if err != nil || process == nil || process.State() != supervisor.WorkerReady {
		contract(t, false, explain("real broker did not reach READY: process=%#v err=%v", process, err))
		return
	}
	contract(t, !strings.Contains(process.Nonce(), "credential-sentinel"), explain("broker credential escaped through READY: %q", process.Nonce()))
	entries, readErr := os.ReadDir(runtimeDir)
	contract(t, readErr == nil && len(entries) == 1 && entries[0].Name() == "broker-control.sock",
		explain("READY must precede no broker-durable state beyond its socket: entries=%v err=%v", entries, readErr))
	session, err := brokerclient.New(fixture.applier).Establish(ctx, brokerclient.ControlRequest{RunID: fixture.runID, RuntimeDir: runtimeDir, ControlToken: "token", At: 2})
	if session != nil {
		defer session.Close()
	}
	contract(t, err == nil && session != nil, explain("brokerclient CI-1 dial was unanswered: %v", err))
}

type epochQueryReceiver interface {
	ReceiveEpochQuery(context.Context, []byte) ([]byte, error)
}

type attemptResultReceiver interface {
	ReceiveAttemptResult(context.Context, []byte) (settle.CarriageDecision, error)
}

type connectorAssignBuilder interface {
	ConnectorAssign() (appipc.ConnectorAssignBody, error)
}

func setStringField(target any, name, value string) bool {
	field := reflect.ValueOf(target).Elem().FieldByName(name)
	if !field.IsValid() || !field.CanSet() {
		return false
	}
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
		return true
	case reflect.Pointer:
		if field.Type().Elem().Kind() != reflect.String {
			return false
		}
		copy := value
		field.Set(reflect.ValueOf(&copy).Convert(field.Type()))
		return true
	default:
		return false
	}
}

func optionalStringField(value any, name string) (*string, bool) {
	field := reflect.ValueOf(value).FieldByName(name)
	if !field.IsValid() {
		return nil, false
	}
	switch field.Kind() {
	case reflect.String:
		text := field.String()
		if text == "" {
			return nil, true
		}
		return &text, true
	case reflect.Pointer:
		if field.IsNil() || field.Type().Elem().Kind() != reflect.String {
			return nil, true
		}
		text := field.Elem().String()
		return &text, true
	default:
		return nil, false
	}
}

// TestCT_C10 binds CT-C10; selectors r9 C10 and connector sequence base; flags MISMATCH, CP, benign.
func TestCT_C10(t *testing.T) {
	left, right := net.Pipe()
	defer left.Close()
	defer right.Close()
	peer, err := testutil.NewFakeConnector(left)
	if err != nil {
		t.Fatal(err)
	}
	result := make(chan struct {
		seq string
		err error
	}, 1)
	go func() {
		seq, sendErr := peer.Send(context.Background(), testutil.Outbound{Type: "ping", Body: appipc.EmptyBody{}})
		result <- struct {
			seq string
			err error
		}{seq, sendErr}
	}()
	if _, err := appipc.ReadFrame(right); err != nil {
		t.Fatal(err)
	}
	select {
	case got := <-result:
		contract(t, got.err == nil && got.seq == "1", explain("test connector starts at seq %q, want 1", got.seq))
	case <-time.After(time.Second):
		t.Fatal("fake connector send blocked")
	}
}
