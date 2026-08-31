package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/config"
	"github.com/The-Frank-Organization/frank/internal/fieldspec"
	"github.com/The-Frank-Organization/frank/internal/record"
	"github.com/The-Frank-Organization/frank/internal/store"
)

func TestRunningConfigUnchangedUntilRestart(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	h.start(t)

	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	before, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools before: %v", err)
	}

	body := mutatedRegistryBody(t, h.root)
	rec := record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "registry config_change",
			"record_kind":     "config_change",
			"member":          "fieldspec",
			"new_digest":      fixtureDigestWithMember(t, h.root, "fieldspec", body),
		},
		Body: string(body),
	}
	h.submit(t, operator, rec)

	materialized, err := os.ReadFile(filepath.Join(h.root, "config", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("read materialized registry: %v", err)
	}
	if string(materialized) != string(body) {
		t.Fatalf("materialized registry differs from submitted body")
	}
	after, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools after: %v", err)
	}
	if after.FormDigest != before.FormDigest {
		t.Fatalf("running registry hot-reloaded: before digest %s after %s", before.FormDigest, after.FormDigest)
	}
}

func TestConfigChangeReadRedactedForNonOperator(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	seatCred := h.mint(t, "seat-a", "implementer")
	h.start(t)

	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	seatClient := h.dial(t, seatCred)
	defer func() { _ = seatClient.Close() }()
	body := mutatedRegistryBody(t, h.root)
	rec := configChangeRelay(t, h.root, body)
	relayID := h.submit(t, operator, rec)

	read := readRelayRaw(t, h.ctx, seatClient, relayID)
	if bytes.Contains(read, body) || bytes.Contains(read, []byte("body")) {
		t.Fatalf("non-operator read leaked member bytes: %s", read)
	}
	var got map[string]any
	if err := json.Unmarshal(read, &got); err != nil {
		t.Fatalf("decode redacted read %s: %v", read, err)
	}
	wantKeys := []string{"envelope", "member", "new_digest", "record_kind", "redacted", "relay_id"}
	var keys []string
	for key := range got {
		keys = append(keys, key)
	}
	sortStrings(keys)
	if !reflect.DeepEqual(keys, wantKeys) {
		t.Fatalf("redacted keys = %v, want %v; read=%s", keys, wantKeys, read)
	}
	if got["relay_id"] != relayID ||
		got["record_kind"] != "config_change" ||
		got["member"] != "fieldspec" ||
		got["new_digest"] != rec.Headers["new_digest"] ||
		got["redacted"] != "config-member-bytes" {
		t.Fatalf("redacted read fields = %#v", got)
	}
	envelope, ok := got["envelope"].(map[string]any)
	if !ok || envelope["from"] != "operator" || envelope["role"] != "operator" || envelope["schema_version"] != float64(1) {
		t.Fatalf("redacted envelope = %#v", got["envelope"])
	}
}

func TestConfigChangeReadFullForOperator(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	h.start(t)

	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	body := mutatedRegistryBody(t, h.root)
	relayID := h.submit(t, operator, configChangeRelay(t, h.root, body))

	read := readRelayRaw(t, h.ctx, operator, relayID)
	var got struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(read, &got); err != nil {
		t.Fatalf("decode operator read %s: %v", read, err)
	}
	if got.Record.Body != string(body) || got.Record.Headers["record_kind"] != "config_change" {
		t.Fatalf("operator record = %+v", got.Record)
	}
}

func TestConfigChangeProjectionsCarryNoMemberBytes(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	seatCred := h.mint(t, "seat-a", "implementer")
	h.start(t)

	seatClient := h.dial(t, seatCred)
	defer func() { _ = seatClient.Close() }()
	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	body := mutatedRegistryBody(t, h.root)
	rec := configChangeRelay(t, h.root, body)
	rec.Envelope.To = "seat-a"
	h.submit(t, operator, rec)

	nudge, err := seatClient.NextPush(h.ctx)
	if err != nil {
		t.Fatalf("NextPush config_change nudge: %v", err)
	}
	if bytes.Contains(nudge, body) {
		t.Fatalf("nudge leaked member bytes: %s", nudge)
	}
	assertTreeLacksBytes(t, filepath.Join(h.root, "projections"), body)
}

func TestPayloadDigestIsClaimNotAuthority(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	h.start(t)

	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	body := mutatedRegistryBody(t, h.root)
	rec := configChangeRelay(t, h.root, body)
	rec.Headers["new_digest"] = "payload-claim-is-not-authority"

	describe, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools: %v", err)
	}
	payload := mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: describe.FormDigest})
	result, err := operator.Call(h.ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit wrong digest: %v", err)
	}
	var outcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode outcome %s: %v", result, err)
	}
	if outcome.State != record.Rejected || outcome.RelayID == "" {
		t.Fatalf("outcome = %+v, want rejected relay", outcome)
	}
	read := readRelayRaw(t, h.ctx, operator, outcome.RelayID)
	if !bytes.Contains(read, []byte("new_digest")) || bytes.Contains(read, []byte(h.root)) {
		t.Fatalf("wrong digest read = %s, want typed path-free new_digest rejection", read)
	}
	materialized := mustReadFile(t, filepath.Join(h.root, "config", "fieldspec", "registry.json"))
	if bytes.Equal(materialized, body) {
		t.Fatalf("wrong digest materialized submitted body")
	}
}

func TestRestartWithNewRegistryBouncesStaleForm(t *testing.T) {
	h := newS4ShimHarness(t)
	operatorCred, err := h.mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	h.start(t)

	operator := h.dial(t, operatorCred)
	before, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools before: %v", err)
	}
	body := mutatedRegistryBody(t, h.root)
	h.submit(t, operator, configChangeRelay(t, h.root, body))
	_ = operator.Close()
	h.stop(t)
	h.start(t)

	operator = h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	stale := record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "stale form",
		},
		Body: "stale",
	}
	payload := mustJSONBytes(t, fieldspec.SubmitPayload{Record: stale, FormDigest: before.FormDigest})
	result, err := operator.Call(h.ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stale form: %v", err)
	}
	var staleOutcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &staleOutcome); err != nil {
		t.Fatalf("decode stale outcome %s: %v", result, err)
	}
	if staleOutcome.State != record.Rejected || staleOutcome.RelayID == "" {
		t.Fatalf("stale outcome = %+v, want rejected relay", staleOutcome)
	}
	read := readRelayRaw(t, h.ctx, operator, staleOutcome.RelayID)
	if !bytes.Contains(read, []byte("form_digest")) || !bytes.Contains(read, []byte("re-render")) {
		t.Fatalf("stale form read = %s, want form_digest re-render bounce", read)
	}

	after, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools after: %v", err)
	}
	if after.FormDigest == before.FormDigest {
		t.Fatalf("restart form digest did not change")
	}
	payload = mustJSONBytes(t, fieldspec.SubmitPayload{Record: stale, FormDigest: after.FormDigest})
	result, err = operator.Call(h.ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit fresh form: %v", err)
	}
	var freshOutcome struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(result, &freshOutcome); err != nil {
		t.Fatalf("decode fresh outcome %s: %v", result, err)
	}
	if freshOutcome.State != record.Accepted {
		t.Fatalf("fresh outcome = %+v, want accepted", freshOutcome)
	}
}

func mutatedRegistryBody(t *testing.T, root string) []byte {
	t.Helper()
	data := mustReadFile(t, filepath.Join(root, "config", "fieldspec", "registry.json"))
	var doc map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("decode registry: %v", err)
	}
	provenance, ok := doc["provenance"].(map[string]any)
	if !ok {
		provenance = map[string]any{}
		doc["provenance"] = provenance
	}
	provenance["s4_test_marker"] = "running-config-unchanged"
	body, err := json.Marshal(doc)
	if err != nil {
		t.Fatalf("marshal mutated registry: %v", err)
	}
	return body
}

func configChangeRelay(t *testing.T, root string, body []byte) record.Record {
	t.Helper()
	return record.Record{
		Headers: map[string]string{
			"PHASE":           "SITREP",
			"AUTHORITY":       "report-only",
			"CEREMONY_TIER":   "medium",
			"EVIDENCE_TARGET": "E1",
			"SUBJECT":         "registry config_change",
			"record_kind":     "config_change",
			"member":          "fieldspec",
			"new_digest":      fixtureDigestWithMember(t, root, "fieldspec", body),
		},
		Body: string(body),
	}
}

func readRelayRaw(t *testing.T, ctx context.Context, client *channel.Client, relayID string) []byte {
	t.Helper()
	args := mustJSONBytes(t, map[string]string{"relay_id": relayID})
	result, err := client.Call(ctx, "read", args)
	if err != nil {
		t.Fatalf("read %s: %v", relayID, err)
	}
	return result
}

func assertTreeLacksBytes(t *testing.T, root string, forbidden []byte) {
	t.Helper()
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if bytes.Contains(data, forbidden) {
			return fmt.Errorf("%s contains forbidden member bytes", path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("projection scan: %v", err)
	}
}

func sortStrings(values []string) {
	sort.Strings(values)
}

func fixtureDigestWithMember(t *testing.T, root, member string, body []byte) string {
	t.Helper()
	pinned, err := config.Load(store.StoreRootConfigPaths(root))
	if err != nil {
		t.Fatalf("load store config: %v", err)
	}
	members := make(map[string][]byte, len(pinned.Members))
	for name, data := range pinned.Members {
		members[name] = append([]byte(nil), data...)
	}
	members[member] = append([]byte(nil), body...)
	return config.Digest(members)
}
