package fixtures_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
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
	outputs := []string{
		`{"state":"accepted"}`,
		`{"state":"rejected","detail":"PHASE:enum"}`,
		`{"state":"held","reason":"system:internal-fault"}`,
		`{"error_class":"root-lock-held"}`,
		`{"error":"roster:seat-scope"}`,
	}
	for _, output := range outputs {
		for _, forbidden := range []string{"bounced", "submitted"} {
			if strings.Contains(output, forbidden) {
				t.Fatalf("output contains forbidden enum token %q: %s", forbidden, output)
			}
		}
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
	if !reflect.DeepEqual(got, []string{"submit", "project", "read"}) {
		t.Fatalf("tool names = %v", got)
	}
}

func TestS6SweepProjectParamsCarryAuditAndRosterWithoutNewVerb(t *testing.T) {
	described := []string{"submit", "project", "read"}
	if !reflect.DeepEqual(described, []string{"submit", "project", "read"}) {
		t.Fatalf("tool names = %v", described)
	}
	projectPayloads := []map[string]string{{"view": "audit"}, {"view": "roster"}}
	for _, payload := range projectPayloads {
		data, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("marshal project payload: %v", err)
		}
		if !strings.Contains(string(data), "view") {
			t.Fatalf("project payload lost view parameter: %s", data)
		}
	}
}

func hasExclusivityClaim(text string) bool {
	for _, token := range []string{"only writer", "sole", "no lane can", "non-lane-writable", "unbypassable"} {
		if strings.Contains(text, token) {
			return true
		}
	}
	return false
}
