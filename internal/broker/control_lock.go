package broker

import (
	"fmt"
	"os"
	"syscall"
)

func verifyControlLock(path string, peer int) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("broker: open control lock: %w", err)
	}
	defer file.Close()
	record := syscall.Flock_t{Type: syscall.F_WRLCK, Whence: 0, Start: 0, Len: 0}
	if err := syscall.FcntlFlock(file.Fd(), syscall.F_GETLK, &record); err != nil {
		return fmt.Errorf("broker: probe control lock: %w", err)
	}
	if record.Type != syscall.F_WRLCK || int(record.Pid) != peer {
		return fmt.Errorf("broker: control lock holder %d does not match peer %d", record.Pid, peer)
	}
	return nil
}
