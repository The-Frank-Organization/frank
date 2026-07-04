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

func TestP1LoopOutcomeDoesNotLeakStorePaths(t *testing.T) {
	root := t.TempDir()
	st, err := store.Open(root)
	if err != nil {
		t.Fatalf("Open store: %v", err)
	}
	loop := engine.New(st, func(context.Context, intake.Cmd) (record.Record, []store.Intent, error) {
		return record.Record{}, nil, errors.New(filepath.Join(root, "records", "leak.json"))
	})
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

func TestP1ToolDescriptionsAndPushFramesDoNotLeakPaths(t *testing.T) {
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
	if err := server.Push([]byte(`{"kind":"recovery-nudge","mailbox":"seat-a"}`)); err != nil {
		t.Fatalf("Push: %v", err)
	}
	push, err := client.NextPush(ctx)
	if err != nil {
		t.Fatalf("NextPush: %v", err)
	}
	outputs := []string{strings.Join(toolList, ","), string(push)}
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
