package tools

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"
)

var (
	ErrInvalidArguments     = errors.New("tools: invalid arguments")
	ErrPathOutsideWorkspace = errors.New("tools: path outside workspace")
	ErrAmbiguousEdit        = errors.New("tools: ambiguous edit")
	ErrMalformedPatch       = errors.New("tools: malformed patch")
)

const TruncationMarker = "\n[output truncated]"

type LocalConfig struct {
	WorkspaceRoot  string
	ToolHome       string
	ToolTemp       string
	ShellPath      string
	ApprovedPath   string
	OutputLimit    int
	TerminateGrace time.Duration
}

// Backend is the one uniform seam used by every local tool.
type Backend interface {
	Read(context.Context, string, int64, int64) (string, error)
	Write(context.Context, string, string) error
	Edit(context.Context, string, string, string, bool) error
	ApplyPatch(context.Context, string) error
	Bash(context.Context, string, time.Duration) (string, error)
	BoundOutput(string) string
}

// LocalBackend is the MVP's single host-filesystem/host-shell backend. It does
// not claim to sandbox bash; it only enforces the documented workspace and
// child-launch hygiene boundaries.
type LocalBackend struct {
	root           string
	home           string
	temp           string
	shell          string
	path           string
	user           string
	outputLimit    int
	terminateGrace time.Duration
	environment    map[string]string
}

func NewLocalBackend(config LocalConfig) (*LocalBackend, error) {
	if config.OutputLimit <= 0 {
		return nil, errors.New("tools: output limit must be positive")
	}
	if config.TerminateGrace <= 0 {
		return nil, errors.New("tools: terminate grace must be positive")
	}
	root, err := resolvedDirectory(config.WorkspaceRoot, false)
	if err != nil {
		return nil, fmt.Errorf("tools: workspace root: %w", err)
	}
	home, err := resolvedDedicatedDirectory(config.ToolHome)
	if err != nil {
		return nil, fmt.Errorf("tools: tool home: %w", err)
	}
	temp, err := resolvedDedicatedDirectory(config.ToolTemp)
	if err != nil {
		return nil, fmt.Errorf("tools: tool temp: %w", err)
	}
	if home == root || temp == root || home == temp || pathWithin(root, home) || pathWithin(home, root) || pathWithin(root, temp) || pathWithin(temp, root) {
		return nil, errors.New("tools: workspace, tool home, and tool temp must be distinct")
	}
	shell, err := filepath.EvalSymlinks(config.ShellPath)
	if err != nil || !filepath.IsAbs(shell) {
		return nil, errors.New("tools: shell path must resolve to an absolute file")
	}
	shellInfo, err := os.Stat(shell)
	if err != nil || !shellInfo.Mode().IsRegular() || shellInfo.Mode().Perm()&0o111 == 0 {
		return nil, errors.New("tools: shell path is not an executable regular file")
	}
	if !utf8.ValidString(root) || !utf8.ValidString(home) || !utf8.ValidString(temp) || !utf8.ValidString(shell) || !utf8.ValidString(config.ApprovedPath) {
		return nil, errors.New("tools: environment source is not valid UTF-8")
	}
	if err := validateApprovedPath(config.ApprovedPath, root); err != nil {
		return nil, err
	}
	currentUser, err := user.Current()
	if err != nil || currentUser.Username == "" || !utf8.ValidString(currentUser.Username) {
		return nil, errors.New("tools: process user is unavailable or invalid UTF-8")
	}
	backend := &LocalBackend{
		root:           root,
		home:           home,
		temp:           temp,
		shell:          shell,
		path:           config.ApprovedPath,
		user:           currentUser.Username,
		outputLimit:    config.OutputLimit,
		terminateGrace: config.TerminateGrace,
	}
	backend.environment = map[string]string{
		"PATH":    backend.path,
		"HOME":    backend.home,
		"TMPDIR":  backend.temp,
		"LANG":    "C.UTF-8",
		"LC_ALL":  "C.UTF-8",
		"TZ":      "UTC",
		"TERM":    "dumb",
		"SHELL":   backend.shell,
		"USER":    backend.user,
		"LOGNAME": backend.user,
		"PWD":     backend.root,
	}
	return backend, nil
}

func (backend *LocalBackend) WorkspaceRoot() string { return backend.root }

func (backend *LocalBackend) Environment() map[string]string {
	result := make(map[string]string, len(backend.environment))
	for name, value := range backend.environment {
		result[name] = value
	}
	return result
}

func (backend *LocalBackend) BoundOutput(output string) string {
	return boundOutput([]byte(output), backend.outputLimit)
}

func (backend *LocalBackend) Read(ctx context.Context, path string, offset, limit int64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if offset < 0 || limit < 0 {
		return "", fmt.Errorf("%w: offset and limit must be non-negative", ErrInvalidArguments)
	}
	resolved, err := backend.resolveExisting(path)
	if err != nil {
		return "", err
	}
	file, err := os.Open(resolved)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if _, err := file.Seek(offset, io.SeekStart); err != nil {
		return "", err
	}
	readLimit := int64(backend.outputLimit + 1)
	if limit > 0 && limit < readLimit {
		readLimit = limit
	}
	contents, err := io.ReadAll(io.LimitReader(file, readLimit))
	if err != nil {
		return "", err
	}
	return boundOutput(contents, backend.outputLimit), nil
}

func (backend *LocalBackend) Write(ctx context.Context, path, contents string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	resolved, err := backend.resolveWritable(path)
	if err != nil {
		return err
	}
	return os.WriteFile(resolved, []byte(contents), 0o644)
}

func (backend *LocalBackend) Edit(ctx context.Context, path, oldText, newText string, replaceAll bool) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if oldText == "" {
		return fmt.Errorf("%w: old_string is empty", ErrInvalidArguments)
	}
	resolved, err := backend.resolveExisting(path)
	if err != nil {
		return err
	}
	contents, err := os.ReadFile(resolved)
	if err != nil {
		return err
	}
	count := strings.Count(string(contents), oldText)
	if count == 0 {
		return fmt.Errorf("tools: edit target not found")
	}
	if !replaceAll && count != 1 {
		return fmt.Errorf("%w: old_string occurs %d times", ErrAmbiguousEdit, count)
	}
	maximum := 1
	if replaceAll {
		maximum = -1
	}
	updated := strings.Replace(string(contents), oldText, newText, maximum)
	return writePreservingMode(resolved, []byte(updated))
}

func (backend *LocalBackend) ApplyPatch(ctx context.Context, patch string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	edits, err := backend.preparePatch(patch)
	if err != nil {
		return err
	}
	for _, edit := range edits {
		if err := ctx.Err(); err != nil {
			return err
		}
		switch edit.kind {
		case patchAdd:
			if err := os.WriteFile(edit.path, edit.contents, 0o644); err != nil {
				return err
			}
		case patchUpdate:
			if err := writePreservingMode(edit.path, edit.contents); err != nil {
				return err
			}
		case patchDelete:
			if err := os.Remove(edit.path); err != nil {
				return err
			}
		}
	}
	return nil
}

func (backend *LocalBackend) Bash(ctx context.Context, command string, timeout time.Duration) (string, error) {
	if command == "" {
		return "", fmt.Errorf("%w: command is empty", ErrInvalidArguments)
	}
	if timeout < 0 {
		return "", fmt.Errorf("%w: timeout is negative", ErrInvalidArguments)
	}
	runContext := ctx
	cancel := func() {}
	if timeout > 0 {
		runContext, cancel = context.WithTimeout(ctx, timeout)
	}
	defer cancel()

	cmd := exec.Command(backend.shell, "-c", command)
	cmd.Dir = backend.root
	cmd.Env = backend.environmentList()
	configureProcess(cmd)
	output := &boundedWriter{limit: backend.outputLimit}
	cmd.Stdout = output
	cmd.Stderr = output
	// Internal IPC descriptors are required to be CLOEXEC at creation. This
	// final sweep also closes any accidentally unmarked descriptor at the exec
	// boundary; only stdin/stdout/stderr are intentionally mapped.
	if err := closeOnExecDescriptors(); err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	wait := make(chan error, 1)
	go func() { wait <- cmd.Wait() }()
	select {
	case err := <-wait:
		return output.String(), err
	case <-runContext.Done():
		terminateProcessGroup(cmd.Process.Pid, syscall.SIGTERM)
		timer := time.NewTimer(backend.terminateGrace)
		select {
		case <-wait:
			if !timer.Stop() {
				<-timer.C
			}
		case <-timer.C:
			terminateProcessGroup(cmd.Process.Pid, syscall.SIGKILL)
			<-wait
		}
		return output.String(), runContext.Err()
	}
}

func (backend *LocalBackend) environmentList() []string {
	names := make([]string, 0, len(backend.environment))
	for name := range backend.environment {
		names = append(names, name)
	}
	sort.Strings(names)
	result := make([]string, 0, len(names))
	for _, name := range names {
		result = append(result, name+"="+backend.environment[name])
	}
	return result
}

func (backend *LocalBackend) resolveExisting(relative string) (string, error) {
	joined, err := backend.lexicalPath(relative)
	if err != nil {
		return "", err
	}
	resolved, err := filepath.EvalSymlinks(joined)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			parent, parentErr := filepath.EvalSymlinks(filepath.Dir(joined))
			if parentErr == nil && !pathWithin(backend.root, parent) {
				return "", ErrPathOutsideWorkspace
			}
		}
		return "", err
	}
	if !pathWithin(backend.root, resolved) {
		return "", ErrPathOutsideWorkspace
	}
	return resolved, nil
}

func (backend *LocalBackend) resolveWritable(relative string) (string, error) {
	joined, err := backend.lexicalPath(relative)
	if err != nil {
		return "", err
	}
	if _, err := os.Lstat(joined); err == nil {
		resolved, evalErr := filepath.EvalSymlinks(joined)
		if evalErr != nil || !pathWithin(backend.root, resolved) {
			return "", ErrPathOutsideWorkspace
		}
		return resolved, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	parent, err := filepath.EvalSymlinks(filepath.Dir(joined))
	if err != nil {
		return "", err
	}
	if !pathWithin(backend.root, parent) {
		return "", ErrPathOutsideWorkspace
	}
	return filepath.Join(parent, filepath.Base(joined)), nil
}

func (backend *LocalBackend) lexicalPath(relative string) (string, error) {
	if relative == "" || filepath.IsAbs(relative) {
		return "", ErrPathOutsideWorkspace
	}
	clean := filepath.Clean(relative)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", ErrPathOutsideWorkspace
	}
	joined := filepath.Join(backend.root, clean)
	if !pathWithin(backend.root, joined) {
		return "", ErrPathOutsideWorkspace
	}
	return joined, nil
}

func resolvedDirectory(path string, create bool) (string, error) {
	if path == "" || !filepath.IsAbs(path) {
		return "", errors.New("path must be absolute")
	}
	if create {
		if err := os.MkdirAll(path, 0o700); err != nil {
			return "", err
		}
		if err := os.Chmod(path, 0o700); err != nil {
			return "", err
		}
	}
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(resolved)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", errors.New("not a directory")
	}
	return resolved, nil
}

func resolvedDedicatedDirectory(path string) (string, error) {
	if path == "" || !filepath.IsAbs(path) {
		return "", errors.New("path must be absolute")
	}
	created := false
	if err := os.Mkdir(path, 0o700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			return "", err
		}
	} else {
		created = true
	}
	resolved, err := resolvedDirectory(path, false)
	if err != nil {
		return "", err
	}
	entries, err := os.ReadDir(resolved)
	if err != nil {
		return "", err
	}
	if len(entries) != 0 {
		return "", errors.New("dedicated directory is not empty")
	}
	if info, statErr := os.Stat(resolved); statErr != nil {
		return "", statErr
	} else if info.Mode().Perm() != 0o700 {
		if !created {
			return "", errors.New("dedicated directory mode is not 0700")
		}
		return "", fmt.Errorf("created dedicated directory has mode %o", info.Mode().Perm())
	}
	return resolved, nil
}

func validateApprovedPath(path, workspace string) error {
	if path == "" {
		return errors.New("tools: approved PATH is empty")
	}
	for _, member := range filepath.SplitList(path) {
		if member == "" || !filepath.IsAbs(member) {
			return errors.New("tools: approved PATH members must be absolute")
		}
		resolved, err := filepath.EvalSymlinks(member)
		if err != nil {
			return fmt.Errorf("tools: approved PATH member: %w", err)
		}
		if pathWithin(workspace, resolved) || pathWithin(resolved, workspace) {
			return errors.New("tools: approved PATH overlaps workspace")
		}
	}
	return nil
}

func pathWithin(root, candidate string) bool {
	relative, err := filepath.Rel(root, candidate)
	return err == nil && relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator))
}

func writePreservingMode(path string, contents []byte) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	return os.WriteFile(path, contents, info.Mode().Perm())
}

func boundOutput(contents []byte, limit int) string {
	if len(contents) <= limit {
		return string(contents)
	}
	return string(contents[:limit]) + TruncationMarker
}

type boundedWriter struct {
	mu        sync.Mutex
	buffer    []byte
	limit     int
	truncated bool
}

func (writer *boundedWriter) Write(data []byte) (int, error) {
	writer.mu.Lock()
	defer writer.mu.Unlock()
	remaining := writer.limit - len(writer.buffer)
	if remaining > 0 {
		count := len(data)
		if count > remaining {
			count = remaining
		}
		writer.buffer = append(writer.buffer, data[:count]...)
	}
	if len(data) > remaining {
		writer.truncated = true
	}
	return len(data), nil
}

func (writer *boundedWriter) String() string {
	writer.mu.Lock()
	defer writer.mu.Unlock()
	result := string(writer.buffer)
	if writer.truncated {
		result += TruncationMarker
	}
	return result
}
