package brokerclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appipc"
)

type DialFunc func(context.Context, string, string) (net.Conn, error)

type ControlRequest struct {
	RunID, RuntimeDir, ControlToken string
	At                              int64
	Dial                            DialFunc
}

type Session struct {
	Conn         net.Conn
	Generation   string
	Outcome      ControlOutcome
	registry     *appipc.Registry
	mu           sync.Mutex
	nextSeq      uint64
	seqExhausted bool
	lock         *controlLock
}

type ControlOutcome string

const (
	ControlAdopted            ControlOutcome = "adopted"
	ControlRejectedLock       ControlOutcome = "rejected-lock"
	ControlRejectedToken      ControlOutcome = "rejected-token"
	ControlRejectedGeneration ControlOutcome = "rejected-generation"
)

type ControlHandshakeError struct {
	Outcome ControlOutcome
}

func (err *ControlHandshakeError) Error() string {
	return fmt.Sprintf("brokerclient: control handshake %s", err.Outcome)
}

var errControlLockAlreadyOpen = errors.New("brokerclient: control lock already open in this process")

type controlLock struct {
	file *os.File
	path string
}

var processControlLocks = struct {
	sync.Mutex
	paths map[string]struct{}
}{paths: make(map[string]struct{})}

func (client *Client) Establish(ctx context.Context, request ControlRequest) (*Session, error) {
	if request.RunID == "" || request.RuntimeDir == "" || request.ControlToken == "" {
		return nil, errors.New("brokerclient: invalid control request")
	}
	lock, err := openControlLock(ctx, filepath.Join(request.RuntimeDir, "broker-control.lock"))
	if err != nil {
		return nil, err
	}
	closeLock := true
	defer func() {
		if closeLock {
			unlockControl(lock)
		}
	}()
	generation, err := client.AdvanceControl(ctx, request.RunID, request.ControlToken, request.At)
	if err != nil {
		return nil, err
	}
	dial := request.Dial
	if dial == nil {
		dial = (&net.Dialer{}).DialContext
	}
	connection, err := dial(ctx, "unix", filepath.Join(request.RuntimeDir, "broker-control.sock"))
	if err != nil {
		return nil, err
	}
	handshake, err := appipc.MarshalJCS(map[string]any{"control_token": request.ControlToken, "control_generation": generation})
	if err != nil {
		_ = connection.Close()
		return nil, err
	}
	if err := appipc.WriteFrame(connection, handshake); err != nil {
		_ = connection.Close()
		return nil, err
	}
	outcome, err := readControlOutcome(ctx, connection)
	if err != nil {
		_ = connection.Close()
		return nil, err
	}
	if outcome != ControlAdopted {
		_ = connection.Close()
		return nil, &ControlHandshakeError{Outcome: outcome}
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		_ = connection.Close()
		return nil, err
	}
	closeLock = false
	return &Session{Conn: connection, Generation: generation, Outcome: outcome, registry: registry, lock: lock}, nil
}

func readControlOutcome(ctx context.Context, connection net.Conn) (ControlOutcome, error) {
	if deadline, ok := ctx.Deadline(); ok {
		if err := connection.SetReadDeadline(deadline); err != nil {
			return "", err
		}
		defer connection.SetReadDeadline(time.Time{})
	}
	raw, err := appipc.ReadFrame(connection)
	if err != nil {
		return "", err
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var reply struct {
		Outcome ControlOutcome `json:"outcome"`
	}
	if decoder.Decode(&reply) != nil || requireJSONEOF(decoder) != nil {
		return "", errors.New("brokerclient: invalid control handshake reply")
	}
	canonical, err := appipc.MarshalJCS(map[string]any{"outcome": string(reply.Outcome)})
	if err != nil || !bytes.Equal(canonical, raw) {
		return "", errors.New("brokerclient: non-canonical control handshake reply")
	}
	switch reply.Outcome {
	case ControlAdopted, ControlRejectedLock, ControlRejectedToken, ControlRejectedGeneration:
		return reply.Outcome, nil
	default:
		return "", errors.New("brokerclient: unknown control handshake outcome")
	}
}

func requireJSONEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		return errors.New("brokerclient: trailing control handshake reply")
	}
	return nil
}

// Propose publishes one durable tuple and applies the total proposal-result
// fold. A transition response waits for its matching install event until the
// bounded proposal deadline, then returns Repropose.
func (session *Session) Propose(ctx context.Context, correlation string, tuple appipc.EpochStateBody) (FoldResult, error) {
	if session == nil || session.Conn == nil || correlation == "" {
		return FoldResult{}, errors.New("brokerclient: invalid proposal session")
	}
	session.mu.Lock()
	defer session.mu.Unlock()
	if session.seqExhausted {
		return FoldResult{}, errors.New("brokerclient: proposal sequence exhausted")
	}
	if session.registry == nil {
		registry, err := appipc.NewProtocolRegistry()
		if err != nil {
			return FoldResult{}, err
		}
		session.registry = registry
	}
	gate := NewAssignGate()
	if err := gate.Propose(correlation, tuple); err != nil {
		return FoldResult{}, err
	}
	sequence := appipc.FormatCounter(session.nextSeq)
	runID, epoch := tuple.RunID, tuple.TurnEpoch
	payload, err := session.registry.Encode(appipc.Envelope{
		V: 1, Channel: appipc.ChannelBroker, Type: "state_proposal", Seq: sequence,
		RunID: &runID, TurnEpoch: &epoch,
		Body: appipc.StateProposalBody{
			ProposalCorrelation: correlation, RunID: tuple.RunID, GenerationID: tuple.GenerationID,
			TurnEpoch: tuple.TurnEpoch, LeaseState: tuple.LeaseState, StateSeq: tuple.StateSeq,
		},
	})
	if err != nil {
		return FoldResult{}, err
	}
	if err := appipc.WriteFrame(session.Conn, payload); err != nil {
		return FoldResult{}, err
	}
	if session.nextSeq == math.MaxUint64 {
		session.seqExhausted = true
	} else {
		session.nextSeq++
	}
	result, err := session.receiveProposalResult(ctx, gate, correlation, tuple, sequence)
	if err != nil {
		return FoldResult{}, err
	}
	return result, nil
}

func (session *Session) receiveProposalResult(ctx context.Context, gate *AssignGate, correlation string, tuple appipc.EpochStateBody, sequence string) (FoldResult, error) {
	deadline := time.Now().Add(appipc.ProposalResultDeadline)
	if contextDeadline, ok := ctx.Deadline(); ok && contextDeadline.Before(deadline) {
		deadline = contextDeadline
	}
	if err := session.Conn.SetReadDeadline(deadline); err != nil {
		return FoldResult{}, err
	}
	defer session.Conn.SetReadDeadline(time.Time{})
	for {
		payload, err := appipc.ReadFrame(session.Conn)
		if err != nil {
			var networkError net.Error
			if errors.As(err, &networkError) && networkError.Timeout() {
				return gate.Deadline(tuple.RunID), nil
			}
			return FoldResult{}, err
		}
		envelope, err := session.registry.Decode(payload)
		if err != nil {
			return FoldResult{}, err
		}
		switch envelope.Type {
		case "state_proposal_result":
			if envelope.Re == nil || *envelope.Re != sequence {
				return FoldResult{}, errors.New("brokerclient: proposal reply correlation mismatch")
			}
			body, ok := envelope.Body.(*appipc.StateProposalResultBody)
			if !ok {
				return FoldResult{}, errors.New("brokerclient: invalid proposal result body")
			}
			fold := gate.Fold(correlation, *body)
			if fold.Action != AwaitEventOrDeadline {
				return fold, nil
			}
		case "epoch_installed":
			body, ok := envelope.Body.(*appipc.EpochInstalledBody)
			if !ok || !gate.InstallEvent(tuple.RunID, *body) {
				return FoldResult{Action: InvariantFault, Loud: true}, fmt.Errorf("brokerclient: epoch install mismatch")
			}
			return FoldResult{Action: OpenAssign}, nil
		default:
			return FoldResult{Action: InvariantFault, Loud: true}, fmt.Errorf("brokerclient: unexpected proposal response %q", envelope.Type)
		}
	}
}

func (session *Session) Close() error {
	if session == nil {
		return nil
	}
	var connectionError error
	if session.Conn != nil {
		connectionError = session.Conn.Close()
		session.Conn = nil
	}
	unlockControl(session.lock)
	session.lock = nil
	return connectionError
}

func openControlLock(ctx context.Context, path string) (*controlLock, error) {
	canonicalPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	canonicalPath = filepath.Clean(canonicalPath)
	processControlLocks.Lock()
	if _, exists := processControlLocks.paths[canonicalPath]; exists {
		processControlLocks.Unlock()
		return nil, errControlLockAlreadyOpen
	}
	processControlLocks.paths[canonicalPath] = struct{}{}
	processControlLocks.Unlock()
	registered := true
	defer func() {
		if registered {
			forgetControlLock(canonicalPath)
		}
	}()

	file, err := os.OpenFile(canonicalPath, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, err
	}
	if err := file.Chmod(0o600); err != nil {
		_ = file.Close()
		return nil, err
	}
	record := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: 0, Len: 0}
	for {
		err = syscall.FcntlFlock(file.Fd(), syscall.F_SETLK, &record)
		if err == nil {
			registered = false
			return &controlLock{file: file, path: canonicalPath}, nil
		}
		if !errors.Is(err, syscall.EACCES) && !errors.Is(err, syscall.EAGAIN) {
			_ = file.Close()
			return nil, err
		}
		timer := time.NewTimer(10 * time.Millisecond)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			_ = file.Close()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}

func unlockControl(lock *controlLock) {
	if lock == nil {
		return
	}
	record := syscall.Flock_t{Type: syscall.F_UNLCK, Whence: 0, Start: 0, Len: 0}
	_ = syscall.FcntlFlock(lock.file.Fd(), syscall.F_SETLK, &record)
	_ = lock.file.Close()
	forgetControlLock(lock.path)
}

func forgetControlLock(path string) {
	processControlLocks.Lock()
	delete(processControlLocks.paths, path)
	processControlLocks.Unlock()
}
