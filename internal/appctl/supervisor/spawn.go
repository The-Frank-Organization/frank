package supervisor

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"sort"
	"syscall"
)

func SanitizeEnv(source map[string]string, allow []string) []string {
	allowed := make(map[string]struct{}, len(allow))
	for _, name := range allow {
		allowed[name] = struct{}{}
	}
	result := make([]string, 0, len(allow))
	for name, value := range source {
		if _, ok := allowed[name]; ok {
			result = append(result, name+"="+value)
		}
	}
	sort.Strings(result)
	return result
}

type DeathPipe struct {
	Parent *os.File
	Child  *os.File
}

func NewDeathPipe() (DeathPipe, error) {
	readEnd, writeEnd, err := os.Pipe()
	if err != nil {
		return DeathPipe{}, err
	}
	return DeathPipe{Parent: writeEnd, Child: readEnd}, nil
}

type ProcessSpec struct {
	Path       string
	Args       []string
	Dir        string
	Env        []string
	ExtraFiles []*os.File
}

func BuildCommand(spec ProcessSpec) (*exec.Cmd, error) {
	if spec.Path == "" || spec.Dir == "" {
		return nil, fmt.Errorf("supervisor: process path and runtime directory are required")
	}
	command := exec.Command(spec.Path, spec.Args...)
	command.Dir = spec.Dir
	command.Env = append([]string(nil), spec.Env...)
	command.ExtraFiles = append([]*os.File(nil), spec.ExtraFiles...)
	return command, nil
}

func PrepareRuntimeDir(path string) error {
	if path == "" {
		return fmt.Errorf("supervisor: empty runtime directory")
	}
	if err := os.MkdirAll(path, 0o700); err != nil {
		return err
	}
	return os.Chmod(path, 0o700)
}

type SocketPair struct {
	Parent net.Conn
	Child  net.Conn
}

func NewSocketPair() (SocketPair, error) {
	fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		return SocketPair{}, err
	}
	syscall.CloseOnExec(fds[0])
	syscall.CloseOnExec(fds[1])
	parentFile := os.NewFile(uintptr(fds[0]), "app-parent")
	childFile := os.NewFile(uintptr(fds[1]), "child-endpoint")
	parent, err := net.FileConn(parentFile)
	_ = parentFile.Close()
	if err != nil {
		_ = childFile.Close()
		return SocketPair{}, err
	}
	child, err := net.FileConn(childFile)
	_ = childFile.Close()
	if err != nil {
		_ = parent.Close()
		return SocketPair{}, err
	}
	return SocketPair{Parent: parent, Child: child}, nil
}
