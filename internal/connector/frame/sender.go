package frame

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"
)

const (
	SendQueueDepth = 64
	SendDeadline   = 5 * time.Second
)

var ErrSendDeadline = errors.New("frame: SEND_DEADLINE exceeded")

type sendRequest struct {
	payload []byte
	result  chan error
}

// Sender serializes writes through a bounded queue. Send blocks until the
// frame is written or its deadline expires; it never silently drops a frame.
type Sender struct {
	writer   io.Writer
	queue    chan sendRequest
	deadline time.Duration
	done     chan struct{}
	close    sync.Once
}

func NewSender(writer io.Writer, queueDepth int, deadline time.Duration) *Sender {
	if queueDepth <= 0 {
		queueDepth = SendQueueDepth
	}
	if deadline <= 0 {
		deadline = SendDeadline
	}
	sender := &Sender{
		writer:   writer,
		queue:    make(chan sendRequest, queueDepth),
		deadline: deadline,
		done:     make(chan struct{}),
	}
	go sender.run()
	return sender
}

func (s *Sender) Send(envelope Envelope) error {
	return s.SendContext(context.Background(), envelope)
}

func (s *Sender) SendContext(ctx context.Context, envelope Envelope) error {
	payload, err := Encode(envelope)
	if err != nil {
		return err
	}
	request := sendRequest{payload: payload, result: make(chan error, 1)}
	timer := time.NewTimer(s.deadline)
	defer timer.Stop()

	select {
	case s.queue <- request:
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return ErrSendDeadline
	case <-s.done:
		return io.ErrClosedPipe
	}

	select {
	case err := <-request.result:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return ErrSendDeadline
	case <-s.done:
		return io.ErrClosedPipe
	}
}

func (s *Sender) Close() {
	s.close.Do(func() {
		close(s.queue)
		<-s.done
	})
}

func (s *Sender) run() {
	defer close(s.done)
	for request := range s.queue {
		request.result <- writeAll(s.writer, request.payload)
	}
}

func writeAll(writer io.Writer, payload []byte) error {
	for len(payload) > 0 {
		written, err := writer.Write(payload)
		if err != nil {
			return err
		}
		if written <= 0 || written > len(payload) {
			return io.ErrShortWrite
		}
		payload = payload[written:]
	}
	return nil
}
