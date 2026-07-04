package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/jackli/frank/internal/channel"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/gate"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
)

type config struct {
	Root     string
	Socket   string
	Registry string
}

func main() {
	root := flag.String("root", "", "store root")
	storeRoot := flag.String("store", ".frank-store", "store root")
	socket := flag.String("socket", "", "unix socket path")
	registry := flag.String("registry", filepath.Join("internal", "fieldspec", "registry.json"), "FieldSpec registry path")
	flag.Parse()
	if *root == "" {
		*root = *storeRoot
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx, config{Root: *root, Socket: *socket, Registry: *registry}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, cfg config) error {
	st, err := store.Open(cfg.Root)
	if err != nil {
		return err
	}
	reg, err := fieldspec.Load(cfg.Registry)
	if err != nil {
		return err
	}
	mgr, err := seat.Open(cfg.Root)
	if err != nil {
		return err
	}
	journal, err := intake.Open(cfg.Root)
	if err != nil {
		return err
	}

	loop := engine.New(st, func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		meta := seat.SeatMeta{Name: cmd.Seat, Role: cmd.Role, IsOperator: cmd.IsOperator}
		return engine.SubmitHandler(st, reg, meta)(ctx, cmd)
	})
	go loop.Run(ctx)

	process := func(cmd intake.Cmd) error {
		if _, err := submitThroughLoop(ctx, loop, cmd); err != nil {
			return err
		}
		return gate.Complete(st)
	}
	if err := frankrecover.RunWithProcessor(cfg.Root, process); err != nil {
		return err
	}

	socket := cfg.Socket
	if socket == "" {
		socket = filepath.Join(cfg.Root, "frank.sock")
	}
	var server *channel.Server
	server, err = channel.ServeAuthenticated(socket, mgr, func(meta seat.SeatMeta) channel.ToolSet {
		return channel.ToolSet{
			Submit: func(ctx context.Context, payload json.RawMessage) (json.RawMessage, error) {
				intakeID, err := journal.Append(intake.Cmd{Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, Verb: "submit", Payload: payload})
				if err != nil {
					return nil, err
				}
				out, err := submitThroughLoop(ctx, loop, intake.Cmd{IntakeID: intakeID, Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, Verb: "submit", Payload: payload})
				if err != nil {
					return nil, err
				}
				if err := gate.Complete(st); err != nil {
					return nil, err
				}
				if server != nil {
					_ = server.Push([]byte(`{"kind":"delivery-nudge"}`))
				}
				return json.Marshal(out)
			},
			Project: func(context.Context, json.RawMessage) (json.RawMessage, error) {
				relayIDs, err := st.Project(meta.Name)
				if err != nil {
					return nil, err
				}
				return json.Marshal(relayIDs)
			},
			Read: func(_ context.Context, payload json.RawMessage) (json.RawMessage, error) {
				var req struct {
					RelayID string `json:"relay_id"`
				}
				if err := json.Unmarshal(payload, &req); err != nil {
					return nil, err
				}
				rec, err := st.Read(req.RelayID)
				if err != nil {
					return nil, err
				}
				return json.Marshal(struct {
					Record        record.Record `json:"record"`
					SchemaVersion int           `json:"schema_version"`
				}{Record: rec, SchemaVersion: rec.Envelope.SchemaVersion})
			},
		}
	})
	if err != nil {
		return err
	}
	defer func() { _ = server.Close() }()
	if seats, err := st.PendingDeliverySeats(); err != nil {
		return err
	} else if len(seats) > 0 {
		frame, err := json.Marshal(struct {
			Kind  string   `json:"kind"`
			Seats []string `json:"seats"`
		}{Kind: "recovery-nudge", Seats: seats})
		if err != nil {
			return err
		}
		if err := server.QueuePush(frame); err != nil {
			return err
		}
	}

	<-ctx.Done()
	return nil
}

func submitThroughLoop(ctx context.Context, loop *engine.Loop, cmd intake.Cmd) (engine.Outcome, error) {
	reply := make(chan engine.Outcome, 1)
	select {
	case loop.In <- engine.Job{Cmd: cmd, ReplyCh: reply}:
	case <-ctx.Done():
		return engine.Outcome{}, ctx.Err()
	}
	select {
	case out := <-reply:
		return out, nil
	case <-ctx.Done():
		return engine.Outcome{}, ctx.Err()
	}
}
