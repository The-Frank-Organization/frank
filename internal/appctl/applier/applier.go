// Package applier owns the app-control store's single serialized write and
// committed-snapshot query loop.
package applier

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/store"
	"github.com/The-Frank-Organization/frank/internal/appipc"
)

var (
	ErrClosed          = errors.New("app control applier: closed")
	ErrInvalidEvent    = errors.New("app control applier: invalid event")
	ErrCounterOverflow = errors.New("app control applier: state sequence exhausted")
	ErrEmission        = errors.New("app control applier: committed emission failed")
)

type Emission struct {
	Kind  string
	Value any
}

type Result struct {
	Value     any
	StateSeq  string
	Emissions []Emission
	// NoMutation marks a classified request that deliberately made no durable
	// change. Such requests commit no state_seq advance.
	NoMutation bool
}

type Response struct {
	Result Result
	Err    error
}

// Event is a typed state transition. Apply runs inside the event's one store
// transaction; the applier advances state_seq in that same transaction.
type Event interface {
	RunID() string
	Apply(context.Context, *store.Tx) (Result, error)
}

type Query interface {
	Read(context.Context, *store.Snapshot) (any, error)
}

type QueryFunc func(context.Context, *store.Snapshot) (any, error)

func (query QueryFunc) Read(ctx context.Context, snapshot *store.Snapshot) (any, error) {
	return query(ctx, snapshot)
}

type Emitter interface {
	Emit(context.Context, Emission) error
}

type EmitterFunc func(context.Context, Emission) error

func (emit EmitterFunc) Emit(ctx context.Context, emission Emission) error {
	return emit(ctx, emission)
}

type Config struct {
	QueueCapacity int
	Emitter       Emitter
}

type requestKind uint8

const (
	requestApply requestKind = iota + 1
	requestRead
)

type request struct {
	kind     requestKind
	ctx      context.Context
	event    Event
	query    Query
	response chan Response
}

type Host struct {
	store     *store.Store
	emitter   Emitter
	requests  chan request
	stop      chan struct{}
	done      chan struct{}
	closeOnce sync.Once
	timers    sync.WaitGroup
}

func New(db *store.Store, config Config) *Host {
	capacity := config.QueueCapacity
	if capacity <= 0 {
		capacity = 64
	}
	host := &Host{
		store: db, emitter: config.Emitter, requests: make(chan request, capacity),
		stop: make(chan struct{}), done: make(chan struct{}),
	}
	go host.loop()
	return host
}

func (host *Host) Apply(ctx context.Context, event Event) (Result, error) {
	if event == nil || event.RunID() == "" {
		return Result{}, ErrInvalidEvent
	}
	return host.submit(request{kind: requestApply, ctx: ctx, event: event, response: make(chan Response, 1)})
}

func (host *Host) Read(ctx context.Context, query Query) (any, error) {
	if query == nil {
		return nil, errors.New("app control applier: nil query")
	}
	result, err := host.submit(request{kind: requestRead, ctx: ctx, query: query, response: make(chan Response, 1)})
	return result.Value, err
}

// After converts timer expiry into an ordinary queued event. The returned
// channel receives exactly one terminal response and is then closed.
func (host *Host) After(ctx context.Context, delay time.Duration, event Event) <-chan Response {
	response := make(chan Response, 1)
	host.timers.Add(1)
	go func() {
		defer host.timers.Done()
		defer close(response)
		timer := time.NewTimer(delay)
		defer timer.Stop()
		select {
		case <-timer.C:
			result, err := host.Apply(ctx, event)
			response <- Response{Result: result, Err: err}
		case <-ctx.Done():
			response <- Response{Err: ctx.Err()}
		case <-host.stop:
			response <- Response{Err: ErrClosed}
		}
	}()
	return response
}

func (host *Host) Close() error {
	host.closeOnce.Do(func() { close(host.stop) })
	host.timers.Wait()
	<-host.done
	return nil
}

func (host *Host) submit(req request) (Result, error) {
	if req.ctx == nil {
		req.ctx = context.Background()
	}
	select {
	case host.requests <- req:
	case <-req.ctx.Done():
		return Result{}, req.ctx.Err()
	case <-host.stop:
		return Result{}, ErrClosed
	}
	select {
	case response := <-req.response:
		return response.Result, response.Err
	case <-req.ctx.Done():
		return Result{}, req.ctx.Err()
	case <-host.stop:
		return Result{}, ErrClosed
	}
}

func (host *Host) loop() {
	defer close(host.done)
	for {
		select {
		case req := <-host.requests:
			host.handle(req)
		case <-host.stop:
			return
		}
	}
}

func (host *Host) handle(req request) {
	select {
	case <-req.ctx.Done():
		req.response <- Response{Err: req.ctx.Err()}
		return
	default:
	}
	switch req.kind {
	case requestApply:
		host.handleApply(req)
	case requestRead:
		var value any
		err := host.store.Read(req.ctx, func(snapshot *store.Snapshot) error {
			var err error
			value, err = req.query.Read(req.ctx, snapshot)
			return err
		})
		req.response <- Response{Result: Result{Value: value}, Err: err}
	default:
		req.response <- Response{Err: errors.New("app control applier: unknown request")}
	}
}

func (host *Host) handleApply(req request) {
	var result Result
	err := host.store.Update(req.ctx, func(tx *store.Tx) error {
		var err error
		result, err = req.event.Apply(req.ctx, tx)
		if err != nil {
			return err
		}
		if !result.NoMutation {
			result.StateSeq, err = bumpStateSeq(req.ctx, tx, req.event.RunID())
		}
		return err
	})
	if err == nil && host.emitter != nil {
		for _, emission := range result.Emissions {
			if emitErr := host.emitter.Emit(req.ctx, emission); emitErr != nil {
				err = fmt.Errorf("%w: %v", ErrEmission, emitErr)
				break
			}
		}
	}
	req.response <- Response{Result: result, Err: err}
}

func bumpStateSeq(ctx context.Context, tx *store.Tx, runID string) (string, error) {
	var stored string
	if err := tx.QueryRowContext(ctx, `SELECT state_seq FROM epochs WHERE run_id=?`, runID).Scan(&stored); err != nil {
		return "", fmt.Errorf("app control applier: read state_seq for %q: %w", runID, err)
	}
	wire, err := appipc.UnpadCounter(stored)
	if err != nil {
		return "", err
	}
	value, err := appipc.ParseCounter(wire)
	if err != nil {
		return "", err
	}
	if value == math.MaxUint64 {
		return "", ErrCounterOverflow
	}
	nextWire := appipc.FormatCounter(value + 1)
	nextStored, err := appipc.PadCounter(nextWire)
	if err != nil {
		return "", err
	}
	result, err := tx.ExecContext(ctx, `UPDATE epochs SET state_seq=? WHERE run_id=? AND state_seq=?`, nextStored, runID, stored)
	if err != nil {
		return "", err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	if rows != 1 {
		return "", errors.New("app control applier: state_seq compare-and-swap failed")
	}
	return nextWire, nil
}
