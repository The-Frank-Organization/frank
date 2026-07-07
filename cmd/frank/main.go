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

const maxDarwinSocketPathBytes = 100

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
	rootLock, err := store.AcquireRoot(cfg.Root)
	if err != nil {
		return err
	}
	defer func() { _ = rootLock.Release() }()
	st, err := store.Open(cfg.Root)
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
	reg := pinned.Registry
	detectorConfig, err := engine.DetectorConfigFromPinned(pinned)
	if err != nil {
		return err
	}
	journal, err := intake.OpenWithConfig(cfg.Root, pinned.Engine)
	if err != nil {
		return err
	}
	liveTables := tables.NewLive(tables.New())
	publishStoreTables := func() (*tables.T, error) {
		tab, err := tables.Build(st)
		if err != nil {
			return nil, err
		}
		liveTables.Publish(tab)
		return tab, nil
	}
	socket := cfg.Socket
	if socket == "" {
		socket = filepath.Join(cfg.Root, "frank.sock")
	}
	var server *channel.Server
	currentAuthGeneration := func(seatName string) string {
		return tables.CurrentAuthGeneration(liveTables.Snapshot(), seatName)
	}
	completeSeatMint := func(rec record.Record, includeReply bool) (engine.OutcomeExtras, error) {
		return completeSeatMintBinding(rec, mgr, server, socket, includeReply)
	}
	handler := func(ctx context.Context, cmd intake.Cmd) (record.Record, []store.Intent, error) {
		meta := seat.SeatMeta{Name: cmd.Seat, Role: cmd.Role, IsOperator: cmd.IsOperator}
		tab := liveTables.Snapshot()
		// Step-1 detection is exactly S1 + S2 + S3 plus other->A. S3 is
		// input-atom-pending until operator config names a declared target field.
		env := fieldspec.RenderEnv{
			ConfigDigest: pinned.Digest,
			KnownA:       engine.KnownADetector(reg, tab, detectorConfig),
			Turn:         turnContextForSeat(st, tab, meta.Name),
		}
		return engine.SubmitHandlerWithRender(st, reg, meta, env, tab)(ctx, cmd)
	}
	completeTurn := func(st *store.Store) error {
		tab := liveTables.Snapshot()
		if err := obligation.CompleteAuto(st, tab); err != nil {
			return err
		}
		if err := frankgc.Pass(st, tab, pinned.Engine); err != nil {
			return err
		}
		_, err := publishStoreTables()
		return err
	}

	process := func(cmd intake.Cmd) error {
		if _, err := publishStoreTables(); err != nil {
			return err
		}
		if cmd.AuthGeneration != "" {
			current := currentAuthGeneration(cmd.Seat)
			if current != "" && current != cmd.AuthGeneration {
				rec := engine.CredentialSupersededRecord(cmd)
				relayID, err := st.Commit(rec, nil)
				if err != nil {
					return err
				}
				rec.Envelope.RelayID = relayID
				tab := liveTables.Snapshot()
				tab.OnCommit(rec)
				liveTables.Publish(tab)
				return completeTurn(st)
			}
		}
		rec, intents, err := handler(ctx, cmd)
		if err != nil {
			return err
		}
		if rec.Envelope.IntakeID == "" {
			rec.Envelope.IntakeID = cmd.IntakeID
		}
		relayID, err := st.Commit(rec, intents)
		if err != nil {
			return err
		}
		rec.Envelope.RelayID = relayID
		tab := liveTables.Snapshot()
		tab.OnCommit(rec)
		liveTables.Publish(tab)
		if err := completeTurn(st); err != nil {
			return err
		}
		_, err = completeSeatMint(rec, false)
		return err
	}
	result, err := frankrecover.RunWithProcessor(cfg.Root, pinned, process)
	if err != nil {
		return err
	}
	postRecoveryTables, err := publishStoreTables()
	if err != nil {
		return err
	}
	if err := completeMissingSeatMintBindings(postRecoveryTables, mgr); err != nil {
		return err
	}

	if err := validateSocketPath(socket); err != nil {
		return err
	}
	var loop *engine.Loop
	var writer *intake.Writer[engine.Outcome]
	if result.Ready != nil {
		loop = engine.New(st, handler, result.Ready)
		loop.Tables = postRecoveryTables
		loop.CurrentAuthGeneration = currentAuthGeneration
		loop.AfterCommit = func(st *store.Store) error {
			if err := frankgc.Pass(st, loop.Tables, pinned.Engine); err != nil {
				return err
			}
			tab, err := tables.Build(st)
			if err != nil {
				return err
			}
			loop.Tables = tab
			liveTables.Publish(tab)
			return nil
		}
		loop.AfterAccepted = func(rec record.Record) (engine.OutcomeExtras, error) {
			return completeSeatMint(rec, true)
		}
		go loop.Run(ctx)
		writer, err = intake.NewWriter[engine.Outcome](journal, pinned.Engine, result.Ready)
		if err != nil {
			return err
		}
		go writer.Run(ctx, loop.In)
	}
	server, err = channel.ServeAuthenticated(socket, mgr, func(meta seat.SeatMeta) channel.ToolSet {
		meta.AuthGeneration = currentAuthGeneration(meta.Name)
		tools := channelTools(ctx, st, reg, pinned.Digest, liveTables, meta, writer, loop, func(out engine.Outcome) {
			if server == nil || out.State != record.Accepted || out.RelayID == "" {
				return
			}
			rec, err := st.Read(out.RelayID)
			if err != nil {
				return
			}
			frame := deliveryNudgeFrame(out.RelayID)
			recipients, err := store.DeliveryRecipients(rec)
			if err != nil {
				return
			}
			for _, recipient := range recipients {
				_ = server.PushTo(recipient, frame)
			}
		})
		if result.Diag != nil {
			return channel.ReadOnlySurface(result.Diag, tools)
		}
		return channel.FullSurface(result.Ready, tools)
	}, channel.WithAuthPushes(func(meta seat.SeatMeta) [][]byte {
		pending, err := st.PendingDeliveryFor(meta.Name)
		if err != nil || !pending {
			return nil
		}
		return [][]byte{recoveryNudgeFrame()}
	}))
	if err != nil {
		return err
	}
	defer func() { _ = server.Close() }()

	<-ctx.Done()
	return nil
}

func validateSocketPath(socket string) error {
	if len(socket) < maxDarwinSocketPathBytes {
		return nil
	}
	return fmt.Errorf("socket path too long for darwin AF_UNIX limit: len=%d max<%d; choose a short path such as /tmp/frank.sock", len(socket), maxDarwinSocketPathBytes)
}

func deliveryNudgeFrame(relayID string) []byte {
	frame, _ := json.Marshal(struct {
		Kind    string `json:"kind"`
		RelayID string `json:"relay_id"`
	}{Kind: "delivery-nudge", RelayID: relayID})
	return frame
}

func recoveryNudgeFrame() []byte {
	frame, _ := json.Marshal(struct {
		Kind string `json:"kind"`
	}{Kind: "recovery-nudge"})
	return frame
}

func channelTools(ctx context.Context, st *store.Store, reg *fieldspec.Registry, configDigest string, liveTables *tables.Live, meta seat.SeatMeta, writer *intake.Writer[engine.Outcome], loop *engine.Loop, nudge func(engine.Outcome)) channel.ToolSet {
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
			tab := liveTables.Snapshot()
			form, digest := reg.Render(
				fieldspec.RenderEnv{ConfigDigest: configDigest, Turn: turnContextForSeat(st, tab, meta.Name)},
				fieldspec.SeatMeta{Name: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator},
				req.Phase,
				req.Tier,
				lineage.RealGrantState(tab),
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
			reply, _, err := writer.Submit(ctx, intake.Cmd{Seat: meta.Name, Role: meta.Role, IsOperator: meta.IsOperator, AuthGeneration: meta.AuthGeneration, Verb: "submit", Payload: payload})
			if err != nil {
				return nil, err
			}
			var out engine.Outcome
			select {
			case out = <-reply:
			case <-ctx.Done():
				return nil, ctx.Err()
			}
			nudge(out)
			return json.Marshal(out)
		},
		Project: func(_ context.Context, payload json.RawMessage) (json.RawMessage, error) {
			var req struct {
				View string `json:"view,omitempty"`
			}
			_ = json.Unmarshal(payload, &req)
			relayIDs, err := st.ProjectView(meta.Name, req.View)
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
			if view.Record.Headers["record_kind"] == "config_change" && !operatorSeat(meta) {
				return json.Marshal(redactedConfigChangeRead(view.Record))
			}
			return json.Marshal(struct {
				Record        record.Record `json:"record"`
				SchemaVersion int           `json:"schema_version"`
				SourceVersion int           `json:"source_schema_version"`
			}{Record: view.Record, SchemaVersion: view.Record.Envelope.SchemaVersion, SourceVersion: view.SourceVersion})
		},
	}
}

type configChangeRedactedEnvelope struct {
	From          string `json:"from"`
	Role          string `json:"role"`
	SchemaVersion int    `json:"schema_version"`
}

type configChangeRedactedView struct {
	RelayID    string                       `json:"relay_id"`
	Envelope   configChangeRedactedEnvelope `json:"envelope"`
	RecordKind string                       `json:"record_kind"`
	Member     string                       `json:"member"`
	NewDigest  string                       `json:"new_digest"`
	Redacted   string                       `json:"redacted"`
}

func redactedConfigChangeRead(rec record.Record) configChangeRedactedView {
	return configChangeRedactedView{
		RelayID: rec.Envelope.RelayID,
		Envelope: configChangeRedactedEnvelope{
			From:          rec.Envelope.From,
			Role:          rec.Envelope.Role,
			SchemaVersion: rec.Envelope.SchemaVersion,
		},
		RecordKind: rec.Headers["record_kind"],
		Member:     rec.Headers["member"],
		NewDigest:  rec.Headers["new_digest"],
		Redacted:   "config-member-bytes",
	}
}

func operatorSeat(meta seat.SeatMeta) bool {
	return meta.IsOperator || meta.Name == "operator" || meta.Role == "operator"
}

func seatToolDescriptions() map[string]string {
	return map[string]string{
		"submit":  "Submit a stamped governance record through the serialized loop.",
		"project": "List committed relay IDs currently visible to this seat mailbox.",
		"read":    "Read an immutable committed relay record by relay ID.",
	}
}

func turnContextForSeat(st *store.Store, tab *tables.T, seatName string) fieldspec.TurnContext {
	if st == nil || tab == nil || seatName == "" {
		return fieldspec.TurnContext{}
	}
	relayIDs, err := st.Project(seatName)
	if err != nil || len(relayIDs) == 0 {
		return fieldspec.TurnContext{}
	}
	var wokenOn string
	for i := len(relayIDs) - 1; i >= 0; i-- {
		rec, ok := tab.ByRelay[relayIDs[i]]
		if ok && rec.Envelope.DeliveryState == record.Accepted {
			wokenOn = relayIDs[i]
			break
		}
	}
	if wokenOn == "" {
		return fieldspec.TurnContext{}
	}
	ctx := fieldspec.TurnContext{WokenOn: wokenOn}
	if rec, ok := tab.ByRelay[wokenOn]; ok {
		ctx.ActiveDispatch = rec.Envelope.DispatchID
	}
	return ctx
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
	if err := requireGenesisTimeMint(cfg.Root); err != nil {
		return err
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

func completeSeatMintBinding(rec record.Record, mgr *seat.Manager, server *channel.Server, endpoint string, includeReply bool) (engine.OutcomeExtras, error) {
	if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" {
		return engine.OutcomeExtras{}, nil
	}
	req, violation := engine.ParseSeatMintBody(rec.Body, rec.Envelope.From)
	if violation != nil {
		return engine.OutcomeExtras{}, fmt.Errorf("accepted seat_mint failed parse: %s:%s", violation.Field, violation.Class)
	}
	cred, err := mgr.MintOrReplace(req.Seat, req.Role, req.IsOperator)
	if err != nil {
		return engine.OutcomeExtras{}, err
	}
	if server != nil {
		server.ForceCloseSeat(req.Seat)
	}
	if !includeReply {
		return engine.OutcomeExtras{}, nil
	}
	return engine.OutcomeExtras{Credential: cred.Value, Endpoint: endpoint}, nil
}

func completeMissingSeatMintBindings(tab *tables.T, mgr *seat.Manager) error {
	if tab == nil {
		return nil
	}
	for _, rec := range tab.Records {
		if rec.Envelope.DeliveryState != record.Accepted || rec.Headers["record_kind"] != "seat_mint" {
			continue
		}
		req, violation := engine.ParseSeatMintBody(rec.Body, rec.Envelope.From)
		if violation != nil {
			return fmt.Errorf("accepted seat_mint failed parse: %s:%s", violation.Field, violation.Class)
		}
		if mgr.CredentialsFor(req.Seat) != 0 {
			continue
		}
		if _, err := mgr.MintOrReplace(req.Seat, req.Role, req.IsOperator); err != nil {
			return err
		}
	}
	return nil
}

func requireGenesisTimeMint(root string) error {
	entries, err := os.ReadDir(filepath.Join(root, "records"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		if entry.Name() != "genesis.json" {
			return errors.New("admin -mint is genesis-time only; submit record_kind seat_mint for live minting")
		}
	}
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
