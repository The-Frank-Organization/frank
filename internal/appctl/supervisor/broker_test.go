package supervisor

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appctl/brokerclient"
)

func TestBrokerReadyRecordIsExactAndFailClosed(t *testing.T) {
	nonce := strings.Repeat("a", 64)
	valid := "BROKER_READY nonce=" + nonce + "\n"
	if got, err := parseBrokerReady(valid); err != nil || got != nonce || len(valid) != 84 {
		t.Fatalf("valid READY nonce=%q len=%d err=%v", got, len(valid), err)
	}
	for _, malformed := range []string{
		"", strings.TrimSuffix(valid, "\n"), "BROKER_READY nonce=" + strings.Repeat("A", 64) + "\n",
		"BROKER_READY nonce=" + strings.Repeat("a", 63) + "\n", valid + "extra\n",
	} {
		if _, err := parseBrokerReady(malformed); err == nil {
			t.Fatalf("malformed READY accepted: %q", malformed)
		}
	}
}

func TestProductionLaunchRecordsSpawnReadyAndCrashFailureClasses(t *testing.T) {
	fixture := newSupervisorFixture(t, "broker-launch-failures")
	client := brokerclient.New(fixture.applier)
	launch := func(binary, instance string, deadline time.Duration) (*BrokerProcess, error) {
		return LaunchBroker(context.Background(), BrokerLaunch{
			BinaryPath: binary, RuntimeDir: filepath.Join(t.TempDir(), "runtime"), ConfigHome: t.TempDir(),
			RunID: fixture.runID, ControlToken: "token-" + instance, At: time.Now().UnixNano(),
			Client: client, Controller: fixture.controller, InstanceID: instance, ReadyDeadline: deadline,
		})
	}
	if process, err := launch(filepath.Join(t.TempDir(), "absent"), "spawn", 5*time.Second); err == nil || process != nil {
		t.Fatalf("spawn failure process=%#v err=%v", process, err)
	}
	malformed := writeBrokerScript(t, "printf 'not-ready\\n'\n")
	if process, err := launch(malformed, "malformed", 5*time.Second); err == nil || process != nil {
		t.Fatalf("malformed READY process=%#v err=%v", process, err)
	}
	absent := writeBrokerScript(t, "sleep 1\n")
	if process, err := launch(absent, "absent", 20*time.Millisecond); err == nil || process != nil {
		t.Fatalf("absent READY process=%#v err=%v", process, err)
	}
	crash := writeBrokerScript(t, "printf 'BROKER_READY nonce="+strings.Repeat("a", 64)+"\\n'\n")
	process, err := launch(crash, "crash", 5*time.Second)
	if err != nil || process == nil || process.State() != WorkerReady {
		t.Fatalf("READY-before-crash process=%#v err=%v", process, err)
	}
	t.Cleanup(func() { _ = process.stdout.Close() })
	deadline := time.Now().Add(time.Second)
	for fixture.count(`SELECT COUNT(*) FROM pending_app_events WHERE run_id=? AND reported_by='broker-supervisor'`, fixture.runID) != 4 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}
	fixture.assertScalar(`SELECT consecutive_failures FROM runs WHERE run_id=?`, fixture.runID, storeCounter(4))
	fixture.assertScalar(`SELECT turn_epoch FROM epochs WHERE run_id=?`, fixture.runID, storeCounter(1))
}

func writeBrokerScript(t *testing.T, body string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "broker-helper")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body), 0o700); err != nil {
		t.Fatal(err)
	}
	return path
}
