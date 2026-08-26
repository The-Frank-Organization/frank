package testutil

import (
	"context"
	"fmt"
	"math"
	"net"
	"sync"
	"time"

	"github.com/jackli/frank/internal/appipc"
)

type Outbound struct {
	Type      string
	Re        *string
	RunID     *string
	TurnEpoch *string
	Body      any
}

type Step struct {
	Send       *Outbound
	ExpectType string
	Inject     []byte
	Close      bool
}

type Peer struct {
	connection net.Conn
	registry   *appipc.Registry
	channel    appipc.Channel

	sendMu       sync.Mutex
	recvMu       sync.Mutex
	nextSeq      uint64
	seqExhausted bool
}

type FakeWorker struct{ *Peer }
type FakeConnector struct{ *Peer }
type FakeBroker struct{ *Peer }

func NewPeer(connection net.Conn, channel appipc.Channel) (*Peer, error) {
	if connection == nil {
		return nil, fmt.Errorf("testutil: connection is required")
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		return nil, err
	}
	return &Peer{connection: connection, registry: registry, channel: channel}, nil
}

func NewFakeWorker(connection net.Conn) (*FakeWorker, error) {
	peer, err := NewPeer(connection, appipc.ChannelCtrlW)
	if err != nil {
		return nil, err
	}
	return &FakeWorker{Peer: peer}, nil
}

func NewFakeConnector(connection net.Conn) (*FakeConnector, error) {
	peer, err := NewPeer(connection, appipc.ChannelCtrlC)
	if err != nil {
		return nil, err
	}
	peer.nextSeq = 1
	return &FakeConnector{Peer: peer}, nil
}

func NewFakeBroker(connection net.Conn) (*FakeBroker, error) {
	peer, err := NewPeer(connection, appipc.ChannelBroker)
	if err != nil {
		return nil, err
	}
	return &FakeBroker{Peer: peer}, nil
}

func (peer *Peer) Send(ctx context.Context, outbound Outbound) (string, error) {
	peer.sendMu.Lock()
	defer peer.sendMu.Unlock()
	if peer.seqExhausted {
		return "", fmt.Errorf("testutil: sequence exhausted")
	}
	if err := setContextWriteDeadline(peer.connection, ctx); err != nil {
		return "", err
	}
	defer clearWriteDeadline(peer.connection)
	sequence := appipc.FormatCounter(peer.nextSeq)
	payload, err := peer.registry.Encode(appipc.Envelope{
		V:         1,
		Channel:   peer.channel,
		Type:      outbound.Type,
		Seq:       sequence,
		Re:        outbound.Re,
		RunID:     outbound.RunID,
		TurnEpoch: outbound.TurnEpoch,
		Body:      outbound.Body,
	})
	if err != nil {
		return "", err
	}
	if err := appipc.WriteFrame(peer.connection, payload); err != nil {
		return "", err
	}
	if peer.nextSeq == math.MaxUint64 {
		peer.seqExhausted = true
		return sequence, nil
	}
	peer.nextSeq++
	return sequence, nil
}

func (peer *Peer) Receive(ctx context.Context) (appipc.Envelope, error) {
	peer.recvMu.Lock()
	defer peer.recvMu.Unlock()
	if err := setContextReadDeadline(peer.connection, ctx); err != nil {
		return appipc.Envelope{}, err
	}
	defer clearReadDeadline(peer.connection)
	payload, err := appipc.ReadFrame(peer.connection)
	if err != nil {
		return appipc.Envelope{}, err
	}
	envelope, err := peer.registry.Decode(payload)
	if err != nil {
		return appipc.Envelope{}, err
	}
	if envelope.Channel != peer.channel {
		return appipc.Envelope{}, fmt.Errorf("testutil: received %s frame on %s peer", envelope.Channel, peer.channel)
	}
	return envelope, nil
}

func (peer *Peer) Inject(ctx context.Context, payload []byte) error {
	peer.sendMu.Lock()
	defer peer.sendMu.Unlock()
	if err := setContextWriteDeadline(peer.connection, ctx); err != nil {
		return err
	}
	defer clearWriteDeadline(peer.connection)
	return appipc.WriteFrame(peer.connection, payload)
}

func (peer *Peer) Run(ctx context.Context, steps []Step) error {
	for index, step := range steps {
		actions := 0
		if step.Send != nil {
			actions++
		}
		if step.ExpectType != "" {
			actions++
		}
		if step.Inject != nil {
			actions++
		}
		if step.Close {
			actions++
		}
		if actions != 1 {
			return fmt.Errorf("testutil: step %d must select exactly one action", index)
		}
		switch {
		case step.Send != nil:
			if _, err := peer.Send(ctx, *step.Send); err != nil {
				return fmt.Errorf("testutil: step %d send: %w", index, err)
			}
		case step.ExpectType != "":
			envelope, err := peer.Receive(ctx)
			if err != nil {
				return fmt.Errorf("testutil: step %d receive: %w", index, err)
			}
			if envelope.Type != step.ExpectType {
				return fmt.Errorf("testutil: step %d received %q, want %q", index, envelope.Type, step.ExpectType)
			}
		case step.Inject != nil:
			if err := peer.Inject(ctx, step.Inject); err != nil {
				return fmt.Errorf("testutil: step %d inject: %w", index, err)
			}
		case step.Close:
			if err := peer.connection.Close(); err != nil {
				return fmt.Errorf("testutil: step %d close: %w", index, err)
			}
		}
	}
	return nil
}

func setContextReadDeadline(connection net.Conn, ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	deadline, ok := ctx.Deadline()
	if !ok {
		return nil
	}
	return connection.SetReadDeadline(deadline)
}

func setContextWriteDeadline(connection net.Conn, ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	deadline, ok := ctx.Deadline()
	if !ok {
		return nil
	}
	return connection.SetWriteDeadline(deadline)
}

func clearReadDeadline(connection net.Conn) {
	_ = connection.SetReadDeadline(time.Time{})
}

func clearWriteDeadline(connection net.Conn) {
	_ = connection.SetWriteDeadline(time.Time{})
}
