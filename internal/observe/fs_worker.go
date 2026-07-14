package observe

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
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
	fsResultDegraded
	fsResultMachinery
)

type fsResult struct {
	kind      fsResultKind
	data      []byte
	detail    string
	timing    string
	count     int
	saturated bool
}

const (
	findRefDepthCeiling       = 32
	findRefFileCeiling        = 10_000
	findRefPerFileByteCeiling = 1 << 20
	findRefTotalByteCeiling   = 256 << 20
	findRefMatchCeiling       = 1_000
	findRefBinaryPrefix       = 8 << 10
)

type findRefLimits struct {
	depth        int
	files        int
	perFileBytes int
	totalBytes   int
	matches      int
}

var defaultFindRefLimits = findRefLimits{
	depth: findRefDepthCeiling, files: findRefFileCeiling,
	perFileBytes: findRefPerFileByteCeiling, totalBytes: findRefTotalByteCeiling,
	matches: findRefMatchCeiling,
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

func (r *Registry) executeFindReferences(selection Selection) fsResult {
	limits := defaultFindRefLimits
	if r.env.findRefLimits != nil {
		limits = *r.env.findRefLimits
	}
	result := r.executeFSWithDetails(
		selection,
		selection.Params["lane_ref"],
		"check-machinery-find-references-breaker-open",
		"check-machinery-find-references-timeout",
		func(rootFD int) fsResult {
			return findReferencesFSOp(rootFD, selection.Params["symbol"], limits, r.fsHook())
		},
	)
	if result.detail == "check-machinery-fs-root-open" || result.detail == "check-machinery-fs-close" {
		result.detail = "check-machinery-find-references-io"
	}
	return result
}

type findRefScan struct {
	symbol    []byte
	limits    findRefLimits
	hook      func(FSStage)
	files     int
	bytes     int
	count     int
	saturated bool
}

func findReferencesFSOp(rootFD int, symbol string, limits findRefLimits, hook func(FSStage)) fsResult {
	scan := &findRefScan{symbol: []byte(symbol), limits: limits, hook: hook}
	if result := scan.directory(rootFD, 0); result.kind != fsResultData {
		return result
	}
	return fsResult{kind: fsResultData, count: scan.count, saturated: scan.saturated}
}

func (s *findRefScan) directory(dirFD, depth int) fsResult {
	dupFD, err := syscall.Dup(dirFD)
	if err != nil {
		return s.machinery("check-machinery-find-references-io")
	}
	dir := os.NewFile(uintptr(dupFD), "find-references-directory")
	callFSHook(s.hook, FSStageDirectoryRead)
	names, err := dir.Readdirnames(-1)
	callFSHook(s.hook, FSStageClose)
	closeErr := dir.Close()
	if err != nil || closeErr != nil {
		return s.machinery("check-machinery-find-references-io")
	}
	sort.Strings(names)
	for _, name := range names {
		if name == "." || name == ".." {
			continue
		}
		callFSHook(s.hook, FSStageFileOpen)
		fd, openErr := syscallOpenat(dirFD, name, os.O_RDONLY|syscall.O_CLOEXEC|syscall.O_NOFOLLOW|syscall.O_NONBLOCK, 0)
		if errors.Is(openErr, syscall.ELOOP) {
			continue
		}
		if openErr != nil {
			return s.machinery("check-machinery-find-references-io")
		}
		entry := os.NewFile(uintptr(fd), "find-references-entry")
		callFSHook(s.hook, FSStageMetadata)
		info, statErr := entry.Stat()
		if statErr != nil {
			callFSHook(s.hook, FSStageClose)
			entry.Close()
			return s.machinery("check-machinery-find-references-io")
		}
		if info.IsDir() {
			if depth >= s.limits.depth {
				callFSHook(s.hook, FSStageClose)
				entry.Close()
				return s.machinery("check-machinery-find-references-depth-ceiling")
			}
			result := s.directory(fd, depth+1)
			callFSHook(s.hook, FSStageClose)
			if closeErr := entry.Close(); closeErr != nil && result.kind == fsResultData {
				return s.machinery("check-machinery-find-references-io")
			}
			if result.kind != fsResultData {
				return result
			}
			continue
		}
		if !info.Mode().IsRegular() {
			callFSHook(s.hook, FSStageClose)
			entry.Close()
			continue
		}
		s.files++
		if s.files > s.limits.files {
			callFSHook(s.hook, FSStageClose)
			entry.Close()
			return s.machinery("check-machinery-find-references-file-ceiling")
		}
		callFSHook(s.hook, FSStageRead)
		data, readErr := io.ReadAll(io.LimitReader(entry, int64(s.limits.perFileBytes)+1))
		callFSHook(s.hook, FSStageClose)
		closeErr := entry.Close()
		if readErr != nil || closeErr != nil {
			return s.machinery("check-machinery-find-references-io")
		}
		prefix := data
		if len(prefix) > findRefBinaryPrefix {
			prefix = prefix[:findRefBinaryPrefix]
		}
		if bytes.IndexByte(prefix, 0) >= 0 {
			continue
		}
		if len(data) > s.limits.perFileBytes || info.Size() > int64(s.limits.perFileBytes) {
			return s.machinery("check-machinery-find-references-per-file-ceiling")
		}
		if s.bytes+len(data) > s.limits.totalBytes {
			return s.machinery("check-machinery-find-references-total-byte-ceiling")
		}
		s.bytes += len(data)
		if !utf8.Valid(data) {
			return fsResult{kind: fsResultDegraded, detail: "scan-domain-incomplete-undecodable"}
		}
		s.addMatches(data)
	}
	return fsResult{kind: fsResultData}
}

func (s *findRefScan) addMatches(data []byte) {
	for offset := 0; offset <= len(data)-len(s.symbol); {
		index := bytes.Index(data[offset:], s.symbol)
		if index < 0 {
			return
		}
		start := offset + index
		end := start + len(s.symbol)
		if (start == 0 || !findRefIdentifierByte(data[start-1])) && (end == len(data) || !findRefIdentifierByte(data[end])) {
			if s.count < s.limits.matches {
				s.count++
			} else {
				s.saturated = true
			}
		}
		offset = start + 1
	}
}

func findRefIdentifierByte(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9' || value == '_' || value == '.'
}

func (s *findRefScan) machinery(detail string) fsResult {
	return fsResult{kind: fsResultMachinery, detail: detail, count: s.count, saturated: s.saturated}
}

func isNonRegularOpenError(err error) bool {
	return errors.Is(err, syscall.ENXIO) || errors.Is(err, syscall.ENODEV) || errors.Is(err, syscall.EOPNOTSUPP)
}
