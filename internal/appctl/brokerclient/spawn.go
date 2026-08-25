package brokerclient

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

func BrokerCommand(path string, args []string, dir string, tokenPipe *os.File) (*exec.Cmd, error) {
	if path == "" || dir == "" {
		return nil, errors.New("brokerclient: path and runtime directory are required")
	}
	command := exec.Command(path, args...)
	command.Dir = dir
	command.Env = []string{}
	if tokenPipe != nil {
		command.ExtraFiles = []*os.File{tokenPipe}
	}
	command.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	return command, nil
}
