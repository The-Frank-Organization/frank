package intake

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackli/frank/internal/config"
	"github.com/jackli/frank/internal/crashpoint"
	"github.com/jackli/frank/internal/fsio"
)

var ErrWriterNotReady = errors.New("intake writer requires ready token")

type Job[T any] struct {
	Cmd     Cmd
	ReplyCh chan T
}

type Writer[T any] struct {
	j        *Journal
	requests chan writeRequest[T]
	next     int
	hashes   map[string]string
}

type writeRequest[T any] struct {
	cmd   Cmd
	reply chan T
	ack   chan writeAck
}

type writeAck struct {
	intakeID string
	err      error
}

func NewWriter[T any](j *Journal, cfg config.EngineConfig, ready any) (*Writer[T], error) {
	if ready == nil {
		return nil, ErrWriterNotReady
	}
	if cfg.SegmentRotateBytes > 0 {
		j.rotateBytes = cfg.SegmentRotateBytes
	}
	entries, err := j.ReadAll()
	if err != nil {
		return nil, err
	}
	w := &Writer[T]{
		j:        j,
		requests: make(chan writeRequest[T], 64),
		next:     1,
		hashes:   map[string]string{},
	}
	for _, entry := range entries {
		if entry.ContentHash != "" {
			w.hashes[entry.ContentHash] = entry.IntakeID
		}
		if n := intakeNumber(entry.IntakeID); n >= w.next {
			w.next = n + 1
		}
	}
	return w, nil
}

func (w *Writer[T]) Submit(ctx context.Context, cmd Cmd) (<-chan T, string, error) {
	reply := make(chan T, 1)
	ack := make(chan writeAck, 1)
	req := writeRequest[T]{cmd: cmd, reply: reply, ack: ack}
	select {
	case w.requests <- req:
	case <-ctx.Done():
		return nil, "", ctx.Err()
	}
	select {
	case out := <-ack:
		if out.err != nil {
			return nil, "", out.err
		}
		return reply, out.intakeID, nil
	case <-ctx.Done():
		return nil, "", ctx.Err()
	}
}

func (w *Writer[T]) Run(ctx context.Context, out chan<- Job[T]) {
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-w.requests:
			job, err := w.prepare(req.cmd, req.reply)
			req.ack <- writeAck{intakeID: job.Cmd.IntakeID, err: err}
			if err != nil {
				continue
			}
			select {
			case out <- job:
			case <-ctx.Done():
				return
			}
		}
	}
}

func (w *Writer[T]) prepare(cmd Cmd, reply chan T) (Job[T], error) {
	if cmd.ContentHash == "" {
		cmd.ContentHash = hashCmd(cmd)
	}
	if existing := w.hashes[cmd.ContentHash]; existing != "" {
		cmd.IntakeID = existing
		return Job[T]{Cmd: cmd, ReplyCh: reply}, nil
	}
	cmd.IntakeID = fmt.Sprintf("intake-%06d", w.next)
	w.next++
	if err := w.j.appendAssigned(cmd); err != nil {
		return Job[T]{Cmd: cmd, ReplyCh: reply}, err
	}
	w.hashes[cmd.ContentHash] = cmd.IntakeID
	return Job[T]{Cmd: cmd, ReplyCh: reply}, nil
}

func (j *Journal) appendAssigned(cmd Cmd) error {
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	seq, path, err := j.activeSegment()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	if err := fsio.AppendFsync(f, data); err != nil {
		_ = f.Close()
		return err
	}
	crashpoint.Hit("post_intake_fsync")
	if err := f.Close(); err != nil {
		return err
	}
	return j.rotateAfterAppend(seq, path)
}

func intakeNumber(id string) int {
	n, err := strconv.Atoi(strings.TrimPrefix(id, "intake-"))
	if err != nil {
		return 0
	}
	return n
}
