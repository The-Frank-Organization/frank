package channel_test

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

	"github.com/The-Frank-Organization/frank/internal/channel"
	"github.com/The-Frank-Organization/frank/internal/seat"
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
	defer func() { _ = serverA.Close() }()

	serverB, err := channel.Serve(seatB, tools)
	if err != nil {
		t.Fatalf("serve seat B: %v", err)
	}
	defer func() { _ = serverB.Close() }()

	clientA, err := channel.Dial(ctx, seatA)
	if err != nil {
		t.Fatalf("dial seat A: %v", err)
	}
	defer func() { _ = clientA.Close() }()

	clientB, err := channel.Dial(ctx, seatB)
	if err != nil {
		t.Fatalf("dial seat B: %v", err)
	}
	defer func() { _ = clientB.Close() }()

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

	if err := clientB.CallExpectNoTool(ctx, "seat-a-only"); err == nil {
		t.Fatalf("seat B unexpectedly reached a non-shared tool")
	}
}

func TestAuthenticatedServerRejectsUnknownCredentialAndSanitizesToolErrors(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-auth-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.ServeAuthenticated(sock, mgr, func(meta seat.SeatMeta) channel.ToolSet {
		if meta.Name != "seat-a" || meta.Role != "implementer" {
			t.Fatalf("unexpected meta: %+v", meta)
		}
		return channel.ToolSet{
			Submit: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return nil, errors.New(filepath.Join(root, "records", "leak.json"))
			},
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
			Read:    func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`{}`), nil },
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	defer func() { _ = server.Close() }()

	if _, err := channel.DialAuthenticated(ctx, sock, "missing"); err == nil {
		t.Fatalf("unknown credential connected")
	}

	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated: %v", err)
	}
	defer func() { _ = client.Close() }()
	_, err = client.Call(ctx, "submit", json.RawMessage(`{}`))
	if err == nil {
		t.Fatalf("submit unexpectedly succeeded")
	}
	if strings.Contains(err.Error(), root) || strings.Contains(err.Error(), "/records/") {
		t.Fatalf("tool error leaked path: %q", err)
	}
}

func TestReadOnlyToolSetListsAndServesOnlyProjectAndRead(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-readonly-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.Serve(sock, channel.ToolSet{
		Project: func(context.Context, json.RawMessage) (json.RawMessage, error) { return json.RawMessage(`[]`), nil },
		Read: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`{"ok":true}`), nil
		},
	})
	if err != nil {
		t.Fatalf("Serve: %v", err)
	}
	defer func() { _ = server.Close() }()
	client, err := channel.Dial(ctx, sock)
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer func() { _ = client.Close() }()

	tools, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("ListTools: %v", err)
	}
	want := []string{"project", "read"}
	if !reflect.DeepEqual(tools, want) {
		t.Fatalf("tools = %v, want %v", tools, want)
	}
	if _, err := client.Call(ctx, "submit", json.RawMessage(`{}`)); err == nil {
		t.Fatalf("submit unexpectedly available on read-only surface")
	}
}
