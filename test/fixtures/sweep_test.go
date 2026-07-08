package fixtures_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

func TestSweepReadmeClaimHonesty(t *testing.T) {
	data, err := os.ReadFile("../../README.md")
	if err != nil {
		t.Fatalf("README missing: %v", err)
	}
	readme := string(data)
	for _, want := range []string{
		"S1 = provenance + transport, not verified work",
		"self_reported",
		"tool-mediated confusion-resistance",
		"D5 residual",
		"pair-Planner grant rendering lands in S3",
		"the S3 registry rides `store.Init`",
		"registry/config evolution on an existing store is live through the operator-authorized §7 record, effective at restart",
	} {
		if !strings.Contains(readme, want) {
			t.Fatalf("README missing honesty phrase %q", want)
		}
	}
	for _, forbidden := range []string{"bounced", "submitted"} {
		if strings.Contains(readme, forbidden) {
			t.Fatalf("README contains forbidden value token %q", forbidden)
		}
	}
	for _, paragraph := range strings.Split(readme, "\n\n") {
		if hasExclusivityClaim(paragraph) && !strings.Contains(paragraph, "D5 residual") && !strings.Contains(paragraph, "governance-surface") {
			t.Fatalf("claim lacks D5/governance qualifier: %q", paragraph)
		}
	}
}

func TestS3SweepGateCategoryAuthorityUsesRegistry(t *testing.T) {
	data, err := os.ReadFile("../../internal/lineage/lineage.go")
	if err != nil {
		t.Fatalf("read lineage.go: %v", err)
	}
	helperName := "isA" + "GateCategory"
	if strings.Contains(string(data), helperName) {
		t.Fatalf("lineage still carries the S1 hard-coded gate-category helper")
	}
}

func TestS3SweepMainThreadsActiveLineageTurnContext(t *testing.T) {
	data, err := os.ReadFile("../../cmd/frank/main.go")
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	main := string(data)
	if strings.Contains(main, "ActiveLineageCandidates(tab, lineage.TurnContext{})") {
		t.Fatalf("main.go still passes an empty lineage turn context into ActiveLineageCandidates")
	}
	if !strings.Contains(main, "turnContextForSeat") {
		t.Fatalf("main.go does not thread the conductor-derived turn context helper")
	}
}

func TestS6SweepEnumFloorAndThreeVerbSurface(t *testing.T) {
	reg, err := fieldspec.Load(filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	if err != nil {
		t.Fatalf("load registry: %v", err)
	}
	registryEnums, err := json.Marshal(reg.NamedEnums)
	if err != nil {
		t.Fatalf("marshal registry enums: %v", err)
	}
	sources := map[string][]byte{
		"record-delivery-state-consts": []byte(strings.Join([]string{record.Accepted, record.Rejected, record.Held}, "\n")),
		"registry-named-enums":         registryEnums,
	}
	for name, data := range s6RealSerializedOutcomeCorpus(t) {
		sources[name] = data
	}
	if leaks := forbiddenEnumLeaks(sources); len(leaks) > 0 {
		t.Fatalf("forbidden enum tokens leaked from real surfaces: %v", leaks)
	}
	if leaks := forbiddenEnumLeaks(map[string][]byte{"plant": []byte(`{"state":"bounced"}`)}); len(leaks) == 0 {
		t.Fatalf("forbidden enum scanner did not bite on planted leak")
	}

	tools := channel.ToolSet{
		Submit:  func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`{}`), nil },
		Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
		Read:    func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`{}`), nil },
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-sweep-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.Serve(sock, tools)
	if err != nil {
		t.Fatalf("Serve: %v", err)
	}
	defer func() { _ = server.Close() }()
	client, err := channel.Dial(context.Background(), sock)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer func() { _ = client.Close() }()
	got, err := client.ListTools(context.Background())
	if err != nil {
		t.Fatalf("ListTools: %v", err)
	}
	if strings.Join(got, ",") != "submit,project,read" {
		t.Fatalf("tool names = %v", got)
	}
}

func TestS6SweepProjectParamsCarryAuditAndRosterWithoutNewVerb(t *testing.T) {
	root := t.TempDir()
	initFixtureStore(t, root)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	bin := buildFrank(t, ctx)
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	operatorCred, err := mgr.Mint("operator", "operator", true)
	if err != nil {
		t.Fatalf("Mint operator: %v", err)
	}
	sock := filepath.Join(os.TempDir(), "frank-s6-sweep-project-"+filepath.Base(root)+".sock")
	t.Cleanup(func() { _ = os.Remove(sock) })
	cmd, stderr := startFrank(t, ctx, bin, root, sock)
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)
	operator, err := channel.DialAuthenticated(ctx, sock, operatorCred.Value)
	if err != nil {
		t.Fatalf("operator dial stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = operator.Close() }()

	describe, err := operator.DescribeTools(ctx, channel.DescribeRequest{Phase: "SITREP", Tier: "medium"})
	if err != nil {
		t.Fatalf("DescribeTools: %v", err)
	}
	if strings.Join(describe.Tools, ",") != "submit,project,read" {
		t.Fatalf("tool names = %v", describe.Tools)
	}
	for _, view := range []string{"audit", "roster"} {
		result, err := operator.Call(ctx, "project", mustJSONBytes(t, map[string]string{"view": view}))
		if err != nil {
			t.Fatalf("project %s: %v", view, err)
		}
		var decoded any
		if err := json.Unmarshal(result, &decoded); err != nil {
			t.Fatalf("project %s returned non-json %s: %v", view, result, err)
		}
	}
}

func s6RealSerializedOutcomeCorpus(t *testing.T) map[string][]byte {
	t.Helper()
	st, err := store.Open(t.TempDir())
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	for _, rec := range []record.Record{
		{Envelope: record.Envelope{RelayID: "accepted-outcome", From: "seat-a", Role: "implementer", DeliveryState: record.Accepted, SchemaVersion: 1}, Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "accepted"}},
		{Envelope: record.Envelope{RelayID: "rejected-outcome", From: "system", Role: "system", DeliveryState: record.Rejected, SchemaVersion: 1}, Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "rejected"}, Body: "PHASE:enum"},
		{Envelope: record.Envelope{RelayID: "held-outcome", From: "system", Role: "system", DeliveryState: record.Held, SchemaVersion: 1}, Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "held"}, Body: "system:internal-fault"},
	} {
		if _, err := st.Commit(rec, nil); err != nil {
			t.Fatalf("Commit %s: %v", rec.Envelope.RelayID, err)
		}
	}
	out := map[string][]byte{}
	for _, relayID := range []string{"accepted-outcome", "rejected-outcome", "held-outcome"} {
		data, err := os.ReadFile(filepath.Join(st.Root, "records", relayID+".json"))
		if err != nil {
			t.Fatalf("read serialized outcome %s: %v", relayID, err)
		}
		out["serialized-"+relayID] = data
	}
	return out
}

func forbiddenEnumLeaks(sources map[string][]byte) []string {
	var leaks []string
	for name, data := range sources {
		text := string(data)
		for _, forbidden := range []string{"bounced", "submitted"} {
			if strings.Contains(text, forbidden) {
				leaks = append(leaks, name+":"+forbidden)
			}
		}
	}
	return leaks
}

func hasExclusivityClaim(text string) bool {
	for _, token := range []string{"only writer", "sole", "no lane can", "non-lane-writable", "unbypassable"} {
		if strings.Contains(text, token) {
			return true
		}
	}
	return false
}
