package channel_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
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

func TestPushToAuthenticatedSeatOnly(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sock, server, credA, credB := serveLifecycleServer(t)
	first, err := channel.DialAuthenticated(ctx, sock, credA.Value)
	if err != nil {
		t.Fatalf("first DialAuthenticated: %v", err)
	}
	defer func() { _ = first.Close() }()
	second, err := channel.DialAuthenticated(ctx, sock, credB.Value)
	if err != nil {
		t.Fatalf("second DialAuthenticated: %v", err)
	}
	defer func() { _ = second.Close() }()

	frame := []byte(`{"kind":"delivery-nudge","relay_id":"relay-b"}`)
	if err := server.PushTo("seat-b", frame); err != nil {
		t.Fatalf("PushTo: %v", err)
	}
	got, err := second.NextPush(ctx)
	if err != nil {
		t.Fatalf("seat-b NextPush: %v", err)
	}
	if string(got) != string(frame) {
		t.Fatalf("seat-b push = %s, want %s", got, frame)
	}
	expectNoLifecyclePush(t, first)
}

func TestForceCloseSeatDropsAuthenticatedChannel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sock, server, credA, credB := serveLifecycleServer(t)
	first, err := channel.DialAuthenticated(ctx, sock, credA.Value)
	if err != nil {
		t.Fatalf("first DialAuthenticated: %v", err)
	}
	second, err := channel.DialAuthenticated(ctx, sock, credB.Value)
	if err != nil {
		t.Fatalf("second DialAuthenticated: %v", err)
	}
	defer func() { _ = second.Close() }()

	server.ForceCloseSeat("seat-a")

	deadline := time.Now().Add(2 * time.Second)
	var lastErr error
	for time.Now().Before(deadline) {
		_, lastErr = first.ListTools(ctx)
		if errors.Is(lastErr, net.ErrClosed) || strings.Contains(lastErr.Error(), "use of closed network connection") || strings.Contains(lastErr.Error(), "broken pipe") {
			break
		}
		time.Sleep(25 * time.Millisecond)
	}
	if lastErr == nil {
		t.Fatalf("force-closed client still answered")
	}
	if tools, err := second.ListTools(ctx); err != nil {
		t.Fatalf("other seat closed unexpectedly: %v", err)
	} else if len(tools) != 1 || tools[0] != "project" {
		t.Fatalf("other seat tools = %v", tools)
	}
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

func expectNoLifecyclePush(t *testing.T, client *channel.Client) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if frame, err := client.NextPush(ctx); err == nil {
		t.Fatalf("unexpected push: %s", frame)
	}
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
