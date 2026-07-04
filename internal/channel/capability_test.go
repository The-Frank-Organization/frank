package channel_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
)

func TestCapability(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tools := channel.ToolSet{
		Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`{"ok":true}`), nil
		},
		Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`[]`), nil
		},
		Read: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`{"record":null}`), nil
		},
	}

	seatA := filepath.Join(t.TempDir(), "seat-a.sock")
	seatB := filepath.Join(t.TempDir(), "seat-b.sock")

	serverA, err := channel.Serve(seatA, tools)
	if err != nil {
		t.Fatalf("serve seat A: %v", err)
	}
	defer serverA.Close()

	serverB, err := channel.Serve(seatB, tools)
	if err != nil {
		t.Fatalf("serve seat B: %v", err)
	}
	defer serverB.Close()

	clientA, err := channel.Dial(ctx, seatA)
	if err != nil {
		t.Fatalf("dial seat A: %v", err)
	}
	defer clientA.Close()

	clientB, err := channel.Dial(ctx, seatB)
	if err != nil {
		t.Fatalf("dial seat B: %v", err)
	}
	defer clientB.Close()

	gotTools, err := clientA.ListTools(ctx)
	if err != nil {
		t.Fatalf("list tools: %v", err)
	}
	wantTools := []string{"submit", "project", "read"}
	if !reflect.DeepEqual(gotTools, wantTools) {
		t.Fatalf("tools = %v, want %v", gotTools, wantTools)
	}

	result, err := clientA.Call(ctx, "submit", json.RawMessage(`{"subject":"hello"}`))
	if err != nil {
		t.Fatalf("call submit: %v", err)
	}
	if string(result) != `{"ok":true}` {
		t.Fatalf("submit result = %s", result)
	}

	if err := serverA.Push([]byte(`{"method":"notifications/nudge"}`)); err != nil {
		t.Fatalf("push: %v", err)
	}
	frame, err := clientA.NextPush(ctx)
	if err != nil {
		t.Fatalf("next push: %v", err)
	}
	if string(frame) != `{"method":"notifications/nudge"}` {
		t.Fatalf("push frame = %s", frame)
	}

	if err := clientB.CallExpectNoTool(ctx, "seat-a-only"); err == nil {
		t.Fatalf("seat B unexpectedly reached a non-shared tool")
	}
}
