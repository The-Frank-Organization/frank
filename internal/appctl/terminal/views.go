package terminal

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/The-Frank-Organization/frank/internal/appctl/applier"
	"github.com/The-Frank-Organization/frank/internal/appctl/store"
)

type viewResult struct {
	lines  []string
	alerts []string
}

type rowSet interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(...any) error
}

func (runner *Runner) view(ctx context.Context, name string, stdout, stderr io.Writer) int {
	value, err := runner.applier.Read(ctx, applier.QueryFunc(func(ctx context.Context, snapshot *store.Snapshot) (any, error) {
		return queryView(ctx, snapshot, name)
	}))
	if err != nil {
		fmt.Fprintf(stderr, "frank-app: %v\n", err)
		return 2
	}
	result := value.(viewResult)
	for _, line := range result.lines {
		fmt.Fprintln(stdout, line)
	}
	for _, alert := range result.alerts {
		fmt.Fprintln(stderr, alert)
	}
	if len(result.alerts) > 0 {
		return 1
	}
	return 0
}

func queryView(ctx context.Context, snapshot *store.Snapshot, name string) (viewResult, error) {
	result := viewResult{}
	var rows rowSet
	var err error
	switch name {
	case "status":
		rows, err = snapshot.QueryContext(ctx, `SELECT run_id,state,run_phase,COALESCE(stop_reason,''),COALESCE(resume_action,''),consecutive_failures,COALESCE(CAST(backoff_until AS TEXT),'') FROM runs ORDER BY run_id`)
	case "attempts":
		rows, err = snapshot.QueryContext(ctx, `SELECT attempt_id,run_id,turn_id,state,logical_surface_digest,COALESCE(frozen_core_digest,''),COALESCE(provider_lowered_tools_digest,''),COALESCE(cancel_point,'') FROM provider_attempts ORDER BY attempt_id`)
	case "tickets":
		rows, err = snapshot.QueryContext(ctx, `SELECT ticket_id,run_id,turn_id,tool_call_id,state,COALESCE(void_reason,''),canonical_tool_name,canonical_args_digest FROM tool_authorizations ORDER BY ticket_id`)
	case "parked":
		rows, err = snapshot.QueryContext(ctx, `SELECT tool_call_id,run_id,turn_id,state,canonical_tool_name,canonical_args_digest FROM tool_calls WHERE state IN ('UNKNOWN_TOOL_OUTCOME','PARTIAL_TOOL_EFFECT') ORDER BY tool_call_id`)
	case "wakes":
		rows, err = snapshot.QueryContext(ctx, `SELECT relay_id,run_id,disposition,COALESCE(admitted_turn_id,'') FROM wake_schedule ORDER BY relay_id`)
	default:
		return result, fmt.Errorf("unknown view %q", name)
	}
	if err != nil {
		return result, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return result, err
	}
	for rows.Next() {
		values := make([]string, len(columns))
		dest := make([]any, len(columns))
		for index := range values {
			dest[index] = &values[index]
		}
		if err := rows.Scan(dest...); err != nil {
			return result, err
		}
		result.lines = append(result.lines, strings.Join(values, "\t"))
		if name == "status" && values[1] == "FAILED" {
			result.alerts = append(result.alerts, fmt.Sprintf("ALERT %s state=FAILED stop_reason=%s", values[0], values[3]))
		}
	}
	return result, rows.Err()
}
