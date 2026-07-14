package observe

import (
	"errors"
	"io"
	"os"
	"syscall"
	"time"
	_ "unsafe" // required by go:linkname
)

const (
	// readCheckTimeout is the E1 soft-expiry prompt point; it is not a silent
	// auto-disposition boundary.
	readCheckTimeout = 5 * time.Second
	// readFileByteCeiling is a durable fail-closed resource bound; it does not
	// change with the soft-expiry prompt behavior.
	readFileByteCeiling = 8 << 20
	readFileChunkSize   = 32 << 10
)

// ReadFileStage preserves the landed read-file test seam while the worker is
// generalized for other descriptor-rooted filesystem operations.
type ReadFileStage string

const (
	ReadFileStageBefore   ReadFileStage = "before"
	ReadFileStageOpen     ReadFileStage = "open"
	ReadFileStageMetadata ReadFileStage = "metadata"
	ReadFileStageRead     ReadFileStage = "read"
)

type readFileResultKind = fsResultKind

const (
	readFileResultData      = fsResultData
	readFileResultAbsent    = fsResultAbsent
	readFileResultRefused   = fsResultRefused
	readFileResultMachinery = fsResultMachinery
)

type readFileResult = fsResult

func (r *Registry) executeReadFile(selection Selection, laneRef, relative string) readFileResult {
	timeout := r.env.ReadTimeout
	if timeout > readCheckTimeout {
		timeout = readCheckTimeout
	}
	workerTimeout := timeout
	if r.env.OnSoftExpiry != nil && r.env.HardCeiling > workerTimeout {
		workerTimeout = r.env.HardCeiling
	}
	hook := r.readFileFSHook()
	result := r.executeFSWithHookDetails(
		selection,
		laneRef,
		"check-machinery-read-file-breaker-open",
		"check-machinery-read-file-timeout",
		hook,
		func(rootFD int) fsResult {
			return readFileFSOp(rootFD, relative, time.Now().Add(workerTimeout), hook)
		},
	)
	if result.detail == "check-machinery-fs-root-open" || result.detail == "check-machinery-fs-close" {
		result.detail = "check-machinery-read-file"
	}
	return result
}

func (r *Registry) readFileFSHook() func(FSStage) {
	return func(stage FSStage) {
		callFSHook(r.env.FSStageHook, stage)
		switch stage {
		case FSStageRootOpen:
			callReadFileHook(r.env.ReadFileStageHook, ReadFileStageBefore)
			callReadFileHook(r.env.ReadFileStageHook, ReadFileStageOpen)
		case FSStageFileOpen:
			callReadFileHook(r.env.ReadFileStageHook, ReadFileStageOpen)
		case FSStageMetadata:
			callReadFileHook(r.env.ReadFileStageHook, ReadFileStageMetadata)
		case FSStageRead:
			callReadFileHook(r.env.ReadFileStageHook, ReadFileStageRead)
		}
	}
}

func readFileFSOp(rootFD int, relative string, deadline time.Time, hook func(FSStage)) fsResult {
	file, err := openRootedNoFollowAt(rootFD, relative, hook)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return fsResult{kind: fsResultAbsent, detail: "read-file-absent"}
		case errors.Is(err, syscall.ELOOP), isNonRegularOpenError(err):
			return fsResult{kind: fsResultRefused, detail: "not-regular-file"}
		default:
			return fsResult{kind: fsResultMachinery, detail: "check-machinery-read-file"}
		}
	}
	closeFile := func(result fsResult) fsResult {
		callFSHook(hook, FSStageClose)
		if err := file.Close(); err != nil && result.kind == fsResultData {
			return fsResult{kind: fsResultMachinery, detail: "check-machinery-read-file"}
		}
		return result
	}

	callFSHook(hook, FSStageMetadata)
	info, err := file.Stat()
	if err != nil {
		return closeFile(fsResult{kind: fsResultMachinery, detail: "check-machinery-read-file"})
	}
	if !info.Mode().IsRegular() {
		return closeFile(fsResult{kind: fsResultRefused, detail: "not-regular-file"})
	}
	if info.Size() > readFileByteCeiling {
		return closeFile(fsResult{kind: fsResultRefused, detail: "read-size-exceeded"})
	}

	data := make([]byte, 0, minInt64(info.Size(), readFileByteCeiling))
	chunk := make([]byte, readFileChunkSize)
	for {
		if !time.Now().Before(deadline) {
			return closeFile(fsResult{kind: fsResultRefused, detail: "read-deadline-exceeded"})
		}
		callFSHook(hook, FSStageRead)
		if !time.Now().Before(deadline) {
			return closeFile(fsResult{kind: fsResultRefused, detail: "read-deadline-exceeded"})
		}
		remaining := readFileByteCeiling + 1 - len(data)
		if remaining < len(chunk) {
			chunk = chunk[:remaining]
		}
		n, readErr := file.Read(chunk)
		data = append(data, chunk[:n]...)
		if len(data) > readFileByteCeiling {
			return closeFile(fsResult{kind: fsResultRefused, detail: "read-size-exceeded"})
		}
		if !time.Now().Before(deadline) {
			return closeFile(fsResult{kind: fsResultRefused, detail: "read-deadline-exceeded"})
		}
		if readErr == io.EOF {
			return closeFile(fsResult{kind: fsResultData, data: data})
		}
		if readErr != nil {
			return closeFile(fsResult{kind: fsResultMachinery, detail: "check-machinery-read-file"})
		}
	}
}

func callReadFileHook(hook func(ReadFileStage), stage ReadFileStage) {
	if hook != nil {
		hook(stage)
	}
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
// private syscall.openat wrapper on the project's bounded Go 1.22+ Darwin and
// Linux compatibility surface; it is not covered by Go 1 compatibility.
//
//go:linkname syscallOpenat syscall.openat
func syscallOpenat(dirfd int, path string, flags int, mode uint32) (int, error)
