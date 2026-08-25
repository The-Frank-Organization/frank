package supervisor

import (
	"context"
	"errors"
	"os"
	"syscall"
	"time"
)

type Health struct {
	interval time.Duration
	misses   int
	lastPong time.Time
}

func NewHealth(interval time.Duration, misses int) *Health {
	return &Health{interval: interval, misses: misses}
}

func (health *Health) Pong(at time.Time) { health.lastPong = at }

func (health *Health) Missed(now time.Time) bool {
	return health == nil || health.interval <= 0 || health.misses <= 0 || health.lastPong.IsZero() || !now.Before(health.lastPong.Add(time.Duration(health.misses)*health.interval))
}

type Child interface {
	Shutdown(context.Context) error
	Signal(os.Signal) error
	Wait(context.Context) error
}

func Terminate(ctx context.Context, child Child, grace time.Duration) error {
	if child == nil || grace <= 0 {
		return errors.New("supervisor: invalid termination")
	}
	_ = child.Shutdown(ctx)
	if waitBounded(ctx, child, grace) == nil {
		return nil
	}
	_ = child.Signal(syscall.SIGTERM)
	if waitBounded(ctx, child, grace) == nil {
		return nil
	}
	_ = child.Signal(syscall.SIGKILL)
	return child.Wait(ctx)
}

func waitBounded(ctx context.Context, child Child, grace time.Duration) error {
	waitCtx, cancel := context.WithTimeout(ctx, grace)
	defer cancel()
	return child.Wait(waitCtx)
}
