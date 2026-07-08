package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

type RootLock struct {
	path string
	file *os.File
	mu   sync.Mutex
	done bool
}

type ErrRootLocked struct {
	ErrorClass    string `json:"error_class"`
	HolderPID     int    `json:"holder_pid,omitempty"`
	HolderStarted string `json:"holder_started,omitempty"`
}

func (e ErrRootLocked) Error() string {
	return fmt.Sprintf("%s: holder_pid=%d holder_started=%s", e.ErrorClass, e.HolderPID, e.HolderStarted)
}

type rootLockInfo struct {
	PID      int               `json:"pid"`
	Started  string            `json:"started"`
	Takeover *rootLockTakeover `json:"takeover,omitempty"`
}

type rootLockTakeover struct {
	Event           string `json:"event"`
	PreviousPID     int    `json:"previous_pid,omitempty"`
	PreviousStarted string `json:"previous_started,omitempty"`
	TakenAt         string `json:"taken_at"`
}

func AcquireRoot(root string) (*RootLock, error) {
	if err := os.MkdirAll(root, 0o755); err != nil {
		return nil, err
	}
	path := filepath.Join(root, "conductor.lock")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		info := readRootLockInfo(path)
		_ = f.Close()
		return nil, ErrRootLocked{ErrorClass: "root-lock-held", HolderPID: info.PID, HolderStarted: info.Started}
	}
	previous := readRootLockInfo(path)
	now := time.Now().UTC().Format(time.RFC3339Nano)
	info := rootLockInfo{PID: os.Getpid(), Started: now}
	if previous.PID != 0 && previous.PID != os.Getpid() {
		info.Takeover = &rootLockTakeover{
			Event:           "TAKEOVER",
			PreviousPID:     previous.PID,
			PreviousStarted: previous.Started,
			TakenAt:         now,
		}
	}
	data, err := json.Marshal(info)
	if err != nil {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
		return nil, err
	}
	data = append(data, '\n')
	if err := f.Truncate(0); err != nil {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
		return nil, err
	}
	if _, err := f.Seek(0, 0); err != nil {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
		return nil, err
	}
	if _, err := f.Write(data); err != nil {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
		return nil, err
	}
	if err := f.Sync(); err != nil {
		_ = syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
		_ = f.Close()
		return nil, err
	}
	return &RootLock{path: path, file: f}, nil
}

func (l *RootLock) Release() error {
	if l == nil {
		return nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.done {
		return nil
	}
	l.done = true
	if err := l.file.Truncate(0); err != nil {
		_ = syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)
		_ = l.file.Close()
		return err
	}
	if _, err := l.file.Seek(0, 0); err != nil {
		_ = syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)
		_ = l.file.Close()
		return err
	}
	if err := l.file.Sync(); err != nil {
		_ = syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)
		_ = l.file.Close()
		return err
	}
	err := syscall.Flock(int(l.file.Fd()), syscall.LOCK_UN)
	if closeErr := l.file.Close(); err == nil {
		err = closeErr
	}
	return err
}

func readRootLockInfo(path string) rootLockInfo {
	data, err := os.ReadFile(path)
	if err != nil {
		return rootLockInfo{}
	}
	var info rootLockInfo
	_ = json.Unmarshal(data, &info)
	return info
}
