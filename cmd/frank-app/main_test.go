package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appctl/terminal"
	"github.com/The-Frank-Organization/frank/internal/appctl/testutil"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestControlTokenMintIsThirtyTwoCSPRNGBytesAsLowerHex(t *testing.T) {
	pattern := regexp.MustCompile(`^[0-9a-f]{64}$`)
	seen := make(map[string]struct{})
	for i := 0; i < 32; i++ {
		token := controlToken()
		if !pattern.MatchString(token) {
			t.Fatalf("control token %q is not 32 bytes as lowercase hex", token)
		}
		seen[token] = struct{}{}
	}
	if len(seen) != 32 {
		t.Fatalf("control token mint repeated: %d distinct of 32", len(seen))
	}
}

func TestExecuteComposesProductionStarterAfterStoreOpen(t *testing.T) {
	called := false
	factory := func(host *applier.Host, runtimeDir string, _ io.Writer) terminal.Starter {
		called = host != nil && runtimeDir != ""
		return starterFunc(func(context.Context, terminal.StartRequest) error { return errors.New("starter reached") })
	}
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	code := execute(context.Background(), []string{"--state-dir", filepath.Join(t.TempDir(), "state"), "run", "start", "--goal", "probe", "--lane", "lane", "--credential-ref", "credential-ref", "--workspace-root", t.TempDir()}, stdout, stderr, func(context.Context, *applier.Host, string) error { return nil }, factory)
	if code == 0 || !called || !strings.Contains(stderr.String(), "starter reached") {
		t.Fatalf("code=%d called=%v stdout=%s stderr=%s", code, called, stdout, stderr)
	}
}

func TestExecuteAlwaysRecoversBeforeCommand(t *testing.T) {
	called := 0
	recoverState := func(context.Context, *applier.Host, string) error {
		called++
		return nil
	}
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	code := execute(context.Background(), []string{"--state-dir", filepath.Join(t.TempDir(), "state"), "status"}, stdout, stderr, recoverState, nil)
	if code != 0 || called != 1 {
		t.Fatalf("code=%d recovery calls=%d stderr=%s", code, called, stderr)
	}
}

func TestExecuteRecoveryFailureIsLoud(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	code := execute(context.Background(), []string{"--state-dir", filepath.Join(t.TempDir(), "state"), "status"}, stdout, stderr, func(context.Context, *applier.Host, string) error {
		return errors.New("durable tuple refused")
	}, nil)
	if code == 0 || !strings.Contains(stderr.String(), "recovery: durable tuple refused") {
		t.Fatalf("code=%d stdout=%s stderr=%s", code, stdout, stderr)
	}
}

func TestProductionRecoveryRunsOnEmptyGenesis(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	code := execute(context.Background(), []string{"--state-dir", filepath.Join(t.TempDir(), "state"), "status"}, stdout, stderr, recoverControlPlane, nil)
	if code != 0 {
		t.Fatalf("code=%d stdout=%s stderr=%s", code, stdout, stderr)
	}
}

func TestProductionRecoveryEstablishesThenPublishesCurrentTuple(t *testing.T) {
	runtimeDir, err := os.MkdirTemp("/tmp", "s13-app-recovery-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(runtimeDir) })
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	db, err := store.Open(ctx, runtimeDir)
	if err != nil {
		t.Fatal(err)
	}
	host := applier.New(db, applier.Config{})
	_, err = host.Apply(ctx, mainEvent{runID: "run", fn: func(ctx context.Context, tx *store.Tx) error {
		if _, err := tx.ExecContext(ctx, `INSERT INTO runs(run_id,manifest_bytes,run_manifest_digest,state,run_phase,consecutive_failures,created_at) VALUES(?,?,?,?,?,?,?)`, "run", []byte("{}"), strings.Repeat("0", 64), "ACTIVE", "established", fmt.Sprintf("%020d", 0), 1); err != nil {
			return err
		}
		if _, err := tx.ExecContext(ctx, `INSERT INTO epochs(run_id,turn_epoch,state_seq) VALUES(?,?,?)`, "run", fmt.Sprintf("%020d", 1), fmt.Sprintf("%020d", 0)); err != nil {
			return err
		}
		_, err := tx.ExecContext(ctx, `INSERT INTO broker_control(singleton,control_token,control_generation,minted_at) VALUES(1,?,?,?)`, "opaque-control-token", fmt.Sprintf("%020d", 0), 1)
		return err
	}})
	if err != nil {
		t.Fatal(err)
	}
	_ = host.Close()
	_ = db.Close()

	listener, err := net.Listen("unix", filepath.Join(runtimeDir, "broker-control.sock"))
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		connection, err := listener.Accept()
		if err != nil {
			done <- err
			return
		}
		defer connection.Close()
		if _, err := appipc.ReadFrame(connection); err != nil {
			done <- err
			return
		}
		handshakeReply, err := appipc.MarshalJCS(map[string]any{"outcome": "adopted"})
		if err != nil {
			done <- err
			return
		}
		if err := appipc.WriteFrame(connection, handshakeReply); err != nil {
			done <- fmt.Errorf("handshake reply: %w", err)
			return
		}
		broker, err := testutil.NewFakeBroker(connection)
		if err != nil {
			done <- err
			return
		}
		request, err := broker.Receive(ctx)
		if err != nil {
			done <- err
			return
		}
		body, ok := request.Body.(*appipc.StateProposalBody)
		if !ok || body.StateSeq != "3" || body.TurnEpoch != "1" || body.LeaseState != appipc.LeaseUnleased {
			done <- fmt.Errorf("proposal tuple=%#v", request.Body)
			return
		}
		tuple := appipc.EpochStateBody{RunID: body.RunID, GenerationID: body.GenerationID, TurnEpoch: body.TurnEpoch, LeaseState: body.LeaseState, StateSeq: body.StateSeq}
		_, err = broker.Send(ctx, testutil.Outbound{Type: "state_proposal_result", Re: &request.Seq, RunID: &body.RunID, TurnEpoch: &body.TurnEpoch, Body: appipc.StateProposalResultBody{ProposalCorrelation: body.ProposalCorrelation, Disposition: appipc.ProposalInstalled, InstalledState: &tuple}})
		done <- err
	}()
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	code := execute(ctx, []string{"--state-dir", runtimeDir, "status"}, stdout, stderr, recoverControlPlane, nil)
	if brokerErr := <-done; code != 0 || brokerErr != nil {
		t.Fatalf("code=%d broker=%v stdout=%s stderr=%s", code, brokerErr, stdout, stderr)
	}
	verified, err := store.Open(ctx, runtimeDir)
	if err != nil {
		t.Fatal(err)
	}
	defer verified.Close()
	var token, generation, failures string
	if err := verified.Read(ctx, func(snapshot *store.Snapshot) error {
		return snapshot.QueryRowContext(ctx, `SELECT b.control_token,b.control_generation,r.consecutive_failures FROM broker_control b CROSS JOIN runs r WHERE b.singleton=1 AND r.run_id='run'`).Scan(&token, &generation, &failures)
	}); err != nil {
		t.Fatal(err)
	}
	if token != "opaque-control-token" || generation != fmt.Sprintf("%020d", 1) || failures != fmt.Sprintf("%020d", 0) {
		t.Fatalf("adoption token=%q generation=%q failures=%q", token, generation, failures)
	}
}

func TestExecuteRequiresStateDirectory(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	if code := execute(context.Background(), []string{"status"}, stdout, stderr, nil, nil); code == 0 || !strings.Contains(stderr.String(), "--state-dir") {
		t.Fatalf("code=%d stdout=%s stderr=%s", code, stdout, stderr)
	}
}

type mainEvent struct {
	runID string
	fn    func(context.Context, *store.Tx) error
}

type starterFunc func(context.Context, terminal.StartRequest) error

func (start starterFunc) Start(ctx context.Context, request terminal.StartRequest) error {
	return start(ctx, request)
}

func (event mainEvent) RunID() string { return event.runID }
func (event mainEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	return applier.Result{}, event.fn(ctx, tx)
}
