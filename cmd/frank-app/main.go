package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/brokerclient"
	"github.com/jackli/frank/internal/appctl/recovery"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appctl/terminal"
	"github.com/jackli/frank/internal/appipc"
)

type recoverFunc func(context.Context, *applier.Host, string) error

func main() {
	os.Exit(execute(context.Background(), os.Args[1:], os.Stdout, os.Stderr, recoverControlPlane, nil))
}

func execute(ctx context.Context, args []string, stdout, stderr io.Writer, recoverState recoverFunc, starter terminal.Starter) int {
	flags := flag.NewFlagSet("frank-app", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	runtimeDir := flags.String("state-dir", "", "private app-control runtime directory")
	if err := flags.Parse(args); err != nil || *runtimeDir == "" {
		fmt.Fprintln(stderr, "frank-app: --state-dir is required before the command")
		return 2
	}
	db, err := store.Open(ctx, *runtimeDir)
	if err != nil {
		fmt.Fprintf(stderr, "frank-app: boot: %v\n", err)
		return 2
	}
	defer db.Close()
	host := applier.New(db, applier.Config{})
	defer host.Close()
	if recoverState == nil {
		recoverState = recoverControlPlane
	}
	if err := recoverState(ctx, host, *runtimeDir); err != nil {
		fmt.Fprintf(stderr, "frank-app: recovery: %v\n", err)
		return 2
	}
	return terminal.New(host, starter).Execute(ctx, flags.Args(), stdout, stderr)
}

func recoverControlPlane(ctx context.Context, host *applier.Host, runtimeDir string) error {
	proposer := &runtimeProposer{host: host, runtimeDir: runtimeDir}
	defer proposer.Close()
	runID, err := firstRecoverableRun(ctx, host)
	if err != nil || runID == "" {
		return err
	}
	if err := proposer.Establish(ctx, runID); err != nil {
		return err
	}
	_, err = recovery.New(host, proposer, generationID, nil).Run(ctx)
	return err
}

type runtimeProposer struct {
	host       *applier.Host
	runtimeDir string
	session    *brokerclient.Session
}

func (proposer *runtimeProposer) Propose(ctx context.Context, correlation string, tuple appipc.EpochStateBody) (brokerclient.FoldResult, error) {
	if proposer.session == nil {
		return brokerclient.FoldResult{}, fmt.Errorf("broker control unavailable: session not established")
	}
	return proposer.session.Propose(ctx, correlation, tuple)
}

func (proposer *runtimeProposer) Establish(ctx context.Context, runID string) error {
	if proposer == nil || proposer.host == nil || proposer.runtimeDir == "" || runID == "" || proposer.session != nil {
		return fmt.Errorf("broker control unavailable: invalid establishment")
	}
	token, err := loadControlToken(ctx, proposer.host)
	if err != nil {
		return err
	}
	session, err := brokerclient.New(proposer.host).Establish(ctx, brokerclient.ControlRequest{
		RunID: runID, RuntimeDir: proposer.runtimeDir, ControlToken: token, At: time.Now().UnixNano(),
	})
	if err != nil {
		return fmt.Errorf("broker control unavailable: %w", err)
	}
	proposer.session = session
	return nil
}

func (proposer *runtimeProposer) Close() error {
	if proposer == nil || proposer.session == nil {
		return nil
	}
	return proposer.session.Close()
}

func loadControlToken(ctx context.Context, host *applier.Host) (string, error) {
	value, err := host.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var token string
		if err := snapshot.QueryRowContext(ctx, `SELECT control_token FROM broker_control WHERE singleton=1`).Scan(&token); err != nil {
			if store.IsNoRows(err) {
				return nil, fmt.Errorf("broker control unavailable: no durable control token")
			}
			return nil, err
		}
		return token, nil
	}))
	if err != nil {
		return "", err
	}
	return value.(string), nil
}

func firstRecoverableRun(ctx context.Context, host *applier.Host) (string, error) {
	value, err := host.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		var runID string
		err := snapshot.QueryRowContext(ctx, `SELECT run_id FROM runs WHERE state NOT IN ('COMPLETED','FAILED','CANCELLED') ORDER BY run_id LIMIT 1`).Scan(&runID)
		if store.IsNoRows(err) {
			return "", nil
		}
		return runID, err
	}))
	if err != nil {
		return "", err
	}
	return value.(string), nil
}

func generationID() string {
	value := make([]byte, 16)
	if _, err := rand.Read(value); err != nil {
		return ""
	}
	return "generation-" + hex.EncodeToString(value)
}
