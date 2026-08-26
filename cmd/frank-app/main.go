package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackli/frank/internal/appctl/applier"
	"github.com/jackli/frank/internal/appctl/brokerclient"
	"github.com/jackli/frank/internal/appctl/manifest"
	"github.com/jackli/frank/internal/appctl/recovery"
	"github.com/jackli/frank/internal/appctl/store"
	"github.com/jackli/frank/internal/appctl/supervisor"
	"github.com/jackli/frank/internal/appctl/terminal"
	"github.com/jackli/frank/internal/appipc"
	connectorcatalog "github.com/jackli/frank/internal/connector/catalog"
	connectorpolicy "github.com/jackli/frank/internal/connector/policy"
)

type recoverFunc func(context.Context, *applier.Host, string) error
type starterFactory func(*applier.Host, string, io.Writer) terminal.Starter

func main() {
	os.Exit(execute(context.Background(), os.Args[1:], os.Stdout, os.Stderr, recoverControlPlane, newProductionStarter))
}

func execute(ctx context.Context, args []string, stdout, stderr io.Writer, recoverState recoverFunc, makeStarter starterFactory) int {
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
	var starter terminal.Starter
	if makeStarter != nil {
		starter = makeStarter(host, *runtimeDir, stdout)
	}
	return terminal.New(host, starter).Execute(ctx, flags.Args(), stdout, stderr)
}

type productionStarter struct {
	host           *applier.Host
	runtimeDir     string
	stdout         io.Writer
	connectorBin   string
	credentialPath string
	catalogPath    string
	policyPath     string
}

func newProductionStarter(host *applier.Host, runtimeDir string, stdout io.Writer) terminal.Starter {
	return &productionStarter{
		host: host, runtimeDir: runtimeDir, stdout: stdout,
		connectorBin: os.Getenv("FRANK_CONNECTOR_BIN"), credentialPath: os.Getenv("FRANK_CONNECTOR_CREDENTIAL"),
		catalogPath: os.Getenv("FRANK_CONNECTOR_CATALOG"), policyPath: os.Getenv("FRANK_CONNECTOR_POLICY"),
	}
}

func (starter *productionStarter) Start(ctx context.Context, request terminal.StartRequest) error {
	if starter == nil || starter.host == nil || starter.runtimeDir == "" || starter.stdout == nil {
		return errors.New("app control start: invalid composition")
	}
	if starter.connectorBin == "" || starter.credentialPath == "" || starter.catalogPath == "" || starter.policyPath == "" {
		return errors.New("app control start: connector configuration is unavailable")
	}
	catalogBytes, err := os.ReadFile(starter.catalogPath)
	if err != nil {
		return fmt.Errorf("app control start: load lane catalog: %w", err)
	}
	loadedCatalog, err := connectorcatalog.Load(catalogBytes)
	if err != nil {
		return fmt.Errorf("app control start: load lane catalog: %w", err)
	}
	var lane connectorcatalog.Lane
	for _, candidate := range loadedCatalog.Lanes {
		if candidate.LaneID == request.Lane {
			lane = candidate
			break
		}
	}
	if lane.LaneID == "" {
		return fmt.Errorf("app control start: lane %q is unavailable", request.Lane)
	}
	policyBytes, err := os.ReadFile(starter.policyPath)
	if err != nil {
		return fmt.Errorf("app control start: load egress policy: %w", err)
	}
	loadedPolicy, err := connectorpolicy.Load(policyBytes, lane)
	if err != nil {
		return fmt.Errorf("app control start: load egress policy: %w", err)
	}
	if err := loadedCatalog.ValidateDeniedHeaders(loadedPolicy.DeniedHeaderNames); err != nil {
		return fmt.Errorf("app control start: compose lane policy: %w", err)
	}
	tools, toolCatalogDigest := productionToolIdentity()
	releaseDigest := sha256Hex([]byte("frank-mvp-development"))
	runID := "run-" + strings.TrimPrefix(generationID(), "generation-")
	laneID := manifest.LaneID{ModelID: lane.ModelID, ProviderID: lane.ProviderID, ServingProfileID: lane.ServingProfileID, CompatMode: lane.CompatMode}
	frozen, err := manifest.Build(manifest.BuildInput{
		RunID: runID, PolicySourceRef: "m3-egress-policy", PolicyDigest: loadedPolicy.Digest, PolicyBytes: policyBytes,
		PolicyPinnedLane: laneID, ToolSet: tools, ToolCatalogDigest: &toolCatalogDigest,
		ProviderLane:   manifest.ProviderLane{LaneID: laneID, LaneCatalogDigest: loadedCatalog.Digest, CredentialRef: request.CredentialRef},
		WorkspaceRoot:  request.WorkspaceRoot,
		ReleaseBinding: &manifest.ReleaseBinding{BoundAtRef: "working-tree", ReleaseDigest: &releaseDigest},
	})
	if err != nil {
		return fmt.Errorf("app control start: build manifest: %w", err)
	}
	gate := manifest.Gate{LockedTools: tools, ShippedToolCatalogDigest: toolCatalogDigest, PolicyBytes: policyBytes, LaneCatalogDigest: loadedCatalog.Digest}
	if _, err := starter.host.Apply(ctx, manifest.FreezeEvent{Frozen: frozen, Gate: gate, SessionLogPath: filepath.Join(starter.runtimeDir, runID+".session"), CreatedAt: time.Now().UnixNano()}); err != nil {
		return fmt.Errorf("app control start: admit run: %w", err)
	}
	assignment, err := frozen.ConnectorAssign()
	if err != nil {
		return fmt.Errorf("app control start: connector assignment: %w", err)
	}
	assignment.ProviderLaneID = lane.LaneID
	assignment.TurnEpoch = "1"
	process, err := supervisor.LaunchConnector(ctx, supervisor.ConnectorLaunch{
		BinaryPath: starter.connectorBin, RuntimeDir: filepath.Dir(starter.credentialPath),
		CredentialPath: starter.credentialPath, CatalogPath: starter.catalogPath, PolicyPath: starter.policyPath,
		BuildInfo: "s16a-wp3", Assignment: assignment,
	})
	if err != nil {
		return fmt.Errorf("app control start: %w", err)
	}
	if process.State() != supervisor.WorkerReady {
		return errors.New("app control start: connector did not reach READY")
	}
	if _, err := fmt.Fprintln(starter.stdout, "CONNECTOR_READY"); err != nil {
		return err
	}
	closeCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := process.Close(closeCtx); err != nil {
		return fmt.Errorf("app control start: stop connector probe: %w", err)
	}
	return nil
}

func productionToolIdentity() ([]manifest.ToolIdentity, string) {
	tools := manifest.StagingToolSet()
	for index := range tools {
		schemaDigest := sha256Hex([]byte("schema:" + tools[index].Name))
		catalogVersion := "mvp-v1"
		tools[index].SchemaDigest = &schemaDigest
		tools[index].CatalogVersion = &catalogVersion
		if strings.HasPrefix(tools[index].Name, "relay.") {
			mappingVersion := "mvp-v1"
			tools[index].MappingVersion = &mappingVersion
		}
	}
	return tools, sha256Hex([]byte(strings.Join(manifest.RatifiedToolNames, "\n")))
}

func sha256Hex(value []byte) string {
	digest := sha256.Sum256(value)
	return hex.EncodeToString(digest[:])
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
	client := brokerclient.New(proposer.host)
	session, err := client.Establish(ctx, brokerclient.ControlRequest{
		RunID: runID, RuntimeDir: proposer.runtimeDir, ControlToken: token, At: time.Now().UnixNano(),
	})
	if err == nil {
		proposer.session = session
		return nil
	}
	brokerBin, configHome := os.Getenv("FRANK_BROKER_BIN"), os.Getenv("FRANK_BROKER_CONFIG_HOME")
	if brokerBin == "" || configHome == "" {
		return fmt.Errorf("broker control unavailable: %w", err)
	}
	newToken, instanceID := controlToken(), generationID()
	if newToken == "" {
		return fmt.Errorf("broker control unavailable: token mint failed")
	}
	process, launchErr := supervisor.LaunchBroker(ctx, supervisor.BrokerLaunch{
		BinaryPath: brokerBin, RuntimeDir: proposer.runtimeDir, ConfigHome: configHome,
		RunID: runID, ControlToken: newToken, At: time.Now().UnixNano(), Client: client,
		Controller: supervisor.New(proposer.host), InstanceID: instanceID,
	})
	if launchErr != nil {
		return fmt.Errorf("broker control unavailable after adoption miss: %w", launchErr)
	}
	session, establishErr := client.Establish(ctx, brokerclient.ControlRequest{
		RunID: runID, RuntimeDir: proposer.runtimeDir, ControlToken: newToken, At: time.Now().UnixNano(),
	})
	if establishErr != nil {
		closeCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = process.Close(closeCtx)
		return fmt.Errorf("broker control unavailable after fresh spawn: %w", establishErr)
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

func controlToken() string {
	value := make([]byte, 32)
	if _, err := rand.Read(value); err != nil {
		return ""
	}
	return hex.EncodeToString(value)
}
