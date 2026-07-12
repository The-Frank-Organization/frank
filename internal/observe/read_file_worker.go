package observe

import (
	"errors"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
	_ "unsafe" // required by go:linkname
)

const (
	// readCheckTimeout is the interim silent deadline and sunsets at s10.
	readCheckTimeout = 5 * time.Second
	// readFileByteCeiling is a durable fail-closed resource bound; it does not
	// sunset with the interim silent-deadline behavior.
	readFileByteCeiling = 8 << 20
	readFileChunkSize   = 32 << 10
)

// ReadFileStage names the four injectable filesystem-operation block points
// used by the E1 fixtures. Hooks run only inside the detachable worker.
type ReadFileStage string

const (
	ReadFileStageBefore   ReadFileStage = "before"
	ReadFileStageOpen     ReadFileStage = "open"
	ReadFileStageMetadata ReadFileStage = "metadata"
	ReadFileStageRead     ReadFileStage = "read"
)

type readFileResultKind uint8

const (
	readFileResultData readFileResultKind = iota
	readFileResultAbsent
	readFileResultRefused
	readFileResultMachinery
)

type readFileResult struct {
	kind   readFileResultKind
	data   []byte
	detail string
	timing string
}

func (r *Registry) executeReadFile(laneRef, relative string) readFileResult {
	if !r.beginReadFile(laneRef) {
		return readFileResult{kind: readFileResultMachinery, detail: "check-machinery-read-file-breaker-open"}
	}

	// The serialized caller performs only state checks, launch, and channel/timer
	// coordination. Root traversal/open/metadata/read/close all stay in worker.
	done := make(chan readFileResult, 1)
	root := r.env.Lanes[laneRef]
	hook := r.env.ReadFileStageHook
	timeout := r.env.ReadTimeout
	if timeout > readCheckTimeout {
		timeout = readCheckTimeout
	}
	go func() {
		done <- readFileWorker(root, relative, timeout, hook)
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	select {
	case result := <-done:
		r.finishReadFile(laneRef)
		return result
	case <-timer.C:
		// Set-before-return is load-bearing: a subsequent same-lane call must
		// observe the open breaker before this machinery fault is returned.
		r.tripReadFileBreaker(laneRef)
		return readFileResult{
			kind: readFileResultMachinery, detail: "check-machinery-read-file-timeout", timing: "timeout",
		}
	}
}

func (r *Registry) beginReadFile(laneRef string) bool {
	r.readFileMu.Lock()
	defer r.readFileMu.Unlock()
	if r.readFileLane[laneRef] != readFileLaneIdle {
		return false
	}
	r.readFileLane[laneRef] = readFileLaneActive
	return true
}

func (r *Registry) finishReadFile(laneRef string) {
	r.readFileMu.Lock()
	defer r.readFileMu.Unlock()
	if r.readFileLane[laneRef] == readFileLaneActive {
		r.readFileLane[laneRef] = readFileLaneIdle
	}
}

func (r *Registry) tripReadFileBreaker(laneRef string) {
	r.readFileMu.Lock()
	r.readFileLane[laneRef] = readFileLaneBreakerOpen
	r.readFileMu.Unlock()
}

func readFileWorker(root, relative string, timeout time.Duration, hook func(ReadFileStage)) readFileResult {
	deadline := time.Now().Add(timeout)
	callReadFileHook(hook, ReadFileStageBefore)

	file, err := openRootedNoFollow(root, relative, hook)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return readFileResult{kind: readFileResultAbsent, detail: "read-file-absent"}
		case errors.Is(err, syscall.ELOOP):
			return readFileResult{kind: readFileResultRefused, detail: "not-regular-file"}
		case isNonRegularOpenError(err):
			return readFileResult{kind: readFileResultRefused, detail: "not-regular-file"}
		default:
			return readFileResult{kind: readFileResultMachinery, detail: "check-machinery-read-file"}
		}
	}
	defer file.Close()

	callReadFileHook(hook, ReadFileStageMetadata)
	info, err := file.Stat()
	if err != nil {
		return readFileResult{kind: readFileResultMachinery, detail: "check-machinery-read-file"}
	}
	if !info.Mode().IsRegular() {
		return readFileResult{kind: readFileResultRefused, detail: "not-regular-file"}
	}
	if info.Size() > readFileByteCeiling {
		return readFileResult{kind: readFileResultRefused, detail: "read-size-exceeded"}
	}

	data := make([]byte, 0, minInt64(info.Size(), readFileByteCeiling))
	chunk := make([]byte, readFileChunkSize)
	for {
		if !time.Now().Before(deadline) {
			return readFileResult{kind: readFileResultRefused, detail: "read-deadline-exceeded"}
		}
		callReadFileHook(hook, ReadFileStageRead)
		if !time.Now().Before(deadline) {
			return readFileResult{kind: readFileResultRefused, detail: "read-deadline-exceeded"}
		}
		remaining := readFileByteCeiling + 1 - len(data)
		if remaining < len(chunk) {
			chunk = chunk[:remaining]
		}
		n, readErr := file.Read(chunk)
		data = append(data, chunk[:n]...)
		if len(data) > readFileByteCeiling {
			return readFileResult{kind: readFileResultRefused, detail: "read-size-exceeded"}
		}
		if !time.Now().Before(deadline) {
			return readFileResult{kind: readFileResultRefused, detail: "read-deadline-exceeded"}
		}
		if readErr == io.EOF {
			return readFileResult{kind: readFileResultData, data: data}
		}
		if readErr != nil {
			return readFileResult{kind: readFileResultMachinery, detail: "check-machinery-read-file"}
		}
	}
}

func openRootedNoFollow(root, relative string, hook func(ReadFileStage)) (*os.File, error) {
	callReadFileHook(hook, ReadFileStageOpen)
	rootFile, err := os.OpenFile(root, os.O_RDONLY|syscall.O_CLOEXEC|syscall.O_DIRECTORY, 0)
	if err != nil {
		return nil, err
	}
	current := rootFile
	components := strings.Split(relative, string(os.PathSeparator))
	for i, component := range components {
		flags := os.O_RDONLY | syscall.O_CLOEXEC | syscall.O_NOFOLLOW | syscall.O_NONBLOCK
		if i < len(components)-1 {
			flags |= syscall.O_DIRECTORY
		}
		callReadFileHook(hook, ReadFileStageOpen)
		fd, openErr := syscallOpenat(int(current.Fd()), component, flags, 0)
		if openErr != nil {
			current.Close()
			return nil, openErr
		}
		next := os.NewFile(uintptr(fd), "read-file-descriptor")
		if closeErr := current.Close(); closeErr != nil {
			next.Close()
			return nil, closeErr
		}
		current = next
	}
	return current, nil
}

func callReadFileHook(hook func(ReadFileStage), stage ReadFileStage) {
	if hook != nil {
		hook(stage)
	}
}

func isNonRegularOpenError(err error) bool {
	return errors.Is(err, syscall.ENXIO) || errors.Is(err, syscall.ENODEV) ||
		errors.Is(err, syscall.EOPNOTSUPP)
}

func minInt64(value int64, ceiling int) int {
	if value <= 0 {
		return 0
	}
	if value > int64(ceiling) {
		return ceiling
	}
	return int(value)
}

// syscall.Openat is not exported on darwin. This declaration matches the
// private syscall.openat wrapper on the project's bounded Go 1.22+
// Darwin/Linux compatibility surface; it is not covered by Go 1 compatibility.
//
//go:linkname syscallOpenat syscall.openat
func syscallOpenat(dirfd int, path string, flags int, mode uint32) (int, error)
