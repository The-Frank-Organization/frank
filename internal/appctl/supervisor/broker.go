package supervisor

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"sync/atomic"
	"time"

	"github.com/jackli/frank/internal/appctl/brokerclient"
)

const BrokerReadyDeadline = 10 * time.Second

var brokerReadyPattern = regexp.MustCompile(`^BROKER_READY nonce=([0-9a-f]{64})\n$`)

type BrokerLaunch struct {
	BinaryPath, RuntimeDir, ConfigHome string
	RunID, ControlToken                string
	At                                 int64
	Client                             *brokerclient.Client
	Controller                         *Controller
	InstanceID                         string
	ReadyDeadline                      time.Duration
}

type BrokerProcess struct {
	command           *exec.Cmd
	stdout            io.ReadCloser
	machine           *Machine
	nonce             string
	controller        *Controller
	runID, instanceID string
	wait              chan error
	closing           atomic.Bool
}

func (process *BrokerProcess) State() WorkerState {
	if process == nil || process.machine == nil {
		return WorkerFailed
	}
	return process.machine.State()
}

func (process *BrokerProcess) Nonce() string {
	if process == nil {
		return ""
	}
	return process.nonce
}

func LaunchBroker(ctx context.Context, launch BrokerLaunch) (*BrokerProcess, error) {
	if launch.BinaryPath == "" || launch.RuntimeDir == "" || launch.ConfigHome == "" || launch.RunID == "" || launch.ControlToken == "" || launch.Client == nil || launch.Controller == nil || launch.InstanceID == "" {
		return nil, errors.New("supervisor: incomplete broker launch")
	}
	if err := PrepareRuntimeDir(launch.RuntimeDir); err != nil {
		return nil, err
	}
	if _, err := launch.Client.AdvanceControl(ctx, launch.RunID, launch.ControlToken, launch.At); err != nil {
		return nil, fmt.Errorf("supervisor: pre-spawn broker control: %w", err)
	}
	tokenRead, tokenWrite, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	defer tokenWrite.Close()
	machine := NewMachine(WorkerAllocated)
	if err := machine.Transition(WorkerSpawning); err != nil {
		_ = tokenRead.Close()
		return nil, err
	}
	command, err := brokerclient.BrokerCommand(launch.BinaryPath, []string{"--config-home", launch.ConfigHome}, launch.RuntimeDir, tokenRead)
	if err != nil {
		_ = tokenRead.Close()
		return nil, err
	}
	stdout, err := command.StdoutPipe()
	if err != nil {
		_ = tokenRead.Close()
		return nil, err
	}
	process := &BrokerProcess{
		command: command, stdout: stdout, machine: machine, controller: launch.Controller,
		runID: launch.RunID, instanceID: launch.InstanceID,
	}
	if err := command.Start(); err != nil {
		_ = tokenRead.Close()
		_ = machine.Transition(WorkerFailed)
		_, recordErr := launch.Controller.RecordBrokerFailure(ctx, BrokerFailureRequest{RunID: launch.RunID, InstanceID: launch.InstanceID, Class: BrokerSpawnFail, At: launch.At})
		return nil, errors.Join(fmt.Errorf("supervisor: start broker: %w", err), recordErr)
	}
	_ = tokenRead.Close()
	if _, err := io.WriteString(tokenWrite, launch.ControlToken+"\n"); err != nil {
		process.fail()
		_, recordErr := launch.Controller.RecordBrokerFailure(ctx, BrokerFailureRequest{RunID: launch.RunID, InstanceID: launch.InstanceID, Class: BrokerSpawnFail, At: launch.At})
		return nil, errors.Join(fmt.Errorf("supervisor: broker token pipe: %w", err), recordErr)
	}
	if err := tokenWrite.Close(); err != nil {
		process.fail()
		return nil, fmt.Errorf("supervisor: close broker token pipe: %w", err)
	}
	deadline := launch.ReadyDeadline
	if deadline <= 0 {
		deadline = BrokerReadyDeadline
	}
	ready := make(chan struct {
		line string
		err  error
	}, 1)
	go func() {
		line, readErr := bufio.NewReader(stdout).ReadString('\n')
		ready <- struct {
			line string
			err  error
		}{line: line, err: readErr}
	}()
	timer := time.NewTimer(deadline)
	defer timer.Stop()
	select {
	case observed := <-ready:
		nonce, parseErr := parseBrokerReady(observed.line)
		if observed.err != nil || parseErr != nil {
			process.fail()
			_, recordErr := launch.Controller.RecordBrokerFailure(ctx, BrokerFailureRequest{RunID: launch.RunID, InstanceID: launch.InstanceID, Class: BrokerMalformedReady, At: launch.At})
			return nil, errors.Join(errors.New("supervisor: invalid broker READY attestation"), recordErr)
		}
		process.nonce = nonce
	case <-ctx.Done():
		process.fail()
		return nil, ctx.Err()
	case <-timer.C:
		process.fail()
		_, recordErr := launch.Controller.RecordBrokerFailure(ctx, BrokerFailureRequest{RunID: launch.RunID, InstanceID: launch.InstanceID, Class: BrokerNoReady, At: launch.At})
		return nil, errors.Join(errors.New("supervisor: broker READY deadline exceeded"), recordErr)
	}
	if err := machine.Transition(WorkerReady); err != nil {
		process.fail()
		return nil, err
	}
	process.wait = make(chan error, 1)
	go func() {
		err := process.command.Wait()
		process.wait <- err
		if !process.closing.Load() {
			_, _ = process.controller.RecordBrokerFailure(context.Background(), BrokerFailureRequest{
				RunID: process.runID, InstanceID: process.instanceID, Class: BrokerReadyCrash, At: time.Now().UnixNano(),
			})
		}
	}()
	return process, nil
}

func parseBrokerReady(line string) (string, error) {
	match := brokerReadyPattern.FindStringSubmatch(line)
	if len(match) != 2 || len(line) != 84 {
		return "", errors.New("supervisor: invalid broker READY attestation")
	}
	return match[1], nil
}

func (process *BrokerProcess) Close(ctx context.Context) error {
	if process == nil || process.command == nil || process.command.Process == nil {
		return nil
	}
	_ = process.machine.Transition(WorkerRetiring)
	process.closing.Store(true)
	_ = process.command.Process.Signal(os.Interrupt)
	select {
	case err := <-process.wait:
		_ = process.stdout.Close()
		_ = process.machine.Transition(WorkerTerminated)
		return err
	case <-ctx.Done():
		_ = process.command.Process.Kill()
		<-process.wait
		_ = process.stdout.Close()
		_ = process.machine.Transition(WorkerTerminated)
		return ctx.Err()
	}
}

func (process *BrokerProcess) fail() {
	_ = process.machine.Transition(WorkerFailed)
	if process.stdout != nil {
		_ = process.stdout.Close()
	}
	if process.command != nil && process.command.Process != nil {
		_ = process.command.Process.Kill()
		_ = process.command.Wait()
	}
}
