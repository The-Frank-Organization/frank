package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/worker/executor"
)

var _ executor.Invoker = (*Registry)(nil)

func TestRegistryDispatchesAllFiveLocalTools(t *testing.T) {
	backend := newTestBackend(t, 4096)
	registry := NewRegistry(backend)
	writeFile(t, backend.WorkspaceRoot(), "old.txt", "before\n")

	tests := []struct {
		name string
		args string
	}{
		{"read", `{"path":"old.txt"}`},
		{"write", `{"path":"new.txt","content":"new\n"}`},
		{"edit", `{"path":"old.txt","old_string":"before","new_string":"after","replace_all":false}`},
		{"apply_patch", `{"patch":"*** Begin Patch\n*** Add File: patched.txt\n+patched\n*** End Patch"}`},
		{"bash", `{"command":"printf shell"}`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := registry.Invoke(context.Background(), executor.Invocation{
				Identity:  executor.Identity{CanonicalToolName: test.name},
				Arguments: []byte(test.args),
			})
			if err != nil {
				t.Fatal(err)
			}
			if _, ok := value.(string); !ok {
				t.Fatalf("result type = %T, want string", value)
			}
		})
	}
}

func TestArgumentsRejectUnknownMissingAndTrailingMembers(t *testing.T) {
	registry := NewRegistry(newTestBackend(t, 4096))
	for _, test := range []struct {
		name string
		args string
	}{
		{"unknown", `{"path":"x","surprise":true}`},
		{"missing", `{}`},
		{"trailing", `{"path":"x"} {}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "read"}, Arguments: []byte(test.args)})
			if !errors.Is(err, ErrInvalidArguments) {
				t.Fatalf("error = %v, want invalid arguments", err)
			}
		})
	}
}

func TestFileToolsStayWithinResolvedWorkspace(t *testing.T) {
	backend := newTestBackend(t, 4096)
	registry := NewRegistry(backend)
	outside := t.TempDir()
	writeFile(t, outside, "x", "outside")
	if err := os.Symlink(outside, filepath.Join(backend.WorkspaceRoot(), "escape")); err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		label string
		name  string
		args  string
	}{
		{"read traversal", "read", `{"path":"../outside"}`},
		{"read absolute", "read", fmt.Sprintf(`{"path":%q}`, filepath.Join(outside, "x"))},
		{"write symlink parent", "write", `{"path":"escape/x","content":"bad"}`},
		{"edit symlink", "edit", `{"path":"escape/x","old_string":"a","new_string":"b"}`},
		{"patch traversal", "apply_patch", `{"patch":"*** Begin Patch\n*** Add File: ../bad\n+x\n*** End Patch"}`},
	} {
		t.Run(test.label, func(t *testing.T) {
			_, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: test.name}, Arguments: []byte(test.args)})
			if !errors.Is(err, ErrPathOutsideWorkspace) {
				t.Fatalf("error = %v, want path-outside-workspace", err)
			}
		})
	}
	assertFile(t, outside, "x", "outside")
}

func TestReadAndBashOutputTruncateWithExplicitMarker(t *testing.T) {
	backend := newTestBackend(t, 8)
	writeFile(t, backend.WorkspaceRoot(), "large.txt", "0123456789abcdef")
	registry := NewRegistry(backend)

	readValue, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "read"}, Arguments: []byte(`{"path":"large.txt"}`)})
	if err != nil {
		t.Fatal(err)
	}
	if got := readValue.(string); got != "01234567"+TruncationMarker {
		t.Fatalf("read result = %q", got)
	}
	bashValue, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "bash"}, Arguments: []byte(`{"command":"printf 0123456789abcdef"}`)})
	if err != nil {
		t.Fatal(err)
	}
	if got := bashValue.(string); got != "01234567"+TruncationMarker {
		t.Fatalf("bash result = %q", got)
	}
}

func TestWriteEditAndApplyPatch(t *testing.T) {
	backend := newTestBackend(t, 4096)
	registry := NewRegistry(backend)
	invoke := func(name, args string) (string, error) {
		t.Helper()
		value, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: name}, Arguments: []byte(args)})
		if value == nil {
			return "", err
		}
		return value.(string), err
	}

	if _, err := invoke("write", `{"path":"a.txt","content":"one one\n"}`); err != nil {
		t.Fatal(err)
	}
	if _, err := invoke("edit", `{"path":"a.txt","old_string":"one","new_string":"two","replace_all":false}`); !errors.Is(err, ErrAmbiguousEdit) {
		t.Fatalf("ambiguous edit error = %v", err)
	}
	if _, err := invoke("edit", `{"path":"a.txt","old_string":"one","new_string":"two","replace_all":true}`); err != nil {
		t.Fatal(err)
	}
	patch := `*** Begin Patch
*** Update File: a.txt
@@
-two two
+updated
*** Add File: b.txt
+added
*** Delete File: doomed.txt
*** End Patch`
	writeFile(t, backend.WorkspaceRoot(), "doomed.txt", "gone\n")
	args, _ := json.Marshal(map[string]string{"patch": patch})
	if _, err := invoke("apply_patch", string(args)); err != nil {
		t.Fatal(err)
	}
	assertFile(t, backend.WorkspaceRoot(), "a.txt", "updated\n")
	assertFile(t, backend.WorkspaceRoot(), "b.txt", "added\n")
	if _, err := os.Stat(filepath.Join(backend.WorkspaceRoot(), "doomed.txt")); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("deleted file still exists: %v", err)
	}
}

func TestMalformedPatchHasNoPartialEffect(t *testing.T) {
	backend := newTestBackend(t, 4096)
	registry := NewRegistry(backend)
	writeFile(t, backend.WorkspaceRoot(), "a.txt", "original\n")
	patch := `*** Begin Patch
*** Update File: a.txt
@@
-original
+changed
*** Update File: missing.txt
@@ malformed
*** End Patch`
	args, _ := json.Marshal(map[string]string{"patch": patch})
	_, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "apply_patch"}, Arguments: args})
	if !errors.Is(err, ErrMalformedPatch) {
		t.Fatalf("error = %v, want malformed patch", err)
	}
	assertFile(t, backend.WorkspaceRoot(), "a.txt", "original\n")
}

func TestBashEnvironmentIsConstructedFromExactAllowlist(t *testing.T) {
	const parentSecret = "must-not-reach-child"
	t.Setenv("PROVIDER_API_KEY", parentSecret)
	t.Setenv("FRANK_CONTROL_TOKEN", parentSecret)
	backend := newTestBackend(t, 64*1024)
	registry := NewRegistry(backend)
	value, err := registry.Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "bash"}, Arguments: []byte(`{"command":"env"}`)})
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(value.(string)), "\n")
	sort.Strings(lines)
	wantNames := []string{"HOME", "LANG", "LC_ALL", "LOGNAME", "PATH", "PWD", "SHELL", "TERM", "TMPDIR", "TZ", "USER"}
	gotNames := make([]string, 0, len(lines))
	for _, line := range lines {
		name, _, found := strings.Cut(line, "=")
		if !found {
			t.Fatalf("malformed env line %q", line)
		}
		// POSIX shells may synthesize SHLVL and _ after their own exec. They
		// are not members of the exact environment presented to the shell.
		if name != "SHLVL" && name != "_" {
			gotNames = append(gotNames, name)
		}
		if strings.Contains(line, parentSecret) {
			t.Fatalf("parent secret reached child in %q", line)
		}
	}
	if strings.Join(gotNames, ",") != strings.Join(wantNames, ",") {
		t.Fatalf("environment names = %v, want %v", gotNames, wantNames)
	}
	environment := backend.Environment()
	actualNames := make([]string, 0, len(environment))
	for name := range environment {
		actualNames = append(actualNames, name)
	}
	sort.Strings(actualNames)
	if strings.Join(actualNames, ",") != strings.Join(wantNames, ",") {
		t.Fatalf("presented environment names = %v, want %v", actualNames, wantNames)
	}
	if environment["HOME"] == os.Getenv("HOME") || environment["HOME"] == "" {
		t.Fatalf("tool HOME is ambient or empty: %q", environment["HOME"])
	}
	if environment["LANG"] != "C.UTF-8" || environment["LC_ALL"] != "C.UTF-8" || environment["TZ"] != "UTC" || environment["TERM"] != "dumb" {
		t.Fatalf("fixed environment values drifted: %#v", environment)
	}
	if environment["PWD"] != backend.WorkspaceRoot() {
		t.Fatalf("PWD = %q, want %q", environment["PWD"], backend.WorkspaceRoot())
	}
	for _, directory := range []string{environment["HOME"], environment["TMPDIR"]} {
		info, err := os.Stat(directory)
		if err != nil {
			t.Fatal(err)
		}
		if info.Mode().Perm() != 0o700 {
			t.Fatalf("mode for %s = %o, want 700", directory, info.Mode().Perm())
		}
	}
}

func TestBashDoesNotInheritUnlistedFileDescriptor(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unix descriptor fixture")
	}
	file, err := os.CreateTemp(t.TempDir(), "sentinel")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	if _, err := file.WriteString("fd-sentinel"); err != nil {
		t.Fatal(err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	fd := int(file.Fd())
	syscall.CloseOnExec(fd)
	// Prove the launcher closes even an accidentally non-CLOEXEC descriptor.
	_, _, errno := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), syscall.F_SETFD, 0)
	if errno != 0 {
		t.Fatal(errno)
	}
	backend := newTestBackend(t, 4096)
	command := fmt.Sprintf(`cat /dev/fd/%d 2>/dev/null || true`, fd)
	args, _ := json.Marshal(map[string]string{"command": command})
	value, err := NewRegistry(backend).Invoke(context.Background(), executor.Invocation{Identity: executor.Identity{CanonicalToolName: "bash"}, Arguments: args})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(value.(string), "fd-sentinel") {
		t.Fatalf("sentinel descriptor was inherited: %q", value)
	}
}

func TestBashCancellationKillsWholeProcessGroup(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unix process-group fixture")
	}
	backend := newTestBackend(t, 4096)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pidFile := filepath.Join(backend.WorkspaceRoot(), "child.pid")
	quoted := strings.ReplaceAll(pidFile, "'", `'\''`)
	command := fmt.Sprintf("sleep 30 & child=$!; printf '%%s' \"$child\" > '%s'; wait", quoted)
	args, _ := json.Marshal(map[string]any{"command": command, "timeout_ms": 10_000})
	done := make(chan error, 1)
	go func() {
		_, err := NewRegistry(backend).Invoke(ctx, executor.Invocation{Identity: executor.Identity{CanonicalToolName: "bash"}, Arguments: args})
		done <- err
	}()

	var childPID int
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		contents, err := os.ReadFile(pidFile)
		if err == nil && len(contents) > 0 {
			childPID, err = strconv.Atoi(string(contents))
			if err != nil {
				t.Fatal(err)
			}
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if childPID == 0 {
		t.Fatal("child process did not start")
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("bash did not terminate by deadline")
	}
	for deadline := time.Now().Add(time.Second); time.Now().Before(deadline); {
		err := syscall.Kill(childPID, 0)
		if errors.Is(err, syscall.ESRCH) {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("process-group child %d survived cancellation", childPID)
}

func newTestBackend(t *testing.T, outputLimit int) *LocalBackend {
	t.Helper()
	base := t.TempDir()
	root := filepath.Join(base, "workspace")
	home := filepath.Join(base, "tool-home")
	tmp := filepath.Join(base, "tool-tmp")
	if err := os.Mkdir(root, 0o700); err != nil {
		t.Fatal(err)
	}
	backend, err := NewLocalBackend(LocalConfig{
		WorkspaceRoot:  root,
		ToolHome:       home,
		ToolTemp:       tmp,
		ShellPath:      "/bin/sh",
		ApprovedPath:   "/usr/bin:/bin",
		OutputLimit:    outputLimit,
		TerminateGrace: 100 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	return backend
}

func writeFile(t *testing.T, root, relative, contents string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(root, relative), []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
}

func assertFile(t *testing.T, root, relative, want string) {
	t.Helper()
	contents, err := os.ReadFile(filepath.Join(root, relative))
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != want {
		t.Fatalf("%s = %q, want %q", relative, contents, want)
	}
}
