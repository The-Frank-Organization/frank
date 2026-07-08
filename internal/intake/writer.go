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
	j           *Journal
	requests    chan writeRequest[T]
	completions chan writeCompletion[T]
	next        int
	hashes      map[string]string
	inFlight    map[string]*inFlight[T]
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

type inFlight[T any] struct {
	intakeID string
	replies  []chan T
}

type writeCompletion[T any] struct {
	hash string
	out  T
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
		j:           j,
		requests:    make(chan writeRequest[T], 64),
		completions: make(chan writeCompletion[T], 64),
		next:        1,
		hashes:      map[string]string{},
		inFlight:    map[string]*inFlight[T]{},
	}
	if n, err := j.nextIntakeNumber(); err == nil {
		w.next = n
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
		case done := <-w.completions:
			if flight := w.inFlight[done.hash]; flight != nil {
				for _, reply := range flight.replies {
					reply <- done.out
				}
				delete(w.inFlight, done.hash)
			}
		case req := <-w.requests:
			job, emit, hash, err := w.prepare(req.cmd, req.reply)
			req.ack <- writeAck{intakeID: job.Cmd.IntakeID, err: err}
			if err != nil || !emit {
				continue
			}
			select {
			case out <- job:
				go w.awaitCompletion(ctx, hash, job.ReplyCh)
			case <-ctx.Done():
				return
			}
		}
	}
}

func (w *Writer[T]) prepare(cmd Cmd, reply chan T) (Job[T], bool, string, error) {
	if cmd.ContentHash == "" {
		cmd.ContentHash = hashCmd(cmd)
	}
	if flight := w.inFlight[cmd.ContentHash]; flight != nil {
		flight.replies = append(flight.replies, reply)
		cmd.IntakeID = flight.intakeID
		return Job[T]{Cmd: cmd}, false, cmd.ContentHash, nil
	}
	if existing := w.hashes[cmd.ContentHash]; existing != "" {
		cmd.IntakeID = existing
	} else {
		cmd.IntakeID = fmt.Sprintf("intake-%06d", w.next)
		w.next++
		if err := w.j.appendAssigned(cmd); err != nil {
			return Job[T]{Cmd: cmd, ReplyCh: reply}, false, cmd.ContentHash, err
		}
		w.hashes[cmd.ContentHash] = cmd.IntakeID
	}
	loopReply := make(chan T, 1)
	w.inFlight[cmd.ContentHash] = &inFlight[T]{intakeID: cmd.IntakeID, replies: []chan T{reply}}
	return Job[T]{Cmd: cmd, ReplyCh: loopReply}, true, cmd.ContentHash, nil
}

func (w *Writer[T]) awaitCompletion(ctx context.Context, hash string, reply <-chan T) {
	select {
	case out := <-reply:
		select {
		case w.completions <- writeCompletion[T]{hash: hash, out: out}:
		case <-ctx.Done():
		}
	case <-ctx.Done():
	}
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
