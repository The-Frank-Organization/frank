package broker

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jackli/frank/internal/appctl/brokerclient"
	"github.com/jackli/frank/internal/appipc"
)

func TestVerifiedControlHandoverOutcomesAndLiveReplacement(t *testing.T) {
	runtimeDir, err := os.MkdirTemp("/tmp", "broker-handover-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(runtimeDir) })
	server := &Server{token: "opaque", registry: mustRegistry(t), proposals: NewProposalEngine()}
	ctx, cancel := context.WithCancel(context.Background())
	readyReader, readyWriter := io.Pipe()
	serveDone := make(chan error, 1)
	go func() {
		err := server.Serve(ctx, runtimeDir, readyWriter)
		_ = readyWriter.CloseWithError(err)
		serveDone <- err
	}()
	if line, err := bufio.NewReader(readyReader).ReadString('\n'); err != nil || !strings.HasPrefix(line, "BROKER_READY nonce=") {
		t.Fatalf("server ready = %q err=%v", line, err)
	}
	t.Cleanup(func() {
		cancel()
		if err := <-serveDone; err != nil {
			t.Errorf("server stop: %v", err)
		}
	})

	socketPath := filepath.Join(runtimeDir, ControlSocketName)
	lockPath := filepath.Join(runtimeDir, "broker-control.lock")
	if outcome := unlockedHandshakeOutcome(t, socketPath, "wrong", "99"); outcome != ControlRejectedLock {
		t.Fatalf("unlocked wrong-token outcome = %q", outcome)
	}
	if outcome := runControllerHelper(t, lockPath, socketPath, "wrong", "1"); outcome != ControlRejectedToken {
		t.Fatalf("wrong-token outcome = %q", outcome)
	}

	old := startControllerHelper(t, lockPath, socketPath, "opaque", "1")
	if old.outcome != ControlAdopted {
		t.Fatalf("initial outcome = %q", old.outcome)
	}
	if _, err := fmt.Fprintln(old.stdin, "unlock"); err != nil {
		t.Fatal(err)
	}
	if !old.scanner.Scan() || old.scanner.Text() != "unlocked" {
		t.Fatalf("old unlock = %q err=%v", old.scanner.Text(), old.scanner.Err())
	}
	if outcome := runControllerHelper(t, lockPath, socketPath, "opaque", "1"); outcome != ControlRejectedGeneration {
		t.Fatalf("stale-generation outcome = %q", outcome)
	}
	if outcome := runControllerHelper(t, lockPath, socketPath, "opaque", "2"); outcome != ControlAdopted {
		t.Fatalf("replacement outcome = %q", outcome)
	}
	if !old.scanner.Scan() || old.scanner.Text() != "closed" {
		t.Fatalf("old session close = %q err=%v", old.scanner.Text(), old.scanner.Err())
	}
	if err := old.command.Wait(); err != nil {
		t.Fatal(err)
	}
	want := []ControlOutcome{ControlRejectedLock, ControlRejectedToken, ControlAdopted, ControlRejectedGeneration, ControlAdopted}
	if got := server.ControlOutcomes(); fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("outcomes = %v, want %v", got, want)
	}
}

func TestControlControllerHelper(t *testing.T) {
	if os.Getenv("FRANK_BROKER_CONTROLLER_HELPER") == "" {
		return
	}
	lockPath, socketPath := os.Getenv("FRANK_BROKER_LOCK_PATH"), os.Getenv("FRANK_BROKER_SOCKET_PATH")
	lock, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		os.Exit(2)
	}
	record := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: 0, Len: 0}
	if err := syscall.FcntlFlock(lock.Fd(), syscall.F_SETLK, &record); err != nil {
		os.Exit(3)
	}
	connection, err := net.Dial("unix", socketPath)
	if err != nil {
		os.Exit(4)
	}
	handshake, _ := appipc.MarshalJCS(map[string]any{
		"control_token": os.Getenv("FRANK_BROKER_TOKEN"), "control_generation": os.Getenv("FRANK_BROKER_GENERATION"),
	})
	if appipc.WriteFrame(connection, handshake) != nil {
		os.Exit(5)
	}
	reply, err := appipc.ReadFrame(connection)
	if err != nil {
		os.Exit(6)
	}
	var result struct {
		Outcome ControlOutcome `json:"outcome"`
	}
	if json.Unmarshal(reply, &result) != nil {
		os.Exit(7)
	}
	fmt.Fprintln(os.Stdout, result.Outcome)
	if os.Getenv("FRANK_BROKER_HOLD") == "" {
		_ = connection.Close()
		_ = syscall.FcntlFlock(lock.Fd(), syscall.F_UNLCK, &record)
		_ = lock.Close()
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() || scanner.Text() != "unlock" {
		os.Exit(8)
	}
	_ = syscall.FcntlFlock(lock.Fd(), syscall.F_UNLCK, &record)
	_ = lock.Close()
	fmt.Fprintln(os.Stdout, "unlocked")
	_ = connection.SetReadDeadline(time.Now().Add(2 * time.Second))
	var one [1]byte
	if _, err := connection.Read(one[:]); err == nil {
		os.Exit(9)
	}
	fmt.Fprintln(os.Stdout, "closed")
}

type controllerHelper struct {
	command *exec.Cmd
	stdin   io.Writer
	scanner *bufio.Scanner
	outcome ControlOutcome
}

func startControllerHelper(t *testing.T, lockPath, socketPath, token, generation string) controllerHelper {
	t.Helper()
	command := exec.Command(os.Args[0], "-test.run=^TestControlControllerHelper$")
	command.Env = append(os.Environ(),
		"FRANK_BROKER_CONTROLLER_HELPER=1", "FRANK_BROKER_HOLD=1", "FRANK_BROKER_LOCK_PATH="+lockPath,
		"FRANK_BROKER_SOCKET_PATH="+socketPath, "FRANK_BROKER_TOKEN="+token, "FRANK_BROKER_GENERATION="+generation,
	)
	stdout, err := command.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdin, err := command.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := command.Start(); err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(stdout)
	if !scanner.Scan() {
		t.Fatalf("controller outcome unavailable: %v", scanner.Err())
	}
	return controllerHelper{command: command, stdin: stdin, scanner: scanner, outcome: ControlOutcome(scanner.Text())}
}

func runControllerHelper(t *testing.T, lockPath, socketPath, token, generation string) ControlOutcome {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestControlControllerHelper$")
	command.Env = append(os.Environ(),
		"FRANK_BROKER_CONTROLLER_HELPER=1", "FRANK_BROKER_LOCK_PATH="+lockPath,
		"FRANK_BROKER_SOCKET_PATH="+socketPath, "FRANK_BROKER_TOKEN="+token, "FRANK_BROKER_GENERATION="+generation,
	)
	output, err := command.Output()
	if err != nil {
		t.Fatalf("controller helper: %v", err)
	}
	line, _, _ := strings.Cut(string(output), "\n")
	return ControlOutcome(line)
}

func unlockedHandshakeOutcome(t *testing.T, socketPath, token, generation string) ControlOutcome {
	t.Helper()
	connection, err := net.Dial("unix", socketPath)
	if err != nil {
		t.Fatal(err)
	}
	defer connection.Close()
	handshake, _ := appipc.MarshalJCS(map[string]any{"control_token": token, "control_generation": generation})
	if err := appipc.WriteFrame(connection, handshake); err != nil {
		t.Fatal(err)
	}
	_ = connection.SetReadDeadline(time.Now().Add(time.Second))
	reply, err := appipc.ReadFrame(connection)
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Outcome ControlOutcome `json:"outcome"`
	}
	if err := json.Unmarshal(reply, &result); err != nil {
		t.Fatal(err)
	}
	return result.Outcome
}

func mustRegistry(t *testing.T) *appipc.Registry {
	t.Helper()
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	return registry
}

func TestControlHandshakeIsCanonicalTokenBoundAndStrictlyIncreasing(t *testing.T) {
	server := &Server{token: "opaque", proposals: NewProposalEngine()}
	handshake := func(token, generation string) []byte {
		wire, err := appipc.MarshalJCS(map[string]any{"control_token": token, "control_generation": generation})
		if err != nil {
			t.Fatal(err)
		}
		return wire
	}
	if generation, err := server.acceptHandshake(handshake("opaque", "2")); err != nil || generation != 2 {
		t.Fatalf("initial handshake=%d err=%v", generation, err)
	}
	if _, err := server.acceptHandshake(handshake("opaque", "2")); err == nil {
		t.Fatal("replayed generation accepted")
	}
	if _, err := server.acceptHandshake(handshake("wrong", "3")); err == nil {
		t.Fatal("wrong token accepted")
	}
	if _, err := server.acceptHandshake([]byte(`{"control_token":"opaque", "control_generation":"3"}`)); err == nil {
		t.Fatal("non-canonical handshake accepted")
	}
}

func TestControlTokenPipeIsSingleLineAndBounded(t *testing.T) {
	for _, invalid := range [][]byte{nil, []byte("token"), []byte("a\nb\n"), append(bytes.Repeat([]byte("x"), 4096), '\n')} {
		if _, err := readControlToken(bytes.NewReader(invalid)); err == nil {
			t.Fatalf("invalid token pipe accepted: len=%d", len(invalid))
		}
	}
	if token, err := readControlToken(bytes.NewBufferString("opaque\n")); err != nil || token != "opaque" {
		t.Fatalf("valid token=%q err=%v", token, err)
	}
}

func TestProductionServerRoundTripsProposalAndTransitionForms(t *testing.T) {
	serverRegistry := mustRegistry(t)
	server := &Server{token: "opaque", registry: serverRegistry, proposals: NewProposalEngine()}
	left, right := net.Pipe()
	defer left.Close()
	done := make(chan error, 1)
	go func() { done <- server.serveVerifiedControl(right) }()
	session := &brokerclient.Session{Conn: left}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bootstrap := tuple("run", "g1", "1", "1")
	if result, err := session.Propose(ctx, "bootstrap", bootstrap); err != nil || result.Action != brokerclient.OpenAssign {
		t.Fatalf("bootstrap result=%#v err=%v", result, err)
	}
	transition := tuple("run", "g2", "2", "2")
	if result, err := session.Propose(ctx, "transition", transition); err != nil || result.Action != brokerclient.OpenAssign {
		t.Fatalf("transition result=%#v err=%v", result, err)
	}
	_ = left.Close()
	if err := <-done; err == nil {
		t.Fatal("closed control connection did not terminate server loop")
	}
}
