package brokerclient

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/jackli/frank/internal/appipc"
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
	registry     *appipc.Registry
	mu           sync.Mutex
	nextSeq      uint64
	seqExhausted bool
	lock         *os.File
}

func (client *Client) Establish(ctx context.Context, request ControlRequest) (*Session, error) {
	if request.RunID == "" || request.RuntimeDir == "" || request.ControlToken == "" {
		return nil, errors.New("brokerclient: invalid control request")
	}
	lock, err := openControlLock(filepath.Join(request.RuntimeDir, "broker-control.lock"))
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
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		_ = connection.Close()
		return nil, err
	}
	closeLock = false
	return &Session{Conn: connection, Generation: generation, registry: registry, lock: lock}, nil
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

func openControlLock(path string) (*os.File, error) {
	lock, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, err
	}
	if err := os.Chmod(path, 0o600); err != nil {
		_ = lock.Close()
		return nil, err
	}
	if err := syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		_ = lock.Close()
		return nil, err
	}
	return lock, nil
}

func unlockControl(lock *os.File) {
	if lock == nil {
		return
	}
	_ = syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
	_ = lock.Close()
}
