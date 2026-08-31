package testutil

import (
	"context"
	"net"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appipc"
)

func TestFakeWorkerSpeaksTypedFramesOverRealSocketpair(t *testing.T) {
	controllerConn, workerConn := socketpair(t)
	defer controllerConn.Close()
	defer workerConn.Close()

	worker, err := NewFakeWorker(workerConn)
	if err != nil {
		t.Fatalf("NewFakeWorker: %v", err)
	}
	controller, err := NewPeer(controllerConn, appipc.ChannelCtrlW)
	if err != nil {
		t.Fatalf("NewPeer: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := worker.Send(ctx, Outbound{Type: "hello", Body: &appipc.HelloBody{PID: 7, BuildInfo: map[string]string{"version": "test"}}}); err != nil {
		t.Fatalf("worker Send hello: %v", err)
	}
	hello, err := controller.Receive(ctx)
	if err != nil {
		t.Fatalf("controller Receive hello: %v", err)
	}
	if _, ok := hello.Body.(*appipc.HelloBody); !ok {
		t.Fatalf("hello body type = %T", hello.Body)
	}

	if _, err := controller.Send(ctx, Outbound{Type: "assign", Body: &appipc.AssignBody{
		RunID: "run-1", TurnEpoch: "1", ManifestDigest: testDigest,
		GenerationID: "gen-1", BrokerWorkerEndpoint: "/runtime/broker.sock",
	}}); err != nil {
		t.Fatalf("controller Send assign: %v", err)
	}
	assign, err := worker.Receive(ctx)
	if err != nil {
		t.Fatalf("worker Receive assign: %v", err)
	}
	if body, ok := assign.Body.(*appipc.AssignBody); !ok || body.GenerationID != "gen-1" {
		t.Fatalf("assign body = %#v", assign.Body)
	}
}

func TestFakeConnectorCanScriptExpectedReceiveAndMalformedInjection(t *testing.T) {
	controllerConn, connectorConn := socketpair(t)
	defer controllerConn.Close()
	defer connectorConn.Close()

	connector, err := NewFakeConnector(connectorConn)
	if err != nil {
		t.Fatalf("NewFakeConnector: %v", err)
	}
	controller, err := NewPeer(controllerConn, appipc.ChannelCtrlC)
	if err != nil {
		t.Fatalf("NewPeer: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- connector.Run(ctx, []Step{
			{ExpectType: "connector_assign"},
			{Send: &Outbound{Type: "connector_ready", Body: &appipc.ConnectorReadyBody{RunID: "run-1", TurnEpoch: "1"}}},
		})
	}()
	if _, err := controller.Send(ctx, Outbound{Type: "connector_assign", Body: &appipc.ConnectorAssignBody{
		RunID: "run-1", TurnEpoch: "1", RunManifestDigest: testDigest, PolicyDigest: testDigest,
		ProviderLaneID: "lane-1", LaneCatalogDigest: testDigest, CredentialRef: "cred-1",
	}}); err != nil {
		t.Fatalf("controller Send connector_assign: %v", err)
	}
	if _, err := controller.Receive(ctx); err != nil {
		t.Fatalf("controller Receive connector_ready: %v", err)
	}
	if err := <-done; err != nil {
		t.Fatalf("connector Run: %v", err)
	}

	if err := connector.Inject(ctx, []byte(`{"v":1,"chan":"ctrl-c","type":"future","seq":"9","body":{}}`)); err != nil {
		t.Fatalf("connector Inject: %v", err)
	}
	if _, err := controller.Receive(ctx); err == nil {
		t.Fatalf("controller accepted injected unknown message")
	}
}

func socketpair(t *testing.T) (net.Conn, net.Conn) {
	t.Helper()
	fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		t.Fatalf("socketpair: %v", err)
	}
	connections := make([]net.Conn, 0, 2)
	for _, fd := range fds {
		file := os.NewFile(uintptr(fd), "appipc-test-socket")
		connection, err := net.FileConn(file)
		_ = file.Close()
		if err != nil {
			for _, open := range connections {
				_ = open.Close()
			}
			t.Fatalf("net.FileConn: %v", err)
		}
		connections = append(connections, connection)
	}
	return connections[0], connections[1]
}

const testDigest = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
