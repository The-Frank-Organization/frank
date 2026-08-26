package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/jackli/frank/internal/connector/control"
	"github.com/jackli/frank/internal/connector/frame"
)

const (
	childModeEnv = "FRANK_CONNECTOR_TEST_CHILD"
	mainCatalog  = `{"lanes":[{"auth":{"auth_header_name":"x-openai-auth","auth_scheme":"bearer"},"compat_mode":"openai-responses","cost":{"effective_time":"2026-07-17T00:00:00Z","input":1,"output":10},"endpoint":"https://api.openai.com/v1/responses","lane_id":"lane-codex-1","limits":{"context":200000,"max_output":100000},"method":"POST","model_id":"gpt-5","observed_at":"2026-07-17T00:00:00Z","profile_facts":{"endpoint_kind":"coding","region":"us-east"},"provider_id":"openai","reasoning":{"effort_levels":["low","medium","high"],"replay_kind":"opaque_item","supported":true},"serving_profile_id":"codex-default","source":"seeded","tool_use":{"strict_schema":true,"supported":true},"wire":{"max_output_tokens_field":"max_output_tokens","server_retention":false,"streaming":true,"usage_in_streaming":true}}],"schema":"m8.lane_catalog.v1"}`
	mainPolicy   = `{"denied_header_names":["authorization","cookie","proxy-authorization","x-api-key","x-openai-auth"],"egress_class":"provider-request","endpoint_allowlist":["https://api.openai.com/v1/responses"],"pinned_lane":"lane-codex-1","schema":"m3.egress_policy.v1"}`
	mainSecret   = "S14_MAIN_SENTINEL_SECRET"
)

func TestMain(testingMain *testing.M) {
	if os.Getenv(childModeEnv) == "1" {
		main()
		return
	}
	os.Exit(testingMain.Run())
}

func TestParseConfigAcceptsPathsAndInheritedDescriptorsOnly(t *testing.T) {
	config, err := parseConfig([]string{
		"-credential", "/runtime/private/credentials.json",
		"-catalog", "/runtime/private/catalog.json",
		"-policy", "/runtime/private/policy.json",
		"-control-fd", "3", "-data-fd", "4", "-death-fd", "5",
		"-runtime-dir", "/runtime/private", "-build-info", "s14-test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if config.controlFD != 3 || config.dataFD != 4 || config.deathFD != 5 || config.credentialPath == "" || config.catalogPath == "" || config.policyPath == "" || config.runtimeDir == "" || config.buildInfo != "s14-test" {
		t.Fatalf("parsed config = %+v", config)
	}
	for _, args := range [][]string{
		{"-credential", "inline-secret"},
		{"-credential", "/a", "-catalog", "/b", "-policy", "/c", "-control-fd", "2", "-data-fd", "4", "-death-fd", "5", "-runtime-dir", "/r", "-build-info", "b"},
		{"-credential", "/a", "-catalog", "/b", "-policy", "/c", "-control-fd", "3", "-data-fd", "3", "-death-fd", "5", "-runtime-dir", "/r", "-build-info", "b"},
		{"-credential", "/runtime/private/../secret", "-catalog", "/runtime/private/catalog", "-policy", "/runtime/private/policy", "-control-fd", "3", "-data-fd", "4", "-death-fd", "5", "-runtime-dir", "/runtime/private", "-build-info", "b"},
		{"-credential", "/runtime/secret", "-catalog", "/runtime/private/catalog", "-policy", "/runtime/private/policy", "-control-fd", "3", "-data-fd", "4", "-death-fd", "5", "-runtime-dir", "/runtime/private", "-build-info", "b"},
	} {
		if _, err := parseConfig(args); err == nil {
			t.Fatalf("invalid argv accepted: %v", args)
		}
	}
}

func TestValidateRuntimeDirRequiresOwnedPrivateRealDirectory(t *testing.T) {
	private := filepath.Join(t.TempDir(), "private")
	if err := os.Mkdir(private, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := validateRuntimeDir(private); err != nil {
		t.Fatalf("private runtime dir: %v", err)
	}
	open := filepath.Join(t.TempDir(), "open")
	if err := os.Mkdir(open, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := validateRuntimeDir(open); !errors.Is(err, errRuntimeDir) {
		t.Fatalf("open runtime dir error = %v", err)
	}
	link := filepath.Join(t.TempDir(), "link")
	if err := os.Symlink(private, link); err != nil {
		t.Fatal(err)
	}
	if err := validateRuntimeDir(link); !errors.Is(err, errRuntimeDir) {
		t.Fatalf("symlink runtime dir error = %v", err)
	}
}

func TestParentDeathPipeCancelsConnectorContext(t *testing.T) {
	readEnd, writeEnd, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer readEnd.Close()
	ctx, cancel := cancelOnParentDeath(context.Background(), readEnd)
	defer cancel()
	if err := writeEnd.Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
		t.Fatal("connector context remained live after parent death EOF")
	}
}

func TestInheritedDescriptorsAreRemarkedCloseOnExec(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	defer writer.Close()
	fd := int(reader.Fd())
	if err := markCloseOnExec(fd); err != nil {
		t.Fatal(err)
	}
	flags, _, errno := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(syscall.F_GETFD), 0)
	if errno != 0 {
		t.Fatal(errno)
	}
	if flags&syscall.FD_CLOEXEC == 0 {
		t.Fatalf("descriptor flags = %#x, want FD_CLOEXEC", flags)
	}
}

func TestCoreDumpsAreDisabledBeforeCredentialLoad(t *testing.T) {
	var original syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_CORE, &original); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := syscall.Setrlimit(syscall.RLIMIT_CORE, &original); err != nil {
			t.Errorf("restore core limit: %v", err)
		}
	})
	if err := disableCoreDumps(); err != nil {
		t.Fatal(err)
	}
	var got syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_CORE, &got); err != nil {
		t.Fatal(err)
	}
	if got.Cur != 0 {
		t.Fatalf("core soft limit = %d, want 0", got.Cur)
	}
}

func TestSpawnHandshakeAndCleanShutdownExposeNoSecret(t *testing.T) {
	runtimeDir := filepath.Join(t.TempDir(), "runtime")
	if err := os.Mkdir(runtimeDir, 0o700); err != nil {
		t.Fatal(err)
	}
	credentialPath := filepath.Join(runtimeDir, "credentials.json")
	catalogPath := filepath.Join(runtimeDir, "catalog.json")
	policyPath := filepath.Join(runtimeDir, "policy.json")
	writeMainArtifact(t, credentialPath, `{"entries":{"cred-main":{"secret":"`+mainSecret+`"}},"schema":"m8.credentials.v1"}`)
	writeMainArtifact(t, catalogPath, mainCatalog)
	writeMainArtifact(t, policyPath, mainPolicy)

	controlFiles := socketFiles(t, "control")
	dataFiles := socketFiles(t, "data")
	deathChild, deathParent, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer deathParent.Close()
	controlConnection, err := net.FileConn(controlFiles[0])
	if err != nil {
		t.Fatal(err)
	}
	_ = controlFiles[0].Close()
	defer controlConnection.Close()
	defer dataFiles[0].Close()
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	command := exec.Command(executable,
		"-credential", credentialPath, "-catalog", catalogPath, "-policy", policyPath,
		"-control-fd", "3", "-data-fd", "4", "-death-fd", "5", "-runtime-dir", runtimeDir, "-build-info", "s14-test",
	)
	command.ExtraFiles = []*os.File{controlFiles[1], dataFiles[1], deathChild}
	command.Env = append(os.Environ(), childModeEnv+"=1")
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	if err := command.Start(); err != nil {
		t.Fatal(err)
	}
	_ = controlFiles[1].Close()
	_ = dataFiles[1].Close()
	_ = deathChild.Close()

	if err := controlConnection.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		t.Fatal(err)
	}
	decoder := frame.NewDecoder(map[string]frame.TypeSpec{"hello": {}, "connector_ready": {}})
	hello, err := decoder.Read(controlConnection)
	if err != nil || hello.Type != "hello" {
		t.Fatalf("hello = %+v, %v; output=%s", hello, err, output.String())
	}
	catalogRaw, _ := os.ReadFile(catalogPath)
	policyRaw, _ := os.ReadFile(policyPath)
	assignment := control.Assign{
		RunID: "run-main", TurnEpoch: 7, RunManifestDigest: mainDigest([]byte("manifest")),
		PolicyDigest: mainDigest(policyRaw), ProviderLaneID: "lane-codex-1",
		LaneCatalogDigest: mainDigest(catalogRaw), CredentialRef: "cred-main",
	}
	writeMainFrame(t, controlConnection, 1, "connector_assign", assignment, assignment)
	ready, err := decoder.Read(controlConnection)
	if err != nil || ready.Type != "connector_ready" {
		t.Fatalf("ready = %+v, %v; output=%s", ready, err, output.String())
	}
	writeMainFrame(t, controlConnection, 2, "shutdown", struct{}{}, assignment)
	wait := make(chan error, 1)
	go func() { wait <- command.Wait() }()
	select {
	case err := <-wait:
		if err != nil {
			t.Fatalf("clean shutdown error = %v; output=%s", err, output.String())
		}
	case <-time.After(2 * time.Second):
		_ = command.Process.Kill()
		t.Fatal("connector did not exit after shutdown")
	}
	if strings.Contains(output.String(), mainSecret) || output.Len() != 0 {
		t.Fatalf("connector emitted process output: %q", output.String())
	}
}

func socketFiles(t *testing.T, name string) [2]*os.File {
	t.Helper()
	fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		t.Fatal(err)
	}
	return [2]*os.File{
		os.NewFile(uintptr(fds[0]), name+"-parent"),
		os.NewFile(uintptr(fds[1]), name+"-child"),
	}
}

func writeMainArtifact(t *testing.T, path, value string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(value), 0o600); err != nil {
		t.Fatal(err)
	}
}

func writeMainFrame(t *testing.T, writer io.Writer, sequence frame.Counter, kind string, body any, assignment control.Assign) {
	t.Helper()
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	epoch := assignment.TurnEpoch
	encoded, err := frame.Encode(frame.Envelope{
		Version: 1, Channel: frame.ChannelControlConnector, Type: kind, Seq: sequence,
		RunID: assignment.RunID, TurnEpoch: &epoch, Body: raw,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := writer.Write(encoded); err != nil {
		t.Fatal(err)
	}
}

func mainDigest(value []byte) string {
	sum := sha256.Sum256(value)
	return hex.EncodeToString(sum[:])
}
