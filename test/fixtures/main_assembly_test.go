package fixtures_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/record"
	"github.com/jackli/frank/internal/seat"
)

func TestFrankBinaryAssemblesAuthenticatedSubmitProjectRead(t *testing.T) {
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	cred, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-main-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	bin := filepath.Join(t.TempDir(), "frank")
	build := exec.CommandContext(ctx, "go", "build", "-o", bin, filepath.Join("..", "..", "cmd", "frank"))
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build frank: %v\n%s", err, out)
	}
	cmd := exec.CommandContext(ctx, bin, "-root", root, "-socket", sock, "-registry", filepath.Join("..", "..", "internal", "fieldspec", "registry.json"))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start frank: %v", err)
	}
	t.Cleanup(func() {
		cancel()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
		_ = cmd.Wait()
	})
	waitForSocket(t, sock)

	client, err := channel.DialAuthenticated(ctx, sock, cred.Value)
	if err != nil {
		t.Fatalf("DialAuthenticated stderr=%s: %v", stderr.String(), err)
	}
	defer func() { _ = client.Close() }()

	payload, _ := json.Marshal(record.Record{
		Envelope: record.Envelope{To: "seat-a", DispatchID: "dispatch-main"},
		Headers:  map[string]string{"PHASE": "SITREP", "AUTHORITY": "report-only", "SUBJECT": "binary path"},
		Body:     "hello",
	})
	result, err := client.Call(ctx, "submit", payload)
	if err != nil {
		t.Fatalf("submit stderr=%s: %v", stderr.String(), err)
	}
	var outcome struct {
		State   string `json:"state"`
		RelayID string `json:"relay_id"`
	}
	if err := json.Unmarshal(result, &outcome); err != nil {
		t.Fatalf("decode outcome %s: %v", result, err)
	}
	if outcome.State != record.Accepted || outcome.RelayID == "" {
		t.Fatalf("outcome = %+v", outcome)
	}

	project, err := client.Call(ctx, "project", nil)
	if err != nil {
		t.Fatalf("project: %v", err)
	}
	var relayIDs []string
	if err := json.Unmarshal(project, &relayIDs); err != nil {
		t.Fatalf("decode project %s: %v", project, err)
	}
	if len(relayIDs) != 1 || relayIDs[0] != outcome.RelayID {
		t.Fatalf("project = %v, want [%s]", relayIDs, outcome.RelayID)
	}

	readArgs, _ := json.Marshal(map[string]string{"relay_id": outcome.RelayID})
	readResult, err := client.Call(ctx, "read", readArgs)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	var read struct {
		Record record.Record `json:"record"`
	}
	if err := json.Unmarshal(readResult, &read); err != nil {
		t.Fatalf("decode read %s: %v", readResult, err)
	}
	if read.Record.Envelope.From != "seat-a" || read.Record.Envelope.Role != "implementer" {
		t.Fatalf("read record not stamped: %+v", read.Record.Envelope)
	}
}

func waitForSocket(t *testing.T, sock string) {
	t.Helper()
	deadline := time.Now().Add(4 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(sock); err == nil {
			return
		}
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatalf("socket %s was not created", sock)
}
