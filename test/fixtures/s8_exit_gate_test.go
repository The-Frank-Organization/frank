package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestS8ExitGateFreshGenesisActivationAndDogfoodLegs(t *testing.T) {
	root := t.TempDir()
	if err := store.Init(root, s8ConfigSources(t, false)); err != nil {
		t.Fatalf("fresh genesis: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if config.PresentLayers(s8PinnedStore(t, root))["observe"] {
		t.Fatal("fresh genesis unexpectedly activated observe")
	}

	v6 := s8FieldspecV6Bytes(t)
	s8CommitOperatorConfigChange(t, st, "fieldspec", v6)
	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v7 registry: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "fieldspec", v7)
	afterV7 := s8PinnedStore(t, root)
	t.Logf("v7_transition_new_digest=%s", afterV7.Digest)

	engineBytes := append([]byte(nil), afterV7.Members["engine"]...)
	var engineDoc map[string]any
	if err := json.Unmarshal(engineBytes, &engineDoc); err != nil {
		t.Fatalf("decode engine member: %v", err)
	}
	engineDoc["present_layers"].(map[string]any)["observe"] = true
	engineBytes, err = json.Marshal(engineDoc)
	if err != nil {
		t.Fatalf("marshal activated engine: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "engine", engineBytes)

	activated := s8PinnedStore(t, root)
	if !config.PresentLayers(activated)["observe"] {
		t.Fatal("observe activation absent after restart load")
	}
	restarted, err := store.Open(root)
	if err != nil {
		t.Fatalf("restart Open: %v", err)
	}
	if err := restarted.ValidateGenesis(activated); err != nil {
		t.Fatalf("phase-0 validation after activation: %v", err)
	}

	meta := seat.SeatMeta{Name: "s8.implementer", Role: "implementer"}
	renderEnv := fieldspec.RenderEnv{ConfigDigest: activated.Digest, PresentLayers: config.PresentLayers(activated)}
	base := record.Record{Headers: map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium",
		"EVIDENCE_TARGET": "E1", "SUBJECT": "dogfood observed send", "TO": `["recipient.planner"]`,
		"ACTIONS_GIT_REF": "branch@candidate", "FINAL_GIT_STATUS_SHORT": "none - clean tree",
	}, Body: "done"}

	failing := s8SubmitWithObservation(t, restarted, activated.Registry, meta, renderEnv, base, observe.PredicateResult{ID: "done-predicate", Predicate: observe.Fail}, "exit-leg-1")
	if failing.rec.Envelope.DeliveryState != record.Rejected || failing.rec.Headers["achieved_evidence"] != "E0" || failing.rec.Headers["record_integrity"] != "observed" {
		t.Fatalf("exit leg 1 record = %#v", failing.rec)
	}
	if failing.intents != nil {
		t.Fatalf("exit leg 1 reached delivery intents: %#v", failing.intents)
	}
	if _, err := restarted.Commit(failing.rec, failing.intents); err != nil {
		t.Fatalf("commit exit leg 1 terminal: %v", err)
	}
	if projected, err := restarted.Project("recipient.planner"); err != nil || len(projected) != 0 {
		t.Fatalf("exit leg 1 recipient projection = %v, err=%v", projected, err)
	}

	passing := s8SubmitWithObservation(t, restarted, activated.Registry, meta, renderEnv, base, observe.PredicateResult{ID: "done-predicate", Predicate: observe.Pass}, "exit-leg-2")
	if passing.rec.Envelope.DeliveryState != record.Accepted || passing.rec.Headers["achieved_evidence"] != "E1" || passing.rec.Headers["target_gap_result"] != "met" || passing.rec.Headers["record_integrity"] != "observed" || passing.rec.Headers["attestation_source"] != "conductor" {
		t.Fatalf("exit leg 2 record = %#v", passing.rec)
	}
	acceptedID, err := restarted.Commit(passing.rec, passing.intents)
	if err != nil {
		t.Fatalf("commit exit leg 2: %v", err)
	}
	projected, err := restarted.Project("recipient.planner")
	if err != nil || len(projected) != 1 || projected[0] != acceptedID {
		t.Fatalf("exit leg 2 recipient projection = %v, want %s, err=%v", projected, acceptedID, err)
	}
}

func TestS8ProductionDogfoodRejectsFalseDoneAndNamesPredicate(t *testing.T) {
	if os.Getenv("FRANK_DOGFOOD_NESTED") != "" {
		return
	}
	root := t.TempDir()
	sources := s8ConfigSources(t, false)
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Second)
	t.Cleanup(cancel)
	bin := buildFrank(t, ctx)
	initCmd := exec.CommandContext(ctx, bin,
		"-root", root, "-registry", sources["fieldspec"], "-engine-config", sources["engine"], "-catalog", sources["catalog"], "-init",
	)
	if out, err := initCmd.CombinedOutput(); err != nil {
		t.Fatalf("production init: %v: %s", err, out)
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	h := &s4ShimHarness{
		root: root, ctx: ctx, cancel: cancel, bin: bin,
		sock: filepath.Join(os.TempDir(), fmt.Sprintf("frank-s8-%d.sock", time.Now().UnixNano())), mgr: mgr,
	}
	t.Cleanup(func() {
		h.stop(t)
		_ = os.Remove(h.sock)
	})
	operatorCred, err := mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	laneCred, err := mgr.Mint("s8.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("Mint lane: %v", err)
	}

	h.start(t)
	operator := h.dial(t, operatorCred)
	v6 := s8FieldspecV6Bytes(t)
	h.submit(t, operator, s8ConfigChangeRecord(t, root, "fieldspec", v6))
	_ = operator.Close()
	h.stop(t)

	h.start(t)
	operator = h.dial(t, operatorCred)
	v6Lane := h.dial(t, laneCred)
	v6Describe, err := v6Lane.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools v6: %v", err)
	}
	_ = v6Lane.Close()
	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v7 registry: %v", err)
	}
	h.submit(t, operator, s8ConfigChangeRecord(t, root, "fieldspec", v7))
	_ = operator.Close()
	h.stop(t)

	h.start(t)
	operator = h.dial(t, operatorCred)
	pinned := s8PinnedStore(t, root)
	var engineDoc map[string]any
	if err := json.Unmarshal(pinned.Members["engine"], &engineDoc); err != nil {
		t.Fatalf("decode engine: %v", err)
	}
	engineDoc["present_layers"].(map[string]any)["observe"] = true
	engineBytes, err := json.Marshal(engineDoc)
	if err != nil {
		t.Fatalf("marshal engine: %v", err)
	}
	h.submit(t, operator, s8ConfigChangeRecord(t, root, "engine", engineBytes))
	_ = operator.Close()
	h.stop(t)

	h.start(t)
	lane := h.dial(t, laneCred)
	defer func() { _ = lane.Close() }()
	rec := s4Relay("recipient.planner", "", "production false done")
	rec.Headers["ACTIONS_GIT_REF"] = "branch@missing"
	rec.Headers["FINAL_GIT_STATUS_SHORT"] = "none - clean tree"
	params, err := fieldspec.CanonicalMarshal(map[string]string{"lane_ref": "repo", "path": "missing-s8-done", "expect": "line:done"})
	if err != nil {
		t.Fatalf("marshal claim params: %v", err)
	}
	claims, err := fieldspec.CanonicalMarshal([]map[string]string{{"claim_ref": "done-predicate", "check_id": "read-file", "params": params}})
	if err != nil {
		t.Fatalf("marshal claims: %v", err)
	}
	rec.Headers["executable_claims"] = claims
	staleRec := rec
	staleRec.Headers = make(map[string]string, len(rec.Headers))
	for key, value := range rec.Headers {
		staleRec.Headers[key] = value
	}
	staleRec.Headers["executable_claims"] = s8Claims(t, map[string]string{"claim_ref": "stale-claim", "check_id": "not-registered", "params": `{}`})
	staleResult, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: staleRec, FormDigest: v6Describe.FormDigest}))
	if err != nil {
		t.Fatalf("stale-v6-form submit: %v", err)
	}
	var staleOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(staleResult, &staleOutcome); err != nil {
		t.Fatalf("decode stale outcome: %v", err)
	}
	staleRead := readRelayRaw(t, ctx, lane, staleOutcome.RelayID)
	if staleOutcome.State != record.Rejected || !bytes.Contains(staleRead, []byte("form_digest:re-render")) || bytes.Contains(staleRead, []byte("unknown-check")) {
		t.Fatalf("stale-v6-form outcome = %+v", staleOutcome)
	}
	describe, err := lane.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools: %v", err)
	}
	result, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("production submit: %v", err)
	}
	var outcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode outcome: %v", err)
	}
	if outcome.State != record.Rejected {
		t.Fatalf("production false-done state = %q, want rejected", outcome.State)
	}
	read := readRelayRaw(t, ctx, lane, outcome.RelayID)
	if !bytes.Contains(read, []byte("done-predicate")) {
		t.Fatalf("production false-done read = %s, want failing predicate", read)
	}

	passing := s4Relay("recipient.planner", "", "production observed done")
	passing.Headers["ACTIONS_GIT_REF"] = "branch@present"
	passing.Headers["FINAL_GIT_STATUS_SHORT"] = "none - clean tree"
	passParams, err := fieldspec.CanonicalMarshal(map[string]string{"lane_ref": "repo", "path": "test/fixtures/s8_exit_gate_test.go", "expect": "line:package fixtures_test"})
	if err != nil {
		t.Fatalf("marshal passing params: %v", err)
	}
	passing.Headers["executable_claims"], err = fieldspec.CanonicalMarshal([]map[string]string{{"claim_ref": "done-predicate", "check_id": "read-file", "params": passParams}})
	if err != nil {
		t.Fatalf("marshal passing claims: %v", err)
	}
	passResult, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: passing, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("production passing submit: %v", err)
	}
	var passOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(passResult, &passOutcome); err != nil {
		t.Fatalf("decode passing outcome: %v", err)
	}
	if passOutcome.State != record.Accepted {
		t.Fatalf("production passing state = %q, want accepted", passOutcome.State)
	}
	passRead := readRelayRaw(t, ctx, lane, passOutcome.RelayID)
	for _, want := range [][]byte{[]byte(`"achieved_evidence":"E1"`), []byte(`"target_gap_result":"met"`), []byte(`"record_integrity":"observed"`), []byte(`"attestation_source":"conductor"`)} {
		if !bytes.Contains(passRead, want) {
			t.Fatalf("production passing read = %s, want %s", passRead, want)
		}
	}
	suiteParams, err := fieldspec.CanonicalMarshal(map[string]string{"target": "dogfood-battery", "expect_green": "true"})
	if err != nil {
		t.Fatalf("marshal suite params: %v", err)
	}
	suiteClaims, err := fieldspec.CanonicalMarshal([]map[string]string{{"claim_ref": "dogfood-battery-green", "check_id": "run-suite", "params": suiteParams}})
	if err != nil {
		t.Fatalf("marshal suite claim: %v", err)
	}
	suiteRecord := s4Relay("recipient.planner", "", "production governed suite")
	suiteRecord.Headers["EVIDENCE_TARGET"] = "E2"
	suiteRecord.Headers["ACTIONS_GIT_REF"] = "branch@candidate"
	suiteRecord.Headers["FINAL_GIT_STATUS_SHORT"] = "none - clean tree"
	suiteRecord.Headers["executable_claims"] = suiteClaims
	suiteResult, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: suiteRecord, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("production suite submit: %v", err)
	}
	var suiteOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(suiteResult, &suiteOutcome); err != nil {
		t.Fatalf("decode suite outcome: %v", err)
	}
	if suiteOutcome.State != record.Accepted {
		t.Fatalf("production suite state = %q, want accepted; read=%s", suiteOutcome.State, readRelayRaw(t, ctx, lane, suiteOutcome.RelayID))
	}
	suiteRead := readRelayRaw(t, ctx, lane, suiteOutcome.RelayID)
	for _, want := range [][]byte{[]byte(`"achieved_evidence":"E2"`), []byte(`"target_gap_result":"met"`), []byte(`"record_integrity":"observed"`)} {
		if !bytes.Contains(suiteRead, want) {
			t.Fatalf("production suite read = %s, want %s", suiteRead, want)
		}
	}
	var suiteEnvelope struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(suiteRead, &suiteEnvelope); err != nil {
		t.Fatalf("decode production suite read: %v", err)
	}
	wantResult := `[{"check_id":"run-suite","claim_ref":"dogfood-battery-green","outcome":"pass"}]`
	if suiteEnvelope.Record.Headers["executable_claim_results"] != wantResult {
		t.Fatalf("production suite results = %q, want %q", suiteEnvelope.Record.Headers["executable_claim_results"], wantResult)
	}

	falseSuiteParams, err := fieldspec.CanonicalMarshal(map[string]string{"target": "dogfood-battery", "expect_green": "false"})
	if err != nil {
		t.Fatalf("marshal false suite params: %v", err)
	}
	falseSuiteClaims, err := fieldspec.CanonicalMarshal([]map[string]string{{"claim_ref": "dogfood-battery-false-done", "check_id": "run-suite", "params": falseSuiteParams}})
	if err != nil {
		t.Fatalf("marshal false suite claim: %v", err)
	}
	falseSuiteRecord := s4Relay("suite-negative.planner", "", "production governed suite false done")
	falseSuiteRecord.Headers["EVIDENCE_TARGET"] = "E2"
	falseSuiteRecord.Headers["ACTIONS_GIT_REF"] = "branch@candidate"
	falseSuiteRecord.Headers["FINAL_GIT_STATUS_SHORT"] = "none - clean tree"
	falseSuiteRecord.Headers["executable_claims"] = falseSuiteClaims
	falseSuiteResult, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: falseSuiteRecord, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("production false suite submit: %v", err)
	}
	var falseSuiteOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(falseSuiteResult, &falseSuiteOutcome); err != nil {
		t.Fatalf("decode false suite outcome: %v", err)
	}
	falseSuiteRead := readRelayRaw(t, ctx, lane, falseSuiteOutcome.RelayID)
	if falseSuiteOutcome.State != record.Rejected || !bytes.Contains(falseSuiteRead, []byte("dogfood-battery-false-done")) {
		t.Fatalf("production false suite outcome = %+v; read=%s", falseSuiteOutcome, falseSuiteRead)
	}
	var falseSuiteEnvelope struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(falseSuiteRead, &falseSuiteEnvelope); err != nil {
		t.Fatalf("decode production false suite read: %v", err)
	}
	wantFalseResult := `[{"check_id":"run-suite","claim_ref":"dogfood-battery-false-done","outcome":"fail"}]`
	if falseSuiteEnvelope.Record.Headers["executable_claim_results"] != wantFalseResult {
		t.Fatalf("production false suite results = %q, want %q", falseSuiteEnvelope.Record.Headers["executable_claim_results"], wantFalseResult)
	}
	projectionStore, err := store.Open(root)
	if err != nil {
		t.Fatalf("open store for false suite projection: %v", err)
	}
	projectedFalseSuite, err := projectionStore.Project("suite-negative.planner")
	if err != nil || len(projectedFalseSuite) != 0 {
		t.Fatalf("false suite recipient projection = %v, err=%v", projectedFalseSuite, err)
	}

}

func TestS8ProductionAbsenceFloorObservesGitAndStampsTargetGap(t *testing.T) {
	root := t.TempDir()
	if err := store.Init(root, s8ConfigSources(t, false)); err != nil {
		t.Fatalf("production absence-floor init: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open production absence-floor store: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "fieldspec", s8FieldspecV6Bytes(t))
	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v7 registry: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "fieldspec", v7)
	pinned := s8PinnedStore(t, root)
	var engineDoc map[string]any
	if err := json.Unmarshal(pinned.Members["engine"], &engineDoc); err != nil {
		t.Fatalf("decode engine: %v", err)
	}
	engineDoc["present_layers"].(map[string]any)["observe"] = true
	engineBytes, err := json.Marshal(engineDoc)
	if err != nil {
		t.Fatalf("marshal engine: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "engine", engineBytes)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	t.Cleanup(cancel)
	h := &s4ShimHarness{
		root: root, ctx: ctx, cancel: cancel, bin: buildFrank(t, ctx),
		sock: filepath.Join(os.TempDir(), fmt.Sprintf("frank-s8-absence-%d.sock", time.Now().UnixNano())),
	}
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("open production absence-floor seats: %v", err)
	}
	h.mgr = mgr
	t.Cleanup(func() {
		h.stop(t)
		_ = os.Remove(h.sock)
	})
	laneCred, err := mgr.Mint("s8.implementer", "implementer", false)
	if err != nil {
		t.Fatalf("mint production absence-floor lane: %v", err)
	}
	h.start(t)
	lane := h.dial(t, laneCred)
	defer func() { _ = lane.Close() }()
	describe, err := lane.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("describe production absence-floor tools: %v", err)
	}

	headOut, err := exec.CommandContext(ctx, "git", "-C", fixtureRepoRoot(t), "rev-parse", "HEAD").Output()
	if err != nil {
		t.Fatalf("resolve production base-observation head: %v", err)
	}
	statusOut, err := exec.CommandContext(ctx, "git", "-C", fixtureRepoRoot(t), "status", "--porcelain", "--untracked-files=normal").Output()
	if err != nil {
		t.Fatalf("read production base-observation status: %v", err)
	}
	declaredStatus := strings.TrimSpace(string(statusOut))
	if declaredStatus == "" {
		declaredStatus = "none - clean tree"
	}
	absentClaims := s4Relay("absence-floor.planner", "", "production absence floor")
	absentClaims.Headers["EVIDENCE_TARGET"] = "E4"
	absentClaims.Headers["ACTIONS_GIT_REF"] = "s8-observe-spine@" + strings.TrimSpace(string(headOut))
	absentClaims.Headers["FINAL_GIT_STATUS_SHORT"] = declaredStatus
	absentResult, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: absentClaims, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("production absence-floor submit: %v", err)
	}
	var absentOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(absentResult, &absentOutcome); err != nil {
		t.Fatalf("decode production absence-floor outcome: %v", err)
	}
	absentRead := readRelayRaw(t, ctx, lane, absentOutcome.RelayID)
	if absentOutcome.State != record.Accepted {
		t.Fatalf("production absence-floor state = %q, want accepted; read=%s", absentOutcome.State, absentRead)
	}
	for _, want := range [][]byte{[]byte(`"achieved_evidence":"E1"`), []byte(`"target_gap_result":"target_gt_achieved"`), []byte(`"record_integrity":"observed"`)} {
		if !bytes.Contains(absentRead, want) {
			t.Fatalf("production absence-floor read = %s, want %s", absentRead, want)
		}
	}

	falseStatus := absentClaims
	falseStatus.Headers = make(map[string]string, len(absentClaims.Headers))
	for key, value := range absentClaims.Headers {
		falseStatus.Headers[key] = value
	}
	if declaredStatus == "none - clean tree" {
		falseStatus.Headers["FINAL_GIT_STATUS_SHORT"] = "M deliberately-not-present"
	} else {
		falseStatus.Headers["FINAL_GIT_STATUS_SHORT"] = "none - clean tree"
	}
	falseStatus.Headers["SUBJECT"] = "production absence floor false status"
	falseStatusResult, err := lane.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: falseStatus, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("production absence-floor false-status submit: %v", err)
	}
	var falseStatusOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(falseStatusResult, &falseStatusOutcome); err != nil {
		t.Fatalf("decode production absence-floor false-status outcome: %v", err)
	}
	falseStatusRead := readRelayRaw(t, ctx, lane, falseStatusOutcome.RelayID)
	if falseStatusOutcome.State != record.Rejected || !bytes.Contains(falseStatusRead, []byte("FINAL_GIT_STATUS_SHORT")) {
		t.Fatalf("production absence-floor false status = %+v; read=%s", falseStatusOutcome, falseStatusRead)
	}
}

func TestS8V6ReaderRefusesV7MarkerBeforeContent(t *testing.T) {
	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read v7 registry: %v", err)
	}
	var planted map[string]any
	if err := json.Unmarshal(v7, &planted); err != nil {
		t.Fatalf("decode v7 registry: %v", err)
	}
	planted["fields"] = "content-must-not-be-interpreted"
	plantedBytes, err := json.Marshal(planted)
	if err != nil {
		t.Fatalf("marshal planted registry: %v", err)
	}
	if _, err := fieldspec.Parse(plantedBytes); err == nil {
		t.Fatal("planted fieldspec content unexpectedly valid")
	}
	if err := config.ValidateFieldspecReaderMarker(plantedBytes, "s7a-fieldspec-v5", "s8-fieldspec-v6"); !errors.Is(err, config.ErrConfigLoad) {
		t.Fatalf("v6 reader marker error = %v, want config-load", err)
	}
	if err := config.ValidateFieldspecReaderMarker(plantedBytes, "s7a-fieldspec-v5", "s8-fieldspec-v6", "s8-fieldspec-v7"); err != nil {
		t.Fatalf("v7 marker preflight interpreted planted content: %v", err)
	}
}

func TestS8ExecutableClaimTypedRejects(t *testing.T) {
	lane := t.TempDir()
	reg := observe.NewRegistry(observe.RegistryEnv{Lanes: map[string]string{"repo": lane}})
	root := t.TempDir()
	if err := store.Init(root, s8ConfigSources(t, false)); err != nil {
		t.Fatalf("init typed-reject store: %v", err)
	}
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("open typed-reject store: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "fieldspec", s8FieldspecV6Bytes(t))
	v7, err := os.ReadFile(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read typed-reject v7 registry: %v", err)
	}
	s8CommitOperatorConfigChange(t, st, "fieldspec", v7)
	pinned := s8PinnedStore(t, root)
	meta := seat.SeatMeta{Name: "s8.implementer", Role: "implementer"}
	env := fieldspec.RenderEnv{ConfigDigest: pinned.Digest, PresentLayers: map[string]bool{"observe": true}}
	baseHeaders := map[string]string{
		"PHASE": "SITREP", "AUTHORITY": "report-only", "CEREMONY_TIER": "medium", "EVIDENCE_TARGET": "E1",
		"SUBJECT": "typed claim rejection", "TO": `["recipient.planner"]`,
		"ACTIONS_GIT_REF": "branch@candidate", "FINAL_GIT_STATUS_SHORT": "none - clean tree",
	}
	_, digest := pinned.Registry.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role}, "SITREP", "medium", fieldspec.ClosedGrantState)
	validParams := s8ClaimParams(t, map[string]string{"lane_ref": "repo", "path": "artifact.txt", "expect": "line:x"})
	for _, tc := range []struct {
		name  string
		raw   string
		class string
		ref   string
	}{
		{name: "unknown check", raw: s8Claims(t, map[string]string{"claim_ref": "unknown-claim", "check_id": "not-registered", "params": `{}`}), class: "unknown-check", ref: "unknown-claim"},
		{name: "invalid params", raw: s8Claims(t, map[string]string{"claim_ref": "safe-claim", "check_id": "read-file", "params": s8ClaimParams(t, map[string]string{"lane_ref": "repo", "path": "/secret/path", "expect": "line:x"})}), class: "check-params-invalid", ref: "safe-claim"},
		{name: "duplicate claim ref", raw: s8Claims(t,
			map[string]string{"claim_ref": "duplicate", "check_id": "read-file", "params": validParams},
			map[string]string{"claim_ref": "duplicate", "check_id": "read-file", "params": validParams},
		), class: "duplicate-claim-ref", ref: "duplicate"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, issue := reg.ValidateClaims(tc.raw); issue == nil || issue.Class != tc.class || issue.ClaimRef != tc.ref {
				t.Fatalf("issue = %#v, want %s/%s", issue, tc.ref, tc.class)
			}
			authoritative := reg.EvaluateClaims(tc.raw, observe.Candidate{})
			if authoritative.Predicate != observe.Fail || authoritative.ID != tc.ref || authoritative.FailureClass != tc.class {
				t.Fatalf("authoritative refusal = %#v, want %s/%s", authoritative, tc.ref, tc.class)
			}
			headers := make(map[string]string, len(baseHeaders)+1)
			for key, value := range baseHeaders {
				headers[key] = value
			}
			headers["executable_claims"] = tc.raw
			payload := mustJSONBytes(t, fieldspec.SubmitPayload{Record: record.Record{Headers: headers}, FormDigest: digest})
			got, intents, err := engine.SubmitHandlerWithClaimRegistry(st, pinned.Registry, meta, env, reg)(context.Background(), intake.Cmd{IntakeID: "typed-" + tc.ref, Seat: meta.Name, Role: meta.Role, Payload: payload})
			if err != nil {
				t.Fatalf("typed claim submit: %v", err)
			}
			if got.Envelope.DeliveryState != record.Rejected || intents != nil || !bytes.Contains([]byte(got.Body), []byte(tc.ref+":"+tc.class)) {
				t.Fatalf("typed claim bounce = state %s intents %#v body %q", got.Envelope.DeliveryState, intents, got.Body)
			}
			if bytes.Contains([]byte(got.Body), []byte("/secret/path")) {
				t.Fatalf("typed claim bounce leaked raw path: %q", got.Body)
			}
		})
	}
}

func TestS8ExecutableClaimAggregationPrecedence(t *testing.T) {
	exec := &s8ClaimExecutor{verdicts: map[string]observe.CheckVerdict{
		"pass":      {Outcome: "pass", RungReached: "E2", Predicate: observe.Pass},
		"unsafe":    {Outcome: "unsafe", RungReached: "none", Predicate: observe.Blocked, FailingDetail: "policy-refused"},
		"machinery": {Outcome: "unsafe", RungReached: "none", Predicate: observe.Blocked, FailingDetail: "executor-timeout"},
		"fail":      {Outcome: "fail", RungReached: "none", Predicate: observe.Fail, FailingDetail: "suite-exit-mismatch"},
		"no-vantage": {Outcome: "skipped", RungReached: "none", Predicate: observe.Blocked,
			FailingDetail: "observation-unavailable"},
	}}
	named := map[string]bool{"pass": true, "unsafe": true, "machinery": true, "fail": true, "no-vantage": true}
	reg := observe.NewRegistry(observe.RegistryEnv{NamedSuites: named, Executor: exec})
	for _, tc := range []struct {
		name      string
		targets   []string
		authority string
		terminal  string
		class     string
		escalate  bool
		integrity string
	}{
		{name: "pass plus unsafe follows integrity", targets: []string{"pass", "unsafe"}, authority: "no", terminal: record.Accepted, integrity: "mixed"},
		{name: "pass plus machinery holds authority", targets: []string{"pass", "machinery"}, authority: "yes", terminal: record.Held, class: "observe-machinery-fault", escalate: true, integrity: "mixed"},
		{name: "pass plus machinery rejects non authority", targets: []string{"pass", "machinery"}, authority: "no", terminal: record.Rejected, class: "observe-machinery-fault", integrity: "mixed"},
		{name: "observed false dominates machinery", targets: []string{"fail", "machinery"}, authority: "yes", terminal: record.Rejected, class: "observed-false", integrity: "mixed"},
		{name: "pass plus no vantage follows integrity", targets: []string{"pass", "no-vantage"}, authority: "no", terminal: record.Accepted, integrity: "mixed"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			exec.calls = nil
			rows := make([]map[string]string, 0, len(tc.targets))
			for i, target := range tc.targets {
				rows = append(rows, map[string]string{
					"claim_ref": fmt.Sprintf("claim-%d", i), "check_id": "run-suite",
					"params": s8ClaimParams(t, map[string]string{"target": target, "expect_green": "true"}),
				})
			}
			raw := s8Claims(t, rows...)
			cand := record.Record{Headers: map[string]string{"authority_class": tc.authority, "EVIDENCE_TARGET": "E2"}}
			result, terminal := observe.Gate(cand, "seat-a", "IMPL", "implementation", observe.Env{
				PresentLayers: map[string]bool{"observe": true},
				Evaluate:      func(candidate observe.Candidate) observe.PredicateResult { return reg.EvaluateClaims(raw, candidate) },
			})
			if terminal != tc.terminal || result.FailureClass != tc.class || result.Escalate != tc.escalate || result.ObservedFields["record_integrity"] != tc.integrity {
				t.Fatalf("terminal = %s, result = %#v", terminal, result)
			}
			if len(exec.calls) != len(tc.targets) {
				t.Fatalf("executor calls = %v, want all %v", exec.calls, tc.targets)
			}
		})
	}
}

type s8ClaimExecutor struct {
	verdicts map[string]observe.CheckVerdict
	calls    []string
}

func (e *s8ClaimExecutor) Spawn(_ observe.CheckEntry, selection observe.Selection) observe.CheckVerdict {
	target := selection.Params["target"]
	e.calls = append(e.calls, target)
	verdict := e.verdicts[target]
	verdict.CheckID = selection.CheckID
	verdict.ClaimRef = selection.ClaimRef
	return verdict
}

func s8ClaimParams(t *testing.T, params map[string]string) string {
	t.Helper()
	raw, err := fieldspec.CanonicalMarshal(params)
	if err != nil {
		t.Fatalf("marshal claim params: %v", err)
	}
	return raw
}

func s8Claims(t *testing.T, rows ...map[string]string) string {
	t.Helper()
	raw, err := fieldspec.CanonicalMarshal(rows)
	if err != nil {
		t.Fatalf("marshal claims: %v", err)
	}
	return raw
}

type s8SubmitResult struct {
	rec     record.Record
	intents []store.Intent
}

func s8SubmitWithObservation(t *testing.T, st *store.Store, reg *fieldspec.Registry, meta seat.SeatMeta, renderEnv fieldspec.RenderEnv, rec record.Record, result observe.PredicateResult, intakeID string) s8SubmitResult {
	t.Helper()
	_, digest := reg.Render(renderEnv, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator}, rec.Headers["PHASE"], rec.Headers["CEREMONY_TIER"], fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal observed submit: %v", err)
	}
	handler := engine.SubmitHandlerWithObservation(st, reg, meta, renderEnv, observe.Env{
		PresentLayers: renderEnv.PresentLayers,
		Evaluate:      func(observe.Candidate) observe.PredicateResult { return result },
	})
	got, intents, err := handler(context.Background(), intake.Cmd{IntakeID: intakeID, Seat: meta.Name, Role: meta.Role, Payload: payload})
	if err != nil {
		t.Fatalf("observed submit: %v", err)
	}
	return s8SubmitResult{rec: got, intents: intents}
}

func s8CommitOperatorConfigChange(t *testing.T, st *store.Store, member string, body []byte) {
	t.Helper()
	pinned := s8PinnedStore(t, st.Root)
	meta := seat.SeatMeta{Name: "operator", Role: "operator", IsOperator: true}
	rec := s8ConfigChangeRecord(t, st.Root, member, body)
	env := fieldspec.RenderEnv{ConfigDigest: pinned.Digest, PresentLayers: config.PresentLayers(pinned)}
	_, digest := pinned.Registry.Render(env, fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: true}, "SITREP", "medium", fieldspec.ClosedGrantState)
	payload, err := json.Marshal(fieldspec.SubmitPayload{Record: rec, FormDigest: digest})
	if err != nil {
		t.Fatalf("marshal config change: %v", err)
	}
	got, intents, err := engine.SubmitHandlerWithRender(st, pinned.Registry, meta, env)(context.Background(), intake.Cmd{
		IntakeID: "activate-" + member, Seat: meta.Name, Role: meta.Role, IsOperator: true, Payload: payload,
	})
	if err != nil {
		t.Fatalf("submit %s config change: %v", member, err)
	}
	if got.Envelope.DeliveryState != record.Accepted {
		t.Fatalf("%s config change = %s, body=%s", member, got.Envelope.DeliveryState, got.Body)
	}
	if _, err := st.Commit(got, intents); err != nil {
		t.Fatalf("commit %s config change: %v", member, err)
	}
}

func s8PinnedStore(t *testing.T, root string) *config.Pinned {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("Load pinned store: %v", err)
	}
	return pinned
}
