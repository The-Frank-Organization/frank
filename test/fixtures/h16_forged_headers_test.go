package fixtures_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/derived"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/observe"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestH16ForgedSystemHeadersFourStateMatrix(t *testing.T) {
	type matrixState struct {
		name      string
		observe   bool
		active    bool
		wantClass func(string) string
	}
	states := []matrixState{
		{name: "S1_pre-active_observe-enabled", observe: true, wantClass: func(value string) string {
			if value != "" {
				return "lane-supplied-system-field"
			}
			return "non-boot-before-active"
		}},
		{name: "S2_pre-active_observe-absent", wantClass: func(string) string { return "non-boot-before-active" }},
		{name: "S3_active_observe-enabled", observe: true, active: true, wantClass: func(value string) string {
			if value != "" {
				return "lane-supplied-system-field"
			}
			return "system-owned"
		}},
		{name: "S4_active_observe-absent", active: true, wantClass: func(string) string { return "system-owned" }},
	}
	values := []struct {
		name  string
		value string
	}{
		{name: "non-empty", value: "forged"},
		{name: "present-empty", value: ""},
	}
	shapes := []string{"hooked", "unhooked"}
	caseCount := 0

	for _, state := range states {
		state := state
		t.Run(state.name, func(t *testing.T) {
			h := newS4ShimHarnessWithSources(t, s8ConfigSources(t, false))
			operatorCred := h.mint(t, "operator", "operator")
			h.start(t)
			operator := h.dial(t, operatorCred)
			seatName := strings.ReplaceAll(strings.ToLower(state.name), "_", "-") + ".implementer"
			mint := submitS6LiveMint(t, h.ctx, operator, seatName, "implementer", false)
			_ = operator.Close()
			if state.observe {
				h.stop(t)
				st, err := store.Open(h.root)
				if err != nil {
					t.Fatalf("open store for observe activation: %v", err)
				}
				pinned := s8PinnedStore(t, h.root)
				var engineDoc map[string]any
				if err := json.Unmarshal(pinned.Members["engine"], &engineDoc); err != nil {
					t.Fatalf("decode engine config: %v", err)
				}
				engineDoc["present_layers"].(map[string]any)["observe"] = true
				engineBytes, err := json.Marshal(engineDoc)
				if err != nil {
					t.Fatalf("marshal engine config: %v", err)
				}
				s8CommitOperatorConfigChange(t, st, "engine", engineBytes)
				h.start(t)
			}
			client, err := channel.DialAuthenticated(h.ctx, h.sock, mint.Credential)
			if err != nil {
				t.Fatalf("dial matrix seat: %v", err)
			}
			defer func() { _ = client.Close() }()
			if state.active {
				h16ActivateMatrixSeat(t, h.ctx, client)
			}

			for _, header := range []string{"hook_contract", "mint_predecessor", "admin_provenance"} {
				for _, value := range values {
					for _, shape := range shapes {
						header, value, shape := header, value, shape
						caseCount++
						t.Run(header+"/"+value.name+"/"+shape, func(t *testing.T) {
							rec := h16ForgedMatrixCandidate(header, value.value, shape, caseCount)
							describe, err := client.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
							if err != nil {
								t.Fatalf("DescribeTools: %v", err)
							}
							result, err := client.Call(h.ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: describe.FormDigest}))
							if err != nil {
								t.Fatalf("raw submit: %v", err)
							}
							var outcome struct {
								State   string `json:"state"`
								RelayID string `json:"relay_id"`
								Detail  string `json:"detail"`
							}
							if err := json.Unmarshal(result, &outcome); err != nil {
								t.Fatalf("decode outcome %s: %v", result, err)
							}
							wantClass := state.wantClass(value.value)
							if outcome.State != record.Rejected || outcome.RelayID == "" || !strings.Contains(outcome.Detail, header+":"+wantClass) {
								t.Fatalf("outcome=%+v, want rejected %s:%s", outcome, header, wantClass)
							}

							st, err := store.Open(h.root)
							if err != nil {
								t.Fatalf("open store: %v", err)
							}
							committed, err := st.Read(outcome.RelayID)
							if err != nil {
								t.Fatalf("read rejected record: %v", err)
							}
							if got, present := committed.Headers[header]; !present || got != value.value {
								t.Fatalf("forged header lost: present=%v value=%q want=%q headers=%v", present, got, value.value, committed.Headers)
							}
							if committed.Envelope.DeliveryState != record.Rejected || committed.Headers["failing_edge"] != "form-validation" {
								t.Fatalf("committed rejection=%+v", committed)
							}
							records := h16Records(t, st)
							if _, member := derived.Fold(records)[outcome.RelayID]; member {
								t.Fatal("rejected forged record entered derived membership")
							}
							chains, err := engine.BuildMintChains(records)
							if err != nil {
								t.Fatalf("build mint chains: %v", err)
							}
							targetSeat := fmt.Sprintf("forged-target-%02d.implementer", caseCount)
							if chain, member := chains[targetSeat]; member || chain.Tip.Envelope.RelayID != "" {
								t.Fatalf("rejected forged record entered mint chain: %+v", chain)
							}
							if committed.Envelope.DeliveryState == record.Accepted || committed.Headers["record_kind"] == "derived-work-attempt" || committed.Headers["record_kind"] == "derived-work-transition" {
								t.Fatalf("rejected case produced accepted/derived work: %+v", committed)
							}
						})
					}
				}
			}
		})
	}
	if caseCount != 48 {
		t.Fatalf("executed cases=%d, want 48", caseCount)
	}
}

func TestH16ForgedHeaderControlsAndRollback(t *testing.T) {
	reg := loadH16Registry(t)
	for _, field := range []string{"achieved_evidence", "executable_claim_results", "authority_class", "surface_intent"} {
		t.Run("S8-control/"+field, func(t *testing.T) {
			if got := observe.LaneSuppliedSystemField(reg, record.Record{Headers: map[string]string{field: "forged"}}); got != field {
				t.Fatalf("LaneSuppliedSystemField=%q, want %q", got, field)
			}
		})
	}
	producer, violation := engine.CompleteObserved(record.Record{Headers: map[string]string{"authority_class": "no"}}, map[string]string{
		"achieved_evidence": "E0", "target_gap_result": "not_applicable", "evidence_integrity": `{}`,
		"record_integrity": "self_reported", "executable_claim_results": `[]`, "egress_scan_result": "not_applicable",
		"degradation_notes": "", "attestation_source": "conductor", "deviated_observed": "no", "bucket_binding_observed": "no",
	})
	if violation != nil {
		t.Fatalf("producer control violation=%+v", violation)
	}
	if value, present := producer.Headers["degradation_notes"]; !present || value != "" {
		t.Fatalf("producer empty member present=%v value=%q", present, value)
	}

	h := newS4ShimHarness(t)
	operatorCred := h.mint(t, "operator", "operator")
	h.start(t)
	operator := h.dial(t, operatorCred)
	defer func() { _ = operator.Close() }()
	accepted := submitS6LiveMint(t, h.ctx, operator, "rollback-accepted.implementer", "implementer", false)
	st, err := store.Open(h.root)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	acceptedRecord, err := st.Read(accepted.RelayID)
	if err != nil {
		t.Fatalf("read accepted hooked record: %v", err)
	}
	if acceptedRecord.Headers["hook_contract"] != derived.HookContractV1 || acceptedRecord.Headers["admin_provenance"] != "" {
		t.Fatalf("ordinary accepted mint headers=%v, want hook_contract=1 and no ceremony provenance", acceptedRecord.Headers)
	}

	for _, value := range []string{"1", "2"} {
		t.Run("raw-forged-hook-contract-"+value, func(t *testing.T) {
			rec := h16ForgedMatrixCandidate("hook_contract", value, "hooked", 90+int(value[0]-'0'))
			describe, err := operator.DescribeTools(h.ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
			if err != nil {
				t.Fatalf("describe: %v", err)
			}
			result, err := operator.Call(h.ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: rec, FormDigest: describe.FormDigest}))
			if err != nil {
				t.Fatalf("submit: %v", err)
			}
			var outcome struct {
				State   string `json:"state"`
				RelayID string `json:"relay_id"`
				Detail  string `json:"detail"`
			}
			if err := json.Unmarshal(result, &outcome); err != nil || outcome.State != record.Rejected || !strings.Contains(outcome.Detail, "hook_contract:system-owned") {
				t.Fatalf("outcome=%s err=%v", result, err)
			}
			committed, err := st.Read(outcome.RelayID)
			if err != nil || committed.Headers["hook_contract"] != value {
				t.Fatalf("committed forged value=%q err=%v record=%+v", value, err, committed)
			}
		})
	}

	future := acceptedRecord
	future.Envelope.RelayID = "accepted-future-contract"
	future.Headers = cloneStringMap(acceptedRecord.Headers)
	future.Headers["hook_contract"] = "2"
	status, present := derived.Fold([]record.Record{future})[future.Envelope.RelayID]
	if !present || status.Status != "unknown" {
		t.Fatalf("future accepted hook contract fold present=%v status=%+v", present, status)
	}
}

func h16ActivateMatrixSeat(t *testing.T, ctx context.Context, client *channel.Client) {
	t.Helper()
	describe, err := client.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("describe boot form: %v", err)
	}
	boot := record.Record{Headers: map[string]string{
		"PHASE": "SITREP", "CEREMONY_TIER": "medium", "SUBJECT": "matrix boot", "charter_loaded": "yes", "dispatch_status": "read",
	}}
	result, err := client.Call(ctx, "submit", mustJSONBytes(t, fieldspec.SubmitPayload{Record: boot, FormDigest: describe.FormDigest}))
	if err != nil {
		t.Fatalf("boot submit: %v", err)
	}
	var outcome struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil || outcome.State != record.Accepted {
		t.Fatalf("boot outcome=%s err=%v", result, err)
	}
}

func h16ForgedMatrixCandidate(header, value, shape string, sequence int) record.Record {
	rec := h16PresenceCandidate()
	rec.Headers[header] = value
	rec.Headers["SUBJECT"] = fmt.Sprintf("forged %s %s %d", header, shape, sequence)
	if shape == "hooked" {
		rec.Headers["record_kind"] = "seat_mint"
		rec.Body = fmt.Sprintf(`{"seat":"forged-target-%02d.implementer","role":"implementer","is_operator":false}`, sequence)
	}
	return rec
}
