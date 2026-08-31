// Package terminal exposes the deliberately small operator surface for the app
// control plane. Reads are committed snapshots and mutations enter the applier.
package terminal

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

type StartRequest struct {
	Goal          string
	Lane          string
	CredentialRef string
	WorkspaceRoot string
}

type Starter interface {
	Start(context.Context, StartRequest) error
}

type Runner struct {
	applier *applier.Host
	starter Starter
	now     func() time.Time
}

func New(host *applier.Host, starter Starter) *Runner {
	return &Runner{applier: host, starter: starter, now: time.Now}
}

func MutatingVerbs() []string {
	return []string{"run cancel", "run start", "run stop"}
}

func (runner *Runner) Execute(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	if runner == nil || runner.applier == nil {
		fmt.Fprintln(stderr, "frank-app: control plane is unavailable")
		return 2
	}
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: frank-app <status|attempts|tickets|parked|wakes|run start|run stop|run cancel>")
		return 2
	}
	var err error
	switch args[0] {
	case "status", "attempts", "tickets", "parked", "wakes":
		if len(args) != 1 {
			err = errors.New("read-only views accept no arguments")
		} else {
			return runner.view(ctx, args[0], stdout, stderr)
		}
	case "run":
		err = runner.runCommand(ctx, args[1:])
	default:
		err = fmt.Errorf("unknown command %q", strings.Join(args, " "))
	}
	if err != nil {
		fmt.Fprintf(stderr, "frank-app: %v\n", err)
		return 2
	}
	return 0
}

func (runner *Runner) runCommand(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return errors.New("run requires start, stop, or cancel")
	}
	switch args[0] {
	case "start":
		return runner.start(ctx, args[1:])
	case "stop", "cancel":
		return runner.end(ctx, args[0], args[1:])
	default:
		return fmt.Errorf("unknown run command %q", args[0])
	}
}

func (runner *Runner) start(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet("run start", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	request := StartRequest{}
	flags.StringVar(&request.Goal, "goal", "", "goal")
	flags.StringVar(&request.Lane, "lane", "", "lane")
	flags.StringVar(&request.CredentialRef, "credential-ref", "", "opaque credential reference")
	flags.StringVar(&request.WorkspaceRoot, "workspace-root", "", "workspace root")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 || request.Goal == "" || request.Lane == "" || request.CredentialRef == "" || request.WorkspaceRoot == "" {
		return errors.New("run start requires --goal, --lane, --credential-ref, and --workspace-root")
	}
	if runner.starter == nil {
		return errors.New("run start is unavailable")
	}
	return runner.starter.Start(ctx, request)
}

func (runner *Runner) end(ctx context.Context, command string, args []string) error {
	flags := flag.NewFlagSet("run "+command, flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	runID := flags.String("run-id", "", "run identifier")
	hard := flags.Bool("hard", false, "hard cancellation")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if flags.NArg() != 0 || *runID == "" {
		return fmt.Errorf("run %s requires --run-id", command)
	}
	if command == "stop" && *hard {
		return errors.New("--hard is valid only for run cancel")
	}
	_, err := runner.applier.Apply(ctx, terminalEvent{
		runID: *runID,
		state: map[string]string{"stop": "INTERRUPTED", "cancel": "CANCELLED"}[command],
		at:    runner.now().UnixNano(),
	})
	return err
}

type terminalEvent struct {
	runID string
	state string
	at    int64
}

func (event terminalEvent) RunID() string { return event.runID }

func (event terminalEvent) Apply(ctx context.Context, tx *store.Tx) (applier.Result, error) {
	result, err := tx.ExecContext(ctx, `UPDATE runs SET state=?,backoff_until=NULL,updated_at=? WHERE run_id=? AND state NOT IN ('COMPLETED','FAILED','CANCELLED')`, event.state, event.at, event.runID)
	if err != nil {
		return applier.Result{}, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return applier.Result{}, err
	}
	if affected != 1 {
		return applier.Result{}, fmt.Errorf("run %q is absent or terminal", event.runID)
	}
	if _, err := tx.ExecContext(ctx, `UPDATE leases SET state='RELEASED',released_at=? WHERE run_id=? AND state='ACTIVE'`, event.at, event.runID); err != nil {
		return applier.Result{}, err
	}
	return applier.Result{}, nil
}
