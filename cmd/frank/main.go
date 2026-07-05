package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/jackli/frank/internal/channel"
	frankconfig "github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/engine"
	"github.com/jackli/frank/internal/fieldspec"
	frankgc "github.com/jackli/frank/internal/gc"
	"github.com/jackli/frank/internal/intake"
	"github.com/jackli/frank/internal/lineage"
	"github.com/jackli/frank/internal/migrate"
	"github.com/jackli/frank/internal/obligation"
	"github.com/jackli/frank/internal/record"
	frankrecover "github.com/jackli/frank/internal/recover"
	"github.com/jackli/frank/internal/seat"
	"github.com/jackli/frank/internal/store"
	"github.com/jackli/frank/internal/tables"
)

type config struct {
	Root           string
	Socket         string
	Registry       string
	EngineConfig   string
	Init           bool
	MintSeat       string
	MintRole       string
	MintOperator   bool
	OperatorSubmit string
	Credential     string
}

func main() {
	root := flag.String("root", "", "store root")
	storeRoot := flag.String("store", ".frank-store", "store root")
	socket := flag.String("socket", "", "unix socket path")
	registry := flag.String("registry", filepath.Join("internal", "fieldspec", "registry.json"), "FieldSpec registry path")
	engineConfig := flag.String("engine-config", "", "engine config path for -init")
	initStore := flag.Bool("init", false, "initialize a store with pinned config")
	mintSeat := flag.String("mint", "", "mint a conductor-internal credential for a seat")
	mintRole := flag.String("role", "implementer", "role for -mint")
	mintOperator := flag.Bool("operator", false, "mint an operator credential")
	operatorSubmit := flag.String("operator-submit", "", "submit a payload JSON file through an authenticated socket")
	credential := flag.String("credential", "", "credential for operator-submit")
	flag.Parse()
	if *root == "" {
		*root = *storeRoot
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	cfg := config{Root: *root, Socket: *socket, Registry: *registry, EngineConfig: *engineConfig, Init: *initStore, MintSeat: *mintSeat, MintRole: *mintRole, MintOperator: *mintOperator, OperatorSubmit: *operatorSubmit, Credential: *credential}
	if err := run(ctx, cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, cfg config) error {
	if cfg.Init {
		if cfg.EngineConfig == "" {
			return errors.New("engine-config required for init")
		}
		return store.Init(cfg.Root, map[string]string{"fieldspec": cfg.Registry, "engine": cfg.EngineConfig})
	}
	if cfg.MintSeat != "" {
		return mintSeat(ctx, cfg)
	}
	if cfg.OperatorSubmit != "" {
		return operatorSubmit(ctx, cfg)
	}
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
	pinned, err := frankconfig.Load(store.StoreRootConfigPaths(cfg.Root))
	if err != nil {
		return err
	}
	journal, err := intake.OpenWithConfig(cfg.Root, pinned.Engine)
	if err != nil {
		return err
	}
	liveTables, err := tables.Build(st)
	if err != nil {
		return err
	}
	handler := func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		meta := seat.SeatMeta{Name: cmd.Seat, Role: cmd.Role, IsOperator: cmd.IsOperator}
		env := fieldspec.RenderEnv{ConfigDigest: pinned.Digest, ParentCandidates: lineage.ActiveLineageCandidates(liveTables, lineage.TurnContext{})}
		return engine.SubmitHandlerWithRender(st, reg, meta, env, liveTables)(ctx, cmd)
	}
	completeTurn := func(st *store.Store) error {
		if err := obligation.CompleteAuto(st, liveTables); err != nil {
			return err
		}
		return frankgc.Pass(st, liveTables, pinned.Engine)
	}

	process := func(cmd intake.Cmd) error {
		rec, intents, err := handler(ctx, cmd)
		if err != nil {
			return err
		}
		if rec.Envelope.IntakeID == "" {
			rec.Envelope.IntakeID = cmd.IntakeID
		}
		if _, err := st.Commit(rec, intents); err != nil {
			return err
		}
		liveTables.OnCommit(rec)
		return completeTurn(st)
	}
	result, err := frankrecover.RunWithProcessor(cfg.Root, pinned, process)
	if err != nil {
		return err
	}

	socket := cfg.Socket
	if socket == "" {
		socket = filepath.Join(cfg.Root, "frank.sock")
	}
	var server *channel.Server
	var loop *engine.Loop
	var writer *intake.Writer[engine.Outcome]
	if result.Ready != nil {
		loop = engine.New(st, handler, result.Ready)
		loop.Tables = liveTables
		loop.AfterCommit = func(st *store.Store) error {
			return frankgc.Pass(st, liveTables, pinned.Engine)
		}
		go loop.Run(ctx)
		writer, err = intake.NewWriter[engine.Outcome](journal, pinned.Engine, result.Ready)
		if err != nil {
			return err
		}
		go writer.Run(ctx, loop.In)
	}
	server, err = channel.ServeAuthenticated(socket, mgr, func(meta seat.SeatMeta) channel.ToolSet {
		tools := channelTools(ctx, st, reg, pinned.Digest, liveTables, meta, writer, loop, func() {
			if server != nil {
				_ = server.Push([]byte(`{"kind":"delivery-nudge"}`))
			}
		})
		if result.Diag != nil {
			return channel.ReadOnlySurface(result.Diag, tools)
		}
		return channel.FullSurface(result.Ready, tools)
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

func channelTools(ctx context.Context, st *store.Store, reg *fieldspec.Registry, configDigest string, liveTables *tables.T, meta seat.SeatMeta, writer *intake.Writer[engine.Outcome], loop *engine.Loop, nudge func()) channel.ToolSet {
	return channel.ToolSet{
		Describe: func(_ context.Context, payload json.RawMessage) (json.RawMessage, error) {
			var req channel.DescribeRequest
			_ = json.Unmarshal(payload, &req)
			if req.Phase == "" {
				req.Phase = "SITREP"
			}
			if req.Tier == "" {
				req.Tier = "medium"
			}
			form, digest := reg.Render(
				fieldspec.RenderEnv{ConfigDigest: configDigest, ParentCandidates: lineage.ActiveLineageCandidates(liveTables, lineage.TurnContext{})},
				fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator},
				req.Phase,
				req.Tier,
				lineage.RealGrantState(liveTables),
			)
			return json.Marshal(channel.DescriptionResponse{
				Tools:        []string{"submit", "project", "read"},
				Descriptions: seatToolDescriptions(),
				SubmitSchema: &form,
				FormDigest:   digest,
			})
		},
		Submit: func(ctx context.Context, payload json.RawMessage) (json.RawMessage, error) {
			if writer == nil {
				return nil, errors.New("submit unavailable")
			}
			reply, _, err := writer.Submit(ctx, intake.Cmd{Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, Verb: "submit", Payload: payload})
			if err != nil {
				return nil, err
			}
			var out engine.Outcome
			select {
			case out = <-reply:
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			nudge()
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
			view, err := (migrate.Reader{Store: st, Registry: migrate.New()}).Read(req.RelayID)
			if err != nil {
				var checksum store.ErrChecksum
				if errors.As(err, &checksum) {
					if loop != nil {
						loop.EnqueueQuarantine(checksum.RelayID)
					}
					return json.Marshal(struct {
						ErrorClass string `json:"error_class"`
						RelayID    string `json:"relay_id"`
					}{ErrorClass: "checksum-mismatch", RelayID: checksum.RelayID})
				}
				var quarantined store.ErrQuarantined
				if errors.As(err, &quarantined) {
					return json.Marshal(struct {
						ErrorClass   string `json:"error_class"`
						RelayID      string `json:"relay_id"`
						IncidentID   string `json:"incident_id,omitempty"`
						FailureClass string `json:"failure_class"`
					}{ErrorClass: "record-quarantined", RelayID: quarantined.RelayID, IncidentID: quarantined.IncidentID, FailureClass: quarantined.FailureClass})
				}
				return nil, err
			}
			return json.Marshal(struct {
				Record        record.Record `json:"record"`
				SchemaVersion int           `json:"schema_version"`
				SourceVersion int           `json:"source_schema_version"`
			}{Record: view.Record, SchemaVersion: view.Record.Envelope.SchemaVersion, SourceVersion: view.SourceVersion})
		},
	}
}

func seatToolDescriptions() map[string]string {
	return map[string]string{
		"submit":  "Submit a stamped governance record through the serialized loop.",
		"project": "List committed relay IDs currently visible to this seat mailbox.",
		"read":    "Read an immutable committed relay record by relay ID.",
	}
}

func mintSeat(ctx context.Context, cfg config) error {
	if cfg.MintRole == "" {
		return errors.New("role required for mint")
	}
	socket := cfg.Socket
	if socket == "" {
		socket = filepath.Join(cfg.Root, "frank.sock")
	}
	if socketIsLive(ctx, socket) {
		return errors.New("conductor is serving; -mint is admin-time only")
	}
	mgr, err := seat.Open(cfg.Root)
	if err != nil {
		return err
	}
	cred, err := mgr.Mint(cfg.MintSeat, cfg.MintRole, cfg.MintOperator)
	if err != nil {
		return err
	}
	fmt.Printf("credential=%s\n", cred.Value)
	return nil
}

func socketIsLive(ctx context.Context, socket string) bool {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "unix", socket)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

func operatorSubmit(ctx context.Context, cfg config) error {
	if cfg.Credential == "" {
		return errors.New("credential required for operator-submit")
	}
	socket := cfg.Socket
	if socket == "" {
		socket = filepath.Join(cfg.Root, "frank.sock")
	}
	payload, err := os.ReadFile(cfg.OperatorSubmit)
	if err != nil {
		return err
	}
	client, err := channel.DialAuthenticated(ctx, socket, cfg.Credential)
	if err != nil {
		return err
	}
	defer func() { _ = client.Close() }()
	result, err := client.Call(ctx, "submit", payload)
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	return nil
}
