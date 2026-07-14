package observe

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// FSStage names injectable filesystem-operation boundaries. Hooks execute only
// inside the detachable worker and exist to prove the serialized caller never
// performs a filesystem syscall.
type FSStage string

const (
	FSStageRootOpen      FSStage = "root-open"
	FSStageDirectoryRead FSStage = "directory-read"
	FSStageFileOpen      FSStage = "file-open"
	FSStageMetadata      FSStage = "metadata"
	FSStageRead          FSStage = "read"
	FSStageClose         FSStage = "close"
)

type fsResultKind uint8

const (
	fsResultData fsResultKind = iota
	fsResultAbsent
	fsResultRefused
	fsResultMachinery
)

type fsResult struct {
	kind   fsResultKind
	data   []byte
	detail string
	timing string
}

type fsLaneState uint8

const (
	fsLaneIdle fsLaneState = iota
	fsLaneActive
	fsLaneBreakerOpen
)

func (r *Registry) executeFS(selection Selection, laneRef string, op func(rootFD int) fsResult) fsResult {
	return r.executeFSWithDetails(selection, laneRef, "check-machinery-fs-breaker-open", "check-machinery-fs-timeout", op)
}

func (r *Registry) executeFSWithDetails(selection Selection, laneRef, breakerDetail, timeoutDetail string, op func(rootFD int) fsResult) fsResult {
	return r.executeFSWithHookDetails(selection, laneRef, breakerDetail, timeoutDetail, r.fsHook(), op)
}

func (r *Registry) executeFSWithHookDetails(selection Selection, laneRef, breakerDetail, timeoutDetail string, hook func(FSStage), op func(rootFD int) fsResult) fsResult {
	if !r.beginFS(laneRef) {
		return fsResult{kind: fsResultMachinery, detail: breakerDetail}
	}
	timeout := r.env.ReadTimeout
	if timeout > readCheckTimeout {
		timeout = readCheckTimeout
	}
	hardCeiling := r.env.HardCeiling
	if hardCeiling < timeout {
		hardCeiling = timeout
	}
	done := make(chan fsResult, 1)
	root := r.env.Lanes[laneRef]
	go func() {
		done <- runFSWorker(root, hook, op)
	}()

	startedAt := time.Now()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case result := <-done:
		r.finishFS(laneRef)
		return result
	case <-timer.C:
		if r.env.OnSoftExpiry != nil {
			return r.resolveFSExpiry(selection, laneRef, done, timeout, hardCeiling, startedAt.Add(hardCeiling), timeoutDetail)
		}
		r.tripFSBreaker(laneRef)
		return fsResult{kind: fsResultMachinery, detail: timeoutDetail, timing: "timeout"}
	}
}

func (r *Registry) resolveFSExpiry(selection Selection, laneRef string, done <-chan fsResult, softExpiry, hardCeiling time.Duration, hardDeadline time.Time, timeoutDetail string) fsResult {
	ctx, cancel := context.WithDeadline(r.env.Context, hardDeadline)
	defer cancel()
	decision := make(chan ExpiryDecision, 1)
	go func() {
		decision <- r.env.OnSoftExpiry(ctx, ExpiryRequest{Selection: selection, SoftExpiry: softExpiry, HardCeiling: hardCeiling})
	}()
	completed := false
	var completedResult fsResult
	for {
		select {
		case completedResult = <-done:
			completed = true
			done = nil
		case picked := <-decision:
			if !completed {
				select {
				case completedResult = <-done:
					completed = true
					done = nil
				default:
				}
			}
			if completed {
				r.finishFS(laneRef)
				return completedResult
			}
			if picked.Action != ExpiryExtend {
				r.tripFSBreaker(laneRef)
				return fsResult{kind: fsResultMachinery, detail: timeoutDetail, timing: "timeout"}
			}
			select {
			case result := <-done:
				r.finishFS(laneRef)
				return result
			case <-ctx.Done():
				r.tripFSBreaker(laneRef)
				return fsResult{kind: fsResultMachinery, detail: timeoutDetail, timing: "timeout"}
			}
		case <-ctx.Done():
			if !completed {
				select {
				case completedResult = <-done:
					completed = true
					done = nil
				default:
				}
			}
			if completed {
				r.finishFS(laneRef)
				return completedResult
			}
			r.tripFSBreaker(laneRef)
			return fsResult{kind: fsResultMachinery, detail: timeoutDetail, timing: "timeout"}
		}
	}
}

func (r *Registry) beginFS(laneRef string) bool {
	r.fsMu.Lock()
	defer r.fsMu.Unlock()
	if r.fsLane[laneRef] != fsLaneIdle {
		return false
	}
	r.fsLane[laneRef] = fsLaneActive
	return true
}

func (r *Registry) finishFS(laneRef string) {
	r.fsMu.Lock()
	defer r.fsMu.Unlock()
	if r.fsLane[laneRef] == fsLaneActive {
		r.fsLane[laneRef] = fsLaneIdle
	}
}

func (r *Registry) tripFSBreaker(laneRef string) {
	r.fsMu.Lock()
	r.fsLane[laneRef] = fsLaneBreakerOpen
	r.fsMu.Unlock()
}

func (r *Registry) fsHook() func(FSStage) {
	return r.env.FSStageHook
}

func runFSWorker(root string, hook func(FSStage), op func(rootFD int) fsResult) fsResult {
	callFSHook(hook, FSStageRootOpen)
	rootFile, err := os.OpenFile(root, os.O_RDONLY|syscall.O_CLOEXEC|syscall.O_DIRECTORY, 0)
	if err != nil {
		return fsResult{kind: fsResultMachinery, detail: "check-machinery-fs-root-open"}
	}
	result := op(int(rootFile.Fd()))
	callFSHook(hook, FSStageClose)
	if err := rootFile.Close(); err != nil && result.kind == fsResultData {
		return fsResult{kind: fsResultMachinery, detail: "check-machinery-fs-close"}
	}
	return result
}

func (r *Registry) executeRootHealth(selection Selection, laneRef string) fsResult {
	return r.executeFS(selection, laneRef, func(rootFD int) fsResult {
		callFSHook(r.fsHook(), FSStageDirectoryRead)
		var buffer [1]byte
		if _, err := syscall.ReadDirent(rootFD, buffer[:]); err != nil {
			return fsResult{kind: fsResultMachinery, detail: "check-machinery-root-observability-directory-read"}
		}
		return fsResult{kind: fsResultData, detail: "root-reachable"}
	})
}

func openRootedNoFollowAt(rootFD int, relative string, hook func(FSStage)) (*os.File, error) {
	if relative == "" || filepath.IsAbs(relative) {
		return nil, syscall.EINVAL
	}
	clean := filepath.Clean(relative)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(os.PathSeparator)) {
		return nil, syscall.EINVAL
	}
	components := strings.Split(clean, string(os.PathSeparator))
	dirFD := rootFD
	var owned *os.File
	for i, component := range components {
		if component == "" || component == "." || component == ".." {
			if owned != nil {
				owned.Close()
			}
			return nil, syscall.EINVAL
		}
		flags := os.O_RDONLY | syscall.O_CLOEXEC | syscall.O_NOFOLLOW | syscall.O_NONBLOCK
		if i < len(components)-1 {
			flags |= syscall.O_DIRECTORY
		}
		callFSHook(hook, FSStageFileOpen)
		fd, err := syscallOpenat(dirFD, component, flags, 0)
		if err != nil {
			if owned != nil {
				owned.Close()
			}
			return nil, err
		}
		next := os.NewFile(uintptr(fd), "rooted-descriptor")
		if owned != nil {
			callFSHook(hook, FSStageClose)
			if err := owned.Close(); err != nil {
				next.Close()
				return nil, err
			}
		}
		owned = next
		dirFD = fd
	}
	return owned, nil
}

func callFSHook(hook func(FSStage), stage FSStage) {
	if hook != nil {
		hook(stage)
	}
}

func isNonRegularOpenError(err error) bool {
	return errors.Is(err, syscall.ENXIO) || errors.Is(err, syscall.ENODEV) || errors.Is(err, syscall.EOPNOTSUPP)
}
