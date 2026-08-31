package channel_test

import (
	"bufio"
	"context"
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/channel"
)

func TestInboundOversizeFrameTypedRefusal(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const limit = 4096
	sock := filepath.Join(os.TempDir(), "frank-frame-inbound-"+time.Now().Format("150405.000000000")+".sock")
	t.Cleanup(func() { _ = os.Remove(sock) })

	server, err := channel.Serve(sock, channel.ToolSet{
		Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.RawMessage(`[]`), nil
		},
	}, channel.WithFrameLimit(limit))
	if err != nil {
		t.Fatalf("Serve: %v", err)
	}
	defer func() { _ = server.Close() }()

	conn, err := (&net.Dialer{}).DialContext(ctx, "unix", sock)
	if err != nil {
		t.Fatalf("DialContext: %v", err)
	}
	defer func() { _ = conn.Close() }()

	reader := bufio.NewReader(conn)
	if _, err := conn.Write([]byte(strings.Repeat("x", limit+1) + "\n")); err != nil {
		t.Fatalf("write oversized frame: %v", err)
	}
	refusal, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("read refusal: %v", err)
	}
	if !strings.Contains(refusal, "frame-too-large") || !strings.Contains(refusal, "4096") {
		t.Fatalf("refusal = %q, want frame-too-large with bound", refusal)
	}

	if _, err := conn.Write([]byte(`{"id":1,"method":"tools/list"}` + "\n")); err != nil {
		t.Fatalf("write tools/list after refusal: %v", err)
	}
	response, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("read tools/list after refusal: %v", err)
	}
	if !strings.Contains(response, `"project"`) {
		t.Fatalf("tools/list response after refusal = %q", response)
	}
}

func TestOutboundOversizeProjectRefusal(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	const limit = 4096
	sock := filepath.Join(os.TempDir(), "frank-frame-outbound-"+time.Now().Format("150405.000000000")+".sock")
	t.Cleanup(func() { _ = os.Remove(sock) })

	server, err := channel.Serve(sock, channel.ToolSet{
		Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
			return json.Marshal(map[string]string{"body": strings.Repeat("x", limit)})
		},
	}, channel.WithFrameLimit(limit))
	if err != nil {
		t.Fatalf("Serve: %v", err)
	}
	defer func() { _ = server.Close() }()

	client, err := channel.Dial(ctx, sock, channel.WithFrameLimit(limit))
	if err != nil {
		t.Fatalf("Dial: %v", err)
	}
	defer func() { _ = client.Close() }()

	_, err = client.Call(ctx, "project", nil)
	if err == nil {
		t.Fatalf("project unexpectedly succeeded")
	}
	if !strings.Contains(err.Error(), "frame-too-large") || !strings.Contains(err.Error(), "narrow the request") || !strings.Contains(err.Error(), "4096") {
		t.Fatalf("project error = %q, want typed frame refusal with hint and bound", err)
	}
	if tools, err := client.ListTools(ctx); err != nil {
		t.Fatalf("ListTools after outbound refusal: %v", err)
	} else if len(tools) != 1 || tools[0] != "project" {
		t.Fatalf("tools after outbound refusal = %v", tools)
	}
}
