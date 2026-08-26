//go:build darwin

package broker

import (
	"errors"
	"net"
	"syscall"
)

const (
	darwinSOLLocal     = 0
	darwinLocalPeerPID = 2
)

func peerPID(connection net.Conn) (int, error) {
	socket, ok := connection.(syscall.Conn)
	if !ok {
		return 0, errors.New("broker: control connection has no raw socket")
	}
	raw, err := socket.SyscallConn()
	if err != nil {
		return 0, err
	}
	var pid int
	var socketErr error
	if err := raw.Control(func(fd uintptr) {
		pid, socketErr = syscall.GetsockoptInt(int(fd), darwinSOLLocal, darwinLocalPeerPID)
	}); err != nil {
		return 0, err
	}
	if socketErr != nil {
		return 0, socketErr
	}
	return pid, nil
}
