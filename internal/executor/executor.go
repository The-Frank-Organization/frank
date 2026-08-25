// Package executor hosts suite-class checks behind the observe gate. Its v1
// claim is deliberately limited to provided-surface absence; same-uid ambient
// OS access remains possible without a separately elected sandbox.
package executor

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/jackli/frank/internal/observe"
)

const AmbientResidual = "same-uid ambient filesystem, network, and process access remains possible without an OS sandbox"

const DefaultHardCeiling = 10 * time.Minute

const diagnosticPrefix = "frank-executor-diagnostic-"

type Suite struct {
	SourceDir    string
	Command      string
	Args         []string
	TimeoutClass string
	Timeout      time.Duration
}

type Config struct {
	Context      context.Context
	TempRoot     string
	OutputLimit  int
	Suites       map[string]Suite
	HardCeiling  time.Duration
	OnSoftExpiry ExpiryDisposition
	// StageHook is a deterministic test seam invoked after the run snapshot is
	// staged and before its manifest is hashed.
	StageHook        func()
	GroupGone        func(int) bool
	GroupVerifyBound time.Duration
}

type ExpiryAction = observe.ExpiryAction
type ExpiryRequest = observe.ExpiryRequest
type ExpiryDecision = observe.ExpiryDecision
type ExpiryDisposition = observe.ExpiryDisposition

const (
	ExpiryKill   = observe.ExpiryKill
	ExpiryExtend = observe.ExpiryExtend
)

type Host struct {
	config Config
	mu     sync.Mutex
	runs   map[string]*run
}

type run struct {
	done    chan struct{}
	verdict observe.CheckVerdict
}

type cappedCapture struct {
	mu        sync.Mutex
	buf       bytes.Buffer
	limit     int
	truncated bool
}

func New(config Config) *Host {
	if config.Context == nil {
		config.Context = context.Background()
	}
	if config.OutputLimit <= 0 {
		config.OutputLimit = 64 << 10
	}
	if config.TempRoot == "" {
		config.TempRoot = "/tmp"
	}
	if resolved, err := filepath.EvalSymlinks(config.TempRoot); err == nil {
		config.TempRoot = resolved
	}
	if config.GroupGone == nil {
		config.GroupGone = processGroupGone
	}
	if config.GroupVerifyBound <= 0 {
		config.GroupVerifyBound = 750 * time.Millisecond
	}
	if config.HardCeiling <= 0 {
		config.HardCeiling = DefaultHardCeiling
	}
	config.Suites = cloneSuites(config.Suites)
	return &Host{config: config, runs: map[string]*run{}}
}

func (h *Host) Spawn(entry observe.CheckEntry, selection observe.Selection) observe.CheckVerdict {
	if entry.Class != "suite" || !entry.ExecutorRequired {
		return fault(selection, "executor-class-refused")
	}
	suite, ok := h.config.Suites[selection.Params["target"]]
	if !ok {
		return fault(selection, "executor-target-refused")
	}
	if suite.TimeoutClass != entry.TimeoutClass || suite.TimeoutClass != "suite_bounded" {
		return fault(selection, "executor-timeout-class-mismatch")
	}
	if suite.Timeout <= 0 || suite.Timeout > 120*time.Second {
		return fault(selection, "executor-timeout-policy-refused")
	}
	workdir, err := os.MkdirTemp(h.config.TempRoot, "frank-executor-")
	if err != nil {
		return fault(selection, "executor-workdir-fault")
	}
	if err := stageTree(suite.SourceDir, workdir); err != nil {
		_ = os.RemoveAll(workdir)
		return fault(selection, "executor-stage-fault")
	}
	if h.config.StageHook != nil {
		h.config.StageHook()
	}
	key, err := manifestKey(entry, selection, suite, workdir)
	if err != nil {
		_ = os.RemoveAll(workdir)
		return fault(selection, "executor-manifest-fault")
	}

	h.mu.Lock()
	if existing := h.runs[key]; existing != nil {
		h.mu.Unlock()
		_ = os.RemoveAll(workdir)
		<-existing.done
		return existing.verdict
	}
	current := &run{done: make(chan struct{})}
	h.runs[key] = current
	h.mu.Unlock()

	verdict := h.execute(selection, suite, workdir, key)
	h.mu.Lock()
	current.verdict = verdict
	close(current.done)
	h.mu.Unlock()
	return verdict
}

func (h *Host) execute(selection observe.Selection, suite Suite, workdir, runKey string) observe.CheckVerdict {
	cleanup := true
	defer func() {
		if cleanup {
			_ = os.RemoveAll(workdir)
		}
	}()
	for _, dir := range []string{".tmp", ".cache/go-build", ".cache/gopath"} {
		if err := os.MkdirAll(filepath.Join(workdir, dir), 0o700); err != nil {
			return fault(selection, "executor-workdir-fault")
		}
	}
	if !validCommand(suite.Command) {
		return fault(selection, "executor-command-refused")
	}
	moduleCache, err := goModuleCachePath()
	if err != nil {
		return fault(selection, "executor-module-cache-miss")
	}
	command := filepath.Join(workdir, filepath.Clean(suite.Command))
	cmd := exec.Command(command, append([]string(nil), suite.Args...)...)
	cmd.Dir = workdir
	cmd.WaitDelay = 2 * time.Second
	cmd.Env = []string{
		"PATH=" + os.Getenv("PATH"),
		"HOME=" + workdir,
		"TMPDIR=" + filepath.Join(workdir, ".tmp"),
		"GOCACHE=" + filepath.Join(workdir, ".cache/go-build"),
		"GOMODCACHE=" + moduleCache,
		"GOPATH=" + filepath.Join(workdir, ".cache/gopath"),
		"GOPROXY=off",
		"GOSUMDB=off",
		"GOTOOLCHAIN=local",
		"GOWORK=off",
		"GOFLAGS=-mod=readonly",
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	capture := &cappedCapture{limit: h.config.OutputLimit}
	cmd.Stdout = capture
	cmd.Stderr = capture
	if err := cmd.Start(); err != nil {
		return fault(selection, "executor-start-fault")
	}
	pgid := cmd.Process.Pid
	waited := make(chan error, 1)
	go func() { waited <- cmd.Wait() }()
	startedAt := time.Now()
	timer := time.NewTimer(suite.Timeout)
	defer timer.Stop()
	select {
	case waitErr := <-waited:
		return h.finalizeRun(selection, runKey, pgid, waitErr, false, false, capture, &cleanup)
	case <-timer.C:
		if h.config.OnSoftExpiry == nil {
			_ = syscall.Kill(-pgid, syscall.SIGKILL)
			waitErr := <-waited
			return h.finalizeRun(selection, runKey, pgid, waitErr, true, false, capture, &cleanup)
		}
		hardCeiling := h.config.HardCeiling
		if hardCeiling < suite.Timeout {
			hardCeiling = suite.Timeout
		}
		hardDeadline := startedAt.Add(hardCeiling)
		expiryCtx, cancel := context.WithDeadline(h.config.Context, hardDeadline)
		decision := make(chan ExpiryDecision, 1)
		go func() {
			decision <- h.config.OnSoftExpiry(expiryCtx, ExpiryRequest{
				Selection: selection, SoftExpiry: suite.Timeout, HardCeiling: hardCeiling,
			})
		}()
		select {
		case waitErr := <-waited:
			cancel()
			return h.finalizeRun(selection, runKey, pgid, waitErr, false, false, capture, &cleanup)
		case picked := <-decision:
			cancel()
			select {
			case waitErr := <-waited:
				return h.finalizeRun(selection, runKey, pgid, waitErr, false, false, capture, &cleanup)
			default:
			}
			extended := picked.Action == ExpiryExtend
			timedOut := !extended
			var waitErr error
			if extended {
				remaining := time.Until(hardDeadline)
				if remaining > 0 {
					hardTimer := time.NewTimer(remaining)
					select {
					case waitErr = <-waited:
						if !hardTimer.Stop() {
							<-hardTimer.C
						}
					case <-hardTimer.C:
						timedOut = true
						extended = false
					}
				} else {
					timedOut = true
					extended = false
				}
			}
			if timedOut {
				_ = syscall.Kill(-pgid, syscall.SIGKILL)
				waitErr = <-waited
			}
			return h.finalizeRun(selection, runKey, pgid, waitErr, timedOut, extended, capture, &cleanup)
		case <-expiryCtx.Done():
			cancel()
			select {
			case waitErr := <-waited:
				return h.finalizeRun(selection, runKey, pgid, waitErr, false, false, capture, &cleanup)
			default:
			}
			_ = syscall.Kill(-pgid, syscall.SIGKILL)
			waitErr := <-waited
			return h.finalizeRun(selection, runKey, pgid, waitErr, true, false, capture, &cleanup)
		}
	}
}

func (h *Host) finalizeRun(selection observe.Selection, runKey string, pgid int, waitErr error, timedOut, extended bool, capture *cappedCapture, cleanup *bool) observe.CheckVerdict {
	if !h.config.GroupGone(pgid) {
		_ = syscall.Kill(-pgid, syscall.SIGKILL)
		if !h.waitGroupGone(pgid) {
			*cleanup = false
			return fault(selection, "executor-survivor")
		}
	}
	if timedOut {
		return faultWithTiming(selection, "executor-timeout", "timeout")
	}

	exitGreen := waitErr == nil || errors.Is(waitErr, exec.ErrWaitDelay)
	expectGreen := selection.Params["expect_green"] == "true"
	verdict := observe.CheckVerdict{
		CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
		Outcome: "pass", RungReached: "E2", Predicate: observe.Pass, Timing: "under-timeout",
	}
	if extended {
		verdict.Timing = "extended"
	}
	if exitGreen != expectGreen {
		if err := h.retainDiagnostic(runKey, capture.tail()); err != nil {
			return fault(selection, "executor-workdir-fault")
		}
		verdict.Outcome = "fail"
		verdict.RungReached = "none"
		verdict.Predicate = observe.Fail
		verdict.FailingDetail = "suite-exit-mismatch"
	} else if capture.wasTruncated() {
		verdict.FailingDetail = "output-truncated"
	}
	return verdict
}

func fault(selection observe.Selection, detail string) observe.CheckVerdict {
	return faultWithTiming(selection, detail, "not-completed")
}

func faultWithTiming(selection observe.Selection, detail, timing string) observe.CheckVerdict {
	return observe.CheckVerdict{
		CheckID: selection.CheckID, ClaimRef: selection.ClaimRef,
		Outcome: "unsafe", RungReached: "none", Predicate: observe.Blocked,
		Timing: timing, FailingDetail: detail,
	}
}

func (c *cappedCapture) Write(data []byte) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	original := len(data)
	if c.limit <= 0 {
		c.truncated = true
		return original, nil
	}
	if len(data) >= c.limit {
		if c.buf.Len() > 0 || len(data) > c.limit {
			c.truncated = true
		}
		c.buf.Reset()
		_, _ = c.buf.Write(data[len(data)-c.limit:])
		return original, nil
	}
	if overflow := c.buf.Len() + len(data) - c.limit; overflow > 0 {
		_ = c.buf.Next(overflow)
		c.truncated = true
	}
	_, _ = c.buf.Write(data)
	return original, nil
}

func (c *cappedCapture) tail() []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]byte(nil), c.buf.Bytes()...)
}

func (c *cappedCapture) wasTruncated() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.truncated
}

func (h *Host) retainDiagnostic(runKey string, data []byte) error {
	target := filepath.Join(h.config.TempRoot, diagnosticPrefix+runKey)
	temporary, err := os.CreateTemp(h.config.TempRoot, diagnosticPrefix+"pending-")
	if err != nil {
		return err
	}
	temporaryName := temporary.Name()
	keep := false
	defer func() {
		_ = temporary.Close()
		if !keep {
			_ = os.Remove(temporaryName)
		}
	}()
	if err := temporary.Chmod(0o600); err != nil {
		return err
	}
	if _, err := temporary.Write(data); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Rename(temporaryName, target); err != nil {
		return err
	}
	keep = true
	return nil
}

func goModuleCachePath() (string, error) {
	if path := os.Getenv("GOMODCACHE"); path != "" {
		return resolvedAbsolutePath(path)
	}
	cmd := exec.Command("go", "env", "GOMODCACHE")
	cmd.Env = []string{"PATH=" + os.Getenv("PATH"), "HOME=" + os.Getenv("HOME"), "GOTOOLCHAIN=local"}
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return resolvedAbsolutePath(strings.TrimSpace(string(output)))
}

func resolvedAbsolutePath(path string) (string, error) {
	if path == "" || !filepath.IsAbs(path) {
		return "", errors.New("host module cache path unavailable")
	}
	if resolved, err := filepath.EvalSymlinks(path); err == nil {
		path = resolved
	}
	return path, nil
}

func processGroupGone(pgid int) bool {
	err := syscall.Kill(-pgid, 0)
	return errors.Is(err, syscall.ESRCH)
}

func (h *Host) waitGroupGone(pgid int) bool {
	deadline := time.Now().Add(h.config.GroupVerifyBound)
	for time.Now().Before(deadline) {
		if h.config.GroupGone(pgid) {
			return true
		}
		_ = syscall.Kill(-pgid, syscall.SIGKILL)
		time.Sleep(10 * time.Millisecond)
	}
	return h.config.GroupGone(pgid)
}

func validCommand(command string) bool {
	if command == "" || filepath.IsAbs(command) {
		return false
	}
	clean := filepath.Clean(command)
	return clean != "." && clean != ".." && !strings.HasPrefix(clean, ".."+string(filepath.Separator))
}

func stageTree(source, target string) error {
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if rel == ".git" || strings.HasPrefix(rel, ".git"+string(filepath.Separator)) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		destination := filepath.Join(target, rel)
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		if info.IsDir() {
			return os.MkdirAll(destination, info.Mode().Perm())
		}
		if !info.Mode().IsRegular() {
			return errors.New("non-regular input refused")
		}
		input, err := os.Open(path)
		if err != nil {
			return err
		}
		output, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode().Perm())
		if err != nil {
			_ = input.Close()
			return err
		}
		_, copyErr := io.Copy(output, input)
		inputCloseErr := input.Close()
		closeErr := output.Close()
		if copyErr != nil {
			return copyErr
		}
		if inputCloseErr != nil {
			return inputCloseErr
		}
		return closeErr
	})
}

func manifestKey(entry observe.CheckEntry, selection observe.Selection, suite Suite, stagedDir string) (string, error) {
	h := sha256.New()
	_, _ = io.WriteString(h, entry.ID+"\x00"+selection.CheckID+"\x00"+selection.ClaimRef+"\x00"+selection.CandidateDigest+"\x00")
	keys := make([]string, 0, len(selection.Params))
	for key := range selection.Params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		_, _ = io.WriteString(h, key+"\x00"+selection.Params[key]+"\x00")
	}
	_, _ = io.WriteString(h, suite.Command+"\x00"+strings.Join(suite.Args, "\x00")+"\x00")
	err := filepath.Walk(stagedDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == ".git" && path != stagedDir {
				return filepath.SkipDir
			}
			return nil
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		if !info.Mode().IsRegular() {
			return errors.New("manifest input refused")
		}
		rel, err := filepath.Rel(stagedDir, path)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		sum := sha256.Sum256(data)
		_, _ = io.WriteString(h, rel+"\x00"+hex.EncodeToString(sum[:])+"\x00")
		return nil
	})
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func cloneSuites(in map[string]Suite) map[string]Suite {
	out := make(map[string]Suite, len(in))
	for key, suite := range in {
		suite.Args = append([]string(nil), suite.Args...)
		out[key] = suite
	}
	return out
}
