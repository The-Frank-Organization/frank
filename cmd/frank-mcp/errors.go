package main

import (
	"errors"
	"io"
	"net"
	"os"
	"strings"
)

const (
	errConductorUnreachable = "shim:conductor-unreachable"
	errAuthFailed           = "shim:auth-failed"
	errConnectionLost       = "shim:connection-lost"
	errFrameTooLarge        = "shim:frame-too-large"
	errProtocol             = "shim:protocol-error"
	errSchemaInvalid        = "schema_invalid"
)

func scrubError(err error) string {
	if err == nil {
		return ""
	}
	msg := err.Error()
	switch {
	case strings.Contains(msg, "frame-too-large"):
		return errFrameTooLarge
	case strings.Contains(msg, "auth:"):
		return errAuthFailed
	case errors.Is(err, net.ErrClosed), errors.Is(err, io.EOF), errors.Is(err, io.ErrClosedPipe),
		strings.Contains(msg, "broken pipe"),
		strings.Contains(msg, "use of closed network connection"),
		strings.Contains(msg, "connection reset by peer"):
		return errConnectionLost
	case errors.Is(err, os.ErrNotExist),
		strings.Contains(msg, "dial unix"),
		strings.Contains(msg, "connect:"),
		strings.Contains(msg, "connect: no such file"),
		strings.Contains(msg, "connection refused"),
		strings.Contains(msg, "conductor"):
		return errConductorUnreachable
	default:
		return errProtocol
	}
}
