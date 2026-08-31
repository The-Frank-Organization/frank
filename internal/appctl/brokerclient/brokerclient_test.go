package brokerclient

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appctl/testutil"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestControlLockUsesBoundedPOSIXRecordLockAndOneDescriptor(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broker-control.lock")
	holder := startControlLockHolder(t, path)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	if lock, err := openControlLock(ctx, path); !errors.Is(err, context.DeadlineExceeded) {
		if lock != nil {
			unlockControl(lock)
		}
		t.Fatalf("contended acquisition err = %v, want deadline", err)
	}
	holder.release(t)

	lock, err := openControlLock(context.Background(), path)
	if err != nil {
		t.Fatal(err)
	}
	defer unlockControl(lock)
	before, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if controlLockAvailableInChild(t, path) {
		t.Fatal("child acquired controller-held record lock")
	}
	if duplicate, err := openControlLock(context.Background(), path); !errors.Is(err, errControlLockAlreadyOpen) {
		if duplicate != nil {
			unlockControl(duplicate)
		}
		t.Fatalf("second same-process open err = %v, want duplicate refusal", err)
	}
	if controlLockAvailableInChild(t, path) {
		t.Fatal("second same-process acquisition attempt dropped the live lock")
	}
	after, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(before, after) {
		t.Fatal("controller lock inode changed during its lifetime")
	}
}

func TestControlLockProcessHelper(t *testing.T) {
	mode, path := os.Getenv("FRANK_CONTROL_LOCK_HELPER"), os.Getenv("FRANK_CONTROL_LOCK_PATH")
	if mode == "" {
		return
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		os.Exit(2)
	}
	defer file.Close()
	record := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: 0, Len: 0}
	if err := syscall.FcntlFlock(file.Fd(), syscall.F_SETLK, &record); err != nil {
		os.Exit(10)
	}
	if mode == "probe" {
		_ = syscall.FcntlFlock(file.Fd(), syscall.F_UNLCK, &record)
		return
	}
	if _, err := fmt.Fprintln(os.Stdout, "locked"); err != nil {
		os.Exit(3)
	}
	_, _ = io.Copy(io.Discard, os.Stdin)
	_ = syscall.FcntlFlock(file.Fd(), syscall.F_UNLCK, &record)
}

type controlLockHolder struct {
	command *exec.Cmd
	stdin   io.Closer
}

func startControlLockHolder(t *testing.T, path string) controlLockHolder {
	t.Helper()
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	_ = file.Close()
	command := exec.Command(os.Args[0], "-test.run=^TestControlLockProcessHelper$")
	command.Env = append(os.Environ(), "FRANK_CONTROL_LOCK_HELPER=hold", "FRANK_CONTROL_LOCK_PATH="+path)
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
	if line, err := bufio.NewReader(stdout).ReadString('\n'); err != nil || line != "locked\n" {
		t.Fatalf("lock helper ready = %q err=%v", line, err)
	}
	return controlLockHolder{command: command, stdin: stdin}
}

func (holder controlLockHolder) release(t *testing.T) {
	t.Helper()
	if err := holder.stdin.Close(); err != nil {
		t.Fatal(err)
	}
	if err := holder.command.Wait(); err != nil {
		t.Fatal(err)
	}
}

func controlLockAvailableInChild(t *testing.T, path string) bool {
	t.Helper()
	command := exec.Command(os.Args[0], "-test.run=^TestControlLockProcessHelper$")
	command.Env = append(os.Environ(), "FRANK_CONTROL_LOCK_HELPER=probe", "FRANK_CONTROL_LOCK_PATH="+path)
	err := command.Run()
	if err == nil {
		return true
	}
	var exitError *exec.ExitError
	if errors.As(err, &exitError) && exitError.ExitCode() == 10 {
		return false
	}
	t.Fatalf("control lock probe failed: %v", err)
	return false
}

func TestProposalFoldAndTwoFormGate(t *testing.T) {
	tuple := appipc.EpochStateBody{RunID: "run", GenerationID: "generation", TurnEpoch: "2", LeaseState: appipc.LeaseLeased, StateSeq: "8"}
	gate := NewAssignGate()
	if err := gate.Propose("correlation", tuple); err != nil {
		t.Fatal(err)
	}
	if err := gate.Propose("second", tuple); err == nil {
		t.Fatal("second outstanding proposal accepted")
	}
	result := gate.Fold("correlation", appipc.StateProposalResultBody{ProposalCorrelation: "correlation", Disposition: appipc.ProposalTransitionStarted})
	if result.Action != AwaitEventOrDeadline || gate.Open(tuple) {
		t.Fatalf("transition-started fold = %#v", result)
	}
	event := appipc.EpochInstalledBody{EpochTransitionID: "transition", GenerationID: tuple.GenerationID, TurnEpoch: tuple.TurnEpoch, StateSeq: tuple.StateSeq}
	if !gate.InstallEvent(tuple.RunID, event) || !gate.Open(tuple) || gate.InstallEvent(tuple.RunID, event) {
		t.Fatal("matching event did not open exactly once")
	}

	mismatch := NewAssignGate()
	_ = mismatch.Propose("mismatch", tuple)
	wrong := tuple
	wrong.StateSeq = "9"
	result = mismatch.Fold("mismatch", appipc.StateProposalResultBody{ProposalCorrelation: "mismatch", Disposition: appipc.ProposalInstalled, InstalledState: &wrong})
	if result.Action != InvariantFault || !result.Loud {
		t.Fatalf("installed mismatch = %#v", result)
	}
	if result := mismatch.Fold("dead", appipc.StateProposalResultBody{ProposalCorrelation: "dead", Disposition: appipc.ProposalRejectedStale}); result.Action != Discard {
		t.Fatalf("dead correlation = %#v", result)
	}
}

func TestDurableControlAdvanceAndBrokerEventDedup(t *testing.T) {
	ctx := context.Background()
	db, err := store.Open(ctx, filepath.Join(t.TempDir(), "runtime"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	seedBrokerRun(t, ctx, host)
	client := New(host)
	first, err := client.AdvanceControl(ctx, "run", "token", 10)
	if err != nil || first != "1" {
		t.Fatalf("first advance = %q err=%v", first, err)
	}
	second, err := client.AdvanceControl(ctx, "run", "token", 11)
	if err != nil || second != "2" {
		t.Fatalf("second advance = %q err=%v", second, err)
	}
	event := Event{BrokerInstanceNonce: "nonce", EventSeq: "3", Type: "epoch_installed", RunID: "run", TurnEpoch: "1", Body: []byte(`{"event":"installed"}`), At: 12}
	ack1, duplicate, err := client.RecordEvent(ctx, event)
	if err != nil || duplicate {
		t.Fatalf("first event = %s duplicate=%v err=%v", ack1, duplicate, err)
	}
	ack2, duplicate, err := client.RecordEvent(ctx, event)
	if err != nil || !duplicate || string(ack1) != string(ack2) {
		t.Fatalf("duplicate event = %s duplicate=%v err=%v", ack2, duplicate, err)
	}
}

func TestFakeBrokerUsesRealFramesAndSpawnIsDetached(t *testing.T) {
	left, right := net.Pipe()
	broker, err := testutil.NewFakeBroker(right)
	if err != nil {
		t.Fatal(err)
	}
	peer, err := testutil.NewPeer(left, appipc.ChannelBroker)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	runID := "run"
	epoch := "1"
	done := make(chan error, 1)
	go func() {
		_, err := broker.Send(ctx, testutil.Outbound{Type: "state_proposal", RunID: &runID, TurnEpoch: &epoch, Body: appipc.StateProposalBody{ProposalCorrelation: "c", RunID: runID, GenerationID: "g", TurnEpoch: epoch, LeaseState: appipc.LeaseLeased, StateSeq: "0"}})
		done <- err
	}()
	envelope, err := peer.Receive(ctx)
	if err != nil || envelope.Type != "state_proposal" || <-done != nil {
		t.Fatalf("fakebroker real frame = %#v err=%v", envelope, err)
	}
	tokenRead, tokenWrite, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer tokenRead.Close()
	defer tokenWrite.Close()
	command, err := BrokerCommand("/bin/frank-broker", []string{"--config-home", "/private/config"}, t.TempDir(), tokenRead)
	if err != nil || command.SysProcAttr == nil || !command.SysProcAttr.Setsid || len(command.ExtraFiles) != 1 || len(command.Env) != 0 || fmt.Sprint(command.Args) != "[/bin/frank-broker --config-home /private/config]" {
		t.Fatalf("BrokerCommand = %#v err=%v", command, err)
	}
}

func TestSessionProposeUsesRealBrokerFrames(t *testing.T) {
	left, right := net.Pipe()
	defer left.Close()
	defer right.Close()
	broker, err := testutil.NewFakeBroker(right)
	if err != nil {
		t.Fatal(err)
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		t.Fatal(err)
	}
	session := &Session{Conn: left, registry: registry}
	tuple := appipc.EpochStateBody{RunID: "run", GenerationID: "generation", TurnEpoch: "2", LeaseState: appipc.LeaseUnleased, StateSeq: "9"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		envelope, err := broker.Receive(ctx)
		if err != nil {
			done <- err
			return
		}
		body, ok := envelope.Body.(*appipc.StateProposalBody)
		if !ok || body.ProposalCorrelation != "correlation" || body.GenerationID != tuple.GenerationID {
			done <- fmt.Errorf("proposal=%#v", envelope)
			return
		}
		done <- sendProposalInstalled(ctx, broker, envelope.Seq, tuple)
	}()
	result, err := session.Propose(ctx, "correlation", tuple)
	if err != nil || result.Action != OpenAssign || <-done != nil {
		t.Fatalf("proposal result=%#v err=%v", result, err)
	}
}

func TestSessionProposalTransitionEventAndDeadline(t *testing.T) {
	tuple := appipc.EpochStateBody{RunID: "run", GenerationID: "generation", TurnEpoch: "2", LeaseState: appipc.LeaseUnleased, StateSeq: "9"}
	t.Run("matching install event opens", func(t *testing.T) {
		left, right := net.Pipe()
		defer left.Close()
		defer right.Close()
		broker, _ := testutil.NewFakeBroker(right)
		registry, _ := appipc.NewProtocolRegistry()
		session := &Session{Conn: left, registry: registry}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		go func() {
			request, _ := broker.Receive(ctx)
			_, _ = broker.Send(ctx, testutil.Outbound{Type: "state_proposal_result", Re: &request.Seq, RunID: &tuple.RunID, TurnEpoch: &tuple.TurnEpoch, Body: appipc.StateProposalResultBody{ProposalCorrelation: "correlation", Disposition: appipc.ProposalTransitionStarted}})
			_, _ = broker.Send(ctx, testutil.Outbound{Type: "epoch_installed", RunID: &tuple.RunID, TurnEpoch: &tuple.TurnEpoch, Body: appipc.EpochInstalledBody{EpochTransitionID: "transition", GenerationID: tuple.GenerationID, TurnEpoch: tuple.TurnEpoch, StateSeq: tuple.StateSeq}})
		}()
		result, err := session.Propose(ctx, "correlation", tuple)
		if err != nil || result.Action != OpenAssign {
			t.Fatalf("result=%#v err=%v", result, err)
		}
	})
	t.Run("deadline requests reproposal", func(t *testing.T) {
		left, right := net.Pipe()
		defer left.Close()
		defer right.Close()
		broker, _ := testutil.NewFakeBroker(right)
		registry, _ := appipc.NewProtocolRegistry()
		session := &Session{Conn: left, registry: registry}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
		defer cancel()
		go func() {
			request, _ := broker.Receive(ctx)
			_, _ = broker.Send(ctx, testutil.Outbound{Type: "state_proposal_result", Re: &request.Seq, RunID: &tuple.RunID, TurnEpoch: &tuple.TurnEpoch, Body: appipc.StateProposalResultBody{ProposalCorrelation: "correlation", Disposition: appipc.ProposalRejectedTransitionActive}})
		}()
		result, err := session.Propose(ctx, "correlation", tuple)
		if err != nil || result.Action != Repropose {
			t.Fatalf("result=%#v err=%v", result, err)
		}
	})
}

func sendProposalInstalled(ctx context.Context, broker *testutil.FakeBroker, requestSeq string, tuple appipc.EpochStateBody) error {
	_, err := broker.Send(ctx, testutil.Outbound{
		Type: "state_proposal_result", Re: &requestSeq, RunID: &tuple.RunID, TurnEpoch: &tuple.TurnEpoch,
		Body: appipc.StateProposalResultBody{ProposalCorrelation: "correlation", Disposition: appipc.ProposalInstalled, InstalledState: &tuple},
	})
	return err
}

func TestEstablishLocksAdvancesBeforeDialAndPresentsControl(t *testing.T) {
	ctx := context.Background()
	runtimeDir, err := os.MkdirTemp("/tmp", "s13-broker-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(runtimeDir) })
	db, err := store.Open(ctx, filepath.Join(runtimeDir, "state"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	seedBrokerRun(t, ctx, host)
	client := New(host)
	socketPath := filepath.Join(runtimeDir, "broker-control.sock")
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	dialObserved := make(chan string, 1)
	dial := func(ctx context.Context, network, address string) (net.Conn, error) {
		var stored string
		_, readErr := host.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
			return nil, snapshot.QueryRowContext(ctx, `SELECT control_generation FROM broker_control WHERE singleton=1`).Scan(&stored)
		}))
		if readErr != nil {
			return nil, readErr
		}
		dialObserved <- stored
		return (&net.Dialer{}).DialContext(ctx, network, address)
	}
	type establishResult struct {
		session *Session
		err     error
	}
	result := make(chan establishResult, 1)
	go func() {
		session, err := client.Establish(ctx, ControlRequest{RunID: "run", RuntimeDir: runtimeDir, ControlToken: "token", At: 20, Dial: dial})
		result <- establishResult{session: session, err: err}
	}()
	connection, err := listener.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer connection.Close()
	handshake, err := appipc.ReadFrame(connection)
	if err != nil {
		t.Fatal(err)
	}
	reply, err := appipc.MarshalJCS(map[string]any{"outcome": "adopted"})
	if err != nil || appipc.WriteFrame(connection, reply) != nil {
		t.Fatalf("write handshake reply: %v", err)
	}
	got := <-result
	if got.err != nil {
		t.Fatal(got.err)
	}
	defer got.session.Close()
	if stored := <-dialObserved; stored != fmt.Sprintf("%020d", 1) {
		t.Fatalf("dial preceded durable advance: %q", stored)
	}
	if string(handshake) != `{"control_generation":"1","control_token":"token"}` {
		t.Fatalf("handshake = %s", handshake)
	}
	if got.session.Generation != "1" {
		t.Fatalf("session generation = %q", got.session.Generation)
	}
	if got.session.Outcome != ControlAdopted {
		t.Fatalf("session outcome = %q", got.session.Outcome)
	}
	if _, err := io.WriteString(got.session.Conn, "x"); err != nil {
		t.Fatal(err)
	}
}

func TestEstablishReturnsTypedRejectedHandshake(t *testing.T) {
	ctx := context.Background()
	runtimeDir := t.TempDir()
	db, err := store.Open(ctx, filepath.Join(runtimeDir, "state"))
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	t.Cleanup(func() { _ = host.Close(); _ = db.Close() })
	seedBrokerRun(t, ctx, host)
	left, right := net.Pipe()
	defer right.Close()
	go func() {
		_, _ = appipc.ReadFrame(right)
		reply, _ := appipc.MarshalJCS(map[string]any{"outcome": "rejected-lock"})
		_ = appipc.WriteFrame(right, reply)
	}()
	_, err = New(host).Establish(ctx, ControlRequest{
		RunID: "run", RuntimeDir: runtimeDir, ControlToken: "token", At: 20,
		Dial: func(context.Context, string, string) (net.Conn, error) { return left, nil },
	})
	var rejection *ControlHandshakeError
	if !errors.As(err, &rejection) || rejection.Outcome != ControlRejectedLock {
		t.Fatalf("typed rejection = %#v err=%v", rejection, err)
	}
}

func seedBrokerRun(t *testing.T, ctx context.Context, host *applier.Host) {
	t.Helper()
	_, err := host.Apply(ctx, brokerEventFunc{func(ctx context.Context, tx *store.Tx) error {
		_, err := tx.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, "run", []byte("{}"), fmt.Sprintf("%064d", 0), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1)
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, "run", fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0))
		return err
	}})
	if err != nil {
		t.Fatal(err)
	}
}

type brokerEventFunc struct {
	fn func(context.Context, *store.Tx) error
}

func (event brokerEventFunc) RunID() string { return "run" }
func (event brokerEventFunc) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, event.fn(ctx, tx)
}
