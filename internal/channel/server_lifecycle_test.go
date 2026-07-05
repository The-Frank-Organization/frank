package channel_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/seat"
)

func TestSecondConnectSameCredentialRejected(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sock, _, credA, _ := serveLifecycleServer(t)
	first, err := channel.DialAuthenticated(ctx, sock, credA.Value)
	if err != nil {
		t.Fatalf("first DialAuthenticated: %v", err)
	}
	defer func() { _ = first.Close() }()

	_, err = channel.DialAuthenticated(ctx, sock, credA.Value)
	if err == nil {
		t.Fatalf("second DialAuthenticated unexpectedly succeeded")
	}
	if !strings.Contains(err.Error(), "auth:channel-active") {
		t.Fatalf("second connect err = %q, want auth:channel-active", err)
	}
	hash := sha256.Sum256([]byte(credA.Value))
	if strings.Contains(err.Error(), credA.Value) || strings.Contains(err.Error(), hex.EncodeToString(hash[:])) {
		t.Fatalf("second connect leaked credential material: %q", err)
	}

	if tools, err := first.ListTools(ctx); err != nil {
		t.Fatalf("first connection after reject: %v", err)
	} else if len(tools) != 1 || tools[0] != "project" {
		t.Fatalf("first tools = %v", tools)
	}
}

func TestDistinctCredentialsBothConnect(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sock, _, credA, credB := serveLifecycleServer(t)
	first, err := channel.DialAuthenticated(ctx, sock, credA.Value)
	if err != nil {
		t.Fatalf("first DialAuthenticated: %v", err)
	}
	defer func() { _ = first.Close() }()

	second, err := channel.DialAuthenticated(ctx, sock, credB.Value)
	if err != nil {
		t.Fatalf("second distinct DialAuthenticated: %v", err)
	}
	defer func() { _ = second.Close() }()
}

func TestProvenDeadRecovery(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sock, _, credA, _ := serveLifecycleServer(t)
	first, err := channel.DialAuthenticated(ctx, sock, credA.Value)
	if err != nil {
		t.Fatalf("first DialAuthenticated: %v", err)
	}
	if err := first.Close(); err != nil {
		t.Fatalf("first Close: %v", err)
	}

	eventuallyDialAuthenticated(t, ctx, sock, credA.Value)
}

func TestKillHostEscapeHatch(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sock, _, credA, _ := serveLifecycleServer(t)
	cmd := exec.Command(os.Args[0], "-test.run=TestCredentialHolderHelper", "--")
	cmd.Env = append(os.Environ(),
		"FRANK_TEST_CREDENTIAL_HOLDER=1",
		"FRANK_TEST_SOCKET="+sock,
		"FRANK_TEST_CREDENTIAL="+credA.Value,
	)
	if err := cmd.Start(); err != nil {
		t.Fatalf("start credential holder: %v", err)
	}
	t.Cleanup(func() {
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
			_, _ = cmd.Process.Wait()
		}
	})

	waitForActiveCredential(t, ctx, sock, credA.Value)
	if err := cmd.Process.Kill(); err != nil {
		t.Fatalf("kill credential holder: %v", err)
	}
	_, _ = cmd.Process.Wait()

	eventuallyDialAuthenticated(t, ctx, sock, credA.Value)
}

func TestCredentialHolderHelper(t *testing.T) {
	if os.Getenv("FRANK_TEST_CREDENTIAL_HOLDER") != "1" {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := channel.DialAuthenticated(ctx, os.Getenv("FRANK_TEST_SOCKET"), os.Getenv("FRANK_TEST_CREDENTIAL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "DialAuthenticated: %v\n", err)
		os.Exit(2)
	}
	defer func() { _ = client.Close() }()
	<-ctx.Done()
}

func serveLifecycleServer(t *testing.T) (string, *channel.Server, seat.Cred, seat.Cred) {
	t.Helper()
	root := t.TempDir()
	mgr, err := seat.Open(root)
	if err != nil {
		t.Fatalf("Open seats: %v", err)
	}
	credA, err := mgr.Mint("seat-a", "implementer", false)
	if err != nil {
		t.Fatalf("Mint seat-a: %v", err)
	}
	credB, err := mgr.Mint("seat-b", "planner", false)
	if err != nil {
		t.Fatalf("Mint seat-b: %v", err)
	}
	sock := filepath.Join(os.TempDir(), fmt.Sprintf("frank-lifecycle-%d.sock", time.Now().UnixNano()))
	t.Cleanup(func() { _ = os.Remove(sock) })
	server, err := channel.ServeAuthenticated(sock, mgr, func(meta seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				return json.Marshal([]string{meta.Name})
			},
		}
	})
	if err != nil {
		t.Fatalf("ServeAuthenticated: %v", err)
	}
	t.Cleanup(func() { _ = server.Close() })
	return sock, server, credA, credB
}

func waitForActiveCredential(t *testing.T, ctx context.Context, sock, credential string) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		client, err := channel.DialAuthenticated(ctx, sock, credential)
		if err != nil {
			if strings.Contains(err.Error(), "auth:channel-active") {
				return
			}
			time.Sleep(25 * time.Millisecond)
			continue
		}
		_ = client.Close()
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatalf("credential never became active")
}

func eventuallyDialAuthenticated(t *testing.T, ctx context.Context, sock, credential string) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	var lastErr error
	for time.Now().Before(deadline) {
		client, err := channel.DialAuthenticated(ctx, sock, credential)
		if err == nil {
			_ = client.Close()
			return
		}
		lastErr = err
		if !strings.Contains(err.Error(), "auth:channel-active") && !errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(25 * time.Millisecond)
			continue
		}
		time.Sleep(25 * time.Millisecond)
	}
	t.Fatalf("DialAuthenticated did not recover: %v", lastErr)
}
