// Package control owns the connector side of the private CTRL-C lifecycle.
// It deliberately carries no worker generation identity.
package control

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/jackli/frank/internal/connector/catalog"
	"github.com/jackli/frank/internal/connector/credentials"
	"github.com/jackli/frank/internal/connector/frame"
	"github.com/jackli/frank/internal/connector/jcs"
	"github.com/jackli/frank/internal/connector/outcome"
	"github.com/jackli/frank/internal/connector/policy"
)

const (
	HandshakeDeadline      = 10 * time.Second
	EpochQueryHoldDeadline = 100 * time.Millisecond
)

var (
	ErrHandshake  = errors.New("connector: handshake failed")
	ErrNotReady   = errors.New("connector: READY gate closed")
	ErrControlEOF = errors.New("connector: CTRL-C EOF")
	ErrShutdown   = errors.New("connector: shutdown requested")
)

type Artifacts struct {
	CredentialPath string
	CatalogPath    string
	PolicyPath     string
}

type Config struct {
	Artifacts      Artifacts
	BuildInfo      string
	PID            int
	HandshakeBound time.Duration
}

type BuildInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	BuiltAt string `json:"built_at"`
}

// Assign is the seven-field connector bootstrap body. It intentionally has no
// generation identifier: m-8 is generation-blind.
type Assign struct {
	RunID             string        `json:"run_id"`
	TurnEpoch         frame.Counter `json:"turn_epoch"`
	RunManifestDigest string        `json:"run_manifest_digest"`
	PolicyDigest      string        `json:"policy_digest"`
	ProviderLaneID    string        `json:"provider_lane_id"`
	LaneCatalogDigest string        `json:"lane_catalog_digest"`
	CredentialRef     string        `json:"credential_ref"`
}

type verificationComparands struct {
	RunManifestDigest string
	PolicyDigest      string
	LaneCatalogDigest string
}

// comparands returns the received bytes without normalization or
// re-derivation; bootstrap verifies loaded artifacts against this exact view.
func (assignment Assign) comparands() verificationComparands {
	return verificationComparands{
		RunManifestDigest: assignment.RunManifestDigest,
		PolicyDigest:      assignment.PolicyDigest,
		LaneCatalogDigest: assignment.LaneCatalogDigest,
	}
}

type FenceResult string

const (
	EpochAllowed FenceResult = "allowed"
	StaleEpoch   FenceResult = "STALE_EPOCH"
	EpochAhead   FenceResult = "EPOCH_AHEAD"
	NotReady     FenceResult = "not_ready"
)

type Session struct {
	control     io.ReadWriteCloser
	sender      *frame.Sender
	decoder     *frame.Decoder
	credentials *credentials.Store
	catalog     *catalog.Catalog
	policy      *policy.Policy
	lane        catalog.Lane
	assignment  Assign

	mu           sync.RWMutex
	ready        bool
	epoch        frame.Counter
	epochChanged chan struct{}
	nextSeq      frame.Counter
	closeOnce    sync.Once
}

func Bootstrap(ctx context.Context, connection io.ReadWriteCloser, config Config) (*Session, error) {
	if connection == nil || config.BuildInfo == "" || len(config.BuildInfo) > 64 || !utf8.ValidString(config.BuildInfo) || config.PID <= 0 {
		return nil, fmt.Errorf("%w: invalid bootstrap configuration", ErrHandshake)
	}
	store, err := credentials.Load(config.Artifacts.CredentialPath)
	if err != nil {
		return nil, err
	}
	catalogRaw, err := os.ReadFile(config.Artifacts.CatalogPath)
	if err != nil {
		return nil, fmt.Errorf("%w: read lane catalog", catalog.ErrInvalidCatalog)
	}
	loadedCatalog, err := catalog.Load(catalogRaw)
	if err != nil {
		return nil, err
	}
	policyRaw, err := os.ReadFile(config.Artifacts.PolicyPath)
	if err != nil || !jcs.IsCanonical(policyRaw) {
		return nil, fmt.Errorf("%w: read or canonicalize policy", policy.ErrPolicyUnavailable)
	}

	session := &Session{
		control: connection,
		sender:  frame.NewSender(connection, frame.SendQueueDepth, frame.SendDeadline),
		decoder: frame.NewDecoder(map[string]frame.TypeSpec{
			"connector_assign": {},
			"epoch_update":     {},
			"ping":             {},
			"shutdown":         {},
		}),
		credentials:  store,
		catalog:      loadedCatalog,
		nextSeq:      1,
		epochChanged: make(chan struct{}),
	}
	failed := true
	defer func() {
		if failed {
			session.Close()
		}
	}()
	if err := session.send(ctx, "hello", "", nil, nil, struct {
		PID       int       `json:"pid"`
		BuildInfo BuildInfo `json:"build_info"`
	}{PID: config.PID, BuildInfo: BuildInfo{Version: config.BuildInfo, Commit: "unknown", BuiltAt: "unknown"}}); err != nil {
		return nil, fmt.Errorf("%w: hello: %v", ErrHandshake, err)
	}

	bound := config.HandshakeBound
	if bound <= 0 {
		bound = HandshakeDeadline
	}
	envelope, err := session.readWithBound(ctx, bound)
	if err != nil {
		return nil, err
	}
	assignment, lane, loadedPolicy, err := verifyAssignment(store, loadedCatalog, policyRaw, envelope)
	if err != nil {
		return nil, err
	}
	session.assignment = assignment
	session.lane = lane
	session.policy = loadedPolicy
	session.epoch = assignment.TurnEpoch
	if err := session.send(ctx, "connector_ready", assignment.RunID, &assignment.TurnEpoch, nil, struct {
		RunID     string        `json:"run_id"`
		TurnEpoch frame.Counter `json:"turn_epoch"`
	}{RunID: assignment.RunID, TurnEpoch: assignment.TurnEpoch}); err != nil {
		return nil, fmt.Errorf("%w: connector_ready: %v", ErrHandshake, err)
	}
	session.ready = true
	failed = false
	return session, nil
}

func verifyAssignment(store *credentials.Store, loadedCatalog *catalog.Catalog, policyRaw []byte, envelope frame.Envelope) (Assign, catalog.Lane, *policy.Policy, error) {
	if envelope.Channel != frame.ChannelControlConnector || envelope.Type != "connector_assign" || envelope.RunID == "" || envelope.TurnEpoch == nil {
		return Assign{}, catalog.Lane{}, nil, fmt.Errorf("%w: invalid connector_assign envelope", ErrHandshake)
	}
	var assignment Assign
	if err := json.Unmarshal(envelope.Body, &assignment); err != nil || assignment.RunID == "" || assignment.ProviderLaneID == "" {
		return Assign{}, catalog.Lane{}, nil, fmt.Errorf("%w: invalid connector_assign body", ErrHandshake)
	}
	comparands := assignment.comparands()
	if assignment.RunID != envelope.RunID || assignment.TurnEpoch != *envelope.TurnEpoch || !validDigest(comparands.RunManifestDigest) ||
		!validDigest(comparands.PolicyDigest) || comparands.PolicyDigest != digestBytes(policyRaw) || !validDigest(comparands.LaneCatalogDigest) || comparands.LaneCatalogDigest != loadedCatalog.Digest {
		return Assign{}, catalog.Lane{}, nil, fmt.Errorf("%w: connector_assign comparand mismatch", ErrHandshake)
	}
	var lane catalog.Lane
	found := false
	for _, candidate := range loadedCatalog.Lanes {
		if candidate.LaneID == assignment.ProviderLaneID {
			lane = candidate
			found = true
			break
		}
	}
	if !found || !credentials.ValidReference(assignment.CredentialRef) || !store.Has(assignment.CredentialRef) {
		return Assign{}, catalog.Lane{}, nil, fmt.Errorf("%w: lane or credential reference mismatch", ErrHandshake)
	}
	loadedPolicy, err := policy.Load(policyRaw, lane)
	if err != nil {
		return Assign{}, catalog.Lane{}, nil, err
	}
	if err := loadedPolicy.VerifyDigest(comparands.PolicyDigest); err != nil {
		return Assign{}, catalog.Lane{}, nil, err
	}
	if err := loadedCatalog.ValidateDeniedHeaders(loadedPolicy.DeniedHeaderNames); err != nil {
		return Assign{}, catalog.Lane{}, nil, err
	}
	return assignment, lane, loadedPolicy, nil
}

func (session *Session) Ready() bool {
	if session == nil {
		return false
	}
	session.mu.RLock()
	defer session.mu.RUnlock()
	return session.ready
}

func (session *Session) Epoch() frame.Counter {
	if session == nil {
		return 0
	}
	session.mu.RLock()
	defer session.mu.RUnlock()
	return session.epoch
}

func (session *Session) RunID() string {
	if session == nil {
		return ""
	}
	return session.assignment.RunID
}

func (session *Session) Lane() catalog.Lane {
	if session == nil {
		return catalog.Lane{}
	}
	return session.lane
}

func (session *Session) Policy() *policy.Policy {
	if session == nil {
		return nil
	}
	return session.policy
}

func (session *Session) Credentials() *credentials.Store {
	if session == nil {
		return nil
	}
	return session.credentials
}

func (session *Session) Assignment() Assign {
	if session == nil {
		return Assign{}
	}
	return session.assignment
}

// SendAttemptResult emits the one-way CTRL-C v2 attempt result. There is no
// reply family for this message.
func (session *Session) SendAttemptResult(ctx context.Context, result outcome.AttemptResultV2) error {
	if session == nil || !session.Ready() {
		return ErrNotReady
	}
	epoch := result.TurnEpoch
	if err := session.send(ctx, "attempt_result", session.assignment.RunID, &epoch, nil, result); err != nil {
		session.failClosed()
		return err
	}
	return nil
}

func (session *Session) FenceDataEpoch(ctx context.Context, presented frame.Counter) (FenceResult, error) {
	if session == nil {
		return NotReady, ErrNotReady
	}
	session.mu.RLock()
	ready, current, changed := session.ready, session.epoch, session.epochChanged
	session.mu.RUnlock()
	if !ready {
		return NotReady, ErrNotReady
	}
	if presented < current {
		return StaleEpoch, nil
	}
	if presented == current {
		return EpochAllowed, nil
	}
	body := struct {
		RunID              string        `json:"run_id"`
		PresentedTurnEpoch frame.Counter `json:"presented_turn_epoch"`
	}{RunID: session.assignment.RunID, PresentedTurnEpoch: presented}
	if err := session.send(ctx, "epoch_query", session.assignment.RunID, &current, nil, body); err != nil {
		session.failClosed()
		return EpochAhead, err
	}
	timer := time.NewTimer(EpochQueryHoldDeadline)
	defer timer.Stop()
	select {
	case <-changed:
	case <-timer.C:
	case <-ctx.Done():
		return EpochAhead, ctx.Err()
	}

	session.mu.RLock()
	ready, current = session.ready, session.epoch
	session.mu.RUnlock()
	if !ready {
		return NotReady, ErrNotReady
	}
	if presented < current {
		return StaleEpoch, nil
	}
	if presented == current {
		return EpochAllowed, nil
	}
	return EpochAhead, nil
}

func (session *Session) HandleControl(ctx context.Context) error {
	if session == nil || !session.Ready() {
		return ErrNotReady
	}
	envelope, err := session.readContext(ctx)
	if err != nil {
		session.failClosed()
		if errors.Is(err, io.EOF) {
			return ErrControlEOF
		}
		return err
	}
	if envelope.Channel != frame.ChannelControlConnector {
		session.failClosed()
		return fmt.Errorf("%w: wrong control channel", ErrHandshake)
	}
	switch envelope.Type {
	case "epoch_update":
		var update struct {
			RunID     string        `json:"run_id"`
			TurnEpoch frame.Counter `json:"turn_epoch"`
		}
		if json.Unmarshal(envelope.Body, &update) != nil || envelope.TurnEpoch == nil || update.RunID != session.assignment.RunID ||
			envelope.RunID != update.RunID || *envelope.TurnEpoch != update.TurnEpoch {
			session.failClosed()
			return fmt.Errorf("%w: malformed epoch_update", ErrHandshake)
		}
		session.mu.Lock()
		if update.TurnEpoch < session.epoch {
			session.ready = false
			close(session.epochChanged)
			session.epochChanged = make(chan struct{})
			session.mu.Unlock()
			return fmt.Errorf("%w: regressed epoch_update", ErrHandshake)
		}
		session.epoch = update.TurnEpoch
		close(session.epochChanged)
		session.epochChanged = make(chan struct{})
		session.mu.Unlock()
		return nil
	case "ping":
		if err := session.send(ctx, "pong", session.assignment.RunID, envelope.TurnEpoch, &envelope.Seq, struct{}{}); err != nil {
			session.failClosed()
			return err
		}
		return nil
	case "shutdown":
		session.failClosed()
		return ErrShutdown
	default:
		session.failClosed()
		return fmt.Errorf("%w: unexpected control type", ErrHandshake)
	}
}

func (session *Session) send(ctx context.Context, kind, runID string, epoch *frame.Counter, replyTo *frame.Counter, body any) error {
	raw, err := json.Marshal(body)
	if err != nil {
		return err
	}
	session.mu.Lock()
	sequence := session.nextSeq
	session.nextSeq++
	session.mu.Unlock()
	return session.sender.SendContext(ctx, frame.Envelope{
		Version: 1, Channel: frame.ChannelControlConnector, Type: kind, Seq: sequence,
		Re: replyTo, RunID: runID, TurnEpoch: epoch, Body: raw,
	})
}

func (session *Session) readWithBound(ctx context.Context, bound time.Duration) (frame.Envelope, error) {
	deadlineContext, cancel := context.WithTimeout(ctx, bound)
	defer cancel()
	return session.readContext(deadlineContext)
}

func (session *Session) readContext(ctx context.Context) (frame.Envelope, error) {
	type result struct {
		envelope frame.Envelope
		err      error
	}
	completed := make(chan result, 1)
	go func() {
		envelope, err := session.decoder.Read(session.control)
		completed <- result{envelope: envelope, err: err}
	}()
	select {
	case got := <-completed:
		return got.envelope, got.err
	case <-ctx.Done():
		_ = session.control.Close()
		return frame.Envelope{}, fmt.Errorf("%w: %v", ErrHandshake, ctx.Err())
	}
}

func (session *Session) failClosed() {
	if session == nil {
		return
	}
	session.mu.Lock()
	session.ready = false
	session.mu.Unlock()
}

func (session *Session) Close() {
	if session == nil {
		return
	}
	session.closeOnce.Do(func() {
		session.failClosed()
		if session.sender != nil {
			session.sender.Close()
		}
		if session.control != nil {
			_ = session.control.Close()
		}
	})
}

func validDigest(value string) bool {
	if len(value) != sha256.Size*2 {
		return false
	}
	decoded, err := hex.DecodeString(value)
	return err == nil && hex.EncodeToString(decoded) == value
}

func digestBytes(value []byte) string {
	sum := sha256.Sum256(value)
	return hex.EncodeToString(sum[:])
}
