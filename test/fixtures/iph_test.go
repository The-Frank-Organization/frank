package fixtures_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/bounce"
	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/store"
)

func TestP1NoPathFamiliesInSeatDeliverableStrings(t *testing.T) {
	outputs := []string{
		bounce.Format(fieldspec.Violation{Field: "PHASE", Class: "enum", Reason: "/store/records/leak"}),
		`{"state":"accepted"}`,
		`["submit","project","read"]`,
	}
	for _, output := range outputs {
		for _, family := range []string{"/records", "/staging", "/outbox", "/binding", "operator-socket"} {
			if strings.Contains(output, family) {
				t.Fatalf("seat-deliverable output leaked %s in %q", family, output)
			}
		}
	}
}

func TestP1S3ExtendedSurfacesDoNotLeakPathFamilies(t *testing.T) {
	outputs := []string{
		bounce.Format(fieldspec.Violation{Field: "form_digest", Class: "re-render", Reason: filepath.Join(t.TempDir(), "records", "leak.json")}),
		migrate.ErrFutureVersion.Error(),
		migrate.ErrUnversioned.Error(),
		fmt.Errorf("%w: 1->2", migrate.ErrMigrationGap).Error(),
	}

	registryPath := filepath.Join(t.TempDir(), "registry.json")
	if err := os.WriteFile(registryPath, []byte(`{
		"version": "bad",
		"provenance": {"owner": "test"},
		"named_enums": {"bool": ["no", "yes"]},
		"fields": [{
			"id": "TRIGGER",
			"layer": "header",
			"owner": "agent_enum_pick",
			"type": "bool",
			"enum_set": "bool",
			"required_when": {"field": "MISSING", "op": "present"}
		}]
	}`), 0o644); err != nil {
		t.Fatalf("write registry fixture: %v", err)
	}
	if _, err := fieldspec.Load(registryPath); err != nil {
		outputs = append(outputs, err.Error())
	} else {
		t.Fatalf("bad registry unexpectedly loaded")
	}

	response, err := json.Marshal(channel.DescriptionResponse{
		Tools: []string{"submit", "project", "read"},
		SubmitSchema: &fieldspec.Form{Fields: map[string]fieldspec.Field{
			"PARENT_DISPATCH_ID": {Type: "id_ref", Options: []string{"parent-1"}, Default: "parent-1", DigestExempt: true},
		}},
		FormDigest: "digest-1",
	})
	if err != nil {
		t.Fatalf("marshal description response: %v", err)
	}
	outputs = append(outputs, string(response))
	assertNoPathFamilies(t, outputs...)
}

func TestP1LoopOutcomeDoesNotLeakStorePaths(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{}, nil, errors.New(filepath.Join(root, "records", "leak.json"))
	}, engine.TestReady())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go loop.Run(ctx)

	reply := make(chan engine.Outcome, 1)
	payload, _ := json.Marshal(record.Record{Headers: map[string]string{"PHASE": "SITREP", "SUBJECT": "leak"}})
	loop.In <- engine.Job{Cmd: intake.Cmd{IntakeID: "i-leak", Seat: "seat-a", Role: "implementer", Payload: payload}, ReplyCh: reply}
	out := <-reply
	if strings.Contains(out.Reason, root) || strings.Contains(out.Reason, "/records/") {
		t.Fatalf("loop outcome leaked path: %+v", out)
	}
}

func assertNoPathFamilies(t *testing.T, outputs ...string) {
	t.Helper()
	for _, output := range outputs {
		for _, family := range []string{"/records", "/staging", "/outbox", "/binding", "operator-socket"} {
			if strings.Contains(output, family) {
				t.Fatalf("P1 output leaked %s in %q", family, output)
			}
		}
	}
}

func TestP1ToolDescriptionsDoNotLeakPaths(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	tools := channel.ToolSet{
		Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`{"state":"accepted"}`), nil
		},
		Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
		Read: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`{"record":null}`), nil
		},
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-p1-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.Serve(sock, tools)
	if err != nil {
		t.Fatalf("Serve: %v", err)
	}
	defer func() { _ = server.Close() }()
	client, err := channel.Dial(ctx, sock)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer func() { _ = client.Close() }()

	toolList, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("ListTools: %v", err)
	}
	if !reflect.DeepEqual(toolList, []string{"submit", "project", "read"}) {
		t.Fatalf("tool list = %v", toolList)
	}
	descriptions, err := client.ToolDescriptions(ctx)
	if err != nil {
		t.Fatalf("ToolDescriptions: %v", err)
	}
	if len(descriptions) != 3 {
		t.Fatalf("descriptions = %#v", descriptions)
	}
	outputs := []string{strings.Join(toolList, ",")}
	for _, desc := range descriptions {
		outputs = append(outputs, desc)
	}
	for _, output := range outputs {
		for _, family := range []string{"/records", "/staging", "/outbox", "/binding", "operator-socket"} {
			if strings.Contains(output, family) {
				t.Fatalf("P1 output leaked %s in %q", family, output)
			}
		}
	}
}
