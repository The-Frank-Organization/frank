//go:build linux

package broker

import (
	"errors"
	"net"
	"syscall"
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
	var credential *syscall.Ucred
	var socketErr error
	if err := raw.Control(func(fd uintptr) {
		credential, socketErr = syscall.GetsockoptUcred(int(fd), syscall.SOL_SOCKET, syscall.SO_PEERCRED)
	}); err != nil {
		return 0, err
	}
	if socketErr != nil || credential == nil {
		return 0, socketErr
	}
	return int(credential.Pid), nil
}
