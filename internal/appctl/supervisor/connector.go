package supervisor

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"

	"github.com/The-Frank-Organization/frank/internal/appipc"
)

const ConnectorReadyDeadline = 10 * time.Second

type ConnectorLaunch struct {
	BinaryPath     string
	RuntimeDir     string
	CredentialPath string
	CatalogPath    string
	PolicyPath     string
	BuildInfo      string
	Assignment     appipc.ConnectorAssignBody
	Env            []string
	ReadyDeadline  time.Duration
}

type ConnectorProcess struct {
	command *exec.Cmd
	control net.Conn
	data    net.Conn
	death   *os.File
	machine *Machine
}

func (process *ConnectorProcess) State() WorkerState {
	if process == nil || process.machine == nil {
		return WorkerFailed
	}
	return process.machine.State()
}

func LaunchConnector(ctx context.Context, launch ConnectorLaunch) (*ConnectorProcess, error) {
	if launch.BinaryPath == "" || launch.RuntimeDir == "" || launch.CredentialPath == "" || launch.CatalogPath == "" || launch.PolicyPath == "" || launch.BuildInfo == "" {
		return nil, errors.New("supervisor: incomplete connector launch")
	}
	if err := PrepareRuntimeDir(launch.RuntimeDir); err != nil {
		return nil, err
	}
	controlParent, controlChild, err := newInheritedSocket("connector-control")
	if err != nil {
		return nil, err
	}
	defer func() {
		if controlChild != nil {
			_ = controlChild.Close()
		}
	}()
	dataParent, dataChild, err := newInheritedSocket("connector-data")
	if err != nil {
		_ = controlParent.Close()
		return nil, err
	}
	defer func() {
		if dataChild != nil {
			_ = dataChild.Close()
		}
	}()
	death, err := NewDeathPipe()
	if err != nil {
		_ = controlParent.Close()
		_ = dataParent.Close()
		return nil, err
	}
	defer func() {
		if death.Child != nil {
			_ = death.Child.Close()
		}
	}()
	machine := NewMachine(WorkerAllocated)
	if err := machine.Transition(WorkerSpawning); err != nil {
		return nil, err
	}
	command, err := BuildCommand(ProcessSpec{
		Path: launch.BinaryPath,
		Args: []string{
			"-credential", launch.CredentialPath,
			"-catalog", launch.CatalogPath,
			"-policy", launch.PolicyPath,
			"-control-fd", strconv.Itoa(3),
			"-data-fd", strconv.Itoa(4),
			"-death-fd", strconv.Itoa(5),
			"-runtime-dir", launch.RuntimeDir,
			"-build-info", launch.BuildInfo,
		},
		Dir: launch.RuntimeDir, Env: append([]string(nil), launch.Env...),
		ExtraFiles: []*os.File{controlChild, dataChild, death.Child},
	})
	if err != nil {
		_ = controlParent.Close()
		_ = dataParent.Close()
		_ = death.Parent.Close()
		return nil, err
	}
	if err := command.Start(); err != nil {
		_ = machine.Transition(WorkerFailed)
		_ = controlParent.Close()
		_ = dataParent.Close()
		_ = death.Parent.Close()
		return nil, fmt.Errorf("supervisor: start connector: %w", err)
	}
	_ = controlChild.Close()
	controlChild = nil
	_ = dataChild.Close()
	dataChild = nil
	_ = death.Child.Close()
	death.Child = nil
	process := &ConnectorProcess{command: command, control: controlParent, data: dataParent, death: death.Parent, machine: machine}
	deadline := launch.ReadyDeadline
	if deadline <= 0 {
		deadline = ConnectorReadyDeadline
	}
	if err := controlParent.SetDeadline(time.Now().Add(deadline)); err != nil {
		process.fail()
		return nil, err
	}
	registry, err := appipc.NewProtocolRegistry()
	if err != nil {
		process.fail()
		return nil, err
	}
	helloBytes, err := appipc.ReadFrame(controlParent)
	if err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: connector hello: %w", err)
	}
	hello, err := registry.Decode(helloBytes)
	if err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: decode connector hello: %w", err)
	}
	if hello.Type != "hello" || hello.Channel != appipc.ChannelCtrlC {
		process.fail()
		return nil, errors.New("supervisor: invalid connector hello")
	}
	runID, epoch := launch.Assignment.RunID, launch.Assignment.TurnEpoch
	assignBytes, err := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "connector_assign", Seq: "1", RunID: &runID, TurnEpoch: &epoch, Body: launch.Assignment})
	if err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: connector assign: %w", err)
	}
	if err := appipc.WriteFrame(controlParent, assignBytes); err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: connector assign: %w", err)
	}
	readyBytes, err := appipc.ReadFrame(controlParent)
	if err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: connector ready: %w", err)
	}
	ready, err := registry.Decode(readyBytes)
	if err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: decode connector ready: %w", err)
	}
	body, ok := ready.Body.(*appipc.ConnectorReadyBody)
	if ready.Type != "connector_ready" || ready.Channel != appipc.ChannelCtrlC || !ok || body.RunID != runID || body.TurnEpoch != epoch {
		process.fail()
		return nil, errors.New("supervisor: invalid connector ready")
	}
	if err := machine.Transition(WorkerReady); err != nil {
		process.fail()
		return nil, err
	}
	_ = controlParent.SetDeadline(time.Time{})
	return process, nil
}

func (process *ConnectorProcess) Close(ctx context.Context) error {
	if process == nil || process.command == nil {
		return nil
	}
	registry, _ := appipc.NewProtocolRegistry()
	_ = process.machine.Transition(WorkerRetiring)
	shutdown, encodeErr := registry.Encode(appipc.Envelope{V: 1, Channel: appipc.ChannelCtrlC, Type: "shutdown", Seq: "2", Body: appipc.EmptyBody{}})
	if encodeErr == nil {
		_ = appipc.WriteFrame(process.control, shutdown)
	}
	wait := make(chan error, 1)
	go func() { wait <- process.command.Wait() }()
	select {
	case err := <-wait:
		_ = process.death.Close()
		_ = process.data.Close()
		_ = process.control.Close()
		_ = process.machine.Transition(WorkerTerminated)
		return err
	case <-ctx.Done():
		_ = process.death.Close()
		_ = process.data.Close()
		_ = process.control.Close()
		_ = process.command.Process.Kill()
		<-wait
		return ctx.Err()
	}
}

func (process *ConnectorProcess) fail() {
	_ = process.machine.Transition(WorkerFailed)
	_ = process.death.Close()
	_ = process.data.Close()
	_ = process.control.Close()
	_ = process.command.Process.Kill()
	_ = process.command.Wait()
}

func newInheritedSocket(name string) (net.Conn, *os.File, error) {
	fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}
	syscall.CloseOnExec(fds[0])
	syscall.CloseOnExec(fds[1])
	parentFile := os.NewFile(uintptr(fds[0]), name+"-parent")
	childFile := os.NewFile(uintptr(fds[1]), name+"-child")
	parent, err := net.FileConn(parentFile)
	_ = parentFile.Close()
	if err != nil {
		_ = childFile.Close()
		return nil, nil, err
	}
	return parent, childFile, nil
}
