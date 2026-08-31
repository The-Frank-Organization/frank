package broker

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/The-Frank-Organization/frank/internal/appipc"
)

const ControlSocketName = "broker-control.sock"

type Server struct {
	token      string
	credential *CredentialSink
	registry   *appipc.Registry
	proposals  *ProposalEngine
	mu         sync.Mutex
	generation uint64
	control    net.Conn
	outcomes   []ControlOutcome
}

type ControlOutcome string

const (
	ControlAdopted            ControlOutcome = "adopted"
	ControlRejectedLock       ControlOutcome = "rejected-lock"
	ControlRejectedToken      ControlOutcome = "rejected-token"
	ControlRejectedGeneration ControlOutcome = "rejected-generation"
)

func NewServer(configHome string, tokenReader io.Reader) (*Server, error) {
	credential, err := LoadCredential(configHome)
	if err != nil {
		return nil, err
	}
	token, err := readControlToken(tokenReader)
	if err != nil {
		return nil, err
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		return nil, err
	}
	return &Server{token: token, credential: credential, registry: registry, proposals: NewProposalEngine()}, nil
}

func (server *Server) Serve(ctx context.Context, runtimeDir string, ready io.Writer) error {
	if server == nil || runtimeDir == "" || ready == nil {
		return errors.New("broker: invalid server configuration")
	}
	socketPath := filepath.Join(runtimeDir, ControlSocketName)
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return err
	}
	defer listener.Close()
	defer os.Remove(socketPath)
	if err := os.Chmod(socketPath, 0o600); err != nil {
		return err
	}
	nonceBytes := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, nonceBytes); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(ready, "BROKER_READY nonce=%s\n", hex.EncodeToString(nonceBytes)); err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = listener.Close()
		server.closeCurrentControl()
	}()
	lockPath := filepath.Join(runtimeDir, "broker-control.lock")
	for {
		connection, acceptErr := listener.Accept()
		if acceptErr != nil {
			if ctx.Err() != nil {
				return nil
			}
			return acceptErr
		}
		go func() {
			if err := server.handleControl(connection, lockPath); err != nil {
				_ = connection.Close()
			}
		}()
	}
}

func (server *Server) handleControl(connection net.Conn, lockPath string) error {
	pid, err := peerPID(connection)
	if err != nil || verifyControlLock(lockPath, pid) != nil {
		server.recordOutcome(ControlRejectedLock)
		_ = writeControlOutcome(connection, ControlRejectedLock)
		return errors.New("broker: rejected control lock")
	}
	handshakeBytes, err := appipc.ReadFrame(connection)
	if err != nil {
		return err
	}
	generation, outcome, err := server.parseHandshake(handshakeBytes)
	if err != nil {
		server.recordOutcome(outcome)
		_ = writeControlOutcome(connection, outcome)
		return err
	}
	old, outcome, err := server.adoptControl(connection, generation)
	if err != nil {
		_ = writeControlOutcome(connection, outcome)
		return err
	}
	if old != nil {
		_ = old.Close()
	}
	if err := writeControlOutcome(connection, ControlAdopted); err != nil {
		server.clearCurrentControl(connection)
		return err
	}
	err = server.serveVerifiedControl(connection)
	server.clearCurrentControl(connection)
	return err
}

func (server *Server) serveVerifiedControl(connection net.Conn) error {
	for {
		payload, err := appipc.ReadFrame(connection)
		if err != nil {
			return err
		}
		envelope, err := server.registry.Decode(payload)
		if err != nil || envelope.Type != "state_proposal" || envelope.Channel != appipc.ChannelBroker {
			return errors.New("broker: invalid control message")
		}
		body, ok := envelope.Body.(*appipc.StateProposalBody)
		if !ok {
			return errors.New("broker: invalid proposal body")
		}
		rawBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		result := server.proposals.Propose(DecodeProposal(rawBody))
		if result == nil {
			return errors.New("broker: uncorrelated proposal fault")
		}
		reply, err := server.registry.Encode(appipc.Envelope{
			V: 1, Channel: appipc.ChannelBroker, Type: "state_proposal_result", Seq: envelope.Seq,
			Re: &envelope.Seq, RunID: envelope.RunID, TurnEpoch: envelope.TurnEpoch, Body: *result,
		})
		if err != nil {
			return err
		}
		if err := appipc.WriteFrame(connection, reply); err != nil {
			return err
		}
		if result.Disposition == appipc.ProposalTransitionStarted && server.proposals.CompleteTransition() {
			installed := appipc.EpochInstalledBody{
				EpochTransitionID: "transition-" + envelope.Seq, GenerationID: body.GenerationID,
				TurnEpoch: body.TurnEpoch, StateSeq: body.StateSeq,
			}
			event, err := server.registry.Encode(appipc.Envelope{
				V: 1, Channel: appipc.ChannelBroker, Type: "epoch_installed", Seq: envelope.Seq,
				RunID: envelope.RunID, TurnEpoch: envelope.TurnEpoch, Body: installed,
			})
			if err != nil {
				return err
			}
			if err := appipc.WriteFrame(connection, event); err != nil {
				return err
			}
		}
	}
}

func (server *Server) acceptHandshake(raw []byte) (uint64, error) {
	generation, outcome, err := server.parseHandshake(raw)
	if err != nil {
		server.recordOutcome(outcome)
		return 0, err
	}
	_, _, err = server.adoptControl(nil, generation)
	return generation, err
}

func (server *Server) parseHandshake(raw []byte) (uint64, ControlOutcome, error) {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var handshake struct {
		Token      string `json:"control_token"`
		Generation string `json:"control_generation"`
	}
	if decoder.Decode(&handshake) != nil || requireEOF(decoder) != nil {
		return 0, ControlRejectedToken, errors.New("broker: rejected control handshake")
	}
	canonical, err := appipc.MarshalJCS(map[string]any{"control_token": handshake.Token, "control_generation": handshake.Generation})
	if err != nil || !bytes.Equal(canonical, raw) {
		return 0, ControlRejectedToken, errors.New("broker: non-canonical control handshake")
	}
	if handshake.Token != server.token {
		return 0, ControlRejectedToken, errors.New("broker: rejected control token")
	}
	generation, err := appipc.ParseCounter(handshake.Generation)
	if err != nil || generation == 0 {
		return 0, ControlRejectedGeneration, errors.New("broker: invalid control generation")
	}
	return generation, "", nil
}

func (server *Server) adoptControl(connection net.Conn, generation uint64) (net.Conn, ControlOutcome, error) {
	server.mu.Lock()
	defer server.mu.Unlock()
	if generation <= server.generation {
		server.outcomes = append(server.outcomes, ControlRejectedGeneration)
		return nil, ControlRejectedGeneration, errors.New("broker: stale control generation")
	}
	old := server.control
	server.generation = generation
	server.control = connection
	server.outcomes = append(server.outcomes, ControlAdopted)
	return old, ControlAdopted, nil
}

func (server *Server) recordOutcome(outcome ControlOutcome) {
	server.mu.Lock()
	server.outcomes = append(server.outcomes, outcome)
	server.mu.Unlock()
}

func (server *Server) ControlOutcomes() []ControlOutcome {
	server.mu.Lock()
	defer server.mu.Unlock()
	return append([]ControlOutcome(nil), server.outcomes...)
}

func (server *Server) clearCurrentControl(connection net.Conn) {
	server.mu.Lock()
	if server.control == connection {
		server.control = nil
	}
	server.mu.Unlock()
	_ = connection.Close()
}

func (server *Server) closeCurrentControl() {
	server.mu.Lock()
	connection := server.control
	server.control = nil
	server.mu.Unlock()
	if connection != nil {
		_ = connection.Close()
	}
}

func writeControlOutcome(connection net.Conn, outcome ControlOutcome) error {
	wire, err := appipc.MarshalJCS(map[string]any{"outcome": string(outcome)})
	if err != nil {
		return err
	}
	return appipc.WriteFrame(connection, wire)
}

func readControlToken(reader io.Reader) (string, error) {
	if reader == nil {
		return "", errors.New("broker: missing control token pipe")
	}
	raw, err := io.ReadAll(io.LimitReader(reader, 4097))
	if err != nil || len(raw) < 2 || len(raw) > 4096 || raw[len(raw)-1] != '\n' || bytes.Count(raw, []byte{'\n'}) != 1 {
		return "", errors.New("broker: invalid control token")
	}
	raw = raw[:len(raw)-1]
	for _, character := range raw {
		if character < 0x21 || character > 0x7e {
			return "", errors.New("broker: invalid control token")
		}
	}
	return string(raw), nil
}
