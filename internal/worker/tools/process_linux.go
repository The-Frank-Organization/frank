//go:build linux

package tools

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func configureProcess(command *exec.Cmd) {
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pdeathsig: syscall.SIGKILL}
}

func terminateProcessGroup(pid int, signal syscall.Signal) {
	_ = syscall.Kill(-pid, signal)
}

func closeOnExecDescriptors() error {
	directory, err := os.Open("/proc/self/fd")
	if err != nil {
		return fmt.Errorf("tools: enumerate descriptors: %w", err)
	}
	names, readErr := directory.Readdirnames(-1)
	closeErr := directory.Close()
	if readErr != nil {
		return fmt.Errorf("tools: enumerate descriptors: %w", readErr)
	}
	if closeErr != nil {
		return fmt.Errorf("tools: close descriptor directory: %w", closeErr)
	}
	for _, name := range names {
		descriptor, parseErr := strconv.Atoi(name)
		if parseErr == nil && descriptor >= 3 {
			syscall.CloseOnExec(descriptor)
		}
	}
	return nil
}
